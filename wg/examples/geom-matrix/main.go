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
		ID: "geom-matrix", Title: "几何与矩阵 API", Group: doc.GroupVisual,
		What: []string{
			"geom 的几何原语：Translate/Scale/Rotate 矩阵、Concat 组合、MapPoint/MapRect 映射",
			"Rect 的交/并/包含与 Color.Lerp——命中测试与 damage 的基础算术",
		},
		Demos: []doc.Demo{
			{Title: "矩阵映射与矩形相交", Note: "平移×旋转×缩放作用于矩形；两半透明矩形的交集描边", Build: func(ctx *ui.BuildContext) ui.Widget {
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					m := geom.Translate(180, 100).Concat(geom.Rotate(math.Pi / 6)).Concat(geom.Scale(1.3, 1.3))
					base := geom.RectXYWH(-40, -24, 80, 48)
					rp := render.NewPath().AddRect(base).Transform(m)
					cv.FillPath(rp, render.Paint{Color: geom.RGBAa(60, 130, 220, 210)})
					p1 := m.MapPoint(base.Center())
					cv.FillOval(geom.RectXYWH(p1.X-4, p1.Y-4, 8, 8), render.Fill(geom.RGBA(240, 80, 60)))
					a := geom.RectXYWH(10, 220, 120, 70)
					b := geom.RectXYWH(80, 250, 120, 70)
					cv.FillRect(a, render.Fill(geom.RGBAa(70, 130, 220, 120)))
					cv.FillRect(b, render.Fill(geom.RGBAa(230, 100, 80, 120)))
					cv.StrokePath(render.NewPath().AddRect(a.Intersect(b)), render.Stroke(geom.RGBA(40, 40, 50), 2))
				}, 360, 320)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("MapPoint(中心)=%v 交集=%v", geom.Pt(0, 0), !geom.RectXYWH(10, 220, 120, 70).Intersect(geom.RectXYWH(80, 250, 120, 70)).IsEmpty())))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "Translate/Scale/Rotate", Type: "func(…) Matrix", Desc: "构造仿射矩阵"},
			{Name: "Matrix.Concat", Type: "func(Matrix) Matrix", Desc: "矩阵组合（右乘）"},
			{Name: "Matrix.Invert", Type: "func() (Matrix, bool)", Desc: "逆矩阵"},
			{Name: "MapPoint/MapRect", Type: "映射", Desc: "点/矩形经过矩阵"},
			{Name: "Rect.Intersect/Union", Type: "集合运算", Desc: "交集/并集（damage 合并用）"},
			{Name: "Color.Lerp", Type: "func(t float32) Color", Desc: "颜色插值"},
		},
	})
}
