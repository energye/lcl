package widgets

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// ── Scrollbar ──────────────────────────────────────────────────────

// Scrollbar 是滚动条组件（可单用，也被 Scroll 内嵌）。
//
// Value / Max / Viewport 都是沿滚动轴的像素量：Max = 内容长 - 视口长。
// 拖动滑块与点击轨道翻页都会调用 OnChange；调用方负责把新值应用到内容。
type Scrollbar struct {
	controlBase
	Axis      scene.Axis
	Length    float32 // 轨道长度（沿轴）
	Thickness float32 // 粗细（垂直=宽，水平=高）；0=8
	Value     float32 // 当前偏移
	Max       float32 // 可滚动距离（内容 - 视口）
	Viewport  float32 // 视口长度（决定滑块比例）
	OnChange  func(v float32)

	thumb    *scene.Node
	trackLen float32
	thumbLen float32
	dragging bool
	dragPtr  float32
	dragVal  float32
}

const (
	scrollbarThickness = 8
	scrollbarMinThumb  = 24
	scrollbarInset     = 2
)

// SetValue 原地更新滑块位置（拖动/滚轮路径用，不重建组件）。
func (b *Scrollbar) SetValue(v float32) {
	if v < 0 {
		v = 0
	}
	if v > b.Max {
		v = b.Max
	}
	if v == b.Value {
		return
	}
	b.Value = v
	b.syncThumb()
}

func (b *Scrollbar) thumbGeometry() (pos, length float32) {
	length = b.trackLen
	content := b.Viewport + b.Max
	if content > 0 && b.trackLen > 0 {
		length = b.trackLen * (b.Viewport / content)
	}
	if length < scrollbarMinThumb {
		length = scrollbarMinThumb
	}
	if length > b.trackLen {
		length = b.trackLen
	}
	span := b.trackLen - length
	if span > 0 && b.Max > 0 {
		pos = (b.Value / b.Max) * span
	}
	return
}

func (b *Scrollbar) syncThumb() {
	if b.thumb == nil {
		return
	}
	pos, _ := b.thumbGeometry()
	if b.Axis == scene.Vertical {
		b.thumb.SetStackPosition(geom.Pt(0, pos))
	} else {
		b.thumb.SetStackPosition(geom.Pt(pos, 0))
	}
}

// syncThumbPaint 交互路径专用：滑块已布局过时只改 offset（标绘制脏，
// 不触发重排），让滚动帧只剩"blit + 条带重画"；未布局时退回布局路径，
// 让尺寸/offset 由 placeStackLike 算定（否则首帧滑块位置错乱）。
//
// 拖动滑块/滚轮是高频操作；走 SetStackPosition 会 MarkNeedsLayout，让整页
// 重排、并阻断滚动 blit（blit 要求无布局脏）——这是"滚动/拖滑块卡顿"的根源。
func (b *Scrollbar) syncThumbPaint() {
	if b.thumb == nil {
		return
	}
	if !b.thumb.LaidOut() {
		b.syncThumb()
		return
	}
	pos, _ := b.thumbGeometry()
	o := geom.Point{}
	if p := b.thumb.Parent(); p != nil {
		o = p.ContentOrigin()
	}
	if b.Axis == scene.Vertical {
		b.thumb.SetOffsetPaint(geom.Pt(o.X, o.Y+pos))
	} else {
		b.thumb.SetOffsetPaint(geom.Pt(o.X+pos, o.Y))
	}
}

func (b *Scrollbar) emit(v float32) {
	if v < 0 {
		v = 0
	}
	if v > b.Max {
		v = b.Max
	}
	if v == b.Value {
		return
	}
	b.Value = v
	b.syncThumbPaint()
	if b.OnChange != nil {
		b.OnChange(v)
	}
}

// Build 实现 ui.Widget。
func (b *Scrollbar) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	thick := b.Thickness
	if thick <= 0 {
		thick = scrollbarThickness
	}
	vertical := b.Axis == scene.Vertical
	b.trackLen = b.Length
	pos, length := b.thumbGeometry()
	b.thumbLen = length

	track := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.SurfaceActive),
		Radius: thick / 2,
	})
	if vertical {
		track.SetSize2(thick, b.Length)
	} else {
		track.SetSize2(b.Length, thick)
	}
	track.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "scrollbar"})

	thumb := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.TextFaint),
		Radius: thick / 2,
	})
	if vertical {
		thumb.SetSize2(thick, length)
	} else {
		thumb.SetSize2(length, thick)
	}
	thumb.SetPositionKeepSize(true)
	if vertical {
		thumb.SetStackPosition(geom.Pt(0, pos))
	} else {
		thumb.SetStackPosition(geom.Pt(pos, 0))
	}
	thumb.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "scrollbar-thumb"})
	track.Add(thumb)
	b.thumb = thumb

	axisOf := func(p geom.Point) float32 {
		if vertical {
			return p.Y
		}
		return p.X
	}
	if ctx.Input != nil {
		// 拖动滑块：用 **根坐标**（滑块自身会动，局部坐标不可靠）。
		ctx.Input.Set(thumb, input.Handler{
			Pointer: func(ev *input.PointerEvent) bool {
				switch ev.Kind {
				case input.PointerDown:
					b.dragging = true
					b.dragPtr = axisOf(ev.PosRoot)
					b.dragVal = b.Value
					return true
				case input.PointerMove:
					if !b.dragging {
						return false
					}
					span := b.trackLen - b.thumbLen
					if span > 0 {
						b.emit(b.dragVal + (axisOf(ev.PosRoot)-b.dragPtr)/span*b.Max)
					}
					return true
				case input.PointerUp, input.PointerCancel:
					b.dragging = false
					return true
				}
				return false
			},
		})
		// 点击轨道翻页（滑块之外的区域；滑块在轨道之上单独命中）。
		ctx.Input.Set(track, input.Handler{
			Pointer: func(ev *input.PointerEvent) bool {
				if !ev.Down() {
					return false
				}
				pos, length := b.thumbGeometry()
				x := axisOf(ev.Pos)
				if x >= pos && x <= pos+length {
					return false
				}
				page := b.Viewport
				if page <= 0 {
					page = b.Max / 4
				}
				if x < pos {
					b.emit(b.Value - page)
				} else {
					b.emit(b.Value + page)
				}
				return true
			},
		})
	}
	return track
}
