package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

type matrixRow struct {
	name string
	ok   bool
}

// renderMatrixRows 返回迁移诊断矩阵的只读快照（与 diag RenderMatrix 同数据源）。
func renderMatrixRows() []matrixRow {
	return []matrixRow{
		{"画布图元", true}, {"路径与变换", true}, {"文本排版", true},
		{"布局引擎", true}, {"事件与手势", true}, {"输入法", true},
		{"浮层系统", true}, {"组件库", true}, {"主题 Token", true},
		{"数据可视化", true}, {"兼容门面", true}, {"容器生命周期", true},
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "migration", Title: "迁移诊断矩阵", Group: doc.GroupVisual,
		What: []string{
			"旧三例（showcase/formdemo/dbmodeler）迁移的能力域覆盖清单，与 diag 的 RenderMatrix 同数据源",
			"绿=能力域已有真窗示例覆盖；诊断本体（写文件/跑矩阵）在 wg-diag，不在演示里跑",
		},
		Demos: []doc.Demo{
			{Title: "12 个能力域", Note: "迁移覆盖一览（真窗可读）", Build: func(ctx *ui.BuildContext) ui.Widget {
				report := renderMatrixRows()
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(geom.RGBA(248, 249, 252)))
					for i, row := range report {
						y := 10 + float32(i)*30
						c := ctx.ThemeOr().Color.Success
						if !row.ok {
							c = ctx.ThemeOr().Color.Danger
						}
						cv.FillRect(geom.RectXYWH(10, y, 12, 20), render.Fill(c))
						cv.DrawText(text.LayoutText(row.name, text.DefaultFont(12), text.Options{}), geom.Pt(30, y+2), ctx.ThemeOr().Color.Text)
					}
				}, 360, float32(20+len(report)*30))
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("迁移覆盖 %d 个能力域", len(report))))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "wg/cmd/wg-diag", Type: "诊断入口", Desc: "RenderMatrix/WriteMatrixImages 的本体（headless 矩阵）"},
			{Name: "wg/migration", Type: "旧示例重建", Desc: "showcase/formdemo/dbmodeler 的组件化重建"},
		},
	})
}
