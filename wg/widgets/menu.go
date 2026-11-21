package widgets

// ── Menu（M11 导航）────────────────────────────────────────────────
//
// antd：<Menu mode="horizontal|vertical" selectedKeys onClick>
// 已对齐：横向/纵向双模式、选中态（横向底部指示条/纵向左侧指示条+
// 浅主色底）、图标（Icon 组件组合复用）、点击回调。
// 未实现（如实标注）：子菜单浮层（Popup）、分组（group）、内联折叠。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
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
}

// Menu 是导航菜单（扁平项；子菜单浮层为后续轮）。
// 有状态组件：Selected 变化由 MarkDirty 触发重建。
type Menu struct {
	controlBase
	Items    []*MenuItem
	Selected string // 选中项 Key
	Mode     MenuMode
	// Width 纵向模式宽度（默认 200；横向模式自适应内容）。
	Width  float32
	OnClick func(key string)

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
	hover  bool
	parent *Menu
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
	}
	children := make([]*scene.Node, 0, len(m.items))
	for _, mi := range m.items {
		children = append(children, mi.Build(ctx))
	}
	var n *scene.Node
	if m.Mode == MenuVertical {
		n = scene.Flex(scene.Vertical, 0, children...)
		n.SetWidth(m.Width)
	} else {
		n = scene.Flex(scene.Horizontal, 0, children...)
	}
	n.SetHit(scene.HitSpec{Label: "menu"})
	m.node = n
	return n
}

// Build 实现 ui.Widget（菜单项）。
//
// 横向：无底色文字页签，选中主色文字+底部 2px 指示条（antd 横向
// 菜单形态）；纵向：整行，选中左侧 3px 主色指示条+浅主色底
// （colorPrimaryBg 系），行高 40。
func (mi *menuItem) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	m := mi.parent
	active := m.Selected == mi.key
	font := text.DefaultFont(th.Font.SizeBase)
	layout := text.LayoutText(mi.label, font, text.Options{})

	if m.Mode == MenuVertical {
		row := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
		row.SetHeight(40)
		bg := geom.Color{}
		if active {
			bg = th.Color.InfoBg // antd 选中浅主色底
		} else if mi.hover {
			bg = th.Color.FillHover
		}
		row.SetBox(scene.BoxSpec{Fill: scene.Solid(bg)})
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
		lc := th.Color.Text
		if active {
			lc = th.Color.Primary
		}
		row.Add(scene.Text(scene.TextSpec{Src: mi.label, Layout: layout, Color: lc}))
		row.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "menu-item", Cursor: scene.CursorPointer})
		mi.attach(ctx, row, mi.tapHandler(ctx, false, func() {
			m.Selected = mi.key
			m.MarkDirty()
			if m.OnClick != nil {
				m.OnClick(mi.key)
			}
		}), false)
		return row
	}

	// 横向：无底色文字页签 + 选中底部指示条。
	w := layout.Size.W + th.Space.LG
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
	if mi.icon != "" {
		ic := (&Icon{Name: mi.icon, Size: 16, Color: lc}).Build(ctx)
		ic.SetMargin(geom.InLTRB(0, (40-16)/2, th.Space.XS, 0))
		n.Add(ic)
	}
	if active {
		ind := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary)}).SetSize2(w, 2)
		ind.SetPositionKeepSize(true).SetStackPosition(geom.Pt(0, 38))
		n.Add(ind)
	}
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "menu-item", Cursor: scene.CursorPointer})
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
