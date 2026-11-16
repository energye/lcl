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
		ID: "clip-path", Title: "路径裁剪", Group: doc.GroupVisual,
		What: []string{
			"scene.Clip 的三种裁剪：矩形=scissor 零成本；圆角/路径=蒙版精确裁剪（AA 由蒙版给）",
			"CustomPaint 内的 ClipPath 也已是精确裁剪（4×4 超采样解析蒙版）",
			"裁剪对组内所有内容生效，包括文字（见 TextFlat 根治后的语义）",
		},
		Demos: []doc.Demo{
			{Title: "矩形 / 圆角 / 三角", Note: "同一块条纹内容，三种裁剪", Build: func(ctx *ui.BuildContext) ui.Widget {
				stripes := func() *scene.Node {
					return scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
						for i := 0; i < 20; i++ {
							cv.FillRect(geom.RectXYWH(0, float32(i)*10, sz.W, 5),
								render.Fill(geom.RGBAa(70, 130, 220, 200)))
						}
					}).SetSize2(160, 160)
				}
				plain := scene.Clip(scene.ClipSpec{Rect: geom.RectXYWH(0, 0, 160, 160)}, stripes())
				rrect := scene.Clip(scene.ClipSpec{Rect: geom.RectXYWH(0, 0, 160, 160), Radius: 40, AA: true}, stripes())
				tri := scene.Clip(scene.ClipSpec{
					Rect: geom.RectXYWH(0, 0, 160, 160),
					Path: render.NewPath().MoveTo(80, 6).LineTo(154, 154).LineTo(6, 154).Close(),
					AA:   true,
				}, stripes())
				row := scene.Flex(scene.Horizontal, 20)
				row.Add(plain)
				row.Add(rrect)
				row.Add(tri)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
		},
		API: []doc.APIRow{
			{Name: "ClipSpec.Rect", Type: "geom.Rect", Desc: "裁剪包围盒（必填）"},
			{Name: "ClipSpec.Radius", Type: "float32", Desc: "圆角半径（>0 走蒙版）"},
			{Name: "ClipSpec.Path", Type: "*render.Path", Desc: "任意路径裁剪（走蒙版）"},
			{Name: "ClipSpec.AA", Type: "bool", Desc: "蒙版抗锯齿开关"},
			{Name: "Canvas.ClipRect", Type: "录制期裁剪", Desc: "scissor 级，零成本"},
			{Name: "Canvas.ClipPath", Type: "录制期裁剪", Desc: "精确路径裁剪"},
		},
	})
}
