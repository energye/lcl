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

type fmCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *fmCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeFormLayout 三布局 100 表单项批量 mount 耗时。
func TestProbeFormLayout(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1000, 900)
	col := scene.Flex(scene.Vertical, 6)
	cap := &fmCap{child: skelWidget(func(ctx *ui.BuildContext) *scene.Node {
		for i := 0; i < 33; i++ {
			v := widgets.FormItem{Label: "垂直", Child: widgets.NewInput(""), Layout: widgets.FormVertical, Required: true}
			col.Add(v.Build(ctx))
			h := widgets.FormItem{Label: "水平", Child: widgets.NewInput(""), LabelColSpan: 6}
			col.Add(h.Build(ctx))
			in := widgets.FormItem{Label: "行内", Child: widgets.NewInput(""), Layout: widgets.FormInline}
			col.Add(in.Build(ctx))
		}
		return col
	})}
	began := time.Now()
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("99 表单项 mount+首帧=%v，阈值 50ms", elapsed)
	}
}
