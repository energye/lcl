package scene

import (
	"math"
	"strconv"

	"github.com/energye/lcl/wg/geom"
)

// Inf 是布局里的"无限"约束上界（滚动内容用）。
//
// 用 float32 的 +Inf 而不是一个大数：`Constraints` 的运算里只要出现它，
// 结果仍然是 Inf，不需要额外的哨兵判断；而 `Min <= Max` 的不变式依然成立。
var Inf = float32(math.Inf(1))

// Constraints 是"约束下传 / 尺寸上传"模型里的约束（Flutter RenderBox 同构）。
//
// 不变式：Min <= Max，且 Min >= 0。构造函数保证这一点，
// 之后所有变换（Tighten/Loosen/Deflate）都保持它。
type Constraints struct {
	MinW, MaxW float32
	MinH, MaxH float32
}

// Unbounded 返回无约束（0..Inf）。
func Unbounded() Constraints {
	return Constraints{0, Inf, 0, Inf}
}

// TightConstraints 返回固定尺寸的紧约束。
func TightConstraints(s geom.Size) Constraints {
	return Constraints{s.W, s.W, s.H, s.H}
}

// LooseConstraints 返回 0..size 的松约束。
func LooseConstraints(s geom.Size) Constraints {
	return Constraints{0, maxf(s.W, 0), 0, maxf(s.H, 0)}
}

// NewConstraints 构造并补齐（保证 Min <= Max）。
func NewConstraints(minW, maxW, minH, maxH float32) Constraints {
	if maxW < minW {
		maxW = minW
	}
	if maxH < minH {
		maxH = minH
	}
	if minW < 0 {
		minW = 0
	}
	if minH < 0 {
		minH = 0
	}
	return Constraints{minW, maxW, minH, maxH}
}

// Constrain 把尺寸夹到约束内。
func (c Constraints) Constrain(s geom.Size) geom.Size {
	return geom.Sz(clampf(s.W, c.MinW, c.MaxW), clampf(s.H, c.MinH, c.MaxH))
}

// Tighten 返回宽高都固定的约束（覆盖式）。
func (c Constraints) Tighten(w, h float32) Constraints {
	return TightConstraints(geom.Sz(w, h))
}

// Loosen 把 Min 归零、保留 Max。
func (c Constraints) Loosen() Constraints {
	return Constraints{0, c.MaxW, 0, c.MaxH}
}

// Deflate 按内边距收缩（负值会放宽，但保证 Min 不为负）。
//
// 双向都夹到 0：内边距大于可用空间时退化为 0 而不是负数，
// 这样 `Constrain` 的结果不会出现"负尺寸"这种下游必崩的值。
func (c Constraints) Deflate(ins geom.Insets) Constraints {
	return NewConstraints(
		maxf(0, c.MinW-ins.Horizontal()), maxf(0, c.MaxW-ins.Horizontal()),
		maxf(0, c.MinH-ins.Vertical()), maxf(0, c.MaxH-ins.Vertical()),
	)
}

// MaxSize 返回最大尺寸（Inf 时装成 0，调用方应先问 HasBounded*）。
func (c Constraints) MaxSize() geom.Size { return geom.Sz(c.MaxW, c.MaxH) }

// MinSize 返回最小尺寸。
func (c Constraints) MinSize() geom.Size { return geom.Sz(c.MinW, c.MinH) }

// HasBoundedWidth 宽度是否有上界。
func (c Constraints) HasBoundedWidth() bool { return !isInf(c.MaxW) }

// HasBoundedHeight 高度是否有上界。
func (c Constraints) HasBoundedHeight() bool { return !isInf(c.MaxH) }

// IsTight 宽高是否都已固定。
func (c Constraints) IsTight() bool {
	return c.MinW == c.MaxW && c.MinH == c.MaxH
}

// IsTightWidth 宽度是否已固定。
func (c Constraints) IsTightWidth() bool { return c.MinW == c.MaxW }

// IsTightHeight 高度是否已固定。
func (c Constraints) IsTightHeight() bool { return c.MinH == c.MaxH }

// String 返回可读形式（诊断用）。
func (c Constraints) String() string {
	return "c(" + ftoa(c.MinW) + ".." + ftoa(c.MaxW) + " × " + ftoa(c.MinH) + ".." + ftoa(c.MaxH) + ")"
}

// ftoa 把 Inf 打印成 "inf"，其余按常规。
func ftoa(v float32) string {
	if isInf(v) {
		return "inf"
	}
	return strconv.FormatFloat(float64(v), 'g', 4, 32)
}

// Axis 是主轴方向。
type Axis uint8

const (
	// Horizontal 水平轴。
	Horizontal Axis = iota
	// Vertical 垂直轴。
	Vertical
	// Both 双轴。scene 自身的滚动视口仍是单轴；这个值给组合级滚动组件
	// （widgets.Scroll 的双轴模式）当标记，表示纵向与横向各自滚动。
	Both
)

// String 返回可读名。
func (a Axis) String() string {
	switch a {
	case Horizontal:
		return "horizontal"
	case Vertical:
		return "vertical"
	default:
		return "both"
	}
}

// Flip 返回相反方向。
func (a Axis) Flip() Axis {
	if a == Horizontal {
		return Vertical
	}
	return Horizontal
}

// Alignment 是 [-1,1]² 的对齐位置：(-1,-1)=左上，(0,0)=居中，(1,1)=右下。
//
// 用连续取值而不是枚举：这样 `Align` 组件可以做 0.25 这种任意偏置，
// 而不需要为每个组合新增常量。
type Alignment struct {
	X, Y float32
}

// 常用对齐常量。
var (
	TopLeft      = Alignment{-1, -1}
	TopCenter    = Alignment{0, -1}
	TopRight     = Alignment{1, -1}
	CenterLeft   = Alignment{-1, 0}
	Center       = Alignment{0, 0}
	CenterRight  = Alignment{1, 0}
	BottomLeft   = Alignment{-1, 1}
	BottomCenter = Alignment{0, 1}
	BottomRight  = Alignment{1, 1}
)

// AlignChild 返回把 child（cw×ch）放进容器（pw×ph）时，按对齐计算的偏移。
func (a Alignment) AlignChild(pw, ph, cw, ch float32) geom.Point {
	return geom.Pt(
		(pw-cw)*(a.X+1)/2,
		(ph-ch)*(a.Y+1)/2,
	)
}

// Resolve 把 [-1,1] 的归一化值映射到 [0,1]（用于渐变/进度）。
func (a Alignment) Resolve() (float32, float32) {
	return (a.X + 1) / 2, (a.Y + 1) / 2
}

// Fit 是子节点在容器内的尺寸策略（对应 Flutter 的 BoxFit）。
type Fit uint8

const (
	// FitNone 保持原始尺寸。
	FitNone Fit = iota
	// FitContain 等比缩放到完全可见。
	FitContain
	// FitCover 等比缩放到铺满（可能溢出）。
	FitCover
	// FitFill 非等比拉伸铺满。
	FitFill
)

func clampf(v, lo, hi float32) float32 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

func maxf(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func minf(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func isInf(v float32) bool { return math.IsInf(float64(v), 1) }

// MainAlign 是 Flex 主轴上的排列方式（对应 CSS justify-content）。
type MainAlign uint8

const (
	// MainStart 主起点对齐。
	MainStart MainAlign = iota
	// MainCenter 居中。
	MainCenter
	// MainEnd 主终点对齐。
	MainEnd
	// MainSpaceBetween 两端对齐，中间均分。
	MainSpaceBetween
	// MainSpaceAround 每项两侧均分。
	MainSpaceAround
	// MainSpaceEvenly 所有间隙相等。
	MainSpaceEvenly
)

// String 返回可读名。
func (a MainAlign) String() string {
	switch a {
	case MainStart:
		return "start"
	case MainCenter:
		return "center"
	case MainEnd:
		return "end"
	case MainSpaceBetween:
		return "space-between"
	case MainSpaceAround:
		return "space-around"
	case MainSpaceEvenly:
		return "space-evenly"
	default:
		return "?"
	}
}

// CrossAlign 是 Flex 交叉轴上的对齐方式（对应 CSS align-items）。
type CrossAlign uint8

const (
	// CrossStart 交叉轴起点对齐。
	CrossStart CrossAlign = iota
	// CrossCenter 居中。
	CrossCenter
	// CrossEnd 交叉轴终点对齐。
	CrossEnd
	// CrossStretch 拉伸铺满交叉轴（有固定尺寸的子项不拉伸）。
	CrossStretch
)

// String 返回可读名。
func (c CrossAlign) String() string {
	switch c {
	case CrossStart:
		return "start"
	case CrossCenter:
		return "center"
	case CrossEnd:
		return "end"
	case CrossStretch:
		return "stretch"
	default:
		return "?"
	}
}

// FlexFit 是 flex 子项的填充方式。
type FlexFit uint8

const (
	// FlexTight 严格占满分配到的空间（Expanded）。
	FlexTight FlexFit = iota
	// FlexLoose 最多占分配到的空间（Flexible）。
	FlexLoose
)
