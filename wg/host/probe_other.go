//go:build !windows

package host

// WindowPixelStats 在非 Windows 平台没有窗口 DC，返回 ok=false。
func WindowPixelStats(hwnd uintptr, step int) (nonWhite, sampled, cw, ch int, ok bool) {
	return 0, 0, 0, 0, false
}
