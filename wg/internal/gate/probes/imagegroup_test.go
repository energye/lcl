//go:build gateprobe

package probes

import (
	"image"
	"image/color"
	"testing"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

type igCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *igCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

func mkIGImg(r uint8) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 20, 15))
	for y := 0; y < 15; y++ {
		for x := 0; x < 20; x++ {
			img.Set(x, y, color.RGBA{r, 100, 200, 255})
		}
	}
	return img
}

// TestProbeImageGroup 多图组 mount + 预览首帧耗时。
func TestProbeImageGroup(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1000, 800)
	g := widgets.NewImageGroup(
		widgets.NewImage(mkIGImg(50), 120, 90),
		widgets.NewImage(mkIGImg(150), 120, 90),
		widgets.NewImage(mkIGImg(250), 120, 90),
	)
	cap := &igCap{child: g}
	began := time.Now()
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("ImageGroup mount+首帧=%v，阈值 50ms", elapsed)
	}
}
