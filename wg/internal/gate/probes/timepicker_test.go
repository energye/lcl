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

type tpCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *tpCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeTimePickerDisabledTime 覆盖 §6A：禁用时段面板首帧 + 禁用格
// 渲染不显著拖慢（禁用判定在 buildColumn 内逐格计算）。
func TestProbeTimePickerDisabledTime(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1000, 800)
	tp := widgets.NewTimePicker()
	tp.Value = time.Date(0, 1, 1, 9, 10, 0, 0, time.Local)
	tp.DisabledTime = func(time.Time) widgets.DisabledTimeSpec {
		return widgets.DisabledTimeSpec{
			DisabledHours: func() []int { return []int{22, 23, 0, 1, 2, 3, 4, 5} },
			DisabledMinutes: func(h int) []int {
				if h == 9 {
					return []int{30, 45}
				}
				return nil
			},
		}
	}
	cap := &tpCap{child: tp}
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	tp.Show(cap.ctx)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	tp.Close()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	began := time.Now()
	tp.Show(cap.ctx)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("禁用时段面板首帧=%v，阈值 16ms", elapsed)
	}
	if n := cap.ctx.Overlay.Count(); n != 1 {
		t.Fatalf("overlay=%d，want 1", n)
	}
}

// TestProbeTimePickerFormatColumns format 列裁剪面板首帧。
func TestProbeTimePickerFormatColumns(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1000, 800)
	tp := widgets.NewTimePicker()
	tp.Format = "15:04"
	tp.Value = time.Date(0, 1, 1, 14, 25, 0, 0, time.Local)
	cap := &tpCap{child: tp}
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	tp.Show(cap.ctx)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	tp.Close()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	began := time.Now()
	tp.Show(cap.ctx)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("列裁剪面板首帧=%v，阈值 16ms", elapsed)
	}
}
