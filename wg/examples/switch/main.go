package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 展示实例包级持有：demo 闭包每次重建都会执行，闭包内新建会把
// 切换/hover 状态随重建丢失（--audit 实测"点了没反应"）。
var (
	sw    = widgets.NewSwitch(true)
	lsw   = widgets.NewSwitch(true)
	dis   = widgets.NewSwitch(false)
	small = widgets.NewSwitch(true)
	words = widgets.NewSwitch(true)
)

func init() {
	lsw.Loading = true
	lsw.Disabled = true
	dis.Disabled = true
	small.Small = true
	words.CheckedChildren = "启用"
	words.UnCheckedChildren = "停用"
}

func main() {
	launcher.Main(doc.Meta{
		ID: "switch", Title: "Switch 开关", Group: doc.GroupControl,
		What: []string{
			"两态开关：轨道+滑块全 scene 图元绘制；M8 对齐 antd——checkedChildren/unCheckedChildren 文案、loading 态、small 尺寸",
			"文案显示在滑块对侧（OnPrimary 色）；loading 时显示旋转指示并忽略点击",
		},
		Demos: []doc.Demo{
			{Title: "基本 / 尺寸", Note: "默认 36×20；Small 24×12", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(sw.Build(ctx))
				row := scene.Flex(scene.Horizontal, 12)
				row.Add(dis.Build(ctx))
				row.Add(small.Build(ctx))
				col.Add(row)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "文案与 loading", Note: "On 开显示「启用」、关显示「停用」；loading 旋转且不可点", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(words.Build(ctx))
				col.Add(lsw.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewSwitch", Type: "func(on bool) *Switch", Desc: "构造"},
			{Name: "On", Type: "bool", Desc: "开关状态"},
			{Name: "CheckedChildren/UnCheckedChildren", Type: "string", Desc: "开/关态文案（antd 同名）"},
			{Name: "Loading", Type: "bool", Default: "false", Desc: "加载中：旋转指示、忽略点击"},
			{Name: "Small", Type: "bool", Default: "false", Desc: "小尺寸（24×12，antd size=small）"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnChange", Type: "func(bool)", Desc: "切换回调"},
		},
	})
}
