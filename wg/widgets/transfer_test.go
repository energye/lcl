package widgets

import (
	"testing"
)

// TestTransferSearchFilter 搜索过滤：默认 title 不区分大小写子串 + 自定义 FilterOption。
func TestTransferSearchFilter(t *testing.T) {
	tf := NewTransfer([]TransferItem{
		{Key: "1", Title: "内容一"},
		{Key: "2", Title: "内容二"},
		{Key: "3", Title: "文章三"},
	})
	// 默认：搜索「内容」→ 内容一/二。
	vis := tf.filterItems(tf.DataSource, "left", "内容")
	if len(vis) != 2 {
		t.Fatalf("默认过滤「内容」应得 2 项，got %d", len(vis))
	}
	// 空搜索 → 全部。
	if got := tf.filterItems(tf.DataSource, "left", ""); len(got) != 3 {
		t.Fatalf("空搜索应得全部 3 项，got %d", len(got))
	}
	// 自定义：只匹配「文章」。
	tf.FilterOption = func(inputValue string, item TransferItem, direction string) bool {
		return item.Title == "文章三"
	}
	vis = tf.filterItems(tf.DataSource, "right", "x")
	if len(vis) != 1 || vis[0].Key != "3" {
		t.Fatalf("自定义过滤应得文章三，got %v", vis)
	}
}

// TestTransferSelectAll 全选：勾选/反选本栏可见项 + 三态聚合。
func TestTransferSelectAll(t *testing.T) {
	tf := NewTransfer([]TransferItem{
		{Key: "1", Title: "内容一"},
		{Key: "2", Title: "内容二"},
		{Key: "3", Title: "文章三"},
	})
	tf.selectAll("左", true)
	for _, it := range tf.DataSource {
		if !tf.checked[it.Key] {
			t.Fatalf("全选后 %s 未勾选", it.Key)
		}
	}
	tf.selectAll("左", false)
	for _, it := range tf.DataSource {
		if tf.checked[it.Key] {
			t.Fatalf("反选后 %s 仍勾选", it.Key)
		}
	}
	// 部分勾选 → indeterminate。
	tf.checked["1"] = true
	checked, indeterminate := tf.selAllState("左")
	if checked || !indeterminate {
		t.Fatalf("部分勾选应 indeterminate，got checked=%v ind=%v", checked, indeterminate)
	}
	// 全勾选 → checked。
	tf.checked["2"] = true
	tf.checked["3"] = true
	checked, indeterminate = tf.selAllState("左")
	if !checked || indeterminate {
		t.Fatalf("全勾选应 checked，got checked=%v ind=%v", checked, indeterminate)
	}
}
