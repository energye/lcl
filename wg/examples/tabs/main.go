package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var tabs *widgets.Tabs

// lazyLabel 把 Label 的 Build 推迟到拿到 ctx 时。
func lazyLabel(text string) ui.Widget {
	return doc.Node(func(ctx *ui.BuildContext) *scene.Node {
		return (&widgets.Label{Text: text}).Build(ctx)
	})
}

func init() {
	tabs = widgets.NewTabs(
		widgets.Tab{Key: "t1", Label: "第一页", Content: lazyLabel("第一页内容")},
		widgets.Tab{Key: "t2", Label: "第二页", Content: lazyLabel("第二页内容")},
		widgets.Tab{Key: "t3", Label: "第三页", Content: lazyLabel("第三页内容")},
	)
}

func main() {
	launcher.Main(doc.Meta{
		ID: "tabs", Title: "Tabs 标签页", Group: doc.GroupControl,
		What: []string{
			"标签页：点击标签切换内容区，当前实现会重建内容子树",
			"计划（M11）：对齐 antd——line/card 类型、closable、切换动画",
		},
		Demos: []doc.Demo{
			{Title: "三页切换", Note: "点击标签切换；Active/OnChange 可受控", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tabs
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTabs", Type: "func(items ...Tab) *Tabs", Desc: "构造；Tab{Key,Label,Content}"},
			{Name: "Items", Type: "[]Tab", Desc: "标签页集合"},
			{Name: "Active", Type: "int", Desc: "当前页下标"},
			{Name: "OnChange", Type: "func(int)", Desc: "切换回调"},
		},
	})
}
