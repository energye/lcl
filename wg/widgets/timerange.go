package widgets

// ── TimeRangePicker（antd TimePicker.RangePicker）────────────────────
//
// 双触发框（起 ~ 止）+ 点开浮层内并排两个时间面板 + 共享底部"确定"。
// 组合复用两个 TimePicker（各自持有列状态与值）：起端面板 pick 后自动
// 切到止端；点确定提交区间（起 > 止自动交换）。两端时间各自独立。
//
// 未实现（如实标注）：disabledTime 的 range 形态（RangeDisabledTime 带
// start/end 类型参数）、order=false 不自动交换、12h 下两端不同 AP 列
// 的对齐细节。

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// TimeRangePicker 是时间区间选择器。
type TimeRangePicker struct {
	controlBase
	Start         TimePicker
	End           TimePicker
	Width         float32
	OnChangeRange func(start, end time.Time)

	handle    overlay.Handle
	fb        feedbackHost
	childNode *scene.Node
	panelOpen bool
	activeEnd int // 0=起 1=止
	curTh     *theme.Data
}

// NewTimeRangePicker 创建时间区间选择器（默认无值）。
func NewTimeRangePicker() *TimeRangePicker {
	r := &TimeRangePicker{}
	r.Start.Placeholder = "开始时间"
	r.End.Placeholder = "结束时间"
	// 起端 pick 后切到止端。
	r.Start.OnValueChange = func() {
		r.activeEnd = 1
	}
	return r
}

// IsOpen 面板是否打开。
func (r *TimeRangePicker) IsOpen() bool { return r.panelOpen }

// Build 实现 ui.Widget：双触发框。
func (r *TimeRangePicker) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(r)
	th := ctx.ThemeOr()
	w := r.Width
	if w <= 0 {
		w = 130
	}
	h := ui.MetricsFor(th, ui.SizeMedium).Height
	row := scene.Flex(scene.Horizontal, th.Space.XS).SetCrossAlign(scene.CrossCenter)
	sep := text.LayoutText("~", text.DefaultFont(th.Font.SizeBase), text.Options{})
	row.Add(scene.Text(scene.TextSpec{Src: "~", Layout: sep, Color: th.Color.TextFaint}))

	mkTrigger := func(tp *TimePicker, hitLabel string) *scene.Node {
		display := tp.displayText()
		txtColor := th.Color.Text
		if tp.Value.IsZero() {
			display = tp.Placeholder
			if display == "" {
				display = "请选择时间"
			}
			txtColor = th.Color.TextPlaceholder
		}
		font := text.DefaultFont(th.Font.SizeBase)
		layout := text.LayoutText(display, font, text.Options{})
		n := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.Surface),
			Border: scene.Stroke(th.Color.Border, 1),
			Radius: th.Radius.Base,
		}).SetSize2(w, h)
		label := scene.Text(scene.TextSpec{Src: display, Layout: layout, Color: txtColor})
		label.SetMargin(geom.InLTRB(11, (h-layout.Size.H)/2, 0, 0))
		n.Add(label)
		ic := (&Icon{Name: "clock", Size: 14, Color: th.Color.TextFaint}).Build(ctx)
		ic.SetPositionKeepSize(true)
		ic.SetStackPosition(geom.Pt(w-11-14, (h-14)/2))
		n.Add(ic)
		n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: hitLabel, Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			anchor := n
			g := &ui.GestureDetector{OnTap: func(geom.Point) {
				if r.panelOpen {
					r.closePanel(ctx)
					return
				}
				r.openPanel(ctx, anchor, th)
			}}
			ctx.Input.Set(n, g.Handler())
		}
		return n
	}

	row.Add(mkTrigger(&r.Start, "time-range-start"))
	row.Add(mkTrigger(&r.End, "time-range-end"))
	row.SetHit(scene.HitSpec{Label: "time-range"})
	r.childNode = row
	return row
}

// Show 程序化打开面板（锚=最近构建的触发框；测试/自动化用）。
func (r *TimeRangePicker) Show(ctx *ui.BuildContext) {
	if r.childNode == nil {
		return
	}
	r.openPanel(ctx, r.childNode, ctx.ThemeOr())
}

// openPanel 打开双面板浮层（两个面板并排 + 共享确定）。
func (r *TimeRangePicker) openPanel(ctx *ui.BuildContext, anchor *scene.Node, th *theme.Data) {
	if r.handle != 0 {
		return
	}
	r.curTh = th
	r.panelOpen = true
	r.activeEnd = 0
	r.Start.curTh = th
	r.End.curTh = th
	panel := r.buildPanel(ctx, th)
	// 浮层宽高需给显式值（W=0 会把 0 传给 Layout 约束，面板宽度算不对）。
	// 宽 = 起列组宽 + 分隔 + 止列组宽。
	_, sw := r.Start.columnsWidth()
	_, ew := r.End.columnsWidth()
	r.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: overlay.PlacementBottomStart,
		Size:      geom.Size{W: sw + 1 + ew, H: 0}, // 高自适应
		Dismiss:   overlay.DismissOutsideClick | overlay.DismissEsc,
		Tag:       "time-range-picker",
		OnClose: func() {
			r.handle = 0
			r.panelOpen = false
		},
		Build: func() *scene.Node { return panel },
	})
}

// closePanel 关闭面板。
func (r *TimeRangePicker) closePanel(ctx *ui.BuildContext) {
	if r.handle != 0 && r.fb.host != nil {
		r.fb.host.Close(r.handle)
	}
	r.handle = 0
	r.panelOpen = false
}

// Close 程序化关闭。
func (r *TimeRangePicker) Close() {
	if r.fb.host != nil {
		r.closePanel(nil)
	}
}

// buildPanel 双面板并排 + 共享 footer（"现在/确定"）。
func (r *TimeRangePicker) buildPanel(ctx *ui.BuildContext, th *theme.Data) *scene.Node {
	outer := scene.Flex(scene.Vertical, 0)
	outer.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	outer.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "time-range-panel"})

	// 两个面板并排（各自列组），中间 1px 分隔。
	panels := scene.Flex(scene.Horizontal, 0)
	startCols, startW := r.Start.buildColumns(ctx, th)
	endCols, endW := r.End.buildColumns(ctx, th)
	panels.Add(startCols)
	panels.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetSize2(1, timeViewportH))
	panels.Add(endCols)
	outer.Add(panels)

	totalW := startW + 1 + endW

	// 底部分隔线 + footer（"现在/确定"）。
	outer.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetSize2(totalW, 1))
	footer := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	footer.SetPadding(geom.InLTRB(th.Space.SM, th.Space.XS, th.Space.SM, th.Space.XS))
	footer.SetWidth(totalW)
	footer.Add(scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight))

	btnNow := &Button{Text: "现在", Type: TypeText, Size: ui.SizeSmall, OnClick: func() {
		now := time.Now()
		if r.activeEnd == 0 {
			r.Start.Value = now
			r.Start.highlightAll(now.Hour(), now.Minute(), now.Second())
		} else {
			r.End.Value = now
			r.End.highlightAll(now.Hour(), now.Minute(), now.Second())
		}
		r.MarkDirty()
	}}
	nowWrap := scene.Flex(scene.Horizontal, 0)
	nowWrap.Add(btnNow.Build(ctx))
	nowWrap.SetHit(scene.HitSpec{Label: "time-range-now"})
	footer.Add(nowWrap)

	btnOk := &Button{Text: "确定", Type: TypeLink, Size: ui.SizeSmall, OnClick: func() {
		s, e := r.Start.Value, r.End.Value
		if !s.IsZero() && !e.IsZero() && e.Before(s) {
			s, e = e, s
		}
		r.Start.Value, r.End.Value = s, e
		r.MarkDirty()
		if r.OnChangeRange != nil {
			r.OnChangeRange(s, e)
		}
		if ctx.Overlay != nil && r.handle != 0 {
			ctx.Overlay.Close(r.handle)
		}
		r.handle = 0
		r.panelOpen = false
	}}
	okWrap := scene.Flex(scene.Horizontal, 0)
	okWrap.Add(btnOk.Build(ctx))
	okWrap.SetHit(scene.HitSpec{Label: "time-range-ok"})
	footer.Add(okWrap)
	outer.Add(footer)
	return outer
}
