package widgets

// ── TreeSelect 树选择（M12 高级录入）─────────────────────────────
//
// antd：<TreeSelect treeData value onChange placeholder>
// 第一轮：**纯组合复用**——Select 的浮层壳（触发框+BandPopup+
// PlacementBottomStart）+ **Tree 组件原样内嵌**（展开/收起/缩进/
// 选中/hover/行交互全部白拿，零重复实现——Flutter DropdownButton
// 内嵌内容树的同款组合法）。
// Tree.OnSelect → 选中 Key → 触发框显示节点 Label + OnChange +
// 关面板（antd TreeSelect 默认选后关；父节点点击仅展开，Tree 自己
// Toggle+OnSelect 同回调——isLeaf 过滤）。Tree 实例由 TreeSelect
// 持有（跨面板开合保留展开状态；跨重建持有——教训⑨）。
//
// 浮层树子树不随根重建（Entry.Build 只在 Show 时调一次）：
// 展开/收起的行数与面板高度变化 → **OnToggle 后 time.AfterFunc(0)
// 重开面板**（Dropdown hideTimer 同款手法——真窗 down 分发中途
// Close+Show 会在帧泵重入下丢新面板内容，探针实证）。
// 行点击用 Tree.TapDownExec（浮层内 tap 配对不可靠——
// Dropdown/DatePicker 轮结论）。
//
// 未实现（如实标注）：多选 tags、搜索过滤、异步加载、受控 open、
// 禁用项、maxCount、树形滚动容器限制（面板高度按行数自适应）。

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// TreeSelect 是树形下拉选择器（Select 壳 + Tree 内容的组合复用）。
type TreeSelect struct {
	controlBase
	Nodes       []TreeNode
	Value       string // 选中节点 Key
	Placeholder string
	Width       float32
	OnChange    func(key string)

	tree      *Tree
	handle    overlay.Handle
	fb        feedbackHost
	childNode *scene.Node
	anchor    *scene.Node
	panelOpen bool
	panelNode *scene.Node // 面板包装层（原地换内容用）
}

// NewTreeSelect 创建树选择器（默认占位"请选择"）。
func NewTreeSelect(nodes []TreeNode) *TreeSelect {
	return &TreeSelect{Nodes: nodes, Placeholder: "请选择"}
}

// findLabel 递归找 Key 对应的节点 Label（找不到返回 Key 本身）。
func (s *TreeSelect) findLabel(key string) string {
	var walk func(nodes []TreeNode) string
	walk = func(nodes []TreeNode) string {
		for _, n := range nodes {
			if n.Key == key {
				return n.Label
			}
			if l := walk(n.Children); l != "" {
				return l
			}
		}
		return ""
	}
	if l := walk(s.Nodes); l != "" {
		return l
	}
	return key
}

// isLeaf 判断 Key 是否叶子节点（Tree 点父节点也会触发 OnSelect——
// 展开与选中同回调，TreeSelect 仅叶子提交+关面板）。
func (s *TreeSelect) isLeaf(key string) bool {
	var walk func(nodes []TreeNode) (bool, bool)
	walk = func(nodes []TreeNode) (bool, bool) {
		for _, n := range nodes {
			if n.Key == key {
				return true, len(n.Children) == 0
			}
			if f, leaf := walk(n.Children); f {
				return f, leaf
			}
		}
		return false, false
	}
	f, leaf := walk(s.Nodes)
	return f && leaf
}

// Build 实现 ui.Widget：触发框（Select 同款形态）。
func (s *TreeSelect) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(s)
	th := ctx.ThemeOr()
	w := s.Width
	if w <= 0 {
		w = 220
	}
	h := ui.MetricsFor(th, ui.SizeMedium).Height
	display := s.Placeholder
	txtColor := th.Color.TextPlaceholder
	if s.Value != "" {
		display, txtColor = s.findLabel(s.Value), th.Color.Text
	}
	font := text.DefaultFont(th.Font.SizeBase)
	layout := text.LayoutText(display, font, text.Options{})

	n := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	}).SetSize2(w, h)
	label := scene.Text(scene.TextSpec{Src: display, Layout: layout, Color: txtColor})
	label.SetMargin(geom.InLTRB(11, (h-layout.Size.H)/2, 0, 0))
	n.Add(label)
	ic := (&Icon{Name: "down", Size: 12, Color: th.Color.TextFaint}).Build(ctx)
	ic.SetPositionKeepSize(true)
	ic.SetStackPosition(geom.Pt(w-11-12, (h-12)/2))
	n.Add(ic)
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "tree-select", Cursor: scene.CursorPointer})

	if ctx.Input != nil {
		anchor := n
		s.anchor = n
		g := &ui.GestureDetector{OnTap: func(geom.Point) {
			if s.panelOpen {
				s.closePanel(ctx)
				return
			}
			s.openPanel(ctx, anchor, th)
		}}
		ctx.Input.Set(n, g.Handler())
	}
	s.childNode = n
	return n
}

// Show 程序化打开面板（锚=最近构建的触发框节点；测试/自动化用）。
func (s *TreeSelect) Show(ctx *ui.BuildContext) {
	if s.childNode == nil {
		return
	}
	s.openPanel(ctx, s.childNode, ctx.ThemeOr())
}

// IsOpen 返回面板是否打开。
func (s *TreeSelect) IsOpen() bool { return s.panelOpen }

// Close 关闭面板。
func (s *TreeSelect) Close() {
	if s.fb.host != nil {
		s.closePanel(nil)
	}
}

// openPanel 打开树面板（Tree 组件原样内嵌——组合复用）。
func (s *TreeSelect) openPanel(ctx *ui.BuildContext, anchor *scene.Node, th *theme.Data) {
	if s.handle != 0 {
		return
	}
	s.panelOpen = true
	if s.tree == nil {
		// Tree 实例跨面板开合持有：展开状态保留（跨重建持有——教训⑨）。
		s.tree = NewTree(s.Nodes)
	}
	s.tree.TapDownExec = true // 浮层内嵌：行点击 down 即执行（tap 配对不可靠）
	s.tree.Selected = s.Value
	s.tree.OnSelect = func(key string) {
			if !s.isLeaf(key) {
			// 父节点：Tree 已自行 Toggle 展开，面板保持。
			return
		}
		s.Value = key
		s.MarkDirty() // 触发框文本随根重建刷新
		// 关面板：ctx.Overlay 直关（DatePicker 结论：fb.host 链路可能为 nil）。
		if ctx.Overlay != nil && s.handle != 0 {
			ctx.Overlay.Close(s.handle)
		}
		s.handle = 0
		s.panelOpen = false
		if s.OnChange != nil {
			s.OnChange(key)
		}
	}
	// 浮层树子树不随根重建（Entry.Build 只在 Show 时调一次）：展开/收起
	// 的行数与面板高度变化 → **原地换面板内容**（RemoveAll+重建子树+
	// Overlay.Relayout 重算高度）。不走 Close+Show 重开：真窗 down 分发
	// 中途重开即使推迟到分发后（time.AfterFunc(0)）仍丢内容（探针实证）。
	s.tree.OnToggle = func(string, bool) {
		time.AfterFunc(0, func() { s.refreshPanel(ctx) })
	}
	w := s.Width
	if w <= 0 {
		w = 220
	}
	s.tree.Width = w - 2
	// Tree 根节点原样内嵌（不覆盖其自身 HitSpec——包一层带面板锚）。
	inner := s.tree.Build(ctx)
	panel := scene.Flex(scene.Vertical, 0)
	panel.Add(inner)
	panel.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "tree-select-panel"})
	s.panelNode = panel
	s.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: overlay.PlacementBottomStart,
		Size:      geom.Size{W: w, H: 0}, // 高自适应（按树行数）
		Dismiss:   overlay.DismissOutsideClick | overlay.DismissEsc,
		Tag:       "treeselect",
		OnClose: func() {
			s.handle = 0
			s.panelOpen = false
		},
		Build: func() *scene.Node { return panel },
	})
}

// refreshPanel 原地重建面板内的树子树（展开/收起后行数变化）。
// 不重开 Entry：内容节点同一批，Overlay.Relayout 重算面板高度。
func (s *TreeSelect) refreshPanel(ctx *ui.BuildContext) {
	if ctx == nil || ctx.Overlay == nil || s.panelNode == nil || s.tree == nil {
		return
	}
	s.panelNode.RemoveAll()
	s.panelNode.Add(s.tree.Build(ctx))
	s.panelNode.MarkNeedsLayout()
	ctx.Overlay.Relayout()
}

// closePanel 关闭面板。
func (s *TreeSelect) closePanel(ctx *ui.BuildContext) {
	if s.handle != 0 && s.fb.host != nil {
		s.fb.host.Close(s.handle)
	}
	s.handle = 0
	s.panelOpen = false
}
