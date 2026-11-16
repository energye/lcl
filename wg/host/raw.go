package host

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
)

// ── 原始消息归一化 ──────────────────────────────────────────────────
//
// 这一层把平台消息翻译成 input 包的事件，并且**只做翻译**：
//   - 它不认 HWND、不认 lcl 类型，所以能完全跑在无窗口自检里；
//   - 它也做状态补偿（Enter/Leave、连击、捕获丢失），因为这些都是
//     "平台不告诉你"的信息，交给每个组件自己记必然记漏。
//
// 平台行为三条（examples/wg 的实证）：
//   - WM_MOUSEMOVE 不带"已离开客户区"，必须靠 TrackMouseEvent + WM_MOUSELEAVE；
//   - 窗口外抬手时 Up 会丢（模态循环把捕获夺走），只有 WM_CAPTURECHANGED 通知；
//   - 滚轮一格是 120，且高精度滚轮会给 120 的零头（触摸板）。

// RawKind 是平台原始消息的归一化分类。
type RawKind uint8

const (
	// RawMove 指针移动（客户区像素坐标）。
	RawMove RawKind = iota
	// RawDown 按键按下。
	RawDown
	// RawUp 按键抬起。
	RawUp
	// RawWheel 滚轮（Wheel 字段为 120 的倍数或零头）。
	RawWheel
	// RawLeave 指针离开客户区（TrackMouseEvent 回调）。
	RawLeave
	// RawCaptureLost 捕获被系统/他人夺走（WM_CAPTURECHANGED）。
	RawCaptureLost
	// RawKeyDown 按键按下。
	RawKeyDown
	// RawKeyUp 按键抬起。
	RawKeyUp
	// RawChar 字符输入（WM_CHAR，已过 TranslateMessage）。
	RawChar
)

// String 返回可读名。
func (k RawKind) String() string {
	switch k {
	case RawMove:
		return "move"
	case RawDown:
		return "down"
	case RawUp:
		return "up"
	case RawWheel:
		return "wheel"
	case RawLeave:
		return "leave"
	case RawCaptureLost:
		return "capture-lost"
	case RawKeyDown:
		return "key-down"
	case RawKeyUp:
		return "key-up"
	case RawChar:
		return "char"
	default:
		return "unknown"
	}
}

// RawEvent 是一条平台消息的归一化视图（不含任何 Windows/LCL 类型）。
type RawEvent struct {
	Kind RawKind
	// X,Y 是客户区**像素**坐标（未乘 DPI）。
	X, Y int32
	// Button 是按钮序号：0 左、1 右、2 中、3 X1、4 X2。
	Button int
	// Buttons 是消息发生前的掩码（仅 RawMove 需要，用于判拖动/hover）。
	Buttons input.Buttons
	// Wheel 是滚轮增量（一格 120，向前为正）。
	Wheel int32
	// Key 是语义键码（VK 映射在 bind_key.go）。
	Key input.KeyCode
	// Text 是 RawChar 的字符。
	Text string
	// Repeat 表示系统自动重复。
	Repeat bool
	// Mods 是修饰键状态。
	Mods input.Modifiers
	// Time 必须来自**原生时间戳**（GetMessageTime），不是处理时刻。
	Time time.Time
}

const (
	// DoubleClickTime 是双击的最大间隔。
	DoubleClickTime = 500 * time.Millisecond
	// DoubleClickSlop 是两次按下之间允许的位移（DIP）。
	DoubleClickSlop = 4
	// WheelStep 是一个滚轮档位滚动的距离（DIP）。
	WheelStep = 48
)

// NormStats 是归一化统计。
type NormStats struct {
	Events  int
	Enters  int
	Leaves  int
	Cancels int
	Clicks  int
	Doubles int
	Dropped int
}

// Normalizer 把平台消息归一化成 input 事件。
//
// 有状态：指针位置、按键掩码、是否在客户区内、上一次按下的时刻与位置。
// 无状态化是不可能的——"是否拖动""是否双击"本质就是差分。
type Normalizer struct {
	dpi DPI

	pos     geom.Point
	hasPos  bool
	buttons input.Buttons
	inside  bool

	lastDownAt  time.Time
	lastDownPos geom.Point
	clickCount  int

	stats NormStats
}

// NewNormalizer 创建归一化器。
func NewNormalizer(d DPI) *Normalizer { return &Normalizer{dpi: d} }

// SetDPI 更新缩放（DPI 变化时调用，不重置状态）。
func (n *Normalizer) SetDPI(d DPI) { n.dpi = d }

// DPI 返回当前缩放。
func (n *Normalizer) DPI() DPI { return n.dpi }

// Stats 返回统计。
func (n *Normalizer) Stats() NormStats { return n.stats }

// Buttons 返回当前按键掩码。
func (n *Normalizer) Buttons() input.Buttons { return n.buttons }

// Inside 表示指针是否在客户区内。
func (n *Normalizer) Inside() bool { return n.inside }

// ClickCount 返回最近一次按下的连击次数。
func (n *Normalizer) ClickCount() int { return n.clickCount }

// Reset 清空全部状态（重挂载、失活、窗口重建时调用）。
//
// 不重置的话，重建后第一次移动会被算成"从旧位置飞过来"，Delta 巨大，
// 任何依赖速度的手势都会误判。
func (n *Normalizer) Reset() {
	n.hasPos = false
	n.buttons = 0
	n.inside = false
	n.lastDownAt = time.Time{}
	n.lastDownPos = geom.Point{}
	n.clickCount = 0
}

// Normalize 归一化一条原始消息，返回 0–2 个事件。
//
// 可能返回两个：指针从窗口外进入时先 Enter 再 Move/Hover。
// 顺序不能反——组件在 Enter 里注册状态、在 Move 里用状态。
func (n *Normalizer) Normalize(raw RawEvent) []input.Event {
	switch raw.Kind {
	case RawMove:
		return n.move(raw)
	case RawDown:
		return n.down(raw)
	case RawUp:
		return n.up(raw)
	case RawWheel:
		return n.wheel(raw)
	case RawLeave:
		return n.leave(raw)
	case RawCaptureLost:
		return n.captureLost(raw)
	case RawKeyDown, RawKeyUp:
		return n.key(raw)
	case RawChar:
		return n.char(raw)
	}
	return nil
}

// toRoot 把客户区像素换算成根空间 DIP。
func (n *Normalizer) toRoot(x, y int32) geom.Point {
	return geom.Pt(n.dpi.DIPI(x), n.dpi.DIPI(y))
}

// pointer 造一个指针事件（Buttons 取当前状态，Delta 取与上次的差）。
func (n *Normalizer) pointer(kind input.PointerKind, p geom.Point, raw RawEvent) *input.PointerEvent {
	e := &input.PointerEvent{
		Kind:    kind,
		Pointer: input.MousePointer,
		Device:  input.DeviceMouse,
		Pos:     p,
		PosRoot: p,
		Buttons: n.buttons,
		Mods:    raw.Mods,
		Time:    raw.Time,
	}
	if n.hasPos {
		e.Delta = p.Sub(n.pos)
	}
	return e
}

// enterIfNeeded 在指针首次进入客户区时补一个 Enter。
func (n *Normalizer) enterIfNeeded(p geom.Point, raw RawEvent) []input.Event {
	if n.inside {
		return nil
	}
	n.inside = true
	n.stats.Enters++
	e := n.pointer(input.PointerEnter, p, raw)
	n.stats.Events++
	return []input.Event{e}
}

// move 处理移动：有按键是 Move（拖动），无按键是 Hover。
func (n *Normalizer) move(raw RawEvent) []input.Event {
	p := n.toRoot(raw.X, raw.Y)
	out := n.enterIfNeeded(p, raw)

	kind := input.PointerHover
	if n.buttons != 0 {
		kind = input.PointerMove
	}
	e := n.pointer(kind, p, raw)
	n.pos = p
	n.hasPos = true
	n.stats.Events++
	return append(out, e)
}

// down 处理按下：顺带合成连击次数。
func (n *Normalizer) down(raw RawEvent) []input.Event {
	p := n.toRoot(raw.X, raw.Y)
	out := n.enterIfNeeded(p, raw)

	// 连击：间隔与位移都在阈值内才累加，否则回到 1。
	if !n.lastDownAt.IsZero() &&
		raw.Time.Sub(n.lastDownAt) <= DoubleClickTime &&
		p.Sub(n.lastDownPos).Len() <= DoubleClickSlop {
		n.clickCount++
	} else {
		n.clickCount = 1
	}
	n.lastDownAt = raw.Time
	n.lastDownPos = p

	n.buttons |= buttonMask(raw.Button)
	n.pos = p
	n.hasPos = true

	e := n.pointer(input.PointerDown, p, raw)
	e.ClickCount = n.clickCount
	n.stats.Events++
	n.stats.Clicks++
	if n.clickCount == 2 {
		n.stats.Doubles++
	}
	return append(out, e)
}

// up 处理抬起。
func (n *Normalizer) up(raw RawEvent) []input.Event {
	p := n.toRoot(raw.X, raw.Y)
	n.buttons &^= buttonMask(raw.Button)
	n.pos = p
	n.hasPos = true
	e := n.pointer(input.PointerUp, p, raw)
	n.stats.Events++
	return []input.Event{e}
}

// wheel 处理滚轮：一格 120 → WheelStep DIP，向前滚为负 Y（内容上移）。
func (n *Normalizer) wheel(raw RawEvent) []input.Event {
	p := n.toRoot(raw.X, raw.Y)
	out := n.enterIfNeeded(p, raw)

	notches := float32(raw.Wheel) / 120
	e := &input.WheelEvent{
		Pos:   p,
		Delta: geom.Pt(0, -notches*WheelStep),
		Mods:  raw.Mods,
		Time:  raw.Time,
	}
	n.pos = p
	n.hasPos = true
	n.stats.Events++
	return append(out, e)
}

// leave 处理离开客户区。
func (n *Normalizer) leave(raw RawEvent) []input.Event {
	if !n.inside {
		return nil
	}
	n.inside = false
	n.stats.Leaves++
	e := n.pointer(input.PointerLeave, n.pos, raw)
	n.stats.Events++
	return []input.Event{e}
}

// captureLost 处理捕获被夺走：必须补 Cancel，否则组件永远停在"按下"状态（C26）。
func (n *Normalizer) captureLost(raw RawEvent) []input.Event {
	had := n.buttons
	n.buttons = 0
	n.inside = false
	n.stats.Cancels++
	e := n.pointer(input.PointerCancel, n.pos, raw)
	// Cancel 要带上"被取消的按键"：组件靠它决定回滚哪个键的视觉状态。
	e.Buttons = had
	n.stats.Events++
	return []input.Event{e}
}

// key 处理按键。
func (n *Normalizer) key(raw RawEvent) []input.Event {
	kind := input.KeyDown
	if raw.Kind == RawKeyUp {
		kind = input.KeyUp
	} else if raw.Repeat {
		kind = input.KeyRepeat
	}
	n.stats.Events++
	return []input.Event{&input.KeyEvent{
		Kind: kind, Code: raw.Key, Mods: raw.Mods, Time: raw.Time,
	}}
}

// char 处理字符输入。
//
// C28：控制字符不得进入 Text；带 Ctrl/Alt 的组合只产生按键、不产生字符。
// 不在这里拦的话 Ctrl+A 会先进输入框（示例踩过），再想撤销就得清空重来。
func (n *Normalizer) char(raw RawEvent) []input.Event {
	if !printable(raw.Text, raw.Mods) {
		n.stats.Dropped++
		return nil
	}
	n.stats.Events++
	return []input.Event{&input.KeyEvent{
		Kind: input.KeyDown, Code: raw.Key, Text: raw.Text, Mods: raw.Mods, Time: raw.Time,
	}}
}

// printable 判定一个字符是否可以作为文本输入。
func printable(s string, mods input.Modifiers) bool {
	if s == "" {
		return false
	}
	if mods.Has(input.ModControl) || mods.Has(input.ModAlt) {
		return false
	}
	for _, r := range s {
		if r < 0x20 || r == 0x7f {
			return false
		}
	}
	return true
}

// buttonMask 把按钮序号换成位。
func buttonMask(i int) input.Buttons {
	switch i {
	case 0:
		return input.ButtonLeft
	case 1:
		return input.ButtonRight
	case 2:
		return input.ButtonMiddle
	case 3:
		return input.ButtonX1
	case 4:
		return input.ButtonX2
	}
	return 0
}
