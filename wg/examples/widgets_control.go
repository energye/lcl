package examples

import (
	"fmt"

	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func init() {
	Register(Example{ID: "button", Name: "Button 按钮", Group: GroupControl, Build: func(ctx *ui.BuildContext) *scene.Node {
		clicks := 0
		main := widgets.NewButton("点我")
		main.OnClick = func() {
			clicks++
			main.Text = fmt.Sprintf("已点击 %d 次", clicks)
			main.MarkDirty()
		}
		row := scene.Flex(scene.Horizontal, 10)
		row.Add(main.Build(ctx))
		disabled := widgets.NewButton("禁用")
		disabled.Disabled = true
		row.Add(disabled.Build(ctx))
		outline := widgets.NewButton("描边")
		outline.Variant = ui.VariantOutline
		row.Add(outline.Build(ctx))
		return Section(ctx, "Button", "点击走 GestureArena → Tap → OnClick；真窗里请用鼠标实际点一下", row)
	}})

	{
		cb := widgets.NewCheckbox("启用动画", true)
		Register(Example{ID: "checkbox", Name: "Checkbox 复选框", Group: GroupControl,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				col := scene.Flex(scene.Vertical, 8)
				col.Add(cb.Build(ctx))
				off := widgets.NewCheckbox("禁用项", false)
				off.Disabled = true
				col.Add(off.Build(ctx))
				return Section(ctx, "Checkbox", "点击切换 Checked，状态由组件实例跨 Build 持有", col)
			}})
	}

	{
		r1 := widgets.NewRadio("选项一", true)
		r2 := widgets.NewRadio("选项二", false)
		r1.OnChange = func(bool) { r2.Checked = !r1.Checked; r2.MarkDirty() }
		r2.OnChange = func(bool) { r1.Checked = !r2.Checked; r1.MarkDirty() }
		Register(Example{ID: "radio", Name: "Radio 单选", Group: GroupControl,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				col := scene.Flex(scene.Vertical, 8)
				col.Add(r1.Build(ctx))
				col.Add(r2.Build(ctx))
				return Section(ctx, "Radio", "互斥切换示例：由回调互改兄弟节点状态", col)
			}})
	}

	{
		sw := widgets.NewSwitch(true)
		Register(Example{ID: "switch", Name: "Switch 开关", Group: GroupControl,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(sw.Build(ctx))
				dis := widgets.NewSwitch(false)
				dis.Disabled = true
				col.Add(dis.Build(ctx))
				return Section(ctx, "Switch", "轨道+滑块用 scene 图元绘制，无子句柄", col)
			}})
	}

	{
		status := widgets.NewButton("等待输入…")
		inp := widgets.NewInput("试试输入中文，或用输入法组合")
		inp.Width = 420
		inp.OnChanged = func(s string) {
			status.Text = "当前内容：" + s
			status.MarkDirty()
		}
		clear := widgets.NewButton("清空")
		clear.OnClick = func() {
			inp.Editor().SetValue(input.EditingValue{})
			inp.Value = ""
			inp.MarkDirty()
		}
		Register(Example{ID: "input", Name: "Input 输入框 / IME", Group: GroupControl,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(inp.Build(ctx))
				row := scene.Flex(scene.Horizontal, 8)
				row.Add(status.Build(ctx))
				row.Add(clear.Build(ctx))
				col.Add(row)
				return Section(ctx, "Input / IME", "真窗里点输入框，切换中文输入法：组合串与选区替换走 C28 链路", col)
			}})
	}

	Register(Example{ID: "progress", Name: "Progress 进度条", Group: GroupControl, Build: func(ctx *ui.BuildContext) *scene.Node {
		col := scene.Flex(scene.Vertical, 10).SetWidth(420)
		for _, v := range []float32{0, 0.25, 0.6, 1} {
			col.Add((&widgets.Progress{Value: v, Width: 400, Height: 10}).Build(ctx))
		}
		return Section(ctx, "Progress", "0 / 25% / 60% / 100% 四档；绘制走 scene 图元", col)
	}})

	{
		slider := widgets.NewSlider(0, 100, 35)
		slider.Width = 420
		Register(Example{ID: "slider", Name: "Slider 滑块", Group: GroupControl,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				return Section(ctx, "Slider", "拖动走 GestureArena：越过 slop 的位置决定新值，避免首帧跳动", slider.Build(ctx))
			}})
	}

	{
		tabs := widgets.NewTabs(
			widgets.Tab{Key: "t1", Label: "第一页", Content: From(func(ctx *ui.BuildContext) *scene.Node {
				return (&widgets.Label{Text: "第一页内容"}).Build(ctx)
			})},
			widgets.Tab{Key: "t2", Label: "第二页", Content: From(func(ctx *ui.BuildContext) *scene.Node {
				return (&widgets.Label{Text: "第二页内容"}).Build(ctx)
			})},
			widgets.Tab{Key: "t3", Label: "第三页", Content: From(func(ctx *ui.BuildContext) *scene.Node {
				return (&widgets.Label{Text: "第三页内容"}).Build(ctx)
			})},
		)
		Register(Example{ID: "tabs", Name: "Tabs 标签页", Group: GroupControl,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				return Section(ctx, "Tabs", "点击标签切换内容；当前实现会重建内容子树", tabs.Build(ctx))
			}})
	}

	{
		tip := &widgets.Tooltip{Text: "悬停 400ms 后显示，移开自动消失", Child: widgets.NewButton("把鼠标停在这里")}
		Register(Example{ID: "tooltip", Name: "Tooltip 提示", Group: GroupControl,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				return Section(ctx, "Tooltip", "走 overlay.Host 浮层；Dismiss(0) 不吞后续点击", tip.Build(ctx))
			}})
	}
}
