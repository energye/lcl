package widgets

import (
	"testing"
	"time"
)

// TestTimePickerUse12Hours 12 小时制：displayText 形态与 24h 还原
//（ap12To24 双向映射语义锁）。
func TestTimePickerUse12Hours(t *testing.T) {
	tp := NewTimePicker()
	tp.Use12Hours = true
	tp.Value = time.Date(0, 1, 1, 15, 20, 0, 0, time.Local)
	if got := tp.displayText(); got != "03:20:00 下午" {
		t.Fatalf("15:20 应显示 03:20:00 下午，got %q", got)
	}
	tp.Value = time.Date(0, 1, 1, 0, 5, 0, 0, time.Local)
	if got := tp.displayText(); got != "12:05:00 上午" {
		t.Fatalf("0:05 应显示 12:05:00 上午，got %q", got)
	}
	tp.Value = time.Date(0, 1, 1, 12, 0, 0, 0, time.Local)
	if got := tp.displayText(); got != "12:00:00 下午" {
		t.Fatalf("12:00 应显示 12:00:00 下午，got %q", got)
	}
	// ap12To24 映射。
	cases := []struct {
		h12, ap, want int
	}{
		{12, 0, 0}, {1, 0, 1}, {11, 0, 11},
		{12, 1, 12}, {1, 1, 13}, {11, 1, 23},
	}
	for _, c := range cases {
		if got := ap12To24(c.h12, c.ap); got != c.want {
			t.Fatalf("ap12To24(%d,%d)=%d 应 %d", c.h12, c.ap, got, c.want)
		}
	}
	// 小时列值集合：12h 顺序 12,1..11。
	tp2 := NewTimePicker()
	tp2.Use12Hours = true
	hv := tp2.hourValues()
	if len(hv) != 12 || hv[0] != 12 || hv[1] != 1 || hv[11] != 11 {
		t.Fatalf("12h 小时列应 [12,1..11]，got %v", hv)
	}
}

// TestTimePickerSteps 步进列值集合：HourStep=2 → 12 格；分钟步长 10 →
// 6 格；0/1 不步进全量。
func TestTimePickerSteps(t *testing.T) {
	if got := stepValues(24, 2); len(got) != 12 || got[1] != 2 {
		t.Fatalf("HourStep=2 应 12 格偶数集合，got %v", got)
	}
	if got := stepValues(60, 10); len(got) != 6 || got[5] != 50 {
		t.Fatalf("MinuteStep=10 应 6 格，got %v", got)
	}
	if got := stepValues(24, 0); len(got) != 24 {
		t.Fatalf("step 0 不步进应 24 格，got %d", len(got))
	}
	tp := NewTimePicker()
	tp.HourStep = 2
	if got := tp.hourValues(); len(got) != 12 {
		t.Fatalf("24h HourStep=2 应 12 格，got %d", len(got))
	}
}
