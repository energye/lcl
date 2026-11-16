package composite

import (
	"math"

	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/scene"
)

// ── 模糊 ────────────────────────────────────────────────────────────
//
// 为什么用三次 box blur 近似高斯：一维高斯卷积是 O(r) 每像素，半径大时
// 直接吃满帧预算；三次 box 是 O(1) 每像素（滑窗累加），视觉上与同 sigma
// 的高斯差不到 1/255（Kutskir 的三次盒近似）。阴影/发光都只需要
// "看起来像高斯"，不需要精确核。
//
// 只处理预乘 BGRA：预乘空间里 blur 四个通道是线性可加的，
// 反预乘再做模糊会在 alpha 边缘产生色晕。

// boxesForGauss 返回三次 box blur 的半径（Kutskir 公式）。
func boxesForGauss(sigma float32, n int) [3]int {
	var out [3]int
	if sigma <= 0 || n <= 0 {
		return out
	}
	wIdeal := math.Sqrt((12*float64(sigma)*float64(sigma)/float64(n)) + 1)
	wl := int(math.Floor(wIdeal))
	if wl%2 == 0 {
		wl--
	}
	lw := wl
	if lw < 1 {
		lw = 1
	}
	wu := wl + 2
	mIdeal := (12*float64(sigma)*float64(sigma) - float64(n*lw*lw) - float64(4*n*lw) - float64(3*n)) / float64(-4*lw-4)
	m := int(math.Round(mIdeal))
	for i := 0; i < n; i++ {
		if i < m {
			out[i] = lw
		} else {
			out[i] = wu
		}
	}
	return out
}

// blurBacking 对一块预乘 BGRA 缓冲做原地模糊（三次 box，水平+垂直）。
//
// 边界取样用夹取（clamp）：阴影/发光通常落在透明边，夹取不会把
// 画面外的内容"拉"进来，也不会在边缘产生暗边。
func blurBacking(b *raster.Backing, sigma float32) {
	if b == nil || sigma <= 0.3 {
		return
	}
	w, h := b.W(), b.H()
	if w <= 1 || h <= 1 {
		return
	}
	boxes := boxesForGauss(sigma, 3)
	tmp := make([]uint32, w*h)
	src := b.Pix()
	for _, r := range boxes {
		if r < 1 {
			continue
		}
		boxH(src, tmp, w, h, r)
		boxV(tmp, src, w, h, r)
	}
}

// boxH 水平方向 box blur：src → dst。
func boxH(src, dst []uint32, w, h, r int) {
	if r < 1 || w < 1 || h < 1 {
		copy(dst, src)
		return
	}
	scale := 1 / float32(2*r+1)
	for y := 0; y < h; y++ {
		row := src[y*w : (y+1)*w]
		out := dst[y*w : (y+1)*w]
		var sa, sr, sg, sb uint32
		// 初始化窗口：左侧越界按 clamp 取 row[0]。
		for i := -r; i <= r; i++ {
			p := row[clampCoord(i, w)]
			sa += p >> 24
			sr += (p >> 16) & 0xFF
			sg += (p >> 8) & 0xFF
			sb += p & 0xFF
		}
		for x := 0; x < w; x++ {
			out[x] = packPremul(sr, sg, sb, sa, scale)
			in := row[clampCoord(x+r+1, w)]
			out2 := row[clampCoord(x-r, w)]
			sa += in>>24 - out2>>24
			sr += (in>>16)&0xFF - (out2>>16)&0xFF
			sg += (in>>8)&0xFF - (out2>>8)&0xFF
			sb += in&0xFF - out2&0xFF
		}
	}
}

// boxV 垂直方向 box blur：src → dst。
func boxV(src, dst []uint32, w, h, r int) {
	if r < 1 || w < 1 || h < 1 {
		copy(dst, src)
		return
	}
	scale := 1 / float32(2*r+1)
	for x := 0; x < w; x++ {
		var sa, sr, sg, sb uint32
		for i := -r; i <= r; i++ {
			p := src[clampCoord(i, h)*w+x]
			sa += p >> 24
			sr += (p >> 16) & 0xFF
			sg += (p >> 8) & 0xFF
			sb += p & 0xFF
		}
		for y := 0; y < h; y++ {
			dst[y*w+x] = packPremul(sr, sg, sb, sa, scale)
			in := src[clampCoord(y+r+1, h)*w+x]
			out := src[clampCoord(y-r, h)*w+x]
			sa += in>>24 - out>>24
			sr += (in>>16)&0xFF - (out>>16)&0xFF
			sg += (in>>8)&0xFF - (out>>8)&0xFF
			sb += in&0xFF - out&0xFF
		}
	}
}

// packPremul 把累加和写回预乘 BGRA。
func packPremul(r, g, b, a uint32, scale float32) uint32 {
	pa := uint32(float32(a)*scale + 0.5)
	if pa > 255 {
		pa = 255
	}
	if pa == 0 {
		return 0
	}
	pr := uint32(float32(r)*scale + 0.5)
	pg := uint32(float32(g)*scale + 0.5)
	pb := uint32(float32(b)*scale + 0.5)
	if pr > pa {
		pr = pa
	}
	if pg > pa {
		pg = pa
	}
	if pb > pa {
		pb = pa
	}
	return pa<<24 | pr<<16 | pg<<8 | pb
}

// alphaMask 抽出 alpha 通道（阴影/发光都只需要覆盖率）。
func alphaMask(b *raster.Backing) []uint8 {
	if b == nil {
		return nil
	}
	w, h := b.W(), b.H()
	m := make([]uint8, w*h)
	for i, p := range b.Pix() {
		m[i] = uint8(p >> 24)
	}
	return m
}

// blurMask 对单通道覆盖率做三次 box 模糊。
func blurMask(m []uint8, w, h int, sigma float32) {
	if len(m) < w*h || sigma <= 0.3 {
		return
	}
	boxes := boxesForGauss(sigma, 3)
	tmp := make([]uint8, w*h)
	for _, r := range boxes {
		if r < 1 {
			continue
		}
		maskBoxH(m, tmp, w, h, r)
		maskBoxV(tmp, m, w, h, r)
	}
}

func maskBoxH(src, dst []uint8, w, h, r int) {
	scale := 1 / float32(2*r+1)
	for y := 0; y < h; y++ {
		row := src[y*w : (y+1)*w]
		out := dst[y*w : (y+1)*w]
		sum := 0
		for i := -r; i <= r; i++ {
			sum += int(row[clampCoord(i, w)])
		}
		for x := 0; x < w; x++ {
			v := float32(sum)*scale + 0.5
			if v < 0 {
				v = 0
			}
			if v > 255 {
				v = 255
			}
			out[x] = uint8(v)
			sum += int(row[clampCoord(x+r+1, w)]) - int(row[clampCoord(x-r, w)])
		}
	}
}

func maskBoxV(src, dst []uint8, w, h, r int) {
	scale := 1 / float32(2*r+1)
	for x := 0; x < w; x++ {
		sum := 0
		for i := -r; i <= r; i++ {
			sum += int(src[clampCoord(i, h)*w+x])
		}
		for y := 0; y < h; y++ {
			v := float32(sum)*scale + 0.5
			if v < 0 {
				v = 0
			}
			if v > 255 {
				v = 255
			}
			dst[y*w+x] = uint8(v)
			sum += int(src[clampCoord(y+r+1, h)*w+x]) - int(src[clampCoord(y-r, h)*w+x])
		}
	}
}

// shadowMargin 返回阴影需要向外扩出的像素数（偏移 + 3σ 覆盖 99.7%）。
func shadowMargin(s *scene.ShadowSpec) int {
	if s == nil {
		return 0
	}
	m := int(math.Abs(float64(s.Dx)))
	if a := int(math.Abs(float64(s.Dy))); a > m {
		m = a
	}
	m += int(math.Ceil(float64(s.Blur) * 3))
	return m
}

// clampCoord 把坐标夹到 [0, n-1]（边界夹取）。
func clampCoord(v, n int) int {
	if v < 0 {
		return 0
	}
	if v >= n {
		return n - 1
	}
	return v
}
