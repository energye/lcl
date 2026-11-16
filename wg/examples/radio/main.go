package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	r1 = widgets.NewRadio("选项一", true)
	r2 = widgets.NewRadio("选项二", false)
)

func init() {
	// 互斥由调用方负责：回调里互改兄弟实例状态（M8 将提供 RadioGroup）。
	r1.OnChange = func(bool) { r2.Checked = !r1.Checked; r2.MarkDirty() }
	r2.OnChange = func(bool) { r1.Checked = !r2.Checked; r1.MarkDirty() }
}

func main() {
	launcher.Main(doc.Meta{
		ID: "radio", Title: "Radio 单选", Group: doc.GroupControl,
		What: []string{
			"单选项：一组内最多选中一个；当前互斥由调用方在回调里维护",
			"计划（M8）：RadioGroup 提供声明式互斥与 antd Button 样式",
		},
		Demos: []doc.Demo{
			{Title: "互斥切换", Note: "回调互改兄弟节点状态", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 8)
				col.Add(r1.Build(ctx))
				col.Add(r2.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewRadio", Type: "func(label string, checked bool) *Radio", Desc: "构造"},
			{Name: "Checked", Type: "bool", Desc: "选中状态"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnChange", Type: "func(bool)", Desc: "切换回调"},
		},
	})
}
