//go:build windows

package gdi

import (
	"syscall"
	"unsafe"
)

// 直接调 gdi32/user32，不经过 cgo 也不经过 LCL。
//
// 为什么不用 lcl.ICanvas：
//   - gdi 后端必须在无窗口环境（CI、golden、--render-matrix）工作，
//     而 LCL 的 Canvas 依赖其运行时与窗口生命周期；
//   - 文字回放需要精确控制 DIB 格式与裁剪（C24），直接调 API 更可控。
//
// 所有句柄都由本文件创建与释放，遵循"谁创建谁释放"（C21 的视图→位图→池
// 顺序在 composite 层体现，这里只管位图与 DC）。

var (
	gdi32 = syscall.NewLazyDLL("gdi32.dll")

	procCreateCompatibleDC = gdi32.NewProc("CreateCompatibleDC")
	procCreateDIBSection   = gdi32.NewProc("CreateDIBSection")
	procSelectObject       = gdi32.NewProc("SelectObject")
	procDeleteObject       = gdi32.NewProc("DeleteObject")
	procDeleteDC           = gdi32.NewProc("DeleteDC")
	procGdiFlush           = gdi32.NewProc("GdiFlush")

	procSetPixel   = gdi32.NewProc("SetPixel")
	procPatBlt     = gdi32.NewProc("PatBlt")
	procPolygon    = gdi32.NewProc("Polygon")
	procPolyline   = gdi32.NewProc("Polyline")
	procCreatePen  = gdi32.NewProc("CreatePen")
	procCreateBrush = gdi32.NewProc("CreateSolidBrush")

	procSetBkMode       = gdi32.NewProc("SetBkMode")
	procSetTextColor    = gdi32.NewProc("SetTextColor")
	procSetTextAlign    = gdi32.NewProc("SetTextAlign")
	procExtTextOutW     = gdi32.NewProc("ExtTextOutW")
	procGetTextExtentPt = gdi32.NewProc("GetTextExtentPoint32W")
	procIntersectClip   = gdi32.NewProc("IntersectClipRect")
	procSaveDC          = gdi32.NewProc("SaveDC")
	procRestoreDC       = gdi32.NewProc("RestoreDC")
)

const (
	dibRGBColors = 0
	psSolid      = 0
	patCopy      = 0x00F00021
	transparent  = 1
	taLeft       = 0
	taBaseline   = 24
	// ETO_CLIPPED = 0x0004：按 ExtTextOutW 的 lpRect 裁剪输出。
	// 注意不是 0x0100——这里曾写错，靠外层的 SaveDC+IntersectClipRect 兜住了裁剪，
	// 于是每段文字多付 3 次 GDI 调用；改成真正的 ETO_CLIPPED 后可以省掉那 3 次。
	etoClipped   = 0x0004
	etoOpaque    = 0x0002
)

// bitmapInfoHeader 对应 BITMAPINFOHEADER（40 字节）。
type bitmapInfoHeader struct {
	Size          uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      uint16
	Compression   uint32
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
}

// point 对应 POINT。
type point struct {
	X, Y int32
}

// rect 对应 RECT。
type rect struct {
	Left, Top, Right, Bottom int32
}

// DIB 是一块 GDI 位图及其内存 DC（top-down，行序与 Go 切片一致）。
type DIB struct {
	// HDC 是内存 DC。
	HDC uintptr
	// HBmp 是 DIB section 位图句柄。
	HBmp uintptr
	// W/H 是像素尺寸。
	W, H int
	// Stride 是行字节数（4 字节对齐）。
	Stride int
	// Bpp 是位深（24 或 32）。
	Bpp int

	old    uintptr // SelectObject 返回的旧位图
	bits   []byte
	closed bool
}

// NewDIB 创建 w×h、bpp 位（24 或 32）的 top-down DIB。
//
// 24bpp 用于**呈现目标**（Pf24bit，契约 C2）；32bpp 用于 GDI 对照后端
// 与临时合成。返回 nil 表示创建失败（无 GDI 会话来宾会话等）。
func NewDIB(w, h, bpp int) *DIB {
	if w <= 0 || h <= 0 || (bpp != 24 && bpp != 32) {
		return nil
	}
	hdc, _, _ := procCreateCompatibleDC.Call(0)
	if hdc == 0 {
		return nil
	}
	bi := bitmapInfoHeader{
		Size:        40,
		Width:       int32(w),
		Height:      int32(-h), // 负数 = top-down
		Planes:      1,
		BitCount:    uint16(bpp),
		Compression: 0,
	}
	var bits unsafe.Pointer
	hbm, _, _ := procCreateDIBSection.Call(hdc, uintptr(unsafe.Pointer(&bi)), dibRGBColors, uintptr(unsafe.Pointer(&bits)), 0, 0)
	if hbm == 0 || bits == nil {
		procDeleteDC.Call(hdc)
		return nil
	}
	old, _, _ := procSelectObject.Call(hdc, hbm)
	stride := ((w*bpp + 31) / 32) * 4
	n := stride * h
	d := &DIB{
		HDC: hdc, HBmp: hbm, W: w, H: h, Stride: stride, Bpp: bpp,
		old: old,
	}
	// 用 unsafe.Slice 得到可读写的字节视图（与 C21 的"视图"概念对应）
	d.bits = unsafe.Slice((*byte)(bits), n)
	return d
}

// Bits 返回像素字节视图（长度 Stride*H）。
func (d *DIB) Bits() []byte {
	if d == nil || d.closed {
		return nil
	}
	return d.bits
}

// Row 返回第 y 行字节。
func (d *DIB) Row(y int) []byte {
	if d == nil || d.closed || y < 0 || y >= d.H {
		return nil
	}
	return d.bits[y*d.Stride : (y+1)*d.Stride]
}

// Flush 确保 GDI 已把像素写入 DIB 内存（GdiFlush 是同步点）。
func (d *DIB) Flush() {
	if d != nil && !d.closed {
		procGdiFlush.Call()
	}
}

// Clear 用指定 COLORREF（0x00BBGGRR）填满。
func (d *DIB) Clear(color uint32) {
	if d == nil || d.closed {
		return
	}
	br, _, _ := procCreateBrush.Call(uintptr(color))
	if br == 0 {
		return
	}
	prev, _, _ := procSelectObject.Call(d.HDC, br)
	procPatBlt.Call(d.HDC, 0, 0, uintptr(d.W), uintptr(d.H), patCopy)
	if prev != 0 {
		procSelectObject.Call(d.HDC, prev)
	}
	procDeleteObject.Call(br)
}

// Delete 释放 DIB（先恢复旧位图，再删位图，最后删 DC——顺序不能反）。
func (d *DIB) Delete() {
	if d == nil || d.closed {
		return
	}
	d.closed = true
	if d.old != 0 {
		procSelectObject.Call(d.HDC, d.old)
		d.old = 0
	}
	if d.HBmp != 0 {
		procDeleteObject.Call(d.HBmp)
		d.HBmp = 0
	}
	if d.HDC != 0 {
		procDeleteDC.Call(d.HDC)
		d.HDC = 0
	}
	d.bits = nil
}

// clipRect 在 DC 上设置裁剪（与当前裁剪求交），返回 SaveDC 的标记用于恢复。
func (d *DIB) clipRect(r [4]float32) uintptr {
	saved, _, _ := procSaveDC.Call(d.HDC)
	procIntersectClip.Call(d.HDC, uintptr(int32(r[0])), uintptr(int32(r[1])), uintptr(int32(r[2])), uintptr(int32(r[3])))
	return saved
}

func (d *DIB) restore(saved uintptr) {
	if saved != 0 {
		procRestoreDC.Call(d.HDC, saved)
	}
}

// fillPolygon 用实心色填充多边形（GDI 使用奇偶规则）。
func (d *DIB) fillPolygon(pts []point, color uint32) bool {
	if len(pts) < 3 {
		return false
	}
	br, _, _ := procCreateBrush.Call(uintptr(color))
	pen, _, _ := procCreatePen.Call(psSolid, 0, uintptr(color))
	if br == 0 || pen == 0 {
		return false
	}
	prevBr, _, _ := procSelectObject.Call(d.HDC, br)
	prevPen, _, _ := procSelectObject.Call(d.HDC, pen)
	ok, _, _ := procPolygon.Call(d.HDC, uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if prevBr != 0 {
		procSelectObject.Call(d.HDC, prevBr)
	}
	if prevPen != 0 {
		procSelectObject.Call(d.HDC, prevPen)
	}
	procDeleteObject.Call(br)
	procDeleteObject.Call(pen)
	return ok != 0
}

// strokePolyline 用指定宽度与颜色画折线。
func (d *DIB) strokePolyline(pts []point, width int, color uint32) bool {
	if len(pts) < 2 {
		return false
	}
	if width < 1 {
		width = 1
	}
	pen, _, _ := procCreatePen.Call(psSolid, uintptr(width), uintptr(color))
	if pen == 0 {
		return false
	}
	prev, _, _ := procSelectObject.Call(d.HDC, pen)
	ok, _, _ := procPolyline.Call(d.HDC, uintptr(unsafe.Pointer(&pts[0])), uintptr(len(pts)))
	if prev != 0 {
		procSelectObject.Call(d.HDC, prev)
	}
	procDeleteObject.Call(pen)
	return ok != 0
}

// fillRect 快速填充矩形（整数像素）。
func (d *DIB) fillRect(x0, y0, x1, y1 int, color uint32) bool {
	if x1 <= x0 || y1 <= y0 {
		return false
	}
	br, _, _ := procCreateBrush.Call(uintptr(color))
	if br == 0 {
		return false
	}
	prev, _, _ := procSelectObject.Call(d.HDC, br)
	ok, _, _ := procPatBlt.Call(d.HDC, uintptr(x0), uintptr(y0), uintptr(x1-x0), uintptr(y1-y0), patCopy)
	if prev != 0 {
		procSelectObject.Call(d.HDC, prev)
	}
	procDeleteObject.Call(br)
	return ok != 0
}
