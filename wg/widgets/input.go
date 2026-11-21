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
	// HitLabel 自定义命中标签（空 = "input"）。嵌入复合组件（如分页跳页框）
	// 时设成非交互标签，避免被自动化审计当独立输入框枚举。
	HitLabel  string
	MaxLines  int
	OnChanged func(string)
	OnSubmit  func(string)

	editor *input.TextEditor
	keys   *input.TextKeyHandler

	// 呈现缓存：编辑时只更新 Text 节点与光标节点（SetText/SetStackPosition），
	// **不整页重建**——每次按键全树 Build 是输入延迟/连发卡顿的主因。
	textNode  *scene.Node
	caretNode *scene.Node
	// countNode 是右下角字数行（单行原地刷新时一并更新）。
	countNode  *scene.Node
	countFont  text.Font
	countColor geom.Color
	countLH    float32
	countMaxW  float32
	refresh    func()
	// scrollX 是单行输入的横向滚动量（保证光标可见）。
	scrollX float32
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

// CaretNode 返回当前插入符节点（诊断/测试用；未聚焦时为 nil）。
func (w *Input) CaretNode() *scene.Node { return w.caretNode }

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
	// antd disabled：边框仍 colorBorder，底转 colorBgContainerDisabled，
	// 无 hover/focus 反馈（光晕也不画）。
	if w.Disabled {
		switch w.Variant {
		case InputFilled:
			return geom.Color{}, c.SurfaceDisabled, 1
		case InputBorderless:
			return geom.Color{}, geom.Color{}, 0
		}
		return c.Border, c.SurfaceDisabled, 1
	}
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
		if hover {
			border = c.WarningHover // antd colorWarningHover #FFC53D
		}
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
	// 本帧结束前 refresh 不可指向旧节点（重建会替换整棵树）。
	w.refresh = nil
	w.textNode, w.caretNode, w.countNode = nil, nil, nil
	if w.Value != e.Value().Text {
		v := input.EditingValue{Text: w.Value, Selection: input.RangeAt(input.RuneLen(w.Value))}
		e.SetValue(v)
	}
	e.OnChanged = func(v input.EditingValue) {
		w.Value = v.Text
		if w.OnChanged != nil {
			w.OnChanged(v.Text)
		}
		// 单行（含 Count 行）：原地更新文本/光标/字数，不整页重建。
		if w.refresh != nil && !w.Multiline {
			w.refresh()
			return
		}
		w.MarkDirty()
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
	// status 反馈图标（antd status 图标：error 红 close-circle /
	// warning 黄 warning-circle，右侧覆盖在 suffix 区前）。
	statusIcon := ""
	statusIconColor := geom.Color{}
	if w.Status == InputStatusError {
		statusIcon, statusIconColor = "close-circle", th.Color.Danger
	} else if w.Status == InputStatusWarning {
		statusIcon, statusIconColor = "warning-circle", th.Color.Warning
	}
	if statusIcon != "" {
		sideSuffix += 18
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
		color = th.Color.TextPlaceholder // antd colorTextPlaceholder #BFBFBF
	}
	font := text.DefaultFont(m.Font)
	lineH := text.Measure("Ag", font).H
	caretIdx := int(v.Selection.End)
	caretX := textWidth(v.Text, caretIdx, font)
	if !w.Multiline {
		// 单行横向滚动：保证光标在可视区内。
		if caretX-w.scrollX > innerW-2 {
			w.scrollX = caretX - (innerW - 2)
		}
		if caretX-w.scrollX < 0 {
			w.scrollX = 0
		}
		if w.scrollX < 0 {
			w.scrollX = 0
		}
	} else {
		w.scrollX = 0
	}
	layout := text.LayoutText(src, font, text.Options{
		MaxWidth: innerW, MaxLines: w.MaxLines, Ellipsis: "…",
	})
	t := scene.Text(scene.TextSpec{
		Src: src, Layout: layout, Color: color,
		SelectionColor: th.Color.Primary,
		SelStart:       v.Selection.Start, SelEnd: v.Selection.End,
		ComposingStart: v.Composing.Start, ComposingEnd: v.Composing.End,
		// 重排元数据必须与本次排版一致：缺 LayoutFont 时场景会用零值字体
		// 重排（字号变大、基线偏移），真窗表现为"重建后文字右下偏移"。
		MaxWidth: innerW, LayoutFont: font, LayoutW: innerW,
	})
	textX := float32(11)
	if w.Prefix != nil {
		textX += 18
		pn := w.Prefix.Build(ctx)
		pn.SetPositionKeepSize(true)
		pn.SetStackPosition(geom.Pt(11, (height-pn.Bounds().H())/2))
		n.Add(pn)
	}
	t.SetMargin(geom.InLTRB(textX-w.scrollX, (height-layout.Size.H)/2, 0, 0))
	n.Add(t)

	sx := width - 11 - sideSuffix
	if statusIcon != "" {
		si := (&Icon{Name: statusIcon, Size: 14, Color: statusIconColor}).Build(ctx)
		si.SetPositionKeepSize(true)
		si.SetStackPosition(geom.Pt(sx, (height-14)/2))
		n.Add(si)
		sx += 18
	}
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
		clear.SetStackPosition(geom.Pt(width-11-16, (height-14)/2))
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
	hitLabel := w.HitLabel
	if hitLabel == "" {
		hitLabel = "input"
	}
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: hitLabel, Cursor: scene.CursorText, Disabled: w.Disabled})

	// 光标（自绘）：聚焦时显示，位置由选区末端决定。
	caretLocal := geom.RectXYWH(textX+caretX-w.scrollX, (height-lineH)/2, 2, lineH)
	w.textNode = t
	if focused {
		caret := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Text)}).SetSize2(2, lineH)
		// Stack 定位默认 keepSize=false（走 Tight 约束，会被拉伸成整框）；
		// 光标必须固定 2×行高。
		caret.SetPositionKeepSize(true)
		caret.SetStackPosition(caretLocal.Min)
		n.Add(caret)
		w.caretNode = caret
	}
	// IME 锚点：宿主布局后换算成客户区坐标（组合窗/候选窗定位）。
	// 以 w 作为 owner：页面上其它未聚焦输入框的 nil 上报不会清掉这里。
	if ctx.ReportIMECaret != nil {
		if focused {
			ctx.ReportIMECaret(w, n, caretLocal)
		} else {
			ctx.ReportIMECaret(w, nil, geom.Rect{})
		}
	}
	// 插入符闪烁交给宿主：按 530ms 切换可见性；编辑时重报会重置相位。
	if ctx.ReportCaret != nil {
		if focused {
			ctx.ReportCaret(w, w.caretNode)
		} else {
			ctx.ReportCaret(w, nil)
		}
	}

	// 原地刷新（单行）：编辑后只更新文本/光标/字数/锚点并请求一帧。
	// 这是"打字/连发不卡"的关键——全树 Build 每键一次在复杂页面上是几十 ms。
	if !w.Multiline {
		opts := text.Options{MaxWidth: innerW, MaxLines: w.MaxLines, Ellipsis: "…"}
		req := ctx.RequestFrame
		w.refresh = func() {
			if w.textNode == nil {
				w.MarkDirty()
				return
			}
			vv := e.Value()
			src2, color2 := vv.Text, th.Color.Text
			if vv.Text == "" {
				src2, color2 = w.Placeholder, th.Color.TextPlaceholder
			}
			lay := text.LayoutText(src2, font, opts)
			// 只重绘不重排：父框尺寸固定，文本变宽/变窄不该触发整页布局。
			w.textNode.SetTextPaint(scene.TextSpec{
				Src: src2, Layout: lay, Color: color2,
				SelectionColor: th.Color.Primary,
				SelStart:       vv.Selection.Start, SelEnd: vv.Selection.End,
				ComposingStart: vv.Composing.Start, ComposingEnd: vv.Composing.End,
				MaxWidth: innerW, LayoutFont: font, LayoutW: innerW,
			})
			cx := textWidth(vv.Text, int(vv.Selection.End), font)
			if cx-w.scrollX > innerW-2 {
				w.scrollX = cx - (innerW - 2)
			}
			if cx-w.scrollX < 0 {
				w.scrollX = 0
			}
			ins := geom.InLTRB(textX-w.scrollX, (height-lay.Size.H)/2, 0, 0)
			w.textNode.SetMarginPaint(ins)
			org := n.ContentOrigin()
			w.textNode.SetOffsetPaint(geom.Pt(org.X+ins.Left, org.Y+ins.Top))
			// 文本变短时旧字形要用框底色覆盖；标父框（连带子树）重绘。
			n.MarkNeedsPaint()
			loc := geom.RectXYWH(textX+cx-w.scrollX, (height-lineH)/2, 2, lineH)
			if w.caretNode != nil {
				w.caretNode.SetStackPositionPaint(loc.Min)
			}
			if ctx.ReportIMECaret != nil {
				ctx.ReportIMECaret(w, n, loc)
			}
			// 编辑/移动光标：恢复可见并重置闪烁相位。
			if ctx.ReportCaret != nil && w.caretNode != nil {
				ctx.ReportCaret(w, w.caretNode)
			}
			// 字数行原地更新：只换文本并标父级重绘（旧数字变短也要覆盖）。
			if w.countNode != nil {
				cnt := fmt.Sprintf("%d", runeLen(vv.Text))
				cl := text.LayoutText(cnt, w.countFont, text.Options{LineHeight: w.countLH})
				w.countNode.SetTextPaint(scene.TextSpec{
					Src: cnt, Layout: cl, Color: w.countColor, LayoutFont: w.countFont,
					MaxWidth: w.countMaxW, LayoutW: w.countMaxW,
				})
				if p := w.countNode.Parent(); p != nil {
					p.MarkNeedsPaint()
				}
			}
			if req != nil {
				req()
			}
		}
	}

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
			// 点击定位光标：仅当 root==输入框本体（无 addon/字数行）时
			// ev.Pos 才相对输入框；多行暂不支持按行定位。
			if !w.Multiline && w.AddonBefore == nil && w.AddonAfter == nil && !w.Count {
				lx := ev.Pos.X - textX + w.scrollX
				e.SetSelection(input.RangeAt(int32(caretFromX(e.Value().Text, lx, font))))
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
		cnt := fmt.Sprintf("%d", runeLen(e.Value().Text))
		cf := text.DefaultFont(th.Font.SizeXS)
		cl := text.LayoutText(cnt, cf, text.Options{LineHeight: th.Font.LineHeight})
		// MaxWidth 钉住自然宽度：否则文本节点会按父级可用宽度撑满，
		// 字数行随页面版心变宽（原地刷新时脏区跟着放大）。
		cn := scene.Text(scene.TextSpec{
			Src: cnt, Layout: cl, Color: th.Color.TextFaint, LayoutFont: cf,
			MaxWidth: width, LayoutW: width,
		})
		col := scene.Flex(scene.Vertical, 2)
		col.Add(root)
		col.Add(cn)
		root = col
		// 字数行参与原地刷新（每键只改文本，不整页重建）。
		w.countNode, w.countFont, w.countColor = cn, cf, th.Color.TextFaint
		w.countLH, w.countMaxW = th.Font.LineHeight, width
	}
	// 处理器与焦点必须挂在输入框本体 n 上：分发器只把 Pointer 投给命中的
	// 最深节点，挂在 addon 的 row / count 的 col 上时点击输入框命中 n，
	// row/col 的处理器不会被调用（--audit 实测：点击不聚焦、键盘无输入）。
	// addon/字数行点击不聚焦，与 antd 一致。
	w.attach(ctx, n, h, !w.Disabled)
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

// textWidth 返回前 n 个 rune 的像素宽（前缀测量，输入框文本通常很短）。
func textWidth(s string, n int, font text.Font) float32 {
	rs := []rune(s)
	if n > len(rs) {
		n = len(rs)
	}
	if n <= 0 {
		return 0
	}
	return text.Measure(string(rs[:n]), font).W
}

// caretFromX 把点击的局部 x 换算成最近的 rune 光标位。
func caretFromX(s string, x float32, font text.Font) int {
	rs := []rune(s)
	prev := float32(0)
	for i := 1; i <= len(rs); i++ {
		cur := text.Measure(string(rs[:i]), font).W
		if x < (prev+cur)/2 {
			return i - 1
		}
		prev = cur
	}
	return len(rs)
}

var _ = theme.Default
