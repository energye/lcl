package examples

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// motionBase 给动画控制器提供稳定的墙钟起点（Step 用绝对时间，不用逐帧累加）。
var motionBase = time.Now()

func init() {
	{
		ease := func(t float32) float32 { return t * t * (3 - 2*t) }
		ctrl := ui.NewController(600*time.Millisecond, ease)
		status := widgets.NewButton("点击播放动画")
		Register(Example{ID: "animation", Name: "动画控制器", Group: GroupVisual,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				sched := ctx.Scheduler
				start := widgets.NewButton("从头播放")
				start.OnClick = func() {
					ctrl.SetValue(0)
					ctrl.Forward(time.Since(motionBase))
					status.Text = "播放中"
					status.MarkDirty()
					if sched != nil {
						sched.RequestFrame()
					}
				}
				bar := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
					// 录制阶段推进插值并请求下一帧：真窗下动画在这里持续走，
					// headless 单帧不会自行推进（由手动帧驱动）。
					if ctrl.Animating() {
						ctrl.Step(time.Since(motionBase))
						if sched != nil {
							sched.RequestFrame()
						}
					}
					cv.FillRRect(geom.RectXYWH(0, 0, sz.W, sz.H), sz.H/2,
						render.Paint{Color: geom.RGBA(210, 215, 225)})
					cv.FillRRect(geom.RectXYWH(0, 0, sz.W*ctrl.Value(), sz.H), sz.H/2,
						render.Paint{Color: geom.RGBA(70, 130, 220)})
				}).SetSize2(420, 18)
				col := scene.Flex(scene.Vertical, 10)
				col.Add(bar)
				row := scene.Flex(scene.Horizontal, 8)
				row.Add(start.Build(ctx))
				row.Add(status.Build(ctx))
				col.Add(row)
				return Section(ctx, "动画控制器", "Controller 只算插值、不自带定时器；Step 与 RequestFrame 由帧循环驱动", col)
			}})
	}

	{
		status := widgets.NewButton("帧信息：—")
		Register(Example{ID: "scheduler", Name: "帧调度", Group: GroupVisual,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				req := widgets.NewButton("请求下一帧")
				req.OnClick = func() {
					if ctx.Scheduler != nil {
						ctx.Scheduler.RequestFrame()
						status.Text = fmt.Sprintf("帧=%d 等待=%v 已武装=%v",
							ctx.Scheduler.Frame(), ctx.Scheduler.Awaiting(), ctx.Scheduler.Armed())
					} else {
						status.Text = "headless：无调度器"
					}
					status.MarkDirty()
				}
				row := scene.Flex(scene.Horizontal, 8)
				row.Add(req.Build(ctx))
				row.Add(status.Build(ctx))
				return Section(ctx, "帧调度", "单 outstanding 帧：重复 RequestFrame 只保留一个待发帧", row)
			}})
	}

	Register(Example{ID: "damage", Name: "damage 合并", Group: GroupVisual, Build: func(ctx *ui.BuildContext) *scene.Node {
		n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
			cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Paint{Color: geom.RGBA(38, 42, 54)})
			for i, r := range []geom.Rect{
				geom.RectXYWH(20, 20, 140, 90),
				geom.RectXYWH(130, 60, 140, 90),
				geom.RectXYWH(300, 20, 100, 130),
			} {
				cv.FillRect(r, render.Paint{
					Color: geom.RGBAa(255, 180, 60, uint8(90+i*40)),
					Style: render.StyleStroke, StrokeWidth: 2,
				})
			}
		}).SetSize2(440, 180)
		return Section(ctx, "damage 合并", "相邻脏矩形合并成一个包围盒，Present 只拷贝该区域", n)
	}})
}
