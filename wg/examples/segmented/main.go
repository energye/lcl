package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	seg = widgets.NewSegmented("日", "周", "月", "年")
	// segBlock block 撑满+图标项（组合复用 Icon）。
	segBlock = func() *widgets.Segmented {
		g := widgets.NewSegmented("列表", "看板", "日历")
		g.Block = true
		g.Width = 480
		g.Icons = []string{"list", "setting", "calendar"}
		g.Selected = 1
		return g
	}()
)

func main() {
	launcher.Main(doc.Meta{
		ID: "segmented", Title: "Segmented 分段选择", Group: doc.GroupDisplay,
		What: []string{
			"一组互斥选项的分段控件：点击切换选中段",
			"手势检测器跨 Build 存活，选中态由实例持有",
		},
		Demos: []doc.Demo{
			{Title: "周期切换", Note: "四段互斥；OnChange 回调下标", Build: func(ctx *ui.BuildContext) ui.Widget {
				return seg
			}},
			{Title: "block 撑满+图标项", Note: "Block=true 各段等分；Icons 左侧图标（组合复用 Icon）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return segBlock
			}},
		},
		API: []doc.APIRow{
			{Name: "NewSegmented", Type: "func(options ...string) *Segmented", Desc: "构造"},
			{Name: "Block", Type: "bool", Desc: "各段等分撑满父宽（antd block）"},
			{Name: "Icons", Type: "[]string", Desc: "与 Options 对齐的左侧图标名"},
			{Name: "Options", Type: "[]string", Desc: "选项文本"},
			{Name: "Selected", Type: "int", Desc: "当前选中下标"},
			{Name: "OnChange", Type: "func(int)", Desc: "切换回调"},
		},
	})
}
