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

// lvSel label-value 分离演示：显示语言名，提交语言代码。
// 提交值文本用 Button.Text+MarkDirty（form demo 校验结果同款已验证通道：
// controlBase 标准 attach 的重建链在审计时序下可靠；TextView/Label 通道
// 实测读旧值）。
var (
	lvSel  = widgets.NewSelect(nil)
	lvStat = widgets.NewButton("提交值：go")
)

func init() {
	lvSel.Width = 260
	lvSel.Selected = 0 // 初始选中 Go：触发框显示唯一文本，供审计锚定位
	lvSel.Items = []widgets.SelectOption{
		{Label: "Go", Value: "go"},
		{Label: "Python", Value: "py"},
		{Label: "Rust", Value: "rust"},
	}
	lvSel.OnChange = func(i int) {
		lvStat.Text = "提交值：" + lvSel.Value()
		lvStat.MarkDirty()
	}
}

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
			"label/value 分离：Label 显示、Value 提交（SelectOption.Value，空=回退 Label）",
		},
		Demos: []doc.Demo{
			{Title: "label/value 分离", Note: "显示 Go/Python/Rust，提交 go/py/rust（SelectOption.Value）", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(lvSel.Build(ctx))
				col.Add(lvStat.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
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
			{Name: "Items", Type: "[]SelectOption{Label, Value, Group}", Desc: "带分组/提交值的选项（Value 空=回退 Label）"},
			{Name: "Value/OptionValue", Type: "func() string / func(i int) string", Desc: "选中项提交值 / 第 i 项提交值（label-value 分离）"},
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
