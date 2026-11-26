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

type wmCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *wmCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeWatermarkMultiline 多行水印首帧耗时（单 CustomPaint 平铺矩阵）。
func TestProbeWatermarkMultiline(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(800, 600)
	w := widgets.NewWatermark(nil, "")
	w.Contents = []string{"机密文件", "2026 年度", "请勿外传"}
	w.Width, w.Height = 480, 240
	cap := &wmCap{child: w}
	began := time.Now()
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("多行水印 mount+首帧=%v，阈值 50ms", elapsed)
	}
}
