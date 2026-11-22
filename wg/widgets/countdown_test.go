package widgets

import (
	"strings"
	"testing"
	"time"

	"github.com/energye/lcl/wg/scene"
)

// nodeContainsText 子树文本查找（测试辅助）。
func nodeContainsText(n *scene.Node, s string) bool {
	found := false
	n.Walk(func(x *scene.Node) bool {
		if sp := x.Text(); sp != nil && sp.Src == s {
			found = true
			return false
		}
		return true
	})
	return found
}

// TestCountdownFormatToken dayjs token 子集格式化全路径。
func TestCountdownFormatToken(t *testing.T) {
	cases := []struct {
		d    time.Duration
		f    string
		want string
	}{
		{2*time.Hour + 30*time.Minute, "HH:mm:ss", "02:30:00"},
		{26*time.Hour + 5*time.Second, "D 天 HH:mm:ss", "1 天 02:00:05"}, // 26h=1 天 2h
		{90 * time.Second, "mm:ss", "01:30"},
		{1*time.Minute + 2*time.Second + 500*time.Millisecond, "mm:ss.SSS", "01:02.500"},
		{0, "HH:mm:ss", "00:00:00"},
		{-time.Hour, "HH:mm:ss", "00:00:00"}, // 负值钳零
		{5 * time.Hour, "H", "5"},            // 不补零总时数
	}
	for _, c := range cases {
		if got := formatDuration(c.d, c.f); got != c.want {
			t.Fatalf("formatDuration(%v, %q) = %q，应 %q", c.d, c.f, got, c.want)
		}
	}
}

// TestCountdownOnFinish Now 注入（真实流逝偏移）+ time.AfterFunc 触发。
func TestCountdownOnFinish(t *testing.T) {
	base := time.Date(2026, 11, 22, 12, 0, 0, 0, time.Local)
	c := NewCountdown("测试", base.Add(2*time.Second))
	start := time.Now()
	c.Now = func() time.Time { return base.Add(time.Since(start)) } // 注入基准+真实流逝
	fired := make(chan struct{}, 1)
	c.OnFinish = func() { fired <- struct{}{} }
	n := c.Build(dataCtx())
	if n == nil {
		t.Fatal("Build 失败")
	}
	// 初始显示以 00:00:0 开头（Build 耗时致秒数 0-2 抖动，形态断言）。
	if !nodeHasTextPrefix(n, "00:00:0") {
		t.Fatal("初始文本应以 00:00:0 开头")
	}
	// AfterFunc 真实触发 OnFinish。
	select {
	case <-fired:
	case <-time.After(4 * time.Second):
		t.Fatal("OnFinish 未触发")
	}
	c.Stop()
}

// nodeHasTextPrefix 子树文本前缀查找。
func nodeHasTextPrefix(n *scene.Node, prefix string) bool {
	found := false
	n.Walk(func(x *scene.Node) bool {
		if sp := x.Text(); sp != nil && strings.HasPrefix(sp.Src, prefix) {
			found = true
			return false
		}
		return true
	})
	return found
}

// TestCountdownCountUp 正计时：Target 为起点，显示已过时长。
func TestCountdownCountUp(t *testing.T) {
	base := time.Date(2026, 11, 22, 12, 0, 0, 0, time.Local)
	c := NewCountdown("计时", base)
	c.Now = func() time.Time { return base.Add(90 * time.Second) }
	c.CountUp = true
	n := c.Build(dataCtx())
	if !nodeContainsText(n, "00:01:30") {
		t.Fatal("正计时应显示 00:01:30")
	}
}

// TestCountdownDayFormat 天级格式 + onChange 回调。
func TestCountdownDayFormat(t *testing.T) {
	base := time.Date(2026, 11, 22, 12, 0, 0, 0, time.Local)
	c := NewCountdown("大促", base.Add(26*time.Hour+5*time.Second))
	c.Now = func() time.Time { return base }
	c.Format = "D 天 HH:mm:ss"
	changed := make(chan time.Duration, 1)
	c.OnChange = func(d time.Duration) { changed <- d }
	n := c.Build(dataCtx())
	if !nodeContainsText(n, "1 天 02:00:05") {
		t.Fatal("天级格式应为 1 天 02:00:05")
	}
	select {
	case d := <-changed:
		want := 26*time.Hour + 5*time.Second
		if d < want-time.Second || d > want {
			t.Fatalf("OnChange 剩余 = %v，应 ≈%v", d, want)
		}
	default:
		t.Fatal("OnChange 未触发")
	}
	c.Stop()
}
