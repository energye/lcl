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

type dpCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *dpCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeDatePickerShowTime 覆盖 §6A：ShowTime 面板打开首帧耗时 +
// 时间列滚动视口不撑高面板（回归 1805px 撑爆的教训）。
func TestProbeDatePickerShowTime(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1000, 800)
	d := widgets.NewDatePicker()
	d.Value = time.Date(2026, 11, 18, 9, 30, 0, 0, time.Local)
	d.ShowTime = true
	cap := &dpCap{child: d}
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	// 预热一次（字体度量冷启动）。
	d.Show(cap.ctx)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	d.Close()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}

	began := time.Now()
	d.Show(cap.ctx)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("ShowTime 面板首帧=%v，阈值 16ms", elapsed)
	}
	entries := cap.ctx.Overlay.Entries()
	if len(entries) != 1 {
		t.Fatalf("overlay=%d，want 1", len(entries))
	}
	// 面板高度必须受控（时间列滚动视口），不得被 24 小时全展开撑爆。
	if h := entries[0].Content().Size().H; h > 520 {
		t.Fatalf("ShowTime 面板高度=%v，超过 520（时间列未滚动裁剪？）", h)
	}
}

// TestProbeDatePickerRangeShowTime 区间+时间联动：面板高度受控（时间列
// 滚动裁剪）+ 首帧达标。
func TestProbeDatePickerRangeShowTime(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1000, 800)
	d := widgets.NewDatePicker()
	d.Range = true
	d.ShowTime = true
	d.Value = time.Date(2026, 11, 10, 8, 15, 0, 0, time.Local)
	d.Value2 = time.Date(2026, 11, 20, 18, 45, 0, 0, time.Local)
	cap := &dpCap{child: d}
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	d.Show(cap.ctx)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	d.Close()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}

	began := time.Now()
	d.Show(cap.ctx)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("区间时间面板首帧=%v，阈值 16ms", elapsed)
	}
	entries := cap.ctx.Overlay.Entries()
	if len(entries) != 1 {
		t.Fatalf("overlay=%d，want 1", len(entries))
	}
	if h := entries[0].Content().Size().H; h > 520 {
		t.Fatalf("区间时间面板高度=%v，超过 520", h)
	}
}
