package widgets

import "testing"

// TestEditableText 编辑：进入/提交/取消三态。
func TestEditableText(t *testing.T) {
	e := NewEditableText("原始")
	var committed, cancelled string
	e.OnChange = func(s string) { committed = s }
	e.OnCancel = func() { cancelled = "yes" }

	if e.Editing() {
		t.Fatal("初始不应处于编辑态")
	}
	e.BeginEdit()
	if !e.Editing() {
		t.Fatal("BeginEdit 后应处于编辑态")
	}
	// 提交（模拟编辑框改了值）。
	e.editor = NewInput("")
	e.editor.Value = "新值"
	e.Commit()
	if e.Editing() {
		t.Fatal("Commit 后应退出编辑态")
	}
	if e.Body != "新值" || committed != "新值" {
		t.Fatalf("提交后 Body=%q committed=%q", e.Body, committed)
	}

	// 取消：值不变、OnCancel 触发。
	e.BeginEdit()
	e.Cancel()
	if e.Editing() || cancelled != "yes" {
		t.Fatalf("取消后 editing=%v cancelled=%q", e.Editing(), cancelled)
	}
	if e.Body != "新值" {
		t.Fatalf("取消不应改 Body，got %q", e.Body)
	}
}
