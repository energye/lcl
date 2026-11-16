package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "steps", Title: "Steps 步骤条", Group: doc.GroupDisplay,
		What: []string{
			"分步引导：当前步及之前高亮，后续置灰",
			"计划（M11）：对齐 antd——status（wait/process/finish/error）、垂直方向、自定义图标",
		},
		Demos: []doc.Demo{
			{Title: "四步进度", Note: "Current=2：第三步进行中", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Steps{Labels: []string{"选择", "配置", "确认", "完成"}, Current: 2}
			}},
		},
		API: []doc.APIRow{
			{Name: "Labels", Type: "[]string", Desc: "步骤名"},
			{Name: "Current", Type: "int", Default: "0", Desc: "当前步下标（0 起）"},
		},
	})
}
