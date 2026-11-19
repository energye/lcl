// Package chart 提供数据驱动的图表组件。
//
// 全部基于 scene.CustomPaintNode + render.Canvas 录制通道（C17 允许），
// 不碰 lcl、不持像素、不启定时器：采样/数值映射在 Build 期算好，
// Paint 闭包只做几何绘制。这让同一份图表能在窗口、headless 与诊断中
// 出同样结果，也便于 --render-matrix 做位置比对。
package chart

import (
	"math"
	"strconv"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── 公共基础 ───────────────────────────────────────────────────────

// Palette 是图表配色。零值由 theme 推导，非零点覆盖。
type Palette struct {
	Series []geom.Color // 多序列/分类色；不够时循环
	Grid   geom.Color
	Axis   geom.Color
	Text   geom.Color
	Fill   geom.Color
	Line   geom.Color
}

func paletteOr(th *theme.Data, p Palette) Palette {
	if len(p.Series) == 0 {
		c := th.Color
		p.Series = []geom.Color{
			c.Primary, c.Success, c.Warning, c.Info, c.Danger,
			geom.RGBA(148, 103, 189), geom.RGBA(227, 119, 194),
		}
	}
	if p.Grid == (geom.Color{}) {
		p.Grid = th.Color.Border
	}
	if p.Axis == (geom.Color{}) {
		p.Axis = th.Color.BorderStrong
	}
	if p.Text == (geom.Color{}) {
		p.Text = th.Color.TextMuted
	}
	if p.Fill == (geom.Color{}) {
		p.Fill = th.Color.Surface
	}
	if p.Line == (geom.Color{}) {
		p.Line = th.Color.Primary
	}
	return p
}

func colorAt(p Palette, i int) geom.Color {
	if len(p.Series) == 0 {
		return p.Line
	}
	return p.Series[i%len(p.Series)]
}

// node 构造固定尺寸的自绘节点。
func node(w, h float32, paint scene.CustomPaint) *scene.Node {
	n := scene.CustomPaintNode(paint)
	n.SetSize2(w, h)
	// 被动定位锚：图表色占比断言用（HitDeferToChild 不影响交互）。
	n.SetHit(scene.HitSpec{Label: "chart-canvas"})
	return n
}

// defaultSize 把非正尺寸补成合理默认值。
func defaultSize(w, h float32) (float32, float32) {
	if w <= 0 {
		w = 320
	}
	if h <= 0 {
		h = 180
	}
	return w, h
}

// niceRange 把 [min,max] 扩到"整齐"范围（含 0 轴，端点留白）。
func niceRange(minV, maxV float64) (float64, float64) {
	if math.IsNaN(minV) || math.IsNaN(maxV) || math.IsInf(minV, 0) || math.IsInf(maxV, 0) {
		return 0, 1
	}
	if minV == maxV {
		if minV == 0 {
			return 0, 1
		}
		if minV > 0 {
			return 0, minV * 1.2
		}
		return minV * 1.2, 0
	}
	pad := (maxV - minV) * 0.08
	lo, hi := minV-pad, maxV+pad
	if minV >= 0 && lo < 0 {
		lo = 0
	}
	if maxV <= 0 && hi > 0 {
		hi = 0
	}
	return lo, hi
}

func valueAt(lo, hi float64, t float32) float64 {
	return lo + float64(t)*(hi-lo)
}

func yOf(v, lo, hi float64, top, bottom float32) float32 {
	span := hi - lo
	if span == 0 {
		return bottom
	}
	t := float32((v - lo) / span)
	return bottom - t*(bottom-top)
}

func fmtInt(v int64) string     { return strconv.FormatInt(v, 10) }
func fmtFloat(v float64) string { return strconv.FormatFloat(v, 'f', 1, 64) }

// ── LineChart：折线 + 面积 + 网格 ────────────────────────────────

// LineChart 是折线图。Series 为多条序列，每条是等长采样值。
type LineChart struct {
	Series    [][]float64
	Width     float32
	Height    float32
	Labels    []string // X 轴刻度（可选）
	FillArea  bool
	ShowDots  bool
	Palette   Palette
	AxisLabel string
}

// Build 实现 ui.Widget。
func (c LineChart) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	p := paletteOr(th, c.Palette)
	w, h := defaultSize(c.Width, c.Height)
	series := c.Series
	labels := append([]string(nil), c.Labels...)
	fillArea, showDots := c.FillArea, c.ShowDots
	axisLabel := c.AxisLabel

	return node(w, h, func(cv render.Canvas, sz geom.Size) {
		const axisW, axisH float32 = 36, 22
		left, right, top, bottom := axisW, sz.W-8, float32(8), sz.H-axisH
		if right <= left || bottom <= top {
			return
		}
		cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(p.Fill))

		lo, hi := chartRange(series)
		// 网格
		const rows = 4
		for i := 0; i <= rows; i++ {
			y := top + float32(i)*(bottom-top)/rows
			cv.DrawLine(left, y, right, y, render.Stroke(p.Grid, 1))
			v := valueAt(lo, hi, 1-float32(i)/rows)
			if l := text.LayoutText(formatTick(v), text.DefaultFont(th.Font.SizeXS), text.Options{}); l != nil {
				cv.DrawText(l, geom.Pt(2, y-l.Size.H/2), p.Text)
			}
		}
		// X 轴
		cv.DrawLine(left, bottom, right, bottom, render.Stroke(p.Axis, 1))
		for i, lab := range labels {
			if len(series) == 0 || len(series[0]) == 0 {
				break
			}
			x := xOf(i, len(series[0]), left, right)
			if l := text.LayoutText(lab, text.DefaultFont(th.Font.SizeXS), text.Options{}); l != nil {
				cv.DrawText(l, geom.Pt(x-l.Size.W/2, bottom+3), p.Text)
			}
		}

		for si, s := range series {
			if len(s) == 0 {
				continue
			}
			col := colorAt(p, si)
			path := render.NewPath()
			for i, v := range s {
				x, y := xOf(i, len(s), left, right), yOf(v, lo, hi, top, bottom)
				if i == 0 {
					path.MoveTo(x, y)
				} else {
					path.LineTo(x, y)
				}
			}
			if fillArea && len(s) > 1 {
				area := render.NewPath()
				area.MoveTo(xOf(0, len(s), left, right), bottom)
				for i, v := range s {
					area.LineTo(xOf(i, len(s), left, right), yOf(v, lo, hi, top, bottom))
				}
				area.LineTo(xOf(len(s)-1, len(s), left, right), bottom)
				area.Close()
				cv.FillPath(area, render.Fill(col.WithAlpha(48)))
			}
			cv.StrokePath(path, render.Stroke(col, 2).WithCap(render.CapRound).WithJoin(render.JoinRound))
			if showDots {
				for i, v := range s {
					x, y := xOf(i, len(s), left, right), yOf(v, lo, hi, top, bottom)
					cv.FillOval(geom.RectXYWH(x-2.5, y-2.5, 5, 5), render.Fill(col))
				}
			}
		}
		if axisLabel != "" {
			if l := text.LayoutText(axisLabel, text.DefaultFont(th.Font.SizeXS), text.Options{}); l != nil {
				cv.DrawText(l, geom.Pt(left, 1), p.Text)
			}
		}
	})
}

func chartRange(series [][]float64) (float64, float64) {
	minV, maxV := math.Inf(1), math.Inf(-1)
	for _, s := range series {
		for _, v := range s {
			if v < minV {
				minV = v
			}
			if v > maxV {
				maxV = v
			}
		}
	}
	if math.IsInf(minV, 1) {
		return 0, 1
	}
	return niceRange(minV, maxV)
}

func xOf(i, n int, left, right float32) float32 {
	if n <= 1 {
		return (left + right) / 2
	}
	return left + float32(i)*((right-left)/float32(n-1))
}

func formatTick(v float64) string {
	if v == math.Trunc(v) {
		return fmtInt(int64(v))
	}
	return fmtFloat(v)
}

// ── BarChart ───────────────────────────────────────────────────────

// BarChart 是分组柱状图。Values 为各组，每组含多序列。
type BarChart struct {
	Groups  [][]float64
	Labels  []string
	Width   float32
	Height  float32
	Palette Palette
}

// Build 实现 ui.Widget。
func (c BarChart) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	p := paletteOr(th, c.Palette)
	w, h := defaultSize(c.Width, c.Height)
	groups := c.Groups
	labels := append([]string(nil), c.Labels...)

	return node(w, h, func(cv render.Canvas, sz geom.Size) {
		const axisW, axisH float32 = 36, 22
		left, right, top, bottom := axisW, sz.W-8, float32(8), sz.H-axisH
		if right <= left || bottom <= top || len(groups) == 0 {
			return
		}
		cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(p.Fill))
		lo, hi := chartRange(groups)
		const rows = 4
		for i := 0; i <= rows; i++ {
			y := top + float32(i)*(bottom-top)/rows
			cv.DrawLine(left, y, right, y, render.Stroke(p.Grid, 1))
			if l := text.LayoutText(formatTick(valueAt(lo, hi, 1-float32(i)/rows)),
				text.DefaultFont(th.Font.SizeXS), text.Options{}); l != nil {
				cv.DrawText(l, geom.Pt(2, y-l.Size.H/2), p.Text)
			}
		}
		cv.DrawLine(left, bottom, right, bottom, render.Stroke(p.Axis, 1))

		slot := (right - left) / float32(len(groups))
		maxSeries := 0
		for _, g := range groups {
			if len(g) > maxSeries {
				maxSeries = len(g)
			}
		}
		barW := slot * 0.7 / float32(maxSeries)
		for gi, g := range groups {
			cx := left + slot*(float32(gi)+0.5)
			for si, v := range g {
				x := cx - slot*0.35 + float32(si)*barW
				y := yOf(v, lo, hi, top, bottom)
				cv.FillRect(geom.RectXYWH(x, y, barW-1, bottom-y), render.Fill(colorAt(p, si)))
			}
			if gi < len(labels) {
				if l := text.LayoutText(labels[gi], text.DefaultFont(th.Font.SizeXS), text.Options{}); l != nil {
					cv.DrawText(l, geom.Pt(cx-l.Size.W/2, bottom+3), p.Text)
				}
			}
		}
	})
}

// ── PieChart / DonutChart ──────────────────────────────────────────

// PieChart 是饼图/环形图。Values 为各项数值，Donut 决定中心镂空。
type PieChart struct {
	Values  []float64
	Labels  []string
	Width   float32
	Height  float32
	Donut   bool
	Palette Palette
}

// Build 实现 ui.Widget。
func (c PieChart) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	p := paletteOr(th, c.Palette)
	w, h := defaultSize(c.Width, c.Height)
	values := append([]float64(nil), c.Values...)
	labels := append([]string(nil), c.Labels...)
	donut := c.Donut

	return node(w, h, func(cv render.Canvas, sz geom.Size) {
		cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(p.Fill))
		if len(values) == 0 {
			return
		}
		sum := 0.0
		for _, v := range values {
			if v > 0 {
				sum += v
			}
		}
		if sum <= 0 {
			return
		}
		// 图例在右侧占 ~96px：饼/环在剩余区域居中，避免图例叠在图形上。
		legendW := float32(0)
		if len(labels) > 0 {
			legendW = 96
		}
		availW := sz.W - 16 - legendW
		side := minf(availW, sz.H-16)
		if side <= 0 {
			return
		}
		rr := geom.RectXYWH(8+(availW-side)/2, (sz.H-side)/2, side, side)
		angle := -float32(math.Pi) / 2
		for i, v := range values {
			if v <= 0 {
				continue
			}
			sweep := float32(v/sum) * float32(math.Pi*2)
			col := colorAt(p, i)
			path := render.NewPath()
			cx, cy := (rr.Min.X+rr.Max.X)/2, (rr.Min.Y+rr.Max.Y)/2
			if donut {
				path.AddArc(rr, angle, sweep)
				inner := geom.RectXYWH(cx-side*0.28/2, cy-side*0.28/2, side*0.28, side*0.28)
				path.AddArc(inner, angle+sweep, -sweep)
				path.Close()
			} else {
				path.MoveTo(cx, cy)
				path.AddArc(rr, angle, sweep)
				path.LineTo(cx, cy)
				path.Close()
			}
			cv.FillPath(path, render.Fill(col))
			angle += sweep
		}
		// 图例
		lx := sz.W - 96
		if lx < sz.W/2 {
			lx = sz.W / 2
		}
		for i, lab := range labels {
			if i >= len(values) {
				break
			}
			y := 12 + float32(i)*16
			if y+14 > sz.H {
				break
			}
			cv.FillRect(geom.RectXYWH(lx, y+3, 8, 8), render.Fill(colorAt(p, i)))
			if l := text.LayoutText(lab, text.DefaultFont(th.Font.SizeXS), text.Options{}); l != nil {
				cv.DrawText(l, geom.Pt(lx+12, y), p.Text)
			}
		}
	})
}

// ── Gauge：环形进度仪表 ───────────────────────────────────────────

// Gauge 是环形进度仪表。Value 归一化到 [0,1]。
type Gauge struct {
	Value   float64
	Label   string
	Width   float32
	Height  float32
	Color   geom.Color
	Palette Palette
}

// Build 实现 ui.Widget。
func (g Gauge) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	p := paletteOr(th, g.Palette)
	w, h := defaultSize(g.Width, g.Height)
	v, lab := g.Value, g.Label
	col := g.Color
	if col == (geom.Color{}) {
		col = p.Line
	}

	return node(w, h, func(cv render.Canvas, sz geom.Size) {
		cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(p.Fill))
		side := minf(sz.W, sz.H) - 12
		if side <= 0 {
			return
		}
		rr := geom.RectXYWH((sz.W-side)/2, (sz.H-side)/2, side, side)
		if v < 0 {
			v = 0
		}
		if v > 1 {
			v = 1
		}
		start := -float32(math.Pi) / 2
		cv.DrawArc(rr, start, float32(math.Pi*2), render.Stroke(p.Grid, 8).WithCap(render.CapRound))
		if v > 0 {
			cv.DrawArc(rr, start, float32(v*math.Pi*2), render.Stroke(col, 8).WithCap(render.CapRound))
		}
		if lab == "" {
			lab = formatPercent(v)
		}
		if l := text.LayoutText(lab, text.DefaultFont(th.Font.SizeBase), text.Options{}); l != nil {
			cv.DrawText(l, geom.Pt((sz.W-l.Size.W)/2, (sz.H-l.Size.H)/2), th.Color.Text)
		}
	})
}

func formatPercent(v float64) string {
	return fmtInt(int64(math.Round(v*100))) + "%"
}

// ── Waveform：密集波形 ────────────────────────────────────────────

// Waveform 是密集采样波形（X 轴均匀，自动缩放）。适合示波器/音频展示。
type Waveform struct {
	Values  []float64
	Width   float32
	Height  float32
	Color   geom.Color
	Palette Palette
}

// Build 实现 ui.Widget。
func (w Waveform) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	p := paletteOr(th, w.Palette)
	width, height := defaultSize(w.Width, w.Height)
	values := append([]float64(nil), w.Values...)
	col := w.Color
	if col == (geom.Color{}) {
		col = p.Line
	}

	return node(width, height, func(cv render.Canvas, sz geom.Size) {
		cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(p.Fill))
		if len(values) < 2 {
			return
		}
		lo, hi := chartRange([][]float64{values})
		mid := sz.H / 2
		path := render.NewPath()
		for i, v := range values {
			x := float32(i) * (sz.W / float32(len(values)-1))
			y := mid - (float32((v-(lo+hi)/2)/(hi-lo)) * mid * 0.9)
			if i == 0 {
				path.MoveTo(x, y)
			} else {
				path.LineTo(x, y)
			}
		}
		cv.DrawLine(0, mid, sz.W, mid, render.Stroke(p.Grid, 1))
		cv.StrokePath(path, render.Stroke(col, 1.5).WithCap(render.CapRound))
	})
}

// ── Scatter：散点图 ───────────────────────────────────────────────

// ScatterPoint 是一个散点。
type ScatterPoint struct {
	X, Y float64
	Size float32
}

// Scatter 是散点图。
type Scatter struct {
	Points  []ScatterPoint
	Width   float32
	Height  float32
	Color   geom.Color
	Palette Palette
}

// Build 实现 ui.Widget。
func (s Scatter) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	p := paletteOr(th, s.Palette)
	w, h := defaultSize(s.Width, s.Height)
	points := append([]ScatterPoint(nil), s.Points...)
	col := s.Color
	if col == (geom.Color{}) {
		col = p.Line
	}

	return node(w, h, func(cv render.Canvas, sz geom.Size) {
		const axisW, axisH float32 = 36, 22
		left, right, top, bottom := axisW, sz.W-8, float32(8), sz.H-axisH
		if right <= left || bottom <= top {
			return
		}
		cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(p.Fill))
		xs, ys := make([]float64, 0, len(points)), make([]float64, 0, len(points))
		for _, pt := range points {
			xs = append(xs, pt.X)
			ys = append(ys, pt.Y)
		}
		xlo, xhi := niceRange(minOf(xs), maxOf(xs))
		ylo, yhi := niceRange(minOf(ys), maxOf(ys))
		const rows = 4
		for i := 0; i <= rows; i++ {
			y := top + float32(i)*(bottom-top)/rows
			cv.DrawLine(left, y, right, y, render.Stroke(p.Grid, 1))
			cv.DrawLine(left+float32(i)*(right-left)/rows, top,
				left+float32(i)*(right-left)/rows, bottom, render.Stroke(p.Grid, 1))
		}
		cv.DrawLine(left, bottom, right, bottom, render.Stroke(p.Axis, 1))
		cv.DrawLine(left, top, left, bottom, render.Stroke(p.Axis, 1))
		for _, pt := range points {
			spanX, spanY := xhi-xlo, yhi-ylo
			if spanX == 0 || spanY == 0 {
				continue
			}
			x := left + float32((pt.X-xlo)/spanX)*(right-left)
			y := bottom - float32((pt.Y-ylo)/spanY)*(bottom-top)
			r := pt.Size
			if r <= 0 {
				r = 3
			}
			cv.FillOval(geom.RectXYWH(x-r, y-r, r*2, r*2), render.Fill(col.WithAlpha(200)))
		}
	})
}

func minOf(v []float64) float64 {
	if len(v) == 0 {
		return 0
	}
	m := v[0]
	for _, x := range v[1:] {
		if x < m {
			m = x
		}
	}
	return m
}

func maxOf(v []float64) float64 {
	if len(v) == 0 {
		return 0
	}
	m := v[0]
	for _, x := range v[1:] {
		if x > m {
			m = x
		}
	}
	return m
}

// ── 运算符辅助（不引 math，保持包极轻） ───────────────────────────

func minf(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

// ── Widget 接口断言 ───────────────────────────────────────────────

var (
	_ ui.Widget = LineChart{}
	_ ui.Widget = BarChart{}
	_ ui.Widget = PieChart{}
	_ ui.Widget = Gauge{}
	_ ui.Widget = Waveform{}
	_ ui.Widget = Scatter{}
)
