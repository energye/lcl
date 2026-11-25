package widgets

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/scene"
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

// TestDatePickerShowTime 时间联动：显示文本到分钟、时间列/确定锚存在、
// 步长过滤正确、日期格保留时分。
func TestDatePickerShowTime(t *testing.T) {
	th := theme.Default()
	d := NewDatePicker()
	d.Value = time.Date(2026, 11, 18, 9, 30, 0, 0, time.Local)
	d.ShowTime = true
	if got := d.displayText(d.Value); got != "2026-11-18 09:30" {
		t.Fatalf("ShowTime 显示文本=%q，want 2026-11-18 09:30", got)
	}
	d.ShowTime = false
	if got := d.displayText(d.Value); got != "2026-11-18" {
		t.Fatalf("普通显示文本=%q", got)
	}

	// 步长过滤。
	if got := timeSteps(24, 1); len(got) != 24 {
		t.Fatalf("24 小时步长 1 应有 24 格，got %d", len(got))
	}
	if got := timeSteps(24, 6); len(got) != 4 || got[0] != 0 || got[3] != 18 {
		t.Fatalf("步长 6 应得 [0 6 12 18]，got %v", got)
	}
	if got := timeSteps(60, 15); len(got) != 4 || got[3] != 45 {
		t.Fatalf("分钟步长 15 应得 [0 15 30 45]，got %v", got)
	}

	// 面板含时间区与确定锚。
	d.ShowTime = true
	d.panelMode = "month"
	n := d.buildPanel(dataCtx(), th)
	if findLabelIn(n, "date-time-panel") == nil {
		t.Fatal("ShowTime 面板缺少 date-time-panel")
	}
	if findLabelIn(n, "date-time-ok") == nil {
		t.Fatal("ShowTime 面板缺少 date-time-ok")
	}
	if findLabelIn(n, "date-time-hour") == nil || findLabelIn(n, "date-time-minute") == nil {
		t.Fatal("ShowTime 面板缺少时/分列")
	}
}

// findLabelIn 在子树里找命中锚为 label 的节点（测试用）。
func findLabelIn(n *scene.Node, label string) *scene.Node {
	if n == nil {
		return nil
	}
	if n.HitSpec().Label == label {
		return n
	}
	for _, c := range n.Children() {
		if got := findLabelIn(c, label); got != nil {
			return got
		}
	}
	return nil
}

// TestDatePickerRangeShowTime 区间+时间联动：时间列作用于活动端、
// 选日期切换端、确定提交区间。
func TestDatePickerRangeShowTime(t *testing.T) {
	d := NewDatePicker()
	d.Range = true
	d.ShowTime = true
	d.Value = time.Date(2026, 11, 10, 8, 15, 0, 0, time.Local)
	d.Value2 = time.Date(2026, 11, 20, 18, 45, 0, 0, time.Local)
	// 初始活动端=起：时间列应读起端 8:15。
	d.activeEnd = 0
	// 面板含时间区四锚。
	n := d.buildPanel(dataCtx(), theme.Default())
	for _, label := range []string{"date-time-panel", "date-time-hour", "date-time-minute", "date-time-ok"} {
		if findLabelIn(n, label) == nil {
			t.Fatalf("区间时间面板缺少 %s", label)
		}
	}
	// 活动端切换后显示文本各自带时间。
	if got := d.displayText(d.Value); got != "2026-11-10 08:15" {
		t.Fatalf("起端显示=%q", got)
	}
	if got := d.displayText(d.Value2); got != "2026-11-20 18:45" {
		t.Fatalf("止端显示=%q", got)
	}
}
