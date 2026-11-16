// Package wg 是 WG Engine 的门面与组合根：它把 host（宿主+画布树）、render
// （录制）、raster（光栅）、composite（合成）、frame（调度）、input（交互）、
// overlay（浮层）、ui/widgets（组件）装配成一个可用的渲染引擎。
//
// # 一句话结构
//
//	Form（1 个 HWND，lcl）
//	└─ ICustomGraphicControl（0 句柄，唯一绘制入口）
//	   └─ Container（唯一父容器，组合根）
//	      └─ 根 Surface（共享帧缓冲 + damage）
//	         ├─ Inline 子画布（面板/滚动视口/列表，共享根像素）
//	         ├─ Layer 子画布（离屏：组透明度/模糊/蒙版/缓存）
//	         └─ Popup 子画布（顶层 z：下拉/菜单/Tooltip/Modal，不新增 HWND）
//
// # 禁止事项（违反即架构违规，checkdeps 会拦）
//
//   - 组件层（ui/widgets）不得出现 lcl / unsafe / ICanvas / IBitmap；
//   - 不许绕过录制直接拿 ICanvas 画像素（CustomPaintNode 也必须走 render.Canvas）；
//   - 不许在光栅期改结构：结构变更进 pending 队列，帧末应用（C18）；
//   - 不许"优化"帧调度：单 outstanding 帧槽是硬约束（C11/C13）。
//
// # 本包是 C17 的唯一例外
//
// 组合根必须在公共签名里出现 lcl.IEngForm（NewContainer(form lcl.IEngForm)），
// 否则调用方要自己写类型断言。因此 C17 的"只有 host/raster 可 import lcl"
// 在根包上放行——但仅限**类型引用**：本包不得调用任何 LCL 绘制/消息 API，
// 所有平台动作都在 host 里完成。
//
// # 文档
//
// 架构与决策见 E:/app/workspace/docs/wg-engine：
// README.md（总览）、01-CHARTER.md（立项）、02-CANVAS.md（画布）、
// 03-ENGINE.md（渲染内核）、04-UI-INTERACTION.md（组件与交互）、
// 05-API-ROADMAP.md（包结构/迁移/里程碑/契约 C1–C28）。
//
// # 当前进度
//
// M0：目录骨架、contracts、checkdeps、backing、ManualClock、Container 空壳、
// host.Surface 树与 damage 上卷、wg-diag --selftest。
// 渲染（render/raster 后端）、布局（scene）、交互（input）在 M1/M2 接入。
package wg
