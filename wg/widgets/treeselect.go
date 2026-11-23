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
// 多选 tags（antd multiple）：触发框渲染可删除 tag（Select 多选
// 同款视觉——灰底圆角+× 删除，按宽度换行、高度自适应），点叶子
// toggle 已选集合**不关面板**（antd 多选语义），Tree.SelectedSet
// 多行高亮；MaxTagCount 超出折叠 "+N"。tag × 与面板行点击同属
// down 分发内改浮层子树 → **AfterFunc(0) 延迟 refreshPanel**（
// Cascader 搜索轮同款教训）。
//
// 未实现（如实标注）：搜索过滤、异步加载、受控 open、
// 禁用项、树形滚动容器限制（面板高度按行数自适应）。

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
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
	Value       string   // 单选：选中节点 Key
	Values      []string // 多选语义（Multiple=true）
	Multiple    bool
	// MaxTagCount 触发框最多显示的 tag 数，超出折叠 "+N"（0=不限制；
	// antd maxTagCount）。
	MaxTagCount int
	Placeholder string
	Width       float32
	OnChange    func(key string)
	// OnValuesChange 多选值变化回调（antd onChange 在 multiple 下返回数组）。
	OnValuesChange func(keys []string)

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

// tagShown 多选触发框折叠计算：返回实际显示的 Key 列表与折叠数。
func (s *TreeSelect) tagShown() (shown []string, overflow int) {
	if s.MaxTagCount > 0 && len(s.Values) > s.MaxTagCount {
		return s.Values[:s.MaxTagCount], len(s.Values) - s.MaxTagCount
	}
	return s.Values, 0
}

// syncTreeSelection 把当前值同步到内嵌 Tree 的高亮集合。
func (s *TreeSelect) syncTreeSelection() {
	if s.tree == nil {
		return
	}
	if s.Multiple {
		s.tree.Selected = ""
		s.tree.SelectedSet = make(map[string]bool, len(s.Values))
		for _, v := range s.Values {
			s.tree.SelectedSet[v] = true
		}
		return
	}
	s.tree.SelectedSet = nil
	s.tree.Selected = s.Value
}

// removeString 返回去掉 key 的新切片（不修改原切片）。
func removeString(list []string, key string) []string {
	next := make([]string, 0, len(list))
	for _, v := range list {
		if v != key {
			next = append(next, v)
		}
	}
	return next
}

// toggleValue 多选：点面板叶子 toggle 已选集合（antd 多选点选不关
// 面板）。调用点在 down 分发内——浮层子树重建走 AfterFunc(0) 延迟。
func (s *TreeSelect) toggleValue(key string, ctx *ui.BuildContext) {
	removed := false
	for _, v := range s.Values {
		if v == key {
			removed = true
			break
		}
	}
	if removed {
		s.Values = removeString(s.Values, key)
	} else {
		s.Values = append(append(make([]string, 0, len(s.Values)+1), s.Values...), key)
	}
	s.syncTreeSelection()
	s.MarkDirty() // 触发框 tags 随根重建刷新
	if s.OnValuesChange != nil {
		s.OnValuesChange(s.Values)
	}
	// 树行高亮变化在浮层子树上——down 分发内不重建（真窗教训），
	// AfterFunc(0) 推迟到分发后原地刷新（Cascader 搜索轮同款）。
	time.AfterFunc(0, func() { s.refreshPanel(ctx) })
}

// removeValue 多选：tag × 删除一个已选 Key。× 的 handler 同属 down
// 分发内——浮层开着时 AfterFunc(0) 延迟同步树高亮。
func (s *TreeSelect) removeValue(key string, ctx *ui.BuildContext) {
	s.Values = removeString(s.Values, key)
	s.syncTreeSelection()
	s.MarkDirty()
	if s.OnValuesChange != nil {
		s.OnValuesChange(s.Values)
	}
	if s.panelNode != nil {
		time.AfterFunc(0, func() { s.refreshPanel(ctx) })
	}
}

// Build 实现 ui.Widget：触发框（Select 同款形态；多选 tags 换行）。
func (s *TreeSelect) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(s)
	th := ctx.ThemeOr()
	w := s.Width
	if w <= 0 {
		w = 220
	}
	var n *scene.Node
	if s.Multiple && len(s.Values) > 0 {
		n = s.buildMultiTrigger(ctx, th, w)
	} else {
		n = s.buildSingleTrigger(ctx, th, w)
	}
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

// buildSingleTrigger 单选触发框（原形态：文本+down 图标）。
func (s *TreeSelect) buildSingleTrigger(ctx *ui.BuildContext, th *theme.Data, w float32) *scene.Node {
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
	return n
}

// buildMultiTrigger 多选触发框：可删除 tag 按宽度换行（Select 多选
// 同款视觉），MaxTagCount 超出折叠 "+N"，高度自适应。
func (s *TreeSelect) buildMultiTrigger(ctx *ui.BuildContext, th *theme.Data, w float32) *scene.Node {
	const (
		tagH      = float32(22)
		tagPad    = float32(6)
		tagGap    = float32(6)
		innerPad  = float32(6)
		arrowZone = float32(26)
	)
	n := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	avail := w - innerPad*2 - arrowZone
	x, y := innerPad, innerPad
	shown, overflow := s.tagShown()
	for _, key := range shown {
		label := s.findLabel(key)
		lv := text.LayoutText(label, text.DefaultFont(th.Font.SizeXS), text.Options{})
		tagW := lv.Size.W + tagPad*2 + 12 // 文本 + 左右内边距 + ×
		if x+tagW > innerPad+avail && x > innerPad {
			x = innerPad
			y += tagH + 4
		}
		tag := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.SurfaceSunken),
			Radius: th.Radius.SM,
		}).SetSize2(tagW, tagH)
		tag.SetPositionKeepSize(true).SetStackPosition(geom.Pt(x, y))
		tv := scene.Text(scene.TextSpec{Src: label, Layout: lv, Color: th.Color.TextMuted})
		tv.SetMargin(geom.InLTRB(tagPad, (tagH-lv.Size.H)/2, 0, 0))
		tag.Add(tv)
		xx := (&Icon{Name: "close", Size: 10, Color: th.Color.TextFaint}).Build(ctx)
		xx.SetPositionKeepSize(true).SetStackPosition(geom.Pt(tagW-14, (tagH-10)/2))
		xx.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "treeselect-tag-x", Cursor: scene.CursorPointer})
		removed, buildCtx := key, ctx
		if ctx.Input != nil {
			ctx.Input.Set(xx, input.Handler{
				Pointer: func(ev *input.PointerEvent) bool {
					if !ev.Down() {
						return false
					}
					s.removeValue(removed, buildCtx)
					return true
				},
			})
		}
		tag.Add(xx)
		tag.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "treeselect-tag", Cursor: scene.CursorPointer})
		n.Add(tag)
		x += tagW + tagGap
	}
	if overflow > 0 {
		chipText := fmt.Sprintf("+%d", overflow)
		cl := text.LayoutText(chipText, text.DefaultFont(th.Font.SizeXS), text.Options{})
		chipW := cl.Size.W + tagPad*2
		chip := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.SurfaceSunken),
			Radius: th.Radius.SM,
		}).SetSize2(chipW, tagH)
		chip.SetPositionKeepSize(true).SetStackPosition(geom.Pt(x, y))
		ct := scene.Text(scene.TextSpec{Src: chipText, Layout: cl, Color: th.Color.TextMuted})
		ct.SetMargin(geom.InLTRB(tagPad, (tagH-cl.Size.H)/2, 0, 0))
		chip.Add(ct)
		chip.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "treeselect-tag-more"})
		n.Add(chip)
	}
	h := y + tagH + innerPad
	if h < 34 {
		h = 34
	}
	n.SetSize2(w, h)
	ic := (&Icon{Name: "down", Size: 12, Color: th.Color.TextFaint}).Build(ctx)
	ic.SetPositionKeepSize(true)
	ic.SetStackPosition(geom.Pt(w-11-12, (h-12)/2))
	n.Add(ic)
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
	s.syncTreeSelection()
	s.tree.OnSelect = func(key string) {
		if !s.isLeaf(key) {
			// 父节点：Tree 已自行 Toggle 展开，面板保持。
			return
		}
		if s.Multiple {
			// 多选点选不关面板（antd 语义）；toggle 内部 AfterFunc 延迟
			// 刷新浮层子树高亮。
			s.toggleValue(key, ctx)
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
