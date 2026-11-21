package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var tl = widgets.NewTimeline([]widgets.TimelineItem{
	{Time: "2026-11-21 09:00", Text: "创建发布单", Color: widgets.TimelineBlue},
	{Time: "2026-11-21 09:30", Text: "测试通过", Color: widgets.TimelineGreen},
	{Time: "2026-11-21 10:00", Text: "预发验证驳回", Color: widgets.TimelineRed},
	{Time: "2026-11-21 11:00", Text: "归档", Color: widgets.TimelineGray},
})

func main() {
	launcher.Main(doc.Meta{
		ID: "timeline", Title: "Timeline 时间轴", Group: doc.GroupData,
		What: []string{
			"垂直时间轴：左轴列（圆点+竖线贯穿）+ 右内容列（时间小字+主文本）",
			"节点色：blue 主色 / green 成功 / red 危险 / gray 灰（antd color 语义）",
			"首行无上延线、尾行无下延线（轴随行数伸缩）",
		},
		Demos: []doc.Demo{
			{Title: "发布流程", Note: "蓝/绿/红/灰四色节点：创建→测试→驳回→归档", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tl
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTimeline", Type: "func(items []TimelineItem) *Timeline", Desc: "构造"},
			{Name: "Items", Type: "[]TimelineItem{Time,Text,Color}", Desc: "时间轴项"},
			{Name: "Color", Type: "string", Desc: "节点色：blue/green/red/gray"},
			{Name: "Width", Type: "float32", Default: "300", Desc: "内容宽"},
		},
	})
}
