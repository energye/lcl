// Package compat 提供旧 API 的兼容层：把既有 wg 里的 TButton / TTab / TListBox
// 等类型映射到新组件，使已有页面可以逐步迁移（M4）。
//
// 兼容层是**单向**的：只做旧签名 → 新组件的适配，不反向污染新 API。
// 每个映射都要有行为对照（截图对比 + 交互回归），否则不算迁移完成。
//
// 依赖：ui / widgets / scene / geom。
//
// 状态：M0 只有包声明；映射表在 M4 落地（见 05-API-ROADMAP.md §4）。
package compat
