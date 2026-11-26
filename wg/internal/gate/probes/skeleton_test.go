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

type skCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *skCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeSkeletonElements 100 个三形态元素批量 mount 耗时。
func TestProbeSkeletonElements(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1200, 900)
	col := scene.Flex(scene.Vertical, 4)
	cap := &skCap{child: skelWidget(func(ctx *ui.BuildContext) *scene.Node {
		for i := 0; i < 33; i++ {
			col.Add(widgets.SkeletonButton{Size: ui.SizeSmall}.Build(ctx))
			col.Add(widgets.SkeletonInput{Width: 120}.Build(ctx))
			col.Add(widgets.SkeletonImage{Width: 24, Height: 24}.Build(ctx))
		}
		return col
	})}
	began := time.Now()
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("99 骨架元素 mount+首帧=%v，阈值 50ms", elapsed)
	}
}

type skelWidget func(*ui.BuildContext) *scene.Node

func (f skelWidget) Build(ctx *ui.BuildContext) *scene.Node { return f(ctx) }
