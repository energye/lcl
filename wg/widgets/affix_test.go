package widgets

import (
	"testing"
)

// TestAffixOffsetBottom 双阈值判定：offsetBottom 吸底、双阈值独立启用
// （未注入的通道不误触发——antd offsetTop/offsetBottom undefined 语义）、
// 顶部超阈值仍优先固定。
func TestAffixOffsetBottom(t *testing.T) {
	// 吸底卡：只注入 BottomOffset（ScrollOffset=nil）——顶部判定不误触发。
	bottom := float32(1000)
	a := NewAffix(nil)
	a.OffsetBottom = 120
	a.BottomOffset = func() float32 { return bottom }
	n := a.Build(testCtx())
	_ = n
	if a.Affixed() {
		t.Fatal("距底 1000>120 不该吸底")
	}
	bottom = 80
	a.Build(testCtx())
	if !a.Affixed() {
		t.Fatal("距底 80≤120 应吸底")
	}
	// 只注入 ScrollOffset（BottomOffset=nil）——底部判定不误触发。
	scroll := float32(0)
	b := NewAffix(nil)
	b.OffsetTop = 100
	b.ScrollOffset = func() float32 { return scroll }
	b.Build(testCtx())
	if b.Affixed() {
		t.Fatal("offset 0<100 不该固定")
	}
	scroll = 150
	b.Build(testCtx())
	if !b.Affixed() {
		t.Fatal("offset 150≥100 应固定")
	}
	// 双阈值并存：底部吸底后顶部无关。
	c := NewAffix(nil)
	c.OffsetTop = 100
	c.OffsetBottom = 50
	scroll2, bottom2 := float32(0), float32(30)
	c.ScrollOffset = func() float32 { return scroll2 }
	c.BottomOffset = func() float32 { return bottom2 }
	c.Build(testCtx())
	if !c.Affixed() {
		t.Fatal("距底 30≤50 应吸底（顶部 offset 0<100 不触发）")
	}
	// OffsetBottom=0 不启用底部判定（即使 BottomOffset 注入）。
	d := NewAffix(nil)
	bottom3 := float32(10)
	d.BottomOffset = func() float32 { return bottom3 }
	d.Build(testCtx())
	if d.Affixed() {
		t.Fatal("OffsetBottom=0 不启用底部判定")
	}
}
