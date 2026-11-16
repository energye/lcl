package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "text-layout", Title: "文本排版与测量", Group: doc.GroupVisual,
		What: []string{
			"text 包只排版不画字：LayoutText 产出 Layout，字形由 GDI 通道回放保 ClearType（C6/C24）",
			"排版选项：MaxWidth 换行 / MaxLines 限行 / Ellipsis 省略；偏移统一按 rune 计算",
			"Measure 测量 + Font.WithWeight/WithSize/WithFamily 派生字体",
		},
		Demos: []doc.Demo{
			{Title: "排版三选项", Note: "混排 / 超长省略 / 多行", Build: func(ctx *ui.BuildContext) ui.Widget {
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
					col.Add(scene.Text(scene.TextSpec{
						Src: s.text,
						Layout: text.LayoutText(s.text, text.DefaultFont(th.Font.SizeBase), text.Options{
							MaxWidth: s.maxWidth, MaxLines: s.maxLines, Ellipsis: s.ellipsis,
						}),
						Color: th.Color.Text,
					}))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "测量与字重", Note: "Measure 三种字体的宽度差异；Layout.LineX 行定位", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				f := text.DefaultFont(th.Font.SizeBase)
				bold := f.WithWeight(text.Bold)
				big := f.WithSize(th.Font.SizeTitle).WithFamily("Segoe UI")
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.DrawText(text.LayoutText("Regular", f, text.Options{}), geom.Pt(10, 8), th.Color.Text)
					cv.DrawText(text.LayoutText("Bold 700", bold, text.Options{}), geom.Pt(10, 40), th.Color.Text)
					cv.DrawText(text.LayoutText("Title 24", big, text.Options{}), geom.Pt(10, 74), th.Color.Primary)
					lay := text.LayoutText("第一行\n第二行\n第三行", f, text.Options{MaxWidth: 300})
					cv.DrawText(lay, geom.Pt(10, 116), th.Color.Text)
					if len(lay.Lines) > 0 {
						cv.DrawLine(10, 116+lay.LineX(0), 300, 116+lay.LineX(0), render.Stroke(th.Color.Border, 1))
					}
				}, 350, 200)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("Measure: Regular=%v Bold=%v Title=%v",
					text.Measure("渲染引擎", f), text.Measure("渲染引擎", bold), text.Measure("渲染引擎", big))))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "text.LayoutText", Type: "func(s string, font Font, opts Options) *Layout", Desc: "排版；Options{MaxWidth,MaxLines,Ellipsis,Align,LineHeight}"},
			{Name: "text.Measure", Type: "func(s string, font Font) Size", Desc: "单行测量"},
			{Name: "Font.WithWeight/WithSize/WithFamily", Type: "派生方法", Desc: "不改原字体的派生"},
			{Name: "Layout.Lines/LineX", Type: "字段 / 方法", Desc: "行数据与行 Y 偏移"},
			{Name: "text.DefaultFont", Type: "func(size float32) Font", Desc: "主题字号的默认字体"},
		},
	})
}

var _ = theme.Default
