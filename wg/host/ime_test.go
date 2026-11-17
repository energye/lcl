package host

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
)

// IME 桥的单测只覆盖两套"看起来能过、实际会丢字"的序列：
// 选区上屏（选中文本后拼音输入）与取消回滚（Esc 后原文还在）。

func TestIMEReplacesSelectionOnCommit(t *testing.T) {
	e := input.NewTextEditor(false)
	e.SetValue(input.EditingValue{Text: "ab", Selection: input.RangeOf(0, 2)})
	b := NewIMEBridge(e)
	b.Handle(IMFEvent{Kind: IMFStartComposition})
	b.Handle(IMFEvent{Kind: IMFUpdateComposition, Text: "hao", Cursor: 3})
	if got := e.Value().Text; got != "hao" {
		t.Fatalf("组合串 = %q，应覆盖整个选区", got)
	}
	b.Handle(IMFEvent{Kind: IMFResult, Text: "好"})
	b.Handle(IMFEvent{Kind: IMFEndComposition})
	if got := e.Value().Text; got != "好" {
		t.Fatalf("上屏结果 = %q（不能残留拼音或旧选区）", got)
	}
}

func TestIMECancelRestoresSelection(t *testing.T) {
	e := input.NewTextEditor(false)
	e.SetValue(input.EditingValue{Text: "前缀选中后缀", Selection: input.RangeOf(2, 4)})
	b := NewIMEBridge(e)
	b.Handle(IMFEvent{Kind: IMFStartComposition})
	b.Handle(IMFEvent{Kind: IMFUpdateComposition, Text: "zh", Cursor: 2})
	if got := e.Value().Text; got != "前缀zh后缀" {
		t.Fatalf("组合中 = %q", got)
	}
	b.Handle(IMFEvent{Kind: IMFEndComposition})
	if got := e.Value().Text; got != "前缀选中后缀" {
		t.Fatalf("取消后 = %q，原文应还回来", got)
	}
}

func TestIMEWithoutStartStillReplacesSelection(t *testing.T) {
	e := input.NewTextEditor(false)
	e.SetValue(input.EditingValue{Text: "abc", Selection: input.RangeOf(1, 3)})
	b := NewIMEBridge(e)
	// 有些输入法不发 START，直接送 UPDATE。
	b.Handle(IMFEvent{Kind: IMFUpdateComposition, Text: "中", Cursor: 1, Time: time.Now()})
	b.Handle(IMFEvent{Kind: IMFResult, Text: "中"})
	if got := e.Value().Text; got != "a中" {
		t.Fatalf("无 START 的上屏结果 = %q", got)
	}
	if b.Stats().Starts != 1 || b.Stats().Commits != 1 {
		t.Fatalf("统计异常: %+v", b.Stats())
	}
}

// resultIME 是"只发上屏结果"的平台 IME 替身。
type resultIME struct{ b *IMEBridge }

func (s *resultIME) Attach() bool { return true }
func (s *resultIME) Detach()      {}
func (s *resultIME) HandleMessage(msg uint32, _, _ uintptr) ([]input.Event, bool) {
	if msg == wmIMEComposition {
		return s.b.Handle(IMFEvent{Kind: IMFResult, Text: "你好"}), true
	}
	return nil, false
}
func (s *resultIME) SetCandidates(geom.Rect) {}

// TestIMECommitCharsNotInsertedTwice 锁定"上屏结果 + 系统补发 WM_CHAR"不双写。
//
// 真机微软拼音实测：透传 WM_IME_COMPOSITION 后，DefWindowProc 会为结果串
// 逐字补发 WM_CHAR；结果串已由桥写入，补发必须吞掉，否则 "你好" 变成
// "你好你好"。
func TestIMECommitCharsNotInsertedTwice(t *testing.T) {
	e := input.NewTextEditor(false)
	b := NewIMEBridge(e)
	sink := &sinkStub{}
	r := NewRouter(NewNormalizer(NewDPI(1)), sink)
	r.SetIME(&resultIME{b: b}, b)

	if r.Handle(Message{Msg: wmIMEComposition, Time: time.Now()}) {
		t.Fatal("IME 消息必须交回 DefWindowProc")
	}
	if got := e.Value().Text; got != "你好" {
		t.Fatalf("上屏结果 = %q", got)
	}
	if !r.Handle(Message{Msg: wmChar, WParam: uintptr('你'), Time: time.Now()}) {
		t.Fatal("补发的 WM_CHAR('你') 应被吞掉")
	}
	if !r.Handle(Message{Msg: wmChar, WParam: uintptr('好'), Time: time.Now()}) {
		t.Fatal("补发的 WM_CHAR('好') 应被吞掉")
	}
	if sink.pushed != 1 {
		t.Fatalf("补发字符不应产生事件，pushed=%d", sink.pushed)
	}
	// 用户的新字符不能被误吞（不匹配即放行；Handler 对普通按键恒返回 true，
	// 用是否产生事件判断）。
	r.Handle(Message{Msg: wmChar, WParam: uintptr('x'), Time: time.Now()})
	if got := e.Value().Text; got != "你好" {
		t.Fatalf("编辑器被二次写入: %q", got)
	}
	if sink.pushed != 2 { // 1 个上屏事件 + 1 个 'x' 字符按键
		t.Fatalf("事件数 = %d", sink.pushed)
	}

	// 新按键应清掉可能残留的抑制队列。
	e2 := input.NewTextEditor(false)
	b2 := NewIMEBridge(e2)
	b2.Handle(IMFEvent{Kind: IMFResult, Text: "ab"})
	r2 := NewRouter(NewNormalizer(NewDPI(1)), &sinkStub{})
	r2.SetIME(&resultIME{b: b2}, b2)
	r2.Handle(Message{Msg: wmKeyDown, WParam: 0x41, Time: time.Now()})
	if r2.bridge.TakeCommitChar('a') {
		t.Fatal("新按键后不应再吞字符")
	}
}

type imeStub struct {
	rect       geom.Rect
	positioned int
}

func (s *imeStub) Attach() bool { return true }
func (s *imeStub) Detach()      {}
func (s *imeStub) HandleMessage(msg uint32, _, _ uintptr) ([]input.Event, bool) {
	if msg == wmIMEStartComposition {
		return nil, true
	}
	return nil, false
}
func (s *imeStub) SetCandidates(r geom.Rect) {
	s.rect, s.positioned = r, s.positioned+1
}

type sinkStub struct {
	pushed int
	frames int
}

func (s *sinkStub) PushEvent(input.Event) bool { s.pushed++; return false }
func (s *sinkStub) RequestFrame()              { s.frames++ }

// TestIMECompositionMessagesPassThrough 锁定"IME 消息不能被吞"。
//
// DefWindowProc 负责组合窗/候选窗的默认处理；吞掉消息后输入法只能把
// 候选框留在默认位置（真窗表现为候选框跑到窗口角落、不跟光标走）。
func TestIMECompositionMessagesPassThrough(t *testing.T) {
	r := NewRouter(NewNormalizer(NewDPI(1)), &sinkStub{})
	ime := &imeStub{}
	r.SetIME(ime, NewIMEBridge(nil))

	if r.Handle(Message{Msg: wmIMEStartComposition, Time: time.Now()}) {
		t.Fatal("IME 组合消息必须交回 DefWindowProc（Handle 应返回 false）")
	}
	if r.Stats().Handled != 1 {
		t.Fatalf("IME 消息应记为已处理: %+v", r.Stats())
	}

	r.SetIMECaret(geom.RectXYWH(10, 20, 2, 16))
	if ime.positioned != 1 || ime.rect.Min.X != 10 || ime.rect.Min.Y != 20 {
		t.Fatalf("光标锚点未送达平台层: %+v", ime)
	}
}
