package software

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
)

// ── 位图绘制 ────────────────────────────────────────────────────────
//
// ImageRef 与 Backing 同为**预乘 BGRA**，所以采样后可以直接走 blendPremul，
// 不需要来回反预乘——这也是当初统一像素格式的收益。
//
// 三档采样：
//   - Nearest：取最近像素（图标放大、像素画）。
//   - Bilinear：源空间双线性（默认，缩小时仍会欠采样）。
//   - Mipmap：没有金字塔；对缩小超过 2 倍的情况用 2×2 盒式预滤近似，
//     把采样足迹摊到 2×2 源像素上，抑制缩小闪烁。这是近似实现，
//     不是真正的 mip 链（文档与 Caps 里如实标注）。
func (b *Backend) drawImage(dst *raster.Backing, img *render.ImageRef, dstRect geom.Rect,
	q render.ImageQuality, clip geom.Rect) raster.Stats {
	var st raster.Stats
	if dst == nil || img == nil || img.W <= 0 || img.H <= 0 || len(img.Pix) == 0 {
		st.Degraded++
		return st
	}
	r := dstRect.Intersect(clip).Intersect(dst.Bounds())
	r = r.RoundOut()
	if r.IsEmpty() {
		return st
	}
	x0, y0 := int(r.Min.X), int(r.Min.Y)
	x1, y1 := int(r.Max.X), int(r.Max.Y)
	if x1 > dst.W() {
		x1 = dst.W()
	}
	if y1 > dst.H() {
		y1 = dst.H()
	}
	if x0 < 0 {
		x0 = 0
	}
	if y0 < 0 {
		y0 = 0
	}
	if x0 >= x1 || y0 >= y1 {
		return st
	}

	// 目标像素中心 → 源像素坐标（连续空间，0 = 第一个像素左边界）。
	invW := float32(img.W) / dstRect.W()
	invH := float32(img.H) / dstRect.H()
	prefilter := q == render.ImageMipmap && (invW > 2 || invH > 2)

	for y := y0; y < y1; y++ {
		v := (float32(y)+0.5-dstRect.Min.Y)*invH - 0.5
		row := dst.Row(y)
		for x := x0; x < x1; x++ {
			u := (float32(x)+0.5-dstRect.Min.X)*invW - 0.5
			var cr, cg, cb, ca uint32
			switch {
			case q == render.ImageNearest:
				cr, cg, cb, ca = samplePremul(img, int(u+0.5), int(v+0.5))
			case prefilter:
				cr, cg, cb, ca = sampleBox2(img, u, v)
			default:
				cr, cg, cb, ca = sampleBilinear(img, u, v)
			}
			if ca == 0 {
				continue
			}
			row[x] = blendPremul(row[x], cr, cg, cb, ca)
			st.Pixels++
		}
		st.Spans++
	}
	return st
}

// clampCoord 把源坐标夹到 [0, n-1]（边缘延展）。
func clampCoord(v, n int) int {
	if v < 0 {
		return 0
	}
	if v >= n {
		return n - 1
	}
	return v
}

// samplePremul 取单个源像素（预乘 BGRA）。
func samplePremul(img *render.ImageRef, x, y int) (r, g, b, a uint32) {
	x = clampCoord(x, img.W)
	y = clampCoord(y, img.H)
	p := img.Pix[y*img.Stride+x]
	return (p >> 16) & 0xFF, (p >> 8) & 0xFF, p & 0xFF, p >> 24
}

// sampleBilinear 在源空间做双线性插值（预乘空间插值即正确的 source-over 近似）。
func sampleBilinear(img *render.ImageRef, u, v float32) (r, g, b, a uint32) {
	fx := floorf(u)
	fy := floorf(v)
	tx := u - fx
	ty := v - fy
	ix, iy := int(fx), int(fy)

	r00, g00, b00, a00 := samplePremul(img, ix, iy)
	r10, g10, b10, a10 := samplePremul(img, ix+1, iy)
	r01, g01, b01, a01 := samplePremul(img, ix, iy+1)
	r11, g11, b11, a11 := samplePremul(img, ix+1, iy+1)

	mix := func(p00, p10, p01, p11 uint32) uint32 {
		top := float32(p00)*(1-tx) + float32(p10)*tx
		bot := float32(p01)*(1-tx) + float32(p11)*tx
		return uint32(top*(1-ty)+bot*ty + 0.5)
	}
	return mix(r00, r10, r01, r11), mix(g00, g10, g01, g11),
		mix(b00, b10, b01, b11), mix(a00, a10, a01, a11)
}

// sampleBox2 对 2×2 源像素做盒式平均（mipmap 的廉价近似）。
func sampleBox2(img *render.ImageRef, u, v float32) (r, g, b, a uint32) {
	ix, iy := int(floorf(u)), int(floorf(v))
	r00, g00, b00, a00 := samplePremul(img, ix, iy)
	r10, g10, b10, a10 := samplePremul(img, ix+1, iy)
	r01, g01, b01, a01 := samplePremul(img, ix, iy+1)
	r11, g11, b11, a11 := samplePremul(img, ix+1, iy+1)
	return (r00 + r10 + r01 + r11 + 2) / 4,
		(g00 + g10 + g01 + g11 + 2) / 4,
		(b00 + b10 + b01 + b11 + 2) / 4,
		(a00 + a10 + a01 + a11 + 2) / 4
}
