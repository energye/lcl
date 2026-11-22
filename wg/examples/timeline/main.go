package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
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

var tlAlt = func() *widgets.Timeline {
	t := widgets.NewTimeline([]widgets.TimelineItem{
		{Time: "09:00", Text: "创建发布单", Color: widgets.TimelineBlue},
		{Time: "09:30", Text: "测试通过", Color: widgets.TimelineGreen},
		{Time: "10:00", Text: "预发验证驳回", Color: widgets.TimelineRed},
		{Time: "11:00", Text: "归档", Color: widgets.TimelineGray},
	})
	t.Mode = widgets.TimelineAlternate
	t.Width = 420
	return t
}()

var tlRight = func() *widgets.Timeline {
	t := widgets.NewTimeline([]widgets.TimelineItem{
		{Time: "09:00", Text: "环境就绪", Color: widgets.TimelineBlue},
		{Time: "10:00", Text: "部署完成", Color: widgets.TimelineGreen},
		{Time: "10:30", Text: "回滚演练", Color: widgets.TimelineGray},
	})
	t.Mode = widgets.TimelineRight
	t.Width = 420
	return t
}()

var tlDot = func() *widgets.Timeline {
	t := widgets.NewTimeline([]widgets.TimelineItem{
		{Time: "09:00", Text: "普通节点", Color: widgets.TimelineBlue},
		{Time: "09:30", Text: "里程碑（自定义 Icon dot）", Color: widgets.TimelineGreen,
			Dot: &widgets.Icon{Name: "check-circle", Size: 18, Color: geom.RGBA(82, 196, 26)}},
		{Time: "10:00", Text: "警告节点", Color: widgets.TimelineGray,
			Dot: &widgets.Icon{Name: "warning-circle", Size: 18, Color: geom.RGBA(250, 140, 22)}},
	})
	t.Width = 300
	return t
}()

func main() {
	launcher.Main(doc.Meta{
		ID: "timeline", Title: "Timeline 时间轴", Group: doc.GroupData,
		What: []string{
			"垂直时间轴：左轴列（圆点+竖线贯穿）+ 右内容列（时间小字+主文本）",
			"节点色：blue 主色 / green 成功 / red 危险 / gray 灰（antd color 语义）",
			"Mode 方向：left 默认 / alternate 交替（奇偶行左右互换）/ right 右轴",
			"Item.Dot 自定义节点（antd dot 任意节点——Icon 组件组合复用）",
			"首行无上延线、尾行无下延线（轴随行数伸缩）",
		},
		Demos: []doc.Demo{
			{Title: "发布流程", Note: "蓝/绿/红/灰四色节点：创建→测试→驳回→归档", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tl
			}},
			{Title: "交替轴", Note: "Mode=alternate：轴居中，内容奇偶行左右互换（antd mode=alternate）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tlAlt
			}},
			{Title: "右轴", Note: "Mode=right：轴居中，全部内容在左（antd mode=right）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tlRight
			}},
			{Title: "自定义 dot", Note: "Item.Dot 任意 Widget（check-circle/warning-circle Icon 组合复用）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tlDot
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTimeline", Type: "func(items []TimelineItem) *Timeline", Desc: "构造"},
			{Name: "Items", Type: "[]TimelineItem{Time,Text,Color,Widget,Dot}", Desc: "时间轴项"},
			{Name: "Color", Type: "string", Desc: "节点色：blue/green/red/gray"},
			{Name: "Mode", Type: "TimelineMode", Default: "TimelineLeft", Desc: "方向：left/alternate/right"},
			{Name: "Item.Dot", Type: "ui.Widget", Desc: "自定义节点（优先于 Color 圆点）"},
			{Name: "Width", Type: "float32", Default: "300", Desc: "内容宽"},
		},
	})
}
