package main

import (
	"fmt"
	"math"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "render-path", Title: "路径构造 API", Group: doc.GroupVisual,
		What: []string{
			"render.Path 的完整构造面：MoveTo/LineTo/QuadTo/CubicTo 与图形直加（Rect/RRect/Oval/Circle/Arc）",
			"路径可 Clone/Transform/Flatten（曲线→折线）/VerbCount/TightBounds",
		},
		Demos: []doc.Demo{
			{Title: "曲线与图形", Note: "二次/三次贝塞尔 + 五种图形直加 + 平移克隆", Build: func(ctx *ui.BuildContext) ui.Widget {
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					p := render.NewPath()
					p.MoveTo(10, 50).QuadTo(60, 5, 110, 50).CubicTo(150, 95, 190, 5, 230, 50)
					cv.StrokePath(p, render.Stroke(geom.RGBA(70, 130, 220), 3))
					cv.FillPath(render.NewPath().AddRect(geom.RectXYWH(10, 80, 50, 36)), render.Fill(geom.RGBA(230, 100, 90)))
					cv.FillPath(render.NewPath().AddRRect(geom.RectXYWH(70, 80, 50, 36), 10), render.Fill(geom.RGBA(90, 190, 140)))
					cv.FillPath(render.NewPath().AddOval(geom.RectXYWH(130, 80, 56, 36)), render.Fill(geom.RGBA(240, 190, 70)))
					cv.FillPath(render.NewPath().AddCircle(240, 98, 18), render.Fill(geom.RGBA(150, 100, 220)))
					cv.FillPath(render.NewPath().AddArc(geom.RectXYWH(275, 80, 56, 56), 0, math.Pi*1.4), render.Fill(geom.RGBAa(60, 180, 200, 200)))
					cl := p.Clone()
					cl = cl.Transform(geom.Translate(0, 125)) // Transform 返回新路径，必须接回
					cv.StrokePath(cl, render.Stroke(geom.RGBA(40, 40, 50), 2))
					_ = cl.VerbCount()
					_ = len(cl.Flatten(0.5, nil))
				}, 350, 270)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("VerbCount/Flatten/TightBounds 可用于诊断与碰撞")))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "MoveTo/LineTo", Type: "func(x, y float32)", Desc: "直线段"},
			{Name: "QuadTo/CubicTo", Type: "func(…)", Desc: "二次/三次贝塞尔"},
			{Name: "Close", Type: "func()", Desc: "闭合子路径"},
			{Name: "AddRect/AddRRect/AddOval/AddCircle/AddArc", Type: "func(…)", Desc: "图形直加"},
			{Name: "Clone/Transform", Type: "func() / func(Matrix)", Desc: "复制与仿射变换"},
			{Name: "Flatten", Type: "func(tol float32, out []Point) []Point", Desc: "曲线扁平化为折线"},
			{Name: "VerbCount/TightBounds", Type: "诊断", Desc: "指令数 / 紧包围盒"},
		},
	})
}
