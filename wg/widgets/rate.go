package widgets

import (
	"math"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// ── Rate（M8 对齐 antd）─────────────────────────────────────────────
//
// antd：<Rate count allowHalf defaultValue onChange tooltips>
// 已对齐：count 数量、整星步进、悬停预览、禁用、清除（点当前值清零）。
// 未实现（如实标注）：allowHalf 半星、tooltips 文案、字符替代。

// Rate 是星级评分。
type Rate struct {
	controlBase
	Value    float32 // 当前评分（整星）
	Count    int     // 星数（默认 5）
	Disabled bool
	OnChange func(float32)

	hoverAt int // 悬停预览到的星（0 = 无）
}

// NewRate 创建评分（默认 5 星）。
func NewRate() *Rate { return &Rate{Count: 5} }

// Build 实现 ui.Widget。
func (r *Rate) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	count := r.Count
	if count <= 0 {
		count = 5
	}
	row := scene.Flex(scene.Horizontal, 4).SetCrossAlign(scene.CrossCenter)
	shown := int(r.Value)
	if r.hoverAt > 0 && !r.Disabled {
		shown = r.hoverAt
	}
	for i := 1; i <= count; i++ {
		color := th.Color.Border // 未选中灰
		if i <= shown {
			color = th.Color.Warning // antd rate 主色是 warning 黄
		}
		idx := i
		fill := i <= shown
		star := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
			p := render.NewPath()
			drawStar(p, sz.W/2, sz.H/2, sz.W*0.42)
			if fill {
				cv.FillPath(p, render.Fill(color))
			}
			cv.StrokePath(p, render.Stroke(color, 1.5).WithJoin(render.JoinRound))
		}).SetSize2(20, 20)
		star.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "rate-star", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			ctx.Input.Set(star, input.Handler{
				Enter: func(*input.PointerEvent) { r.hoverAt = idx; r.MarkDirty() },
				Leave: func(*input.PointerEvent) {
					if r.hoverAt == idx {
						r.hoverAt = 0
						r.MarkDirty()
					}
				},
				Pointer: func(ev *input.PointerEvent) bool {
					if r.Disabled || !ev.Down() {
						return false
					}
					next := float32(idx)
					if int(next) == int(r.Value) {
						next = 0 // antd：再点当前值清除
					}
					r.Value = next
					r.MarkDirty()
					if r.OnChange != nil {
						r.OnChange(next)
					}
					return true
				},
			})
		}
		row.Add(star)
	}
	r.attach(ctx, row, input.Handler{}, false)
	return row
}

// drawStar 画五角星（外接圆半径 rr，中心 cx,cy）。
func drawStar(p *render.Path, cx, cy, rr float32) {
	const inner = 0.424 // 内接圆比例
	var pts []geom.Point
	for i := 0; i < 10; i++ {
		ang := -math.Pi/2 + float32(i)*math.Pi/5
		rad := rr
		if i%2 == 1 {
			rad = rr * inner
		}
		pts = append(pts, geom.Pt(cx+rad*float32(math.Cos(float64(ang))), cy+rad*float32(math.Sin(float64(ang)))))
	}
	p.MoveTo(pts[0].X, pts[0].Y)
	for _, pt := range pts[1:] {
		p.LineTo(pt.X, pt.Y)
	}
	p.Close()
}
