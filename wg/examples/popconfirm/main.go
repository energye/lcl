package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	confirm := widgets.NewPopconfirm("确定删除这一项？", widgets.NewButton("删除"))
	launcher.Main(doc.Meta{
		ID: "popconfirm", Title: "Popconfirm 确认气泡", Group: doc.GroupFeedback,
		What: []string{
			"锚定在触发元素旁的确认气泡：比 Modal 轻量，适合危险操作二次确认",
			"OutsideClick / Esc 关闭；240×110",
		},
		Demos: []doc.Demo{
			{Title: "删除确认", Note: "点「删除」弹出气泡", Build: func(ctx *ui.BuildContext) ui.Widget {
				return confirm
			}},
		},
		API: []doc.APIRow{
			{Name: "NewPopconfirm", Type: "func(title string, child ui.Widget) *Popconfirm", Desc: "构造"},
			{Name: "Content", Type: "string", Desc: "确认正文"},
			{Name: "Placement", Type: "枚举", Desc: "气泡方位"},
			{Name: "Confirm/Cancel", Type: "string", Desc: "按钮文字"},
			{Name: "OnConfirm/OnCancel", Type: "func()", Desc: "按钮回调"},
		},
	})
}
