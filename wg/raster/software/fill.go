package software

import (
	"sort"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
)

// coverage 使用 8 条子扫描线做垂直采样，水平方向用解析覆盖率。
//
// 为什么混合而不是全超采样：
//   - 纯 4×4/8×8 超采样在近水平边仍会出现台阶，要消除得 16×16（16 倍成本）；
//   - 水平解析覆盖率对近垂直边是精确的（这正是细竖线、文字边缘的主要场景）；
//   - 垂直 8 级对近水平边给出 1/8 的量化，配 1px 级图元视觉上足够。
//
// 这是 03-ENGINE.md §7 里"按 tile 做覆盖率"的务实版本：M1 先按整行做，
// 等 M2 有真实场景数据再决定是否引入 tile 与代际标记。
const subSamples = 8

// edge 是一条设备空间的填充边（已按 y 升序归一）。
type edge struct {
	x0, y0 float32 // 上端点
	x1, y1 float32 // 下端点
	dir    int32   // +1 向下（increasing y），-1 向上
	dxdy   float32 // dx/dy
}

// filler 是可复用的扫描线填充器。
//
// 三层缓冲各自按帧复用，热路径零分配：
//
//	edges  —— 全部边（一次 fill 内的所有轮廓）
//	order  —— 按 y0 升序的边下标（活跃边表的上界游标）
//	active —— 当前行真正与扫描线相交的边下标
//	cov    —— 行覆盖率（宽 + 2）
type filler struct {
	edges  []edge
	order  []int32
	active []int32
	xs     []float32
	dirs   []int32
	cov    []uint16
	// wmul 是单次跨度写入覆盖率缓冲的权重：普通行 1（sub 条子扫描线各加一次），
	// 相干行 sub（一次写满，跳过其余 sub-1 次累加）。
	wmul   uint16
	aa     bool
	stats  raster.Stats
	pixTol float32

	// mask 是路径裁剪蒙版（nil 表示无）；maskW 是它的行宽。
	// 只在有 OpClipPath 的指令流里非 nil，热路径用一次判空挡住。
	mask  []uint8
	maskW int
}

func newFiller() *filler {
	return &filler{wmul: 1, aa: true}
}

// maskCoverage 取蒙版覆盖率；越界视为 0（裁掉了）。
func maskCoverage(m []uint8, w, x, y int) uint32 {
	if w <= 0 {
		return 0
	}
	i := y*w + x
	if i < 0 || i >= len(m) {
		return 0
	}
	return uint32(m[i])
}

func (f *filler) reset() {
	f.edges = f.edges[:0]
	f.stats = raster.Stats{}
}

// addPolyline 追加一条折线的边（自动闭合）。
func (f *filler) addPolyline(pts []geom.Point) {
	n := len(pts)
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ {
		a := pts[i]
		b := pts[(i+1)%n]
		f.addEdge(a, b)
	}
}

// addEdge 追加一条边，跳过水平线（对扫描线无贡献）。
func (f *filler) addEdge(a, b geom.Point) {
	if a.Y == b.Y {
		return
	}
	dir := int32(1)
	if a.Y > b.Y {
		a, b = b, a
		dir = -1
	}
	f.edges = append(f.edges, edge{
		x0: a.X, y0: a.Y, x1: b.X, y1: b.Y,
		dir:  dir,
		dxdy: (b.X - a.X) / (b.Y - a.Y),
	})
}

// buildOrder 生成按 y0 升序的边下标表（活跃边表的上界）。
//
// 用插入排序而不是 sort.Slice：边数在千级以内，且边按折线顺序追加时
// y0 已近似递增，插入排序在这种输入上接近 O(n)，并且**零分配**
//（sort.Slice 会为 Swapper 分配）。
func (f *filler) buildOrder() {
	f.order = f.order[:0]
	for i := range f.edges {
		f.order = append(f.order, int32(i))
	}
	// 按 y0 排序。边数少（<64，矩形/圆角矩形等常规图元）用插入排序
	// （缓存友好、无闭包开销）；多时必须 O(n log n)——progress 圆环轮
	// 实测：4.4 万条描边边的弧线用插入排序一次 fill 1.2s（占总耗时 85%），
	// 换 sort.Slice 后降到 ~4ms。
	if len(f.order) < 64 {
		for i := 1; i < len(f.order); i++ {
			v := f.order[i]
			vy := f.edges[v].y0
			j := i - 1
			for j >= 0 && f.edges[f.order[j]].y0 > vy {
				f.order[j+1] = f.order[j]
				j--
			}
			f.order[j+1] = v
		}
		return
	}
	sort.Slice(f.order, func(a, b int) bool {
		return f.edges[f.order[a]].y0 < f.edges[f.order[b]].y0
	})
}

// fill 用给定颜色/渐变填充已累积的边，写入 dst 的 clip 区域。
//
// 活跃边表是这里的核心：每条边只在它真正跨越的行里参与扫描线求交。
// 没有活跃边表时，每行每子扫描线都要遍历**全部**边——描边轮廓动辄
// 三四百条边，而一行通常只有个位数条真正相交，这是 M1 首轮 --bench
// 只有 30 FPS 的主因（见 05-API-ROADMAP.md §12）。
func (f *filler) fill(dst *raster.Backing, clip geom.Rect, rule render.FillRule, paint *render.Paint, inv geom.Matrix) {
	if len(f.edges) == 0 {
		return
	}
	// 计算边的包围盒并与裁剪求交，避免遍历整幅图。
	minY, maxY := f.edges[0].y0, f.edges[0].y1
	minX, maxX := f.edges[0].x0, f.edges[0].x0
	for _, e := range f.edges {
		if e.y0 < minY {
			minY = e.y0
		}
		if e.y1 > maxY {
			maxY = e.y1
		}
		if e.x0 < minX {
			minX = e.x0
		}
		if e.x0 > maxX {
			maxX = e.x0
		}
		if e.x1 < minX {
			minX = e.x1
		}
		if e.x1 > maxX {
			maxX = e.x1
		}
	}
	box := geom.RectLTRB(minX, minY, maxX, maxY).Intersect(clip)
	if box.IsEmpty() {
		return
	}
	x0 := int(box.Min.X)
	x1 := int(box.Max.X) + 1
	y0 := int(box.Min.Y)
	y1 := int(box.Max.Y) + 1
	if x0 < 0 {
		x0 = 0
	}
	if y0 < 0 {
		y0 = 0
	}
	if x1 > dst.W() {
		x1 = dst.W()
	}
	if y1 > dst.H() {
		y1 = dst.H()
	}
	if x0 >= x1 || y0 >= y1 {
		return
	}

	// 覆盖率缓冲按行分配一次（宽度 + 2 给跨边界像素留余量）。
	w := x1 - x0
	if cap(f.cov) < w+2 {
		f.cov = make([]uint16, w+2)
	}
	f.cov = f.cov[:w+2]

	opaque := paint != nil && paint.IsOpaque()
	hasShader := paint != nil && paint.Shader != nil
	basePremul := uint32(0)
	sr0, sg0, sb0, sa0 := uint32(0), uint32(0), uint32(0), uint32(0)
	if !hasShader {
		basePremul = paint.Color.PremulBGRA()
		sr0 = (basePremul >> 16) & 0xFF
		sg0 = (basePremul >> 8) & 0xFF
		sb0 = basePremul & 0xFF
		sa0 = basePremul >> 24
	}
	// sub 必须是编译期常量：热循环里 c/sub 才会被折成移位，
	// 用字段会让每像素付一次整数除法（实测一帧多出 1.5ms 以上）。
	invSub := 1.0 / float32(subSamples)

	f.buildOrder()
	f.active = f.active[:0]
	next := 0

	for y := y0; y < y1; y++ {
		rowTop := float32(y)
		rowBot := rowTop + 1

		// ① 增补：上端点落在本行以下的边全部进入活跃表
		for next < len(f.order) {
			ei := f.order[next]
			if f.edges[ei].y0 >= rowBot {
				break
			}
			f.active = append(f.active, ei)
			next++
		}
		// ② 剔除：下端点已不低于行顶的边离开活跃表
		if n := len(f.active); n > 0 {
			k := 0
			for _, ei := range f.active {
				if f.edges[ei].y1 > rowTop {
					f.active[k] = ei
					k++
				}
			}
			f.active = f.active[:k]
		}
		if len(f.active) < 2 {
			continue // 少于两条边不可能构成跨度
		}

		// 清空本行覆盖率
		for i := range f.cov {
			f.cov[i] = 0
		}

		// 相干行判定：整行里没有斜边，也没有边端点落在行内，
		// 那么 sub 条子扫描线求出的交点完全相同——算一次就够。
		// 矩形与圆角矩形的内部行全部属于这一类，能省掉 7/8 的覆盖率累加。
		coherent := true
		for _, ei := range f.active {
			e := &f.edges[ei]
			if e.dxdy != 0 || e.y0 > rowTop || e.y1 < rowBot {
				coherent = false
				break
			}
		}

		rowTouched := false
		if coherent {
			f.wmul = uint16(subSamples)
			rowTouched = f.spans(rowTop+0.5, rule, x0, w)
			f.wmul = 1
		} else {
			for k := 0; k < subSamples; k++ {
				sy := rowTop + (float32(k)+0.5)*invSub
				if f.spans(sy, rule, x0, w) {
					rowTouched = true
				}
			}
		}

		if !rowTouched {
			continue
		}
		// 混合整行
		dstRow := dst.Row(y)
		for i := 0; i < w; i++ {
			c := f.cov[i]
			if c == 0 {
				continue
			}
			a := uint32(c) / uint32(subSamples)
			if !f.aa {
				a = 255
			}
			if a == 0 {
				continue
			}
			px := x0 + i
			// 路径裁剪：覆盖率再乘蒙版（矩形裁剪不会走到这里，mask 为 nil）。
			if f.mask != nil {
				ma := maskCoverage(f.mask, f.maskW, px, y)
				if ma == 0 {
					continue
				}
				if ma != 255 {
					a = a * ma / 255
					if a == 0 {
						continue
					}
				}
			}
			if hasShader {
				// 渐变：逐像素采样（设备中心 → 用户空间）
				lp := inv.MapPoint(geom.Point{X: float32(px) + 0.5, Y: rowTop + 0.5})
				col := paint.Shader.Sample(paint.Shader.GradientT(lp))
				cr, cg, cb, ca := premulScaled(col, a)
				dstRow[px] = blendPremul(dstRow[px], cr, cg, cb, ca)
				f.stats.Pixels++
				continue
			}
			if a == 255 {
				if opaque {
					dstRow[px] = basePremul
				} else {
					dstRow[px] = blendPremul(dstRow[px], sr0, sg0, sb0, sa0)
				}
				f.stats.Pixels++
				continue
			}
			dstRow[px] = blendPremul(dstRow[px],
				sr0*a/255, sg0*a/255, sb0*a/255, sa0*a/255)
			f.stats.Pixels++
		}
		f.stats.Spans++
	}
}

// spans 在一条扫描线上求交、排序、按填充规则生成跨度并累加覆盖率。
// 返回 true 表示本行有跨度被写入。
func (f *filler) spans(sy float32, rule render.FillRule, x0, w int) bool {
	// 收集交点（只在活跃边里）
	f.xs = f.xs[:0]
	f.dirs = f.dirs[:0]
	for _, ei := range f.active {
		e := &f.edges[ei]
		if sy < e.y0 || sy >= e.y1 {
			continue
		}
		f.xs = append(f.xs, e.x0+(sy-e.y0)*e.dxdy)
		f.dirs = append(f.dirs, e.dir)
	}
	if len(f.xs) < 2 {
		return false
	}
	// 排序：交点数少（<32）用插入排序（常数小），多时用 Go 排序
	// O(k log k)——弧线描边一行交点可达数百，插入排序 O(k²) 是
	// progress 圆环首帧 2.3s 的主因之一（实测 360° 弧 fill 1.7s）。
	if len(f.xs) < 32 {
		for i := 1; i < len(f.xs); i++ {
			x, d := f.xs[i], f.dirs[i]
			j := i - 1
			for j >= 0 && f.xs[j] > x {
				f.xs[j+1], f.dirs[j+1] = f.xs[j], f.dirs[j]
				j--
			}
			f.xs[j+1], f.dirs[j+1] = x, d
		}
	} else {
		idx := make([]int, len(f.xs))
		for i := range idx {
			idx[i] = i
		}
		sort.Slice(idx, func(a, b int) bool { return f.xs[idx[a]] < f.xs[idx[b]] })
		xs2 := make([]float32, len(f.xs))
		dirs2 := make([]int32, len(f.dirs))
		for i, id := range idx {
			xs2[i], dirs2[i] = f.xs[id], f.dirs[id]
		}
		copy(f.xs, xs2)
		copy(f.dirs, dirs2)
	}
	// 按填充规则生成跨度
	touched := false
	winding := int32(0)
	spanStart := float32(0)
	inSpan := false
	for i, x := range f.xs {
		var inside bool
		if rule == render.EvenOdd {
			winding++ // 奇偶规则只数交叉次数
			inside = winding&1 == 1
		} else {
			winding += f.dirs[i]
			inside = winding != 0
		}
		if !inSpan && inside {
			spanStart = x
			inSpan = true
			continue
		}
		if inSpan && !inside {
			f.accumulate(spanStart, x, x0, w)
			inSpan = false
			touched = true
		}
	}
	// 收尾：未闭合的跨度（理论上闭合路径不会出现）延伸到裁剪右边界
	if inSpan {
		f.accumulate(spanStart, float32(x0+w), x0, w)
		touched = true
	}
	return touched
}

// accumulate 把一个跨度累加进覆盖率缓冲。
//
// 水平方向用解析覆盖率（跨度与像素的交叠长度），因此细竖线的边缘是精确的。
//
// 归一化约定：**每次子扫描线贡献 0..255**，全部子扫描线累加后再除以 sub。
// 不要在这里预先除 sub——那会和外面的除法叠加，把整幅图压成 1/sub 的亮度
//（被 TestFillOval 抓过一次：alpha 恒为 31）。
// f.wmul 只在相干行里取 sub，用来一次写满补偿被跳过的那 sub-1 条子扫描线。
func (f *filler) accumulate(xa, xb float32, x0, w int) {
	if xb <= xa {
		return
	}
	f.stats.Spans++
	// 转到缓冲下标
	xa -= float32(x0)
	xb -= float32(x0)
	if xb <= 0 || xa >= float32(w) {
		return
	}
	if xa < 0 {
		xa = 0
	}
	if xb > float32(w) {
		xb = float32(w)
	}
	i0 := int(xa)
	i1 := int(xb)
	m := f.wmul
	if i0 == i1 {
		// 落在单个像素内
		f.cov[i0] += uint16((xb-xa)*255*float32(m) + 0.5)
		return
	}
	// 首像素
	f.cov[i0] += uint16((float32(i0+1)-xa)*255*float32(m) + 0.5)
	// 中间整像素
	full := 255 * m
	for i := i0 + 1; i < i1; i++ {
		f.cov[i] += full
	}
	// 末像素
	if i1 < w {
		f.cov[i1] += uint16((xb-float32(i1))*255*float32(m) + 0.5)
	}
}

// premulScaled 把颜色预乘并按覆盖率缩放。
func premulScaled(c geom.Color, cov uint32) (sr, sg, sb, sa uint32) {
	if c.A == 0 || cov == 0 {
		return 0, 0, 0, 0
	}
	a := uint32(c.A) * cov / 255
	sr = uint32(c.R) * a / 255
	sg = uint32(c.G) * a / 255
	sb = uint32(c.B) * a / 255
	sa = a
	return
}

// blendPremul 做 source-over 混合（源与目标都是预乘）。
func blendPremul(dst, sr, sg, sb, sa uint32) uint32 {
	if sa == 0 {
		return dst
	}
	if sa >= 255 {
		return 0xFF000000 | sr<<16 | sg<<8 | sb
	}
	inv := 255 - sa
	da := dst >> 24
	dr := (dst >> 16) & 0xFF
	dg := (dst >> 8) & 0xFF
	db := dst & 0xFF
	oa := sa + da*inv/255
	orr := sr + dr*inv/255
	og := sg + dg*inv/255
	ob := sb + db*inv/255
	if oa > 255 {
		oa = 255
	}
	return oa<<24 | orr<<16 | og<<8 | ob
}
