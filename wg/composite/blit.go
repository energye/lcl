package composite

import (
	"math"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
)

// blendBacking 把 src 以 source-over 混合到 dst 的 (dx,dy)。
//
// 参数含义（三者相乘就是"这个源像素最终的有效 alpha"）：
//   - src 自带 alpha（内容本身的半透明）；
//   - alpha 是层透明度（组透明，0..1）；
//   - mask 是裁剪覆盖率（nil 表示不裁剪）；
//   - clip 是设备空间的矩形裁剪（滚动视口等继承裁剪在这里收口——
//     它不烘进层内容，所以缓存命中的贴图同样被正确裁剪）。
//
// 全部在**预乘**空间里算：src 已经是预乘的，所以"再乘一个因子 f"
// 等价于 RGB 与 A 同时乘 f，不需要反预乘再预乘（那会丢精度）。
//
// 返回实际参与混合的像素数（统计用）。
func blendBacking(dst, src *raster.Backing, dx, dy int, alpha float32, mask []uint8, clip geom.Rect) int {
	if dst == nil || src == nil {
		return 0
	}
	f0 := int(clamp01(alpha)*255 + 0.5)
	if f0 == 0 {
		return 0
	}
	if mask != nil && len(mask) < src.W()*src.H() {
		mask = nil // 蒙版尺寸不匹配：宁可不裁也不能越界读
	}
	sw, sh := src.W(), src.H()
	// 有效混合区 = clip（取整向外）∩ dst 边界，换算成 src 的行列区间。
	// 循环前收敛成区间而不是逐像素判断：热循环里一次比较都不加。
	cx0, cy0 := int(math.Floor(float64(clip.Min.X))), int(math.Floor(float64(clip.Min.Y)))
	cx1, cy1 := int(math.Ceil(float64(clip.Max.X))), int(math.Ceil(float64(clip.Max.Y)))
	dw, dh := dst.W(), dst.H()
	if cx0 < 0 {
		cx0 = 0
	}
	if cy0 < 0 {
		cy0 = 0
	}
	if cx1 > dw {
		cx1 = dw
	}
	if cy1 > dh {
		cy1 = dh
	}
	sx0, sx1 := clampi(cx0-dx, 0, sw), clampi(cx1-dx, 0, sw)
	sy0, sy1 := clampi(cy0-dy, 0, sh), clampi(cy1-dy, 0, sh)
	if sx0 >= sx1 || sy0 >= sy1 {
		return 0
	}
	n := 0
	for sy := sy0; sy < sy1; sy++ {
		srow := src.Row(sy)
		drow := dst.Row(dy + sy)
		var mrow []uint8
		if mask != nil {
			mrow = mask[sy*sw : (sy+1)*sw]
		}
		for sx := sx0; sx < sx1; sx++ {
			sp := srow[sx]
			sa0 := int(sp >> 24)
			if sa0 == 0 {
				continue
			}
			f := f0
			if mrow != nil {
				m := int(mrow[sx])
				if m == 0 {
					continue
				}
				if m != 255 {
					f = f * m / 255
				}
			}
			if f == 0 {
				continue
			}
			sa := sa0
			if f != 255 {
				sa = sa0 * f / 255
				if sa == 0 {
					continue
				}
			}
			tx := dx + sx
			d := drow[tx]
			if f == 255 && sa == 255 {
				// 完全不透明：直接替换，省掉两次乘法
				drow[tx] = sp
				n++
				continue
			}
			sr := int((sp>>16)&0xFF) * f / 255
			sg := int((sp>>8)&0xFF) * f / 255
			sb := int(sp&0xFF) * f / 255
			inv := 255 - sa
			orr := sr + int((d>>16)&0xFF)*inv/255
			og := sg + int((d>>8)&0xFF)*inv/255
			ob := sb + int(d&0xFF)*inv/255
			oa := sa + int(d>>24)*inv/255
			if orr > 255 {
				orr = 255
			}
			if og > 255 {
				og = 255
			}
			if ob > 255 {
				ob = 255
			}
			if oa > 255 {
				oa = 255
			}
			drow[tx] = uint32(oa)<<24 | uint32(orr)<<16 | uint32(og)<<8 | uint32(ob)
			n++
		}
	}
	return n
}

// clampi 把 v 钳到 [lo,hi]。
func clampi(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
