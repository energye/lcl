package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// tip 包级持有：闭包内新建会让 hover 态随重建丢失。
var tip = &widgets.Tooltip{
	Text:  "悬停 400ms 后显示，移开自动消失",
	Child: widgets.NewButton("把鼠标停在这里"),
}

// tipMulti 多行文本卡实例（\n 换行；富内容形态已拆分至 Popover）。
var tipMulti = &widgets.Tooltip{
	Child: widgets.NewButton("悬停看多行提示"),
	Text:  "第一步：填写表单字段\n第二步：点击提交按钮",
}

func main() {
	launcher.Main(doc.Meta{
		ID: "tooltip", Title: "Tooltip 提示", Group: doc.GroupControl,
		What: []string{
			"悬停提示：走 overlay.Host 浮层（BandTooltip 层带）",
			"Dismiss(0) 不吞后续点击——浮层消失当帧即可点到下面的元素",
			"**富内容形态已拆分至 Popover**（白底 Title/Content——antd 语义边界：Tooltip 深底短文本、Popover 白底富内容，M9 尾轮）",
		},
		Demos: []doc.Demo{
			{Title: "基本悬停", Note: "Enter 延迟显示、Leave 消失", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tip
			}},
			{Title: "多行文本", Note: "Text 支持 \\n 换行（深底形态保持——富内容请用 Popover）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tipMulti
			}},
		},
		API: []doc.APIRow{
			{Name: "Child", Type: "ui.Widget", Desc: "触发元素"},
			{Name: "Text", Type: "string", Desc: "提示文本（\\n 多行；富内容形态见 Popover）"},
			{Name: "Placement", Type: "枚举", Desc: "弹出方位"},
			{Name: "Build", Type: "ui.Widget", Desc: "值类型组件，直接内嵌"},
		},
	})
}
