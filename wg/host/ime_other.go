//go:build !windows

package host

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
)

// windowsIME 在其他平台没有等价物，这里只保留同名空实现，
// 让上层代码不需要在每个平台分叉。GTK 的 GtkIMContext 在 M5 接。
type windowsIME struct{}

// newPlatformIME 在非 Windows 平台返回 nil（无平台 IME）。
func newPlatformIME(hwnd uintptr, b *IMEBridge) platformIME { return nil }

func (w *windowsIME) Attach() bool { return false }
func (w *windowsIME) Detach()      {}
func (w *windowsIME) HandleMessage(msg uint32, wParam, lParam uintptr) ([]input.Event, bool) {
	return nil, false
}
func (w *windowsIME) SetCandidates(r geom.Rect) {}
