package composite

import (
	"errors"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
)

// TextPainter 是文字回放通道（由 gdi 包实现并注入）。
//
// 为什么用接口注入而不是直接 import gdi：
//   - 保持 composite 的平台无关性（依赖表里 composite 不含 gdi）；
//   - headless 自检可以注入空实现，从而在非 Windows 上跑完整帧流程。
type TextPainter interface {
	// PaintText 在 dst 上回放 dl 中的文字指令，只处理与 clip 相交的部分。
	PaintText(dst *raster.Backing, dl *render.DisplayList, clip geom.Rect) int
}

// Options 是合成器选项。
type Options struct {
	// Backend 是形状光栅后端（必填）。
	Backend raster.Backend
	// Text 是文字通道（可为 nil，此时文字不出图但会计入 TextSkipped）。
	Text TextPainter
	// Width/Height 是初始尺寸。
	Width, Height int
	// Background 是清屏色（默认不透明白）。
	Background geom.Color
}

// Stats 是一帧的合成统计。
type Stats struct {
	// Raster 是后端统计。
	Raster raster.Stats
	// TextLines 是文字通道绘制的行数。
	TextLines int
	// TextSkipped 表示有文字但没装文字通道（headless 场景）。
	TextSkipped bool
	// DamageArea 是本帧提交的面积（像素²）。
	DamageArea float32
	// DamageRects 是本帧提交的矩形数。
	DamageRects int
	// Presented 表示本帧是否真的提交了（静止帧为 false，对应 C14）。
	Presented bool
	// Duration 是本帧耗时。
	Duration time.Duration
	// Frame 是帧序号。
	Frame int
}

// Compositor 负责"把指令表变成像素"并维护 damage。
//
// 一帧的顺序是固定的（对应 03-ENGINE.md §9）：
//
//	accumulate damage → clear(damage, bg) → backend.Draw(damage) → text pass → commit
//
// 注意清屏只清 damage 区：整屏清屏会让"局部重绘"失去意义（多画一倍像素），
// 也会把未变化的静态内容擦掉再重画（示例里 24.7ms 的整窗合成正是这么来的）。
type Compositor struct {
	backend  raster.Backend
	text     TextPainter
	dst      *raster.Backing
	bg       geom.Color
	damage   Region
	last     Region
	frame    int
	lastSize geom.Size
}

// New 创建合成器。
func New(opts Options) (*Compositor, error) {
	if opts.Backend == nil {
		return nil, errors.New("wg/composite: 缺少光栅后端")
	}
	if opts.Width <= 0 || opts.Height <= 0 {
		return nil, errors.New("wg/composite: 尺寸非法")
	}
	bg := opts.Background
	if bg.A == 0 && opts.Background == (geom.Color{}) {
		bg = geom.RGBA(255, 255, 255)
	}
	c := &Compositor{
		backend: opts.Backend,
		text:    opts.Text,
		bg:      bg,
	}
	if err := c.Resize(opts.Width, opts.Height); err != nil {
		return nil, err
	}
	return c, nil
}

// Backing 返回合成目标（32bpp 预乘 BGRA）。调用方只读，不得修改。
func (c *Compositor) Backing() *raster.Backing { return c.dst }

// Backend 返回光栅后端（诊断用）。
func (c *Compositor) Backend() raster.Backend { return c.backend }

// SetTextPainter 注入/替换文字通道。
func (c *Compositor) SetTextPainter(t TextPainter) { c.text = t }

// Resize 改变合成目标尺寸，并把整屏标脏（尺寸变化必须全重画）。
func (c *Compositor) Resize(w, h int) error {
	if w <= 0 || h <= 0 {
		return errors.New("wg/composite: 尺寸非法")
	}
	dst := raster.NewBacking(w, h)
	if dst == nil {
		return errors.New("wg/composite: 帧缓冲分配失败")
	}
	c.dst = dst
	c.lastSize = geom.Sz(float32(w), float32(h))
	c.damage.Clear()
	c.damage.Add(dst.Bounds())
	c.last.Clear()
	return nil
}

// Invalidate 标记一块区域需要重绘（区域坐标 = 合成目标坐标）。
func (c *Compositor) Invalidate(r geom.Rect) {
	if r.IsEmpty() || c.dst == nil {
		return
	}
	c.damage.Add(r.Intersect(c.dst.Bounds()))
}

// InvalidateAll 标记整屏。
func (c *Compositor) InvalidateAll() {
	if c.dst == nil {
		return
	}
	c.damage.Add(c.dst.Bounds())
}

// Damage 返回当前累计的脏区（只读）。
func (c *Compositor) Damage() *Region { return &c.damage }

// LastPresented 返回上一帧提交的区域。
func (c *Compositor) LastPresented() *Region { return &c.last }

// Draw 执行一帧。
//
// dl 为 nil 表示"本帧没有新指令"（例如只有 damage 需要清屏），仍然会提交。
// 返回的 Stats.Presented=false 表示没有可提交内容（静止帧）。
func (c *Compositor) Draw(dl *render.DisplayList) (Stats, error) {
	began := time.Now()
	var st Stats
	if c.dst == nil {
		return st, errors.New("wg/composite: 合成目标未初始化")
	}
	st.Frame = c.frame

	// 注意：这里**不**把 dl.Bounds() 自动并入脏区。
	//
	// damage 的唯一来源是上层（画布树/场景）的失效通知。如果录制完就把
	// 包围盒当地次脏区，那么"每帧重录"会让每帧都全量重画，C14（静止零帧）
	// 立刻失效——而重录本身在 M2 之后只在脏帧发生，所以这条捷径只会误导。
	// 忘了 Invalidate 的后果由 --damage 的面积断言暴露，而不是靠这里兜底。

	if c.damage.IsEmpty() {
		// 静止帧：不提交（C14）
		st.Duration = time.Since(began)
		c.frame++
		return st, nil
	}

	clip := c.damage.Bounds().Intersect(c.dst.Bounds())
	if clip.IsEmpty() {
		c.damage.Clear()
		st.Duration = time.Since(began)
		c.frame++
		return st, nil
	}

	// 1) 只清脏区，不做整屏清屏
	for _, rc := range c.damage.Rects() {
		c.dst.FillRect(rc, c.bg)
	}

	// 2) 形状光栅
	if dl != nil && !dl.Empty() {
		st.Raster = c.backend.Draw(c.dst, dl, clip)
	}

	// 3) 文字回放（像素已就位，设备空间）
	if dl != nil && dl.HasText() {
		if c.text != nil {
			st.TextLines = c.text.PaintText(c.dst, dl, clip)
		} else {
			st.TextSkipped = true
		}
	}

	// 4) 提交：记录本帧区域后清空 damage（一次提交，C13）
	c.last.Clear()
	c.last.AddRegion(&c.damage)
	st.DamageArea = c.damage.Area()
	st.DamageRects = c.damage.Len()
	st.Presented = true
	c.damage.Clear()

	st.Duration = time.Since(began)
	c.frame++
	return st, nil
}
