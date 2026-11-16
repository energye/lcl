//go:build windows

package host

import (
	"syscall"
	"unsafe"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
)

// ── Windows 输入法（imm32） ─────────────────────────────────────────
//
// 只做三件事：开上下文、取组合串、摆候选窗。所有状态机在 IMEBridge 里，
// 这样这部分即使出错也只影响"取不到数据"，不会让编辑模型错乱。
//
// 两条实测约束（§B7.2）：
//   - 候选窗坐标必须是 **Form 客户区**坐标；给屏幕坐标会让候选框
//     跑到屏幕左上角（示例踩过，很难看出是坐标系问题）；
//   - WM_IME_* 必须透传给 DefWindowProc，否则系统会在客户区再画一份
//     默认的组合窗，与自绘的虚线框重叠。

const (
	gcsCompReadStr = 0x0001
	gcsCompStr     = 0x0008
	gcsCursorPos   = 0x0080
	gcsResultStr   = 0x0800

	cfsPoint = 0x0002

	wmIMEStartComposition = 0x010D
	wmIMEEndComposition   = 0x010E
	wmIMEComposition      = 0x010F
)

var (
	imm32 = syscall.NewLazyDLL("imm32.dll")

	procImmGetContext            = imm32.NewProc("ImmGetContext")
	procImmReleaseContext        = imm32.NewProc("ImmReleaseContext")
	procImmGetCompositionStringW = imm32.NewProc("ImmGetCompositionStringW")
	procImmSetCompositionWindow  = imm32.NewProc("ImmSetCompositionWindow")
)

// point 是 Windows POINT。
type point struct{ X, Y int32 }

// rect32 是 Windows RECT。
type rect32 struct{ Left, Top, Right, Bottom int32 }

// compositionForm 是 COMPOSITIONFORM。
type compositionForm struct {
	style      uint32
	currentPos point
	area       rect32
	_          [4]byte
}

// windowsIME 是 imm32 平台层。
type windowsIME struct {
	hwnd   uintptr
	himc   uintptr
	bridge *IMEBridge
}

// newPlatformIME 创建平台 IME（Windows）。
func newPlatformIME(hwnd uintptr, b *IMEBridge) platformIME {
	return &windowsIME{hwnd: hwnd, bridge: b}
}

// Attach 取得输入上下文。取不到时返回 false（无 IME 环境，不是错误）。
func (w *windowsIME) Attach() bool {
	if w.hwnd == 0 {
		return false
	}
	h, _, _ := procImmGetContext.Call(w.hwnd)
	w.himc = h
	return h != 0
}

// Detach 释放输入上下文。
func (w *windowsIME) Detach() {
	if w.himc != 0 {
		procImmReleaseContext.Call(w.hwnd, w.himc)
		w.himc = 0
	}
}

// compositionString 读组合串（UTF-16 → UTF-8）。
func (w *windowsIME) compositionString(flag uint32) string {
	if w.himc == 0 {
		return ""
	}
	n, _, _ := procImmGetCompositionStringW.Call(w.himc, uintptr(flag), 0, 0)
	if int32(n) <= 0 {
		return ""
	}
	units := make([]uint16, int(n)/2)
	p := unsafe.Pointer(&units[0])
	procImmGetCompositionStringW.Call(w.himc, uintptr(flag), uintptr(p), n)
	// 结果按 UTF-16LE 解，长度以返回字节数为准（可能不含结尾 0）。
	return utf16ToString(units)
}

// cursorPos 读组合串内光标（UTF-16 码元偏移）。
func (w *windowsIME) cursorPos() int32 {
	if w.himc == 0 {
		return 0
	}
	n, _, _ := procImmGetCompositionStringW.Call(w.himc, gcsCursorPos, 0, 0)
	return int32(n)
}

// HandleMessage 把一条 IME 消息翻译成归一化事件。
//
// 返回 handled=false 时调用方必须继续 DefWindowProc：
// 消息不该被吞，否则输入法的默认行为（比如系统候选窗）会失效。
func (w *windowsIME) HandleMessage(msg uint32, wParam, lParam uintptr) ([]input.Event, bool) {
	if w.himc == 0 {
		return nil, false
	}
	switch msg {
	case wmIMEStartComposition:
		return w.bridge.Handle(IMFEvent{Kind: IMFStartComposition}), true

	case wmIMEComposition:
		flags := uint32(lParam)
		var out []input.Event
		if flags&gcsResultStr != 0 {
			if s := w.compositionString(gcsResultStr); s != "" {
				out = append(out, w.bridge.Handle(IMFEvent{Kind: IMFResult, Text: s})...)
			}
		}
		if flags&(gcsCompStr|gcsCompReadStr) != 0 {
			s := w.compositionString(gcsCompStr)
			out = append(out, w.bridge.Handle(IMFEvent{
				Kind: IMFUpdateComposition, Text: s, Cursor: w.cursorPos(),
			})...)
		}
		return out, true

	case wmIMEEndComposition:
		return w.bridge.Handle(IMFEvent{Kind: IMFEndComposition}), true
	}
	return nil, false
}

// SetCandidates 把候选窗摆到光标处（客户区坐标）。
func (w *windowsIME) SetCandidates(r geom.Rect) {
	if w.himc == 0 {
		return
	}
	cf := compositionForm{
		style:      cfsPoint,
		currentPos: point{X: int32(r.Min.X), Y: int32(r.Max.Y)},
	}
	procImmSetCompositionWindow.Call(w.himc, uintptr(unsafe.Pointer(&cf)))
}

// utf16ToString 把 UTF-16 码元序列转成 Go 字符串（去结尾 0）。
func utf16ToString(units []uint16) string {
	for i, u := range units {
		if u == 0 {
			units = units[:i]
			break
		}
	}
	buf := make([]rune, 0, len(units))
	for i := 0; i < len(units); i++ {
		u := units[i]
		if u >= 0xD800 && u <= 0xDBFF && i+1 < len(units) {
			lo := units[i+1]
			if lo >= 0xDC00 && lo <= 0xDFFF {
				buf = append(buf, rune(0x10000+(int32(u-0xD800)<<10)+int32(lo-0xDC00)))
				i++
				continue
			}
		}
		buf = append(buf, rune(u))
	}
	return string(buf)
}
