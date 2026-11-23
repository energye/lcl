
package widgets

import (
	"testing"
)

// TestSelectMaxTagCount 折叠：MaxTagCount=2 截断显示前 2 个+overflow 计数，
// 0 不限制；tagShown 语义与 TreeSelect 一致。
func TestSelectMaxTagCount(t *testing.T) {
	s := NewSelect([]string{"渲染", "布局", "输入", "浮层", "主题"})
	s.Multiple = true
	s.Values = []string{"渲染", "输入", "浮层", "主题"}
	s.MaxTagCount = 2
	shown, overflow := s.tagShown()
	if len(shown) != 2 || shown[0] != "渲染" || shown[1] != "输入" {
		t.Fatalf("MaxTagCount=2 应显示前 2 个，got %v", shown)
	}
	if overflow != 2 {
		t.Fatalf("折叠数应为 2，got %d", overflow)
	}
	s.MaxTagCount = 0
	shown, overflow = s.tagShown()
	if len(shown) != 4 || overflow != 0 {
		t.Fatalf("MaxTagCount=0 不限制，got %v overflow %d", shown, overflow)
	}
	// 恰好等于 MaxTagCount 时无折叠。
	s.MaxTagCount = 4
	shown, overflow = s.tagShown()
	if overflow != 0 {
		t.Fatalf("值数=MaxTagCount 不折叠，got overflow %d", overflow)
	}
}
