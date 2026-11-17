package widgets

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── Button（M8 对齐 antd v5）────────────────────────────────────────
//
// antd：<Button type="primary|default|dashed|text|link" danger size
//
//	loading block shape icon onClick>
//
// 视觉对齐点：主色三态/hover 转主色的描边/dashed 虚线边框/text 按钮的
// fill 底/link 的主色文字/danger 系（bg #FF4D4F、hover #FF7875、
// active #D9363E、default-danger 悬停底 #FFF1F0）/禁用专色不降透明度/
// paddingInline 15（small 7）/loading 旋转指示。
// 未实现（如实标注）：href 链接跳转、ghost 修饰符、图标两态按钮。

// ButtonType 是按钮类型（antd Button.type）。
type ButtonType uint8

const (
	// TypeDefault 默认按钮（灰底描边）。
	TypeDefault ButtonType = iota
	// TypePrimary 主按钮（主色实心）。
	TypePrimary
	// TypeDashed 虚线按钮。
	TypeDashed
	// TypeText 文字按钮。
	TypeText
	// TypeLink 链接按钮。
	TypeLink
)

// ButtonShape 是按钮形状（antd Button.shape）。
type ButtonShape uint8

const (
	// ShapeDefault 默认圆角。
	ShapeDefault ButtonShape = iota
	// ShapeCircle 圆形（图标按钮常用）。
	ShapeCircle
	// ShapeRound 胶囊形。
	ShapeRound
)

// Button 是按钮。点击与 Enter/Space 走同一个 OnClick。
type Button struct {
	controlBase
	Text     string
	Type     ButtonType
	Size     ui.Size
	Shape    ButtonShape
	Danger   bool // 危险按钮（语义色整体切到 Danger 系）
	Loading  bool // 加载中：显示旋转指示并忽略点击
	Block    bool // 宽度占满父级交叉轴
	Disabled bool
	Icon     *Icon // 前置图标（Loading 时被旋转指示替代）
	OnClick  func()
}

// NewButton 创建默认按钮。
func NewButton(text string) *Button { return &Button{Text: text, Size: ui.SizeMedium} }

// buttonBase 供 ControlGroup/auxiliary 逻辑复用
var buttonSpinBase = time.Now()

// variantOf 把 antd type 映射到样式系统的 Variant。
func (b *Button) variantOf() ui.Variant {
	switch b.Type {
	case TypePrimary:
		return ui.VariantSolid
	case TypeDashed:
		return ui.VariantDashed
	case TypeText:
		return ui.VariantGhost
	case TypeLink:
		return ui.VariantLink
	}
	return ui.VariantOutline
}

// dangerize 把解析后的配色整体切到 Danger 系（antd danger 语义）。
// 规则：实心=Error 实底；描边/虚线=Error 文字与边框；文字/链接=Error 文字。
func (b *Button) dangerize(th *theme.Data, st ui.Style, v ui.Variant, s ui.Status) ui.Style {
	c := th.Color
	switch v {
	case ui.VariantSolid:
		st.Background, st.Foreground, st.Border = c.Danger, c.OnDanger, c.Danger
		switch s {
		case ui.StatusHover:
			st.Background = c.DangerHover
		case ui.StatusActive:
			st.Background = c.DangerActive
		}
	default:
		st.Foreground = c.Danger
		if st.BorderWidth > 0 {
			st.Border = c.Danger
		}
		switch s {
		case ui.StatusHover:
			st.Foreground = c.DangerHover
			if st.BorderWidth > 0 {
				st.Border = c.DangerHover
			}
			if v == ui.VariantOutline || v == ui.VariantDashed {
				st.Background = c.DangerBg
			}
		case ui.StatusActive:
			st.Foreground = c.DangerActive
			if st.BorderWidth > 0 {
				st.Border = c.DangerActive
			}
		}
	}
	return st
}

// Build 实现 ui.Widget。
func (b *Button) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	v := b.variantOf()
	st := ui.StyleOf(th, v, b.Size, b.states(b.Disabled || b.Loading))
	if b.Danger {
		st = b.dangerize(th, st, v, st2status(b.states(b.Disabled || b.Loading)))
	}
	m := ui.MetricsFor(th, b.Size)
	font := text.DefaultFont(m.Font)

	// 内容宽度：图标(或 loading) + 间距 + 文本。
	var iconW float32
	if b.Loading || b.Icon != nil {
		iconW = m.Font * 1.1
	}
	gap := float32(0)
	if iconW > 0 && b.Text != "" {
		gap = m.Gap
	}
	layout := text.LayoutText(b.Text, font, text.Options{})
	textW := layout.Size.W

	w := textW + iconW + gap + m.Padding.Min.X + m.Padding.Max.X
	h := m.Height
	switch b.Shape {
	case ShapeCircle:
		if w < h {
			w = h
		}
	case ShapeRound:
	}
	spec := scene.BoxSpec{Fill: scene.Solid(st.Background), Radius: st.Radius}
	switch b.Shape {
	case ShapeCircle, ShapeRound:
		spec.Radius = h / 2
	}
	if st.BorderWidth > 0 {
		if st.Dashed {
			spec.Border = render.Dashed(st.Border, st.BorderWidth, 5, 3)
		} else {
			spec.Border = scene.Stroke(st.Border, st.BorderWidth)
		}
	}
	n := scene.Box(spec)
	if b.Block {
		n.SetFlex(1, scene.FlexTight)
	} else {
		n.SetWidth(w)
	}
	n.SetHeight(h)
	n.SetOpacity(st.Opacity)
	n.SetHit(scene.HitSpec{
		Behavior: scene.HitOpaque, Label: "button", Cursor: scene.CursorPointer,
	})

	// 内容水平居中（内容宽 < 按钮宽时），垂直绝对居中。
	contentX := (w - textW - iconW - gap) / 2
	if contentX < m.Padding.Min.X {
		contentX = m.Padding.Min.X
	}
	if iconW > 0 {
		ic := b.iconFor(ctx)
		in := ic.Build(ctx)
		in.SetPositionKeepSize(true)
		in.SetStackPosition(geom.Pt(contentX, (h-iconW)/2))
		n.Add(in)
	}
	if b.Text != "" {
		label := scene.Text(scene.TextSpec{Src: b.Text, Layout: layout, Color: st.Foreground})
		label.SetMargin(geom.InLTRB(contentX+iconW+gap, (h-layout.Size.H)/2, 0, 0))
		n.Add(label)
	}

	activate := func() {
		if b.Disabled || b.Loading {
			return
		}
		if b.OnClick != nil {
			b.OnClick()
		}
	}
	h2 := b.tapHandler(ctx, b.Disabled || b.Loading, activate)
	h2.Key = func(e *input.KeyEvent) bool {
		if e.Kind == input.KeyUp || b.Disabled || b.Loading {
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
	b.attach(ctx, n, h2, !b.Disabled)

	// Loading 续帧：旋转指示需要帧泵持续推进（真窗；headless 单帧静止）。
	if b.Loading && ctx.Scheduler != nil {
		ctx.Scheduler.RequestFrameIn(33 * time.Millisecond)
	}
	return n
}

// iconFor 返回前置图标：Loading 时用旋转的 loading 图标，否则用户图标。
func (b *Button) iconFor(ctx *ui.BuildContext) *Icon {
	if b.Loading {
		m := ui.MetricsFor(ctx.ThemeOr(), b.Size)
		return &Icon{Name: "loading", Size: m.Font * 1.1, Color: b.dangerFg(ctx), Spin: true}
	}
	ic := b.Icon
	ic.Size = m0size(ctx, ic.Size)
	ic.Color = b.dangerFg(ctx)
	return ic
}

func (b *Button) dangerFg(ctx *ui.BuildContext) geom.Color {
	if b.Danger {
		return ctx.ThemeOr().Color.Danger
	}
	return ui.StyleOf(ctx.ThemeOr(), b.variantOf(), b.Size, b.states(b.Disabled || b.Loading)).Foreground
}

func m0size(ctx *ui.BuildContext, s float32) float32 {
	if s > 0 {
		return s
	}
	return ui.MetricsFor(ctx.ThemeOr(), ui.SizeMedium).Font * 1.1
}

// st2status 把状态位集折成单值（避免在外部重复优先级逻辑）。
func st2status(s ui.StateSet) ui.Status { return s.Status() }
