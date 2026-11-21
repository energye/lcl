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
		ID: "popover", Title: "Popover 气泡卡片", Group: doc.GroupFeedback,
		What: []string{
			"点击触发：Child 点击开/再点关+OutsideClick（DismissOutsideClick）",
			"浅底卡片（白底+描边+阴影——与 Tooltip 深底区分，antd Popover 白底）",
			"Title 加粗行+Content Widget（尺寸按内容实测高度自适应）",
		},
		Demos: []doc.Demo{
			{Title: "点击弹出卡片", Note: "点击按钮开/再点关；标题+富内容 Widget", Build: func(ctx *ui.BuildContext) ui.Widget {
				pop := &widgets.Popover{
					Child: widgets.NewButton("点我看详情"),
					Title: "卡片标题",
					Content: doc.Node(func(c *ui.BuildContext) *scene.Node {
						col := scene.Flex(scene.Vertical, 6)
						col.Add((&widgets.Label{Text: "这里是富内容区域："}).Build(c))
						col.Add((&widgets.Label{Text: "支持任意 Widget 组合排版。"}).Build(c))
						return col
					}),
					Width: 260,
				}
				return pop
			}},
		},
		API: []doc.APIRow{
			{Name: "Child", Type: "ui.Widget", Desc: "触发元素（点击开/再点关）"},
			{Name: "Title", Type: "string", Desc: "卡片标题（加粗行）"},
			{Name: "Content", Type: "ui.Widget", Desc: "卡片内容（Widget 组合）"},
			{Name: "Width", Type: "float32", Default: "260", Desc: "卡片宽度"},
			{Name: "Close/IsOpen", Type: "func() / func() bool", Desc: "关闭 / 状态查询"},
		},
	})
}
