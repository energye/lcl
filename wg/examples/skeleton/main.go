package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "skeleton", Title: "Skeleton 骨架屏", Group: doc.GroupDisplay,
		What: []string{
			"加载占位灰条：Phase 外驱或 Animated=true 呼吸（1.4s/圈，透明度 0.45..1，antd active 节奏）",
			"计划（M10）：avatar/title/paragraph 组合形态",
		},
		Demos: []doc.Demo{
			{Title: "四行占位", Note: "Rows=4 宽 360", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Skeleton{Rows: 4, Width: 360, Height: 14}
			}},
			{Title: "呼吸态", Note: "Animated=true：墙钟相位驱动透明度摆动", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Skeleton{Rows: 3, Width: 360, Height: 14, Animated: true}
			}},
		},
		API: []doc.APIRow{
			{Name: "Rows", Type: "int", Default: "3", Desc: "灰条行数"},
			{Name: "Width", Type: "float32", Desc: "行宽"},
			{Name: "Height", Type: "float32", Default: "14", Desc: "行高"},
			{Name: "Phase", Type: "float32", Default: "0", Desc: "呼吸相位 0..1，由调用方每帧推进"},
			{Name: "Animated", Type: "bool", Default: "false", Desc: "呼吸态（墙钟驱动，组件级动画声明）"},
		},
	})
}
