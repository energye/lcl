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

	cur      int
	dots     []*scene.Node
	autoplay *time.Timer
}

// NewCarousel 创建走马灯。
func NewCarousel(slides []ui.Widget) *Carousel {
	return &Carousel{Slides: slides}
}

// Build 实现 ui.Widget。
func (c *Carousel) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(c)
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
		if ctx.Input != nil {
			// 页不可见时不响应点击（hit 测试对 !Visible 节点跳过）。
			ctx.Input.Set(page, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				return i == c.cur
			}})
		}
		stack.Add(page)
	}
	col.Add(stack)

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
				c.goTo(i, th)
				return true
			}})
		}
		dotsRow.Add(dot)
	}
	col.Add(dotsRow)

	c.startAutoplay(th)
	return col
}

// goTo 切到第 i 页（可见性原地切换+圆点重建刷新）。
func (c *Carousel) goTo(i int, th *theme.Data) {
	if i < 0 || i >= len(c.Slides) || i == c.cur {
		return
	}
	c.cur = i
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
