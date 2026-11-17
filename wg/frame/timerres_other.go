//go:build !windows

package frame

// ensureTimerResolution 在非 Windows 平台无操作（帧闹钟精度由平台保证）。
func ensureTimerResolution() {}
