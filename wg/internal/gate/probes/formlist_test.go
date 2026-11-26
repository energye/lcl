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

type flCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *flCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeFormList 动态行容器：100 行 mount + Add/Remove 重绘耗时。
func TestProbeFormList(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(900, 900)
	l := widgets.NewFormList("items", "a", "b")
	for i := 0; i < 99; i++ {
		l.Add()
	}
	cap := &flCap{child: l}
	began := time.Now()
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("FormList 100 行 mount+首帧=%v，阈值 50ms", elapsed)
	}
	// Add 重绘。
	l.Add()
	began = time.Now()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("FormList Add 重绘=%v，阈值 16ms", elapsed)
	}
}
