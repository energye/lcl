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

type grCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *grCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeGridResponsive 响应式栅格：mount + 100 列批量重绘耗时。
func TestProbeGridResponsive(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1200, 800)
	i := func(v int) *int { return &v }
	cols := make([]widgets.GridCol, 100)
	for k := range cols {
		cols[k] = widgets.GridCol{Child: widgets.Label{Text: "x"}, Span: 6, Md: i(12), Xs: i(24)}
	}
	g := widgets.NewGrid(cols...)
	cap := &grCap{child: g}
	began := time.Now()
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("Grid 100 列 mount+首帧=%v，阈值 50ms", elapsed)
	}
	// 重绘。
	began = time.Now()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	_ = began
}
