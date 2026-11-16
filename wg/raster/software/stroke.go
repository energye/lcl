package software

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
)

// 描边实现策略：stroke-to-fill（把描边转成多边形再按填充规则画）。
//
// 为什么不做"沿扫描线求带符号距离"的解析描边：
//   - 解析描边对自交、尖角、虚线的处理非常复杂（需要距离场与平局规则）；
//   - stroke-to-fill 复用同一条填充路径，AA 质量与填充完全一致；
//   - 组合出的多边形用 nonzero 规则求并，天然处理"线段重叠"。
//
// 代价是拐角处多边形会重叠，靠 nonzero 规则合并——所以每个多边形都必须
// 保证方向一致（统一成逆时针），否则环绕数会抵消出空洞。
// 这就是 appendPoly 里那道 signedArea 检查的原因，删掉它会导致
// "斜线交叉处出现透明缝"。

// strokePolygons 把折线集转成描边多边形集（局部坐标，未变换）。
func strokePolygons(polys []render.Polyline, paint *render.Paint, out [][]geom.Point) [][]geom.Point {
	width := paint.StrokeWidthOr()
	if width <= 0 {
		return out
	}
	hw := width * 0.5
	dash := paint.Dash
	if len(dash) > 0 {
		polys = applyDash(polys, dash, paint.DashPhase)
	}
	for _, pl := range polys {
		out = strokeOne(pl, hw, paint.Cap, paint.Join, paint.MiterLimit, out)
	}
	return out
}

// strokeOne 处理单条折线。
func strokeOne(pl render.Polyline, hw float32, capStyle render.Cap, join render.Join, miterLimit float32, out [][]geom.Point) [][]geom.Point {
	pts := pl.Pts
	// 去掉连续重复点，避免零长度线段产生 NaN 方向
	pts = dedup(pts)
	n := len(pts)
	if n == 0 {
		return out
	}
	if n == 1 {
		// 退化成一个点：圆头/方头画点，平头不画
		if capStyle == render.CapRound {
			out = appendCircle(out, pts[0], hw, 12)
		}
		return out
	}
	if miterLimit <= 0 {
		miterLimit = 4
	}

	// 线段
	segCount := n - 1
	if pl.Closed && n > 2 {
		segCount = n
	}
	for i := 0; i < segCount; i++ {
		a := pts[i]
		b := pts[(i+1)%n]
		out = appendSegment(out, a, b, hw)
	}

	// 拐角
	joins := n - 2
	if pl.Closed && n > 2 {
		joins = n
	}
	for i := 0; i < joins; i++ {
		prev := pts[i%n]
		cur := pts[(i+1)%n]
		next := pts[(i+2)%n]
		if pl.Closed && n == 2 {
			break
		}
		out = appendJoin(out, prev, cur, next, hw, join, miterLimit)
	}

	// 端帽（闭合折线没有端帽）
	if !pl.Closed {
		out = appendCap(out, pts[0], pts[1], hw, capStyle, true)
		out = appendCap(out, pts[n-1], pts[n-2], hw, capStyle, false)
	}
	return out
}

// appendSegment 追加一条线段对应的四边形。
func appendSegment(out [][]geom.Point, a, b geom.Point, hw float32) [][]geom.Point {
	d := b.Sub(a)
	l := d.Len()
	if l == 0 {
		return out
	}
	d = d.Mul(1 / l)
	nrm := geom.Pt(-d.Y*hw, d.X*hw)
	return appendPoly(out, []geom.Point{
		a.Add(nrm), b.Add(nrm), b.Sub(nrm), a.Sub(nrm),
	})
}

// appendJoin 追加拐角补丁。
//
// 两侧都补：内侧虽然被线段四边形盖住，但多补一个三角形不会有副作用
// （nonzero 下重叠区域环绕数更大，仍然是填充），却省掉了"判断哪一侧是外侧"
// 这个容易出符号错误的分支。
func appendJoin(out [][]geom.Point, prev, cur, next geom.Point, hw float32, join render.Join, miterLimit float32) [][]geom.Point {
	d0 := cur.Sub(prev)
	d1 := next.Sub(cur)
	l0, l1 := d0.Len(), d1.Len()
	if l0 == 0 || l1 == 0 {
		return out
	}
	d0 = d0.Mul(1 / l0)
	d1 = d1.Mul(1 / l1)

	if join == render.JoinRound {
		// 圆角：一个圆覆盖面两侧的所有缺口
		return appendCircle(out, cur, hw, circleSegs(hw))
	}

	n0 := geom.Pt(-d0.Y*hw, d0.X*hw)
	n1 := geom.Pt(-d1.Y*hw, d1.X*hw)
	for _, s := range []float32{1, -1} {
		a := cur.Add(n0.Mul(s))
		b := cur.Add(n1.Mul(s))
		if join == render.JoinMiter {
			if p, ok := lineIntersect(a, d0, b, d1); ok {
				if p.Sub(cur).Len() <= miterLimit*hw {
					out = appendPoly(out, []geom.Point{a, p, b, cur})
					continue
				}
			}
		}
		// 斜切兜底
		out = appendPoly(out, []geom.Point{cur, a, b})
	}
	return out
}

// appendCap 追加端帽。atStart 表示这是起点帽（方向取反向）。
func appendCap(out [][]geom.Point, end, inner geom.Point, hw float32, capStyle render.Cap, atStart bool) [][]geom.Point {
	switch capStyle {
	case render.CapRound:
		return appendCircle(out, end, hw, circleSegs(hw))
	case render.CapSquare:
		d := end.Sub(inner)
		l := d.Len()
		if l == 0 {
			return out
		}
		d = d.Mul(1 / l)
		if atStart {
			d = d.Mul(-1) // 起点帽向外延伸
		}
		nrm := geom.Pt(-d.Y*hw, d.X*hw)
		return appendPoly(out, []geom.Point{
			end.Add(nrm), end.Add(nrm).Add(d.Mul(hw)),
			end.Sub(nrm).Add(d.Mul(hw)), end.Sub(nrm),
		})
	default: // CapButt
		return out
	}
}

// appendCircle 追加正多边形近似的圆。
func appendCircle(out [][]geom.Point, c geom.Point, r float32, segs int) [][]geom.Point {
	if segs < 6 {
		segs = 6
	}
	pts := make([]geom.Point, segs)
	for i := 0; i < segs; i++ {
		ang := 6.283185307179586 * float32(i) / float32(segs)
		pts[i] = geom.Pt(c.X+r*cosf(ang), c.Y+r*sinf(ang))
	}
	return appendPoly(out, pts)
}

// appendPoly 追加多边形，并保证逆时针（正面积）方向。
func appendPoly(out [][]geom.Point, pts []geom.Point) [][]geom.Point {
	if len(pts) < 3 {
		return out
	}
	if signedArea(pts) < 0 {
		for i, j := 0, len(pts)-1; i < j; i, j = i+1, j-1 {
			pts[i], pts[j] = pts[j], pts[i]
		}
	}
	return append(out, pts)
}

// signedArea 返回多边形有符号面积（逆时针为正，y 向下的坐标系里符号相反，
// 但这里只用来判方向是否一致，绝对值不重要）。
func signedArea(pts []geom.Point) float32 {
	a := float32(0)
	for i := 0; i < len(pts); i++ {
		p := pts[i]
		q := pts[(i+1)%len(pts)]
		a += p.X*q.Y - q.X*p.Y
	}
	return a * 0.5
}

// lineIntersect 求两条直线交点（p + t*d 形式）。
func lineIntersect(p, dp, q, dq geom.Point) (geom.Point, bool) {
	den := dp.X*dq.Y - dp.Y*dq.X
	if absf(den) < 1e-6 {
		return geom.Point{}, false
	}
	t := ((q.X-p.X)*dq.Y - (q.Y-p.Y)*dq.X) / den
	return p.Add(dp.Mul(t)), true
}

// applyDash 按虚线模式切分折线。
func applyDash(polys []render.Polyline, dash []float32, phase float32) []render.Polyline {
	// 只支持偶数长度模式（实/空成对）；奇数长度按"重复一遍"补全（SVG 语义）
	if len(dash) == 0 {
		return polys
	}
	pattern := dash
	if len(pattern)%2 == 1 {
		cp := make([]float32, 0, len(pattern)*2)
		cp = append(cp, pattern...)
		cp = append(cp, pattern...)
		pattern = cp
	}
	total := float32(0)
	for _, v := range pattern {
		if v <= 0 {
			return polys // 非法模式，退化成实线
		}
		total += v
	}
	if total <= 0 {
		return polys
	}
	// 相位归一到 [0,total)
	ph := phase
	for ph < 0 {
		ph += total
	}
	for ph >= total {
		ph -= total
	}

	var out []render.Polyline
	for _, pl := range polys {
		pts := pl.Pts
		if len(pts) < 2 {
			continue
		}
		// 沿路径累计长度，按模式取实段
		idx := 0
		remain := pattern[0]
		on := true
		// 消耗相位
		for ph > 0 {
			if ph >= remain {
				ph -= remain
				idx = (idx + 1) % len(pattern)
				remain = pattern[idx]
				on = !on
				continue
			}
			remain -= ph
			ph = 0
		}
		cur := []geom.Point{pts[0]}
		for i := 1; i < len(pts); i++ {
			a, b := pts[i-1], pts[i]
			segLen := b.Sub(a).Len()
			travelled := float32(0)
			for travelled < segLen {
				step := remain
				if step > segLen-travelled {
					step = segLen - travelled
				}
				travelled += step
				remain -= step
				pt := a.Add(b.Sub(a).Mul(travelled / segLen))
				if on {
					cur = append(cur, pt)
				}
				if remain <= 0 {
					if on && len(cur) >= 2 {
						out = append(out, render.Polyline{Pts: cur})
					}
					cur = []geom.Point{pt}
					idx = (idx + 1) % len(pattern)
					remain = pattern[idx]
					on = !on
				}
			}
		}
		if on && len(cur) >= 2 {
			out = append(out, render.Polyline{Pts: cur})
		}
	}
	return out
}

// dedup 去掉连续重复点。
func dedup(pts []geom.Point) []geom.Point {
	if len(pts) < 2 {
		return pts
	}
	out := pts[:1]
	for _, p := range pts[1:] {
		last := out[len(out)-1]
		if absf(p.X-last.X) > 1e-5 || absf(p.Y-last.Y) > 1e-5 {
			out = append(out, p)
		}
	}
	// 闭合折线首尾重合时去掉尾点（避免零长度段）
	if len(out) > 2 {
		first := out[0]
		last := out[len(out)-1]
		if absf(first.X-last.X) <= 1e-5 && absf(first.Y-last.Y) <= 1e-5 {
			out = out[:len(out)-1]
		}
	}
	return out
}

// circleSegs 按半径决定圆的分段数：半径越大越细，避免大圆看出棱角。
func circleSegs(r float32) int {
	n := int(r * 6)
	if n < 8 {
		n = 8
	}
	if n > 48 {
		n = 48
	}
	return n
}

func absf(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}
