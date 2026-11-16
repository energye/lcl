package composite

import (
	"math"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/scene"
)

// ── 组滤镜 ──────────────────────────────────────────────────────────
//
// 阴影/发光/模糊都在**层局部空间**完成，最后作为一张整图参与合成。
// 这样做的好处：缓存键仍然是"尺寸 + 内容签名"，滤镜结果随层一起缓存，
// 静止的卡片投影不会每帧重算卷积。
//
// 三条渲染顺序固定：阴影 → 发光 → 内容。阴影和发光都在内容下方，
// 内容永远不被染脏。

// filterPad 返回本层为滤镜需要向外扩出的像素数。
func filterPad(l *Layer) int {
	pad := 0
	if l.boxShadow != nil {
		if m := shadowMargin(l.boxShadow); m > pad {
			pad = m
		}
	}
	if l.filter != nil {
		if l.filter.BlurSigma > 0 {
			if m := int(math.Ceil(float64(l.filter.BlurSigma) * 3)); m > pad {
				pad = m
			}
		}
		if l.filter.GlowRadius > 0 {
			if m := int(math.Ceil(float64(l.filter.GlowRadius) * 3)); m > pad {
				pad = m
			}
		}
		if l.filter.Shadow != nil {
			if m := shadowMargin(l.filter.Shadow); m > pad {
				pad = m
			}
		}
	}
	return pad
}

// fillRRectMask 把圆角矩形的覆盖率写进 mask（无 AA 的解析判定）。
//
// 只用于阴影形状：阴影随后会被模糊，边缘的 1 像素量化在模糊后不可见，
// 因此不需要 buildMask 那套 4×4 超采样。
func fillRRectShadowMask(m []uint8, w, h int, r geom.Rect, radius float32) {
	radius = minf(radius, minf(r.W(), r.H())*0.5)
	if radius <= 0 {
		fillRectMaskLocal(m, w, h, r)
		return
	}
	// 四个圆角的圆心。
	cx0, cy0 := r.Min.X+radius, r.Min.Y+radius
	cx1, cy1 := r.Max.X-radius, r.Max.Y-radius
	rr := radius * radius
	x0, y0 := maxi(int(r.Min.X), 0), maxi(int(r.Min.Y), 0)
	x1, y1 := mini(int(r.Max.X+0.5), w), mini(int(r.Max.Y+0.5), h)
	for y := y0; y < y1; y++ {
		py := float32(y) + 0.5
		for x := x0; x < x1; x++ {
			px := float32(x) + 0.5
			inside := true
			if px < cx0 && py < cy0 { // 左上
				dx, dy := px-cx0, py-cy0
				inside = dx*dx+dy*dy <= rr
			} else if px > cx1 && py < cy0 { // 右上
				dx, dy := px-cx1, py-cy0
				inside = dx*dx+dy*dy <= rr
			} else if px < cx0 && py > cy1 { // 左下
				dx, dy := px-cx0, py-cy1
				inside = dx*dx+dy*dy <= rr
			} else if px > cx1 && py > cy1 { // 右下
				dx, dy := px-cx1, py-cy1
				inside = dx*dx+dy*dy <= rr
			}
			if inside {
				m[y*w+x] = 255
			}
		}
	}
}

func fillRectMaskLocal(m []uint8, w, h int, r geom.Rect) {
	x0, y0 := int(r.Min.X+0.5), int(r.Min.Y+0.5)
	x1, y1 := int(r.Max.X+0.5), int(r.Max.Y+0.5)
	for y := maxi(y0, 0); y < mini(y1, h); y++ {
		for x := maxi(x0, 0); x < mini(x1, w); x++ {
			m[y*w+x] = 255
		}
	}
}

// offsetMask 把覆盖率按 (dx,dy) 平移，返回新数组（越界按 0）。
func offsetMask(src []uint8, w, h, dx, dy int) []uint8 {
	if dx == 0 && dy == 0 {
		return src
	}
	out := make([]uint8, w*h)
	for y := 0; y < h; y++ {
		sy := y - dy
		if sy < 0 || sy >= h {
			continue
		}
		for x := 0; x < w; x++ {
			sx := x - dx
			if sx < 0 || sx >= w {
				continue
			}
			out[y*w+x] = src[sy*w+sx]
		}
	}
	return out
}

// blendMaskColor 把一张覆盖率蒙版按指定颜色 source-over 混合进 dst。
//
// 蒙版与 dst 同尺寸同原点（滤镜缓冲已按 pad 扩过），因此不需要坐标换算。
func blendMaskColor(dst *raster.Backing, m []uint8, c geom.Color) int {
	if dst == nil || len(m) < dst.W()*dst.H() || c.A == 0 {
		return 0
	}
	w, h := dst.W(), dst.H()
	n := 0
	for y := 0; y < h; y++ {
		row := dst.Row(y)
		mrow := m[y*w : (y+1)*w]
		for x := 0; x < w; x++ {
			cov := uint32(mrow[x])
			if cov == 0 {
				continue
			}
			a := cov * uint32(c.A) / 255
			if a == 0 {
				continue
			}
			sr := uint32(c.R) * a / 255
			sg := uint32(c.G) * a / 255
			sb := uint32(c.B) * a / 255
			row[x] = blendPremulPx(row[x], sr, sg, sb, a)
			n++
		}
	}
	return n
}

// blendBackingPremul 把 src 以 source-over 叠进 dst（两者都是预乘，alpha=1）。
func blendBackingPremul(dst, src *raster.Backing, dx, dy int) {
	if dst == nil || src == nil {
		return
	}
	sw, sh := src.W(), src.H()
	dw, dh := dst.W(), dst.H()
	for sy := 0; sy < sh; sy++ {
		ty := dy + sy
		if ty < 0 || ty >= dh {
			continue
		}
		srow := src.Row(sy)
		drow := dst.Row(ty)
		for sx := 0; sx < sw; sx++ {
			tx := dx + sx
			if tx < 0 || tx >= dw {
				continue
			}
			p := srow[sx]
			if p == 0 {
				continue
			}
			drow[tx] = blendPremulPx(drow[tx], (p>>16)&0xFF, (p>>8)&0xFF, p&0xFF, p>>24)
		}
	}
}

// blendPremulPx 是预乘 BGRA 的 source-over。
func blendPremulPx(dst, sr, sg, sb, sa uint32) uint32 {
	if sa == 0 {
		return dst
	}
	if sa >= 255 {
		return 0xFF000000 | sr<<16 | sg<<8 | sb
	}
	ia := 255 - sa
	dr := (dst >> 16) & 0xFF
	dg := (dst >> 8) & 0xFF
	db := dst & 0xFF
	da := dst >> 24
	r := sr + dr*ia/255
	g := sg + dg*ia/255
	b := sb + db*ia/255
	a := sa + da*ia/255
	return a<<24 | r<<16 | g<<8 | b
}

// applyFilters 在 content 上应用阴影/发光/模糊，把最终画面写进 dst。
//
// content 与 dst 同尺寸（都已按 pad 扩过），content 的内容位于 (pad,pad)。
func applyFilters(l *Layer, pad int, content, dst *raster.Backing) {
	if l == nil || content == nil || dst == nil {
		return
	}
	w, h := content.W(), content.H()

	// 1) 盒阴影：形状来自节点矩形 + 圆角，不是内容 alpha。
	if l.boxShadow != nil {
		m := make([]uint8, w*h)
		nr := geom.RectXYWH(0, 0, l.nodeSize.W, l.nodeSize.H).Offset(float32(pad), float32(pad))
		fillRRectShadowMask(m, w, h, nr, l.boxRadius)
		m = offsetMask(m, w, h, roundf(l.boxShadow.Dx), roundf(l.boxShadow.Dy))
		if l.boxShadow.Blur > 0 {
			blurMask(m, w, h, l.boxShadow.Blur)
		}
		blendMaskColor(dst, m, l.boxShadow.Color)
	}

	// 2) 组阴影：形状来自内容 alpha。
	if l.filter != nil && l.filter.Shadow != nil && l.filter.Shadow.Color.A > 0 {
		m := alphaMask(content)
		m = offsetMask(m, w, h, roundf(l.filter.Shadow.Dx), roundf(l.filter.Shadow.Dy))
		if l.filter.Shadow.Blur > 0 {
			blurMask(m, w, h, l.filter.Shadow.Blur)
		}
		blendMaskColor(dst, m, l.filter.Shadow.Color)
	}

	// 3) 发光：内容 alpha 的扩散，压在内容下方。
	if l.filter != nil && l.filter.GlowRadius > 0 {
		m := alphaMask(content)
		blurMask(m, w, h, l.filter.GlowRadius*0.5)
		c := l.filter.GlowColor
		if c.A == 0 {
			c = geom.RGBAa(255, 255, 255, 255)
		}
		blendMaskColor(dst, m, c)
	}

	// 4) 组模糊：作用在内容本身（阴影已被单独处理，不会被二次模糊）。
	if l.filter != nil && l.filter.BlurSigma > 0 {
		blurBacking(content, l.filter.BlurSigma)
	}

	// 5) 内容叠上去。content 与 dst 同尺寸同原点（内容都已落在 pad 偏移处），
	// 所以是 1:1 叠加，不能再按 pad 平移一次。
	blendBackingPremul(dst, content, 0, 0)
}

// filterEffect 返回节点上的组滤镜（可能为 nil）。
func filterEffect(n *scene.Node) *scene.EffectSpec {
	if n == nil {
		return nil
	}
	return n.Effect()
}
