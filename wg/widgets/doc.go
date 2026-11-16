// Package widgets 是组件库：basic / layout / control / nav / data / feedback /
// form / chart 八个分组。
//
// 每个组件只组合 ui 的基座与 scene 的节点，不写自己的布局算法、不碰像素、
// 不做自己的命中测试——命中与手势由 input 统一处理。
//
// 依赖：ui / input / theme（以及 scene/geom/render 的类型）。
//
// 状态：M0 只有包声明；P1 组件（12 个）在 M3、P2 与图表在 M4 落地
// （见 04-UI-INTERACTION.md §5 组件清单）。
package widgets
