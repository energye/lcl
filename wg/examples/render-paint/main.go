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
		ID: "render-paint", Title: "画笔样式 API", Group: doc.GroupVisual,
		What: []string{
			"Paint 的派生面：Cap（端帽）/Join（拐角）/MiterLimit/FillRule/虚线/抗锯齿开关",
			"WithCap/WithJoin 等派生方法返回新 Paint，不改原值",
		},
		Demos: []doc.Demo{
			{Title: "端帽、拐角与虚线", Note: "三种 Cap × 三种 Join + EvenOdd 填充 + 虚线", Build: func(ctx *ui.BuildContext) ui.Widget {
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(geom.RGBA(250, 250, 252)))
					for i, cap := range []render.Cap{render.CapButt, render.CapRound, render.CapSquare} {
						y := float32(24 + i*34)
						cv.DrawLine(20, y, 150, y, render.Stroke(geom.RGBA(40, 90, 180), 12).WithCap(cap))
					}
					for i, j := range []render.Join{render.JoinMiter, render.JoinBevel, render.JoinRound} {
						y := float32(140 + i*44)
						p := render.NewPath().MoveTo(20, y+30).LineTo(60, y).LineTo(100, y+30)
						cv.StrokePath(p, render.Stroke(geom.RGBA(210, 90, 60), 10).WithJoin(j).WithMiterLimit(4))
					}
					star := render.NewPath().MoveTo(250, 140).LineTo(272, 205).LineTo(205, 165).LineTo(295, 165).LineTo(228, 205).Close()
					cv.FillPath(star, render.Fill(geom.RGBA(130, 80, 200)).WithFillRule(render.EvenOdd))
					cv.DrawLine(20, 290, 350, 290, render.Dashed(geom.RGBA(60, 60, 70), 2, 6, 4))
				}, 380, 310)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("Cap=%s/%s/%s Join=%s/%s/%s FillRule=%s/%s",
					render.CapButt, render.CapRound, render.CapSquare,
					render.JoinMiter, render.JoinBevel, render.JoinRound, render.NonZero, render.EvenOdd)))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "WithCap", Type: "func(Cap) Paint", Desc: "CapButt / CapRound / CapSquare"},
			{Name: "WithJoin", Type: "func(Join) Paint", Desc: "JoinMiter / JoinBevel / JoinRound"},
			{Name: "WithMiterLimit", Type: "func(float32) Paint", Desc: "尖角长度上限（超出转 Bevel）"},
			{Name: "WithFillRule", Type: "func(FillRule) Paint", Desc: "NonZero / EvenOdd"},
			{Name: "render.Dashed", Type: "func(color, width, dash, gap)", Desc: "虚线描边"},
			{Name: "WithoutAA", Type: "func() Paint", Desc: "关闭抗锯齿（像素级对照用）"},
		},
	})
}
