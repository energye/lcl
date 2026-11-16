package examples

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func init() {
	Register(Example{ID: "canvas-primitives", Name: "画布图元", Group: GroupVisual, Build: func(ctx *ui.BuildContext) *scene.Node {
		n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
			cv.FillRect(geom.RectXYWH(10, 10, 100, 60), render.Paint{Color: geom.RGBA(64, 120, 220)})
			cv.FillRRect(geom.RectXYWH(130, 10, 100, 60), 14, render.Paint{Color: geom.RGBA(220, 100, 90)})
			cv.FillOval(geom.RectXYWH(250, 10, 80, 60), render.Paint{Color: geom.RGBA(90, 190, 140)})
			cv.DrawLine(10, 90, 330, 90, render.Stroke(geom.RGBA(200, 200, 210), 2))
			cv.DrawArc(geom.RectXYWH(30, 110, 80, 80), 0, 1.6*3.1415926, render.Stroke(geom.RGBA(230, 170, 60), 6))
			p := render.NewPath()
			p.MoveTo(160, 110).LineTo(240, 110).LineTo(200, 190).Close()
			cv.FillPath(p, render.Paint{Color: geom.RGBA(150, 100, 220)})
			p2 := render.NewPath()
			p2.MoveTo(260, 110).LineTo(330, 180).LineTo(330, 110).Close()
			cv.StrokePath(p2, render.Stroke(geom.RGBA(60, 180, 200), 2))
		}).SetSize2(360, 220)
		return Section(ctx, "画布图元", "矩形/圆角/椭圆/线/弧/填充路径/描边路径，全部走同一 Canvas", n)
	}})

	Register(Example{ID: "canvas-transform", Name: "变换与裁剪", Group: GroupVisual, Build: func(ctx *ui.BuildContext) *scene.Node {
		n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
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
		}).SetSize2(380, 220)
		return Section(ctx, "变换与裁剪", "Save/Restore 配对；裁剪与当前变换求交", n)
	}})

	Register(Example{ID: "text-layout", Name: "文本排版", Group: GroupVisual, Build: func(ctx *ui.BuildContext) *scene.Node {
		th := ctx.ThemeOr()
		col := scene.Flex(scene.Vertical, 10).SetWidth(520)
		samples := []struct {
			text     string
			maxWidth float32
			maxLines int
			ellipsis string
		}{
			{"中文与 English 混排，偏移统一按 rune 计算", 420, 1, ""},
			{"这一行故意很长很长很长很长很长很长很长很长很长很长，末尾应该出现省略号", 420, 1, "…"},
			{"第一行\n第二行\n第三行", 420, 3, ""},
		}
		for _, s := range samples {
			n := scene.Text(scene.TextSpec{
				Src: s.text,
				Layout: text.LayoutText(s.text, text.DefaultFont(th.Font.SizeBase), text.Options{
					MaxWidth: s.maxWidth, MaxLines: s.maxLines, Ellipsis: s.ellipsis,
				}),
				Color: th.Color.Text,
			})
			col.Add(n)
		}
		return Section(ctx, "文本排版", "MaxWidth/MaxLines/Ellipsis；UTF-16 只在 host 层换算", col)
	}})

	Register(Example{ID: "layer-alpha", Name: "图层与半透明混合", Group: GroupVisual, Build: func(ctx *ui.BuildContext) *scene.Node {
		root := scene.Stack().SetSize2(420, 180)
		root.Add(box(ctx, "不透明底", geom.RGBA(40, 50, 80), 400, 160))
		for i := 0; i < 4; i++ {
			layer := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBAa(240, 100, 60, 90))},
				(&widgets.Label{Text: "α", Color: geom.RGBA(255, 255, 255)}).Build(ctx))
			layer.SetSize2(180, 70).SetMargin(geom.InLTRB(float32(20+i*45), float32(20+i*22), 0, 0))
			root.Add(layer)
		}
		return Section(ctx, "图层与半透明", "重叠 α 混合在合成阶段完成，不产生额外窗口句柄", root)
	}})

	Register(Example{ID: "theme-tokens", Name: "主题 Token", Group: GroupVisual, Build: func(ctx *ui.BuildContext) *scene.Node {
		th := ctx.ThemeOr()
		grid := scene.Grid(2, 8).SetWidth(520)
		pairs := []struct {
			name string
			c    geom.Color
		}{
			{"Primary", th.Color.Primary}, {"Success", th.Color.Success},
			{"Warning", th.Color.Warning}, {"Danger", th.Color.Danger},
			{"Surface", th.Color.Surface}, {"SurfaceSunken", th.Color.SurfaceSunken},
			{"Border", th.Color.Border}, {"TextMuted", th.Color.TextMuted},
		}
		for _, p := range pairs {
			cell := scene.Flex(scene.Horizontal, 8).SetHeight(40).SetFlex(1, scene.FlexTight)
			cell.SetBox(scene.BoxSpec{Fill: scene.Solid(p.c)})
			cell.Add((&widgets.Label{Text: p.name}).Build(ctx))
			grid.Add(cell)
		}
		return Section(ctx, "主题 Token", "组件只读 theme.Data，不硬编码颜色", grid)
	}})
}
