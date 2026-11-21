package widgets

// ── Timeline 时间轴（M12 长尾）───────────────────────────────────
//
// antd：<Timeline items={[{content, color}]}> 垂直时间轴：
// 左轴列（圆点+竖线贯穿）+ 右内容列（文本）。节点色：blue（默认
// 主色）/green/gred/gray（antd color 语义）。第一轮内容为文本
// （Time+Text 合并行）；alternate 交替/右轴/自定义 dot 未实现
// （如实标注）。

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// TimelineNodeColor 节点色（antd color 语义子集）。
const (
	TimelineBlue = "blue"
	TimelineGreen = "green"
	TimelineRed   = "red"
	TimelineGray  = "gray"
)

// TimelineItem 是时间轴项。
type TimelineItem struct {
	Time  string // 小字时间（如 2026-11-21 10:24）
	Text  string // 主文本
	Color string // 节点色（TimelineBlue/Green/Red/Gray）
}

// Timeline 是垂直时间轴。
type Timeline struct {
	Items []TimelineItem
	Width float32 // 内容宽（默认 300）
}

// NewTimeline 创建时间轴。
func NewTimeline(items []TimelineItem) *Timeline {
	return &Timeline{Items: items, Width: 300}
}

// dotColor 节点色映射。
func (t *Timeline) dotColor(th *theme.Data, color string) geom.Color {
	switch color {
	case TimelineGreen:
		return th.Color.Success
	case TimelineRed:
		return th.Color.Danger
	case TimelineGray:
		return th.Color.TextFaint
	default: // blue
		return th.Color.Primary
	}
}

// Build 实现 ui.Widget。
func (t *Timeline) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	w := t.Width
	if w <= 0 {
		w = 300
	}

	col := scene.Flex(scene.Vertical, 0)
	col.SetWidth(w)
	col.SetHit(scene.HitSpec{Label: "timeline"})

	const (
		axisX   = float32(9)  // 轴中心 x（dot 10px 直径）
		lineW   = float32(2)
		rowH    = float32(56)
		dotY    = float32(6)
	)
	for i, it := range t.Items {
		it := it
		rowH := rowH
		if i == len(t.Items)-1 {
			rowH = 40 // 尾行无下延线
		}
		row := scene.Flex(scene.Horizontal, 0)
		row.SetWidth(w).SetHeight(rowH)

		// 轴列：竖线（上段+下段）+ 圆点（行内绝对定位）。
		axis := scene.Box(scene.BoxSpec{}).SetSize2(20, rowH)
		lineColor := th.Color.Border
		if i > 0 { // 非首行：上段线
			upLen := dotY
			if i == 0 {
				upLen = 0
			}
			up := scene.Box(scene.BoxSpec{Fill: scene.Solid(lineColor)}).SetSize2(lineW, upLen)
			up.SetPositionKeepSize(true)
			up.SetStackPosition(geom.Pt(axisX-lineW/2, 0))
			axis.Add(up)
		}
		if i < len(t.Items)-1 { // 非尾行：下段线
			down := scene.Box(scene.BoxSpec{Fill: scene.Solid(lineColor)}).SetSize2(lineW, rowH-dotY-10)
			down.SetPositionKeepSize(true)
			down.SetStackPosition(geom.Pt(axisX-lineW/2, dotY+10))
			axis.Add(down)
		}
		dot := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(t.dotColor(th, it.Color)),
			Radius: 5,
		}).SetSize2(10, 10)
		dot.SetPositionKeepSize(true)
		dot.SetStackPosition(geom.Pt(axisX-5, dotY))
		dot.SetHit(scene.HitSpec{Label: "timeline-dot"})
		axis.Add(dot)
		row.Add(axis)

		// 内容列：时间小字 + 主文本。
		ccol := scene.Flex(scene.Vertical, 2)
		ccol.SetWidth(w - 20)
		if it.Time != "" {
			tl := text.LayoutText(it.Time, text.DefaultFont(th.Font.SizeXS), text.Options{})
			tv := scene.Text(scene.TextSpec{Src: it.Time, Layout: tl, Color: th.Color.TextFaint})
			tv.SetMargin(geom.InLTRB(8, dotY-2, 0, 0))
			tv.SetHit(scene.HitSpec{Label: "timeline-time"})
			ccol.Add(tv)
		}
		ml := text.LayoutText(it.Text, text.DefaultFont(th.Font.SizeBase), text.Options{})
		mv := scene.Text(scene.TextSpec{Src: it.Text, Layout: ml, Color: th.Color.Text})
		mv.SetMargin(geom.InLTRB(8, 0, 0, 0))
		mv.SetHit(scene.HitSpec{Label: "timeline-item"})
		ccol.Add(mv)
		row.Add(ccol)
		col.Add(row)

		_ = fmt.Sprintf // fmt 保留（后续轮内容格式化）
	}
	return col
}
