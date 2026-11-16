package examples

import (
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets/chart"
	"math"
)

func init() {
	Register(Example{ID: "chart-line", Name: "折线图", Group: GroupChart, Build: func(ctx *ui.BuildContext) *scene.Node {
		c := chart.LineChart{
			Series: [][]float64{
				{3, 5, 4, 8, 7, 11, 9},
				{2, 3, 6, 5, 9, 8, 12},
			},
			Width: 560, Height: 240, Labels: []string{"一", "二", "三", "四", "五", "六", "日"},
			FillArea: true, ShowDots: true, AxisLabel: "访问量",
		}
		return Section(ctx, "LineChart", "多序列、面积填充、网格与坐标轴；数值映射走纯函数", c.Build(ctx))
	}})

	Register(Example{ID: "chart-bar", Name: "柱状图", Group: GroupChart, Build: func(ctx *ui.BuildContext) *scene.Node {
		c := chart.BarChart{
			Groups: [][]float64{{12, 18, 9, 24}, {15, 11, 20, 14}},
			Labels: []string{"Q1", "Q2", "Q3", "Q4"},
			Width:  560, Height: 240,
		}
		return Section(ctx, "BarChart", "分组柱，全正区间不强行从 0 起以保留差异", c.Build(ctx))
	}})

	Register(Example{ID: "chart-pie", Name: "饼图 / 环图", Group: GroupChart, Build: func(ctx *ui.BuildContext) *scene.Node {
		c := chart.PieChart{
			Values: []float64{38, 27, 21, 14},
			Labels: []string{"直接", "搜索", "推荐", "社交"},
			Width:  260, Height: 260, Donut: true,
		}
		return Section(ctx, "PieChart", "Donut 开关切换实心/环形，扇形用 render.Path 弧线", c.Build(ctx))
	}})

	Register(Example{ID: "chart-gauge", Name: "仪表盘", Group: GroupChart, Build: func(ctx *ui.BuildContext) *scene.Node {
		c := chart.Gauge{Value: 0.72, Label: "72%", Width: 320, Height: 180}
		return Section(ctx, "Gauge", "半环仪表，值域钳制到 [0,1]", c.Build(ctx))
	}})

	Register(Example{ID: "chart-wave", Name: "波形图", Group: GroupChart, Build: func(ctx *ui.BuildContext) *scene.Node {
		vals := make([]float64, 128)
		for i := range vals {
			t := float64(i) / 12
			vals[i] = 0.5*math.Sin(t*1.8) + 0.3*math.Sin(t*4.1) + 0.2*math.Sin(t*0.7)
		}
		c := chart.Waveform{Values: vals, Width: 560, Height: 180}
		return Section(ctx, "Waveform", "密集折线，数据驱动、无定时器（Phase 由调用方更新）", c.Build(ctx))
	}})

	Register(Example{ID: "chart-scatter", Name: "散点图", Group: GroupChart, Build: func(ctx *ui.BuildContext) *scene.Node {
		c := chart.Scatter{
			Points: []chart.ScatterPoint{
				{X: 1, Y: 2, Size: 5}, {X: 2, Y: 5, Size: 8}, {X: 3, Y: 3, Size: 5},
				{X: 4, Y: 8, Size: 11}, {X: 5, Y: 6, Size: 7}, {X: 6, Y: 9, Size: 9},
			},
			Width: 520, Height: 240,
		}
		return Section(ctx, "Scatter", "点大小随第三维数据变化", c.Build(ctx))
	}})
}
