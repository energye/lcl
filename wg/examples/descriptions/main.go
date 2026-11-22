package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var (
	descBordered = widgets.NewDescriptions([]widgets.DescriptionItem{
		{Label: "用户名", Content: "张三"},
		{Label: "城市", Content: "杭州"},
		{Label: "手机号", Content: "138-1234-5678"},
		{Label: "邮箱", Content: "zhangsan@example.com"},
		{Label: "性别", Content: "男"},
		// Widget 内容（antd children 任意节点——深化轮）。
		{Label: "状态", Widget: &widgets.Tag{Text: "活跃", Color: geom.RGBA(82, 196, 26)}},
	})
	descPlain = widgets.NewDescriptions([]widgets.DescriptionItem{
		{Label: "创建时间", Content: "2026-11-21 10:24"},
		{Label: "版本", Content: "v1.0.0"},
		{Label: "状态", Content: "运行中"},
		{Label: "负责人", Content: "李四"},
	})
	descVert = func() *widgets.Descriptions {
		d := widgets.NewDescriptions([]widgets.DescriptionItem{
			{Label: "实例 ID", Content: "i-2ze8f3a1"},
			{Label: "规格", Content: "ecs.g7.xlarge"},
			{Label: "创建时间", Content: "2026-11-20 08:30"},
		})
		d.Title = "实例详情（竖排）"
		d.Vertical = true
		d.Bordered = true
		d.Width = 400
		return d
	}()
	descSpan = func() *widgets.Descriptions {
		d := widgets.NewDescriptions([]widgets.DescriptionItem{
			{Label: "备注", Content: "跨双列的长备注内容——span=2", Span: 2},
			{Label: "配置", Content: "8C16G"},
			{Label: "系统", Content: "Ubuntu 24.04"},
			{Label: "操作", Content: "", Span: 2,
				Extra: &widgets.Button{Text: "编辑", Type: widgets.TypeLink, Size: ui.SizeSmall}},
		})
		d.Title = "跨列与操作区"
		d.Columns = 2
		d.Bordered = true
		d.Width = 560
		return d
	}()
	descSmall = func() *widgets.Descriptions {
		d := widgets.NewDescriptions([]widgets.DescriptionItem{
			{Label: "区域", Content: "cn-hangzhou"},
			{Label: "可用区", Content: "cn-hangzhou-a"},
		})
		d.Size = widgets.DescSizeSmall
		d.Columns = 2
		d.Bordered = true
		d.Width = 400
		return d
	}()
)

func init() {
	descBordered.Title = "用户信息"
	descBordered.Columns = 2
	descBordered.Bordered = true
	descBordered.Width = 560
	descPlain.Columns = 2
	descPlain.Width = 560
}

func main() {
	launcher.Main(doc.Meta{
		ID: "descriptions", Title: "Descriptions 描述列表", Group: doc.GroupData,
		What: []string{
			"标题行 + 网格布局（label 格+content 格交替，**Flex 行支持 span 加权**）",
			"双形态：Bordered（label 灰底格+描边，antd bordered）/默认（label 灰字冒号+疏排）",
			"Vertical 竖排（label 上 content 下）/ Size 三档行高 / Item.Span 跨列 / Extra 操作区",
		},
		Demos: []doc.Demo{
			{Title: "带边框", Note: "Bordered=true：label 灰底格+描边网格（column=2）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return descBordered
			}},
			{Title: "默认无边框", Note: "label 灰字冒号+正文疏排", Build: func(ctx *ui.BuildContext) ui.Widget {
				return descPlain
			}},
			{Title: "竖排", Note: "Vertical=true：label 上 content 下（antd layout=vertical）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return descVert
			}},
			{Title: "跨列与操作区", Note: "Item.Span=2 占双列；Item.Extra 操作链接（antd extra）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return descSpan
			}},
			{Title: "小尺寸", Note: "Size=small：行高 24（antd size）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return descSmall
			}},
		},
		API: []doc.APIRow{
			{Name: "NewDescriptions", Type: "func(items []DescriptionItem) *Descriptions", Desc: "构造"},
			{Name: "Items", Type: "[]DescriptionItem{Label,Content,Widget,Span,Extra}", Desc: "描述项"},
			{Name: "Title / Extra", Type: "string / ui.Widget", Desc: "标题 / 全局操作区（标题行右侧）"},
			{Name: "Columns", Type: "int", Default: "2", Desc: "每行组数（antd column）"},
			{Name: "Bordered", Type: "bool", Default: "false", Desc: "边框形态（label 灰底格+网格描边）"},
			{Name: "Vertical", Type: "bool", Default: "false", Desc: "竖排（antd layout=vertical）"},
			{Name: "Size", Type: "DescriptionsSize", Default: "Default", Desc: "行高三档（Default 34/Middle 30/Small 24）"},
			{Name: "Item.Span", Type: "int", Default: "1", Desc: "跨列数（antd span，行尾截断）"},
			{Name: "Width", Type: "float32", Default: "560", Desc: "总宽"},
		},
	})
}
