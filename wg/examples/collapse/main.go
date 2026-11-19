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

// 多开实例包级持有：demo 闭包每次重建都会执行，闭包内新建会让展开状态
// 随重建丢失（--audit 实测"点了没反应"）。
var collapseMulti = widgets.NewCollapse(
	widgets.CollapsePanel{Title: "多开 · 第一节", Content: &widgets.Label{Text: "多开 · 第一节内容"}},
	widgets.CollapsePanel{Title: "多开 · 第二节", Content: &widgets.Label{Text: "多开 · 第二节内容"}},
	widgets.CollapsePanel{Title: "多开 · 第三节", Content: &widgets.Label{Text: "多开 · 第三节内容"}},
)

func init() {
	collapseMulti.Multiple = true
	collapseMulti.OpenSet = []bool{true, false, true}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "collapse", Title: "Collapse 折叠面板", Group: doc.GroupDisplay,
		What: []string{
			"可展开/收起的内容面板：点击表头切换，表头宽度走 CrossStretch",
			"单开手风琴（Open(-1) 全收）与多开模式（Multiple+OpenSet，各面板独立开关）",
		},
		Demos: []doc.Demo{
			{Title: "三节手风琴", Note: "点击表头展开/收起（单开，同时只开一节）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return collapse
			}},
			{Title: "多开模式", Note: "Multiple=true：各面板独立开关，可同时展开多节", Build: func(ctx *ui.BuildContext) ui.Widget {
				return collapseMulti
			}},
		},
		API: []doc.APIRow{
			{Name: "NewCollapse", Type: "func(panels ...CollapsePanel) *Collapse", Desc: "构造；CollapsePanel{Title,Content}"},
			{Name: "Panels", Type: "[]CollapsePanel", Desc: "面板集合"},
			{Name: "Open", Type: "int", Default: "-1（全收）", Desc: "单开模式的展开下标"},
			{Name: "Multiple", Type: "bool", Default: "false", Desc: "多开模式（antd 默认）：各面板独立开关"},
			{Name: "OpenSet", Type: "[]bool", Desc: "多开模式下各面板的展开状态（与 Panels 对齐）"},
			{Name: "OnChange", Type: "func(int)", Desc: "展开状态变化回调"},
		},
	})
}
