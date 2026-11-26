//go:build gateprobe

package probes

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

type pgCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *pgCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeProgressRailColor 100 个自定义轨道色进度批量 mount 耗时。
func TestProbeProgressRailColor(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1000, 900)
	rail := geom.RGBA(186, 224, 255)
	col := scene.Flex(scene.Vertical, 4)
	cap := &pgCap{child: skelWidget(func(ctx *ui.BuildContext) *scene.Node {
		for i := 0; i < 100; i++ {
			p := widgets.Progress{Value: 0.4, Width: 200, RailColor: rail}
			col.Add(p.Build(ctx))
		}
		return col
	})}
	began := time.Now()
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("100 进度 mount+首帧=%v，阈值 50ms", elapsed)
	}
}
