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
		ID: "progress", Title: "Progress 进度条", Group: doc.GroupControl,
		What: []string{
			"线形进度条：绘制走 scene 图元，Value ∈ [0,1]",
			"计划（M10）：对齐 antd v5——circle/dashboard 形态与 status 语义色",
		},
		Demos: []doc.Demo{
			{Title: "四档进度", Note: "0 / 25% / 60% / 100%", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10).SetWidth(420)
				for _, v := range []float32{0, 0.25, 0.6, 1} {
					col.Add((&widgets.Progress{Value: v, Width: 400, Height: 10}).Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "Value", Type: "float32", Default: "0", Desc: "进度 0..1"},
			{Name: "Width/Height", Type: "float32", Desc: "轨道尺寸"},
			{Name: "Color", Type: "geom.Color", Default: "Color.Primary", Desc: "进度色"},
		},
	})
}
