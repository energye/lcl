package input

import "testing"

// TextKeyHandler 是 C28 的唯一落点。这里的单测不重复 diag 的链路断言，
// 只守"映射表本身"：一旦有人把 Tab 接成插制表符、把 Ctrl+A 接成插 'a'，
// 这里会直接红，不用等整条门禁。

func newKeyEditor(multi bool) (*TextEditor, *TextKeyHandler) {
	e := NewTextEditor(multi)
	return e, NewTextKeyHandler(e)
}

func TestKeyHandlerTextGoesToEditor(t *testing.T) {
	e, h := newKeyEditor(false)
	if !h.Handle(&KeyEvent{Text: "中"}) || e.Value().Text != "中" {
		t.Fatalf("文本未进入编辑模型: %q", e.Value().Text)
	}
}

func TestKeyHandlerControlCharsNeverEnterText(t *testing.T) {
	e, h := newKeyEditor(false)
	for _, s := range []string{"\b", "\t", "\r", "\x01"} {
		if !h.Handle(&KeyEvent{Text: s}) {
			t.Fatalf("%q 未被消费", s)
		}
	}
	if e.Value().Text != "" {
		t.Fatalf("控制字符进入了文本: %q", e.Value().Text)
	}
}

func TestKeyHandlerCtrlASelectsAll(t *testing.T) {
	e, h := newKeyEditor(false)
	e.SetValue(EditingValue{Text: "abcde", Selection: RangeAt(2)})
	if !h.Handle(&KeyEvent{Code: CodeA, Mods: ModControl}) {
		t.Fatal("Ctrl+A 未消费")
	}
	if r := e.Value().Selection; r.Start != 0 || r.End != 5 {
		t.Fatalf("选区 = %+v", r)
	}
	if e.Value().Text != "abcde" {
		t.Fatalf("Ctrl+A 改写了文本: %q", e.Value().Text)
	}
}

func TestKeyHandlerCtrlAltStaysSemantic(t *testing.T) {
	e, h := newKeyEditor(false)
	if h.Handle(&KeyEvent{Code: CodeS, Mods: ModAlt}) {
		t.Fatal("未实现的 Alt 组合不该被悄悄消费")
	}
	if e.Value().Text != "" {
		t.Fatal("快捷键污染了文本")
	}
}

func TestKeyHandlerBackspaceDeletesRune(t *testing.T) {
	e, h := newKeyEditor(false)
	e.SetValue(EditingValue{Text: "你中", Selection: RangeAt(2)})
	h.Handle(&KeyEvent{Code: CodeBackspace})
	if e.Value().Text != "你" {
		t.Fatalf("退格结果 = %q", e.Value().Text)
	}
}

func TestKeyHandlerEnterIsActionNotText(t *testing.T) {
	e, h := newKeyEditor(false)
	var got []Action
	e.OnAction = func(a Action) { got = append(got, a) }
	h.Handle(&KeyEvent{Code: CodeEnter})
	if e.Value().Text != "" {
		t.Fatalf("单行回车插入了文本: %q", e.Value().Text)
	}
	if len(got) != 1 || got[0] != ActionDone {
		t.Fatalf("动作 = %v", got)
	}
}

func TestKeyHandlerTabIsFocusNotText(t *testing.T) {
	e, h := newKeyEditor(true)
	var got []Action
	e.OnAction = func(a Action) { got = append(got, a) }
	h.Handle(&KeyEvent{Code: CodeTab})
	if e.Value().Text != "" {
		t.Fatalf("多行框的 Tab 也不该写字: %q", e.Value().Text)
	}
	if len(got) != 1 || got[0] != ActionNext {
		t.Fatalf("动作 = %v", got)
	}
}

func TestKeyHandlerOnActionCanTakeOver(t *testing.T) {
	e, h := newKeyEditor(true)
	h.OnAction = func(a Action, _ *KeyEvent) bool { return true } // 外层已处理
	var got []Action
	e.OnAction = func(a Action) { got = append(got, a) }
	h.Handle(&KeyEvent{Code: CodeEnter})
	if e.Value().Text != "" || len(got) != 0 {
		t.Fatalf("外层接管后仍执行了默认动作: %q %v", e.Value().Text, got)
	}
}

func TestKeyHandlerCtrlWordMovement(t *testing.T) {
	e, h := newKeyEditor(false)
	// "你好 world" 共 8 个 rune：末尾光标是 8，词 w 的起点是 3。
	e.SetValue(EditingValue{Text: "你好 world", Selection: RangeAt(8)})
	h.Handle(&KeyEvent{Code: CodeLeft, Mods: ModControl})
	if c := e.Value().Cursor(); c != 3 {
		t.Fatalf("Ctrl+← 光标 = %d，应为 3", c)
	}
	h.Handle(&KeyEvent{Code: CodeRight, Mods: ModControl})
	if c := e.Value().Cursor(); c != 8 {
		t.Fatalf("Ctrl+→ 光标 = %d，应为 8", c)
	}
}
