package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	status = widgets.NewButton("未点击")
	crumb  = &widgets.Breadcrumb{Items: []string{"首页", "资源管理", "渲染引擎"}}
	// crumbMenu 独立实例（demo 卡级联索引互不干扰；首项文本独特避免
	// 与卡 1「首页」text 锚撞车）。
	crumbMenu = &widgets.Breadcrumb{
		Items: []string{"控制台", "应用中心", "路由管理"},
		Menus: [][]string{
			{"General", "Layout", "Navigation"},
			nil,
			{"Layer", "Frame", "Animation"},
		},
	}
	menuStatus = widgets.NewButton("菜单未点击")
)

func init() {
	crumb.OnClick = func(i int) {
		status.Text = "点击了：" + crumb.Items[i]
		status.MarkDirty()
	}
	crumbMenu.OnMenuClick = func(itemIndex int, label string) {
		menuStatus.Text = "菜单：" + crumbMenu.Items[itemIndex] + " / " + label
		menuStatus.MarkDirty()
	}
	crumbMenu.OnClick = func(i int) {
		if i < len(crumbMenu.Menus) && len(crumbMenu.Menus[i]) > 0 {
			return
		}
		menuStatus.Text = "点击了：" + crumbMenu.Items[i]
		menuStatus.MarkDirty()
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "breadcrumb", Title: "Breadcrumb 面包屑", Group: doc.GroupDisplay,
		What: []string{
			"层级路径导航：最后一项为当前页不可点，其余项可点击回调",
			"下拉菜单（items.menu）：带菜单的项渲染 down 图标，点击弹 BandPopup 浮层，子项 down 即执行回调并关面板",
			"分隔符固定为 ›（计划 M11 支持自定义）",
		},
		Demos: []doc.Demo{
			{Title: "三级路径", Note: "点击非末项回显", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(crumb.Build(ctx))
				col.Add(status.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "下拉菜单", Note: "带菜单的项（控制台/路由管理）有 down 图标，点击弹下拉；子项点击回显", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(crumbMenu.Build(ctx))
				col.Add(menuStatus.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "Items", Type: "[]string", Desc: "路径项，最后一项为当前页"},
			{Name: "OnClick", Type: "func(int)", Desc: "点击回调（参数为项下标）"},
			{Name: "Menus", Type: "[][]string", Desc: "按项下标挂下拉菜单（antd items.menu）；空=无菜单"},
			{Name: "OnMenuClick", Type: "func(itemIndex int, label string)", Desc: "下拉子项点击回调"},
		},
	})
}
