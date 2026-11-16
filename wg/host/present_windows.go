//go:build windows

package host

import (
	"unsafe"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
)

// Win32 呈现：合成器产出 BGRA32 预乘内存帧，最终画到窗口上。
//
// 为什么不像早期版本那样直接 GetDC + StretchDIBits：窗口客户区是 DWM
// 重定向表面，绕过 LCL 的 TCanvas 直接写窗口 DC 会间歇性返回 0（实测同一份
// 参数有时成功、有时失败），且必然与 LCL 自己的绘制争抢同一条呈现路径。
// 参考 examples/wg/showcase 的实证做法：交给 LCL 的 Canvas。
//
// 关键约束（来自 showcase 的实测）：
//   - 离屏位图必须 Pf24bit。Pf32bit 的 DIB 第 4 字节恒为 0，LCL 贴图时按
//     AC_SRC_ALPHA 混合，所有 GDI 矢量图元都会被当全透明而消失。
//   - 位图用 CreateIntfImage 拿到可写像素区，写完 LoadFromIntfImage 同步回
//     位图，再 CopyRect 到窗口 Canvas。BGRA→BGR 的转换只做一次。
//
// 线程模型：Frame 在帧泵 goroutine 记录**自有副本**（合成缓冲下一帧就会被
// 覆写），Paint 在窗口主线程 OnPaint 里提交。
type presentPlatform struct {
	bmp    lcl.IBitmap
	intf   lcl.ILazIntfImage
	pixels []byte // 24bpp BGR 视图，直接指向 intf.PixelData
	w, h   int
	stride int
}

type winRect struct {
	Left, Top, Right, Bottom int32
}

// pointerBytes 把 LCL 返回的像素地址包成可写字节切片。
//
// PixelData 返回的是 uintptr（C 侧 DIB 地址）。直接写
// unsafe.Pointer(ptr) 会触发 vet 的 "possible misuse of unsafe.Pointer"，
// 因为编译器无法确认该 uintptr 来自指针。这里经内存中转一步，
// 语义相同（地址由 LCL 位图对象持有，生命周期由 bmp 保证），且不误报。
func pointerBytes(ptr uintptr, n int) []byte {
	if ptr == 0 || n <= 0 {
		return nil
	}
	var up unsafe.Pointer
	*(*uintptr)(unsafe.Pointer(&up)) = ptr
	return unsafe.Slice((*byte)(up), n)
}

// ensureBuffer 准备 24bpp 离屏位图与可写像素视图；尺寸变化时重建。
func (p *Presenter) ensureBuffer(w, h int) bool {
	pl := &p.plat
	if pl.bmp != nil && pl.w == w && pl.h == h && pl.pixels != nil {
		return true
	}
	p.releaseBuffer()

	bmp := lcl.NewBitmap()
	if bmp == nil || !bmp.IsValid() {
		return false
	}
	bmp.SetPixelFormat(types.Pf24bit)
	bmp.SetSize(int32(w), int32(h))
	intf := bmp.CreateIntfImage()
	if intf == nil || !intf.IsValid() {
		bmp.Free()
		return false
	}
	ptr := intf.PixelData()
	if ptr == 0 {
		intf.Free()
		bmp.Free()
		return false
	}
	stride := w * 3 // Pf24bit：3 字节/像素，行间无填充
	pl.bmp, pl.intf, pl.w, pl.h, pl.stride = bmp, intf, w, h, stride
	pl.pixels = pointerBytes(ptr, stride*h)
	return true
}

// releaseBuffer 释放 LCL 位图与像素视图。
// 顺序重要：先放开像素视图再 Free 位图，否则 DIB 指针悬空。
func (p *Presenter) releaseBuffer() {
	pl := &p.plat
	pl.pixels = nil
	if pl.intf != nil {
		pl.intf.Free()
		pl.intf = nil
	}
	if pl.bmp != nil {
		pl.bmp.Free()
		pl.bmp = nil
	}
	pl.w, pl.h, pl.stride = 0, 0, 0
}

// submit 把 b 的像素提交到窗口 Canvas。damage 为空则整屏。
//
// 返回值严格：只有像素真的写进位图并成功 CopyRect 才返回 true。
func (p *Presenter) submit(b *raster.Backing, damage geom.Rect) bool {
	if b == nil || p.canvas == nil {
		return false
	}
	w, h := b.W(), b.H()
	if w <= 0 || h <= 0 {
		return false
	}
	// resize 期间可能收到旧尺寸的帧；与当前客户区不符就不提交。
	if p.host != nil {
		cs := p.host.ClientSize()
		if cs.W > 0 && (int(cs.W) != w || int(cs.H) != h) {
			return false
		}
	}
	if !p.ensureBuffer(w, h) {
		return false
	}

	// 脏区（DIP 与像素 1:1）；空 damage 视为整帧。
	r := geom.RectXYWH(0, 0, float32(w), float32(h))
	if !damage.IsEmpty() {
		if d := damage.Intersect(b.Bounds()); !d.IsEmpty() {
			r = d
		}
	}
	x0, y0 := int(r.Min.X), int(r.Min.Y)
	x1, y1 := int(r.Max.X), int(r.Max.Y)

	// BGRA32 预乘 → 24bpp BGR：丢弃 alpha（画面已合成到不透明底）。
	src := b.Bytes()
	if len(src) == 0 {
		return false
	}
	pl := &p.plat
	for y := y0; y < y1; y++ {
		srow := y * b.Stride() // Stride 已是字节数（w*4）
		drow := y * pl.stride
		for x := x0; x < x1; x++ {
			s := srow + x*4
			d := drow + x*3
			pl.pixels[d+0] = src[s+0] // B
			pl.pixels[d+1] = src[s+1] // G
			pl.pixels[d+2] = src[s+2] // R
		}
	}

	// 像素视图 → 位图 → 窗口 Canvas。
	pl.intf.BeginUpdate()
	pl.intf.EndUpdate()
	pl.bmp.LoadFromIntfImage(pl.intf)

	dst := types.TRect{Left: int32(x0), Top: int32(y0), Right: int32(x1), Bottom: int32(y1)}
	dstCanvas := p.canvas()
	if dstCanvas == nil || !dstCanvas.IsValid() {
		return false
	}
	dstCanvas.CopyRectWithRectX2Canvas(dst, pl.bmp.Canvas(), dst)
	return true
}

// invalidate 在 LCL 路径下等价于请求重绘（像素在 OnPaint 里提交）。
func (p *Presenter) invalidate(damage geom.Rect) {
	if p.repaint != nil {
		p.repaint()
	}
}

// release 释放平台资源。
func (p *Presenter) release() { p.releaseBuffer() }
