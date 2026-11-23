package widgets

import (
	"testing"
)

// TestTreeSelectMultiToggle 多选 toggle：点选加入、再点移出（antd 多选
// 点选不关面板语义——回调拿到完整数组）。
func TestTreeSelectMultiToggle(t *testing.T) {
	s := NewTreeSelect([]TreeNode{
		{Key: "comp", Label: "组件", Children: []TreeNode{
			{Key: "input", Label: "输入框"},
			{Key: "select", Label: "选择器"},
		}},
		{Key: "nav", Label: "导航", Children: []TreeNode{
			{Key: "top", Label: "顶部导航"},
		}},
	})
	s.Multiple = true
	s.tree = NewTree(s.Nodes) // 真流程 openPanel 先建 tree（OnSelect 只在浮层内触发）
	var got [][]string
	s.OnValuesChange = func(keys []string) { got = append(got, append([]string(nil), keys...)) }

	s.toggleValue("input", nil)
	if len(s.Values) != 1 || s.Values[0] != "input" {
		t.Fatalf("首次点选应加入 input，got %v", s.Values)
	}
	if len(got) != 1 || len(got[0]) != 1 || got[0][0] != "input" {
		t.Fatalf("首次回调应得 [input]，got %v", got)
	}
	// syncTreeSelection 应把值同步进内嵌 Tree 的 SelectedSet。
	if s.tree == nil || !s.tree.SelectedSet["input"] {
		t.Fatalf("syncTreeSelection 应写入 tree.SelectedSet")
	}
	s.toggleValue("top", nil)
	if len(s.Values) != 2 {
		t.Fatalf("第二次点选应加入 top，got %v", s.Values)
	}
	s.toggleValue("input", nil)
	if len(s.Values) != 1 || s.Values[0] != "top" {
		t.Fatalf("再点已选应移出，got %v", s.Values)
	}
	if s.tree.SelectedSet["input"] {
		t.Fatalf("移出后 tree.SelectedSet 应同步清掉 input")
	}
	if len(got) != 3 || len(got[2]) != 1 || got[2][0] != "top" {
		t.Fatalf("第三次回调应得 [top]，got %v", got)
	}
}

// TestTreeSelectMultiRemoveAndMaxTag × 删除 + MaxTagCount 折叠。
func TestTreeSelectMultiRemoveAndMaxTag(t *testing.T) {
	s := NewTreeSelect([]TreeNode{
		{Key: "a", Label: "甲"},
		{Key: "b", Label: "乙"},
		{Key: "c", Label: "丙"},
	})
	s.Multiple = true
	var calls int
	s.OnValuesChange = func([]string) { calls++ }

	s.Values = []string{"a", "b", "c"}
	s.removeValue("b", nil)
	if len(s.Values) != 2 || s.Values[0] != "a" || s.Values[1] != "c" {
		t.Fatalf("删除 b 应剩 [a c]，got %v", s.Values)
	}
	if calls != 1 {
		t.Fatalf("removeValue 应回调一次，got %d", calls)
	}

	s.Values = []string{"a", "b", "c", "d"}
	s.MaxTagCount = 2
	shown, overflow := s.tagShown()
	if len(shown) != 2 || shown[0] != "a" || shown[1] != "b" {
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
}
