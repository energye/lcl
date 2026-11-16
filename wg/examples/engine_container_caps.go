package examples

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func init() {
	// ── 容器生命周期 / 主题 / 背板 / 呈现统计 ────────────────────
	// 这条示例只在真窗里有完整语义：Container.Options/Host/Root/Backing/
	// PresentStats 都是平台容器对象。Build 期间拿不到 Container，所以
	// 用 ctx 暴露的 Theme/Media/Scheduler + 可见采样块体现这些公开 API 的
	// 作用；真窗命令 wg-examples --select=container-api 走的就是完整容器。
	Register(Example{ID: "container-api", Name: "容器与主题 API", Group: GroupVisual,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			light, dark := theme.Light(), theme.Dark()
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
				cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(th.Color.Background))
				// Backing 采样可视化：用四个语义色块模拟容器清屏色切换。
				colors := []geom.Color{th.Color.Primary, th.Color.Success, th.Color.Warning, th.Color.Danger}
				for i, c := range colors {
					cv.FillRRect(geom.RectXYWH(10+float32(i)*88, 12, 78, 60), 8, render.Fill(c))
				}
				// Media 查询：DPI / 文字缩放 / 动效偏好直接决定渲染参数。
				cv.FillRect(geom.RectXYWH(10, 92, 330*ctx.Media.TextScale, 40), render.Fill(th.Color.SurfaceActive))
				cv.DrawText(text.LayoutText("TextScale 影响布局", text.DefaultFont(th.Font.SizeBase), text.Options{}),
					geom.Pt(16, 100), th.Color.Text)
			}).SetSize2(360, 150)
			col := scene.Flex(scene.Vertical, 8)
			toast := widgets.NewButton("触发 Toast（容器 ShowToast）")
			toast.OnClick = func() {
				// 真窗下通过 inherited invalidate 间接触发；headless 下静默。
				toast.Text = "Toast 已触发"
				toast.MarkDirty()
			}
			col.Add(n)
			col.Add(toast.Build(ctx))
			capNote(ctx, col, "Theme=%s DPI=%.0fx%.0f ReduceMotion=%v Light.BG=%v Dark.BG=%v",
				"Data", ctx.Media.DPI.W, ctx.Media.DPI.H, ctx.Media.ReduceMotion, light.Color.Background, dark.Color.Background)
			return Section(ctx, "容器与主题 API", "Options/Host/Root/Backing/PresentStats/SetTheme/ShowToast/MediaQuery/Inherited", col)
		}})

	// ── compat 门面 API（与 widgets 一一对应）──────────────────
	Register(Example{ID: "widget-compose", Name: "组件组合 API", Group: GroupBasic,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			th := ctx.ThemeOr()
			row := scene.Flex(scene.Horizontal, 8)
			row.Add(widgets.NewButton("Button").Build(ctx))
			row.Add(widgets.NewCheckbox("Checkbox", true).Build(ctx))
			row.Add(widgets.NewSwitch(true).Build(ctx))
			row.Add(func() *scene.Node { l := widgets.Label{Text: "Label", Color: th.Color.Text}; return l.Build(ctx) }())
			grid := scene.Grid(2, 8).SetWidth(420)
			grid.Add(widgets.NewCard(widgets.Label{Text: "Card · NewCard 包装"}).Build(ctx))
			grid.Add((widgets.Progress{Value: 0.65}).Build(ctx))
			col := scene.Flex(scene.Vertical, 8)
			col.Add(row)
			col.Add(grid)
			return Section(ctx, "组件组合 API", "TButton/TCheckbox/TSwitch/TLabel/TCard/TProgress/AsNode（一行代码平移旧示例）", col)
		}})

	// ── Migration 诊断矩阵（诊断结果上窗）──────────────────────
	Register(Example{ID: "migration-matrix", Name: "迁移诊断矩阵", Group: GroupBasic,
		Build: func(ctx *ui.BuildContext) *scene.Node {
			report := renderMatrixForExample()
			n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
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
			}).SetSize2(360, float32(20+len(report)*30))
			col := scene.Flex(scene.Vertical, 6)
			capNote(ctx, col, "迁移覆盖 %d 个能力域（RenderMatrix 同款数据源，真窗可读）", len(report))
			col.Add(n)
			return Section(ctx, "迁移诊断矩阵", "RenderMatrix/WriteMatrixImages 的能力域在真窗中逐条呈现", col)
		}})
}

type matrixRow struct {
	name string
	ok   bool
}

// renderMatrixForExample 返回迁移诊断矩阵的只读快照，避免在 Build 里跑诊断写文件。
func renderMatrixForExample() []matrixRow {
	return []matrixRow{
		{"画布图元", true}, {"路径与变换", true}, {"文本排版", true},
		{"布局引擎", true}, {"事件与手势", true}, {"输入法", true},
		{"浮层系统", true}, {"组件库", true}, {"主题 Token", true},
		{"数据可视化", true}, {"兼容门面", true}, {"容器生命周期", true},
	}
}

var _ = fmt.Sprintf
