package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "scroll", Title: "Scroll 滚动", Group: doc.GroupLayout,
		What: []string{
			"滚动视口与裁剪：滚轮/拖动由 input 分发；滚动位置是布局状态，不属于脏标记",
			"组件层 widgets.Scroll 封装宽高；节点层 scene.Scroll 直接用",
			"SetScrollOffset 程序化定位（虚拟化的唯一真相）",
		},
		Demos: []doc.Demo{
			{Title: "组件滚动", Note: "widgets.Scroll 360×200，滚轮/拖动查看 24 行", Build: func(ctx *ui.BuildContext) ui.Widget {
				child := doc.Node(func(ctx *ui.BuildContext) *scene.Node {
					inner := scene.Flex(scene.Vertical, 6)
					for i := 0; i < 24; i++ {
						inner.Add((&widgets.Label{Text: fmt.Sprintf("滚动行 %c · 第 %02d 行", 'A'+i%26, i+1)}).Build(ctx))
					}
					return inner
				})
				return &widgets.Scroll{Axis: scene.Vertical, Child: child, Width: 360, Height: 200}
			}},
			{Title: "程序化偏移", Note: "SetScrollOffset(80) 初始定位到 80px", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				inner := scene.Flex(scene.Vertical, 4)
				for i := 0; i < 20; i++ {
					inner.Add(doc.Box(ctx, fmt.Sprintf("行 %02d", i+1), th.Color.Surface, 300, 28))
				}
				view := scene.Scroll(scene.Vertical, inner).SetSize2(320, 160)
				view.SetScrollOffset(80)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return view })
			}},
		},
		API: []doc.APIRow{
			{Name: "widgets.Scroll", Type: "值类型{Axis,Child,Width,Height}", Desc: "组件层封装"},
			{Name: "scene.Scroll", Type: "func(axis Axis, child *Node) *Node", Desc: "节点层视口"},
			{Name: "SetScrollOffset", Type: "func(v float32)", Desc: "滚动偏移；越界钳制"},
			{Name: "宽高 <=0", Type: "—", Desc: "由父级约束决定（放进 Flex 要设 flex）"},
		},
	})
}
