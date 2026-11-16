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
	nodes, keys, _ := ui.KeyedList(l.Items,
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
	return n
}

// Nodes 返回当前列表项节点（诊断用，只读语义：不要修改返回值）。
func (l *List) Nodes() []*scene.Node { return l.nodes }
