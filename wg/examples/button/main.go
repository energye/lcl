package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 有状态演示：实例放包级变量，跨 Build 保住交互状态。
var demoBtn = widgets.NewButton("点我")

func init() {
	demoBtn.OnClick = func() {
		demoBtn.Text = "已点击"
		demoBtn.MarkDirty()
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "button", Title: "Button 按钮", Group: doc.GroupControl,
		What: []string{
			"触发一个即时操作的按钮；点击走 GestureDetector → Tap → OnClick",
			"变体（实心/描边/浅底）与禁用态由 Variant/Disabled 控制",
			"键盘可达：获得焦点后 Enter/Space 触发",
		},
		Demos: []doc.Demo{
			{Title: "点击计数", Note: "真窗里请用鼠标实际点一下：状态由组件实例持有，重建不丢", Build: func(ctx *ui.BuildContext) ui.Widget {
				return demoBtn
			}},
			{Title: "变体与禁用", Note: "VariantOutline 描边样式；Disabled 后不响应点击且降透明度", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 10)
				disabled := widgets.NewButton("禁用")
				disabled.Disabled = true
				outline := widgets.NewButton("描边")
				outline.Variant = ui.VariantOutline
				row.Add(disabled.Build(ctx))
				row.Add(outline.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewButton", Type: "func(text string) *Button", Desc: "构造按钮"},
			{Name: "Text", Type: "string", Desc: "按钮文字，改后需 MarkDirty"},
			{Name: "Variant", Type: "ui.Variant", Default: "VariantSolid", Desc: "Solid/Outline/Ghost/Soft 四种变体"},
			{Name: "Size", Type: "ui.Size", Default: "SizeMedium(32px)", Desc: "Small 24 / Medium 32 / Large 40 控件高度"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态：不响应事件、降透明度"},
			{Name: "OnClick", Type: "func()", Desc: "点击回调（GestureArena 裁决 Tap）"},
			{Name: "MarkDirty", Type: "func()", Desc: "状态改变后请求重建"},
		},
	})
}
