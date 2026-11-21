package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// tip 包级持有：闭包内新建会让 hover 态随重建丢失。
var tip = &widgets.Tooltip{
	Text:  "悬停 400ms 后显示，移开自动消失",
	Child: widgets.NewButton("把鼠标停在这里"),
}

func main() {
	launcher.Main(doc.Meta{
		ID: "tooltip", Title: "Tooltip 提示", Group: doc.GroupControl,
		What: []string{
			"悬停提示：走 overlay.Host 浮层（BandTooltip 层带）",
			"Dismiss(0) 不吞后续点击——浮层消失当帧即可点到下面的元素",
			"富内容：Content Widget（实测宽+高度自适应）+Title 加粗标题行",
		},
		Demos: []doc.Demo{
			{Title: "基本悬停", Note: "Enter 延迟显示、Leave 消失", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tip
			}},
			{Title: "富内容", Note: "Content Widget：标题+多级文案浮层（尺寸按内容实测）", Build: func(ctx *ui.BuildContext) ui.Widget {
				rich := &widgets.Tooltip{
					Child: widgets.NewButton("悬停看富内容"),
					Title: "操作指引",
					Content: doc.Node(func(c *ui.BuildContext) *scene.Node {
						col := scene.Flex(scene.Vertical, 4)
						col.Add((&widgets.Label{Text: "1. 填写表单字段", Color: geom.RGBA(255, 255, 255)}).Build(c))
						col.Add((&widgets.Label{Text: "2. 点击提交按钮", Color: geom.RGBA(255, 255, 255)}).Build(c))
						col.Add((&widgets.Label{Text: "3. 等待校验结果", Color: geom.RGBA(255, 255, 255)}).Build(c))
						return col
					}),
				}
				return rich
			}},
		},
		API: []doc.APIRow{
			{Name: "Child", Type: "ui.Widget", Desc: "触发元素"},
			{Name: "Text", Type: "string", Desc: "提示内容"},
			{Name: "Title", Type: "string", Desc: "富内容形态：加粗标题行"},
			{Name: "Content", Type: "ui.Widget", Desc: "富内容 Widget（优先于 Text，尺寸按内容实测）"},
			{Name: "Placement", Type: "枚举", Desc: "弹出方位"},
			{Name: "Build", Type: "ui.Widget", Desc: "值类型组件，直接内嵌"},
		},
	})
}
