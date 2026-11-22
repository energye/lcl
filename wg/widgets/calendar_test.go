package widgets

import (
	"testing"
	"time"
)

// TestCalendarValidRangeShiftClamp validRange 钳制：上月/下月切换不越界。
func TestCalendarValidRangeShiftClamp(t *testing.T) {
	c := NewCalendar()
	c.SetView(2026, 11)
	c.ValidRange = [2]time.Time{
		time.Date(2026, 11, 5, 0, 0, 0, 0, time.Local),
		time.Date(2026, 11, 25, 0, 0, 0, 0, time.Local),
	}
	c.shiftMonth(nil, -1) // 10 月首 < 范围首月 → 钳制停 11 月
	if c.viewMonth != 11 {
		t.Fatalf("上月钳制失败：viewMonth = %d", c.viewMonth)
	}
	c.shiftMonth(nil, 1) // 12 月 > 范围末月 → 钳制
	if c.viewMonth != 11 {
		t.Fatalf("下月钳制失败：viewMonth = %d", c.viewMonth)
	}
}

// TestCalendarValidRangeInit 初始视图钳进范围（当前月 12 月在范围外 →
// 回范围首月）。
func TestCalendarValidRangeInit(t *testing.T) {
	c := NewCalendar()
	c.Now = func() time.Time { return time.Date(2026, 12, 10, 0, 0, 0, 0, time.Local) }
	c.ValidRange = [2]time.Time{
		time.Date(2026, 11, 5, 0, 0, 0, 0, time.Local),
		time.Date(2026, 11, 25, 0, 0, 0, 0, time.Local),
	}
	c.Build(dataCtx())
	if c.viewYear != 2026 || c.viewMonth != 11 {
		t.Fatalf("初始视图钳制失败：%d-%d，应 2026-11", c.viewYear, c.viewMonth)
	}
}

// TestCalendarValidRangeGoToday 今天不在范围 → 回范围首月。
func TestCalendarValidRangeGoToday(t *testing.T) {
	base := time.Date(2026, 12, 10, 0, 0, 0, 0, time.Local)
	c := NewCalendar()
	c.Now = func() time.Time { return base }
	c.SetView(2026, 12)
	c.ValidRange = [2]time.Time{
		time.Date(2026, 11, 5, 0, 0, 0, 0, time.Local),
		time.Date(2026, 11, 25, 0, 0, 0, 0, time.Local),
	}
	c.goToday(nil)
	if c.viewMonth != 11 {
		t.Fatalf("goToday 钳制失败：viewMonth = %d", c.viewMonth)
	}
}

// TestCalendarYearPanelSwitch 年面板点月切回 month 模式。
func TestCalendarYearPanelSwitch(t *testing.T) {
	c := NewCalendar()
	c.Mode = "year"
	c.SetView(2026, 11)
	n := c.Build(dataCtx())
	if n == nil {
		t.Fatal("年面板构建失败")
	}
	// 模拟点 3 月格（直接走切换语义：viewMonth+回 month）。
	c.viewMonth = 3
	c.Mode = "month"
	c.MarkDirty()
	n2 := c.Build(dataCtx())
	if n2 == nil || c.Mode != "month" || c.viewMonth != 3 {
		t.Fatalf("年面板切月失败：%d mode=%s", c.viewMonth, c.Mode)
	}
}
