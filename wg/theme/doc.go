// Package theme 是设计令牌树：颜色、字体、间距、圆角、阴影、动效时长，
// 以及浮层层带 Band（放在这里而不是 overlay，避免 host ↔ overlay 成环）。
//
// 主题是**不可变数据**：切换主题生成新快照并触发整树重绘，
// 不在原对象上原地改字段（否则缓存与 diff 都会算错）。
//
// 依赖：geom。
//
// 状态：M0 只有包声明；Token 树与内置亮/暗主题在 M3 落地（见 04-UI-INTERACTION.md §2）。
package theme
