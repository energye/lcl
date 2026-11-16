package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets/chart"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "chart-bar", Title: "BarChart 柱状图", Group: doc.GroupChart,
		What: []string{
			"分组柱状图：每组内多根柱子并排，槽宽占用 70%",
			"全正区间不强行从 0 起，以保留数据间差异",
		},
		Demos: []doc.Demo{
			{Title: "四季度两组", Note: "分组柱对比", Build: func(ctx *ui.BuildContext) ui.Widget {
				return &chart.BarChart{
					Groups: [][]float64{{12, 18, 9, 24}, {15, 11, 20, 14}},
					Labels: []string{"Q1", "Q2", "Q3", "Q4"},
					Width:  560, Height: 240,
				}
			}},
		},
		API: []doc.APIRow{
			{Name: "Groups", Type: "[][]float64", Desc: "分组数据（每组多柱）"},
			{Name: "Labels", Type: "[]string", Desc: "X 轴组标签"},
			{Name: "Width/Height", Type: "float32", Desc: "画布尺寸"},
		},
	})
}
