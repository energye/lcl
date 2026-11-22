package widgets

// ── Timeline 时间轴（M12 长尾）───────────────────────────────────
//
// antd：<Timeline items={[{content, color}]}> 垂直时间轴：
// 左轴列（圆点+竖线贯穿）+ 右内容列（文本）。节点色：blue（默认
// 主色）/green/gred/gray（antd color 语义）。第一轮内容为文本
// （Time+Text 合并行）；alternate 交替/右轴/自定义 dot 未实现
// （如实标注）。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// TimelineNodeColor 节点色（antd color 语义子集）。
const (
	TimelineBlue  = "blue"
	TimelineGreen = "green"
	TimelineRed   = "red"
	TimelineGray  = "gray"
)

// TimelineItem 是时间轴项。
type TimelineItem struct {
	Time  string // 小字时间（如 2026-11-21 10:24）
	Text  string // 主文本
	Color string // 节点色（TimelineBlue/Green/Red/Gray）
	// Widget 主文本 Widget（antd content 任意节点；非空时优先于
	// Text——深化轮）。
	Widget ui.Widget
	// Dot 自定义节点（antd dot 任意节点——10px 圆点被替换为该
	// Widget 原样渲染，轴断线仍按默认 dot 高度占位）。非空时优先于
	// Color 圆点。
	Dot ui.Widget
}

// TimelineMode 是时间轴方向（antd mode 语义）。
type TimelineMode uint8

const (
	// TimelineLeft 左轴右内容（默认）。
	TimelineLeft TimelineMode = iota
	// TimelineAlternate 交替：内容奇偶行左右互换（antd mode="alternate"）。
	TimelineAlternate
	// TimelineRight 右轴左内容（antd mode="right"——全部内容在左）。
	TimelineRight
)

// Timeline 是垂直时间轴。
type Timeline struct {
	Items []TimelineItem
	Width float32 // 内容宽（默认 300）
	// Mode 方向（默认左轴）。
	Mode TimelineMode
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
		axisX    = float32(9) // 轴中心 x（dot 10px 直径）
		lineW    = float32(2)
		baseRowH = float32(56)
		dotY     = float32(6)
	)
	// 内容列宽：alternate/right 模式轴居中，左右各留半宽。
	halfW := (w - 20) / 2
	for i, it := range t.Items {
		it := it
		rowH := baseRowH
		if i == len(t.Items)-1 {
			rowH = 40 // 尾行无下延线
		}
		// 本行内容侧：left=恒右；right=恒左；alternate=奇偶互换（首行右）。
		sideRight := true
		if t.Mode == TimelineRight {
			sideRight = false
		} else if t.Mode == TimelineAlternate {
			sideRight = i%2 == 0
		}
		// 轴 x：left 模式轴在左缘（axisX=9）；alternate/right 轴居中。
		ax := axisX
		if t.Mode != TimelineLeft {
			ax = w / 2
		}
		// alternate/right 模式 row 用 **Stack**（轴+内容全部绝对定位——
		// Flex 子节点 SetStackPosition 不生效，实测踩坑）；left 保持
		// Flex（轴窄列+内容 Flex 排布）。
		var row *scene.Node
		if t.Mode != TimelineLeft {
			row = scene.Stack().SetSize2(w, rowH)
		} else {
			row = scene.Flex(scene.Horizontal, 0)
			row.SetWidth(w).SetHeight(rowH)
		}

		// 轴列：竖线（上段+下段）+ 圆点（行内绝对定位）。alternate/
		// right 时轴层横跨全宽（Stack 内即全宽容器）；left 时窄列。
		var axis *scene.Node
		if t.Mode != TimelineLeft {
			axis = scene.Box(scene.BoxSpec{}).SetSize2(w, rowH)
		} else {
			axis = scene.Box(scene.BoxSpec{}).SetSize2(20, rowH)
		}
		lineColor := th.Color.Border
		if i > 0 { // 非首行：上段线
			upLen := dotY
			if i == 0 {
				upLen = 0
			}
			up := scene.Box(scene.BoxSpec{Fill: scene.Solid(lineColor)}).SetSize2(lineW, upLen)
			up.SetPositionKeepSize(true)
			up.SetStackPosition(geom.Pt(ax-lineW/2, 0))
			axis.Add(up)
		}
		if i < len(t.Items)-1 { // 非尾行：下段线
			down := scene.Box(scene.BoxSpec{Fill: scene.Solid(lineColor)}).SetSize2(lineW, rowH-dotY-10)
			down.SetPositionKeepSize(true)
			down.SetStackPosition(geom.Pt(ax-lineW/2, dotY+10))
			axis.Add(down)
		}
		// 节点：Dot Widget 非空时原样渲染（自定义 dot），否则默认圆点。
		if it.Dot != nil {
			dn := it.Dot.Build(ctx)
			dn.SetPositionKeepSize(true)
			dn.SetStackPosition(geom.Pt(ax-10, dotY-1))
			dn.SetHit(scene.HitSpec{Label: "timeline-dot"})
			axis.Add(dn)
		} else {
			dot := scene.Box(scene.BoxSpec{
				Fill:   scene.Solid(t.dotColor(th, it.Color)),
				Radius: 5,
			}).SetSize2(10, 10)
			dot.SetPositionKeepSize(true)
			dot.SetStackPosition(geom.Pt(ax-5, dotY))
			dot.SetHit(scene.HitSpec{Label: "timeline-dot"})
			axis.Add(dot)
		}
		row.Add(axis)

		// 内容列：时间小字 + 主文本；alternate/right 时内容在另一侧。
		ccol := scene.Flex(scene.Vertical, 2)
		ccolW := w - 20
		if t.Mode != TimelineLeft {
			ccolW = halfW
		}
		ccol.SetWidth(ccolW)
		if it.Time != "" {
			tl := text.LayoutText(it.Time, text.DefaultFont(th.Font.SizeXS), text.Options{})
			tv := scene.Text(scene.TextSpec{Src: it.Time, Layout: tl, Color: th.Color.TextFaint})
			tv.SetMargin(geom.InLTRB(8, dotY-2, 0, 0))
			tv.SetHit(scene.HitSpec{Label: "timeline-time"})
			ccol.Add(tv)
		}
		if it.Widget != nil {
			// Widget 主内容（antd content 任意节点——深化轮）。
			if wn := it.Widget.Build(ctx); wn != nil {
				wn.SetMargin(geom.InLTRB(8, 0, 0, 0))
				wn.SetHit(scene.HitSpec{Label: "timeline-item"})
				ccol.Add(wn)
			}
		} else {
			ml := text.LayoutText(it.Text, text.DefaultFont(th.Font.SizeBase), text.Options{})
			mv := scene.Text(scene.TextSpec{Src: it.Text, Layout: ml, Color: th.Color.Text})
			mv.SetMargin(geom.InLTRB(8, 0, 0, 0))
			mv.SetHit(scene.HitSpec{Label: "timeline-item"})
			ccol.Add(mv)
		}
		// alternate/right 模式：Stack 布局下内容绝对定位贴对应侧
		//（sideRight=true 内容右列 x=ax+14；false 内容左列 x=0——
		// antd mode=right 内容左列文字左对齐）。
		if t.Mode != TimelineLeft {
			ccol.SetPositionKeepSize(true)
			if sideRight {
				ccol.SetStackPosition(geom.Pt(ax+14, 0))
			} else {
				ccol.SetStackPosition(geom.Pt(0, 0))
			}
		}
		row.Add(ccol)
		col.Add(row)

	}
	return col
}
