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

	// 预览变换状态（antd TransformType 的 scale/rotate 子集）。
	previewScale  float32          // 1 = 原始
	previewRotate int              // 90° 步数（0..3）
	bodyNode      *scene.Node      // 预览图像节点（变换后原地重绘）
	showCtx       *ui.BuildContext // 预览期 ctx（程序化 Close 用）
	ref           *render.ImageRef // 位图解码缓存（避免每次打开预览重解码）
	previewW      float32
	previewH      float32
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
		l := text.LayoutText(txt, text.DefaultFont(th.Font.SizeBase), text.Options{})
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
	im.showCtx = ctx
	im.previewScale = 1
	im.previewRotate = 0
	pw := im.Width * 2.5
	if pw > 640 {
		pw = 640
	}
	ph := pw * im.Height / im.Width
	im.previewW, im.previewH = pw, ph

	if im.ref == nil {
		im.ref = render.ImageFrom(im.Src)
	}
	ref := im.ref
	// 预览图像按当前变换绘制（缩放+旋转，以画布中心为原点）。
	body := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		sc := im.previewScale
		if sc <= 0 {
			sc = 1
		}
		cv.Save()
		cv.Translate(sz.W/2, sz.H/2)
		if im.previewRotate != 0 {
			cv.Rotate(float32(im.previewRotate) * 1.5707963) // 90° × n
		}
		cv.Scale(sc, sc)
		cv.Translate(-sz.W/2, -sz.H/2)
		cv.DrawImage(ref, geom.RectXYWH(0, 0, sz.W, sz.H), render.ImageBilinear)
		cv.Restore()
	}).SetSize2(pw, ph)
	im.bodyNode = body

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

	// 底部工具条（antd preview toolbar）：放大/缩小/左转/右转/复位。
	// 按钮为半透明白色图标，居中一排。
	tools := []struct {
		icon  string
		label string
		act   func()
	}{
		{"zoom-out", "image-zoom-out", func() { im.applyZoom(1 / 1.5) }},
		{"zoom-in", "image-zoom-in", func() { im.applyZoom(1.5) }},
		{"undo", "image-rotate-left", func() { im.applyRotate(-1) }},
		{"reload", "image-rotate-right", func() { im.applyRotate(1) }},
	}
	const toolSize = float32(28)
	const gap = float32(8)
	toolW := float32(len(tools))*toolSize + float32(len(tools)-1)*gap
	tx := (pw - toolW) / 2
	ty := ph - 40
	for _, t := range tools {
		ic := (&Icon{Name: t.icon, Size: 16, Color: geom.RGBA(255, 255, 255)}).Build(ctx)
		ic.SetPositionKeepSize(true)
		ic.SetStackPosition(geom.Pt(tx, ty))
		ic.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: t.label, Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			act := t.act
			ctx.Input.Set(ic, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				if !ev.Down() {
					return false
				}
				act()
				return true
			}})
		}
		box.Add(ic)
		tx += toolSize + gap
	}

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

// applyZoom 按倍数缩放预览（钳制 [1, 50]），原地重绘图像。
func (im *Image) applyZoom(factor float32) {
	if im.previewScale <= 0 {
		im.previewScale = 1
	}
	s := im.previewScale * factor
	if s < 1 {
		s = 1
	}
	if s > 50 {
		s = 50
	}
	im.previewScale = s
	im.repaintPreview()
}

// applyRotate 旋转预览（dir=±1，90° 步进，0..3 循环）。
func (im *Image) applyRotate(dir int) {
	im.previewRotate = (im.previewRotate + dir + 4) % 4
	im.repaintPreview()
}

// repaintPreview 原地重绘预览图像节点（变换状态变化后调用）。
func (im *Image) repaintPreview() {
	if im.bodyNode != nil {
		im.bodyNode.MarkNeedsPaint()
	}
}

// Show 程序化打开预览（锚=最近构建的图像节点；测试/自动化用）。
func (im *Image) Show(ctx *ui.BuildContext) {
	if im.childNode == nil {
		return
	}
	im.openPreview(ctx, im.childNode, ctx.ThemeOr())
}

// IsOpen 预览是否打开。
func (im *Image) IsOpen() bool { return im.panelOpen }

// Close 程序化关闭预览。
func (im *Image) Close() {
	im.closePreview(im.showCtx)
}

// closePreview 关闭预览。
func (im *Image) closePreview(ctx *ui.BuildContext) {
	if im.handle != 0 && ctx != nil && ctx.Overlay != nil {
		ctx.Overlay.Close(im.handle)
	}
	im.handle = 0
	im.panelOpen = false
}
