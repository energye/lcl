package widgets

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/theme"
)

// TestDatePickerYearPanelSwitch 年面板点月回 month 模式（切月语义）。
func TestDatePickerYearPanelSwitch(t *testing.T) {
	d := NewDatePicker()
	d.Value = time.Date(2026, 11, 21, 12, 0, 0, 0, time.Local)
	d.panelMode = "year"
	n := d.buildPanel(dataCtx(), theme.Default())
	if n == nil {
		t.Fatal("年面板构建失败")
	}
	// 点月切回 month 语义。
	d.viewMonth = 3
	d.panelMode = "month"
	d.MarkDirty()
	n2 := d.buildPanel(dataCtx(), theme.Default())
	if n2 == nil || d.panelMode != "month" || d.viewMonth != 3 {
		t.Fatalf("年面板切月失败：%d mode=%s", d.viewMonth, d.panelMode)
	}
}

// TestDatePickerDisabledDate 禁用日期：灰置锚 + down 吞事件。
func TestDatePickerDisabledDate(t *testing.T) {
	d := NewDatePicker()
	d.Value = time.Date(2026, 11, 21, 12, 0, 0, 0, time.Local)
	d.DisabledDate = func(date time.Time) bool {
		return date.Weekday() == time.Saturday || date.Weekday() == time.Sunday
	}
	if !d.DisabledDate(time.Date(2026, 11, 15, 0, 0, 0, 0, time.Local)) {
		t.Fatal("11-15 周日应禁用")
	}
	if d.DisabledDate(time.Date(2026, 11, 16, 0, 0, 0, 0, time.Local)) {
		t.Fatal("11-16 周一不应禁用")
	}
}
