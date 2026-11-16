package render

import "github.com/energye/lcl/wg/geom"

// Style 是绘制方式。
type Style uint8

const (
	// StyleFill 填充。
	StyleFill Style = iota
	// StyleStroke 描边。
	StyleStroke
)

// String 返回可读名。
func (s Style) String() string {
	if s == StyleStroke {
		return "stroke"
	}
	return "fill"
}

// Cap 是线端样式（对齐 Skia/Cairo）。
type Cap uint8

const (
	// CapButt 平头：不做任何延伸。
	CapButt Cap = iota
	// CapRound 圆头：以线宽一半为半径加半圆。
	CapRound
	// CapSquare 方头：延伸线宽一半。
	CapSquare
)

// String 返回可读名。
func (c Cap) String() string {
	switch c {
	case CapRound:
		return "round"
	case CapSquare:
		return "square"
	default:
		return "butt"
	}
}

// Join 是拐角样式。
type Join uint8

const (
	// JoinMiter 尖角：受 MiterLimit 约束，超限退化为 bevel。
	JoinMiter Join = iota
	// JoinRound 圆角。
	JoinRound
	// JoinBevel 斜切。
	JoinBevel
)

// String 返回可读名。
func (j Join) String() string {
	switch j {
	case JoinRound:
		return "round"
	case JoinBevel:
		return "bevel"
	default:
		return "miter"
	}
}

// FillRule 是填充规则。
type FillRule uint8

const (
	// NonZero 非零环绕（默认，适合自交与复合轮廓）。
	NonZero FillRule = iota
	// EvenOdd 奇偶规则（适合环形/镂空）。
	EvenOdd
)

// String 返回可读名。
func (f FillRule) String() string {
	if f == EvenOdd {
		return "even-odd"
	}
	return "non-zero"
}

// Spread 是渐变在 [0,1] 之外的延展方式。
type Spread uint8

const (
	// SpreadPad 钳制到端点色（默认）。
	SpreadPad Spread = iota
	// SpreadRepeat 重复。
	SpreadRepeat
	// SpreadReflect 镜像重复。
	SpreadReflect
)

// Stop 是渐变色标。
type Stop struct {
	Offset float32
	Color  geom.Color
}

// ShaderKind 是渐变类型。
type ShaderKind uint8

const (
	// ShaderNone 表示无渐变（纯色）。
	ShaderNone ShaderKind = iota
	// ShaderLinear 线性渐变：沿 Start→End 投影取 t。
	ShaderLinear
	// ShaderRadial 径向渐变：按到 Center 的距离 / Radius 取 t。
	ShaderRadial
)

// String 返回可读名。
func (k ShaderKind) String() string {
	switch k {
	case ShaderLinear:
		return "linear"
	case ShaderRadial:
		return "radial"
	default:
		return "none"
	}
}

// Shader 是渐变定义（值类型，便于放进 side table 无额外分配）。
//
// 只做两种：线性与径向。扫描渐变（conic）在图表里有需求，但 M4 再补——
// 它的采样逻辑与另外两种不同，不适合在 M1 抢工期。
type Shader struct {
	Kind ShaderKind
	// Start/End 是线性渐变的两个端点（局部坐标）。
	Start, End geom.Point
	// Center 是径向渐变的圆心，Radius 是半径。
	Center geom.Point
	Radius float32
	// Stops 必须按 Offset 升序；采样时线性插值。
	Stops  []Stop
	Spread Spread
}

// NewLinearGradient 构造线性渐变。
func NewLinearGradient(from, to geom.Point, stops ...Stop) *Shader {
	return &Shader{Kind: ShaderLinear, Start: from, End: to, Stops: stops, Spread: SpreadPad}
}

// NewRadialGradient 构造径向渐变。
func NewRadialGradient(center geom.Point, radius float32, stops ...Stop) *Shader {
	return &Shader{Kind: ShaderRadial, Center: center, Radius: radius, Stops: stops, Spread: SpreadPad}
}

// Sample 在归一化位置 t 采样颜色（t 会按 Spread 处理）。
func (s *Shader) Sample(t float32) geom.Color {
	if s == nil || len(s.Stops) == 0 {
		return geom.Color{}
	}
	t = s.applySpread(t)
	stops := s.Stops
	if t <= stops[0].Offset {
		return stops[0].Color
	}
	last := stops[len(stops)-1]
	if t >= last.Offset {
		return last.Color
	}
	// 二分定位区间：色标通常很少，线性扫描也够，但二分更稳。
	lo, hi := 0, len(stops)-1
	for hi-lo > 1 {
		mid := (lo + hi) / 2
		if stops[mid].Offset <= t {
			lo = mid
		} else {
			hi = mid
		}
	}
	a, b := stops[lo], stops[hi]
	span := b.Offset - a.Offset
	if span <= 0 {
		return b.Color
	}
	return a.Color.Lerp(b.Color, (t-a.Offset)/span)
}

// applySpread 把 t 映射回 [0,1]。
func (s *Shader) applySpread(t float32) float32 {
	switch s.Spread {
	case SpreadRepeat:
		if t >= 0 && t <= 1 {
			return t
		}
		t -= floorf(t) // 取小数部分，负数也能落到 [0,1)
		return t
	case SpreadReflect:
		if t >= 0 && t <= 1 {
			return t
		}
		k := floorf(t)
		t -= k
		if int(k)&1 != 0 {
			t = 1 - t
		}
		return t
	default: // SpreadPad
		if t < 0 {
			return 0
		}
		if t > 1 {
			return 1
		}
		return t
	}
}

// GradientT 返回点 p 在渐变中的归一化位置。
func (s *Shader) GradientT(p geom.Point) float32 {
	switch s.Kind {
	case ShaderLinear:
		d := s.End.Sub(s.Start)
		den := d.Dot(d)
		if den == 0 {
			return 0
		}
		return p.Sub(s.Start).Dot(d) / den
	case ShaderRadial:
		if s.Radius == 0 {
			return 0
		}
		return p.Sub(s.Center).Len() / s.Radius
	default:
		return 0
	}
}

// IsTranslucent 表示渐变是否含半透明色标（决定能否走不透明快速路径）。
func (s *Shader) IsTranslucent() bool {
	if s == nil {
		return false
	}
	for _, st := range s.Stops {
		if !st.Color.IsOpaque() {
			return true
		}
	}
	return false
}

// Paint 是一次绘制的全部参数。
//
// 设计上刻意扁平（无接口）：Paint 存在 DisplayList 的 side table 里，
// 值类型可以直接拷贝与比较，避免接口装箱与逃逸。
type Paint struct {
	// Color 是纯色；有 Shader 时作为兜底（Shader 采样失败时用）。
	Color geom.Color
	// Style 决定填充还是描边。
	Style Style
	// StrokeWidth 是描边宽度（DIP）。
	StrokeWidth float32
	// Cap/Join/MiterLimit 是描边细项。
	Cap        Cap
	Join       Join
	MiterLimit float32
	// Dash 是虚线模式（DIP，nil 表示实线）。
	Dash      []float32
	DashPhase float32
	// FillRule 是填充规则。
	FillRule FillRule
	// Shader 为 nil 时用 Color。
	Shader *Shader
	// AntiAlias 关闭后走硬边（覆盖率二值化），用于 A/B 对照。
	AntiAlias bool
}

// Fill 返回一个填充画笔。
func Fill(c geom.Color) Paint {
	return Paint{Color: c, Style: StyleFill, AntiAlias: true, MiterLimit: 4}
}

// FillWith 返回带渐变的填充画笔。
func FillWith(s *Shader) Paint {
	return Paint{Style: StyleFill, Shader: s, AntiAlias: true, MiterLimit: 4, Color: geom.RGBA(0, 0, 0)}
}

// Stroke 返回一个描边画笔。
func Stroke(c geom.Color, width float32) Paint {
	return Paint{
		Color: c, Style: StyleStroke, StrokeWidth: width,
		Cap: CapButt, Join: JoinMiter, MiterLimit: 4, AntiAlias: true,
	}
}

// Dashed 返回虚线描边画笔（dash 单位为 DIP，成对出现：实/空）。
func Dashed(c geom.Color, width float32, dash ...float32) Paint {
	p := Stroke(c, width)
	p.Dash = dash
	return p
}

// WithCap 设置线端。
func (p Paint) WithCap(c Cap) Paint { p.Cap = c; return p }

// WithJoin 设置拐角。
func (p Paint) WithJoin(j Join) Paint { p.Join = j; return p }

// WithMiterLimit 设置尖角上限。
func (p Paint) WithMiterLimit(v float32) Paint { p.MiterLimit = v; return p }

// WithFillRule 设置填充规则。
func (p Paint) WithFillRule(r FillRule) Paint { p.FillRule = r; return p }

// WithoutAA 关闭抗锯齿（A/B 对照用）。
func (p Paint) WithoutAA() Paint { p.AntiAlias = false; return p }

// EffectiveColor 返回不透明快速路径可用的纯色（不可用时 alpha=0）。
//
// 不透明纯色可以跳过逐像素混合，直接写像素——这是矩形背景的主路径。
func (p Paint) EffectiveColor() (geom.Color, bool) {
	if p.Shader != nil {
		return geom.Color{}, false
	}
	if p.Color.IsOpaque() {
		return p.Color, true
	}
	return p.Color, false
}

// IsOpaque 表示该画笔在不透明前提下绘制（纯色不透明）。
func (p Paint) IsOpaque() bool {
	return p.Shader == nil && p.Color.IsOpaque()
}

// StrokeWidthOr 返回有效线宽：未设置时按 1 DIP 处理，避免零宽描边不可见。
func (p Paint) StrokeWidthOr() float32 {
	if p.Style != StyleStroke {
		return 0
	}
	if p.StrokeWidth <= 0 {
		return 1
	}
	return p.StrokeWidth
}

// floorf 是本地实现，避免 render 依赖 math（保持包极轻）。
func floorf(v float32) float32 {
	i := float32(int64(v))
	if v < 0 && i != v {
		i--
	}
	return i
}
