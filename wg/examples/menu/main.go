package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var vSub = widgets.NewMenu(
	&widgets.MenuItem{Key: "grp", Label: "项目分组", Group: true,
		Children: []*widgets.MenuItem{
			{Key: "g1", Label: "成员一"},
			{Key: "g2", Label: "成员二"},
		}},
	&widgets.MenuItem{Key: "nav", Label: "导航配置", Icon: "setting",
		Children: []*widgets.MenuItem{
			{Key: "nav-a", Label: "布局设置"},
			{Key: "nav-b", Label: "主题设置"},
		}},
	&widgets.MenuItem{Key: "about", Label: "关于", Icon: "info-circle"},
)

var (
	hMenu = widgets.NewMenu(
		&widgets.MenuItem{Key: "home", Label: "首页", Icon: "home"},
		&widgets.MenuItem{Key: "docs", Label: "文档", Icon: "edit"},
		&widgets.MenuItem{Key: "settings", Label: "设置", Icon: "setting",
			Children: []*widgets.MenuItem{
				{Key: "set-a", Label: "个人设置"},
				{Key: "set-b", Label: "系统设置", Children: []*widgets.MenuItem{
					{Key: "set-b-1", Label: "权限管理"},
					{Key: "set-b-2", Label: "操作日志"},
				}},
			}},
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
	vSub.Mode = widgets.MenuVertical
	vSub.Selected = "nav-a"
	vSub.OpenKeys = map[string]bool{"nav": true}
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
			{Title: "子菜单与分组", Note: "纵向内联展开：点击父项切换展开（▾/▸）；分组标题置灰不可点", Build: func(ctx *ui.BuildContext) ui.Widget {
				return vSub
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
