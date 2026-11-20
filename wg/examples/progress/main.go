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
		ID: "progress", Title: "Progress 进度条", Group: doc.GroupControl,
		What: []string{
			"线形进度条：M8 对齐 antd——status 语义色（success/exception）、满进度自动转绿、showInfo 百分比",
			"未实现（M10 计划）：circle/dashboard 形态、trail 色自定义",
		},
		Demos: []doc.Demo{
			{Title: "四档进度", Note: "0 / 25% / 60% / 100%（满进度自动转 success 绿）；末行用默认尺寸（高 8）", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10).SetWidth(480)
				for _, v := range []float32{0, 0.25, 0.6, 1} {
					col.Add((&widgets.Progress{Value: v, Width: 400, Height: 10}).Build(ctx))
				}
				col.Add((&widgets.Progress{Value: 0.3}).Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "status 与百分比", Note: "showInfo 显示右侧百分比；exception 红", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10).SetWidth(480)
				ok := widgets.Progress{Value: 0.7, Width: 400, Height: 10, ShowInfo: true}
				bad := widgets.Progress{Value: 0.4, Width: 400, Height: 10, ShowInfo: true, Status: widgets.ProgressException}
				col.Add(ok.Build(ctx))
				col.Add(bad.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "circle 圆环", Note: "Shape=ProgressCircle：灰轨+主色进度弧+中心百分比（CustomPaint 弧线）", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 24)
				row.Add((&widgets.Progress{Value: 0.75, Shape: widgets.ProgressCircle, ShowInfo: true}).Build(ctx))
				row.Add((&widgets.Progress{Value: 1, Shape: widgets.ProgressCircle, ShowInfo: true}).Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "dashboard 仪表盘", Note: "Shape=ProgressDashboard：270° 缺口朝下（antd dashboard）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					return (&widgets.Progress{Value: 0.6, Shape: widgets.ProgressDashboard, ShowInfo: true}).Build(c)
				})
			}},
		},
		API: []doc.APIRow{
			{Name: "Value", Type: "float32", Default: "0", Desc: "进度 0..1"},
			{Name: "Width/Height", Type: "float32", Desc: "轨道尺寸"},
			{Name: "Status", Type: "widgets.ProgressStatus", Default: "Normal", Desc: "Normal/Success/Exception（满进度自动 Success）"},
			{Name: "ShowInfo", Type: "bool", Default: "false", Desc: "右侧（线形）/中心（圆环）百分比文字"},
			{Name: "Shape", Type: "ProgressShape", Default: "Line", Desc: "Line/Circle/Dashboard 三形态"},
			{Name: "StrokeWidth", Type: "float32", Default: "6", Desc: "圆环描边宽（antd strokeWidth）"},
		},
	})
}
