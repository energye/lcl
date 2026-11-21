package widgets

// ── DatePicker（M12 高级录入）──────────────────────────────────────
//
// antd：<DatePicker value onChange placeholder>
// 本轮（第一轮）：触发框（Input 形态）+ 点击弹日历面板 + 日期格选择
// （6×7 Grid）+ 月切换（‹ ›）。未实现：范围选择/年月面板/禁用日期/
// 时间联动（后续轮）。
//
// 组合复用：Icon（左右箭头）、Grid 布局通道、BandPopup+PlacementBottom
// （Dropdown 同浮层通道）。

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// DatePicker 是日期选择器（面板骨架轮：日/月粒度）。
type DatePicker struct {
	controlBase
	Value        time.Time
	Placeholder  string
	Width        float32
	OnChange     func(t time.Time)
	Disabled     bool
	PlaceholderT string // 面板无值时显示的占位（内部）

	handle    overlay.Handle
	fb        feedbackHost
	childNode *scene.Node
	viewYear  int  // 面板当前年份
	viewMonth int  // 面板当前月份（1-12）
	panelOpen bool
}

// NewDatePicker 创建日期选择器（默认当前日期）。
func NewDatePicker() *DatePicker {
	now := time.Now()
	return &DatePicker{Value: now, Placeholder: "请选择日期", PlaceholderT: "请选择日期"}
}

// Build 实现 ui.Widget：触发框（Input 形态）。
func (d *DatePicker) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(d)
	th := ctx.ThemeOr()
	w := d.Width
	if w <= 0 {
		w = 200
	}
	h := ui.MetricsFor(th, ui.SizeMedium).Height
	display := d.Value.Format("2006-01-02")
	if d.Value.IsZero() {
		display = d.Placeholder
	}
	font := text.DefaultFont(th.Font.SizeBase)
	layout := text.LayoutText(display, font, text.Options{})

	n := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	}).SetSize2(w, h)
	txtColor := th.Color.Text
	if d.Value.IsZero() {
		txtColor = th.Color.TextPlaceholder
	}
	label := scene.Text(scene.TextSpec{Src: display, Layout: layout, Color: txtColor})
	label.SetMargin(geom.InLTRB(11, (h-layout.Size.H)/2, 0, 0))
	n.Add(label)
	// 右缘日历图标（组合复用 Icon）。
	ic := (&Icon{Name: "calendar", Size: 14, Color: th.Color.TextFaint}).Build(ctx)
	ic.SetPositionKeepSize(true)
	ic.SetStackPosition(geom.Pt(w-11-14, (h-14)/2))
	n.Add(ic)
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "date-picker", Cursor: scene.CursorPointer})

	if ctx.Input != nil {
		anchor := n
		g := &ui.GestureDetector{OnTap: func(geom.Point) {
			if d.panelOpen {
				d.closePanel(ctx)
				return
			}
			d.openPanel(ctx, anchor, th)
		}}
		ctx.Input.Set(n, g.Handler())
	}
	d.childNode = n
	return n
}

// openPanel 打开日历面板（BandPopup+PlacementBottomStart，Dropdown 同通道）。
func (d *DatePicker) openPanel(ctx *ui.BuildContext, anchor *scene.Node, th *theme.Data) {
	if d.handle != 0 {
		return
	}
	if d.Value.IsZero() {
		now := time.Now()
		d.viewYear, d.viewMonth = now.Year(), int(now.Month())
	} else {
		d.viewYear, d.viewMonth = d.Value.Year(), int(d.Value.Month())
	}
	d.panelOpen = true
	panel := d.buildPanel(ctx, th)
	d.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: overlay.PlacementBottomStart,
		Size:      geom.Size{W: 280, H: 0}, // 高度自适应
		Dismiss:   overlay.DismissOutsideClick,
		Tag:       "datepicker",
		OnClose: func() {
			d.handle = 0
			d.panelOpen = false
		},
		Build: func() *scene.Node { return panel },
	})
}

// closePanel 关闭面板。
func (d *DatePicker) closePanel(ctx *ui.BuildContext) {
	if d.handle != 0 && d.fb.host != nil {
		d.fb.host.Close(d.handle)
	}
	d.handle = 0
	d.panelOpen = false
}

// buildPanel 构建日历面板：月标题+切换箭头+星期行+6×7 日期格。
func (d *DatePicker) buildPanel(ctx *ui.BuildContext, th *theme.Data) *scene.Node {
	panel := scene.Flex(scene.Vertical, th.Space.XS).SetWidth(280)
	panel.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	panel.SetPadding(geom.InLTRB(th.Space.SM, th.Space.SM, th.Space.SM, th.Space.SM))

	// 月标题行：‹ 2026年11月 ›（切换箭头用 Icon 组件组合复用）。
	head := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	head.SetHit(scene.HitSpec{Label: "date-panel-head"})
	title := fmt.Sprintf("%d年%d月", d.viewYear, d.viewMonth)
	tl := text.LayoutText(title, text.DefaultFont(th.Font.SizeBase), text.Options{})
	ht := scene.Text(scene.TextSpec{Src: title, Layout: tl, Color: th.Color.Text})
	ht.SetHit(scene.HitSpec{Label: "date-panel-title"})
	head.Add(ht)
	sp := scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight)
	head.Add(sp)
	for _, dir := range []struct {
		icon string
		dx   int
	}{
		{"right", -1}, // ‹ 上一月（图标向右箭头旋转无——直接用 right 表示向前后退）
		{"left", 1},   // › 下一月
	} {
		ic := (&Icon{Name: dir.icon, Size: 14, Color: th.Color.TextMuted}).Build(ctx)
		ic.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "date-nav", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			mv := dir.dx
			g := &ui.GestureDetector{OnTap: func(geom.Point) {
				d.shiftMonth(mv)
			}}
			ctx.Input.Set(ic, g.Handler())
		}
		head.Add(ic)
	}
	panel.Add(head)

	// 星期行（一~日）。
	week := scene.Flex(scene.Horizontal, 0)
	for _, w := range []string{"一", "二", "三", "四", "五", "六", "日"} {
		cell := scene.Flex(scene.Vertical, 0).SetCrossAlign(scene.CrossCenter).SetWidth(38)
		l := text.LayoutText(w, text.DefaultFont(th.Font.SizeXS), text.Options{})
		cell.Add(scene.Text(scene.TextSpec{Src: w, Layout: l, Color: th.Color.TextFaint}))
		week.Add(cell)
	}
	panel.Add(week)

	// 日期格 6×7：Grid 等分（组合复用 Grid 通道）。
	first := time.Date(d.viewYear, time.Month(d.viewMonth), 1, 0, 0, 0, 0, time.Local)
	// 周一为首：offset=(weekday+6)%7
	offset := (int(first.Weekday()) + 6) % 7
	daysInMonth := time.Date(d.viewYear, time.Month(d.viewMonth)+1, 0, 0, 0, 0, 0, time.Local).Day()
	today := time.Now()
	grid := scene.Grid(7, 0)
	for i := 0; i < offset; i++ {
		grid.Add(scene.Empty())
	}
	for day := 1; day <= daysInMonth; day++ {
		day := day
		t := time.Date(d.viewYear, time.Month(d.viewMonth), day, 0, 0, 0, 0, time.Local)
		selected := !d.Value.IsZero() &&
			d.Value.Year() == t.Year() && d.Value.Month() == t.Month() && d.Value.Day() == t.Day()
		isToday := today.Year() == t.Year() && today.Month() == t.Month() && today.Day() == t.Day()
		cell := scene.Flex(scene.Vertical, 0).SetCrossAlign(scene.CrossCenter).SetWidth(38).SetHeight(28)
		bg := geom.Color{}
		fg := th.Color.Text
		if selected {
			bg, fg = th.Color.Primary, th.Color.OnPrimary
		} else if isToday {
			fg = th.Color.Primary
		}
		cell.SetBox(scene.BoxSpec{Fill: scene.Solid(bg), Radius: th.Radius.SM})
		l := text.LayoutText(fmt.Sprintf("%d", day), text.DefaultFont(th.Font.SizeSM), text.Options{})
		txt := scene.Text(scene.TextSpec{Src: fmt.Sprintf("%d", day), Layout: l, Color: fg})
		txt.SetMargin(geom.InLTRB(0, (28-l.Size.H)/2, 0, 0))
		cell.Add(txt)
		cell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "date-cell", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			// **down 即执行**（Dropdown 轮结论：浮层内 tap 配对永不触发）。
			ctx.Input.Set(cell, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				if !ev.Down() {
					return false
				}
				d.Value = t
				d.MarkDirty()
				// 关面板**直接用 ctx.Overlay**（d.fb.host 在该链路为 nil——
				// captureHost 依赖 Build 期 ctx.Overlay 注入，浮层 cell 的
				// down 闭包捕获的 ctx 同源但实测 host=nil，根因待查；
				// ctx.Overlay 直关是同对象、必有效）。
				if ctx.Overlay != nil && d.handle != 0 {
					ctx.Overlay.Close(d.handle)
				}
				d.handle = 0
				d.panelOpen = false
				if d.OnChange != nil {
					d.OnChange(t)
				}
				return true
			}})
		}
		grid.Add(cell)
	}
	panel.Add(grid)
	panel.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "date-panel"})
	return panel
}

// shiftMonth 切换面板月份（dx=±1），跨年自动进退位。
func (d *DatePicker) shiftMonth(dx int) {
	m := d.viewMonth + dx
	if m < 1 {
		m = 12
		d.viewYear--
	} else if m > 12 {
		m = 1
		d.viewYear++
	}
	d.viewMonth = m
	d.MarkDirty()
}

// Close 关闭面板。
func (d *DatePicker) Close() {
	if d.fb.host != nil {
		d.closePanel(nil)
	}
}
