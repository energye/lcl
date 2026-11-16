package input

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/energye/lcl/wg/geom"
)

// ── 文本输入 ────────────────────────────────────────────────────────
//
// 分工（04-UI-INTERACTION.md §B7.1）：
//   - TextEditor 是**模型**：文本、选区、组合串、撤销；
//   - TextInputConnection 是**平台连接**：host 实现，把编辑状态告诉 IME；
//   - TextInputClient 是平台回调进来的**接口**（本文件由 TextEditor 实现）。
//
// 偏移单位统一用 **rune**，不是 UTF-16 码元、也不是字节：
//   - 字节偏移在中文/emoji 上会把一个字符切成两半（切出无效 UTF-8）；
//   - UTF-16 偏移在 Go 里每一步都要转换，容易漏。
//   Windows IME 给的是 UTF-16 码元，转换集中在 host/ime_*.go 一处，
//   上层永远只见 rune 偏移。

// TextRange 是文本区间（rune 偏移，Start ≤ End）。
type TextRange struct {
	Start int32
	End   int32
}

// RangeAt 返回 i 处的空区间（光标）。
func RangeAt(i int32) TextRange { return TextRange{Start: i, End: i} }

// RangeOf 返回 [start,end) 区间（自动纠正顺序）。
func RangeOf(start, end int32) TextRange {
	if start > end {
		start, end = end, start
	}
	return TextRange{Start: start, End: end}
}

// Len 返回区间长度。
func (r TextRange) Len() int32 {
	if r.End <= r.Start {
		return 0
	}
	return r.End - r.Start
}

// IsEmpty 表示光标（非选区）。
func (r TextRange) IsEmpty() bool { return r.End <= r.Start }

// Valid 判定区间是否落在长度 n 的文本内。
func (r TextRange) Valid(n int32) bool {
	return r.Start >= 0 && r.Start <= r.End && r.End <= n
}

// Clamp 把区间夹到 [0,n]。
func (r TextRange) Clamp(n int32) TextRange {
	s, e := r.Start, r.End
	if s < 0 {
		s = 0
	}
	if e < 0 {
		e = 0
	}
	if s > n {
		s = n
	}
	if e > n {
		e = n
	}
	if s > e {
		s = e
	}
	return TextRange{Start: s, End: e}
}

// Action 是输入法/键盘触发的语义动作。
type Action uint8

const (
	// ActionNone 无动作。
	ActionNone Action = iota
	// ActionDone 完成（单行回车、失焦）。
	ActionDone
	// ActionNext 下一个字段。
	ActionNext
	// ActionSearch 搜索。
	ActionSearch
	// ActionGo 前往。
	ActionGo
	// ActionNewline 换行（多行）。
	ActionNewline
)

// String 返回可读名。
func (a Action) String() string {
	switch a {
	case ActionDone:
		return "done"
	case ActionNext:
		return "next"
	case ActionSearch:
		return "search"
	case ActionGo:
		return "go"
	case ActionNewline:
		return "newline"
	default:
		return "none"
	}
}

// EditingValue 是编辑状态快照（不可变语义：每次编辑产生新值）。
//
// Composing 是**Text 内部**的组合区间（不是另一份文本）：IME 组合串
// 与已确认文本共存于同一个字符串，光标与选区才能用同一套偏移表达。
// 把组合串单独存放会立刻出现"两个光标"问题（一个是文本里的，一个是组合串里的）。
type EditingValue struct {
	Text      string
	Selection TextRange
	Composing TextRange
}

// Cursor 返回光标位置。
func (v EditingValue) Cursor() int32 { return v.Selection.End }

// HasSelection 表示存在选区。
func (v EditingValue) HasSelection() bool { return !v.Selection.IsEmpty() }

// SelectedText 返回选中的文本。
func (v EditingValue) SelectedText() string {
	r := v.Selection.Clamp(int32(len([]rune(v.Text))))
	runes := []rune(v.Text)
	return string(runes[r.Start:r.End])
}

// ComposingText 返回组合串。
func (v EditingValue) ComposingText() string {
	if v.Composing.IsEmpty() {
		return ""
	}
	r := v.Composing.Clamp(int32(len([]rune(v.Text))))
	runes := []rune(v.Text)
	return string(runes[r.Start:r.End])
}

// Len 返回文本的 rune 数。
func (v EditingValue) Len() int32 { return int32(len([]rune(v.Text))) }

// normalize 修好越界区间（外部直接赋值时可能给出无效值）。
func (v EditingValue) normalize() EditingValue {
	n := v.Len()
	v.Selection = v.Selection.Clamp(n)
	v.Composing = v.Composing.Clamp(n)
	if v.Composing.Start < 0 || v.Composing.End < 0 {
		v.Composing = TextRange{}
	}
	return v
}

// EditStats 是编辑统计。
type EditStats struct {
	Inserts    int
	Deletes    int
	Commits    int
	Actions    int
	Dropped    int
	Undos      int
	Selections int
}

// TextInputConnection 是编辑模型与平台键盘/输入法之间的连接。
//
// 由 host 实现（Windows 下最终落到 imm32）。单独抽出来是为了让
// 无窗口自检能用一个假的 connection 跑完整编辑流程——IME 行为没法在
// CI 里点出来，但"编辑模型对组合串的处理"完全可以。
type TextInputConnection interface {
	// SetEditingState 告诉平台当前文本与选区（决定候选框与光标）。
	SetEditingState(v EditingValue)
	// SetClientRect 告诉平台光标矩形（候选框贴着光标，§B7.1）。
	SetClientRect(r geom.Rect)
	// Close 断开（提交或取消组合串由实现决定，默认提交）。
	Close()
}

// TextEditor 是可编辑文本的模型，同时实现 TextInputClient。
//
// 为什么模型与控件分离：同一个编辑模型要服务输入框、密码框、富文本的
// 纯文本模式、以及表格内联编辑。把它写进控件里这些场景各要一份，
// 而"光标在哪、选区多大、组合串到哪"的规则必须处处一致。
type TextEditor struct {
	value     EditingValue
	conn      TextInputConnection
	multiline bool

	undo  []EditingValue
	undos int

	// OnChanged 在每次编辑后调用（含组合串变化）。
	OnChanged func(EditingValue)
	// OnAction 在语义动作发生时调用。
	OnAction func(Action)
	// CursorRect 由控件提供，用于把光标矩形交给平台（候选框定位）。
	CursorRect func(EditingValue) geom.Rect

	stats EditStats
}

// NewTextEditor 创建编辑模型。
func NewTextEditor(multiline bool) *TextEditor {
	return &TextEditor{multiline: multiline}
}

// Multiline 表示是否多行。
func (e *TextEditor) Multiline() bool { return e.multiline }

// Value 返回当前编辑状态。
func (e *TextEditor) Value() EditingValue { return e.value }

// Stats 返回统计。
func (e *TextEditor) Stats() EditStats { return e.stats }

// Connected 表示是否已接上平台键盘。
func (e *TextEditor) Connected() bool { return e.conn != nil }

// ── 连接生命周期 ────────────────────────────────────────────────────

// Connect 接上平台连接（焦点进入）：立刻同步一次编辑状态与光标矩形。
func (e *TextEditor) Connect(c TextInputConnection) {
	if c == nil || e.conn == c {
		return
	}
	if e.conn != nil {
		e.conn.Close()
	}
	e.conn = c
	e.sync()
}

// Disconnect 断开（焦点离开）：先提交组合串再关连接。
//
// 顺序不能反：先关连接的话，IME 会把未提交的组合串丢掉，
// 用户打了一半的拼音就凭空消失。
func (e *TextEditor) Disconnect() {
	if e.conn == nil {
		return
	}
	e.CommitComposing()
	e.conn.Close()
	e.conn = nil
}

// sync 把状态推给平台。
func (e *TextEditor) sync() {
	if e.conn == nil {
		return
	}
	e.conn.SetEditingState(e.value)
	if e.CursorRect != nil {
		e.conn.SetClientRect(e.CursorRect(e.value))
	}
}

// changed 通知外部并同步平台。
func (e *TextEditor) changed() {
	if e.OnChanged != nil {
		e.OnChanged(e.value)
	}
	e.sync()
}

// pushUndo 记录撤销点。
func (e *TextEditor) pushUndo() {
	e.undo = append(e.undo, e.value)
	if len(e.undo) > 64 {
		e.undo = e.undo[1:]
	}
}

// Undo 回退一次编辑。
func (e *TextEditor) Undo() bool {
	if len(e.undo) == 0 {
		return false
	}
	e.value = e.undo[len(e.undo)-1]
	e.undo = e.undo[:len(e.undo)-1]
	e.stats.Undos++
	e.changed()
	return true
}

// ── 值操作 ──────────────────────────────────────────────────────────

// SetValue 外部赋值（受控输入/初始化）。会清空撤销栈。
func (e *TextEditor) SetValue(v EditingValue) {
	e.value = v.normalize()
	e.undo = e.undo[:0]
	e.changed()
}

// SetSelection 设置选区（光标）。
func (e *TextEditor) SetSelection(r TextRange) {
	e.value.Selection = r.Clamp(e.value.Len())
	e.stats.Selections++
	e.changed()
}

// SelectAll 全选。
func (e *TextEditor) SelectAll() {
	e.SetSelection(RangeOf(0, e.value.Len()))
}

// MoveCursor 移动光标（delta 为 rune 数，clamp 到文本范围）。
func (e *TextEditor) MoveCursor(delta int32) {
	c := e.value.Cursor() + delta
	if c < 0 {
		c = 0
	}
	if n := e.value.Len(); c > n {
		c = n
	}
	e.SetSelection(RangeAt(c))
}

// ── TextInputClient ─────────────────────────────────────────────────

// InsertText 插入文本（替换选区）。组合期间的插入会先提交组合串。
//
// 控制字符在单行模式下被丢弃（多行只允许 \n）：IME 偶尔会送 \r\n，
// 直接插进去会让"单行输入框里出现换行"这种诡异现象。
func (e *TextEditor) InsertText(s string) {
	if s == "" {
		return
	}
	s = e.filter(s)
	if s == "" {
		e.stats.Dropped++
		return
	}
	e.pushUndo()
	// 组合串未提交时先落定：否则把新字符插到组合区间里，IME 会再送一次
	// 组合串，表现为"打的字重复一遍"。
	if !e.value.Composing.IsEmpty() {
		e.value.Composing = TextRange{}
		e.stats.Commits++
	}
	runes := []rune(e.value.Text)
	sel := e.value.Selection.Clamp(int32(len(runes)))
	ins := []rune(s)
	out := make([]rune, 0, len(runes)+len(ins))
	out = append(out, runes[:sel.Start]...)
	out = append(out, ins...)
	out = append(out, runes[sel.End:]...)
	e.value.Text = string(out)
	c := sel.Start + int32(len(ins))
	e.value.Selection = RangeAt(c)
	e.stats.Inserts++
	e.changed()
}

// filter 过滤不可输入的字符。
func (e *TextEditor) filter(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch {
		case r == '\n' || r == '\r':
			if !e.multiline {
				continue
			}
			if r == '\r' {
				continue // 统一成 \n
			}
			b.WriteRune(r)
		case r == '\t':
			if !e.multiline {
				continue
			}
			b.WriteRune(r)
		case unicode.IsControl(r):
			continue
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

// DeleteBackward 删除光标前一个字符（或选区）。
func (e *TextEditor) DeleteBackward() { e.deleteSpan(-1) }

// DeleteForward 删除光标后一个字符（或选区）。
func (e *TextEditor) DeleteForward() { e.deleteSpan(1) }

// deleteSpan 实现前/后删除。delta = -1 向前，+1 向后。
func (e *TextEditor) deleteSpan(delta int32) {
	if !e.value.Selection.IsEmpty() {
		e.replaceSelection("")
		return
	}
	c := e.value.Cursor()
	n := e.value.Len()
	target := c + delta
	if target < 0 || target > n || target == c {
		return
	}
	e.pushUndo()
	runes := []rune(e.value.Text)
	out := make([]rune, 0, len(runes)-1)
	out = append(out, runes[:min32(c, target)]...)
	out = append(out, runes[min32(c, target)+1:]...)
	e.value.Text = string(out)
	e.value.Selection = RangeAt(min32(c, target))
	e.value.Composing = TextRange{}
	e.stats.Deletes++
	e.changed()
}

// replaceSelection 用 s 替换当前选区。
func (e *TextEditor) replaceSelection(s string) {
	e.pushUndo()
	runes := []rune(e.value.Text)
	sel := e.value.Selection.Clamp(int32(len(runes)))
	ins := []rune(s)
	out := make([]rune, 0, len(runes)+len(ins))
	out = append(out, runes[:sel.Start]...)
	out = append(out, ins...)
	out = append(out, runes[sel.End:]...)
	e.value.Text = string(out)
	c := sel.Start + int32(len(ins))
	e.value.Selection = RangeAt(c)
	e.value.Composing = TextRange{}
	e.stats.Deletes++
	e.changed()
}

// SetComposingRegion 更新组合串：用 text 替换 [start,end)。
//
// IME 每敲一个字母都会重送整个组合串（拼音 "zhong" 的中间态是 "z"、"zh"…），
// 所以语义是"替换上一次的组合区间"，不是"追加"。
func (e *TextEditor) SetComposingRegion(start, end int32, text string) {
	e.pushUndo()
	runes := []rune(e.value.Text)
	r := RangeOf(start, end).Clamp(int32(len(runes)))

	// 组合串要替换的是"上一次的组合区间"，与传入区间合并：
	// 有些 IME 只给当前词的范围，链式输入时旧组合串会残留。
	if old := e.value.Composing; !old.IsEmpty() && old.Clamp(int32(len(runes))) != r {
		if old.Start < r.Start {
			r.Start = old.Start
		}
		if old.End > r.End {
			r.End = old.End
		}
	}

	ins := []rune(text)
	out := make([]rune, 0, len(runes)+len(ins))
	out = append(out, runes[:r.Start]...)
	out = append(out, ins...)
	out = append(out, runes[r.End:]...)
	e.value.Text = string(out)

	c := r.Start + int32(len(ins))
	e.value.Selection = RangeAt(c)
	e.value.Composing = RangeOf(r.Start, c)
	e.changed()
}

// CommitComposing 落定组合串（清空组合区间，文本保留）。
func (e *TextEditor) CommitComposing() {
	if e.value.Composing.IsEmpty() {
		return
	}
	e.value.Composing = TextRange{}
	e.stats.Commits++
	e.changed()
}

// CommitComposingWith 用 s 替换组合区间并落定——IME 上屏走这条。
//
// 为什么不是 "CommitComposing + InsertText"：组合串（拼音 "zhong"）是
// **预编辑**，它只是临时占据文本位置；上屏时要用汉字"中"把它顶掉。
// 分成两步做会得到 "zhong中"，用户看到的是自己的拼音被留在输入框里。
func (e *TextEditor) CommitComposingWith(s string) {
	r := e.value.Composing.Clamp(e.value.Len())
	runes := []rune(e.value.Text)
	ins := []rune(s)
	out := make([]rune, 0, len(runes)-int(r.Len())+len(ins))
	out = append(out, runes[:r.Start]...)
	out = append(out, ins...)
	out = append(out, runes[r.End:]...)
	e.value.Text = string(out)
	c := r.Start + int32(len(ins))
	e.value.Selection = RangeAt(c)
	e.value.Composing = TextRange{}
	e.stats.Commits++
	e.changed()
}

// CancelComposing 取消组合串：把它从文本里删掉（Esc / IME 取消）。
func (e *TextEditor) CancelComposing() {
	r := e.value.Composing.Clamp(e.value.Len())
	if r.IsEmpty() {
		return
	}
	e.value.Composing = TextRange{}
	runes := []rune(e.value.Text)
	out := make([]rune, 0, len(runes))
	out = append(out, runes[:r.Start]...)
	out = append(out, runes[r.End:]...)
	e.value.Text = string(out)
	e.value.Selection = RangeAt(r.Start)
	e.changed()
}

// SelectionChanged 接收平台的选区变化。
//
// 平台会把"我们刚同步过去的选区"再回送一次，这类回声必须过滤：
// 不过滤会与本地编辑互相覆盖，表现为光标乱跳（尤其是输入法组合期间）。
func (e *TextEditor) SelectionChanged(r TextRange) {
	r = r.Clamp(e.value.Len())
	if r == e.value.Selection {
		return
	}
	e.value.Selection = r
	e.changed()
}

// PerformAction 处理语义动作。
//
// 单行模式下 ActionNewline 视为 ActionDone：单行输入框里按回车
// 应该是"提交"，不是"插入换行"（插不进去，也不该没反应）。
func (e *TextEditor) PerformAction(a Action) {
	e.stats.Actions++
	if a == ActionNewline {
		if e.multiline {
			e.InsertText("\n")
			return
		}
		a = ActionDone
	}
	if e.OnAction != nil {
		e.OnAction(a)
	}
}

// ── 小工具 ──────────────────────────────────────────────────────────

func min32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// RuneLen 返回 s 的 rune 数（供 host 做 UTF-16 换算的对照）。
func RuneLen(s string) int32 { return int32(utf8.RuneCountInString(s)) }
