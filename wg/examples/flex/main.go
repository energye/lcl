package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "flex", Title: "Flex 弹性布局", Group: doc.GroupLayout,
		What: []string{
			"布局引擎的主轴系统：主轴分布（5 种对齐）+ 交叉轴对齐（4 种）+ flex 权重分配",
			"组件层有 Row/Column 两个值类型封装；节点层直接用 scene.Flex",
			"约束系统：固定尺寸 + flex 权重共同分配主轴剩余空间；Tight/Loose/Deflate 协商",
		},
		Demos: []doc.Demo{
			{Title: "Row / Column 组件", Note: "widgets 值类型封装，声明式组装", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := widgets.Row{Gap: 12, Children: []ui.Widget{
					&widgets.Label{Text: "A"}, &widgets.Label{Text: "B"}, &widgets.Label{Text: "C"},
				}}
				spread := widgets.Row{Gap: 8, MainAlign: scene.MainSpaceBetween, Children: []ui.Widget{
					&widgets.Label{Text: "左对齐"}, &widgets.Label{Text: "居中"}, &widgets.Label{Text: "右对齐"},
				}}
				col := widgets.Column{Gap: 8, Children: []ui.Widget{
					&widgets.Label{Text: "第一行"}, &widgets.Label{Text: "第二行"}, &widgets.Label{Text: "第三行"},
				}}
				return widgets.Column{Gap: 12, Children: []ui.Widget{&row, &spread, &col}}
			}},
			{Title: "主轴分布", Note: "Start / Center / End / SpaceBetween / SpaceAround", Build: func(ctx *ui.BuildContext) ui.Widget {
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
						row.Add(doc.Box(ctx, fmt.Sprintf("%d", i+1), th.Color.Primary, 48, 28))
					}
					col.Add((&widgets.Label{Text: align.name}).Build(ctx))
					col.Add(row)
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "交叉轴对齐", Note: "Stretch 让子节点占满交叉轴；固定宽 0 会被当成 0 而非自动", Build: func(ctx *ui.BuildContext) ui.Widget {
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
						b := doc.Box(ctx, "x", th.Color.Info, 0, 0)
						if item.a != scene.CrossStretch {
							b.SetHeight(28)
						}
						b.SetFlex(1, scene.FlexTight)
						row.Add(b)
					}
					col.Add((&widgets.Label{Text: item.name}).Build(ctx))
					col.Add(row)
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "约束与拉伸", Note: "固定尺寸 + flex 权重共同分配主轴剩余空间", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				row := scene.Flex(scene.Horizontal, 8).SetWidth(560).SetHeight(80)
				row.Add(doc.Box(ctx, "固定 80", th.Color.Primary, 80, 60))
				row.Add(doc.Box(ctx, "flex=1", th.Color.Info, 0, 60).SetFlex(1, scene.FlexTight))
				row.Add(doc.Box(ctx, "flex=2", th.Color.Success, 0, 60).SetFlex(2, scene.FlexTight))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "约束协商 API", Note: "Tight/Loose/Deflate 的数值演示（CustomPaint 内直接调用约束函数）", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				tight := scene.TightConstraints(geom.Sz(120, 60))
				loose := scene.NewConstraints(40, 200, 30, 120)
				def := loose.Deflate(geom.InLTRB(10, 5, 10, 5))
				t := text.DefaultFont(th.Font.SizeSM)
				n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
					lay := text.LayoutText("Tight 120x60", t, text.Options{})
					cv.DrawText(lay, geom.Pt(10, 6), th.Color.Text)
					cv.StrokePath(render.NewPath().AddRect(geom.RectXYWH(10, 28, tight.Constrain(geom.Sz(260, 200)).W, tight.Constrain(geom.Sz(260, 200)).H)),
						render.Stroke(th.Color.Primary, 2))
					cv.DrawText(text.LayoutText("Loose/Deflate", t, text.Options{}), geom.Pt(10, 110), th.Color.Text)
					cv.StrokePath(render.NewPath().AddRect(geom.RectXYWH(10, 132, def.Constrain(geom.Sz(260, 200)).W, def.Constrain(geom.Sz(260, 200)).H)),
						render.Stroke(th.Color.Success, 2))
				}).SetSize2(300, 220)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("Tight=%v Loose=%v Deflate=%v",
					tight.MaxSize(), loose.MaxSize(), def.MaxSize())))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "scene.Flex", Type: "func(axis Axis, gap float32) *Node", Desc: "弹性容器；Add 追加子节点"},
			{Name: "SetMainAlign", Type: "func(MainAlign)", Desc: "Start/Center/End/SpaceBetween/SpaceAround"},
			{Name: "SetCrossAlign", Type: "func(CrossAlign)", Desc: "Stretch/Center/Start/End"},
			{Name: "SetFlex", Type: "func(weight float32, tightness)", Desc: "权重分配剩余空间；FlexTight 表示尽量占满"},
			{Name: "widgets.Row/Column", Type: "值类型", Desc: "组件层封装：Gap/Children/MainAlign/CrossAlign"},
			{Name: "TightConstraints", Type: "func(Size) Constraints", Desc: "紧约束（强制占满）"},
			{Name: "NewConstraints", Type: "func(minW,maxW,minH,maxH) Constraints", Desc: "松约束"},
			{Name: "Constraints.Deflate", Type: "func(Insets) Constraints", Desc: "内缩约束"},
		},
	})
}
