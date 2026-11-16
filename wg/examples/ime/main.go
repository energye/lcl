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
	inp    = widgets.NewInput("点这里，然后切换中文输入法")
	status = widgets.NewButton("等待组合串…")
)

func init() { inp.Width = 420 }

func main() {
	launcher.Main(doc.Meta{
		ID: "ime", Title: "IME 组合与选区", Group: doc.GroupInput,
		What: []string{
			"文本输入法链路（C28）：TextKeyHandler 与 IME 结果分离——组合串经宿主 IME 事件进编辑器",
			"组合串提交时替换当前选区，而不是追加到末尾",
			"演示按钮在 headless 下模拟 ni hao → 你好 的完整组合流程；真窗可直接用系统输入法",
		},
		Demos: []doc.Demo{
			{Title: "组合串替换选区", Note: "点「模拟」看 headless 模拟结果；真窗里直接切输入法", Build: func(ctx *ui.BuildContext) ui.Widget {
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
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "TextEditor.SetComposingRegion", Type: "func(start, end int, text string)", Desc: "组合串区间（rune 偏移）"},
			{Name: "CommitComposingWith", Type: "func(result string)", Desc: "提交组合串并替换选区"},
			{Name: "c.IME().Handle", Type: "host.IMFEvent", Desc: "平台 IME 事件入口（Start/Update/Result/End）"},
			{Name: "CursorRect", Type: "func() geom.Rect", Desc: "候选框定位（当前为整个输入框矩形）"},
		},
	})
}
