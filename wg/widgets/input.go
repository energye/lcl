package widgets

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// ── Input ──────────────────────────────────────────────────────────

// Input 是单行/多行文本输入框。
//
// 编辑状态放在 input.TextEditor（可跨 Build 存活），控件只负责呈现与
// 把 Key/Focus/IME 接到正确的模型上。重建时不会丢光标、选区、撤销栈。
type Input struct {
	controlBase
	Placeholder string
	Value       string
	Multiline   bool
	Disabled    bool
	Width       float32
	Height      float32
	// MaxLines 只对多行生效。
	MaxLines  int
	OnChanged func(string)
	OnSubmit  func(string)

	editor *input.TextEditor
	keys   *input.TextKeyHandler
}

// NewInput 创建输入框。
func NewInput(placeholder string) *Input {
	return &Input{Placeholder: placeholder, Width: 200, Height: ui.MetricsFor(nil, ui.SizeMedium).Height}
}

// Editor 返回编辑模型（诊断/测试用）。
func (w *Input) Editor() *input.TextEditor {
	w.ensureEditor()
	return w.editor
}

// FocusNode 返回持久焦点节点（焦点管理/诊断用）。
func (w *Input) FocusNode() *input.FocusNode {
	if w.focus == nil {
		w.focus = input.NewFocusNode(true)
	}
	return w.focus
}

func (w *Input) ensureEditor() *input.TextEditor {
	if w.editor == nil {
		w.editor = input.NewTextEditor(w.Multiline)
		w.keys = input.NewTextKeyHandler(w.editor)
	} else if w.editor.Multiline() != w.Multiline {
		// 单行 ↔ 多行是结构变化，换模型但保留文本。
		old := w.editor.Value().Text
		w.editor = input.NewTextEditor(w.Multiline)
		w.keys = input.NewTextKeyHandler(w.editor)
		w.editor.SetValue(input.EditingValue{Text: old, Selection: input.RangeAt(input.RuneLen(old))})
	}
	return w.editor
}

// Build 实现 ui.Widget。
func (w *Input) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	e := w.ensureEditor()
	if w.Value != e.Value().Text {
		v := input.EditingValue{Text: w.Value, Selection: input.RangeAt(input.RuneLen(w.Value))}
		e.SetValue(v)
	}
	e.OnChanged = func(v input.EditingValue) {
		w.Value = v.Text
		w.MarkDirty()
		if w.OnChanged != nil {
			w.OnChanged(v.Text)
		}
	}
	e.OnAction = func(a input.Action) {
		if a == input.ActionDone && w.OnSubmit != nil {
			w.OnSubmit(w.Value)
		}
	}
	e.CursorRect = func(input.EditingValue) geom.Rect {
		return geom.RectXYWH(0, 0, w.Width, w.Height)
	}

	st := ui.StyleOf(th, ui.VariantOutline, ui.SizeMedium, w.states(w.Disabled))
	m := ui.MetricsFor(th, ui.SizeMedium)
	height := w.Height
	if height <= 0 {
		height = m.Height
	}
	width := w.Width
	if width <= 0 {
		width = 200
	}
	spec := scene.BoxSpec{Fill: scene.Solid(th.Color.Surface), Radius: st.Radius}
	spec.Border = scene.Stroke(st.Border, 1)
	n := scene.Box(spec).SetSize2(width, height)
	n.SetOpacity(st.Opacity)

	v := e.Value()
	src := v.Text
	color := st.Foreground
	if src == "" {
		src = w.Placeholder
		color = th.Color.TextFaint
	}
	font := text.DefaultFont(m.Font)
	layout := text.LayoutText(src, font, text.Options{
		MaxWidth: width - m.Padding.Min.X - m.Padding.Max.X,
		MaxLines: w.MaxLines, Ellipsis: "…",
	})
	t := scene.Text(scene.TextSpec{
		Src: src, Layout: layout, Color: color,
		SelectionColor: th.Color.Primary,
		SelStart:       v.Selection.Start, SelEnd: v.Selection.End,
		ComposingStart: v.Composing.Start, ComposingEnd: v.Composing.End,
	})
	t.SetMargin(geom.InLTRB(m.Padding.Min.X, (height-layout.Size.H)/2, 0, 0))
	n.Add(t)
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "input", Cursor: scene.CursorText})

	h := input.Handler{
		Key: func(ev *input.KeyEvent) bool {
			if w.Disabled {
				return false
			}
			if w.keys.Handle(ev) {
				// Tab 在编辑模型里是 ActionNext，但它不认识焦点管理器。
				// 组件必须显式转发，否则输入框里按 Tab 光标不动（最常见的"焦点陷阱"）。
				if ev.Code == input.CodeTab && ctx.Focus != nil {
					ctx.Focus.HandleKey(ev)
				}
				return true
			}
			return ctx.Focus != nil && ctx.Focus.HandleKey(ev)
		},
		Focus: func(ev *input.FocusEvent) bool {
			if ev.Kind == input.FocusIn {
				if ctx.SetTextEditor != nil {
					ctx.SetTextEditor(e)
				}
			} else if ctx.SetTextEditor != nil {
				ctx.SetTextEditor(nil)
			}
			return true
		},
		Pointer: func(ev *input.PointerEvent) bool {
			if w.Disabled || !ev.Down() {
				return false
			}
			if ctx.Focus != nil && w.focus != nil {
				ctx.Focus.Focus(w.focus)
			}
			return true
		},
	}
	w.attach(ctx, n, h, !w.Disabled)
	return n
}

// ── Progress ───────────────────────────────────────────────────────

// Progress 是进度条。Value 取值 0…1，越界自动夹住。
type Progress struct {
	Value  float32
	Width  float32
	Height float32
}

// Build 实现 ui.Widget。
func (p Progress) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	w, h := p.Width, p.Height
	if w <= 0 {
		w = 160
	}
	if h <= 0 {
		h = 6
	}
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken), Radius: h / 2}).SetSize2(w, h)
	v := p.Value
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	if v > 0 {
		fill := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary), Radius: h / 2}).
			SetSize2(w*v, h)
		n.Add(fill)
	}
	return n
}

// ── Slider ─────────────────────────────────────────────────────────

// Slider 是拖动滑块。坐标换算只用事件自带的节点局部坐标，
// 不自己反算矩阵——那会和命中用的矩阵分家。
type Slider struct {
	controlBase
	Min, Max float32
	Value    float32
	Width    float32
	Disabled bool
	OnChange func(float32)
}

// NewSlider 创建滑块。
func NewSlider(min, max, value float32) *Slider {
	return &Slider{Min: min, Max: max, Value: value, Width: 160}
}

// Build 实现 ui.Widget。
func (b *Slider) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	w := b.Width
	if w <= 0 {
		w = 160
	}
	const (
		trackH = float32(4)
		knob   = float32(14)
	)
	track := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border), Radius: trackH / 2}).
		SetSize2(w, trackH)
	track.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "slider", Cursor: scene.CursorPointer})

	lo, hi := b.Min, b.Max
	if hi <= lo {
		hi = lo + 1
	}
	v := b.Value
	if v < lo {
		v = lo
	}
	if v > hi {
		v = hi
	}
	ratio := (v - lo) / (hi - lo)
	fill := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary), Radius: trackH / 2}).
		SetSize2(w*ratio, trackH)
	track.Add(fill)
	knobX := w*ratio - knob/2
	if knobX < 0 {
		knobX = 0
	}
	if knobX > w-knob {
		knobX = w - knob
	}
	k := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary), Radius: knob / 2}).SetSize2(knob, knob)
	k.SetPositionKeepSize(true).SetStackPosition(geom.Pt(knobX, (trackH-knob)/2))
	track.Add(k)

	apply := func(pos geom.Point) {
		if b.Disabled {
			return
		}
		r := pos.X / w
		if r < 0 {
			r = 0
		}
		if r > 1 {
			r = 1
		}
		next := lo + r*(hi-lo)
		if next == b.Value {
			return
		}
		b.Value = next
		b.MarkDirty()
		if b.OnChange != nil {
			b.OnChange(next)
		}
	}
	g := &ui.GestureDetector{
		OnPanStart: func(pos geom.Point) { apply(pos) },
		OnPan:      func(pos, _ geom.Point) { apply(pos) },
	}
	b.tap = g
	h := g.Handler()
	h.Key = func(e *input.KeyEvent) bool {
		if e.Kind == input.KeyUp || b.Disabled {
			return false
		}
		step := (hi - lo) / 20
		switch e.Code {
		case input.CodeLeft:
			apply(geom.Pt((b.Value-step-lo)/(hi-lo)*w, 0))
			return true
		case input.CodeRight:
			apply(geom.Pt((b.Value+step-lo)/(hi-lo)*w, 0))
			return true
		}
		return ctx.Focus != nil && ctx.Focus.HandleKey(e)
	}
	b.attach(ctx, track, h, !b.Disabled)
	return track
}
