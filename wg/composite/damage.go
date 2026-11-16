package composite

import "github.com/energye/lcl/wg/geom"

// Region 是一组不重叠矩形的并集，用于 damage 与裁剪。
//
// 为什么不用单矩形：局部重绘经常出现"两个不相邻的小块"（例如按钮 A 与
// 远处的输入框同时变脏）。用单矩形会把中间大片无关区域一起重画，
// 直接抵消掉 damage 的收益。
//
// 为什么不用通用多边形区域：Windows 的 HRGN 精度只有像素，且跨平台需要
// 自己实现扫描线布尔运算——M1 用"矩形列表 + 合并"已经能覆盖 UI 场景，
// 复杂度留给真正需要的时候（M5 若有 4K 全屏动画再说）。
type Region struct {
	rects []geom.Rect
}

// NewRegion 创建空区域。
func NewRegion() *Region { return &Region{} }

// RegionOf 用单矩形创建区域。
func RegionOf(r geom.Rect) *Region {
	rg := &Region{}
	rg.Add(r)
	return rg
}

// Add 把矩形并入区域（相邻/重叠的会被合并）。
func (r *Region) Add(rc geom.Rect) {
	if rc.IsEmpty() {
		return
	}
	rc = rc.RoundOut()
	// 先尝试与已有矩形合并（含相邻情况：共边时合并可减少后续重绘次数）
	for i := 0; i < len(r.rects); i++ {
		if merged, ok := mergeRect(r.rects[i], rc); ok {
			r.rects[i] = merged
			r.normalize()
			return
		}
	}
	r.rects = append(r.rects, rc)
	r.normalize()
}

// AddRegion 并入另一个区域。
func (r *Region) AddRegion(o *Region) {
	if o == nil {
		return
	}
	for _, rc := range o.rects {
		r.Add(rc)
	}
}

// Intersect 与矩形求交（就地修改）。
func (r *Region) Intersect(clip geom.Rect) {
	if clip.IsEmpty() {
		r.rects = r.rects[:0]
		return
	}
	out := r.rects[:0]
	for _, rc := range r.rects {
		if v := rc.Intersect(clip); !v.IsEmpty() {
			out = append(out, v)
		}
	}
	r.rects = out
}

// Bounds 返回外包矩形。
func (r *Region) Bounds() geom.Rect {
	var b geom.Rect
	for _, rc := range r.rects {
		b = b.Union(rc)
	}
	return b
}

// Area 返回总面积（像素²）。
func (r *Region) Area() float32 {
	a := float32(0)
	for _, rc := range r.rects {
		a += rc.W() * rc.H()
	}
	return a
}

// IsEmpty 表示没有区域。
func (r *Region) IsEmpty() bool { return r == nil || len(r.rects) == 0 }

// Len 返回矩形个数。
func (r *Region) Len() int {
	if r == nil {
		return 0
	}
	return len(r.rects)
}

// Rects 返回矩形列表（只读）。
func (r *Region) Rects() []geom.Rect {
	if r == nil {
		return nil
	}
	return r.rects
}

// Clear 清空。
func (r *Region) Clear() { r.rects = r.rects[:0] }

// Contains 判定点是否在区域内。
func (r *Region) Contains(p geom.Point) bool {
	for _, rc := range r.rects {
		if rc.Contains(p) {
			return true
		}
	}
	return false
}

// Clone 深拷贝。
func (r *Region) Clone() *Region {
	out := &Region{rects: make([]geom.Rect, len(r.rects))}
	copy(out.rects, r.rects)
	return out
}

// normalize 合并所有可合并的矩形，直到稳定。
func (r *Region) normalize() {
	for changed := true; changed; {
		changed = false
		for i := 0; i < len(r.rects); i++ {
			for j := i + 1; j < len(r.rects); j++ {
				if m, ok := mergeRect(r.rects[i], r.rects[j]); ok {
					r.rects[i] = m
					r.rects = append(r.rects[:j], r.rects[j+1:]...)
					changed = true
					j--
				}
			}
		}
	}
}

// mergeRect 判定两矩形能否用更少的矩形表示（重叠或共边且对齐）。
//
// 返回的合并矩形必须**等于**两者的并集，否则会多画——所以只在
// 完全共边或包含时合并，不做保守近似。
func mergeRect(a, b geom.Rect) (geom.Rect, bool) {
	// 包含关系
	if a.ContainsRect(b) {
		return a, true
	}
	if b.ContainsRect(a) {
		return b, true
	}
	u := a.Union(b)
	// 同一行带且横向**重叠或相接**才合并。
	//
	// 只判"y 范围相同"就合并是错的：那会把 (0,0,10,10) 与 (100,0,10,10)
	// 合并成 (0,0,110,10)，凭空多画中间 90px 的空白——damage 会被放大，
	// C20 的"面积断言"立刻失真（被 TestRegionOcclude 抓住）。
	if a.Min.Y == b.Min.Y && a.Max.Y == b.Max.Y && a.Max.X >= b.Min.X && b.Max.X >= a.Min.X {
		return u, true
	}
	// 同一列带且纵向重叠或相接
	if a.Min.X == b.Min.X && a.Max.X == b.Max.X && a.Max.Y >= b.Min.Y && b.Max.Y >= a.Min.Y {
		return u, true
	}
	return geom.ZeroRect, false
}

// Occlude 用覆盖矩形剔除区域中被完全遮住的子区域（保守实现）。
//
// M1 只处理"某个矩形完全落在覆盖者内"的情况；部分遮挡保留原样
// （宁可多画，不可少画——少画就是花屏）。
func (r *Region) Occlude(cover geom.Rect) {
	if cover.IsEmpty() {
		return
	}
	out := r.rects[:0]
	for _, rc := range r.rects {
		if cover.ContainsRect(rc) {
			continue
		}
		if v := subtractRect(rc, cover); !v.IsEmpty() {
			out = append(out, v)
		}
	}
	r.rects = out
}

// subtractRect 返回 a 减去 b 的一部分（只处理 b 覆盖 a 一整条边的简单情形）。
func subtractRect(a, b geom.Rect) geom.Rect {
	if !a.Intersects(b) {
		return a
	}
	if a.ContainsRect(b) {
		// 中间挖洞无法用单矩形表达：返回原矩形（保守）
		return a
	}
	// b 从左侧切入
	if b.Min.X <= a.Min.X && b.Max.X < a.Max.X {
		return geom.RectLTRB(b.Max.X, a.Min.Y, a.Max.X, a.Max.Y)
	}
	// b 从右侧切入
	if b.Max.X >= a.Max.X && b.Min.X > a.Min.X {
		return geom.RectLTRB(a.Min.X, a.Min.Y, b.Min.X, a.Max.Y)
	}
	// b 从上方切入
	if b.Min.Y <= a.Min.Y && b.Max.Y < a.Max.Y {
		return geom.RectLTRB(a.Min.X, b.Max.Y, a.Max.X, a.Max.Y)
	}
	// b 从下方切入
	if b.Max.Y >= a.Max.Y && b.Min.Y > a.Min.Y {
		return geom.RectLTRB(a.Min.X, a.Min.Y, a.Max.X, b.Min.Y)
	}
	return a
}
