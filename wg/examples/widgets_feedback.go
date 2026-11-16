package examples

import (
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// captureCtx 在 Build 时记下上下文，供后续点击回调打开浮层。
// 上下文里的 Overlay/Input/Theme 指针由 Container 长期持有，跨帧有效。
func captureCtx(slot **ui.BuildContext, ctx *ui.BuildContext) { *slot = ctx }

func init() {
	{
		var ctxRef *ui.BuildContext
		open := widgets.NewButton("打开模态框")
		open.OnClick = func() {
			if ctxRef == nil || ctxRef.Overlay == nil {
				return
			}
			widgets.NewModal("示例模态框", From(func(ctx *ui.BuildContext) *scene.Node {
				return (&widgets.Label{Text: "模态遮罩 + 焦点域；Esc / 点遮罩关闭"}).Build(ctx)
			})).Show(ctxRef)
		}
		Register(Example{ID: "modal", Name: "Modal 模态框", Group: GroupFeedback,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				captureCtx(&ctxRef, ctx)
				return Section(ctx, "Modal", "BandModal + 遮罩 + TrapFocus，真窗里按 Esc 可关闭", open.Build(ctx))
			}})
	}

	{
		var ctxRef *ui.BuildContext
		open := widgets.NewButton("从左侧滑出抽屉")
		open.OnClick = func() {
			if ctxRef == nil || ctxRef.Overlay == nil {
				return
			}
			widgets.NewDrawer("侧边抽屉", From(func(ctx *ui.BuildContext) *scene.Node {
				return (&widgets.Label{Text: "Drawer 从边缘滑入，锚定视口一侧"}).Build(ctx)
			})).Show(ctxRef)
		}
		Register(Example{ID: "drawer", Name: "Drawer 抽屉", Group: GroupFeedback,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				captureCtx(&ctxRef, ctx)
				return Section(ctx, "Drawer", "左/右/上/下四向，关闭策略同 Modal", open.Build(ctx))
			}})
	}

	{
		var ctxRef *ui.BuildContext
		open := widgets.NewButton("弹出通知")
		open.OnClick = func() {
			if ctxRef == nil || ctxRef.Overlay == nil {
				return
			}
			n := widgets.NewNotification("保存成功", "数据已写入本地")
			n.Duration = 0
			n.Show(ctxRef)
		}
		Register(Example{ID: "notification", Name: "Notification 通知", Group: GroupFeedback,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				captureCtx(&ctxRef, ctx)
				return Section(ctx, "Notification", "BandToast + 超时自动消失，不抢焦点", open.Build(ctx))
			}})
	}

	{
		confirm := widgets.NewPopconfirm("确定删除这一项？", widgets.NewButton("删除"))
		Register(Example{ID: "popconfirm", Name: "Popconfirm 确认气泡", Group: GroupFeedback,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				return Section(ctx, "Popconfirm", "锚定在按钮上的确认气泡，OutsideClick / Esc 关闭", confirm.Build(ctx))
			}})
	}
}
