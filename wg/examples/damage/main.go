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
		ID: "damage", Title: "Damage 脏区", Group: doc.GroupVisual,
		What: []string{
			"脏区合并：相邻脏矩形合并成一个包围盒，Present 只拷贝该区域（第二个性能杠杆）",
			"遮挡剔除：被完全覆盖的矩形可以从脏区里剔除（Occlude 语义）",
			"scene.Node.DirtyRect 与 Rect.Intersect 是 composite 层做 damage 合并读的同一份几何",
		},
		Demos: []doc.Demo{
			{Title: "相邻脏矩形", Note: "三个矩形区域合并为一个包围盒提交", Build: func(ctx *ui.BuildContext) ui.Widget {
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Paint{Color: geom.RGBA(38, 42, 54)})
					for i, r := range []geom.Rect{
						geom.RectXYWH(20, 20, 140, 90),
						geom.RectXYWH(130, 60, 140, 90),
						geom.RectXYWH(300, 20, 100, 130),
					} {
						cv.FillRect(r, render.Paint{
							Color: geom.RGBAa(255, 180, 60, uint8(90+i*40)),
							Style: render.StyleStroke, StrokeWidth: 2,
						})
					}
				}, 440, 180)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return n })
			}},
			{Title: "并集与遮挡", Note: "三个真实节点的 DirtyRect 并集（绿框）；A∩B 重叠区（红框）", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				a := doc.Box(ctx, "A", geom.RGBAa(70, 130, 220, 90), 120, 90)
				b := doc.Box(ctx, "B", geom.RGBAa(230, 100, 80, 90), 120, 90)
				c := doc.Box(ctx, "C", geom.RGBAa(90, 190, 140, 90), 120, 90)
				root := scene.Stack().SetSize2(360, 170)
				a.SetStackPosition(geom.Pt(10, 10))
				b.SetStackPosition(geom.Pt(90, 60))
				c.SetStackPosition(geom.Pt(200, 20))
				a.MarkNeedsPaint()
				b.MarkNeedsPaint()
				c.MarkNeedsPaint()
				root.Add(a, b, c)
				union := root.DirtyRect()
				occluded := b.BoundsInRoot().Intersect(a.BoundsInRoot())
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(geom.RGBA(250, 250, 252)))
					for _, rc := range []geom.Rect{a.BoundsInRoot(), b.BoundsInRoot(), c.BoundsInRoot()} {
						cv.FillRect(rc, render.Fill(geom.RGBAa(120, 140, 200, 70)))
					}
					cv.StrokePath(render.NewPath().AddRect(union), render.Stroke(th.Color.Success, 2))
					cv.StrokePath(render.NewPath().AddRect(occluded), render.Stroke(th.Color.Danger, 2))
				}, 380, 190)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("脏区并集=%v 面积=%.0f 重叠=%v", union, union.W()*union.H(), !occluded.IsEmpty())))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "Node.MarkNeedsPaint", Type: "func()", Desc: "标脏节点（触发重录与脏区）"},
			{Name: "Node.DirtyRect", Type: "func() geom.Rect", Desc: "子树脏区（并集）"},
			{Name: "Rect.Intersect", Type: "func(Rect) Rect", Desc: "交集（遮挡剔除）"},
			{Name: "只重绘并集", Type: "契约", Desc: "Present 只拷贝 damage 区域（C13/C14）"},
		},
	})
}
