// Package input 是交互层：原生消息归一化后的输入泵、命中测试、
// capture/target/bubble 路由、GestureArena 裁决、焦点树、文本输入连接与快捷键。
//
// 它与渲染解耦：不持有 Canvas，不碰像素，只操作事件与状态。
// 原生结构体（RawPointer/RawKey/RawWheel/RawIME）定义在 host/raw.go，
// 由本包的 Manager 实现 host.RawSink 后注入——这样 host 不必 import input，
// 避免 host ↔ input 成环。
//
// 依赖：host / scene / frame / render / geom（不含 ui）。
//
// 状态：M0 只有包声明；事件与命中在 M2 落地（见 04-UI-INTERACTION.md Part B）。
package input
