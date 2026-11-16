package input

// TextKeyHandler 把引擎按键翻译成编辑动作。
//
// 为什么需要这一层：C28 要求"控制字符不进 Text、语义键只由 Code 处理"。
// 如果每个输入框自己写 switch，第一版都会漏掉一种组合——
// 典型的是 Ctrl+A 先被当成字符 'a' 插进文本，或者 Tab 在输入框里变成空格。
// 把映射收在一处，组件只接一个 Handle，契约也只需要断言这一处。
//
// 它不碰 TextEditor 以外的东西：焦点切换、提交表单这些"语义键的外延"
// 通过 OnAction 回调交给外层，编辑内核因此不需要认识 FocusManager。
type TextKeyHandler struct {
	editor *TextEditor
	// OnAction 在 Enter/Tab 这类语义动作上被调用。
	// 返回 true 表示已消费（例如焦点管理者切了字段）；返回 false 时
	// 编辑器仍会执行默认编辑动作（多行 Enter 的换行）。
	OnAction func(a Action, ev *KeyEvent) bool
	// OnUnhandled 收到本处理器不认识的键时调用，给滚动/快捷键层兜底。
	OnUnhandled func(ev *KeyEvent) bool
}

// NewTextKeyHandler 创建处理器。
func NewTextKeyHandler(e *TextEditor) *TextKeyHandler {
	return &TextKeyHandler{editor: e}
}

// Editor 返回绑定的编辑模型。
func (h *TextKeyHandler) Editor() *TextEditor { return h.editor }

// Handle 处理一条按键；返回是否消费。
//
// 判定顺序是刻意的：**先看 Text，再看 Code**。
// 一条 KeyEvent 不该同时携带可打印文本和语义键；若平台胶水层误送了
// （例如 IME 组合期间混进 WM_CHAR），以文本为准，避免删掉刚上屏的字。
func (h *TextKeyHandler) Handle(ev *KeyEvent) bool {
	if ev == nil || h.editor == nil {
		return false
	}
	if ev.Text != "" {
		// 归一化层已过滤控制字符；这里再过一遍编辑模型（多行策略），
		// 让"直接构造 KeyEvent"的调用方也无法把 \r 塞进单行框。
		if h.editor.filter(ev.Text) == "" {
			return true // 控制字符被吞掉，但事件算已处理
		}
		h.editor.InsertText(ev.Text)
		return true
	}

	// Ctrl/Alt 组合只有语义，没有文本。目前内核只实现全选；
	// 其余组合留给注册在前面的快捷键处理器（复制/粘贴/撤销）。
	//
	// 编辑键必须先放行：Ctrl+←/→ 是按词移动，Ctrl+Backspace 是删词，
	// 它们落在"编辑"而不是"快捷键"这一侧。若在这里一刀切拦截，
	// 用户按 Ctrl+← 会毫无反应——而这正是最容易漏测的组合。
	if (ev.Mods.Has(ModControl) || ev.Mods.Has(ModAlt)) && !isEditingKey(ev.Code) {
		if ev.Code == CodeA && ev.Mods.Has(ModControl) {
			h.editor.SelectAll()
			return true
		}
		return h.unhandled(ev)
	}

	switch ev.Code {
	case CodeBackspace:
		h.editor.DeleteBackward()
		return true
	case CodeDelete:
		h.editor.DeleteForward()
		return true
	case CodeLeft:
		h.moveWord(ev, -1)
		return true
	case CodeRight:
		h.moveWord(ev, 1)
		return true
	case CodeHome:
		h.editor.SetSelection(RangeAt(0))
		return true
	case CodeEnd:
		h.editor.SetSelection(RangeAt(h.editor.Value().Len()))
		return true
	case CodeEnter:
		// 单行框的 PerformAction 会把 Newline 降级成 Done（不插换行）。
		h.action(ActionNewline, ev)
		return true
	case CodeTab:
		// Tab 永远是焦点导航，不在文本里插入制表符——
		// 连多行框也不插：缩进有专门的命令，直接插 \t 会和输入法打架。
		h.action(ActionNext, ev)
		return true
	case CodeEscape:
		h.editor.CancelComposing()
		h.action(ActionDone, ev)
		return true
	}
	return h.unhandled(ev)
}

// moveWord 按 Ctrl 时跨一个词，否则逐 rune。
//
// 中文没有空格分词，按"连续非空白算一个词"处理：Ctrl+← 在
// "你好 world" 里从末尾先跳到 world 起点，再跳到开头，符合直觉。
func (h *TextKeyHandler) moveWord(ev *KeyEvent, dir int32) {
	if !ev.Mods.Has(ModControl) {
		h.editor.MoveCursor(dir)
		return
	}
	s := []rune(h.editor.Value().Text)
	c := int(h.editor.Value().Cursor())
	if dir < 0 {
		for c > 0 && isSpace(s[c-1]) {
			c--
		}
		for c > 0 && !isSpace(s[c-1]) {
			c--
		}
	} else {
		for c < len(s) && isSpace(s[c]) {
			c++
		}
		for c < len(s) && !isSpace(s[c]) {
			c++
		}
	}
	h.editor.SetSelection(RangeAt(int32(c)))
}

func (h *TextKeyHandler) action(a Action, ev *KeyEvent) {
	if h.OnAction != nil && h.OnAction(a, ev) {
		return
	}
	h.editor.PerformAction(a)
}

func (h *TextKeyHandler) unhandled(ev *KeyEvent) bool {
	return h.OnUnhandled != nil && h.OnUnhandled(ev)
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// isEditingKey 列出"带修饰键也仍归编辑模型管"的键。
func isEditingKey(c KeyCode) bool {
	switch c {
	case CodeBackspace, CodeDelete, CodeLeft, CodeRight, CodeHome, CodeEnd,
		CodeEnter, CodeTab, CodeEscape:
		return true
	}
	return false
}
