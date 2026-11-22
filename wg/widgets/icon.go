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
	// info-body / exclamation：实心圆底内的纯图形（Modal 预设确认框用——
	// antd 图标=实底圆白图形，圆底由组合方 Box 画，这里只画本体）。
	"info-body": {
		stroke: func(b iconBrush) { b.seg(12, 11, 12, 17.5) },
		fill:   func(b iconBrush) { b.circle(12, 7, 1.4) },
	},
	"exclamation": {
		stroke: func(b iconBrush) { b.seg(12, 4.5, 12, 13.5) },
		fill:   func(b iconBrush) { b.circle(12, 17.5, 1.4) },
	},
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

	// ── 第一百零九轮扩展（26 → 84）：方向 ─────────────────────────
	"up-circle": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(8, 14, 12, 10, 16, 14)
	}},
	"down-circle": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(8, 10, 12, 14, 16, 10)
	}},
	"left-circle": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(14, 8, 10, 12, 14, 16)
	}},
	"right-circle": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(10, 8, 14, 12, 10, 16)
	}},
	"double-left": {stroke: func(b iconBrush) {
		b.seg(11, 6, 6, 12, 11, 18).seg(18, 6, 13, 12, 18, 18)
	}},
	"double-right": {stroke: func(b iconBrush) {
		b.seg(6, 6, 11, 12, 6, 18).seg(13, 6, 18, 12, 13, 18)
	}},
	"arrow-left": {stroke: func(b iconBrush) {
		b.seg(20, 12, 4, 12).seg(10, 6, 4, 12, 10, 18)
	}},
	"arrow-right": {stroke: func(b iconBrush) {
		b.seg(4, 12, 20, 12).seg(14, 6, 20, 12, 14, 18)
	}},

	// ── 数据 ─────────────────────────────────────────────────────────
	"database": {stroke: func(b iconBrush) {
		b.p.AddOval(geom.RectXYWH(5, 3, 14, 5))
		b.seg(5, 5.5, 5, 12).seg(5, 12, 5, 18.5)
		b.seg(19, 5.5, 19, 12).seg(19, 12, 19, 18.5)
		b.p.AddArc(geom.RectXYWH(5, 9.5, 14, 5), 3.14159, 3.14159)
		b.p.AddArc(geom.RectXYWH(5, 16, 14, 5), 3.14159, 3.14159)
	}},
	"folder": {stroke: func(b iconBrush) {
		b.seg(3.5, 5, 10, 5, 12, 7.5, 20.5, 7.5, 20.5, 19, 3.5, 19, 3.5, 5)
	}},
	"folder-open": {stroke: func(b iconBrush) {
		b.seg(3.5, 5, 10, 5, 12, 7.5, 20.5, 7.5, 20.5, 10)
		b.seg(20.5, 10, 18.5, 19, 3.5, 19, 3.5, 5)
	}},
	"file": {stroke: func(b iconBrush) {
		b.seg(6, 3, 14, 3, 18, 7, 18, 21, 6, 21, 6, 3)
		b.seg(14, 3, 14, 7, 18, 7)
	}},
	"file-text": {stroke: func(b iconBrush) {
		b.seg(6, 3, 14, 3, 18, 7, 18, 21, 6, 21, 6, 3)
		b.seg(14, 3, 14, 7, 18, 7)
		b.seg(9, 11, 15, 11).seg(9, 14, 15, 14).seg(9, 17, 13, 17)
	}},
	"image": {stroke: func(b iconBrush) {
		b.seg(3.5, 3.5, 20.5, 3.5, 20.5, 20.5, 3.5, 20.5, 3.5, 3.5)
		b.circle(9, 9, 1.8)
		b.seg(5, 18, 11, 12, 14.5, 15.5, 17, 13, 20, 16.5)
	}},
	"mail": {stroke: func(b iconBrush) {
		b.seg(3, 5, 21, 5, 21, 19, 3, 19, 3, 5)
		b.seg(3.5, 6, 12, 12.5, 20.5, 6)
	}},
	"link": {stroke: func(b iconBrush) {
		b.seg(10, 14, 14, 10)
		b.seg(8.5, 9, 7, 7.5, 7, 5.5, 8.5, 4, 10.5, 4, 12.5, 5.5, 14, 7)
		b.seg(10, 6.5, 12.5, 4)
		b.seg(15.5, 15, 17, 16.5, 17, 18.5, 15.5, 20, 13.5, 20, 11.5, 18.5, 10, 17)
		b.seg(14, 17.5, 11.5, 20)
	}},
	"paper-clip": {stroke: func(b iconBrush) {
		b.seg(8, 12.5, 15, 5.5, 18, 8.5, 10, 16.5, 7, 13.5, 14.5, 6)
	}},
	"printer": {stroke: func(b iconBrush) {
		b.seg(7, 8, 7, 3.5, 17, 3.5, 17, 8)
		b.seg(5, 8, 19, 8, 19, 15, 17, 15)
		b.seg(7, 13, 7, 20.5, 17, 20.5, 17, 13)
		b.seg(7, 15, 5, 15, 5, 8)
	}},

	// ── 操作 ─────────────────────────────────────────────────────────
	"lock": {stroke: func(b iconBrush) {
		b.seg(5.5, 11, 18.5, 11, 18.5, 20, 5.5, 20, 5.5, 11)
		b.seg(8, 11, 8, 7, 9.5, 4.5, 12, 4.5, 14.5, 4.5, 16, 7, 16, 11)
		b.seg(12, 14, 12, 16.5)
	}},
	"unlock": {stroke: func(b iconBrush) {
		b.seg(5.5, 11, 18.5, 11, 18.5, 20, 5.5, 20, 5.5, 11)
		b.seg(8, 11, 8, 7, 9.5, 4.5, 12, 4.5, 14.5, 4.5, 16, 7, 16, 8)
		b.seg(12, 14, 12, 16.5)
	}},
	"key": {stroke: func(b iconBrush) {
		b.circle(8, 14, 4)
		b.seg(11, 11, 20, 2).seg(16, 6, 18.5, 8.5).seg(14, 8, 16, 10)
	}},
	"poweroff": {stroke: func(b iconBrush) {
		b.p.AddArc(geom.RectXYWH(5, 5, 14, 14), -2.2, 4.4)
		b.seg(12, 3, 12, 11)
	}},
	"reload": {stroke: func(b iconBrush) {
		b.p.AddArc(geom.RectXYWH(5, 5, 14, 14), 0.6, 4.6)
		b.seg(18, 3.5, 19.3, 8, 14.8, 9)
	}},
	"redo": {stroke: func(b iconBrush) {
		b.seg(4, 18, 6.5, 11, 15, 10.5)
		b.seg(12, 6, 16.5, 10.5, 12, 15)
	}},
	"undo": {stroke: func(b iconBrush) {
		b.seg(20, 18, 17.5, 11, 9, 10.5)
		b.seg(12, 6, 7.5, 10.5, 12, 15)
	}},
	"zoom-in": {stroke: func(b iconBrush) {
		b.circle(10.5, 10.5, 6.5)
		b.seg(15.5, 15.5, 20.5, 20.5)
		b.seg(7.8, 10.5, 13.2, 10.5).seg(10.5, 7.8, 10.5, 13.2)
	}},
	"zoom-out": {stroke: func(b iconBrush) {
		b.circle(10.5, 10.5, 6.5)
		b.seg(15.5, 15.5, 20.5, 20.5)
		b.seg(7.8, 10.5, 13.2, 10.5)
	}},
	"fullscreen": {stroke: func(b iconBrush) {
		b.seg(4, 9, 4, 4, 9, 4).seg(15, 4, 20, 4, 20, 9)
		b.seg(20, 15, 20, 20, 15, 20).seg(9, 20, 4, 20, 4, 15)
	}},
	"fullscreen-exit": {stroke: func(b iconBrush) {
		b.seg(9, 4, 9, 9, 4, 9).seg(15, 9, 20, 9, 20, 4)
		b.seg(20, 15, 20, 20, 15, 20).seg(9, 20, 4, 20, 4, 15)
	}},
	"drag": {stroke: func(b iconBrush) {
		b.seg(12, 3, 12, 9).seg(9, 6, 12, 3, 15, 6)
		b.seg(12, 15, 12, 21).seg(9, 18, 12, 21, 15, 18)
		b.seg(3, 12, 9, 12).seg(6, 9, 3, 12, 6, 15)
		b.seg(15, 12, 21, 12).seg(18, 9, 21, 12, 18, 15)
	}},

	// ── 编辑 ─────────────────────────────────────────────────────────
	"form": {stroke: func(b iconBrush) {
		b.seg(5, 4.5, 19, 4.5).seg(5, 10, 12, 10)
		b.seg(17.5, 9.5, 20, 12, 15.5, 19, 12.5, 20, 12, 16.5, 17.5, 9.5)
		b.seg(5, 15.5, 10, 15.5)
	}},
	"highlight": {stroke: func(b iconBrush) {
		b.seg(12.5, 4.5, 18, 10, 11, 17, 7.5, 17, 7, 13.5, 12.5, 4.5)
		b.seg(5, 20.5, 10, 20.5)
	}},
	"bold": {stroke: func(b iconBrush) {
		b.seg(7, 4, 7, 20)
		b.seg(7, 4, 13, 4, 16.5, 6.5, 16.5, 10.5, 13, 12, 7, 12)
		b.seg(7, 12, 14, 12, 17, 14.5, 17, 17.5, 14, 20, 7, 20)
	}},
	"italic": {stroke: func(b iconBrush) {
		b.seg(10, 4, 19, 4).seg(5, 20, 14, 20).seg(14.5, 4, 9.5, 20)
	}},
	"underline": {stroke: func(b iconBrush) {
		b.seg(7, 4, 7, 11, 8.5, 15.5, 12, 17, 15.5, 15.5, 17, 11, 17, 4)
		b.seg(5, 20.5, 19, 20.5)
	}},
	"align-left": {stroke: func(b iconBrush) {
		b.seg(4, 5, 20, 5).seg(4, 10, 14, 10).seg(4, 15, 20, 15).seg(4, 20, 14, 20)
	}},
	"align-center": {stroke: func(b iconBrush) {
		b.seg(5, 5, 19, 5).seg(8, 10, 16, 10).seg(5, 15, 19, 15).seg(8, 20, 16, 20)
	}},
	"align-right": {stroke: func(b iconBrush) {
		b.seg(4, 5, 20, 5).seg(10, 10, 20, 10).seg(4, 15, 20, 15).seg(10, 20, 20, 20)
	}},

	// ── 状态 / 媒体 ─────────────────────────────────────────────────
	"play-circle": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(10, 8.5, 15.5, 12, 10, 15.5, 10, 8.5)
	}},
	"pause-circle": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(9.5, 8.5, 9.5, 15.5).seg(14.5, 8.5, 14.5, 15.5)
	}},
	"sound": {stroke: func(b iconBrush) {
		b.seg(4, 9.5, 8, 9.5, 13, 5, 13, 19, 8, 14.5, 4, 14.5, 4, 9.5)
		b.p.AddArc(geom.RectXYWH(13, 7, 8, 10), -0.9, 1.8)
	}},
	"wifi": {stroke: func(b iconBrush) {
		b.p.AddArc(geom.RectXYWH(4, 6, 16, 16), -2.3, 4.6)
		b.p.AddArc(geom.RectXYWH(7, 9, 10, 10), -2.3, 4.6)
	}, fill: func(b iconBrush) { b.circle(12, 18.5, 1.6) }},
	"cloud": {stroke: func(b iconBrush) {
		b.seg(7, 18, 6, 18, 4, 16, 4, 13, 6, 11, 8.5, 11)
		b.seg(9, 8, 12, 6, 15, 7, 17, 10, 19.5, 11.5, 20, 13.5, 19, 16, 17, 18, 7, 18)
	}},
	"cloud-download": {stroke: func(b iconBrush) {
		b.seg(7, 13, 6, 13, 4, 11, 4, 8.5, 6, 6.5, 8.5, 6.5)
		b.seg(9, 3.5, 12, 1.5, 15, 2.5, 17, 5.5, 19.5, 7, 20, 9, 19, 11.5, 17, 13, 7, 13)
		b.seg(12, 11, 12, 18).seg(9, 15.5, 12, 18.5, 15, 15.5)
	}},
	"cloud-upload": {stroke: func(b iconBrush) {
		b.seg(7, 17, 6, 17, 4, 15, 4, 12.5, 6, 10.5, 8.5, 10.5)
		b.seg(9, 7.5, 12, 5.5, 15, 6.5, 17, 9.5, 19.5, 11, 20, 13, 19, 15.5, 17, 17, 7, 17)
		b.seg(12, 20, 12, 13).seg(9, 15.5, 12, 12.5, 15, 15.5)
	}},
	"thunderbolt": {stroke: func(b iconBrush) {
		b.seg(13, 2.5, 5.5, 13.5, 11, 13.5, 10, 21.5, 18.5, 10.5, 13, 10.5, 13, 2.5)
	}},
	"smile": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.p.AddArc(geom.RectXYWH(8, 10, 8, 8), 0.5, 2.1)
	}, fill: func(b iconBrush) {
		b.circle(9, 9.5, 1.1)
		b.circle(15, 9.5, 1.1)
	}},

	// ── 其他 ─────────────────────────────────────────────────────────
	"shopping-cart": {stroke: func(b iconBrush) {
		b.seg(3, 4.5, 6, 4.5, 8.5, 15, 19, 15, 21, 7.5, 7, 7.5)
	}, fill: func(b iconBrush) {
		b.circle(10, 19, 1.6)
		b.circle(17.5, 19, 1.6)
	}},
	"car": {stroke: func(b iconBrush) {
		b.seg(4, 13, 6.5, 7.5, 17.5, 7.5, 20, 13, 20, 17, 4, 17, 4, 13)
		b.seg(6.5, 7.5, 4, 13)
	}, fill: func(b iconBrush) {
		b.circle(7.5, 17.5, 1.6)
		b.circle(16.5, 17.5, 1.6)
	}},
	"phone": {stroke: func(b iconBrush) {
		b.seg(6, 3.5, 10, 3.5, 10.5, 8, 8.5, 10, 14, 15.5, 16, 13.5, 20.5, 14, 20.5, 18, 17, 20.5, 13, 18, 6, 11, 3.5, 7, 6, 3.5)
	}},
	"bell": {stroke: func(b iconBrush) {
		b.seg(6, 16, 6, 10, 8, 5.5, 12, 4.5, 16, 5.5, 18, 10, 18, 16, 20, 18.5, 4, 18.5, 6, 16)
		b.seg(10, 21, 12, 19.5, 14, 21)
	}},
	"book": {stroke: func(b iconBrush) {
		b.seg(4, 5, 4, 19, 6, 20.5, 20, 20.5, 20, 4, 6, 4, 4, 5)
		b.seg(8, 4, 8, 20.5)
	}},
	"tag-o": {stroke: func(b iconBrush) {
		b.seg(4, 4, 11, 4, 20, 13, 13, 20, 4, 11, 4, 4)
		b.circle(8, 8, 1.3)
	}},
	"audit": {stroke: func(b iconBrush) {
		b.seg(5, 3.5, 19, 3.5, 19, 20.5, 5, 20.5, 5, 3.5)
		b.seg(9, 7.5, 15, 7.5).seg(9, 11, 15, 11)
		b.seg(9, 15, 11, 15).seg(13, 18, 15, 14)
	}},
	"global": {stroke: func(b iconBrush) {
		b.circle(12, 12, 8.5)
		b.seg(3.5, 12, 20.5, 12)
		b.p.AddOval(geom.RectXYWH(8, 3.5, 8, 17))
	}},
	"apartment": {stroke: func(b iconBrush) {
		b.seg(12, 3, 12, 7).seg(7, 7, 17, 7)
		b.seg(7, 7, 7, 11).seg(17, 7, 17, 11)
		b.seg(4.5, 14, 9.5, 14, 9.5, 20.5, 4.5, 20.5, 4.5, 14)
		b.seg(14.5, 14, 19.5, 14, 19.5, 20.5, 14.5, 20.5, 14.5, 14)
		b.seg(7, 11, 7, 14).seg(17, 11, 17, 14)
	}},
	"team": {stroke: func(b iconBrush) {
		b.circle(9, 8.5, 3.2)
		b.seg(3.5, 19.5, 4, 16.5, 6, 14.5, 9, 14, 12, 14.5, 14, 16.5, 14.5, 19.5)
		b.circle(16.5, 9, 2.6)
		b.seg(16, 14.5, 18.5, 15.5, 20, 17.5, 20.5, 19.5)
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
	const spinPeriod = 800 * time.Millisecond
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
			// Spin：相位在**录制时**按墙钟计算——纯 Build 时固定相位在静止窗
			// 缓存下会永远静止（--audit 实测"加载图标不旋转"）。
			spinPhase := float32(time.Since(iconBase)%spinPeriod) / float32(spinPeriod)
			arc := render.NewPath().AddArc(geom.RectXYWH(4.5, 4.5, 15, 15),
				spinPhase*2*3.14159265, 4.2) // AddArc 第二参是扫掠角（非结束角）：相位只动 start，弧长恒 240°
			arc = arc.Transform(m) // Transform 返回新路径（不修改自身），必须接回
			cv.StrokePath(arc, render.Stroke(color, 2*k).WithCap(render.CapRound).WithJoin(render.JoinRound))
			return
		}
		if def.stroke != nil {
			p := render.NewPath()
			def.stroke(iconBrush{p})
			p = p.Transform(m) // 同上：丢弃返回值 = 图标按 24 视口原坐标绘制（M8 起）
			cv.StrokePath(p, render.Stroke(color, 2*k).WithCap(render.CapRound).WithJoin(render.JoinRound))
		}
		if def.fill != nil {
			p := render.NewPath()
			def.fill(iconBrush{p})
			p = p.Transform(m)
			cv.FillPath(p, render.Fill(color))
		}
	}).SetSize2(size, size)
	n.SetHit(scene.HitSpec{Label: "icon"})
	return n
}
