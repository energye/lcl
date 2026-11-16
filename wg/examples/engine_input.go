package examples

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func init() {
	{
		status := widgets.NewButton("点击矩形试试")
		Register(Example{ID: "hit-path", Name: "命中路径", Group: GroupInput,
			Build: func(ctx *ui.BuildContext) *scene.Node {
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
				return Section(ctx, "命中路径", "真窗里移动/点击矩形：Enter/Leave 走差分，Pointer 走捕获与冒泡", col)
			}})
	}

	{
		status := widgets.NewButton("等待手势…")
		var detector *ui.GestureDetector // 跨 Build 存活，避免闭包被 GC
		Register(Example{ID: "gesture", Name: "手势竞技场", Group: GroupInput,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				th := ctx.ThemeOr()
				area := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceActive)}).SetSize2(420, 160)
				area.Add((&widgets.Label{Text: "单击 / 长按 / 拖动"}).Build(ctx))
				if detector == nil {
					detector = &ui.GestureDetector{
						OnTap: func(p geom.Point) {
							status.Text = fmt.Sprintf("Tap @ (%.0f, %.0f)", p.X, p.Y)
							status.MarkDirty()
						},
						OnLongPress: func(p geom.Point) { status.Text = "LongPress"; status.MarkDirty() },
						OnPanStart:  func(p geom.Point) { status.Text = "PanStart"; status.MarkDirty() },
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
				if ctx.Input != nil {
					ctx.Input.Set(area, detector.Handler())
				}
				col := scene.Flex(scene.Vertical, 10)
				col.Add(area)
				col.Add(status.Build(ctx))
				return Section(ctx, "手势竞技场", "单胜者：Tap / LongPress / Pan 互斥，Accept 是唯一开火点", col)
			}})
	}

	{
		inputs := []*widgets.Input{
			widgets.NewInput("第一个输入框"),
			widgets.NewInput("第二个输入框"),
			widgets.NewInput("第三个输入框"),
		}
		for _, in := range inputs {
			in.Width = 360
		}
		Register(Example{ID: "focus", Name: "焦点域与 Tab", Group: GroupInput,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				col := scene.Flex(scene.Vertical, 10)
				for _, in := range inputs {
					col.Add(in.Build(ctx))
				}
				return Section(ctx, "焦点域与 Tab", "真窗里 Tab / Shift+Tab 顺序切换；焦点改变后 Key/IME 投给新节点", col)
			}})
	}

	{
		status := widgets.NewButton("按任意键（先点一下此区域）")
		Register(Example{ID: "keyboard", Name: "键盘事件", Group: GroupInput,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				th := ctx.ThemeOr()
				area := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceActive)}).SetSize2(420, 90)
				area.Add((&widgets.Label{Text: "键盘目标区域"}).Build(ctx))
				if ctx.Input != nil {
					ctx.Input.Set(area, input.Handler{
						Key: func(ev *input.KeyEvent) bool {
							status.Text = fmt.Sprintf("Kind=%d code=%d text=%q mods=%v", ev.Kind, ev.Code, ev.Text, ev.Mods)
							status.MarkDirty()
							return true
						},
					})
				}
				col := scene.Flex(scene.Vertical, 10)
				col.Add(area)
				col.Add(status.Build(ctx))
				return Section(ctx, "键盘事件", "VK → Code 归一化；焦点节点未处理时沿父链冒泡", col)
			}})
	}

	{
		inp := widgets.NewInput("点这里，然后切换中文输入法")
		inp.Width = 420
		status := widgets.NewButton("等待组合串…")
		Register(Example{ID: "ime", Name: "IME 组合与选区", Group: GroupInput,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				compose := widgets.NewButton("模拟：拼音 ni hao → 你好")
				compose.OnClick = func() {
					ed := inp.Editor()
					ed.SetValue(input.EditingValue{Text: "hello"})
					ed.SetSelection(input.TextRange{Start: 0, End: 5})
					ed.SetComposingRegion(0, 5, "ni hao")
					ed.CommitComposingWith("你好")
					status.Text = "编辑结果：" + ed.Value().Text
					status.MarkDirty()
					inp.MarkDirty()
				}
				col := scene.Flex(scene.Vertical, 10)
				col.Add(inp.Build(ctx))
				row := scene.Flex(scene.Horizontal, 8)
				row.Add(compose.Build(ctx))
				row.Add(status.Build(ctx))
				col.Add(row)
				return Section(ctx, "IME 组合与选区", "C28：组合串提交时替换选区，而不是追加到末尾", col)
			}})
	}

	{
		inp := widgets.NewInput("试试 Backspace / Delete / Ctrl+A / Ctrl+←→")
		inp.Width = 460
		Register(Example{ID: "text-keys", Name: "文本编辑按键", Group: GroupInput,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				return Section(ctx, "文本编辑按键", "TextKeyHandler 统一处理编辑键；Enter/Tab 是语义动作", inp.Build(ctx))
			}})
	}

	{
		tip := &widgets.Tooltip{Text: "浮层按 Band 排序：Popup < Toast < Modal", Child: widgets.NewButton("悬停查看浮层层级")}
		sel := widgets.NewSelect([]string{"浮层 A", "浮层 B", "浮层 C"})
		sel.Width = 200
		Register(Example{ID: "overlay", Name: "浮层与关闭策略", Group: GroupInput,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				row := scene.Flex(scene.Horizontal, 16)
				row.Add(tip.Build(ctx))
				row.Add(sel.Build(ctx))
				return Section(ctx, "浮层与关闭策略", "OutsideClick / Esc / AnchorGone 由 overlay.Host 统一裁决", row)
			}})
	}
}
