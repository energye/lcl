package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets/chart"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "chart-pie", Title: "PieChart 饼图 / 环图", Group: doc.GroupChart,
		What: []string{
			"饼图与环形图一体：Donut 开关切换，扇形用 render.Path 弧线",
			"右侧自动生成图例",
		},
		Demos: []doc.Demo{
			{Title: "环形图", Note: "Donut=true 内径 0.28", Build: func(ctx *ui.BuildContext) ui.Widget {
				return &chart.PieChart{
					Values: []float64{38, 27, 21, 14},
					Labels: []string{"直接", "搜索", "推荐", "社交"},
					Width:  260, Height: 260, Donut: true,
				}
			}},
		},
		API: []doc.APIRow{
			{Name: "Values", Type: "[]float64", Desc: "各扇区值"},
			{Name: "Labels", Type: "[]string", Desc: "扇区标签（进图例）"},
			{Name: "Donut", Type: "bool", Default: "false", Desc: "true 为环形（内径 0.28）"},
			{Name: "Width/Height", Type: "float32", Desc: "画布尺寸"},
		},
	})
}
