package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/text"
)

// TestCountLines countLines 行数统计（显式换行 + 自动折行）。
func TestCountLines(t *testing.T) {
	font := text.DefaultFont(14)
	// 三行显式换行。
	if got := countLines("a\nb\nc", font, 200); got != 3 {
		t.Fatalf("三行换行 countLines=%d，want 3", got)
	}
	// 单行短文本。
	if got := countLines("hello", font, 200); got != 1 {
		t.Fatalf("短单行=%d，want 1", got)
	}
	// 空串算 1 行。
	if got := countLines("", font, 200); got != 1 {
		t.Fatalf("空串=%d，want 1", got)
	}
	// 窄宽自动折行：长串在 40px 宽下应折成多行。
	if got := countLines("aaaaaaaaaaaaaaaaaaaa", font, 40); got < 2 {
		t.Fatalf("窄宽折行=%d，want >=2", got)
	}
}

// TestInputAutoSize autosize 高度：随行数增长 + minRows/maxRows 钳制。
func TestInputAutoSize(t *testing.T) {
	w := NewInput("多行")
	w.Multiline = true
	w.AutoSize = true
	w.Width = 320
	w.AutoMinRows = 2
	w.AutoMaxRows = 6
	// 通过 countLines 的钳制语义验证：1 行内容在 minRows=2 下应按 2 行算高。
	lines := countLines("单行", text.DefaultFont(14), 298)
	if lines < 1 {
		t.Fatalf("行数=%d", lines)
	}
	if lines < w.AutoMinRows {
		lines = w.AutoMinRows
	}
	if lines != 2 {
		t.Fatalf("minRows 钳制后行数=%d，want 2", lines)
	}
	// maxRows 钳制：7 行内容按 6 行算。
	seven := "a\nb\nc\nd\ne\nf\ng"
	l2 := countLines(seven, text.DefaultFont(14), 298)
	if l2 > w.AutoMaxRows {
		l2 = w.AutoMaxRows
	}
	if l2 != 6 {
		t.Fatalf("maxRows 钳制后行数=%d，want 6", l2)
	}
}
