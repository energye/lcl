package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var inputs = []*widgets.Input{
	widgets.NewInput("第一个输入框"),
	widgets.NewInput("第二个输入框"),
	widgets.NewInput("第三个输入框"),
}

func init() {
	for _, in := range inputs {
		in.Width = 360
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "focus", Title: "焦点域与 Tab", Group: doc.GroupInput,
		What: []string{
			"FocusManager 管理 Tab 遍历：真窗里 Tab / Shift+Tab 顺序切换",
			"焦点改变后 Key/IME 投给新节点（SetTextEditor 通知平台输入法）",
			"组件用 FocusNode 参与焦点链，重建后由 controlBase 重指向新节点",
		},
		Demos: []doc.Demo{
			{Title: "三个输入框", Note: "Tab / Shift+Tab 切换；焦点框高亮", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				for _, in := range inputs {
					col.Add(in.Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "FocusManager", Type: "ctx.Focus", Desc: "焦点域（Tab 遍历/SetCurrent）"},
			{Name: "Input.FocusNode", Type: "字段", Desc: "组件的焦点节点"},
			{Name: "SetTextEditor", Type: "回调", Desc: "焦点进出时通知平台 IME（C28）"},
			{Name: "重建安全", Type: "约定", Desc: "focusNode 重建后自动重指向新节点"},
		},
	})
}
