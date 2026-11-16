package render

import (
	"image"

	"github.com/energye/lcl/wg/geom"
)

// ImageQuality 是缩放采样质量。
type ImageQuality uint8

const (
	// ImageNearest 最近邻（图标放大、像素画）。
	ImageNearest ImageQuality = iota
	// ImageBilinear 双线性（默认）。
	ImageBilinear
	// ImageMipmap 预乘 mip 采样（大幅缩小，M2 落地）。
	ImageMipmap
)

// ImageRef 是绘制用的位图引用。
//
// 像素格式固定为**预乘 BGRA**（与 raster.Backing 一致），这样软件后端
// 可以零转换直接混色。来源（PNG/JPEG/截图）在解码时统一转成这个格式。
//
// M1 只做类型与解码，软件后端的 DrawImage 在 M2 落地（当前计入 degraded）。
type ImageRef struct {
	W, H   int
	Stride int
	// Pix 是预乘 BGRA 像素；索引方式 y*Stride+x。
	Pix []uint32
}

// NewImageRef 分配一张透明位图。
func NewImageRef(w, h int) *ImageRef {
	if w <= 0 || h <= 0 {
		return nil
	}
	return &ImageRef{W: w, H: h, Stride: w, Pix: make([]uint32, w*h)}
}

// Bounds 返回像素尺寸（DIP 由调用方决定缩放）。
func (img *ImageRef) Bounds() geom.Rect {
	if img == nil {
		return geom.ZeroRect
	}
	return geom.RectXYWH(0, 0, float32(img.W), float32(img.H))
}

// ImageFrom 把标准库 image.Image 转成 ImageRef（预乘 BGRA）。
//
// 逐像素转换，属于一次性成本（解码后缓存 ImageRef 复用）。
func ImageFrom(src image.Image) *ImageRef {
	if src == nil {
		return nil
	}
	b := src.Bounds()
	w, h := b.Dx(), b.Dy()
	out := NewImageRef(w, h)
	if out == nil {
		return nil
	}
	for y := 0; y < h; y++ {
		row := out.Pix[y*out.Stride : y*out.Stride+w]
		for x := 0; x < w; x++ {
			r, g, bl, a := src.At(b.Min.X+x, b.Min.Y+y).RGBA()
			c := geom.Color{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(bl >> 8), A: uint8(a >> 8)}
			row[x] = c.PremulBGRA()
		}
	}
	return out
}
