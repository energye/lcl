package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// pagedList 包级持有：翻页 pageNo 状态跨重建（闭包内新建会归页，教训⑨）。
var pagedList = func() *widgets.List {
	lst := widgets.NewList()
	lst.PageSize = 6
	lst.EmptyText = "暂无数据"
	for i := 1; i <= 23; i++ {
		lst.Items = append(lst.Items, widgets.ListItem{
			Key:    ui.Key(fmt.Sprintf("item-%d", i)),
			Widget: &widgets.Label{Text: fmt.Sprintf("条目 %02d", i)},
		})
	}
	return lst
}()

func main() {
	launcher.Main(doc.Meta{
		ID: "list", Title: "List 键控列表", Group: doc.GroupBasic,
		What: []string{
			"同层同 Key 的子项跨帧复用节点：布局与合成缓存都保留",
			"适合内容顺序会增删/重排的静态列表；超大数据量用 Table（行虚拟化）",
		},
		Demos: []doc.Demo{
			{Title: "分页列表", Note: "23 项 × 每页 6：底部分页器翻页（组合复用 Pagination）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return pagedList // 包级持有：页码状态跨重建（教训⑨）
			}},
			{Title: "空态", Note: "Items 空 + EmptyText → Empty 组件组合复用", Build: func(ctx *ui.BuildContext) ui.Widget {
				lst := widgets.NewList()
				lst.EmptyText = "暂无数据"
				return lst
			}},
			{Title: "键控复用", Note: "ListItem.Key 稳定的项在重建时复用节点", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.NewList(
					widgets.ListItem{Key: "a", Widget: &widgets.Label{Text: "键控项 A"}},
					widgets.ListItem{Key: "b", Widget: &widgets.Label{Text: "键控项 B"}},
					widgets.ListItem{Key: "c", Widget: &widgets.Label{Text: "键控项 C（键控复用，不做虚拟滚动）"}},
				)
			}},
		},
		API: []doc.APIRow{
			{Name: "NewList", Type: "func(items ...ListItem) *List", Desc: "构造键控列表"},
			{Name: "Items", Type: "[]ListItem{Key, Widget}", Desc: "列表项；Key 是复用身份"},
			{Name: "Gap", Type: "float32", Default: "0", Desc: "项间距"},
			{Name: "Nodes", Type: "func() []*scene.Node", Desc: "诊断用：返回当前构建的节点"},
			{Name: "ui.KeyedList", Type: "接口", Desc: "List 实现键控 diff（ui/keyed.go 最小复用）"},
		},
	})
}
