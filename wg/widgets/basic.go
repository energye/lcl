// Package widgets 的 P1 基础与布局组件。
//
// 组件只做三件事：读主题 token、 newNode、把子节点递进来。
// 布局交给 scene 的约束系统，命中交给 input，像素交给 render——
// 组件里出现自己算坐标或自己填颜色的代码就是设计走样。
package widgets

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
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

// ── TextView（可原地更新的文本） ────────────────────────────────────

// TextView 是可在原地更新文本的静态文本。
//
// Label 是值对象：改文本只能整页重建，每键回显会把整页 Build 一遍。
// TextView 持有场景节点，SetText 只改文本节点并标绘制脏（不重排），
// 适合输入回显、计数、状态文案这类高频变化。
//
// Width>0 时外层是固定宽框（超出省略）：原地更新成立的前提是"文本变化
// 不影响父级布局"。需要随内容伸缩的文本请用 Label。
type TextView struct {
	ui.Base
	Text     string
	Size     float32
	Color    geom.Color
	Width    float32 // >0 固定宽，超出省略
	Height   float32
	MaxLines int
	Ellipsis string

	node *scene.Node
	box  *scene.Node
	req  func()
	font text.Font
	opts text.Options
	// color 是 Build 时解析后的颜色（零值 Color 落主题 Text）。
	color geom.Color
}

// NewTextView 创建可原地更新的文本。
func NewTextView(s string) *TextView { return &TextView{Text: s} }

// SetText 原地换文本：只标绘制脏并请求一帧，不触发整页重建。
func (v *TextView) SetText(s string) {
	v.Text = s
	if v.node == nil {
		return
	}
	lay := text.LayoutText(s, v.font, v.opts)
	v.node.SetTextPaint(scene.TextSpec{Src: s, Layout: lay, Color: v.color})
	// 旧字形可能比新文本宽：标外层框（连带子树）重绘。
	if v.box != nil {
		v.box.MarkNeedsPaint()
	}
	if v.req != nil {
		v.req()
	}
}

// Build 实现 ui.Widget。
func (v *TextView) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	size := v.Size
	if size <= 0 {
		size = th.Font.SizeBase
	}
	color := v.Color
	if color == (geom.Color{}) {
		color = th.Color.Text
	}
	ell := v.Ellipsis
	if ell == "" {
		ell = "…"
	}
	font := text.DefaultFont(size)
	opts := text.Options{MaxWidth: v.Width, MaxLines: v.MaxLines, Ellipsis: ell}
	lay := text.LayoutText(v.Text, font, opts)
	t := scene.Text(scene.TextSpec{
		Src: v.Text, Layout: lay, Color: color, LayoutFont: font, LayoutW: v.Width,
	})
	box := scene.Box(scene.BoxSpec{})
	if v.Width > 0 {
		box.SetWidth(v.Width)
	}
	if v.Height > 0 {
		box.SetHeight(v.Height)
	}
	box.Add(t)
	v.node, v.box, v.req = t, box, ctx.RequestFrame
	v.font, v.opts, v.color = font, opts, color
	return box
}

// ── Scroll ─────────────────────────────────────────────────────────

// Scroll 是带滚动条的滚动视口。宽高 <=0 时由父级约束决定（放进 Flex 要设 flex）。
//
// Offset 跨重建保留；ShowScrollbar = nil 时内容溢出才显示滚动条。
type Scroll struct {
	controlBase
	Axis          scene.Axis
	Child         ui.Widget
	Width, Height float32
	// ShowScrollbar：nil = 内容溢出时自动显示；true 始终显示；false 永不显示。
	ShowScrollbar *bool
	// BarThickness 滚动条粗细（0 = 8）。
	BarThickness float32
	// Offset 是滚动偏移（跨重建保留）。
	Offset float32

	view *scene.Node
	bar  *Scrollbar
}

// Build 实现 ui.Widget。
func (s *Scroll) Build(ctx *ui.BuildContext) *scene.Node {
	vertical := s.Axis == scene.Vertical
	w, h := s.Width, s.Height
	if w < 0 {
		w = 0
	}
	if h < 0 {
		h = 0
	}

	child := buildChild(ctx, s.Child)
	view := scene.Scroll(s.Axis, child)
	if w > 0 && h > 0 {
		view.SetSize2(w, h)
	}
	// 视口自身必须可命中：滚轮/空白区域的事件要经过这条路径；
	// 纯布局节点与文本默认不参与命中，不设它滚轮就永远收不到。
	view.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "scroll"})

	// 预测量内容尺寸（与 placeScroll 相同约束）：滚动条比例与偏移钳制要用。
	cs := scene.Unbounded()
	switch {
	case vertical && w > 0:
		cs = scene.NewConstraints(w, w, 0, scene.Inf)
	case !vertical && h > 0:
		cs = scene.NewConstraints(0, scene.Inf, h, h)
	}
	child.Layout(cs)
	content := child.Size()
	viewport, max := h, content.H-h
	if !vertical {
		viewport, max = w, content.W-w
	}
	if max < 0 {
		max = 0
	}
	if s.Offset < 0 {
		s.Offset = 0
	}
	if s.Offset > max {
		s.Offset = max
	}
	view.SetScrollOffset(s.Offset) // 布局后会按真实内容再钳一次

	thick := s.BarThickness
	if thick <= 0 {
		thick = scrollbarThickness
	}
	if s.bar == nil {
		s.bar = &Scrollbar{Axis: s.Axis}
	}
	bar := s.bar
	bar.Axis = s.Axis
	bar.Thickness = thick
	bar.Max = max
	bar.Viewport = viewport
	bar.Value = s.Offset
	if vertical {
		bar.Length = h - 2*scrollbarInset
	} else {
		bar.Length = w - 2*scrollbarInset
	}
	bar.OnChange = func(v float32) {
		s.Offset = v
		if s.view != nil {
			s.view.SetScrollOffset(v)
		}
		if ctx.RequestFrame != nil {
			ctx.RequestFrame()
		}
	}

	// 滚轮：按轴滚动；到边界返回 false，让外层滚动接手（嵌套滚动）。
	scrollBy := func(dy float32) bool {
		next := s.Offset + dy
		if next < 0 {
			next = 0
		}
		if next > bar.Max {
			next = bar.Max
		}
		if next == s.Offset {
			return false
		}
		s.Offset = next
		if s.view != nil {
			s.view.SetScrollOffset(next)
		}
		if s.bar != nil {
			s.bar.SetValue(next)
		}
		if ctx.RequestFrame != nil {
			ctx.RequestFrame()
		}
		return true
	}

	show := max > 0
	if s.ShowScrollbar != nil {
		show = *s.ShowScrollbar
	}
	s.view = view
	if !show {
		if ctx.Input != nil {
			ctx.Input.Set(view, input.Handler{Wheel: func(ev *input.WheelEvent) bool {
				step := ev.Delta.Y
				if !vertical {
					step = ev.Delta.X
				}
				if step == 0 {
					return false
				}
				return scrollBy(step)
			}})
		}
		return view
	}

	stack := scene.Stack()
	if w > 0 && h > 0 {
		stack.SetSize2(w, h)
	}
	view.SetPositionKeepSize(true)
	view.SetStackPosition(geom.Pt(0, 0))
	stack.Add(view)

	barNode := bar.Build(ctx)
	barNode.SetPositionKeepSize(true)
	if vertical {
		barNode.SetStackPosition(geom.Pt(w-thick-scrollbarInset, scrollbarInset))
	} else {
		barNode.SetStackPosition(geom.Pt(scrollbarInset, h-thick-scrollbarInset))
	}
	stack.Add(barNode)

	if ctx.Input != nil {
		ctx.Input.Set(stack, input.Handler{Wheel: func(ev *input.WheelEvent) bool {
			step := ev.Delta.Y
			if !vertical {
				step = ev.Delta.X
			}
			if step == 0 {
				return false
			}
			return scrollBy(step)
		}})
	}
	return stack
}
