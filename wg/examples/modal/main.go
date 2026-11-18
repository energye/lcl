package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var ctxRef *ui.BuildContext

// open 包级持有：闭包内新建会让 hover 态随重建丢失。
var open = widgets.NewButton("打开模态框")

func init() {
	open.OnClick = func() {
		if ctxRef == nil || ctxRef.Overlay == nil {
			return
		}
		widgets.NewModal("示例模态框", doc.Node(func(ctx *ui.BuildContext) *scene.Node {
			return (&widgets.Label{Text: "模态遮罩 + 焦点域；Esc / 点遮罩关闭"}).Build(ctx)
		})).Show(ctxRef)
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "modal", Title: "Modal 模态框", Group: doc.GroupFeedback,
		What: []string{
			"模态对话框：BandModal 层带 + 遮罩 + TrapFocus 焦点域",
			"关闭：Esc / 点遮罩 / Confirm / Cancel 回调里调 Close",
			"计划（M10）：自定义 footer、loading 确认、info/success 系列预设",
		},
		Demos: []doc.Demo{
			{Title: "打开模态", Note: "真窗里按 Esc 或点遮罩可关闭", Build: func(ctx *ui.BuildContext) ui.Widget {
				doc.CaptureCtx(&ctxRef, ctx)
				return open
			}},
		},
		API: []doc.APIRow{
			{Name: "NewModal", Type: "func(title, content string|ui.Widget) *Modal", Desc: "构造"},
			{Name: "Confirm/Cancel", Type: "string", Default: "确定 / 取消", Desc: "按钮文字"},
			{Name: "Width", Type: "float32", Default: "400", Desc: "对话框宽度（居中 400×220）"},
			{Name: "OnConfirm/OnCancel", Type: "func()", Desc: "按钮回调"},
			{Name: "Show", Type: "func(ctx *ui.BuildContext)", Desc: "打开（需 Build 期捕获的 ctx）"},
			{Name: "Close/IsOpen", Type: "func() / func() bool", Desc: "关闭 / 状态查询"},
		},
	})
}
