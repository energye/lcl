package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var frameStatus = widgets.NewButton("帧信息：—")

func main() {
	launcher.Main(doc.Meta{
		ID: "frame-scheduler", Title: "帧调度", Group: doc.GroupVisual,
		What: []string{
			"帧调度器（C14）：单 outstanding 帧——重复 RequestFrame 只保留一个待发帧",
			"静止零帧：没有脏区就没有帧；60Hz 对齐 / RequestFrameIn 定时帧",
			"耗时分布用中位数与 P90 报告（见 --bench 矩阵）",
		},
		Demos: []doc.Demo{
			{Title: "请求下一帧", Note: "点击后显示帧号/等待/已武装状态", Build: func(ctx *ui.BuildContext) ui.Widget {
				req := widgets.NewButton("请求下一帧")
				req.OnClick = func() {
					if ctx.Scheduler != nil {
						ctx.Scheduler.RequestFrame()
						frameStatus.Text = fmt.Sprintf("帧=%d 等待=%v 已武装=%v",
							ctx.Scheduler.Frame(), ctx.Scheduler.Awaiting(), ctx.Scheduler.Armed())
					} else {
						frameStatus.Text = "headless：无调度器"
					}
					frameStatus.MarkDirty()
				}
				row := scene.Flex(scene.Horizontal, 8)
				row.Add(req.Build(ctx))
				row.Add(frameStatus.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "耗时分布", Note: "帧耗时样本的柱状分布，超 50ms 标红（中位/P90 算法与 --bench 一致）", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				deltas := []float64{8, 12, 16, 20, 24, 100}
				sorted := append([]float64(nil), deltas...)
				for i := 1; i < len(sorted); i++ {
					for j := i; j > 0 && sorted[j] < sorted[j-1]; j-- {
						sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
					}
				}
				median := sorted[len(sorted)/2]
				p90 := sorted[int(float64(len(sorted)-1)*0.9+0.5)]
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
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
				}, 390, 165)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("中位=%.0fms P90=%.0fms", median, p90)))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "Scheduler.RequestFrame", Type: "func()", Desc: "请求下一帧（单 outstanding）"},
			{Name: "RequestFrameIn", Type: "func(d time.Duration)", Desc: "定时帧（动画节流用）"},
			{Name: "Frame/Awaiting/Armed", Type: "查询", Desc: "帧号 / 待发 / 已武装"},
			{Name: "静止零帧（C14）", Type: "契约", Desc: "没有脏区就不出帧不提交"},
		},
	})
}
