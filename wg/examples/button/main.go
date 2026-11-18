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
// 样式展示按钮同样要包级持有——demo 闭包每次重建都会执行，闭包内
// widget.NewButton 会换新实例，hover/pressed 态随重建丢失（--audit 实测
// "按下无反馈、hover 闪烁"）。闭包内只做实例的 Build。
var (
	clicked  = widgets.NewButton("点我")
	loading  = widgets.NewButton("加载 2 秒")
	spinning bool

	typesRow    []*widgets.Button
	dangerRow   []*widgets.Button
	shapeRow    []*widgets.Button
	blockBtn    *widgets.Button
	disabledRow []*widgets.Button
)

// btn 构造一个展示按钮（样式通过 fn 就地设置）。
func btn(text string, fn func(*widgets.Button)) *widgets.Button {
	b := widgets.NewButton(text)
	if fn != nil {
		fn(b)
	}
	return b
}

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

	typesRow = []*widgets.Button{
		btn("Primary", func(b *widgets.Button) { b.Type = widgets.TypePrimary }),
		btn("Default", nil),
		btn("Dashed", func(b *widgets.Button) { b.Type = widgets.TypeDashed }),
		btn("Text", func(b *widgets.Button) { b.Type = widgets.TypeText }),
		btn("Link", func(b *widgets.Button) { b.Type = widgets.TypeLink }),
	}
	dangerRow = []*widgets.Button{
		btn("Danger Primary", func(b *widgets.Button) { b.Type, b.Danger = widgets.TypePrimary, true }),
		btn("Danger Default", func(b *widgets.Button) { b.Danger = true }),
		btn("Danger Text", func(b *widgets.Button) { b.Type, b.Danger = widgets.TypeText, true }),
	}
	shapeRow = []*widgets.Button{
		btn("Search", func(b *widgets.Button) {
			b.Type = widgets.TypePrimary
			b.Icon = widgets.NewIcon("search")
		}),
		btn("", func(b *widgets.Button) {
			b.Type = widgets.TypePrimary
			b.Shape = widgets.ShapeCircle
			b.Icon = widgets.NewIcon("plus")
		}),
		btn("胶囊", func(b *widgets.Button) { b.Shape = widgets.ShapeRound }),
		btn("Small", func(b *widgets.Button) { b.Size = ui.SizeSmall }),
		btn("Large", func(b *widgets.Button) { b.Size = ui.SizeLarge }),
	}
	blockBtn = btn("Block 占满", func(b *widgets.Button) { b.Block = true })
	disabledRow = []*widgets.Button{
		btn("禁用主按钮", func(b *widgets.Button) { b.Type, b.Disabled = widgets.TypePrimary, true }),
		btn("禁用默认", func(b *widgets.Button) { b.Disabled = true }),
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
				for _, b := range typesRow {
					row.Add(b.Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "danger 危险按钮", Note: "实心/描边/文字三种 danger 配色", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 10)
				for _, b := range dangerRow {
					row.Add(b.Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "点击与加载", Note: "点「点我」回显；点「加载 2 秒」看 loading 旋转与点击抑制", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 10)
				row.Add(clicked.Build(ctx))
				row.Add(loading.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "图标 / 形状 / 尺寸 / block", Note: "icon 前置；circle 圆形；round 胶囊；small/large；block 占满", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				row := scene.Flex(scene.Horizontal, 10)
				for _, b := range shapeRow {
					row.Add(b.Build(ctx))
				}
				col.Add(row)
				col.Add(blockBtn.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "禁用", Note: "antd 禁用专色（不降透明度）：bg #FAFAFA、文字 #BFBFBF", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 10)
				for _, b := range disabledRow {
					row.Add(b.Build(ctx))
				}
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
