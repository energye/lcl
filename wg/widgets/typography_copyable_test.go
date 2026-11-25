package widgets

import "testing"

// TestCopyableText 复制：OnCopy 回调、已复制态、CopyText 覆盖。
func TestCopyableText(t *testing.T) {
	var got string
	c := NewCopyableText("正文内容")
	c.OnCopy = func(s string) { got = s }
	c.doCopy()
	if got != "正文内容" {
		t.Fatalf("OnCopy 收到=%q，want 正文内容", got)
	}
	if !c.Copied() {
		t.Fatal("复制后应进入已复制态")
	}

	// CopyText 覆盖 Body。
	var got2 string
	c2 := NewCopyableText("显示正文")
	c2.CopyText = "实际复制"
	c2.OnCopy = func(s string) { got2 = s }
	c2.doCopy()
	if got2 != "实际复制" {
		t.Fatalf("CopyText 覆盖后 OnCopy=%q，want 实际复制", got2)
	}

	// 无 OnCopy 时也应进入已复制态（不 panic）。
	c3 := NewCopyableText("x")
	c3.doCopy()
	if !c3.Copied() {
		t.Fatal("无回调也应进入已复制态")
	}
}
