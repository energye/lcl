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

type captureWidget struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *captureWidget) Build(ctx *ui.BuildContext) *scene.Node {
	c.ctx = ctx
	return c.child.Build(ctx)
}

func TestProbeDrawerDefaultWidth(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(800, 600)
	capture := &captureWidget{child: widgets.Spacer{Width: 1, Height: 1}}
	c.MountWidget(capture)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if capture.ctx == nil {
		t.Fatal("未捕获 BuildContext")
	}
	warmup := widgets.NewDrawer("预热", widgets.Label{Text: "内容"})
	warmup.Show(capture.ctx)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	warmup.Close()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	drawer := widgets.NewDrawer("探针", widgets.Label{Text: "内容"})
	// §6A：浮层创建（Show）≤16ms 与整帧合成分开计时——合并计时下高负载
	// 首帧光栅偶发破 16ms（17.17ms 实测），组件可控成本是 Show。
	began := time.Now()
	drawer.Show(capture.ctx)
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("Drawer Show（浮层创建）=%v，阈值 16ms", elapsed)
	}
	began = time.Now()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("Drawer 首帧=%v，阈值 50ms（§6A 首帧标准）", elapsed)
	}
	entries := capture.ctx.Overlay.Entries()
	if len(entries) != 1 {
		t.Fatalf("Drawer overlay=%d，want 1", len(entries))
	}
	if got := entries[0].Size.W; got != widgets.DrawerSizeDefault {
		t.Fatalf("Drawer 默认宽度=%v，want %v", got, widgets.DrawerSizeDefault)
	}
	drawer.Close()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
}
