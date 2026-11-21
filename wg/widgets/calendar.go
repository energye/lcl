package widgets

// ── Calendar 日历（M12 长尾）─────────────────────────────────────
//
// antd：<Calendar fullscreen value onSelect onPanelChange>
// 整页日历：页头（年月标题 + 上月/今天/下月按钮）+ 星期表头 +
// 6×7 日期格（Flex 行方案——DatePicker 轮实证的零歧义通道；页面级
// 虽可用 Grid，但与浮层面板保持同一构建路径降低心智负担）。
// 节点色：today 主色描边、选中主色实底；点击格 onSelect。
// 未实现（如实标注）：Month/Year 模式切换、农历/节假日、单元格
// 自定义内容（dateCellRender）、locale、validRange。

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// Calendar 是整页日历（全屏形态）。
type Calendar struct {
	controlBase
	Value    time.Time // 选中日期（零值=未选）
	OnSelect func(date time.Time)
	// OnPanelChange 面板月份切换回调。
	OnPanelChange func(month time.Time)
	Width         float32

	viewYear  int
	viewMonth int
	navInited bool
	// Now current moment source (injectable fixed value for audit/regression cross-day determinism; nil defaults to time.Now).
	Now func() time.Time
}

// NewCalendar 创建整页日历（默认当前月）。
func NewCalendar() *Calendar {
	return &Calendar{}
}

// Build 实现 ui.Widget。
func (c *Calendar) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(c)
	th := ctx.ThemeOr()
	w := c.Width
	if w <= 0 {
		w = 640
	}
	now := c.NowNow()
	if !c.navInited {
		c.viewYear, c.viewMonth = now.Year(), int(now.Month())
		c.navInited = true
	}

	col := scene.Flex(scene.Vertical, 0)
	col.SetWidth(w)
	col.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	col.SetHit(scene.HitSpec{Label: "calendar"})

	// 页头：年月大字 + 上月/今天/下月。
	head := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	head.SetWidth(w)
	head.SetPadding(geom.InLTRB(th.Space.MD, th.Space.SM, th.Space.MD, th.Space.SM))
	title := fmt.Sprintf("%d年%d月", c.viewYear, c.viewMonth)
	tl := text.LayoutText(title, text.DefaultFont(th.Font.SizeLG), text.Options{})
	ht := scene.Text(scene.TextSpec{Src: title, Layout: tl, Color: th.Color.Text})
	ht.SetHit(scene.HitSpec{Label: "calendar-title"})
	head.Add(ht)
	head.Add(scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight))
	mkNav := func(label string, hit string, fn func()) {
		b := &Button{Text: label, Size: ui.SizeSmall, OnClick: fn}
		wrap := scene.Flex(scene.Horizontal, 0)
		wrap.Add(b.Build(ctx))
		wrap.SetHit(scene.HitSpec{Label: hit})
		head.Add(wrap)
	}
	mkNav("上月", "calendar-nav-prev", func() { c.shiftMonth(ctx, -1) })
	mkNav("今天", "calendar-nav-today", func() { c.goToday(ctx) })
	mkNav("下月", "calendar-nav-next", func() { c.shiftMonth(ctx, 1) })
	col.Add(head)

	// 星期表头。
	week := scene.Flex(scene.Horizontal, 0)
	week.SetWidth(w)
	week.SetPadding(geom.InLTRB(6, 0, 6, 0))
	for _, wd := range []string{"一", "二", "三", "四", "五", "六", "日"} {
		cell := scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight).SetHeight(30)
		l := text.LayoutText(wd, text.DefaultFont(th.Font.SizeSM), text.Options{})
		tv := scene.Text(scene.TextSpec{Src: wd, Layout: l, Color: th.Color.TextMuted})
		tv.SetMargin(geom.InLTRB(8, (30-l.Size.H)/2, 0, 0))
		cell.Add(tv)
		week.Add(cell)
	}
	col.Add(week)

	// 日期区：6 行 Flex（DatePicker 轮实证通道）。
	first := time.Date(c.viewYear, time.Month(c.viewMonth), 1, 0, 0, 0, 0, time.Local)
	offset := (int(first.Weekday()) + 6) % 7 // 周一首
	daysInMonth := time.Date(c.viewYear, time.Month(c.viewMonth)+1, 0, 0, 0, 0, 0, time.Local).Day()
	today := c.NowNow()
	grid := scene.Flex(scene.Vertical, 0)
	grid.SetWidth(w)
	grid.SetPadding(geom.InLTRB(6, 0, 6, th.Space.XS))
	weekRow := scene.Flex(scene.Horizontal, 0).SetWidth(w)
	daysThisRow := offset
	for i := 0; i < offset; i++ {
		weekRow.Add(scene.Empty())
	}
	flushRow := func() {
		if daysThisRow > 0 {
			grid.Add(weekRow)
		}
		weekRow = scene.Flex(scene.Horizontal, 0).SetWidth(w)
		daysThisRow = 0
	}
	for day := 1; day <= daysInMonth; day++ {
		day := day
		date := time.Date(c.viewYear, time.Month(c.viewMonth), day, 0, 0, 0, 0, time.Local)
		isToday := today.Year() == date.Year() && today.Month() == date.Month() && today.Day() == date.Day()
		isSel := !c.Value.IsZero() &&
			c.Value.Year() == date.Year() && c.Value.Month() == date.Month() && c.Value.Day() == date.Day()

		cellW := (w - 12) / 7
		cell := scene.Box(scene.BoxSpec{
			Fill: func() render.Paint {
				if isSel {
					return scene.Solid(th.Color.Primary)
				}
				return render.Paint{}
			}(),
			Border: func() render.Paint {
				if isToday && !isSel {
					return scene.Stroke(th.Color.Primary, 1)
				}
				return scene.Stroke(th.Color.Border, 1)
			}(),
		}).SetSize2(cellW, 52)
		txtColor := th.Color.Text
		if isSel {
			txtColor = th.Color.OnPrimary
		} else if isToday {
			txtColor = th.Color.Primary
		}
		l := text.LayoutText(fmt.Sprintf("%d", day), text.DefaultFont(th.Font.SizeBase), text.Options{})
		num := scene.Text(scene.TextSpec{Src: fmt.Sprintf("%d", day), Layout: l, Color: txtColor})
		num.SetMargin(geom.InLTRB(8, 6, 0, 0))
		cell.Add(num)
		cell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "calendar-cell", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			ctx.Input.Set(cell, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				if !ev.Down() {
					return false
				}
				c.Value = date
				c.MarkDirty()
				if c.OnSelect != nil {
					c.OnSelect(date)
				}
				return true
			}})
		}
		weekRow.Add(cell)
		daysThisRow++
		if daysThisRow == 7 {
			flushRow()
		}
	}
	flushRow()
	col.Add(grid)
	return col
}

// shiftMonth 面板月份切换（±1，跨年进退位）。
func (c *Calendar) shiftMonth(ctx *ui.BuildContext, dx int) {
	m := c.viewMonth + dx
	if m < 1 {
		m = 12
		c.viewYear--
	} else if m > 12 {
		m = 1
		c.viewYear++
	}
	c.viewMonth = m
	c.MarkDirty()
	if c.OnPanelChange != nil {
		c.OnPanelChange(time.Date(c.viewYear, time.Month(c.viewMonth), 1, 0, 0, 0, 0, time.Local))
	}
}

// NowNow current moment (Now injection or time.Now).
func (c *Calendar) NowNow() time.Time {
	if c.Now != nil {
		return c.Now()
	}
	return time.Now()
}

// goToday 回当月并选中今天。
func (c *Calendar) goToday(ctx *ui.BuildContext) {
	now := c.NowNow()
	c.viewYear, c.viewMonth = now.Year(), int(now.Month())
	c.Value = now
	c.MarkDirty()
	if c.OnSelect != nil {
		c.OnSelect(now)
	}
	if c.OnPanelChange != nil {
		c.OnPanelChange(now)
	}
}
