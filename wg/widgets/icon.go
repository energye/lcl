package widgets

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// ── Icon ───────────────────────────────────────────────────────────
//
// 内置线性图标集（@ant-design/icons 的 outline 风格近似：24×24 视口、
// 2px 圆帽描边）。图形是 wg 自绘的 render.Path——**近似**而非官方 SVG
// 逐点复刻；需要精确图标时用 SetPath 注入自定义路径。
//
// 图标按名字取用，与 antd 图标名对齐（Plus/Search/Close/…）。

// Icon 是图标组件。
type Icon struct {
	Name  string // 图标名（见 iconPainters 的键）
	Size  float32
	Color geom.Color // 零值取主题 Text；Button 内部传 Foreground 实现 currentColor
	Spin  bool       // 旋转动画（loading）
}

// NewIcon 创建图标。
func NewIcon(name string) *Icon { return &Icon{Name: name} }

// iconDef 是单个图标的绘制定义（24×24 视口坐标）。
type iconDef struct {
	stroke func(b iconBrush) // 描边路径（可空）
	fill   func(b iconBrush) // 填充路径（可空，如圆点）
}

// iconBrush 给绘制函数提供链式路径构造。
type iconBrush struct{ p *render.Path }

func (b iconBrush) move(x, y float32) iconBrush { b.p.MoveTo(x, y); return b }
func (b iconBrush) line(x, y float32) iconBrush { b.p.LineTo(x, y); return b }

// seg 追加一段折线（点列）。
func (b iconBrush) seg(pts ...float32) iconBrush {
	for i := 0; i+1 < len(pts); i += 2 {
		if i == 0 {
			b.p.MoveTo(pts[0], pts[1])
			continue
		}
		b.p.LineTo(pts[i], pts[i+1])
	}
	return b
}

func (b iconBrush) circle(cx, cy, r float32) iconBrush {
	b.p.AddCircle(cx, cy, r)
	return b
}

// iconPainters 是内置图标集。坐标都在 24×24 视口内。
var iconPainters = map[string]iconDef{
	"plus":  {stroke: func(b iconBrush) { b.seg(12, 5, 12, 19).seg(5, 12, 19, 12) }},
	"minus": {stroke: func(b iconBrush) { b.seg(5, 12, 19, 12) }},
	"close": {stroke: func(b iconBrush) { b.seg(6, 6, 18, 18).seg(18, 6, 6, 18) }},
	"check": {stroke: func(b iconBrush) { b.seg(4.5, 12.5, 9.5, 17.5, 19.5, 6.5) }},
	"search": {stroke: func(b iconBrush) {
		b.circle(10.5, 10.5, 6.5)
		b.seg(15.5, 15.5, 20.5, 20.5)
	}},
	"up":    {stroke: func(b iconBrush) { b.seg(6, 15, 12, 9, 18, 15) }},
	"down":  {stroke: func(b iconBrush) { b.seg(6, 9, 12, 15, 18, 9) }},
	"left":  {stroke: func(b iconBrush) { b.seg(15, 6, 9, 12, 15, 18) }},
	"right": {stroke: func(b iconBrush) { b.seg(9, 6, 15, 12, 9, 18) }},
	"menu":  {stroke: func(b iconBrush) { b.seg(4, 6, 20, 6).seg(4, 12, 20, 12).seg(4, 18, 20, 18) }},
	"calendar": {stroke: func(b iconBrush) {
		b.move(4, 6).line(20, 6)
		b.seg(8, 3, 8, 7).seg(16, 3, 16, 7)
		b.seg(4.5, 4.5, 19.5, 4.5, 19.5, 19.5, 4.5, 19.5, 4.5, 4.5)
	}},
	"clock": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(12, 7.5, 12, 12.5, 15.5, 14.5)
	}},
	"user": {stroke: func(b iconBrush) {
		b.circle(12, 8, 3.8)
		b.seg(4.5, 20, 4.5, 19, 5.5, 16.5, 8, 15, 12, 15, 16, 15, 18.5, 16.5, 19.5, 19, 19.5, 20)
	}},
	"home": {stroke: func(b iconBrush) {
		b.seg(3.5, 11, 12, 3.5, 20.5, 11)
		b.seg(5.5, 9.5, 5.5, 20, 10, 20, 10, 14, 14, 14, 14, 20, 18.5, 20, 18.5, 9.5)
	}},
	"eye": {stroke: func(b iconBrush) {
		b.seg(2.5, 12, 5, 8.5, 8.5, 6.2, 12, 5.5, 15.5, 6.2, 19, 8.5, 21.5, 12)
		b.seg(21.5, 12, 19, 15.5, 15.5, 17.8, 12, 18.5, 8.5, 17.8, 5, 15.5, 2.5, 12)
		b.circle(12, 12, 3)
	}},
	"delete": {stroke: func(b iconBrush) {
		b.seg(4, 7, 20, 7).seg(9.5, 7, 9.5, 4.5, 14.5, 4.5, 14.5, 7)
		b.seg(6, 7, 7, 20, 17, 20, 18, 7)
		b.seg(10, 11, 10, 16.5).seg(14, 11, 14, 16.5)
	}},
	"edit": {stroke: func(b iconBrush) {
		b.seg(4, 20, 4.5, 15.5, 15.5, 4.5, 19.5, 8.5, 8.5, 19.5, 4, 20)
	}},
	"setting": {stroke: func(b iconBrush) {
		b.circle(12, 12, 7.5)
		b.circle(12, 12, 3)
		b.seg(12, 2.5, 12, 4.5).seg(12, 19.5, 12, 21.5)
		b.seg(2.5, 12, 4.5, 12).seg(19.5, 12, 21.5, 12)
		b.seg(5.3, 5.3, 6.7, 6.7).seg(17.3, 17.3, 18.7, 18.7)
		b.seg(18.7, 5.3, 17.3, 6.7).seg(6.7, 17.3, 5.3, 18.7)
	}},
	"info-circle": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(12, 11, 12, 16)
	}, fill: func(b iconBrush) { b.circle(12, 7.8, 1.1) }},
	"question-circle": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(9.2, 9.2, 9.8, 7.8, 11.2, 7.2, 13, 7.6, 13.8, 9, 13.4, 10.6, 12, 11.4, 12, 13)
	}, fill: func(b iconBrush) { b.circle(12, 16.2, 1.1) }},
	"warning-circle": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(12, 7.5, 12, 13)
	}, fill: func(b iconBrush) { b.circle(12, 16.2, 1.1) }},
	"check-circle": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(8, 12.5, 10.8, 15.2, 16.2, 9.2)
	}},
	"close-circle": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(9, 9, 15, 15).seg(15, 9, 9, 15)
	}},
	"download": {stroke: func(b iconBrush) {
		b.seg(12, 3.5, 12, 15).seg(7, 10.5, 12, 15.5, 17, 10.5)
		b.seg(4, 16.5, 4, 20.5, 20, 20.5, 20, 16.5)
	}},
	"upload": {stroke: func(b iconBrush) {
		b.seg(12, 15.5, 12, 4).seg(7, 8.5, 12, 3.5, 17, 8.5)
		b.seg(4, 16.5, 4, 20.5, 20, 20.5, 20, 16.5)
	}},
	"copy": {stroke: func(b iconBrush) {
		b.seg(9, 9, 20, 9, 20, 20, 9, 20, 9, 9)
		b.seg(15, 9, 15, 4, 4, 4, 4, 15, 9, 15)
	}},
	"more": {fill: func(b iconBrush) {
		b.circle(5, 12, 1.6)
		b.circle(12, 12, 1.6)
		b.circle(19, 12, 1.6)
	}},
	"star": {stroke: func(b iconBrush) {
		b.seg(12, 3.5, 14.8, 8.8, 20.8, 9.7, 16.4, 13.9, 17.5, 19.9, 12, 17, 6.5, 19.9, 7.6, 13.9, 3.2, 9.7, 9.2, 8.8, 12, 3.5)
		b.p.Close()
	}},
	"heart": {stroke: func(b iconBrush) {
		b.seg(12, 20, 6.5, 15, 4, 11, 3.5, 8.2, 5.2, 5.8, 8.5, 5.5, 10.7, 6.8, 12, 9, 13.3, 6.8, 15.5, 5.5, 18.8, 5.8, 20.5, 8.2, 20, 11, 17.5, 15, 12, 20)
		b.p.Close()
	}},
	"loading": {stroke: func(b iconBrush) {
		b.p.AddArc(geom.RectXYWH(4.5, 4.5, 15, 15), 0, 4.2)
	}},
}

// IconNames 返回内置图标名（诊断/文档用）。
func IconNames() []string {
	names := make([]string, 0, len(iconPainters))
	for n := range iconPainters {
		names = append(names, n)
	}
	// 排序保证输出稳定
	for i := 1; i < len(names); i++ {
		for j := i; j > 0 && names[j] < names[j-1]; j-- {
			names[j], names[j-1] = names[j-1], names[j]
		}
	}
	return names
}

// iconBase 是墙钟起点（Spin 用绝对时间，与动画窗同约定）。
var iconBase = time.Now()

// Build 实现 ui.Widget。
func (ic *Icon) Build(ctx *ui.BuildContext) *scene.Node {
	size := ic.Size
	if size <= 0 {
		size = 16
	}
	color := ic.Color
	if color == (geom.Color{}) {
		color = ctx.ThemeOr().Color.Text
	}
	var spinPhase float32
	if ic.Spin {
		const period = 800 * time.Millisecond
		spinPhase = float32(time.Since(iconBase)%period) / float32(period)
	}
	def := iconPainters[ic.Name]
	n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		if def.stroke == nil && def.fill == nil {
			// 未知图标：画一个橙色占位块，问题在视觉上直接暴露。
			cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(geom.RGBA(250, 173, 20)))
			return
		}
		k := sz.W / 24
		m := geom.Scale(k, k).Concat(geom.Translate(1, 1))
		if ic.Spin {
			// Spin：3/4 圆弧绕中心旋转（启动 loading 动画的组件负责续帧）。
			arc := render.NewPath().AddArc(geom.RectXYWH(4.5, 4.5, 15, 15),
				spinPhase*2*3.14159265, spinPhase*2*3.14159265+4.2)
			arc.Transform(m)
			cv.StrokePath(arc, render.Stroke(color, 2*k).WithCap(render.CapRound).WithJoin(render.JoinRound))
			return
		}
		if def.stroke != nil {
			p := render.NewPath()
			def.stroke(iconBrush{p})
			p.Transform(m)
			cv.StrokePath(p, render.Stroke(color, 2*k).WithCap(render.CapRound).WithJoin(render.JoinRound))
		}
		if def.fill != nil {
			p := render.NewPath()
			def.fill(iconBrush{p})
			p.Transform(m)
			cv.FillPath(p, render.Fill(color))
		}
	}).SetSize2(size, size)
	return n
}
