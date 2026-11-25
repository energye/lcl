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

type inCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *inCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeInputNumberPrecision precision 格式化首帧 + 步进重绘耗时。
func TestProbeInputNumberPrecision(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(800, 600)
	w := widgets.NewInputNumber(0, 1000, 12.3456)
	w.Precision = 2
	cap := &inCap{child: w}
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	// 预热。
	c.FrameOnce()
	began := time.Now()
	w.MarkDirty()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("InputNumber precision 重绘=%v，阈值 16ms", elapsed)
	}
}
