package widgets

// ── Tour 引导（M12 长尾）────────────────────────────────────────
//
// antd：<Tour steps open current onClose> 新手引导：
// Open 时浮层显示当前步卡片（标题+描述+步骤指示+上一步/下一步/关闭）。
// 第一轮**居中卡片**（AnchorNone+PlacementCenter，Modal 遮罩）——
// target 锚定卡（贴目标元素+遮罩镂空）未实现（如实标注）。
// 步骤切换：**time.AfterFunc(0) 重开面板**（TreeSelect 轮教训——
// down 分发中途重开丢内容；Tour 按钮为 GestureDetector tap，同款
// 防御）。

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// TourStep 是引导步骤。
type TourStep struct {
	Title string
	Desc  string
}

// Tour 是新手引导。
type Tour struct {
	controlBase
	Steps    []TourStep
	OnFinish func() // 最后一步点"完成"或任意时刻关闭
	Width    float32

	Open    bool // 是否展示中（程序化/按钮开启）
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

// Build 实现 ui.Widget：开启引导的触发按钮（antd 由 open 受控——
// 这里按钮是演示入口；程序化 Open() 亦可）。
func (t *Tour) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(t)
	th := ctx.ThemeOr()
	w := t.Width
	if w <= 0 {
		w = 300
	}
	btn := &Button{Text: "开始引导", Type: TypePrimary, OnClick: func() {
		t.current = 0
		t.Open = true
		t.showCard(ctx, th)
	}}
	btnWrap := scene.Flex(scene.Horizontal, 0)
	btnWrap.Add(btn.Build(ctx))
	btnWrap.SetHit(scene.HitSpec{Label: "tour-open"})
	t.childNode = btnWrap
	return btnWrap
}

// Open 方法名与字段冲突——展示入口用 Start。
// Start 从第一步开始展示。
func (t *Tour) Start(ctx *ui.BuildContext) {
	t.current = 0
	t.Open = true
	if t.childNode != nil {
		t.showCard(ctx, ctx.ThemeOr())
	}
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

// showCard 弹当前步卡片（BandModal 居中+遮罩）。
func (t *Tour) showCard(ctx *ui.BuildContext, th *theme.Data) {
	if t.handle != 0 || ctx == nil || ctx.Overlay == nil || t.Open == false {
		return
	}
	t.panelOpen = true
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
	dl := text.LayoutText(step.Desc, text.DefaultFont(th.Font.SizeSM), text.Options{})
	desc := scene.Text(scene.TextSpec{Src: step.Desc, Layout: dl, Color: th.Color.TextMuted})
	desc.SetHit(scene.HitSpec{Label: "tour-desc"})
	card.Add(desc)

	// 指示行：步数 + 上一步/下一步（最后一步为完成）。
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
		nav("上一步", "tour-prev", func() { t.stepTo(ctx, th, t.current-1) })
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
		nav("下一步", "tour-next", func() { t.stepTo(ctx, th, t.current+1) })
	}
	card.Add(foot)
	card.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "tour-card"})

	t.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandModal,
		Anchor:    overlay.Anchor{Kind: overlay.AnchorNone},
		Placement: overlay.PlacementCenter,
		Size:      geom.Size{W: w},
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
}

// stepTo 切步：AfterFunc(0) 重开面板（TreeSelect 轮教训——重开必经
// 延迟；Tour 按钮为 tap 非浮层 down，仍统一防御）。
func (t *Tour) stepTo(ctx *ui.BuildContext, th *theme.Data, idx int) {
	t.current = idx
	if ctx != nil && ctx.Overlay != nil && t.handle != 0 {
		ctx.Overlay.Close(t.handle)
		t.handle = 0
	}
	time.AfterFunc(0, func() { t.showCard(ctx, th) })
}

// closeCard 关闭卡片。
func (t *Tour) closeCard(ctx *ui.BuildContext) {
	if ctx != nil && ctx.Overlay != nil && t.handle != 0 {
		ctx.Overlay.Close(t.handle)
	}
	t.handle = 0
	t.panelOpen = false
}
