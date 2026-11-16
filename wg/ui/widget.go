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
	Theme     *theme.Data
	Media     MediaQuery
	Focus     *input.FocusManager
	Overlay   *overlay.Host
	Scheduler *frame.Scheduler
	// Inherited 读取祖先提供的配置（主题/媒体/滚动）。
	Inherited func(key any) any
}

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
