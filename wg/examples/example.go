// Package examples 是 wg 引擎的可运行示例集与逐项容器测试。
//
// 每个 Example 是一个独立场景：组件示例展示一个组件的完整用法，引擎功能
// 示例展示一条引擎能力（布局/图层/命中/IME…）。所有示例统一用 ui.Widget
// 契约产出节点，因此同一份代码既能跑在真窗 Container 上，也能跑在 headless
// Container 里做自动化断言——这是本包存在的意义。
//
// 真窗入口：go run ./wg/cmd/wg-examples
// 逐项容器测试：go test ./wg/examples/...
package examples

import (
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 分组名。画廊按 Group 分栏，测试报告也用它做归类。
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

// Example 是一条可运行示例。
//
// Build 与 ui.Widget.Build 同签名，因此任何示例都能直接塞进
// MountWidget——容器测试走的就是这条路，不另造测试专用上下文。
type Example struct {
	ID    string // 稳定标识，测试名与画廊定位都用它
	Name  string // 显示名
	Group string
	Build func(*ui.BuildContext) *scene.Node
}

// All 是全部示例，注册表按分组插入顺序排列（画廊顺序即此顺序）。
var All []Example

// Register 登记一条示例。重复 ID 会被拒绝，避免测试名冲突。
func Register(e Example) {
	for _, old := range All {
		if old.ID == e.ID {
			panic("examples: duplicate ID " + e.ID)
		}
	}
	All = append(All, e)
}

// widgetExample 把 Example 包装成 ui.Widget，供 Container.MountWidget 使用。
// 容器测试与真窗画廊都通过它进入同一条构建链路。
type widgetExample struct {
	build func(*ui.BuildContext) *scene.Node
}

func (w widgetExample) Build(ctx *ui.BuildContext) *scene.Node { return w.build(ctx) }

// Widget 返回可作为 Container.MountWidget 参数的组件。
func (e Example) Widget() ui.Widget { return widgetExample{build: e.Build} }

// From 把一段构建函数包装成 ui.Widget，便于把 scene 节点树交给
// 需要 Child 的组合组件（Scroll/Card/Modal 等）。每次 Build 都返回新节点。
func From(build func(*ui.BuildContext) *scene.Node) ui.Widget {
	return widgetExample{build: build}
}

// ByID 按 ID 查找示例。
func ByID(id string) (Example, bool) {
	for _, e := range All {
		if e.ID == id {
			return e, true
		}
	}
	return Example{}, false
}

// ── 示例里的通用排版工具 ────────────────────────────────────────────

// Heading 产出示例标题。
func Heading(ctx *ui.BuildContext, text string) *scene.Node {
	return (&widgets.Label{Text: text, Size: ctx.ThemeOr().Font.SizeTitle}).Build(ctx)
}

// Note 产出一行说明。
func Note(ctx *ui.BuildContext, text string) *scene.Node {
	return (&widgets.Label{Text: text, Size: ctx.ThemeOr().Font.SizeBase}).Build(ctx)
}

// Section 把标题、说明和内容纵向排成一页，供各示例复用。
func Section(ctx *ui.BuildContext, title, note string, body *scene.Node) *scene.Node {
	col := scene.Flex(scene.Vertical, 10).SetFlex(1, scene.FlexTight)
	col.Add(Heading(ctx, title))
	col.Add(Note(ctx, note))
	if body != nil {
		col.Add(body)
	}
	return col
}
