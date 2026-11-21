package widgets

// ── Watermark 水印（M12 长尾）────────────────────────────────────
//
// antd：<Watermark content child gap rotate font>
// Stack 底层 Child + **水印层 CustomPaint**（Save→Rotate→网格
// DrawText→Restore——canvas 原生旋转，antd 默认 -22°）。
// alpha 语义：antd 默认 rgba(0,0,0,0.15)——geom.RGBAa(0,0,0,38)。
// 未实现（如实标注）：多行 content、图片水印（Content 图片）、
// z-index/防篡改（MutationObserver 对应）、zIndex 定制。

import (
	"math"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// Watermark 是水印（平铺旋转文本矩阵盖在 Child 上）。
type Watermark struct {
	// Child 被覆盖内容。
	Child ui.Widget
	// Content 水印文本（单行）。
	Content string
	// GapX/GapY 水印间距（默认 100/80）。
	GapX, GapY float32
	// FontSize 水印字号（默认 14）。
	FontSize float32
	// RotateDeg 旋转角（默认 -22，antd 默认）。
	RotateDeg float32
	// Width/Height 显式尺寸（Stack 布局必需）。
	Width, Height float32
}

// NewWatermark 创建水印。
func NewWatermark(child ui.Widget, content string) *Watermark {
	return &Watermark{Child: child, Content: content, GapX: 100, GapY: 80, FontSize: 14, RotateDeg: -22}
}

// Build 实现 ui.Widget。
func (w *Watermark) Build(ctx *ui.BuildContext) *scene.Node {
	width := w.Width
	if width <= 0 {
		width = 480
	}
	height := w.Height
	if height <= 0 {
		height = 240
	}

	stack := scene.Stack().SetSize2(width, height)
	stack.SetHit(scene.HitSpec{Label: "watermark"})
	if w.Child != nil {
		if child := w.Child.Build(ctx); child != nil {
			child.SetStackPosition(geom.Pt(0, 0)).SetPositionKeepSize(false)
			stack.Add(child)
		}
	}

	// 水印层：旋转平铺文本矩阵（CustomPaint 单节点——O(1) 节点）。
	content := w.Content
	gx, gy := w.GapX, w.GapY
	if gx <= 0 {
		gx = 100
	}
	if gy <= 0 {
		gy = 80
	}
	fs := w.FontSize
	if fs <= 0 {
		fs = 14
	}
	deg := w.RotateDeg
	if deg == 0 {
		deg = -22 // 未设置时 antd 默认
	}
	rad := float32(float64(deg) * math.Pi / 180)
	font := text.DefaultFont(fs)
	layout := text.LayoutText(content, font, text.Options{})
	wc := geom.RGBAa(0, 0, 0, 38) // antd rgba(0,0,0,0.15)

	// 覆盖范围按对角线扩展（旋转后仍盖满）。
	diag := float32(math.Sqrt(float64(width*width + height*height)))
	paint := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		cv.Save()
		cv.Translate(sz.W/2, sz.H/2)
		cv.Rotate(rad)
		half := diag / 2
		for y := -half; y <= half; y += gy {
			for x := -half; x <= half; x += gx {
				cv.DrawText(layout, geom.Pt(x-layout.Size.W/2, y-layout.Size.H/2), wc)
			}
		}
		cv.Restore()
	}).SetSize2(width, height)
	paint.SetPositionKeepSize(true)
	paint.SetStackPosition(geom.Pt(0, 0))
	paint.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "watermark-layer"})
	stack.Add(paint)
	return stack
}
