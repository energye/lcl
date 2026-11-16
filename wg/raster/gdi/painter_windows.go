//go:build windows

package gdi

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
)

// Painter 把 GDI 文字通道接到 composite 的 TextPainter 接口上。
//
// 它内部维护一块 24bpp DIB（ClearType 需要 24/32bpp 的真实设备格式，且在
// 32bpp 上 GDI 会写坏 alpha），回放文字后把结果搬回目标缓冲。两条路径：
//
//   - 目标区域全不透明（主缓冲）：单趟白底往返，文字由 GDI 直接混进
//     已有像素，成本与实现都最简；
//   - 目标含透明像素（组层离屏）：黑白双底各渲染一次，从两趟的差里
//     逐通道提取文字覆盖率，按预乘合成回去。GDI 文字只写 RGB，直接
//     画进透明缓冲会把矩形写成不透明——双趟提取是文字能进组层
//     （透明度/裁剪蒙版/滤镜/缓存全部自动生效）的根本。
type Painter struct {
	dib   *DIB
	saved []byte // 黑底趟快照（BGR 三字节/像素，与 dib 同尺寸）
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
			p.saved = nil
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

	// 快路径：目标区域全不透明（主缓冲清屏色不透明），单趟白底往返。
	if backingOpaque(dst, p.areas) {
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

	// alpha 路径：黑白双底各渲染一次，逐通道提取覆盖率后按预乘合成。
	//
	// 数学：黑底趟 P_b = a·C（预乘文字色），白底趟 P_w = P_b + (1-a)·255，
	// 所以 (P_w-P_b) 就是 255·(1-a)——"背景保留系数"，逐通道精确，
	// ClearType 彩色边缘保留，重叠文字栈的组合覆盖率也精确。
	// 两趟必须回放同一份 dl（同字体同坐标），几何才逐像素对齐。
	for _, a := range p.areas {
		fillDIBRect(p.dib, a, 0x00)
	}
	n := TextPass(p.dib, dl, clip)
	if n == 0 {
		return 0
	}
	p.dib.Flush()
	p.saved = saveDIBRects(p.dib, p.areas, p.saved)
	for _, a := range p.areas {
		fillDIBRect(p.dib, a, 0xFF)
	}
	TextPass(p.dib, dl, clip)
	p.dib.Flush()
	for _, a := range p.areas {
		compositeTextAlpha(dst, p.dib, p.saved, a)
	}
	return n
}

// backingOpaque 报告各矩形区域是否全部不透明（alpha=255）。
// 只扫文字覆盖的像素，主缓冲一次命中，成本可忽略。
func backingOpaque(dst *raster.Backing, areas []geom.Rect) bool {
	for _, r := range areas {
		x0, y0, x1, y1 := clampRect(r, dst.W(), dst.H(), dst.W(), dst.H())
		for y := y0; y < y1; y++ {
			row := dst.Row(y)
			for x := x0; x < x1; x++ {
				if row[x]>>24 != 255 {
					return false
				}
			}
		}
	}
	return true
}

// fillDIBRect 把 DIB 的矩形区域填成灰度 v（BGR 三字节同值）。
func fillDIBRect(d *DIB, r geom.Rect, v byte) {
	x0, y0, x1, y1 := clampRect(r, d.W, d.H, d.W, d.H)
	for y := y0; y < y1; y++ {
		row := d.Row(y)
		for x := x0; x < x1; x++ {
			row[x*3] = v
			row[x*3+1] = v
			row[x*3+2] = v
		}
	}
}

// saveDIBRects 把 DIB 的各矩形区域拷进 buf（BGR 三字节/像素），
// 返回复用/扩容后的 buf。作为黑底趟的快照。
func saveDIBRects(d *DIB, areas []geom.Rect, buf []byte) []byte {
	if len(buf) < d.W*d.H*3 {
		buf = make([]byte, d.W*d.H*3)
	}
	for _, r := range areas {
		x0, y0, x1, y1 := clampRect(r, d.W, d.H, d.W, d.H)
		for y := y0; y < y1; y++ {
			row := d.Row(y)
			copy(buf[(y*d.W+x0)*3:(y*d.W+x1)*3], row[x0*3:x1*3])
		}
	}
	return buf
}

// compositeTextAlpha 把黑底趟快照（saved）与白底趟（dib）的差作为文字
// 覆盖率，按预乘 source-over 合成进 dst 的矩形区域。
//
// 逐通道：inv = P_w - P_b 是 255·(1-a_ch)，out = P_b + inv·dst/255。
// alpha 字节取三通道覆盖率的最大值（inv 取最小），保证预乘不变式
// （RGB ≤ A）；ClearType 彩色边缘处三通道 a 不同，逐通道合成保住了
// 次像素着色，这是单覆盖率近似做不到的。
func compositeTextAlpha(dst *raster.Backing, d *DIB, saved []byte, r geom.Rect) {
	x0, y0, x1, y1 := clampRect(r, dst.W(), dst.H(), d.W, d.H)
	w := d.W
	for y := y0; y < y1; y++ {
		srow := d.Row(y)
		brow := saved[(y*w+x0)*3 : (y*w+x1)*3]
		drow := dst.Row(y)
		for x := x0; x < x1; x++ {
			i := (x - x0) * 3
			// 黑底趟 P_b（saved），白底趟 P_w（dib）
			pbb, pbg, pbr := int(brow[i]), int(brow[i+1]), int(brow[i+2])
			pwb := int(srow[x*3])
			pwg := int(srow[x*3+1])
			pwr := int(srow[x*3+2])
			invb, invg, invr := pwb-pbb, pwg-pbg, pwr-pbr // 255·(1-a_ch)
			a := 255 - mini3(invb, invg, invr)            // 覆盖率最大通道
			dp := drow[x]
			da := int(dp >> 24)
			or := pbr + invr*int((dp>>16)&0xFF)/255
			og := pbg + invg*int((dp>>8)&0xFF)/255
			ob := pbb + invb*int(dp&0xFF)/255
			oa := a + da*(255-a)/255
			// 量化取整可能让单通道略超 oa，夹一次保住预乘不变式。
			if or > oa {
				or = oa
			}
			if og > oa {
				og = oa
			}
			if ob > oa {
				ob = oa
			}
			drow[x] = uint32(oa)<<24 | uint32(or)<<16 | uint32(og)<<8 | uint32(ob)
		}
	}
}

func mini3(a, b, c int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	return a
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
	p.saved = nil
}

func min32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}
