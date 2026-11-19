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
		ID: "spin", Title: "Spin 加载指示", Group: doc.GroupDisplay,
		What: []string{
			"旋转弧线加载指示：CustomPaintNode + render.Path 弧线绘制",
			"Phase 明确交调用方控制（帧循环推进），组件内无定时器、headless 可确定性出帧",
			"Spin=true 自动旋转（墙钟相位 + 组件级动画声明，外部不必推帧）",
			"计划（M10）：tip 文案与容器包裹形态",
		},
		Demos: []doc.Demo{
			{Title: "四个相位", Note: "Phase 0 / 0.25 / 0.5 / 0.75 的弧线起点", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 24)
				for _, p := range []float32{0, 0.25, 0.5, 0.75} {
					row.Add((&widgets.Spin{Size: 28, Phase: p}).Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "自动旋转", Note: "Spin=true：墙钟相位 800ms/圈，静止窗持续重录", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(*ui.BuildContext) *scene.Node {
					return (&widgets.Spin{Size: 28, Spin: true}).Build(ctx)
				})
			}},
		},
		API: []doc.APIRow{
			{Name: "Size", Type: "float32", Default: "28", Desc: "指示器直径"},
			{Name: "Phase", Type: "float32", Default: "0", Desc: "旋转相位 0..1，由调用方推进"},
			{Name: "Spin", Type: "bool", Default: "false", Desc: "自动旋转（800ms/圈，组件级动画声明）"},
		},
	})
}
