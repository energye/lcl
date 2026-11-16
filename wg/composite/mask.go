package composite

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
)

// 蒙版超采样：每个像素取 4×4 个点。
//
// 为什么是覆盖率蒙版而不是"给每个像素算精确面积"：圆角矩形与任意路径的
// 精确面积需要解析求交，代码量与出错面都远大于收益。4×4 = 17 级覆盖率
// 在圆角边缘上肉眼已分不出台阶，而代价只是每像素 16 次点在内测试——
// 而且只在**缓存未命中**时算一次（蒙版随离屏一起缓存）。
const maskSamples = 4

// buildMask 计算裁剪蒙版：返回 w×h 的覆盖率数组（0=全透明，255=全不透明）。
//
// 坐标系：离屏缓冲的 (0,0) 对应 layer.bounds.Min，所以层的裁剪矩形要
// 先减去 bounds.Min 才对得上像素下标。这一步错了会表现为"圆角跑到外面"，
// 而且只在层有偏移时才现形——测试必须覆盖带偏移的层。
func buildMask(l *Layer, w, h int) []uint8 {
	if l == nil || w <= 0 || h <= 0 || l.clip == ClipNone || l.clip == ClipRect {
		// 矩形裁剪由指令流的 scissor 精确完成，不需要蒙版。
		return nil
	}
	off := l.bounds.RoundOut().Min
	r := l.clipRect.Offset(-off.X, -off.Y)

	m := make([]uint8, w*h)
	if l.clip == ClipRRect {
		fillRRectMask(m, w, h, r, l.clipRadius)
		return m
	}
	if l.clipPath != nil {
		polys := l.clipPath.Flatten(0.25, nil)
		for i := range polys {
			for j := range polys[i].Pts {
				polys[i].Pts[j] = polys[i].Pts[j].Sub(off)
			}
		}
		fillPathMask(m, w, h, polys)
		return m
	}
	// 既非圆角也非路径：退化为矩形（全 255 在裁剪区内）。
	fillRectMask(m, w, h, r)
	return m
}

// fillRectMask 填充矩形蒙版（逐行区间，无 AA——矩形边界本就是像素边界）。
func fillRectMask(m []uint8, w, h int, r geom.Rect) {
	x0, y0, x1, y1 := int(r.Min.X+0.5), int(r.Min.Y+0.5), int(r.Max.X+0.5), int(r.Max.Y+0.5)
	for y := maxi(y0, 0); y < mini(y1, h); y++ {
		for x := maxi(x0, 0); x < mini(x1, w); x++ {
			m[y*w+x] = 255
		}
	}
}

// fillRRectMask 填充圆角矩形蒙版。
//
// 内部矩形（内缩 radius）整块 255，只有四条边带需要超采样——
// 这是把"每像素 16 次"降到"边缘附近每像素 16 次"的关键。
func fillRRectMask(m []uint8, w, h int, r geom.Rect, radius float32) {
	radius = minf(radius, minf(r.W(), r.H())*0.5)
	if radius <= 0 {
		fillRectMask(m, w, h, r)
		return
	}
	inner := r.Inset(radius)
	sx := float32(1.0) / maskSamples
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			px := geom.RectXYWH(float32(x), float32(y), 1, 1)
			if inner.ContainsRect(px) {
				m[y*w+x] = 255
				continue
			}
			if !r.Expand(1).Intersects(px) {
				continue
			}
			hits := 0
			for j := 0; j < maskSamples; j++ {
				sy := float32(y) + (float32(j)+0.5)*sx
				for i := 0; i < maskSamples; i++ {
					sxPt := float32(x) + (float32(i)+0.5)*sx
					if scene.PointInRRect(geom.Pt(sxPt-r.Min.X, sy-r.Min.Y), r.Size(), radius) {
						hits++
					}
				}
			}
			if hits > 0 {
				m[y*w+x] = uint8(hits * 255 / (maskSamples * maskSamples))
			}
		}
	}
}

// fillPathMask 填充任意路径蒙版（4×4 超采样 + 非零环绕规则）。
func fillPathMask(m []uint8, w, h int, polys []render.Polyline) {
	if len(polys) == 0 {
		return
	}
	var bounds geom.Rect
	first := true
	for _, p := range polys {
		for _, pt := range p.Pts {
			if first {
				bounds = geom.RectLTRB(pt.X, pt.Y, pt.X, pt.Y)
				first = false
				continue
			}
			bounds = geom.RectLTRB(
				minf(bounds.Min.X, pt.X), minf(bounds.Min.Y, pt.Y),
				maxf(bounds.Max.X, pt.X), maxf(bounds.Max.Y, pt.Y),
			)
		}
	}
	if first {
		return
	}
	sx := float32(1.0) / maskSamples
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			px := geom.RectXYWH(float32(x), float32(y), 1, 1)
			if !bounds.Expand(1).Intersects(px) {
				continue
			}
			hits := 0
			for j := 0; j < maskSamples; j++ {
				sy := float32(y) + (float32(j)+0.5)*sx
				for i := 0; i < maskSamples; i++ {
					sxPt := float32(x) + (float32(i)+0.5)*sx
					if windingContains(polys, geom.Pt(sxPt, sy)) {
						hits++
					}
				}
			}
			if hits > 0 {
				m[y*w+x] = uint8(hits * 255 / (maskSamples * maskSamples))
			}
		}
	}
}

// windingContains 用非零环绕规则判定点是否在多边形集合内。
func windingContains(polys []render.Polyline, p geom.Point) bool {
	winding := 0
	for _, poly := range polys {
		pts := poly.Pts
		n := len(pts)
		if n < 3 {
			continue
		}
		closed := poly.Closed
		for i := 0; i < n; i++ {
			a := pts[i]
			var b geom.Point
			if i+1 < n {
				b = pts[i+1]
			} else if closed {
				b = pts[0]
			} else {
				break
			}
			if a.Y <= p.Y {
				if b.Y > p.Y && cross(a, b, p) > 0 {
					winding++
				}
			} else if b.Y <= p.Y && cross(a, b, p) < 0 {
				winding--
			}
		}
	}
	return winding != 0
}

// cross 返回 (b-a) × (p-a) 的 z 分量。
func cross(a, b, p geom.Point) float32 {
	return (b.X-a.X)*(p.Y-a.Y) - (p.X-a.X)*(b.Y-a.Y)
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxi(a, b int) int {
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

func maxf(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
