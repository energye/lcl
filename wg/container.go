package wg

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
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
	"github.com/energye/lcl/wg/ui"
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
	// widget 是声明式内容源；widgetNode 是它最近一次 Build 出的节点。
	// 只支持一个根 widget（表单页就是一棵树），与 Mount(*scene.Node) 二选一。
	widget     ui.Widget
	widgetNode *scene.Node
	ov         *overlay.Host
	focus      *input.FocusManager
	disp       *input.Dispatcher
	ime        *host.IMEBridge // 平台输入法桥；当前编辑模型由聚焦的输入框切换
	th         *theme.Data
	// inherited 是沿组件树向下传的配置（BuildContext.Inherited 的数据源）。
	inherited map[any]any
	// layoutAll 表示下一帧必须整树重排（挂载/换主题/resize）。
	layoutAll bool

	// layerTree 是当前帧的层树（缓存命中的层跨帧复用）。
	layerTree *composite.Tree
	// presenter 是呈现链路的最后一公里：把合成出的 Backing 提交到窗口 DC。
	// headless 容器为 nil（合成结果只在内存里，供出图/自检读取）。
	presenter *host.Presenter
	// binder 是平台消息接管器；Run 里安装，WM_PAINT 时驱动 presenter.Paint。
	binder *host.Binder

	// 帧泵（Run 启动；无窗口模式下由调用方手动 FrameOnce/Tick）
	sched *frame.Scheduler // schedAtomic 是给平台回调（WM_SIZE/WM_PAINT）用的无锁快照。
	//
	// 为什么需要：这些回调在消息线程上，不能为了读一个指针去抢 c.mu——
	// 整帧合成会持有 c.mu 20–30ms，实测把 WM_SIZE 处理拖成 avg 25ms /
	// max 108ms，窗口边框跟不上鼠标（一卡一卡）。平台回调只做
	// "请求一次"（s.RequestResize / RequestFrame），调度器自身线程安全。
	schedAtomic atomic.Pointer[frame.Scheduler]
	ticker      *frame.ManualTicker

	// lastCursor 是最近一次命中的光标（平台层据此换鼠标指针）。
	lastCursor scene.Cursor

	index     int
	last      time.Duration
	requested atomic.Bool
	// lastDamage 是最近一帧提交的脏区（诊断与门禁用）。
	lastDamage geom.Rect
	// deferred 记录"帧内有人请求了下一帧"。帧内回调已经持锁，
	// 直接走 RequestFrame 会自己和自己抢锁（死锁），因此先存标志、帧末合并。
	deferred atomic.Bool
	// animating 是"本窗有活跃动画"标志（Container.SetAnimating）：置位后
	// 帧泵按 60Hz 节拍标脏场景并续帧（动态组件驱动渲染）；未置位的窗
	// 严格按需渲染——有变化才出帧，静止零帧（C14）。
	animating atomic.Bool
	// viewportDep 记录"上次 Build 是否读取了视口"（ViewportSize）。
	// resize 只对依赖视口的页面重建；固定布局页面只重排布局，免全树
	// Build（成熟框架语义：只为声明依赖的组件重建）。
	viewportDep bool
	// renderHWND 是主线程渲染消息的目标窗口（0 = 无窗口，帧线程内联合成）。
	renderHWND uintptr
	// renderReq 是当前在途的渲染请求（单 outstanding，与帧槽一致）。
	renderReq atomic.Pointer[renderRequest]
	// renderCostUs 是最近一次主线程整帧（布局+合成+呈现）的耗时。
	// 超过一帧预算时，投递下一帧前留 inputBreath 空档给输入
	// （showcase §9：posted 渲染消息若永不停歇会饿死硬件输入）。
	renderCostUs atomic.Int64
	// imeAnchor 是输入控件上报的光标锚点（节点 + 局部矩形），
	// 布局后换算成客户区坐标交给平台输入法（组合窗/候选窗定位）。
	// imeAnchorOwner/caretOwner 是登记者身份（见 BuildContext 的同名钩子）。
	imeAnchor      *scene.Node
	imeAnchorLocal geom.Rect
	imeAnchorOwner any
	imeCaret       geom.Rect
	imeCaretOK     bool

	// 插入符闪烁：caretNode 非 nil 时按系统节奏切换可见性。
	caretNode  *scene.Node
	caretOwner any
	caretOn    bool
	caretNext  time.Duration
	caretArmed bool
	st         state
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
		ime:   host.NewIMEBridge(nil),
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
	c.widget, c.widgetNode = nil, nil // 直接式树替换声明式树
	c.swapContent(root)
	c.mu.Unlock()
	c.RequestFrame()
}

// MountWidget 挂载一棵声明式组件树。
//
// 与 Mount(*scene.Node) 的关系：两者共用同一个内容槽，后挂的替换先挂的。
// 直接式节点（图表/画布）仍可走 Mount；表单页这类有状态的界面走这里。
//
// SetState 之后组件由帧循环自动重建——不需要调用方手动 Mount 第二次。
// 重建发生在帧内、布局之前，因此新节点的尺寸当帧就能生效。
func (c *Container) MountWidget(w ui.Widget) {
	if w == nil {
		c.mu.Lock()
		c.widget, c.widgetNode = nil, nil
		c.mu.Unlock()
		c.Mount(nil)
		return
	}
	c.mu.Lock()
	c.ensureScene()
	c.viewportDep = false
	ctx := c.buildContext()
	node := w.Build(ctx)
	if r, ok := w.(interface{ Rebuild() }); ok {
		r.Rebuild() // 首次 Build 也要记账，否则诊断计数比实际少一次
	}
	if a, ok := w.(interface{ Attach(func()) }); ok {
		a.Attach(c.requestFrame)
	}
	c.widget, c.widgetNode = w, node
	c.swapContent(node)
	c.mu.Unlock()
	c.RequestFrame()
}

// SetInherited 设置沿组件树向下传递的值（主题之外的配置）。
// key 用私有类型的值，避免不同组件包之间撞键。
func (c *Container) SetInherited(key, value any) {
	c.mu.Lock()
	if c.inherited == nil {
		c.inherited = make(map[any]any)
	}
	c.inherited[key] = value
	c.mu.Unlock()
}

// Widget 返回当前挂载的组件（未挂载时为 nil）。
func (c *Container) Widget() ui.Widget {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.widget
}

// buildContext 组装 Build 环境（调用方持锁）。
func (c *Container) buildContext() *ui.BuildContext {
	scale := c.opts.Scale
	if scale <= 0 {
		scale = 1
	}
	th := c.th
	if th == nil {
		th = c.opts.Theme
	}
	if th == nil {
		th = theme.Default()
	}
	inherited := c.inherited
	return &ui.BuildContext{
		Theme: th,
		Media: ui.MediaQuery{
			DPI:   geom.Size{W: 96 * scale, H: 96 * scale},
			Scale: scale, TextScale: 1,
		},
		Focus: c.focus, Input: c.disp, Overlay: c.ov, Scheduler: c.sched,
		SetTextEditor: func(e *input.TextEditor) { c.ime.SetEditor(e) },
		Inherited: func(key any) any {
			if inherited == nil {
				return nil
			}
			return inherited[key]
		},
		// 深层组件的状态变化 → 根 widget 置脏 + 请求帧。		// rebuildDirtyWidget 只扫根 widget 的脏位；深层组件直调 MarkDirty
		// 不请求帧的话，静止窗口（C14）永不重绘——真窗表现为"点了没反应"。
		// 置根脏必须用 SetDirtyOnly：根的 invalidate 往往就是本回调，
		// 走 MarkDirty 会形成 MarkDirty→InvalidateRoot→MarkDirty 无限递归。
		// 根 widget 必须**实时读**：buildContext 在首次 Build 前被调用，
		// 那时 c.widget 还没赋值——闭包捕获会拿到 nil，深层组件的置脏
		// 全部 no-op（动画窗只出两帧的根因）。
		RequestFrame: c.requestFrame,
		InvalidateRoot: func() {
			if c.mu.TryLock() {
				w := c.widget
				c.mu.Unlock()
				if w != nil {
					if r, ok := w.(interface{ SetDirtyOnly() }); ok {
						r.SetDirtyOnly()
					}
				}
				c.requestFrame()
				return
			}
			// 帧内（本 goroutine 持锁，TryLock 必败）：直接置根脏，
			// 并经 deferred 让帧尾补一帧（见 FrameOnce 尾部的消费）。
			if w := c.widget; w != nil {
				if r, ok := w.(interface{ SetDirtyOnly() }); ok {
					r.SetDirtyOnly()
				}
			}
			c.deferred.Store(true)
		},
		Viewport: func() geom.Size {
			if c.host == nil {
				return geom.Size{}
			}
			return c.host.ClientSize()
		}(),
		// 视口依赖登记：本次 Build 有组件调用 ViewportSize() 即置位。
		ViewportUse: func() { c.viewportDep = true },
		// IME 光标锚点：输入控件在 Build 里上报（局部坐标，布局后换算）。
		// 按 owner 归属：未聚焦输入框的 nil 上报不得清掉已聚焦控件的锚点。
		ReportIMECaret: func(owner any, n *scene.Node, local geom.Rect) {
			if n == nil {
				if owner == nil || c.imeAnchorOwner == owner {
					c.imeAnchor, c.imeAnchorLocal, c.imeAnchorOwner = nil, geom.Rect{}, nil
				}
				return
			}
			c.imeAnchor, c.imeAnchorLocal, c.imeAnchorOwner = n, local, owner
		},
		// 插入符闪烁：输入控件上报光标节点，宿主负责 530ms 明暗切换。
		ReportCaret: c.setCaretNode,
	}
}

// rebuildDirtyWidget 在帧首重建脏组件（调用方持锁）。
//
// 直接在这里换 c.content 而不调 Mount：Mount 会抢 c.mu，而本方法运行在
// FrameOnce 的锁内。两条路径的"换树后要做什么"因此抽到 swapContent 里共用。
func (c *Container) rebuildDirtyWidget() bool {
	if c.widget == nil {
		return false
	}
	d, ok := c.widget.(interface{ Dirty() bool })
	if !ok || !d.Dirty() {
		return false
	}
	// 每次 Build 重新计算视口依赖（组件可能条件性读取）。
	c.viewportDep = false
	ctx := c.buildContext()
	node := c.widget.Build(ctx)
	if r, ok := c.widget.(interface{ Rebuild() }); ok {
		r.Rebuild()
	}
	c.widgetNode = node
	c.swapContent(node)
	return true
}

// swapContent 替换内容槽并同步组合根/分发器（调用方持锁）。
func (c *Container) swapContent(node *scene.Node) {
	old := c.content
	c.content = node
	c.layoutAll = true
	if c.scene != nil {
		c.rebuildScene()
		if c.disp != nil {
			c.disp.SetRoot(c.scene)
		}
	}
	if old != nil && old != node {
		old.ClearDirtyTree()
		// 旧树的 handler 登记必须清掉：重建每次产出全新节点，不清就是
		// 无界增长（悬停切换、逐键输入都是重建）。分发器会保留活跃
		// 捕获与悬停链引用的节点——手势要在"按下那一刻的路径"上完成。
		if c.disp != nil {
			c.disp.ForgetTree(old)
		}
	}
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

// Walk 在容器锁内遍历组合根（内容 + 浮层）。
//
// 供工作线程（审计等）安全读取场景树：Scene() 返回的指针在锁外遍历会
// 与主线程布局（帧内写 size/offset）并发读写同一 Node，实测偶发
// "unexpected signal during runtime execution"（--audit 全量第 36 窗）。
// 回调内只读节点属性，**不得调用会重新获取容器锁的方法**（Scene/Snapshot 等）。
func (c *Container) Walk(fn func(*scene.Node) bool) {
	if fn == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.ensureScene() || c.scene == nil {
		return
	}
	c.scene.Walk(fn)
}

// rebuildScene 按当前内容与浮层重建组合根的子节点（调用方持锁）。
func (c *Container) rebuildScene() {
	if c.scene == nil {
		return
	}
	c.scene.RemoveAll()
	if c.content != nil {
		// 内容必须铺满视口：Stack 对普通子节点用 LooseConstraints（按内容
		// 自然尺寸），会让只含窄导航+文字的页面缩成一小条、右侧内容宽 0。
		// SetStackPosition(0,0) 让内容成为 positioned 且 keepSize=false，
		// 走 TightConstraints(视口)——这正是"父容器铺满"的语义（浮层遮罩同理）。
		c.content.SetStackPosition(geom.Pt(0, 0)).SetPositionKeepSize(false)
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
	// 焦点管理器必须挂到分发器：否则 SetCurrent 只改内部状态，
	// Key/IME 仍投给旧节点（表现为"输入框有焦点框但打不出字"）。
	c.focus.Attach(c.disp)
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

// IME 返回输入法桥。聚焦的输入框通过 BuildContext.SetTextEditor 切换它，
// 平台消息经宿主路由进入这里；headless 下可以直接 Handle 注入消息做门禁。
func (c *Container) IME() *host.IMEBridge {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ime
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
	// 字号/行高/内边距都是度量，换主题必须让组件重新 Build。
	if m, ok := c.widget.(interface{ MarkDirty() }); ok {
		m.MarkDirty()
	}
	if c.scene != nil {
		c.scene.Relayout(c.rootConstraints())
	}
	// resize 过渡帧的填边色跟随主题（否则深色主题拖动时会闪浅色条带）。
	if c.presenter != nil {
		c.presenter.SetResizePresent(c.opts.ResizePresent == ResizePresentStretch, c.resizeBackground())
	}
	c.mu.Unlock()
	c.RequestFrame()
}

// resizeBackground 返回 resize 过渡帧填边用的背景色。
//
// 优先 Options.Background；零值跟随当前主题背景（深色主题下不会闪白条），
// 与根节点铺底颜一致是"先露底色、下一帧覆盖"观感正确的前提。
func (c *Container) resizeBackground() geom.Color {
	bg := c.opts.Background
	if bg == (geom.Color{}) {
		if c.opts.Theme != nil {
			bg = c.opts.Theme.Color.Background
		} else {
			bg = theme.Default().Color.Background
		}
	}
	return bg
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
	ov, d, f, root := c.ov, c.disp, c.focus, c.scene
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

	// 点击非焦点子树 → 失焦（标准行为）：先记按下前的焦点与命中目标，
	// 分发后若焦点没被控件主动迁移、且目标不在焦点子树内，就清焦点。
	var before *input.FocusNode
	var hit *scene.Node
	if ev.Down() && f != nil {
		before = f.Current()
		if before != nil && before.Node != nil && root != nil {
			hit = input.HitTest(root, ev.PosRoot).Target()
		}
	}
	handled := d.Pointer(ev)
	if before != nil && f.Current() == before && !nodeWithin(hit, before.Node) {
		f.Blur()
	}
	return handled
}

// updateIMECaretLocked 把输入控件上报的光标锚点换算成客户区坐标（持锁）。
//
// 只认仍在当前场景树里的节点：重建换树后旧节点若没被新 Build 重新上报，
// 就清空锚点（候选窗不动，但不会再定位到幽灵光标）。
func (c *Container) updateIMECaretLocked() {
	n := c.imeAnchor
	if n == nil || c.scene == nil {
		c.imeCaretOK = false
		return
	}
	var chain []*scene.Node
	for x := n; x != nil; x = x.Parent() {
		chain = append(chain, x)
	}
	if chain[len(chain)-1] != c.scene {
		c.imeCaretOK = false
		return
	}
	m := geom.Identity
	for i := len(chain) - 1; i >= 0; i-- {
		m = chain[i].FrameMatrix(m)
	}
	p := m.MapPoint(c.imeAnchorLocal.Min)
	c.imeCaret = geom.RectXYWH(p.X, p.Y, c.imeAnchorLocal.W(), c.imeAnchorLocal.H())
	c.imeCaretOK = true
}

// syncIMECaret 把最新光标矩形交给平台输入法（仅主线程调用）。
func (c *Container) syncIMECaret() {
	c.mu.Lock()
	ok, r, b := c.imeCaretOK, c.imeCaret, c.binder
	c.mu.Unlock()
	if ok && b != nil {
		b.SetIMECaret(r)
	}
}

// caretBlinkPeriod 是插入符明暗切换间隔（Windows 默认 530ms）。
const caretBlinkPeriod = 530 * time.Millisecond

// setCaretNode 记录/重置插入符节点（Build 与编辑时调用；node=nil 且
// owner 是当前登记者时停止闪烁）。
//
// owner 归属是必需的：一个页面上有多个输入框，未聚焦者也会上报 nil；
// 不区分 owner 时，后构建的未聚焦输入框会把已聚焦输入框的插入符
// 从闪烁队列里清掉（真窗表现：刚点进输入框/重新聚焦时光标不闪）。
//
// 与 IME 锚点同一线程约定：Build 在帧内（持锁），编辑刷新在消息线程；
// P0 下二者都是窗口主线程，不发生并发写。
func (c *Container) setCaretNode(owner any, n *scene.Node) {
	if n == nil {
		if owner != nil && c.caretOwner != owner {
			return
		}
		c.caretNode, c.caretOwner, c.caretArmed = nil, nil, false
		c.caretOn = true
		return
	}
	c.caretOwner = owner
	if c.caretNode != n {
		c.caretNode = n
		c.caretOn = true
		n.SetVisible(true)
	} else if !c.caretOn {
		// 同一个节点上的编辑/移动：立即恢复可见并重置相位（标准光标行为）。
		c.caretOn = true
		n.SetVisible(true)
	}
	c.caretNext = c.clock.Now() + caretBlinkPeriod
	c.caretArmed = true
}

// updateCaretBlinkLocked 在帧首推进闪烁相位（调用方持锁）。
func (c *Container) updateCaretBlinkLocked(now time.Duration) {
	if c.caretNode == nil {
		return
	}
	// 换树后旧光标节点可能已不在场景里（新内容没有输入框/组件没上报 nil）。
	// 锚点用"是否仍在当前场景"兜底，避免给孤儿节点持续排闪烁帧。
	if !c.caretNodeAlive(c.caretNode) {
		c.caretNode, c.caretOwner, c.caretArmed = nil, nil, false
		c.caretOn = true
		return
	}
	if !c.caretArmed {
		c.caretNext = now + caretBlinkPeriod
		c.caretArmed = true
		return
	}
	if now < c.caretNext {
		return
	}
	c.caretOn = !c.caretOn
	c.caretNode.SetVisible(c.caretOn)
	c.caretNext = now + caretBlinkPeriod
}

// caretNodeAlive 判断节点是否仍挂在当前场景树上。
func (c *Container) caretNodeAlive(n *scene.Node) bool {
	if n == nil || c.scene == nil {
		return false
	}
	for x := n; x != nil; x = x.Parent() {
		if x == c.scene {
			return true
		}
	}
	return false
}

// scheduleCaretBlink 申请下一次插入符切换的帧（帧末调用）。
func (c *Container) scheduleCaretBlink() {
	if c.caretNode == nil {
		return
	}
	if c.scene != nil && !c.caretNodeAlive(c.caretNode) {
		return
	}
	s := c.schedAtomic.Load()
	if s == nil {
		return
	}
	d := c.caretNext - c.clock.Now()
	if !c.caretArmed || d < 0 {
		d = 0
	}
	if d > caretBlinkPeriod {
		d = caretBlinkPeriod
	}
	s.RequestFrameIn(d)
}

// BlurFocus 实现 host.FocusSink：窗口失活时清焦点。
func (c *Container) BlurFocus() {
	c.mu.Lock()
	f := c.focus
	c.mu.Unlock()
	if f != nil {
		f.Blur()
	}
}

var _ host.FocusSink = (*Container)(nil)

// nodeWithin 判断 n 是否在 root 子树内（含自身）。
func nodeWithin(n, root *scene.Node) bool {
	for x := n; x != nil; x = x.Parent() {
		if x == root {
			return true
		}
	}
	return false
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

// ResizeRequest 实现 host.ResizeSink：把平台尺寸变化转给节流调度器。
//
// 不能直接调 Resize：拖动窗口时会来几十条 WM_SIZE，每次都重建合成缓冲
// 既慢又会让画面抖动。交给 scheduler 节流+尾随：拖动中每帧（~16.7ms）
// 跟随一次，停下后尾随保证最后尺寸落地（纯尾随在连续 WM_SIZE 下永不触发）。
//
// 无锁：WM_SIZE 在消息线程上，只读 schedAtomic 并登记请求；真正应用尺寸
// 在调度器的定时器 goroutine 上执行，消息线程绝不等整帧合成。
func (c *Container) ResizeRequest(w, h int) {
	if s := c.schedAtomic.Load(); s != nil {
		s.RequestResize(w, h)
		return
	}
	// 还没有调度器（启动早期）：直接同步一次，保证首帧尺寸正确。
	c.Resize(w, h)
	c.RequestFrame()
}

// DPIChanged 实现 host.DPISink：DPI 变更在容器锁内应用。
//
// 平台回调（WM_DPICHANGED）在消息线程上只传事件；场景重排、组件重建与
// 根画布失效都在 c.mu 下进行，与 resize goroutine、帧线程串行——消息
// 线程直接改 Surface 会与整帧合成竞争（-race 实测）。文字度量/布局可能
// 依赖新缩放，沿用 SetTheme 的保守整树重建。
func (c *Container) DPIChanged(scale float32) {
	c.mu.Lock()
	if c.st == stateClosed || c.host == nil {
		c.mu.Unlock()
		return
	}
	c.host.SetScale(scale)
	c.layoutAll = true
	if m, ok := c.widget.(interface{ MarkDirty() }); ok {
		m.MarkDirty()
	}
	if c.scene != nil {
		c.scene.Relayout(c.rootConstraints())
	}
	c.host.Root().Invalidate(c.host.Root().LocalBounds())
	c.requested.Store(true)
	c.mu.Unlock()
	c.RequestFrame()
}

var _ host.DPISink = (*Container)(nil)

// Resize 修改客户区尺寸并重建帧缓冲。
//
// 这里**不做旧帧缩放**：resize 节流窗口是一帧，每次尺寸落地后
// fireResize 都会立刻申请一帧整屏重绘（Resize 已 Invalidate 全根 +
// layoutAll + widget MarkDirty），过渡帧由 DWM 拉伸上一张已提交帧兜底。
// 旧版逐像素最近邻铺满（4.5M 像素 ×2 次整数除法）实测 2880×1557 一次
// 要 46–55ms，全部花在消息线程上——最大化卡顿的主因，且那些像素从未
// 被提交（Presenter 只提交与客户区同尺寸的帧），已删。
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
	// 组件可能按 MediaQuery 分支（窄屏换布局），尺寸变了就得重建。
	//
	// 但只重建**声明了视口依赖**的页面（上次 Build 调过 ViewportSize）：
	// 固定布局页面 resize 时只做布局重排，免全树 Build（成熟框架语义）。
	if m, ok := c.widget.(interface{ MarkDirty() }); ok && (c.viewportDep || c.widgetNode == nil) {
		m.MarkDirty()
	}
}

// RequestFrame 请求在下一帧处理（无 damage 的场景，如异步结果回填）。
//
// 除了记标志，还要**装填帧闹钟**：帧泵是"按需装填"的（C14），
// 没有这一句的话帧泵根本不会被叫醒——界面会停在上一帧，
// 直到别的原因（拖动窗口、鼠标划过）触发了帧。
//
// 无锁：输入/平台回调在消息线程调用它，不能为了读 sched 指针去抢 c.mu
// （整帧合成持有 c.mu 20–30ms）。
func (c *Container) RequestFrame() {
	c.requested.Store(true)
	if s := c.schedAtomic.Load(); s != nil {
		s.RequestFrame()
	}
}

// requestFrame 是"可能已经在帧内"的安全版本。
//
// 浮层宿主的 OnChange 既可能被用户代码调用（帧外），也可能被帧内的
// Relayout 触发。帧内不能直接装帧（本帧末会把 requested 清掉），
// 用 TryLock 区分：帧外直接装下一拍，帧内记 deferred 由帧末合并。
func (c *Container) requestFrame() {
	c.requested.Store(true)
	if c.mu.TryLock() {
		c.mu.Unlock()
		if s := c.schedAtomic.Load(); s != nil {
			c.requestNextTick(s)
		}
		return
	}
	c.deferred.Store(true)
}

// requestNextTick 申请"下一个 60Hz 边界"的帧（动画/自驱动续帧专用）。
//
// 为什么不是固定的"当前时刻 +16ms"：相对延迟会把每帧的合成耗时累加进
// 周期（帧尾 +16ms → 实测 29ms/帧 ≈ 34fps）。钉在墙钟的绝对 16.667ms
// 网格上，周期恒为 60Hz；超过一拍的慢帧自然跳到下一个边界（不补帧，
// 与 showcase/Flutter 的 vsync 语义一致）。
// 手动时钟保持固定 16ms：确定性门禁按 16ms 片推进，不引入网格相位。
func (c *Container) requestNextTick(s *frame.Scheduler) {
	s.RequestFrameIn(nextTickDelay(c.clock.Kind(), c.clock.Now(), frame.DefaultStep))
}

// nextTickDelay 计算"到下一个 60Hz 边界"的延迟。
//
// 墙钟：钉在绝对网格（delay ∈ (0, step]）；手动时钟保持固定 16ms 片长，
// 确定性门禁不引入网格相位（详见 requestNextTick 的说明）。
func nextTickDelay(kind frame.ClockKind, now, step time.Duration) time.Duration {
	if kind == frame.Manual {
		return 16 * time.Millisecond
	}
	delay := step - now%step
	if delay <= 0 {
		delay = step
	}
	return delay
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

// Snapshot 在容器锁内拷贝当前合成缓冲，返回独立的 32bpp BGRA 预乘位图。
//
// 与 Backing 的区别：Backing 返回的是主线程复用的活缓冲，跨线程在锁外读取
// 会与下一帧光栅写竞争。审计等需要在工作线程采样像素的场景必须用本方法，
// 拷贝在锁内完成，返回的缓冲与合成器再无关联。未出帧时返回 nil。
func (c *Container) Snapshot() *raster.Backing {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.comp == nil {
		return nil
	}
	b := c.comp.Backing()
	if b == nil || b.W() <= 0 || b.H() <= 0 {
		return nil
	}
	out := raster.NewBacking(b.W(), b.H())
	if out == nil {
		return nil
	}
	copy(out.Pix(), b.Pix())
	return out
}

// ReleasePresent 在主线程释放呈现层的 LCL 资源（位图/像素视图）。
//
// 必须在 Application.Run 返回前调用：Run 之后 LCL 运行时开始销毁，
// 再 Free 位图会访问违例。Application.Terminate 不会触发 OnClose，
// 所以退出路径要显式调这个。可重复调用（内部置空）。
func (c *Container) ReleasePresent() {
	c.mu.Lock()
	p := c.presenter
	c.mu.Unlock()
	if p != nil {
		p.Close()
	}
}

// PresentNow 立即把最近一帧提交到窗口 DC（必须在主线程调用）。
//
// 正常路径是 WM_PAINT → Presenter.Paint；但刚显示窗口时系统可能还没投递
// paint，自检/截图如果只看 WM_PAINT 会得到"painted=0"的假阴性。这个入口
// 让调用方能同步确认像素真的落地，不依赖消息时机。
func (c *Container) PresentNow() bool {
	c.mu.Lock()
	p := c.presenter
	c.mu.Unlock()
	if p == nil {
		return false
	}
	return p.Paint()
}

// PresentStats 返回呈现链路统计（帧提交到窗口 DC 的次数）。
// headless 容器没有 Presenter，返回零值。
func (c *Container) PresentStats() host.PresentStats {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.presenter == nil {
		return host.PresentStats{}
	}
	return c.presenter.Stats()
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

	// 浮层计时/动画标脏在容器锁内推进：帧线程与 resize goroutine（异步
	// 应用尺寸）会并发改场景树，锁外 MarkNeedsPaint 与整树布局竞争
	// （-race 实测）。
	if c.ov != nil && c.ov.Tick(time.Now()) > 0 {
		c.requested.Store(true)
	}
	if c.animating.Load() && c.scene != nil {
		c.scene.MarkNeedsPaint()
	}
	// 插入符闪烁：到点切换可见性并让脏区自然收敛到光标矩形。
	c.updateCaretBlinkLocked(now)

	// 组件重建要在布局之前：SetState 改了内容，当帧就得用新尺寸排版，
	// 否则界面会晚一帧才跟上状态（肉眼见"点一下慢半拍"）。
	if c.rebuildDirtyWidget() {
		c.requested.Store(true)
	}

	root := c.host.Root()
	if root != nil {
		st.Surfaces = root.Count()
	}

	if err := c.ensureCompositor(); err != nil {
		// 还没有尺寸（headless 未 Resize）：只做结构统计，与 M0 行为一致
		if root != nil {
			st.Damage = root.Damage()
			st.Dirty = root.HasDamage() || root.StructureDirty() || c.requested.Load()
			if st.Dirty && st.Damage.IsEmpty() {
				st.Damage = root.LocalBounds()
			}
		}
		st.FrameDuration = time.Since(began)
		c.index++
		c.last = now
		c.requested.Store(false)
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
	// IME 光标锚点：布局后把控件局部矩形换算成客户区坐标。
	c.updateIMECaretLocked()

	// 结构变化（增删画布/改尺寸）必须整屏重画：布局可能整体位移。
	dirty := root.HasDamage() || root.StructureDirty() || c.requested.Load() || sceneDirty.Any()
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

	// 两条绘制路径：
	//   - 场景树 → Build 层树 → DrawTree：组透明/路径裁剪/滤镜/层缓存都在这条；
	//   - 外部 Painter（c.paint）→ 扁平指令表 → Draw：兼容逃生舱，不带层。
	// 早期这里只走扁平路径，导致层能力（阴影模糊、组透明度、蒙版裁剪）
	// 在真窗里从未生效——单测与真窗表现不一致的根源。
	var (
		cs  composite.Stats
		err error
	)
	if dirty {
		if c.paint == nil && c.scene != nil {
			tree := composite.Build(c.scene, composite.BuildOptions{Cache: true})
			c.layerTree = tree
			cs, err = c.comp.DrawTree(tree)
		} else {
			c.rec.Reset()
			if c.scene != nil {
				c.scene.Record(c.rec)
			}
			if c.paint != nil {
				c.paint(c.rec)
			}
			cs, err = c.comp.Draw(c.rec.Finish())
		}
	} else {
		cs, err = c.comp.Draw(nil)
	}
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
		// 最后一公里：把本帧合成结果交给 Presenter，驱动 WM_PAINT 提交到窗口。
		// 必须在清脏之前取 Damage（st.Damage 已保存），且只在真窗下存在。
		if c.presenter != nil {
			c.presenter.Frame(c.comp.Backing(), st.Damage)
		}
	}

	st.FrameDuration = time.Since(began)
	c.index++
	c.last = now
	c.requested.Store(false)
	// 帧内的置脏/请求（组件 Build 尾的 MarkDirty、浮层 Relayout 回调）
	// 必须落到下一帧：无论本帧是否提交成功都补一帧——动画页靠这条
	// "每帧重建自己"续命；普通组件帧内 MarkDirty 也同样需要下帧生效。
	// （旧语义"没提交成功才补"会把动画页的续帧吞掉，动画停在第二帧。）
	if c.deferred.Swap(false) {
		c.requested.Store(true)
		if s := c.sched; s != nil {
			// 续帧钉在 60Hz 绝对网格（见 requestNextTick）；相对 16ms 会把
			// 合成耗时累加进周期，且帧内自驱动的动画页没有这一拍间隔就会
			// 全速空转（deferred 链实测跑出 400fps 拉满 CPU）。
			c.requestNextTick(s)
		}
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

// Start 启动本窗的帧泵/消息接管/呈现器，但**不进入平台消息循环**。
//
// 多窗口入口：平台消息循环全局只有一条（主窗 Run 里的 Application.Run），
// 每个窗的帧泵/事件分发/呈现全部按容器隔离——Start 之后各窗即可独立
// 接收输入、独立出帧（静止零帧 C14，动画窗 60Hz 单 outstanding）。
// 主窗调 Run()（= Start + Application.Run）；其余窗调 Start()。
//
// 窗口关闭（OnClose）会停本窗帧泵并释放呈现层；主窗 Run 返回即程序退出。
func (c *Container) Start() error {
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
			// P0：组合在**窗口主线程**执行（与输入同线程），帧时钟只投递
			// 并等待完成。headless/无 HWND 时退化为帧线程内联。
			if c.renderHWND != 0 {
				return c.renderOnMain()
			}
			return c.frameOnceAndContinue()
		},
		OnResize: c.fireResize,
	})
	if err != nil {
		c.mu.Unlock()
		return err
	}
	c.sched = sched
	c.schedAtomic.Store(sched)
	c.st = stateRunning
	renderHWND := uintptr(0)
	if c.form != nil {
		renderHWND = uintptr(c.form.Handle())
	}
	c.renderHWND = renderHWND
	c.mu.Unlock()

	// 平台消息必须在消息循环启动前接管。此前 Run 只跑了帧泵，
	// 窗口消息仍走 LCL 默认处理：鼠标键盘到不了分发器，输入法更是完全失联。
	binder := host.NewBinder(c.host, c)
	binder.SetIMEBridge(c.ime)
	binder.SetRender(c.renderFrame)
	// 呈现链路：真窗需要一个 Presenter 把合成出的 Backing 画到窗口上。
	// 像素提交走 LCL 无句柄控件的 OnPaint（见 host.BindCanvas），
	// 不直接写窗口 DC —— 后者在 DWM 重定向表面上会间歇性失败。
	if c.form != nil && c.host.Bound() {
		p := host.NewPresenter(uintptr(c.form.Handle()), c.host)
		p.SetCanvas(func() lcl.ICanvas {
			if cc := c.host.CanvasControl(); cc != nil {
				return cc.Canvas()
			}
			return nil
		})
		p.SetRepaint(c.host.Repaint)
		// 过渡帧语义：默认精确（内容 1:1 + 背景填边），整窗缩放需显式开启。
		p.SetResizePresent(c.opts.ResizePresent == ResizePresentStretch, c.resizeBackground())
		// 帧泵在 goroutine 里合成完就 post 一条高优先级呈现消息（posted 优先
		// 于硬件输入与 WM_PAINT，拖动 resize 时不会被饿死）。
		p.SetPost(func() { host.PostPresent(uintptr(c.form.Handle())) })
		c.host.OnPaint = func(cv lcl.ICanvas) {
			_ = cv // Presenter 自己取 Canvas，这里只触发提交
			p.Paint()
		}
		if err := c.host.BindCanvas(); err != nil {
			fmt.Fprintf(os.Stderr, "wg: 画布控件绑定失败: %v\n", err)
		}
		c.mu.Lock()
		c.presenter = p
		c.mu.Unlock()
		binder.SetPresenter(p)
		// 本窗关闭：停帧泵 + 释放呈现位图。OnClose 是主线程回调，
		// 此时 LCL 仍完整可用，sched.Stop/P.Close 都安全。
		// （主窗 Run 返回后的整体退出走进程回收，不经此路径。）
		if f := c.form; f != nil {
			f.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
				sched.Stop()
				p.Close()
				c.mu.Lock()
				c.st = stateClosed
				c.sched = nil
				c.schedAtomic.Store(nil)
				c.presenter = nil
				c.binder = nil
				c.mu.Unlock()
			})
		}
	}
	if err := binder.Install(); err != nil {
		c.mu.Lock()
		c.sched = nil
		c.schedAtomic.Store(nil)
		c.st = stateCreated
		c.presenter = nil
		c.mu.Unlock()
		sched.Stop()
		return err
	}
	c.mu.Lock()
	c.binder = binder
	c.mu.Unlock()

	sched.RequestFrame()
	return nil
}

// Run 启动帧泵并进入消息循环，阻塞直到窗口关闭或 Close 被调用。
//
// 帧泵的形状：**按需装填**。没有脏帧、没有浮层计时、没有输入时不装闹钟，
// 于是静止界面一帧都不出（C14）——这是"CPU 占用 0%"与"每帧 0.1ms 但每秒
// 60 次"之间的区别，也是笔记本风扇转不转的区别。
//
// 多窗口：本方法 = Start + 平台消息循环，作为**主窗**入口；
// 其余窗在 Run 之前调 Start()（见 Start 的注释）。
//
// 无窗口容器（NewHeadless）没有消息循环可进，返回 ErrNilForm；
// 那种场景请用 FrameOnce + ManualTicker 自己驱动。
func (c *Container) Run() error {
	// Start 已建好本窗的帧泵/消息接管/呈现器（含 OnClose 清理）；
	// 这里只进平台消息循环（全局唯一），返回即程序退出。
	if err := c.Start(); err != nil {
		return err
	}
	// 进入平台消息循环：帧闹钟在别的 goroutine 上触发，消息循环负责
	// 输入与窗口生命周期，两者不互相阻塞（否则拖动窗口时动画会停）。
	lcl.Application.Run()

	c.mu.Lock()
	s := c.sched
	c.sched = nil
	c.schedAtomic.Store(nil)
	c.st = stateClosed
	c.mu.Unlock()
	if s != nil {
		s.Stop()
	}
	return nil
}

// renderRequest 是一次"主线程整帧渲染"请求（单 outstanding）。
type renderRequest struct {
	done       chan struct{}
	dispatched atomic.Bool
	presented  bool
}

// frameOnceAndContinue 执行一帧并续动画节拍（渲染入口的唯一实现）。
func (c *Container) frameOnceAndContinue() bool {
	st, ferr := c.FrameOnce()
	if ferr == nil && c.animating.Load() {
		// 60Hz 节拍续帧：动态组件驱动渲染；静止窗不进此分支（C14）。
		// 钉在绝对帧网格上（requestNextTick），慢帧跳过错过的边界。
		if s := c.schedAtomic.Load(); s != nil {
			c.requestNextTick(s)
		}
	}
	if ferr == nil {
		// 光标闪烁是低频率自驱动帧：按相位续装闹钟（静止窗也只每 530ms 一帧）。
		c.scheduleCaretBlink()
	}
	return ferr == nil && st.Presented
}

// renderOnMain 把一帧投递到窗口主线程执行并等待完成（帧时钟 goroutine）。
//
// P0 方案①：布局/录制/合成/呈现与输入分发同线程，消除场景树的穿线竞态。
// 消息循环尚未启动（窗口启动早期）时退化为本 goroutine 内联合成一次。
func (c *Container) renderOnMain() bool {
	req := &renderRequest{done: make(chan struct{}, 1)}
	c.renderReq.Store(req)
	// 慢帧后先留空队列让输入先走，再投下一帧（showcase §9 的 inputBreath）。
	if c.renderCostUs.Load() > int64(frame.DefaultStep) {
		time.Sleep(1500 * time.Microsecond)
	}
	if !host.PostRender(c.renderHWND) {
		return c.frameOnceAndContinue()
	}
	select {
	case <-req.done:
		return req.presented
	case <-time.After(250 * time.Millisecond):
		if req.dispatched.CompareAndSwap(false, true) {
			// 主线程还没进消息循环：先内联合成（此处尚无输入，无竞态窗口）。
			return c.frameOnceAndContinue()
		}
		// 主线程已开始渲染但较慢：等它完成（闭环，不叠加第二帧）。
		<-req.done
		return req.presented
	}
}

// renderFrame 在窗口主线程执行一帧（Binder 收到 wmRequestRender）。
func (c *Container) renderFrame() {
	req := c.renderReq.Load()
	if req == nil || !req.dispatched.CompareAndSwap(false, true) {
		return
	}
	c.mu.Lock()
	p := c.presenter
	c.mu.Unlock()
	if p != nil {
		// 渲染方在主线程：present 不需要再走 posted 消息。
		p.SetSuppressPost(true)
	}
	began := time.Now()
	presented := c.frameOnceAndContinue()
	if p != nil {
		p.SetSuppressPost(false)
		if presented {
			p.Present() // 同步 Invalidate+Update → Paint
		}
	}
	c.renderCostUs.Store(int64(time.Since(began)))
	c.syncIMECaret() // 组合窗/候选窗跟随光标
	req.presented = presented
	req.done <- struct{}{}
}

// fireResize 是节流前沿或尾随到期后的真正 resize。
//
// 尺寸取**当前客户区**而不是请求快照：WM_SIZE 在消息线程已即时更新
// host.client，异步 fire 若用 1–2ms 前的快照会把客户区回滚到旧值
// （host.Resize 会无条件写回），导致合成尺寸与窗口来回摆动。
func (c *Container) fireResize(w, h int) {
	if cs := c.host.ClientSize(); cs.W > 0 && cs.H > 0 {
		w, h = int(cs.W), int(cs.H)
	}
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
			// 浮层计时/动画标脏由 FrameOnce 在容器锁内统一推进。
			return c.frameOnceAndContinue()
		},
		OnResize: c.fireResize,
	})
	if err != nil {
		c.mu.Unlock()
		return nil
	}
	c.sched = sched
	c.schedAtomic.Store(sched)
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
	c.schedAtomic.Store(nil)
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

// SetAnimating 设置"本窗有活跃动画"标志：置位后帧泵按 60Hz 节拍标脏
// 场景并续帧（动态组件驱动渲染）；清除后恢复严格按需渲染——有变化才出
// 帧，静止零帧（C14）。多窗隔离：每个容器独立持有自己的标志。
func (c *Container) SetAnimating(on bool) { c.animating.Store(on) }

// Animating 返回动画活跃标志。
func (c *Container) Animating() bool { return c.animating.Load() }
