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
		ID: "layer-alpha", Title: "图层与半透明混合", Group: doc.GroupVisual,
		What: []string{
			"重叠 α 混合在合成阶段完成（预乘 BGRA source-over），不产生额外窗口句柄",
			"带 alpha 的 Box 直接参与混合；组级透明度用 scene.Opacity 离屏合成",
		},
		Demos: []doc.Demo{
			{Title: "递进半透明", Note: "四层 α=90 的橙层依次叠放", Build: func(ctx *ui.BuildContext) ui.Widget {
				root := scene.Stack().SetSize2(420, 180)
				root.Add(doc.Box(ctx, "不透明底", geom.RGBA(40, 50, 80), 400, 160))
				for i := 0; i < 4; i++ {
					layer := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBAa(240, 100, 60, 90))},
						(&widgets.Label{Text: "α", Color: geom.RGBA(255, 255, 255)}).Build(ctx))
					layer.SetSize2(180, 70).SetMargin(geom.InLTRB(float32(20+i*45), float32(20+i*22), 0, 0))
					root.Add(layer)
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return root })
			}},
		},
		API: []doc.APIRow{
			{Name: "geom.RGBAa", Type: "func(r,g,b,a uint8) Color", Desc: "带 alpha 颜色（预乘发生在写入像素时）"},
			{Name: "Color.PremulBGRA", Type: "func() uint32", Desc: "预乘 BGRA 像素值"},
			{Name: "scene.Opacity", Type: "组透明", Desc: "整组统一乘 alpha（文字同样生效，见 TextFlat 根治）"},
		},
	})
}
