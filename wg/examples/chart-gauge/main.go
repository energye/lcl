package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets/chart"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "chart-gauge", Title: "Gauge 仪表盘", Group: doc.GroupChart,
		What: []string{
			"半环仪表：值域钳制到 [0,1]，环宽 8，中央显示百分比或自定义 Label",
		},
		Demos: []doc.Demo{
			{Title: "72%", Note: "默认按百分比显示", Build: func(ctx *ui.BuildContext) ui.Widget {
				return &chart.Gauge{Value: 0.72, Label: "72%", Width: 320, Height: 180}
			}},
		},
		API: []doc.APIRow{
			{Name: "Value", Type: "float32", Default: "0", Desc: "0..1，越界钳制"},
			{Name: "Label", Type: "string", Default: "百分比", Desc: "中央文字"},
			{Name: "Width/Height", Type: "float32", Desc: "画布尺寸"},
		},
	})
}
