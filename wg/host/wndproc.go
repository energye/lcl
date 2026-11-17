package host

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
)

// ── 消息路由 ────────────────────────────────────────────────────────
//
// 把窗口消息解码成 RawEvent、喂给 Normalizer、再推给应用。**不含平台类型**，
// 所以能在无窗口自检里跑完整的"消息 → 事件 → 状态"链路：这条链路的 bug
// （坐标没换、滚轮方向反、Up 丢失）恰恰是最难靠手点发现的。
//
// 平台胶水（SetOnWndProc / GetKeyState / TrackMouseEvent）在 bind_*.go。

// 窗口消息编号（只列本层用到的）。
const (
	wmKeyDown        = 0x0100
	wmKeyUp          = 0x0101
	wmChar           = 0x0102
	wmSysKeyDown     = 0x0104
	wmSysKeyUp       = 0x0105
	wmSysChar        = 0x0106
	wmMouseMove      = 0x0200
	wmLButtonDown    = 0x0201
	wmLButtonUp      = 0x0202
	wmLButtonDblClk  = 0x0203
	wmRButtonDown    = 0x0204
	wmRButtonUp      = 0x0205
	wmRButtonDblClk  = 0x0206
	wmMButtonDown    = 0x0207
	wmMButtonUp      = 0x0208
	wmMButtonDblClk  = 0x0209
	wmMouseWheel     = 0x020A
	wmCaptureChanged = 0x0215
	wmMouseLeave     = 0x02A3
	wmSize           = 0x0005
	wmSetFocus       = 0x0007
	wmKillFocus      = 0x0008
	wmActivate       = 0x0006
	wmDpiChanged     = 0x02E0
	wmSettingChange  = 0x001A
	wmTimer          = 0x0113
	wmEraseBkgnd     = 0x0014
	wmPaint          = 0x000F

	// 输入法消息（imm32.h）。组合消息必须继续走 DefWindowProc：
	// 平台默认处理创建的组合窗/候选窗才是用户看到的那一层。
	wmIMEStartComposition = 0x010D
	wmIMEEndComposition   = 0x010E
	wmIMEComposition      = 0x010F
	wmIMESetContext       = 0x0281
)

// 鼠标消息 wParam 的键位。
const (
	mkLButton  = 0x0001
	mkRButton  = 0x0002
	mkShift    = 0x0004
	mkControl  = 0x0008
	mkMButton  = 0x0010
	mkXButton1 = 0x0020
	mkXButton2 = 0x0040
)

// Message 是一条窗口消息（已去掉平台类型）。
type Message struct {
	Msg    uint32
	WParam uintptr
	LParam uintptr
	// Time 是原生时间戳（GetMessageTime），不能填处理时刻。
	Time time.Time
}

// EventSink 接收归一化事件（由 wg.Container 实现）。
type EventSink interface {
	// PushEvent 投递一个事件，返回是否被消费。
	PushEvent(ev input.Event) bool
	// RequestFrame 请求一帧（事件改变了状态就需要重绘）。
	RequestFrame()
}

// ResizeSink 是可选的尺寸变化接口。
//
// 只有 Host 更新 client 是不够的：合成缓冲、场景根节点、光栅缓存都必须
// 跟着换尺寸，否则 Presenter 因"帧尺寸 ≠ 客户区"一直拒绝提交，窗口就停在
// 首帧（真窗表现为底部/右侧一条白边，painted 不再增长）。
// Container 实现它，把尺寸变化转给节流的 scheduler。
type ResizeSink interface {
	ResizeRequest(w, h int)
}

// DPISink 由容器实现：DPI 变更必须走容器路径（c.mu 下重排/重建/重绘）。
//
// 平台回调只传事件：Surface 树不是线程安全的，消息线程直接改根画布会与
// 帧线程的合成竞争（真窗 -race 实测）。
type DPISink interface {
	DPIChanged(scale float32)
}

// FocusSink 由容器实现：窗口失活（WA_INACTIVE）时必须清焦点。
//
// 否则切到别的应用再回来，输入框还显示焦点框、IME 还挂在旧编辑模型上。
type FocusSink interface {
	BlurFocus()
}

// RouterStats 是路由统计。
type RouterStats struct {
	Messages int
	Decoded  int
	Handled  int
	Ignored  int
	Resizes  int
	DPI      int
}

// Router 解码消息并驱动输入链路。
type Router struct {
	norm   *Normalizer
	bridge *IMEBridge
	ime    platformIME
	sink   EventSink
	mods   input.Modifiers

	// toClient 把屏幕坐标换算成客户区坐标（滚轮消息自带屏幕坐标）。
	toClient func(x, y int32) (int32, int32)

	// OnResize 在收到尺寸变化时调用（节流在 frame.Scheduler）。
	OnResize func(w, h int)
	// OnDPI 在 DPI 变化时调用。
	OnDPI func(dpi DPI)

	stats RouterStats
}

// NewRouter 创建路由。
func NewRouter(norm *Normalizer, sink EventSink) *Router {
	return &Router{norm: norm, sink: sink}
}

// SetIME 挂上平台输入法。
func (r *Router) SetIME(ime platformIME, b *IMEBridge) {
	r.ime = ime
	r.bridge = b
}

// SetScreenToClient 设置屏幕 → 客户区换算（滚轮消息必需）。
func (r *Router) SetScreenToClient(fn func(x, y int32) (int32, int32)) { r.toClient = fn }

// AttachIME 取得输入法上下文（窗口获焦/接管消息后调用）。
func (r *Router) AttachIME() {
	if r.ime != nil {
		r.ime.Attach()
	}
}

// DetachIME 释放输入法上下文。
func (r *Router) DetachIME() {
	if r.ime != nil {
		r.ime.Detach()
	}
}

// SetIMECaret 把组合窗/候选窗摆到光标处（客户端像素坐标）。
func (r *Router) SetIMECaret(rect geom.Rect) {
	if r.ime != nil {
		r.ime.SetCandidates(rect)
	}
}

// SetMods 由胶水层按 GetKeyState 更新修饰键状态。
func (r *Router) SetMods(m input.Modifiers) { r.mods = m }

// Mods 返回当前修饰键状态。
func (r *Router) Mods() input.Modifiers { return r.mods }

// Norm 返回归一化器。
func (r *Router) Norm() *Normalizer { return r.norm }

// Stats 返回统计。
func (r *Router) Stats() RouterStats { return r.stats }

// Handle 处理一条消息，返回是否消费（false 表示调用方应继续 DefWindowProc）。
func (r *Router) Handle(m Message) bool {
	r.stats.Messages++

	// 上屏抑制：透传 WM_IME_COMPOSITION 后，系统会为 GCS_RESULTSTR 逐字
	// 补发 WM_CHAR。结果串已经写入编辑器，这些补发必须吞掉，否则二次上屏。
	if (m.Msg == wmChar || m.Msg == wmSysChar) && r.bridge != nil {
		if r.bridge.TakeCommitChar(uint16(m.WParam)) {
			r.stats.Handled++
			return true
		}
	}
	// 新的按键意味着上一次上屏的补发序列已经结束（或该 IME 根本不补发）。
	if (m.Msg == wmKeyDown || m.Msg == wmSysKeyDown) && r.bridge != nil {
		r.bridge.ClearSuppressed()
	}

	// 先给 IME：组合期间键盘输入属于输入法，不该再走按键路径，
	// 否则"拼音输入时按 A 会在输入框里出现字母"。
	//
	// 消费后仍返回 false：WM_IME_* 必须继续给 DefWindowProc。系统在默认
	// 处理里创建/维护组合窗与候选窗——吞掉消息后 IME 只能把候选框留在
	// 默认位置（真窗表现为候选框跑到窗口左上角/固定不动）。
	if r.ime != nil {
		if evs, ok := r.ime.HandleMessage(m.Msg, m.WParam, m.LParam); ok {
			r.stats.Handled++
			r.emit(evs...)
			return false
		}
	}

	switch m.Msg {
	case wmSetFocus:
		// 窗口获焦：补一次 IME 上下文（Install 时可能还拿不到）。
		r.AttachIME()
		return false

	case wmActivate:
		// 失活（WA_INACTIVE）：清焦点，避免回来时输入框仍持焦点/IME。
		if uint32(m.WParam&0xFFFF) == 0 {
			if fs, ok := r.sink.(FocusSink); ok {
				fs.BlurFocus()
			}
		}
		return false

	case wmSize:
		w := int(uint16(m.LParam & 0xFFFF))
		h := int(uint16((m.LParam >> 16) & 0xFFFF))
		r.stats.Resizes++
		if r.OnResize != nil {
			r.OnResize(w, h)
		}
		// 节流后的真正重建：更新合成缓冲/场景/缓存。
		if rs, ok := r.sink.(ResizeSink); ok {
			rs.ResizeRequest(w, h)
		}
		// WM_SIZE 本身要透传（窗口类需要它做默认处理）。
		return false

	case wmEraseBkgnd:
		// C3：吃掉 ERASEBKGND。不吃的后果是系统先刷一遍背景色，
		// 再由我们整块覆盖 —— 一帧里两次全窗写，肉眼可见闪烁。
		return true

	case wmPaint:
		if r.sink != nil {
			r.sink.RequestFrame()
		}
		return true
	}

	raws := r.decode(m)
	if len(raws) == 0 {
		r.stats.Ignored++
		return false
	}
	r.stats.Decoded += len(raws)
	r.stats.Handled++
	for _, raw := range raws {
		r.emit(r.norm.Normalize(raw)...)
	}
	return true
}

// emit 把事件推给 sink，并在"有消费"时请求一帧。
func (r *Router) emit(evs ...input.Event) {
	if r.sink == nil {
		return
	}
	consumed := false
	for _, ev := range evs {
		if r.sink.PushEvent(ev) {
			consumed = true
		}
	}
	if consumed {
		r.sink.RequestFrame()
	}
}

// decode 把消息解码成 0–2 条原始事件。
func (r *Router) decode(m Message) []RawEvent {
	// 坐标：lParam 低/高 16 位是有符号的客户区坐标。
	cx, cy := clientXY(m.LParam)
	btn := buttonFromMsg(m.Msg)
	t := m.Time
	if t.IsZero() {
		t = time.Now()
	}

	switch m.Msg {
	case wmMouseMove:
		return []RawEvent{{
			Kind: RawMove, X: cx, Y: cy,
			Buttons: buttonsFromWParam(m.WParam), Mods: r.mods, Time: t,
		}}

	// 双击消息也必须当"按下"解码：LCL 窗口类带 CS_DBLCLKS，快速连点时
	// Windows 只把后续按下投成 WM_*BUTTONDBLCLK。以前只认 *_DOWN，第二次
	// 点击被丢弃（随后孤立的 Up 也凑不成手势）——连点 8 次只触发 3 次。
	case wmLButtonDown, wmRButtonDown, wmMButtonDown,
		wmLButtonDblClk, wmRButtonDblClk, wmMButtonDblClk:
		return []RawEvent{{Kind: RawDown, X: cx, Y: cy, Button: btn, Mods: r.mods, Time: t}}

	case wmLButtonUp, wmRButtonUp, wmMButtonUp:
		return []RawEvent{{Kind: RawUp, X: cx, Y: cy, Button: btn, Mods: r.mods, Time: t}}

	case wmMouseWheel:
		// 滚轮消息的 lParam 是**屏幕**坐标，必须先换成客户区。
		sx, sy := clientXY(m.LParam)
		if r.toClient != nil {
			sx, sy = r.toClient(sx, sy)
		}
		delta := int32(int16((m.WParam >> 16) & 0xFFFF))
		return []RawEvent{{Kind: RawWheel, X: sx, Y: sy, Wheel: delta, Mods: r.mods, Time: t}}

	case wmMouseLeave:
		return []RawEvent{{Kind: RawLeave, Time: t}}

	case wmCaptureChanged:
		return []RawEvent{{Kind: RawCaptureLost, Time: t}}

	case wmKeyDown, wmSysKeyDown:
		vk := uint16(m.WParam)
		return []RawEvent{{
			Kind: RawKeyDown, Key: VKToCode(vk),
			Repeat: m.LParam&(1<<30) != 0, Mods: r.mods, Time: t,
		}}

	case wmKeyUp, wmSysKeyUp:
		return []RawEvent{{Kind: RawKeyUp, Key: VKToCode(uint16(m.WParam)), Mods: r.mods, Time: t}}

	case wmChar, wmSysChar:
		ch := rune(uint16(m.WParam))
		return []RawEvent{{Kind: RawChar, Text: string(ch), Mods: r.mods, Time: t}}

	case wmDpiChanged:
		r.stats.DPI++
		if r.OnDPI != nil {
			// 新 DPI 在 wParam 低 16 位。
			r.OnDPI(DPIFor(int(uint16(m.WParam & 0xFFFF))))
		}
		return nil
	}
	return nil
}

// clientXY 从 lParam 解出有符号的客户区坐标。
func clientXY(l uintptr) (int32, int32) {
	x := int32(int16(uint16(l & 0xFFFF)))
	y := int32(int16(uint16((l >> 16) & 0xFFFF)))
	return x, y
}

// buttonFromMsg 返回消息对应的按钮序号。
func buttonFromMsg(msg uint32) int {
	switch msg {
	case wmLButtonDown, wmLButtonUp, wmLButtonDblClk:
		return 0
	case wmRButtonDown, wmRButtonUp, wmRButtonDblClk:
		return 1
	case wmMButtonDown, wmMButtonUp, wmMButtonDblClk:
		return 2
	}
	return 0
}

// buttonsFromWParam 把鼠标消息 wParam 的键位换算成 Buttons。
func buttonsFromWParam(w uintptr) input.Buttons {
	var b input.Buttons
	if w&mkLButton != 0 {
		b |= input.ButtonLeft
	}
	if w&mkRButton != 0 {
		b |= input.ButtonRight
	}
	if w&mkMButton != 0 {
		b |= input.ButtonMiddle
	}
	if w&mkXButton1 != 0 {
		b |= input.ButtonX1
	}
	if w&mkXButton2 != 0 {
		b |= input.ButtonX2
	}
	return b
}

// modsFromWParam 把鼠标消息 wParam 的修饰位换算成 Modifiers。
func modsFromWParam(w uintptr) input.Modifiers {
	return modsFrom(w&mkShift != 0, w&mkControl != 0, false)
}

// ClientSize 从 lParam 解出 WM_SIZE 的宽高（低 16 位宽，高 16 位高）。
func ClientSize(l uintptr) geom.Size {
	return geom.Sz(float32(uint16(l&0xFFFF)), float32(uint16((l>>16)&0xFFFF)))
}
