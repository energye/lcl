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

	grouped = widgets.NewSelect(nil)
	multi   = widgets.NewSelect(nil)
	mStat   = widgets.NewButton("多选：[]")
)

func init() {
	sel.Width = 260
	sel.OnChange = func(i int) {
		status.Text = fmt.Sprintf("已选：%d · %s", i, sel.Options[i])
		status.MarkDirty()
	}
	grouped.Width, grouped.ShowSearch = 260, true
	grouped.Items = []widgets.SelectOption{
		{Label: "软件光栅化", Group: "渲染后端"},
		{Label: "GDI 混合", Group: "渲染后端"},
		{Label: "预乘 BGRA", Group: "像素格式"},
		{Label: "ClearType", Group: "文字通道"},
	}
	multi.Width, multi.Multiple = 260, true
	multi.Options = []string{"渲染", "布局", "输入", "浮层", "主题"}
	multi.Values = []string{"渲染", "输入"}
	multi.OnValues = func(vs []string) {
		multi.Values = vs
		mStat.Text = fmt.Sprintf("多选：%v", vs)
		mStat.MarkDirty()
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "select", Title: "Select 下拉选择", Group: doc.GroupData,
		What: []string{
			"下拉选择：弹层走 overlay.Host（BandPopup），OutsideClick/Esc 关闭",
			"M8 对齐 antd：showSearch 搜索过滤、multiple 多选（选中不收起）、分组标题行",
			"未实现：label/value 分离（Label 即 Value）、异步加载、虚拟滚动、option 自定义渲染",
		},
		Demos: []doc.Demo{
			{Title: "基本单选", Note: "点开下拉选择；选中回显", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(sel.Build(ctx))
				col.Add(status.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "分组 + 搜索", Note: "ShowSearch 过滤 contains；Group 标题行置灰", Build: func(ctx *ui.BuildContext) ui.Widget {
				return grouped
			}},
			{Title: "多选", Note: "选中行显示对勾、选中不收起；触发框逗号连接", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(multi.Build(ctx))
				col.Add(mStat.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewSelect", Type: "func(options []string) *Select", Desc: "构造（单选字符串项）"},
			{Name: "Items", Type: "[]SelectOption{Label, Group}", Desc: "带分组的选项（设置后优先于 Options）"},
			{Name: "ShowSearch", Type: "bool", Default: "false", Desc: "弹层顶部搜索行（antd showSearch）"},
			{Name: "Multiple", Type: "bool", Default: "false", Desc: "多选（antd multiple；选中不收起）"},
			{Name: "Selected/Values", Type: "int / []string", Desc: "单选下标 / 多选集合"},
			{Name: "Width/MaxHeight", Type: "float32", Default: "180 / 240", Desc: "触发框宽 / 弹层最大高"},
			{Name: "Placeholder", Type: "string", Default: "请选择", Desc: "未选占位"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnChange/OnValues", Type: "回调", Desc: "单选下标 / 多选集合"},
			{Name: "SetSelected/Close", Type: "func(int) / func()", Desc: "程序化选择 / 收起"},
		},
	})
}
