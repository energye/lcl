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

type tyCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *tyCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeCopyableText 复制文本首帧 + 点击后重绘耗时。
func TestProbeCopyableText(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(800, 600)
	ct := widgets.NewCopyableText("可复制文本")
	cap := &tyCap{child: ct}
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.FrameOnce()
	began := time.Now()
	ct.MarkDirty()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("CopyableText 重绘=%v，阈值 16ms", elapsed)
	}
}
