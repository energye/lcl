package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var seg = widgets.NewSegmented("日", "周", "月", "年")

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
		},
		API: []doc.APIRow{
			{Name: "NewSegmented", Type: "func(options ...string) *Segmented", Desc: "构造"},
			{Name: "Options", Type: "[]string", Desc: "选项文本"},
			{Name: "Selected", Type: "int", Desc: "当前选中下标"},
			{Name: "OnChange", Type: "func(int)", Desc: "切换回调"},
		},
	})
}
