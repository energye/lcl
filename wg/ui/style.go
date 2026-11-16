package ui

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/theme"
)

// ── 样式系统 ────────────────────────────────────────────────────────
//
// 组件的视觉 = Variant × Size × Status。三维而不是"每个组件一套颜色字段"：
// 否则 Button 的 hover 色、Input 的 hover 色、Select 的 hover 色各写一遍，
// 改一次主题要改十几处，而且必然有不一致的那一个。
//
// 硬规则（§A7 组件目录纪律）：颜色只来自 theme.Data，组件不硬编码色值。

// Variant 是视觉变体。
type Variant uint8

const (
	// VariantSolid 实心（主按钮）。
	VariantSolid Variant = iota
	// VariantOutline 描边（次要按钮）。
	VariantOutline
	// VariantGhost 无底色（文字按钮/图标按钮）。
	VariantGhost
	// VariantSoft 浅底（标签/选中态）。
	VariantSoft
)

// String 返回可读名。
func (v Variant) String() string {
	switch v {
	case VariantSolid:
		return "solid"
	case VariantOutline:
		return "outline"
	case VariantGhost:
		return "ghost"
	case VariantSoft:
		return "soft"
	}
	return "unknown"
}

// Size 是尺寸档。
type Size uint8

const (
	// SizeSmall 小（表格内、工具栏）。
	SizeSmall Size = iota
	// SizeMedium 中（默认）。
	SizeMedium
	// SizeLarge 大（表单主操作）。
	SizeLarge
)

// Status 是交互状态。
type Status uint8

const (
	// StatusNormal 常态。
	StatusNormal Status = iota
	// StatusHover 悬停。
	StatusHover
	// StatusActive 按下/激活。
	StatusActive
	// StatusDisabled 禁用。
	StatusDisabled
	// StatusSelected 选中（列表项、Tab）。
	StatusSelected
)

// String 返回可读名。
func (s Status) String() string {
	switch s {
	case StatusNormal:
		return "normal"
	case StatusHover:
		return "hover"
	case StatusActive:
		return "active"
	case StatusDisabled:
		return "disabled"
	case StatusSelected:
		return "selected"
	}
	return "unknown"
}

// StateSet 是交互状态位集（一个节点可能同时 hover+selected）。
type StateSet uint8

const (
	// StateHover 悬停中。
	StateHover StateSet = 1 << iota
	// StateActive 按下中。
	StateActive
	// StateFocus 持有焦点。
	StateFocus
	// StateDisabled 禁用。
	StateDisabled
	// StateSelected 选中。
	StateSelected
)

// Has 判定是否包含某状态。
func (s StateSet) Has(v StateSet) bool { return s&v != 0 }

// Status 把位集折叠成单值（优先级：禁用 > 按下 > 悬停 > 选中 > 常态）。
//
// 为什么要有优先级：一个禁用的选中项被 hover 时，显示的应该是"禁用"，
// 而不是"悬停"——否则用户会以为可以点。
func (s StateSet) Status() Status {
	switch {
	case s.Has(StateDisabled):
		return StatusDisabled
	case s.Has(StateActive):
		return StatusActive
	case s.Has(StateHover):
		return StatusHover
	case s.Has(StateSelected):
		return StatusSelected
	}
	return StatusNormal
}

// Style 是一次解析后的视觉参数。
type Style struct {
	Background geom.Color
	Foreground geom.Color
	Border     geom.Color
	// BorderWidth 为 0 表示不画边框。
	BorderWidth float32
	// Radius 是圆角。
	Radius float32
	// Padding 是内边距。
	Padding geom.Rect
	// Opacity 是整体不透明度（禁用时降低）。
	Opacity float32
}

// ResolveStyle 按 (Variant, Status) 解析配色。
//
// 颜色全部取自 theme.Data.Color 的语义 token：组件只知道"我要实心/描边"，
// 不知道具体色值，换主题不需要改组件。
func ResolveStyle(t *theme.Data, v Variant, st Status) Style {
	if t == nil {
		t = theme.Default()
	}
	c := t.Color
	out := Style{Opacity: 1, Radius: t.Radius.Base}
	switch v {
	case VariantSolid:
		out.Background, out.Foreground, out.Border = c.Primary, c.OnPrimary, c.Primary
		switch st {
		case StatusHover:
			out.Background = c.PrimaryHover
		case StatusActive:
			out.Background = c.PrimaryActive
		case StatusDisabled:
			out.Background, out.Foreground = c.SurfaceSunken, c.TextFaint
			out.Opacity = 0.6
		case StatusSelected:
			out.Background = c.PrimaryActive
		}
	case VariantOutline:
		out.Background, out.Foreground, out.Border = c.Background, c.Text, c.Border
		out.BorderWidth = 1
		switch st {
		case StatusHover:
			out.Background = c.SurfaceHover
			out.Border = c.BorderStrong
		case StatusActive:
			out.Background = c.SurfaceActive
		case StatusDisabled:
			out.Foreground, out.Border = c.TextFaint, c.Border
			out.Opacity = 0.6
		case StatusSelected:
			out.Border, out.Foreground = c.Primary, c.Primary
		}
	case VariantGhost:
		out.Background, out.Foreground = geom.Color{}, c.Text
		switch st {
		case StatusHover:
			out.Background = c.SurfaceHover
		case StatusActive:
			out.Background = c.SurfaceActive
		case StatusDisabled:
			out.Foreground = c.TextFaint
			out.Opacity = 0.6
		case StatusSelected:
			out.Foreground = c.Primary
		}
	case VariantSoft:
		out.Background, out.Foreground = c.SurfaceSunken, c.Text
		switch st {
		case StatusHover:
			out.Background = c.SurfaceHover
		case StatusActive:
			out.Background = c.SurfaceActive
		case StatusDisabled:
			out.Foreground = c.TextFaint
			out.Opacity = 0.6
		case StatusSelected:
			out.Background, out.Foreground = c.Primary, c.OnPrimary
		}
	}
	return out
}

// Metrics 是尺寸档对应的度量。
type Metrics struct {
	Height  float32
	Padding geom.Rect
	Font    float32
	Gap     float32
}

// MetricsFor 返回尺寸档的度量。
func MetricsFor(t *theme.Data, s Size) Metrics {
	if t == nil {
		t = theme.Default()
	}
	sp := t.Space
	switch s {
	case SizeSmall:
		return Metrics{Height: 24, Padding: geom.RectLTRB(sp.SM, 0, sp.SM, 0), Font: 12, Gap: sp.XS}
	case SizeLarge:
		return Metrics{Height: 40, Padding: geom.RectLTRB(sp.LG, 0, sp.LG, 0), Font: 15, Gap: sp.MD}
	}
	return Metrics{Height: 32, Padding: geom.RectLTRB(sp.MD, 0, sp.MD, 0), Font: 13, Gap: sp.SM}
}

// StyleOf 解析完整视觉参数（配色 + 度量）。
func StyleOf(t *theme.Data, v Variant, s Size, states StateSet) Style {
	if t == nil {
		t = theme.Default()
	}
	out := ResolveStyle(t, v, states.Status())
	m := MetricsFor(t, s)
	out.Padding = m.Padding
	if s == SizeSmall {
		out.Radius = t.Radius.SM
	}
	return out
}
