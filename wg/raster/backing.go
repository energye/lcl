// Package raster 提供光栅化后端与帧缓冲（backing store）。
//
// 本包只依赖 render（绘制指令/图元模型）与 geom，**不 import lcl**：
// 与 LCL 位图的对接由 host 负责（host 把 DIB 指针交给 Backing.Bytes 视图）。
// 这样 headless 自检可以完全脱离窗口与 LCL 运行。
//
// M0 只落地帧缓冲（backing）与格式约定；光栅化在 M1 接入
// （software/ 扫描线 AA 与 gdi/ 快速路径两个后端）。
package raster

import (
	"unsafe"

	"github.com/energye/lcl/wg/geom"
)

// Format 是帧缓冲像素格式。
type Format uint8

const (
	// FormatBGRA32Premul 是唯一的中间格式：32bpp，小端内存序 B,G,R,A，
	// **预乘 alpha**。与 Windows DIB（BI_RGB 32bpp）和 lcl Pf32bit 内存布局一致，
	// 因此可以直接零拷贝交给 GDI / LCL 位图。
	FormatBGRA32Premul Format = iota

	// FormatBGR24 是最终呈现格式（lcl Pf24bit，契约 C2）。
	// 它只出现在 host 的呈现路径上，由合成结果转换而来，不是中间缓冲格式。
	FormatBGR24
)

// String 返回可读名。
func (f Format) String() string {
	switch f {
	case FormatBGRA32Premul:
		return "bgra32-premul"
	case FormatBGR24:
		return "bgr24"
	default:
		return "unknown"
	}
}

// BytesPerPixel 返回每像素字节数。
func (f Format) BytesPerPixel() int {
	switch f {
	case FormatBGRA32Premul:
		return 4
	case FormatBGR24:
		return 3
	default:
		return 0
	}
}

// Backing 是一块 32bpp 预乘 BGRA 像素缓冲。
//
// 生命周期约束（契约 C21）：释放顺序必须是 视图(Surface) → 位图(Bitmap) → 池(Pool)。
// Backing 自身不持有任何视图，也不注册 finalizer，释放完全由持有者显式控制。
type Backing struct {
	w, h   int
	stride int
	pix    []uint32
	format Format
	// opaque 表示本缓冲当前是否全不透明（预乘下 A==255），
	// 用于跳过多余的 alpha 混合（快速路径）。
	opaque bool
}

// NewBacking 分配 w×h 的帧缓冲；尺寸非法时返回 nil。
//
// stride 固定为 w*4（32bpp 天然 4 字节对齐，无需 padding）。
func NewBacking(w, h int) *Backing {
	if w <= 0 || h <= 0 {
		return nil
	}
	// 防御整数溢出：单缓冲上限 1GiB。
	if int64(w)*int64(h)*4 > 1<<30 {
		return nil
	}
	return &Backing{
		w:      w,
		h:      h,
		stride: w * 4,
		pix:    make([]uint32, w*h),
		format: FormatBGRA32Premul,
		// make 出来的像素为 0，即全透明黑，因此初始不满足 opaque。
		opaque: false,
	}
}

// NewBackingFrom 用已有像素切片构造帧缓冲（用于复用池中的缓冲）。
// 切片的长度必须恰好是 w*h。
func NewBackingFrom(w, h int, pix []uint32) *Backing {
	if w <= 0 || h <= 0 || len(pix) != w*h {
		return nil
	}
	b := &Backing{w: w, h: h, stride: w * 4, pix: pix, format: FormatBGRA32Premul}
	b.opaque = false
	return b
}

// W 返回宽度（像素）。
func (b *Backing) W() int { return b.w }

// H 返回高度（像素）。
func (b *Backing) H() int { return b.h }

// Stride 返回行字节数。
func (b *Backing) Stride() int { return b.stride }

// Format 返回像素格式。
func (b *Backing) Format() Format { return b.format }

// Opaque 表示整块缓冲当前是否完全不透明。
func (b *Backing) Opaque() bool { return b.opaque }

// Bounds 返回像素矩形。
func (b *Backing) Bounds() geom.Rect { return geom.RectXYWH(0, 0, float32(b.w), float32(b.h)) }

// Pix 返回可写像素切片（长度 w*h，行优先）。
func (b *Backing) Pix() []uint32 { return b.pix }

// Row 返回第 y 行的像素切片；y 越界返回 nil。
func (b *Backing) Row(y int) []uint32 {
	if y < 0 || y >= b.h {
		return nil
	}
	return b.pix[y*b.w : (y+1)*b.w]
}

// Bytes 返回底层字节视图（长度 h*stride），可直接作为 DIB 像素指针交给 GDI。
//
// 这是本包唯一需要 unsafe 的地方：uint32 切片 → 字节切片，零拷贝。
// 调用方不得在 Backing 存活期外持有该切片。
func (b *Backing) Bytes() []byte {
	if len(b.pix) == 0 {
		return nil
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(&b.pix[0])), len(b.pix)*4)
}

// Clear 把整块缓冲清零（全透明黑）。
func (b *Backing) Clear() {
	for i := range b.pix {
		b.pix[i] = 0
	}
	b.opaque = false
}

// Fill 用非预乘颜色填满整块缓冲。
func (b *Backing) Fill(c geom.Color) {
	b.FillRect(b.Bounds(), c)
}

// FillRect 用非预乘颜色填充矩形区域（裁到缓冲内，**替换**像素而非混合）。
func (b *Backing) FillRect(r geom.Rect, c geom.Color) {
	x0, y0, x1, y1, ok := b.clipRect(r)
	if !ok {
		return
	}
	v := c.PremulBGRA()
	for y := y0; y < y1; y++ {
		row := b.pix[y*b.w : y*b.w+b.w]
		for x := x0; x < x1; x++ {
			row[x] = v
		}
	}
	// 不透明标记的维护规则：
	//   - 半透明填充 → 整块不再保证不透明；
	//   - 不透明填充且覆盖全幅 → 整块不透明；
	//   - 不透明填充但只覆盖局部 → 保持原状（原来是 opaque 就还是不透明）。
	if !c.IsOpaque() {
		b.opaque = false
	} else if x0 == 0 && y0 == 0 && x1 == b.w && y1 == b.h {
		b.opaque = true
	}
}

// BlendRect 用非预乘颜色以 source-over 混合矩形区域。
func (b *Backing) BlendRect(r geom.Rect, c geom.Color) {
	if c.IsTransparent() {
		return
	}
	if c.IsOpaque() {
		b.FillRect(r, c)
		return
	}
	x0, y0, x1, y1, ok := b.clipRect(r)
	if !ok {
		return
	}
	sr, sg, sb, sa := uint32(c.R), uint32(c.G), uint32(c.B), uint32(c.A)
	for y := y0; y < y1; y++ {
		row := b.pix[y*b.w : y*b.w+b.w]
		for x := x0; x < x1; x++ {
			dst := row[x]
			da := dst >> 24
			dr := (dst >> 16) & 0xFF
			dg := (dst >> 8) & 0xFF
			db := dst & 0xFF
			// src 预乘
			spr := (sr*sa + 128) / 255
			spg := (sg*sa + 128) / 255
			spb := (sb*sa + 128) / 255
			// dst 已预乘；out = src + dst*(1-sa)
			inv := 255 - sa
			oa := sa + da*inv/255
			orr := spr + dr*inv/255
			og := spg + dg*inv/255
			ob := spb + db*inv/255
			row[x] = oa<<24 | orr<<16 | og<<8 | ob
		}
	}
	b.opaque = false
}

// CopyRect 把 src 的矩形区域不透明拷贝到本缓冲的 (dx,dy)。
//
// 这是呈现路径的核心操作：合成结果到 Pf24bit 的那一步必须是一次不透明拷贝
// （契约 C13：帧=帧槽，单次提交），不做 read-modify-write。
func (b *Backing) CopyRect(src *Backing, srcRect geom.Rect, dx, dy int) bool {
	if src == nil {
		return false
	}
	sx0, sy0, sx1, sy1, ok := src.clipRect(srcRect)
	if !ok {
		return false
	}
	w, h := sx1-sx0, sy1-sy0
	tx, ty := dx, dy
	// 目标裁剪：保持宽高比例直接截断
	if tx < 0 {
		sx0 -= tx
		w += tx
		tx = 0
	}
	if ty < 0 {
		sy0 -= ty
		h += ty
		ty = 0
	}
	if tx+w > b.w {
		w = b.w - tx
	}
	if ty+h > b.h {
		h = b.h - ty
	}
	if w <= 0 || h <= 0 {
		return false
	}
	for y := 0; y < h; y++ {
		dst := b.pix[(ty+y)*b.w+tx : (ty+y)*b.w+tx+w]
		srow := src.pix[(sy0+y)*src.w+sx0 : (sy0+y)*src.w+sx0+w]
		copy(dst, srow)
	}
	return true
}

// clipRect 把 DIP 矩形取整并向内裁到缓冲范围，返回半开像素区间。
func (b *Backing) clipRect(r geom.Rect) (x0, y0, x1, y1 int, ok bool) {
	r = r.RoundOut()
	x0 = int(r.Min.X)
	y0 = int(r.Min.Y)
	x1 = int(r.Max.X)
	y1 = int(r.Max.Y)
	if x0 < 0 {
		x0 = 0
	}
	if y0 < 0 {
		y0 = 0
	}
	if x1 > b.w {
		x1 = b.w
	}
	if y1 > b.h {
		y1 = b.h
	}
	if x0 >= x1 || y0 >= y1 {
		return 0, 0, 0, 0, false
	}
	return x0, y0, x1, y1, true
}
