package frame

import (
	"sync"
	"testing"
	"time"

	"github.com/energye/lcl/wg/geom"
)

func TestManualClockDeterministic(t *testing.T) {
	a := NewManualClock(0, 5*time.Millisecond)
	b := NewManualClock(0, 5*time.Millisecond)
	for i := 0; i < 10; i++ {
		if x, y := a.Advance(), b.Advance(); x != y {
			t.Fatalf("第 %d 步不一致: %v vs %v", i, x, y)
		}
	}
	if a.Now() != 50*time.Millisecond {
		t.Fatalf("Now = %v", a.Now())
	}
	if a.Kind() != Manual || a.Kind().String() != "manual" {
		t.Fatal("Kind 错误")
	}
	a.Reset(0)
	if a.Now() != 0 || a.Ticks() != 0 {
		t.Fatal("Reset 未生效")
	}
}

func TestManualClockDefaultStep(t *testing.T) {
	c := NewManualClock(0, 0)
	if c.Step() != DefaultStep {
		t.Fatalf("默认步长 = %v", c.Step())
	}
	// 60Hz 步长应接近 16.67ms
	if d := DefaultStep.Seconds(); d < 0.0166 || d > 0.0167 {
		t.Fatalf("60Hz 步长 = %v", d)
	}
	c.SetStep(0) // 非法值应被忽略
	if c.Step() != DefaultStep {
		t.Fatal("非法步长不该生效")
	}
	c.SetStep(time.Millisecond)
	c.Advance()
	if c.Now() != time.Millisecond {
		t.Fatalf("改步长后 Now = %v", c.Now())
	}
	c.AdvanceBy(-5) // 负值不前进
	if c.Now() != time.Millisecond {
		t.Fatal("负值不该前进")
	}
	if c.AdvanceBy(2*time.Millisecond) != 3*time.Millisecond {
		t.Fatalf("AdvanceBy = %v", c.Now())
	}
}

func TestWallClockMonotonic(t *testing.T) {
	c := NewWallClock()
	if c.Kind() != Wall || c.Kind().String() != "wall" {
		t.Fatal("Kind 错误")
	}
	first := c.Now()
	time.Sleep(2 * time.Millisecond)
	if second := c.Now(); second <= first {
		t.Fatalf("墙钟未单调递增: %v → %v", first, second)
	}
}

func TestManualClockConcurrent(t *testing.T) {
	c := NewManualClock(0, time.Millisecond)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				c.Advance()
				_ = c.Now()
			}
		}()
	}
	wg.Wait()
	if c.Ticks() != 400 {
		t.Fatalf("并发推进后 ticks = %d", c.Ticks())
	}
}

func TestMedian(t *testing.T) {
	if got := Median(nil); got != 0 {
		t.Fatalf("空输入 = %v", got)
	}
	if got := Median([]time.Duration{5}); got != 5 {
		t.Fatalf("单元素 = %v", got)
	}
	if got := Median([]time.Duration{16, 17, 16, 300, 16}); got != 16 {
		t.Fatalf("奇数中位数 = %v", got)
	}
	// 排序后 [1,2,3,4]，偶数个取中间两个的均值 = 2.5ms
	if got := Median([]time.Duration{4 * time.Millisecond, 1 * time.Millisecond, 3 * time.Millisecond, 2 * time.Millisecond}); got != 2500*time.Microsecond {
		t.Fatalf("偶数中位数 = %v", got)
	}
}

func TestPercentile(t *testing.T) {
	d := []time.Duration{16, 16, 16, 100, 200}
	if got := Percentile(d, 0); got != 16 {
		t.Fatalf("P0 = %v", got)
	}
	if got := Percentile(d, 100); got != 200 {
		t.Fatalf("P100 = %v", got)
	}
	if got := Percentile(d, 50); got != 16 {
		t.Fatalf("P50 = %v", got)
	}
	if got := Percentile(nil, 50); got != 0 {
		t.Fatalf("空输入 = %v", got)
	}
}

func TestStatsString(t *testing.T) {
	s := Stats{Index: 3, Dirty: true, Damage: geom.RectXYWH(10, 20, 30, 40)}
	str := s.String()
	if str == "" || !containsAll(str, []string{"frame=3", "dirty", "surfaces="}) {
		t.Fatalf("Stats.String = %q", str)
	}
	if (Stats{}).Idle() != true {
		t.Fatal("零值应为 idle")
	}
}

func TestKindString(t *testing.T) {
	if Compositor.String() != "compositor" {
		t.Fatal("Compositor.String 错误")
	}
}

func containsAll(s string, subs []string) bool {
	for _, sub := range subs {
		found := false
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
