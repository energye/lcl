package wg

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/theme"
)

// Version 是引擎版本。里程碑后缀标识能力边界。
const Version = "0.1.0-m4"

// Stage 是当前里程碑名（与 contracts.Stage 对应，这里用字符串避免上层依赖）。
const Stage = "M4"

// Options 是容器选项。
//
// 取值约定：**零值 Options{} 表示"全部默认"**（AA 开启、缓存 64MiB、层数上限 64）。
// 若要关闭某项默认能力（例如 A/B 对比时关 AA），请基于 DefaultOptions() 改：
//
//	opts := wg.DefaultOptions()
//	opts.AA = false
//	c := wg.NewContainer(form, opts)
//
// 直接写 Options{AA: false} 与 Options{} 等价，仍然会得到默认开启的 AA——
// 这是为了避免"忘了填选项就关掉了抗锯齿"这类静默降级。
type Options struct {
	// Title 仅用于诊断输出（真实窗口标题仍由窗体决定）。
	Title string

	// AA 开启抗锯齿。零值 Options 视为 true；显式关闭请用 DefaultOptions()。
	AA bool
	// DebugDamage 把本帧脏区描成红色，用于确认"只重绘该重绘的地方"。
	DebugDamage bool
	// DebugLayers 在每个合成层的边界画 1px 边框，用于确认层数与嵌套。
	DebugLayers bool

	// Scale 覆盖 DPI 缩放（0 = 跟随窗体 DPI）。
	//
	// 只在需要固定出图的场景使用（golden 测试、截图回归）。
	Scale float32

	// CacheBytes 是光栅缓存预算上限（0 = 默认 64MiB）。
	CacheBytes int64
	// MaxLayers 是单帧合成层数上限（0 = 默认 64），超出时按 Reason 优先级降级。
	MaxLayers int

	// FrameInterval 是目标帧间隔（0 = 跟随显示器刷新率）。
	FrameInterval time.Duration

	// Background 是清屏色（零值 = 不透明白）。
	//
	// 清屏只发生在脏区，不是整屏：这是"局部重绘"能真正省下像素的前提。
	Background geom.Color

	// Theme 是初始主题（零值 = theme.Default()）。
	// 运行期换主题用 Container.SetTheme。
	Theme *theme.Data
}

// DefaultOptions 返回推荐默认值。
func DefaultOptions() Options {
	return Options{
		AA:         true,
		CacheBytes: 64 << 20,
		MaxLayers:  64,
	}
}

// normalized 补齐零值，返回可直接使用的选项。
func (o Options) normalized() Options {
	if !o.AA && o == (Options{}) {
		o = DefaultOptions()
	}
	if o.CacheBytes <= 0 {
		o.CacheBytes = 64 << 20
	}
	if o.MaxLayers <= 0 {
		o.MaxLayers = 64
	}
	if o.FrameInterval < 0 {
		o.FrameInterval = 0
	}
	if o.Background == (geom.Color{}) {
		o.Background = geom.RGBA(255, 255, 255)
	}
	return o
}
