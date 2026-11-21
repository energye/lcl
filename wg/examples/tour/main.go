package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）；目标锚节点由 Build 闭包捕获后注入 Steps。
var tour = widgets.NewTour(nil)
var tourCtx *ui.BuildContext // Build 闭包捕获（探针/程序化 Start 用）
var (
	b1, b2, b3 *widgets.Button // 目标按钮池化实例（教训⑨第五次）
)

// demoBuild 构造目标行+注入 Steps（每次重建注入最新节点）。
func demoBuild(ctx *ui.BuildContext) *scene.Node {
	tourCtx = ctx
	// 目标按钮池化：Build 内新建实例按压态随重建丢失（教训⑨第五次）。
	if b1 == nil {
		b1 = widgets.NewButton("目标一：创建")
		b2 = widgets.NewButton("目标二：审核")
		b3 = widgets.NewButton("目标三：发布")
	}
	nb1 := b1.Build(ctx)
	nb2 := b2.Build(ctx)
	nb3 := b3.Build(ctx)
	row := scene.Flex(scene.Horizontal, 16).SetCrossAlign(scene.CrossCenter)
	row.Add(nb1)
	row.Add(nb2)
	row.Add(nb3)
	note := (&widgets.Label{Text: "遮罩挖洞露出目标按钮，卡片贴目标下方"}).Build(ctx)
	tour.Steps = []widgets.TourStep{
		{Title: "第一步：创建", Desc: "点击「创建」开始新任务。", TargetNode: nb1},
		{Title: "第二步：审核", Desc: "任务创建后进入审核队列。", TargetNode: nb2},
		{Title: "第三步：发布", Desc: "审核通过即可发布上线。", TargetNode: nb3},
	}
	col := scene.Flex(scene.Vertical, 24)
	col.SetPadding(geom.InLTRB(24, 24, 24, 24))
	col.Add(row)
	col.Add(note)
	col.Add(tour.Build(ctx))
	return col
}

func main() {
	launcher.Main(doc.Meta{
		ID: "tour", Title: "Tour 引导", Group: doc.GroupData,
		What: []string{
			"**target 锚定卡**：遮罩挖洞（四块矩形绕开目标 rect，antd 遮罩 rgba(0,0,0,0.45)）+卡片贴目标下方（空间不足翻上方、横向钳制）",
			"TargetNode 由 demo Build 闭包捕获目标组件节点直接传递；无 target 步骤卡片居中",
			"步骤切换 AfterFunc(0) 重开防御（TreeSelect 轮教训）；Esc/完成退出",
		},
		Demos: []doc.Demo{
			{Title: "锚定三步引导", Note: "遮罩挖洞露出目标按钮；卡片贴目标下方；下一步跳下一目标", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(*ui.BuildContext) *scene.Node { return demoBuild(ctx) })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTour", Type: "func(steps []TourStep) *Tour", Desc: "构造"},
			{Name: "Steps", Type: "[]TourStep{Title,Desc,TargetNode}", Desc: "步骤（TargetNode 锚定目标）"},
			{Name: "Start / Close / IsOpen", Type: "func(...)", Desc: "程序化开启/关闭/查询"},
			{Name: "OnFinish", Type: "func()", Desc: "完成回调"},
			{Name: "Width", Type: "float32", Default: "300", Desc: "卡片宽"},
		},
	})
}
