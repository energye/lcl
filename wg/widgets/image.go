package widgets

// ── Image 图片（M12 长尾）────────────────────────────────────────
//
// antd：<Image width height src preview fallback>
// 引擎无网络/文件 IO：Src 为内存位图（render.ImageRef，宿主注入
// ——文件/网络由平台层读入后传入）；Src 为 nil 时渲染**加载失败
// 占位**（灰底+文本，antd fallback 形态）。Preview：点击原图弹
// **BandModal 预览**（遮罩+居中大图+右上关闭+Esc——Modal 组件同
// overlay 模式）。
// 未实现（如实标注）：加载中状态、懒加载、旋转/缩放操作、
// preview 多图切换、宽高自适应容器模式。

import (
	"image"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// Image 是图片（显示+点击预览）。
type Image struct {
	// Src 内存位图（nil=加载失败占位）。
	controlBase
	Src image.Image
	// Width/Height 显示尺寸。
	Width, Height float32
	// Radius 圆角（默认 4）。
	Radius float32
	// Preview 点击原图放大预览（默认 true）。
	Preview      bool
	FallbackText string // Src 为 nil 时的占位文本（默认"图片加载失败"）

	handle    overlay.Handle
	fb        feedbackHost
	childNode *scene.Node
	panelOpen bool
}

// NewImage 创建图片组件。
func NewImage(src image.Image, w, h float32) *Image {
	return &Image{Src: src, Width: w, Height: h, Radius: 4, Preview: true}
}

// Build 实现 ui.Widget。
func (im *Image) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(im)
	th := ctx.ThemeOr()
	w := im.Width
	h := im.Height
	if w <= 0 {
		w = 200
	}
	if h <= 0 {
		h = 150
	}

	var node *scene.Node
	if im.Src != nil {
		ref := render.ImageFrom(im.Src)
		radius := im.Radius
		if radius == 0 {
			radius = 4
		}
		node = scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
			cv.FillRRect(geom.RectXYWH(0, 0, sz.W, sz.H), radius, render.Fill(th.Color.SurfaceSunken))
			cv.DrawImage(ref, geom.RectXYWH(0, 0, sz.W, sz.H), render.ImageBilinear)
		}).SetSize2(w, h)
	} else {
		// 加载失败占位（antd fallback：灰底+提示文本）。
		node = scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.SurfaceSunken),
			Border: scene.Stroke(th.Color.Border, 1),
			Radius: 4,
		}).SetSize2(w, h)
		txt := im.FallbackText
		if txt == "" {
			txt = "图片加载失败"
		}
		l := text.LayoutText(txt, text.DefaultFont(th.Font.SizeSM), text.Options{})
		tv := scene.Text(scene.TextSpec{Src: txt, Layout: l, Color: th.Color.TextFaint})
		tv.SetMargin(geom.InLTRB(0, (h-l.Size.H)/2, 0, 0))
		tv.SetPositionKeepSize(true)
		tv.SetStackPosition(geom.Pt(w/2-l.Size.W/2, 0))
		node.Add(tv)
	}
	node.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "image", Cursor: scene.CursorPointer})

	if ctx.Input != nil && im.Preview {
		anchor := node
		g := &ui.GestureDetector{OnTap: func(geom.Point) {
			if im.panelOpen {
				im.closePreview(ctx)
				return
			}
			im.openPreview(ctx, anchor, th)
		}}
		ctx.Input.Set(node, g.Handler())
	}
	im.childNode = node
	return node
}

// openPreview 弹 BandModal 预览（遮罩+居中大图+右上关闭+Esc）。
func (im *Image) openPreview(ctx *ui.BuildContext, anchor *scene.Node, th *theme.Data) {
	if im.handle != 0 || im.Src == nil || ctx == nil || ctx.Overlay == nil {
		return
	}
	th2 := th
	if th2 == nil {
		th2 = ctx.ThemeOr()
	}
	im.panelOpen = true
	pw := im.Width * 2.5
	if pw > 640 {
		pw = 640
	}
	ph := pw * im.Height / im.Width

	ref := render.ImageFrom(im.Src)
	body := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		cv.DrawImage(ref, geom.RectXYWH(0, 0, sz.W, sz.H), render.ImageBilinear)
	}).SetSize2(pw, ph)

	box := scene.Stack()
	box.SetSize2(pw, ph)
	box.Add(body)
	// 右上关闭钮。
	closeIcon := (&Icon{Name: "close", Size: 18, Color: geom.RGBA(255, 255, 255)}).Build(ctx)
	closeIcon.SetPositionKeepSize(true)
	closeIcon.SetStackPosition(geom.Pt(pw-30, 10))
	closeIcon.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "image-close", Cursor: scene.CursorPointer})
	box.Add(closeIcon)
	if ctx.Input != nil {
		ctx.Input.Set(closeIcon, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
			if !ev.Down() {
				return false
			}
			im.closePreview(ctx)
			return true
		}})
	}
	box.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "image-preview"})

	im.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandModal,
		Anchor:    overlay.Anchor{Kind: overlay.AnchorNone},
		Placement: overlay.PlacementCenter,
		Size:      geom.Size{W: pw, H: ph},
		Modal:     true,
		TrapFocus: true,
		Dismiss:   overlay.DismissEsc,
		Tag:       "image-preview",
		OnClose: func() {
			im.handle = 0
			im.panelOpen = false
		},
		Build: func() *scene.Node { return box },
	})
}

// closePreview 关闭预览。
func (im *Image) closePreview(ctx *ui.BuildContext) {
	if im.handle != 0 && ctx != nil && ctx.Overlay != nil {
		ctx.Overlay.Close(im.handle)
	}
	im.handle = 0
	im.panelOpen = false
}
