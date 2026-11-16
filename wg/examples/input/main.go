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
	status = widgets.NewButton("等待输入…")
	edit   = widgets.NewInput("试试 Backspace / Delete / Ctrl+A / Ctrl+←→")
)

func init() {
	inp.Width = 420
	inp.OnChanged = func(s string) {
		status.Text = "当前内容：" + s
		status.MarkDirty()
	}
	edit.Width = 460
}

func main() {
	launcher.Main(doc.Meta{
		ID: "input", Title: "Input 输入框", Group: doc.GroupControl,
		What: []string{
			"单行/多行文本输入：TextEditor 持有文本/选区/撤销栈，跨重建不丢",
			"对齐 antd：variant（outlined/filled/borderless）、status（error/warning）、prefix/suffix、addonBefore/addonAfter、count、allowClear、三档尺寸",
			"IME 走 C28 链路：组合串提交时替换选区；编辑键由 TextKeyHandler 统一处理",
			"未实现：status 反馈图标、autosize",
		},
		Demos: []doc.Demo{
			{Title: "受控输入与 IME", Note: "真窗里点输入框并切换中文输入法；OnChanged 实时回显", Build: func(ctx *ui.BuildContext) ui.Widget {
				clear := widgets.NewButton("清空")
				clear.OnClick = func() {
					inp.Editor().SetValue(input.EditingValue{})
					inp.Value = ""
					inp.MarkDirty()
				}
				col := scene.Flex(scene.Vertical, 10)
				col.Add(inp.Build(ctx))
				row := scene.Flex(scene.Horizontal, 8)
				row.Add(status.Build(ctx))
				row.Add(clear.Build(ctx))
				col.Add(row)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "三种 variant", Note: "outlined（默认）/ filled（灰底）/ borderless（无框）", Build: func(ctx *ui.BuildContext) ui.Widget {
				o := widgets.NewInput("outlined")
				o.Width = 260
				f := widgets.NewInput("filled")
				f.Width, f.Variant = 260, widgets.InputFilled
				bl := widgets.NewInput("borderless")
				bl.Width, bl.Variant = 260, widgets.InputBorderless
				return widgets.Column{Gap: 8, Children: []ui.Widget{o, f, bl}}
			}},
			{Title: "status 与尺寸", Note: "error 红边 / warning 黄边；small/middle/large", Build: func(ctx *ui.BuildContext) ui.Widget {
				err := widgets.NewInput("error 状态")
				err.Width, err.Status = 260, widgets.InputStatusError
				warn := widgets.NewInput("warning 状态")
				warn.Width, warn.Status = 260, widgets.InputStatusWarning
				sm := widgets.NewInput("small")
				sm.Width, sm.Size = 260, ui.SizeSmall
				lg := widgets.NewInput("large")
				lg.Width, lg.Size = 260, ui.SizeLarge
				return widgets.Column{Gap: 8, Children: []ui.Widget{err, warn, sm, lg}}
			}},
			{Title: "前后缀与清除", Note: "prefix 图标 + suffix 图标；输入后 hover 显示 allowClear 清除", Build: func(ctx *ui.BuildContext) ui.Widget {
				prefixed := widgets.NewInput("站内搜索")
				prefixed.Width, prefixed.AllowClear = 300, true
				prefixed.Prefix = &widgets.Icon{Name: "search", Size: 14}
				suffixed := widgets.NewInput("金额")
				suffixed.Width = 300
				suffixed.Suffix = &widgets.Icon{Name: "info-circle", Size: 14}
				return widgets.Column{Gap: 8, Children: []ui.Widget{prefixed, suffixed}}
			}},
			{Title: "addon 与 count", Note: "addonBefore/addonAfter 灰块；Count 右下角字数", Build: func(ctx *ui.BuildContext) ui.Widget {
				withAddon := widgets.NewInput("example.com 前缀")
				withAddon.Width = 300
				withAddon.AddonAfter = doc.Node(func(ctx *ui.BuildContext) *scene.Node {
					return (&widgets.Label{Text: ".com", Size: 13, Color: ctx.ThemeOr().Color.TextMuted}).Build(ctx)
				})
				counted := widgets.NewInput("看右下角字数")
				counted.Width, counted.Count = 300, true
				return widgets.Column{Gap: 12, Children: []ui.Widget{withAddon, counted}}
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
