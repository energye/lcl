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
var (
	open    = widgets.NewButton("弹出通知")
	okBtn   = widgets.NewButton("成功通知")
	warnBtn = widgets.NewButton("警告通知")
)

func init() {
	open.OnClick = func() {
		if ctxRef == nil || ctxRef.Overlay == nil {
			return
		}
		n := widgets.NewNotification("保存成功", "数据已写入本地")
		n.Duration = 0 // 0 = 不自动消失（默认 3s）
		n.Show(ctxRef)
	}
	// type 语义图标（messagePalette 复用：绿勾/黄叹号——antd notification
	// 图标在标题左侧）。
	okBtn.OnClick = func() {
		if ctxRef == nil || ctxRef.Overlay == nil {
			return
		}
		n := widgets.NewNotification("操作成功", "部署任务已完成")
		n.Type = widgets.MessageSuccess
		n.Show(ctxRef)
	}
	warnBtn.OnClick = func() {
		if ctxRef == nil || ctxRef.Overlay == nil {
			return
		}
		n := widgets.NewNotification("磁盘告警", "C 盘剩余空间不足 10%")
		n.Type = widgets.MessageWarning
		n.Show(ctxRef)
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "notification", Title: "Notification 通知", Group: doc.GroupFeedback,
		What: []string{
			"全局通知：BandToast 层带，右上角弹出，超时自动消失、不抢焦点",
			"Duration=0 表示常驻；Type 语义图标（info/success/warning/error，messagePalette 复用）",
		},
		Demos: []doc.Demo{
			{Title: "弹出通知", Note: "常驻通知（需手动关），验证 BandToast 层带", Build: func(ctx *ui.BuildContext) ui.Widget {
				doc.CaptureCtx(&ctxRef, ctx)
				return open
			}},
			{Title: "语义图标", Note: "Type=success 绿勾 / warning 黄叹号——标题左侧图标（复用 Message 色板）", Build: func(ctx *ui.BuildContext) ui.Widget {
				doc.CaptureCtx(&ctxRef, ctx)
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					row := scene.Flex(scene.Horizontal, 12)
					row.Add(okBtn.Build(c))
					row.Add(warnBtn.Build(c))
					return row
				})
			}},
		},
		API: []doc.APIRow{
			{Name: "NewNotification", Type: "func(title, body string) *Notification", Desc: "构造（280×72 右上角）"},
			{Name: "Title/Body", Type: "string", Desc: "标题与正文"},
			{Name: "Type", Type: "MessageType", Default: "零值=无图标", Desc: "语义图标（success 绿勾/warning 黄叹号/error 红叉/loading 旋转）"},
			{Name: "Duration", Type: "time.Duration", Default: "3s", Desc: "超时；0 = 常驻"},
			{Name: "Show/Close/IsOpen", Type: "同 Modal", Desc: "打开/关闭/状态"},
		},
	})
}
