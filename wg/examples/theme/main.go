package main

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "theme", Title: "主题 Token", Group: doc.GroupVisual,
		What: []string{
			"组件只读 theme.Data，不硬编码颜色（唯一色源纪律）；换肤 = 换一个 *Data",
			"M7 值级对齐 antd v5：Primary #1677FF、Text #1F1F1F、Border #D9D9D9、控件高 24/32/40",
			"antd 的 rgba 语义色换算为白底实色等价（GDI 文字通道丢 alpha 的物理约束）",
		},
		Demos: []doc.Demo{
			{Title: "语义色板", Note: "组件只读 theme.Data；暗色主题同键换值", Build: func(ctx *ui.BuildContext) ui.Widget {
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
					cell.Add(doc.Note(ctx, p.name))
					grid.Add(cell)
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return grid })
			}},
			{Title: "明暗主题与缓动曲线", Note: "Light/Dark 对照 + 四条缓动曲线", Build: func(ctx *ui.BuildContext) ui.Widget {
				light := theme.Light()
				dark := theme.Dark()
				curves := []struct {
					name string
					fn   theme.Curve
				}{
					{"Linear", theme.Linear}, {"EaseIn", theme.EaseIn},
					{"EaseOut", theme.EaseOut}, {"EaseInOut", theme.EaseInOut},
				}
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
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
					for i, c := range []geom.Color{light.Color.Primary, light.Color.Background, dark.Color.Primary, dark.Color.Background} {
						cv.FillRect(geom.RectXYWH(10+float32(i)*90, 110, 80, 40), render.Fill(c))
					}
				}, 400, 165)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("Light.BG=%v Dark.BG=%v Elevation[2].Blur=%.1f",
					light.Color.Background, dark.Color.Background, light.ElevationAt(2).Blur)))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "样式解析 API", Note: "StyleOf/ResolveStyle/MetricsFor 是组件取配色与尺寸的唯一入口", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				im := ui.NewImplicit(300*time.Millisecond, theme.EaseOut, 0)
				im.To(1, 0)
				tween := ui.Tween{From: 40, To: 200}
				solid := ui.ResolveStyle(th, ui.VariantSolid, ui.StatusNormal)
				outline := ui.StyleOf(th, ui.VariantOutline, ui.SizeSmall, ui.StateHover)
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.FillRRect(geom.RectXYWH(10, 10, 340, 44), solid.Radius, render.Fill(solid.Background))
					cv.FillRRect(geom.RectXYWH(10, 66, 220, 44), outline.Radius,
						render.Stroke(outline.Border, outline.BorderWidth))
					cv.FillRect(geom.RectXYWH(10, 130, tween.At(0.5), 24), render.Fill(geom.RGBA(70, 130, 220)))
					cv.FillRect(geom.RectXYWH(10, 166, 340*im.Value(), 24), render.Fill(geom.RGBA(230, 120, 70)))
				}, 370, 210)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("Tween(0.5)=%.1f Implicit.Target=%.1f Solid=%v Outline=%v",
					tween.At(0.5), im.Target(), solid.Background, outline.Border)))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "theme.Light/Dark", Type: "func() *Data", Desc: "内置亮/暗主题（antd v5 对齐）"},
			{Name: "Data.Color/Radius/Space/Font/Motion/Elevation/Z", Type: "Token 组", Desc: "全部设计 token"},
			{Name: "ElevationAt", Type: "func(i int) ShadowSpec", Desc: "第 i 层阴影（越界取末层）"},
			{Name: "ui.StyleOf", Type: "func(th, Variant, Size, StateSet) Style", Desc: "组件样式解析入口"},
			{Name: "ui.MetricsFor", Type: "func(th, Size) Metrics", Desc: "控件高度/内边距/字号（24/32/40）"},
			{Name: "ui.ResolveStyle", Type: "func(th, Variant, Status) Style", Desc: "无状态样式解析"},
			{Name: "ui.NewImplicit/Tween", Type: "动画原语", Desc: "见 animation 窗"},
		},
	})
}
