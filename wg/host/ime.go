package host

import (
	"time"
	"unicode/utf16"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
)

// ── 输入法桥 ────────────────────────────────────────────────────────
//
// 平台把 IME 消息变成"组合串的开始/更新/结束"，应用关心的是
// "把组合串画在光标处、带上虚线、提交时落定"。两者之间的翻译与状态机
// 放在这里，平台部分（imm32 / GTK IM / NSTextInputClient）只负责取数据。
//
// 这样切分的原因很实际：IME 行为在 CI 里点不出来，但"组合串状态机"
// 完全可以用构造的消息跑——M2 的 --inputtest 就靠这个覆盖上屏/取消/替换。
//
// 三条平台事实（§B7.2）：
//   - 候选窗位置必须用 **Form 客户区**坐标（ClientToParent），不是屏幕坐标；
//   - WM_IME_* 必须先透传（DefWindowProc），否则默认输入法窗口会自己画一份；
//   - 组合串要应用自己画（虚线+下划线），不能让系统画在客户区上。

// IMFKind 是平台 IME 消息的归一化分类。
type IMFKind uint8

const (
	// IMFStartComposition 开始组合（WM_IME_STARTCOMPOSITION）。
	IMFStartComposition IMFKind = iota
	// IMFUpdateComposition 组合串更新（WM_IME_COMPOSITION with GCS_COMPSTR）。
	IMFUpdateComposition
	// IMFEndComposition 结束组合（WM_IME_ENDCOMPOSITION）。
	IMFEndComposition
	// IMFResult 上屏结果（WM_IME_COMPOSITION with GCS_RESULTSTR）。
	IMFResult
)

// String 返回可读名。
func (k IMFKind) String() string {
	switch k {
	case IMFStartComposition:
		return "start"
	case IMFUpdateComposition:
		return "update"
	case IMFEndComposition:
		return "end"
	case IMFResult:
		return "result"
	default:
		return "unknown"
	}
}

// IMFEvent 是一条平台 IME 消息（已去掉 Windows 类型）。
type IMFEvent struct {
	Kind IMFKind
	// Text 是组合串/上屏串（UTF-16 已被平台解出，这里收 UTF-8）。
	Text string
	// Cursor 是组合串内的光标（**UTF-16 码元**偏移，来自 GCS_CURSORPOS）。
	Cursor int32
	Time   time.Time
}

// platformIME 是平台输入法的最小接口。
//
// 抽象只在这一点（§B7.3）：Windows 走 imm32，GTK 走 GtkIMContext（M5），
// macOS 走 NSTextInputClient（M5）。上层只见 IMEEvent 与 IMEBridge。
type platformIME interface {
	// Attach 取得平台输入上下文（拿不到返回 false，不是错误）。
	Attach() bool
	// Detach 释放上下文。
	Detach()
	// HandleMessage 把一条原生消息翻译成归一化事件；false 表示"不归我管"。
	HandleMessage(msg uint32, wParam, lParam uintptr) ([]input.Event, bool)
	// SetCandidates 把候选窗摆到光标处（客户区坐标）。
	SetCandidates(r geom.Rect)
}

// Composition 是组合状态。
type Composition struct {
	// Active 表示正在组合。
	Active bool
	// Text 是当前组合串。
	Text string
	// Cursor 是组合串内光标（rune 偏移）。
	Cursor int32
	// Start 是组合串在编辑文本中的起始位置（rune 偏移）。
	Start int32
	// First 表示还没有把第一次 Update 写进文本。
	// 有选区时第一次 Update 要覆盖整个选区，不是从光标末尾追加。
	First bool
	// Replaced 是第一次 Update 覆盖掉的原文，取消组合时必须还回去。
	Replaced string
}

// IMEStats 是 IME 桥的统计。
type IMEStats struct {
	Starts   int
	Updates  int
	Commits  int
	Cancels  int
	Empty    int
	UTF16Fix int
}

// IMEBridge 把平台 IME 消息翻译成对 TextEditor 的编辑操作。
//
// 不实现 TextInputClient：那是平台 →? 的方向；这里走的是
// 平台消息 → 编辑模型，方向相反容易绕晕，所以分开。
type IMEBridge struct {
	editor *input.TextEditor
	comp   Composition
	stats  IMEStats
	// now 可注入（门禁里用固定时钟）。
	now func() time.Time
}

// NewIMEBridge 创建桥。
func NewIMEBridge(e *input.TextEditor) *IMEBridge {
	return &IMEBridge{editor: e, now: time.Now}
}

// SetClock 注入时钟。
func (b *IMEBridge) SetClock(fn func() time.Time) {
	if fn != nil {
		b.now = fn
	}
}

// Composition 返回当前组合状态。
func (b *IMEBridge) Composition() Composition { return b.comp }

// Active 表示正在组合。
func (b *IMEBridge) Active() bool { return b.comp.Active }

// Stats 返回统计。
func (b *IMEBridge) Stats() IMEStats { return b.stats }

// Handle 处理一条平台 IME 消息，返回等价的 input 事件（供分发器广播）。
//
// 三条规则：
//   - Start 只记状态，不动文本（此时组合串通常还是空）；
//   - Update 用 **替换** 语义写组合区间，不是追加（IME 每次送全量）；
//   - Result 先把组合串落定，再插入上屏串——顺序反了会重复一遍。
func (b *IMEBridge) Handle(ev IMFEvent) []input.Event {
	if b.editor == nil {
		return nil
	}
	switch ev.Kind {
	case IMFStartComposition:
		b.stats.Starts++
		// 有选区时从选区**起点**开始组合，而不是 Cursor（=选区末端）。
		// 否则"选中 ab 后打中文"会得到 "ab好"：拼音追加在末尾，选区纹丝不动。
		v := b.editor.Value()
		b.comp = Composition{Active: true, Start: v.Selection.Start, First: true}
		return nil

	case IMFUpdateComposition:
		b.stats.Updates++
		text := ev.Text
		if !b.comp.Active {
			// 有些输入法不发 START 直接发 UPDATE：补一次开始。
			b.stats.Starts++
			b.comp = Composition{Active: true, Start: b.editor.Value().Selection.Start, First: true}
		}
		cur, fixed := utf16CursorToRune(text, ev.Cursor)
		if fixed {
			b.stats.UTF16Fix++
		}
		// 组合区间 = 上一次起始位置 + 旧组合串长度（若尚未写入则为光标处）。
		// 第一次写入时若编辑器有选区，改为覆盖整个选区，并记下原文以便取消时回滚。
		start := b.comp.Start
		end := start + input.RuneLen(b.comp.Text)
		if b.comp.First && b.comp.Text == "" {
			if v := b.editor.Value(); !v.Selection.IsEmpty() {
				sel := v.Selection.Clamp(v.Len())
				start, end = sel.Start, sel.End
				runes := []rune(v.Text)
				b.comp.Replaced = string(runes[sel.Start:sel.End])
			}
		}
		b.editor.SetComposingRegion(start, end, text)

		v := b.editor.Value()
		// 起始位置以编辑器实际写入位置为准：IME 给的 start 可能因为
		// 选区替换而失效，编辑器内部做了 clamp。
		newStart := v.Composing.Start
		b.comp.Text = text
		b.comp.Cursor = cur
		b.comp.Start = newStart
		b.comp.First = false
		return []input.Event{&input.IMEEvent{
			Kind: input.IMEComposition, Text: text, Cursor: cur,
			Rect: b.candidatesRect(), Time: b.stamp(ev.Time),
		}}

	case IMFResult:
		// 上屏要**替换**预编辑串：拼音 "zhong" 被汉字"中"顶掉。
		// 若此前没有组合（有些输入法直接送结果），退化为普通插入。
		if b.comp.Active || !b.editor.Value().Composing.IsEmpty() {
			b.editor.CommitComposingWith(ev.Text)
		} else if ev.Text != "" {
			b.editor.InsertText(ev.Text)
		}
		b.comp = Composition{}
		b.stats.Commits++
		return []input.Event{&input.IMEEvent{
			Kind: input.IMECommit, Text: ev.Text,
			Rect: b.candidatesRect(), Time: b.stamp(ev.Time),
		}}

	case IMFEndComposition:
		// 结束组合时组合串已经通过 Result 上屏；若没有 Result（用户取消），
		// 这里必须把它撤掉，否则光标处留下一串拼音。
		if b.comp.Active {
			if b.comp.Text != "" {
				b.editor.CancelComposing()
			}
			// 组合期间顶掉的选中文本要还回去：用户按 Esc 取消输入，
			// 不该连原先选中的内容一起丢掉。
			if b.comp.Replaced != "" {
				b.editor.InsertText(b.comp.Replaced)
			}
			if b.comp.Text != "" || b.comp.Replaced != "" {
				b.stats.Cancels++
			} else {
				b.stats.Empty++
			}
		}
		b.comp = Composition{}
		return []input.Event{&input.IMEEvent{
			Kind: input.IMECancel, Rect: b.candidatesRect(), Time: b.stamp(ev.Time),
		}}
	}
	return nil
}

// stamp 返回事件时间（平台没给就用注入的时钟）。
func (b *IMEBridge) stamp(t time.Time) time.Time {
	if !t.IsZero() {
		return t
	}
	if b.now != nil {
		return b.now()
	}
	return time.Now()
}

// candidatesRect 返回候选窗应该贴的位置：光标矩形下沿。
//
// 用光标矩形而不是输入框矩形：输入框长一行、光标在第 30 个字时，
// 候选窗贴在框左边会离眼睛很远（§B7.1）。
func (b *IMEBridge) candidatesRect() geom.Rect {
	if b.editor == nil || b.editor.CursorRect == nil {
		return geom.ZeroRect
	}
	return b.editor.CursorRect(b.editor.Value())
}

// ── UTF-16 ↔ rune ──────────────────────────────────────────────────

// utf16CursorToRune 把 UTF-16 码元偏移换算成 rune 偏移。
//
// 必须换算：Windows 给的是 GCS_CURSORPOS（UTF-16 码元），而编辑模型用 rune。
// 组合串里有 emoji/生僻字（代理对）时两者差 1，直接用会让光标落在半个字符上。
// 返回 fixed 表示确实发生了换算差异（诊断用）。
func utf16CursorToRune(s string, u16 int32) (int32, bool) {
	if u16 <= 0 {
		return 0, false
	}
	units := len(utf16.Encode([]rune(s)))
	if int(u16) >= units {
		r := input.RuneLen(s)
		return r, int(r) != int(u16)
	}
	// 逐码元累加，遇到代理对算 1 个 rune。
	var runes int32
	var used int32
	for _, r := range s {
		w := int32(1)
		if r > 0xFFFF {
			w = 2
		}
		if used+w > u16 {
			break
		}
		used += w
		runes++
	}
	return runes, runes != u16
}

// UTF16Len 返回 s 的 UTF-16 码元数（平台侧长度对账用）。
func UTF16Len(s string) int32 { return int32(len(utf16.Encode([]rune(s)))) }

// UTF16Slice 取 s 的 [start,end) UTF-16 码元切片（解 GCS_COMPSTR 用）。
func UTF16Slice(s string, start, end int32) string {
	units := utf16.Encode([]rune(s))
	if start < 0 {
		start = 0
	}
	if end > int32(len(units)) {
		end = int32(len(units))
	}
	if start >= end {
		return ""
	}
	return string(utf16.Decode(units[start:end]))
}

// IsComposing 判定编辑器是否正在组合（供 TextNode 决定是否画虚线）。
func IsComposing(e *input.TextEditor) bool {
	if e == nil {
		return false
	}
	return !e.Value().Composing.IsEmpty()
}

// CompositionUnderline 返回组合串的区间（画虚线用），无组合返回退化区间。
func CompositionUnderline(e *input.TextEditor) input.TextRange {
	if e == nil {
		return input.TextRange{}
	}
	return e.Value().Composing
}
