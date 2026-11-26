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

type elCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *elCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeEllipsisText 省略文本：mount + 展开切换重绘耗时。
func TestProbeEllipsisText(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(800, 600)
	e := widgets.NewEllipsisText("这是一段很长的说明文字，用来演示多行省略与展开收起效果，确保超过两行触发省略。")
	e.Rows = 2
	e.MaxWidth = 300
	e.Expandable = true
	cap := &elCap{child: e}
	began := time.Now()
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("EllipsisText mount+首帧=%v，阈值 50ms", elapsed)
	}
	// 展开切换重绘。
	e.Toggle()
	began = time.Now()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("EllipsisText 展开切换=%v，阈值 16ms", elapsed)
	}
}
