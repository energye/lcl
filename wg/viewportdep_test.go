package wg

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// vpPage 是可计数的测试页：useViewport 决定 Build 是否声明视口依赖。
type vpPage struct {
	ui.Base
	builds      int
	useViewport bool
}

func (p *vpPage) Build(ctx *ui.BuildContext) *scene.Node {
	p.builds++
	size := geom.Sz(300, 200)
	if p.useViewport {
		size = ctx.ViewportSize()
	}
	return scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
	}).SetSize2(size.W, size.H)
}

// TestViewportDependencyRebuild 锁定「按需重建」语义：
// 只有声明了视口依赖的页面在 resize 时重建；固定布局页面只重排布局。
func TestViewportDependencyRebuild(t *testing.T) {
	t.Run("fixed-layout-no-rebuild", func(t *testing.T) {
		c := NewHeadless(DefaultOptions())
		defer c.Close()
		c.Resize(400, 300)
		p := &vpPage{}
		c.MountWidget(p)
		if p.builds != 1 {
			t.Fatalf("首次挂载 Build=%d，应为 1", p.builds)
		}
		c.Resize(800, 600)
		if _, err := c.FrameOnce(); err != nil {
			t.Fatal(err)
		}
		if p.builds != 1 {
			t.Fatalf("固定布局页面 resize 不应重建，Build=%d", p.builds)
		}
		b := c.Backing()
		if b == nil || b.W() != 800 || b.H() != 600 {
			t.Fatalf("布局应跟随视口：backing=%v", b)
		}
	})

	t.Run("viewport-dependent-rebuilds", func(t *testing.T) {
		c := NewHeadless(DefaultOptions())
		defer c.Close()
		c.Resize(400, 300)
		p := &vpPage{useViewport: true}
		c.MountWidget(p)
		if p.builds != 1 {
			t.Fatalf("首次挂载 Build=%d，应为 1", p.builds)
		}
		c.Resize(800, 600)
		if _, err := c.FrameOnce(); err != nil {
			t.Fatal(err)
		}
		if p.builds != 2 {
			t.Fatalf("声明视口依赖的页面 resize 应重建一次，Build=%d", p.builds)
		}
		// 再次 resize 继续重建。
		c.Resize(1000, 700)
		if _, err := c.FrameOnce(); err != nil {
			t.Fatal(err)
		}
		if p.builds != 3 {
			t.Fatalf("再次 resize 应再重建，Build=%d", p.builds)
		}
	})
}
