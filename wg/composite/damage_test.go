package composite

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
)

func TestRegionMergesBands(t *testing.T) {
	r := NewRegion()
	r.Add(geom.RectXYWH(0, 0, 10, 10))
	r.Add(geom.RectXYWH(10, 0, 10, 10)) // 同带相邻 → 合并
	if r.Len() != 1 {
		t.Fatalf("同带相邻应合并，实际 %d 个矩形: %v", r.Len(), r.Rects())
	}
	if got := r.Bounds(); got != geom.RectLTRB(0, 0, 20, 10) {
		t.Fatalf("合并后 = %v", got)
	}
	if got := r.Area(); got != 200 {
		t.Fatalf("面积 = %v", got)
	}
}

func TestRegionKeepsDisjoint(t *testing.T) {
	r := NewRegion()
	r.Add(geom.RectXYWH(0, 0, 10, 10))
	r.Add(geom.RectXYWH(100, 100, 10, 10))
	if r.Len() != 2 {
		t.Fatalf("不相邻矩形不该合并: %v", r.Rects())
	}
	if got := r.Area(); got != 200 {
		t.Fatalf("面积 = %v", got)
	}
	if got := r.Bounds(); got != geom.RectLTRB(0, 0, 110, 110) {
		t.Fatalf("外包 = %v", got)
	}
}

func TestRegionContainsAndIntersect(t *testing.T) {
	r := RegionOf(geom.RectXYWH(10, 10, 20, 20))
	if !r.Contains(geom.Pt(15, 15)) {
		t.Fatal("应包含")
	}
	if r.Contains(geom.Pt(5, 5)) {
		t.Fatal("不该包含")
	}
	r.Intersect(geom.RectXYWH(0, 0, 20, 20))
	if got := r.Bounds(); got != geom.RectLTRB(10, 10, 20, 20) {
		t.Fatalf("裁剪后 = %v", got)
	}
	r.Intersect(geom.RectXYWH(100, 100, 10, 10))
	if !r.IsEmpty() {
		t.Fatalf("完全裁掉后应为空: %v", r.Rects())
	}
}

func TestRegionOcclude(t *testing.T) {
	r := NewRegion()
	r.Add(geom.RectXYWH(0, 0, 10, 10))     // 被完全覆盖
	r.Add(geom.RectXYWH(50, 50, 10, 10))   // 不受影响
	r.Add(geom.RectXYWH(100, 0, 40, 10))   // 左侧被覆盖
	r.Occlude(geom.RectXYWH(0, 0, 110, 10))
	// 第一个消失；第三个从 x=110 开始保留 30 宽
	if r.Area() != 100+300 {
		t.Fatalf("遮挡后面积 = %v，矩形 = %v", r.Area(), r.Rects())
	}
}

func TestRegionCloneAndClear(t *testing.T) {
	r := RegionOf(geom.RectXYWH(0, 0, 5, 5))
	c := r.Clone()
	c.Add(geom.RectXYWH(100, 100, 5, 5))
	if r.Len() != 1 {
		t.Fatal("Clone 应为深拷贝")
	}
	if c.Len() != 2 {
		t.Fatal("副本应有两个矩形")
	}
	r.Clear()
	if !r.IsEmpty() {
		t.Fatal("Clear 后应为空")
	}
}
