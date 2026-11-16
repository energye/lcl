package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	status   = widgets.NewButton("等待手势…")
	detector *ui.GestureDetector // 跨 Build 存活，避免闭包被 GC
)

func init() {
	detector = &ui.GestureDetector{
		OnTap: func(p geom.Point) {
			status.Text = fmt.Sprintf("Tap @ (%.0f, %.0f)", p.X, p.Y)
			status.MarkDirty()
		},
		OnLongPress: func(geom.Point) { status.Text = "LongPress"; status.MarkDirty() },
		OnPanStart:  func(geom.Point) { status.Text = "PanStart"; status.MarkDirty() },
		OnPan: func(_, d geom.Point) {
			status.Text = fmt.Sprintf("Pan Δ(%.0f, %.0f)", d.X, d.Y)
			status.MarkDirty()
		},
		OnPanEnd: func(_, v geom.Point) {
			status.Text = fmt.Sprintf("PanEnd v=(%.0f, %.0f)", v.X, v.Y)
			status.MarkDirty()
		},
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "gesture", Title: "手势竞技场", Group: doc.GroupInput,
		What: []string{
			"GestureArena 单胜者裁决：Tap / LongPress / Pan 互斥，Accept 是唯一开火点",
			"竞技场每次按下新建；越界 slop 之前不判定，避免首帧抖动",
		},
		Demos: []doc.Demo{
			{Title: "三种手势", Note: "真窗里单击 / 长按 / 拖动同一块区域", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				area := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceActive)}).SetSize2(420, 160)
				area.Add((&widgets.Label{Text: "单击 / 长按 / 拖动"}).Build(ctx))
				if ctx.Input != nil {
					ctx.Input.Set(area, detector.Handler())
				}
				col := scene.Flex(scene.Vertical, 10)
				col.Add(area)
				col.Add(status.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "ui.GestureDetector", Type: "结构体", Desc: "OnTap/OnLongPress/OnPanStart/OnPan/OnPanEnd"},
			{Name: "Handler", Type: "func() input.Handler", Desc: "包装成事件处理器"},
			{Name: "竞技场语义", Type: "契约", Desc: "单胜者互斥；Tap 与 Pan 以移动距离裁决"},
			{Name: "实例持有", Type: "约定", Desc: "检测器跨 Build 存活（包级变量）"},
		},
	})
}
