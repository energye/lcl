package diag

import (
	"errors"
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
)

// ── --keyselftest（C28）─────────────────────────────────────────────
//
// 覆盖"平台消息 → 归一化 → 分发 → 文本编辑"整条链路。
// 断言对象不是"按键有没有送到"，而是**两个通道有没有串线**：
//
//	Text 通道只装可打印字符（含中文）
//	Code 通道只装语义（退格 / Tab / 回车 / 快捷键）
//
// 这两条线一旦串了，症状都是"偶尔多一个字/少一个字"，肉眼极难复现。

// KeyCase 是单条断言的结果。
type KeyCase struct {
	Name string
	OK   bool
	Got  string
	Note string
}

// KeyResult 是 --keyselftest 的汇总。
type KeyResult struct {
	Cases   int
	Failed  int
	Elapsed time.Duration
	Details []KeyCase
}

// OK 判定是否全绿（门禁要求 ≥12 项）。
func (r KeyResult) OK() bool { return r.Failed == 0 && r.Cases >= 12 }

func (r KeyResult) String() string {
	return fmt.Sprintf("keyselftest: %d 项，失败 %d，用时 %v", r.Cases, r.Failed, r.Elapsed)
}

var kBase = time.Date(2026, 11, 16, 10, 0, 0, 0, time.UTC)

func kAt(ms int) time.Time { return kBase.Add(time.Duration(ms) * time.Millisecond) }

const (
	kKeyDn  = 0x0100
	kKeyUp  = 0x0101
	kChar   = 0x0102
	kSysDn  = 0x0104
	vkBack  = 0x08
	vkTab   = 0x09
	vkEnter = 0x0D
	vkEsc   = 0x1B
	vkA     = 0x41
)

// keyRig 是"路由 + 分发 + 编辑模型"的最小闭环。
type keyRig struct {
	norm   *host.Normalizer
	router *host.Router
	sink   *keySink
	disp   *input.Dispatcher
	editor *input.TextEditor
	keys   *input.TextKeyHandler
	ime    *host.IMEBridge
	acts   []input.Action
}

type keySink struct {
	disp   *input.Dispatcher
	events []input.Event
	keys   []*input.KeyEvent
	frames int
}

func (s *keySink) PushEvent(ev input.Event) bool {
	s.events = append(s.events, ev)
	switch e := ev.(type) {
	case *input.KeyEvent:
		s.keys = append(s.keys, e)
		return s.disp.Key(e)
	case *input.IMEEvent:
		return s.disp.IME(e)
	}
	return false
}

func (s *keySink) RequestFrame() { s.frames++ }

func newKeyRig(multiline bool) *keyRig {
	btn := scene.Box(scene.BoxSpec{Fill: render.Fill(geom.RGBA(200, 200, 200))}).
		SetSize2(80, 80).SetStackPosition(geom.Pt(10, 10)).
		SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "field"})
	root := scene.Stack(btn).SetSize2(100, 100)
	root.Layout(scene.Unbounded())

	r := &keyRig{}
	r.norm = host.NewNormalizer(host.NewDPI(1))
	r.disp = input.NewDispatcher(root)
	r.disp.SetFocus(btn)
	r.sink = &keySink{disp: r.disp}
	r.router = host.NewRouter(r.norm, r.sink)
	r.editor = input.NewTextEditor(multiline)
	r.keys = input.NewTextKeyHandler(r.editor)
	r.ime = host.NewIMEBridge(r.editor)
	r.editor.OnAction = func(a input.Action) { r.acts = append(r.acts, a) }
	r.disp.Set(btn, input.Handler{
		Key: func(e *input.KeyEvent) bool { return r.keys.Handle(e) },
		IME: func(e *input.IMEEvent) bool { return true },
	})
	return r
}

func (r *keyRig) text(s string, ms int) *keyRig {
	r.router.Handle(host.Message{Msg: kChar, WParam: uintptr([]rune(s)[0]), Time: kAt(ms)})
	return r
}

func (r *keyRig) key(msg uint32, vk int, ms int) *keyRig {
	r.router.Handle(host.Message{Msg: msg, WParam: uintptr(vk), Time: kAt(ms)})
	return r
}

func (r *keyRig) mod(m input.Modifiers) *keyRig { r.router.SetMods(m); return r }

// KeyTest 跑文本/按键分离门禁。
func KeyTest() (KeyResult, error) {
	began := time.Now()
	var res KeyResult
	cases := []struct {
		name string
		run  func() (string, error)
	}{
		{"ctrl-a-no-text", caseKeyCtrlA},
		{"ctrl-alt-drop-char", caseKeyCtrlAltChar},
		{"backspace-code-only", caseKeyBackspaceCode},
		{"enter-tab-code-only", caseKeyEnterTabCode},
		{"control-chars-dropped", caseKeyControlChars},
		{"utf8-char-passes", caseKeyUTF8Passes},
		{"backspace-deletes-rune", caseKeyBackspaceRune},
		{"delete-forward-rune", caseKeyDeleteRune},
		{"select-all-replace", caseKeySelectAllReplace},
		{"enter-single-vs-multi", caseKeyEnterModes},
		{"tab-never-inserts", caseKeyTabNoInsert},
		{"ime-chinese-commit", caseKeyIMEChinese},
		{"ime-cancel-no-pinyin", caseKeyIMECancel},
		{"paste-filters-control", caseKeyPasteFilter},
	}
	for _, c := range cases {
		got, err := c.run()
		d := KeyCase{Name: c.name, OK: err == nil, Got: got}
		if err != nil {
			d.Note = err.Error()
			res.Failed++
		}
		res.Cases++
		res.Details = append(res.Details, d)
	}
	res.Elapsed = time.Since(began)
	return res, nil
}

func caseKeyCtrlA() (string, error) {
	r := newKeyRig(false).text("a", 0).text("b", 16).mod(input.ModControl).key(kKeyDn, vkA, 32)
	keys := r.sink.keys
	last := keys[len(keys)-1]
	got := fmt.Sprintf("events=%d code=%v text=%q sel=[%d,%d)",
		len(keys), last.Code, last.Text, r.editor.Value().Selection.Start, r.editor.Value().Selection.End)
	if len(keys) != 3 || last.Text != "" || last.Code != input.CodeA {
		return got, errors.New("Ctrl+A 不应产生文本事件，只应有 CodeA")
	}
	if r.editor.Value().Text != "ab" || r.editor.Value().Selection.End != 2 {
		return got, errors.New("Ctrl+A 未全选，或文本被污染")
	}
	return got, nil
}

func caseKeyCtrlAltChar() (string, error) {
	r := newKeyRig(false).mod(input.ModControl | input.ModAlt)
	r.router.Handle(host.Message{Msg: kChar, WParam: uintptr('s'), Time: kAt(0)})
	got := fmt.Sprintf("events=%d text=%q", len(r.sink.events), r.editor.Value().Text)
	if len(r.sink.events) != 0 || r.editor.Value().Text != "" {
		return got, errors.New("Ctrl+Alt+S 不该进 Text 通道")
	}
	return got, nil
}

func caseKeyBackspaceCode() (string, error) {
	r := newKeyRig(false).text("a", 0).key(kKeyDn, vkBack, 16)
	ke := r.sink.keys[len(r.sink.keys)-1]
	got := fmt.Sprintf("code=%v text=%q editor=%q", ke.Code, ke.Text, r.editor.Value().Text)
	if ke.Code != input.CodeBackspace || ke.Text != "" {
		return got, errors.New("Backspace 必须只有 Code，没有 Text")
	}
	if r.editor.Value().Text != "" {
		return got, errors.New("Backspace 未删除字符")
	}
	return got, nil
}

func caseKeyEnterTabCode() (string, error) {
	r := newKeyRig(false).text("a", 0).key(kKeyDn, vkEnter, 16).key(kKeyDn, vkTab, 32)
	got := ""
	for _, ke := range r.sink.keys[1:] {
		got += fmt.Sprintf("%v/%q ", ke.Code, ke.Text)
		if ke.Text != "" {
			return got, errors.New("Enter/Tab 携带了 Text")
		}
	}
	if r.sink.keys[1].Code != input.CodeEnter || r.sink.keys[2].Code != input.CodeTab {
		return got, errors.New("Enter/Tab 的 Code 错误")
	}
	if r.editor.Value().Text != "a" {
		return got, errors.New("语义键改写了文本")
	}
	return got, nil
}

func caseKeyControlChars() (string, error) {
	r := newKeyRig(false)
	for _, s := range []rune{'\b', '\t', '\r', '\x01', '\x7f'} {
		r.router.Handle(host.Message{Msg: kChar, WParam: uintptr(s), Time: kAt(0)})
	}
	got := fmt.Sprintf("dropped=%d text=%q", r.norm.Stats().Dropped, r.editor.Value().Text)
	if r.norm.Stats().Dropped != 5 || r.editor.Value().Text != "" {
		return got, errors.New("控制字符进入了文本")
	}
	return got, nil
}

func caseKeyUTF8Passes() (string, error) {
	r := newKeyRig(false)
	// BMP 汉字在 WM_CHAR 里是一个 UTF-16 单元；逐字发送才符合真实消息序列。
	for i, s := range []rune{'a', '中', '文'} {
		r.router.Handle(host.Message{Msg: kChar, WParam: uintptr(s), Time: kAt(i * 16)})
	}
	got := fmt.Sprintf("text=%q events=%d", r.editor.Value().Text, len(r.sink.keys))
	if r.editor.Value().Text != "a中文" {
		return got, errors.New("UTF-8 字符未能上屏")
	}
	return got, nil
}

func caseKeyBackspaceRune() (string, error) {
	r := newKeyRig(false)
	r.editor.SetValue(input.EditingValue{Text: "你中", Selection: input.RangeAt(2)})
	r.keys.Handle(&input.KeyEvent{Code: input.CodeBackspace})
	got := fmt.Sprintf("text=%q", r.editor.Value().Text)
	if r.editor.Value().Text != "你" {
		return got, errors.New("退格按字节删会留下半个汉字")
	}
	return got, nil
}

func caseKeyDeleteRune() (string, error) {
	r := newKeyRig(false)
	r.editor.SetValue(input.EditingValue{Text: "你中好", Selection: input.RangeAt(1)})
	r.keys.Handle(&input.KeyEvent{Code: input.CodeDelete})
	got := fmt.Sprintf("text=%q", r.editor.Value().Text)
	if r.editor.Value().Text != "你好" {
		return got, errors.New("Delete 未按 rune 删除")
	}
	return got, nil
}

func caseKeySelectAllReplace() (string, error) {
	r := newKeyRig(false).text("a", 0).text("b", 16).mod(input.ModControl).key(kKeyDn, vkA, 32)
	r.mod(0)
	r.ime.Handle(host.IMFEvent{Kind: host.IMFStartComposition, Time: kAt(48)})
	r.ime.Handle(host.IMFEvent{Kind: host.IMFUpdateComposition, Text: "hao", Time: kAt(64)})
	r.ime.Handle(host.IMFEvent{Kind: host.IMFResult, Text: "好", Time: kAt(80)})
	got := fmt.Sprintf("text=%q", r.editor.Value().Text)
	if r.editor.Value().Text != "好" {
		return got, errors.New("全选后 IME 上屏未替换整段")
	}
	return got, nil
}

func caseKeyEnterModes() (string, error) {
	single := newKeyRig(false)
	single.keys.Handle(&input.KeyEvent{Code: input.CodeEnter})
	multi := newKeyRig(true)
	multi.keys.Handle(&input.KeyEvent{Code: input.CodeEnter})
	got := fmt.Sprintf("single=%q multi=%q acts=%v/%v",
		single.editor.Value().Text, multi.editor.Value().Text, single.acts, multi.acts)
	if single.editor.Value().Text != "" || len(single.acts) != 1 || single.acts[0] != input.ActionDone {
		return got, errors.New("单行 Enter 应提交，不插换行")
	}
	// 多行 Enter 走编辑模型的换行分支（不产生语义动作回调），这是刻意的：
	// 换行是文本编辑，不该顺带触发"提交表单"。
	if multi.editor.Value().Text != "\n" || len(multi.acts) != 0 {
		return got, errors.New("多行 Enter 应只换行，不提交")
	}
	return got, nil
}

func caseKeyTabNoInsert() (string, error) {
	r := newKeyRig(true)
	r.keys.Handle(&input.KeyEvent{Code: input.CodeTab})
	got := fmt.Sprintf("text=%q actions=%v", r.editor.Value().Text, r.acts)
	if r.editor.Value().Text != "" {
		return got, errors.New("Tab 不得写入文本")
	}
	if len(r.acts) != 1 || r.acts[0] != input.ActionNext {
		return got, errors.New("Tab 应触发 ActionNext（焦点导航）")
	}
	return got, nil
}

func caseKeyIMEChinese() (string, error) {
	r := newKeyRig(false).text("你", 0)
	r.ime.Handle(host.IMFEvent{Kind: host.IMFStartComposition, Time: kAt(16)})
	r.ime.Handle(host.IMFEvent{Kind: host.IMFUpdateComposition, Text: "zhong", Cursor: 5, Time: kAt(32)})
	mid := r.editor.Value().Text
	r.ime.Handle(host.IMFEvent{Kind: host.IMFResult, Text: "中", Time: kAt(48)})
	r.ime.Handle(host.IMFEvent{Kind: host.IMFEndComposition, Time: kAt(64)})
	got := fmt.Sprintf("during=%q after=%q", mid, r.editor.Value().Text)
	if mid != "你zhong" {
		return got, errors.New("组合串未进入预编辑区")
	}
	if r.editor.Value().Text != "你中" {
		return got, errors.New("上屏结果不精确（残留拼音或多字）")
	}
	return got, nil
}

func caseKeyIMECancel() (string, error) {
	r := newKeyRig(false).text("你", 0)
	r.ime.Handle(host.IMFEvent{Kind: host.IMFStartComposition, Time: kAt(16)})
	r.ime.Handle(host.IMFEvent{Kind: host.IMFUpdateComposition, Text: "zh", Cursor: 2, Time: kAt(32)})
	r.ime.Handle(host.IMFEvent{Kind: host.IMFEndComposition, Time: kAt(48)})
	got := fmt.Sprintf("text=%q cancels=%d", r.editor.Value().Text, r.ime.Stats().Cancels)
	if r.editor.Value().Text != "你" {
		return got, errors.New("取消组合后残留拼音")
	}
	return got, nil
}

func caseKeyPasteFilter() (string, error) {
	r := newKeyRig(false)
	r.editor.InsertText("你\x01好\r\n")
	got := fmt.Sprintf("text=%q", r.editor.Value().Text)
	if r.editor.Value().Text != "你好" {
		return got, errors.New("粘贴内容中的控制字符未被过滤")
	}
	return got, nil
}
