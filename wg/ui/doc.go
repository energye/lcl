// Package ui 是组件基座：Widget/Base/BuildContext、样式系统（Variant/Size/Status/
// StateSet）、动画（Controller/Tween/Curve）、声明式手势包装与 KeyedList。
//
// 只有两层：Widget → Node（不做 Flutter 的 Element 三层）。有状态组件持有
// State 并长期存活，按 Key 复用；列表用 KeyedList 做最小 diff。
//
// 铁律（契约 C17）：本包不得出现 lcl / unsafe / ICanvas / IBitmap，
// 组件只 Build 节点 + 调 SetState，不碰像素。
//
// 依赖：scene / input / overlay / frame / render / text / theme。
//
// 状态：M0 只有包声明；基座与 P1 组件在 M3 落地（见 04-UI-INTERACTION.md Part A）。
package ui
