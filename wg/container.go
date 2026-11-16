package wg

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/wg/composite"
	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/raster/gdi"
	"github.com/energye/lcl/wg/raster/software"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
)

// Painter 是"把一帧的内容录进 Canvas"的回调。
//
// M1 的过渡形态：场景树（scene）与组件（ui）在 M2/M3 才落地，在那之前
// 由调用方直接往录制器里画。这不是临时 API——它的形状与将来的
// scene.Node.Paint(canvas) 一致，迁移时只需把回调体换成树的遍历。
type Painter func(c render.Canvas)

// state 是容器生命周期。
type state uint8

const (
	stateCreated state = iota
	stateRunning
	stateClosed
)

// String 返回可读名。
func (s state) String() string {
	switch s {
	case stateCreated:
		return "created"
	case stateRunning:
		return "running"
	case stateClosed:
		return "closed"
	default:
		return "?"
	}
}

// Container 是**唯一的父容器**，也是整个引擎的组合根。
//
// 它持有：宿主（Form + 根画布）、帧时钟、光栅后端、合成器与录制器。
// 使用者只与 Container 打交道，不直接接触 Surface 树与像素。
//
// 线程模型：所有公开方法都是线程安全的（内部单锁）。但**结构性变更**
// （Resize/Close）只允许在帧外调用；帧内变更必须走 RequestFrame
// 投递到下一帧（契约 C18）。
type Container struct {
	mu    sync.Mutex
	form  lcl.IEngForm
	opts  Options
	host  *host.Host
	clock frame.FrameClock

	// 渲染管线（首次出帧时惰性装配）
	backend         raster.Backend
	backendOverride raster.Backend
	textPainter     composite.TextPainter
	comp            *composite.Compositor
	rec             *render.Recorder
	paint           Painter

	// 内容树（M2 起是主路径；M1 的 Painter 仍然可用，两者可叠加）
	scene   *scene.Node // 组合根：内容 + 浮层根
	content *scene.Node
	ov      *overlay.Host
	focus   *input.FocusManager
	disp    *input.Dispatcher
	th      *theme.Data
	// layoutAll 表示下一帧必须整树重排（挂载/换主题/resize）。
	layoutAll bool

	// 帧泵（Run 启动；无窗口模式下由调用方手动 FrameOnce/Tick）
	sched  *frame.Scheduler
	ticker *frame.ManualTicker

	// lastCursor 是最近一次命中的光标（平台层据此换鼠标指针）。
	lastCursor scene.Cursor

	index     int
	last      time.Duration
	requested bool
	// lastDamage 是最近一帧提交的脏区（诊断与门禁用）。
	lastDamage geom.Rect
	// deferred 记录"帧内有人请求了下一帧"。帧内回调已经持锁，
	// 直接走 RequestFrame 会自己和自己抢锁（死锁），因此先存标志、帧末合并。
	deferred atomic.Bool
	st       state
}

// NewContainer 创建容器并绑定窗体。
//
// form 为 nil 时得到一个**无窗口容器**（等价于 NewHeadless）：根画布尺寸为 0，
// 可手动 Resize，用于自检与出图。这条路径不访问任何 LCL 符号。
func NewContainer(form lcl.IEngForm, opts Options) *Container {
	c := &Container{
		form:  form,
		opts:  opts.normalized(),
		clock: frame.NewWallClock(),
		st:    stateCreated,
	}
	c.host = host.New(form)
	return c
}

// NewHeadless 创建一个不绑定窗口的容器，用于无显示器环境（CI、golden 出图、
// 契约自检）。尺寸由 Resize 指定。
func NewHeadless(opts Options) *Container { return NewContainer(nil, opts) }

// Form 返回绑定的窗体；无窗口模式下为 nil。
func (c *Container) Form() lcl.IEngForm { return c.form }

// Options 返回生效的选项（已补齐默认值）。
func (c *Container) Options() Options { return c.opts }

// Host 返回宿主（画布树入口）。上层一般只需要 Root()。
func (c *Container) Host() *host.Host { return c.host }

// Root 返回根画布；容器关闭后返回 nil。
func (c *Container) Root() *host.Surface {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.st == stateClosed || c.host == nil {
		return nil
	}
	return c.host.Root()
}

// State 返回生命周期状态名（诊断用）。
func (c *Container) State() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.st.String()
}

// Closed 表示容器已关闭。
func (c *Container) Closed() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.st == stateClosed
}

// Clock 返回当前帧时钟。
func (c *Container) Clock() frame.FrameClock {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.clock
}

// SetClock 替换帧时钟（必须在 Run 之前调用）。
func (c *Container) SetClock(clk frame.FrameClock) {
	if clk == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.st == stateRunning {
		return
	}
	c.clock = clk
	c.last = clk.Now()
}

// SetPainter 设置每帧的内容录制回调。
func (c *Container) SetPainter(p Painter) {
	c.mu.Lock()
	c.paint = p
	c.mu.Unlock()
}

// SetBackend 替换光栅后端（A/B 对照与诊断用，必须在首次出帧前调用）。
func (c *Container) SetBackend(b raster.Backend) {
	c.mu.Lock()
	c.backendOverride = b
	c.mu.Unlock()
}

// ── 内容树 ──────────────────────────────────────────────────────────
//
// 树的形状：**组合根 Stack → [内容, 浮层根]**。
//
// 为什么浮层挂在同一棵树里而不是独立的第二棵树：
//   - 命中测试只走一遍，浮层天然"在前面"（同级后者在上），
//     不需要"先问浮层再问内容"的两段式判断，也就不会漏掉某一层；
//   - 合成层树也只有一棵，浮层的 z 与内容的 z 在同一个 Band 空间里排序；
//   - 悬浮/光标/捕获的路径都是"根→目标"，两棵树各算一次必然对不齐。

// Mount 挂载内容树根节点（nil 表示卸载）。
//
// 挂载后下一次出帧会整树重排。重复挂载会替换旧根，旧树的脏标记不再生效。
func (c *Container) Mount(root *scene.Node) {
	c.mu.Lock()
	old := c.content
	c.content = root
	c.layoutAll = true
	c.rebuildScene()
	if c.disp != nil && c.scene != nil {
		c.disp.SetRoot(c.scene)
	}
	c.mu.Unlock()
	if old != nil && old != root {
		old.ClearDirtyTree()
	}
	c.RequestFrame()
}

// Content 返回内容树根（未挂载时为 nil）。
func (c *Container) Content() *scene.Node {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.content
}

// Scene 返回组合根（内容 + 浮层）。命中测试与录制都从它开始。
func (c *Container) Scene() *scene.Node {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.ensureScene()
	return c.scene
}

// rebuildScene 按当前内容与浮层重建组合根的子节点（调用方持锁）。
func (c *Container) rebuildScene() {
	if c.scene == nil {
		return
	}
	c.scene.RemoveAll()
	if c.content != nil {
		c.scene.Add(c.content)
	}
	if c.ov != nil {
		c.scene.Add(c.ov.Root())
	}
}

// ensureScene 惰性装配场景/浮层/焦点/分发器（调用方持锁）。
//
// 需要客户区尺寸：Compositor 与根节点的尺寸都来自它，0 尺寸时无处安放。
// 返回是否就绪。
func (c *Container) ensureScene() bool {
	if c.scene != nil {
		return true
	}
	if c.host == nil {
		return false
	}
	sz := c.host.ClientSize()
	if sz.W <= 0 || sz.H <= 0 {
		return false
	}
	if c.th == nil {
		c.th = c.opts.Theme
		if c.th == nil {
			c.th = theme.Default()
		}
	}
	c.scene = scene.Stack().SetSize2(sz.W, sz.H)
	// 树变脏 → 申请下一帧。挂在组合根上，内容与浮层的任何变更都能叫醒帧泵。
	c.scene.SetInvalidateHook(c.requestFrame)
	c.focus = input.NewFocusManager(input.NewFocusNode(false))
	c.ov = overlay.New(overlay.Options{
		Viewport: c.viewport,
		Focus:    c.focus,
		Now:      func() time.Time { return time.Now() },
		// 浮层集合变化必然要重画：靠这次请求帧把新浮层推到屏幕上。
		OnChange: c.requestFrame,
	})
	c.disp = input.NewDispatcher(c.scene)
	c.disp.OnCursorChanged(func(cur scene.Cursor) {
		// 光标暂由宿主平台层落地（M3 接 Binder.SetCursor）。
		c.lastCursor = cur
	})
	c.rebuildScene()
	return true
}

// viewport 返回根空间可用区域。
func (c *Container) viewport() geom.Rect {
	if c.host == nil {
		return geom.ZeroRect
	}
	sz := c.host.ClientSize()
	return geom.RectXYWH(0, 0, sz.W, sz.H)
}

// Overlay 返回浮层宿主（首次调用时装配）。
//
// 未就绪（还没有尺寸）时返回 nil——调用方必须先 Resize/Show 再弹浮层，
// 否则浮层没有视口可定位。
func (c *Container) Overlay() *overlay.Host {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.ensureScene() {
		return nil
	}
	return c.ov
}

// Focus 返回焦点管理器（键盘导航与焦点域）。
func (c *Container) Focus() *input.FocusManager {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.ensureScene() {
		return nil
	}
	return c.focus
}

// Dispatcher 返回事件分发器（给节点挂处理器用）。
func (c *Container) Dispatcher() *input.Dispatcher {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.ensureScene() {
		return nil
	}
	return c.disp
}

// SetHandler 给节点挂事件处理器（等价于 Dispatcher().Set）。
func (c *Container) SetHandler(n *scene.Node, h input.Handler) {
	if d := c.Dispatcher(); d != nil {
		d.Set(n, h)
	}
}

// ── 主题 ────────────────────────────────────────────────────────────

// Theme 返回当前主题（未设置时为默认浅色）。
func (c *Container) Theme() *theme.Data {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.th == nil {
		if c.opts.Theme != nil {
			return c.opts.Theme
		}
		return theme.Default()
	}
	return c.th
}

// SetTheme 替换主题。
//
// "换主题"在渲染上等价于"整树重排 + 重录"：颜色不进布局，但字号/行高/
// 内边距进度量，所以不能只标绘制脏。整树重排的代价只在切换那一帧。
func (c *Container) SetTheme(t *theme.Data) {
	if t == nil {
		t = theme.Default()
	}
	c.mu.Lock()
	c.th = t
	c.opts.Theme = t
	c.layoutAll = true
	if c.scene != nil {
		c.scene.Relayout(c.rootConstraints())
	}
	c.mu.Unlock()
	c.RequestFrame()
}

// rootConstraints 返回根节点的布局约束（调用方持锁）。
func (c *Container) rootConstraints() scene.Constraints {
	sz := geom.Size{}
	if c.host != nil {
		sz = c.host.ClientSize()
	}
	return scene.TightConstraints(sz)
}

// ── 浮层便捷入口 ────────────────────────────────────────────────────

// ShowToast 弹一条自动消失的提示，返回句柄（未就绪时为 0）。
//
// 这是"引擎提供默认行为"的一个样本：位置、层带、关闭策略、停留时长
// 都不该由调用方每次手写——写十次必然有九种尺寸。
func (c *Container) ShowToast(msg string, d time.Duration) overlay.Handle {
	ov := c.Overlay()
	if ov == nil {
		return 0
	}
	if d <= 0 {
		d = 3 * time.Second
	}
	vp := c.viewport()
	const (
		w = float32(260)
		h = float32(40)
	)
	return ov.Show(overlay.Entry{
		Band:      theme.BandToast,
		Anchor:    overlay.PointAnchor(geom.Pt((vp.W()-w)/2, vp.Max.Y-h-48)),
		Placement: overlay.PlacementPoint,
		Size:      geom.Size{W: w, H: h},
		Dismiss:   overlay.DismissTimeout,
		Duration:  d,
		Tag:       "toast",
		Build: func() *scene.Node {
			th := c.Theme()
			n := scene.Box(scene.BoxSpec{
				Fill:   scene.Solid(th.Color.Surface),
				Radius: th.Radius.Base,
				Border: scene.Stroke(th.Color.Border, 1),
			})
			n.SetSize2(w, h)
			n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "toast"})
			tx := scene.Text(scene.TextSpec{
				Src:    msg,
				Color:  th.Color.Text,
				Layout: text.LayoutText(msg, text.DefaultFont(th.Font.SizeBase), text.Options{}),
			})
			// Box 的子节点靠 margin + 对齐放置（不是 Stack 的绝对定位）。
			tx.SetMargin(geom.Insets{Left: 12, Top: (h - tx.Size().H) / 2})
			n.Add(tx)
			return n
		},
	})
}

// ── 输入 ────────────────────────────────────────────────────────────

// PushEvent 实现 host.EventSink：接收平台层归一化后的事件。
func (c *Container) PushEvent(ev input.Event) bool {
	if ev == nil {
		return false
	}
	switch e := ev.(type) {
	case *input.PointerEvent:
		return c.HandlePointer(e)
	case *input.KeyEvent:
		// 按键先给浮层（Esc 关闭），未消费再走焦点链。
		if c.handleOverlayKey(e) {
			return true
		}
		if d := c.Dispatcher(); d != nil {
			return d.Key(e)
		}
	case *input.WheelEvent:
		if d := c.Dispatcher(); d != nil {
			return d.Wheel(e)
		}
	case *input.IMEEvent:
		if d := c.Dispatcher(); d != nil {
			return d.IME(e)
		}
	}
	return false
}

// HandlePointer 处理指针事件，返回是否被消费。
//
// 顺序固定：**浮层 → 内容树**。浮层命中就透传不下去（浮层是最上面的一层纸）；
// 浮层未命中也不代表事件白给——OutsideClick 会在这里把它关掉，那次点击
// 同样不该落到下面的按钮上。
func (c *Container) HandlePointer(ev *input.PointerEvent) bool {
	if ev == nil {
		return false
	}
	c.mu.Lock()
	ready := c.ensureScene()
	ov, d := c.ov, c.disp
	c.mu.Unlock()
	if !ready {
		return false
	}
	if ov != nil && !ov.HandlePointer(ev, ev.PosRoot) {
		return true
	}
	if d == nil {
		return false
	}
	return d.Pointer(ev)
}

// HandleKey 处理键盘事件（浮层优先）。
func (c *Container) HandleKey(ev *input.KeyEvent) bool {
	if ev == nil {
		return false
	}
	if c.handleOverlayKey(ev) {
		return true
	}
	if d := c.Dispatcher(); d != nil {
		return d.Key(ev)
	}
	return false
}

// handleOverlayKey 把按键交给浮层（仅 Esc 有关闭语义）。
func (c *Container) handleOverlayKey(ev *input.KeyEvent) bool {
	c.mu.Lock()
	ov := c.ov
	c.mu.Unlock()
	return ov != nil && ov.HandleKey(ev)
}

// TickTime 推进浮层计时与动画（由帧泵每帧调用，无窗口模式下由门禁手动调）。
func (c *Container) TickTime(now time.Time) {
	c.mu.Lock()
	ov := c.ov
	c.mu.Unlock()
	if ov != nil && ov.Tick(now) > 0 {
		c.RequestFrame()
	}
}

// Resize 修改客户区尺寸并重建帧缓冲。
func (c *Container) Resize(w, h int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.st == stateClosed || c.host == nil {
		return
	}
	c.host.Resize(w, h)
	if c.comp != nil {
		_ = c.comp.Resize(w, h)
	}
	c.host.Root().Invalidate(c.host.Root().LocalBounds())
	if c.scene != nil {
		// 根节点跟着视口走；整树重排因为换行/份额都会变。
		c.scene.SetSize2(float32(w), float32(h))
		c.layoutAll = true
	}
}

// RequestFrame 请求在下一帧处理（无 damage 的场景，如异步结果回填）。
//
// 除了记标志，还要**装填帧闹钟**：帧泵是"按需装填"的（C14），
// 没有这一句的话帧泵根本不会被叫醒——界面会停在上一帧，
// 直到别的原因（拖动窗口、鼠标划过）触发了帧。
func (c *Container) RequestFrame() {
	c.mu.Lock()
	if c.st == stateClosed {
		c.mu.Unlock()
		return
	}
	c.requested = true
	s := c.sched
	c.mu.Unlock()
	if s != nil {
		s.RequestFrame()
	}
}

// requestFrame 是"可能已经在帧内"的安全版本。
//
// 浮层宿主的 OnChange 既可能被用户代码调用（帧外），也可能被帧内的
// Relayout 触发。后者持着容器锁，再走 RequestFrame 就是自己等自己。
func (c *Container) requestFrame() {
	if c.mu.TryLock() {
		if c.st != stateClosed {
			c.requested = true
			if s := c.sched; s != nil {
				c.mu.Unlock()
				s.RequestFrame()
				return
			}
		}
		c.mu.Unlock()
		return
	}
	c.deferred.Store(true)
}

// ensureCompositor 装配渲染管线（调用方持锁）。
func (c *Container) ensureCompositor() error {
	if c.comp != nil {
		return nil
	}
	size := c.host.ClientSize()
	w, h := int(size.W), int(size.H)
	if w <= 0 || h <= 0 {
		return ErrNotReady
	}
	be := c.backendOverride
	if be == nil {
		be = software.New(software.Options{DisableAA: !c.opts.AA})
	}
	c.backend = be
	c.textPainter = gdi.NewPainter()
	comp, err := composite.New(composite.Options{
		Backend:    be,
		Text:       c.textPainter,
		Width:      w,
		Height:     h,
		Background: c.opts.Background,
	})
	if err != nil {
		return err
	}
	c.comp = comp
	c.rec = render.NewRecorder(render.NewDisplayList(128))
	return nil
}

// Backing 返回合成结果（32bpp 预乘 BGRA，只读）；未出帧时为 nil。
func (c *Container) Backing() *raster.Backing {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.comp == nil {
		return nil
	}
	return c.comp.Backing()
}

// LastDamage 返回最近一帧提交的脏区（未提交过时为零矩形）。
func (c *Container) LastDamage() geom.Rect {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.lastDamage
}

// FrameOnce 走一遍完整帧流程并返回统计。
//
// 顺序（与 03-ENGINE.md §9 一致）：
//
//	damage 汇总 → 清脏区 → 录制（仅脏帧）→ 光栅 → 文字回放 → 提交 → 清画布脏标记
//
// 静止帧（无 damage、无结构变更、无帧请求）直接返回，不录制也不光栅（C14）。
func (c *Container) FrameOnce() (frame.Stats, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.st == stateClosed || c.host == nil {
		return frame.Stats{}, ErrClosed
	}

	began := time.Now()
	now := c.clock.Now()
	var st frame.Stats
	st.Index = c.index
	st.Now = now
	if c.index > 0 {
		st.Delta = now - c.last
	}

	root := c.host.Root()
	if root != nil {
		st.Surfaces = root.Count()
	}

	if err := c.ensureCompositor(); err != nil {
		// 还没有尺寸（headless 未 Resize）：只做结构统计，与 M0 行为一致
		if root != nil {
			st.Damage = root.Damage()
			st.Dirty = root.HasDamage() || root.StructureDirty() || c.requested
			if st.Dirty && st.Damage.IsEmpty() {
				st.Damage = root.LocalBounds()
			}
		}
		st.FrameDuration = time.Since(began)
		c.index++
		c.last = now
		c.requested = false
		return st, nil
	}

	if root == nil {
		st.FrameDuration = time.Since(began)
		c.index++
		return st, nil
	}

	// 场景树：布局 → 脏区 → 录制。（M1 的 Painter 路径保留，可与场景树叠加）
	var (
		sceneDirty scene.Dirty
		needLayout bool
	)
	if c.ensureScene() {
		if c.layoutAll {
			c.scene.Relayout(c.rootConstraints())
			c.layoutAll = false
			needLayout = true
		}
		sceneDirty = c.scene.AnyDirty()
		if sceneDirty.Layout {
			c.scene.Layout(c.rootConstraints())
			needLayout = true
			sceneDirty = c.scene.AnyDirty()
			sceneDirty.Layout = false
		}
		if needLayout && c.ov != nil {
			c.ov.Relayout() // 布局变了，锚点跟着走
		}
	}

	// 结构变化（增删画布/改尺寸）必须整屏重画：布局可能整体位移。
	dirty := root.HasDamage() || root.StructureDirty() || c.requested || sceneDirty.Any()
	switch {
	case root.StructureDirty() || needLayout:
		c.comp.InvalidateAll()
	case !root.Damage().IsEmpty():
		c.comp.Invalidate(root.Damage())
	case sceneDirty.Any():
		if r := c.scene.DirtyRect(); !r.IsEmpty() {
			c.comp.Invalidate(r)
		} else {
			c.comp.InvalidateAll()
		}
	}

	var dl *render.DisplayList
	if dirty {
		c.rec.Reset()
		if c.scene != nil {
			c.scene.Record(c.rec)
		}
		if c.paint != nil {
			c.paint(c.rec)
		}
		dl = c.rec.Finish()
	}

	cs, err := c.comp.Draw(dl)
	if err != nil {
		return st, err
	}
	st.Presented = cs.Presented
	st.Dirty = cs.Presented
	st.RasterOps = cs.Raster.Ops
	st.Degraded = cs.Raster.Degraded
	st.TextLines = cs.TextLines
	st.Nodes = countNodes(c.scene)
	st.Damage = c.comp.LastPresented().Bounds()
	if cs.Presented {
		c.lastDamage = st.Damage
	}

	if cs.Presented {
		root.ClearDamageTree()
		root.ClearStructureDirty()
		if c.scene != nil {
			c.scene.ClearDirtyTree()
		}
	}

	st.FrameDuration = time.Since(began)
	c.index++
	c.last = now
	c.requested = false
	// 帧内发生的浮层变化：本帧已把它们录进画面，没提交成功才补一帧。
	if c.deferred.Swap(false) && !st.Presented {
		c.requested = true
	}
	return st, nil
}

// countNodes 统计场景节点数（诊断用）。
func countNodes(n *scene.Node) int {
	if n == nil {
		return 0
	}
	c := 1
	for _, ch := range n.Children() {
		c += countNodes(ch)
	}
	return c
}

// Frames 返回已走过的帧数。
func (c *Container) Frames() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.index
}

// Run 启动帧泵并进入消息循环，阻塞直到窗口关闭或 Close 被调用。
//
// 帧泵的形状：**按需装填**。没有脏帧、没有浮层计时、没有输入时不装闹钟，
// 于是静止界面一帧都不出（C14）——这是"CPU 占用 0%"与"每帧 0.1ms 但每秒
// 60 次"之间的区别，也是笔记本风扇转不转的区别。
//
// 无窗口容器（NewHeadless）没有消息循环可进，返回 ErrNilForm；
// 那种场景请用 FrameOnce + ManualTicker 自己驱动。
func (c *Container) Run() error {
	c.mu.Lock()
	if c.st == stateClosed {
		c.mu.Unlock()
		return ErrClosed
	}
	if c.form == nil {
		c.mu.Unlock()
		return ErrNilForm
	}
	if c.sched != nil {
		c.mu.Unlock()
		return nil // 已在跑
	}
	sched, err := frame.NewScheduler(frame.Options{
		Clock:  c.clock,
		Ticker: frame.NewTicker(),
		OnFrame: func(now, delta time.Duration) bool {
			c.TickTime(time.Now())
			st, ferr := c.FrameOnce()
			return ferr == nil && st.Presented
		},
		OnResize: c.fireResize,
	})
	if err != nil {
		c.mu.Unlock()
		return err
	}
	c.sched = sched
	c.st = stateRunning
	c.mu.Unlock()

	sched.RequestFrame()
	// 进入平台消息循环：帧闹钟在别的 goroutine 上触发，消息循环负责
	// 输入与窗口生命周期，两者不互相阻塞（否则拖动窗口时动画会停）。
	lcl.Application.Run()

	c.mu.Lock()
	s := c.sched
	c.sched = nil
	c.st = stateClosed
	c.mu.Unlock()
	if s != nil {
		s.Stop()
	}
	return nil
}

// fireResize 是去抖窗口到期后的真正 resize。
func (c *Container) fireResize(w, h int) {
	c.Resize(w, h)
	c.RequestFrame()
}

// Scheduler 返回帧泵（未运行时为 nil）。
func (c *Container) Scheduler() *frame.Scheduler {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.sched
}

// RunManual 启动一个由调用方手动推进的帧泵（无窗口门禁用）。
//
// 与 Run 的差别只有时钟与闹钟的来源：逻辑完全同一条路径，
// 否则门禁测的就不是真实运行的那套代码了。
func (c *Container) RunManual(clock *frame.ManualClock) *frame.ManualTicker {
	c.mu.Lock()
	if c.st == stateClosed {
		c.mu.Unlock()
		return nil
	}
	tk := frame.NewManualTicker(clock)
	c.clock = clock
	c.ticker = tk
	sched, err := frame.NewScheduler(frame.Options{
		Clock:        clock,
		Ticker:       tk,
		ResizeTicker: tk,
		OnFrame: func(now, delta time.Duration) bool {
			c.TickTime(time.Now())
			st, ferr := c.FrameOnce()
			return ferr == nil && st.Presented
		},
		OnResize: c.fireResize,
	})
	if err != nil {
		c.mu.Unlock()
		return nil
	}
	c.sched = sched
	c.st = stateRunning
	c.mu.Unlock()
	sched.RequestFrame()
	return tk
}

// Close 关闭容器并释放整条管线。
//
// 释放顺序遵守 C21：先解绑视图（Surface），再释放位图（Backing/DIB），
// 最后归还池（M2 引入池后在此回收）。
func (c *Container) Close() {
	c.mu.Lock()
	if c.st == stateClosed {
		c.mu.Unlock()
		return
	}
	c.st = stateClosed
	h := c.host
	tp := c.textPainter
	s := c.sched
	tk := c.ticker
	c.host = nil
	c.comp = nil
	c.rec = nil
	c.sched = nil
	c.ticker = nil
	c.ov = nil
	c.scene = nil
	c.content = nil
	c.disp = nil
	c.focus = nil
	c.mu.Unlock()

	if s != nil {
		s.Stop()
	}
	if tk != nil {
		tk.Stop()
	}
	if tp != nil {
		if cl, ok := tp.(interface{ Close() }); ok {
			cl.Close()
		}
	}
	if h == nil {
		return
	}
	if root := h.Root(); root != nil {
		// 视图解绑：从叶子往上摘；Detach 会改父的 children，所以先取快照。
		var detach func(*host.Surface)
		detach = func(s *host.Surface) {
			kids := make([]*host.Surface, len(s.Children()))
			copy(kids, s.Children())
			for _, child := range kids {
				detach(child)
			}
			s.Detach()
		}
		detach(root)
	}
}
