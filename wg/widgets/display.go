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
	// 定位锚（自动化测试选择器；DeferToChild 不参与命中，不影响交互）。
	n.SetHit(scene.HitSpec{Label: "card"})
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
		dot := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: th.Radius.Pill}).SetSize2(8, 8)
		dot.SetHit(scene.HitSpec{Label: "badge"})
		return dot
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
	n.SetHit(scene.HitSpec{Label: "badge"})
	return n
}

// Tag 是标签（浅底圆角文字）。
type Tag struct {
	Text  string
	Color geom.Color // 零值时取主题 Primary
}

// Build 实现 ui.Widget。
//
// antd Tag 两态：**默认**（无 color）= 灰底 colorFillQuaternary + colorBorder
// 描边 + colorText 文字；**彩色**（显式 Color）= 浅底（基色 40% 透明）+ 同色
// 文字 + 基色淡描边。
func (t Tag) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	textColor, fillColor, borderColor := th.Color.Text, th.Color.SurfaceSunken, th.Color.Border
	if t.Color != (geom.Color{}) {
		textColor = t.Color
		fillColor = t.Color.WithAlpha(40)
		borderColor = t.Color.WithAlpha(80)
	}
	font := text.DefaultFont(th.Font.SizeSM)
	layout := text.LayoutText(t.Text, font, text.Options{})
	w := layout.Size.W + th.Space.MD
	h := layout.Size.H + th.Space.SM
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(fillColor), Border: scene.Stroke(borderColor, 1), Radius: th.Radius.SM}).SetSize2(w, h)
	label := scene.Text(scene.TextSpec{Src: t.Text, Layout: layout, Color: textColor})
	label.SetMargin(geom.InLTRB(th.Space.SM, (h-layout.Size.H)/2, 0, 0))
	n.Add(label)
	n.SetHit(scene.HitSpec{Label: "tag"})
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
	n.SetHit(scene.HitSpec{Label: "avatar"})
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
	n.SetHit(scene.HitSpec{Label: "empty"})
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
			// SurfaceActive(#F0F0F0)：SurfaceSunken 与演示舞台同色，灰条会隐形。
			Fill: scene.Solid(th.Color.SurfaceActive), Radius: th.Radius.SM,
		}).SetSize2(rw, h))
	}
	n.SetHit(scene.HitSpec{Label: "skeleton"})
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
	// antd 步骤条：圆点一行、步骤间有连接线（已完成段主色、待办段灰），
	// 标签在圆点下方居中。
	n := scene.Flex(scene.Horizontal, 0)
	n.SetCrossAlign(scene.CrossStart)
	for i, label := range s.Labels {
		step := scene.Flex(scene.Vertical, th.Space.XS)
		step.SetCrossAlign(scene.CrossCenter)
		fill := th.Color.SurfaceActive
		fg := th.Color.TextMuted
		// antd wait 步的圆点是**空心描边圆**（透明底 + 灰边），完成/进行步
		// 是主色实心圆。此前 wait 用实心灰底，与 antd 官方对照不符。
		wait := i > s.Current
		if i <= s.Current {
			fill, fg = th.Color.Primary, th.Color.OnPrimary
		}
		circle := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: 10}).SetSize2(20, 20)
		if wait {
			circle = scene.Box(scene.BoxSpec{Border: scene.Stroke(th.Color.BorderStrong, 2), Radius: 10}).SetSize2(20, 20)
		}
		circle.SetHit(scene.HitSpec{Label: "steps-circle"}) // 被动锚：步骤圆点断言
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
		if i < len(s.Labels)-1 {
			line := th.Color.Border
			if i < s.Current {
				line = th.Color.Primary
			}
			n.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(line), Radius: 1}).
				SetSize2(40, 2).SetMargin(geom.InLTRB(4, 9, 4, 0)))
		}
	}
	n.SetHit(scene.HitSpec{Label: "steps"})
	return n
}

// ── Segmented ──────────────────────────────────────────────────────

// Segmented 是分段选择器。选中段用 Primary 底色。
type Segmented struct {
	controlBase
	Options  []string
	Selected int
	OnChange func(int)

	// items 跨重建持有：分段项的 hover 态与手势实例必须稳定。
	items []*segmentItem
}

// NewSegmented 创建分段选择器。
func NewSegmented(options ...string) *Segmented { return &Segmented{Options: options} }

// Build 实现 ui.Widget。
func (s *Segmented) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(s)
	th := ctx.ThemeOr()
	if len(s.items) != len(s.Options) {
		s.items = make([]*segmentItem, len(s.Options))
		for i := range s.items {
			s.items[i] = &segmentItem{index: i, parent: s}
		}
	}
	children := make([]*scene.Node, 0, len(s.Options))
	for i, opt := range s.Options {
		s.items[i].label = opt
		children = append(children, s.items[i].Build(ctx))
	}
	row := scene.Flex(scene.Horizontal, 0, children...)
	// 轨道用 SurfaceActive（#F0F0F0，与 antd 分段容器一致）——
	// SurfaceSunken(#FAFAFA) 与演示舞台同色，轨道会隐形。
	n := scene.Box(scene.BoxSpec{
		Fill: scene.Solid(th.Color.SurfaceActive), Radius: th.Radius.Base,
	}, row)
	n.SetPadding(geom.InLTRB(2, 2, 2, 2))
	// 被动定位锚（HitDeferToChild 不影响交互）：容器轨道色断言用。
	n.SetHit(scene.HitSpec{Label: "segmented"})
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
	// antd Segmented：灰轨道上的选中项是**白色滑块**（带轻投影）+ 深色
	// 文字，不是主色实心块；未选项 hover 轻微加深（colorFill 系反馈）。
	fill := geom.Color{}
	fg := th.Color.Text
	shadow := (*scene.ShadowSpec)(nil)
	if selected {
		fill = th.Color.Surface
		shadow = &scene.ShadowSpec{Dx: 0, Dy: 1, Blur: 3, Color: th.Color.Shadow}
	} else if s.hover {
		fill = th.Color.FillHover
	}
	font := text.DefaultFont(th.Font.SizeSM)
	layout := text.LayoutText(s.label, font, text.Options{})
	w := layout.Size.W + th.Space.LG
	const h = float32(26)
	item := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: th.Radius.SM, Shadow: shadow}).SetSize2(w, h)
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
	Open     int    // 单开（默认）：展开的面板下标，-1=全收
	Multiple bool   // 多开模式（antd 默认）：Multiple 模式下各面板独立开关
	OpenSet  []bool // Multiple 模式下各面板的展开状态（与 Panels 对齐）
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
	// Multiple 模式：OpenSet 与 Panels 对齐（Panels 可变长）。
	if c.Multiple && len(c.OpenSet) != len(c.Panels) {
		c.OpenSet = make([]bool, len(c.Panels))
	}
	col := scene.Flex(scene.Vertical, th.Space.XS)
	// 面板通栏铺满父级宽度（antd 折叠面板是全宽边框容器，不是紧凑按钮盒）。
	col.SetCrossAlign(scene.CrossStretch)
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
	if c.parent.Multiple {
		open = c.parent.OpenSet[c.index]
	}
	// antd 折叠面板：面板=通栏边框容器；表头一行"标题居左、箭头居右"
	// （SpaceBetween，宽度随父级自适应，箭头不再写死 x=320 飘在面板外）。
	layout := text.LayoutText(panel.Title, text.DefaultFont(th.Font.SizeBase), text.Options{})
	label := scene.Text(scene.TextSpec{Src: panel.Title, Layout: layout, Color: th.Color.Text})
	label.SetMargin(geom.InLTRB(th.Space.MD, 0, 0, 0))
	arrow := "▸"
	if open {
		arrow = "▾"
	}
	al := text.LayoutText(arrow, text.DefaultFont(th.Font.SizeSM), text.Options{})
	at := scene.Text(scene.TextSpec{Src: arrow, Layout: al, Color: th.Color.TextMuted})
	at.SetMargin(geom.InLTRB(0, 0, th.Space.MD, 0))
	header := scene.Flex(scene.Horizontal, 0, label, at)
	header.SetCrossAlign(scene.CrossCenter).SetMainAlign(scene.MainSpaceBetween)
	header.SetHeight(36)
	header.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "collapse", Cursor: scene.CursorPointer})
	parent, index := c.parent, c.index
	hnd := c.tapHandler(ctx, false, func() {
		if parent.Multiple {
			parent.OpenSet[index] = !parent.OpenSet[index]
			// 兼容 Open 单开语义：多开模式下把最后一个展开项同步给 Open。
			parent.Open = -1
			for i, v := range parent.OpenSet {
				if v {
					parent.Open = i
				}
			}
			parent.MarkDirty()
			if parent.OnChange != nil {
				parent.OnChange(parent.Open)
			}
			return
		}
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
	// 内容区：表头下方带分隔线与内边距（antd 形态）。
	rootChildren := []*scene.Node{header}
	if open && panel.Content != nil {
		wrap := scene.Box(scene.BoxSpec{})
		wrap.SetPadding(geom.InLTRB(th.Space.MD, th.Space.SM, th.Space.MD, th.Space.MD))
		wrap.Add(buildChild(ctx, panel.Content))
		rootChildren = append(rootChildren, wrap)
	}
	root := scene.Flex(scene.Vertical, 0, rootChildren...)
	root.SetCrossAlign(scene.CrossStretch)
	// 定位 + 紧约束：内容根被拉满面板宽度（表头 SpaceBetween 才有全宽）。
	root.SetPositionKeepSize(false)
	root.SetStackPosition(geom.Pt(0, 0))
	panelBox := scene.Box(scene.BoxSpec{
		Fill: scene.Solid(th.Color.Surface), Border: scene.Stroke(th.Color.Border, 1), Radius: th.Radius.Base,
	}, root)
	// 被动定位锚（HitDeferToChild 不影响交互）：面板级内容存在性断言用
	// （表头锚的子树不含内容文字——内容挂在面板容器下）。
	panelBox.SetHit(scene.HitSpec{Label: "collapse-panel"})
	return panelBox
}
