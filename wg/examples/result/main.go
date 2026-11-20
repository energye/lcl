package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "result", Title: "Result 结果页", Group: doc.GroupFeedback,
		What: []string{
			"四状态大图标：success 绿 / error 红 / warning 黄 / info 主色（语义色圆底）",
			"标题/副标题两级文案 + Extra 操作区（Button 组件组合复用）",
		},
		Demos: []doc.Demo{
			{Title: "成功与错误", Note: "大图标+标题+副标题+操作区", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 16)
				ok := &widgets.Result{
					Status:   widgets.ResultSuccess,
					Title:    "提交成功",
					SubTitle: "表单内容已保存，可在历史记录中查看。",
					Extra: []ui.Widget{
						widgets.NewButton("查看记录"),
						widgets.NewButton("继续填写"),
					},
				}
				bad := &widgets.Result{
					Status:   widgets.ResultError,
					Title:    "提交失败",
					SubTitle: "网络连接超时，请检查网络后重试。",
					Extra:    []ui.Widget{widgets.NewButton("重试")},
				}
				col.Add(ok.Build(ctx))
				col.Add(bad.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "警告与信息", Note: "无操作区的精简形态", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 16)
				col.Add((&widgets.Result{
					Status: widgets.ResultWarning,
					Title:  "存储空间告急",
				}).Build(ctx))
				col.Add((&widgets.Result{
					Status: widgets.ResultInfo,
					Title:  "新版本已就绪",
				}).Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "Status", Type: "ResultStatus", Desc: "Success/Error/Warning/Info"},
			{Name: "Title/SubTitle", Type: "string", Desc: "标题 / 副标题"},
			{Name: "Extra", Type: "[]ui.Widget", Desc: "操作区（通常 Button 组合复用）"},
			{Name: "Width", Type: "float32", Default: "420", Desc: "整列宽度"},
		},
	})
}
