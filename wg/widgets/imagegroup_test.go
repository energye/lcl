package widgets

import (
	"image"
	"image/color"
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// imgCap 测试用 Widget 包装。
type imgCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *imgCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

func mkTestImg(r, g, b uint8) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, 20, 15))
	for y := 0; y < 15; y++ {
		for x := 0; x < 20; x++ {
			img.SetNRGBA(x, y, color.NRGBA{R: r, G: g, B: b, A: 255})
		}
	}
	return img
}

// TestImageGroupStep 多图切换：循环、计数、OnChange。
func TestImageGroupStep(t *testing.T) {
	g := NewImageGroup(
		NewImage(mkTestImg(255, 0, 0), 100, 75),
		NewImage(mkTestImg(0, 255, 0), 100, 75),
		NewImage(mkTestImg(0, 0, 255), 100, 75),
	)
	var changes [][2]int
	g.OnChange = func(cur, prev int) { changes = append(changes, [2]int{cur, prev}) }

	g.current = 0
	// 右移 → 1。
	g.step(dataCtx(), 1)
	if g.Current() != 1 {
		t.Fatalf("右移后 current=%d，want 1", g.Current())
	}
	// 右移 → 2。
	g.step(dataCtx(), 1)
	if g.Current() != 2 {
		t.Fatalf("再右移 current=%d，want 2", g.Current())
	}
	// 越界右移 → 回绕 0。
	g.step(dataCtx(), 1)
	if g.Current() != 0 {
		t.Fatalf("越界右移 current=%d，want 0（循环）", g.Current())
	}
	// 越界左移 → 回绕 2。
	g.step(dataCtx(), -1)
	if g.Current() != 2 {
		t.Fatalf("越界左移 current=%d，want 2（循环）", g.Current())
	}
	// OnChange 收到 4 次，最后一次 (2, 0)。
	if len(changes) != 4 {
		t.Fatalf("OnChange 次数=%d，want 4", len(changes))
	}
	if changes[3] != [2]int{2, 0} {
		t.Fatalf("末次 OnChange=%v，want [2 0]", changes[3])
	}
	// 计数文案 1-based。
	g.current = 0
	if got := g.countText(); got != "1/3" {
		t.Fatalf("countText=%q，want 1/3", got)
	}
	g.current = 2
	if got := g.countText(); got != "3/3" {
		t.Fatalf("countText=%q，want 3/3", got)
	}
}

// TestImageGroupRefCache refOf 懒解码 + 越界 nil。
func TestImageGroupRefCache(t *testing.T) {
	g := NewImageGroup(
		NewImage(mkTestImg(1, 2, 3), 50, 40),
		nil,
	)
	if r := g.refOf(0); r == nil {
		t.Fatal("第 0 张应有 ImageRef")
	}
	// 缓存命中（第二次返回同一指针）。
	r1 := g.refOf(0)
	r2 := g.refOf(0)
	if r1 != r2 {
		t.Fatal("refOf 应缓存同一指针")
	}
	// nil 项返回 nil。
	if g.refOf(1) != nil {
		t.Fatal("nil Thumbs 项应返回 nil")
	}
	// 越界 nil。
	if g.refOf(99) != nil {
		t.Fatal("越界应返回 nil")
	}
	_ = render.ImageBilinear
	_ = geom.Pt
}
