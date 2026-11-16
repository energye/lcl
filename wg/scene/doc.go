// Package scene 是渲染树与布局：RenderNode、Constraints、measure/arrange。
//
// 它只描述"多大、放哪、谁在上"，**不持有像素**，也不碰 lcl。
// 15 种封闭节点（Box/Text/Path/Image/Flex/Stack/Wrap/Grid/Scroll/Clip/
// Transform/Opacity/Effect/Surface/CustomPaint）在 M2 落地。
//
// 依赖：geom / render / text / theme。
//
// 状态：M0 只有包声明；约束布局与节点在 M2 落地（见 03-ENGINE.md §3–§4）。
package scene
