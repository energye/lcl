package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "card", Title: "Card 卡片", Group: doc.GroupDisplay,
		What: []string{
			"内容容器：标题/内容/底部三段式，圆角+边框+可选阴影",
			"Elevated 控制投影；宽高由父级约束决定（放进 Flex 给 flex）",
			"计划（M9）：对齐 antd——actions 操作区、cover 封面、hover 可点",
		},
		Demos: []doc.Demo{
			{Title: "三段式卡片", Note: "标题/正文/底部 + Elevated 阴影", Build: func(ctx *ui.BuildContext) ui.Widget {
				card := widgets.NewCard(&widgets.Label{Text: "卡片正文 · 内容由 ui.Widget 组合"})
				card.Title = &widgets.Label{Text: "卡片标题"}
				card.Footer = &widgets.Label{Text: "底部区域"}
				return card
			}},
		},
		API: []doc.APIRow{
			{Name: "NewCard", Type: "func(content ui.Widget) *Card", Desc: "构造（默认带阴影）"},
			{Name: "Title", Type: "ui.Widget", Desc: "标题区（下方有分隔线）"},
			{Name: "Content", Type: "ui.Widget", Desc: "正文区"},
			{Name: "Footer", Type: "ui.Widget", Desc: "底部操作区"},
			{Name: "Padding", Type: "geom.Insets", Default: "Space.LG", Desc: "内边距"},
			{Name: "Elevated", Type: "bool", Default: "NewCard=true", Desc: "是否投影（Elevation[0]）"},
		},
	})
}
