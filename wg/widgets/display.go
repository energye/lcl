// Package widgets 的 P2 展示型组件。
//
// 这一组都是"纯呈现"组件：读主题 token、把 props 折成 scene 节点，
// 不持有像素、不碰 lcl。需要交互的组件挂 Handler，图形（Spin 等）
// 走 CustomPaintNode + render.Canvas——这是 C17 允许的录制通道。
package widgets

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// ── Card ───────────────────────────────────────────────────────────

// Card 是一张带标题与可选底部操作区的卡片。
//
// 宽高由父级约束决定：放进 Flex 时给 flex，卡片自己只按内容增长。
type Card struct {
	Title    ui.Widget
	Content  ui.Widget
	Footer   ui.Widget
	Padding  geom.Insets
	Elevated bool
}

// NewCard 创建卡片。
func NewCard(content ui.Widget) *Card { return &Card{Content: content, Elevated: true} }

// Build 实现 ui.Widget。
func (c Card) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	pad := c.Padding
	if pad == (geom.Insets{}) {
		pad = geom.InLTRB(th.Space.LG, th.Space.LG, th.Space.LG, th.Space.LG)
	}
	spec := scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.LG,
	}
	if c.Elevated && len(th.Elevation) > 0 {
		e := th.Elevation[0]
		spec.Shadow = &scene.ShadowSpec{Dx: e.Dx, Dy: e.Dy, Blur: e.Blur, Color: e.Color}
	}
	col := scene.Flex(scene.Vertical, th.Space.MD)
	if c.Title != nil {
		col.Add(buildChild(ctx, c.Title))
		col.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetHeight(1).SetFlex(1, scene.FlexTight))
	}
	if c.Content != nil {
		col.Add(buildChild(ctx, c.Content))
	}
	if c.Footer != nil {
		col.Add(buildChild(ctx, c.Footer))
	}
	n := scene.Box(spec, col)
	n.SetPadding(pad)
	return n
}

// ── Badge / Tag ────────────────────────────────────────────────────

// Badge 是角标（小圆点或数字），常作为其它组件的装饰子组件。
type Badge struct {
	Text      string
	Count     int
	Dot       bool
	Color     geom.Color
	TextColor geom.Color
}

// Build 实现 ui.Widget。
func (b Badge) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	fill := b.Color
	if fill == (geom.Color{}) {
		fill = th.Color.Danger
	}
	if b.Dot {
		return scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: th.Radius.Pill}).SetSize2(8, 8)
	}
	label := b.Text
	if label == "" && b.Count > 0 {
		label = fmt.Sprintf("%d", b.Count)
	}
	font := text.DefaultFont(th.Font.SizeXS)
	layout := text.LayoutText(label, font, text.Options{})
	w := layout.Size.W + th.Space.MD
	h := layout.Size.H + th.Space.XS*2
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: h / 2}).SetSize2(w, h)
	fg := b.TextColor
	if fg == (geom.Color{}) {
		fg = th.Color.OnDanger
	}
	t := scene.Text(scene.TextSpec{Src: label, Layout: layout, Color: fg})
	t.SetMargin(geom.InLTRB(th.Space.SM, (h-layout.Size.H)/2, 0, 0))
	n.Add(t)
	return n
}

// Tag 是标签（浅底圆角文字）。
type Tag struct {
	Text  string
	Color geom.Color // 零值时取主题 Primary
}

// Build 实现 ui.Widget。
func (t Tag) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	fg := t.Color
	if fg == (geom.Color{}) {
		fg = th.Color.Primary
	}
	font := text.DefaultFont(th.Font.SizeSM)
	layout := text.LayoutText(t.Text, font, text.Options{})
	w := layout.Size.W + th.Space.MD
	h := layout.Size.H + th.Space.SM
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(fg.WithAlpha(40)), Radius: th.Radius.SM}).SetSize2(w, h)
	label := scene.Text(scene.TextSpec{Src: t.Text, Layout: layout, Color: fg})
	label.SetMargin(geom.InLTRB(th.Space.SM, (h-layout.Size.H)/2, 0, 0))
	n.Add(label)
	return n
}

// ── Avatar ─────────────────────────────────────────────────────────

// Avatar 是圆形头像：有文字时取首字符，否则显示纯色圆。
type Avatar struct {
	Text  string
	Size  float32
	Color geom.Color
}

// Build 实现 ui.Widget。
func (a Avatar) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	s := a.Size
	if s <= 0 {
		s = 32
	}
	fill := a.Color
	if fill == (geom.Color{}) {
		fill = th.Color.Primary
	}
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: s / 2}).SetSize2(s, s)
	if a.Text != "" {
		r := []rune(a.Text)
		label := string(r[:1])
		font := text.DefaultFont(s * 0.45)
		layout := text.LayoutText(label, font, text.Options{})
		t := scene.Text(scene.TextSpec{Src: label, Layout: layout, Color: th.Color.OnPrimary})
		t.SetMargin(geom.InLTRB((s-layout.Size.W)/2, (s-layout.Size.H)/2, 0, 0))
		n.Add(t)
	}
	return n
}

// ── Empty / Skeleton / Spin ────────────────────────────────────────

// Empty 是空状态占位（图标圆 + 主文案 + 可选副文案）。
type Empty struct {
	Text    string
	Subtext string
	Width   float32
}

// Build 实现 ui.Widget。
func (e Empty) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	w := e.Width
	if w <= 0 {
		w = 220
	}
	n := scene.Box(scene.BoxSpec{}).SetSize2(w, 96)
	icon := scene.Box(scene.BoxSpec{
		Fill: scene.Solid(th.Color.SurfaceSunken), Radius: 20,
	}).SetSize2(40, 40)
	icon.SetStackPosition(geom.Pt((w-40)/2, 0)).SetPositionKeepSize(true)
	n.Add(icon)
	label := e.Text
	if label == "" {
		label = "暂无数据"
	}
	font := text.DefaultFont(th.Font.SizeBase)
	layout := text.LayoutText(label, font, text.Options{})
	t := scene.Text(scene.TextSpec{Src: label, Layout: layout, Color: th.Color.TextMuted})
	t.SetStackPosition(geom.Pt((w-layout.Size.W)/2, 48)).SetPositionKeepSize(true)
	n.Add(t)
	if e.Subtext != "" {
		sub := text.LayoutText(e.Subtext, text.DefaultFont(th.Font.SizeSM), text.Options{})
		st := scene.Text(scene.TextSpec{Src: e.Subtext, Layout: sub, Color: th.Color.TextFaint})
		st.SetStackPosition(geom.Pt((w-sub.Size.W)/2, 70)).SetPositionKeepSize(true)
		n.Add(st)
	}
	return n
}

// Skeleton 是加载骨架（n 条灰色圆角条，最后一条较短）。
type Skeleton struct {
	Rows   int
	Width  float32
	Height float32
}

// Build 实现 ui.Widget。
func (s Skeleton) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	rows := s.Rows
	if rows <= 0 {
		rows = 3
	}
	w, h := s.Width, s.Height
	if w <= 0 {
		w = 240
	}
	if h <= 0 {
		h = 12
	}
	n := scene.Flex(scene.Vertical, th.Space.MD)
	for i := 0; i < rows; i++ {
		rw := w
		if i == rows-1 {
			rw = w * 0.6
		}
		n.Add(scene.Box(scene.BoxSpec{
			Fill: scene.Solid(th.Color.SurfaceSunken), Radius: th.Radius.SM,
		}).SetSize2(rw, h))
	}
	return n
}

// Spin 是加载指示器（环形缺口；缺口位置由 Phase 决定，外部动画驱动）。
type Spin struct {
	Size  float32
	Phase float32 // 0..1
}

// Build 实现 ui.Widget。
func (s Spin) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	size := s.Size
	if size <= 0 {
		size = 20
	}
	n := scene.CustomPaintNode(func(c render.Canvas, sz geom.Size) {
		p := render.NewPath()
		p.AddArc(geom.RectXYWH(2, 2, sz.W-4, sz.H-4), s.Phase*6.2831853, 4.7)
		c.StrokePath(p, render.Stroke(th.Color.Primary, 2).WithCap(render.CapRound))
	})
	n.SetSize2(size, size)
	return n
}

// ── Steps ───────────────────────────────────────────────────────────

// Steps 是步骤条。Current 为当前步（0 基），之前的步骤显示对勾。
type Steps struct {
	Labels  []string
	Current int
}

// Build 实现 ui.Widget。
func (s Steps) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	n := scene.Flex(scene.Horizontal, th.Space.MD)
	for i, label := range s.Labels {
		step := scene.Flex(scene.Vertical, th.Space.XS)
		fill := th.Color.SurfaceSunken
		fg := th.Color.TextMuted
		if i <= s.Current {
			fill, fg = th.Color.Primary, th.Color.OnPrimary
		}
		circle := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: 10}).SetSize2(20, 20)
		cap := fmt.Sprintf("%d", i+1)
		if i < s.Current {
			cap = "✓"
		}
		fl := text.LayoutText(cap, text.DefaultFont(th.Font.SizeXS), text.Options{})
		ct := scene.Text(scene.TextSpec{Src: cap, Layout: fl, Color: fg})
		ct.SetMargin(geom.InLTRB((20-fl.Size.W)/2, (20-fl.Size.H)/2, 0, 0))
		circle.Add(ct)
		step.Add(circle)
		ll := text.LayoutText(label, text.DefaultFont(th.Font.SizeSM), text.Options{})
		step.Add(scene.Text(scene.TextSpec{Src: label, Layout: ll, Color: th.Color.Text}))
		n.Add(step)
	}
	return n
}

// ── Segmented ──────────────────────────────────────────────────────

// Segmented 是分段选择器。选中段用 Primary 底色。
type Segmented struct {
	controlBase
	Options  []string
	Selected int
	OnChange func(int)
}

// NewSegmented 创建分段选择器。
func NewSegmented(options ...string) *Segmented { return &Segmented{Options: options} }

// Build 实现 ui.Widget。
func (s *Segmented) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(s)
	th := ctx.ThemeOr()
	children := make([]*scene.Node, 0, len(s.Options))
	for i, opt := range s.Options {
		seg := &segmentItem{index: i, label: opt, parent: s}
		children = append(children, seg.Build(ctx))
	}
	row := scene.Flex(scene.Horizontal, 0, children...)
	n := scene.Box(scene.BoxSpec{
		Fill: scene.Solid(th.Color.SurfaceSunken), Radius: th.Radius.Base,
	}, row)
	n.SetPadding(geom.InLTRB(2, 2, 2, 2))
	s.node = n
	return n
}

// segmentItem 是单个分段（内部组件）。
type segmentItem struct {
	controlBase
	index  int
	label  string
	parent *Segmented
}

func (s *segmentItem) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	selected := s.index == s.parent.Selected
	fill := geom.Color{}
	fg := th.Color.Text
	if selected {
		fill, fg = th.Color.Primary, th.Color.OnPrimary
	}
	font := text.DefaultFont(th.Font.SizeSM)
	layout := text.LayoutText(s.label, font, text.Options{})
	w := layout.Size.W + th.Space.LG
	const h = float32(26)
	item := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: th.Radius.SM}).SetSize2(w, h)
	label := scene.Text(scene.TextSpec{Src: s.label, Layout: layout, Color: fg})
	label.SetMargin(geom.InLTRB(th.Space.SM, (h-layout.Size.H)/2, 0, 0))
	item.Add(label)
	item.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "segment", Cursor: scene.CursorPointer})
	parent, index := s.parent, s.index
	hnd := s.tapHandler(ctx, false, func() {
		if parent.Selected == index {
			return
		}
		parent.Selected = index
		parent.MarkDirty()
		if parent.OnChange != nil {
			parent.OnChange(index)
		}
	})
	s.attach(ctx, item, hnd, false)
	return item
}

// ── Pagination ─────────────────────────────────────────────────────

// Pagination 是分页器。Page 从 1 开始。
type Pagination struct {
	controlBase
	Page     int
	Total    int // 总页数
	OnChange func(int)
}

// NewPagination 创建分页器。
func NewPagination(page, total int) *Pagination {
	if page < 1 {
		page = 1
	}
	return &Pagination{Page: page, Total: total}
}

// Build 实现 ui.Widget。
func (p *Pagination) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	row := scene.Flex(scene.Horizontal, th.Space.XS)
	for _, page := range p.visiblePages() {
		row.Add((&pageButton{page: page, current: page == p.Page, parent: p,
			label: fmt.Sprintf("%d", page)}).Build(ctx))
	}
	p.node = row
	return row
}

func (p *Pagination) visiblePages() []int {
	total := p.Total
	if total < 1 {
		return []int{1}
	}
	lo, hi := 1, total
	if total > 7 {
		lo, hi = p.Page-2, p.Page+2
		if lo < 1 {
			lo, hi = 1, 5
		}
		if hi > total {
			hi, lo = total, total-4
		}
	}
	out := make([]int, 0, hi-lo+1)
	for i := lo; i <= hi; i++ {
		out = append(out, i)
	}
	return out
}

// pageButton 是单个页码按钮（内部组件）。
type pageButton struct {
	controlBase
	page    int
	current bool
	label   string
	parent  *Pagination
}

func (b *pageButton) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	fill := geom.Color{}
	fg := th.Color.Text
	if b.current {
		fill, fg = th.Color.Primary, th.Color.OnPrimary
	}
	layout := text.LayoutText(b.label, text.DefaultFont(th.Font.SizeSM), text.Options{})
	w := layout.Size.W + th.Space.MD
	if w < 26 {
		w = 26
	}
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: th.Radius.SM}).SetSize2(w, 26)
	t := scene.Text(scene.TextSpec{Src: b.label, Layout: layout, Color: fg})
	t.SetMargin(geom.InLTRB((w-layout.Size.W)/2, (26-layout.Size.H)/2, 0, 0))
	n.Add(t)
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "page", Cursor: scene.CursorPointer})
	parent, page := b.parent, b.page
	hnd := b.tapHandler(ctx, false, func() {
		if parent.Page == page {
			return
		}
		parent.Page = page
		parent.MarkDirty()
		if parent.OnChange != nil {
			parent.OnChange(page)
		}
	})
	b.attach(ctx, n, hnd, false)
	return n
}

// ── Breadcrumb ─────────────────────────────────────────────────────

// Breadcrumb 是面包屑导航。最后一项为当前项（高亮且不可点）。
type Breadcrumb struct {
	Items   []string
	OnClick func(int)
	node    *scene.Node
	taps    []*ui.GestureDetector // 跨 Build 存活，避免回调闭包被 GC
}

// Node 返回最近一次构建出的场景节点。
func (b *Breadcrumb) Node() *scene.Node { return b.node }

// Build 实现 ui.Widget。
func (b *Breadcrumb) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	row := scene.Flex(scene.Horizontal, th.Space.SM)
	b.taps = b.taps[:0]
	for i, label := range b.Items {
		if i > 0 {
			sep := text.LayoutText("›", text.DefaultFont(th.Font.SizeBase), text.Options{})
			row.Add(scene.Text(scene.TextSpec{Src: "›", Layout: sep, Color: th.Color.TextFaint}))
		}
		color := th.Color.TextMuted
		last := i == len(b.Items)-1
		if last {
			color = th.Color.Text
		}
		layout := text.LayoutText(label, text.DefaultFont(th.Font.SizeBase), text.Options{})
		item := scene.Text(scene.TextSpec{Src: label, Layout: layout, Color: color})
		if !last && b.OnClick != nil {
			idx := i
			g := &ui.GestureDetector{OnTap: func(geom.Point) { b.OnClick(idx) }}
			b.taps = append(b.taps, g)
			item.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "breadcrumb", Cursor: scene.CursorPointer})
			if ctx.Input != nil {
				ctx.Input.Set(item, g.Handler())
			}
		}
		row.Add(item)
	}
	b.node = row
	return row
}

// ── Collapse ───────────────────────────────────────────────────────

// CollapsePanel 是折叠面板的一项。
type CollapsePanel struct {
	Title   string
	Content ui.Widget
}

// Collapse 是可折叠面板组。Open 为展开项索引，-1 表示全部收起。
type Collapse struct {
	controlBase
	Panels   []CollapsePanel
	Open     int
	OnChange func(int)
}

// NewCollapse 创建折叠面板组（默认全部收起）。
func NewCollapse(panels ...CollapsePanel) *Collapse {
	return &Collapse{Panels: panels, Open: -1}
}

// Build 实现 ui.Widget。
func (c *Collapse) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(c)
	th := ctx.ThemeOr()
	col := scene.Flex(scene.Vertical, th.Space.XS)
	for i := range c.Panels {
		col.Add((&collapseItem{index: i, parent: c}).Build(ctx))
	}
	c.node = col
	return col
}

// collapseItem 是单个折叠项（内部组件）。
type collapseItem struct {
	controlBase
	index  int
	parent *Collapse
}

func (c *collapseItem) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	panel := c.parent.Panels[c.index]
	open := c.index == c.parent.Open
	header := scene.Box(scene.BoxSpec{
		Fill: scene.Solid(th.Color.Surface), Border: scene.Stroke(th.Color.Border, 1), Radius: th.Radius.Base,
	}).SetHeight(36)
	layout := text.LayoutText(panel.Title, text.DefaultFont(th.Font.SizeBase), text.Options{})
	label := scene.Text(scene.TextSpec{Src: panel.Title, Layout: layout, Color: th.Color.Text})
	label.SetMargin(geom.InLTRB(th.Space.MD, (36-layout.Size.H)/2, 0, 0))
	header.Add(label)
	arrow := "▸"
	if open {
		arrow = "▾"
	}
	al := text.LayoutText(arrow, text.DefaultFont(th.Font.SizeSM), text.Options{})
	at := scene.Text(scene.TextSpec{Src: arrow, Layout: al, Color: th.Color.TextMuted})
	at.SetPositionKeepSize(true).SetStackPosition(geom.Pt(320, (36-al.Size.H)/2))
	header.Add(at)
	header.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "collapse", Cursor: scene.CursorPointer})
	parent, index := c.parent, c.index
	hnd := c.tapHandler(ctx, false, func() {
		next := -1
		if parent.Open != index {
			next = index
		}
		parent.Open = next
		parent.MarkDirty()
		if parent.OnChange != nil {
			parent.OnChange(next)
		}
	})
	c.attach(ctx, header, hnd, false)
	root := scene.Flex(scene.Vertical, 0, header)
	if open && panel.Content != nil {
		root.Add(buildChild(ctx, panel.Content))
	}
	return root
}
