package widgets

import (
	"sort"
	"strconv"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
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
	// Render 自定义单元格渲染（antd render 的组件级对应）：返回
	// Widget（Tag/Badge/按钮组等组合）。非空时优先于 Formatter 与
	// 默认格式化；列宽建议 Fixed/Width（Render 内容宽不可静态测量）。
	Render func(row int, value any) ui.Widget
	// Fixed 冻结列（antd fixed）："" 不冻结 / "left" 左缘冻结 /
	// "right" 右缘冻结——横向滚动时该列不随内容滚动（叠层实现）。
	Fixed string
	// Filters 筛选选项（antd filters）：非空时列头渲染筛选图标，点击
	// 弹出选项菜单（Dropdown 组合复用），选中值经 OnFilter 回调。
	Filters []string
}

// TableExpandable 是表格展开行配置（antd expandable）。
type TableExpandable struct {
	// ExpandedRowRender 展开内容渲染（antd expandedRowRender）：返回
	// Widget（描述文本/表单/按钮组等组合）。每次构建都会调用（与
	// 单元格 Render 同款契约）。非空时启用展开行。
	ExpandedRowRender func(row int) ui.Widget
	// DefaultExpandRows 初始展开的行（数据索引；antd defaultExpandAllRows
	// 的按行版）。展开态按数据索引记忆，排序后跟随数据行。
	DefaultExpandRows []int
	// ExpandRowByClick 点击整行切换展开（antd expandRowByClick）；开启后
	// 首列不再渲染展开图标。
	ExpandRowByClick bool
	// ExpandedRowHeight 展开行高（<=0 时用默认 48）。
	ExpandedRowHeight float32
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
	// RowLabel 自定义行命中标签（空 = "table-row"）。
	// 纯展示表（文档 API 参数表）设成非交互标签，自动化审计据此跳过。
	RowLabel string

	// ShowScrollbar：nil = 内容溢出自动显示；true 始终；false 隐藏（仍可滚轮）。
	ShowScrollbar *bool
	// ShowPagination 底部分页器（antd pagination，**Pagination 组件组合
	// 复用**）：开启后 Page 从 1 起、每页 PageSize 行（默认 10）。
	ShowPagination bool
	// Filtered 列筛选值（列下标→选中值，空串=清除筛选）。
	Filtered        map[int]string
	filterDropdowns []*Dropdown
	// OnFilter 筛选变化回调（列下标、选中值）。
	OnFilter func(col int, value string)
	// PageSize 每页行数（ShowPagination 时生效，默认 10）。
	PageSize  int
	pageNo    int         // 当前页（1 起，内部状态）
	pagePager *Pagination // 分页器实例池（跨重建持有）

	// Expandable 展开行配置（antd expandable）。ExpandedRowRender 非空时
	// 启用；展开态按数据索引记忆（排序/翻页跟随数据行）。
	Expandable TableExpandable
	expanded   map[int]bool // 展开态（数据索引集，跨 Build 持有）
	expandTaps []*ui.GestureDetector

	// CellSpan 单元格行/列合并（antd onCell 的 rowSpan/colSpan 直译）：
	// 返回该格 rowSpan 与 colSpan——1=普通格；0=**该格不渲染**（被合并
	// 格覆盖，antd「设值为 0 时不会渲染」同语义）；>1=跨 rowSpan 行 /
	// colSpan 列（格子加高加宽覆盖被合并格，被覆盖格渲染占位空格）。
	// 跨行合并基于**数据行号**（排序/筛选/翻页后合并关系按数据行走，
	// 显示序错位与 antd 排序破坏合并同——如实标注）。nil=全普通格。
	CellSpan func(row, col int) (rowSpan, colSpan int)

	hoverRow int // 行悬停（antd 行 hover 浅色反馈）；-1 = 无

	offset    float32
	sortCache []int
	sortDirty bool
	sorts     []*ui.GestureDetector // 跨 Build 存活
	rowTaps   []*ui.GestureDetector
	// 虚拟窗口与滚动条（原地滚动用）
	rowNodes     []*scene.Node
	freezeNodes  []*scene.Node // 左冻结列行的同步副本（横向滚动不随动）
	freezeRNodes []*scene.Node // 右冻结列行的同步副本（镜像左冻结）
	rowVis       []int
	builtFirst   int
	builtLast    int
	lastIdx      []int     // 最近一次 buildInner 生效的 idx（排序筛选分页后）
	lastRowY     []float32 // lastIdx 的行 y 前缀和（展开行不恒定高定位用）
	// 右冻结层宿主定位缓存（buildInner 产出、Build 包装定位）。
	lastFreezeRight  *scene.Node
	lastFreezeRightW float32
	bar              *Scrollbar
	reqFrame         func()
}

// DefaultTableHeight 是未设置 Height 时的视口高。
const DefaultTableHeight float32 = 240

// DefaultExpandRowH 是未设置 ExpandedRowHeight 时的展开行高。
const DefaultExpandRowH float32 = 48

// expandH 返回展开行高。
func (t *Table) expandH() float32 {
	if t.Expandable.ExpandedRowHeight > 0 {
		return t.Expandable.ExpandedRowHeight
	}
	return DefaultExpandRowH
}

// expandEnabled 返回展开行是否启用。
func (t *Table) expandEnabled() bool {
	return t.Expandable.ExpandedRowRender != nil
}

// isExpanded 返回数据行是否处于展开态。
func (t *Table) isExpanded(row int) bool { return t.expanded[row] }

// toggleExpand 切换展开态并重建（展开改变行高布局，走 MarkDirty 全量
// 重建前缀和——展开切换不频繁，不做增量）。
func (t *Table) toggleExpand(row int) {
	if t.expanded == nil {
		t.expanded = map[int]bool{}
	}
	if t.expanded[row] {
		delete(t.expanded, row)
	} else {
		t.expanded[row] = true
	}
	t.MarkDirty()
}

// rowHeightAt 返回视口序 vi 行的总高度（行高 + 展开内容高）。
func (t *Table) rowHeightAt(vi int, idx []int) float32 {
	h := t.RowHeight
	if t.expandEnabled() && t.isExpanded(idx[vi]) {
		h += t.expandH()
	}
	return h
}

// rowYPrefixSum 计算视口行的 y 前缀和（len(idx)+1 项；rowY[k] 为第 k 行
// 顶缘相对内容区顶的偏移）。展开使行高不再恒定，虚拟化定位统一走它。
func (t *Table) rowYPrefixSum(idx []int) []float32 {
	rowY := make([]float32, len(idx)+1)
	for vi := range idx {
		rowY[vi+1] = rowY[vi] + t.rowHeightAt(vi, idx)
	}
	return rowY
}

// NewTable 创建表格。
func NewTable(cols []TableColumn, rows [][]any) *Table {
	return &Table{
		Columns: cols, Rows: rows, Width: 640, RowHeight: 30, HeaderH: 32,
		Overscan: 2, SortColumn: -1, Selected: -1, hoverRow: -1, sortDirty: true,
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
	// 展开行启用时行高不恒定，用构建时缓存的前缀和定位；否则线性公式。
	y := func(vi int) float32 {
		if t.lastRowY != nil && vi < len(t.lastRowY)-1 {
			return t.lastRowY[vi]
		}
		return float32(vi) * t.RowHeight
	}
	for i, n := range t.rowNodes {
		n.SetStackPosition(geom.Pt(0, t.HeaderH+y(t.rowVis[i])-v))
	}
	for i, n := range t.freezeNodes {
		if i < len(t.rowVis) {
			n.SetStackPosition(geom.Pt(0, t.HeaderH+y(t.rowVis[i])-v))
		}
	}
	for i, n := range t.freezeRNodes {
		if i < len(t.rowVis) {
			n.SetStackPosition(geom.Pt(0, t.HeaderH+y(t.rowVis[i])-v))
		}
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
// 展开行启用时行高不恒定：先按前缀和找偏移命中的首行，再向后累计到
// 视口底；否则用 O(1) 线性公式。
func (t *Table) rowWindow() (int, int) {
	vh := t.ViewportHeight()
	n := len(t.Rows)
	if t.expandEnabled() {
		rowY := t.rowYPrefixSum(t.lastIdx)
		// 二分：找最后一个 y <= offset 的行
		first := 0
		lo, hi := 0, len(t.lastIdx)
		for lo < hi {
			mid := (lo + hi) / 2
			if rowY[mid+1] <= t.offset {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		first = lo - t.Overscan
		if first < 0 {
			first = 0
		}
		// 从 first 向后累计到覆盖视口高
		last := first
		for last < len(t.lastIdx) && rowY[last]-rowY[first] < vh+t.expandH() {
			last++
		}
		last += t.Overscan
		if last > len(t.lastIdx) {
			last = len(t.lastIdx)
		}
		return first, last
	}
	first := int(t.offset/t.RowHeight) - t.Overscan
	if first < 0 {
		first = 0
	}
	last := first + int(vh/t.RowHeight) + t.Overscan*2 + 2
	if last > n {
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
	if t.expandEnabled() {
		rowY := t.rowYPrefixSum(t.lastIdx)
		total = t.HeaderH + rowY[len(rowY)-1]
	}
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
// hScrollNeeded 判断是否需要横向滚动：各列内容宽总和 > Width。
func (t *Table) hScrollNeeded(fontSize float32) bool {
	_, total := t.contentWidths(fontSize)
	return total > t.Width
}

// contentWidths 折算各列"内容宽"（Auto=内容最大宽、Fixed=像素、
// Flex=Min 或内容宽）与总和——横向滚动判定与布局共用。
func (t *Table) contentWidths(fontSize float32) ([]float32, float32) {
	n := len(t.Columns)
	if n == 0 {
		return nil, 0
	}
	out := make([]float32, n)
	font := text.DefaultFont(fontSize)
	total := float32(0)
	for i, c := range t.Columns {
		w := c.Width
		switch c.Mode {
		case ColumnFixed:
		default: // Auto/Flex 都取内容最大宽（Flex 上限 Min/Max 仍生效）
			for r := range t.Rows {
				if lw := text.Measure(t.cellText(i, r), font).W; lw > w {
					w = lw
				}
			}
		}
		if c.Min > 0 && w < c.Min {
			w = c.Min
		}
		if c.Max > 0 && w > c.Max {
			w = c.Max
		}
		out[i] = w + 12
		total += out[i]
	}
	return out, total
}

func (t *Table) columnWidths(fontSize float32) []float32 {
	n := len(t.Columns)
	if n == 0 {
		return nil
	}
	out := make([]float32, n)
	remaining := t.Width
	flexTotal := float32(0)
	font := text.DefaultFont(fontSize)
	// 横向滚动模式（列内容总宽>Width）：Auto/Flex 全部退回内容宽，
	// 不做剩余分摊——列总宽超视口交给横向滚动。
	if t.hScrollNeeded(fontSize) {
		cw, _ := t.contentWidths(fontSize)
		return cw
	}
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
	idx := t.lastIdx
	if idx == nil {
		idx = t.sortedIndex()
	}
	first, last := t.rowWindow()
	if first > len(idx) {
		first = len(idx)
	}
	if last > len(idx) {
		last = len(idx)
	}
	if first > last {
		first = last
	}
	return idx[first:last]
}

// Build 实现 ui.Widget。
func (t *Table) Build(ctx *ui.BuildContext) *scene.Node {
	// 行选中/排序走 MarkDirty：不接线的话静止窗口（C14）点击后
	// 只置脏位不请求帧，界面表现为"点了没反应"（--audit 实测）。
	ctx.BindState(t)
	th := ctx.ThemeOr()
	idx := t.sortedIndex()
	// 展开行启用：DefaultExpandRows 初始化展开态（首次构建生效）。
	if t.expandEnabled() && t.expanded == nil && len(t.Expandable.DefaultExpandRows) > 0 {
		t.expanded = map[int]bool{}
		for _, r := range t.Expandable.DefaultExpandRows {
			t.expanded[r] = true
		}
	}
	// 筛选应用：Filtered 非空列按单元格文本过滤（筛选项与文本精确匹配）。
	for col2, fv := range t.Filtered {
		if fv == "" {
			continue
		}
		var kept []int
		for _, r := range idx {
			if t.cellText(col2, r) == fv {
				kept = append(kept, r)
			}
		}
		idx = kept
	}

	// 分页模式：排序结果按页切片（Pagination 组件组合复用管页码），
	// 页内行子集仍走虚拟化窗口。
	if t.ShowPagination && t.PageSize > 0 {
		ps := t.PageSize
		total := len(idx)
		maxPage := (total + ps - 1) / ps
		if maxPage < 1 {
			maxPage = 1
		}
		if t.pageNo > maxPage {
			t.pageNo = maxPage
		}
		if t.pageNo < 1 {
			t.pageNo = 1
		}
		lo := (t.pageNo - 1) * ps
		if lo > total {
			lo = total
		}
		hi := lo + ps
		if hi > total {
			hi = total
		}
		idx = idx[lo:hi]
		inner, freeze, _ := t.buildInner(ctx, th, idx)
		tableArea := scene.Stack().SetSize2(t.Width, t.ViewportHeight())
		tableArea.Add(inner)
		if freeze != nil {
			tableArea.Add(freeze)
		}
		// 分页模式同样叠右冻结层（贴视口右缘）。
		if freezeR := t.lastFreezeRight; freezeR != nil {
			freezeR.SetPositionKeepSize(true)
			freezeR.SetStackPosition(geom.Pt(t.Width-t.lastFreezeRightW, 0))
			tableArea.Add(freezeR)
		}
		col := scene.Flex(scene.Vertical, th.Space.SM)
		col.Add(tableArea)
		// 底部分页器行（Pagination 组件组合复用；Total=总条目数语义）。
		if t.pagePager == nil {
			t.pagePager = &Pagination{}
		}
		t.pagePager.PageSize = t.PageSize
		t.pagePager.Total = len(t.Rows)
		t.pagePager.ShowTotal = true // antd 表格分页默认显示总条数
		t.pagePager.Page = t.pageNo
		t.pagePager.OnChange = func(n int) {
			t.pageNo = n
			t.MarkDirty()
		}
		pager := t.pagePager.Build(ctx)
		pagerRow := scene.Flex(scene.Horizontal, 0).SetMainAlign(scene.MainEnd)
		pagerRow.Add(pager)
		col.Add(pagerRow)
		return col
	}

	inner, freeze, hscroll := t.buildInner(ctx, th, idx)
	if !hscroll {
		return inner
	}
	// 横向滚动包装：内容行进 Horizontal Scroll（视口宽=Table.Width）；
	// 冻结列叠层覆盖其上不随滚动（Z 序：Scroll 下、freeze 上）。
	wrapped := scene.Scroll(scene.Horizontal, inner)
	wrapped.SetWidth(t.Width)
	wrapped.SetHeight(t.ViewportHeight())
	wrapped.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "table-hscroll"})
	// 横向滚轮：scene.Scroll 节点本身无 handler，Delta.X 需在此消费
	// （SetScrollOffset 自动 clamp 到内容范围；Delta.Y 冒泡给页面）。
	if ctx.Input != nil {
		ctx.Input.Set(wrapped, input.Handler{Wheel: func(ev *input.WheelEvent) bool {
			if ev.Delta.X == 0 {
				return false
			}
			wrapped.SetScrollOffset(wrapped.ScrollOffset() + ev.Delta.X)
			if ctx.RequestFrame != nil {
				ctx.RequestFrame()
			}
			return true
		}})
	}
	stack := scene.Stack().SetSize2(t.Width, t.ViewportHeight())
	stack.Add(wrapped)
	if freeze != nil {
		freeze.SetPositionKeepSize(true)
		freeze.SetStackPosition(geom.Pt(0, 0))
		stack.Add(freeze)
	}
	// 右冻结层贴视口右缘（buildInner 不知宿主 x，此处定位）。
	if freezeR := t.lastFreezeRight; freezeR != nil {
		freezeR.SetPositionKeepSize(true)
		freezeR.SetStackPosition(geom.Pt(t.Width-t.lastFreezeRightW, 0))
		stack.Add(freezeR)
	}
	stack.SetHit(scene.HitSpec{Label: "table-hscroll-wrap"})
	return stack
}

// buildInner 构建表格本体（表头+虚拟行[+滚动条]），idx 为已排序（可能
// 已分页切片）的数据行下标。
// spanCache 是 buildInner 期间的合并格速查（数据行→列→(rowSpan,colSpan)）。
type spanCache struct {
	t     *Table
	rows  map[int]map[int][2]int
	colsN int
}

// at 查 (row,col) 合并跨度（CellSpan 每格只调一次缓存；负值钳 1）。
func (s *spanCache) at(row, col int) (rs, cs int) {
	rs, cs = 1, 1
	if s.t.CellSpan == nil || col < 0 || col >= s.colsN {
		return
	}
	m, ok := s.rows[row]
	if !ok {
		m = map[int][2]int{}
		s.rows[row] = m
	}
	if v, ok := m[col]; ok {
		return v[0], v[1]
	}
	rs, cs = s.t.CellSpan(row, col)
	if rs < 0 {
		rs = 1
	}
	if cs < 0 {
		cs = 1
	}
	m[col] = [2]int{rs, cs}
	return
}

// covered 判断 (row,col) 是否被上方/左侧合并格覆盖（span 0 格不渲染）。
func (s *spanCache) covered(row, col int) bool {
	for r := row - 1; r >= 0 && r >= row-16; r-- {
		if rs, _ := s.at(r, col); rs > 1 && r+rs > row {
			return true
		}
	}
	for c := col - 1; c >= 0 && c >= col-16; c-- {
		if _, cs := s.at(row, c); cs > 1 && c+cs > col {
			return true
		}
	}
	return false
}

func (t *Table) buildInner(ctx *ui.BuildContext, th *theme.Data, idx []int) (*scene.Node, *scene.Node, bool) {
	t.lastIdx = idx
	t.lastRowY = nil
	if t.expandEnabled() {
		t.lastRowY = t.rowYPrefixSum(idx)
	}
	widths := t.columnWidths(th.Font.SizeBase)
	hscroll := t.hScrollNeeded(th.Font.SizeBase)

	col := scene.Flex(scene.Vertical, 0).SetWidth(t.Width).SetHeight(t.ViewportHeight())
	col.SetOverflow(true)
	// 横向滚动模式：容器与行宽=列内容总宽（列不再分摊 Flex/剩余）。
	contentW := float32(0)
	for _, w := range widths {
		contentW += w
	}
	freezeW := float32(0)
	freezeRightW := float32(0)
	var freeze, freezeRight *scene.Node
	if hscroll {
		col.SetWidth(contentW)
		widths, _ = t.contentWidths(th.Font.SizeBase)
		for _, w := range widths {
			_ = w
		}
		for ci := range t.Columns {
			switch t.Columns[ci].Fixed {
			case "left":
				freezeW += widths[ci]
			case "right":
				freezeRightW += widths[ci]
			}
		}
		if freezeW > 0 {
			freeze = scene.Flex(scene.Vertical, 0).SetWidth(freezeW).SetHeight(t.ViewportHeight())
			freeze.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "table-freeze"})
		}
		if freezeRightW > 0 {
			freezeRight = scene.Flex(scene.Vertical, 0).SetWidth(freezeRightW).SetHeight(t.ViewportHeight())
			freezeRight.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "table-freeze-right"})
		}
	}
	t.lastFreezeRight = freezeRight
	t.lastFreezeRightW = freezeRightW

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
		l := text.LayoutText(label, text.DefaultFont(th.Font.SizeBase), text.Options{})
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
		// 列筛选：Filters 非空时列头右缘漏斗图标，点击弹出选项菜单
		// （**Dropdown 组件组合复用**——菜单项点击回调设筛选值）。
		if len(c.Filters) > 0 && ctx.Input != nil {
			colIdx := ci
			filters := c.Filters
			funnel := (&Icon{Name: "search", Size: 13,
				Color: map[bool]geom.Color{true: th.Color.Primary}[t.Filtered[ci] != ""]}).Build(ctx)
			funnel.SetPositionKeepSize(true)
			funnel.SetStackPosition(geom.Pt(w-18, (t.HeaderH-13)/2))
			funnel.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "table-filter", Cursor: scene.CursorPointer})
			fd := &Dropdown{Items: func() []*MenuItem {
				out := []*MenuItem{{Key: "", Label: "全部"}}
				for _, f := range filters {
					out = append(out, &MenuItem{Key: f, Label: f})
				}
				return out
			}()}
			fd.OnClick = func(key string) {
				if t.Filtered == nil {
					t.Filtered = make(map[int]string)
				}
				t.Filtered[colIdx] = key
				t.MarkDirty()
				if t.OnFilter != nil {
					t.OnFilter(colIdx, key)
				}
			}
			// 漏斗图标+菜单项的手势：down 弹出（**ShowAt 显式锚**——
			// 筛选场景 Dropdown 无 Child，Show 取 childNode 会静默
			// return——真窗点漏斗永不弹菜单的断链修复）。
			gf := &ui.GestureDetector{OnTap: func(geom.Point) {
				fd.ShowAt(ctx, funnel)
			}}
			t.sorts = append(t.sorts, gf)
			ctx.Input.Set(funnel, gf.Handler())
			t.filterDropdowns = append(t.filterDropdowns, fd)
			cell.Add(funnel)
		}
		header.Add(cell)
	}
	col.Add(header)

	// 右冻结表头副本：只渲染 Fixed=="right" 的表头单元格（白底盖住滚动
	// 内容，左缘描边做冻结分界线——与左冻结右缘线镜像）。
	if freezeRight != nil {
		hRow := scene.Flex(scene.Horizontal, 0).SetWidth(freezeRightW).SetHeight(t.HeaderH)
		hRow.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface)})
		hx := float32(0)
		for ci, w := range widths {
			if t.Columns[ci].Fixed != "right" {
				continue
			}
			hcell := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken), Border: scene.Stroke(th.Color.Border, 1)}).SetWidth(w).SetHeight(t.HeaderH)
			hl := text.LayoutText(t.Columns[ci].Title, text.DefaultFont(th.Font.SizeBase), text.Options{})
			htv := scene.Text(scene.TextSpec{Src: t.Columns[ci].Title, Layout: hl, Color: th.Color.TextMuted})
			htv.SetMargin(geom.InLTRB(6, (t.HeaderH-hl.Size.H)/2, 0, 0))
			hcell.Add(htv)
			hcell.SetStackPosition(geom.Pt(hx, 0))
			hcell.SetPositionKeepSize(true)
			hcell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "table-frozen-head-right"})
			hRow.Add(hcell)
			hx += w
		}
		// 冻结分界线：副本行左缘 2px 竖线（左冻结右缘线镜像）。
		edge := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetWidth(2).SetHeight(t.HeaderH)
		edge.SetStackPosition(geom.Pt(0, 0))
		edge.SetPositionKeepSize(true)
		hRow.Add(edge)
		hRow.SetStackPosition(geom.Pt(0, 0))
		hRow.SetPositionKeepSize(true)
		freezeRight.Add(hRow)
	}

	// 虚拟行：只构建可见区 ± Overscan
	vh := t.ViewportHeight()
	var first, last int
	if t.expandEnabled() {
		// 展开行启用：行高不恒定，前缀和定位可见窗口（rowY 长度 = len(idx)+1）。
		lo, hi := 0, len(idx)
		for lo < hi {
			mid := (lo + hi) / 2
			if t.lastRowY[mid+1] <= t.offset {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		first = lo - t.Overscan
		if first < 0 {
			first = 0
		}
		last = first
		for last < len(idx) && t.lastRowY[last]-t.lastRowY[first] < vh+t.expandH() {
			last++
		}
		last += t.Overscan
		if last > len(idx) {
			last = len(idx)
		}
	} else {
		first = int(t.offset/t.RowHeight) - t.Overscan
		if first < 0 {
			first = 0
		}
		last = first + int(vh/t.RowHeight) + t.Overscan*2 + 2
		if last > len(idx) {
			last = len(idx)
		}
	}
	t.rowTaps = t.rowTaps[:0]
	t.expandTaps = t.expandTaps[:0]
	t.rowNodes = t.rowNodes[:0]
	t.freezeNodes = t.freezeNodes[:0]
	t.freezeRNodes = t.freezeRNodes[:0]
	t.rowVis = t.rowVis[:0]
	t.builtFirst, t.builtLast = first, last
	t.reqFrame = ctx.RequestFrame
	rowLabel := t.RowLabel
	if rowLabel == "" {
		rowLabel = "table-row"
	}
	// 行 y 定位：展开启用走前缀和，否则线性公式。
	rowYAt := func(vi int) float32 {
		if t.lastRowY != nil {
			return t.lastRowY[vi]
		}
		return float32(vi) * t.RowHeight
	}
	// 合并格速查：spanAt/covered（数据行→列缓存；行循环/覆盖判断共用）。
	sp := &spanCache{t: t, rows: map[int]map[int][2]int{}, colsN: len(t.Columns)}
	for vi := first; vi < last; vi++ {
		dataRow := idx[vi]
		expanded := t.expandEnabled() && t.isExpanded(dataRow)
		// 展开图标列：ExpandRowByClick=false 时首列前插图标单元格
		//（antd expandable 每行首列的折叠按钮）。列宽模型不变——图标
		// 占用首列单元格左缘内边距区。
		row := scene.Flex(scene.Horizontal, 0).SetWidth(t.Width).SetHeight(t.RowHeight)
		// 行底三态：选中 Selection > hover FillHover > 斑马 SurfaceSunken。
		// zebra 记录本行斑马底（hover 原地迁移时保留斑马底仍判选中）。
		bg := geom.Color{}
		zebra := false
		if dataRow == t.Selected {
			bg = th.Color.Selection
		} else if dataRow == t.hoverRow {
			// antd 行悬停：浅色底（colorFillTertiary 系 FillHover）。
			bg = th.Color.FillHover
		} else if vi%2 == 1 {
			bg = th.Color.SurfaceSunken
			zebra = true
		}
		row.SetBox(scene.BoxSpec{Fill: scene.Solid(bg)})
		for ci, w := range widths {
			// 合并语义（antd onCell rowSpan/colSpan）：span 0=格不渲染
			//（被合并格覆盖）；colSpan>1 格加宽吃掉右侧列；rowSpan>1 格
			// 脱离行流绝对定位跨行覆盖。占位逻辑：covered 格渲染透明
			// 占位保住列 x 布局（不挂 hit——点击落在覆盖格上）。
			rs, cs := sp.at(dataRow, ci)
			if sp.covered(dataRow, ci) {
				ph := scene.Box(scene.BoxSpec{}).SetWidth(w).SetHeight(t.RowHeight)
				row.Add(ph)
				continue
			}
			cellW := w
			if cs > 1 {
				for k := 1; k < cs && ci+k < len(widths); k++ {
					cellW += widths[ci+k]
				}
			}
			cell := scene.Box(scene.BoxSpec{}).SetWidth(cellW).SetHeight(t.RowHeight)
			if rs > 1 {
				// 跨行格：脱离行流，绝对定位到 bar 层（高度=rs 行高，
				// 白底盖住下方被合并行——同冻结列叠层手法）。
				spanH := float32(rs) * t.RowHeight
				cell = scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface)}).SetWidth(cellW).SetHeight(spanH)
			}
			// 展开图标：首列且非行点击模式时渲染 caret（right=收起/down=展开）。
			if ci == 0 && t.expandEnabled() && !t.Expandable.ExpandRowByClick {
				name := "right"
				if expanded {
					name = "down"
				}
				ic := (&Icon{Name: name, Size: 12, Color: th.Color.TextFaint}).Build(ctx)
				ic.SetPositionKeepSize(true)
				ic.SetStackPosition(geom.Pt(8, (t.RowHeight-12)/2))
				ic.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "table-expand", Cursor: scene.CursorPointer})
				r0 := dataRow
				g := &ui.GestureDetector{OnTap: func(geom.Point) { t.toggleExpand(r0) }}
				t.expandTaps = append(t.expandTaps, g)
				ctx.Input.Set(ic, g.Handler())
				cell.Add(ic)
			}
			if c := t.Columns[ci]; c.Render != nil {
				// antd render：组件级单元格渲染（Tag/按钮组等组合）。
				// 行数组列数可能少于列定义（cellText 有界保护同款语义）。
				var val any
				if dataRow < len(t.Rows) && ci < len(t.Rows[dataRow]) {
					val = t.Rows[dataRow][ci]
				}
				if wgt := c.Render(dataRow, val); wgt != nil {
					rn := wgt.Build(ctx)
					padL := float32(6)
					if ci == 0 && t.expandEnabled() && !t.Expandable.ExpandRowByClick {
						padL = 24 // 首列给展开图标让位
					}
					rn.SetMargin(geom.InLTRB(padL, (t.RowHeight-rn.Size().H)/2, 0, 0))
					cell.Add(rn)
				}
			} else {
				txt := t.cellText(ci, dataRow)
				l := text.LayoutText(txt, text.DefaultFont(th.Font.SizeBase), text.Options{})
				tv := scene.Text(scene.TextSpec{Src: txt, Layout: l, Color: th.Color.Text})
				padL := float32(6)
				if ci == 0 && t.expandEnabled() && !t.Expandable.ExpandRowByClick {
					padL = 24
				}
				tv.SetMargin(geom.InLTRB(padL, (t.RowHeight-l.Size.H)/2, 0, 0))
				cell.Add(tv)
			}
			if ctx.Input != nil {
				cell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: rowLabel, Cursor: scene.CursorPointer})
				r := dataRow
				g := &ui.GestureDetector{OnTap: func(geom.Point) {
					t.Selected = r
					if t.expandEnabled() && t.Expandable.ExpandRowByClick {
						t.toggleExpand(r)
					}
					t.MarkDirty()
					if t.OnSelect != nil {
						t.OnSelect(r)
					}
				}}
				h := g.Handler()
				// 行悬停：Enter/Leave 原地迁移行底（TimePicker SetBox 推广
				// ——不再 MarkDirty 整树重建）。hover 底 FillHover 只在非
				// 选中行叠加（选中行底恒 Selection）；斑马行迁出时回
				// SurfaceSunken。hoverRow 字段仍维护——选中/排序等真重建
				// 场景消费同一状态。
				h.Enter = func(*input.PointerEvent) {
					if t.hoverRow != r && r != t.Selected {
						t.hoverRow = r
						row.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.FillHover)})
						row.MarkNeedsPaint()
					} else if t.hoverRow != r {
						t.hoverRow = r
					}
				}
				h.Leave = func(*input.PointerEvent) {
					if t.hoverRow == r {
						t.hoverRow = -1
						if r != t.Selected {
							base := geom.Color{}
							if zebra {
								base = th.Color.SurfaceSunken
							}
							row.SetBox(scene.BoxSpec{Fill: scene.Solid(base)})
							row.MarkNeedsPaint()
						}
					}
				}
				t.rowTaps = append(t.rowTaps, g)
				ctx.Input.Set(cell, h)
			}
			if rs > 1 {
				// 跨行格：绝对定位覆盖（x=前序列宽前缀和，y=本行顶延伸
				// rs 行；白底盖住被合并行——同冻结列叠层手法）。
				// 注意：跨行格与展开行/斑马行并存时白底覆盖斑马——antd
				// 合并格视觉即整块底色，语义一致。
				cx := float32(0)
				for k := 0; k < ci && k < len(widths); k++ {
					cx += widths[k]
				}
				cell.SetStackPosition(geom.Pt(cx, t.HeaderH+rowYAt(vi)-t.offset))
				cell.SetPositionKeepSize(true)
				col.Add(cell)
				t.rowNodes = append(t.rowNodes, cell) // 随行节点集维护（滚动重建窗口一致）
			} else {
				row.Add(cell)
			}
		}
		row.SetStackPosition(geom.Pt(0, t.HeaderH+rowYAt(vi)-t.offset))
		row.SetPositionKeepSize(true)
		col.Add(row)
		t.rowNodes = append(t.rowNodes, row)
		t.rowVis = append(t.rowVis, vi)
		// 展开内容行：紧跟父行，整行宽、浅底（antd 展开行 fillQuaternary
		// 视觉 ≈ 白底 FillHover 合成色）。
		if expanded {
			eRow := scene.Flex(scene.Horizontal, 0).SetWidth(t.Width).SetHeight(t.expandH())
			eRow.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.FillHover)})
			if wgt := t.Expandable.ExpandedRowRender(dataRow); wgt != nil {
				en := wgt.Build(ctx)
				en.SetMargin(geom.InLTRB(24, 8, 12, 8))
				eRow.Add(en)
			}
			eRow.SetHit(scene.HitSpec{Label: "table-expanded-row"})
			eRow.SetStackPosition(geom.Pt(0, t.HeaderH+rowYAt(vi)+t.RowHeight-t.offset))
			eRow.SetPositionKeepSize(true)
			col.Add(eRow)
		}
		// 左冻结列副本行：只渲染 Fixed=="left" 的单元格（白底盖住滚动
		// 内容，右缘描边做冻结分界线）。横向滚动时此副本不随内容平移。
		if freezeW > 0 {
			fRow := scene.Flex(scene.Horizontal, 0).SetWidth(freezeW).SetHeight(t.RowHeight)
			fRow.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface)})
			fx := float32(0)
			for ci, w := range widths {
				if t.Columns[ci].Fixed != "left" {
					fx += w
					continue
				}
				fcell := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface)}).SetWidth(w).SetHeight(t.RowHeight)
				if c := t.Columns[ci]; c.Render != nil {
					var fval any
					if dataRow < len(t.Rows) && ci < len(t.Rows[dataRow]) {
						fval = t.Rows[dataRow][ci]
					}
					if wgt := c.Render(dataRow, fval); wgt != nil {
						rn := wgt.Build(ctx)
						rn.SetMargin(geom.InLTRB(6, (t.RowHeight-rn.Size().H)/2, 0, 0))
						fcell.Add(rn)
					}
				} else {
					ftxt := t.cellText(ci, dataRow)
					fl := text.LayoutText(ftxt, text.DefaultFont(th.Font.SizeBase), text.Options{})
					ftv := scene.Text(scene.TextSpec{Src: ftxt, Layout: fl, Color: th.Color.Text})
					ftv.SetMargin(geom.InLTRB(6, (t.RowHeight-fl.Size.H)/2, 0, 0))
					fcell.Add(ftv)
				}
				fcell.SetStackPosition(geom.Pt(fx, 0))
				fcell.SetPositionKeepSize(true)
				// 冻结副本单元格独立锚（gapBetween 验证"不随滚"——滚动层
				// 的 table-row 锚随内容平移，冻结层恒定）。
				fcell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "table-frozen-cell"})
				fRow.Add(fcell)
				fx += w
			}
			// 冻结分界线：副本行右缘 2px 竖线。
			edge := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetWidth(2).SetHeight(t.RowHeight)
			edge.SetStackPosition(geom.Pt(freezeW-2, 0))
			edge.SetPositionKeepSize(true)
			fRow.Add(edge)
			fRow.SetStackPosition(geom.Pt(0, t.HeaderH+rowYAt(vi)-t.offset))
			fRow.SetPositionKeepSize(true)
			freeze.Add(fRow)
			t.freezeNodes = append(t.freezeNodes, fRow)
		}
		// 右冻结列副本行：只渲染 Fixed=="right" 的单元格（白底盖住滚动
		// 内容，左缘描边做冻结分界线）。镜像左冻结叠层。
		if freezeRightW > 0 {
			fRow := scene.Flex(scene.Horizontal, 0).SetWidth(freezeRightW).SetHeight(t.RowHeight)
			fRow.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface)})
			fx := float32(0)
			for ci, w := range widths {
				if t.Columns[ci].Fixed != "right" {
					continue
				}
				fcell := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface)}).SetWidth(w).SetHeight(t.RowHeight)
				if c := t.Columns[ci]; c.Render != nil {
					var fval any
					if dataRow < len(t.Rows) && ci < len(t.Rows[dataRow]) {
						fval = t.Rows[dataRow][ci]
					}
					if wgt := c.Render(dataRow, fval); wgt != nil {
						rn := wgt.Build(ctx)
						rn.SetMargin(geom.InLTRB(6, (t.RowHeight-rn.Size().H)/2, 0, 0))
						fcell.Add(rn)
					}
				} else {
					ftxt := t.cellText(ci, dataRow)
					fl := text.LayoutText(ftxt, text.DefaultFont(th.Font.SizeBase), text.Options{})
					ftv := scene.Text(scene.TextSpec{Src: ftxt, Layout: fl, Color: th.Color.Text})
					ftv.SetMargin(geom.InLTRB(6, (t.RowHeight-fl.Size.H)/2, 0, 0))
					fcell.Add(ftv)
				}
				fcell.SetStackPosition(geom.Pt(fx, 0))
				fcell.SetPositionKeepSize(true)
				fcell.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "table-frozen-cell-right"})
				fRow.Add(fcell)
				fx += w
			}
			// 冻结分界线：副本行左缘 2px 竖线。
			edge := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetWidth(2).SetHeight(t.RowHeight)
			edge.SetStackPosition(geom.Pt(0, 0))
			edge.SetPositionKeepSize(true)
			fRow.Add(edge)
			fRow.SetStackPosition(geom.Pt(0, t.HeaderH+rowYAt(vi)-t.offset))
			fRow.SetPositionKeepSize(true)
			freezeRight.Add(fRow)
			t.freezeRNodes = append(t.freezeRNodes, fRow)
		}
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
	show := max > 0 && !hscroll // 横向滚动模式分页替代纵向滚动，双滚动条互斥
	if t.ShowScrollbar != nil {
		show = *t.ShowScrollbar
	}
	if !show || hscroll {
		// 横向滚动模式：col 进 Horizontal Scroll，冻结叠层覆盖其上。
		wrapped := scene.Scroll(scene.Horizontal, col)
		wrapped.SetWidth(t.Width)
		wrapped.SetHeight(t.ViewportHeight())
		wrapped.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "table-hscroll"})
		stack := scene.Stack().SetSize2(t.Width, t.ViewportHeight())
		stack.Add(wrapped)
		if freeze != nil {
			// 冻结层 HitOpaque 挡住下层 wrapped 的 Wheel handler——Delta.X
			// 在此转发给 wrapped（Delta.Y 返回 false 冒泡给页面纵向滚动）。
			if ctx.Input != nil {
				wh := wrapped
				ctx.Input.Set(freeze, input.Handler{Wheel: func(ev *input.WheelEvent) bool {
					if ev.Delta.X == 0 {
						return false
					}
					wh.SetScrollOffset(wh.ScrollOffset() + ev.Delta.X)
					if ctx.RequestFrame != nil {
						ctx.RequestFrame()
					}
					return true
				}})
			}
			freeze.SetPositionKeepSize(true)
			freeze.SetStackPosition(geom.Pt(0, 0))
			stack.Add(freeze)
		}
		if freezeRight != nil {
			// 右冻结层贴视口右缘（左冻结镜像）；wheel 转发同左冻结。
			if ctx.Input != nil {
				wh := wrapped
				ctx.Input.Set(freezeRight, input.Handler{Wheel: func(ev *input.WheelEvent) bool {
					if ev.Delta.X == 0 {
						return false
					}
					wh.SetScrollOffset(wh.ScrollOffset() + ev.Delta.X)
					if ctx.RequestFrame != nil {
						ctx.RequestFrame()
					}
					return true
				}})
			}
			freezeRight.SetPositionKeepSize(true)
			freezeRight.SetStackPosition(geom.Pt(t.Width-freezeRightW, 0))
			stack.Add(freezeRight)
		}
		stack.SetHit(scene.HitSpec{Label: "table-hscroll-wrap"})
		return stack, freeze, hscroll
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
	return stack, freeze, hscroll
}
