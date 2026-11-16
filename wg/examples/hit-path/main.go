package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var status = widgets.NewButton("点击矩形试试")

func main() {
	launcher.Main(doc.Meta{
		ID: "hit-path", Title: "命中路径", Group: doc.GroupInput,
		What: []string{
			"指针事件按命中路径分发：捕获（按下后锁定）→ 目标 → 冒泡",
			"Enter/Leave 由 Dispatcher 的路径差分合成，Hover 事件也会重算",
			"形状与绘制共用同一份几何（HitSpec），命中不需要额外描述",
		},
		Demos: []doc.Demo{
			{Title: "矩形命中区域", Note: "真窗里移动/点击矩形，观察 Enter/Leave/Pointer 回调", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				area := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceActive)}).SetSize2(420, 140)
				area.Add((&widgets.Label{Text: "命中区域（捕获 → 目标 → 冒泡）"}).Build(ctx))
				if ctx.Input != nil {
					ctx.Input.Set(area, input.Handler{
						Pointer: func(ev *input.PointerEvent) bool {
							status.Text = fmt.Sprintf("%s @ (%.0f, %.0f) button=%v",
								ev.Kind.String(), ev.Pos.X, ev.Pos.Y, ev.Buttons)
							status.MarkDirty()
							return true
						},
						Enter: func(*input.PointerEvent) { status.Text = "鼠标进入"; status.MarkDirty() },
						Leave: func(*input.PointerEvent) { status.Text = "鼠标离开"; status.MarkDirty() },
					})
				}
				col := scene.Flex(scene.Vertical, 10)
				col.Add(area)
				col.Add(status.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "Dispatcher.Set", Type: "func(node, Handler)", Desc: "给节点挂事件处理器"},
			{Name: "input.Handler", Type: "结构体", Desc: "Capture/Pointer/Bubble/Enter/Leave/Wheel/Key/IME/Focus"},
			{Name: "PointerEvent", Type: "{Kind,Pos,Buttons…}", Desc: "归一化指针事件"},
			{Name: "捕获语义", Type: "契约", Desc: "按下后锁定目标，不因移动重新命中"},
			{Name: "Cancel 自愈", Type: "契约 C12", Desc: "指针异常抬起时自动取消手势"},
		},
	})
}
