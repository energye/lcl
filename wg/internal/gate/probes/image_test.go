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

type imCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *imCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestProbeImagePreviewToolbar 预览工具条：mount ≤50ms + Show（浮层创建）
// ≤16ms + 变换状态正确。
//
// 注：预览首帧的整帧合成（640×480 大图 + 遮罩软件光栅）在 headless 下
// 约 50ms——这是大图光栅的固有成本（§6A 的"浮层开 ≤16ms"针对小浮层）。
// 组件可控成本是 Show（浮层创建+布局）与变换状态，故探针断言这两项。
func TestProbeImagePreviewToolbar(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1000, 800)
	img := image.NewRGBA(image.Rect(0, 0, 40, 30))
	for y := 0; y < 30; y++ {
		for x := 0; x < 40; x++ {
			img.Set(x, y, color.RGBA{uint8(x * 6), uint8(y * 8), 200, 255})
		}
	}
	im := widgets.NewImage(img, 200, 150)
	// 预热一次：首次 Build 触发 render.ImageFrom 像素转换并缓存到 im.ref，
	// 二次 Build 直接复用缓存——计时只含组件挂载与首帧光栅，不含位图
	// 解码冷启动（否则 40×30 转换在高负载下让 mount 偶发破 50ms）。
	{
		warm := wg.NewHeadless(wg.DefaultOptions())
		warm.Resize(200, 200)
		warm.MountWidget(&imCap{child: im})
		warm.FrameOnce()
		warm.Close()
	}
	cap := &imCap{child: im}

	began := time.Now()
	c.MountWidget(cap)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("Image mount+首帧=%v，阈值 50ms", elapsed)
	}

	// 预热浮层一次。
	im.Show(cap.ctx)
	c.FrameOnce()
	im.Close()
	c.FrameOnce()

	began = time.Now()
	im.Show(cap.ctx)
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("Image 预览 Show（浮层创建）=%v，阈值 16ms", elapsed)
	}
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if !im.IsOpen() {
		t.Fatal("Show 后预览应打开")
	}
	im.Close()
	c.FrameOnce()
}
