package widgets

// ── Grid 响应式栅格（antd Row/Col）──────────────────────────────────
//
// 24 栅格布局：每个 GridCol 占 Span 列（默认按内容等分），Offset 左偏移；
// 响应式断点（antd）：xs<576 / sm≥576 / md≥768 / lg≥992 / xl≥1200。
// 断点字段为指针（nil = 继承更大断点，antd Col 语义）。
//
// 实现：逐行 Flex（每行最多 24 栅格宽），每列宽 = (span/24)×行宽 - 间距。
// 视口宽经 ViewportSize() 读取并登记依赖（resize 时容器重建本组件）。
//
// 未实现（如实标注）：gutter 的横纵分离、Flex 模式（justify/align）、
// useBreakpoint 钩子。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// GridCol 是一列（antd Col）。
type GridCol struct {
	Child ui.Widget
	// Span 占用的 24 栅格数（默认 24 = 整行；0 也按 24）。
	Span int
	// Offset 左侧空出的栅格数。
	Offset int
	// 响应式断点 span（nil = 不覆盖）。按 antd 升序取"当前及以下最大命中"。
	Xs, Sm, Md, Lg, Xl *int
}

// Grid 是 24 栅格响应式布局（antd Row + Col）。
type Grid struct {
	Cols []GridCol
	// Gap 列间距（DIP，默认 16）。
	Gap float32
}

// NewGrid 创建栅格。
func NewGrid(cols ...GridCol) *Grid { return &Grid{Cols: cols} }

// breakpointSpan 按视口宽解析某列的 span 与 offset（antd 断点语义：
// 取"当前宽度命中的最大断点"的 span；该断点未设则向下回退）。
func breakpointSpan(c GridCol, vw float32) (span, offset int) {
	span = c.Span
	if span <= 0 {
		span = 24
	}
	offset = c.Offset
	// 从大到小找第一个命中且非 nil 的 span 覆盖。
	type bp struct {
		min float32
		sp  *int
	}
	bps := []bp{{1200, c.Xl}, {992, c.Lg}, {768, c.Md}, {576, c.Sm}, {0, c.Xs}}
	for _, b := range bps {
		if vw >= b.min && b.sp != nil {
			span = *b.sp
			break
		}
	}
	if span <= 0 {
		span = 24
	}
	if span > 24 {
		span = 24
	}
	return span, offset
}

// Build 实现 ui.Widget。
func (g *Grid) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	vp := ctx.ViewportSize() // 登记视口依赖（resize 重排）
	vw := vp.W
	if vw <= 0 {
		vw = 1024
	}
	gap := g.Gap
	if gap <= 0 {
		gap = th.Space.XL
	}
	// 行宽 = 视口宽 - 版心边距（与 doc.Page 的 40×2 一致；无版心时全宽）。
	rowW := vw - 80
	if rowW <= 0 {
		rowW = vw
	}
	// 单栅格宽（含间距折算：n 列有 n-1 个 gap）。
	unit := (rowW + gap) / 24

	col := scene.Flex(scene.Vertical, gap)
	col.SetWidth(rowW)
	col.SetHit(scene.HitSpec{Label: "grid"})

	var row *scene.Node
	used := 0
	flush := func() {
		if row != nil {
			col.Add(row)
			row = nil
			used = 0
		}
	}
	for _, c := range g.Cols {
		span, offset := breakpointSpan(c, vw)
		// 偏移 + 占位超出行宽则换行。
		if used+offset+span > 24 && used > 0 {
			flush()
		}
		if row == nil {
			row = scene.Flex(scene.Horizontal, gap)
			row.SetWidth(rowW)
		}
		if offset > 0 {
			row.Add(scene.Box(scene.BoxSpec{}).SetSize2(float32(offset)*unit, 1))
			used += offset
		}
		w := float32(span)*unit - gap
		if w < 0 {
			w = 0
		}
		cell := scene.Box(scene.BoxSpec{})
		cell.SetWidth(w)
		if c.Child != nil {
			if n := c.Child.Build(ctx); n != nil {
				cell.Add(n)
			}
		}
		cell.SetHit(scene.HitSpec{Label: "grid-col"})
		row.Add(cell)
		used += span
		if used >= 24 {
			flush()
		}
	}
	flush()
	return col
}

var _ = geom.Pt
