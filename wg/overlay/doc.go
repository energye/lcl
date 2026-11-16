// Package overlay 是浮层宿主：下拉、菜单、Tooltip、Popover、Modal、Toast 的
// 层带（Band）、锚点定位、翻转与关闭策略。
//
// 它**不碰像素**：浮层内容是一棵 scene 子树，交给 Popup 画布绘制，
// 由本包决定位置与 z 序。这样弹窗不再需要新 HWND，也不会脱离父窗口裁剪。
//
// 依赖：host / scene / input / theme。Band 定义在 theme 里（避免 host ↔ overlay 成环）。
//
// 状态：M0 只有包声明；Entry/Show/Close 在 M2 落地（见 04-UI-INTERACTION.md Part A §7）。
package overlay
