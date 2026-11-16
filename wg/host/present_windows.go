//go:build windows

package host

import (
	"syscall"
	"unsafe"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
)

// Win32 呈现：把合成器的 BGRA32 预乘内存一次性提交到窗口 DC。
//
// 为什么是 StretchDIBits 而不是 LCL Canvas.CopyRect：
//   - 内存布局天然匹配：Backing 是行优先 BGRA32（小端 B,G,R,A），
//     32bpp top-down DIB（biHeight 为负）的字节序完全一致，零转换；
//   - 单次 SRCCOPY 不透明拷贝，不做 alpha 混合（契约 C2 的最终呈现格式）；
//   - 不依赖 LCL Canvas 对象（Binder 已完全接管 WndProc，LCL 绘制被屏蔽）。
var (
	presentGDI32      = syscall.NewLazyDLL("gdi32.dll")
	procStretchDIBits = presentGDI32.NewProc("StretchDIBits")
)

const (
	dibRGBColors = 0
	srcCopy      = 0x00CC0020
)

// presentBitmapInfo 对应 BITMAPINFO（仅 BITMAPINFOHEADER，32bpp 无颜色表）。
type presentBitmapInfo struct {
	Size          uint32
	Width         int32
	Height        int32 // 负数 = top-down，与 Go 行序一致
	Planes        uint16
	BitCount      uint16
	Compression   uint32
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
}

type winRect struct {
	Left, Top, Right, Bottom int32
}

// submit 把 b 的像素提交到窗口 DC。damage 为空则整屏。
func (p *Presenter) submit(b *raster.Backing, damage geom.Rect) bool {
	if b == nil || p.hwnd == 0 {
		return false
	}
	w, h := b.W(), b.H()
	if w <= 0 || h <= 0 {
		return false
	}
	// resize 期间可能收到旧尺寸的帧；与当前客户区不符就不提交，
	// 下一帧会用新尺寸的缓冲重画（否则 StretchDIBits 会拉伸出花屏）。
	if p.host != nil {
		cs := p.host.ClientSize()
		if cs.W > 0 && (int(cs.W) != w || int(cs.H) != h) {
			return false
		}
	}

	// 脏区 → 源/目标矩形（backing 与客户区 1:1 像素）。
	sx, sy, sw, sh := 0, 0, w, h
	if !damage.IsEmpty() {
		r := damage.Intersect(b.Bounds())
		if r.IsEmpty() {
			return true
		}
		sx, sy = int(r.Min.X), int(r.Min.Y)
		sw, sh = int(r.W()), int(r.H())
		if sw <= 0 || sh <= 0 {
			return true
		}
	}

	hdc, _, _ := procGetDC.Call(p.hwnd)
	if hdc == 0 {
		return false
	}
	defer procReleaseDC.Call(p.hwnd, hdc)

	bi := presentBitmapInfo{
		Size:        40,
		Width:       int32(w),
		Height:      int32(-h), // top-down
		Planes:      1,
		BitCount:    32,
		Compression: 0, // BI_RGB
	}
	bits := b.Bytes()
	if len(bits) == 0 {
		return false
	}
	_, _, err := procStretchDIBits.Call(
		hdc,
		uintptr(int32(sx)), uintptr(int32(sy)), uintptr(int32(sw)), uintptr(int32(sh)),
		uintptr(int32(sx)), uintptr(int32(sy)), uintptr(int32(sw)), uintptr(int32(sh)),
		uintptr(unsafe.Pointer(&bits[0])),
		uintptr(unsafe.Pointer(&bi)),
		uintptr(dibRGBColors),
		uintptr(srcCopy),
	)
	return err == nil || err == syscall.Errno(0)
}

// invalidate 把脏区标为无效，驱动系统下一轮发 WM_PAINT。
// bErase=false：整帧提交会覆盖脏区，擦背景只会多一次全窗写（闪烁根因）。
func (p *Presenter) invalidate(damage geom.Rect) {
	if p.hwnd == 0 {
		return
	}
	var ptr uintptr
	if !damage.IsEmpty() {
		r := winRect{
			Left:   int32(damage.Min.X),
			Top:    int32(damage.Min.Y),
			Right:  int32(damage.Max.X),
			Bottom: int32(damage.Max.Y),
		}
		ptr = uintptr(unsafe.Pointer(&r))
	}
	procInvalidateRect.Call(p.hwnd, ptr, 0)
}

func (p *Presenter) release() { /* StretchDIBits 无持久 GDI 对象 */ }
