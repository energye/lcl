package widgets

import (
	"sort"
	"strconv"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// ── Table：虚拟表格 + 排序 + 三档列宽 ─────────────────────────────
//
// Q9 定案：列宽三档。
//   - ColumnAuto：取该列所有单元格文本 intrinsic 最大宽，受 Min/Max 约束；
//   - ColumnFixed：固定像素；
//   - ColumnFlex：按权重分配扣除 auto/fixed 后的剩余宽度。
//
// 行虚拟化：只构建可见区 ± Overscan 行的节点，滚动偏移是唯一真相。
// 行按 vi 位置复用场景节点（每帧重建但 Key 稳定），避免 10 万行撑爆节点树。

// ColumnWidthMode 是列宽模式（Q9）。
type ColumnWidthMode uint8

const (
	// ColumnAuto 按内容最大宽。
	ColumnAuto ColumnWidthMode = iota
	// ColumnFixed 固定像素宽。
	ColumnFixed
	// ColumnFlex 按权重分配剩余。
	ColumnFlex
)

// TableColumn 是表格列定义。
type TableColumn struct {
	Title    string
	Mode     ColumnWidthMode
	Width    float32 // fixed 时的像素宽；auto 时的初始宽
	Flex     float32 // flex 权重
	Min, Max float32 // auto/flex 的上下限（0 表示不限）
	Sortable bool
	Align    scene.Alignment
	// Formatter 自定义单元格文本；返回空串时回退到默认格式化。
	Formatter func(row int, value any) string
}

// Table 是虚拟表格。Rows 为行数据；每行是 []any，与 Columns 对齐。
// SortColumn >= 0 时按该列排序；Asc 决定方向（稳定排序）。
type Table struct {
	controlBase
	Columns    []TableColumn
	Rows       [][]any
	Width      float32
	Height     float32 // 视口高；<=0 时用 DefaultTableHeight
	RowHeight  float32
	HeaderH    float32
	Overscan   int
	SortColumn int
	Asc        bool
	Selected   int // 选中行索引（数据索引，-1 无）
	OnSort     func(col int, asc bool)
	OnSelect   func(row int)

	// ShowScrollbar：nil = 内容溢出自动显示；true 始终；false 隐藏（仍可滚轮）。
	ShowScrollbar *bool

	offset    float32
	sortCache []int
	sortDirty bool
	sorts     []*ui.GestureDetector // 跨 Build 存活
	rowTaps   []*ui.GestureDetector
	// 虚拟窗口与滚动条（原地滚动用）
	rowNodes   []*scene.Node
	rowVis     []int
	builtFirst int
	builtLast  int
	bar        *Scrollbar
	reqFrame   func()
}

// DefaultTableHeight 是未设置 Height 时的视口高。
const DefaultTableHeight float32 = 240

// NewTable 创建表格。
func NewTable(cols []TableColumn, rows [][]any) *Table {
	return &Table{
		Columns: cols, Rows: rows, Width: 640, RowHeight: 30, HeaderH: 32,
		Overscan: 2, SortColumn: -1, Selected: -1, sortDirty: true,
	}
}

// SetScrollOffset 设置垂直滚动偏移（虚拟化的唯一真相）。
//
// 可见窗口不变时只平移行节点（不重建，拖动/滚轮才顺滑）；
// 越过已构建的虚拟窗口才置脏重建。
func (t *Table) SetScrollOffset(v float32) {
	if v < 0 {
		v = 0
	}
	if max := t.maxOffset(); v > max {
		v = max
	}
	if t.offset == v {
		return
	}
	t.offset = v
	first, last := t.rowWindow()
	if t.rowNodes == nil || first < t.builtFirst || last > t.builtLast {
		t.MarkDirty()
		return
	}
	for i, n := range t.rowNodes {
		n.SetStackPosition(geom.Pt(0, t.HeaderH+float32(t.rowVis[i])*t.RowHeight-v))
	}
	if t.node != nil {
		t.node.MarkNeedsPaint()
	}
	if t.bar != nil {
		t.bar.SetValue(v)
	}
	if t.reqFrame != nil {
		t.reqFrame()
	}
}

// rowWindow 返回当前偏移下应构建的行窗口 [first,last)。
func (t *Table) rowWindow() (int, int) {
	vh := t.ViewportHeight()
	first := int(t.offset/t.RowHeight) - t.Overscan
	if first < 0 {
		first = 0
	}
	last := first + int(vh/t.RowHeight) + t.Overscan*2 + 2
	if n := len(t.Rows); last > n {
		last = n
	}
	return first, last
}

// ScrollOffset 返回当前垂直滚动偏移。
func (t *Table) ScrollOffset() float32 { return t.offset }

// ViewportHeight 返回视口高。
func (t *Table) ViewportHeight() float32 {
	if t.Height > 0 {
		return t.Height
	}
	return DefaultTableHeight
}

func (t *Table) maxOffset() float32 {
	total := t.HeaderH + float32(len(t.Rows))*t.RowHeight
	h := t.ViewportHeight()
	if total <= h {
		return 0
	}
	return total - h
}

// Sort 按列排序（稳定排序）。
func (t *Table) Sort(col int, asc bool) {
	if col < 0 || col >= len(t.Columns) {
		return
	}
	t.SortColumn, t.Asc = col, asc
	t.sortDirty = true
	t.MarkDirty()
	if t.OnSort != nil {
		t.OnSort(col, asc)
	}
}

func (t *Table) sortedIndex() []int {
	if !t.sortDirty && t.sortCache != nil && len(t.sortCache) == len(t.Rows) {
		return t.sortCache
	}
	idx := make([]int, len(t.Rows))
	for i := range idx {
		idx[i] = i
	}
	if t.SortColumn >= 0 && t.SortColumn < len(t.Columns) {
		col := t.SortColumn
		sort.SliceStable(idx, func(i, j int) bool {
			ri, rj := t.Rows[idx[i]], t.Rows[idx[j]]
			var a, b string
			if col < len(ri) {
				a = cellString(ri[col])
			}
			if col < len(rj) {
				b = cellString(rj[col])
			}
			if t.Asc {
				return a < b
			}
			return a > b
		})
	}
	t.sortCache, t.sortDirty = idx, false
	return idx
}

func cellString(v any) string {
	switch x := v.(type) {
	case nil:
		return ""
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case int64:
		return strconv.FormatInt(x, 10)
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return ""
	}
}

func (t *Table) cellText(col, row int) string {
	if col < 0 || col >= len(t.Columns) || row < 0 || row >= len(t.Rows) {
		return ""
	}
	c := t.Columns[col]
	if c.Formatter != nil {
		if s := c.Formatter(row, t.Rows[row][col]); s != "" {
			return s
		}
	}
	return cellString(t.Rows[row][col])
}

// columnWidths 按 Q9 三档计算列宽（像素数组）。
func (t *Table) columnWidths(fontSize float32) []float32 {
	n := len(t.Columns)
	if n == 0 {
		return nil
	}
	out := make([]float32, n)
	remaining := t.Width
	flexTotal := float32(0)
	font := text.DefaultFont(fontSize)
	for i, c := range t.Columns {
		switch c.Mode {
		case ColumnFixed:
			out[i] = c.Width
			remaining -= c.Width
		case ColumnFlex:
			flexTotal += c.Flex
		default: // ColumnAuto
			w := c.Width
			for r := range t.Rows {
				if lw := text.Measure(t.cellText(i, r), font).W; lw > w {
					w = lw
				}
			}
			if c.Min > 0 && w < c.Min {
				w = c.Min
			}
			if c.Max > 0 && w > c.Max {
				w = c.Max
			}
			out[i] = w + 12 // 左右内边距
			remaining -= out[i]
		}
	}
	if flexTotal > 0 && remaining > 0 {
		for i, c := range t.Columns {
			if c.Mode != ColumnFlex {
				continue
			}
			w := remaining * (c.Flex / flexTotal)
			if c.Min > 0 && w < c.Min {
				w = c.Min
			}
			if c.Max > 0 && w > c.Max {
				w = c.Max
			}
			out[i] = w
		}
	}
	return out
}

// VisibleRows 返回当前可见的数据行索引（供诊断/自动化断言）。
func (t *Table) VisibleRows() []int {
	idx := t.sortedIndex()
	vh := t.ViewportHeight()
	first := int(t.offset/t.RowHeight) - t.Overscan
	if first < 0 {
		first = 0
	}
	last := first + int(vh/t.RowHeight) + t.Overscan*2 + 2
	if last > len(idx) {
		last = len(idx)
	}
	if first > len(idx) {
		first = len(idx)
	}
	return idx[first:last]
}

// Build 实现 ui.Widget。
func (t *Table) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	widths := t.columnWidths(th.Font.SizeSM)
	idx := t.sortedIndex()

	col := scene.Flex(scene.Vertical, 0).SetWidth(t.Width).SetHeight(t.ViewportHeight())
	col.SetOverflow(true)

	// 表头
	header := scene.Flex(scene.Horizontal, 0).SetHeight(t.HeaderH)
	for ci, c := range t.Columns {
		w := widths[ci]
		cell := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.SurfaceSunken),
			Border: scene.Stroke(th.Color.Border, 1),
		}).SetWidth(w).SetHeight(t.HeaderH)
		label := c.Title
		if c.Sortable {
			if ci == t.SortColumn {
				label += map[bool]string{true: " ↑", false: " ↓"}[t.Asc]
			} else {
				label += " ↕"
			}
		}
		l := text.LayoutText(label, text.DefaultFont(th.Font.SizeSM), text.Options{})
		tv := scene.Text(scene.TextSpec{Src: label, Layout: l, Color: th.Color.TextMuted})
		tv.SetMargin(geom.InLTRB(6, (t.HeaderH-l.Size.H)/2, 0, 0))
		cell.Add(tv)
		if c.Sortable && ctx.Input != nil {
			cell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "table-sort", Cursor: scene.CursorPointer})
			cc, cAsc := ci, (ci == t.SortColumn && !t.Asc)
			g := &ui.GestureDetector{OnTap: func(geom.Point) { t.Sort(cc, cAsc) }}
			t.sorts = append(t.sorts, g)
			ctx.Input.Set(cell, g.Handler())
		}
		header.Add(cell)
	}
	col.Add(header)

	// 虚拟行：只构建可见区 ± Overscan
	vh := t.ViewportHeight()
	first := int(t.offset/t.RowHeight) - t.Overscan
	if first < 0 {
		first = 0
	}
	last := first + int(vh/t.RowHeight) + t.Overscan*2 + 2
	if last > len(idx) {
		last = len(idx)
	}
	t.rowTaps = t.rowTaps[:0]
	t.rowNodes = t.rowNodes[:0]
	t.rowVis = t.rowVis[:0]
	t.builtFirst, t.builtLast = first, last
	t.reqFrame = ctx.RequestFrame
	for vi := first; vi < last; vi++ {
		dataRow := idx[vi]
		row := scene.Flex(scene.Horizontal, 0).SetWidth(t.Width).SetHeight(t.RowHeight)
		bg := geom.Color{}
		if dataRow == t.Selected {
			bg = th.Color.Selection
		} else if vi%2 == 1 {
			bg = th.Color.SurfaceSunken
		}
		row.SetBox(scene.BoxSpec{Fill: scene.Solid(bg)})
		for ci, w := range widths {
			cell := scene.Box(scene.BoxSpec{}).SetWidth(w).SetHeight(t.RowHeight)
			txt := t.cellText(ci, dataRow)
			l := text.LayoutText(txt, text.DefaultFont(th.Font.SizeSM), text.Options{})
			tv := scene.Text(scene.TextSpec{Src: txt, Layout: l, Color: th.Color.Text})
			tv.SetMargin(geom.InLTRB(6, (t.RowHeight-l.Size.H)/2, 0, 0))
			cell.Add(tv)
			if ctx.Input != nil {
				cell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "table-row", Cursor: scene.CursorPointer})
				r := dataRow
				g := &ui.GestureDetector{OnTap: func(geom.Point) {
					t.Selected = r
					t.MarkDirty()
					if t.OnSelect != nil {
						t.OnSelect(r)
					}
				}}
				t.rowTaps = append(t.rowTaps, g)
				ctx.Input.Set(cell, g.Handler())
			}
			row.Add(cell)
		}
		row.SetStackPosition(geom.Pt(0, t.HeaderH+float32(vi)*t.RowHeight-t.offset))
		row.SetPositionKeepSize(true)
		col.Add(row)
		t.rowNodes = append(t.rowNodes, row)
		t.rowVis = append(t.rowVis, vi)
	}
	t.node = col

	// 滚轮：越界时返回 false，让外层滚动接手（嵌套滚动）。
	if ctx.Input != nil {
		ctx.Input.Set(col, input.Handler{Wheel: func(ev *input.WheelEvent) bool {
			if ev.Delta.Y == 0 {
				return false
			}
			before := t.offset
			t.SetScrollOffset(t.offset + ev.Delta.Y)
			return t.offset != before
		}})
	}

	max := t.maxOffset()
	show := max > 0
	if t.ShowScrollbar != nil {
		show = *t.ShowScrollbar
	}
	if !show {
		return col
	}
	if t.bar == nil {
		t.bar = &Scrollbar{}
	}
	t.bar.Axis = scene.Vertical
	t.bar.Max = max
	t.bar.Viewport = vh
	t.bar.Length = vh - 2*scrollbarInset
	t.bar.Value = t.offset
	t.bar.OnChange = func(v float32) { t.SetScrollOffset(v) }

	stack := scene.Stack().SetSize2(t.Width, vh)
	col.SetPositionKeepSize(true)
	col.SetStackPosition(geom.Pt(0, 0))
	stack.Add(col)
	barNode := t.bar.Build(ctx)
	barNode.SetPositionKeepSize(true)
	barNode.SetStackPosition(geom.Pt(t.Width-scrollbarThickness-scrollbarInset, scrollbarInset))
	stack.Add(barNode)
	return stack
}
