//go:build gateprobe

package probes

import (
	"image"
	"image/color"
	"testing"

	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// TestProbeLedgerErratum 台账勘误回归：Avatar.Src 图片头像与 Badge.Child
// 包裹 Avatar 均能构建（证明缺口条目确已交付，非代码回退）。
func TestProbeLedgerErratum(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			img.Set(x, y, color.RGBA{uint8(x * 8), uint8(y * 8), 200, 255})
		}
	}
	// Avatar.Src 图片头像。
	av := &widgets.Avatar{Src: img, Size: 36}
	if n := av.Build((&ui.BuildContext{Theme: theme.Default()})); n == nil {
		t.Fatal("Avatar.Src 图片头像构建失败（台账缺口应已交付）")
	}
	// Badge.Child 包裹 Avatar（图片头像场景）。
	bg := &widgets.Badge{Count: 3, Child: av}
	if n := bg.Build(&ui.BuildContext{Theme: theme.Default()}); n == nil {
		t.Fatal("Badge.Child 包裹构建失败（台账缺口应已交付）")
	}
}
