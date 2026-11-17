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
// 真机验证结论（2026-11-17，微软拼音 HKL 0x08040804）：
//   - 候选窗坐标必须是 **Form 客户区**坐标；给屏幕坐标会让候选框
//     跑到屏幕左上角（示例踩过，很难看出是坐标系问题）；
//   - WM_IME_* 必须透传给 DefWindowProc，由系统维护组合窗/候选窗；
//     在默认处理**之后**重新 ImmSet* 才能压住系统给的默认位置；
//   - HIMC 不能长期缓存：输入法切换/失焦会换上下文，旧句柄上的
//     ImmSet* 静默失败（候选框不跟光标的元凶之一）；
//   - 透传的完成消息会让系统按结果串逐码元补发 WM_CHAR，
//     结果串已由 IMEBridge 写入，必须在上层吞掉补发（否则二次上屏）。

const (
	gcsCompReadStr = 0x0001
	gcsCompStr     = 0x0008
	gcsCursorPos   = 0x0080
	gcsResultStr   = 0x0800

	cfsPoint         = 0x0002
	cfsForcePosition = 0x0020
	cfsCandidatePos  = 0x0040
)

var (
	imm32 = syscall.NewLazyDLL("imm32.dll")

	procImmGetContext            = imm32.NewProc("ImmGetContext")
	procImmReleaseContext        = imm32.NewProc("ImmReleaseContext")
	procImmGetCompositionStringW = imm32.NewProc("ImmGetCompositionStringW")
	procImmSetCompositionWindow  = imm32.NewProc("ImmSetCompositionWindow")
	procImmSetCandidateWindow    = imm32.NewProc("ImmSetCandidateWindow")
)

// point 是 Windows POINT。
type point struct{ X, Y int32 }

// rect32 是 Windows RECT。
type rect32 struct{ Left, Top, Right, Bottom int32 }

// compositionForm 是 COMPOSITIONFORM（布局与 C 一致，勿加填充）。
type compositionForm struct {
	style      uint32
	currentPos point
	area       rect32
}

// candidateForm 是 CANDIDATEFORM。
type candidateForm struct {
	index      uint32
	style      uint32
	currentPos point
	area       rect32
}

// windowsIME 是 imm32 平台层。
type windowsIME struct {
	hwnd   uintptr
	bridge *IMEBridge
}

// newPlatformIME 创建平台 IME（Windows）。
func newPlatformIME(hwnd uintptr, b *IMEBridge) platformIME {
	return &windowsIME{hwnd: hwnd, bridge: b}
}

// context 取当前输入上下文（**按次获取**，调用方负责 release）。
//
// 不能长期缓存 HIMC：输入法切换、窗口失焦/获焦时系统会替换窗口的输入
// 上下文（WM_IME_SETCONTEXT 携带新句柄），旧句柄上的
// ImmSetCompositionWindow / ImmSetCandidateWindow 会**静默失败**。
// 真窗表现正是"候选框不跟光标，切换输入法后更明显"。
// canvas_ime 的已验证实现也是每次调用 ImmGetContext。
func (w *windowsIME) context() uintptr {
	if w.hwnd == 0 {
		return 0
	}
	h, _, _ := procImmGetContext.Call(w.hwnd)
	return h
}

// release 归还上下文。
func (w *windowsIME) release(h uintptr) {
	if w.hwnd != 0 && h != 0 {
		procImmReleaseContext.Call(w.hwnd, h)
	}
}

// Attach 检查输入上下文是否可用（真正取用按次进行）。
func (w *windowsIME) Attach() bool {
	h := w.context()
	if h == 0 {
		return false
	}
	w.release(h)
	return true
}

// Detach 无需释放：上下文按次获取/归还。
func (w *windowsIME) Detach() {}

// compositionString 读组合串（UTF-16 → UTF-8）。
func (w *windowsIME) compositionString(h uintptr, flag uint32) string {
	if h == 0 {
		return ""
	}
	n, _, _ := procImmGetCompositionStringW.Call(h, uintptr(flag), 0, 0)
	if int32(n) <= 0 {
		return ""
	}
	units := make([]uint16, int(n)/2)
	p := unsafe.Pointer(&units[0])
	procImmGetCompositionStringW.Call(h, uintptr(flag), uintptr(p), n)
	// 结果按 UTF-16LE 解，长度以返回字节数为准（可能不含结尾 0）。
	return utf16ToString(units)
}

// cursorPos 读组合串内光标（UTF-16 码元偏移）。
func (w *windowsIME) cursorPos(h uintptr) int32 {
	if h == 0 {
		return 0
	}
	n, _, _ := procImmGetCompositionStringW.Call(h, gcsCursorPos, 0, 0)
	return int32(n)
}

// HandleMessage 把一条 IME 消息翻译成归一化事件。
//
// 返回 handled=false 时调用方必须继续 DefWindowProc：
// 消息不该被吞，否则输入法的默认行为（比如系统候选窗）会失效。
func (w *windowsIME) HandleMessage(msg uint32, wParam, lParam uintptr) ([]input.Event, bool) {
	h := w.context()
	if h == 0 {
		return nil, false
	}
	defer w.release(h)
	switch msg {
	case wmIMEStartComposition:
		return w.bridge.Handle(IMFEvent{Kind: IMFStartComposition}), true

	case wmIMEComposition:
		flags := uint32(lParam)
		var out []input.Event
		if flags&gcsResultStr != 0 {
			if s := w.compositionString(h, gcsResultStr); s != "" {
				out = append(out, w.bridge.Handle(IMFEvent{Kind: IMFResult, Text: s})...)
			}
		}
		if flags&(gcsCompStr|gcsCompReadStr) != 0 {
			s := w.compositionString(h, gcsCompStr)
			out = append(out, w.bridge.Handle(IMFEvent{
				Kind: IMFUpdateComposition, Text: s, Cursor: w.cursorPos(h),
			})...)
		}
		return out, true

	case wmIMEEndComposition:
		return w.bridge.Handle(IMFEvent{Kind: IMFEndComposition}), true
	}
	return nil, false
}

// SetCandidates 把组合窗与候选窗都摆到光标处（**Form 客户区**坐标）。
//
// 两条实测约束（继承 canvas_ime）：
//   - ptCurrentPos 是 hwnd 客户区坐标，不是屏幕坐标（给屏幕坐标候选框会
//     跑到屏幕左上角）；
//   - 只设 COMPOSITIONFORM 不够：候选列表由 CANDIDATEFORM 决定，必须
//     同时 ImmSetCandidateWindow，否则候选框仍停在默认位置。
func (w *windowsIME) SetCandidates(r geom.Rect) {
	h := w.context()
	if h == 0 {
		return
	}
	defer w.release(h)
	// 组合串放在光标行**顶端**（与 canvas_ime 的已验证实现一致：真实
	// 输入法把组合窗锚在行首，候选列表再往下 20px；用行底会让候选框
	// 整体偏低一行的位置）。
	pt := point{X: int32(r.Min.X), Y: int32(r.Min.Y)}
	cf := compositionForm{
		style:      cfsPoint | cfsForcePosition,
		currentPos: pt,
		area: rect32{
			Left: pt.X, Top: pt.Y,
			Right: pt.X + 200, Bottom: pt.Y + 28,
		},
	}
	procImmSetCompositionWindow.Call(h, uintptr(unsafe.Pointer(&cf)))

	// 候选列表贴在光标行下方。
	candPos := point{X: pt.X, Y: pt.Y + 20}
	cand := candidateForm{
		style:      cfsCandidatePos,
		currentPos: candPos,
		area: rect32{
			Left: candPos.X, Top: candPos.Y,
			Right: candPos.X + 320, Bottom: candPos.Y + 160,
		},
	}
	procImmSetCandidateWindow.Call(h, uintptr(unsafe.Pointer(&cand)))
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
