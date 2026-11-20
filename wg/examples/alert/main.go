package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// closableDemo 包级持有：闭包内新建会让状态随重建丢失（教训⑨）。
var closableAlive = true

func main() {
	launcher.Main(doc.Meta{
		ID: "alert", Title: "Alert 提醒", Group: doc.GroupFeedback,
		What: []string{
			"四类型语义：info 主色底 / success 绿底 / warning 黄底 / error 红底（antd colorXxxBg）",
			"showIcon 左侧类型图标（Icon 组件组合复用）+ 标题/描述两级文案",
			"closable 可关闭：右上 × 点击触发 OnClose",
		},
		Demos: []doc.Demo{
			{Title: "四类型", Note: "info/success/warning/error 浅色语义底+图标", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				for _, a := range []widgets.AlertType{
					widgets.AlertInfo, widgets.AlertSuccess, widgets.AlertWarning, widgets.AlertError,
				} {
					col.Add((&widgets.Alert{
						Type:     a,
						Title:    "提示标题",
						ShowIcon: true,
						Width:    420,
					}).Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "带描述与关闭", Note: "标题+描述两级文案；点 × 真实收起提示", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				if closableAlive {
					a := &widgets.Alert{
						Type:        widgets.AlertWarning,
						Title:       "磁盘空间不足",
						Description: "系统盘剩余 2.1 GB，建议清理临时文件后重试。",
						ShowIcon:    true,
						Closable:    true,
						Width:       420,
					}
					a.OnClose = func() {
						closableAlive = false
						// 整页重建由演示框架的下一帧处理；这里标记组件树需要重绘。
						a.MarkDirty()
					}
					col.Add(a.Build(ctx))
				} else {
					col.Add((&widgets.Label{Text: "提示已关闭"}).Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "Type", Type: "AlertType", Desc: "Info/Success/Warning/Error（语义底色+图标）"},
			{Name: "Title/Description", Type: "string", Desc: "主文案 / 副文案（弱化色）"},
			{Name: "ShowIcon", Type: "bool", Default: "false", Desc: "左侧类型图标（antd showIcon）"},
			{Name: "Closable/OnClose", Type: "bool / func()", Desc: "右上 × 可关闭与回调"},
			{Name: "Width", Type: "float32", Default: "420", Desc: "提示条宽度"},
		},
	})
}
