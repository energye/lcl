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
		ID: "raster-backing", Title: "像素缓冲与统计", Group: doc.GroupVisual,
		What: []string{
			"预乘 BGRA 像素缓冲的读写与 source-over 混合——raster.Backing 的公开面在 render.ImageRef 上的等价演示",
			"非透明像素统计与包围盒：等价 raster.Backing 的 AlphaCount/AlphaBounds（damage 与命中几何的基础）",
		},
		Demos: []doc.Demo{
			{Title: "手写像素 + 混合", Note: "直接写预乘 BGRA：不透明蓝块 + α=140 红块混合", Build: func(ctx *ui.BuildContext) ui.Widget {
				const bw, bh = 120, 90
				buf := render.NewImageRef(bw, bh)
				for i := range buf.Pix {
					buf.Pix[i] = geom.RGBA(245, 245, 248).PremulBGRA()
				}
				fillPx := func(x0, y0, x1, y1 int, c geom.Color, blend bool) {
					for y := y0; y < y1; y++ {
						for x := x0; x < x1; x++ {
							if x < 0 || y < 0 || x >= bw || y >= bh {
								continue
							}
							dst := buf.Pix[y*buf.Stride+x]
							if !blend {
								buf.Pix[y*buf.Stride+x] = c.PremulBGRA()
								continue
							}
							sa := uint32(c.A)
							ia := 255 - sa
							r := (uint32(c.R)*sa + ((dst >> 16) & 0xff) * ia) / 255
							g := (uint32(c.G)*sa + ((dst >> 8) & 0xff) * ia) / 255
							b := (uint32(c.B)*sa + (dst & 0xff) * ia) / 255
							buf.Pix[y*buf.Stride+x] = r<<16 | g<<8 | b | 0xFF000000
						}
					}
				}
				fillPx(10, 10, 60, 50, geom.RGBA(60, 120, 220), false)
				fillPx(40, 30, 90, 70, geom.RGBAa(230, 80, 60, 140), true)
				count, minX, minY, maxX, maxY := 0, bw, bh, 0, 0
				for y := 0; y < bh; y++ {
					for x := 0; x < bw; x++ {
						if buf.Pix[y*buf.Stride+x]>>24 == 0 {
							continue
						}
						count++
						if x < minX {
							minX = x
						}
						if y < minY {
							minY = y
						}
						if x > maxX {
							maxX = x
						}
						if y > maxY {
							maxY = y
						}
					}
				}
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.DrawImage(buf, geom.RectXYWH(10, 10, 240, 180), render.ImageNearest)
				}, 270, 200)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("Stride=%d Bounds=%v 非透明=%d 包围盒=(%d,%d)-(%d,%d)",
					buf.Stride, buf.Bounds(), count, minX, minY, maxX, maxY)))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "ImageRef.Pix", Type: "[]uint32", Desc: "预乘 BGRA 像素（0xAARRGGBB 布局）"},
			{Name: "ImageRef.Stride", Type: "int", Desc: "行字节数（w*4），不要再乘 4"},
			{Name: "source-over 混合", Type: "out = s + d·(1-a)", Desc: "预乘空间直接算，无需反预乘"},
			{Name: "AlphaCount/AlphaBounds", Type: "等价能力", Desc: "raster.Backing 的非透明统计（damage/命中用）"},
		},
	})
}
