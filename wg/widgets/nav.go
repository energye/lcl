package widgets

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── Tabs ───────────────────────────────────────────────────────────

// Tab 是一页签。Content 为 nil 时只渲染标签条（内容由外部按 Active 自己切）。
type Tab struct {
	Key     ui.Key
	Label   string
	Content ui.Widget
}

// Tabs 是标签页。它是有状态组件：Active 改变后由 MarkDirty 触发重建。
type Tabs struct {
	ui.Base
	Items    []Tab
	Active   int
	OnChange func(int)

	// headers 跨重建持有：页签的 hover 态与手势实例必须稳定，
	// 每帧新建会让 hover 高亮闪烁、按住期间的重建丢失手势（--audit 实测）。
	headers []*tabHeader
	node    *scene.Node
}

// NewTabs 创建标签页。
func NewTabs(items ...Tab) *Tabs { return &Tabs{Items: items} }

// Node 返回最近一次构建的根节点（诊断/自动化点击用）。
func (t *Tabs) Node() *scene.Node { return t.node }

// Build 实现 ui.Widget。
func (t *Tabs) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(t)
	th := ctx.ThemeOr()
	if len(t.Items) == 0 {
		return scene.Empty()
	}
	if t.Active < 0 {
		t.Active = 0
	}
	if t.Active >= len(t.Items) {
		t.Active = len(t.Items) - 1
	}
	m := ui.MetricsFor(th, ui.SizeMedium)
	if len(t.headers) != len(t.Items) {
		t.headers = make([]*tabHeader, len(t.Items))
		for i := range t.headers {
			t.headers[i] = &tabHeader{index: i, tabs: t}
		}
	}
	for i, item := range t.Items {
		t.headers[i].label = item.Label
		t.headers[i].metrics = m
	}
	headers := make([]ui.Widget, len(t.headers))
	for i, hdr := range t.headers {
		headers[i] = hdr
	}
	bar := scene.Flex(scene.Horizontal, th.Space.SM, buildChildren(ctx, headers)...)
	col := scene.Flex(scene.Vertical, th.Space.MD, bar)
	if c := buildChild(ctx, t.Items[t.Active].Content); c != nil {
		col.Add(c)
	}
	t.node = col
	return col
}

// tabHeader 是单个页签（内部组件，不导出）。
type tabHeader struct {
	controlBase
	index   int
	label   string
	style   ui.Style
	metrics ui.Metrics
	tabs    *Tabs
}

func (h *tabHeader) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	font := text.DefaultFont(h.metrics.Font)
	layout := text.LayoutText(h.label, font, text.Options{})
	w := layout.Size.W + h.metrics.Padding.Min.X + h.metrics.Padding.Max.X
	// antd line 型页签：无底色按钮盒，文字色区分选中/悬停，选中带底部
	// 下划线指示条（此前是实心主色圆角块，与官方 line 型不符）。
	n := scene.Box(scene.BoxSpec{}).SetSize2(w, h.metrics.Height)
	col := th.Color.Text
	if h.index == h.tabs.Active {
		col = th.Color.Primary
	} else if h.hover {
		col = th.Color.PrimaryHover
	}
	label := scene.Text(scene.TextSpec{Src: h.label, Layout: layout, Color: col})
	label.SetMargin(geom.InLTRB(h.metrics.Padding.Min.X, (h.metrics.Height-layout.Size.H)/2, 0, 0))
	n.Add(label)
	if h.index == h.tabs.Active {
		ind := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary)}).SetSize2(w, 2)
		ind.SetPositionKeepSize(true).SetStackPosition(geom.Pt(0, h.metrics.Height-2))
		n.Add(ind)
	}
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "tab", Cursor: scene.CursorPointer})

	tabs, index := h.tabs, h.index
	hd := h.tapHandler(ctx, false, func() {
		if tabs.Active == index {
			return
		}
		tabs.Active = index
		tabs.MarkDirty()
		if tabs.OnChange != nil {
			tabs.OnChange(index)
		}
	})
	h.attach(ctx, n, hd, true)
	return n
}

// ── Tooltip ────────────────────────────────────────────────────────

// Tooltip 给子组件挂一个悬停提示。
//
// 浮层的定位/关闭/z 序全部交给 overlay.Host；这里只负责
// "何时显示、显示什么"，不在组件里写全局鼠标监听或计时器。
type Tooltip struct {
	Child     ui.Widget
	Text      string
	Placement overlay.Placement
	handle    overlay.Handle
}

// Build 实现 ui.Widget。
func (t *Tooltip) Build(ctx *ui.BuildContext) *scene.Node {
	child := buildChild(ctx, t.Child)
	if child == nil {
		return scene.Empty()
	}
	// 合并而不是覆盖：子组件（比如 Button）已经注册了指针处理器。
	h := input.Handler{}
	if ctx.Input != nil {
		if old, ok := ctx.Input.Handler(child); ok {
			h = old
		}
	}
	show := func() { t.show(ctx, child) }
	hide := func() { t.hide(ctx) }
	oldEnter, oldLeave := h.Enter, h.Leave
	h.Enter = func(e *input.PointerEvent) {
		if oldEnter != nil {
			oldEnter(e)
		}
		show()
	}
	h.Leave = func(e *input.PointerEvent) {
		hide()
		if oldLeave != nil {
			oldLeave(e)
		}
	}
	if ctx.Input != nil {
		ctx.Input.Set(child, h)
	}
	return child
}

func (t *Tooltip) show(ctx *ui.BuildContext, anchor *scene.Node) {
	if ctx.Overlay == nil || t.handle != 0 || t.Text == "" {
		return
	}
	th := ctx.ThemeOr()
	text := t.Text
	placement := t.Placement
	t.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandTooltip,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: placement,
		Size:      geom.Size{W: float32(len([]rune(text)))*13 + 16, H: 30},
		// 不用 OutsideClick：那会让"先悬停再点按钮"的第二次点击被
		// 关闭浮层吃掉（浮层吞掉的那次点击不下传，按钮看起来失灵）。
		// Tooltip 只由 Enter/Leave 管，永远不参与点击路由。
		Dismiss: overlay.Dismiss(0),
		Tag:     "tooltip",
		OnClose: func() { t.handle = 0 },
		Build: func() *scene.Node {
			m := ui.MetricsFor(th, ui.SizeSmall)
			layout := textLayout(text, m.Font)
			n := scene.Box(scene.BoxSpec{
				Fill:   scene.Solid(th.Color.Surface),
				Border: scene.Stroke(th.Color.Border, 1),
				Radius: th.Radius.Base,
			}).SetSize2(layout.Size.W+16, 30)
			label := scene.Text(scene.TextSpec{Src: text, Layout: layout, Color: th.Color.Text})
			label.SetMargin(geom.InLTRB(8, (30-layout.Size.H)/2, 0, 0))
			n.Add(label)
			n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "tooltip"})
			return n
		},
	})
}

func (t *Tooltip) hide(ctx *ui.BuildContext) {
	if ctx.Overlay == nil || t.handle == 0 {
		return
	}
	ctx.Overlay.Close(t.handle)
	t.handle = 0
}

// textLayout 复用了 Label 的排版参数，但避免在这里引入 ui.Widget 递归。
func textLayout(s string, size float32) *text.Layout {
	return text.LayoutText(s, text.DefaultFont(size), text.Options{})
}
