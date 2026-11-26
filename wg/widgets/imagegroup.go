package widgets

// ── ImageGroup（antd Image.PreviewGroup）────────────────────────────
//
// 多张图并排缩略图，点击任一张打开预览；预览内左右切换按钮循环切图，
// 底部显示 current/total 计数，切换回调 OnChange。
//
// 组合复用：缩略图用 Image 组件（各自可预览，但组内点击统一走本组预览）；
// 预览体沿用 Image 的 CustomPaint 位图绘制。
//
// 未实现（如实标注）：预览内缩放/旋转工具条（Image 单体已有，组内未接）、
// countRender 自定义计数、items 的 crossOrigin 等非画布属性。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ImageGroup 是多图预览组（antd Image.PreviewGroup）。
type ImageGroup struct {
	controlBase
	// Thumbs 缩略图（每项是已配置好的 Image；点击统一打开本组预览）。
	Thumbs []*Image
	// Gap 缩略图间距（默认 8）。
	Gap float32
	// OnChange 切图回调（antd preview.onChange：(current, prevCurrent)）。
	OnChange func(current, prev int)

	handle    overlay.Handle
	fb        feedbackHost
	showCtx   *ui.BuildContext
	current   int // 预览当前索引
	panelOpen bool
	bodyNode  *scene.Node
	countNode *scene.Node
	refs      []*render.ImageRef // 解码缓存（与 Thumbs 一一对应）
}

// NewImageGroup 创建多图预览组。
func NewImageGroup(thumbs ...*Image) *ImageGroup {
	return &ImageGroup{Thumbs: thumbs, Gap: 8}
}

// Current 返回预览当前索引（诊断/断言用）。
func (g *ImageGroup) Current() int { return g.current }

// Build 实现 ui.Widget：并排缩略图。
func (g *ImageGroup) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(g)
	th := ctx.ThemeOr()
	gap := g.Gap
	if gap <= 0 {
		gap = th.Space.MD
	}
	row := scene.Flex(scene.Horizontal, gap).SetCrossAlign(scene.CrossCenter)
	for i, t := range g.Thumbs {
		if t == nil {
			continue
		}
		// 缩略图强制关闭自身预览（点击统一走组预览）。
		t.Preview = false
		n := t.Build(ctx)
		if n == nil {
			continue
		}
		n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "image-group-thumb", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			idx := i
			ctx.Input.Set(n, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				if !ev.Down() {
					return false
				}
				g.open(ctx, n, th, idx)
				return true
			}})
		}
		row.Add(n)
	}
	row.SetHit(scene.HitSpec{Label: "image-group"})
	return row
}

// open 打开预览并定位到 idx。
func (g *ImageGroup) open(ctx *ui.BuildContext, anchor *scene.Node, th *theme.Data, idx int) {
	if g.handle != 0 || ctx == nil || ctx.Overlay == nil || len(g.Thumbs) == 0 {
		return
	}
	g.showCtx = ctx
	g.panelOpen = true
	g.current = idx
	// 预览尺寸按当前图（封顶 640 宽）。
	t := g.Thumbs[g.current]
	pw := t.Width * 2.5
	if pw > 640 {
		pw = 640
	}
	if pw <= 0 {
		pw = 400
	}
	ph := pw * t.Height / t.Width
	if t.Width <= 0 {
		ph = pw * 0.75
	}

	box := g.buildPreview(ctx, th, pw, ph)
	g.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandModal,
		Anchor:    overlay.Anchor{Kind: overlay.AnchorNone},
		Placement: overlay.PlacementCenter,
		Size:      geom.Size{W: pw, H: ph},
		Modal:     true,
		TrapFocus: true,
		Dismiss:   overlay.DismissEsc,
		Tag:       "image-group-preview",
		OnClose: func() {
			g.handle = 0
			g.panelOpen = false
		},
		Build: func() *scene.Node { return box },
	})
}

// buildPreview 构建预览体（当前图 + 左右切换 + 计数 + 关闭）。
func (g *ImageGroup) buildPreview(ctx *ui.BuildContext, th *theme.Data, pw, ph float32) *scene.Node {
	box := scene.Stack().SetSize2(pw, ph)
	box.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "image-group-preview"})

	// 当前图（CustomPaint 位图）：闭包内按 g.current 取 ref——切图后
	// MarkNeedsPaint 重绘即生效（Image.applyZoom 同款原地刷新）。
	body := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		if r := g.refOf(g.current); r != nil {
			cv.DrawImage(r, geom.RectXYWH(0, 0, sz.W, sz.H), render.ImageBilinear)
		}
	}).SetSize2(pw, ph)
	g.bodyNode = body
	box.Add(body)

	// 右上关闭钮。
	closeIcon := (&Icon{Name: "close", Size: 18, Color: geom.RGBA(255, 255, 255)}).Build(ctx)
	closeIcon.SetPositionKeepSize(true)
	closeIcon.SetStackPosition(geom.Pt(pw-30, 10))
	closeIcon.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "image-close", Cursor: scene.CursorPointer})
	if ctx.Input != nil {
		ctx.Input.Set(closeIcon, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
			if !ev.Down() {
				return false
			}
			g.close(ctx)
			return true
		}})
	}
	box.Add(closeIcon)

	// 左右切换按钮（antd preview prev/next）。
	mkArrow := func(icon, label string, delta int) *scene.Node {
		ic := (&Icon{Name: icon, Size: 20, Color: geom.RGBA(255, 255, 255)}).Build(ctx)
		ic.SetPositionKeepSize(true)
		ic.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: label, Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			ctx.Input.Set(ic, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				if !ev.Down() {
					return false
				}
				g.step(ctx, delta)
				return true
			}})
		}
		return ic
	}
	prev := mkArrow("left", "image-group-prev", -1)
	prev.SetStackPosition(geom.Pt(12, ph/2-10))
	box.Add(prev)
	next := mkArrow("right", "image-group-next", 1)
	next.SetStackPosition(geom.Pt(pw-32, ph/2-10))
	box.Add(next)

	// 底部计数 current/total。
	cnt := text.LayoutText(g.countText(), text.DefaultFont(th.Font.SizeBase), text.Options{})
	ct := scene.Text(scene.TextSpec{Src: g.countText(), Layout: cnt, Color: geom.RGBA(255, 255, 255)})
	ct.SetPositionKeepSize(true)
	ct.SetStackPosition(geom.Pt((pw-cnt.Size.W)/2, ph-26))
	ct.SetHit(scene.HitSpec{Label: "image-group-count"})
	g.countNode = ct
	box.Add(ct)
	return box
}

// countText 计数文案 "current/total"（1-based）。
func (g *ImageGroup) countText() string {
	return itoa(g.current+1) + "/" + itoa(len(g.Thumbs))
}

// refOf 第 i 张的解码缓存（懒解码）。
func (g *ImageGroup) refOf(i int) *render.ImageRef {
	if i < 0 || i >= len(g.Thumbs) || g.Thumbs[i] == nil || g.Thumbs[i].Src == nil {
		return nil
	}
	if len(g.refs) != len(g.Thumbs) {
		g.refs = make([]*render.ImageRef, len(g.Thumbs))
	}
	if g.refs[i] == nil {
		g.refs[i] = render.ImageFrom(g.Thumbs[i].Src)
	}
	return g.refs[i]
}

// step 循环切图（delta=±1），原地重绘预览体。
func (g *ImageGroup) step(ctx *ui.BuildContext, delta int) {
	if len(g.Thumbs) == 0 {
		return
	}
	prev := g.current
	g.current = (g.current + delta + len(g.Thumbs)) % len(g.Thumbs)
	g.repaint(ctx)
	if g.OnChange != nil {
		g.OnChange(g.current, prev)
	}
}

// repaint 原地换当前图：位图节点重绘（闭包读 g.current）+ 计数文字
// 重设 + 整页重建刷触发框（MarkDirty 让根重建同步计数显示）。
func (g *ImageGroup) repaint(ctx *ui.BuildContext) {
	if g.bodyNode != nil {
		g.bodyNode.MarkNeedsPaint()
	}
	if g.countNode != nil {
		txt := g.countText()
		lay := text.LayoutText(txt, text.DefaultFont(14), text.Options{})
		g.countNode.SetText(scene.TextSpec{Src: txt, Layout: lay, Color: geom.RGBA(255, 255, 255)})
		g.countNode.MarkNeedsPaint()
	}
}

// close 关闭预览。
func (g *ImageGroup) close(ctx *ui.BuildContext) {
	if g.handle != 0 && ctx != nil && ctx.Overlay != nil {
		ctx.Overlay.Close(g.handle)
	}
	g.handle = 0
	g.panelOpen = false
}

// Close 程序化关闭。
func (g *ImageGroup) Close() { g.close(g.showCtx) }

// IsOpen 预览是否打开。
func (g *ImageGroup) IsOpen() bool { return g.panelOpen }

// itoa 小整数转字符串（避免引入 strconv 冲突；仅 0..999）。
func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	var buf [4]byte
	i := len(buf)
	neg := v < 0
	if neg {
		v = -v
	}
	for v > 0 {
		i--
		buf[i] = byte('0' + v%10)
		v /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
