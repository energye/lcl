package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var collapse = widgets.NewCollapse(
	widgets.CollapsePanel{Title: "第一节 · 渲染管线", Content: &widgets.Label{Text: "RenderNode → Canvas → DisplayList → Layer → damage"}},
	widgets.CollapsePanel{Title: "第二节 · 交互链路", Content: &widgets.Label{Text: "命中路径 → Dispatcher → GestureArena → 回调"}},
	widgets.CollapsePanel{Title: "第三节 · 输入与 IME", Content: &widgets.Label{Text: "C28：TextKeyHandler 与 IME 结果分离"}},
)

func main() {
	launcher.Main(doc.Meta{
		ID: "collapse", Title: "Collapse 折叠面板", Group: doc.GroupDisplay,
		What: []string{
			"可展开/收起的内容面板：点击表头切换，表头宽度走 CrossStretch",
			"当前为单开手风琴（Open(-1) 全收）；计划（M9）支持多开",
		},
		Demos: []doc.Demo{
			{Title: "三节手风琴", Note: "点击表头展开/收起", Build: func(ctx *ui.BuildContext) ui.Widget {
				return collapse
			}},
		},
		API: []doc.APIRow{
			{Name: "NewCollapse", Type: "func(panels ...CollapsePanel) *Collapse", Desc: "构造；CollapsePanel{Title,Content}"},
			{Name: "Panels", Type: "[]CollapsePanel", Desc: "面板集合"},
			{Name: "Open", Type: "int", Default: "-1（全收）", Desc: "展开的面板下标"},
			{Name: "OnChange", Type: "func(int)", Desc: "展开状态变化回调"},
		},
	})
}
