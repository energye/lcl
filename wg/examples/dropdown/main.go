package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	trigger    = widgets.NewButton("悬停下拉")
	ddStatus   = widgets.NewButton("未选择")
	lastChoice = "未选择"
)

func init() {
	trigger.OnClick = func() {}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "dropdown", Title: "Dropdown 下拉菜单", Group: doc.GroupControl,
		What: []string{
			"悬停触发：Enter 显示浮层/Leave 隐藏（Dismiss(0) 不吞后续点击，Tooltip 同机制）",
			"BandPopup 层带+PlacementBottomStart；菜单项点击回调并关闭浮层",
			"白底圆角+阴影浮层（Menu 纵向项同款形态）",
		},
		Demos: []doc.Demo{
			{Title: "悬停下拉", Note: "悬停按钮显示菜单项；点击项回调并关闭", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 12)
				dd := widgets.NewDropdown(trigger,
					&widgets.MenuItem{Key: "copy", Label: "复制", Icon: "copy"},
					&widgets.MenuItem{Key: "edit", Label: "编辑", Icon: "edit"},
					&widgets.MenuItem{Key: "del", Label: "删除", Icon: "delete"},
				)
				dd.OnClick = func(key string) {
					lastChoice = "已选择：" + key
					ddStatus.Text = lastChoice
					ddStatus.MarkDirty()
				}
				row := scene.Flex(scene.Horizontal, 16)
				row.Add(dd.Build(ctx))
				row.Add(ddStatus.Build(ctx))
				col.Add(row)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewDropdown", Type: "func(child ui.Widget, items ...*MenuItem) *Dropdown", Desc: "构造（Items 复用 MenuItem 结构）"},
			{Name: "OnClick", Type: "func(key string)", Desc: "菜单项点击回调（点击后浮层自动关闭）"},
			{Name: "Close/IsOpen", Type: "func() / func() bool", Desc: "立即关闭 / 状态查询"},
		},
	})
}
