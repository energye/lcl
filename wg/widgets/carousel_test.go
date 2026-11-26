package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/ui"
)

// TestCarouselStep 箭头循环切换：越界回绕。
func TestCarouselStep(t *testing.T) {
	c := NewCarousel([]ui.Widget{&Label{Text: "a"}, &Label{Text: "b"}, &Label{Text: "c"}})
	ctx := dataCtx()

	// 初始 0，右箭头 → 1。
	c.step(1, ctx)
	if c.cur != 1 {
		t.Fatalf("右移后 cur=%d，want 1", c.cur)
	}
	// 继续右移 → 2。
	c.step(1, ctx)
	if c.cur != 2 {
		t.Fatalf("再右移 cur=%d，want 2", c.cur)
	}
	// 越界右移 → 回绕 0。
	c.step(1, ctx)
	if c.cur != 0 {
		t.Fatalf("越界右移 cur=%d，want 0（回绕）", c.cur)
	}
	// 越界左移 → 回绕末尾 2。
	c.step(-1, ctx)
	if c.cur != 2 {
		t.Fatalf("越界左移 cur=%d，want 2（回绕）", c.cur)
	}
}

// TestCarouselFadeState fade 启动后 fading=true，动画推进到位后复位。
func TestCarouselFadeState(t *testing.T) {
	c := NewCarousel([]ui.Widget{&Label{Text: "a"}, &Label{Text: "b"}})
	c.Fade = true
	ctx := dataCtx()

	c.startFade(ctx)
	if !c.fading {
		t.Fatal("startFade 后 fading 应为 true")
	}
	// 非 Fade 时 startFade 不应启动。
	c2 := NewCarousel([]ui.Widget{&Label{Text: "a"}})
	c2.startFade(ctx)
	if c2.fading {
		t.Fatal("Fade=false 时不应启动 fade")
	}
}
