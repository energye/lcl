package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets/chart"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "chart-scatter", Title: "Scatter 散点图", Group: doc.GroupChart,
		What: []string{
			"双轴散点：点大小随第三维数据变化，轴范围 niceRange 归整",
		},
		Demos: []doc.Demo{
			{Title: "三维散点", Note: "X/Y 定位、Size 控制点大小", Build: func(ctx *ui.BuildContext) ui.Widget {
				return &chart.Scatter{
					Points: []chart.ScatterPoint{
						{X: 1, Y: 2, Size: 5}, {X: 2, Y: 5, Size: 8}, {X: 3, Y: 3, Size: 5},
						{X: 4, Y: 8, Size: 11}, {X: 5, Y: 6, Size: 7}, {X: 6, Y: 9, Size: 9},
					},
					Width: 520, Height: 240,
				}
			}},
		},
		API: []doc.APIRow{
			{Name: "Points", Type: "[]ScatterPoint{X,Y,Size}", Desc: "散点数据"},
			{Name: "Width/Height", Type: "float32", Desc: "画布尺寸"},
		},
	})
}
