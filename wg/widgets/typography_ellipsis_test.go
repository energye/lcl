package widgets

import "testing"

// TestEllipsisText 省略展开：展开态切换 + 按钮文案。
func TestEllipsisText(t *testing.T) {
	e := NewEllipsisText("长文本内容")
	e.Rows = 2
	e.Expandable = true

	var expanded bool
	e.OnExpand = func(v bool) { expanded = v }

	if e.Expanded() {
		t.Fatal("初始不应展开")
	}
	if got := e.symbolLabel(); got != "展开" {
		t.Fatalf("收起态按钮=%q，want 展开", got)
	}
	// 模拟点击切换。
	e.expanded = !e.expanded
	e.OnExpand(e.expanded)
	if !e.Expanded() || !expanded {
		t.Fatalf("展开后 Expanded=%v 回调=%v", e.Expanded(), expanded)
	}
	if got := e.symbolLabel(); got != "收起" {
		t.Fatalf("展开态按钮=%q，want 收起", got)
	}

	// 自定义 symbol。
	e2 := NewEllipsisText("x")
	e2.Symbol = "更多"
	if got := e2.symbolLabel(); got != "更多" {
		t.Fatalf("自定义 symbol 收起态=%q，want 更多", got)
	}
	e2.expanded = true
	if got := e2.symbolLabel(); got != "收起更多" {
		t.Fatalf("自定义 symbol 展开态=%q，want 收起更多", got)
	}
}
