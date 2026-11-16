package widgets

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── Input（M8 对齐 antd v5）─────────────────────────────────────────
//
// antd：<Input variant="outlined|filled|borderless" status="error|warning"
//
//	size prefix suffix addonBefore addonAfter count allowClear>
//
// 视觉对齐：outlined 边框 #D9D9D9、hover/focus 转主色 + 主色光晕；
// filled 灰底无边框；borderless 纯文本；status 用 error/warning 边框色；
// 控件高 24/32/40；禁用底 #FAFAFA、文字 #BFBFBF。
// 未实现（如实标注）：status 的反馈图标、textarea 的 showCount 行、autosize。

// InputVariant 是输入框形态（antd Input.variant）。
type InputVariant uint8

const (
	// InputOutlined 默认描边形态。
	InputOutlined InputVariant = iota
	// InputFilled 灰底形态。
	InputFilled
	// InputBorderless 无边框形态。
	InputBorderless
)

// InputStatus 是校验状态（antd Input.status）。
type InputStatus uint8

const (
	// StatusOK 常态。
	StatusOK InputStatus = iota
	// InputStatusError 错误（红边）。
	InputStatusError
	// InputStatusWarning 警告（黄边）。
	InputStatusWarning
)

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
	Variant     InputVariant
	Status      InputStatus
	Size        ui.Size
	Prefix      ui.Widget // 框内左侧
	Suffix      ui.Widget // 框内右侧（allowClear 时清除图标在其后）
	AllowClear  bool
	AddonBefore ui.Widget // 框外左侧灰块
	AddonAfter  ui.Widget // 框外右侧灰块
	Count       bool      // 右下角显示字数
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

// statusBorder 折算边框/底色（antd 状态语义）。
func (w *Input) statusBorder(th *theme.Data, hover bool, focused bool) (geom.Color, geom.Color, float32) {
	c := th.Color
	border, bg := c.Border, c.Surface
	switch w.Variant {
	case InputFilled:
		border, bg = geom.Color{}, c.SurfaceSunken // #FAFAFA
		if hover {
			bg = c.SurfaceActive // #F0F0F0
		}
		if focused {
			bg = c.Surface
		}
	case InputBorderless:
		return geom.Color{}, geom.Color{}, 0
	}
	if w.Status == InputStatusError {
		border = c.Danger
		if hover {
			border = c.DangerHover
		}
	} else if w.Status == InputStatusWarning {
		border = c.Warning
	}
	if hover && w.Status == StatusOK {
		border = c.PrimaryHover // antd hover #4096FF
	}
	if focused && w.Status == StatusOK {
		border = c.Primary
	}
	return border, bg, 1
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

	m := ui.MetricsFor(th, w.Size)
	height := w.Height
	if height <= 0 {
		height = m.Height
	}
	width := w.Width
	if width <= 0 {
		width = 200
	}
	focused := w.focus != nil && w.focus.Focused()
	hover := w.hover
	border, bg, bw := w.statusBorder(th, hover && !w.Disabled, focused)

	// 框内两侧的装饰宽度（prefix/suffix/clear）。
	sideW := float32(0)
	if w.Prefix != nil {
		sideW += 18
	}
	showClear := w.AllowClear && w.Value != "" && !w.Disabled
	sideSuffix := float32(0)
	if w.Suffix != nil {
		sideSuffix += 18
	}
	if showClear {
		sideSuffix += 16
	}
	sideW += sideSuffix

	innerW := width - 22 - sideW // 22 = 左右基础内边距
	if innerW < 24 {
		innerW = 24
	}
	spec := scene.BoxSpec{Fill: scene.Solid(bg), Radius: st2radius(w.Size, th)}
	if bw > 0 {
		spec.Border = scene.Stroke(border, bw)
	}
	if focused && w.Variant == InputOutlined && w.Status == StatusOK {
		// antd 焦点光晕：主色 10% 的 2px 环，用零偏移阴影近似。
		spec.Shadow = &scene.ShadowSpec{Dx: 0, Dy: 0, Blur: 3, Color: geom.RGBAa(22, 119, 255, 26)}
	}
	if w.Disabled {
		spec.Shadow = nil
	}
	n := scene.Box(spec).SetSize2(width, height)
	n.SetOpacity(1)

	v := e.Value()
	src := v.Text
	color := th.Color.Text
	if w.Disabled {
		color = th.Color.BorderStrong
	}
	if src == "" {
		src = w.Placeholder
		color = th.Color.TextFaint
	}
	font := text.DefaultFont(m.Font)
	layout := text.LayoutText(src, font, text.Options{
		MaxWidth: innerW, MaxLines: w.MaxLines, Ellipsis: "…",
	})
	t := scene.Text(scene.TextSpec{
		Src: src, Layout: layout, Color: color,
		SelectionColor: th.Color.Primary,
		SelStart:       v.Selection.Start, SelEnd: v.Selection.End,
		ComposingStart: v.Composing.Start, ComposingEnd: v.Composing.End,
	})
	textX := float32(11)
	if w.Prefix != nil {
		textX += 18
		pn := w.Prefix.Build(ctx)
		pn.SetPositionKeepSize(true)
		pn.SetStackPosition(geom.Pt(11, (height-pn.Bounds().H())/2))
		n.Add(pn)
	}
	t.SetMargin(geom.InLTRB(textX, (height-layout.Size.H)/2, 0, 0))
	n.Add(t)

	sx := width - 11 - sideSuffix
	if w.Suffix != nil {
		sn := w.Suffix.Build(ctx)
		sn.SetPositionKeepSize(true)
		sn.SetStackPosition(geom.Pt(sx, (height-sn.Bounds().H())/2))
		n.Add(sn)
		sx += 18
	}
	if showClear {
		clear := (&Icon{Name: "close-circle", Size: 14,
			Color: th.Color.TextFaint}).Build(ctx)
		clear.SetPositionKeepSize(true)
		clear.SetStackPosition(geom.Pt(width - 11 - 16, (height-14)/2))
		// 清除图标自己吃点击：挂在同一 Box 上的独立命中区。
		clear.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "input-clear", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			ctx.Input.Set(clear, input.Handler{
				Pointer: func(ev *input.PointerEvent) bool {
					if !ev.Down() {
						return false
					}
					e.SetValue(input.EditingValue{})
					w.Value = ""
					w.MarkDirty()
					if w.OnChanged != nil {
						w.OnChanged("")
					}
					return true
				},
			})
		}
		n.Add(clear)
	}
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "input", Cursor: scene.CursorText})

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
			w.MarkDirty() // 焦点框变化需要重绘
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

	// 外层组合：addon 块 + 输入框 (+ count 行)。
	var root *scene.Node
	if w.AddonBefore != nil || w.AddonAfter != nil {
		row := scene.Flex(scene.Horizontal, 0).SetCrossAlign(scene.CrossStretch)
		if w.AddonBefore != nil {
			row.Add(addonBox(ctx, w.AddonBefore, height))
		}
		row.Add(n)
		if w.AddonAfter != nil {
			row.Add(addonBox(ctx, w.AddonAfter, height))
		}
		root = row
	} else {
		root = n
	}
	if w.Count {
		col := scene.Flex(scene.Vertical, 2)
		col.Add(root)
		col.Add((&Label{Text: fmt.Sprintf("%d", runeLen(e.Value().Text)), Size: th.Font.SizeXS,
			Color: th.Color.TextFaint}).Build(ctx))
		root = col
	}
	w.attach(ctx, root, h, !w.Disabled)
	return root
}

// addonBox 是 addon 前后缀的灰块容器。
func addonBox(ctx *ui.BuildContext, w ui.Widget, height float32) *scene.Node {
	th := ctx.ThemeOr()
	box := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.SurfaceSunken),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	box.SetPadding(geom.InLTRB(11, 0, 11, 0))
	box.SetHeight(height)
	box.Add(w.Build(ctx))
	box.SetOpacity(1)
	// addon 文字置灰
	return box
}

// st2radius 折算控件圆角（antd：middle/default 6、small 4）。
func st2radius(s ui.Size, th *theme.Data) float32 {
	if s == ui.SizeSmall {
		return th.Radius.SM
	}
	return th.Radius.Base
}

// runeLen 是 rune 计数（antd count 语义）。
func runeLen(s string) int { return len([]rune(s)) }

var _ = theme.Default
