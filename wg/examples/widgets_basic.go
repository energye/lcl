package examples

import (
	"fmt"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func init() {
	Register(Example{ID: "label", Name: "Label 文本", Group: GroupBasic, Build: func(ctx *ui.BuildContext) *scene.Node {
		th := ctx.ThemeOr()
		col := scene.Flex(scene.Vertical, 8).SetWidth(560)
		for _, s := range []struct {
			text string
			size float32
		}{
			{"标题 22px", th.Font.SizeTitle},
			{"大号 18px", th.Font.SizeXL},
			{"正文 13px · 中文排版走 rune 偏移", th.Font.SizeBase},
			{"小号 12px", th.Font.SizeSM},
			{"极小 11px", th.Font.SizeXS},
		} {
			col.Add((&widgets.Label{Text: s.text, Size: s.size}).Build(ctx))
		}
		return Section(ctx, "Label", "文本排版：rune 偏移，UTF-16 只在 host 层换算", col)
	}})

	Register(Example{ID: "divider", Name: "Divider 分隔线", Group: GroupBasic, Build: func(ctx *ui.BuildContext) *scene.Node {
		col := scene.Flex(scene.Vertical, 10).SetWidth(420)
		col.Add((&widgets.Label{Text: "上方内容"}).Build(ctx))
		col.Add((&widgets.Divider{Length: 400}).Build(ctx))
		col.Add((&widgets.Label{Text: "下方内容"}).Build(ctx))
		col.Add((&widgets.Divider{Length: 400, Thickness: 3}).Build(ctx))
		col.Add((&widgets.Label{Text: "加粗分隔线"}).Build(ctx))
		return Section(ctx, "Divider", "水平/竖直分隔线，厚度与长度可配", col)
	}})

	Register(Example{ID: "spacer", Name: "Spacer / SizedBox", Group: GroupBasic, Build: func(ctx *ui.BuildContext) *scene.Node {
		row := scene.Flex(scene.Horizontal, 0)
		row.Add((&widgets.Label{Text: "左"}).Build(ctx))
		row.Add((&widgets.Spacer{Width: 120}).Build(ctx))
		row.Add((&widgets.Label{Text: "右 · 中间 120px 空隙"}).Build(ctx))
		box := (&widgets.SizedBox{Width: 260, Height: 60, Child: &widgets.Label{Text: "固定 260×60"}}).Build(ctx)
		box.SetBox(scene.BoxSpec{Fill: scene.Solid(ctx.ThemeOr().Color.Surface)})
		col := scene.Flex(scene.Vertical, 12)
		col.Add(row)
		col.Add(box)
		return Section(ctx, "Spacer / SizedBox", "弹性留白与固定尺寸盒子", col)
	}})

	Register(Example{ID: "padding", Name: "Padding 内边距", Group: GroupBasic, Build: func(ctx *ui.BuildContext) *scene.Node {
		inner := scene.Box(scene.BoxSpec{Fill: scene.Solid(ctx.ThemeOr().Color.Surface)},
			(&widgets.Label{Text: "内边距 24 / 12 / 24 / 12"}).Build(ctx))
		inner.SetSize2(420, 56)
		n := (&widgets.Padding{Insets: geom.InLTRB(24, 12, 24, 12), Child: &widgets.Label{Text: "被 padding 包裹"}}).Build(ctx)
		col := scene.Flex(scene.Vertical, 12)
		col.Add(inner)
		col.Add(n)
		return Section(ctx, "Padding", "Insets 顺序为 LTRB；节点自身尺寸随子内容增长", col)
	}})

	Register(Example{ID: "row", Name: "Row 横向布局", Group: GroupBasic, Build: func(ctx *ui.BuildContext) *scene.Node {
		row := widgets.Row{Gap: 12, Children: []ui.Widget{
			&widgets.Label{Text: "A"},
			&widgets.Label{Text: "B"},
			&widgets.Label{Text: "C"},
		}}
		spread := widgets.Row{Gap: 8, MainAlign: scene.MainSpaceBetween, Children: []ui.Widget{
			&widgets.Label{Text: "左对齐"},
			&widgets.Label{Text: "居中"},
			&widgets.Label{Text: "右对齐"},
		}}
		col := scene.Flex(scene.Vertical, 12).SetWidth(520)
		col.Add(row.Build(ctx))
		col.Add(spread.Build(ctx))
		return Section(ctx, "Row", "主轴横向，支持 space-between / space-around 等对齐", col)
	}})

	Register(Example{ID: "column", Name: "Column 纵向布局", Group: GroupBasic, Build: func(ctx *ui.BuildContext) *scene.Node {
		col := widgets.Column{Gap: 8, Children: []ui.Widget{
			&widgets.Label{Text: "第一行"},
			&widgets.Label{Text: "第二行"},
			&widgets.Label{Text: "第三行"},
		}}
		return Section(ctx, "Column", "主轴纵向，交叉轴拉伸由 scene 约束系统决定", col.Build(ctx))
	}})

	{
		list := widgets.NewList(
			widgets.ListItem{Key: "a", Widget: &widgets.Label{Text: "键控项 A"}},
			widgets.ListItem{Key: "b", Widget: &widgets.Label{Text: "键控项 B"}},
			widgets.ListItem{Key: "c", Widget: &widgets.Label{Text: "键控项 C（键控复用，不做虚拟滚动）"}},
		)
		Register(Example{ID: "list", Name: "List 键控列表", Group: GroupBasic,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				return Section(ctx, "List", "同层同 Key 复用节点，保留布局与缓存", list.Build(ctx))
			}})
	}

	Register(Example{ID: "scroll", Name: "Scroll 滚动", Group: GroupBasic, Build: func(ctx *ui.BuildContext) *scene.Node {
		child := From(func(ctx *ui.BuildContext) *scene.Node {
			inner := scene.Flex(scene.Vertical, 6)
			for i := 0; i < 24; i++ {
				inner.Add((&widgets.Label{Text: fmt.Sprintf("滚动行 %c · 第 %02d 行", 'A'+i%26, i+1)}).Build(ctx))
			}
			return inner
		})
		n := (&widgets.Scroll{Axis: scene.Vertical, Child: child, Width: 360, Height: 200}).Build(ctx)
		return Section(ctx, "Scroll", "滚动视口与裁剪；滚轮/拖动由 input 分发", n)
	}})
}
