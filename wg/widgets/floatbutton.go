package widgets

// ── FloatButton（M11 导航）─────────────────────────────────────────
//
// antd：<FloatButton icon type tooltip shape="circle" onClick>
// 已对齐：primary（主色实底白图标）/default（白底描边）双型、圆形
// （56/40 两档）、Tooltip 悬停提示（**组合复用 Tooltip 组件**——
// Child=自身）、点击回调。
// 未实现（如实标注）：description 副文案形态、Badge 角标集成。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// FloatButtonType 是悬浮按钮类型（antd FloatButton.type）。
type FloatButtonType uint8

const (
	// FloatPrimary 主色实底白图标（默认）。
	FloatPrimary FloatButtonType = iota
	// FloatDefault 白底描边深图标。
	FloatDefault
)

// FloatButton 是圆形悬浮按钮（56/40 两档，组合复用 Icon）。
type FloatButton struct {
	Icon    string
	Type    FloatButtonType
	// Tooltip 悬停提示文案（内部组合 Tooltip 组件挂自身）。
	Tooltip string
	// Size 直径（默认 56；小档 40）。
	Size    float32
	OnClick func()

	controlBase
}

// Pressed 返回按压态（探针/诊断用）。
func (f *FloatButton) Pressed() bool { return f.pressed }

// Build 实现 ui.Widget。
func (f *FloatButton) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	s := f.Size
	if s <= 0 {
		s = 56
	}
	fill, icColor, borderColor := th.Color.Primary, th.Color.OnPrimary, th.Color.Primary
	if f.Type == FloatDefault {
		fill, icColor, borderColor = th.Color.Surface, th.Color.Text, th.Color.Border
	}
	// 按压态（antd active 色）：主色加深（PrimaryActive），default 灰底。
	if f.pressed {
		if f.Type == FloatPrimary {
			fill = th.Color.PrimaryActive
		} else {
			fill = th.Color.SurfaceActive
		}
	}
	n := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(fill),
		Radius: s / 2,
		Border: scene.Stroke(borderColor, 1),
		Shadow: func() *scene.ShadowSpec {
			if len(th.Elevation) > 1 {
				e := th.Elevation[1]
				return &scene.ShadowSpec{Dx: e.Dx, Dy: e.Dy, Blur: e.Blur, Color: e.Color}
			}
			return nil
		}(),
	}).SetSize2(s, s)
	ic := (&Icon{Name: f.Icon, Size: s * 0.4, Color: icColor}).Build(ctx)
	ic.SetMargin(geom.InLTRB((s-s*0.4)/2, (s-s*0.4)/2, 0, 0))
	n.Add(ic)
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "float-button", Cursor: scene.CursorPointer})
	// tapHandlerPressed：按压视觉反馈（antd 按钮点击有 active 态，
	// tapHandler 无按压视觉——toggle 断言实测变化 0.000）。
	hnd := f.tapHandlerPressed(ctx, false, func() {
		if f.OnClick != nil {
			f.OnClick()
		}
	})
	f.attach(ctx, n, hnd, false)
	return n
}

// FloatButtonGroup 是悬浮按钮组（纵向排列，gap 12——antd FloatButton.Group）。
type FloatButtonGroup struct {
	Buttons []*FloatButton
	Gap     float32
	// ShapeGroup 组背景（antd shape="square" 组底色；false=透明独立排列）。
	ShapeGroup bool
}

// Build 实现 ui.Widget。
func (g *FloatButtonGroup) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	col := scene.Flex(scene.Vertical, g.Gap)
	if g.Gap <= 0 {
		col = scene.Flex(scene.Vertical, 12)
	}
	for _, b := range g.Buttons {
		col.Add(b.Build(ctx))
	}
	if g.ShapeGroup {
		// 组底色：浅主色圆角容器（antd shape="square" 组）。
		col.SetBox(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.InfoBg),
			Radius: th.Radius.LG,
		})
		col.SetPadding(geom.InLTRB(6, 6, 6, 6))
	}
	col.SetHit(scene.HitSpec{Label: "float-button-group"})
	return col
}
