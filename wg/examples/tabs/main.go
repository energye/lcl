package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	tabs *widgets.Tabs
	// cardTabs card 型（包级持有：实例池跨重建）。
	cardTabs = widgets.NewTabs(
		widgets.Tab{Key: "c1", Label: "文档", Content: lazyLabel("文档内容")},
		widgets.Tab{Key: "c2", Label: "设置", Content: lazyLabel("设置内容")},
		widgets.Tab{Key: "c3", Label: "关于", Content: lazyLabel("关于内容")},
	)
	// closableTabs 可关闭页签（OnEdit 真实移除）。
	closableTabs = widgets.NewTabs(
		widgets.Tab{Key: "e1", Label: "标签一", Content: lazyLabel("标签一内容"), Closable: true},
		widgets.Tab{Key: "e2", Label: "标签二", Content: lazyLabel("标签二内容"), Closable: true},
		widgets.Tab{Key: "e3", Label: "固定页", Content: lazyLabel("固定页内容")},
	)
)

func init() {
	cardTabs.Type = widgets.TabsCard
	closableTabs.OnEdit = func(key ui.Key, action string) {
		if action != "remove" {
			return
		}
		for i, it := range closableTabs.Items {
			if it.Key == key {
				closableTabs.Items = append(closableTabs.Items[:i], closableTabs.Items[i+1:]...)
				break
			}
		}
		if closableTabs.Active >= len(closableTabs.Items) {
			closableTabs.Active = len(closableTabs.Items) - 1
		}
		closableTabs.MarkDirty()
	}
}

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
			"card 卡片型：灰底描边方块页签、选中白底+顶部主色线（antd type=card）",
			"closable 可关闭：页签右侧 × 热区点击触发 OnEdit(remove)",
		},
		Demos: []doc.Demo{
			{Title: "三页切换", Note: "line 线型（默认）：主色文字+底部指示条", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tabs
			}},
			{Title: "card 卡片型", Note: "Type=TabsCard：灰底描边方块、选中白底+顶部主色线", Build: func(ctx *ui.BuildContext) ui.Widget {
				return cardTabs
			}},
			{Title: "closable 可关闭", Note: "页签右侧 × 热区点击移除（OnEdit 回调）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return closableTabs
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTabs", Type: "func(items ...Tab) *Tabs", Desc: "构造；Tab{Key,Label,Content}"},
			{Name: "Items", Type: "[]Tab", Desc: "标签页集合"},
			{Name: "Active", Type: "int", Desc: "当前页下标"},
			{Name: "OnChange", Type: "func(int)", Desc: "切换回调"},
			{Name: "Type", Type: "TabsType", Desc: "TabsLine 线型 / TabsCard 卡片型"},
			{Name: "Tab.Closable", Type: "bool", Desc: "页签右侧 × 可关闭"},
			{Name: "OnEdit", Type: "func(key, action)", Desc: "页签编辑回调（closable × 触发 remove）"},
		},
	})
}
