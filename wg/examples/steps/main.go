package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "steps", Title: "Steps 步骤条", Group: doc.GroupDisplay,
		What: []string{
			"分步引导：当前步及之前高亮，后续置灰",
			"垂直方向（direction=vertical）：步骤纵排、连接线竖直、标签右侧",
			"自定义图标（items.icon）：Icons 按步设图标名，32×32 容器状态着色（process 主色实底白 icon/finish 主色 icon/wait 灰/error 红）",
		},
		Demos: []doc.Demo{
			{Title: "四步进度", Note: "Current=2：第三步进行中", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.NewSteps([]string{"选择", "配置", "确认", "完成"}, 2)
			}},
			{Title: "error 状态", Note: "第二步出错（红圆+×，antd status=error）", Build: func(ctx *ui.BuildContext) ui.Widget {
				st := widgets.NewSteps([]string{"选择", "配置", "确认", "完成"}, 2)
				st.ErrorStep = 1
				return st
			}},
			{Title: "垂直方向", Note: "Vertical=true：步骤纵排、连接线竖直、标签右侧", Build: func(ctx *ui.BuildContext) ui.Widget {
				st := widgets.NewSteps([]string{"选择", "配置", "确认", "完成"}, 2)
				st.Vertical = true
				return st
			}},
			{Title: "自定义图标", Note: "Icons=[user,setting,shopping-cart,check-circle]：32×32 容器随状态着色（antd steps-demo-icon）", Build: func(ctx *ui.BuildContext) ui.Widget {
				st := widgets.NewSteps([]string{"账户", "设置", "支付", "完成"}, 2)
				st.Icons = []string{"user", "setting", "shopping-cart", "check-circle"}
				return st
			}},
		},
		API: []doc.APIRow{
			{Name: "Labels", Type: "[]string", Desc: "步骤名"},
			{Name: "Current", Type: "int", Default: "0", Desc: "当前步下标（0 起）"},
			{Name: "ErrorStep", Type: "int", Default: "-1", Desc: "错误步下标（红圆+×，antd status=error）"},
			{Name: "Icons", Type: "[]string", Desc: "每步自定义图标名（antd items.icon；空=默认数字/✓/×）"},
		},
	})
}
