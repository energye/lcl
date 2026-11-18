package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 展示控件包级持有：闭包内新建会让勾选/开关/hover 状态随重建丢失。
var (
	cbtn = widgets.NewButton("Button")
	cbx  = widgets.NewCheckbox("Checkbox", true)
	csw  = widgets.NewSwitch(true)
	card = widgets.NewCard(widgets.Label{Text: "Card · NewCard 包装"})
	prog = widgets.Progress{Value: 0.65}
)

func main() {
	launcher.Main(doc.Meta{
		ID: "compat", Title: "组件组合 API", Group: doc.GroupVisual,
		What: []string{
			"widgets 新构造器（NewButton/NewCheckbox/…）是组件的唯一正门",
			"wg/compat 是旧 wg 风格 API（TButton/TCheckbox/TSwitch/TLabel/TCard/TProgress）的门面：只做旧签名→新组件的单向适配",
			"旧示例页面可以一行平移；新代码请直接用 widgets",
		},
		Demos: []doc.Demo{
			{Title: "一行组合", Note: "同一份组件树既可跑真窗也可跑 headless（ui.Widget 契约）", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				row := scene.Flex(scene.Horizontal, 8)
				row.Add(cbtn.Build(ctx))
				row.Add(cbx.Build(ctx))
				row.Add(csw.Build(ctx))
				row.Add((&widgets.Label{Text: "Label", Color: th.Color.Text}).Build(ctx))
				grid := scene.Grid(2, 8).SetWidth(420)
				grid.Add(card.Build(ctx))
				grid.Add(prog.Build(ctx))
				col := scene.Flex(scene.Vertical, 8)
				col.Add(row)
				col.Add(grid)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "compat.TButton 等", Type: "旧 API 门面", Desc: "旧 wg 风格 → widgets 新组件的单向适配"},
			{Name: "compat.AsNode", Type: "func(ui.Widget) *scene.Node", Desc: "组件 → 场景节点"},
			{Name: "ui.Widget", Type: "接口", Desc: "Build(ctx) *scene.Node——真窗/headless 共用契约"},
		},
	})
}
