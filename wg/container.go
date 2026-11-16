package wg

import (
	"sync"
	"time"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/host"
)

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
// 它持有：宿主（Form + 根画布）、帧时钟、以及（后续）调度器、输入泵、
// 浮层宿主与主题。使用者只与 Container 打交道，不直接接触 Surface 树。
//
// 线程模型：所有公开方法都是线程安全的（内部单锁）。但**结构性变更**
// （Mount/Detach/Resize）只允许在帧外调用；帧内变更必须走 RequestFrame
// 投递到下一帧（契约 C18）。
type Container struct {
	mu   sync.Mutex
	form lcl.IEngForm
	opts Options
	host *host.Host
	// clock 是帧时间源；默认墙钟。
	clock frame.FrameClock

	// index 是下一个帧序号。
	index int
	// last 是上一帧的时钟快照，用于计算 Delta。
	last time.Duration

	// requested 表示有非 damage 原因的帧请求（动画、异步结果回填）。
	requested bool

	st state
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
func NewHeadless(opts Options) *Container {
	return NewContainer(nil, opts)
}

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
//
// headless 自检与 golden 出图必须注入 ManualClock，否则出图不可复现。
func (c *Container) SetClock(clk frame.FrameClock) {
	if clk == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.st == stateRunning {
		return // 运行中换时钟会让动画时间跳变，禁止
	}
	c.clock = clk
	c.last = clk.Now()
}

// Resize 修改无窗口容器的客户区尺寸（绑定了窗体的容器由 host 的 WndProc 驱动）。
func (c *Container) Resize(w, h int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.st == stateClosed || c.host == nil {
		return
	}
	c.host.Resize(w, h)
}

// RequestFrame 请求在下一帧处理（不动画、无 damage 的场景，如异步结果回填）。
//
// 这是唯一允许跨线程调用的"触发绘制"入口：它只置一个标志位，
// 真正的结构变更仍必须在帧外由持有者串行完成（C18）。
func (c *Container) RequestFrame() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.st == stateClosed {
		return
	}
	c.requested = true
}

// FrameOnce 走一遍帧流程并返回统计，用于 headless 自检与确定性出帧。
//
// M0 语义：**只统计，不提交**——本阶段还没有光栅与呈现，因此不清除脏区
// （没有真的把像素画上去就清 damage，会让下一帧漏画，属于典型脏区 bug）。
// M1 接入光栅后，这里会在一次不透明提交之后 ClearDamageTree（C13）。
func (c *Container) FrameOnce() (frame.Stats, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.st == stateClosed || c.host == nil {
		return frame.Stats{}, ErrClosed
	}

	began := time.Now()
	now := c.clock.Now()
	st := frame.Stats{Index: c.index, Now: now}
	if c.index > 0 {
		st.Delta = now - c.last
	}

	root := c.host.Root()
	if root != nil {
		st.Surfaces = root.Count()
		st.Damage = root.Damage()
		st.Dirty = root.HasDamage() || root.StructureDirty() || c.requested
		// 结构性脏但没有具体 damage（例如首次挂载、尺寸变化）时，
		// 脏区取整块根画布，避免"该画却没画"。
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

// Frames 返回已走过的帧数。
func (c *Container) Frames() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.index
}

// Run 启动帧泵与输入泵，阻塞直到 Close。
//
// M0 未接入调度器（契约 C11 的单 outstanding 帧槽在 M2 落地），
// 因此这里返回 ErrNotReady 而不是静默返回——静默返回会让调用方以为
// 窗口已经在画，实际一帧都没有。
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
	c.st = stateRunning
	c.mu.Unlock()
	return ErrNotReady
}

// Close 关闭容器并释放画布树。
//
// 释放顺序遵守 C21：先解绑视图（Surface），再释放位图，最后归还池。
// M0 没有位图与池，因此只做视图解绑。
func (c *Container) Close() {
	c.mu.Lock()
	if c.st == stateClosed {
		c.mu.Unlock()
		return
	}
	c.st = stateClosed
	h := c.host
	c.host = nil
	c.mu.Unlock()

	if h == nil {
		return
	}
	if root := h.Root(); root != nil {
		// 视图解绑：从叶子往上摘，避免父被摘掉后子仍持有 parent 指针。
		// 注意 Detach 会修改父的 children，因此必须先取快照再递归。
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
