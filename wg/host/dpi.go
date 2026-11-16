package host

import "github.com/energye/lcl/wg/geom"

// ── DPI ────────────────────────────────────────────────────────────
//
// 两条平台事实决定了这一层的存在：
//   - LCL 的 Canvas 用**像素**，布局用 **DIP**。混用会让 125% 缩放下
//     所有文字偏 1/4 像素，AA 一糊就"看起来像没对齐"。
//   - GDI 文字和 1px 描边不做亚像素定位：坐标必须落在整像素上，
//     否则字会发虚（实测 8× 截图最明显的差别之一）。

// baseDPI 是缩放基准（100%）。
const baseDPI = 96

// DPI 是缩放上下文。零值等价于 100%。
type DPI struct{ scale float32 }

// NewDPI 由缩放因子创建（96 DPI = 1.0）。
//
// 非法值（0、负数、NaN）一律回落到 1.0：DPI 来自宿主，
// 它算错时整个界面会缩成 0 像素，静默故障比崩溃更难查。
func NewDPI(scale float32) DPI {
	if !(scale > 0) || scale > 8 {
		return DPI{scale: 1}
	}
	return DPI{scale: scale}
}

// DPIFor 由整数 DPI 创建（96 → 1.0，144 → 1.5）。
func DPIFor(dpi int) DPI {
	if dpi <= 0 {
		return DPI{scale: 1}
	}
	return NewDPI(float32(dpi) / baseDPI)
}

// Scale 返回缩放因子。
func (d DPI) Scale() float32 {
	if d.scale <= 0 {
		return 1
	}
	return d.scale
}

// Percent 返回缩放百分比（100 / 125 / 150 …）。
func (d DPI) Percent() int { return int(d.Scale()*100 + 0.5) }

// IsIdentity 表示 100%。
func (d DPI) IsIdentity() bool { return d.Scale() == 1 }

// Pixels 把 DIP 换算为像素（浮点，用于绘制路径）。
func (d DPI) Pixels(dip float32) float32 { return dip * d.Scale() }

// DIP 把像素换算为 DIP（输入坐标永远走这条）。
func (d DPI) DIP(px float32) float32 { return px / d.Scale() }

// PixelsI 把 DIP 换算成整像素（四舍五入）。
func (d DPI) PixelsI(dip float32) int32 {
	return int32(dip*d.Scale() + 0.5)
}

// DIPI 把整像素换算成 DIP。
func (d DPI) DIPI(px int32) float32 { return float32(px) / d.Scale() }

// SnapPoint 把点对齐到整像素。
func (d DPI) SnapPoint(p geom.Point) geom.Point {
	return geom.Pt(float32(d.PixelsI(p.X))/d.Scale(), float32(d.PixelsI(p.Y))/d.Scale())
}

// SnapRect 把矩形四边对齐到整像素。
//
// 分别处理两条边而不是"对齐左上 + 圆整宽高"：宽高再圆整会引入 ±1px
// 累积误差，表格里几十个单元格叠起来就是明显的错位。
func (d DPI) SnapRect(r geom.Rect) geom.Rect {
	x0 := d.PixelsI(r.Min.X)
	y0 := d.PixelsI(r.Min.Y)
	x1 := d.PixelsI(r.Max.X)
	y1 := d.PixelsI(r.Max.Y)
	if x1 < x0 {
		x1 = x0
	}
	if y1 < y0 {
		y1 = y0
	}
	return geom.RectXYWH(d.DIPI(x0), d.DIPI(y0), d.DIPI(x1-x0), d.DIPI(y1-y0))
}
