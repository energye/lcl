package widgets

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// TestTimePickerUse12Hours 12 小时制：displayText 形态与 24h 还原
// （ap12To24 双向映射语义锁）。
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

func TestTimePickerSelectionIndices(t *testing.T) {
	tests := []struct {
		name                 string
		setup                func(*TimePicker)
		want                 [4]int
		pickCol, pickIdx     int
		wantHour, wantMinute int
	}{
		{
			name: "12h 午夜",
			setup: func(tp *TimePicker) {
				tp.Use12Hours = true
				tp.Value = time.Date(0, 1, 1, 0, 0, 0, 0, time.Local)
			},
			want: [4]int{0, 0, 0, 0}, pickCol: 0, pickIdx: 1, wantHour: 1,
		},
		{
			name: "12h 下午",
			setup: func(tp *TimePicker) {
				tp.Use12Hours = true
				tp.Value = time.Date(0, 1, 1, 15, 20, 0, 0, time.Local)
			},
			want: [4]int{3, 20, 0, 1}, pickCol: -1,
		},
		{
			name: "分钟步进",
			setup: func(tp *TimePicker) {
				tp.MinuteStep = 10
				tp.Value = time.Date(0, 1, 1, 9, 20, 0, 0, time.Local)
			},
			want: [4]int{9, 2, 0, -1}, pickCol: 1, pickIdx: 4, wantHour: 9, wantMinute: 40,
		},
		{
			name: "步进集合无当前值",
			setup: func(tp *TimePicker) {
				tp.MinuteStep = 10
				tp.Value = time.Date(0, 1, 1, 9, 25, 0, 0, time.Local)
			},
			want: [4]int{9, -1, 0, -1}, pickCol: 1, pickIdx: 3, wantHour: 9, wantMinute: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &ui.BuildContext{Theme: theme.Default()}
			tp := NewTimePicker()
			tt.setup(tp)
			tp.curTh = ctx.ThemeOr()
			tp.buildPanel(ctx, ctx.ThemeOr())
			if tp.selIdx != tt.want {
				t.Fatalf("selIdx=%v, want %v", tp.selIdx, tt.want)
			}
			if tt.pickCol >= 0 {
				tp.pick(tt.pickCol, tt.pickIdx)
				if tp.Value.Hour() != tt.wantHour || tp.Value.Minute() != tt.wantMinute {
					t.Fatalf("pick 后时刻=%02d:%02d, want %02d:%02d", tp.Value.Hour(), tp.Value.Minute(), tt.wantHour, tt.wantMinute)
				}
			}
		})
	}
}

func TestTimePickerHighlightAllUsesCellIndices(t *testing.T) {
	ctx := &ui.BuildContext{Theme: theme.Default()}
	tp := NewTimePicker()
	tp.Use12Hours = true
	tp.MinuteStep = 10
	tp.Value = time.Date(0, 1, 1, 15, 20, 0, 0, time.Local)
	tp.curTh = ctx.ThemeOr()
	tp.buildPanel(ctx, ctx.ThemeOr())
	tp.highlightAll(0, 35, 0)
	if tp.selIdx[0] != 0 || tp.selIdx[1] != 5 || tp.selIdx[3] != 0 {
		t.Fatalf("highlightAll 后 selIdx=%v, want [0 5 0 0]", tp.selIdx)
	}
}

// TestTimePickerDisabledTime 禁用时段：disabledSet 各列取值 + 12h 换算。
func TestTimePickerDisabledTime(t *testing.T) {
	tp := NewTimePicker()
	tp.Value = time.Date(0, 1, 1, 9, 10, 0, 0, time.Local)
	tp.DisabledTime = func(time.Time) DisabledTimeSpec {
		return DisabledTimeSpec{
			DisabledHours: func() []int { return []int{22, 23, 0} },
			DisabledMinutes: func(h int) []int {
				if h == 9 {
					return []int{30, 45}
				}
				return nil
			},
			DisabledSeconds: func(h, m int) []int {
				if h == 9 && m == 10 {
					return []int{5}
				}
				return nil
			},
		}
	}
	// 时列：22/23/0 禁用。
	hs := tp.disabledSet(0)
	if !hs[22] || !hs[23] || !hs[0] || hs[9] {
		t.Fatalf("时列禁用集合错误: %v", hs)
	}
	// 分列：当前小时 9 → 30/45 禁用。
	ms := tp.disabledSet(1)
	if !ms[30] || !ms[45] || ms[10] {
		t.Fatalf("分列禁用集合错误: %v", ms)
	}
	// 秒列：h=9,m=10 → 5 禁用。
	ss := tp.disabledSet(2)
	if !ss[5] || ss[10] {
		t.Fatalf("秒列禁用集合错误: %v", ss)
	}
	// 无 DisabledTime 时返回 nil。
	tp2 := NewTimePicker()
	if tp2.disabledSet(0) != nil {
		t.Fatal("无 DisabledTime 应返回 nil")
	}
}
