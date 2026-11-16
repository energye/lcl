package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var sw = widgets.NewSwitch(true)

func main() {
	launcher.Main(doc.Meta{
		ID: "switch", Title: "Switch 开关", Group: doc.GroupControl,
		What: []string{
			"两态开关：轨道+滑块全部用 scene 图元绘制，无子句柄",
			"点击或按 Space 切换；OnChange 通知外部",
		},
		Demos: []doc.Demo{
			{Title: "开/关与禁用", Note: "选中态为主色轨道，关闭态为灰轨道", Build: func(ctx *ui.BuildContext) ui.Widget {
				dis := widgets.NewSwitch(false)
				dis.Disabled = true
				col := scene.Flex(scene.Vertical, 10)
				col.Add(sw.Build(ctx))
				col.Add(dis.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewSwitch", Type: "func(on bool) *Switch", Desc: "构造"},
			{Name: "On", Type: "bool", Desc: "开关状态"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnChange", Type: "func(bool)", Desc: "切换回调"},
		},
	})
}
