package widgets

// ── Descriptions 描述列表（M12 长尾）─────────────────────────────
//
// antd：<Descriptions title bordered column={n} items>
// 第一轮：标题行 + Grid(columns×2) 网格（**组合复用 scene.Grid**）+
// 双形态（Bordered：label 灰底格/全网格描边；默认：label 灰字+
// 冒号 + 无边框疏排）。Items 首轮为文本值（antd children 的
// Widget 化后续轮）。
// 未实现（如实标注）：content 为 Widget、layout 竖排（vertical）、
// size 尺寸、extra 操作区、colon 开关、多值 spans 合并。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// DescriptionItem 是描述项（antd items 项的文本子集）。
type DescriptionItem struct {
	Label   string
	Content string
}

// Descriptions 是描述列表。
type Descriptions struct {
	controlBase
	Title    string
	Items    []DescriptionItem
	Columns  int // 每行组数（antd column；默认 2）
	Bordered bool
	Width    float32
}

// NewDescriptions 创建描述列表。
func NewDescriptions(items []DescriptionItem) *Descriptions {
	return &Descriptions{Items: items, Columns: 2}
}

// Build 实现 ui.Widget。
func (d *Descriptions) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(d)
	th := ctx.ThemeOr()
	nCols := d.Columns
	if nCols <= 0 {
		nCols = 2
	}
	w := d.Width
	if w <= 0 {
		w = 560
	}

	col := scene.Flex(scene.Vertical, th.Space.SM)
	col.SetWidth(w)
	col.SetHit(scene.HitSpec{Label: "descriptions"})

	if d.Title != "" {
		tl := text.LayoutText(d.Title, text.DefaultFont(th.Font.SizeLG), text.Options{})
		title := scene.Text(scene.TextSpec{Src: d.Title, Layout: tl, Color: th.Color.Text})
		title.SetHit(scene.HitSpec{Label: "descriptions-title"})
		col.Add(title)
	}

	// 网格：每行 nCols 组（label 格+content 格交替）。
	gridCols := nCols * 2
	cellW := (w - th.Space.XS*float32(gridCols-1)) / float32(gridCols)
	var cells []*scene.Node
	for _, it := range d.Items {
		cells = append(cells, d.labelCell(ctx, th, it.Label, cellW))
		cells = append(cells, d.valueCell(ctx, th, it.Content, cellW))
	}
	gap := th.Space.XS
	if d.Bordered {
		gap = 0 // bordered 网格密排（共享边框线）
	}
	grid := scene.Grid(gridCols, gap, cells...)
	grid.SetHit(scene.HitSpec{Label: "descriptions-grid"})
	col.Add(grid)
	return col
}

// labelCell label 格：bordered 灰底描边格；默认灰字+冒号。
func (d *Descriptions) labelCell(ctx *ui.BuildContext, th *theme.Data, label string, w float32) *scene.Node {
	box := scene.Box(scene.BoxSpec{})
	box.SetWidth(w).SetHeight(34)
	if d.Bordered {
		box.SetBox(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.FillHover),
			Border: scene.Stroke(th.Color.Border, 1),
		})
	}
	txt := label
	if !d.Bordered {
		txt += ":"
	}
	l := text.LayoutText(txt, text.DefaultFont(th.Font.SizeSM), text.Options{})
	t := scene.Text(scene.TextSpec{Src: txt, Layout: l, Color: th.Color.TextMuted})
	t.SetMargin(geom.InLTRB(10, (34-l.Size.H)/2, 0, 0))
	t.SetHit(scene.HitSpec{Label: "desc-label"})
	box.Add(t)
	return box
}

// valueCell content 格：bordered 描边白底格；默认正文。
func (d *Descriptions) valueCell(ctx *ui.BuildContext, th *theme.Data, content string, w float32) *scene.Node {
	box := scene.Box(scene.BoxSpec{})
	box.SetWidth(w).SetHeight(34)
	if d.Bordered {
		box.SetBox(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.Surface),
			Border: scene.Stroke(th.Color.Border, 1),
		})
	}
	l := text.LayoutText(content, text.DefaultFont(th.Font.SizeSM), text.Options{})
	t := scene.Text(scene.TextSpec{Src: content, Layout: l, Color: th.Color.Text})
	t.SetMargin(geom.InLTRB(10, (34-l.Size.H)/2, 0, 0))
	t.SetHit(scene.HitSpec{Label: "desc-value"})
	box.Add(t)
	return box
}

