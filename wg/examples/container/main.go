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
	"github.com/energye/lcl/wg/widgets"
)

var ctxRef *ui.BuildContext

func main() {
	toast := widgets.NewButton("触发 Toast（容器 ShowToast）")
	toast.OnClick = func() {
		toast.Text = "Toast 已触发（真窗）"
		toast.MarkDirty()
	}
	launcher.Main(doc.Meta{
		ID: "container", Title: "容器与生命周期", Group: doc.GroupVisual,
		What: []string{
			"Container 是组合根：MountWidget 挂载、帧泵 FrameOnce、Presenter 提交、Binder 接管消息",
			"退出协议：先 ReleasePresent 再 Application.Terminate（Terminate 后释放位图会 0xc0000005）",
			"BuildContext 暴露 Theme/Media/Scheduler/Inherited；本窗即跑在一个真 Container 里",
		},
		Demos: []doc.Demo{
			{Title: "Media 查询与清屏色", Note: "TextScale 影响布局宽度；四个色块对应容器清屏色切换演示", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				light, dark := theme.Light(), theme.Dark()
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(th.Color.Background))
					colors := []geom.Color{th.Color.Primary, th.Color.Success, th.Color.Warning, th.Color.Danger}
					for i, c := range colors {
						cv.FillRRect(geom.RectXYWH(10+float32(i)*88, 12, 78, 60), 8, render.Fill(c))
					}
					cv.FillRect(geom.RectXYWH(10, 92, 330*ctx.Media.TextScale, 40), render.Fill(th.Color.SurfaceActive))
					cv.DrawText(text.LayoutText("TextScale 影响布局", text.DefaultFont(th.Font.SizeBase), text.Options{}),
						geom.Pt(16, 100), th.Color.Text)
				}, 360, 150)
				col := scene.Flex(scene.Vertical, 8)
				col.Add(n)
				col.Add(toast.Build(ctx))
				col.Add(doc.Note(ctx, fmt.Sprintf("DPI=%.0fx%.0f ReduceMotion=%v Light.BG=%v Dark.BG=%v",
					ctx.Media.DPI.W, ctx.Media.DPI.H, ctx.Media.ReduceMotion, light.Color.Background, dark.Color.Background)))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "独立 main 骨架", Note: "每个功能窗都是这个骨架（launcher.Main 内部完成）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return &widgets.Label{
					Text: "lcl.Init → NewForm → wg.NewContainer(form, opts) → MountWidget(doc.Page) → container.Run()\n退出：ReleasePresent() → Application.Terminate()",
					Size: 13,
				}
			}},
		},
		API: []doc.APIRow{
			{Name: "wg.NewContainer", Type: "func(form, opts) *Container", Desc: "组合根；form=nil 为 headless"},
			{Name: "MountWidget/Mount", Type: "挂载", Desc: "挂 ui.Widget / scene.Node"},
			{Name: "Run", Type: "func() error", Desc: "帧泵+消息循环（内部已调 Application.Run）"},
			{Name: "FrameOnce", Type: "func() (frame.Stats, error)", Desc: "单帧（headless 驱动）"},
			{Name: "ReleasePresent", Type: "func()", Desc: "Terminate 前必须调用（呈现位图释放）"},
			{Name: "PresentNow/Backing/Frames", Type: "呈现查询", Desc: "--shot 管线的基础"},
			{Name: "SetTheme", Type: "func(*theme.Data)", Desc: "换肤（整树重排+重建）"},
		},
	})
}
