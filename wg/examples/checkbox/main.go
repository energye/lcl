package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var cb = widgets.NewCheckbox("启用动画", true)

func main() {
	launcher.Main(doc.Meta{
		ID: "checkbox", Title: "Checkbox 复选框", Group: doc.GroupControl,
		What: []string{
			"二态勾选项：点击或按 Space 切换 Checked",
			"状态由组件实例跨 Build 持有；OnChange 通知外部",
		},
		Demos: []doc.Demo{
			{Title: "勾选与禁用", Note: "点击切换；禁用项不响应且降透明度", Build: func(ctx *ui.BuildContext) ui.Widget {
				off := widgets.NewCheckbox("禁用项", false)
				off.Disabled = true
				col := scene.Flex(scene.Vertical, 8)
				col.Add(cb.Build(ctx))
				col.Add(off.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewCheckbox", Type: "func(label string, checked bool) *Checkbox", Desc: "构造"},
			{Name: "Checked", Type: "bool", Desc: "勾选状态"},
			{Name: "Label", Type: "string", Desc: "右侧文字"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnChange", Type: "func(bool)", Desc: "切换回调，参数为新状态"},
		},
	})
}
