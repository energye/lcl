//go:build gateprobe

package probes

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

type inputCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *inputCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeInputAutoSize autosize 多行首帧 + 高度随行数增长。
func TestProbeInputAutoSize(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(800, 600)
	w := widgets.NewInput("多行")
	w.Multiline = true
	w.AutoSize = true
	w.AutoMinRows = 2
	w.AutoMaxRows = 6
	w.Width = 320
	w.Value = "第一行\n第二行\n第三行"
	w.Editor().SetValue(input.EditingValue{Text: w.Value, Selection: input.RangeAt(input.RuneLen(w.Value))})
	cap := &inputCap{child: w}
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.FrameOnce()
	began := time.Now()
	w.MarkDirty()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("Input autosize 重绘=%v，阈值 16ms", elapsed)
	}
}
