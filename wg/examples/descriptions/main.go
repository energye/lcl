package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var (
	descBordered = widgets.NewDescriptions([]widgets.DescriptionItem{
		{Label: "用户名", Content: "张三"},
		{Label: "城市", Content: "杭州"},
		{Label: "手机号", Content: "138-1234-5678"},
		{Label: "邮箱", Content: "zhangsan@example.com"},
		{Label: "性别", Content: "男"},
		{Label: "备注", Content: "内测用户"},
	})
	descPlain = widgets.NewDescriptions([]widgets.DescriptionItem{
		{Label: "创建时间", Content: "2026-11-21 10:24"},
		{Label: "版本", Content: "v1.0.0"},
		{Label: "状态", Content: "运行中"},
		{Label: "负责人", Content: "李四"},
	})
)

func init() {
	descBordered.Title = "用户信息"
	descBordered.Columns = 2
	descBordered.Bordered = true
	descBordered.Width = 560
	descPlain.Columns = 2
	descPlain.Width = 560
}

func main() {
	launcher.Main(doc.Meta{
		ID: "descriptions", Title: "Descriptions 描述列表", Group: doc.GroupData,
		What: []string{
			"标题行 + **Grid(columns×2) 网格组合复用**（label 格+content 格交替）",
			"双形态：Bordered（label 灰底格+全网格描边，antd bordered）/默认（label 灰字冒号+疏排）",
			"antd column 语义：每行组数（首轮 Items 为文本值）",
		},
		Demos: []doc.Demo{
			{Title: "带边框", Note: "Bordered=true：label 灰底格+描边网格（column=2）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return descBordered
			}},
			{Title: "默认无边框", Note: "label 灰字冒号+正文疏排", Build: func(ctx *ui.BuildContext) ui.Widget {
				return descPlain
			}},
		},
		API: []doc.APIRow{
			{Name: "NewDescriptions", Type: "func(items []DescriptionItem) *Descriptions", Desc: "构造"},
			{Name: "Items", Type: "[]DescriptionItem{Label,Content}", Desc: "描述项（首轮文本值）"},
			{Name: "Title", Type: "string", Desc: "标题行"},
			{Name: "Columns", Type: "int", Default: "2", Desc: "每行组数（antd column）"},
			{Name: "Bordered", Type: "bool", Default: "false", Desc: "边框形态（label 灰底格+网格描边）"},
			{Name: "Width", Type: "float32", Default: "560", Desc: "总宽"},
		},
	})
}
