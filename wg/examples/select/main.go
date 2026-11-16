package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	sel    = widgets.NewSelect([]string{"软件光栅化", "GDI 混合", "自动选择"})
	status = widgets.NewButton("尚未选择")
)

func init() {
	sel.Width = 260
	sel.OnChange = func(i int) {
		status.Text = fmt.Sprintf("已选：%d · %s", i, sel.Options[i])
		status.MarkDirty()
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "select", Title: "Select 下拉选择", Group: doc.GroupData,
		What: []string{
			"下拉选择：选项列表走 overlay.Host（BandPopup 层带）",
			"关闭策略：OutsideClick / Esc 由 overlay.Host 统一裁决",
			"计划（M8）：对齐 antd——搜索、多选、分组、自定义 option 渲染",
		},
		Demos: []doc.Demo{
			{Title: "字符串选项", Note: "选中后回显到下方按钮", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(sel.Build(ctx))
				col.Add(status.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewSelect", Type: "func(options []string) *Select", Desc: "构造"},
			{Name: "Options", Type: "[]string", Desc: "选项文本"},
			{Name: "Selected", Type: "int", Default: "-1（未选）", Desc: "选中下标"},
			{Name: "Width/MaxHeight", Type: "float32", Desc: "触发框宽 / 下拉最大高"},
			{Name: "Placeholder", Type: "string", Desc: "未选时占位文字"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnChange", Type: "func(int)", Desc: "选择回调"},
			{Name: "SetSelected/Close", Type: "func(int) / func()", Desc: "程序化选择 / 收起下拉"},
		},
	})
}
