package host

import (
	"errors"
	"testing"

	"github.com/energye/lcl/wg/geom"
)

func TestAppendAndDetach(t *testing.T) {
	root := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 100, 100))
	a := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 10, 10))
	b := NewSurface(SurfaceLayer, geom.RectXYWH(0, 0, 10, 10))

	if err := root.Append(a); err != nil {
		t.Fatal(err)
	}
	if err := root.Append(b); err != nil {
		t.Fatal(err)
	}
	if got := len(root.Children()); got != 2 {
		t.Fatalf("子节点数 = %d", got)
	}
	if a.Index() != 0 || b.Index() != 1 {
		t.Fatalf("z 序错误: a=%d b=%d", a.Index(), b.Index())
	}
	if a.Depth() != 1 || root.Depth() != 0 {
		t.Fatalf("深度错误")
	}
	if !a.Detach() {
		t.Fatal("Detach 应返回 true")
	}
	if a.Parent() != nil || len(root.Children()) != 1 {
		t.Fatal("Detach 后状态错误")
	}
	if a.Detach() {
		t.Fatal("重复 Detach 应返回 false")
	}
	if !root.Remove(b) {
		t.Fatal("Remove 应返回 true")
	}
	if root.Remove(b) {
		t.Fatal("重复 Remove 应返回 false")
	}
}

func TestAppendRejectsCycle(t *testing.T) {
	root := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 100, 100))
	child := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 50, 50))
	if err := root.Append(child); err != nil {
		t.Fatal(err)
	}
	if err := root.Append(nil); !errors.Is(err, ErrNilSurface) {
		t.Fatalf("nil 应返回 ErrNilSurface，实际 %v", err)
	}
	if err := root.Append(root); !errors.Is(err, ErrSurfaceCycle) {
		t.Fatalf("自挂应返回 ErrSurfaceCycle，实际 %v", err)
	}
	if err := child.Append(root); !errors.Is(err, ErrSurfaceCycle) {
		t.Fatalf("祖先挂到后代应返回 ErrSurfaceCycle，实际 %v", err)
	}
	if err := root.Append(child); !errors.Is(err, ErrSurfaceAttached) {
		t.Fatalf("重复挂接应返回 ErrSurfaceAttached，实际 %v", err)
	}
	if !root.IsAncestorOf(child) || child.IsAncestorOf(root) {
		t.Fatal("IsAncestorOf 判定错误")
	}
}

func TestInsertAtReorders(t *testing.T) {
	root := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 100, 100))
	a := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 10, 10))
	b := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 10, 10))
	c := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 10, 10))
	_ = root.Append(a)
	_ = root.Append(b)
	if err := root.InsertAt(0, c); err != nil {
		t.Fatal(err)
	}
	order := root.Children()
	if order[0] != c || order[1] != a || order[2] != b {
		t.Fatalf("插入后 z 序错误")
	}
}

func TestDamageBubblesUp(t *testing.T) {
	root := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 200, 200))
	mid := NewSurface(SurfaceInline, geom.RectXYWH(20, 30, 100, 100))
	leaf := NewSurface(SurfaceInline, geom.RectXYWH(5, 7, 40, 40))
	if err := root.Append(mid); err != nil {
		t.Fatal(err)
	}
	if err := mid.Append(leaf); err != nil {
		t.Fatal(err)
	}
	root.ClearDamageTree()

	leaf.Invalidate(geom.RectXYWH(0, 0, 10, 10))
	// 叶(0,0,10,10) → 中(5,7,15,17) → 根(25,37,35,47)
	want := geom.RectLTRB(25, 37, 35, 47)
	if got := root.Damage(); got != want {
		t.Fatalf("根脏区 = %v，期望 %v", got, want)
	}
	if got := mid.Damage(); got != geom.RectLTRB(5, 7, 15, 17) {
		t.Fatalf("中间层脏区 = %v", got)
	}
}

func TestDamageClipped(t *testing.T) {
	root := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 100, 100))
	child := NewSurface(SurfaceInline, geom.RectXYWH(90, 90, 20, 20))
	if err := root.Append(child); err != nil {
		t.Fatal(err)
	}
	root.ClearDamageTree()

	// 超出子画布的部分必须被裁掉
	child.Invalidate(geom.RectXYWH(0, 0, 100, 100))
	want := geom.RectLTRB(90, 90, 100, 100) // 子 bounds 与根 bounds 双重裁剪
	if got := root.Damage(); got != want {
		t.Fatalf("根脏区 = %v，期望 %v", got, want)
	}

	// 关闭裁剪后不上裁，但仍受父裁剪限制
	child.SetClip(false)
	root.ClearDamageTree()
	child.Invalidate(geom.RectXYWH(0, 0, 30, 30))
	if got := root.Damage(); got != geom.RectLTRB(90, 90, 100, 100) {
		t.Fatalf("关闭裁剪后根脏区 = %v", got)
	}
}

func TestDamageWithMatrix(t *testing.T) {
	root := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 100, 100))
	child := NewSurface(SurfaceLayer, geom.RectXYWH(10, 10, 20, 20))
	if err := root.Append(child); err != nil {
		t.Fatal(err)
	}
	// 平移 5,5 相对父
	child.SetMatrix(geom.Translate(5, 5))
	root.ClearDamageTree()
	child.Invalidate(geom.RectXYWH(0, 0, 4, 4))
	// (10+5, 10+5) 起 4x4
	if got := root.Damage(); got != geom.RectLTRB(15, 15, 19, 19) {
		t.Fatalf("带平移的脏区 = %v", got)
	}
}

func TestStructureDirtyPropagates(t *testing.T) {
	root := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 100, 100))
	child := NewSurface(SurfaceLayer, geom.RectXYWH(0, 0, 10, 10))
	if err := root.Append(child); err != nil {
		t.Fatal(err)
	}
	root.ClearStructureDirty()
	if root.StructureDirty() {
		t.Fatal("清除后仍为脏")
	}
	child.SetBounds(geom.RectXYWH(0, 0, 20, 20))
	if !root.StructureDirty() {
		t.Fatal("子几何变化未传播到根")
	}
	// 无变化时不应重复标脏
	root.ClearStructureDirty()
	child.SetBounds(geom.RectXYWH(0, 0, 20, 20))
	if root.StructureDirty() {
		t.Fatal("相同 bounds 不应标脏")
	}
}

func TestAbsoluteGeometry(t *testing.T) {
	root := NewSurface(SurfaceInline, geom.RectXYWH(0, 0, 100, 100))
	mid := NewSurface(SurfaceLayer, geom.RectXYWH(10, 20, 50, 50))
	leaf := NewSurface(SurfaceInline, geom.RectXYWH(3, 4, 10, 10))
	_ = root.Append(mid)
	_ = mid.Append(leaf)

	if got := leaf.AbsoluteOrigin(); got != (geom.Point{X: 13, Y: 24}) {
		t.Fatalf("绝对原点 = %v", got)
	}
	if got := leaf.BoundsInRoot(); got != geom.RectLTRB(13, 24, 23, 34) {
		t.Fatalf("根坐标包围盒 = %v", got)
	}
	if n := root.Count(); n != 3 {
		t.Fatalf("节点数 = %d", n)
	}
	// Walk 返回 false 应立即终止
	seen := 0
	root.Walk(func(*Surface) bool {
		seen++
		return seen < 2
	})
	if seen != 2 {
		t.Fatalf("Walk 提前终止失效，访问了 %d 个", seen)
	}
}

func TestSurfaceKind(t *testing.T) {
	if !SurfaceLayer.Offscreen() || !SurfaceOffscreen.Offscreen() {
		t.Fatal("Layer/Offscreen 应需要独立缓冲")
	}
	if SurfaceInline.Offscreen() || SurfacePopup.Offscreen() {
		t.Fatal("Inline/Popup 默认不申请独立缓冲")
	}
	if SurfacePopup.String() != "popup" {
		t.Fatal("String 错误")
	}
}
