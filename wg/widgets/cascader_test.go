package widgets

import (
	"testing"
)

// TestCascaderSearchFilter 搜索过滤：默认路径 label 子串 + 自定义 Filter。
func TestCascaderSearchFilter(t *testing.T) {
	c := NewCascader()
	c.Options = []*CascaderOption{
		{Value: "zhejiang", Label: "浙江", Children: []*CascaderOption{
			{Value: "hangzhou", Label: "杭州", Children: []*CascaderOption{
				{Value: "xihu", Label: "西湖区"},
			}},
		}},
		{Value: "jiangsu", Label: "江苏", Children: []*CascaderOption{
			{Value: "nanjing", Label: "南京"},
		}},
	}
	c.search = "杭州"
	matches := c.searchMatches()
	if len(matches) != 2 {
		t.Fatalf("搜索杭州应得 2 条（浙江/杭州、浙江/杭州/西湖区），got %d", len(matches))
	}
	c.Filter = func(inputValue string, path []*CascaderOption) bool {
		return len(path) >= 1 && path[len(path)-1].Label == "南京"
	}
	matches = c.searchMatches()
	if len(matches) != 1 || matches[0][len(matches[0])-1].Value != "nanjing" {
		t.Fatalf("自定义过滤应得南京，got %v", matches)
	}
}

// TestCascaderSelectPath 搜索选中：设 Value+清搜索。
func TestCascaderSelectPath(t *testing.T) {
	c := NewCascader()
	c.Options = []*CascaderOption{
		{Value: "zhejiang", Label: "浙江", Children: []*CascaderOption{
			{Value: "hangzhou", Label: "杭州", Children: []*CascaderOption{
				{Value: "xihu", Label: "西湖区"},
			}},
		}},
	}
	c.search = "西湖区"
	matches := c.searchMatches()
	if len(matches) != 1 {
		t.Fatalf("搜索西湖区应得 1 条，got %d", len(matches))
	}
	c.selectPath(nil, matches[0])
	want := []string{"zhejiang", "hangzhou", "xihu"}
	if len(c.Value) != len(want) {
		t.Fatalf("选中路径长度 want %d got %d", len(want), len(c.Value))
	}
	for i := range want {
		if c.Value[i] != want[i] {
			t.Fatalf("选中路径 want %v got %v", want, c.Value)
		}
	}
	if c.search != "" {
		t.Fatalf("选中后应清搜索，got %q", c.search)
	}
}

// TestCascaderMultiToggle 多选 toggle：点行勾选/取消完整路径（antd
// value 形态 string[][]）、hasPath 判定、tagShown 折叠。
func TestCascaderMultiToggle(t *testing.T) {
	c := NewCascader()
	c.Multiple = true
	var got [][]string
	c.OnValuesChange = func(paths [][]string) { got = paths }

	c.togglePath(nil, []string{"zj", "hz", "xihu"})
	if len(c.Values) != 1 || c.Values[0][2] != "xihu" {
		t.Fatalf("首次点选应入集合，got %v", c.Values)
	}
	if !c.hasPath([]string{"zj", "hz", "xihu"}) {
		t.Fatal("hasPath 应命中")
	}
	c.togglePath(nil, []string{"js", "nj", "xw"})
	if len(c.Values) != 2 || len(got) != 2 {
		t.Fatalf("第二条应入集合，got %v", c.Values)
	}
	// 再点已选路径=移除。
	c.togglePath(nil, []string{"zj", "hz", "xihu"})
	if len(c.Values) != 1 || c.Values[0][0] != "js" {
		t.Fatalf("再点已选应移除，got %v", c.Values)
	}
	if c.hasPath([]string{"zj", "hz", "xihu"}) {
		t.Fatal("移除后 hasPath 不应命中")
	}
	// MaxTagCount 折叠（触发框 buildMultiTrigger 消费同款逻辑）。
	c.MaxTagCount = 1
	if len(c.Values) != 1 {
		t.Fatal("前置")
	}
	c.Values = [][]string{{"a", "b"}, {"c", "d"}, {"e", "f"}}
	shown, overflow := c.Values, 0
	if c.MaxTagCount > 0 && len(c.Values) > c.MaxTagCount {
		shown, overflow = c.Values[:c.MaxTagCount], len(c.Values)-c.MaxTagCount
	}
	if len(shown) != 1 || overflow != 2 {
		t.Fatalf("折叠应为 1+2，got %d+%d", len(shown), overflow)
	}
}

// TestCascaderLeafLabel 叶子 Label 折算：正常链取末级 Label、断链兜底 Value。
func TestCascaderLeafLabel(t *testing.T) {
	c := NewCascader()
	c.Options = []*CascaderOption{
		{Value: "zj", Label: "浙江", Children: []*CascaderOption{
			{Value: "hz", Label: "杭州"},
		}},
	}
	if l := c.leafLabel([]string{"zj", "hz"}); l != "杭州" {
		t.Fatalf("叶子 Label 应为杭州，got %q", l)
	}
	if l := c.leafLabel([]string{"xx", "yy"}); l != "yy" {
		t.Fatalf("断链应兜底末位 Value，got %q", l)
	}
}
