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

// DescriptionItem 是描述项（antd items 项）。
type DescriptionItem struct {
	Label   string
	Content string
	// Widget 内容 Widget（antd children 任意节点；非空时优先于
	// Content 文本——深化轮）。
	Widget ui.Widget
	// Span 跨列数（antd span：该项占的 column 格数，>=1；默认 1）。
	// 影响该项 content 格在行内的宽度权重（label 占 1）。
	Span int
	// Extra 操作区（antd extra：该项 content 右侧追加 Widget）。
	Extra ui.Widget
}

// DescriptionsSize 是描述列表尺寸（antd size）。
type DescriptionsSize uint8

const (
	// DescSizeDefault 默认行高 34。
	DescSizeDefault DescriptionsSize = iota
	// DescSizeMiddle 中号行高 30。
	DescSizeMiddle
	// DescSizeSmall 小号行高 24。
	DescSizeSmall
)

// Descriptions 是描述列表。
type Descriptions struct {
	controlBase
	Title    string
	Items    []DescriptionItem
	Columns  int // 每行组数（antd column；默认 2）
	Bordered bool
	// Vertical 竖排布局（antd layout="vertical"）：label 上 content
	// 下逐项纵排（每项独占行宽）。
	Vertical bool
	// Size 尺寸（默认/中/小三档行高）。
	Size DescriptionsSize
	// Extra 全局操作区（antd extra：标题行右侧 Widget）。
	Extra ui.Widget
	Width float32
}

// NewDescriptions 创建描述列表。
func NewDescriptions(items []DescriptionItem) *Descriptions {
	return &Descriptions{Items: items, Columns: 2}
}

// rowH 当前尺寸的单元格行高。
func (d *Descriptions) rowH() float32 {
	switch d.Size {
	case DescSizeMiddle:
		return 30
	case DescSizeSmall:
		return 24
	default:
		return 34
	}
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
	h := d.rowH()

	col := scene.Flex(scene.Vertical, th.Space.SM)
	col.SetWidth(w)
	col.SetHit(scene.HitSpec{Label: "descriptions"})

	if d.Title != "" || d.Extra != nil {
		headRow := scene.Flex(scene.Horizontal, th.Space.SM)
		headRow.SetCrossAlign(scene.CrossCenter)
		if d.Title != "" {
			tl := text.LayoutText(d.Title, text.DefaultFont(th.Font.SizeLG), text.Options{})
			title := scene.Text(scene.TextSpec{Src: d.Title, Layout: tl, Color: th.Color.Text})
			title.SetHit(scene.HitSpec{Label: "descriptions-title"})
			headRow.Add(title)
		}
		if d.Extra != nil {
			sp := scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight)
			headRow.Add(sp)
			if en := d.Extra.Build(ctx); en != nil {
				en.SetHit(scene.HitSpec{Label: "desc-extra"})
				headRow.Add(en)
			}
		}
		col.Add(headRow)
	}

	if d.Vertical {
		// 竖排：label 上（灰字）+ content 下，逐项纵排独占行宽。
		for _, it := range d.Items {
			itemCol := scene.Flex(scene.Vertical, 2)
			itemCol.SetWidth(w)
			if d.Bordered {
				itemCol.SetBox(scene.BoxSpec{Border: scene.Stroke(th.Color.Border, 1)})
			}
			l := text.LayoutText(it.Label, text.DefaultFont(th.Font.SizeBase), text.Options{})
			lt := scene.Text(scene.TextSpec{Src: it.Label, Layout: l, Color: th.Color.TextMuted})
			lt.SetMargin(geom.InLTRB(10, 2, 0, 0))
			lt.SetHit(scene.HitSpec{Label: "desc-label"})
			itemCol.Add(lt)
			if d.Bordered {
				lt.SetMargin(geom.InLTRB(10, 4, 0, 0))
			}
			vc := d.valueCell(ctx, th, it, w-20, h)
			vc.SetWidth(w)
			vc.SetHeight(h + 4)
			itemCol.Add(vc)
			col.Add(itemCol)
		}
		return col
	}

	// 横排：逐行 Flex（支持 span 加权——Grid 等宽格不支持跨列，改为
	// 每行一个 Flex Horizontal，label 固定宽、content 按 span 权重
	// 分摊剩余）。
	cellW := (w - th.Space.XS*float32(nCols-1)) / float32(nCols)
	labelW := cellW * 0.42
	rowItems := 0
	var row *scene.Node
	flushRow := func() {
		if row != nil && rowItems > 0 {
			col.Add(row)
		}
		row = nil
		rowItems = 0
	}
	for i := 0; i < len(d.Items); i++ {
		it := d.Items[i]
		if rowItems == 0 {
			row = scene.Flex(scene.Horizontal, 0)
			if d.Bordered {
				row.SetBox(scene.BoxSpec{Border: scene.Stroke(th.Color.Border, 1)})
			}
		}
		span := it.Span
		if span < 1 {
			span = 1
		}
		if rowItems+span > nCols {
			span = nCols - rowItems // 行尾截断（antd 同语义）
		}
		lc := d.labelCell(ctx, th, it.Label, labelW, h)
		vc := d.valueCell(ctx, th, it, 0, h)
		contentW := cellW*float32(span) - labelW
		if d.Bordered {
			contentW += th.Space.XS
		}
		vc.SetWidth(contentW)
		// extra 追加在 content 格右侧。
		if it.Extra != nil {
			pair := scene.Flex(scene.Horizontal, th.Space.SM)
			pair.SetWidth(contentW)
			vc.SetWidth(contentW - 60)
			if en := it.Extra.Build(ctx); en != nil {
				en.SetHit(scene.HitSpec{Label: "desc-extra"})
				pair.Add(vcc(vc), pairExtra(en))
			} else {
				pair.Add(vcc(vc))
			}
			row.Add(pair)
		} else {
			row.Add(lc)
			row.Add(vc)
		}
		rowItems += span
		if rowItems >= nCols {
			flushRow()
		}
	}
	flushRow()
	return col
}

// vcc/vccPair 透传节点（内部小工具，避免闭包捕获歧义）。
func vcc(n *scene.Node) *scene.Node { return n }

func pairExtra(en *scene.Node) *scene.Node { return en }

// labelCell label 格：bordered 灰底描边格；默认灰字+冒号。
func (d *Descriptions) labelCell(ctx *ui.BuildContext, th *theme.Data, label string, w, h float32) *scene.Node {
	box := scene.Box(scene.BoxSpec{})
	box.SetWidth(w).SetHeight(h)
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
	l := text.LayoutText(txt, text.DefaultFont(th.Font.SizeBase), text.Options{})
	t := scene.Text(scene.TextSpec{Src: txt, Layout: l, Color: th.Color.TextMuted})
	t.SetMargin(geom.InLTRB(10, (h-l.Size.H)/2, 0, 0))
	t.SetHit(scene.HitSpec{Label: "desc-label"})
	box.Add(t)
	return box
}

// valueCell content 格：bordered 描边白底格；默认正文。
// Widget 非空时渲染 Widget（antd children 任意节点——深化轮）。
func (d *Descriptions) valueCell(ctx *ui.BuildContext, th *theme.Data, it DescriptionItem, w, h float32) *scene.Node {
	box := scene.Box(scene.BoxSpec{})
	box.SetWidth(w).SetHeight(h)
	if d.Bordered {
		box.SetBox(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.Surface),
			Border: scene.Stroke(th.Color.Border, 1),
		})
	}
	if it.Widget != nil {
		if wn := it.Widget.Build(ctx); wn != nil {
			wn.SetMargin(geom.InLTRB(10, (h-wn.Size().H)/2, 0, 0))
			wn.SetHit(scene.HitSpec{Label: "desc-value"})
			box.Add(wn)
		}
		return box
	}
	l := text.LayoutText(it.Content, text.DefaultFont(th.Font.SizeBase), text.Options{})
	t := scene.Text(scene.TextSpec{Src: it.Content, Layout: l, Color: th.Color.Text})
	t.SetMargin(geom.InLTRB(10, (h-l.Size.H)/2, 0, 0))
	t.SetHit(scene.HitSpec{Label: "desc-value"})
	box.Add(t)
	return box
}
