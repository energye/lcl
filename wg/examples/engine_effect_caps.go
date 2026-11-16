package examples

import (
	"fmt"
	"math"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func init() {
	// ── 阴影模糊 ───────────────────────────────────────────────
	Register(Example{ID: "shadow-blur", Name: "阴影模糊", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			row := scene.Flex(scene.Horizontal, 24).SetWidth(700).SetHeight(170)
			// 同一形状依次加大 Blur，肉眼就能看出软硬差异。
			for i, blur := range []float32{0, 4, 10, 20} {
				card := scene.Box(scene.BoxSpec{
					Fill:   scene.Solid(th.Color.Surface),
					Radius: 10,
					Shadow: &scene.ShadowSpec{
						Dx: 0, Dy: 6, Blur: blur,
						Color: geom.RGBAa(10, 20, 45, 190),
					},
				}, (&widgets.Label{Text: fmt.Sprintf("Blur=%.0f", blur)}).Build(ctx)).SetSize2(150, 110)
				card.SetMargin(geom.InLTRB(0, 24, 0, 0))
				row.Add(card)
				_ = i
			}
			// 浅色衬底：白卡 + 白底会让阴影几乎不可见，看不见就谈不上验收。
			canvas := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken)})
			canvas.SetPadding(geom.InLTRB(24, 8, 24, 24))
			canvas.Add(row)
			return Section(ctx, "阴影模糊", "Blur=0 走硬阴影（零成本）；Blur>0 走离屏三次 box 卷积，约等于同 σ 高斯", canvas)
		}})

	// ── 组模糊 / 发光 ──────────────────────────────────────────
	Register(Example{ID: "effect-blur-glow", Name: "组模糊与发光", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			dot := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary), Radius: 12}).SetSize2(56, 56)
			glow := scene.Effect(scene.EffectSpec{
				GlowRadius: 14,
				GlowColor:  geom.RGBAa(70, 130, 220, 255),
			}, dot)
			glow.SetMargin(geom.InLTRB(20, 20, 20, 20))

			sharp := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Danger), Radius: 8}).SetSize2(90, 70)
			blurred := scene.Effect(scene.EffectSpec{BlurSigma: 5},
				scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Danger), Radius: 8}).SetSize2(90, 70))

			row := scene.Flex(scene.Horizontal, 40).SetWidth(700)
			row.SetPadding(geom.InLTRB(10, 10, 10, 10))
			row.Add(glow)
			col := scene.Flex(scene.Vertical, 10)
			col.Add(sharp)
			col.Add((&widgets.Label{Text: "BlurSigma=5 ↓"}).Build(ctx))
			col.Add(blurred)
			row.Add(col)
			canvas := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken)})
			canvas.SetPadding(geom.InLTRB(24, 16, 24, 24))
			canvas.Add(row)
			return Section(ctx, "组模糊与发光", "EffectSpec.BlurSigma 模糊内容本身；GlowRadius 扩散 alpha 并压在内容下方", canvas)
		}})

	// ── 路径裁剪 ───────────────────────────────────────────────
	Register(Example{ID: "clip-path", Name: "路径裁剪", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			// 同一块条纹内容，分别用矩形、圆角、三角形裁剪。
			stripes := From(func(c *ui.BuildContext) *scene.Node {
				return scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
					for i := 0; i < 20; i++ {
						cv.FillRect(geom.RectXYWH(0, float32(i)*10, sz.W, 5),
							render.Fill(geom.RGBAa(70, 130, 220, 200)))
					}
				}).SetSize2(160, 160)
			})
			plain := scene.Clip(scene.ClipSpec{Rect: geom.RectXYWH(0, 0, 160, 160)}, stripes.Build(ctx))
			rrect := scene.Clip(scene.ClipSpec{Rect: geom.RectXYWH(0, 0, 160, 160), Radius: 40, AA: true}, stripes.Build(ctx))
			tri := scene.Clip(scene.ClipSpec{
				Rect: geom.RectXYWH(0, 0, 160, 160),
				Path: render.NewPath().MoveTo(80, 6).LineTo(154, 154).LineTo(6, 154).Close(),
				AA:   true,
			}, stripes.Build(ctx))

			row := scene.Flex(scene.Horizontal, 20)
			row.Add(plain)
			row.Add(rrect)
			row.Add(tri)
			col := scene.Flex(scene.Vertical, 8)
			col.Add(row)
			capNote(ctx, col, "矩形=scissor 零成本；圆角/路径=蒙版精确裁剪（AA 由蒙版给）")
			_ = th
			return Section(ctx, "路径裁剪", "scene.Clip 的 Radius/Path 走蒙版；CustomPaint 内 ClipPath 也已是精确裁剪", col)
		}})

	// ── 软阴影卡片（组件级组合）────────────────────────────────
	Register(Example{ID: "shadow-card-stack", Name: "软阴影卡片层叠", Group: GroupDisplay,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			stack := scene.Stack().SetSize2(560, 280)
			for i := 0; i < 5; i++ {
				card := scene.Box(scene.BoxSpec{
					Fill:   scene.Solid(th.Color.Surface),
					Radius: 12,
					Shadow: &scene.ShadowSpec{
						Dx: 0, Dy: float32(4 + i*2), Blur: float32(6 + i*3),
						Color: geom.RGBAa(10, 20, 45, 170),
					},
				}, (&widgets.Label{
					Text:  fmt.Sprintf("层 %d · blur=%.0f", i+1, float32(6+i*3)),
					Size:  th.Font.SizeBase,
					Color: th.Color.Text,
				}).Build(ctx)).SetSize2(300, 70)
				card.SetStackPosition(geom.Pt(float32(20+i*32), float32(20+i*34)))
				stack.Add(card)
			}
			canvas := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken)})
			canvas.SetPadding(geom.InLTRB(24, 8, 24, 24))
			canvas.Add(stack)
			return Section(ctx, "软阴影卡片层叠", "阴影随层递增；每张卡片是一个滤镜层，静止时随层缓存复用", canvas)
		}})

	// ── 阴影/发光/模糊与主题联动 ───────────────────────────────
	Register(Example{ID: "effect-theme-tokens", Name: "阴影 Token 与层级", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			grid := scene.Grid(3, 14).SetWidth(620)
			for i := 0; i < 6; i++ {
				sp := th.ElevationAt(i)
				card := scene.Box(scene.BoxSpec{
					Fill:   scene.Solid(th.Color.Surface),
					Radius: 10,
					Shadow: &scene.ShadowSpec{Dx: sp.Dx, Dy: sp.Dy, Blur: sp.Blur, Color: sp.Color},
				}, (&widgets.Label{Text: fmt.Sprintf("Elevation %d", i), Size: th.Font.SizeSM}).Build(ctx)).SetSize2(190, 90)
				grid.Add(card)
			}
			col := scene.Flex(scene.Vertical, 8)
			col.Add(grid)
			capNote(ctx, col, "theme.Data.ElevationAt(i) 给出 Dx/Dy/Blur/Color，组件不硬编码阴影")
			return Section(ctx, "阴影 Token 与层级", "ElevationAt 的 Blur 直接驱动滤镜层；换主题即换阴影", col)
		}})

	// 引用，避免 text/theme/math 未使用（示例按需保留）。
	_ = text.DefaultFont
	_ = theme.EaseOut
	_ = math.Pi
}
