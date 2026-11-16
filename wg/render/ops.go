// Package render 是绘制指令层：Canvas 接口、录制器、不可变 DisplayList、
// Paint/Path/Image 模型。它是"画什么"的唯一表达方式。
//
// 铁律：
//   - 上层（scene/ui/widgets）只能通过 Canvas 表达绘制，不得直接拿 lcl.ICanvas；
//   - DisplayList 一经录制不可变，属性变化必须重录并换 version（契约 C23）；
//   - 每条指令要么被执行，要么显式降级并计数，不得静默丢弃（契约 C22）。
//
// 依赖：只允许 geom / text / theme。不 import lcl、不 import scene。
package render

import "github.com/energye/lcl/wg/geom"

// Op 是指令种类。
//
// 指令集刻意做小：所有图元最终都归约成"填充一个 Path"或"描边一个 Path"，
// Rect/RRect/Oval/Arc/Line 只是构造 Path 的语法糖。这样：
//   - 光栅器只需要实现两条路径（fill/stroke），图元扩展不增加后端工作量；
//   - GPU 后端将来只需实现一个 path renderer；
//   - golden 比对不会因为"某图元只在某个后端实现"而失真。
//
// 例外是 Rect：轴对齐矩形走专用快速路径（不建边表，直接算 span 覆盖率），
// 这是唯一值得特殊化的形状——面板、背景、分割线全是它。
type Op uint8

const (
	// OpSave 压栈当前变换与裁剪。
	OpSave Op = iota
	// OpRestore 出栈。
	OpRestore
	// OpConcat 左乘变换矩阵。
	OpConcat
	// OpClipRect 与当前裁剪求交（矩形裁剪可精确执行）。
	OpClipRect
	// OpClipPath 按路径裁剪（M1 降级为包围盒裁剪并计数）。
	OpClipPath
	// OpFillRect 轴对齐矩形填充快速路径。
	OpFillRect
	// OpFillPath 填充路径。
	OpFillPath
	// OpStrokePath 描边路径。
	OpStrokePath
	// OpDrawText 文字回放（设备空间，像素写回后执行）。
	OpDrawText
	// OpDrawImage 绘制位图（M2 落地，M1 计入 degraded）。
	OpDrawImage
)

// String 返回可读名（诊断与日志用）。
func (o Op) String() string {
	switch o {
	case OpSave:
		return "save"
	case OpRestore:
		return "restore"
	case OpConcat:
		return "concat"
	case OpClipRect:
		return "clip-rect"
	case OpClipPath:
		return "clip-path"
	case OpFillRect:
		return "fill-rect"
	case OpFillPath:
		return "fill-path"
	case OpStrokePath:
		return "stroke-path"
	case OpDrawText:
		return "draw-text"
	case OpDrawImage:
		return "draw-image"
	default:
		return "unknown"
	}
}

// Command 是一条指令。
//
// 布局取舍（对应开放问题 Q2：定长 + side table）：
// 指令本身定长，路径/画笔/文字放 side table 用索引引用。定长让 Replay
// 无分支且缓存友好；side table 让 Paint 的切片字段（Dash、渐变 stops）
// 不撑大指令。
type Command struct {
	// Op 是指令种类。
	Op Op
	// Rect 是图元矩形（FillRect/ClipRect 用）或图元包围盒（诊断用）。
	Rect geom.Rect
	// Radius 是圆角半径（RRect 快速路径）。
	Radius float32
	// Start/Sweep 是圆弧起止角（弧度），仅诊断与快速路径用。
	Start, Sweep float32
	// Matrix 是录制该条指令时的**绝对**变换。
	//
	// 关键决策：录制期就把变换累积进指令，光栅期不再维护矩阵栈。
	// 好处是任意单条指令可以脱离上下文重放（命中测试、缓存复用、
	// 局部重绘都依赖这一点）；代价是 Save/Restore 只影响后续指令的录制。
	// 文字因此天然处于设备空间——这正是 C24 要的。
	Matrix geom.Matrix
	// Path 是 side table 索引（-1 表示无）。
	Path int32
	// Paint 是 side table 索引（-1 表示无）。
	Paint int32
	// Text 是文字 side table 索引（-1 表示无）。
	Text int32
	// Image 是图片 side table 索引（-1 表示无）。
	Image int32
	// Quality 是位图采样质量（仅 OpDrawImage 有效）。
	Quality ImageQuality
}

// hasPath 表示该指令引用路径。
func (c Command) hasPath() bool { return c.Path >= 0 }

// hasPaint 表示该指令引用画笔。
func (c Command) hasPaint() bool { return c.Paint >= 0 }
