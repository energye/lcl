package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// Scroll 实例在 Build 外持有：Offset 跨重建保留（每帧新建会把滚动位置丢掉）。
var (
	vscroll = &widgets.Scroll{Axis: scene.Vertical, Width: 360, Height: 220}
	hscroll = &widgets.Scroll{Axis: scene.Horizontal, Width: 360, Height: 96}
	both    = &widgets.Scroll{Axis: scene.Both, Width: 360, Height: 160}
)

func main() {
	vscroll.Child = doc.Node(func(ctx *ui.BuildContext) *scene.Node {
		inner := scene.Flex(scene.Vertical, 6)
		for i := 0; i < 24; i++ {
			inner.Add((&widgets.Label{Text: fmt.Sprintf("滚动行 %c · 第 %02d 行", 'A'+i%26, i+1)}).Build(ctx))
		}
		return inner
	})

	hscroll.Child = doc.Node(func(ctx *ui.BuildContext) *scene.Node {
		row := scene.Flex(scene.Horizontal, 8)
		for i := 0; i < 16; i++ {
			row.Add(doc.Box(ctx, fmt.Sprintf("列 %02d", i+1), ctx.ThemeOr().Color.Surface, 90, 28))
		}
		return row
	})

	both.Child = doc.Node(func(ctx *ui.BuildContext) *scene.Node {
		col := scene.Flex(scene.Vertical, 6)
		for r := 0; r < 12; r++ {
			row := scene.Flex(scene.Horizontal, 8)
			for i := 0; i < 10; i++ {
				row.Add(doc.Box(ctx, fmt.Sprintf("%c·%02d", 'A'+r, i+1), ctx.ThemeOr().Color.Surface, 90, 28))
			}
			col.Add(row)
		}
		return col
	})

	launcher.Main(doc.Meta{
		ID: "scroll", Title: "Scroll 滚动", Group: doc.GroupLayout,
		What: []string{
			"完整滚动条：滑块拖动、轨道翻页、滚轮；内容溢出时自动出现（ShowScrollbar 可强制显隐）",
			"滚动视口 = 裁剪 + 偏移；纵向内容宽度固定、横向内容高度固定（与布局一致）",
			"Both 双轴：内容宽 = max(视口宽, 自然宽)，两条都钉在视口边缘；各示例窗的页面滚动就是它",
			"Offset/OffsetX 跨重建保留（组件实例必须在 Build 外持有）；越过边界时滚轮事件交给外层滚动",
		},
		Demos: []doc.Demo{
			{Title: "纵向滚动（带滚动条）", Note: "滚轮 / 拖滑块 / 点轨道翻页", Build: func(ctx *ui.BuildContext) ui.Widget {
				return vscroll
			}},
			{Title: "横向滚动", Note: "滚轮（横向轴）或拖动底部滑块", Build: func(ctx *ui.BuildContext) ui.Widget {
				return hscroll
			}},
			{Title: "双轴滚动（Both）", Note: "纵向滚轮 + 横向滚轮（触控板横扫）；两条都钉在视口边缘", Build: func(ctx *ui.BuildContext) ui.Widget {
				return both
			}},
			{Title: "强制隐藏滚动条", Note: "ShowScrollbar=false：仍可滚轮滚动", Build: func(ctx *ui.BuildContext) ui.Widget {
				hide := false
				return &widgets.Scroll{Axis: scene.Vertical, Width: 360, Height: 160,
					ShowScrollbar: &hide, Child: doc.Node(func(ctx *ui.BuildContext) *scene.Node {
						inner := scene.Flex(scene.Vertical, 4)
						for i := 0; i < 20; i++ {
							inner.Add((&widgets.Label{Text: fmt.Sprintf("无条行 %02d", i+1)}).Build(ctx))
						}
						return inner
					})}
			}},
		},
		API: []doc.APIRow{
			{Name: "widgets.Scroll", Type: "struct{Axis,Child,Width,Height,ShowScrollbar,BarThickness,Offset,OffsetX}", Desc: "滚动视口 + 滚动条（Axis=Both 时双轴）"},
			{Name: "widgets.Scrollbar", Type: "struct{Axis,Length,Thickness,Value,Max,Viewport,OnChange}", Desc: "独立滚动条（可单用）"},
			{Name: "SetScrollOffset", Type: "func(v float32)", Desc: "程序化定位；内容未测量时先记录、布局后钳制"},
			{Name: "ShowScrollbar", Type: "*bool", Desc: "nil=溢出自动；true 始终；false 隐藏（仍可滚轮）"},
		},
	})
}
