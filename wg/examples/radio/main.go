package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	r1 = widgets.NewRadio("选项一", true)
	r2 = widgets.NewRadio("选项二", false)
	r3 = widgets.NewRadio("禁用项", true)
	rg = &widgets.RadioGroup{}
)

func init() {
	// 手动互斥（保留旧用法；组见第二个演示）。
	r1.OnChange = func(bool) { r2.Checked = !r1.Checked; r2.MarkDirty() }
	r2.OnChange = func(bool) { r1.Checked = !r2.Checked; r1.MarkDirty() }
	r3.Disabled = true
	rg.Options = []string{"软件光栅", "GDI 混合", "自动"}
	rg.Value = "自动"
	rg.OnChange = func(s string) { rg.Value = s }
}

func main() {
	launcher.Main(doc.Meta{
		ID: "radio", Title: "Radio 单选", Group: doc.GroupControl,
		What: []string{
			"单选项：M8 对齐 antd——Radio.Group 管互斥（options/value/onChange）",
			"选中态为主色实心圆点（antd 视觉）；散用 Radio 时互斥仍由调用方维护",
		},
		Demos: []doc.Demo{
			{Title: "Radio.Group", Note: "组管互斥与选中值", Build: func(ctx *ui.BuildContext) ui.Widget {
				return rg
			}},
			{Title: "散用 Radio", Note: "回调互改兄弟实例（兼容旧用法）；disabled 禁用专色", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Row{Gap: 12, Children: []ui.Widget{r1, r2, r3}}
			}},
		},
		API: []doc.APIRow{
			{Name: "NewRadio", Type: "func(label string, checked bool) *Radio", Desc: "构造"},
			{Name: "Checked/Disabled", Type: "bool", Desc: "选中 / 禁用"},
			{Name: "OnChange", Type: "func(bool)", Desc: "切换回调"},
			{Name: "RadioGroup", Type: "struct{Options, Value, Disabled, OnChange}", Desc: "组：互斥与选中值由组管理"},
		},
	})
}
