package widgets

import (
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// ListItem 是列表项。Key 决定复用：同 Key 的节点保留滚动位置/局部缓存。
type ListItem struct {
	Key    ui.Key
	Widget ui.Widget
}

// List 是垂直列表。
//
// 它持有上一次 Build 的节点与键，因此是**有状态组件**（嵌入 ui.Base）。
// 不做虚拟滚动——P1 的列表只保证键控复用；虚拟滚动是 M4 的 P2 范围。
type List struct {
	ui.Base
	Items []ListItem
	Gap   float32
	// PageSize 分页（antd pagination）：>0 时每页渲染 PageSize 项，
	// 底部 **Pagination 组件组合复用**（Table 分页同款通道）。
	PageSize int
	// EmptyText 空态文案（Items 空时显示，**Empty 组件组合复用**；
	// 空=不渲染空态块）。
	EmptyText string

	pageNo    int
	pagePager *Pagination

	nodes []*scene.Node
	keys  []ui.Key
}

// NewList 创建列表。
func NewList(items ...ListItem) *List { return &List{Items: items} }

// Build 实现 ui.Widget。
//
// 键控复用的意义不只是省一次分配：输入框列表中间插一项时，若同 Key 的
// 节点被重建，正在输入的那个框会丢焦点、丢组合串。这里必须复用节点。
func (l *List) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(l)
	// 空态：Items 空且配置了 EmptyText——Empty 组件组合复用。
	if len(l.Items) == 0 && l.EmptyText != "" {
		return (&Empty{Text: l.EmptyText, Width: 300}).Build(ctx)
	}
	items := l.Items
	// 分页切片：只渲染当前页项（KeyedList 键控复用照常生效）。
	if l.PageSize > 0 && len(items) > l.PageSize {
		maxPage := (len(items) + l.PageSize - 1) / l.PageSize
		if l.pageNo < 1 {
			l.pageNo = 1
		}
		if l.pageNo > maxPage {
			l.pageNo = maxPage
		}
		lo := (l.pageNo - 1) * l.PageSize
		hi := lo + l.PageSize
		if hi > len(items) {
			hi = len(items)
		}
		items = items[lo:hi]
	}
	nodes, keys, _ := ui.KeyedList(items,
		func(it ListItem) ui.Key { return it.Key },
		l.nodes, l.keys,
		func(it ListItem, _ int, reuse *scene.Node) *scene.Node {
			if reuse != nil {
				return reuse
			}
			if it.Widget == nil {
				return scene.Empty()
			}
			return it.Widget.Build(ctx)
		})
	l.nodes, l.keys = nodes, keys
	n := scene.Flex(scene.Vertical, l.Gap, nodes...)
	// 底部分页器行（Pagination 组件组合复用；Total=总条目数语义）。
	if l.PageSize > 0 && len(l.Items) > l.PageSize {
		if l.pagePager == nil {
			l.pagePager = &Pagination{}
		}
		l.pagePager.PageSize = l.PageSize
		l.pagePager.Total = len(l.Items)
		l.pagePager.ShowTotal = true
		l.pagePager.Page = l.pageNo
		l.pagePager.OnChange = func(pg int) {
			l.pageNo = pg
			l.MarkDirty()
		}
		pager := l.pagePager.Build(ctx)
		pagerRow := scene.Flex(scene.Horizontal, 0).SetMainAlign(scene.MainEnd)
		pagerRow.Add(pager)
		wrap := scene.Flex(scene.Vertical, l.Gap, n)
		wrap.Add(pagerRow)
		wrap.SetHit(scene.HitSpec{Label: "list"})
		return wrap
	}
	n.SetHit(scene.HitSpec{Label: "list"})
	return n
}

// Nodes 返回当前列表项节点（诊断用，只读语义：不要修改返回值）。
func (l *List) Nodes() []*scene.Node { return l.nodes }
