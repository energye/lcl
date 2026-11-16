package examples

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func box(ctx *ui.BuildContext, label string, color geom.Color, w, h float32) *scene.Node {
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(color)}, (&widgets.Label{Text: label, Color: geom.RGBA(255, 255, 255)}).Build(ctx))
	if w > 0 || h > 0 {
		n.SetSize2(w, h)
	}
	return n
}

func init() {
	Register(Example{ID: "layout-stack", Name: "Stack 层叠", Group: GroupLayout, Build: func(ctx *ui.BuildContext) *scene.Node {
		th := ctx.ThemeOr()
		root := scene.Stack().SetSize2(420, 160)
		root.Add(box(ctx, "底层", th.Color.Primary, 400, 140))
		root.Add(scene.Stack().SetMargin(geom.InLTRB(40, 40, 0, 0)).SetSize2(320, 100))
		inner := scene.Stack(
			box(ctx, "中层", th.Color.Warning, 300, 90),
		).SetMargin(geom.InLTRB(20, 20, 0, 0))
		root.Add(inner)
		return Section(ctx, "Stack 层叠", "后加的子节点盖在前面；无句柄、同一画布内合成", root)
	}})

	Register(Example{ID: "layout-flex-main", Name: "Flex 主轴分布", Group: GroupLayout, Build: func(ctx *ui.BuildContext) *scene.Node {
		th := ctx.ThemeOr()
		col := scene.Flex(scene.Vertical, 10).SetWidth(560)
		for _, align := range []struct {
			name string
			a    scene.MainAlign
		}{
			{"Start", scene.MainStart}, {"Center", scene.MainCenter},
			{"End", scene.MainEnd}, {"SpaceBetween", scene.MainSpaceBetween},
			{"SpaceAround", scene.MainSpaceAround},
		} {
			row := scene.Flex(scene.Horizontal, 4).SetMainAlign(align.a).SetHeight(36)
			for i := 0; i < 3; i++ {
				row.Add(box(ctx, fmt.Sprintf("%d", i+1), th.Color.Primary, 48, 28))
			}
			col.Add((&widgets.Label{Text: align.name}).Build(ctx))
			col.Add(row)
		}
		return Section(ctx, "Flex 主轴", "Start / Center / End / SpaceBetween / SpaceAround", col)
	}})

	Register(Example{ID: "layout-flex-cross", Name: "Flex 交叉轴对齐", Group: GroupLayout, Build: func(ctx *ui.BuildContext) *scene.Node {
		th := ctx.ThemeOr()
		col := scene.Flex(scene.Vertical, 10).SetWidth(560)
		for _, item := range []struct {
			name string
			a    scene.CrossAlign
		}{
			{"Stretch（默认）", scene.CrossStretch},
			{"Center", scene.CrossCenter},
			{"Start", scene.CrossStart},
			{"End", scene.CrossEnd},
		} {
			row := scene.Flex(scene.Horizontal, 6).SetCrossAlign(item.a).SetHeight(52)
			row.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken)})
			for i := 0; i < 3; i++ {
				b := box(ctx, "x", th.Color.Info, 0, 0) // 不固定尺寸：交给拉伸/对齐决定
				if item.a != scene.CrossStretch {
					b.SetHeight(28)
				}
				b.SetFlex(1, scene.FlexTight)
				row.Add(b)
			}
			col.Add((&widgets.Label{Text: item.name}).Build(ctx))
			col.Add(row)
		}
		return Section(ctx, "Flex 交叉轴", "Stretch 让子节点占满交叉轴；固定宽 0 会被当成 0 而非自动", col)
	}})

	Register(Example{ID: "layout-grid", Name: "Grid 网格", Group: GroupLayout, Build: func(ctx *ui.BuildContext) *scene.Node {
		th := ctx.ThemeOr()
		grid := scene.Grid(3, 8).SetWidth(520)
		for i := 0; i < 12; i++ {
			grid.Add(box(ctx, fmt.Sprintf("%d", i+1), th.Color.Primary, 0, 52).SetFlex(1, scene.FlexTight))
		}
		return Section(ctx, "Grid", "固定列数 + 间距；列宽由约束系统分配", grid)
	}})

	Register(Example{ID: "layout-wrap", Name: "Wrap 流式换行", Group: GroupLayout, Build: func(ctx *ui.BuildContext) *scene.Node {
		th := ctx.ThemeOr()
		wrap := scene.Wrap(scene.Horizontal, 6, 6).SetWidth(520)
		for _, t := range []string{"无句柄画布", "DisplayList", "Layer 树", "damage", "GestureArena", "IME", "焦点域", "浮层", "主题"} {
			wrap.Add(box(ctx, t, th.Color.SurfaceActive, 0, 32).SetFlex(0, scene.FlexTight).SetPadding(geom.InLTRB(8, 4, 8, 4)))
		}
		return Section(ctx, "Wrap", "主轴排满自动换行，runGap 控制行间距", wrap)
	}})

	Register(Example{ID: "layout-constraints", Name: "约束与拉伸", Group: GroupLayout, Build: func(ctx *ui.BuildContext) *scene.Node {
		th := ctx.ThemeOr()
		row := scene.Flex(scene.Horizontal, 8).SetWidth(560).SetHeight(80)
		row.Add(box(ctx, "固定 80", th.Color.Primary, 80, 60))
		row.Add(box(ctx, "flex=1", th.Color.Info, 0, 60).SetFlex(1, scene.FlexTight))
		row.Add(box(ctx, "flex=2", th.Color.Success, 0, 60).SetFlex(2, scene.FlexTight))
		return Section(ctx, "约束与拉伸", "固定尺寸 + flex 权重共同分配主轴剩余空间", row)
	}})

	Register(Example{ID: "layout-scroll-offset", Name: "滚动偏移", Group: GroupLayout, Build: func(ctx *ui.BuildContext) *scene.Node {
		th := ctx.ThemeOr()
		inner := scene.Flex(scene.Vertical, 4)
		for i := 0; i < 20; i++ {
			inner.Add(box(ctx, fmt.Sprintf("行 %02d", i+1), th.Color.Surface, 300, 28))
		}
		view := scene.Scroll(scene.Vertical, inner).SetSize2(320, 160)
		view.SetScrollOffset(80)
		return Section(ctx, "滚动偏移", "SetScrollOffset 改变视口位置；滚动位置是布局状态，不属于脏标记", view)
	}})
}
