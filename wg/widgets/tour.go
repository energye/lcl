package widgets

// ── Tour 引导（M12 长尾+深化轮）─────────────────────────────────
//
// antd：<Tour steps open current onClose> 新手引导：
// **target 锚定卡**——遮罩挖洞（四块矩形绕开目标 rect，antd 遮罩
// rgba(0,0,0,0.45)）+ 卡片贴目标下方（空间不足翻上方、横向钳制）；
// TargetNode 由 demo Build 闭包捕获（目标组件 Build 返回的节点直接
// 传递）。无 target 的步骤卡片居中（BandModal AnchorNone+Center）。
// 步骤切换：**time.AfterFunc(0) 重开面板**（TreeSelect 轮教训——
// 重开防御统一）。
// 未实现（如实标注）：自定义 primary 动作、indicatorsText、
// 挖洞区可交互（点击被 Modal 吸收）、平滑滚动跟随。

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// TourStep 是引导步骤。
type TourStep struct {
	Title string
	Desc  string
	// TargetNode 目标锚节点（demo Build 闭包捕获目标组件 Build 返回
	// 的节点；nil=无目标，卡片居中）。
	TargetNode *scene.Node
}

// Tour 是新手引导。
type Tour struct {
	controlBase
	Steps    []TourStep
	OnFinish func() // 最后一步点"完成"
	Width    float32

	Open    bool
	current int

	handle    overlay.Handle
	fb        feedbackHost
	childNode *scene.Node
	panelOpen bool
}

// NewTour 创建引导。
func NewTour(steps []TourStep) *Tour {
	return &Tour{Steps: steps, Width: 300}
}

// Build 实现 ui.Widget：开启引导的触发按钮。
func (t *Tour) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(t)
	btn := &Button{Text: "开始引导", Type: TypePrimary, OnClick: func() {
		t.current = 0
		t.Open = true
		t.showCard(ctx)
	}}
	btnWrap := scene.Flex(scene.Horizontal, 0)
	btnWrap.Add(btn.Build(ctx))
	btnWrap.SetHit(scene.HitSpec{Label: "tour-open"})
	t.childNode = btnWrap
	return btnWrap
}

// Start 从第一步开始展示（锚节点须已 Build）。
func (t *Tour) Start(ctx *ui.BuildContext) {
	t.current = 0
	t.Open = true
	t.showCard(ctx)
}

// IsOpen 返回引导是否展示中。
func (t *Tour) IsOpen() bool { return t.panelOpen }

// Close 关闭引导。
func (t *Tour) Close() {
	if t.fb.host != nil {
		t.closeCard(nil)
	}
	t.Open = false
}

// buildCard 构建当前步卡片节点（锚定/居中两模式复用）。
func (t *Tour) buildCard(ctx *ui.BuildContext, th *theme.Data) *scene.Node {
	w := t.Width
	if w <= 0 {
		w = 300
	}
	step := t.Steps[t.current]
	card := scene.Flex(scene.Vertical, th.Space.SM)
	card.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Radius: th.Radius.Base,
	})
	card.SetPadding(geom.InLTRB(th.Space.LG, th.Space.LG, th.Space.LG, th.Space.LG))
	card.SetWidth(w)

	tl := text.LayoutText(step.Title, text.DefaultFont(th.Font.SizeLG), text.Options{})
	title := scene.Text(scene.TextSpec{Src: step.Title, Layout: tl, Color: th.Color.Text})
	title.SetHit(scene.HitSpec{Label: "tour-title"})
	card.Add(title)
	dl := text.LayoutText(step.Desc, text.DefaultFont(th.Font.SizeBase), text.Options{})
	desc := scene.Text(scene.TextSpec{Src: step.Desc, Layout: dl, Color: th.Color.TextMuted})
	desc.SetHit(scene.HitSpec{Label: "tour-desc"})
	card.Add(desc)

	foot := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	ind := fmt.Sprintf("%d / %d", t.current+1, len(t.Steps))
	il := text.LayoutText(ind, text.DefaultFont(th.Font.SizeSM), text.Options{})
	indText := scene.Text(scene.TextSpec{Src: ind, Layout: il, Color: th.Color.TextMuted})
	indText.SetHit(scene.HitSpec{Label: "tour-indicator"})
	foot.Add(indText)
	foot.Add(scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight))

	nav := func(label string, anchor string, fn func()) {
		b := &Button{Text: label, Size: ui.SizeSmall, OnClick: fn}
		wrap := scene.Flex(scene.Horizontal, 0)
		wrap.Add(b.Build(ctx))
		wrap.SetHit(scene.HitSpec{Label: anchor})
		foot.Add(wrap)
	}
	isLast := t.current == len(t.Steps)-1
	if t.current > 0 {
		nav("上一步", "tour-prev", func() { t.stepTo(ctx, t.current-1) })
	}
	if isLast {
		nav("完成", "tour-finish", func() {
			t.closeCard(ctx)
			t.Open = false
			if t.OnFinish != nil {
				t.OnFinish()
			}
		})
	} else {
		nav("下一步", "tour-next", func() { t.stepTo(ctx, t.current+1) })
	}
	card.Add(foot)
	card.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "tour-card"})
	return card
}

// showCard 弹当前步：有 TargetNode → 遮罩挖洞+卡片贴目标；
// 无 → 居中卡（BandModal AnchorNone+Center）。
func (t *Tour) showCard(ctx *ui.BuildContext) {
	if t.handle != 0 || ctx == nil || ctx.Overlay == nil || !t.Open {
		return
	}
	if t.current < 0 || t.current >= len(t.Steps) {
		return
	}
	th := ctx.ThemeOr()
	t.panelOpen = true
	vp := ctx.ViewportSize()
	step := t.Steps[t.current]
	card := t.buildCard(ctx, th)

	if step.TargetNode == nil {
		t.handle = ctx.Overlay.Show(overlay.Entry{
			Band:      theme.BandModal,
			Anchor:    overlay.Anchor{Kind: overlay.AnchorNone},
			Placement: overlay.PlacementCenter,
			Size:      geom.Size{W: t.Width},
			Modal:     true,
			TrapFocus: true,
			Dismiss:   overlay.DismissEsc,
			Tag:       "tour",
			OnClose: func() {
				t.handle = 0
				t.panelOpen = false
			},
			Build: func() *scene.Node { return card },
		})
		return
	}

	// 锚定模式：全屏 content（Stack）= 遮罩挖洞四块 + 卡片贴目标。
	tr := step.TargetNode.BoundsInRoot()
	cardX := tr.Min.X
	if cardX+t.Width > vp.W {
		cardX = vp.W - t.Width
	}
	if cardX < 0 {
		cardX = 0
	}
	// 卡高实测（LooseConstraints 量一次）用于上/下放置判定。
	card.Layout(scene.LooseConstraints(geom.Sz(t.Width, 10000)))
	ch := card.Size().H
	gap := float32(8)
	cardY := tr.Max.Y + gap
	if cardY+ch > vp.H {
		cardY = tr.Min.Y - gap - ch
		if cardY < 0 {
			cardY = tr.Max.Y + gap
		}
	}

	mask := geom.RGBAa(0, 0, 0, 115) // antd 遮罩 rgba(0,0,0,0.45)
	content := scene.Stack().SetSize2(vp.W, vp.H)
	// 遮罩挖洞：四块矩形绕开目标 rect（CustomPaint 单节点）。
	mp := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		cv.FillRect(geom.RectXYWH(0, 0, sz.W, tr.Min.Y), render.Fill(mask))
		cv.FillRect(geom.RectXYWH(0, tr.Max.Y, sz.W, sz.H-tr.Max.Y), render.Fill(mask))
		cv.FillRect(geom.RectXYWH(0, tr.Min.Y, tr.Min.X, tr.H()), render.Fill(mask))
		cv.FillRect(geom.RectXYWH(tr.Max.X, tr.Min.Y, sz.W-tr.Max.X, tr.H()), render.Fill(mask))
	}).SetSize2(vp.W, vp.H)
	mp.SetPositionKeepSize(true)
	mp.SetStackPosition(geom.Pt(0, 0))
	mp.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "tour-mask"})
	content.Add(mp)

	card.SetPositionKeepSize(true)
	card.SetStackPosition(geom.Pt(cardX, cardY))
	content.Add(card)


	t.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandModal,
		Anchor:    overlay.Anchor{Kind: overlay.AnchorNone},
		Placement: overlay.PlacementCenter,
		Size:      geom.Size{W: vp.W, H: vp.H}, // 全屏 content
		Modal:     true,
		TrapFocus: true,
		Dismiss:   overlay.DismissEsc,
		Tag:       "tour",
		OnClose: func() {
			t.handle = 0
			t.panelOpen = false
		},
		Build: func() *scene.Node { return content },
	})
}

// stepTo 切步：AfterFunc(0) 重开面板（重开防御统一）。
func (t *Tour) stepTo(ctx *ui.BuildContext, idx int) {
	t.current = idx
	if ctx != nil && ctx.Overlay != nil && t.handle != 0 {
		ctx.Overlay.Close(t.handle)
		t.handle = 0
	}
	time.AfterFunc(0, func() { t.showCard(ctx) })
}

// closeCard 关闭卡片。
func (t *Tour) closeCard(ctx *ui.BuildContext) {
	if ctx != nil && ctx.Overlay != nil && t.handle != 0 {
		ctx.Overlay.Close(t.handle)
	}
	t.handle = 0
	t.panelOpen = false
}
