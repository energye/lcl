package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "padding-spacer", Title: "Padding / Spacer / SizedBox", Group: doc.GroupBasic,
		What: []string{
			"三个无绘制的布局原语：Padding 包内边距、Spacer 弹性留白、SizedBox 固定尺寸",
			"全部是值类型字面量，Build 时折成 scene 约束节点",
		},
		Demos: []doc.Demo{
			{Title: "Spacer 弹性留白", Note: "在 Row 里撑出固定宽度的空隙", Build: func(ctx *ui.BuildContext) ui.Widget {
				return &widgets.Row{Gap: 0, Children: []ui.Widget{
					&widgets.Label{Text: "左"},
					&widgets.Spacer{Width: 120},
					&widgets.Label{Text: "右 · 中间 120px 空隙"},
				}}
			}},
			{Title: "SizedBox 固定尺寸", Note: "把子内容约束进 260×60", Build: func(ctx *ui.BuildContext) ui.Widget {
				box := (&widgets.SizedBox{Width: 260, Height: 60, Child: &widgets.Label{Text: "固定 260×60"}}).Build(ctx)
				box.SetBox(scene.BoxSpec{Fill: scene.Solid(ctx.ThemeOr().Color.Surface)})
				return doc.Node(func(*ui.BuildContext) *scene.Node { return box })
			}},
			{Title: "Padding 内边距", Note: "Insets 顺序为 LTRB；节点自身尺寸随子内容增长", Build: func(ctx *ui.BuildContext) ui.Widget {
				return &widgets.Padding{Insets: geom.InLTRB(24, 12, 24, 12),
					Child: &widgets.Label{Text: "被 padding 包裹（24/12/24/12）"}}
			}},
		},
		API: []doc.APIRow{
			{Name: "Padding.Insets", Type: "geom.Insets", Desc: "左/上/右/下内边距（InLTRB）"},
			{Name: "Padding.Child", Type: "ui.Widget", Desc: "被包裹的子组件"},
			{Name: "Spacer.Width/Height", Type: "float32", Default: "0", Desc: "弹性留白尺寸"},
			{Name: "SizedBox.Width/Height", Type: "float32", Desc: "固定尺寸；Child 可省"},
		},
	})
}
