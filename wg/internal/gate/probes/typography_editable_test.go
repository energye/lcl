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

type etCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *etCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeEditableText 可编辑文本：mount + 切编辑态重绘耗时。
func TestProbeEditableText(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(800, 600)
	et := widgets.NewEditableText("可编辑文本")
	cap := &etCap{child: et}
	began := time.Now()
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("EditableText mount+首帧=%v，阈值 50ms", elapsed)
	}
	// 切编辑态重绘。
	et.BeginEdit()
	began = time.Now()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("EditableText 切编辑态=%v，阈值 16ms", elapsed)
	}
}
