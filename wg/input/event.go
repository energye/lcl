// Package input 是输入层：归一化事件、命中测试、分发（捕获→目标→冒泡）、
// 指针捕获与 hover 差分。
//
// 两条边界：
//   - **不碰原生消息**。原生 → 归一化的映射在 host（SetOnMouseDown、
//     WM_MOUSELEAVE、imm32 等），本包只认已经归一化的时间戳与 DIP 坐标。
//     这样无窗口自检可以完整跑一遍命中与分发，不需要开窗口。
//   - **不改场景树**。命中只读，状态（捕获、hover、光标）都在 Dispatcher 上。
//
// 坐标约定：事件坐标是**根/画布空间**（DIP，未乘 DPI）。命中测试沿树向下
// 逐层求逆映射，因此旋转/缩放过的节点也能精确命中。
package input

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
)

// EventKind 是事件大类（分派时先按它选通道）。
type EventKind uint8

const (
	// EventPointer 指针（鼠标/触摸/笔）。
	EventPointer EventKind = iota
	// EventKey 键盘。
	EventKey
	// EventWheel 滚轮。
	EventWheel
	// EventIME 输入法。
	EventIME
	// EventFocus 焦点。
	EventFocus
)

// String 返回可读名。
func (k EventKind) String() string {
	switch k {
	case EventPointer:
		return "pointer"
	case EventKey:
		return "key"
	case EventWheel:
		return "wheel"
	case EventIME:
		return "ime"
	case EventFocus:
		return "focus"
	default:
		return "unknown"
	}
}

// Event 是事件的最小公共契约。
type Event interface {
	EventKind() EventKind
}

// PointerKind 是指针事件类型。
type PointerKind uint8

const (
	// PointerDown 按下。
	PointerDown PointerKind = iota
	// PointerMove 移动（有按键时也走这里，`Buttons` 说明按了谁）。
	PointerMove
	// PointerUp 抬起。
	PointerUp
	// PointerCancel 取消（被系统夺走焦点、进入模态循环等）。
	PointerCancel
	// PointerEnter 指针进入节点。
	PointerEnter
	// PointerLeave 指针离开节点。
	PointerLeave
	// PointerHover 无按键悬浮（"hover 只影响状态，不请求帧"）。
	PointerHover
)

// String 返回可读名。
func (k PointerKind) String() string {
	switch k {
	case PointerDown:
		return "down"
	case PointerMove:
		return "move"
	case PointerUp:
		return "up"
	case PointerCancel:
		return "cancel"
	case PointerEnter:
		return "enter"
	case PointerLeave:
		return "leave"
	case PointerHover:
		return "hover"
	default:
		return "unknown"
	}
}

// PointerID 是指针标识。鼠标恒为 MousePointer。
type PointerID uint32

// MousePointer 是鼠标的固定指针 id（与触摸/笔区分）。
const MousePointer PointerID = 1

// DeviceKind 是输入设备。
type DeviceKind uint8

const (
	// DeviceMouse 鼠标。
	DeviceMouse DeviceKind = iota
	// DeviceTouch 触摸。
	DeviceTouch
	// DevicePen 笔。
	DevicePen
)

// Buttons 是按键位集。
type Buttons uint8

const (
	// ButtonLeft 左键。
	ButtonLeft Buttons = 1 << iota
	// ButtonRight 右键。
	ButtonRight
	// ButtonMiddle 中键。
	ButtonMiddle
	// ButtonX1 侧键 1。
	ButtonX1
	// ButtonX2 侧键 2。
	ButtonX2
)

// Has 判定是否包含某个键。
func (b Buttons) Has(v Buttons) bool { return b&v != 0 }

// Modifiers 是修饰键位集。
type Modifiers uint8

const (
	// ModShift Shift。
	ModShift Modifiers = 1 << iota
	// ModControl Ctrl。
	ModControl
	// ModAlt Alt。
	ModAlt
	// ModMeta Win/Cmd。
	ModMeta
)

// Has 判定是否包含某个修饰键。
func (m Modifiers) Has(v Modifiers) bool { return m&v != 0 }

// PointerEvent 是指针事件。
//
// Time 必须来自**原生时间戳**（不是处理时刻）：手势判定、双击间隔、
// 速度估计全依赖它，用 time.Now() 会把派发延迟算进用户的动作时长里。
type PointerEvent struct {
	Kind     PointerKind
	Pointer  PointerID
	Device   DeviceKind
	Pos      geom.Point // 当前空间（分发时由 Dispatcher 换算到节点空间）
	PosRoot  geom.Point // 根空间，供浮层定位
	Delta    geom.Point
	Buttons  Buttons
	Mods     Modifiers
	Pressure float32
	Time     time.Time
	// ClickCount 是本次按下的连击次数（1 = 单击，2 = 双击，其余事件为 0）。
	//
	// 由 host 依"间隔 + 位移"合成，不直接用 WM_LBUTTONDBLCLK：平台的双击
	// 判定用的是系统鼠标设置（500ms/4px 之外的另一种阈值），与应用的 slop
	// 不一致时会出现"系统说双击、组件判单击"；而且触摸/笔根本不发 DBLCLK。
	ClickCount int
}

// EventKind 实现 Event。
func (*PointerEvent) EventKind() EventKind { return EventPointer }

// Down 表示这是按下事件。
func (e *PointerEvent) Down() bool { return e.Kind == PointerDown }

// Up 表示这是抬起事件。
func (e *PointerEvent) Up() bool { return e.Kind == PointerUp }

// Cancel 表示这是取消事件。
func (e *PointerEvent) Cancel() bool { return e.Kind == PointerCancel }

// Drag 表示按下状态下的移动。
func (e *PointerEvent) Drag() bool { return e.Kind == PointerMove && e.Buttons != 0 }

// KeyCode 是归一化物理键码（映射表在 host，本包只认语义）。
type KeyCode uint16

// 常用键码（其余由 host 映射到这些语义键或透传 VK）。
//
// 前缀 Code 是刻意的：KeyKind 里已有 KeyDown/KeyUp（事件类型），
// 键码再叫 KeyUp 就会撞名——"抬起事件"和"↑键"是两件事，名字必须分开。
//
// 为什么不用平台 VK 直传：快捷键表（Ctrl+S）必须跨平台一致，
// 而 VK 是 Windows 的值，GTK/macOS 各有自己的编码。语义键码是
// 平台层唯一的映射终点（映射表在 host/keys.go）。
const (
	CodeUnknown KeyCode = iota
	CodeBackspace
	CodeTab
	CodeEnter
	CodeEscape
	CodeSpace
	CodeDelete
	CodeInsert
	CodeLeft
	CodeRight
	CodeUp
	CodeDown
	CodeHome
	CodeEnd
	CodePageUp
	CodePageDown
	CodeShift
	CodeControl
	CodeAlt
	CodeMeta
	CodeCapsLock
)

// 数字行 0–9。
const (
	Code0 KeyCode = 100 + iota
	Code1
	Code2
	Code3
	Code4
	Code5
	Code6
	Code7
	Code8
	Code9
)

// 字母 A–Z。
const (
	CodeA KeyCode = 200 + iota
	CodeB
	CodeC
	CodeD
	CodeE
	CodeF
	CodeG
	CodeH
	CodeI
	CodeJ
	CodeK
	CodeL
	CodeM
	CodeN
	CodeO
	CodeP
	CodeQ
	CodeR
	CodeS
	CodeT
	CodeU
	CodeV
	CodeW
	CodeX
	CodeY
	CodeZ
)

// 功能键 F1–F12。
const (
	CodeF1 KeyCode = 300 + iota
	CodeF2
	CodeF3
	CodeF4
	CodeF5
	CodeF6
	CodeF7
	CodeF8
	CodeF9
	CodeF10
	CodeF11
	CodeF12
)

// 标点与符号（按美式键盘物理位置，不随布局变化）。
const (
	CodeMinus KeyCode = 400 + iota
	CodeEqual
	CodeBracketLeft
	CodeBracketRight
	CodeBackslash
	CodeSemicolon
	CodeQuote
	CodeComma
	CodePeriod
	CodeSlash
	CodeBackquote
	CodeNumpadAdd
	CodeNumpadSub
	CodeNumpadMul
	CodeNumpadDiv
	CodeNumpadDecimal
)

// KeyKind 是键盘事件类型。
type KeyKind uint8

const (
	// KeyDown 按下。
	KeyDown KeyKind = iota
	// KeyUp 抬起。
	KeyUp
	// KeyRepeat 系统重复。
	KeyRepeat
)

// KeyEvent 是键盘事件。
//
// **文本与按键分离**：Code 决定语义（退格/回车/Tab/方向键/快捷键），
// Text 只承载可打印字符。示例已实证 UTF8KeyPress 会把 0x08/0x0D/0x09
// 一起送进来，不分离就会出现"退格变方块、Ctrl+A 被吃掉"。
type KeyEvent struct {
	Kind KeyKind
	Code KeyCode
	Text string
	Mods Modifiers
	Time time.Time
}

// EventKind 实现 Event。
func (*KeyEvent) EventKind() EventKind { return EventKey }

// Printable 表示是否携带可插入文本。
func (e *KeyEvent) Printable() bool { return e.Text != "" }

// WheelEvent 是滚轮事件。
type WheelEvent struct {
	Pos   geom.Point
	Delta geom.Point
	Mods  Modifiers
	Time  time.Time
}

// EventKind 实现 Event。
func (*WheelEvent) EventKind() EventKind { return EventWheel }

// IMEKind 是输入法事件类型。
type IMEKind uint8

const (
	// IMEComposition 预编辑串更新。
	IMEComposition IMEKind = iota
	// IMECommit 上屏。
	IMECommit
	// IMECancel 取消预编辑。
	IMECancel
)

// IMEEvent 是输入法事件。Cursor 是预编辑串内的光标（rune 偏移）。
type IMEEvent struct {
	Kind   IMEKind
	Text   string
	Cursor int32
	Rect   geom.Rect
	Time   time.Time
}

// EventKind 实现 Event。
func (*IMEEvent) EventKind() EventKind { return EventIME }

// FocusKind 是焦点事件类型。
type FocusKind uint8

const (
	// FocusIn 获得焦点。
	FocusIn FocusKind = iota
	// FocusOut 失去焦点。
	FocusOut
)

// FocusEvent 是焦点事件。Node 为 nil 表示焦点移出整棵树。
type FocusEvent struct {
	Kind FocusKind
	Node *scene.Node
	Time time.Time
}

// EventKind 实现 Event。
func (*FocusEvent) EventKind() EventKind { return EventFocus }

// String 返回键码的可读名（诊断与 --inputtest 输出用）。
func (c KeyCode) String() string {
	switch c {
	case CodeUnknown:
		return "unk"
	case CodeBackspace:
		return "backspace"
	case CodeTab:
		return "tab"
	case CodeEnter:
		return "enter"
	case CodeEscape:
		return "escape"
	case CodeSpace:
		return "space"
	case CodeDelete:
		return "delete"
	case CodeInsert:
		return "insert"
	case CodeLeft:
		return "left"
	case CodeRight:
		return "right"
	case CodeUp:
		return "up"
	case CodeDown:
		return "down"
	case CodeHome:
		return "home"
	case CodeEnd:
		return "end"
	case CodePageUp:
		return "pageup"
	case CodePageDown:
		return "pagedown"
	case CodeShift:
		return "shift"
	case CodeControl:
		return "control"
	case CodeAlt:
		return "alt"
	case CodeMeta:
		return "meta"
	case CodeCapsLock:
		return "capslock"
	}
	if c >= Code0 && c <= Code9 {
		return string(rune('0' + (c - Code0)))
	}
	if c >= CodeA && c <= CodeZ {
		return string(rune('a' + (c - CodeA)))
	}
	if c >= CodeF1 && c <= CodeF12 {
		return "f" + itoa(int(c-CodeF1)+1)
	}
	switch c {
	case CodeMinus:
		return "-"
	case CodeEqual:
		return "="
	case CodeBracketLeft:
		return "["
	case CodeBracketRight:
		return "]"
	case CodeBackslash:
		return "\\"
	case CodeSemicolon:
		return ";"
	case CodeQuote:
		return "'"
	case CodeComma:
		return ","
	case CodePeriod:
		return "."
	case CodeSlash:
		return "/"
	case CodeBackquote:
		return "`"
	case CodeNumpadAdd:
		return "numpad+"
	case CodeNumpadSub:
		return "numpad-"
	case CodeNumpadMul:
		return "numpad*"
	case CodeNumpadDiv:
		return "numpad/"
	case CodeNumpadDecimal:
		return "numpad."
	}
	return "code(" + itoa(int(c)) + ")"
}

// itoa 是极小整数转字符串（避免为诊断引入 strconv 依赖热路径）。
func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	neg := v < 0
	if neg {
		v = -v
	}
	var b [12]byte
	i := len(b)
	for v > 0 {
		i--
		b[i] = byte('0' + v%10)
		v /= 10
	}
	if neg {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}
