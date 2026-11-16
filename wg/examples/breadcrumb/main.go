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
)

func init() {
	crumb.OnClick = func(i int) {
		status.Text = "点击了：" + crumb.Items[i]
		status.MarkDirty()
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "breadcrumb", Title: "Breadcrumb 面包屑", Group: doc.GroupDisplay,
		What: []string{
			"层级路径导航：最后一项为当前页不可点，其余项可点击回调",
			"分隔符固定为 ›（计划 M11 支持自定义）",
		},
		Demos: []doc.Demo{
			{Title: "三级路径", Note: "点击非末项回显", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(crumb.Build(ctx))
				col.Add(status.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "Items", Type: "[]string", Desc: "路径项，最后一项为当前页"},
			{Name: "OnClick", Type: "func(int)", Desc: "点击回调（参数为项下标）"},
		},
	})
}
