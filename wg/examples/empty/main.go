package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "empty", Title: "Empty 空状态", Group: doc.GroupDisplay,
		What: []string{
			"数据为空时的占位：插图与两级文案全部用 scene 图元画出",
			"默认文案「暂无数据」",
		},
		Demos: []doc.Demo{
			{Title: "基本用法", Note: "主文案 + 次级说明", Build: func(ctx *ui.BuildContext) ui.Widget {
				e := widgets.Empty{Text: "暂无数据", Subtext: "列表还没有任何条目", Width: 320}
				return e
			}},
		},
		API: []doc.APIRow{
			{Name: "Text", Type: "string", Default: "暂无数据", Desc: "主文案"},
			{Name: "Subtext", Type: "string", Desc: "次级说明"},
			{Name: "Width", Type: "float32", Desc: "占位宽度"},
		},
	})
}
