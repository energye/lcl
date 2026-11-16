package render

import "github.com/energye/lcl/wg/geom"

// verb 是路径指令。
type verb uint8

const (
	verbMove verb = iota
	verbLine
	verbQuad
	verbCubic
	verbClose
)

// Path 是可复用的路径对象：由直线、二次/三次贝塞尔、圆弧组成。
//
// 为什么不用"点列表 + 曲线标记"的即时模型：路径会被多次重放（每帧录制、
// 缓存重用、命中测试），保留原始曲线比保留扁平化后的折线更省内存也更准。
// 扁平化发生在光栅或命中阶段，容差由调用方给。
//
// 坐标系：路径点处于**录制时的局部坐标**，光栅期由 Command.Matrix 变换到设备。
type Path struct {
	verbs []verb
	// pts 按 verb 存放：move/line 1 点，quad 2 点，cubic 3 点。
	pts []geom.Point

	bounds   geom.Rect
	computed bool
}

// NewPath 创建空路径。
func NewPath() *Path { return &Path{} }

// Reset 清空路径以便复用（避免每帧分配），返回自身以便链式构建。
func (p *Path) Reset() *Path {
	p.verbs = p.verbs[:0]
	p.pts = p.pts[:0]
	p.bounds = geom.ZeroRect
	p.computed = false
	return p
}

// MoveTo 开始新的子路径。
func (p *Path) MoveTo(x, y float32) *Path {
	p.verbs = append(p.verbs, verbMove)
	p.pts = append(p.pts, geom.Pt(x, y))
	p.computed = false
	return p
}

// LineTo 追加直线段。
func (p *Path) LineTo(x, y float32) *Path {
	if len(p.verbs) == 0 {
		return p.MoveTo(x, y)
	}
	p.verbs = append(p.verbs, verbLine)
	p.pts = append(p.pts, geom.Pt(x, y))
	p.computed = false
	return p
}

// QuadTo 追加二次贝塞尔（cx,cy 为控制点）。
func (p *Path) QuadTo(cx, cy, x, y float32) *Path {
	if len(p.verbs) == 0 {
		p.MoveTo(cx, cy)
	}
	p.verbs = append(p.verbs, verbQuad)
	p.pts = append(p.pts, geom.Pt(cx, cy), geom.Pt(x, y))
	p.computed = false
	return p
}

// CubicTo 追加三次贝塞尔。
func (p *Path) CubicTo(c1x, c1y, c2x, c2y, x, y float32) *Path {
	if len(p.verbs) == 0 {
		p.MoveTo(c1x, c1y)
	}
	p.verbs = append(p.verbs, verbCubic)
	p.pts = append(p.pts, geom.Pt(c1x, c1y), geom.Pt(c2x, c2y), geom.Pt(x, y))
	p.computed = false
	return p
}

// Close 闭合当前子路径。
func (p *Path) Close() *Path {
	if len(p.verbs) == 0 {
		return p
	}
	p.verbs = append(p.verbs, verbClose)
	return p
}

// AddRect 追加一个矩形（独立子路径，方向为顺时针）。
func (p *Path) AddRect(r geom.Rect) *Path {
	if r.IsEmpty() {
		return p
	}
	return p.MoveTo(r.Min.X, r.Min.Y).
		LineTo(r.Max.X, r.Min.Y).
		LineTo(r.Max.X, r.Max.Y).
		LineTo(r.Min.X, r.Max.Y).
		Close()
}

// AddRRect 追加圆角矩形。半径超过半边长时按半边长夹紧。
func (p *Path) AddRRect(r geom.Rect, radius float32) *Path {
	if r.IsEmpty() {
		return p
	}
	maxR := minf(r.W(), r.H()) * 0.5
	radius = clampf(radius, 0, maxR)
	if radius <= 0.01 {
		return p.AddRect(r)
	}
	// 用三次贝塞尔近似四分之一圆，kappa 为经典 0.5522847。
	const k = 0.5522847
	c := radius * k
	x0, y0, x1, y1 := r.Min.X, r.Min.Y, r.Max.X, r.Max.Y

	p.MoveTo(x0+radius, y0)
	p.LineTo(x1-radius, y0)
	p.CubicTo(x1-radius+c, y0, x1, y0+radius-c, x1, y0+radius)
	p.LineTo(x1, y1-radius)
	p.CubicTo(x1, y1-radius+c, x1-radius+c, y1, x1-radius, y1)
	p.LineTo(x0+radius, y1)
	p.CubicTo(x0+radius-c, y1, x0, y1-radius+c, x0, y1-radius)
	p.LineTo(x0, y0+radius)
	p.CubicTo(x0, y0+radius-c, x0+radius-c, y0, x0+radius, y0)
	return p.Close()
}

// AddOval 追加椭圆（内接于 r）。
func (p *Path) AddOval(r geom.Rect) *Path {
	if r.IsEmpty() {
		return p
	}
	cx, cy := (r.Min.X+r.Max.X)*0.5, (r.Min.Y+r.Max.Y)*0.5
	rx, ry := r.W()*0.5, r.H()*0.5
	const k = 0.5522847
	kx, ky := rx*k, ry*k

	p.MoveTo(cx+rx, cy)
	p.CubicTo(cx+rx, cy-ky, cx+kx, cy-ry, cx, cy-ry)
	p.CubicTo(cx-kx, cy-ry, cx-rx, cy-ky, cx-rx, cy)
	p.CubicTo(cx-rx, cy+ky, cx-kx, cy+ry, cx, cy+ry)
	p.CubicTo(cx+kx, cy+ry, cx+rx, cy+ky, cx+rx, cy)
	return p.Close()
}

// AddArc 追加圆弧（内接于 r），角度为弧度，从 start 扫过 sweep。
//
// 不做"连到圆心"的饼形，也不做弦闭合：M1 的用法是描边弧线
// （进度环、仪表盘），需要开放子路径；要饼形由调用方自己 MoveTo 圆心。
func (p *Path) AddArc(r geom.Rect, start, sweep float32) *Path {
	if r.IsEmpty() || sweep == 0 {
		return p
	}
	cx, cy := (r.Min.X+r.Max.X)*0.5, (r.Min.Y+r.Max.Y)*0.5
	rx, ry := r.W()*0.5, r.H()*0.5

	// 每段不超过 90°，用三次贝塞尔近似。
	segs := int(absf(sweep)/(3.14159265/2)) + 1
	step := sweep / float32(segs)
	// 单段贝塞尔的控制点系数：4/3 * tan(Δ/4)
	alpha := 4.0 / 3.0 * tanf(step/4)

	a0 := start
	x0 := cx + rx*cosf(a0)
	y0 := cy + ry*sinf(a0)
	if len(p.verbs) == 0 {
		p.MoveTo(x0, y0)
	} else {
		p.LineTo(x0, y0)
	}
	for i := 0; i < segs; i++ {
		a1 := a0 + step
		x1 := cx + rx*cosf(a1)
		y1 := cy + ry*sinf(a1)
		// 切向
		t0x, t0y := -rx*sinf(a0), ry*cosf(a0)
		t1x, t1y := -rx*sinf(a1), ry*cosf(a1)
		p.CubicTo(
			x0+alpha*t0x, y0+alpha*t0y,
			x1-alpha*t1x, y1-alpha*t1y,
			x1, y1,
		)
		a0, x0, y0 = a1, x1, y1
	}
	return p
}

// AddCircle 追加整圆。
func (p *Path) AddCircle(cx, cy, r float32) *Path {
	return p.AddOval(geom.RectXYWH(cx-r, cy-r, 2*r, 2*r))
}

// IsEmpty 表示路径没有任何绘制内容。
func (p *Path) IsEmpty() bool { return len(p.verbs) == 0 }

// Verbs 返回指令序列长度（诊断用）。
func (p *Path) VerbCount() int { return len(p.verbs) }

// Bounds 返回保守包围盒（用控制点，不扁平化——快且不会小于真实包围盒）。
//
// 注意：控制点包围盒可能明显大于曲线实际范围（例如大弧度圆弧），
// 所以它只适合"是否需要裁剪"的粗判；精确包围盒用 TightBounds。
func (p *Path) Bounds() geom.Rect {
	if p.computed {
		return p.bounds
	}
	p.bounds = boundPoints(p.pts)
	p.computed = true
	return p.bounds
}

// boundPoints 用最小/最大坐标累积包围盒。
//
// 不能用 Rect.Union 累积：单点矩形的宽高为 0，会被 IsEmpty 判成空，
// Union 于是不断丢弃——这正是"椭圆包围盒退化成 0×0、整条路径被裁没"的
// 根本原因（被 TestFillOval 抓住）。
func boundPoints(pts []geom.Point) geom.Rect {
	if len(pts) == 0 {
		return geom.ZeroRect
	}
	minX, minY := pts[0].X, pts[0].Y
	maxX, maxY := minX, minY
	for _, pt := range pts[1:] {
		if pt.X < minX {
			minX = pt.X
		}
		if pt.X > maxX {
			maxX = pt.X
		}
		if pt.Y < minY {
			minY = pt.Y
		}
		if pt.Y > maxY {
			maxY = pt.Y
		}
	}
	return geom.RectLTRB(minX, minY, maxX, maxY)
}

// TightBounds 返回扁平化后的精确包围盒（光栅与命中用）。
func (p *Path) TightBounds(tolerance float32) geom.Rect {
	polys := p.Flatten(tolerance, nil)
	var all []geom.Point
	for _, pl := range polys {
		all = append(all, pl.Pts...)
	}
	return boundPoints(all)
}

// Polyline 是一条扁平化后的折线。
type Polyline struct {
	Pts    []geom.Point
	Closed bool
}

// Flatten 把路径扁平化为折线，容差 tolerance 为最大弦高（DIP）。
//
// buf 可复用以减少分配；返回的折线可能引用 buf 的底层数组。
func (p *Path) Flatten(tolerance float32, buf []Polyline) []Polyline {
	if tolerance <= 0 {
		tolerance = 0.25
	}
	out := buf[:0]
	var cur *Polyline
	var last geom.Point
	var start geom.Point

	begin := func(pt geom.Point) {
		out = append(out, Polyline{Pts: make([]geom.Point, 0, 8)})
		cur = &out[len(out)-1]
		cur.Pts = append(cur.Pts, pt)
		last, start = pt, pt
	}

	// pi 是点游标：与 verbs 的下标**不**一一对应
	//（quad 吃 2 个点、cubic 吃 3 个），必须单独维护。
	pi := 0
	for _, v := range p.verbs {
		switch v {
		case verbMove:
			begin(p.pts[pi])
			pi++
		case verbLine:
			pt := p.pts[pi]
			pi++
			if cur == nil {
				begin(last)
			}
			cur.Pts = append(cur.Pts, pt)
			last = pt
		case verbQuad:
			c, e := p.pts[pi], p.pts[pi+1]
			pi += 2
			if cur == nil {
				begin(last)
			}
			cur.Pts = flattenQuad(cur.Pts, last, c, e, tolerance, 0)
			last = e
		case verbCubic:
			c1, c2, e := p.pts[pi], p.pts[pi+1], p.pts[pi+2]
			pi += 3
			if cur == nil {
				begin(last)
			}
			cur.Pts = flattenCubic(cur.Pts, last, c1, c2, e, tolerance, 0)
			last = e
		case verbClose:
			if cur != nil {
				cur.Closed = true
				cur.Pts = append(cur.Pts, start)
				last = start
				cur = nil
			}
		}
	}
	return out
}

// flattenQuad 递归细分二次贝塞尔，直到弦高小于容差。
func flattenQuad(out []geom.Point, p0, c, p1 geom.Point, tol float32, depth int) []geom.Point {
	if depth >= 10 || quadFlatEnough(p0, c, p1, tol) {
		return append(out, p1)
	}
	// de Casteljau 中点分割
	p01 := mid(p0, c)
	p12 := mid(c, p1)
	p012 := mid(p01, p12)
	out = flattenQuad(out, p0, p01, p012, tol, depth+1)
	return flattenQuad(out, p012, p12, p1, tol, depth+1)
}

// flattenCubic 递归细分三次贝塞尔。
func flattenCubic(out []geom.Point, p0, c1, c2, p1 geom.Point, tol float32, depth int) []geom.Point {
	if depth >= 10 || cubicFlatEnough(p0, c1, c2, p1, tol) {
		return append(out, p1)
	}
	p01 := mid(p0, c1)
	p12 := mid(c1, c2)
	p23 := mid(c2, p1)
	p012 := mid(p01, p12)
	p123 := mid(p12, p23)
	p0123 := mid(p012, p123)
	out = flattenCubic(out, p0, p01, p012, p0123, tol, depth+1)
	return flattenCubic(out, p0123, p123, p23, p1, tol, depth+1)
}

func mid(a, b geom.Point) geom.Point { return geom.Pt((a.X+b.X)*0.5, (a.Y+b.Y)*0.5) }

// quadFlatEnough 判定二次贝塞尔是否已足够平（控制点到弦的距离）。
func quadFlatEnough(p0, c, p1 geom.Point, tol float32) bool {
	return distToLine(c, p0, p1) <= tol
}

// cubicFlatEnough 判定三次贝塞尔是否已足够平。
func cubicFlatEnough(p0, c1, c2, p1 geom.Point, tol float32) bool {
	d1 := distToLine(c1, p0, p1)
	d2 := distToLine(c2, p0, p1)
	return d1 <= tol && d2 <= tol
}

// distToLine 返回点到线段所在直线的距离。
func distToLine(p, a, b geom.Point) float32 {
	d := b.Sub(a)
	l := d.Len()
	if l == 0 {
		return p.Sub(a).Len()
	}
	// 叉积 / 长度
	cross := d.X*(p.Y-a.Y) - d.Y*(p.X-a.X)
	return absf(cross) / l
}

// Transform 返回应用矩阵后的新路径（不修改自身）。
func (p *Path) Transform(m geom.Matrix) *Path {
	out := &Path{verbs: make([]verb, len(p.verbs)), pts: make([]geom.Point, len(p.pts))}
	copy(out.verbs, p.verbs)
	for i, pt := range p.pts {
		out.pts[i] = m.MapPoint(pt)
	}
	return out
}

// Clone 深拷贝路径（side table 复用同一路径时防止意外修改）。
func (p *Path) Clone() *Path {
	out := &Path{verbs: make([]verb, len(p.verbs)), pts: make([]geom.Point, len(p.pts))}
	copy(out.verbs, p.verbs)
	copy(out.pts, p.pts)
	out.bounds, out.computed = p.bounds, p.computed
	return out
}

func minf(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func clampf(v, lo, hi float32) float32 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func absf(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}
