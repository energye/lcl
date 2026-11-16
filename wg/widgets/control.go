package widgets

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── 控件公共基座 ───────────────────────────────────────────────────

// controlBase 管理控件的焦点节点与交互状态。
//
// 焦点节点必须跨 Build 存活：每次重建都新建 FocusNode 会让 Tab 链
// 长度随着重建次数增长，而且焦点会莫名其妙丢失。
type controlBase struct {
	ui.Base
	node    *scene.Node
	focus   *input.FocusNode
	hover   bool
	pressed bool
	tap     *ui.GestureDetector
}

// Node 返回最近一次构建出的场景节点（诊断/自动化点击用）。
func (b *controlBase) Node() *scene.Node { return b.node }

func (b *controlBase) states(disabled bool) ui.StateSet {
	var s ui.StateSet
	if disabled {
		s |= ui.StateDisabled
	}
	if b.hover {
		s |= ui.StateHover
	}
	if b.pressed {
		s |= ui.StateActive
	}
	if b.focus != nil && b.focus.Focused() {
		s |= ui.StateFocus
	}
	return s
}

// attach 注册事件处理器并把持久焦点节点指向新构建的节点。
//
// 重建后 Dispatcher 的焦点仍指向**旧**节点，必须手动重定向：
// FocusManager.SetCurrent 看到"还是同一个 FocusNode"会直接返回，
// 不补这一步，输入框重建一次后就再也收不到按键。
func (b *controlBase) attach(ctx *ui.BuildContext, node *scene.Node, h input.Handler, canFocus bool) {
	b.node = node
	if ctx.Input != nil {
		ctx.Input.Set(node, h)
	}
	if !canFocus {
		return
	}
	if b.focus == nil {
		b.focus = input.NewFocusNode(true)
	}
	old := b.focus.Node
	b.focus.Node = node
	if ctx.Focus == nil {
		return
	}
	root := ctx.Focus.Root()
	if root == nil {
		root = input.NewFocusNode(false)
		ctx.Focus.SetRoot(root)
	}
	if b.focus.Parent == nil {
		root.AddChild(b.focus)
	}
	if b.focus.Focused() && old != node && ctx.Input != nil {
		ctx.Input.SetFocus(node)
	}
}

// tapHandler 生成"点一下就回调"的手势处理器。
func (b *controlBase) tapHandler(ctx *ui.BuildContext, disabled bool, onTap func()) input.Handler {
	g := &ui.GestureDetector{}
	if !disabled && onTap != nil {
		g.OnTap = func(geom.Point) { onTap() }
	}
	h := g.Handler()
	h.Enter = func(*input.PointerEvent) {
		if !disabled {
			b.hover = true
			b.MarkDirty()
		}
	}
	h.Leave = func(*input.PointerEvent) {
		if b.hover || b.pressed {
			b.hover, b.pressed = false, false
			b.MarkDirty()
		}
	}
	// GestureDetector 每次按下都会替换成员，闭包需要长期存活；
	// 存在控件上而不是局部变量，避免 GC 后回调消失。
	b.tap = g
	return h
}

// ── Button ─────────────────────────────────────────────────────────

// Button 是按钮。点击与 Enter/Space 走同一个 OnClick。
type Button struct {
	controlBase
	Text     string
	Variant  ui.Variant
	Size     ui.Size
	Disabled bool
	OnClick  func()
}

// NewButton 创建按钮。
func NewButton(text string) *Button { return &Button{Text: text, Size: ui.SizeMedium} }

// Build 实现 ui.Widget。
func (b *Button) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	st := ui.StyleOf(th, b.Variant, b.Size, b.states(b.Disabled))
	m := ui.MetricsFor(th, b.Size)
	font := text.DefaultFont(m.Font)
	layout := text.LayoutText(b.Text, font, text.Options{})

	w := layout.Size.W + m.Padding.Min.X + m.Padding.Max.X
	if w < m.Height {
		w = m.Height
	}
	spec := scene.BoxSpec{Fill: scene.Solid(st.Background), Radius: st.Radius}
	if st.BorderWidth > 0 {
		spec.Border = scene.Stroke(st.Border, st.BorderWidth)
	}
	n := scene.Box(spec).SetSize2(w, m.Height)
	n.SetOpacity(st.Opacity)
	n.SetHit(scene.HitSpec{
		Behavior: scene.HitOpaque, Label: "button", Cursor: scene.CursorPointer,
	})
	label := scene.Text(scene.TextSpec{Src: b.Text, Layout: layout, Color: st.Foreground})
	label.SetMargin(geom.InLTRB(m.Padding.Min.X, (m.Height-layout.Size.H)/2, 0, 0))
	n.Add(label)

	activate := func() {
		if !b.Disabled && b.OnClick != nil {
			b.OnClick()
		}
	}
	h := b.tapHandler(ctx, b.Disabled, activate)
	h.Key = func(e *input.KeyEvent) bool {
		if e.Kind == input.KeyUp || b.Disabled {
			return false
		}
		if e.Code == input.CodeEnter || e.Code == input.CodeSpace {
			activate()
			return true
		}
		if ctx.Focus != nil {
			return ctx.Focus.HandleKey(e)
		}
		return false
	}
	b.attach(ctx, n, h, !b.Disabled)
	return n
}

// ── Checkbox / Radio / Switch ──────────────────────────────────────

// Checkbox 是复选框。
type Checkbox struct {
	controlBase
	Checked  bool
	Disabled bool
	Label    string
	OnChange func(bool)
}

// NewCheckbox 创建复选框。
func NewCheckbox(label string, checked bool) *Checkbox {
	return &Checkbox{Label: label, Checked: checked}
}

// Build 实现 ui.Widget。
func (b *Checkbox) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	st := ui.StyleOf(th, ui.VariantOutline, ui.SizeSmall, b.states(b.Disabled))
	box := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(st.Border, 1),
		Radius: th.Radius.SM,
	}).SetSize2(16, 16)
	if b.Checked {
		mark := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary), Radius: 2}).SetSize2(8, 8)
		mark.SetPositionKeepSize(true).SetStackPosition(geom.Pt(4, 4))
		box.Add(mark)
	}
	box.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "checkbox", Cursor: scene.CursorPointer})
	box.SetOpacity(st.Opacity)

	toggle := func() {
		if b.Disabled {
			return
		}
		b.Checked = !b.Checked
		b.MarkDirty()
		if b.OnChange != nil {
			b.OnChange(b.Checked)
		}
	}
	h := b.tapHandler(ctx, b.Disabled, toggle)
	h.Key = func(e *input.KeyEvent) bool {
		if e.Code == input.CodeSpace && e.Kind != input.KeyUp && !b.Disabled {
			toggle()
			return true
		}
		return ctx.Focus != nil && ctx.Focus.HandleKey(e)
	}
	b.attach(ctx, box, h, !b.Disabled)
	return box
}

// Radio 是单选钮（同一组的互斥由调用方的 OnChange 负责）。
type Radio struct {
	controlBase
	Checked  bool
	Disabled bool
	Label    string
	OnChange func(bool)
}

// NewRadio 创建单选钮。
func NewRadio(label string, checked bool) *Radio {
	return &Radio{Label: label, Checked: checked}
}

// Build 实现 ui.Widget。
func (b *Radio) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	st := ui.StyleOf(th, ui.VariantOutline, ui.SizeSmall, b.states(b.Disabled))
	n := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(st.Border, 1),
		Radius: 8,
	}).SetSize2(16, 16)
	if b.Checked {
		dot := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary), Radius: 4}).SetSize2(8, 8)
		dot.SetPositionKeepSize(true).SetStackPosition(geom.Pt(4, 4))
		n.Add(dot)
	}
	n.SetOpacity(st.Opacity)
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "radio", Cursor: scene.CursorPointer})

	selectIt := func() {
		if b.Disabled || b.Checked {
			return
		}
		b.Checked = true
		b.MarkDirty()
		if b.OnChange != nil {
			b.OnChange(true)
		}
	}
	h := b.tapHandler(ctx, b.Disabled, selectIt)
	b.attach(ctx, n, h, !b.Disabled)
	return n
}

// Switch 是开关。
type Switch struct {
	controlBase
	On       bool
	Disabled bool
	OnChange func(bool)
}

// NewSwitch 创建开关。
func NewSwitch(on bool) *Switch { return &Switch{On: on} }

// Build 实现 ui.Widget。
func (b *Switch) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	st := ui.StyleOf(th, ui.VariantSoft, ui.SizeSmall, b.states(b.Disabled))
	const (
		trackW = float32(36)
		trackH = float32(20)
		knob   = float32(16)
	)
	track := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(bgOf(th, st, b.On)),
		Radius: trackH / 2,
	}).SetSize2(trackW, trackH)
	knobX := float32(2)
	if b.On {
		knobX = trackW - knob - 2
	}
	k := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface), Radius: knob / 2}).SetSize2(knob, knob)
	k.SetPositionKeepSize(true).SetStackPosition(geom.Pt(knobX, 2))
	track.Add(k)
	track.SetOpacity(st.Opacity)
	track.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "switch", Cursor: scene.CursorPointer})

	toggle := func() {
		if b.Disabled {
			return
		}
		b.On = !b.On
		b.MarkDirty()
		if b.OnChange != nil {
			b.OnChange(b.On)
		}
	}
	h := b.tapHandler(ctx, b.Disabled, toggle)
	b.attach(ctx, track, h, !b.Disabled)
	return track
}

// bgOf 根据布尔状态取语义色（不用组件自己的色值）。
func bgOf(th *theme.Data, st ui.Style, on bool) geom.Color {
	if !on {
		return th.Color.SurfaceSunken
	}
	if st.Background == (geom.Color{}) {
		return th.Color.Primary
	}
	return st.Background
}
