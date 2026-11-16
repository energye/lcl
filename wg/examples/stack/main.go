package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "stack", Title: "Stack 层叠", Group: doc.GroupLayout,
		What: []string{
			"层叠布局：后加的子节点盖在前面，无句柄、同一画布内合成",
			"SetStackPosition 精确落位；SetMargin 相对父级偏移",
		},
		Demos: []doc.Demo{
			{Title: "三层叠放", Note: "底层 400×140 → 中层偏移 → 顶层；Z 序=添加顺序", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				root := scene.Stack().SetSize2(420, 160)
				root.Add(doc.Box(ctx, "底层", th.Color.Primary, 400, 140))
				root.Add(scene.Stack().SetMargin(geom.InLTRB(40, 40, 0, 0)).SetSize2(320, 100))
				inner := scene.Stack(
					doc.Box(ctx, "中层", th.Color.Warning, 300, 90),
				).SetMargin(geom.InLTRB(20, 20, 0, 0))
				root.Add(inner)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return root })
			}},
		},
		API: []doc.APIRow{
			{Name: "scene.Stack", Type: "func(children ...*Node) *Node", Desc: "层叠容器"},
			{Name: "SetSize2", Type: "func(w, h float32)", Desc: "Stack 需要显式尺寸（不按子内容收缩）"},
			{Name: "SetStackPosition", Type: "func(geom.Point)", Desc: "子节点精确落位"},
			{Name: "SetMargin", Type: "func(geom.Insets)", Desc: "相对父级的边距偏移"},
		},
	})
}
