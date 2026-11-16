package main

import (
	"time"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 有状态演示：实例放包级变量，跨 Build 保住交互状态。
var (
	clicked  = widgets.NewButton("点我")
	loading  = widgets.NewButton("加载 2 秒")
	spinning bool
)

func init() {
	clicked.OnClick = func() {
		clicked.Text = "已点击"
		clicked.MarkDirty()
	}
	loading.OnClick = func() {
		if spinning {
			return
		}
		spinning = true
		loading.Loading = true
		loading.MarkDirty()
		time.AfterFunc(2*time.Second, func() {
			spinning = false
			loading.Loading = false
			loading.MarkDirty()
		})
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "button", Title: "Button 按钮", Group: doc.GroupControl,
		What: []string{
			"触发一个即时操作的按钮；点击走 GestureDetector → Tap → OnClick",
			"type 对齐 antd：primary / default / dashed / text / link 五种",
			"danger 危险语义（实心/描边/文字三套配色）、loading 加载态（旋转指示并忽略点击）、block 占满、shape 圆形/胶囊",
			"视觉对齐：hover 转主色、禁用专色不降透明度、paddingInline 15（small 7）、控件高 24/32/40",
			"未实现：href 跳转、ghost 修饰符（如实标注）",
		},
		Demos: []doc.Demo{
			{Title: "五种 type", Note: "primary / default / dashed / text / link", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 10)
				p := widgets.NewButton("Primary")
				p.Type = widgets.TypePrimary
				row.Add(p.Build(ctx))
				row.Add(widgets.NewButton("Default").Build(ctx))
				d := widgets.NewButton("Dashed")
				d.Type = widgets.TypeDashed
				row.Add(d.Build(ctx))
				t := widgets.NewButton("Text")
				t.Type = widgets.TypeText
				row.Add(t.Build(ctx))
				l := widgets.NewButton("Link")
				l.Type = widgets.TypeLink
				row.Add(l.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "danger 危险按钮", Note: "实心/描边/文字三种 danger 配色", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 10)
				p := widgets.NewButton("Danger Primary")
				p.Type, p.Danger = widgets.TypePrimary, true
				row.Add(p.Build(ctx))
				d := widgets.NewButton("Danger Default")
				d.Danger = true
				row.Add(d.Build(ctx))
				t := widgets.NewButton("Danger Text")
				t.Type, t.Danger = widgets.TypeText, true
				row.Add(t.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "点击与加载", Note: "点「点我」回显；点「加载 2 秒」看 loading 旋转与点击抑制", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 10)
				row.Add(clicked.Build(ctx))
				row.Add(loading.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "图标 / 形状 / 尺寸 / block", Note: "icon 前置；circle 圆形；round 胶囊；small/large；block 占满", Build: func(ctx *ui.BuildContext) ui.Widget {
				search := widgets.NewButton("Search")
				search.Type = widgets.TypePrimary
				search.Icon = widgets.NewIcon("search")
				circle := widgets.NewButton("")
				circle.Type = widgets.TypePrimary
				circle.Shape = widgets.ShapeCircle
				circle.Icon = widgets.NewIcon("plus")
				round := widgets.NewButton("胶囊")
				round.Shape = widgets.ShapeRound
				small := widgets.NewButton("Small")
				small.Size = ui.SizeSmall
				large := widgets.NewButton("Large")
				large.Size = ui.SizeLarge
				col := scene.Flex(scene.Vertical, 10)
				col.Add(func() *scene.Node {
					row := scene.Flex(scene.Horizontal, 10)
					row.Add(search.Build(ctx))
					row.Add(circle.Build(ctx))
					row.Add(round.Build(ctx))
					row.Add(small.Build(ctx))
					row.Add(large.Build(ctx))
					return row
				}())
				block := widgets.NewButton("Block 占满")
				block.Block = true
				col.Add(block.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "禁用", Note: "antd 禁用专色（不降透明度）：bg #FAFAFA、文字 #BFBFBF", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 10)
				p := widgets.NewButton("禁用主按钮")
				p.Type, p.Disabled = widgets.TypePrimary, true
				row.Add(p.Build(ctx))
				d := widgets.NewButton("禁用默认")
				d.Disabled = true
				row.Add(d.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
		},
		API: []doc.APIRow{
			{Name: "Type", Type: "widgets.ButtonType", Default: "TypeDefault", Desc: "primary/default/dashed/text/link（antd Button.type）"},
			{Name: "Text", Type: "string", Desc: "按钮文字，改后需 MarkDirty"},
			{Name: "Size", Type: "ui.Size", Default: "SizeMedium(32px)", Desc: "Small 24 / Medium 32 / Large 40 控件高度"},
			{Name: "Shape", Type: "widgets.ButtonShape", Default: "ShapeDefault", Desc: "default/circle/round"},
			{Name: "Danger", Type: "bool", Default: "false", Desc: "危险语义：实心=Error 实底，描边/文字=Error 前景"},
			{Name: "Loading", Type: "bool", Default: "false", Desc: "加载中：旋转指示、忽略点击"},
			{Name: "Block", Type: "bool", Default: "false", Desc: "宽度占满父级交叉轴"},
			{Name: "Icon", Type: "*widgets.Icon", Desc: "前置图标（Loading 时被旋转指示替代）"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态（antd disabled 专色）"},
			{Name: "OnClick", Type: "func()", Desc: "点击回调；Enter/Space 同触发"},
			{Name: "MarkDirty", Type: "func()", Desc: "状态改变后请求重建"},
		},
	})
}
