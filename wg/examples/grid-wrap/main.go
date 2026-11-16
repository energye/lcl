package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "grid-wrap", Title: "Grid 网格 / Wrap 流式", Group: doc.GroupLayout,
		What: []string{
			"Grid：固定列数网格，列宽由约束系统分配",
			"Wrap：主轴排满自动换行，runGap 控制行间距，适合标签流",
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
			{Name: "SetWidth", Type: "func(w float32)", Desc: "Grid/Wrap 都需要显式宽度才有换行边界"},
		},
	})
}
