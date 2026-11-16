package examples

import (
	"math"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

func init() {
	// ── scene 特殊节点：PathNode / ImageNode / Clip / Transform / Opacity / Effect ──
	Register(Example{ID: "scene-nodes", Name: "场景节点构造器", Group: GroupLayout,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			tri := render.NewPath().MoveTo(0, 46).LineTo(46, 0).LineTo(92, 46).Close()
			pathN := scene.PathNode(scene.PathSpec{Path: tri, Paint: render.Fill(th.Color.Primary)}).SetSize2(100, 50)
			img := render.NewImageRef(24, 24)
			for y := 0; y < 24; y++ {
				for x := 0; x < 24; x++ {
					img.Pix[y*img.Stride+x] = geom.RGBA(uint8(x*10), uint8(y*10), 200).PremulBGRA()
				}
			}
			imageN := scene.ImageNode(scene.ImageSpec{Image: img, Fit: scene.FitFill}).SetSize2(80, 50)
			clipN := scene.Clip(scene.ClipSpec{Radius: 14}, box(ctx, "Clip 圆角", th.Color.Info, 140, 60))
			transN := scene.Transform(geom.Rotate(math.Pi/12), box(ctx, "Transform", th.Color.Warning, 130, 50))
			opN := scene.Opacity(0.45, box(ctx, "Opacity 45%", th.Color.Success, 120, 50))
			effN := scene.Effect(scene.EffectSpec{Shadow: &scene.ShadowSpec{Dx: 3, Dy: 4, Blur: 6, Color: geom.RGBAa(0, 0, 0, 90)}},
				box(ctx, "Effect 阴影", th.Color.Surface, 130, 50))
			surfN := scene.SurfaceNode(box(ctx, "SurfaceNode", th.Color.SurfaceActive, 130, 50))

			grid := scene.Wrap(scene.Horizontal, 12, 12).SetWidth(680)
			for _, n := range []*scene.Node{pathN, imageN, clipN, transN, opN, effN, surfN} {
				grid.Add(n)
			}
			return Section(ctx, "场景节点构造器", "PathNode/ImageNode/Clip/Transform/Opacity/Effect/SurfaceNode/CustomPaintNode", grid)
		}})

	// ── scene 约束与对齐 API ───────────────────────────────────
	Register(Example{ID: "scene-constraints", Name: "约束与对齐 API", Group: GroupLayout,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			tight := scene.TightConstraints(geom.Sz(120, 60))
			loose := scene.NewConstraints(40, 200, 30, 120)
			def := loose.Deflate(geom.InLTRB(10, 5, 10, 5))
			t := text.DefaultFont(th.Font.SizeSM)
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				lay := text.LayoutText("Tight 120x60", t, text.Options{})
				cv.DrawText(lay, geom.Pt(10, 6), th.Color.Text)
				cv.StrokePath(render.NewPath().AddRect(geom.RectXYWH(10, 28, tight.Constrain(geom.Sz(260, 200)).W, tight.Constrain(geom.Sz(260, 200)).H)),
					render.Stroke(th.Color.Primary, 2))
				cv.DrawText(text.LayoutText("Loose/Deflate", t, text.Options{}), geom.Pt(10, 110), th.Color.Text)
				cv.StrokePath(render.NewPath().AddRect(geom.RectXYWH(10, 132, def.Constrain(geom.Sz(260, 200)).W, def.Constrain(geom.Sz(260, 200)).H)),
					render.Stroke(th.Color.Success, 2))
			}).SetSize2(300, 220)
			col := scene.Flex(scene.Vertical, 6)
			col.Add(n)
			capNote(ctx, col, "Tight=%v Loose=%v Deflate=%v", tight.MaxSize(), loose.MaxSize(), def.MaxSize())
			return Section(ctx, "约束与对齐 API", "TightConstraints/LooseConstraints/NewConstraints/Deflate/Tighten/Loosen/IsTight/Alignment", col)
		}})

	// ── 脏区合并（scene.Node.DirtyRect 公开面）──────────────────
	Register(Example{ID: "composite-region", Name: "脏区合并与遮挡", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			// 用三个真实节点制造脏区，再由 DirtyRect 求并集：
			// 这正是 composite 层做 damage 合并时读的同一份几何。
			a := box(ctx, "A", geom.RGBAa(70, 130, 220, 90), 120, 90)
			b := box(ctx, "B", geom.RGBAa(230, 100, 80, 90), 120, 90)
			c := box(ctx, "C", geom.RGBAa(90, 190, 140, 90), 120, 90)
			root := scene.Stack().SetSize2(360, 170)
			a.SetStackPosition(geom.Pt(10, 10))
			b.SetStackPosition(geom.Pt(90, 60))
			c.SetStackPosition(geom.Pt(200, 20))
			a.MarkNeedsPaint()
			b.MarkNeedsPaint()
			c.MarkNeedsPaint()
			root.Add(a, b, c)
			union := root.DirtyRect()
			// 遮挡：被完全覆盖的矩形可以从脏区里剔除（Occlude 语义）。
			occluded := b.BoundsInRoot().Intersect(a.BoundsInRoot())
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(geom.RGBA(250, 250, 252)))
				for _, rc := range []geom.Rect{a.BoundsInRoot(), b.BoundsInRoot(), c.BoundsInRoot()} {
					cv.FillRect(rc, render.Fill(geom.RGBAa(120, 140, 200, 70)))
				}
				cv.StrokePath(render.NewPath().AddRect(union), render.Stroke(th.Color.Success, 2))
				cv.StrokePath(render.NewPath().AddRect(occluded), render.Stroke(th.Color.Danger, 2))
			}).SetSize2(380, 190)
			col := scene.Flex(scene.Vertical, 6)
			capNote(ctx, col, "脏区并集=%v 面积=%.0f 重叠=%v", union, union.W()*union.H(), !occluded.IsEmpty())
			col.Add(n)
			return Section(ctx, "脏区合并与遮挡", "DirtyRect 求并集 / Rect.Intersect 求重叠 / 只重绘并集区域", col)
		}})

	// ── 帧调度与耗时统计（BuildContext.Scheduler 公开面）──────
	Register(Example{ID: "frame-clock", Name: "帧调度与耗时统计", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			var frameNo int
			var awaiting, armed bool
			if ctx.Scheduler != nil {
				frameNo = ctx.Scheduler.Frame()
				awaiting = ctx.Scheduler.Awaiting()
				armed = ctx.Scheduler.Armed()
			}
			// 与 diag.Median/Percentile 同算法的耗时分布（示例只用公开 math）。
			deltas := []float64{8, 12, 16, 20, 24, 100}
			sorted := append([]float64(nil), deltas...)
			for i := 1; i < len(sorted); i++ {
				for j := i; j > 0 && sorted[j] < sorted[j-1]; j-- {
					sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
				}
			}
			median := sorted[len(sorted)/2]
			p90 := sorted[int(float64(len(sorted)-1)*0.9+0.5)]
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(geom.RGBA(248, 249, 252)))
				for i, d := range deltas {
					x := 10 + float32(i)*60
					h := float32(d)
					c := th.Color.Primary
					if d > 50 {
						c = th.Color.Danger
					}
					cv.FillRect(geom.RectXYWH(x, 150-h, 44, h), render.Fill(c))
				}
			}).SetSize2(390, 165)
			col := scene.Flex(scene.Vertical, 6)
			capNote(ctx, col, "帧=%d 等待=%v 已武装=%v 中位=%.0fms P90=%.0fms",
				frameNo, awaiting, armed, median, p90)
			col.Add(n)
			return Section(ctx, "帧调度与耗时统计", "Scheduler.Frame/Awaiting/Armed/RequestFrameIn + 中位/P90 耗时分布", col)
		}})

	// ── text 测量与字体 API ────────────────────────────────────
	Register(Example{ID: "text-metrics", Name: "文本测量与字体 API", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			f := text.DefaultFont(th.Font.SizeBase)
			bold := f.WithWeight(text.Bold)
			big := f.WithSize(th.Font.SizeTitle).WithFamily("Segoe UI")
			var wNormal, wBold, wBig geom.Size
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				wNormal = text.Measure("渲染引擎", f)
				wBold = text.Measure("渲染引擎", bold)
				wBig = text.Measure("渲染引擎", big)
				cv.DrawText(text.LayoutText("Regular", f, text.Options{}), geom.Pt(10, 8), th.Color.Text)
				cv.DrawText(text.LayoutText("Bold 700", bold, text.Options{}), geom.Pt(10, 40), th.Color.Text)
				cv.DrawText(text.LayoutText("Title 22", big, text.Options{}), geom.Pt(10, 74), th.Color.Primary)
				lay := text.LayoutText("第一行\n第二行\n第三行", f, text.Options{MaxWidth: 300})
				cv.DrawText(lay, geom.Pt(10, 116), th.Color.Text)
				if len(lay.Lines) > 0 {
					cv.DrawLine(10, 116+lay.LineX(0), 300, 116+lay.LineX(0), render.Stroke(th.Color.Border, 1))
				}
			}).SetSize2(350, 200)
			col := scene.Flex(scene.Vertical, 6)
			capNote(ctx, col, "Normal=%v Bold=%v Title=%v Lines=%d", wNormal, wBold, wBig, 3)
			col.Add(n)
			return Section(ctx, "文本测量与字体 API", "Measure/Font.WithWeight/WithSize/WithFamily/Weight/Align/Layout.LineX/ClearMeasureCache", col)
		}})

	// ── theme 明暗与缓动 API ───────────────────────────────────
	Register(Example{ID: "theme-curves", Name: "主题切换与缓动曲线", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			light := theme.Light()
			dark := theme.Dark()
			curves := []struct {
				name string
				fn   theme.Curve
			}{
				{"Linear", theme.Linear}, {"EaseIn", theme.EaseIn},
				{"EaseOut", theme.EaseOut}, {"EaseInOut", theme.EaseInOut},
			}
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				for i, c := range curves {
					ox, oy := float32(10+i*95), float32(10)
					cv.FillRect(geom.RectXYWH(ox, oy, 84, 84), render.Fill(geom.RGBA(248, 249, 252)))
					prev := geom.Pt(ox+6, oy+78)
					for k := 1; k <= 20; k++ {
						t := float32(k) / 20
						pt := geom.Pt(ox+6+t*72, oy+78-c.fn(t)*68)
						cv.DrawLine(prev.X, prev.Y, pt.X, pt.Y, render.Stroke(geom.RGBA(60, 110, 210), 2))
						prev = pt
					}
				}
				// 明暗 token 对照色块
				for i, c := range []geom.Color{light.Color.Primary, light.Color.Background, dark.Color.Primary, dark.Color.Background} {
					cv.FillRect(geom.RectXYWH(10+float32(i)*90, 110, 80, 40), render.Fill(c))
				}
			}).SetSize2(400, 165)
			col := scene.Flex(scene.Vertical, 6)
			capNote(ctx, col, "Light.BG=%v Dark.BG=%v Elevation[2].Blur=%.1f",
				light.Color.Background, dark.Color.Background, light.ElevationAt(2).Blur)
			col.Add(n)
			return Section(ctx, "主题切换与缓动曲线", "Light/Dark/Curve(EaseIn/Out/InOut/Linear)/ElevationAt/样式 Token", col)
		}})

	// ── ui Implicit / Tween / Style 能力 ───────────────────────
	Register(Example{ID: "ui-implicit-style", Name: "隐式动画与样式解析", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			im := ui.NewImplicit(300*time.Millisecond, theme.EaseOut, 0)
			im.To(1, 0)
			tween := ui.Tween{From: 40, To: 200}
			solid := ui.ResolveStyle(th, ui.VariantSolid, ui.StatusNormal)
			outline := ui.StyleOf(th, ui.VariantOutline, ui.SizeSmall, ui.StateHover)
			m := ui.MetricsFor(th, ui.SizeLarge)
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				cv.FillRect(geom.RectXYWH(10, 10, 340, 44), render.FillWith(nil))
				cv.FillRRect(geom.RectXYWH(10, 10, 340, 44), solid.Radius, render.Fill(solid.Background))
				cv.FillRRect(geom.RectXYWH(10, 66, 220, 44), outline.Radius,
					render.Stroke(outline.Border, outline.BorderWidth))
				cv.FillRect(geom.RectXYWH(10, 130, tween.At(0.5), 24), render.Fill(geom.RGBA(70, 130, 220)))
				cv.FillRect(geom.RectXYWH(10, 166, 340*im.Value(), 24), render.Fill(geom.RGBA(230, 120, 70)))
				_ = m
			}).SetSize2(370, 210)
			col := scene.Flex(scene.Vertical, 6)
			capNote(ctx, col, "Tween(0.5)=%.1f Implicit.Target=%.1f Solid=%v Outline=%v",
				tween.At(0.5), im.Target(), solid.Background, outline.Border)
			col.Add(n)
			return Section(ctx, "隐式动画与样式解析", "Implicit/Tween/Lerp/ResolveStyle/StyleOf/MetricsFor/StateSet/Variant", col)
		}})
}
