//go:build windows

package host

import (
	"unsafe"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/colors"
	"github.com/energye/lcl/wg/geom"
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
	// 行跨距必须实测：24bpp 的 TRawImage 行按 4 字节对齐（实测 W=101 →
	// stride=304 而非 303）。按 w*3 写非 4 对齐宽度会让每行错位 1–3 字节，
	// 画面出现逐行偏移（灰/彩色斜纹）——只有 W%4==0 的窗口才恰好正常。
	stride := w * 3
	if p0, p1 := intf.GetDataLineStart(0), intf.GetDataLineStart(1); p1 > p0 {
		stride = int(p1 - p0)
	}
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

// submit 把 staged（24bpp BGR）提交到窗口 Canvas。
//
// 遵循成熟框架的 resize 语义：
//   - 尺寸吻合：有新帧只更新脏区；没有新帧（被遮挡后重露）整帧重放。
//   - 尺寸失配（过渡帧）：默认精确模式——内容 1:1 **绝不缩放**，只画
//     帧/客户区交集，多出的条带填背景色；只有显式 ResizePresentStretch
//     才整窗缩放。
//
// 主线程只做：行 memcpy（紧凑 stride → 位图对齐 stride）、
// LoadFromIntfImage、CopyRect/FillRect；逐像素转换已在帧线程完成。
func (p *Presenter) submit() bool {
	p.mu.RLock()
	st := p.staged
	w, h := p.sw, p.sh
	damage := p.damage
	dirty := p.dirty
	gen := p.frames
	stretch := p.stretchResize
	bg := p.background
	p.mu.RUnlock()
	if st == nil || p.canvas == nil || w <= 0 || h <= 0 || len(st) < w*3*h {
		return false
	}
	cw, ch := w, h
	scale := false
	if p.host != nil {
		cs := p.host.ClientSize()
		if cs.W > 0 && cs.H > 0 {
			cw, ch = int(cs.W), int(cs.H)
			scale = cw != w || ch != h
		}
	}
	if !p.ensureBuffer(w, h) {
		return false
	}

	// 1:1 更新区：有新帧只搬脏区；没有新帧（重露/系统重绘）整帧。
	r := geom.RectXYWH(0, 0, float32(w), float32(h))
	if dirty && !damage.IsEmpty() {
		if d := damage.Intersect(r); !d.IsEmpty() {
			r = d
		}
	}
	x0, y0 := int(r.Min.X), int(r.Min.Y)
	x1, y1 := int(r.Max.X), int(r.Max.Y)

	// 行拷贝：staged 紧凑 BGR → 位图行（真实 stride，4 字节对齐）。
	pl := &p.plat
	p.mu.RLock()
	if pl.stride > 0 && len(p.staged) >= w*3*h {
		for y := y0; y < y1; y++ {
			srow := y * w * 3
			drow := y * pl.stride
			copy(pl.pixels[drow+x0*3:drow+x1*3], p.staged[srow+x0*3:srow+x1*3])
		}
	}
	p.mu.RUnlock()

	// 像素视图 → 位图 → 窗口 Canvas。
	pl.intf.BeginUpdate()
	pl.intf.EndUpdate()
	pl.bmp.LoadFromIntfImage(pl.intf)

	dstCanvas := p.canvas()
	if dstCanvas == nil || !dstCanvas.IsValid() {
		return false
	}
	if scale {
		if stretch {
			// 显式整窗缩放（ResizePresentStretch）。
			dst := types.TRect{Left: 0, Top: 0, Right: int32(cw), Bottom: int32(ch)}
			src := types.TRect{Left: 0, Top: 0, Right: int32(w), Bottom: int32(h)}
			dstCanvas.CopyRectWithRectX2Canvas(dst, pl.bmp.Canvas(), src)
			p.markPresented(gen)
			return true
		}
		// 精确模式：1:1 交集 + 背景填边（几何见 exactPresentPlan）。
		plan := exactPresentPlan(w, h, cw, ch)
		cp := plan.Copy
		dst := types.TRect{Left: 0, Top: 0, Right: int32(cp.Max.X), Bottom: int32(cp.Max.Y)}
		dstCanvas.CopyRectWithRectX2Canvas(dst, pl.bmp.Canvas(), dst)
		if len(plan.Fill) > 0 {
			if bg.A == 0 && bg == (geom.Color{}) {
				bg = geom.RGBA(255, 255, 255)
			}
			br := dstCanvas.BrushToBrush()
			br.SetColor(colors.RGBToColor(bg.R, bg.G, bg.B))
			br.SetStyle(types.BsSolid)
			for _, f := range plan.Fill {
				dstCanvas.FillRectWithIntX4(
					int32(f.Min.X), int32(f.Min.Y), int32(f.Max.X), int32(f.Max.Y))
			}
		}
		p.markPresented(gen)
		return true
	}
	dst := types.TRect{Left: int32(x0), Top: int32(y0), Right: int32(x1), Bottom: int32(y1)}
	dstCanvas.CopyRectWithRectX2Canvas(dst, pl.bmp.Canvas(), dst)
	p.markPresented(gen)
	return true
}

// markPresented 清掉"有新像素"标记；提交期间若又来了新帧（frames 变化）
// 则保留标记，下一次提交仍按新帧处理。
func (p *Presenter) markPresented(gen int) {
	p.mu.Lock()
	if p.frames == gen {
		p.dirty = false
	}
	p.mu.Unlock()
}

// invalidate 请求把新帧提交到窗口。
//
// 优先走 posted 消息（高优先级：GetMessage 顺序 sent > posted > 输入 >
// WM_PAINT，拖动/连续输入下 WM_PAINT 会被饿死）；没有 post 通道时退回
// LCL Repaint。
func (p *Presenter) invalidate(damage geom.Rect) {
	p.mu.RLock()
	post := p.post
	repaint := p.repaint
	p.mu.RUnlock()
	if post != nil {
		// 闭环：上一帧还在队列里或正在画 → 只记 missed（合成缓冲区里已是
		// 最新像素，下一次投递自然取到），不往队列里再压一条。
		if !p.presenting.CompareAndSwap(false, true) {
			p.missed.Store(true)
			return
		}
		// 慢帧后先空出队列让输入先走，再投下一帧（showcase §9）。
		p.breathIfSlow()
		post()
		return
	}
	if repaint != nil {
		repaint()
	}
}

// release 释放平台资源。
func (p *Presenter) release() { p.releaseBuffer() }
