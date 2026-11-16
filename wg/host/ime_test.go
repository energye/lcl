package host

import (
	"testing"
	"time"

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
