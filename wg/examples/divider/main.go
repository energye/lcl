package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "divider", Title: "Divider 分隔线", Group: doc.GroupBasic,
		What: []string{
			"区块之间的水平/竖直分隔线，纯 scene 图元绘制、无子句柄",
			"厚度与长度可配；用于内容分组的视觉边界",
		},
		Demos: []doc.Demo{
			{Title: "水平分隔线", Note: "默认 1px；Thickness 可加粗", Build: func(ctx *ui.BuildContext) ui.Widget {
				return &widgets.Column{Gap: 10, Children: []ui.Widget{
					&widgets.Label{Text: "上方内容"},
					&widgets.Divider{Length: 400},
					&widgets.Label{Text: "下方内容"},
					&widgets.Divider{Length: 400, Thickness: 3},
					&widgets.Label{Text: "加粗分隔线"},
				}}
			}},
		},
		API: []doc.APIRow{
			{Name: "Vertical", Type: "bool", Default: "false", Desc: "true 时为竖直分隔线"},
			{Name: "Length", Type: "float32", Default: "父级约束", Desc: "分隔线主轴长度"},
			{Name: "Thickness", Type: "float32", Default: "1", Desc: "线宽"},
			{Name: "Color", Type: "geom.Color", Default: "Color.Border", Desc: "线色，零值取主题"},
		},
	})
}
