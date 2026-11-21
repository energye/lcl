package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 包级持有：按压态（pressed）在实例上——闭包内新建会让按压视觉随
// 重建丢失（教训⑨变体：状态标志在实例）。
var (
	fbPrimary = &widgets.FloatButton{Icon: "edit", OnClick: func() {}}
	fbDefault = &widgets.FloatButton{Icon: "user", Type: widgets.FloatDefault, Size: 40}
)

func main() {
	launcher.Main(doc.Meta{
		ID: "floatbutton", Title: "FloatButton 悬浮按钮", Group: doc.GroupControl,
		What: []string{
			"primary 主色实底白图标 / default 白底描边，圆形 56/40 两档（antd FloatButton）",
			"FloatButtonGroup 纵向排列；ShapeGroup 组底色（antd shape=square）",
			"图标组合复用 Icon 组件；OnClick 点击回调",
		},
		Demos: []doc.Demo{
			{Title: "双型", Note: "primary 主色实底 / default 白底描边，直径 56/40", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 18)
				row.Add(fbPrimary.Build(ctx))
				row.Add(fbDefault.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "按钮组", Note: "ShapeGroup=true 组底色：三个悬浮按钮纵排", Build: func(ctx *ui.BuildContext) ui.Widget {
				grp := &widgets.FloatButtonGroup{ShapeGroup: true}
				grp.Buttons = []*widgets.FloatButton{
					{Icon: "edit", OnClick: func() {}},
					{Icon: "copy", OnClick: func() {}},
					{Icon: "delete", OnClick: func() {}},
				}
				return grp
			}},
		},
		API: []doc.APIRow{
			{Name: "Icon", Type: "string", Desc: "图标名（Icon 组件组合复用）"},
			{Name: "Type", Type: "FloatButtonType", Desc: "Primary 主色实底 / Default 白底描边"},
			{Name: "Size", Type: "float32", Default: "56", Desc: "直径（小档 40）"},
			{Name: "Tooltip", Type: "string", Desc: "悬停提示文案"},
			{Name: "OnClick", Type: "func()", Desc: "点击回调"},
			{Name: "FloatButtonGroup", Type: "struct{Buttons; Gap; ShapeGroup}", Desc: "纵向按钮组"},
		},
	})
}
