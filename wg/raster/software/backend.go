// Package software 是默认光栅后端：纯 Go 扫描线覆盖率抗锯齿。
//
// 它不 import lcl，可以在无窗口、无图形设备的环境下出图，
// 因此 golden 比对、契约自检、CI 全部跑在它上面。
//
// 能力边界（诚实上报，缺失的指令由上层计 degraded，见 C22）：
//   - 支持：填充/描边/虚线、线性/径向渐变、矩形裁剪、任意仿射变换；
//   - 不支持：路径裁剪（降级为包围盒）、位图绘制（M2）、扫描渐变（M4）。
//
// 文字不在本后端：字形回放必须由 GDI 在像素写回后于设备空间完成，
// 才能保留 ClearType（C6/C24）。文字指令由 composite 的文字通道执行。
package software

import (
	"math"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
)

// Options 是后端选项。
type Options struct {
	// DisableAA 关闭抗锯齿（A/B 对照用，对应 WG_AA=0）。
	DisableAA bool
	// Tolerance 是曲线扁平化容差（设备像素）；0 表示 0.25。
	Tolerance float32
}

// Backend 是软件光栅后端。
//
// 它持有可复用的扫描线缓冲与折线缓冲，因此**不是并发安全的**：
// 一个 Backend 只能被一个光栅序列使用（当前等于"一个窗口一帧"）。
type Backend struct {
	aa  bool
	tol float32

	fill filler

	polys   []render.Polyline
	outline [][]geom.Point
	// scratch 是"变换后的折线"暂存缓冲：addPolyline 只读取点、不保留引用，
	// 因此同一条折线可以整批复用，避免每条折线都 make 一个新切片。
	scratch []geom.Point

	// textOps 记录本次绘制遇到的文字指令数（由文字通道执行，不计 degraded）。
	textOps int

	// mask 是路径裁剪蒙版（nil 表示无路径裁剪）；maskW 是行宽。
	// 矩形裁剪走 scissor，不进这里——只有 OpClipPath 才会建蒙版。
	mask  []uint8
	maskW int
	// maskStack 保存 Save/Restore 之间的蒙版（nil 也要能压栈）。
	maskStack [][]uint8
}

// New 创建软件后端。
func New(opts Options) *Backend {
	tol := opts.Tolerance
	if tol <= 0 {
		tol = 0.25
	}
	b := &Backend{aa: !opts.DisableAA, tol: tol}
	b.fill = *newFiller()
	b.fill.aa = b.aa
	return b
}

// Name 实现 raster.Backend。
func (b *Backend) Name() string { return "software" }

// Caps 实现 raster.Backend。
func (b *Backend) Caps() raster.Caps {
	return raster.Caps{
		AA:             b.aa,
		Stroke:         true,
		Dash:           true,
		LinearGradient: true,
		RadialGradient: true,
		ClipRect:       true,
		ClipPath:       false, // M1 降级为包围盒
		Text:           false, // 文字走 GDI 文字通道
		Image:          true,
		Transform:      true,
	}
}

// TextOps 返回上次 Draw 遇到的文字指令数。
func (b *Backend) TextOps() int { return b.textOps }

// CoverageBytes 返回覆盖率缓冲占用的字节数。
//
// 这是契约 C5"AA 就地混色、不额外分配整帧 alpha 缓冲"的可测量证据：
// 按行实现时该值约为 2×行宽，按整帧实现时会等于 2×宽×高。
func (b *Backend) CoverageBytes() int { return cap(b.fill.cov) * 2 }

// Draw 实现 raster.Backend。
func (b *Backend) Draw(dst *raster.Backing, dl *render.DisplayList, clip geom.Rect) raster.Stats {
	var total raster.Stats
	if dst == nil || dl == nil || dl.Empty() {
		return total
	}
	b.textOps = 0

	// 裁剪栈：指令里的矩阵是绝对值，裁剪也按设备空间逐步求交。
	var clips []geom.Rect
	cur := clip
	if cur.IsEmpty() {
		cur = dst.Bounds()
	}

	for _, cmd := range dl.Commands() {
		total.Ops++
		switch cmd.Op {
		case render.OpSave:
			clips = append(clips, cur)
			b.maskStack = append(b.maskStack, b.mask)
		case render.OpRestore:
			if n := len(clips); n > 0 {
				cur = clips[n-1]
				clips = clips[:n-1]
			}
			if n := len(b.maskStack); n > 0 {
				b.mask = b.maskStack[n-1]
				b.maskStack = b.maskStack[:n-1]
			}
		case render.OpConcat:
			// 矩阵已固化在每条指令里，这里无需处理

		case render.OpClipRect:
			cur = cur.Intersect(cmd.Rect)
		case render.OpClipPath:
			// 包围盒先收 scissor（省掉大量无用像素），再叠一层精确路径蒙版。
			cur = cur.Intersect(cmd.Rect)
			if p := dl.Path(cmd.Path); p != nil {
				b.mask = b.buildPathMask(p, cmd.Matrix, dst.W(), dst.H(), cur)
				b.maskW = dst.W()
			}

		case render.OpFillRect:
			paint := dl.Paint(cmd.Paint)
			if paint == nil {
				continue
			}
			st := b.fillRect(dst, cmd.Rect, paint, cur)
			total.Add(st)

		case render.OpFillPath:
			path := dl.Path(cmd.Path)
			paint := dl.Paint(cmd.Paint)
			if path == nil || paint == nil {
				total.Degraded++
				continue
			}
			st := b.fillPath(dst, path, cmd.Matrix, paint, cur)
			total.Add(st)

		case render.OpStrokePath:
			path := dl.Path(cmd.Path)
			paint := dl.Paint(cmd.Paint)
			if path == nil || paint == nil {
				total.Degraded++
				continue
			}
			st := b.strokePath(dst, path, cmd.Matrix, paint, cur)
			total.Add(st)

		case render.OpDrawText:
			b.textOps++

		case render.OpDrawImage:
			img := dl.Image(cmd.Image)
			if img == nil {
				total.Degraded++
				continue
			}
			st := b.drawImage(dst, img, cmd.Rect, cmd.Quality, cur)
			total.Add(st)
		}
	}
	return total
}

// fillPath 填充路径。
func (b *Backend) fillPath(dst *raster.Backing, path *render.Path, m geom.Matrix, paint *render.Paint, clip geom.Rect) raster.Stats {
	scale := m.ScaleFactor()
	polys := path.Flatten(b.tol/maxf(scale, 0.01), b.polys)
	b.polys = polys

	b.fill.reset()
	b.fill.aa = b.aa && paint.AntiAlias
	b.fill.mask, b.fill.maskW = b.mask, b.maskW
	for _, pl := range polys {
		b.fill.addPolyline(b.transform(pl.Pts, m))
	}
	inv, ok := m.Invert()
	if !ok {
		inv = geom.Identity
	}
	b.fill.fill(dst, clip, paint.FillRule, paint, inv)
	return b.fill.stats
}

// strokePath 描边路径（stroke-to-fill）。
func (b *Backend) strokePath(dst *raster.Backing, path *render.Path, m geom.Matrix, paint *render.Paint, clip geom.Rect) raster.Stats {
	scale := m.ScaleFactor()
	polys := path.Flatten(b.tol/maxf(scale, 0.01), b.polys)
	b.polys = polys

	b.outline = strokePolygons(polys, paint, b.outline[:0])

	b.fill.reset()
	b.fill.aa = b.aa && paint.AntiAlias
	b.fill.mask, b.fill.maskW = b.mask, b.maskW
	for _, poly := range b.outline {
		b.fill.addPolyline(b.transform(poly, m))
	}
	// 描边一律 nonzero：线段四边形与拐角补丁会重叠，靠环绕数求并
	inv, ok := m.Invert()
	if !ok {
		inv = geom.Identity
	}
	b.fill.fill(dst, clip, render.NonZero, paint, inv)
	return b.fill.stats
}

// fillRect 轴对齐矩形填充：解析覆盖率，比走扫描线更精确也更快。
//
// 关键优化：把矩形拆成**内接整像素块**与**边带**两部分。
// 内接块里每个像素的覆盖率恒为 1，逐像素解析覆盖率是纯浪费——
// 一屏 1440×880 的清屏如果走逐像素路径要 6.4ms，批量铺只要 0.6ms，
// 而清屏正是每帧面积最大的一条指令（见 05-API-ROADMAP.md §12）。
// 边带（左/右各 ≤1 列、上/下各 ≤1 行）仍然逐像素算，保证亚像素精度不变。
func (b *Backend) fillRect(dst *raster.Backing, r geom.Rect, paint *render.Paint, clip geom.Rect) raster.Stats {
	r = r.Intersect(clip)
	if r.IsEmpty() {
		return raster.Stats{}
	}
	x0 := int(floorf(r.Min.X))
	x1 := int(ceilf(r.Max.X))
	y0 := int(floorf(r.Min.Y))
	y1 := int(ceilf(r.Max.Y))
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
		return raster.Stats{Ops: 1}
	}

	// 有路径蒙版时不能走"批量铺内接块"的快速路径：逐像素覆盖率要乘蒙版。
	if b.mask != nil {
		p := render.NewPath().AddRect(r)
		return b.fillPath(dst, p, geom.Identity, paint, clip)
	}

	var st raster.Stats
	shader := paint.Shader
	base := paint.Color.PremulBGRA()
	opaque := paint.IsOpaque()
	sr0 := (base >> 16) & 0xFF
	sg0 := (base >> 8) & 0xFF
	sb0 := base & 0xFF
	sa0 := base >> 24

	// 内接整像素范围 [ix0,ix1)×[iy0,iy1)：这些像素完全落在矩形内部，覆盖率 1。
	ix0, ix1 := x0, x1
	iy0, iy1 := y0, y1
	if b.aa && paint.AntiAlias {
		ix0 = clampInt(int(ceilf(r.Min.X)), x0, x1)
		ix1 = clampInt(int(floorf(r.Max.X)), x0, x1)
		iy0 = clampInt(int(ceilf(r.Min.Y)), y0, y1)
		iy1 = clampInt(int(floorf(r.Max.Y)), y0, y1)
	}

	if iy0 < iy1 && ix0 < ix1 {
		n := ix1 - ix0
		for y := iy0; y < iy1; y++ {
			row := dst.Row(y)
			switch {
			case shader != nil:
				// 渐变横向变化，只能逐像素采样
				for x := ix0; x < ix1; x++ {
					lp := geom.Pt(float32(x)+0.5, float32(y)+0.5)
					col := shader.Sample(shader.GradientT(lp))
					cr, cg, cb, ca := premulScaled(col, 255)
					row[x] = blendPremul(row[x], cr, cg, cb, ca)
				}
			case opaque || sa0 == 255:
				// 覆盖满、颜色不透明 → 纯存储，无读改写
				for x := ix0; x < ix1; x++ {
					row[x] = base
				}
			default:
				for x := ix0; x < ix1; x++ {
					row[x] = blendPremul(row[x], sr0, sg0, sb0, sa0)
				}
			}
			st.Pixels += n
		}
		st.Spans += iy1 - iy0
	}

	// 边带：只在开启 AA 时存在（关闭 AA 时整块都是满覆盖，已由上面的内接块覆盖）
	if ix0 == x0 && ix1 == x1 && iy0 == y0 && iy1 == y1 {
		return st
	}
	// 上下不满整行的两段：整宽逐像素
	for y := y0; y < y1; y++ {
		if y >= iy0 && y < iy1 {
			continue
		}
		coverRow(dst, r, y, x0, x1, base, opaque, shader, &st)
	}
	// 内部行的左右边带：通常各只有 1 列
	for y := iy0; y < iy1; y++ {
		coverRow(dst, r, y, x0, ix0, base, opaque, shader, &st)
		coverRow(dst, r, y, ix1, x1, base, opaque, shader, &st)
	}
	return st
}

// coverRow 把矩形 r 在第 y 行的 [xa,xb) 列按解析覆盖率混合进 dst。
//
// 只在边带上调用：内部像素已由 fillRect 批量铺过，这里再写会重复混合。
func coverRow(dst *raster.Backing, r geom.Rect, y, xa, xb int, base uint32, opaque bool, shader *render.Shader, st *raster.Stats) {
	if xa >= xb {
		return
	}
	cy := coverageAxis(float32(y), r.Min.Y, r.Max.Y)
	if cy <= 0 {
		return
	}
	row := dst.Row(y)
	sr0 := (base >> 16) & 0xFF
	sg0 := (base >> 8) & 0xFF
	sb0 := base & 0xFF
	sa0 := base >> 24
	for x := xa; x < xb; x++ {
		cx := coverageAxis(float32(x), r.Min.X, r.Max.X)
		if cx <= 0 {
			continue
		}
		a := uint32(cy*cx*255 + 0.5)
		if a == 0 {
			continue
		}
		if shader != nil {
			lp := geom.Pt(float32(x)+0.5, float32(y)+0.5)
			col := shader.Sample(shader.GradientT(lp))
			cr, cg, cb, ca := premulScaled(col, a)
			row[x] = blendPremul(row[x], cr, cg, cb, ca)
		} else if a == 255 {
			if opaque {
				row[x] = base
			} else {
				row[x] = blendPremul(row[x], sr0, sg0, sb0, sa0)
			}
		} else {
			row[x] = blendPremul(row[x], sr0*a/255, sg0*a/255, sb0*a/255, sa0*a/255)
		}
		st.Pixels++
	}
	st.Spans++
}

// clampInt 把 v 夹到 [lo, hi]。
func clampInt(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// coverageAxis 返回像素区间 [i, i+1) 与 [a, b) 的重叠比例。
func coverageAxis(i, a, b float32) float32 {
	lo := maxf(a, i)
	hi := minf(b, i+1)
	if hi <= lo {
		return 0
	}
	return hi - lo
}

// transform 把折线点变换到设备空间，复用 scratch 缓冲（热路径零分配）。
//
// 前提：调用方紧接着把结果交给 addPolyline/fill.addPolyline，且不保留引用。
func (b *Backend) transform(pts []geom.Point, m geom.Matrix) []geom.Point {
	out := b.scratch[:0]
	if cap(out) < len(pts) {
		out = make([]geom.Point, 0, len(pts))
	}
	for _, p := range pts {
		out = append(out, m.MapPoint(p))
	}
	b.scratch = out
	return out
}

func maxf(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func minf(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func floorf(v float32) float32 {
	i := float32(int64(v))
	if v < 0 && i != v {
		i--
	}
	return i
}

func ceilf(v float32) float32 { return -floorf(-v) }

func cosf(v float32) float32 { return float32(math.Cos(float64(v))) }
func sinf(v float32) float32 { return float32(math.Sin(float64(v))) }
