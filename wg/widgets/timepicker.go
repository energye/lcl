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
// 未实现（如实标注）：use12Hours / hourStep·minuteStep·secondStep /
// disabledTime 系列 / hideDisabledOptions / allowClear 清除图标 /
// format 列裁剪（列固定三列）/ RangePicker / 受控 open / placement
// 四向 / changeOnScroll / 无值时滚动定位到当前时刻。

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

// 面板尺寸令牌（antd design token：列 56 宽、cell 28 高、视口 224）。
const (
	timeColW      = float32(56)
	timeCellH     = float32(28)
	timeViewportH = float32(224)
)

// timeCellInfo 一格的节点与排版（原地改色所需）。
type timeCellInfo struct {
	cell   *scene.Node
	txt    *scene.Node
	src    string
	layout *text.Layout
}

// TimePicker 是时间选择器（时/分/秒三列面板）。
type TimePicker struct {
	controlBase
	Value       time.Time
	Placeholder string
	Width       float32
	OnChange    func(t time.Time)

	handle    overlay.Handle
	fb        feedbackHost
	childNode *scene.Node
	panelOpen bool
	curTh     *theme.Data

	// 面板打开期间的 cell 表与选中下标（选中/hover 原地改，零重建）。
	cellsByCol [3][]*timeCellInfo
	selIdx     [3]int // -1 = 该列未选中
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
	display := t.Value.Format("15:04:05")
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

// selectedParts 当前选中的时/分/秒（无值 {-1,-1,-1}）。
func (t *TimePicker) selectedParts() [3]int {
	if t.Value.IsZero() {
		return [3]int{-1, -1, -1}
	}
	return [3]int{t.Value.Hour(), t.Value.Minute(), t.Value.Second()}
}

// openPanel 打开三列面板（BandPopup+PlacementBottomStart，Select 同通道）。
func (t *TimePicker) openPanel(ctx *ui.BuildContext, anchor *scene.Node, th *theme.Data) {
	if t.handle != 0 {
		return
	}
	t.curTh = th
	t.panelOpen = true
	t.selIdx = t.selectedParts()
	panel := t.buildPanel(ctx, th)
	t.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: overlay.PlacementBottomStart,
		Size:      geom.Size{W: timeColW*3 + 2, H: 0}, // 高自适应（三列+footer）
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

// buildPanel 构建面板：三列（时/分/秒）+列分隔+底部分隔线+"现在/确定"。
func (t *TimePicker) buildPanel(ctx *ui.BuildContext, th *theme.Data) *scene.Node {
	panel := scene.Flex(scene.Vertical, 0)
	panel.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	panel.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "time-panel"})

	// 三列：Scroll 组合复用（Select 选项列表同款视口）；列间 1px 分隔。
	cols := scene.Flex(scene.Horizontal, 0)
	cellLabels := []string{"time-cell-h", "time-cell-m", "time-cell-s"}
	counts := []int{24, 60, 60}
	sel := t.selectedParts()
	for i := 0; i < 3; i++ {
		col := t.buildColumn(ctx, th, i, counts[i], cellLabels[i])
		scroll := &Scroll{Axis: scene.Vertical, Width: timeColW, Height: timeViewportH,
			Child: nodeFn(func(ctx *ui.BuildContext) *scene.Node { return col })}
		// 选中项尽量居中：offset = idx*cellH - (viewport-cellH)/2。
		if sel[i] >= 0 {
			off := float32(sel[i])*timeCellH - (timeViewportH-timeCellH)/2
			if off < 0 {
				off = 0
			}
			scroll.Offset = off
		}
		wrap := scene.Flex(scene.Vertical, 0)
		wrap.Add(scroll.Build(ctx))
		wrap.SetHit(scene.HitSpec{Label: fmt.Sprintf("time-col-%d", i)})
		cols.Add(wrap)
		if i < 2 {
			cols.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetSize2(1, timeViewportH))
		}
	}
	panel.Add(cols)

	// 底部："现在/确定"（Button 文本按钮组合复用）。
	panel.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetSize2(timeColW*3+2, 1))
	footer := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	footer.SetPadding(geom.InLTRB(th.Space.SM, th.Space.XS, th.Space.SM, th.Space.XS))
	footer.SetWidth(timeColW*3 + 2)
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

// buildColumn 构建一列的 cell 纵向列表（24/60/60 个，28px 高）。
func (t *TimePicker) buildColumn(ctx *ui.BuildContext, th *theme.Data, colIdx, count int, cellLabel string) *scene.Node {
	sel := t.selectedParts()
	col := scene.Flex(scene.Vertical, 0)
	t.cellsByCol[colIdx] = make([]*timeCellInfo, 0, count)
	for v := 0; v < count; v++ {
		v := v
		src := fmt.Sprintf("%02d", v)
		l := text.LayoutText(src, text.DefaultFont(th.Font.SizeSM), text.Options{})
		info := &timeCellInfo{src: src, layout: l}
		info.txt = scene.Text(scene.TextSpec{Src: src, Layout: l, Color: th.Color.Text})
		info.txt.SetMargin(geom.InLTRB((timeColW-l.Size.W)/2, (timeCellH-l.Size.H)/2, 0, 0))
		cell := scene.Box(scene.BoxSpec{})
		cell.SetSize2(timeColW, timeCellH)
		cell.Add(info.txt)
		cell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: cellLabel, Cursor: scene.CursorPointer})
		info.cell = cell
		t.cellsByCol[colIdx] = append(t.cellsByCol[colIdx], info)
		if sel[colIdx] == v {
			t.selIdx[colIdx] = v
			t.paintSelected(info, true)
		}
		if ctx.Input != nil {
			ctx.Input.Set(cell, input.Handler{
				Enter: func(*input.PointerEvent) {
					if t.selIdx[colIdx] != v { // 选中项不叠加 hover 底
						info.cell.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.FillHover)})
						info.cell.MarkNeedsPaint()
					}
				},
				Leave: func(*input.PointerEvent) {
					if t.selIdx[colIdx] != v {
						info.cell.SetBox(scene.BoxSpec{})
						info.cell.MarkNeedsPaint()
					}
				},
				Pointer: func(ev *input.PointerEvent) bool {
					if !ev.Down() {
						return false
					}
					// down 即执行（浮层内 tap 配对不可靠结论复用）。
					t.pick(colIdx, v)
					return true
				},
			})
		}
		col.Add(cell)
	}
	return col
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

// pick 选中某列某值：高亮原地迁移 + 合成完整时刻。
func (t *TimePicker) pick(colIdx, v int) {
	if t.selIdx[colIdx] >= 0 {
		t.paintSelected(t.cellsByCol[colIdx][t.selIdx[colIdx]], false)
	}
	t.selIdx[colIdx] = v
	t.paintSelected(t.cellsByCol[colIdx][v], true)
	// 合成完整时刻：未选列保持原值（无值列取 0）。
	parts := t.selectedParts()
	if t.Value.IsZero() {
		parts = [3]int{0, 0, 0}
	}
	parts[colIdx] = v
	nv := time.Date(0, 1, 1, parts[0], parts[1], parts[2], 0, time.Local)
	t.Value = nv
	t.MarkDirty() // 触发框文本随根重建刷新
	if t.OnChange != nil {
		t.OnChange(nv)
	}
}

// highlightAll "现在"：三列高亮整体迁移到当前时刻（原地，零重建）。
func (t *TimePicker) highlightAll(h, m, s int) {
	vals := [3]int{h, m, s}
	for i := 0; i < 3; i++ {
		if t.selIdx[i] >= 0 {
			t.paintSelected(t.cellsByCol[i][t.selIdx[i]], false)
		}
		t.selIdx[i] = vals[i]
		t.paintSelected(t.cellsByCol[i][vals[i]], true)
	}
}
