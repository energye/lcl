package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	page = widgets.NewPagination(3, 8)

	// pageSized：Total=50 条、PageSize=10 → 5 页；切换每页条数页码窗口重算。
	pageSized = func() *widgets.Pagination {
		p := widgets.NewPagination(1, 50)
		p.PageSize = 10
		p.ShowTotal = true
		p.ShowSizeChanger = true
		return p
	}()

	// pageJump：快速跳页，回车提交（超范围钳制）。
	pageJump = func() *widgets.Pagination {
		p := widgets.NewPagination(1, 20)
		p.ShowQuickJumper = true
		return p
	}()
)

func main() {
	launcher.Main(doc.Meta{
		ID: "pagination", Title: "Pagination 分页", Group: doc.GroupDisplay,
		What: []string{
			"页码导航：只显示当前页附近 ±2 页，首末页常驻可跳转",
			"pageSize 模式：Total 为总条目数，切换每页条数页码窗口重算",
			"快速跳页：输入页码回车直达，超出范围自动钳制",
		},
		Demos: []doc.Demo{
			{Title: "8 页导航", Note: "当前第 3 页", Build: func(ctx *ui.BuildContext) ui.Widget {
				return page
			}},
			{Title: "pageSize 切换", Note: "Total=50 条 PageSize=10 → 5 页，切换每页条数重算页码", Build: func(ctx *ui.BuildContext) ui.Widget {
				return pageSized
			}},
			{Title: "快速跳页", Note: "ShowQuickJumper：输入页码回车直达（超范围钳制）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return pageJump
			}},
		},
		API: []doc.APIRow{
			{Name: "NewPagination", Type: "func(page, total int) *Pagination", Desc: "构造"},
			{Name: "Page", Type: "int", Desc: "当前页（1 起）"},
			{Name: "Total", Type: "int", Desc: "总页数；PageSize>0 时为总条目数"},
			{Name: "PageSize", Type: "int", Default: "0", Desc: "每页条数（>0 启用条目模式）"},
			{Name: "ShowTotal", Type: "bool", Desc: "显示\"共 X 条\"（需 PageSize>0）"},
			{Name: "ShowSizeChanger", Type: "bool", Desc: "显示\"X 条/页\"切换器（需 PageSize>0）"},
			{Name: "ShowQuickJumper", Type: "bool", Desc: "显示\"跳至 X 页\"输入框（回车提交）"},
			{Name: "SetPage", Type: "func(n int)", Desc: "程序化跳页（钳制+回调）"},
			{Name: "OnChange", Type: "func(int)", Desc: "页码切换回调"},
			{Name: "OnPageSize", Type: "func(int)", Desc: "每页条数变化回调"},
		},
	})
}
