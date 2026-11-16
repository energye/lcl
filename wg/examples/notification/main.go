package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var ctxRef *ui.BuildContext

func main() {
	open := widgets.NewButton("弹出通知")
	open.OnClick = func() {
		if ctxRef == nil || ctxRef.Overlay == nil {
			return
		}
		n := widgets.NewNotification("保存成功", "数据已写入本地")
		n.Duration = 0 // 0 = 不自动消失（默认 3s）
		n.Show(ctxRef)
	}
	launcher.Main(doc.Meta{
		ID: "notification", Title: "Notification 通知", Group: doc.GroupFeedback,
		What: []string{
			"全局通知：BandToast 层带，右上角弹出，超时自动消失、不抢焦点",
			"Duration=0 表示常驻；计划（M10）：info/success/warning/error 类型与图标",
		},
		Demos: []doc.Demo{
			{Title: "弹出通知", Note: "常驻通知（需手动关），验证 BandToast 层带", Build: func(ctx *ui.BuildContext) ui.Widget {
				doc.CaptureCtx(&ctxRef, ctx)
				return open
			}},
		},
		API: []doc.APIRow{
			{Name: "NewNotification", Type: "func(title, body string) *Notification", Desc: "构造（280×72 右上角）"},
			{Name: "Title/Body", Type: "string", Desc: "标题与正文"},
			{Name: "Duration", Type: "time.Duration", Default: "3s", Desc: "超时；0 = 常驻"},
			{Name: "Show/Close/IsOpen", Type: "同 Modal", Desc: "打开/关闭/状态"},
		},
	})
}
