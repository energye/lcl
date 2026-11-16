package main

import (
	"time"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// motionBase 给动画控制器提供稳定的墙钟起点（Step 用绝对时间，不逐帧累加）。
var motionBase = time.Now()

var (
	ctrl   = ui.NewController(600*time.Millisecond, func(t float32) float32 { return t * t * (3 - 2 * t) })
	status = widgets.NewButton("点击播放动画")
)

func init() {
	status.OnClick = func() {
		status.Text = "播放中"
		status.MarkDirty()
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "animation", Title: "动画与隐式插值", Group: doc.GroupVisual,
		What: []string{
			"Controller 只算插值、不自带定时器：Step 与 RequestFrame 由帧循环驱动",
			"Implicit 是隐式动画（目标值渐近）；Tween 是纯区间插值原语",
			"ReduceMotion 偏好由 ctx.Motion(d) 统一归零时长",
		},
		Demos: []doc.Demo{
			{Title: "Controller 进度条", Note: "点「点击播放动画」再点进度条下的按钮从头播放", Build: func(ctx *ui.BuildContext) ui.Widget {
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
				bar := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
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
				}, 420, 18)
				col := scene.Flex(scene.Vertical, 10)
				col.Add(bar)
				row := scene.Flex(scene.Horizontal, 8)
				row.Add(start.Build(ctx))
				row.Add(status.Build(ctx))
				col.Add(row)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "Implicit / Tween / 样式解析", Note: "隐式动画与 Tween 插值（样式解析详见 theme 窗）", Build: func(ctx *ui.BuildContext) ui.Widget {
				im := ui.NewImplicit(300*time.Millisecond, theme.EaseOut, 0)
				im.To(1, 0)
				tween := ui.Tween{From: 40, To: 200}
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.FillRect(geom.RectXYWH(10, 130, tween.At(0.5), 24), render.Fill(geom.RGBA(70, 130, 220)))
					cv.FillRect(geom.RectXYWH(10, 166, 340*im.Value(), 24), render.Fill(geom.RGBA(230, 120, 70)))
				}, 370, 210)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, "Tween(0.5)=120 隐式动画：目标值渐近，帧循环推进"))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "ui.NewController", Type: "func(d, ease) *Controller", Desc: "动画控制器（插值计算）"},
			{Name: "Controller.Step/Forward", Type: "func(time.Time/Duration)", Desc: "推进（绝对时间）"},
			{Name: "Controller.Animating/Value", Type: "查询", Desc: "是否在动 / 当前值"},
			{Name: "ui.NewImplicit", Type: "func(d, ease, init) *Implicit", Desc: "隐式动画：To 设目标，值渐近"},
			{Name: "ui.Tween", Type: "{From, To}.At(t)", Desc: "区间插值原语"},
			{Name: "ctx.Scheduler.RequestFrame", Type: "帧请求", Desc: "驱动下一帧（单 outstanding 帧）"},
		},
	})
}
