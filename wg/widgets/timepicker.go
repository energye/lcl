package widgets

// ── TimePicker（M12 高级录入）──────────────────────────────────────
//
// antd：<TimePicker value onChange placeholder format="HH:mm:ss">
// 第一轮：触发框（Input 形态+clock 图标）+ 面板三列（时/分/秒，
// antd 令牌：列 56px、cell 28px 高、视口 224px）+ 选择高亮（选中
// 主色实底、hover colorFillTertiary）+ 底部"现在/确定"（Button
// TypeText/TypeLink 组合复用——antd footer 即文本按钮）。
//
// antd needConfirm 默认形态（time picker 面板有确定钮）：**选择不关
// 面板**（onChange 已触发、高亮迁移），"确定"/Esc/外部点击才关闭——
// 与 DatePicker 单选（选完即关）的核心行为差异，spec 专断言。
//
// 浮层面板的动态刷新不走根重建（Entry.Build 只在 Show 时调一次，
// 面板子树不在重建链上）：选中/hover 一律 **SetBox/SetTextPaint
// 原地 mutation**（O(1)，零重建——§6A 极致要求）。
//
// 组合复用：Icon（clock）、Scroll（三列滚动视口，Select 选项列表
// 同款）、BandPopup+PlacementBottomStart（Select/DatePicker 同通道）、
// Button（footer 文本按钮）。
// use12Hours（第一百二十九轮）：12 小时列+上午/下午第 4 列（antd
// use12Hours——触发框 h:mm:ss AM/PM 形态；12 列值 12,1..11 与 AM/PM
// 组合还原 24h）；hourStep/minuteStep/secondStep 列生成按步长过滤
//（antd hourStep 等——值集合为步长倍数）。
// 未实现（如实标注）：disabledTime 系列 / hideDisabledOptions /
// allowClear 清除图标 / format 列裁剪（列固定）/ RangePicker / 受控
// open / placement 四向 / changeOnScroll / 无值时滚动定位到当前时刻。

import (
	"fmt"
	"strings"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// 面板尺寸令牌（antd design token：列 56 宽、cell 28 高、视口 224）。
const (
	timeColW      = float32(56)
	timeCellH     = float32(28)
	timeViewportH = float32(224)
)

// DisabledTimeSpec 是禁用时段（antd DisabledTime 返回结构）。
// 各回调返回该列不可选的值列表；nil = 该列无禁用。
type DisabledTimeSpec struct {
	DisabledHours   func() []int
	DisabledMinutes func(hour int) []int
	DisabledSeconds func(hour, minute int) []int
}

// timeCellInfo 一格的节点与排版（原地改色所需）。
type timeCellInfo struct {
	cell   *scene.Node
	txt    *scene.Node
	src    string
	value  int
	layout *text.Layout
}

// TimePicker 是时间选择器（时/分/秒三列面板）。
type TimePicker struct {
	controlBase
	Value       time.Time
	Placeholder string
	Width       float32
	OnChange    func(t time.Time)
	// Use12Hours 12 小时制（antd use12Hours）：面板加上午/下午第 4 列，
	// 小时列 12 格（值 12,1..11），触发框显示 "hh:mm:ss 上午/下午"。
	Use12Hours bool
	// HourStep/MinuteStep/SecondStep 列值步长（antd hourStep 等——列
	// 只渲染步长倍数；0/1=不步进）。
	HourStep   int
	MinuteStep int
	SecondStep int
	// DisabledTime 禁用时段回调（antd disabledTime）：按当前值返回各列
	// 不可选值；nil = 不禁用。禁用格灰置且点击无效。
	DisabledTime func(now time.Time) DisabledTimeSpec
	// Format 时间格式（Go 布局串，antd format 对应）：省略分钟则分列消失、
	// 省略秒则秒列消失（antd hide-column 语义）。空 = 默认 15:04:05
	// （12h 为 03:04:05 下午）。
	Format string

	handle    overlay.Handle
	fb        feedbackHost
	childNode *scene.Node
	panelOpen bool
	curTh     *theme.Data

	// 面板打开期间的 cell 表与选中下标（选中/hover 原地改，零重建）。
	// 12h 模式 4 列：0=小时(0..11 映射)、1..2=分/秒、3=上午/下午。
	cellsByCol [4][]*timeCellInfo
	selIdx     [4]int // -1 = 该列未选中；3 列模式第 4 列恒 -1

	// OnValueChange 内部值变化钩子（区间选择器监听每端变化切端用；
	// 不影响公开 OnChange 语义）。
	OnValueChange func()
}

// NewTimePicker 创建时间选择器（默认无值——显示占位）。
func NewTimePicker() *TimePicker {
	return &TimePicker{Placeholder: "请选择时间"}
}

// Build 实现 ui.Widget：触发框（Input 形态）。
func (t *TimePicker) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(t)
	th := ctx.ThemeOr()
	t.curTh = th
	w := t.Width
	if w <= 0 {
		w = 120
	}
	h := ui.MetricsFor(th, ui.SizeMedium).Height
	display := t.displayText()
	txtColor := th.Color.Text
	if t.Value.IsZero() {
		display, txtColor = t.Placeholder, th.Color.TextPlaceholder
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
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "time-picker", Cursor: scene.CursorPointer})

	if ctx.Input != nil {
		anchor := n
		g := &ui.GestureDetector{OnTap: func(geom.Point) {
			if t.panelOpen {
				t.closePanel(ctx)
				return
			}
			t.openPanel(ctx, anchor, th)
		}}
		ctx.Input.Set(n, g.Handler())
	}
	t.childNode = n
	return n
}

// Show 程序化打开面板（锚=最近构建的触发框节点；测试/自动化用）。
func (t *TimePicker) Show(ctx *ui.BuildContext) {
	if t.childNode == nil {
		return
	}
	t.openPanel(ctx, t.childNode, ctx.ThemeOr())
}

// IsOpen 返回面板是否打开。
func (t *TimePicker) IsOpen() bool { return t.panelOpen }

// Close 关闭面板。
func (t *TimePicker) Close() {
	if t.fb.host != nil {
		t.closePanel(nil)
	}
}

// effectiveFormat 生效的格式串：Format 非空用它，否则默认（12h 用
// h:mm:ss a 形态、24h 用 15:04:05）。
func (t *TimePicker) effectiveFormat() string {
	if t.Format != "" {
		return t.Format
	}
	if t.Use12Hours {
		return "03:04:05 下午"
	}
	return "15:04:05"
}

// columns 返回可见列（按 format 裁剪）：[showHour, showMinute, showSecond]。
// 小时列恒显示（antd TimePicker 至少保留小时）。判定用"基准时间 + 变化
// 该字段看输出是否变"：分钟取两个不同分钟、秒取两个不同秒，格式化后若
// 结果不同则说明该字段参与输出（列显示）。比字符串匹配布局标记可靠
// （"15:04" 的 5 会误命中秒、":" 会干扰）。
func (t *TimePicker) columns() [3]bool {
	f := t.effectiveFormat()
	base := time.Date(2026, 1, 2, 13, 45, 56, 0, time.Local)
	otherMin := time.Date(2026, 1, 2, 13, 7, 56, 0, time.Local)
	otherSec := time.Date(2026, 1, 2, 13, 45, 8, 0, time.Local)
	showMin := base.Format(f) != otherMin.Format(f)
	showSec := base.Format(f) != otherSec.Format(f)
	return [3]bool{true, showMin, showSec}
}

// displayText 触发框显示：按 Format（默认 24h "15:04:05"；12h "hh:mm:ss 上午/下午"）。
func (t *TimePicker) displayText() string {
	f := t.effectiveFormat()
	if t.Use12Hours {
		// Go 布局里中文"下午"是字面量，会恒显下午——用标准 PM 标记格式化
		// 后替换为中文，保证 AM/PM 正确。
		out := t.Value.Format(strings.ReplaceAll(f, "下午", "PM"))
		out = strings.ReplaceAll(out, "PM", "下午")
		out = strings.ReplaceAll(out, "AM", "上午")
		return out
	}
	return t.Value.Format(f)
}

// selectedParts 当前选中的时/分/秒（无值 {-1,-1,-1}）。
func (t *TimePicker) selectedParts() [3]int {
	if t.Value.IsZero() {
		return [3]int{-1, -1, -1}
	}
	return [3]int{t.Value.Hour(), t.Value.Minute(), t.Value.Second()}
}

// stepValues 列值集合：步长倍数（antd hourStep 等——0/1 视为 1）。
func stepValues(count, step int) []int {
	if step <= 0 || step == 1 {
		out := make([]int, count)
		for i := range out {
			out[i] = i
		}
		return out
	}
	var out []int
	for v := 0; v < count; v += step {
		out = append(out, v)
	}
	return out
}

// hourValues 小时列值集合：24h=0..23；12h=12,1..11（antd 12 列顺序，
// 显示 %02d、内部按 12h 值映射）。
func (t *TimePicker) hourValues() []int {
	if t.Use12Hours {
		out := []int{12}
		for i := 1; i <= 11; i++ {
			out = append(out, i)
		}
		if t.HourStep > 1 {
			var filtered []int
			for _, v := range out {
				if (v-1)%t.HourStep == 0 {
					filtered = append(filtered, v)
				}
			}
			return filtered
		}
		return out
	}
	return stepValues(24, t.HourStep)
}

// openPanel 打开三列面板（BandPopup+PlacementBottomStart，Select 同通道）。
func (t *TimePicker) openPanel(ctx *ui.BuildContext, anchor *scene.Node, th *theme.Data) {
	if t.handle != 0 {
		return
	}
	t.curTh = th
	t.panelOpen = true
	panel := t.buildPanel(ctx, th)
	colN := 3
	w := timeColW*3 + 2
	if t.Use12Hours {
		colN = 4
		w = timeColW*4 + 3
	}
	_ = colN
	t.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: overlay.PlacementBottomStart,
		Size:      geom.Size{W: w, H: 0}, // 高自适应（列+footer）
		Dismiss:   overlay.DismissOutsideClick | overlay.DismissEsc,
		Tag:       "timepicker",
		OnClose: func() {
			t.handle = 0
			t.panelOpen = false
		},
		Build: func() *scene.Node { return panel },
	})
}

// closePanel 关闭面板。
func (t *TimePicker) closePanel(ctx *ui.BuildContext) {
	if t.handle != 0 && t.fb.host != nil {
		t.fb.host.Close(t.handle)
	}
	t.handle = 0
	t.panelOpen = false
}

// columnsWidth 计算列组宽（不构建节点；供区间选择器算浮层宽）。
func (t *TimePicker) columnsWidth() (cols int, width float32) {
	vis := t.columns()
	cols = 1
	if vis[1] {
		cols++
	}
	if vis[2] {
		cols++
	}
	if t.Use12Hours {
		cols++
	}
	return cols, timeColW*float32(cols) + float32(cols-1)
}

// buildColumns 构建列组（时/分/秒+可选 AP，含列间分隔），返回节点与总宽。
// 单面板与区间双面板共用；sel 为当前选中时/分/秒（-1 表示无）。
func (t *TimePicker) buildColumns(ctx *ui.BuildContext, th *theme.Data) (*scene.Node, float32) {
	cols := scene.Flex(scene.Horizontal, 0)
	type colSpec struct {
		colIdx int
		label  string
		vals   []int
	}
	vis := t.columns()
	specs := []colSpec{{0, "time-cell-h", t.hourValues()}}
	if vis[1] {
		specs = append(specs, colSpec{1, "time-cell-m", stepValues(60, t.MinuteStep)})
	}
	if vis[2] {
		specs = append(specs, colSpec{2, "time-cell-s", stepValues(60, t.SecondStep)})
	}
	if t.Use12Hours {
		specs = append(specs, colSpec{3, "time-cell-ap", []int{0, 1}})
	}
	colN := len(specs)
	width := timeColW*float32(colN) + float32(colN-1)
	sel := t.selectedParts()
	t.selIdx = [4]int{-1, -1, -1, -1}
	if t.Use12Hours && sel[0] >= 0 {
		t.selIdx[3] = hourToAP(sel[0])
	}
	for k, sp := range specs {
		i := sp.colIdx
		col := t.buildColumn(ctx, th, i, sp.vals, sp.label)
		scroll := &Scroll{Axis: scene.Vertical, Width: timeColW, Height: timeViewportH,
			Child: nodeFn(func(ctx *ui.BuildContext) *scene.Node { return col })}
		if selected := t.selIdx[i]; selected >= 0 {
			off := float32(selected)*timeCellH - (timeViewportH-timeCellH)/2
			if off < 0 {
				off = 0
			}
			scroll.Offset = off
		}
		wrap := scene.Flex(scene.Vertical, 0)
		wrap.Add(scroll.Build(ctx))
		wrap.SetHit(scene.HitSpec{Label: fmt.Sprintf("time-col-%d", i)})
		cols.Add(wrap)
		if k < colN-1 {
			cols.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetSize2(1, timeViewportH))
		}
	}
	return cols, width
}

// buildPanel 构建面板：三列（时/分/秒）+列分隔+底部分隔线+"现在/确定"。
func (t *TimePicker) buildPanel(ctx *ui.BuildContext, th *theme.Data) *scene.Node {
	panel := scene.Flex(scene.Vertical, 0)
	panel.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	panel.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "time-panel"})

	// 列：时/分/秒（12h 加上午/下午第 4 列）；Scroll 组合复用（Select
	// 选项列表同款视口）；列间 1px 分隔。format 裁剪：不显示的列直接跳过
	// （列号语义保持 0=时/1=分/2=秒/3=AP，pick 与 selIdx 不受影响）。
	cols, panelW := t.buildColumns(ctx, th)
	panel.Add(cols)

	// 底部："现在/确定"（Button 文本按钮组合复用）。
	panel.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetSize2(panelW, 1))
	footer := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	footer.SetPadding(geom.InLTRB(th.Space.SM, th.Space.XS, th.Space.SM, th.Space.XS))
	footer.SetWidth(panelW)
	sp := scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight)
	footer.Add(sp)
	btnNow := &Button{Text: "现在", Type: TypeText, Size: ui.SizeSmall, OnClick: func() {
		now := time.Now()
		t.Value = now
		t.highlightAll(now.Hour(), now.Minute(), now.Second())
		t.MarkDirty() // 触发框文本随根重建刷新
		if t.OnChange != nil {
			t.OnChange(now)
		}
	}}
	// footer 按钮包被动锚：Button 自带 hit label（"button"）会撞全页
	// 按钮——审计锚唯一化（DatePicker 锚撞车同款处置）。
	nowWrap := scene.Flex(scene.Horizontal, 0)
	nowWrap.Add(btnNow.Build(ctx))
	nowWrap.SetHit(scene.HitSpec{Label: "time-now"})
	footer.Add(nowWrap)
	btnOk := &Button{Text: "确定", Type: TypeLink, Size: ui.SizeSmall, OnClick: func() {
		// ctx.Overlay 直关（DatePicker 结论：fb.host 链路可能为 nil）。
		if ctx.Overlay != nil && t.handle != 0 {
			ctx.Overlay.Close(t.handle)
		}
		t.handle = 0
		t.panelOpen = false
	}}
	okWrap := scene.Flex(scene.Horizontal, 0)
	okWrap.Add(btnOk.Build(ctx))
	okWrap.SetHit(scene.HitSpec{Label: "time-ok"})
	footer.Add(okWrap)
	panel.Add(footer)
	return panel
}

// disabledSet 返回某列（colIdx）当前禁用的值集合。hour/minute 取当前
// 选中值（无选中取 0），供 disabledMinutes/Seconds 的联动过滤。
func (t *TimePicker) disabledSet(colIdx int) map[int]bool {
	if t.DisabledTime == nil {
		return nil
	}
	sel := t.selectedParts()
	hh, mm := 0, 0
	if sel[0] >= 0 {
		hh = sel[0]
	}
	if sel[1] >= 0 {
		mm = sel[1]
	}
	spec := t.DisabledTime(t.nowValue())
	var vals []int
	switch colIdx {
	case 0:
		if spec.DisabledHours != nil {
			vals = spec.DisabledHours()
		}
	case 1:
		if spec.DisabledMinutes != nil {
			vals = spec.DisabledMinutes(hh)
		}
	case 2:
		if spec.DisabledSeconds != nil {
			vals = spec.DisabledSeconds(hh, mm)
		}
	}
	if len(vals) == 0 {
		return nil
	}
	set := make(map[int]bool, len(vals))
	for _, v := range vals {
		set[v] = true
	}
	return set
}

// nowValue 当前值（无值取今天 0 点，供 DisabledTime 参数）。
func (t *TimePicker) nowValue() time.Time {
	if !t.Value.IsZero() {
		return t.Value
	}
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

// buildColumn 构建一列的 cell 纵向列表（vals 为列值集合，28px 高）。
// 12h 小时列（colIdx 0 且 Use12Hours）：显示 %02d（12,01..11）、值存
// 12h 表示；上午/下午列显示中文、值 0=上午/1=下午。pick 负责还原 24h。
func (t *TimePicker) buildColumn(ctx *ui.BuildContext, th *theme.Data, colIdx int, vals []int, cellLabel string) *scene.Node {
	sel := t.selectedParts()
	col := scene.Flex(scene.Vertical, 0)
	t.cellsByCol[colIdx] = make([]*timeCellInfo, 0, len(vals))
	isHour12 := colIdx == 0 && t.Use12Hours
	isAP := colIdx == 3
	disabled := t.disabledSet(colIdx)
	for idx, v := range vals {
		idx, v := idx, v
		src := fmt.Sprintf("%02d", v)
		if isAP {
			if v == 0 {
				src = "上午"
			} else {
				src = "下午"
			}
		}
		// 12h 小时列显示值 12,01..11；disabledSet 用 24h 语义，这里换算。
		disVal := v
		if isHour12 {
			if v == 12 {
				disVal = 0
			} else {
				disVal = v
			}
		}
		cellDisabled := disabled != nil && disabled[disVal]
		l := text.LayoutText(src, text.DefaultFont(th.Font.SizeBase), text.Options{})
		info := &timeCellInfo{src: src, value: v, layout: l}
		txtColor := th.Color.Text
		if cellDisabled {
			txtColor = th.Color.TextFaint
		}
		info.txt = scene.Text(scene.TextSpec{Src: src, Layout: l, Color: txtColor})
		info.txt.SetMargin(geom.InLTRB((timeColW-l.Size.W)/2, (timeCellH-l.Size.H)/2, 0, 0))
		cell := scene.Box(scene.BoxSpec{})
		cell.SetSize2(timeColW, timeCellH)
		cell.Add(info.txt)
		if cellDisabled {
			// 禁用格：独立锚（审计断言用）+ 吞掉点击（不选中、也不让
			// down 冒泡触发浮层 OutsideClick 关闭——antd 禁用格点击无效）。
			cell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "time-cell-disabled"})
			if ctx.Input != nil {
				ctx.Input.Set(cell, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
					return ev.Down()
				}})
			}
			info.cell = cell
			t.cellsByCol[colIdx] = append(t.cellsByCol[colIdx], info)
			col.Add(cell)
			continue
		}
		cell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: cellLabel, Cursor: scene.CursorPointer})
		info.cell = cell
		t.cellsByCol[colIdx] = append(t.cellsByCol[colIdx], info)
		// 选中判定：时列 24h 直等；12h 换算（12→0 点、1..11→13..23 侧依
		// AM/PM）；上午/下午列按 selIdx[3]（buildPanel 已填）。
		selected := false
		if isAP {
			selected = t.selIdx[3] == idx
		} else if isHour12 {
			if sel[0] >= 0 {
				h12 := sel[0] % 12
				if h12 == 0 {
					h12 = 12
				}
				selected = h12 == v && t.selIdx[3] == hourToAP(sel[0])
			}
		} else {
			selected = sel[colIdx] == v
		}
		if selected {
			t.selIdx[colIdx] = idx
			t.paintSelected(info, true)
		}
		if ctx.Input != nil {
			ctx.Input.Set(cell, input.Handler{
				Enter: func(*input.PointerEvent) {
					if t.selIdx[colIdx] != idx { // 选中项不叠加 hover 底
						info.cell.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.FillHover)})
						info.cell.MarkNeedsPaint()
					}
				},
				Leave: func(*input.PointerEvent) {
					if t.selIdx[colIdx] != idx {
						info.cell.SetBox(scene.BoxSpec{})
						info.cell.MarkNeedsPaint()
					}
				},
				Pointer: func(ev *input.PointerEvent) bool {
					if !ev.Down() {
						return false
					}
					// down 即执行（浮层内 tap 配对不可靠结论复用）。
					t.pick(colIdx, idx)
					return true
				},
			})
		}
		col.Add(cell)
	}
	return col
}

// hourToAP 24h 小时的上午/下午侧（0=上午 1=下午）。
func hourToAP(h int) int {
	if h >= 12 {
		return 1
	}
	return 0
}

// pick 选中某列某值：高亮原地迁移 + 合成完整时刻。12h 模式：小时列值
// （12,1..11）+上午/下午列还原 24h；分/秒列原值。
func (t *TimePicker) pick(colIdx, idx int) {
	if t.selIdx[colIdx] >= 0 {
		t.paintSelected(t.cellsByCol[colIdx][t.selIdx[colIdx]], false)
	}
	t.selIdx[colIdx] = idx
	info := t.cellsByCol[colIdx][idx]
	v := info.value
	t.paintSelected(info, true)
	// 合成完整时刻：未选列保持原值（无值列取 0）。
	parts := t.selectedParts()
	if t.Value.IsZero() {
		parts = [3]int{0, 0, 0}
	}
	if t.Use12Hours {
		// 由列集合推导 24h 小时：h12 + 上午/下午侧。
		if colIdx == 0 {
			h12 := v
			ap := 0
			if t.selIdx[3] >= 0 {
				ap = t.selIdx[3]
			} else {
				ap = hourToAP(parts[0])
			}
			parts[0] = ap12To24(h12, ap)
		} else if colIdx == 3 {
			h12 := parts[0] % 12
			if h12 == 0 {
				h12 = 12
			}
			parts[0] = ap12To24(h12, v)
		} else {
			parts[colIdx] = v
		}
	} else {
		parts[colIdx] = v
	}
	nv := time.Date(0, 1, 1, parts[0], parts[1], parts[2], 0, time.Local)
	t.Value = nv
	t.MarkDirty() // 触发框文本随根重建刷新
	if t.OnChange != nil {
		t.OnChange(nv)
	}
	if t.OnValueChange != nil {
		t.OnValueChange()
	}
}

// ap12To24 12h 值（12,1..11）+上午/下午侧 → 24h 小时。
func ap12To24(h12, ap int) int {
	if ap == 1 { // 下午
		if h12 == 12 {
			return 12
		}
		return h12 + 12
	}
	// 上午
	if h12 == 12 {
		return 0
	}
	return h12
}

// paintSelected 原地切换一格的选中样式（主色实底+白字 / 常态）。
func (t *TimePicker) paintSelected(info *timeCellInfo, sel bool) {
	th := t.curTh
	spec := scene.BoxSpec{}
	color := th.Color.Text
	if sel {
		spec = scene.BoxSpec{Fill: scene.Solid(th.Color.Primary)}
		color = th.Color.OnPrimary
	}
	info.cell.SetBox(spec)
	info.txt.SetTextPaint(scene.TextSpec{Src: info.src, Layout: info.layout, Color: color})
	info.cell.MarkNeedsPaint()
}

// highlightAll "现在"：列高亮整体迁移到当前时刻（原地，零重建）。
func (t *TimePicker) highlightAll(h, m, s int) {
	vals := [3]int{h, m, s}
	for i := 0; i < 3; i++ {
		if t.selIdx[i] >= 0 {
			t.paintSelected(t.cellsByCol[i][t.selIdx[i]], false)
		}
		target := vals[i]
		if t.Use12Hours && i == 0 {
			target = h % 12
			if target == 0 {
				target = 12
			}
		}
		idx := -1
		for j, info := range t.cellsByCol[i] {
			if info.value == target {
				idx = j
				break
			}
		}
		if idx < 0 {
			idx = len(t.cellsByCol[i]) - 1
		}
		t.selIdx[i] = idx
		t.paintSelected(t.cellsByCol[i][idx], true)
	}
	if t.Use12Hours {
		if t.selIdx[3] >= 0 {
			t.paintSelected(t.cellsByCol[3][t.selIdx[3]], false)
		}
		ap := hourToAP(h)
		t.selIdx[3] = ap
		t.paintSelected(t.cellsByCol[3][ap], true)
	}
}
