package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var app = widgets.NewApp(nil)

func main() {
	launcher.Main(doc.Meta{
		ID: "app", Title: "App 应用容器", Group: doc.GroupControl,
		What: []string{
			"**全局上下文容器**：反馈类 API（message/notification/modal）经 App 实例注入——避免静态方法（antd App 语义）",
			"组合复用 Message（单例语义——重复 Show 先关前一条，BandToast 层带）",
			"未实现：notification/modal 上下文透传、ConfigProvider 主题、useApp hooks 语义（实例方法替代）",
		},
		Demos: []doc.Demo{
			{Title: "消息注入", Note: "点按钮 → app.ShowMessage(Success) → BandToast 消息（单例语义）", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 12)
				btn := &widgets.Button{Text: "触发全局消息", Type: widgets.TypePrimary, OnClick: func() {
					app.ShowMessage(widgets.MessageSuccess, "保存成功：所有更改已同步")
				}}
				col.Add(btn.Build(ctx))
				col.Add((&widgets.Label{Text: "App 组合暴露反馈宿主——消息 ctx 由最近一次 Build 捕获"}).Build(ctx))
				app.Child = doc.Node(func(*ui.BuildContext) *scene.Node { return col })
				return app
			}},
		},
		API: []doc.APIRow{
			{Name: "NewApp", Type: "func(child ui.Widget) *App", Desc: "构造（应用内容）"},
			{Name: "ShowMessage", Type: "func(typ MessageType, text string)", Desc: "全局消息（ctx 由最近 Build 捕获）"},
			{Name: "MessageDuration", Type: "int", Desc: "消息时长毫秒（0 默认 3000）"},
			{Name: "LastMessage", Type: "func() string", Desc: "最近消息文本（诊断）"},
		},
	})
}
