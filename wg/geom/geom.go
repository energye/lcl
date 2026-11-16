// Package geom 提供 WG Engine 的纯几何类型：点、尺寸、矩形、矩阵、颜色。
//
// 本包零依赖，不 import 任何其它 wg 包，也不 import lcl。
// 布局只算 DIP（device independent pixel），乘 DPI 发生在绘制期，
// 因此所有坐标与尺寸统一使用 float32（见 02-CANVAS.md §4.2）。
//
// 约定：
//   - 矩形半开区间：[Min.X, Max.X) × [Min.Y, Max.Y)。
//   - 矩阵为 2D 仿射：(x', y') = (a*x + c*y + e, b*x + d*y + f)，与 Skia/Cairo 一致。
//   - 颜色为**非预乘** sRGB8，预乘转换发生在光栅边界（geom.Color.PremulBGRA）。
package geom

// Point 是二维点/向量，单位 DIP。
type Point struct {
	X, Y float32
}

// Pt 构造 Point。
func Pt(x, y float32) Point { return Point{X: x, Y: y} }

// Add 返回 p+v。
func (p Point) Add(v Point) Point { return Point{p.X + v.X, p.Y + v.Y} }

// Sub 返回 p-v。
func (p Point) Sub(v Point) Point { return Point{p.X - v.X, p.Y - v.Y} }

// Mul 返回 p*k。
func (p Point) Mul(k float32) Point { return Point{p.X * k, p.Y * k} }

// Dot 返回点积。
func (p Point) Dot(v Point) float32 { return p.X*v.X + p.Y*v.Y }

// Len 返回向量长度。
func (p Point) Len() float32 { return sqrt(p.X*p.X + p.Y*p.Y) }

// Size 是宽高，单位 DIP。
type Size struct {
	W, H float32
}

// Sz 构造 Size。
func Sz(w, h float32) Size { return Size{W: w, H: h} }

// IsEmpty 表示宽或高 <= 0。
func (s Size) IsEmpty() bool { return s.W <= 0 || s.H <= 0 }

// Rect 是轴对齐矩形，半开区间。
type Rect struct {
	Min, Max Point
}

// RectXYWH 用左上角与宽高构造 Rect。
func RectXYWH(x, y, w, h float32) Rect {
	return Rect{Min: Point{x, y}, Max: Point{x + w, y + h}}
}

// RectLTRB 用四边构造 Rect。
func RectLTRB(l, t, r, b float32) Rect { return Rect{Min: Point{l, t}, Max: Point{r, b}} }

// ZeroRect 是空矩形（Min=Max=0）。
var ZeroRect = Rect{}

// W 返回宽度（不裁剪负宽）。
func (r Rect) W() float32 { return r.Max.X - r.Min.X }

// H 返回高度。
func (r Rect) H() float32 { return r.Max.Y - r.Min.Y }

// Size 返回宽高。
func (r Rect) Size() Size { return Size{W: r.W(), H: r.H()} }

// Left/Top/Right/Bottom 为四边别名。
func (r Rect) Left() float32   { return r.Min.X }
func (r Rect) Top() float32    { return r.Min.Y }
func (r Rect) Right() float32  { return r.Max.X }
func (r Rect) Bottom() float32 { return r.Max.Y }

// Center 返回中心点。
func (r Rect) Center() Point {
	return Point{(r.Min.X + r.Max.X) * 0.5, (r.Min.Y + r.Max.Y) * 0.5}
}

// IsEmpty 表示矩形没有正面积。
func (r Rect) IsEmpty() bool { return r.Max.X <= r.Min.X || r.Max.Y <= r.Min.Y }

// Offset 平移。
func (r Rect) Offset(dx, dy float32) Rect {
	return Rect{Min: Point{r.Min.X + dx, r.Min.Y + dy}, Max: Point{r.Max.X + dx, r.Max.Y + dy}}
}

// OffsetPt 用点平移。
func (r Rect) OffsetPt(d Point) Rect { return r.Offset(d.X, d.Y) }

// Inset 四边内缩（负数即外扩）。
func (r Rect) Inset(d float32) Rect { return r.InsetXY(d, d) }

// InsetXY 分别内缩 X/Y。
func (r Rect) InsetXY(dx, dy float32) Rect {
	return Rect{Min: Point{r.Min.X + dx, r.Min.Y + dy}, Max: Point{r.Max.X - dx, r.Max.Y - dy}}
}

// InsetEdges 按四边分别内缩。
func (r Rect) InsetEdges(l, t, rt, b float32) Rect {
	return Rect{Min: Point{r.Min.X + l, r.Min.Y + t}, Max: Point{r.Max.X - rt, r.Max.Y - b}}
}

// Contains 判定点是否落在半开区间内。
func (r Rect) Contains(p Point) bool {
	return p.X >= r.Min.X && p.X < r.Max.X && p.Y >= r.Min.Y && p.Y < r.Max.Y
}

// ContainsRect 判定另一矩形是否完全包含。
func (r Rect) ContainsRect(o Rect) bool {
	return o.Min.X >= r.Min.X && o.Min.Y >= r.Min.Y && o.Max.X <= r.Max.X && o.Max.Y <= r.Max.Y
}

// Intersects 判定是否与 o 有正面积交集。
func (r Rect) Intersects(o Rect) bool {
	return r.Min.X < o.Max.X && o.Min.X < r.Max.X && r.Min.Y < o.Max.Y && o.Min.Y < r.Max.Y
}

// Intersect 返回交集；无交集时返回 ZeroRect。
func (r Rect) Intersect(o Rect) Rect {
	if !r.Intersects(o) {
		return ZeroRect
	}
	return Rect{
		Min: Point{maxf(r.Min.X, o.Min.X), maxf(r.Min.Y, o.Min.Y)},
		Max: Point{minf(r.Max.X, o.Max.X), minf(r.Max.Y, o.Max.Y)},
	}
}

// Union 返回外包矩形；任一为空时返回另一个。
func (r Rect) Union(o Rect) Rect {
	if r.IsEmpty() {
		return o
	}
	if o.IsEmpty() {
		return r
	}
	return Rect{
		Min: Point{minf(r.Min.X, o.Min.X), minf(r.Min.Y, o.Min.Y)},
		Max: Point{maxf(r.Max.X, o.Max.X), maxf(r.Max.Y, o.Max.Y)},
	}
}

// Expand 四边外扩 d，等价 Inset(-d)。
func (r Rect) Expand(d float32) Rect { return r.Inset(-d) }

// RoundOut 向外取整到整数像素（damage 对齐用）。
func (r Rect) RoundOut() Rect {
	if r.IsEmpty() {
		return ZeroRect
	}
	return Rect{
		Min: Point{floorf(r.Min.X), floorf(r.Min.Y)},
		Max: Point{ceilf(r.Max.X), ceilf(r.Max.Y)},
	}
}

// RoundIn 向内取整到整数像素。
func (r Rect) RoundIn() Rect {
	out := Rect{Min: Point{ceilf(r.Min.X), ceilf(r.Min.Y)}, Max: Point{floorf(r.Max.X), floorf(r.Max.Y)}}
	if out.IsEmpty() {
		return ZeroRect
	}
	return out
}

// Insets 是四边内边距，单位 DIP。
type Insets struct {
	Left, Top, Right, Bottom float32
}

// In 构造四边相同的 Insets。
func In(v float32) Insets { return Insets{v, v, v, v} }

// InXY 分别构造水平/垂直 Insets。
func InXY(h, v float32) Insets { return Insets{h, v, h, v} }

// InLTRB 按四边构造。
func InLTRB(l, t, r, b float32) Insets { return Insets{l, t, r, b} }

// Horizontal 返回左右之和。
func (i Insets) Horizontal() float32 { return i.Left + i.Right }

// Vertical 返回上下之和。
func (i Insets) Vertical() float32 { return i.Top + i.Bottom }

// Deflate 按 Insets 收缩矩形。
func (i Insets) Deflate(r Rect) Rect { return r.InsetEdges(i.Left, i.Top, i.Right, i.Bottom) }

// Inflate 按 Insets 扩张矩形。
func (i Insets) Inflate(r Rect) Rect { return r.InsetEdges(-i.Left, -i.Top, -i.Right, -i.Bottom) }

// Matrix 是 2D 仿射矩阵，列主序习惯的六元组。
type Matrix struct {
	A, B, C, D, E, F float32
}

// Identity 是单位矩阵。
var Identity = Matrix{A: 1, D: 1}

// IsIdentity 判定是否为单位矩阵。
func (m Matrix) IsIdentity() bool {
	return m == Identity
}

// Translate 构造平移矩阵。
func Translate(dx, dy float32) Matrix { return Matrix{A: 1, D: 1, E: dx, F: dy} }

// Scale 构造缩放矩阵。
func Scale(sx, sy float32) Matrix { return Matrix{A: sx, D: sy} }

// ScaleUniform 构造等比缩放矩阵。
func ScaleUniform(s float32) Matrix { return Scale(s, s) }

// Rotate 构造旋转矩阵，rad 为弧度。
func Rotate(rad float32) Matrix {
	s, c := sinf(rad), cosf(rad)
	return Matrix{A: c, B: s, C: -s, D: c}
}

// Concat 返回 m∘o（先应用 o，再应用 m）。
func (m Matrix) Concat(o Matrix) Matrix {
	return Matrix{
		A: m.A*o.A + m.C*o.B,
		B: m.B*o.A + m.D*o.B,
		C: m.A*o.C + m.C*o.D,
		D: m.B*o.C + m.D*o.D,
		E: m.A*o.E + m.C*o.F + m.E,
		F: m.B*o.E + m.D*o.F + m.F,
	}
}

// MapPoint 变换点。
func (m Matrix) MapPoint(p Point) Point {
	return Point{
		X: m.A*p.X + m.C*p.Y + m.E,
		Y: m.B*p.X + m.D*p.Y + m.F,
	}
}

// MapVector 变换向量（忽略平移）。
func (m Matrix) MapVector(v Point) Point {
	return Point{X: m.A*v.X + m.C*v.Y, Y: m.B*v.X + m.D*v.Y}
}

// MapRect 变换轴对齐矩形，返回外包矩形（旋转时变大，符合预期）。
func (m Matrix) MapRect(r Rect) Rect {
	if r.IsEmpty() {
		return ZeroRect
	}
	p0 := m.MapPoint(r.Min)
	p1 := m.MapPoint(Point{r.Max.X, r.Min.Y})
	p2 := m.MapPoint(r.Max)
	p3 := m.MapPoint(Point{r.Min.X, r.Max.Y})
	minX, maxX := minf(minf(p0.X, p1.X), minf(p2.X, p3.X)), maxf(maxf(p0.X, p1.X), maxf(p2.X, p3.X))
	minY, maxY := minf(minf(p0.Y, p1.Y), minf(p2.Y, p3.Y)), maxf(maxf(p0.Y, p1.Y), maxf(p2.Y, p3.Y))
	return RectLTRB(minX, minY, maxX, maxY)
}

// IsAxisAligned 判定矩阵是否只含平移与正缩放（无旋转/斜切）。
func (m Matrix) IsAxisAligned() bool {
	return m.B == 0 && m.C == 0
}

// Invert 求逆；不可逆时第二返回值为 false。
func (m Matrix) Invert() (Matrix, bool) {
	det := m.A*m.D - m.B*m.C
	if det == 0 {
		return Identity, false
	}
	id := 1 / det
	return Matrix{
		A: m.D * id,
		B: -m.B * id,
		C: -m.C * id,
		D: m.A * id,
		E: (m.C*m.F - m.D*m.E) * id,
		F: (m.B*m.E - m.A*m.F) * id,
	}, true
}

// ScaleFactor 返回平均缩放因子（用于线宽/模糊半径换算）。
func (m Matrix) ScaleFactor() float32 {
	sx := sqrt(m.A*m.A + m.B*m.B)
	sy := sqrt(m.C*m.C + m.D*m.D)
	return (sx + sy) * 0.5
}

// Color 是非预乘 sRGB8 颜色。
type Color struct {
	R, G, B, A uint8
}

// RGBA 构造不透明颜色。
func RGBA(r, g, b uint8) Color { return Color{r, g, b, 255} }

// RGBAa 构造带 alpha 颜色。
func RGBAa(r, g, b, a uint8) Color { return Color{r, g, b, a} }

// Hex 用 0xRRGGBB 构造不透明颜色。
func Hex(rgb uint32) Color {
	return Color{uint8(rgb >> 16), uint8(rgb >> 8), uint8(rgb), 255}
}

// HexA 用 0xAARRGGBB 构造颜色。
func HexA(argb uint32) Color {
	return Color{uint8(argb >> 16), uint8(argb >> 8), uint8(argb), uint8(argb >> 24)}
}

// WithAlpha 返回替换 alpha 的颜色。
func (c Color) WithAlpha(a uint8) Color { c.A = a; return c }

// IsOpaque 判定是否完全不透明。
func (c Color) IsOpaque() bool { return c.A == 255 }

// IsTransparent 判定是否完全透明。
func (c Color) IsTransparent() bool { return c.A == 0 }

// Pre-multiply 用的整数常量。
const (
	// mulDiv255 的舍入补偿：x*255/255 精确。
	half = 128
)

// PremulBGRA 返回预乘后的 BGRA 打包值（小端内存布局为 B,G,R,A，
// 与 Windows DIB / lcl Pf32bit 的内存顺序一致）。
func (c Color) PremulBGRA() uint32 {
	if c.A == 0 {
		return 0
	}
	if c.A == 255 {
		return uint32(c.R)<<16 | uint32(c.G)<<8 | uint32(c.B) | 0xFF000000
	}
	a := uint32(c.A)
	r := (uint32(c.R)*a + half) / 255
	g := (uint32(c.G)*a + half) / 255
	b := (uint32(c.B)*a + half) / 255
	return r<<16 | g<<8 | b | a<<24
}

// Lerp 线性插值（用于动画与渐变）。
func (c Color) Lerp(o Color, t float32) Color {
	if t <= 0 {
		return c
	}
	if t >= 1 {
		return o
	}
	return Color{
		R: uint8(float32(c.R) + (float32(o.R)-float32(c.R))*t + 0.5),
		G: uint8(float32(c.G) + (float32(o.G)-float32(c.G))*t + 0.5),
		B: uint8(float32(c.B) + (float32(o.B)-float32(c.B))*t + 0.5),
		A: uint8(float32(c.A) + (float32(o.A)-float32(c.A))*t + 0.5),
	}
}

// Clamp01 把 v 限制到 [0,1]。
func Clamp01(v float32) float32 { return minf(1, maxf(0, v)) }

// Clamp 把 v 限制到 [lo,hi]。
func Clamp(v, lo, hi float32) float32 { return minf(hi, maxf(lo, v)) }
