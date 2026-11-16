package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "effect-blur-glow", Title: "组模糊与发光", Group: doc.GroupVisual,
		What: []string{
			"EffectSpec.BlurSigma 模糊内容本身（CSS filter:blur 语义，文字同样被模糊）",
			"GlowRadius 扩散内容 alpha 并压在内容下方（发光）；组阴影基于内容 alpha 形状",
			"三条渲染顺序固定：阴影 → 发光 → 内容，内容永远不被染脏",
		},
		Demos: []doc.Demo{
			{Title: "发光与模糊", Note: "蓝色圆 GlowRadius=14；红色方块 BlurSigma=5 对照未模糊", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				dot := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary), Radius: 12}).SetSize2(56, 56)
				glow := scene.Effect(scene.EffectSpec{
					GlowRadius: 14,
					GlowColor:  geom.RGBAa(70, 130, 220, 255),
				}, dot)
				glow.SetMargin(geom.InLTRB(20, 20, 20, 20))

				blurred := scene.Effect(scene.EffectSpec{BlurSigma: 5},
					scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Danger), Radius: 8}).SetSize2(90, 70))

				row := scene.Flex(scene.Horizontal, 40).SetWidth(700)
				row.SetPadding(geom.InLTRB(10, 10, 10, 10))
				row.Add(glow)
				col := scene.Flex(scene.Vertical, 10)
				col.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Danger), Radius: 8}).SetSize2(90, 70))
				col.Add((&widgets.Label{Text: "BlurSigma=5 ↓"}).Build(ctx))
				col.Add(blurred)
				row.Add(col)
				canvas := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken)})
				canvas.SetPadding(geom.InLTRB(24, 16, 24, 24))
				canvas.Add(row)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return canvas })
			}},
		},
		API: []doc.APIRow{
			{Name: "EffectSpec.BlurSigma", Type: "float32", Desc: "内容模糊（真高斯近似，pad=3σ）"},
			{Name: "EffectSpec.GlowRadius", Type: "float32", Desc: "发光半径（内容 alpha 扩散）"},
			{Name: "EffectSpec.GlowColor", Type: "geom.Color", Default: "白", Desc: "发光颜色"},
			{Name: "EffectSpec.Shadow", Type: "*ShadowSpec", Desc: "组阴影（内容 alpha 形状，文字参与投影）"},
			{Name: "scene.Effect", Type: "func(EffectSpec, child) *Node", Desc: "组滤镜节点（离屏卷积）"},
		},
	})
}
