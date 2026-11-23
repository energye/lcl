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
