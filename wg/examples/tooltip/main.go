package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	tip := &widgets.Tooltip{
		Text:  "悬停 400ms 后显示，移开自动消失",
		Child: widgets.NewButton("把鼠标停在这里"),
	}
	launcher.Main(doc.Meta{
		ID: "tooltip", Title: "Tooltip 提示", Group: doc.GroupControl,
		What: []string{
			"悬停提示：走 overlay.Host 浮层（BandTooltip 层带）",
			"Dismiss(0) 不吞后续点击——浮层消失当帧即可点到下面的元素",
			"计划（M9）：富内容 Tooltip 与 Popover 拆分",
		},
		Demos: []doc.Demo{
			{Title: "基本悬停", Note: "Enter 延迟显示、Leave 消失", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tip
			}},
		},
		API: []doc.APIRow{
			{Name: "Child", Type: "ui.Widget", Desc: "触发元素"},
			{Name: "Text", Type: "string", Desc: "提示内容"},
			{Name: "Placement", Type: "枚举", Desc: "弹出方位"},
			{Name: "Build", Type: "ui.Widget", Desc: "值类型组件，直接内嵌"},
		},
	})
}
