// Package doc 是 M7 的文档窗框架：一窗一种，窗内三段式——
// 是什么/何时使用（文字）→ 演示块×N（标题+效果说明+可交互 demo）→ API 属性表。
//
// 每个功能是 examples/<feature>/main.go（package main，单独 go run），
// main 里组装一份 Meta 交给 launcher.Run；本包负责把 Meta 渲染成文档页。
//
// 演示中的有状态组件请在 main 包里用包级变量持有实例（参照 widgets.DemoForm
// 的做法）：容器重建时 Build 会被再次调用，只有包级实例能跨重建保住交互状态。
package doc

import (
	"github.com/energye/lcl/wg/ui"
)

// 分组名。与旧 examples 注册表的九组一致，launcher --meta 汇总时用它归类。
const (
	GroupBasic    = "基础组件"
	GroupControl  = "输入控制"
	GroupDisplay  = "展示组件"
	GroupData     = "数据组件"
	GroupFeedback = "反馈浮层"
	GroupChart    = "图表"
	GroupLayout   = "布局引擎"
	GroupVisual   = "绘制与合成"
	GroupInput    = "事件与输入"
)

// APIRow 是 API 属性表的一行：参数名 / 类型或签名 / 默认值 / 说明。
type APIRow struct {
	Name    string
	Type    string
	Default string // 可空
	Desc    string
}

// Demo 是一个可交互演示块。
//
// Build 每次文档页重建都会被调用；有状态的组件实例请在 main 包
// 用包级变量持有，Build 里直接返回实例。
type Demo struct {
	Title string // 演示名，如"主要按钮"
	Note  string // 效果说明，讲清这块演示展示了什么
	Build func(ctx *ui.BuildContext) ui.Widget
}

// Meta 是一个功能窗的完整文档描述。
//
// ID 必须与目录名一致（examples/<ID>/main.go），--selfcheck 的枚举门禁
// 靠这条约定把"目录"与"窗口"对上。
type Meta struct {
	ID    string   // 目录名，如 button
	Title string   // 显示名，如 "Button 按钮"
	Group string   // 分组常量之一
	What  []string // 是什么/何时使用，一行一条
	Demos []Demo
	API   []APIRow // API 属性表；能力窗放 scene/render 构造函数，组件窗放属性/事件
}

// Registered 是本进程内已登记的功能窗（launcher 与测试用）。
var Registered []Meta

// Register 登记一个功能窗。重复 ID 直接 panic——文档窗的唯一性问题
// 必须在启动时就暴露，而不是靠门禁事后抓。
func Register(m Meta) {
	for _, old := range Registered {
		if old.ID == m.ID {
			panic("doc: duplicate feature ID " + m.ID)
		}
	}
	Registered = append(Registered, m)
}
