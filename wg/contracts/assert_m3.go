package contracts

import (
	"fmt"

	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/input"
)

// 本文件是 M3 到期契约的断言实现。
//
// C28 是"文本与按键分离"。它不像 C27 那样只看状态机，也不像 C1 那样看像素：
// 它约束的是**两个通道不许串线**——
//
//	WM_CHAR / IME 结果 → Text（可打印字符，UTF-8）
//	WM_KEYDOWN / VK   → Code（语义：退格、Tab、回车、快捷键）
//
// 串线的表现极其顽固：退格会变成方块、Ctrl+A 会往输入框插 'a'、
// 中文上屏会留下拼音。这些都不能靠"看起来对了"验收，必须逐字断言。

// ── C28 文本与按键分离，控制字符不进 Text ────────────────────────

func assertTextKeySeparation() error {
	norm := host.NewNormalizer(host.NewDPI(1))
	t0 := m2Clock(0)

	// ① 控制字符走 WM_CHAR 时必须被归一化层丢掉。
	for _, s := range []string{"\b", "\t", "\r", "\x01", "\x7f"} {
		evs := norm.Normalize(host.RawEvent{Kind: host.RawChar, Text: s, Time: t0})
		if len(evs) != 0 {
			return fmt.Errorf("控制字符 %q 进入了 Text 事件", s)
		}
	}
	if norm.Stats().Dropped != 5 {
		return fmt.Errorf("丢弃计数 = %d，应为 5", norm.Stats().Dropped)
	}

	// ② 带 Ctrl/Alt 的字母同理：Ctrl+A 是快捷键，不是文本 'a'。
	if evs := norm.Normalize(host.RawEvent{
		Kind: host.RawChar, Text: "a", Mods: input.ModControl, Time: t0,
	}); len(evs) != 0 {
		return fmt.Errorf("Ctrl+A 产生了文本事件")
	}
	// ③ 反向对照：普通字母与中文必须能进 Text（不能"为了安全全丢"）。
	for _, s := range []string{"a", "中文"} {
		evs := norm.Normalize(host.RawEvent{Kind: host.RawChar, Text: s, Time: t0})
		if len(evs) != 1 {
			return fmt.Errorf("%q 未能进入 Text 事件", s)
		}
		ke, ok := evs[0].(*input.KeyEvent)
		if !ok || ke.Text != s {
			return fmt.Errorf("%q 的 Text 事件内容错误: %#v", s, evs[0])
		}
	}

	// ④ 语义键必须有 Code、绝不带 Text。这里走 VK → Code 的真实映射，
	//    而不是直接构造 Code，才能同时验到平台层。
	for _, c := range []struct {
		vk   uint16
		want input.KeyCode
		name string
	}{
		{0x08, input.CodeBackspace, "Backspace"},
		{0x09, input.CodeTab, "Tab"},
		{0x0D, input.CodeEnter, "Enter"},
		{0x1B, input.CodeEscape, "Esc"},
		{0x41, input.CodeA, "A"},
	} {
		evs := norm.Normalize(host.RawEvent{
			Kind: host.RawKeyDown, Key: host.VKToCode(c.vk), Time: t0,
		})
		if len(evs) != 1 {
			return fmt.Errorf("%s 未产生按键事件", c.name)
		}
		ke := evs[0].(*input.KeyEvent)
		if ke.Code != c.want || ke.Text != "" {
			return fmt.Errorf("%s: Code=%v Text=%q（应只有 Code）", c.name, ke.Code, ke.Text)
		}
	}

	// ⑤ 落到编辑模型：Backspace 按 rune 删（中文不能被删成半个字）。
	e := input.NewTextEditor(false)
	kh := input.NewTextKeyHandler(e)
	e.SetValue(input.EditingValue{Text: "中文ab", Selection: input.RangeAt(4)})
	if !kh.Handle(&input.KeyEvent{Code: input.CodeBackspace}) || e.Value().Text != "中文a" {
		return fmt.Errorf("Backspace 未按 rune 删除: %q", e.Value().Text)
	}

	// ⑥ Ctrl+A 全选后打字：整段被替换，而不是插到光标处。
	if !kh.Handle(&input.KeyEvent{Code: input.CodeA, Mods: input.ModControl}) {
		return fmt.Errorf("Ctrl+A 未被消费")
	}
	if r := e.Value().Selection; r.Start != 0 || r.End != 3 {
		return fmt.Errorf("Ctrl+A 选区 = [%d,%d)", r.Start, r.End)
	}
	if !kh.Handle(&input.KeyEvent{Text: "好"}) || e.Value().Text != "好" {
		return fmt.Errorf("全选后输入未替换: %q", e.Value().Text)
	}

	// ⑦ Enter/Tab 只触发语义动作，不许把控制字符写进文本。
	var acts []input.Action
	e.OnAction = func(a input.Action) { acts = append(acts, a) }
	kh.Handle(&input.KeyEvent{Code: input.CodeEnter})
	kh.Handle(&input.KeyEvent{Code: input.CodeTab})
	if e.Value().Text != "好" {
		return fmt.Errorf("Enter/Tab 污染了文本: %q", e.Value().Text)
	}
	if len(acts) != 2 || acts[0] != input.ActionDone || acts[1] != input.ActionNext {
		return fmt.Errorf("语义动作 = %v，应为 [done next]", acts)
	}

	// ⑧ 中文上屏：组合串必须被汉字顶掉，且不残留拼音。
	e2 := input.NewTextEditor(false)
	br := host.NewIMEBridge(e2)
	br.Handle(host.IMFEvent{Kind: host.IMFStartComposition})
	br.Handle(host.IMFEvent{Kind: host.IMFUpdateComposition, Text: "zhong"})
	br.Handle(host.IMFEvent{Kind: host.IMFResult, Text: "中"})
	if e2.Value().Text != "中" || e2.Value().Composing.Start != 0 {
		return fmt.Errorf("IME 上屏结果 = %q 组合区 = %+v", e2.Value().Text, e2.Value().Composing)
	}
	return nil
}
