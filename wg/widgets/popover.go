package widgets

// ── Popover（M9 导航/展示）─────────────────────────────────────────
//
// antd：<Popover title content trigger="click">
// 已对齐：点击触发（Toggle：Child 点击开/再点关+OutsideClick）、
// BandPopup 层带+浅底卡片（**白底浅描边**——与 Tooltip 深底区分，
// antd Popover 是白底）、Title 加粗行+Content Widget（尺寸按内容
// 实测高度自适应）。
// 未实现（如实标注）：箭头指向、受控 open 属性。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// Popover 是点击弹出的富内容卡片（浅底——与 Tooltip 深底区分）。
type Popover struct {
	// Child 触发元素（点击开/再点关）。
	Child ui.Widget
	// Title 卡片标题（加粗行）。
	Title string
	// Content 卡片内容（Widget，尺寸按内容实测高度自适应）。
	Content ui.Widget
	Placement overlay.Placement
	Width     float32

	handle    overlay.Handle
	fb        feedbackHost
	childNode *scene.Node
	open      bool
}

// Build 实现 ui.Widget。
func (p *Popover) Build(ctx *ui.BuildContext) *scene.Node {
	child := buildChild(ctx, p.Child)
	p.childNode = child
	if child == nil {
		return scene.Empty()
	}
	// 合并 Child 已注册 handler（不覆盖触发元素自身点击）。
	h := input.Handler{}
	if ctx.Input != nil {
		if old, ok := ctx.Input.Handler(child); ok {
			h = old
		}
	}
	oldTap := h.Pointer
	h.Pointer = func(ev *input.PointerEvent) bool {
		if ev.Down() {
			if p.open {
				p.hide(ctx)
			} else {
				p.show(ctx, child)
			}
			// 吞掉本次点击：不透传给 Child 自身语义（Popover 点击开）。
			return true
		}
		if oldTap != nil {
			return oldTap(ev)
		}
		return false
	}
	if ctx.Input != nil {
		ctx.Input.Set(child, h)
	}
	return child
}

// show 显示卡片（BandPopup+浅底 Title/Content）。
func (p *Popover) show(ctx *ui.BuildContext, anchor *scene.Node) {
	if ctx.Overlay == nil || p.handle != 0 {
		return
	}
	p.fb = captureHost(ctx)
	th := p.fb.th
	placement := p.Placement
	if placement == 0 {
		placement = overlay.PlacementBottom
	}
	col := scene.Flex(scene.Vertical, th.Space.SM)
	col.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
		Shadow: func() *scene.ShadowSpec {
			if len(th.Elevation) > 1 {
				e := th.Elevation[1]
				return &scene.ShadowSpec{Dx: e.Dx, Dy: e.Dy, Blur: e.Blur, Color: e.Color}
			}
			return nil
		}(),
	})
	col.SetPadding(geom.InLTRB(th.Space.MD, th.Space.SM, th.Space.MD, th.Space.SM))
	w := p.Width
	if w <= 0 {
		w = 260
	}
	col.SetWidth(w)
	if p.Title != "" {
		tl := textLayout(p.Title, th.Font.SizeBase)
		col.Add(scene.Text(scene.TextSpec{Src: p.Title, Layout: tl, Color: th.Color.Text}))
	}
	if p.Content != nil {
		col.Add(buildChild(ctx, p.Content))
	}
	col.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "popover"})
	p.open = true
	p.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: placement,
		Size:      geom.Size{W: w, H: 0}, // 高度自适应
		Dismiss:   overlay.DismissOutsideClick,
		Tag:       "popover",
		OnClose: func() {
			p.handle = 0
			p.open = false
		},
		Build: func() *scene.Node { return col },
	})
}

// hide 隐藏卡片。
func (p *Popover) hide(ctx *ui.BuildContext) {
	if ctx.Overlay == nil || p.handle == 0 {
		return
	}
	ctx.Overlay.Close(p.handle)
}

// Close 立即关闭卡片。
func (p *Popover) Close() { p.hide(nil) }

// IsOpen 返回卡片是否打开。
func (p *Popover) IsOpen() bool { return p.handle != 0 }

