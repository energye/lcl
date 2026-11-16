// Package text 是文本子系统：Font、TextRun、断行、省略、对齐、度量缓存、
// 字形缓存与选区几何。
//
// 它只做排版计算，**不画字**：字形回放交给光栅后端在设备空间完成，
// 以保留 ClearType（契约 C6/C24）。旋转或非整数缩放时降级为灰度 AA。
//
// v1 不做 shaping（HarfBuzz/bidi），先覆盖中英文、emoji、省略与换行；
// 接口预留 Shaper，M5 条件接入。
//
// 依赖：geom。
//
// 状态：M0 只有包声明；Font/TextLayout 在 M1 落地（见 03-ENGINE.md §6）。
package text
