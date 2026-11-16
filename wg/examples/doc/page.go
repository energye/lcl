package doc

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// contentWidth 按视口算版心宽：随窗口自适应，钳制在 [640, 1200]。
// 窗口 resize 时容器会重建根 widget，Build 拿到新的 Viewport 自动重排。
func contentWidth(ctx *ui.BuildContext) float32 {
	w := ctx.Viewport.W - 80 // 版心边距 40×2
	if w < 640 {
		w = 640
	}
	if w > 1200 {
		w = 1200
	}
	return w
}

// page 是文档页渲染器：把 Meta 渲染成"说明 → 演示块 → API 表"三段式页面。
type page struct {
	ui.Base
	meta Meta
}

// Page 把 Meta 包装成可挂载的文档页 Widget。
func Page(m Meta) ui.Widget { return &page{meta: m} }

// Build 实现 ui.Widget。
//
// 页面高度超过视口时整页滚动；版心宽随视口自适应（resize 时容器重建根
// widget，本方法拿到新的 Viewport 自动重排）。
func (p *page) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	w := contentWidth(ctx)
	col := scene.Flex(scene.Vertical, 16)
	col.SetWidth(w)

	// ① 标题区：组件名 + 分组脚注。
	col.Add(widgets.Label{Text: p.meta.Title, Size: th.Font.SizeTitle}.Build(ctx))
	col.Add(widgets.Label{
		Text:  "分组 " + p.meta.Group + " · 目录 examples/" + p.meta.ID,
		Size:  th.Font.SizeSM,
		Color: th.Color.TextFaint,
	}.Build(ctx))

	// ② 是什么/何时使用。
	if len(p.meta.What) > 0 {
		col.Add(sectionTitle(ctx, "何时使用"))
		for _, line := range p.meta.What {
			col.Add(widgets.Label{
				Text:     "· " + line,
				MaxWidth: w - 32,
				Color:    th.Color.Text,
			}.Build(ctx))
		}
	}

	// ③ 演示块×N：每块 = 标题 + 效果说明 + 可交互 demo 舞台。
	if len(p.meta.Demos) > 0 {
		col.Add(sectionTitle(ctx, "使用演示"))
		for _, d := range p.meta.Demos {
			col.Add(p.demoCard(ctx, d))
		}
	}

	// ④ API 属性表。
	if len(p.meta.API) > 0 {
		col.Add(sectionTitle(ctx, "API"))
		col.Add(apiTable(ctx, p.meta.API))
	}

	// 版心边距由外层 Box 的 padding 给（Scroll 子节点不吃 margin，
	// 交叉轴被视口约束直接拉回 0）。
	page := scene.Scroll(scene.Vertical, col)
	root := scene.Box(scene.BoxSpec{}, page)
	root.SetPadding(geom.InLTRB(40, 24, 40, 32))
	return root
}

// sectionTitle 是段落标题。
func sectionTitle(ctx *ui.BuildContext, text string) *scene.Node {
	th := ctx.ThemeOr()
	return widgets.Label{Text: text, Size: th.Font.SizeXL}.Build(ctx)
}

// demoCard 渲染一个演示块。舞台底色用 SurfaceSunken，与页面背景区分开。
func (p *page) demoCard(ctx *ui.BuildContext, d Demo) *scene.Node {
	w := contentWidth(ctx)
	th := ctx.ThemeOr()
	head := scene.Flex(scene.Vertical, 4)
	head.Add(widgets.Label{Text: d.Title, Size: th.Font.SizeLG}.Build(ctx))
	if d.Note != "" {
		head.Add(widgets.Label{
			Text:     d.Note,
			MaxWidth: w - 48,
			Size:     th.Font.SizeSM,
			Color:    th.Color.TextMuted,
		}.Build(ctx))
	}

	// 舞台：内嵌 demo 的浅底容器。demo 为 nil 时只出说明（纯文字演示位）。
	stage := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.SurfaceSunken),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	stage.SetPadding(geom.InLTRB(16, 16, 16, 16))
	if d.Build != nil {
		row := scene.Flex(scene.Horizontal, 12)
		row.SetCrossAlign(scene.CrossCenter)
		row.Add(d.Build(ctx).Build(ctx))
		stage.Add(row)
	}

	card := scene.Flex(scene.Vertical, 12)
	card.SetWidth(w)
	card.Add(head)
	card.Add(stage)
	box := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.LG,
	}, card)
	box.SetPadding(geom.InLTRB(16, 16, 16, 16))
	return box
}

// apiTable 把 APIRow 渲染成 widgets.Table（参数/类型/默认值/说明）。
func apiTable(ctx *ui.BuildContext, rows []APIRow) *scene.Node {
	w := contentWidth(ctx)
	cols := []widgets.TableColumn{
		{Title: "参数", Mode: widgets.ColumnFixed, Width: 150},
		{Title: "类型 / 签名", Mode: widgets.ColumnFixed, Width: 220},
		{Title: "默认值", Mode: widgets.ColumnFixed, Width: 110},
		{Title: "说明", Mode: widgets.ColumnFlex, Flex: 1, Min: 260, Max: 420},
	}
	data := make([][]any, 0, len(rows))
	for _, r := range rows {
		def := r.Default
		if def == "" {
			def = "—" // 空单元格也保持可读（"无默认值"）
		}
		data = append(data, []any{r.Name, r.Type, def, r.Desc})
	}
	t := widgets.NewTable(cols, data)
	t.Width = w
	// 视口高按行数算：表头 32 + 行高 30×N，上下限兜底。
	t.Height = 32 + 30*float32(len(rows)) + 8
	if t.Height < 96 {
		t.Height = 96
	}
	if t.Height > 420 {
		t.Height = 420 // 更长的表在表格内部滚动
	}
	return t.Build(ctx)
}
