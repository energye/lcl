package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "canvas-primitives", Title: "画布图元", Group: doc.GroupVisual,
		What: []string{
			"render.Canvas 的全部图元：矩形/圆角矩形/椭圆/线/弧/填充路径/描边路径",
			"所有图元走同一条无句柄录制链：CustomPaintNode → DisplayList → 软件光栅",
		},
		Demos: []doc.Demo{
			{Title: "七种图元", Note: "矩形 / 圆角 / 椭圆 / 线 / 弧 / 填充路径 / 描边路径", Build: func(ctx *ui.BuildContext) ui.Widget {
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.FillRect(geom.RectXYWH(10, 10, 100, 60), render.Paint{Color: geom.RGBA(64, 120, 220)})
					cv.FillRRect(geom.RectXYWH(130, 10, 100, 60), 14, render.Paint{Color: geom.RGBA(220, 100, 90)})
					cv.FillOval(geom.RectXYWH(250, 10, 80, 60), render.Paint{Color: geom.RGBA(90, 190, 140)})
					cv.DrawLine(10, 90, 330, 90, render.Stroke(geom.RGBA(200, 200, 210), 2))
					cv.DrawArc(geom.RectXYWH(30, 110, 80, 80), 0, 1.6*3.1415926, render.Stroke(geom.RGBA(230, 170, 60), 6))
					p := render.NewPath()
					p.MoveTo(160, 110).LineTo(240, 110).LineTo(200, 190).Close()
					cv.FillPath(p, render.Paint{Color: geom.RGBA(150, 100, 220)})
					p2 := render.NewPath()
					p2.MoveTo(260, 110).LineTo(330, 180).LineTo(330, 110).Close()
					cv.StrokePath(p2, render.Stroke(geom.RGBA(60, 180, 200), 2))
				}, 360, 220)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return n })
			}},
		},
		API: []doc.APIRow{
			{Name: "FillRect/FillRRect/FillOval", Type: "func(rect, …, Paint)", Desc: "填充图元"},
			{Name: "DrawLine/DrawArc", Type: "func(…, Stroke)", Desc: "线条图元"},
			{Name: "FillPath/StrokePath", Type: "func(*Path, Paint/Stroke)", Desc: "路径图元"},
			{Name: "render.Paint", Type: "{Color, Style, StrokeWidth, Shader…}", Desc: "填充或描边样式"},
			{Name: "render.Stroke/Fill", Type: "构造助手", Desc: "带颜色/宽度构造 Paint"},
			{Name: "doc.CustomPaint", Type: "演示助手", Desc: "固定尺寸 CustomPaintNode"},
		},
	})
}
