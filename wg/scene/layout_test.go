package scene

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/text"
)

// TestThreeDirtyFlagsAreSeparate 是 M2 最容易写错的一条：
// 三档脏标记必须独立，否则"改个 hover 颜色"会触发整棵树重排。
func TestThreeDirtyFlagsAreSeparate(t *testing.T) {
	leaf := Box(BoxSpec{Fill: render.Fill(geom.RGBA(1, 2, 3))}).SetSize2(10, 10)
	root := Flex(Vertical, 0, leaf)
	root.Layout(Unbounded())
	root.ClearDirtyTree()

	leaf.MarkNeedsPaint()
	if leaf.NeedsLayout() {
		t.Fatal("MarkNeedsPaint 不应标布局脏")
	}
	if root.NeedsLayout() {
		t.Fatal("MarkNeedsPaint 不应冒泡到父节点")
	}
	if !leaf.NeedsPaint() {
		t.Fatal("MarkNeedsPaint 后应处于绘制脏")
	}

	root.ClearDirtyTree()
	leaf.MarkNeedsLayout()
	if !root.NeedsLayout() {
		t.Fatal("MarkNeedsLayout 应冒泡到父节点")
	}

	root.ClearDirtyTree()
	leaf.MarkNeedsComposite()
	if leaf.NeedsLayout() || leaf.NeedsPaint() {
		t.Fatal("MarkNeedsComposite 不应顺带标布局/绘制脏")
	}
}

// TestLayoutBoundaryStopsBubbling 验证布局边界：脏标记在边界处停住。
func TestLayoutBoundaryStopsBubbling(t *testing.T) {
	inner := Box(BoxSpec{}).SetSize2(10, 10)
	boundary := Box(BoxSpec{}).SetLayoutBoundary(true)
	boundary.Add(inner)
	root := Flex(Vertical, 0, boundary)
	root.Layout(Unbounded())
	root.ClearDirtyTree()

	inner.MarkNeedsLayout()
	if !boundary.NeedsLayout() {
		t.Fatal("边界自身应变脏")
	}
	if root.NeedsLayout() {
		t.Fatal("脏标记不应越过布局边界")
	}
}

// TestLayoutCacheHits 验证"约束不变则不重排"。
func TestLayoutCacheHits(t *testing.T) {
	n := Flex(Vertical, 4, Box(BoxSpec{}).SetSize2(10, 10))
	c := LooseConstraints(geom.Sz(100, 100))
	n.Layout(c)
	first := n.Children()[0].Offset()
	// 手动把子节点位置改掉，再跑一次：如果走了缓存，位置不会被重算
	n.Children()[0].offset = geom.Pt(999, 999)
	got := n.Layout(c)
	if got != n.Size() {
		t.Fatal("缓存命中应返回同一尺寸")
	}
	if n.Children()[0].Offset() != geom.Pt(999, 999) {
		t.Fatal("约束不变且不脏时不应重新摆放子节点")
	}
	_ = first
	// 标脏后必须重算
	n.Children()[0].MarkNeedsLayout()
	n.Layout(c)
	if n.Children()[0].Offset() == geom.Pt(999, 999) {
		t.Fatal("标脏后应重新摆放")
	}
}

// TestValidateCatchesKindSpecMismatch 验证 Validate 能抓住"声明与数据不一致"。
func TestValidateCatchesKindSpecMismatch(t *testing.T) {
	n := newNode(KindText) // 故意不给 text spec
	if err := n.Validate(); err == nil {
		t.Fatal("KindText 缺 text spec 应被 Validate 抓到")
	}
	ok := Text(TextSpec{Src: "hi", Layout: text.LayoutText("hi", text.DefaultFont(12), text.Options{})})
	if err := ok.Validate(); err != nil {
		t.Fatalf("合法文本节点不应报错: %v", err)
	}
	// 单子节点容器塞两个子节点
	bad := Clip(ClipSpec{}, Box(BoxSpec{}))
	bad.Add(Box(BoxSpec{}))
	if err := bad.Validate(); err == nil {
		t.Fatal("单子节点容器塞两个子节点应被 Validate 抓到")
	}
}

// TestRecordAppliesOffsetAndPadding 守住"子节点 offset 已含父 padding"这条约定。
func TestRecordAppliesOffsetAndPadding(t *testing.T) {
	child := Box(BoxSpec{Fill: render.Fill(geom.RGBA(0, 0, 0))}).SetSize2(10, 10)
	root := Box(BoxSpec{}).SetPadding(geom.In(6)).Add(child)
	root.Layout(LooseConstraints(geom.Sz(100, 100)))

	if got, want := child.Offset(), geom.Pt(6, 6); got != want {
		t.Fatalf("子节点偏移应含 padding：得到 %v 期望 %v", got, want)
	}
	dl := root.RecordTree(nil)
	if dl.Empty() {
		t.Fatal("录制结果不应为空")
	}
	// 填充矩形的设备坐标 = padding + 子节点自身（0,0）→ (6,6)
	b := dl.Bounds()
	if b.Min.X < 5.9 || b.Min.Y < 5.9 {
		t.Fatalf("矩形应在 (6,6) 起，实际包围盒 %v", b)
	}
}

// TestLocalHitShapes 验证命中形状与绘制几何一致。
func TestLocalHitShapes(t *testing.T) {
	r := Box(BoxSpec{}).SetSize2(20, 20)
	r.Layout(Unbounded())
	r.SetHit(HitSpec{Shape: ShapeRect})
	if !r.LocalHit(geom.Pt(1, 1)) || r.LocalHit(geom.Pt(-1, 1)) {
		t.Fatal("矩形命中判定错误")
	}

	rr := Box(BoxSpec{}).SetSize2(20, 20)
	rr.Layout(Unbounded())
	rr.SetHit(HitSpec{Shape: ShapeRRect, Radius: 10})
	// 圆角半径 10 = 半宽，四角必须落在圆外
	if rr.LocalHit(geom.Pt(0.5, 0.5)) {
		t.Fatal("圆角矩形的方角不应命中")
	}
	if !rr.LocalHit(geom.Pt(10, 10)) {
		t.Fatal("圆心应命中")
	}

	c := Box(BoxSpec{}).SetSize2(20, 20)
	c.Layout(Unbounded())
	c.SetHit(HitSpec{Shape: ShapeCircle})
	if c.LocalHit(geom.Pt(0.5, 0.5)) {
		t.Fatal("圆形的方角不应命中")
	}
	if !c.LocalHit(geom.Pt(10, 10)) {
		t.Fatal("圆心应命中")
	}
}

// TestScrollClampAndFollow 验证滚动偏移的 clamp 与内容跟随。
func TestScrollClampAndFollow(t *testing.T) {
	content := Flex(Vertical, 0, Box(BoxSpec{}).SetSize2(10, 300))
	s := Scroll(Vertical, content)
	s.Layout(TightConstraints(geom.Sz(50, 100)))
	if got := s.ScrollContentSize().H; !approxf(got, 300) {
		t.Fatalf("内容高应为 300，实际 %v", got)
	}
	s.SetScrollOffset(50)
	if !approxf(content.Offset().Y, -50) {
		t.Fatalf("内容应上移 50，实际 %v", content.Offset().Y)
	}
	s.SetScrollOffset(10000)
	if got := s.ScrollOffset(); !approxf(got, 200) {
		t.Fatalf("越界偏移应 clamp 到 200，实际 %v", got)
	}
}

// TestConstraintsInvariants 守住 Constraints 的不变式。
func TestConstraintsInvariants(t *testing.T) {
	c := NewConstraints(50, 10, 30, 20) // 故意倒置
	if c.MinW > c.MaxW || c.MinH > c.MaxH {
		t.Fatalf("NewConstraints 应修正倒置: %v", c)
	}
	if got := c.Constrain(geom.Sz(5, 5)); got != (geom.Size{W: 50, H: 30}) {
		t.Fatalf("Constrain 应取下界，实际 %v", got)
	}
	d := c.Deflate(geom.In(100)) // 内边距大于可用空间
	if d.MinW < 0 || d.MinH < 0 {
		t.Fatalf("Deflate 不应产生负约束: %v", d)
	}
}

func approxf(a, b float32) bool {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d <= 0.01
}
