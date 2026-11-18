package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	inp    = widgets.NewInput("试试输入中文，或用输入法组合")
	status = widgets.NewTextView("当前内容：")
	edit   = widgets.NewInput("试试 Backspace / Delete / Ctrl+A / Ctrl+←→")
	clear  = widgets.NewButton("清空")

	// 各 demo 的交互控件实例在包级持有：demo 闭包在每次页面重建时都会
	// 重新执行（产出全新节点），实例一换，焦点/文本/勾选态就跟着丢——
	// "点击聚焦→重建→焦点丢失"就是实例不持有导致的。闭包内只做
	// instance.Build(ctx)（模式与 scroll 示例一致）。
	variants = [3]*widgets.Input{
		widgets.NewInput("outlined"),
		widgets.NewInput("filled"),
		widgets.NewInput("borderless"),
	}
	statuses = [4]*widgets.Input{
		widgets.NewInput("error 状态"),
		widgets.NewInput("warning 状态"),
		widgets.NewInput("small"),
		widgets.NewInput("large"),
	}
	prefixed = widgets.NewInput("站内搜索")
	suffixed = widgets.NewInput("金额")
	addoned  = widgets.NewInput("example.com 前缀")
	counted  = widgets.NewInput("看右下角字数")
)

func init() {
	inp.OnChanged = func(s string) {
		status.SetText("当前内容：" + s)
	}
	clear.OnClick = func() {
		inp.Editor().SetValue(input.EditingValue{})
		inp.Value = ""
		inp.MarkDirty()
	}
}

func init() {
	inp.Width = 420
	// 固定宽 + 原地更新：每键回显只重绘这一行，不触发整页重建。
	status.Width = 520
	inp.OnChanged = func(s string) {
		status.SetText("当前内容：" + s)
	}
	edit.Width = 460

	variants[0].Width = 260
	variants[1].Width, variants[1].Variant = 260, widgets.InputFilled
	variants[2].Width, variants[2].Variant = 260, widgets.InputBorderless
	statuses[0].Width, statuses[0].Status = 260, widgets.InputStatusError
	statuses[1].Width, statuses[1].Status = 260, widgets.InputStatusWarning
	statuses[2].Width, statuses[2].Size = 260, ui.SizeSmall
	statuses[3].Width, statuses[3].Size = 260, ui.SizeLarge
	prefixed.Width, prefixed.AllowClear = 300, true
	prefixed.Prefix = &widgets.Icon{Name: "search", Size: 14}
	suffixed.Width = 300
	suffixed.Suffix = &widgets.Icon{Name: "info-circle", Size: 14}
	addoned.Width = 300
	addoned.AddonAfter = doc.Node(func(ctx *ui.BuildContext) *scene.Node {
		return (&widgets.Label{Text: ".com", Size: 13, Color: ctx.ThemeOr().Color.TextMuted}).Build(ctx)
	})
	counted.Width, counted.Count = 300, true
}

func main() {
	launcher.Main(doc.Meta{
		ID: "input", Title: "Input 输入框", Group: doc.GroupControl,
		What: []string{
			"单行/多行文本输入：TextEditor 持有文本/选区/撤销栈，跨重建不丢",
			"对齐 antd：variant（outlined/filled/borderless）、status（error/warning）、prefix/suffix、addonBefore/addonAfter、count、allowClear、三档尺寸",
			"IME 走 C28 链路：组合串提交时替换选区；编辑键由 TextKeyHandler 统一处理",
			"未实现：status 反馈图标、autosize；多行内核已支持（换行/组合/撤销），按行定位与内部滚动待补",
		},
		Demos: []doc.Demo{
			{Title: "受控输入与 IME", Note: "真窗里点输入框并切换中文输入法；OnChanged 实时回显", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(inp.Build(ctx))
				row := scene.Flex(scene.Horizontal, 8)
				row.Add(status.Build(ctx))
				row.Add(clear.Build(ctx))
				col.Add(row)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "三种 variant", Note: "outlined（默认）/ filled（灰底）/ borderless（无框）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Column{Gap: 8, Children: []ui.Widget{variants[0], variants[1], variants[2]}}
			}},
			{Title: "status 与尺寸", Note: "error 红边 / warning 黄边；small/middle/large", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Column{Gap: 8, Children: []ui.Widget{statuses[0], statuses[1], statuses[2], statuses[3]}}
			}},
			{Title: "前后缀与清除", Note: "prefix 图标 + suffix 图标；输入后 hover 显示 allowClear 清除", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Column{Gap: 8, Children: []ui.Widget{prefixed, suffixed}}
			}},
			{Title: "addon 与 count", Note: "addonBefore/addonAfter 灰块；Count 右下角字数", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Column{Gap: 12, Children: []ui.Widget{addoned, counted}}
			}},
			{Title: "文本编辑按键", Note: "编辑键统一处理；Enter/Tab 是语义动作不进入文本", Build: func(ctx *ui.BuildContext) ui.Widget {
				return edit
			}},
		},
		API: []doc.APIRow{
			{Name: "NewInput", Type: "func(placeholder string) *Input", Desc: "构造"},
			{Name: "Variant", Type: "widgets.InputVariant", Default: "InputOutlined", Desc: "outlined/filled/borderless（antd variant）"},
			{Name: "Status", Type: "widgets.InputStatus", Default: "StatusOK", Desc: "error 红边 / warning 黄边（antd status）"},
			{Name: "Size", Type: "ui.Size", Default: "SizeMedium", Desc: "控件高 24/32/40"},
			{Name: "Prefix/Suffix", Type: "ui.Widget", Desc: "框内左右装饰"},
			{Name: "AddonBefore/AddonAfter", Type: "ui.Widget", Desc: "框外灰块"},
			{Name: "AllowClear", Type: "bool", Default: "false", Desc: "有值时显示清除图标"},
			{Name: "Count", Type: "bool", Default: "false", Desc: "右下角字数（rune 计数）"},
			{Name: "Placeholder/Value", Type: "string", Desc: "占位 / 当前文本"},
			{Name: "Multiline/MaxLines", Type: "bool / int", Desc: "多行模式与最大行数"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态（antd disabled 专色）"},
			{Name: "OnChanged/OnSubmit", Type: "func(string)", Desc: "变化 / Enter 提交"},
			{Name: "Editor/FocusNode", Type: "取模型", Desc: "编辑器与焦点节点"},
		},
	})
}
