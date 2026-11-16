package main

import (
	"math"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets/chart"
)

func main() {
	vals := make([]float64, 128)
	for i := range vals {
		t := float64(i) / 12
		vals[i] = 0.5*math.Sin(t*1.8) + 0.3*math.Sin(t*4.1) + 0.2*math.Sin(t*0.7)
	}
	launcher.Main(doc.Meta{
		ID: "chart-wave", Title: "Waveform 波形图", Group: doc.GroupChart,
		What: []string{
			"密集采样折线：中线归一化，数据驱动、无定时器",
		},
		Demos: []doc.Demo{
			{Title: "合成波形", Note: "三个正弦叠加，128 采样点", Build: func(ctx *ui.BuildContext) ui.Widget {
				return &chart.Waveform{Values: vals, Width: 560, Height: 180}
			}},
		},
		API: []doc.APIRow{
			{Name: "Values", Type: "[]float64", Desc: "密集采样值（中线归一）"},
			{Name: "Width/Height", Type: "float32", Desc: "画布尺寸"},
		},
	})
}
