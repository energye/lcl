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
	// Ghost 透明背景（antd background-ghost）：文字/边框按类型取色，
	// 背景恒透明；text/link 变体不允许（忽略，同 antd warn 后不生效）。
	Ghost bool
	// IconPosition 图标位于文字之前（默认）或之后（antd iconPosition）。
	IconPosition ButtonIconPosition
	Icon         *Icon // 图标（Loading 时被旋转指示替代）
	OnClick      func()
}

// ButtonIconPosition 是图标位置（antd Button.iconPosition）。
type ButtonIconPosition uint8

const (
	// IconStart 图标在文字前（默认）。
	IconStart ButtonIconPosition = iota
	// IconEnd 图标在文字后。
	IconEnd
)

// isTwoCNChars 判定恰为两个汉字（antd two-chinese-chars 的 letter-spacing 规则）。
func isTwoCNChars(s string) bool {
	r := []rune(s)
	if len(r) != 2 {
		return false
	}
	for _, c := range r {
		if c < 0x4E00 || c > 0x9FA5 {
			return false
		}
	}
	return true
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
	// antd loading 不是禁用态：保持类型色，仅 opacity 0.65 + 忽略点击。
	st := ui.StyleOf(th, v, b.Size, b.states(b.Disabled))
	if b.Danger {
		st = b.dangerize(th, st, v, st2status(b.states(b.Disabled)))
	}
	if b.Loading {
		st.Opacity = 0.65
	}
	// antd ghost：背景恒透明；text/link 变体忽略 ghost。
	if b.Ghost && v != ui.VariantGhost && v != ui.VariantLink {
		status := st2status(b.states(b.Disabled || b.Loading))
		st.Background = geom.Color{}
		switch {
		case b.Disabled || b.Loading:
			st.Foreground, st.Border = th.Color.BorderStrong, th.Color.Border
		case b.Danger:
			st.Foreground, st.Border = th.Color.Danger, th.Color.Danger
			switch status {
			case ui.StatusHover:
				st.Foreground, st.Border = th.Color.DangerHover, th.Color.DangerHover
			case ui.StatusActive:
				st.Foreground, st.Border = th.Color.DangerActive, th.Color.DangerActive
			}
		case v == ui.VariantSolid:
			st.Foreground, st.Border = th.Color.Primary, th.Color.Primary
			switch status {
			case ui.StatusHover:
				st.Foreground, st.Border = th.Color.PrimaryHover, th.Color.PrimaryHover
			case ui.StatusActive:
				st.Foreground, st.Border = th.Color.PrimaryActive, th.Color.PrimaryActive
			}
		default:
			// default ghost：文字/边框 = colorBgContainer（Surface 白），无 hover 变化。
			st.Foreground, st.Border = th.Color.Surface, th.Color.Surface
		}
	}
	m := ui.MetricsFor(th, b.Size)
	font := text.DefaultFont(m.Font)

	// 内容宽度：图标(或 loading) + 间距 + 文本。
	// antd 5.29.3：图标 1em（=contentFontSize），iconGap = marginXS(8) 所有尺寸一致。
	var iconW float32
	if b.Loading || b.Icon != nil {
		iconW = m.Font
	}
	gap := float32(0)
	if iconW > 0 && b.Text != "" {
		gap = th.Space.MD
	}
	// antd icon-only：无文字且带图标时宽 = 控件高。
	iconOnly := b.Text == "" && iconW > 0
	// antd two-chinese-chars：无图标、恰两个汉字、非 text/link。
	twoCN := iconW == 0 && isTwoCNChars(b.Text) && b.Type != TypeText && b.Type != TypeLink
	layout := text.LayoutText(b.Text, font, text.Options{})
	textW := layout.Size.W
	var cnSpacing float32
	if twoCN {
		r := []rune(b.Text)
		l1 := text.LayoutText(string(r[0]), font, text.Options{})
		l2 := text.LayoutText(string(r[1]), font, text.Options{})
		cnSpacing = 0.34 * m.Font
		textW = l1.Size.W + cnSpacing + l2.Size.W
	}

	w := textW + iconW + gap + m.Padding.Min.X + m.Padding.Max.X
	h := m.Height
	if iconOnly {
		w = h
	}
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
		Disabled: b.Disabled || b.Loading,
	})

	// 内容水平居中（内容宽 < 按钮宽时），垂直绝对居中。
	// 钳制只在内容溢出边框盒（contentX<0，被外部挤压）时贴 padding 兜底；
	// icon-only 的 w=32 < 内容+双 padding，(32-14)/2=9 是正确居中，
	// 不能按 padding 钳（否则左 15px 右 3px，加号明显偏右）。
	contentX := (w - textW - iconW - gap) / 2
	if contentX < 0 {
		contentX = m.Padding.Min.X
	}
	textX := contentX
	iconX := contentX
	if iconW > 0 {
		if b.IconPosition == IconEnd {
			iconX = contentX + textW + gap // 图标在右
		} else {
			textX = contentX + iconW + gap // 图标在左
		}
		ic := b.iconFor(ctx)
		in := ic.Build(ctx)
		in.SetPositionKeepSize(true)
		in.SetStackPosition(geom.Pt(iconX, (h-iconW)/2))
		n.Add(in)
	}
	if twoCN {
		// 两字中文：逐字排版，中间 0.34em 字距（GDI 无 letter-spacing）。
		r := []rune(b.Text)
		l1 := text.LayoutText(string(r[0]), font, text.Options{})
		l2 := text.LayoutText(string(r[1]), font, text.Options{})
		t1 := scene.Text(scene.TextSpec{Src: string(r[0]), Layout: l1, Color: st.Foreground})
		t1.SetMargin(geom.InLTRB(textX, (h-l1.Size.H)/2, 0, 0))
		n.Add(t1)
		t2 := scene.Text(scene.TextSpec{Src: string(r[1]), Layout: l2, Color: st.Foreground})
		t2.SetMargin(geom.InLTRB(textX+l1.Size.W+cnSpacing, (h-l2.Size.H)/2, 0, 0))
		n.Add(t2)
	} else if b.Text != "" {
		label := scene.Text(scene.TextSpec{Src: b.Text, Layout: layout, Color: st.Foreground})
		label.SetMargin(geom.InLTRB(textX, (h-layout.Size.H)/2, 0, 0))
		label.SetHit(scene.HitSpec{Label: "button-label"}) // 定位锚（DeferToChild）
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
	h2 := b.tapHandlerPressed(ctx, b.Disabled || b.Loading, activate)
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

	// Loading 旋转：声明活跃动画——容器 60Hz 标脏**重录**（不重建），
	// 图标相位在录制时按墙钟计算，因此每个重录帧角度都在推进。
	if ctx.SetAnimating != nil {
		ctx.SetAnimating(b, b.Loading)
	}
	return n
}

// iconFor 返回前置图标：Loading 时用旋转的 loading 图标，否则用户图标。
func (b *Button) iconFor(ctx *ui.BuildContext) *Icon {
	if b.Loading {
		m := ui.MetricsFor(ctx.ThemeOr(), b.Size)
		return &Icon{Name: "loading", Size: m.Font, Color: b.dangerFg(ctx), Spin: true}
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
	return ui.MetricsFor(ctx.ThemeOr(), ui.SizeMedium).Font
}

// st2status 把状态位集折成单值（避免在外部重复优先级逻辑）。
func st2status(s ui.StateSet) ui.Status { return s.Status() }
