//go:build windows

package gdi

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
)

// Painter 把 GDI 文字通道接到 composite 的 TextPainter 接口上。
//
// 它内部维护一块 24bpp DIB：先把合成结果转成 BGR（ClearType 需要 24/32bpp
// 的真实设备格式，且在 32bpp 上 GDI 会写坏 alpha），回放文字后再转回
// 预乘 BGRA。这次往返是刻意的，也是"文字必须在像素写回后回放"的物理体现。
type Painter struct {
	dib   *DIB
	areas []geom.Rect
}

// NewPainter 创建文字通道。
func NewPainter() *Painter { return &Painter{} }

// PaintText 实现 composite.TextPainter。
//
// 只转换**文字覆盖的矩形**，不做整帧往返：
// 一屏 1440×880 的 32bpp↔24bpp 来回转换约 250 万像素，实测占一帧开销的一半以上；
// 而文字通常只占几千像素。按矩形转换后，同一场景的转换量下降两个数量级。
//
// 转换与回放都是**批量**的：先把所有文字区域的像素搬进 DIB，再一次性回放
// 全部文字（TextPass 内部按"变化才重设"维护 GDI 状态），最后一次性搬回。
// 逐区域回放会让字体在每个区域开头被重新 SelectObject，触发 GDI 字体缓存抖动。
func (p *Painter) PaintText(dst *raster.Backing, dl *render.DisplayList, clip geom.Rect) int {
	if dst == nil || dl == nil || !dl.HasText() {
		return 0
	}
	if p.dib == nil || p.dib.W != dst.W() || p.dib.H != dst.H() {
		if p.dib != nil {
			p.dib.Delete()
		}
		p.dib = NewDIB(dst.W(), dst.H(), 24)
		if p.dib == nil {
			return 0
		}
	}

	// 汇总文字指令覆盖的矩形（与裁剪求交），这就是需要格式转换的区域。
	p.areas = p.areas[:0]
	for _, cmd := range dl.Commands() {
		if cmd.Op != render.OpDrawText {
			continue
		}
		if a := cmd.Rect.Intersect(clip); !a.IsEmpty() {
			p.areas = append(p.areas, a.RoundOut())
		}
	}
	if len(p.areas) == 0 {
		return 0
	}

	for _, a := range p.areas {
		toBGR24Rect(dst, p.dib, a)
	}
	n := TextPass(p.dib, dl, clip)
	if n == 0 {
		return 0
	}
	p.dib.Flush()
	for _, a := range p.areas {
		fromBGR24Rect(p.dib, dst, a)
	}
	return n
}

// toBGR24Rect 把矩形区域从预乘 BGRA 转成 24bpp BGR（与白底合成）。
func toBGR24Rect(src *raster.Backing, d *DIB, r geom.Rect) {
	x0, y0, x1, y1 := clampRect(r, src.W(), src.H(), d.W, d.H)
	for y := y0; y < y1; y++ {
		srow := src.Row(y)
		drow := d.Row(y)
		for x := x0; x < x1; x++ {
			px := srow[x]
			a := px >> 24
			var rr, gg, bb uint32
			switch {
			case a == 255:
				rr, gg, bb = (px>>16)&0xFF, (px>>8)&0xFF, px&0xFF
			case a == 0:
				rr, gg, bb = 255, 255, 255
			default:
				// 预乘色与白底合成：out = premul + (255 - a)
				//（因为 premul 已经等于 unpremul*a/255，白底贡献 255*(255-a)/255）
				inv := uint32(255 - a)
				rr = min32(255, ((px>>16)&0xFF)+inv)
				gg = min32(255, ((px>>8)&0xFF)+inv)
				bb = min32(255, (px&0xFF)+inv)
			}
			drow[x*3] = byte(bb)
			drow[x*3+1] = byte(gg)
			drow[x*3+2] = byte(rr)
		}
	}
}

// fromBGR24Rect 把矩形区域从 24bpp BGR 转回不透明预乘 BGRA。
func fromBGR24Rect(d *DIB, dst *raster.Backing, r geom.Rect) {
	x0, y0, x1, y1 := clampRect(r, dst.W(), dst.H(), d.W, d.H)
	for y := y0; y < y1; y++ {
		srow := d.Row(y)
		drow := dst.Row(y)
		for x := x0; x < x1; x++ {
			bb := uint32(srow[x*3])
			gg := uint32(srow[x*3+1])
			rr := uint32(srow[x*3+2])
			drow[x] = 0xFF000000 | rr<<16 | gg<<8 | bb
		}
	}
}

// clampRect 把矩形裁到两边的像素范围内，返回半开区间。
func clampRect(r geom.Rect, w1, h1, w2, h2 int) (x0, y0, x1, y1 int) {
	x0, y0 = int(r.Min.X), int(r.Min.Y)
	x1, y1 = int(r.Max.X+0.5), int(r.Max.Y+0.5)
	if x0 < 0 {
		x0 = 0
	}
	if y0 < 0 {
		y0 = 0
	}
	if x1 > w1 {
		x1 = w1
	}
	if x1 > w2 {
		x1 = w2
	}
	if y1 > h1 {
		y1 = h1
	}
	if y1 > h2 {
		y1 = h2
	}
	return
}

// Close 释放内部 DIB。
func (p *Painter) Close() {
	if p.dib != nil {
		p.dib.Delete()
		p.dib = nil
	}
}

func min32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}
