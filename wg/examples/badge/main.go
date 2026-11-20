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
		ID: "badge", Title: "Badge 徽标", Group: doc.GroupDisplay,
		What: []string{
			"角标：计数/文本/圆点三态，可与其他节点组合使用",
			"默认 Danger 红底；overflowCount 溢出省略（>99 显示 99+）",
			"status 语义色（Success 绿/Processing 主色/Error 红）",
			"角标集成：Badge.Child 包裹组件，角标叠右上角出血（antd 核心用法）",
			"OverflowCount 显示上限：超过显示 N+（默认 99，可调如 10→10+）",
		},
		Demos: []doc.Demo{
			{Title: "三态角标", Note: "计数 3 / 文本 NEW / 圆点；溢出 150 显示 99+", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 18)
				row.Add((&widgets.Badge{Count: 3}).Build(ctx))
				row.Add((&widgets.Badge{Text: "NEW"}).Build(ctx))
				row.Add((&widgets.Badge{Dot: true}).Build(ctx))
				row.Add((&widgets.Badge{Count: 150}).Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "status 语义色", Note: "Success 绿 / Processing 主色 / Error 红（antd status）", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 18)
				for _, st := range []widgets.BadgeStatus{widgets.BadgeStatusSuccess, widgets.BadgeStatusProcessing, widgets.BadgeStatusError} {
					row.Add((&widgets.Badge{Text: "status", Status: st}).Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "角标集成", Note: "Badge.Child 包裹 Avatar/Button，角标右上角出血", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 24)
				row.Add((&widgets.Badge{Count: 5, Child: &widgets.Avatar{Icon: "user", Size: 40}}).Build(ctx))
				row.Add((&widgets.Badge{Dot: true, Child: widgets.NewButton("消息")}).Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "OverflowCount", Note: "上限 10：Count=15 显示 10+（默认 99→99+）", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 18)
				row.Add((&widgets.Badge{Count: 15, OverflowCount: 10}).Build(ctx))
				row.Add((&widgets.Badge{Count: 150}).Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
		},
		API: []doc.APIRow{
			{Name: "Count", Type: "int", Desc: "计数值（设置后优先于 Text）"},
			{Name: "Text", Type: "string", Desc: "文本角标，如 NEW"},
			{Name: "Dot", Type: "bool", Desc: "只显示小圆点"},
			{Name: "Color", Type: "geom.Color", Default: "Color.Danger", Desc: "底色"},
			{Name: "TextColor", Type: "geom.Color", Default: "Color.OnDanger", Desc: "文字色"},
			{Name: "Status", Type: "widgets.BadgeStatus", Default: "Default", Desc: "Success/Processing/Error 语义色（Color 显式时优先）"},
			{Name: "Child", Type: "ui.Widget", Desc: "被包裹组件（角标叠右上角出血）"},
			{Name: "OverflowCount", Type: "int", Default: "99", Desc: "显示上限（超过显示 N+）"},
		},
	})
}
