package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 四类型消息实例（包级持有：浮层句柄跨重建）。
var (
	msgInfo    = widgets.NewMessage(widgets.MessageInfo, "这是一条信息提示")
	msgSuccess = widgets.NewMessage(widgets.MessageSuccess, "操作成功")
	msgWarning = widgets.NewMessage(widgets.MessageWarning, "警告：输入内容可能丢失")
	msgError   = widgets.NewMessage(widgets.MessageError, "错误：网络连接失败")
	msgLoading = widgets.NewMessage(widgets.MessageLoading, "正在加载…")
)

// showFor 用 ctxRef 显示指定消息（demo 闭包外捕获 ctx）。
var ctxRef *ui.BuildContext

func main() {
	launcher.Main(doc.Meta{
		ID: "message", Title: "Message 全局提示", Group: doc.GroupFeedback,
		What: []string{
			"全局消息：顶部居中白底圆角浮层，语义色图标+文本（BandToast 层带）",
			"自动消失（DismissTimeout，默认 3s）；重复显示先关前一条（单例语义）",
			"loading 类型图标旋转（Spin 组合复用）",
		},
		Demos: []doc.Demo{
			{Title: "四类型触发", Note: "点击按钮显示对应消息（顶部居中，3s 自动消失）", Build: func(ctx *ui.BuildContext) ui.Widget {
				doc.CaptureCtx(&ctxRef, ctx)
				col := scene.Flex(scene.Vertical, 10)
				type row struct {
					btn *widgets.Button
					m   *widgets.Message
				}
				for _, r0 := range []row{
					{widgets.NewButton("显示信息"), msgInfo},
					{widgets.NewButton("显示成功"), msgSuccess},
					{widgets.NewButton("显示警告"), msgWarning},
					{widgets.NewButton("显示错误"), msgError},
					{widgets.NewButton("显示加载"), msgLoading},
				} {
					r := r0 // go 1.20 range 变量复用：闭包捕获前必须复制
					r.btn.OnClick = func() {
						if ctxRef != nil && ctxRef.Overlay != nil {
							r.m.Show(ctxRef)
						}
					}
					col.Add(r.btn.Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewMessage", Type: "func(typ MessageType, text string) *Message", Desc: "构造"},
			{Name: "Type", Type: "MessageType", Desc: "Info/Success/Warning/Error/Loading"},
			{Name: "Show", Type: "func(ctx *ui.BuildContext)", Desc: "顶部居中显示（需 Build 期捕获的 ctx）"},
			{Name: "Duration", Type: "time.Duration", Default: "3s", Desc: "自动消失时长"},
			{Name: "Close/IsOpen", Type: "func() / func() bool", Desc: "立即关闭 / 状态查询"},
		},
	})
}
