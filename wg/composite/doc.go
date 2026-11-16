// Package composite 负责一帧的合成：Layer 树、合成原因（Reason）、
// 光栅缓存、damage 区域与遮挡剔除、以及最终的一次不透明提交。
//
// 关键约束：每个 Layer 必须声明 Reason（组透明度/模糊/蒙版/缓存/独立滚动），
// 没有 Reason 的层是纯开销——示例里"整窗 alpha 合成 24.7ms"就是这么来的
// （契约 C19）。默认路径下内容直接画进根帧缓冲，不产生任何离屏层。
//
// 依赖：scene / raster / render / geom。
//
// 状态：M0 只有包声明；Layer 树与 damage 在 M1–M2 落地（见 02-CANVAS.md §6）。
package composite
