package widgets

import (
	"strconv"
	"strings"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// ── InputNumber（M8 对齐 antd）──────────────────────────────────────
//
// antd：<InputNumber min max step precision onChange>
// 已对齐：min/max 钳制、step 步进（按钮 + 聚焦后 ↑/↓ 键）、右侧上下
// 步进器、禁用；非法输入回退上一个合法值。
// 未实现：precision 小数位格式化（按输入原样显示）、键盘长按连发。

// InputNumber 是数字输入框。
type InputNumber struct {
	controlBase
	Min, Max float32
	Step     float32 // 默认 1
	Value    float32
	Disabled bool
	Size     ui.Size
	Width    float32
	OnChange func(float32)

	editor *input.TextEditor
	keys   *input.TextKeyHandler
}

// NewInputNumber 创建数字输入框。
func NewInputNumber(min, max, value float32) *InputNumber {
	return &InputNumber{Min: min, Max: max, Value: value, Step: 1,
		Size: ui.SizeMedium, Width: 120}
}

func (w *InputNumber) clamp(v float32) float32 {
	if v < w.Min {
		return w.Min
	}
	if v > w.Max {
		return w.Max
	}
	return v
}

func (w *InputNumber) stepBy(dir float32) {
	d := w.Step
	if d <= 0 {
		d = 1
	}
	next := w.clamp(w.Value + dir*d)
	if next == w.Value {
		return
	}
	w.Value = next
	w.MarkDirty()
	if w.OnChange != nil {
		w.OnChange(next)
	}
}

// Build 实现 ui.Widget。
func (w *InputNumber) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	if w.editor == nil {
		w.editor = input.NewTextEditor(false)
		w.keys = input.NewTextKeyHandler(w.editor)
	}
	cur := fmtFloat(w.Value)
	if w.editor.Value().Text != cur {
		w.editor.SetValue(input.EditingValue{Text: cur,
			Selection: input.RangeAt(input.RuneLen(cur))})
	}
	w.editor.OnChanged = func(v input.EditingValue) {
		s := strings.TrimSpace(v.Text)
		if f, err := strconv.ParseFloat(s, 32); err == nil {
			nv := w.clamp(float32(f))
			if nv != w.Value {
				w.Value = nv
				if w.OnChange != nil {
					w.OnChange(nv)
				}
			}
			return
		}
		// 非法输入：回退上一个合法值（光标到末尾）。
		w.editor.SetValue(input.EditingValue{Text: fmtFloat(w.Value),
			Selection: input.RangeAt(input.RuneLen(fmtFloat(w.Value)))})
	}
	e := w.editor

	m := ui.MetricsFor(th, w.Size)
	height := m.Height
	width := w.Width
	if width <= 0 {
		width = 120
	}
	border, bg := th.Color.Border, th.Color.Surface
	if w.hover && !w.Disabled {
		border = th.Color.PrimaryHover
	}
	if w.focus != nil && w.focus.Focused() && !w.Disabled {
		border = th.Color.Primary
	}
	const stepW = float32(18)
	box := scene.Box(scene.BoxSpec{Fill: scene.Solid(bg), Border: scene.Stroke(border, 1), Radius: th.Radius.Base}).
		SetSize2(width, height)

	font := text.DefaultFont(m.Font)
	layout := text.LayoutText(e.Value().Text, font, text.Options{MaxWidth: width - 24 - stepW})
	t := scene.Text(scene.TextSpec{Src: e.Value().Text, Layout: layout, Color: th.Color.Text})
	t.SetMargin(geom.InLTRB(11, (height-layout.Size.H)/2, 0, 0))
	box.Add(t)
	box.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "input-number", Cursor: scene.CursorText, Disabled: w.Disabled})

	// 右侧步进器：上/下两个半区。
	for i, name := range []string{"up", "down"} {
		ic := (&Icon{Name: name, Size: 10, Color: th.Color.TextMuted}).Build(ctx)
		ic.SetPositionKeepSize(true)
		ic.SetStackPosition(geom.Pt(width-stepW+3, float32(i)*(height/2)+(height/2-10)/2))
		ic.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "input-number-" + name, Cursor: scene.CursorPointer})
		dir := float32(1)
		if i == 1 {
			dir = -1
		}
		if ctx.Input != nil {
			ctx.Input.Set(ic, input.Handler{
				Pointer: func(ev *input.PointerEvent) bool {
					if w.Disabled || !ev.Down() {
						return false
					}
					w.stepBy(dir)
					return true
				},
			})
		}
		box.Add(ic)
	}
	// 分隔线
	sep := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetSize2(1, height-8)
	sep.SetPositionKeepSize(true).SetStackPosition(geom.Pt(width-stepW-1, 4))
	box.Add(sep)

	h := input.Handler{
		Enter: func(*input.PointerEvent) { w.hover = true; w.MarkDirty() },
		Leave: func(*input.PointerEvent) {
			if w.hover {
				w.hover = false
				w.MarkDirty()
			}
		},
		Key: func(ev *input.KeyEvent) bool {
			if w.Disabled {
				return false
			}
			if ev.Kind != input.KeyUp {
				switch ev.Code {
				case input.CodeUp:
					w.stepBy(1)
					return true
				case input.CodeDown:
					w.stepBy(-1)
					return true
				}
			}
			if w.keys.Handle(ev) {
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
			w.MarkDirty()
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
	w.attach(ctx, box, h, !w.Disabled)
	return box
}

// fmtFloat 数字文本（整数不带小数点）。
func fmtFloat(v float32) string {
	if v == float32(int64(v)) {
		return strconv.FormatInt(int64(v), 10)
	}
	return strconv.FormatFloat(float64(v), 'f', -1, 32)
}
