package wg

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// findTextNode 找 Src 匹配的文本节点（诊断 hover 重建后的位置）。
func findTextNode(root *scene.Node, src string) *scene.Node {
	var found *scene.Node
	root.Walk(func(n *scene.Node) bool {
		if t := n.Text(); t != nil && t.Src == src {
			found = n
			return false
		}
		return true
	})
	return found
}

// TestInputHoverRebuildKeepsTextLayout 锁定"hover 重建后文本度量/位置不变"。
//
// 旧实现 Input 的 TextSpec 没带 LayoutFont：场景重排会用零值字体重排，
// 字号与行高都变，真窗表现为鼠标移出输入框后文本整体右下偏移。
func TestInputHoverRebuildKeepsTextLayout(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 260)

	in := widgets.NewInput("")
	in.Width = 200
	c.MountWidget(&ipPage{child: widgets.Column{Children: []ui.Widget{in}}})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.focus.Focus(in.FocusNode())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}

	now := time.Now()
	for _, ch := range "abc" {
		c.PushEvent(&input.KeyEvent{Kind: input.KeyDown, Text: string(ch), Time: now})
	}
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	n1 := findTextNode(c.Scene(), "abc")
	if n1 == nil {
		t.Fatal("找不到文本节点")
	}
	b1 := n1.BoundsInRoot()
	f1 := n1.Text().Layout.Font.Size
	s1 := n1.Text().Layout.Size

	hover := func(x, y float32) {
		pt := geom.Pt(x, y)
		c.PushEvent(&input.PointerEvent{
			Kind: input.PointerMove, Pointer: input.MousePointer,
			Pos: pt, PosRoot: pt, Time: now,
		})
		if _, err := c.FrameOnce(); err != nil {
			t.Fatal(err)
		}
	}

	hover(50, 16)   // enter
	hover(300, 200) // leave
	n2 := findTextNode(c.Scene(), "abc")
	if n2 == nil {
		t.Fatal("移出后找不到文本节点")
	}
	b2 := n2.BoundsInRoot()
	f2 := n2.Text().Layout.Font.Size
	s2 := n2.Text().Layout.Size
	if b1.Min != b2.Min {
		t.Fatalf("文本位置发生偏移: %v → %v", b1.Min, b2.Min)
	}
	if f1 != f2 || s1 != s2 {
		t.Fatalf("重建后文本度量变了: 字体 %v→%v，排版 %v→%v", f1, f2, s1, s2)
	}
}
