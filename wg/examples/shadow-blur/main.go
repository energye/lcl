package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "shadow-blur", Title: "阴影与层级", Group: doc.GroupVisual,
		What: []string{
			"盒阴影：Blur=0 走硬阴影（零成本）；Blur>0 走离屏三次 box 卷积，约等于同 σ 高斯",
			"阴影形状来自节点矩形+圆角（不是内容 alpha），所以卡片上的文字保持锐利",
			"theme.Data.ElevationAt(i) 给出 Dx/Dy/Blur/Color——组件不硬编码阴影，换主题即换阴影",
		},
		Demos: []doc.Demo{
			{Title: "Blur 递进", Note: "0/4/10/20 的软硬差异；浅色衬底让阴影可见", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				row := scene.Flex(scene.Horizontal, 24).SetWidth(700).SetHeight(170)
				for _, blur := range []float32{0, 4, 10, 20} {
					card := scene.Box(scene.BoxSpec{
						Fill:   scene.Solid(th.Color.Surface),
						Radius: 10,
						Shadow: &scene.ShadowSpec{Dx: 0, Dy: 6, Blur: blur, Color: geom.RGBAa(10, 20, 45, 190)},
					}, (&widgets.Label{Text: fmt.Sprintf("Blur=%.0f", blur)}).Build(ctx)).SetSize2(150, 110)
					card.SetMargin(geom.InLTRB(0, 24, 0, 0))
					row.Add(card)
				}
				canvas := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken)})
				canvas.SetPadding(geom.InLTRB(24, 8, 24, 24))
				canvas.Add(row)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return canvas })
			}},
			{Title: "软阴影卡片层叠", Note: "阴影随层递增；每张卡片是一个滤镜层，静止时随层缓存复用", Build: func(ctx *ui.BuildContext) ui.Widget {
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
				return doc.Node(func(*ui.BuildContext) *scene.Node { return canvas })
			}},
			{Title: "Elevation Token", Note: "theme 的层级阴影阶梯直接驱动滤镜层", Build: func(ctx *ui.BuildContext) ui.Widget {
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
				return doc.Node(func(*ui.BuildContext) *scene.Node { return grid })
			}},
		},
		API: []doc.APIRow{
			{Name: "BoxSpec.Shadow", Type: "*scene.ShadowSpec{Dx,Dy,Blur,Color}", Desc: "盒阴影参数"},
			{Name: "Blur=0", Type: "硬阴影", Desc: "零成本路径：无离屏、无卷积"},
			{Name: "Blur>0", Type: "滤镜层", Desc: "离屏 + 三次 box 卷积 ≈ 同 σ 高斯"},
			{Name: "theme.ElevationAt", Type: "func(i int) ShadowSpec", Desc: "层级阴影 token"},
			{Name: "组阴影/发光", Type: "EffectSpec", Desc: "基于内容 alpha，见 effect-blur-glow 窗"},
		},
	})
}
