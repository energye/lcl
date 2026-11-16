// Package diag 是诊断命令矩阵的实现：--render-matrix、--layout、--hit、
// --gesturetest、--bench、--schedtest、--inputtest、--damage-vis、--layers 等。
//
// 命令的唯一入口是 cmd/wg-diag（main 只做参数解析与退出码，逻辑都放这里，
// 这样每个矩阵都能被 go test 直接调用）。
//
// 依赖：任意（诊断需要观测所有层）。
//
// 状态：M0 只有 --selftest / --headless 两项，其余命令随里程碑补齐
// （见 05-API-ROADMAP.md §6 验收命令矩阵）。
package diag
