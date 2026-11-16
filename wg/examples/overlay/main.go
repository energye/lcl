package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	tip := &widgets.Tooltip{Text: "浮层按 Band 排序：Popup < Toast < Modal", Child: widgets.NewButton("悬停查看浮层层级")}
	sel := widgets.NewSelect([]string{"浮层 A", "浮层 B", "浮层 C"})
	sel.Width = 200
	launcher.Main(doc.Meta{
		ID: "overlay", Title: "浮层与关闭策略", Group: doc.GroupInput,
		What: []string{
			"浮层系统：overlay.Host 统一管理，按 Band 层带排序（Popup 300 / Tooltip 400 / Modal 500 / Toast 600）",
			"关闭策略三件套由 Host 裁决：OutsideClick（点外关闭）/ Esc / AnchorGone（锚点消失）",
			"Dismiss(0) 不吞后续点击：浮层消失当帧即可点到下面的元素",
		},
		Demos: []doc.Demo{
			{Title: "Tooltip + Select", Note: "两种层带的浮层并存；Esc/点外关闭", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 16)
				row.Add(tip.Build(ctx))
				row.Add(sel.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
		},
		API: []doc.APIRow{
			{Name: "overlay.Host", Type: "ctx.Overlay", Desc: "浮层宿主（Show/Dismiss/裁决）"},
			{Name: "theme.Band", Type: "层带常量", Desc: "Popup 300 / Tooltip 400 / Modal 500 / Toast 600"},
			{Name: "关闭策略", Type: "契约", Desc: "OutsideClick / Esc / AnchorGone 统一裁决"},
			{Name: "Dismiss(0)", Type: "语义", Desc: "关闭浮层但不吞当帧点击"},
		},
	})
}
