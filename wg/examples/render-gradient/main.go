package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "render-gradient", Title: "渐变与采样", Group: doc.GroupVisual,
		What: []string{
			"线性/径向渐变：8 子扫描线解析覆盖率 + 渐变色标插值（真实现，非降级）",
			"Shader.Sample(t) 可在 CPU 侧直接采样色标；FillWith 把渐变挂到 Paint",
		},
		Demos: []doc.Demo{
			{Title: "线性 + 径向 + 采样条", Note: "底部 8 块色条是 Sample(t) 的 CPU 采样结果", Build: func(ctx *ui.BuildContext) ui.Widget {
				stops := []render.Stop{
					{Offset: 0, Color: geom.RGBA(255, 90, 90)},
					{Offset: 0.5, Color: geom.RGBA(90, 200, 120)},
					{Offset: 1, Color: geom.RGBA(70, 120, 240)},
				}
				lin := render.NewLinearGradient(geom.Pt(10, 20), geom.Pt(330, 20), stops...)
				rad := render.NewRadialGradient(geom.Pt(180, 180), 80, stops...)
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.FillRRect(geom.RectXYWH(10, 10, 330, 60), 10, render.FillWith(lin))
					cv.FillOval(geom.RectXYWH(100, 100, 160, 160), render.FillWith(rad))
					for i := 0; i < 8; i++ {
						t := float32(i) / 7
						cv.FillRect(geom.RectXYWH(10+float32(i)*42, 285, 38, 24), render.Fill(lin.Sample(t)))
					}
				}, 360, 320)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("Sample(0)=%v Sample(1)=%v IsTranslucent=%v",
					lin.Sample(0), lin.Sample(1), lin.IsTranslucent())))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewLinearGradient", Type: "func(from, to geom.Point, stops ...Stop)", Desc: "线性渐变"},
			{Name: "NewRadialGradient", Type: "func(center geom.Point, r float32, stops ...Stop)", Desc: "径向渐变"},
			{Name: "Stop", Type: "{Offset float32, Color geom.Color}", Desc: "色标"},
			{Name: "Sample", Type: "func(t float32) geom.Color", Desc: "CPU 侧采样（色标插值）"},
			{Name: "IsTranslucent", Type: "func() bool", Desc: "渐变是否含半透明"},
			{Name: "FillWith", Type: "func(Shader) Paint", Desc: "把渐变挂到 Paint"},
			{Name: "Spread", Type: "字段", Desc: "越界色扩散模式"},
		},
	})
}
