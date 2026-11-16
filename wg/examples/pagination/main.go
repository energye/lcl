package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var page = widgets.NewPagination(3, 8)

func main() {
	launcher.Main(doc.Meta{
		ID: "pagination", Title: "Pagination 分页", Group: doc.GroupDisplay,
		What: []string{
			"页码导航：只显示当前页附近 ±2 页，首末页常驻可跳转",
			"计划（M11）：pageSize 切换、快速跳页输入框",
		},
		Demos: []doc.Demo{
			{Title: "8 页导航", Note: "当前第 3 页", Build: func(ctx *ui.BuildContext) ui.Widget {
				return page
			}},
		},
		API: []doc.APIRow{
			{Name: "NewPagination", Type: "func(page, total int) *Pagination", Desc: "构造"},
			{Name: "Page", Type: "int", Desc: "当前页（1 起）"},
			{Name: "Total", Type: "int", Desc: "总页数"},
			{Name: "OnChange", Type: "func(int)", Desc: "页码切换回调"},
		},
	})
}
