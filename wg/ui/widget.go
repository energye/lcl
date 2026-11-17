package ui

import (
	"time"

	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
)

// ── 组件基座 ────────────────────────────────────────────────────────
//
// 只有两层：**Widget → Node**（不抄 Flutter 的 Element 三层）。
// 理由：Go 没有 JSX，三层带来的 diff 灵活性在这里的收益低于复杂度；
// 组件实例本身长期存活、持有状态、可被父级按 Key 复用，它**就是** Element。
//
// 两条并存的写法（§A1.1）：
//   - 声明式：Build(ctx) → Node，常规组件；
//   - 直接式：持有 *Node 原地改属性（图表、画布图元），跳过重建但不是跳过管线。

// Key 是节点复用标识。
//
// 同层同 Key → 复用节点（保留布局/缓存/滚动位置）；无 Key → 按位置复用。
type Key string

// Widget 是组件。
type Widget interface {
	// Build 产出节点树。每次调用都应返回**新**节点，不要复用旧节点。
	Build(ctx *BuildContext) *scene.Node
}

// MediaQuery 是媒体查询（DPI/尺寸/文字缩放/动效偏好）。
type MediaQuery struct {
	DPI       geom.Size
	Scale     float32
	TextScale float32
	// ReduceMotion 为真时动画时长归零但仍走状态变更（§A4 规则 5）。
	ReduceMotion bool
}

// BuildContext 是 Build 期间可用的环境。
type BuildContext struct {
	Theme *theme.Data
	Media MediaQuery
	Focus *input.FocusManager
	// Input 是事件分发器：组件用它注册自己的 Handler。
	Input *input.Dispatcher
	// SetTextEditor 在输入框获得/失去焦点时通知平台 IME。
	// nil 表示没有平台输入法（headless 或还没绑定窗体）。
	SetTextEditor func(*input.TextEditor)
	Overlay       *overlay.Host
	Scheduler     *frame.Scheduler
	// Inherited 读取祖先提供的配置（主题/媒体/滚动）。
	Inherited func(key any) any

	// RequestFrame 请求下一帧（组件用它驱动动画/异步后重绘）。
	// nil 表示宿主不驱动帧泵（部分 headless 测试）。
	RequestFrame func()
	// InvalidateRoot 把**根 widget** 置脏并请求一帧。
	//
	// 深层组件的状态变化必须走这里：容器只重建"根 widget 的 Build 结果"，
	// 深层组件只调 MarkDirty（不置根脏）的话，静止窗口（C14）永远不会
	// 重绘——真窗里表现为"点了没反应"（M8 对拍发现的引擎级缺陷）。
	// nil 表示宿主不重建组件树（部分 headless 测试）。
	InvalidateRoot func()
	// Viewport 是当前视口尺寸（DIP）。Build 期据此做自适应排版。
	// 零值表示未知（部分 headless 测试）。
	//
	// **直接读字段不会登记"依赖视口"**：resize 时容器不会为你重建
	// （成熟框架语义：只为声明了依赖的组件重建）。需要响应视口变化的
	// 组件请调用 ViewportSize()。
	Viewport geom.Size
	// ViewportUse 由宿主注入：ViewportSize() 调用它登记依赖。
	ViewportUse func()
	// ReportIMECaret 由输入控件在 Build 时上报光标锚点：node 是控件节点，
	// local 是相对该节点的光标矩形（像素）。传 node=nil 表示登记者放弃光标。
	//
	// owner 是登记者的身份（通常是控件自身）：页面上可能有多个输入框，
	// 未聚焦的输入框也会按"我没有光标"上报 nil。宿主必须按 owner 归属
	// 处理，否则后构建的未聚焦输入框会把已聚焦输入框刚上报的锚点清掉
	//（真窗表现：刚点进输入框时 IME 候选框/插入符定位丢失）。
	ReportIMECaret func(owner any, node *scene.Node, local geom.Rect)
	// ReportCaret 由输入控件上报"正在闪烁的插入符节点"。非 nil 时宿主
	// 按系统节奏（530ms）切换其可见性并按时申请下一帧；node=nil 且
	// owner 与当前登记者一致时才停止闪烁（语义同 ReportIMECaret）。
	// 编辑/移动光标时重新上报同一节点会重置相位（刚敲完字不会立刻变暗）。
	ReportCaret func(owner any, node *scene.Node)
}

// ViewportSize 返回视口尺寸（DIP）并登记"本组件依赖视口"。
//
// resize 时容器只会重建登记过依赖的组件树；读字段不登记。
func (c *BuildContext) ViewportSize() geom.Size {
	if c == nil {
		return geom.Size{}
	}
	if c.ViewportUse != nil {
		c.ViewportUse()
	}
	return c.Viewport
}

// ---- M8 补充字段由单独编辑添加 ----

// ThemeOr 返回主题（nil 时返回默认主题）。
func (c *BuildContext) ThemeOr() *theme.Data {
	if c == nil || c.Theme == nil {
		return theme.Default()
	}
	return c.Theme
}

// Motion 返回动效时长（考虑 ReduceMotion）。
func (c *BuildContext) Motion(d time.Duration) time.Duration {
	if c != nil && c.Media.ReduceMotion {
		return 0
	}
	return d
}

// Base 是可持有状态的组件嵌入体（等价于 Flutter 的 State，但不需要两份对象）。
type Base struct {
	key     Key
	dirty   bool
	mounted bool
	// invalidate 在 SetState 后请求一帧（由宿主注入）。
	invalidate func()
	builds     int
}

// SetKey 设置复用标识。
func (b *Base) SetKey(k Key) *Base { b.key = k; return b }

// Key 返回复用标识。
func (b *Base) Key() Key { return b.key }

// Dirty 表示需要重建。
func (b *Base) Dirty() bool { return b.dirty }

// ClearDirty 清除脏标记（宿主在重建后调用）。
func (b *Base) ClearDirty() { b.dirty = false }

// Builds 返回累计重建次数（诊断用）。
func (b *Base) Builds() int { return b.builds }

// Mounted 表示已挂载。
func (b *Base) Mounted() bool { return b.mounted }

// Attach 注入"请求一帧"的回调（挂载时调用）。
func (b *Base) Attach(invalidate func()) {
	b.invalidate = invalidate
	if !b.mounted {
		b.mounted = true
	}
}

// Detach 解挂。
func (b *Base) Detach() {
	b.mounted = false
	b.invalidate = nil
}

// SetState 执行 fn 并标脏，下一帧重建。
//
// 为什么必须包一层 fn 而不是让调用方自己标脏：忘记标脏是最常见的一类
// "改了数据界面不动"的 bug，而包一层可以让"改状态"和"要重建"绑死。
func (b *Base) SetState(fn func()) {
	if fn != nil {
		fn()
	}
	b.dirty = true
	if b.invalidate != nil {
		b.invalidate()
	}
}

// MarkDirty 只标脏不执行（用于异步回调里改完字段后）。
func (b *Base) MarkDirty() {
	b.dirty = true
	if b.invalidate != nil {
		b.invalidate()
	}
}

// Rebuild 在一次重建后记账。
func (b *Base) Rebuild() {
	b.dirty = false
	b.builds++
}

// BindState 把有状态组件接到宿主：其 MarkDirty/SetState 会请求一帧并把
// 根 widget 置脏（深层组件不接线的话，静止窗口 C14 下点击永不重绘）。
// 走 controlBase 的控件已在 attach 里接线；其余有状态组件在 Build 里调用。
func (ctx *BuildContext) BindState(b interface{ Attach(func()) }) {
	if ctx == nil || b == nil {
		return
	}
	switch {
	case ctx.InvalidateRoot != nil:
		b.Attach(ctx.InvalidateRoot)
	case ctx.RequestFrame != nil:
		b.Attach(ctx.RequestFrame)
	}
}

// SetDirtyOnly 只置脏位，不触发 invalidate 回调。
//
// 供容器的 InvalidateRoot 使用：容器要把根 widget 标记为"需要重建"，
// 但根自己的 invalidate 往往就是"请求帧/再置脏"——走 MarkDirty 会形成
// MarkDirty→InvalidateRoot→MarkDirty 的无限递归（M8 真窗探针实测爆内存）。
func (b *Base) SetDirtyOnly() {
	b.dirty = true
}
