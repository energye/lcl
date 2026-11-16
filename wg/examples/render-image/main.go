package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "render-image", Title: "位图与采样质量", Group: doc.GroupVisual,
		What: []string{
			"DrawImage 三档采样：Nearest 硬边缘 / Bilinear 双线性 / Mipmap 2×2 盒式预滤（近似，非真 mip 金字塔）",
			"ImageFrom 包装标准库 image；NewImageRef 直接持有预乘 BGRA 像素",
		},
		Demos: []doc.Demo{
			{Title: "三档采样对照", Note: "64×64 渐变图放大 150×150：Nearest / Bilinear / Mipmap", Build: func(ctx *ui.BuildContext) ui.Widget {
				img := image.NewNRGBA(image.Rect(0, 0, 64, 64))
				for y := 0; y < 64; y++ {
					for x := 0; x < 64; x++ {
						img.SetNRGBA(x, y, color.NRGBA{R: uint8(x * 4), G: uint8(y * 4), B: uint8((x + y) * 2), A: 255})
					}
				}
				ref := render.ImageFrom(img)
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.DrawImage(ref, geom.RectXYWH(10, 10, 150, 150), render.ImageNearest)
					cv.DrawImage(ref, geom.RectXYWH(180, 10, 150, 150), render.ImageBilinear)
					cv.DrawImage(ref, geom.RectXYWH(10, 175, 150, 150), render.ImageMipmap)
				}, 350, 340)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("Bounds=%v W=%d H=%d", ref.Bounds(), ref.W, ref.H)))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "ImageFrom", Type: "func(image.Image) *ImageRef", Desc: "标准库图像 → 预乘 BGRA"},
			{Name: "NewImageRef", Type: "func(w, h int) *ImageRef", Desc: "直接分配（Pix/Stride 可写）"},
			{Name: "DrawImage", Type: "func(ref, rect, quality)", Desc: "位图绘制；缩放采样"},
			{Name: "ImageNearest/Bilinear/Mipmap", Type: "ImageQuality", Desc: "三档采样；Mipmap 为 2×2 预滤近似"},
		},
	})
}
