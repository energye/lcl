package wg

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/frame"
)

// TestNextTickDelay 锁定续帧节拍：墙钟钉绝对 60Hz 网格，手动时钟固定 16ms。
func TestNextTickDelay(t *testing.T) {
	step := frame.DefaultStep
	cases := []struct {
		kind frame.ClockKind
		now  time.Duration
		want time.Duration
	}{
		// 墙钟：整拍起点 → 等一整拍（不能 0，否则贴着边界空转）
		{frame.Wall, 0, step},
		{frame.Wall, step, step},
		// 墙钟：帧耗时后的相位 → 补到下一个边界，周期恒为 60Hz
		{frame.Wall, 5 * time.Millisecond, step - 5*time.Millisecond},
		{frame.Wall, step + 15*time.Millisecond, step - 15*time.Millisecond},
		// 手动时钟：固定 16ms（确定性门禁按 16ms 片推进）
		{frame.Manual, 5 * time.Millisecond, 16 * time.Millisecond},
		{frame.Manual, 0, 16 * time.Millisecond},
	}
	for _, c := range cases {
		if got := nextTickDelay(c.kind, c.now, step); got != c.want {
			t.Errorf("nextTickDelay(%v, %v) = %v，应为 %v", c.kind, c.now, got, c.want)
		}
	}
}
