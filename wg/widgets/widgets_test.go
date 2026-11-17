package widgets

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

func testCtx() *ui.BuildContext { return &ui.BuildContext{Theme: theme.Default()} }

func TestDividerSize(t *testing.T) {
	h := NewDivider().Build(testCtx())
	h.Layout(scene.Unbounded())
	if s := h.Size(); s.W != 120 || s.H != 1 {
		t.Fatalf("水平分隔线尺寸 = %+v", s)
	}
	v := (&Divider{Vertical: true, Length: 40, Thickness: 2}).Build(testCtx())
	v.Layout(scene.Unbounded())
	if s := v.Size(); s.W != 2 || s.H != 40 {
		t.Fatalf("竖直分隔线尺寸 = %+v", s)
	}
}

func TestRowColumnLayout(t *testing.T) {
	row := Row{Gap: 10, Children: []ui.Widget{
		Spacer{Width: 20, Height: 20}, Spacer{Width: 20, Height: 20},
	}}.Build(testCtx())
	row.Layout(scene.TightConstraints(geom.Size{W: 200, H: 50}))
	ch := row.Children()
	if len(ch) != 2 || ch[1].Bounds().Min.X != 30 {
		t.Fatalf("Row 第二项 x = %v，应为 30", ch[1].Bounds().Min.X)
	}

	col := Column{Gap: 5, Children: []ui.Widget{
		Spacer{Width: 10, Height: 10}, Spacer{Width: 10, Height: 10},
	}}.Build(testCtx())
	col.Layout(scene.TightConstraints(geom.Size{W: 50, H: 200}))
	ch = col.Children()
	if len(ch) != 2 || ch[1].Bounds().Min.Y != 15 {
		t.Fatalf("Column 第二项 y = %v，应为 15", ch[1].Bounds().Min.Y)
	}
}

func TestPaddingIsPartOfIntrinsicSize(t *testing.T) {
	n := Padding{
		Insets: geom.InLTRB(5, 7, 5, 7),
		Child:  Spacer{Width: 20, Height: 10},
	}.Build(testCtx())
	sz := n.Layout(scene.Unbounded())
	if sz.W != 30 || sz.H != 24 {
		t.Fatalf("带 padding 的固有尺寸 = %+v，应为 30×24", sz)
	}
}

func TestListReusesSameKey(t *testing.T) {
	l := NewList(
		ListItem{Key: "a", Widget: Spacer{Width: 1, Height: 1}},
		ListItem{Key: "b", Widget: Spacer{Width: 1, Height: 1}},
	)
	l.Build(testCtx())
	if len(l.Nodes()) != 2 {
		t.Fatalf("首次构建节点数 = %d", len(l.Nodes()))
	}
	a := l.Nodes()[0]
	l.Items = append([]ListItem{{Key: "x", Widget: Spacer{}}}, l.Items...)
	l.Build(testCtx())
	if l.Nodes()[1] != a {
		t.Fatal("同 Key 节点未复用（输入框会丢焦点）")
	}
}

func TestScrollKeepsAxis(t *testing.T) {
	n := (&Scroll{Axis: scene.Horizontal, Width: 100, Height: 40,
		Child: Spacer{Width: 400, Height: 20}}).Build(testCtx())
	n.Layout(scene.Unbounded())
	if s := n.Size(); s.W != 100 || s.H != 40 {
		t.Fatalf("视口尺寸 = %+v", s)
	}
	n.Layout(scene.TightConstraints(geom.Size{W: 100, H: 40}))
	n.SetScrollOffset(20)
	if got := n.Size().W; got != 100 {
		t.Fatalf("滚动后视口尺寸变了: %v", got)
	}
}

func TestLabelUsesThemeFont(t *testing.T) {
	n := Label{Text: "中文", Size: 18}.Build(testCtx())
	if n.Size().W <= 0 || n.Size().H <= 0 {
		t.Fatalf("文本固有尺寸 = %+v", n.Size())
	}
}
