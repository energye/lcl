// Package widgets 的 P1 基础与布局组件。
//
// 组件只做三件事：读主题 token、 newNode、把子节点递进来。
// 布局交给 scene 的约束系统，命中交给 input，像素交给 render——
// 组件里出现自己算坐标或自己填颜色的代码就是设计走样。
package widgets

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// buildChild 把可选的子组件变成节点。
func buildChild(ctx *ui.BuildContext, w ui.Widget) *scene.Node {
	if w == nil {
		return nil
	}
	return w.Build(ctx)
}

func buildChildren(ctx *ui.BuildContext, ws []ui.Widget) []*scene.Node {
	out := make([]*scene.Node, 0, len(ws))
	for _, w := range ws {
		if n := buildChild(ctx, w); n != nil {
			out = append(out, n)
		}
	}
	return out
}

// ── Divider ────────────────────────────────────────────────────────

// Divider 是一条分隔线。
type Divider struct {
	ui.Base
	// Vertical 为真时竖直，默认为水平。
	Vertical bool
	// Length 是主线长度；<=0 时水平取 120、竖直取 32。
	Length float32
	// Thickness 是粗细；<=0 时用 1。
	Thickness float32
	// Color 零值时取主题边框色。
	Color geom.Color
}

// NewDivider 创建水平分隔线。
func NewDivider() *Divider { return &Divider{} }

// Build 实现 ui.Widget。
func (d *Divider) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	color := d.Color
	if color == (geom.Color{}) {
		color = th.Color.Border
	}
	length, thick := d.Length, d.Thickness
	if length <= 0 {
		if d.Vertical {
			length = 32
		} else {
			length = 120
		}
	}
	if thick <= 0 {
		thick = 1
	}
	if d.Vertical {
		return scene.Box(scene.BoxSpec{Fill: scene.Solid(color)}).SetSize2(thick, length)
	}
	return scene.Box(scene.BoxSpec{Fill: scene.Solid(color)}).SetSize2(length, thick)
}

// ── Spacer / SizedBox ──────────────────────────────────────────────

// Spacer 占据固定空间，不绘制任何东西。
type Spacer struct {
	Width, Height float32
}

// Build 实现 ui.Widget。
func (s Spacer) Build(*ui.BuildContext) *scene.Node {
	return scene.Empty().SetSize2(s.Width, s.Height)
}

// SizedBox 给子组件一个固定尺寸。
type SizedBox struct {
	Width, Height float32
	Child         ui.Widget
}

// Build 实现 ui.Widget。
func (s SizedBox) Build(ctx *ui.BuildContext) *scene.Node {
	n := scene.Box(scene.BoxSpec{})
	n.SetSize2(s.Width, s.Height)
	if child := buildChild(ctx, s.Child); child != nil {
		n.Add(child)
	}
	return n
}

// ── Padding ────────────────────────────────────────────────────────

// Padding 在子组件四周加内边距。
//
// 用 Box 而不是"手动 SetPosition(child, l, t)"：padding 必须参与
// intrinsic 尺寸计算，否则父级 Flex 会把带 padding 的子项当成没有边距。
type Padding struct {
	Insets geom.Insets
	Child  ui.Widget
}

// Build 实现 ui.Widget。
func (p Padding) Build(ctx *ui.BuildContext) *scene.Node {
	n := scene.Box(scene.BoxSpec{})
	n.SetPadding(p.Insets)
	if child := buildChild(ctx, p.Child); child != nil {
		n.Add(child)
	}
	return n
}

// ── Row / Column ───────────────────────────────────────────────────

// Row 水平排列子组件。
type Row struct {
	Gap        float32
	Children   []ui.Widget
	MainAlign  scene.MainAlign
	CrossAlign scene.CrossAlign
}

// Build 实现 ui.Widget。
func (r Row) Build(ctx *ui.BuildContext) *scene.Node {
	n := scene.Flex(scene.Horizontal, r.Gap, buildChildren(ctx, r.Children)...)
	n.SetMainAlign(r.MainAlign).SetCrossAlign(r.CrossAlign)
	return n
}

// Column 垂直排列子组件。
type Column struct {
	Gap        float32
	Children   []ui.Widget
	MainAlign  scene.MainAlign
	CrossAlign scene.CrossAlign
}

// Build 实现 ui.Widget。
func (c Column) Build(ctx *ui.BuildContext) *scene.Node {
	n := scene.Flex(scene.Vertical, c.Gap, buildChildren(ctx, c.Children)...)
	n.SetMainAlign(c.MainAlign).SetCrossAlign(c.CrossAlign)
	return n
}

// ── Label ──────────────────────────────────────────────────────────

// Label 是一段静态文本。
type Label struct {
	Text string
	// Size <=0 时用主题 SizeBase。
	Size     float32
	Color    geom.Color
	MaxWidth float32
	MaxLines int
	Ellipsis string
}

// Build 实现 ui.Widget。
func (l Label) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	size := l.Size
	if size <= 0 {
		size = th.Font.SizeBase
	}
	color := l.Color
	if color == (geom.Color{}) {
		color = th.Color.Text
	}
	font := text.DefaultFont(size)
	layout := text.LayoutText(l.Text, font, text.Options{
		MaxWidth: l.MaxWidth, MaxLines: l.MaxLines, Ellipsis: l.Ellipsis,
		LineHeight: th.Font.LineHeight,
	})
	return scene.Text(scene.TextSpec{
		Src: l.Text, Layout: layout, Color: color,
		LayoutFont: font, LayoutW: l.MaxWidth,
	})
}

// ── Scroll ─────────────────────────────────────────────────────────

// Scroll 是滚动视口。宽高 <=0 时由父级约束决定（放进 Flex 要设 flex）。
type Scroll struct {
	Axis          scene.Axis
	Child         ui.Widget
	Width, Height float32
}

// Build 实现 ui.Widget。
func (s Scroll) Build(ctx *ui.BuildContext) *scene.Node {
	n := scene.Scroll(s.Axis, buildChild(ctx, s.Child))
	if s.Width > 0 || s.Height > 0 {
		n.SetSize2(s.Width, s.Height)
	}
	return n
}
