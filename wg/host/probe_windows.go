//go:build windows

package host

import (
	"syscall"
	"unsafe"
)

// ── 窗口级像素探针 ──────────────────────────────────────────────────
//
// 为什么需要它：Container.Backing() 是**内存**合成缓冲，它里面有内容
// 不代表屏幕上真的有像素。M6 的真窗空白 bug 就是因为只看内存而误判通过。
// 这里直接 GetDC(hwnd) + GetPixel 采样真实窗口客户区，给出屏幕级证据。

var (
	probeUser32 = syscall.NewLazyDLL("user32.dll")
	probeGDI32  = syscall.NewLazyDLL("gdi32.dll")

	procProbeGetDC     = probeUser32.NewProc("GetDC")
	procProbeReleaseDC = probeUser32.NewProc("ReleaseDC")
	procProbeGetClient = probeUser32.NewProc("GetClientRect")
	procProbeGetPixel  = probeGDI32.NewProc("GetPixel")
)

type probeRect struct {
	Left, Top, Right, Bottom int32
}

// WindowPixelStats 采样窗口客户区的真实像素。
//
// 返回非白采样点数与总采样点数；客户区为空或 GetDC 失败时返回 ok=false。
// step 是采样步长（像素），建议 8：兼顾覆盖率与 GetPixel 的调用开销。
func WindowPixelStats(hwnd uintptr, step int) (nonWhite, sampled, cw, ch int, ok bool) {
	if hwnd == 0 {
		return 0, 0, 0, 0, false
	}
	if step < 1 {
		step = 8
	}
	var rc probeRect
	procProbeGetClient.Call(hwnd, uintptr(unsafe.Pointer(&rc)))
	cw, ch = int(rc.Right-rc.Left), int(rc.Bottom-rc.Top)
	if cw <= 0 || ch <= 0 {
		return 0, 0, cw, ch, false
	}
	hdc, _, _ := procProbeGetDC.Call(hwnd)
	if hdc == 0 {
		return 0, 0, cw, ch, false
	}
	defer procProbeReleaseDC.Call(hwnd, hdc)

	clrInvalid := uintptr(0xFFFFFFFF)
	for y := 0; y < ch; y += step {
		for x := 0; x < cw; x += step {
			c, _, _ := procProbeGetPixel.Call(hdc, uintptr(int32(x)), uintptr(int32(y)))
			if c == clrInvalid {
				continue
			}
			sampled++
			// COLORREF 是 0x00BBGGRR。
			r := byte(c & 0xFF)
			g := byte((c >> 8) & 0xFF)
			b := byte((c >> 16) & 0xFF)
			if r != 255 || g != 255 || b != 255 {
				nonWhite++
			}
		}
	}
	return nonWhite, sampled, cw, ch, true
}
