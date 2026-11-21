package widgets

// ── App 应用容器（M12 收尾）──────────────────────────────────────
//
// antd：<App><content/></App> 全局上下文容器——**避免静态方法**：
// 反馈类 API（message/notification/modal）经 App 实例注入而非散落
// 全局。第一轮组合复用 **Message**（单例语义——重复 Show 先关前
// 一条）；notification/modal 实例透传后续轮。
// 未实现（如实标注）：notification/modal 上下文、ConfigProvider
// 主题透传、useApp hooks 语义（引擎无 hooks——实例方法替代）。

import (
	"time"

	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// App 是应用上下文容器（组合暴露反馈宿主）。
type App struct {
	Child ui.Widget
	// MessageDuration 全局消息时长（毫秒；0 默认 3000）。
	MessageDuration int
	Width           float32

	message  *Message
	lastCtx  *ui.BuildContext
	lastText string
}

// NewApp 创建应用容器。
func NewApp(child ui.Widget) *App {
	return &App{Child: child}
}

// Build 实现 ui.Widget。
func (a *App) Build(ctx *ui.BuildContext) *scene.Node {
	a.lastCtx = ctx
	w := a.Width
	if w > 0 {
	}
	var n *scene.Node
	if a.Child != nil {
		n = a.Child.Build(ctx)
	} else {
		n = scene.Empty()
	}
	wrap := scene.Flex(scene.Vertical, 0)
	wrap.SetHit(scene.HitSpec{Label: "app"})
	if w > 0 {
		wrap.SetWidth(w)
	}
	wrap.Add(n)
	return wrap
}

// ShowMessage 显示全局消息（Message 单例语义——重复调用关前一条；
// ctx 由最近一次 Build 捕获）。
func (a *App) ShowMessage(typ MessageType, text string) {
	if a.lastCtx == nil {
		return
	}
	if a.message == nil {
		a.message = &Message{}
		if a.MessageDuration > 0 {
			a.message.Duration = time.Duration(a.MessageDuration) * time.Millisecond
		}
	}
	a.message.Type = typ
	a.message.Text = text
	a.lastText = text
	a.message.Show(a.lastCtx)
}

// LastMessage 最近一次消息文本（诊断用）。
func (a *App) LastMessage() string { return a.lastText }
