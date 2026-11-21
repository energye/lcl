package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var (
	tfBase = widgets.NewTransfer([]widgets.TransferItem{
		{Key: "a", Title: "内容一"},
		{Key: "b", Title: "内容二"},
		{Key: "c", Title: "内容三"},
		{Key: "d", Title: "内容四"},
		{Key: "e", Title: "内容五"},
	})
	tfInit = widgets.NewTransfer([]widgets.TransferItem{
		{Key: "a", Title: "内容一"},
		{Key: "b", Title: "内容二"},
		{Key: "c", Title: "内容三"},
		{Key: "d", Title: "内容四"},
	})
	tfLog   = ""
	tfInited bool
)

func main() {
	launcher.Main(doc.Meta{
		ID: "transfer", Title: "Transfer 穿梭框", Group: doc.GroupControl,
		What: []string{
			"双栏（源/目标）+ 中间操作列（> < » «——Button 组合复用+被动锚唯一锚）",
			"列表项 **Checkbox 组件勾选**（池化实例——Tree.Checkable 先例）+ 点行整体切换 + 标题行勾选/总数计数",
			"TargetKeys 驱动左右分栏；OnChange(targetKeys, direction, moveKeys)",
		},
		Demos: []doc.Demo{
			{Title: "基础", Note: "勾选左侧项 → > 右移；目标栏项 ← 左移；» 全右", Build: func(ctx *ui.BuildContext) ui.Widget {
				tfBase.OnChange = func(keys []string, dir widgets.TransferDirection, move []string) {
					tfLog = fmt.Sprintf("moved %v", move)
				}
				return tfBase
			}},
			{Title: "带初值", Note: "TargetKeys=[b,d]：内容二/四在右栏", Build: func(ctx *ui.BuildContext) ui.Widget {
				if !tfInited {
					tfInit.TargetKeys = []string{"b", "d"}
					tfInited = true
				}
				return tfInit
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTransfer", Type: "func(items []TransferItem) *Transfer", Desc: "构造"},
			{Name: "DataSource", Type: "[]TransferItem", Desc: "数据（Key/Title）"},
			{Name: "TargetKeys", Type: "[]string", Desc: "右栏 Key 集（驱动分栏）"},
			{Name: "OnChange", Type: "func(keys, dir, moveKeys)", Desc: "移动回调（右移/左移/全右/全左）"},
			{Name: "Width", Type: "float32", Default: "220", Desc: "单栏宽"},
			{Name: "move / moveAll", Type: "func(dir)", Desc: "程序化移动（勾选项 / 全部）"},
		},
	})
}
