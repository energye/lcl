package main

import (
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
}

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
