package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var tour = widgets.NewTour([]widgets.TourStep{
	{Title: "欢迎", Desc: "这是主工作台——左侧为功能导航，右侧为内容区。"},
	{Title: "创建任务", Desc: "点击顶部「新建」按钮即可创建你的第一个任务。"},
	{Title: "完成", Desc: "引导结束——随时点「开始引导」重新查看。"},
})

func main() {
	launcher.Main(doc.Meta{
		ID: "tour", Title: "Tour 引导", Group: doc.GroupData,
		What: []string{
			"**BandModal 居中步骤卡**（遮罩+标题/描述/步骤指示 x/y+上一步/下一步/完成）",
			"步骤切换 **AfterFunc(0) 重开面板**（TreeSelect 轮教训——重开防御统一）；最后一步变「完成」",
			"Esc/完成/关闭退出；首轮居中卡片——target 锚定卡（贴目标+遮罩镂空）未实现",
		},
		Demos: []doc.Demo{
			{Title: "三步引导", Note: "点「开始引导」→ 三步卡片 → 完成回调", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tour
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTour", Type: "func(steps []TourStep) *Tour", Desc: "构造"},
			{Name: "Steps", Type: "[]TourStep{Title,Desc}", Desc: "引导步骤"},
			{Name: "Start / Close / IsOpen", Type: "func(...)", Desc: "程序化开启/关闭/查询"},
			{Name: "OnFinish", Type: "func()", Desc: "完成回调"},
			{Name: "Width", Type: "float32", Default: "300", Desc: "卡片宽"},
		},
	})
}
