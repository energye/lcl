package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	cb   = widgets.NewCheckbox("启用动画", true)
	pref = &widgets.CheckboxGroup{}
)

func init() {
	pref.Options = []string{"光栅化", "文字通道", "浮层系统"}
	pref.Values = []string{"光栅化"}
	pref.OnChange = func(vs []string) { pref.Values = vs }
}

func main() {
	launcher.Main(doc.Meta{
		ID: "checkbox", Title: "Checkbox 复选框", Group: doc.GroupControl,
		What: []string{
			"二态勾选：点击或按 Space 切换；M8 对齐 antd：indeterminate 半选态、Checkbox.Group 组",
			"勾选态为主色底 + 对勾图标（antd 视觉）；禁用专色",
		},
		Demos: []doc.Demo{
			{Title: "基本 / 半选 / 禁用", Note: "Indeterminate 横杠态由父级聚合状态驱动", Build: func(ctx *ui.BuildContext) ui.Widget {
				off := widgets.NewCheckbox("禁用项", false)
				off.Disabled = true
				mid := widgets.NewCheckbox("半选（父级部分选中）", false)
				mid.Indeterminate = true
				col := scene.Flex(scene.Vertical, 8)
				col.Add(cb.Build(ctx))
				col.Add(mid.Build(ctx))
				col.Add(off.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "Checkbox.Group", Note: "组管选中集合（antd options/value/onChange）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return pref
			}},
		},
		API: []doc.APIRow{
			{Name: "NewCheckbox", Type: "func(label string, checked bool) *Checkbox", Desc: "构造"},
			{Name: "Checked", Type: "bool", Desc: "勾选状态"},
			{Name: "Indeterminate", Type: "bool", Default: "false", Desc: "半选横杠态（antd indeterminate）"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnChange", Type: "func(bool)", Desc: "切换回调"},
			{Name: "CheckboxGroup", Type: "struct{Options, Values, Disabled, OnChange}", Desc: "组：选中集合由组管理"},
		},
	})
}
