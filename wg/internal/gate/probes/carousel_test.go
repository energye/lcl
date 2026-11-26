//go:build gateprobe

package probes

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

type crCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *crCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeCarouselArrowsFade 箭头+fade 首帧 + 切页重绘耗时。
func TestProbeCarouselArrowsFade(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(800, 600)
	cr := widgets.NewCarousel([]ui.Widget{
		&widgets.Label{Text: "A"}, &widgets.Label{Text: "B"}, &widgets.Label{Text: "C"},
	})
	cr.Arrows = true
	cr.Fade = true
	cap := &crCap{child: cr}
	began := time.Now()
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("Carousel mount+首帧=%v，阈值 50ms", elapsed)
	}
	// 切页重绘。
	cr.Next()
	began = time.Now()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("Carousel 切页=%v，阈值 16ms", elapsed)
	}
}
