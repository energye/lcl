package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "card", Title: "Card 卡片", Group: doc.GroupDisplay,
		What: []string{
			"内容容器：标题/内容/底部三段式，圆角+边框+可选阴影",
			"Elevated 控制投影；宽高由父级约束决定（放进 Flex 给 flex）",
			"actions 操作区（顶部 hairline 分隔、横排均分、直贴底缘）",
			"cover 封面：顶部通栏图元区（直贴卡片边缘，antd cover 语义）",
		},
		Demos: []doc.Demo{
			{Title: "三段式卡片", Note: "标题/正文/底部 + Elevated 阴影", Build: func(ctx *ui.BuildContext) ui.Widget {
				card := widgets.NewCard(&widgets.Label{Text: "卡片正文 · 内容由 ui.Widget 组合"})
				card.Title = &widgets.Label{Text: "卡片标题"}
				card.Footer = &widgets.Label{Text: "底部区域"}
				return card
			}},
			{Title: "actions 操作区", Note: "底部横排均分、直贴底缘（antd actions）", Build: func(ctx *ui.BuildContext) ui.Widget {
				card := widgets.NewCard(&widgets.Label{Text: "带操作区的卡片"})
				card.Title = &widgets.Label{Text: "actions 卡片"}
				card.Actions = []ui.Widget{
					&widgets.Label{Text: "收藏"},
					&widgets.Label{Text: "编辑"},
					&widgets.Label{Text: "分享"},
				}
				return card
			}},
			{Title: "cover 封面卡片", Note: "顶部通栏 Avatar 图标圆（组合复用）+ 标题/内容", Build: func(ctx *ui.BuildContext) ui.Widget {
				card := widgets.NewCard(&widgets.Label{Text: "封面下方是标题与正文"})
				card.Title = &widgets.Label{Text: "cover 卡片"}
				card.Cover = doc.Node(func(c *ui.BuildContext) *scene.Node {
					row := scene.Flex(scene.Horizontal, 0).SetHeight(120).SetCrossAlign(scene.CrossCenter).SetMainAlign(scene.MainCenter)
					row.Add((&widgets.Avatar{Icon: "user", Size: 64}).Build(c))
					return row
				})
				return card
			}},
		},
		API: []doc.APIRow{
			{Name: "NewCard", Type: "func(content ui.Widget) *Card", Desc: "构造（默认带阴影）"},
			{Name: "Title", Type: "ui.Widget", Desc: "标题区（下方有分隔线）"},
			{Name: "Content", Type: "ui.Widget", Desc: "正文区"},
			{Name: "Footer", Type: "ui.Widget", Desc: "底部操作区"},
			{Name: "Cover", Type: "ui.Widget", Desc: "封面区（顶部通栏无内边距）"},
			{Name: "Padding", Type: "geom.Insets", Default: "Space.LG", Desc: "内边距"},
			{Name: "Elevated", Type: "bool", Default: "NewCard=true", Desc: "是否投影（Elevation[0]）"},
		},
	})
}
