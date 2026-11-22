package widgets

// ── Calendar 日历（M12 长尾）─────────────────────────────────────
//
// antd：<Calendar fullscreen value onSelect onPanelChange>
// 整页日历：页头（年月标题 + 上月/今天/下月按钮）+ 星期表头 +
// 6×7 日期格（Flex 行方案——DatePicker 轮实证的零歧义通道；页面级
// 虽可用 Grid，但与浮层面板保持同一构建路径降低心智负担）。
// 节点色：today 主色描边、选中主色实底；点击格 onSelect。
// dateCellRender/validRange/Month-Year 模式已交付（第一百零九/
// 一十五轮）；未实现（如实标注）：农历/节假日、locale。

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
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
	// DateCellRender 自定义日期格内容（antd cellRender 的"数字+追加"
	// 形态）：非空时该格在日期数字下方渲染返回的 Widget；返回 nil =
	// 无自定义内容。每次构建调用（与单元格 Render 同款契约）。
	DateCellRender func(date time.Time) ui.Widget
	// Mode 面板模式（antd mode：month 日历默认 / year 年面板——12
	// 月格点选切月）。
	Mode string
	// ValidRange 有效范围（antd validRange：范围外日期灰置不可选、
	// 月份切换钳制在范围内）。
	ValidRange [2]time.Time

	viewYear  int
	viewMonth int
	navInited bool
	// navBtns 导航按钮池（教训⑨：Build 内新建按压态随重建丢失——
	// 多实例页启发式审计逐一点击时 press=0 FAIL）。
	navBtns [4]*Button
	// yearBtn 年面板返回按钮池。
	yearBtn *Button
	// Now current moment source (injectable fixed value for audit/regression cross-day determinism; nil defaults to time.Now).
	Now func() time.Time
}

// NewCalendar 创建整页日历（默认当前月）。
func NewCalendar() *Calendar {
	return &Calendar{}
}

// SetView 设置初始视图年月（宿主跨天确定性注入用——viewYear/viewMonth
// 私有，审计/回归经此固定面板月）。
func (c *Calendar) SetView(year, month int) {
	c.viewYear, c.viewMonth, c.navInited = year, month, true
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
		// validRange 非空：初始视图钳进范围。
		if !c.ValidRange[0].IsZero() && !c.ValidRange[1].IsZero() {
			start, end := c.ValidRange[0], c.ValidRange[1]
			cur := time.Date(c.viewYear, time.Month(c.viewMonth), 1, 0, 0, 0, 0, time.Local)
			if cur.Before(time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, time.Local)) {
				c.viewYear, c.viewMonth = start.Year(), int(start.Month())
			} else if cur.After(time.Date(end.Year(), end.Month(), 1, 0, 0, 0, 0, time.Local)) {
				c.viewYear, c.viewMonth = end.Year(), int(end.Month())
			}
		}
		c.navInited = true
	}
	// year 模式：年面板（12 月格）独立构建路径。
	if c.Mode == "year" {
		return c.buildYearPanel(ctx, th, w)
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
	mkNav := func(label string, hit string, idx int, fn func()) {
		if c.navBtns[idx] == nil {
			c.navBtns[idx] = &Button{Size: ui.SizeSmall}
		}
		b := c.navBtns[idx]
		b.Text = label
		b.OnClick = fn
		wrap := scene.Flex(scene.Horizontal, 0)
		wrap.Add(b.Build(ctx))
		wrap.SetHit(scene.HitSpec{Label: hit})
		head.Add(wrap)
	}
	mkNav("上月", "calendar-nav-prev", 0, func() { c.shiftMonth(ctx, -1) })
	mkNav("今天", "calendar-nav-today", 1, func() { c.goToday(ctx) })
	mkNav("下月", "calendar-nav-next", 2, func() { c.shiftMonth(ctx, 1) })
	// 模式切换（antd mode select）：月面板 ↔ 年面板。
	modeLabel := "年面板"
	if c.Mode == "year" {
		modeLabel = "月面板"
	}
	mkNav(modeLabel, "calendar-nav-mode", 3, func() {
		if c.Mode == "year" {
			c.Mode = "month"
		} else {
			c.Mode = "year"
		}
		c.MarkDirty()
	})
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
		// validRange 范围外日期：灰置（TextFaint）不可选。
		outOfRange := !c.ValidRange[0].IsZero() && !c.ValidRange[1].IsZero() &&
			(date.Before(c.ValidRange[0]) || date.After(c.ValidRange[1]))
		if outOfRange {
			txtColor = th.Color.TextFaint
		}
		l := text.LayoutText(fmt.Sprintf("%d", day), text.DefaultFont(th.Font.SizeBase), text.Options{})
		num := scene.Text(scene.TextSpec{Src: fmt.Sprintf("%d", day), Layout: l, Color: txtColor})
		num.SetMargin(geom.InLTRB(8, 6, 0, 0))
		cell.Add(num)
		// 自定义日期格内容：数字下方追加（antd cellRender 语义——数字
		// 仍在，自定义内容渲染在其下）。nil 返回=纯数字格。
		if c.DateCellRender != nil {
			if wgt := c.DateCellRender(date); wgt != nil {
				wn := wgt.Build(ctx)
				wn.SetMargin(geom.InLTRB(8, 2, 2, 0))
				cell.Add(wn)
			}
		}
		cell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "calendar-cell", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			ctx.Input.Set(cell, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				if !ev.Down() {
					return false
				}
				// validRange 范围外日期不可选（antd 语义）。
				if outOfRange {
					return true
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

// shiftMonth 面板月份切换（±1，跨年进退位）；validRange 非空时钳制
// 在范围内（antd 语义：面板切换不越界）。
func (c *Calendar) shiftMonth(ctx *ui.BuildContext, dx int) {
	m := c.viewMonth + dx
	if m < 1 {
		m = 12
		c.viewYear--
	} else if m > 12 {
		m = 1
		c.viewYear++
	}
	// validRange 钳制：目标月首越界则停在本月。
	if !c.ValidRange[0].IsZero() && !c.ValidRange[1].IsZero() {
		cand := time.Date(c.viewYear, time.Month(m), 1, 0, 0, 0, 0, time.Local)
		start := time.Date(c.ValidRange[0].Year(), c.ValidRange[0].Month(), 1, 0, 0, 0, 0, time.Local)
		end := time.Date(c.ValidRange[1].Year(), c.ValidRange[1].Month(), 1, 0, 0, 0, 0, time.Local)
		if cand.Before(start) || cand.After(end) {
			return
		}
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
	// validRange 钳制：今天不在范围内时回范围内首月。
	if !c.ValidRange[0].IsZero() && !c.ValidRange[1].IsZero() && (now.Before(c.ValidRange[0]) || now.After(c.ValidRange[1])) {
		c.viewYear, c.viewMonth = c.ValidRange[0].Year(), int(c.ValidRange[0].Month())
	}
	c.Value = now
	c.MarkDirty()
	if c.OnSelect != nil {
		c.OnSelect(now)
	}
	if c.OnPanelChange != nil {
		c.OnPanelChange(now)
	}
}

// buildYearPanel 年面板（antd mode="year"）：4×3 月格网格，点选切到
// 对应月的 month 面板；validRange 范围外月份灰置不可点。
func (c *Calendar) buildYearPanel(ctx *ui.BuildContext, th *theme.Data, w float32) *scene.Node {
	col := scene.Flex(scene.Vertical, th.Space.SM)
	col.SetWidth(w)
	col.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	col.SetHit(scene.HitSpec{Label: "calendar"})

	// 页头：年份大字 + 返回月面板。
	head := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	head.SetWidth(w)
	head.SetPadding(geom.InLTRB(th.Space.MD, th.Space.SM, th.Space.MD, th.Space.SM))
	title := fmt.Sprintf("%d年", c.viewYear)
	tl := text.LayoutText(title, text.DefaultFont(th.Font.SizeLG), text.Options{})
	ht := scene.Text(scene.TextSpec{Src: title, Layout: tl, Color: th.Color.Text})
	ht.SetHit(scene.HitSpec{Label: "calendar-year-title"})
	head.Add(ht)
	head.Add(scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight))
	if c.yearBtn == nil {
		c.yearBtn = &Button{Size: ui.SizeSmall}
	}
	c.yearBtn.Text = "月面板"
	c.yearBtn.OnClick = func() {
		c.Mode = "month"
		c.MarkDirty()
	}
	head.Add(c.yearBtn.Build(ctx))
	col.Add(head)

	// 4×3 月格。
	grid := scene.Grid(4, th.Space.SM)
	cur := time.Date(c.viewYear, time.Month(c.viewMonth), 1, 0, 0, 0, 0, time.Local)
	for m := 1; m <= 12; m++ {
		m := m
		cell := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.Surface),
			Border: scene.Stroke(th.Color.Border, 1),
		}).SetHeight(64)
		// validRange 范围外月份灰置不可点。
		md := time.Date(c.viewYear, time.Month(m), 1, 0, 0, 0, 0, time.Local)
		outOfRange := !c.ValidRange[0].IsZero() && !c.ValidRange[1].IsZero() &&
			(md.Before(time.Date(c.ValidRange[0].Year(), c.ValidRange[0].Month(), 1, 0, 0, 0, 0, time.Local)) ||
				md.After(time.Date(c.ValidRange[1].Year(), c.ValidRange[1].Month(), 1, 0, 0, 0, 0, time.Local)))
		isCur := m == c.viewMonth
		txtColor := th.Color.Text
		if outOfRange {
			txtColor = th.Color.TextFaint
		} else if isCur {
			txtColor = th.Color.Primary
		}
		l := text.LayoutText(fmt.Sprintf("%d月", m), text.DefaultFont(th.Font.SizeBase), text.Options{})
		tv := scene.Text(scene.TextSpec{Src: fmt.Sprintf("%d月", m), Layout: l, Color: txtColor})
		tv.SetMargin(geom.InLTRB(12, (64-l.Size.H)/2, 0, 0))
		cell.Add(tv)
		cell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "calendar-month-cell", Cursor: scene.CursorPointer})
		if ctx.Input != nil && !outOfRange {
			ctx.Input.Set(cell, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				if !ev.Down() {
					return false
				}
				c.viewMonth = m
				c.Mode = "month" // 选中月回到日历面板（antd 语义）
				c.MarkDirty()
				if c.OnPanelChange != nil {
					c.OnPanelChange(time.Date(c.viewYear, time.Month(m), 1, 0, 0, 0, 0, time.Local))
				}
				return true
			}})
		}
		_ = cur
		grid.Add(cell)
	}
	col.Add(grid)
	return col
}
