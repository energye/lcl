package render

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/text"
)

// Canvas 是绘制接口：组件与渲染树只用它表达"画什么"。
//
// 它是一组**命令**而不是像素操作：实现可以是录制器（本包）、软件光栅
// 适配器、GDI 适配器或未来的 GPU 后端。契约 C17 要求上层只能用这个接口，
// 不得直接接触 lcl.ICanvas。
//
// 坐标与变换：
//   - 传入的坐标处于当前用户空间；
//   - Save/Restore 成对使用，管矩阵与裁剪；
//   - 裁剪是"与当前裁剪求交"，不可放大回退。
//
// 所有方法都不返回错误：绘制不能因为一条指令失败而中断整帧——
// 无法执行的指令计入 degraded 并由 --render-matrix 显式报出来（C22）。
type Canvas interface {
	// Save 压栈矩阵与裁剪。
	Save()
	// Restore 出栈；多余调用被忽略（不让一个组件的 bug 破坏整帧）。
	Restore()

	// Translate 平移。
	Translate(dx, dy float32)
	// Scale 缩放。
	Scale(sx, sy float32)
	// Rotate 旋转（弧度）。
	Rotate(rad float32)
	// Concat 左乘变换矩阵。
	Concat(m geom.Matrix)

	// ClipRect 与当前裁剪求交。
	ClipRect(r geom.Rect)
	// ClipPath 按路径裁剪（M1 降级为包围盒并计数）。
	ClipPath(p *Path)

	// FillRect 填充轴对齐矩形（快速路径）。
	FillRect(r geom.Rect, paint Paint)
	// FillRRect 填充圆角矩形。
	FillRRect(r geom.Rect, radius float32, paint Paint)
	// FillOval 填充椭圆。
	FillOval(r geom.Rect, paint Paint)
	// DrawLine 画线（按 paint.Style 填充或描边）。
	DrawLine(x0, y0, x1, y1 float32, paint Paint)
	// DrawArc 画圆弧（内接于 r，角度为弧度）。
	DrawArc(r geom.Rect, start, sweep float32, paint Paint)

	// FillPath 填充路径。
	FillPath(p *Path, paint Paint)
	// StrokePath 描边路径。
	StrokePath(p *Path, paint Paint)
	// DrawPath 按 paint.Style 分派到 FillPath/StrokePath。
	DrawPath(p *Path, paint Paint)

	// DrawText 绘制已排版文本（文本基线由 Layout 内算，origin 为左上角）。
	DrawText(l *text.Layout, origin geom.Point, color geom.Color)
	// DrawImage 绘制位图。
	DrawImage(img *ImageRef, dst geom.Rect, q ImageQuality)
}
