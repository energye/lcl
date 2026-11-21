package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var tbl *widgets.Table

func init() {
	cols := []widgets.TableColumn{
		{Title: "字段", Mode: widgets.ColumnAuto, Sortable: true},
		{Title: "类型", Mode: widgets.ColumnAuto},
		{Title: "主键", Mode: widgets.ColumnFixed, Width: 60},
		{Title: "说明", Mode: widgets.ColumnFlex, Flex: 1},
	}
	rows := [][]any{
		{"id", "INTEGER", "是", "主键"},
		{"name", "TEXT", "否", "名称"},
		{"created_at", "DATETIME", "否", "创建时间"},
		{"owner_id", "INTEGER", "否", "外键 → user.id"},
	}
	tbl = widgets.NewTable(cols, rows)
	tbl.Width, tbl.Height = 620, 180

	// tblPage 分页卡（ShowPagination）：23 行 × 每页 5 行（Pagination
	// 组件组合复用，Total=总条目数语义）。
	pcols := []widgets.TableColumn{
		{Title: "序号", Mode: widgets.ColumnFixed, Width: 60},
		{Title: "字段", Mode: widgets.ColumnAuto}, // 不 Sortable：table-sort 锚与四列混排卡撞车
		{Title: "说明", Mode: widgets.ColumnFlex, Flex: 1},
	}
	var prows [][]any
	for i := 1; i <= 23; i++ {
		prows = append(prows, []any{i, fmt.Sprintf("field_%02d", i), "第 " + fmt.Sprintf("%02d", i) + " 个字段"})
	}
	tblPage = widgets.NewTable(pcols, prows)
	tblPage.Width, tblPage.Height = 620, 180
	tblPage.ShowPagination = true
	tblPage.PageSize = 5

	// tblFreeze 固定列卡：列总宽 900 > 视口 620 → 横向滚动，
	// 左两列 Fixed="left" 冻结（叠层不随滚动）。
	fcols := []widgets.TableColumn{
		{Title: "字段", Mode: widgets.ColumnAuto, Fixed: "left"},
		{Title: "主键", Mode: widgets.ColumnFixed, Width: 70, Fixed: "left"},
		{Title: "类型", Mode: widgets.ColumnFixed, Width: 110},
		{Title: "默认值", Mode: widgets.ColumnFixed, Width: 120},
		{Title: "说明", Mode: widgets.ColumnFixed, Width: 200},
		{Title: "备注", Mode: widgets.ColumnFixed, Width: 200},
	}
	var frows [][]any
	for _, r := range rows {
		frows = append(frows, []any{r[0], r[2], r[1], "—", "说明 " + r[0].(string), "备注 " + r[0].(string)})
	}
	tblFreeze = widgets.NewTable(fcols, frows)
	tblFreeze.Width = 620
	tblFreeze.Height = 150
}

var (
	tblPage   *widgets.Table
	tblFreeze *widgets.Table
)

func main() {
	launcher.Main(doc.Meta{
		ID: "table", Title: "Table 表格", Group: doc.GroupData,
		What: []string{
			"虚拟化表格：只构建可见区 ±Overscan 行的节点，十万行也不撑爆节点树",
			"列宽三档（Q9）：Auto 按内容 / Fixed 固定 / Flex 按权重分配",
			"点击可排序列的表头排序（稳定排序）；斑马纹 + 行选中",
			"计划（M9）：固定列/列拖宽/行悬停/分页/筛选/自定义单元格渲染",
		},
		Demos: []doc.Demo{
			{Title: "分页表格", Note: "23 行 × 每页 5 行：底部分页器翻页（组合复用 Pagination）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tblPage
			}},
			{Title: "固定列", Note: "列总宽 900>视口 620：横向滚动，左两列冻结不随动（antd fixed=left）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tblFreeze
			}},
			{Title: "四列混排", Note: "Auto+Fixed+Flex 三档列宽；点击「字段」表头排序", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tbl
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTable", Type: "func(cols []TableColumn, rows [][]any) *Table", Desc: "构造"},
			{Name: "Columns", Type: "[]TableColumn", Desc: "列定义（Title/Mode/Width/Flex/Sortable/Align/Formatter）"},
			{Name: "Rows", Type: "[][]any", Desc: "行数据；单元格 any → 默认格式化为文本"},
			{Name: "Width/Height", Type: "float32", Default: "640 / 240", Desc: "视口尺寸"},
			{Name: "RowHeight", Type: "float32", Default: "30", Desc: "行高"},
			{Name: "SortColumn/Asc", Type: "int / bool", Desc: "排序状态（点表头切换）"},
			{Name: "OnSort", Type: "func(col int, asc bool)", Desc: "排序回调"},
			{Name: "OnSelect", Type: "func(row int)", Desc: "行选中回调"},
			{Name: "SetScrollOffset", Type: "func(v float32)", Desc: "滚动偏移（虚拟化唯一真相）"},
		},
	})
}
