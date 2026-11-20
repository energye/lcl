package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var ctxRef *ui.BuildContext

// open/预设按钮包级持有：闭包内新建会让 hover 态随重建丢失。
var (
	open = widgets.NewButton("打开模态框")
	// 预设系列：info/success/warning/error 四档（antd Modal 预设确认框）。
	infoBtn    = widgets.NewButton("信息提示")
	successBtn = widgets.NewButton("操作成功")
	warningBtn = widgets.NewButton("警告")
	errorBtn   = widgets.NewButton("错误")
)

func init() {
	open.OnClick = func() {
		if ctxRef == nil || ctxRef.Overlay == nil {
			return
		}
		widgets.NewModal("示例模态框", doc.Node(func(ctx *ui.BuildContext) *scene.Node {
			return (&widgets.Label{Text: "模态遮罩 + 焦点域；Esc / 点遮罩关闭"}).Build(ctx)
		})).Show(ctxRef)
	}
	// 预设构造器：图标圆底（语义色）+标题+内容+单"知道了"按钮。
	show := func(kind widgets.ModalKind, title string) func() {
		return func() {
			if ctxRef == nil || ctxRef.Overlay == nil {
				return
			}
			widgets.NewModalPreset(kind, title, doc.Node(func(ctx *ui.BuildContext) *scene.Node {
				return (&widgets.Label{Text: "这是一段预设提示内容"}).Build(ctx)
			}), nil).Show(ctxRef)
		}
	}
	infoBtn.OnClick = show(widgets.ModalInfo, "信息提示")
	successBtn.OnClick = show(widgets.ModalSuccess, "操作成功")
	warningBtn.OnClick = show(widgets.ModalWarning, "警告")
	errorBtn.OnClick = show(widgets.ModalError, "错误")
}

func main() {
	launcher.Main(doc.Meta{
		ID: "modal", Title: "Modal 模态框", Group: doc.GroupFeedback,
		What: []string{
			"模态对话框：BandModal 层带 + 遮罩 + TrapFocus 焦点域",
			"关闭：Esc / 点遮罩 / Confirm / Cancel 回调里调 Close",
			"预设确认框：Modal.info/success/warning/error——状态图标+单「知道了」按钮",
		},
		Demos: []doc.Demo{
			{Title: "打开模态", Note: "真窗里按 Esc 或点遮罩可关闭", Build: func(ctx *ui.BuildContext) ui.Widget {
				doc.CaptureCtx(&ctxRef, ctx)
				return open
			}},
			{Title: "预设系列", Note: "Modal.info/success/warning/error：状态图标圆底+单「知道了」按钮", Build: func(ctx *ui.BuildContext) ui.Widget {
				doc.CaptureCtx(&ctxRef, ctx)
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					row := scene.Flex(scene.Horizontal, 12)
					row.Add(infoBtn.Build(c))
					row.Add(successBtn.Build(c))
					row.Add(warningBtn.Build(c))
					row.Add(errorBtn.Build(c))
					return row
				})
			}},
		},
		API: []doc.APIRow{
			{Name: "NewModal", Type: "func(title, content string|ui.Widget) *Modal", Desc: "构造"},
			{Name: "Confirm/Cancel", Type: "string", Default: "确定 / 取消", Desc: "按钮文字"},
			{Name: "Width", Type: "float32", Default: "400", Desc: "对话框宽度（居中 400×220）"},
			{Name: "OnConfirm/OnCancel", Type: "func()", Desc: "按钮回调"},
			{Name: "Show", Type: "func(ctx *ui.BuildContext)", Desc: "打开（需 Build 期捕获的 ctx）"},
			{Name: "Close/IsOpen", Type: "func() / func() bool", Desc: "关闭 / 状态查询"},
			{Name: "NewModalPreset", Type: "func(kind ModalKind, title string, content ui.Widget, onOk func()) *Modal", Desc: "预设确认框（图标+单 OK 按钮）"},
			{Name: "Kind", Type: "ModalKind", Desc: "ModalInfo/Success/Warning/Error"},
			{Name: "OKText", Type: "string", Default: "知道了", Desc: "预设按钮文字"},
		},
	})
}
