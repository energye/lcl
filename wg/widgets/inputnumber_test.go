package widgets

import "testing"

// TestInputNumberPrecision precision 小数位格式化（四舍五入 + -1 不限制）。
func TestInputNumberPrecision(t *testing.T) {
	cases := []struct {
		v    float32
		p    int
		want string
	}{
		{1.23456, 2, "1.23"},
		{12.3456, 2, "12.35"},
		{1, 2, "1.00"},
		{1.5, 0, "2"},
		{2.5, 0, "2"},
		{1.23456, -1, "1.23456"},
		{20, -1, "20"},
	}
	for _, c := range cases {
		w := NewInputNumber(-1000, 1000, c.v)
		w.Precision = c.p
		if got := w.displayText(); got != c.want {
			t.Errorf("v=%v p=%d displayText=%q，want %q", c.v, c.p, got, c.want)
		}
	}
	// 默认 Precision=-1 不限制。
	if got := NewInputNumber(0, 100, 1).Precision; got != -1 {
		t.Fatalf("默认 Precision=%d，want -1", got)
	}
}
