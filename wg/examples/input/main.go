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
			"IME 走 C28 链路：组合串经 SetComposingRegion 写入 Composing 区间，提交时替换选区",
			"编辑键（Backspace/Delete/Ctrl+A/Ctrl+←→）由 TextKeyHandler 统一处理，Enter/Tab 是语义动作",
			"计划（M8）：对齐 antd v5——variant/status/prefix/suffix/addon/字数统计",
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
			{Title: "文本编辑按键", Note: "编辑键统一处理；Enter/Tab 是语义动作不进入文本", Build: func(ctx *ui.BuildContext) ui.Widget {
				return edit
			}},
		},
		API: []doc.APIRow{
			{Name: "NewInput", Type: "func(placeholder string) *Input", Desc: "构造"},
			{Name: "Placeholder", Type: "string", Desc: "占位文字"},
			{Name: "Value", Type: "string", Desc: "当前文本（与 Editor 同步）"},
			{Name: "Multiline", Type: "bool", Default: "false", Desc: "多行模式"},
			{Name: "Width/Height", Type: "float32", Default: "主题控件尺寸", Desc: "显式尺寸"},
			{Name: "MaxLines", Type: "int", Default: "1", Desc: "多行最大行数"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnChanged", Type: "func(string)", Desc: "文本变化回调"},
			{Name: "OnSubmit", Type: "func(string)", Desc: "Enter 提交（单行）"},
			{Name: "Editor", Type: "func() *input.TextEditor", Desc: "取文本编辑器（选区/组合/撤销）"},
			{Name: "FocusNode", Type: "字段", Desc: "焦点节点，供 FocusManager 定位"},
		},
	})
}
