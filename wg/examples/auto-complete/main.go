package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 展示实例包级持有：demo 闭包内新建会让输入/焦点状态随重建丢失。
var (
	ac = widgets.NewAutoComplete([]string{
		"软件光栅化", "GDI 混合", "无句柄画布", "DisplayList", "Layer 树",
		"脏区合并", "GestureArena", "IME 组合", "ClearType", "预乘 BGRA",
	})
	acDis = widgets.NewAutoComplete([]string{"软件光栅化", "GDI 混合"})
)

func init() {
	acDis.Width = 240
	acDis.Disabled = true
}

func main() {
	launcher.Main(doc.Meta{
		ID: "auto-complete", Title: "AutoComplete 自动完成", Group: doc.GroupControl,
		What: []string{
			"输入过滤建议：contains 匹配（大小写不敏感），弹层走 overlay（BandPopup）",
			"选中后回填输入框并触发 OnSelect",
			"未实现：自定义 option 渲染、分类建议",
		},
		Demos: []doc.Demo{
			{Title: "渲染术语建议", Note: "输入「光」过滤出「软件光栅化」；点建议回填", Build: func(ctx *ui.BuildContext) ui.Widget {
				return ac
			}},
			{Title: "禁用态", Note: "Disabled 后不弹建议、不可聚焦", Build: func(ctx *ui.BuildContext) ui.Widget {
				return acDis
			}},
		},
		API: []doc.APIRow{
			{Name: "NewAutoComplete", Type: "func(suggestions []string) *AutoComplete", Desc: "构造"},
			{Name: "Suggestions", Type: "[]string", Desc: "建议池"},
			{Name: "Value", Type: "string", Desc: "当前文本"},
			{Name: "Width", Type: "float32", Default: "240", Desc: "输入框宽"},
			{Name: "Placeholder", Type: "string", Desc: "占位文字"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnSelect", Type: "func(string)", Desc: "选中建议回调"},
			{Name: "OnChanged", Type: "func(string)", Desc: "文本变化回调"},
		},
	})
}
