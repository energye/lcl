package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/scene"
)

func anchorItems() []AnchorItem {
	return []AnchorItem{{Key: "a", Title: "第一部分"}, {Key: "b", Title: "第二部分"}, {Key: "c", Title: "第三部分"}}
}

// nodeHasAnchorText 子树文本查找。
func nodeHasAnchorText(n *scene.Node, s string) bool {
	found := false
	n.Walk(func(x *scene.Node) bool {
		if sp := x.Text(); sp != nil && sp.Src == s {
			found = true
			return false
		}
		return true
	})
	return found
}

// TestAnchorInitial 首项默认高亮。
func TestAnchorInitial(t *testing.T) {
	a := NewAnchor(anchorItems()...)
	n := a.Build(dataCtx())
	if a.ActiveKey != "a" {
		t.Fatalf("默认高亮 = %q，应 a", a.ActiveKey)
	}
	if !nodeHasAnchorText(n, "第二部分") {
		t.Fatal("锚项标题缺失")
	}
}

// TestAnchorClick 点击切换高亮+回调。
func TestAnchorClick(t *testing.T) {
	a := NewAnchor(anchorItems()...)
	var changed string
	a.OnChange = func(k string) { changed = k }
	a.Build(dataCtx())
	a.activate("b")
	if a.ActiveKey != "b" || changed != "b" {
		t.Fatalf("点击后高亮 = %q changed = %q", a.ActiveKey, changed)
	}
	// 同 key 重复激活不回调。
	changed = ""
	a.activate("b")
	if changed != "" {
		t.Fatal("重复激活不应回调")
	}
}

// TestAnchorScrollSync 滚动位置注入→最近区块翻转。
func TestAnchorScrollSync(t *testing.T) {
	a := NewAnchor(anchorItems()...)
	a.Build(dataCtx())
	// 均分假设：a=0, b=200, c=400。
	a.ScrollPos = func() float32 { return 210 }
	a.Bounds = 5
	a.syncScroll()
	if a.ActiveKey != "b" {
		t.Fatalf("滚动 210 应高亮 b，实际 %q", a.ActiveKey)
	}
	// 回滚到首区块。
	a.ScrollPos = func() float32 { return 10 }
	a.syncScroll()
	if a.ActiveKey != "a" {
		t.Fatalf("滚动 10 应高亮 a，实际 %q", a.ActiveKey)
	}
}

// TestAnchorHorizontal 横排锚项。
func TestAnchorHorizontal(t *testing.T) {
	a := NewAnchor(anchorItems()...)
	a.Direction = AnchorHorizontal
	n := a.Build(dataCtx())
	if !nodeHasAnchorText(n, "第三部分") {
		t.Fatal("横排锚项缺失")
	}
}
