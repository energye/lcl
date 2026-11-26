package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/theme"
)

// TestProgressRailColor railColor 解析：自定义优先，零值回退 SurfaceActive。
func TestProgressRailColor(t *testing.T) {
	th := theme.Default()

	// 零值：回退 SurfaceActive。
	p := Progress{Value: 0.4}
	if got := p.railColor(th); got != th.Color.SurfaceActive {
		t.Fatalf("零值 railColor=%v，want SurfaceActive", got)
	}

	// 自定义：优先。
	custom := geom.RGBA(186, 224, 255)
	p2 := Progress{Value: 0.4, RailColor: custom}
	if got := p2.railColor(th); got != custom {
		t.Fatalf("自定义 railColor=%v，want %v", got, custom)
	}

	// 线形与圆环共用同一 railColor（构建不 panic，锚存在）。
	for _, shape := range []ProgressShape{ProgressLine, ProgressCircle, ProgressDashboard} {
		pp := Progress{Value: 0.4, Shape: shape, RailColor: custom}
		if n := pp.Build(dataCtx()); n == nil {
			t.Fatalf("shape=%d 构建失败", shape)
		}
	}
}
