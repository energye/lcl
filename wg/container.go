package wg

import (
	"sync"
	"time"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/wg/composite"
	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/raster/gdi"
	"github.com/energye/lcl/wg/raster/software"
	"github.com/energye/lcl/wg/render"
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
	mu   sync.Mutex
	form lcl.IEngForm
	opts Options
	host *host.Host
	clock frame.FrameClock

	// 渲染管线（首次出帧时惰性装配）
	backend         raster.Backend
	backendOverride raster.Backend
	textPainter     composite.TextPainter
	comp            *composite.Compositor
	rec             *render.Recorder
	paint           Painter

	index     int
	last      time.Duration
	requested bool
	st        state
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
}

// RequestFrame 请求在下一帧处理（无 damage 的场景，如异步结果回填）。
func (c *Container) RequestFrame() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.st == stateClosed {
		return
	}
	c.requested = true
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

	// 结构变化（增删画布/改尺寸）必须整屏重画：布局可能整体位移。
	dirty := root.HasDamage() || root.StructureDirty() || c.requested
	if root.StructureDirty() {
		c.comp.InvalidateAll()
	} else if d := root.Damage(); !d.IsEmpty() {
		c.comp.Invalidate(d)
	}

	var dl *render.DisplayList
	if dirty {
		c.rec.Reset()
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
	st.Damage = c.comp.LastPresented().Bounds()

	if cs.Presented {
		root.ClearDamageTree()
		root.ClearStructureDirty()
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
// M1 仍未接入调度器（单 outstanding 帧槽在 M2 落地），因此返回 ErrNotReady
// 而不是静默返回——静默返回会让调用方以为窗口已经在画。
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
	c.host = nil
	c.comp = nil
	c.rec = nil
	c.mu.Unlock()

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
