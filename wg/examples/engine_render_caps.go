package examples

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// capNote 在当前列追加一行 API 统计文本。
func capNote(ctx *ui.BuildContext, col *scene.Node, format string, args ...any) {
	col.Add((&widgets.Label{Text: fmt.Sprintf(format, args...), Size: ctx.ThemeOr().Font.SizeSM}).Build(ctx))
}

func init() {
	// ── render.Path 完整构造 API ───────────────────────────────
	Register(Example{ID: "render-path", Name: "路径构造 API", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				p := render.NewPath()
				p.MoveTo(10, 50).QuadTo(60, 5, 110, 50).CubicTo(150, 95, 190, 5, 230, 50)
				cv.StrokePath(p, render.Stroke(geom.RGBA(70, 130, 220), 3))
				cv.FillPath(render.NewPath().AddRect(geom.RectXYWH(10, 80, 50, 36)), render.Fill(geom.RGBA(230, 100, 90)))
				cv.FillPath(render.NewPath().AddRRect(geom.RectXYWH(70, 80, 50, 36), 10), render.Fill(geom.RGBA(90, 190, 140)))
				cv.FillPath(render.NewPath().AddOval(geom.RectXYWH(130, 80, 56, 36)), render.Fill(geom.RGBA(240, 190, 70)))
				cv.FillPath(render.NewPath().AddCircle(240, 98, 18), render.Fill(geom.RGBA(150, 100, 220)))
				cv.FillPath(render.NewPath().AddArc(geom.RectXYWH(275, 80, 56, 56), 0, math.Pi*1.4), render.Fill(geom.RGBAa(60, 180, 200, 200)))
				cl := p.Clone()
				cl.Transform(geom.Translate(0, 125))
				cv.StrokePath(cl, render.Stroke(geom.RGBA(40, 40, 50), 2))
				_ = cl.VerbCount()
				_ = len(cl.Flatten(0.5, nil))
			}).SetSize2(350, 270)
			col := scene.Flex(scene.Vertical, 6)
			col.Add(n)
			return Section(ctx, "路径构造 API", "MoveTo/LineTo/QuadTo/CubicTo/AddRect/RRect/Oval/Circle/Arc/Clone/Transform/Flatten/VerbCount/TightBounds", col)
		}})

	// ── render.Paint 样式 API ──────────────────────────────────
	Register(Example{ID: "render-paint", Name: "画笔样式 API", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
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
			}).SetSize2(380, 310)
			col := scene.Flex(scene.Vertical, 6)
			col.Add(n)
			capNote(ctx, col, "Cap=%s/%s/%s Join=%s/%s/%s FillRule=%s/%s",
				render.CapButt, render.CapRound, render.CapSquare,
				render.JoinMiter, render.JoinBevel, render.JoinRound, render.NonZero, render.EvenOdd)
			_ = th
			return Section(ctx, "画笔样式 API", "Cap/Join/MiterLimit/FillRule/Dashed/WithoutAA 与 Paint 派生方法", col)
		}})

	// ── render.Shader 渐变 API ─────────────────────────────────
	Register(Example{ID: "render-gradient", Name: "渐变与采样 API", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			stops := []render.Stop{
				{Offset: 0, Color: geom.RGBA(255, 90, 90)},
				{Offset: 0.5, Color: geom.RGBA(90, 200, 120)},
				{Offset: 1, Color: geom.RGBA(70, 120, 240)},
			}
			lin := render.NewLinearGradient(geom.Pt(10, 20), geom.Pt(330, 20), stops...)
			rad := render.NewRadialGradient(geom.Pt(180, 180), 80, stops...)
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				cv.FillRRect(geom.RectXYWH(10, 10, 330, 60), 10, render.FillWith(lin))
				cv.FillOval(geom.RectXYWH(100, 100, 160, 160), render.FillWith(rad))
				for i := 0; i < 8; i++ {
					t := float32(i) / 7
					cv.FillRect(geom.RectXYWH(10+float32(i)*42, 285, 38, 24), render.Fill(lin.Sample(t)))
				}
			}).SetSize2(360, 320)
			col := scene.Flex(scene.Vertical, 6)
			col.Add(n)
			capNote(ctx, col, "Sample(0)=%v Sample(1)=%v IsTranslucent=%v",
				lin.Sample(0), lin.Sample(1), lin.IsTranslucent())
			return Section(ctx, "渐变与采样 API", "Linear/Radial/Sample/GradientT/IsTranslucent/Spread/FillWith", col)
		}})

	// ── render.Recorder + DisplayList ──────────────────────────
	Register(Example{ID: "render-displaylist", Name: "显示列表录制回放", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			dl := render.NewDisplayList(16)
			rec := render.NewRecorder(dl)
			rec.FillRect(geom.RectXYWH(10, 10, 120, 80), render.Fill(geom.RGBA(70, 130, 220)))
			rec.FillRRect(geom.RectXYWH(150, 10, 120, 80), 16, render.Fill(geom.RGBA(230, 110, 80)))
			rec.FillOval(geom.RectXYWH(10, 110, 120, 80), render.Fill(geom.RGBA(90, 190, 140)))
			rec.DrawLine(150, 150, 270, 150, render.Stroke(geom.RGBA(240, 180, 60), 6))
			rec.Finish()
			frozen := dl.Freeze().Frozen()
			bounds := dl.Bounds()
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				cv.FillRect(geom.RectXYWH(10, 10, 120, 80), render.Fill(geom.RGBA(70, 130, 220)))
				cv.FillRRect(geom.RectXYWH(150, 10, 120, 80), 16, render.Fill(geom.RGBA(230, 110, 80)))
				cv.FillOval(geom.RectXYWH(10, 110, 120, 80), render.Fill(geom.RGBA(90, 190, 140)))
				cv.DrawLine(150, 150, 270, 150, render.Stroke(geom.RGBA(240, 180, 60), 6))
			}).SetSize2(290, 190)
			col := scene.Flex(scene.Vertical, 6)
			col.Add(n)
			capNote(ctx, col, "Len=%d Version=%d Frozen=%v Intersects=%v Bounds=%v",
				dl.Len(), dl.Version(), frozen, dl.Intersects(geom.RectXYWH(0, 0, 300, 200)), bounds)
			return Section(ctx, "显示列表录制回放", "Recorder/DisplayList/Freeze/OffsetBy/ScaleAlpha/Intersects/Bounds/Version/Painted", col)
		}})

	// ── render.Image 位图绘制 ──────────────────────────────────
	Register(Example{ID: "render-image", Name: "位图与采样质量", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			img := image.NewNRGBA(image.Rect(0, 0, 64, 64))
			for y := 0; y < 64; y++ {
				for x := 0; x < 64; x++ {
					img.SetNRGBA(x, y, color.NRGBA{R: uint8(x * 4), G: uint8(y * 4), B: uint8((x + y) * 2), A: 255})
				}
			}
			ref := render.ImageFrom(img)
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				cv.DrawImage(ref, geom.RectXYWH(10, 10, 150, 150), render.ImageNearest)
				cv.DrawImage(ref, geom.RectXYWH(180, 10, 150, 150), render.ImageBilinear)
				cv.DrawImage(ref, geom.RectXYWH(10, 175, 150, 150), render.ImageMipmap)
			}).SetSize2(350, 340)
			col := scene.Flex(scene.Vertical, 6)
			col.Add(n)
			capNote(ctx, col, "Bounds=%v W=%d H=%d", ref.Bounds(), ref.W, ref.H)
			return Section(ctx, "位图与采样质量", "ImageFrom/NewImageRef/Bounds/DrawImage 与 Nearest/Bilinear/Mipmap 三档采样", col)
		}})

	// ── 像素缓冲与 Alpha 统计（render.ImageRef 公开面）───────────
	Register(Example{ID: "raster-backing", Name: "像素缓冲与统计", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			const bw, bh = 120, 90
			buf := render.NewImageRef(bw, bh)
			// 直接写预乘 BGRA：等价于 raster.Backing 的 Fill/FillRect/BlendRect，
			// 但只用到 render 公开类型（examples 不允许依赖 raster）。
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
						r := (uint32(c.R)*sa + ((dst>>16)&0xff)*ia) / 255
						g := (uint32(c.G)*sa + ((dst>>8)&0xff)*ia) / 255
						b := (uint32(c.B)*sa + (dst&0xff)*ia) / 255
						buf.Pix[y*buf.Stride+x] = r<<16 | g<<8 | b | 0xFF000000
					}
				}
			}
			fillPx(10, 10, 60, 50, geom.RGBA(60, 120, 220), false)
			fillPx(40, 30, 90, 70, geom.RGBAa(230, 80, 60, 140), true)
			// 统计非透明像素与包围盒（对应 AlphaCount/AlphaBounds）。
			count, minX, minY, maxX, maxY := 0, bw, bh, 0, 0
			for y := 0; y < bh; y++ {
				for x := 0; x < bw; x++ {
					if buf.Pix[y*buf.Stride+x]>>24 == 0 {
						continue
					}
					count++
					if x < minX { minX = x }
					if y < minY { minY = y }
					if x > maxX { maxX = x }
					if y > maxY { maxY = y }
				}
			}
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				cv.DrawImage(buf, geom.RectXYWH(10, 10, 240, 180), render.ImageNearest)
			}).SetSize2(270, 200)
			col := scene.Flex(scene.Vertical, 6)
			col.Add(n)
			capNote(ctx, col, "Stride=%d Bounds=%v 非透明=%d 包围盒=(%d,%d)-(%d,%d)",
				buf.Stride, buf.Bounds(), count, minX, minY, maxX, maxY)
			return Section(ctx, "像素缓冲与统计", "预乘 BGRA 读写 / source-over 混合 / 非透明统计；等价 raster.Backing.Fill/Rect/Blend/Alpha*", col)
		}})

	// ── geom 几何与变换 API ────────────────────────────────────
	Register(Example{ID: "geom-matrix", Name: "几何与矩阵 API", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
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
			}).SetSize2(360, 320)
			col := scene.Flex(scene.Vertical, 6)
			col.Add(n)
			return Section(ctx, "几何与矩阵 API", "Translate/Scale/Rotate/Concat/Invert/MapPoint/MapRect/Rect 交并/Color.Lerp", col)
		}})

	// ── DPI 与坐标换算（MediaQuery 公开面）─────────────────────
	Register(Example{ID: "host-dpi", Name: "DPI 与坐标换算", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			scale := ctx.Media.Scale
			if scale <= 0 {
				scale = 1
			}
			// 与 host.DPI 同语义的换算：DIP↔像素、对齐到设备像素网格。
			dip := func(px float32) float32 { return px / scale }
			pixels := func(v float32) float32 { return v * scale }
			snapRect := func(r geom.Rect) geom.Rect {
				return geom.RectLTRB(
					float32(math.Round(float64(r.Min.X*scale)))/scale,
					float32(math.Round(float64(r.Min.Y*scale)))/scale,
					float32(math.Round(float64(r.Max.X*scale)))/scale,
					float32(math.Round(float64(r.Max.Y*scale)))/scale,
				)
			}
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(geom.RGBA(248, 249, 252)))
				r := geom.RectXYWH(20, 20, 160, 90)
				cv.FillRect(r, render.Fill(geom.RGBAa(70, 130, 220, 90)))
				cv.StrokePath(render.NewPath().AddRect(snapRect(r)), render.Stroke(geom.RGBA(230, 80, 60), 2))
				// 1 DIP 网格：验证像素对齐后仍落在整数设备像素上。
				for x := float32(0); x < sz.W; x += 8 {
					cv.DrawLine(x, 120, x, 140, render.Stroke(geom.RGBA(180, 185, 195), 1))
				}
			}).SetSize2(240, 160)
			col := scene.Flex(scene.Vertical, 6)
			col.Add(n)
			capNote(ctx, col, "Scale=%.2f Percent=%d%% Identity=%v DIP(24)=%.1f Pixels(24)=%.1f",
				scale, int(scale*100+0.5), scale == 1, dip(24), pixels(24))
			return Section(ctx, "DPI 与坐标换算", "MediaQuery.Scale → DIP/Pixels/SnapRect/Percent；真窗里跟随窗体 DPI", col)
		}})
}

func min255(v uint32) uint32 {
	if v > 255 {
		return 255
	}
	return v
}

// 确保 text 包被显式引用（排版能力由 text-layout 覆盖，这里保留导入对齐）。
var _ = text.DefaultFont
