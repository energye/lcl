package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var respGrid = widgets.NewGrid(
	widgets.GridCol{Child: widgets.Label{Text: "卡片 1"}, Span: 6, Md: intp(12), Xs: intp(24)},
	widgets.GridCol{Child: widgets.Label{Text: "卡片 2"}, Span: 6, Md: intp(12), Xs: intp(24)},
	widgets.GridCol{Child: widgets.Label{Text: "卡片 3"}, Span: 6, Md: intp(12), Xs: intp(24)},
	widgets.GridCol{Child: widgets.Label{Text: "卡片 4"}, Span: 6, Md: intp(12), Xs: intp(24)},
)

func intp(v int) *int { return &v }

func main() {
	launcher.Main(doc.Meta{
		ID: "grid-wrap", Title: "Grid 网格 / Wrap 流式", Group: doc.GroupLayout,
		What: []string{
			"Grid：固定列数网格，列宽由约束系统分配",
			"Wrap：主轴排满自动换行，runGap 控制行间距，适合标签流",
			"widgets.Grid：24 栅格 + 响应式断点（xs/sm/md/lg/xl），按视口宽自适应列数",
		},
		Demos: []doc.Demo{
			{Title: "Grid 3 列", Note: "12 个单元格，列宽 flex 均分", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				grid := scene.Grid(3, 8).SetWidth(520)
				for i := 0; i < 12; i++ {
					grid.Add(doc.Box(ctx, fmt.Sprintf("%d", i+1), th.Color.Primary, 0, 52).SetFlex(1, scene.FlexTight))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return grid })
			}},
			{Title: "响应式栅格", Note: "widgets.Grid：24 栅格 + 断点（lg 4 列/md 2 列/xs 1 列，按视口宽自适应）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return respGrid
			}},
			{Title: "Wrap 流式换行", Note: "标签排满自动换行，行距 6", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				wrap := scene.Wrap(scene.Horizontal, 6, 6).SetWidth(520)
				for _, t := range []string{"无句柄画布", "DisplayList", "Layer 树", "damage", "GestureArena", "IME", "焦点域", "浮层", "主题"} {
					wrap.Add(doc.Box(ctx, t, th.Color.SurfaceActive, 0, 32).SetFlex(0, scene.FlexTight).SetPadding(geom.InLTRB(8, 4, 8, 4)))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return wrap })
			}},
		},
		API: []doc.APIRow{
			{Name: "scene.Grid", Type: "func(columns, gap int) *Node", Desc: "固定列数网格"},
			{Name: "scene.Wrap", Type: "func(axis Axis, gap, runGap float32) *Node", Desc: "流式换行容器；runGap 是行间距"},
			{Name: "widgets.NewGrid/GridCol", Type: "func(cols ...GridCol) *Grid", Desc: "24 栅格响应式布局；Span/Offset + Xs/Sm/Md/Lg/Xl 断点"},
			{Name: "SetWidth", Type: "func(w float32)", Desc: "Grid/Wrap 都需要显式宽度才有换行边界"},
		},
	})
}
