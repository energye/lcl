// Package theme 是设计 Token 树：颜色、圆角、间距、字体、动效、阴影、层带。
//
// 它只依赖 geom（纯值类型），因此可以被任何层引用而不引入环。
// 组件不写颜色/尺寸字面量，一律从这里取（契约 C9 的延伸：唯一色源）。
//
// 为什么 Token 而不是"主题对象"：
//   - 组件是纯函数 `ButtonStyle(variant, size, status, state, theme)`，
//     可单测、可快照、无隐藏状态；
//   - 换肤只换 Data，不需要重建组件树；
//   - 局部覆盖用 ui.ThemeProvider 写入 Inherited 链，不引全局变量。
package theme

import (
	"time"

	"github.com/energye/lcl/wg/geom"
)

// Band 是层带（堆叠上下文）的整数桶，取值见 02-CANVAS.md §5.2。
//
// 它定义在 theme 而不是 host/overlay：overlay 依赖 host，若 Band 也住在
// host，overlay 与 theme 之间会出现"谁定义谁"的纠缠。host 只需接受
// int32 的层带（`Surface.SetBand(int32)`），因为 Band 的底层就是 int32。
//
// 禁止使用任意浮点 z：桶内按创建序号排序，避免"永远差 0.1"的调试地狱。
type Band int32

const (
	BandBase    Band = 0   // 背景 / 网格 / 画布底
	BandContent Band = 100 // 主内容
	BandSticky  Band = 200 // 吸顶表头 / 悬浮按钮
	BandPopup   Band = 300 // Dropdown / ContextMenu / Select
	BandTooltip Band = 400 // Tooltip / Popover
	BandModal   Band = 500 // 遮罩 + Dialog / Drawer
	BandToast   Band = 600 // Message / Notification
	BandDebug   Band = 900 // 性能叠加层 / damage 可视化
)

// String 返回可读名（诊断输出用）。
func (b Band) String() string {
	switch b {
	case BandBase:
		return "base"
	case BandContent:
		return "content"
	case BandSticky:
		return "sticky"
	case BandPopup:
		return "popup"
	case BandTooltip:
		return "tooltip"
	case BandModal:
		return "modal"
	case BandToast:
		return "toast"
	case BandDebug:
		return "debug"
	default:
		return "band"
	}
}

// Curve 是缓动曲线：输入 0..1，输出 0..1。
type Curve func(t float32) float32

// 内置缓动。都做 0..1 钳制，调用方不必自己 clamp。
func Linear(t float32) float32 { return clamp01(t) }

// EaseOut 是二次缓出（减速），UI 里最常用的默认曲线。
func EaseOut(t float32) float32 {
	t = clamp01(t)
	return 1 - (1-t)*(1-t)
}

// EaseIn 是二次缓入（加速），用于"消失"。
func EaseIn(t float32) float32 {
	t = clamp01(t)
	return t * t
}

// EaseInOut 是三次缓入缓出，用于位移。
func EaseInOut(t float32) float32 {
	t = clamp01(t)
	if t < 0.5 {
		return 4 * t * t * t
	}
	u := -2*t + 2
	return 1 - u*u*u/2
}

func clamp01(t float32) float32 {
	if t < 0 {
		return 0
	}
	if t > 1 {
		return 1
	}
	return t
}

// ColorTokens 是语义色板。
//
// 命名按**用途**而不是色相：组件写 `Color.Surface`，不写 `Color.Gray50`，
// 这样换肤不需要改组件，暗色主题也不用逐个改映射表。
type ColorTokens struct {
	// 主色三态
	Primary       geom.Color
	PrimaryHover  geom.Color
	PrimaryActive geom.Color
	OnPrimary     geom.Color

	// 表面
	Background    geom.Color // 窗口底
	Surface       geom.Color // 卡片/输入框
	SurfaceHover  geom.Color
	SurfaceActive geom.Color
	SurfaceSunken geom.Color // 内凹区（表格斑马、代码块）

	// 描边
	Border       geom.Color
	BorderStrong geom.Color

	// 文字
	Text      geom.Color
	TextMuted geom.Color
	TextFaint geom.Color
	OnDark    geom.Color

	// 语义
	Danger    geom.Color
	OnDanger  geom.Color
	Warning   geom.Color
	Success   geom.Color
	Info      geom.Color
	FocusRing geom.Color
	Selection geom.Color

	// Overlay 是模态遮罩色（带 alpha）
	Overlay geom.Color

	// Shadow 是默认阴影色（带 alpha）
	Shadow geom.Color
}

// RadiusTokens 是圆角。
type RadiusTokens struct {
	SM   float32 // 4
	Base float32 // 6
	LG   float32 // 10
	Pill float32 // 999（用于胶囊）
}

// SpaceTokens 是间距标尺（4 的倍数）。
type SpaceTokens struct {
	XS  float32 // 2
	SM  float32 // 4
	MD  float32 // 8
	LG  float32 // 12
	XL  float32 // 16
	XXL float32 // 24
}

// FontTokens 是字体标尺。
//
// 权重用 int32 而不是 text.Weight：theme 要保持零依赖（只依赖 geom），
// 由 ui/scene 在构造 text.Font 时转换。
type FontTokens struct {
	Family     string
	SizeXS     float32 // 11
	SizeSM     float32 // 12
	SizeBase   float32 // 13
	SizeLG     float32 // 15
	SizeXL     float32 // 18
	SizeTitle  float32 // 22
	LineHeight float32 // 行高倍数，1.5
	WeightReg  int32   // 400
	WeightMed  int32   // 500
	WeightBold int32   // 700
}

// MotionTokens 是动效时长与曲线。
//
// 组件禁止硬编码 `200 * time.Millisecond`：时长必须来自这里，
// 这样"减少动效"设置可以一次性把时长归零而不改组件。
type MotionTokens struct {
	Fast      time.Duration // 120ms
	Mid       time.Duration // 200ms
	Slow      time.Duration // 320ms
	Easing    Curve         // 默认缓动
	EaseIn    Curve
	EaseOut   Curve
	EaseInOut Curve
}

// ShadowSpec 是一层阴影（离屏模糊的参数）。
type ShadowSpec struct {
	Dx, Dy float32
	Blur   float32
	Color  geom.Color
}

// ZTokens 是层带别名，方便写成 `theme.Z.Popup`。
type ZTokens struct {
	Base    Band
	Content Band
	Sticky  Band
	Popup   Band
	Tooltip Band
	Modal   Band
	Toast   Band
	Debug   Band
}

// Data 是一整套主题。
//
// 约定：`Data` 是不可变的——换肤 = 换一个 `*Data`，不是原地改字段。
// 这样才能安全地被多棵树共享，也不需要订阅/失效机制。
type Data struct {
	Name      string
	Dark      bool
	Color     ColorTokens
	Radius    RadiusTokens
	Space     SpaceTokens
	Font      FontTokens
	Motion    MotionTokens
	Elevation []ShadowSpec
	Z         ZTokens
}

func defaultZ() ZTokens {
	return ZTokens{
		Base: BandBase, Content: BandContent, Sticky: BandSticky,
		Popup: BandPopup, Tooltip: BandTooltip, Modal: BandModal,
		Toast: BandToast, Debug: BandDebug,
	}
}

func defaultMotion() MotionTokens {
	return MotionTokens{
		Fast:      120 * time.Millisecond,
		Mid:       200 * time.Millisecond,
		Slow:      320 * time.Millisecond,
		Easing:    EaseOut,
		EaseIn:    EaseIn,
		EaseOut:   EaseOut,
		EaseInOut: EaseInOut,
	}
}

// Light 返回内置亮色主题。
func Light() *Data {
	return &Data{
		Name: "light",
		Color: ColorTokens{
			Primary:       geom.RGBA(59, 130, 246),
			PrimaryHover:  geom.RGBA(37, 99, 235),
			PrimaryActive: geom.RGBA(29, 78, 216),
			OnPrimary:     geom.RGBA(255, 255, 255),

			Background:    geom.RGBA(250, 250, 252),
			Surface:       geom.RGBA(255, 255, 255),
			SurfaceHover:  geom.RGBA(246, 248, 251),
			SurfaceActive: geom.RGBA(238, 242, 247),
			SurfaceSunken: geom.RGBA(244, 246, 249),

			Border:       geom.RGBA(226, 232, 240),
			BorderStrong: geom.RGBA(203, 213, 225),

			Text:      geom.RGBA(15, 23, 42),
			TextMuted: geom.RGBA(100, 116, 139),
			TextFaint: geom.RGBA(148, 163, 184),
			OnDark:    geom.RGBA(248, 250, 252),

			Danger:    geom.RGBA(220, 38, 38),
			OnDanger:  geom.RGBA(255, 255, 255),
			Warning:   geom.RGBA(217, 119, 6),
			Success:   geom.RGBA(22, 163, 74),
			Info:      geom.RGBA(14, 165, 233),
			FocusRing: geom.RGBA(59, 130, 246),
			Selection: geom.RGBA(191, 219, 254),

			Overlay: geom.RGBAa(15, 23, 42, 115),
			Shadow:  geom.RGBAa(15, 23, 42, 26),
		},
		Radius: RadiusTokens{SM: 4, Base: 6, LG: 10, Pill: 999},
		Space:  SpaceTokens{XS: 2, SM: 4, MD: 8, LG: 12, XL: 16, XXL: 24},
		Font: FontTokens{
			Family: "Segoe UI", LineHeight: 1.5,
			SizeXS: 11, SizeSM: 12, SizeBase: 13, SizeLG: 15, SizeXL: 18, SizeTitle: 22,
			WeightReg: 400, WeightMed: 500, WeightBold: 700,
		},
		Motion: defaultMotion(),
		Elevation: []ShadowSpec{
			{Dx: 0, Dy: 1, Blur: 2, Color: geom.RGBAa(15, 23, 42, 15)},
			{Dx: 0, Dy: 2, Blur: 8, Color: geom.RGBAa(15, 23, 42, 26)},
			{Dx: 0, Dy: 8, Blur: 24, Color: geom.RGBAa(15, 23, 42, 36)},
		},
		Z: defaultZ(),
	}
}

// Dark 返回内置暗色主题。语义键与亮色完全一致，只换值。
func Dark() *Data {
	return &Data{
		Name: "dark",
		Dark: true,
		Color: ColorTokens{
			Primary:       geom.RGBA(96, 165, 250),
			PrimaryHover:  geom.RGBA(147, 197, 253),
			PrimaryActive: geom.RGBA(59, 130, 246),
			OnPrimary:     geom.RGBA(15, 23, 42),

			Background:    geom.RGBA(15, 23, 42),
			Surface:       geom.RGBA(30, 41, 59),
			SurfaceHover:  geom.RGBA(38, 51, 71),
			SurfaceActive: geom.RGBA(47, 62, 85),
			SurfaceSunken: geom.RGBA(22, 32, 48),

			Border:       geom.RGBA(51, 65, 85),
			BorderStrong: geom.RGBA(71, 85, 105),

			Text:      geom.RGBA(241, 245, 249),
			TextMuted: geom.RGBA(148, 163, 184),
			TextFaint: geom.RGBA(100, 116, 139),
			OnDark:    geom.RGBA(248, 250, 252),

			Danger:    geom.RGBA(248, 113, 113),
			OnDanger:  geom.RGBA(15, 23, 42),
			Warning:   geom.RGBA(251, 191, 36),
			Success:   geom.RGBA(74, 222, 128),
			Info:      geom.RGBA(56, 189, 248),
			FocusRing: geom.RGBA(96, 165, 250),
			Selection: geom.RGBA(30, 64, 175),

			Overlay: geom.RGBAa(0, 0, 0, 140),
			Shadow:  geom.RGBAa(0, 0, 0, 102),
		},
		Radius: RadiusTokens{SM: 4, Base: 6, LG: 10, Pill: 999},
		Space:  SpaceTokens{XS: 2, SM: 4, MD: 8, LG: 12, XL: 16, XXL: 24},
		Font: FontTokens{
			Family: "Segoe UI", LineHeight: 1.5,
			SizeXS: 11, SizeSM: 12, SizeBase: 13, SizeLG: 15, SizeXL: 18, SizeTitle: 22,
			WeightReg: 400, WeightMed: 500, WeightBold: 700,
		},
		Motion: defaultMotion(),
		Elevation: []ShadowSpec{
			{Dx: 0, Dy: 1, Blur: 2, Color: geom.RGBAa(0, 0, 0, 77)},
			{Dx: 0, Dy: 2, Blur: 8, Color: geom.RGBAa(0, 0, 0, 115)},
			{Dx: 0, Dy: 8, Blur: 24, Color: geom.RGBAa(0, 0, 0, 140)},
		},
		Z: defaultZ(),
	}
}

// Default 返回默认主题（亮色）。
func Default() *Data { return Light() }

// ElevationAt 返回第 i 层阴影；越界时返回最后一层（不 panic，缺省语义明确）。
func (d *Data) ElevationAt(i int) ShadowSpec {
	if d == nil || len(d.Elevation) == 0 {
		return ShadowSpec{}
	}
	if i < 0 {
		i = 0
	}
	if i >= len(d.Elevation) {
		i = len(d.Elevation) - 1
	}
	return d.Elevation[i]
}
