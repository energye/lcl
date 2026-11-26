package widgets

// ── Carousel 走马灯（M12 长尾）───────────────────────────────────
//
// antd：<Carousel autoplay dots> 横向走马灯：Stack 叠层切换（SetVisible
// 翻页——零成本切换，无动画轮实现滑动过渡）+ 底部圆点指示（当前点
// 主色实底放大）+ **点击圆点切换** + autoplay 定时轮播（默认关）。
// 未实现（如实标注）：滑动过渡动画、arrows 箭头、vertical 纵向、
// fade/dots 变体、Swipe 触摸手势。

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// Carousel 走马灯。
type Carousel struct {
	controlBase
	// Slides 每页内容（Widget 列表）。
	Slides []ui.Widget
	// Autoplay 自动轮播间隔（0=关闭；>0 毫秒）。审计确定性下默认 0，
	// 演示可开（如 3000）。
	Autoplay int
	Width    float32
	Height   float32
	// Arrows 显示左右切换箭头（antd arrows，5.17.0）。
	Arrows bool
	// Fade 淡入淡出过渡（antd fade/effect="fade"）：切页时新页淡入。
	Fade bool

	cur      int
	dots     []*scene.Node
	autoplay *time.Timer
	showCtx  *ui.BuildContext // 最近 Build ctx（程序化 Next/Prev 用）
	// fade 动画（切页时新页透明度 0→1 插值）。
	fadeAnim *ui.Controller
	fadeNode *scene.Node
	fading   bool
}

// NewCarousel 创建走马灯。
func NewCarousel(slides []ui.Widget) *Carousel {
	return &Carousel{Slides: slides}
}

// Build 实现 ui.Widget。
func (c *Carousel) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(c)
	c.showCtx = ctx
	th := ctx.ThemeOr()
	w := c.Width
	if w <= 0 {
		w = 480
	}
	h := c.Height
	if h <= 0 {
		h = 180
	}
	if c.cur >= len(c.Slides) {
		c.cur = 0
	}

	col := scene.Flex(scene.Vertical, 0)
	col.SetWidth(w)
	col.SetHit(scene.HitSpec{Label: "carousel"})

	// 页叠层。
	stack := scene.Stack().SetSize2(w, h)
	stack.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "carousel-slides"})
	for i, sl := range c.Slides {
		i := i
		page := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.SurfaceSunken),
			Border: scene.Stroke(th.Color.Border, 1),
		}).SetSize2(w, h)
		if inner := sl.Build(ctx); inner != nil {
			inner.SetMargin(geom.InLTRB(0, (h-inner.Size().H)/2, 0, 0))
			inner.SetPositionKeepSize(true)
			inner.SetStackPosition(geom.Pt((w-inner.Size().W)/2, (h-inner.Size().H)/2))
			page.Add(inner)
		}
		page.SetVisible(i == c.cur)
		if c.Fade && i == c.cur && c.fading && c.fadeAnim != nil {
			page.SetOpacity(c.fadeAnim.Value())
			c.fadeNode = page
		}
		if ctx.Input != nil {
			// 页不可见时不响应点击（hit 测试对 !Visible 节点跳过）。
			ctx.Input.Set(page, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				return i == c.cur
			}})
		}
		stack.Add(page)
	}
	col.Add(stack)

	// 左右切换箭头（antd arrows）：叠在页层两侧，半透明深底圆钮。
	if c.Arrows && len(c.Slides) > 1 {
		arrowRow := scene.Stack().SetSize2(w, h)
		mkArrow := func(icon, label string, delta int) *scene.Node {
			ic := (&Icon{Name: icon, Size: 18, Color: th.Color.OnDark}).Build(ctx)
			btn := scene.Box(scene.BoxSpec{
				Fill:   scene.Solid(geom.RGBAa(0, 0, 0, 90)),
				Radius: 16,
			}).SetSize2(32, 32)
			ic.SetPositionKeepSize(true)
			ic.SetStackPosition(geom.Pt(7, 7))
			btn.Add(ic)
			btn.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: label, Cursor: scene.CursorPointer})
			if ctx.Input != nil {
				ctx.Input.Set(btn, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
					if !ev.Down() {
						return false
					}
					c.step(delta, ctx)
					return true
				}})
			}
			return btn
		}
		prev := mkArrow("left", "carousel-prev", -1)
		prev.SetPositionKeepSize(true)
		prev.SetStackPosition(geom.Pt(th.Space.MD, (h-32)/2))
		arrowRow.Add(prev)
		next := mkArrow("right", "carousel-next", 1)
		next.SetPositionKeepSize(true)
		next.SetStackPosition(geom.Pt(w-th.Space.MD-32, (h-32)/2))
		arrowRow.Add(next)
		col.Add(arrowRow)
	}

	// 圆点指示（底部居中）。
	dotsRow := scene.Flex(scene.Horizontal, th.Space.SM)
	dotsRow.SetWidth(w)
	dotsRow.SetCrossAlign(scene.CrossCenter)
	dotsRow.SetPadding(geom.InLTRB(0, th.Space.SM, 0, 0))
	for i := range c.Slides {
		i := i
		cur := i == c.cur
		size := float32(8)
		if cur {
			size = 10
		}
		fill := th.Color.TextFaint
		if cur {
			fill = th.Color.Primary
		}
		dot := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(fill),
			Radius: size / 2,
		}).SetSize2(size, size)
		dot.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "carousel-dot", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			ctx.Input.Set(dot, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				if !ev.Down() {
					return false
				}
				c.goTo(i, ctx)
				return true
			}})
		}
		dotsRow.Add(dot)
	}
	col.Add(dotsRow)

	if c.fading {
		c.driveFade()
	}
	c.startAutoplay(th)
	return col
}

// Next 切到下一页（antd next，循环）。
func (c *Carousel) Next() { c.step(1, c.showCtx) }

// Prev 切到上一页（antd prev，循环）。
func (c *Carousel) Prev() { c.step(-1, c.showCtx) }

// GoTo 切到第 i 页（antd goTo）。
func (c *Carousel) GoTo(i int) {
	if i < 0 || i >= len(c.Slides) {
		return
	}
	c.cur = i
	c.startFade(c.showCtx)
	c.MarkDirty()
}

// Current 返回当前页索引（诊断/断言用）。
func (c *Carousel) Current() int { return c.cur }

// step 按 delta 循环切换（箭头用）。
func (c *Carousel) step(delta int, ctx *ui.BuildContext) {
	if len(c.Slides) == 0 {
		return
	}
	next := (c.cur + delta + len(c.Slides)) % len(c.Slides)
	c.cur = next
	c.startFade(ctx)
	c.MarkDirty()
}

// startFade 启动切页淡入（Fade=true 时新页 opacity 0→1）。
func (c *Carousel) startFade(ctx *ui.BuildContext) {
	if !c.Fade || ctx == nil {
		return
	}
	if c.fadeAnim == nil {
		c.fadeAnim = ui.NewController(0, theme.EaseOut)
	}
	c.fadeAnim.SetDuration(ctx.Motion(300 * time.Millisecond))
	c.fadeAnim.SetValue(0)
	c.fading = true
	c.fadeAnim.Forward(time.Since(animBase))
	if ctx.Scheduler != nil {
		ctx.Scheduler.RequestFrame()
	}
}

// driveFade 推进淡入一帧（递归 AfterFunc 续帧）。
func (c *Carousel) driveFade() {
	if c.fadeAnim == nil {
		return
	}
	if c.fadeAnim.Step(time.Since(animBase)) {
		time.AfterFunc(16*time.Millisecond, func() {
			if c.fading {
				c.MarkDirty()
			}
		})
		return
	}
	c.fading = false
	c.MarkDirty()
}

// goTo 切到第 i 页（可见性原地切换+圆点重建刷新）。
func (c *Carousel) goTo(i int, ctx *ui.BuildContext) {
	if i < 0 || i >= len(c.Slides) || i == c.cur {
		return
	}
	c.cur = i
	c.startFade(ctx)
	c.MarkDirty()
}

// startAutoplay 自动轮播（Autoplay>0 时；timer 跨重建持有）。
func (c *Carousel) startAutoplay(th *theme.Data) {
	if c.Autoplay <= 0 {
		return
	}
	if c.autoplay != nil {
		c.autoplay.Stop()
	}
	c.autoplay = time.AfterFunc(time.Duration(c.Autoplay)*time.Millisecond, func() {
		next := (c.cur + 1) % len(c.Slides)
		c.cur = next
		c.MarkDirty()
		c.startAutoplay(th)
	})
}
