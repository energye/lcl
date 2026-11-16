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
		ID: "canvas-transform", Title: "变换与裁剪", Group: doc.GroupVisual,
		What: []string{
			"Canvas 状态机：Save/Restore 必须配对；裁剪与当前变换求交",
			"Translate/Rotate/Scale 叠加成当前矩阵，后续图元全部被变换",
		},
		Demos: []doc.Demo{
			{Title: "旋转条纹 + 缩放椭圆", Note: "8 次旋转叠加 + 局部裁剪；右侧 Save/Translate/Scale/Restore", Build: func(ctx *ui.BuildContext) ui.Widget {
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.Save()
					cv.ClipRect(geom.RectXYWH(10, 10, 220, 180))
					cv.Translate(120, 100)
					for i := 0; i < 8; i++ {
						cv.Rotate(0.38)
						cv.FillRect(geom.RectXYWH(-30, -8, 60, 16), render.Paint{Color: geom.RGBAa(70, 130, 220, uint8(90+i*20))})
					}
					cv.Restore()
					cv.FillRect(geom.RectXYWH(250, 20, 100, 100), render.Paint{Color: geom.RGBA(200, 90, 110)})
					cv.Save()
					cv.Translate(300, 70)
					cv.Scale(0.5, 0.5)
					cv.Translate(300, 70)
					cv.FillOval(geom.RectXYWH(0, 0, 100, 100), render.Paint{Color: geom.RGBA(255, 255, 255)})
					cv.Restore()
				}, 380, 220)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return n })
			}},
		},
		API: []doc.APIRow{
			{Name: "Save/Restore", Type: "func()", Desc: "状态栈配对（矩阵+裁剪）"},
			{Name: "ClipRect", Type: "func(geom.Rect)", Desc: "与当前裁剪求交"},
			{Name: "Translate/Rotate/Scale", Type: "func(…)", Desc: "叠加当前变换"},
			{Name: "scene.Transform", Type: "节点级变换", Desc: "不进 Canvas 状态机的节点级方案"},
		},
	})
}
