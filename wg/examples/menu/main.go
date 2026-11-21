package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	hMenu = widgets.NewMenu(
		&widgets.MenuItem{Key: "home", Label: "首页", Icon: "home"},
		&widgets.MenuItem{Key: "docs", Label: "文档", Icon: "edit"},
		&widgets.MenuItem{Key: "settings", Label: "设置", Icon: "setting"},
	)
	vMenu = widgets.NewMenu(
		&widgets.MenuItem{Key: "mail", Label: "邮件", Icon: "edit"},
		&widgets.MenuItem{Key: "files", Label: "文件", Icon: "copy"},
		&widgets.MenuItem{Key: "users", Label: "用户", Icon: "user"},
	)
)

func init() {
	hMenu.Mode = widgets.MenuHorizontal
	hMenu.Selected = "home"
	vMenu.Mode = widgets.MenuVertical
	vMenu.Selected = "mail"
}

func main() {
	launcher.Main(doc.Meta{
		ID: "menu", Title: "Menu 导航菜单", Group: doc.GroupControl,
		What: []string{
			"横向模式（顶部导航）：主色文字+底部 2px 指示条（antd 横向菜单）",
			"纵向模式（侧边导航）：选中左侧 3px 主色指示条+浅主色底（InfoBg）",
			"图标组合复用 Icon 组件；OnClick 点击回调",
		},
		Demos: []doc.Demo{
			{Title: "横向菜单", Note: "Mode=Horizontal：点击切换选中（指示条跟随）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return hMenu
			}},
			{Title: "纵向菜单", Note: "Mode=Vertical：宽 200，选中浅主色底+左指示条", Build: func(ctx *ui.BuildContext) ui.Widget {
				return vMenu
			}},
		},
		API: []doc.APIRow{
			{Name: "NewMenu", Type: "func(items ...*MenuItem) *Menu", Desc: "构造"},
			{Name: "Items", Type: "[]*MenuItem{Key,Label,Icon}", Desc: "菜单项（Icon 组合复用 Icon 组件）"},
			{Name: "Mode", Type: "MenuMode", Desc: "Horizontal 横向 / Vertical 纵向"},
			{Name: "Selected", Type: "string", Desc: "选中项 Key（空=首项）"},
			{Name: "OnClick", Type: "func(key string)", Desc: "点击回调"},
			{Name: "Width", Type: "float32", Default: "200", Desc: "纵向模式宽度"},
		},
	})
}
