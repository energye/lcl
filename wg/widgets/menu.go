package widgets

// ── Menu（M11 导航）────────────────────────────────────────────────
//
// antd：<Menu mode="horizontal|vertical" selectedKeys onClick>
// 已对齐：横向/纵向双模式、选中态（横向底部指示条/纵向左侧指示条+
// 浅主色底）、图标（Icon 组件组合复用）、点击回调。
// 横向 Popup 子菜单已对齐（M9 尾轮）：顶级项有 Children → **hover
// 展开** BandPopup（Dropdown 同款延迟隐藏 150ms——鼠标移向浮层的
// Leave 间隙不闪断，浮层项 Enter 取消定时器）；子项 down 即执行
// （浮层 tap 配对不可靠结论）选中+回调+关面板。
// 键盘导航已对齐（第一百二十轮，antd roving tabindex）：纵向模式
// 整菜单一个焦点节点，↓↑ 在**可见项**间移动焦点、→ 展开/← 收起、
// Enter 选中（父项=切换展开）；焦点项 InfoBg 浅底（antd focus 视觉
// 同选中底色系但指示条不动——指示条只跟 Selected）。点击项自请焦点
// 并跟随 FocusedKey。
// 未实现（如实标注）：横向模式键盘导航（↓ 展开子面板）、分组
// （group）、内联折叠、多层嵌套子菜单（横向子面板只渲染一层）。

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// MenuMode 是菜单方向（antd Menu.mode）。
type MenuMode uint8

const (
	// MenuHorizontal 横向（顶部导航）：选中底部主色指示条。
	MenuHorizontal MenuMode = iota
	// MenuVertical 纵向（侧边导航）：选中左侧 3px 主色指示条 + 浅主色底。
	MenuVertical
)

// MenuItem 是菜单项。
type MenuItem struct {
	Key   string
	Label string
	// Icon 左侧图标名（Icon 组件组合复用；空=无图标）。
	Icon string
	// Children 子项（antd SubMenu）：非空时为父项——纵向模式**内联展开**
	// （缩进渲染子项+展开箭头），横向模式点击弹出子菜单浮层。
	Children []*MenuItem
	// Group 分组标题（antd itemGroup）：Label 渲染为置灰标题行，
	// 不可点、不参与选中。
	Group bool
}

// Menu 是导航菜单（扁平项；子菜单浮层为后续轮）。
// 有状态组件：Selected 变化由 MarkDirty 触发重建。
type Menu struct {
	controlBase
	Items    []*MenuItem
	Selected string // 选中项 Key
	Mode     MenuMode
	// Width 纵向模式宽度（默认 200；横向模式自适应内容）。
	Width   float32
	OnClick func(key string)

	// OpenKeys 展开的子菜单 Key 集合（纵向内联展开状态）。
	OpenKeys map[string]bool

	// FocusedKey 键盘焦点项 Key（roving tabindex：↓↑ 移动、点击跟随）；
	// 空=无键盘焦点（焦点视觉不渲染）。
	FocusedKey string

	// items 跨重建持有：页签的 hover 态与手势实例必须稳定（教训⑨）。
	items []*menuItem
	node  *scene.Node
}

// NewMenu 创建菜单。
func NewMenu(items ...*MenuItem) *Menu {
	return &Menu{Items: items, Width: 200}
}

// menuItem 是单个菜单项（内部组件，实例池跨重建）。
type menuItem struct {
	controlBase
	index  int
	key    string
	label  string
	icon   string
	depth  int
	hover  bool
	parent *Menu
	// 横向 Popup 子菜单：hasKids=有子项；kids 递归数据；hideTimer
	// 延迟隐藏（Dropdown 同款）；subHandles 各级面板 handle。
	hasKids   bool
	kids      []menuKid
	hideTimer *time.Timer
	// subAnchor 子面板锚（最近构建的行节点）。
	subAnchor  *scene.Node
	subHandles map[int]overlay.Handle
}

// menuKid 子菜单项数据（横向浮层渲染，**递归支持多层嵌套**）。
type menuKid struct {
	key   string
	label string
	kids  []menuKid
}

// kidsFrom 递归拷贝 MenuItem.Children 为浮层渲染数据。
func kidsFrom(items []*MenuItem) []menuKid {
	out := make([]menuKid, 0, len(items))
	for _, it := range items {
		out = append(out, menuKid{key: it.Key, label: it.Label, kids: kidsFrom(it.Children)})
	}
	return out
}

// navItem 是键盘导航的可见项快照（key+是否父项）。
type navItem struct {
	key    string
	parent bool // 有子项（→ 展开 / Enter 切换展开）
}

// navItems 收集当前**可见**的可导航项（分组标题跳过、未展开父项的
// 子项跳过——antd 键盘导航只在可见项间移动）。
func (m *Menu) navItems() []navItem {
	var out []navItem
	var walk func(items []*MenuItem)
	walk = func(items []*MenuItem) {
		for _, it := range items {
			if it.Group {
				continue
			}
			out = append(out, navItem{key: it.Key, parent: len(it.Children) > 0})
			if len(it.Children) > 0 && m.OpenKeys[it.Key] {
				walk(it.Children)
			}
		}
	}
	walk(m.Items)
	return out
}

// attachKeyboard 给纵向菜单挂焦点节点与按键处理（roving tabindex）。
// 整菜单一个 FocusNode：焦点在菜单整体上，FocusedKey 标记项内位置
// （antd Menu 同构——焦点框在菜单容器，导航键在项间移动）。
func (m *Menu) attachKeyboard(ctx *ui.BuildContext, n *scene.Node) {
	h := input.Handler{
		Key: func(ev *input.KeyEvent) bool {
			if ev.Kind == input.KeyUp {
				return false
			}
			if m.handleNavKey(ev, ctx) {
				return true
			}
			return ctx.Focus != nil && ctx.Focus.HandleKey(ev)
		},
	}
	m.attach(ctx, n, h, true)
}

// handleNavKey 处理导航键：↓↑ 移动焦点、→ 展开、← 收起、Enter 选中。
// 返回是否消费。
func (m *Menu) handleNavKey(ev *input.KeyEvent, ctx *ui.BuildContext) bool {
	nav := m.navItems()
	if len(nav) == 0 {
		return false
	}
	idx := -1
	for i, it := range nav {
		if it.key == m.FocusedKey {
			idx = i
			break
		}
	}
	switch ev.Code {
	case input.CodeDown:
		if idx < 0 {
			m.FocusedKey = nav[0].key
		} else if idx < len(nav)-1 {
			m.FocusedKey = nav[idx+1].key
		}
	case input.CodeUp:
		if idx < 0 {
			m.FocusedKey = nav[0].key
		} else if idx > 0 {
			m.FocusedKey = nav[idx-1].key
		}
	case input.CodeRight:
		if idx < 0 {
			return false
		}
		if nav[idx].parent && !m.OpenKeys[nav[idx].key] {
			if m.OpenKeys == nil {
				m.OpenKeys = make(map[string]bool)
			}
			m.OpenKeys[nav[idx].key] = true
		} else if nav[idx].parent && m.OpenKeys[nav[idx].key] {
			// 已展开：→ 进首个子项（antd SubMenu 键盘语义）。
			m.FocusedKey = m.firstChildKey(nav[idx].key)
		} else {
			return false
		}
	case input.CodeLeft:
		if idx < 0 {
			return false
		}
		if m.closeParentOf(nav[idx].key) {
			// 收起后焦点回父项（antd 语义）。
			m.FocusedKey = m.parentKeyOf(nav[idx].key)
		} else {
			return false
		}
	case input.CodeEnter:
		if idx < 0 {
			return false
		}
		if nav[idx].parent {
			if m.OpenKeys == nil {
				m.OpenKeys = make(map[string]bool)
			}
			m.OpenKeys[nav[idx].key] = !m.OpenKeys[nav[idx].key]
		} else {
			if m.Selected == nav[idx].key {
				return true
			}
			m.Selected = nav[idx].key
			if m.OnClick != nil {
				m.OnClick(nav[idx].key)
			}
		}
	default:
		return false
	}
	m.MarkDirty()
	return true
}

// firstChildKey 返回父项 Key 的首个子项 Key（展开状态下子项可见）。
func (m *Menu) firstChildKey(parentKey string) string {
	var walk func(items []*MenuItem) string
	walk = func(items []*MenuItem) string {
		for _, it := range items {
			if it.Key == parentKey {
				if len(it.Children) > 0 && !it.Children[0].Group {
					return it.Children[0].Key
				}
				return ""
			}
			if r := walk(it.Children); r != "" {
				return r
			}
		}
		return ""
	}
	return walk(m.Items)
}

// parentKeyOf 返回子项的父项 Key（根层返回 ""）。
func (m *Menu) parentKeyOf(key string) string {
	var walk func(items []*MenuItem) string
	walk = func(items []*MenuItem) string {
		for _, it := range items {
			for _, c := range it.Children {
				if c.Key == key {
					return it.Key
				}
			}
			if r := walk(it.Children); r != "" {
				return r
			}
		}
		return ""
	}
	return walk(m.Items)
}

// closeParentOf 收起包含 key 的最近父项；返回是否收起动作发生。
func (m *Menu) closeParentOf(key string) bool {
	p := m.parentKeyOf(key)
	if p == "" || !m.OpenKeys[p] {
		return false
	}
	m.OpenKeys[p] = false
	return true
}

// Build 实现 ui.Widget。
func (m *Menu) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(m)
	if m.Selected == "" && len(m.Items) > 0 {
		m.Selected = m.Items[0].Key
	}
	if len(m.items) != len(m.Items) {
		m.items = make([]*menuItem, len(m.Items))
		for i := range m.items {
			m.items[i] = &menuItem{index: i, parent: m}
		}
	}
	for i, it := range m.Items {
		m.items[i].key = it.Key
		m.items[i].label = it.Label
		m.items[i].icon = it.Icon
		// 横向 Popup 子菜单：子项数据递归拷贝（多层嵌套）。
		mi := m.items[i]
		mi.hasKids = len(it.Children) > 0
		mi.kids = kidsFrom(it.Children)
	}
	if m.Mode == MenuVertical {
		n := scene.Flex(scene.Vertical, 0)
		n.SetWidth(m.Width)
		m.buildVerticalItems(ctx, n, m.Items, 0)
		n.SetHit(scene.HitSpec{Label: "menu"})
		m.node = n
		m.attachKeyboard(ctx, n)
		return n
	}
	children := make([]*scene.Node, 0, len(m.items))
	for _, mi := range m.items {
		children = append(children, mi.Build(ctx))
	}
	n := scene.Flex(scene.Horizontal, 0, children...)
	n.SetHit(scene.HitSpec{Label: "menu"})
	m.node = n
	return n
}

// buildVerticalItems 递归构建纵向菜单项（含分组标题/子菜单内联展开）。
func (m *Menu) buildVerticalItems(ctx *ui.BuildContext, col *scene.Node, items []*MenuItem, depth int) {
	for _, it := range items {
		if it.Group {
			// 分组标题行：置灰小字、不可点（伪 MenuItem 走标题渲染）。
			mi := &menuItem{key: it.Key, label: it.Label, parent: m, index: len(m.items)}
			col.Add(mi.buildGroupTitle(ctx, depth))
			m.buildVerticalItems(ctx, col, it.Children, depth)
			continue
		}
		mi := &menuItem{key: it.Key, label: it.Label, icon: it.Icon, parent: m, index: len(m.items), depth: depth}
		col.Add(mi.Build(ctx))
		if len(it.Children) > 0 {
			open := m.OpenKeys[it.Key]
			// 展开箭头（▸/▾）替代 icon 位：父项渲染箭头提示。
			if open {
				m.buildVerticalItems(ctx, col, it.Children, depth+1)
			}
		}
	}
}

// buildGroupTitle 构建分组标题行：置灰小字、不可点、缩进同子项。
func (mi *menuItem) buildGroupTitle(ctx *ui.BuildContext, depth int) *scene.Node {
	th := ctx.ThemeOr()
	row := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	row.SetHeight(32)
	font := text.DefaultFont(th.Font.SizeXS)
	layout := text.LayoutText(mi.label, font, text.Options{})
	t := scene.Text(scene.TextSpec{Src: mi.label, Layout: layout, Color: th.Color.TextFaint})
	t.SetMargin(geom.InLTRB(float32(depth)*14+12, (32-layout.Size.H)/2, 0, 0))
	row.Add(t)
	return row
}

// Build 实现 ui.Widget（菜单项）。
//
// 横向：无底色文字页签，选中主色文字+底部 2px 指示条（antd 横向
// 菜单形态）；纵向：整行，选中左侧 3px 主色指示条+浅主色底
// （colorPrimaryBg 系），行高 40。
func (mi *menuItem) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	m := mi.parent
	// active：本项被选中，或（横向 SubMenu 语义）其子项被选中——
	// antd 选中子项后父 SubMenu 同步高亮。
	active := m.Selected == mi.key
	if !active && m.Mode == MenuHorizontal {
		for _, kid := range mi.kids {
			if kid.key == m.Selected {
				active = true
				break
			}
		}
	}
	font := text.DefaultFont(th.Font.SizeBase)
	layout := text.LayoutText(mi.label, font, text.Options{})

	if m.Mode == MenuVertical {
		row := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
		row.SetHeight(40)
		// 焦点项（antd roving focus 视觉）：浅主色底——与选中同色系，
		// 但左指示条不动（指示条只跟 Selected）；非焦点项才画 hover 底。
		focused := m.FocusedKey == mi.key
		bg := geom.Color{}
		if active || focused {
			bg = th.Color.InfoBg // antd 选中/焦点浅主色底
		} else if mi.hover {
			bg = th.Color.FillHover
		}
		row.SetBox(scene.BoxSpec{Fill: scene.Solid(bg)})
		// 缩进（子项 depth>0）。
		if mi.depth > 0 {
			row.Add(scene.Box(scene.BoxSpec{}).SetWidth(float32(mi.depth) * 14).SetHeight(1))
		}
		if active {
			ind := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary)}).SetWidth(3).SetHeight(24)
			row.Add(ind)
		}
		if mi.icon != "" {
			ic := (&Icon{Name: mi.icon, Size: 16, Color: func() geom.Color {
				if active {
					return th.Color.Primary
				}
				return th.Color.Text
			}()}).Build(ctx)
			row.Add(ic)
		}
		// 展开箭头（父项）：▾ 展开 / ▸ 收起。
		isParent := false
		arrow := ""
		for _, it := range m.Items {
			if it.Key == mi.key && len(it.Children) > 0 {
				isParent = true
				if m.OpenKeys[mi.key] {
					arrow = "▾"
				} else {
					arrow = "▸"
				}
			}
		}
		if isParent {
			al := text.LayoutText(arrow, text.DefaultFont(th.Font.SizeXS), text.Options{})
			row.Add(scene.Text(scene.TextSpec{Src: arrow, Layout: al, Color: th.Color.TextMuted}))
		}
		lc := th.Color.Text
		if active {
			lc = th.Color.Primary
		}
		row.Add(scene.Text(scene.TextSpec{Src: mi.label, Layout: layout, Color: lc}))
		row.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "menu-item", Cursor: scene.CursorPointer})
		// hover 原地化（TimePicker SetBox 推广）：hover 只影响 bg 底
		//（FillHover），active 底/indicator 不受 hover 干扰——选中项
		// 不叠 hover 底。回调捕获的 row/active 是本次构建的闭包值，
		// 选中态重建后新闭包自动跟上。
		mi.setHoverVisual(func(node *scene.Node, hover bool) {
			if active {
				return
			}
			if hover {
				node.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.FillHover)})
			} else {
				node.SetBox(scene.BoxSpec{})
			}
			node.MarkNeedsPaint()
		})
		// 点击跟随焦点（roving tabindex——antd 键盘鼠标焦点统一）+ 自请
		// 焦点（键盘注入先 click 聚焦——Input 同款链路）。
		h := mi.tapHandler(ctx, false, func() {
			m.FocusedKey = mi.key
			if isParent {
				// 父项点击=切换展开（antd inline SubMenu）。
				if m.OpenKeys == nil {
					m.OpenKeys = make(map[string]bool)
				}
				m.OpenKeys[mi.key] = !m.OpenKeys[mi.key]
				m.MarkDirty()
				return
			}
			m.Selected = mi.key
			m.MarkDirty()
			if m.OnClick != nil {
				m.OnClick(mi.key)
			}
		})
		oldDown := h.Pointer
		h.Pointer = func(ev *input.PointerEvent) bool {
			if ev != nil && ev.Down() && ctx.Focus != nil && m.focus != nil {
				ctx.Focus.Focus(m.focus)
			}
			return oldDown != nil && oldDown(ev)
		}
		mi.attach(ctx, row, h, false)
		return row
	}

	// 横向：无底色文字页签 + 选中底部指示条。
	w := layout.Size.W + th.Space.LG
	if mi.hasKids {
		// 子菜单箭头（antd 横向 SubMenu 下箭头）。
		w += 12 + th.Space.XS
	}
	n := scene.Box(scene.BoxSpec{}).SetSize2(w, 40)
	lc := th.Color.Text
	if active {
		lc = th.Color.Primary
	} else if mi.hover {
		lc = th.Color.PrimaryHover
	}
	label := scene.Text(scene.TextSpec{Src: mi.label, Layout: layout, Color: lc})
	label.SetMargin(geom.InLTRB(th.Space.SM, (40-layout.Size.H)/2, 0, 0))
	n.Add(label)
	// hover 原地化：文字色 Text↔PrimaryHover 原地刷（active 时文字恒
	// 主色不参与 hover——与纵向选中项纪律一致）。
	mi.setHoverVisual(func(node *scene.Node, hover bool) {
		if active {
			return
		}
		c := th.Color.Text
		if hover {
			c = th.Color.PrimaryHover
		}
		label.SetTextPaint(scene.TextSpec{Src: mi.label, Layout: layout, Color: c})
		label.MarkNeedsPaint()
	})
	if mi.icon != "" {
		ic := (&Icon{Name: mi.icon, Size: 16, Color: lc}).Build(ctx)
		ic.SetMargin(geom.InLTRB(0, (40-16)/2, th.Space.XS, 0))
		n.Add(ic)
	}
	if mi.hasKids {
		ic := (&Icon{Name: "down", Size: 11, Color: lc}).Build(ctx)
		ic.SetPositionKeepSize(true)
		ic.SetStackPosition(geom.Pt(w-12-4, (40-11)/2))
		n.Add(ic)
	}
	if active {
		ind := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary)}).SetSize2(w, 2)
		ind.SetPositionKeepSize(true).SetStackPosition(geom.Pt(0, 38))
		n.Add(ind)
	}
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "menu-item", Cursor: scene.CursorPointer})
	mi.subAnchor = n
	if mi.hasKids {
		// 横向 SubMenu：**hover 展开**（Dropdown 同款——Enter 立即弹、
		// Leave 延迟 150ms 关，浮层项 Enter 取消）；父项点击=开/关面板。
		h := mi.tapHandler(ctx, false, func() {
			if mi.handleAt(1) != 0 {
				mi.hideSubNow(ctx)
			} else {
				mi.showSub(ctx, n, th)
			}
		})
		oldEnter, oldLeave := h.Enter, h.Leave
		h.Enter = func(e *input.PointerEvent) {
			if oldEnter != nil {
				oldEnter(e)
			}
			mi.showSub(ctx, n, th)
		}
		h.Leave = func(e *input.PointerEvent) {
			if oldLeave != nil {
				oldLeave(e)
			}
			mi.hideSubDelayed(ctx)
		}
		mi.attach(ctx, n, h, false)
		return n
	}
	mi.attach(ctx, n, mi.tapHandler(ctx, false, func() {
		if m.Selected == mi.key {
			return
		}
		m.Selected = mi.key
		m.MarkDirty()
		if m.OnClick != nil {
			m.OnClick(mi.key)
		}
	}), false)
	return n
}

// showSub 弹出横向子菜单浮层（level 1；子级经 showSubKids 递归）。
func (mi *menuItem) showSub(ctx *ui.BuildContext, anchor *scene.Node, th *theme.Data) {
	mi.showSubKids(ctx, anchor, th, mi.kids, 1)
}

// showSubKids 弹出某级子面板（level 从 1 起；面板项有 kids → hover
// **递归展开右侧**——PlacementRightStart，全深度支持）。
func (mi *menuItem) showSubKids(ctx *ui.BuildContext, anchor *scene.Node, th *theme.Data, kids []menuKid, level int) {
	m := mi.parent
	if mi.hideTimer != nil {
		mi.hideTimer.Stop()
		mi.hideTimer = nil
	}
	if ctx == nil || ctx.Overlay == nil || len(kids) == 0 {
		return
	}
	// 切分支：关闭比 level 深的既有面板（保留本级）。
	mi.hideFrom(ctx, level)
	if mi.handleAt(level) != 0 {
		return
	}
	panel := scene.Flex(scene.Vertical, 0)
	panel.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.SM,
		Shadow: func() *scene.ShadowSpec {
			if len(th.Elevation) > 1 {
				e := th.Elevation[1]
				return &scene.ShadowSpec{Dx: e.Dx, Dy: e.Dy, Blur: e.Blur, Color: e.Color}
			}
			return nil
		}(),
	})
	panel.SetWidth(160)
	panel.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "menu-sub"})
	for _, kid := range kids {
		kid := kid
		active := m.Selected == kid.key
		row := scene.Box(scene.BoxSpec{})
		row.SetWidth(160)
		row.SetHeight(34)
		lc := th.Color.Text
		if active {
			lc = th.Color.Primary
		}
		l := text.LayoutText(kid.label, text.DefaultFont(th.Font.SizeBase), text.Options{})
		txt := scene.Text(scene.TextSpec{Src: kid.label, Layout: l, Color: lc})
		txt.SetMargin(geom.InLTRB(12, (34-l.Size.H)/2, 0, 0))
		row.Add(txt)
		if len(kid.kids) > 0 {
			ic := (&Icon{Name: "right", Size: 10, Color: th.Color.TextFaint}).Build(ctx)
			ic.SetPositionKeepSize(true)
			ic.SetStackPosition(geom.Pt(160-14, (34-10)/2))
			row.Add(ic)
		}
		row.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "menu-sub-item", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			// down 即执行（浮层 tap 配对不可靠结论复用）；Enter 取消
			// Leave 的延迟隐藏（Dropdown 同款）+ **有子级时递归展开右侧**。
			ctx.Input.Set(row, input.Handler{
				Enter: func(*input.PointerEvent) {
					if mi.hideTimer != nil {
						mi.hideTimer.Stop()
						mi.hideTimer = nil
					}
					if len(kid.kids) > 0 {
						mi.showSubKids(ctx, row, th, kid.kids, level+1)
					}
					row.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.FillHover)})
					row.MarkNeedsPaint()
				},
				Leave: func(*input.PointerEvent) {
					row.SetBox(scene.BoxSpec{})
					row.MarkNeedsPaint()
				},
				Pointer: func(ev *input.PointerEvent) bool {
					if !ev.Down() {
						return false
					}
					if len(kid.kids) > 0 {
						// 有子级：down 展开右侧（与 hover 一致）。
						mi.showSubKids(ctx, row, th, kid.kids, level+1)
						return true
					}
					m.Selected = kid.key
					m.MarkDirty()
					mi.hideSubNow(ctx)
					if m.OnClick != nil {
						m.OnClick(kid.key)
					}
					return true
				},
			})
		}
		panel.Add(row)
	}
	placement := overlay.PlacementBottomStart
	if level > 1 {
		placement = overlay.PlacementRightStart // 二级起右侧展开（antd 竖向子菜单位）
	}
	mi.setHandle(level, ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: placement,
		Size:      geom.Size{W: 160, H: 0}, // 高自适应
		Dismiss:   overlay.DismissOutsideClick | overlay.DismissEsc,
		Tag:       "menu-sub",
		OnClose: func() {
			mi.clearHandle(level)
		},
		Build: func() *scene.Node { return panel },
	}))
}

// handleAt 取某级面板 handle（0=未开）。
func (mi *menuItem) handleAt(level int) overlay.Handle {
	if mi.subHandles == nil {
		return 0
	}
	return mi.subHandles[level]
}

// setHandle 记录某级面板 handle（同时关闭更深级别——切分支语义）。
func (mi *menuItem) setHandle(level int, h overlay.Handle) {
	if mi.subHandles == nil {
		mi.subHandles = map[int]overlay.Handle{}
	}
	for lv := range mi.subHandles {
		if lv > level {
			delete(mi.subHandles, lv)
		}
	}
	mi.subHandles[level] = h
}

// clearHandle 清某级 handle 记录。
func (mi *menuItem) clearHandle(level int) {
	if mi.subHandles != nil {
		delete(mi.subHandles, level)
	}
}

// hideFrom 关闭 level 及更深的全部面板。
func (mi *menuItem) hideFrom(ctx *ui.BuildContext, level int) {
	if mi.subHandles == nil {
		return
	}
	for lv, h := range mi.subHandles {
		if lv >= level && h != 0 && ctx != nil && ctx.Overlay != nil {
			ctx.Overlay.Close(h)
		}
		if lv >= level {
			delete(mi.subHandles, lv)
		}
	}
}

// hideSubDelayed 延迟 150ms 关闭全部面板（鼠标移向浮层的 Leave 间隙
// 不闪断；面板项 Enter 取消）。
func (mi *menuItem) hideSubDelayed(ctx *ui.BuildContext) {
	if mi.hideTimer != nil {
		mi.hideTimer.Stop()
	}
	mi.hideTimer = time.AfterFunc(150*time.Millisecond, func() { mi.hideSubNow(ctx) })
}

// hideSubNow 立即关闭全部层级面板。
func (mi *menuItem) hideSubNow(ctx *ui.BuildContext) {
	if mi.hideTimer != nil {
		mi.hideTimer.Stop()
		mi.hideTimer = nil
	}
	mi.hideFrom(ctx, 1)
}
