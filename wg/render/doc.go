// Package render 是绘制指令层：Canvas 接口、录制器、不可变 DisplayList、
// Paint/Path/Image 模型。它是"画什么"的唯一表达方式。
//
// 铁律：
//   - 上层（scene/ui/widgets）只能通过 Canvas 表达绘制，不得直接拿 lcl.ICanvas；
//   - DisplayList 一经录制不可变，属性变化必须重录并换 version（契约 C23）；
//   - 每条指令要么被执行，要么显式降级并计数，不得静默丢弃（契约 C22）。
//
// 依赖：只允许 geom / text / theme。不 import lcl、不 import scene。
//
// 状态：M0 只有包声明；Canvas/Recorder/DisplayList 在 M1 落地
// （见 03-ENGINE.md §5–§7）。
package render
