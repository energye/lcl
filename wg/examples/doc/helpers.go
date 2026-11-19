package doc

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// ── 演示编写用的共享助手 ────────────────────────────────────────────
//
// 功能窗是独立的 package main，彼此不能共享代码；这些助手放在 doc 包里
// 供所有窗口复用，避免每个 main.go 各抄一份。

// Box 造一块带文字的实色矩形（演示占位块）。
func Box(ctx *ui.BuildContext, label string, color geom.Color, w, h float32) *scene.Node {
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(color)},
		(&widgets.Label{Text: label, Color: geom.RGBA(255, 255, 255)}).Build(ctx))
	if w > 0 || h > 0 {
		n.SetSize2(w, h)
	}
	return n
}

// Note 产出一行小号说明文字（演示块内的补充说明）。
func Note(ctx *ui.BuildContext, text string) *scene.Node {
	return (&widgets.Label{Text: text, Size: ctx.ThemeOr().Font.SizeSM}).Build(ctx)
}

// nodeWidget 把一段构建函数包装成 ui.Widget（demo 返回值用）。
type nodeWidget struct {
	build func(*ui.BuildContext) *scene.Node
}

func (w nodeWidget) Build(ctx *ui.BuildContext) *scene.Node { return w.build(ctx) }

// Node 把 scene 构建函数包装成 Demo.Build 可返回的 ui.Widget。
func Node(build func(*ui.BuildContext) *scene.Node) ui.Widget {
	return nodeWidget{build: build}
}

// CaptureCtx 在 Build 时记下上下文，供后续点击回调打开浮层。
// 上下文里的 Overlay/Input/Theme 指针由 Container 长期持有，跨帧有效。
func CaptureCtx(slot **ui.BuildContext, ctx *ui.BuildContext) { *slot = ctx }

// CustomPaint 造一个固定尺寸的自绘节点（C17 允许的录制通道）。
func CustomPaint(fn func(cv render.Canvas, sz geom.Size), w, h float32) *scene.Node {
	n := scene.CustomPaintNode(fn).SetSize2(w, h)
	// 被动定位锚（HitDeferToChild 不影响交互）：演示窗的色占比断言用。
	n.SetHit(scene.HitSpec{Label: "custom-canvas"})
	return n
}
