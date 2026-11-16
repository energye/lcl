package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets/chart"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "chart-line", Title: "LineChart 折线图", Group: doc.GroupChart,
		What: []string{
			"多序列折线图：可选面积填充、数据点、网格与坐标轴",
			"数值映射是纯函数；配色默认从主题推导（Palette 可覆盖）",
		},
		Demos: []doc.Demo{
			{Title: "双序列", Note: "面积填充 + 数据点 + 周标签", Build: func(ctx *ui.BuildContext) ui.Widget {
				return &chart.LineChart{
					Series: [][]float64{
						{3, 5, 4, 8, 7, 11, 9},
						{2, 3, 6, 5, 9, 8, 12},
					},
					Width: 560, Height: 240, Labels: []string{"一", "二", "三", "四", "五", "六", "日"},
					FillArea: true, ShowDots: true, AxisLabel: "访问量",
				}
			}},
		},
		API: []doc.APIRow{
			{Name: "Series", Type: "[][]float64", Desc: "序列数据（每条一线）"},
			{Name: "Labels", Type: "[]string", Desc: "X 轴标签"},
			{Name: "Width/Height", Type: "float32", Desc: "画布尺寸"},
			{Name: "FillArea", Type: "bool", Desc: "面积填充（alpha 48）"},
			{Name: "ShowDots", Type: "bool", Desc: "显示数据点"},
			{Name: "AxisLabel", Type: "string", Desc: "轴说明文字"},
			{Name: "Palette", Type: "[]geom.Color", Default: "主题推导", Desc: "序列配色覆盖"},
		},
	})
}
