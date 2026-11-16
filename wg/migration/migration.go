// Package migration 用新引擎组件重建三个旧示例的核心场景。
//
// 旧示例（examples/wg/{showcase,canvas_ime,dbmodeler}）在 M4 前保持可运行，
// 是行为基线。这里不是逐行翻译那 12k 行代码，而是把每个示例**展示的能力**
// 用新引擎的组件与节点重新表达，让同一场景在新管线上出帧、可交互：
//
//   - showcase：9 卡片 + 5 类图表 → Card + widgets/chart
//   - canvas_ime：输入框 + IME 状态 → P1 Input + Button + 状态面板
//   - dbmodeler：表结构画布 → Table + CustomPaintNode 关系连线
//
// 场景都是零 lcl 的 ui.Widget，可在窗口、headless 与诊断里复用。
// 逐像素对比沿用 --render-matrix（图元级）；这里做结构与交互链路断言。
package migration

import (
	"fmt"
	"math"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
	"github.com/energye/lcl/wg/widgets/chart"
)

// ── Showcase：玻璃卡片 + 波形/仪表/折线/柱状/饼图 ────────────────

// Showcase 是迁移后的展示台页面。
type Showcase struct {
	ui.Base
	Wave  chart.Waveform
	Gauge chart.Gauge
	Line  chart.LineChart
	Bar   chart.BarChart
	Pie   chart.PieChart
	Tip   *widgets.Tooltip
}

// NewShowcase 构造展示台（控件只建一次，重建时复用）。
func NewShowcase() *Showcase {
	s := &Showcase{
		Wave: chart.Waveform{
			Values: waveSamples(180), Width: 430, Height: 170,
		},
		Gauge: chart.Gauge{Value: 0.72, Label: "72%", Width: 430, Height: 170},
		Line: chart.LineChart{
			Series: [][]float64{
				{3, 5, 4, 8, 6, 9, 7}, {2, 3, 6, 4, 7, 5, 8}, {1, 2, 3, 5, 4, 6, 7},
			},
			Labels: []string{"一", "二", "三", "四", "五", "六", "七"},
			Width: 430, Height: 180, FillArea: true, ShowDots: true,
		},
		Bar: chart.BarChart{
			Groups: [][]float64{{3, 5}, {6, 2}, {4, 7}, {8, 3}},
			Labels: []string{"Q1", "Q2", "Q3", "Q4"}, Width: 430, Height: 180,
		},
		Pie: chart.PieChart{
			Values: []float64{4, 3, 2, 1}, Labels: []string{"A", "B", "C", "D"},
			Width: 430, Height: 170, Donut: true,
		},
	}
	s.Tip = &widgets.Tooltip{Text: "新引擎录制的图表", Child: widgets.NewButton("性能说明")}
	return s
}

// Build 实现 ui.Widget。
func (s *Showcase) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	root := scene.Flex(scene.Vertical, th.Space.MD).SetWidth(1440)
	root.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Background)})
	root.SetPadding(geom.InLTRB(24, 16, 24, 16))

	root.Add(widgets.Label{Text: "无句柄画布渲染引擎 · 展示台", Size: th.Font.SizeTitle}.Build(ctx))

	row1 := scene.Flex(scene.Horizontal, th.Space.MD)
	row1.Add(cardWith(ctx, "动态波形 · 180 点", s.Wave))
	row1.Add(cardWith(ctx, "环形仪表 · 旋转星形", s.Gauge))
	row1.Add(cardWith(ctx, "渲染性能实测", s.Tip))
	root.Add(row1)

	row2 := scene.Flex(scene.Horizontal, th.Space.MD)
	row2.Add(cardWith(ctx, "实时折线 · 3 条曲线", s.Line))
	row2.Add(cardWith(ctx, "动态柱状图", s.Bar))
	root.Add(row2)

	row3 := scene.Flex(scene.Horizontal, th.Space.MD)
	row3.Add(cardWith(ctx, "环形饼图", s.Pie))
	row3.Add(cardWith(ctx, "架构与操作", widgets.Label{Text: "RenderNode → Canvas → DisplayList → Layer → damage"}))
	root.Add(row3)
	return root
}

func cardWith(ctx *ui.BuildContext, title string, child ui.Widget) *scene.Node {
	c := widgets.NewCard(child)
	c.Title = widgets.Label{Text: title}
	return c.Build(ctx)
}

func waveSamples(n int) []float64 {
	out := make([]float64, n)
	for i := range out {
		t := float64(i) / float64(n)
		out[i] = 0.5*math.Sin(t*18) + 0.3*math.Sin(t*41) + 0.2*math.Sin(t*7)
	}
	return out
}

// ── CanvasIME：输入框 + IME 状态 ─────────────────────────────────

// CanvasIME 是迁移后的输入法演示页。
type CanvasIME struct {
	ui.Base
	Input    *widgets.Input
	Status   *widgets.Label
	Progress widgets.Progress
	Clear    *widgets.Button
	Submit   *widgets.Button
}

// NewCanvasIME 构造输入法演示页。
func NewCanvasIME() *CanvasIME {
	c := &CanvasIME{
		Input:    widgets.NewInput("在这里输入中文…"),
		Status:   &widgets.Label{Text: "IME 就绪"},
		Progress: widgets.Progress{Value: 0.4, Width: 400, Height: 8},
		Clear:    widgets.NewButton("清空输入"),
		Submit:   widgets.NewButton("模拟组合串"),
	}
	c.Input.Width = 560
	c.Input.OnChanged = func(v string) {
		c.Status.Text = fmt.Sprintf("已输入 %d 字", len([]rune(v)))
	}
	c.Clear.OnClick = func() {
		c.Input.Value = ""
		c.Input.Editor().SetValue(input.EditingValue{})
		c.Input.MarkDirty()
	}
	c.Submit.OnClick = func() {
		// 模拟一次 IME 组合串上屏：设组合区再提交成中文。
		e := c.Input.Editor()
		e.InsertText("ni hao")
		e.SetComposingRegion(0, 6, "ni hao")
		e.CommitComposingWith("你好")
		c.Input.MarkDirty()
	}
	return c
}

// Build 实现 ui.Widget。
func (c *CanvasIME) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	root := scene.Flex(scene.Vertical, th.Space.MD).SetWidth(960)
	root.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Background)})
	root.SetPadding(geom.InLTRB(40, 22, 40, 22))
	root.Add(widgets.Label{Text: "无句柄画布 · 输入法", Size: th.Font.SizeTitle}.Build(ctx))
	root.Add(widgets.Label{Text: "焦点链 / IME 组合 / 选区替换统一走 C28", Size: th.Font.SizeBase}.Build(ctx))
	root.Add(c.Input.Build(ctx))
	btnRow := scene.Flex(scene.Horizontal, th.Space.SM)
	btnRow.Add(c.Submit.Build(ctx))
	btnRow.Add(c.Clear.Build(ctx))
	root.Add(btnRow)
	root.Add(c.Progress.Build(ctx))
	root.Add(c.Status.Build(ctx))
	return root
}

// ── DBModeler：表结构 + 关系画布 ────────────────────────────────

// DBField 是模型字段。
type DBField struct {
	Name string
	Type string
	PK   bool
}

// DBTable 是模型表。
type DBTable struct {
	Name   string
	Fields []DBField
}

// DBModeler 是迁移后的数据库建模页。
type DBModeler struct {
	ui.Base
	Tables []DBTable
	Table  *widgets.Table
}

// NewDBModeler 构造建模页。
func NewDBModeler() *DBModeler {
	tables := []DBTable{
		{Name: "users", Fields: []DBField{{"id", "int", true}, {"name", "varchar", false}, {"email", "varchar", false}}},
		{Name: "orders", Fields: []DBField{{"id", "int", true}, {"user_id", "int", false}, {"amount", "decimal", false}}},
		{Name: "items", Fields: []DBField{{"id", "int", true}, {"order_id", "int", false}, {"sku", "varchar", false}}},
	}
	cols := []widgets.TableColumn{
		{Title: "字段", Mode: widgets.ColumnAuto},
		{Title: "类型", Mode: widgets.ColumnAuto},
		{Title: "主键", Mode: widgets.ColumnFixed, Width: 60},
	}
	var rows [][]any
	for _, t := range tables {
		for _, f := range t.Fields {
			rows = append(rows, []any{t.Name + "." + f.Name, f.Type, f.PK})
		}
	}
	tbl := widgets.NewTable(cols, rows)
	tbl.Width, tbl.Height = 620, 300
	return &DBModeler{Tables: tables, Table: tbl}
}

// Build 实现 ui.Widget。
func (d *DBModeler) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	root := scene.Flex(scene.Horizontal, th.Space.LG).SetWidth(1100)
	root.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Background)})
	root.SetPadding(geom.InLTRB(24, 16, 24, 16))

	left := scene.Flex(scene.Vertical, th.Space.SM)
	left.Add(widgets.Label{Text: "数据库模型", Size: th.Font.SizeTitle}.Build(ctx))
	left.Add(d.Table.Build(ctx))
	root.Add(left)

	// 关系画布：CustomPaintNode 直接录制连线（dbmodeler/canvasview.go 的通道）
	canvas := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		drawRelations(cv, sz, d.Tables)
	}).SetSize2(420, 300)
	root.Add(canvas)
	return root
}

// drawRelations 画表之间的关系连线。
func drawRelations(cv render.Canvas, sz geom.Size, tables []DBTable) {
	th := geom.RGBA(120, 140, 170)
	paint := render.Stroke(th, 2).WithCap(render.CapRound)
	// 三张表等距横排，画 FK 连线
	centers := make([]geom.Point, len(tables))
	for i := range tables {
		centers[i] = geom.Pt(60+float32(i)*150, 60)
	}
	for i := 1; i < len(centers); i++ {
		cv.DrawLine(centers[i-1].X+40, centers[i-1].Y+20,
			centers[i].X-40, centers[i].Y+20, paint)
	}
}
