package widgets

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// ── Slider（M8 对齐 antd：marks 刻度 + range 双滑块）────────────────
//
// antd：<Slider range marks={{0:'0℃',50:'50℃'}} step tooltipOpen>
// 已对齐：单/双滑块、marks 刻度（线+文字）、键盘步进、禁用。
// 未实现（如实标注）：tooltip 取值气泡、step 分档吸附、垂直方向。

// SliderMarks 是刻度表：值 → 文字。
type SliderMarks map[float32]string

// Slider 是滑块。Range=true 时在 [Value, Value2] 双端取值。
type Slider struct {
	controlBase
	Min, Max  float32
	Value     float32
	Value2    float32 // range 模式的第二端
	Range     bool
	Step      float32 // 0 = 连续
	Marks     SliderMarks
	Width     float32
	Disabled  bool
	OnChange  func(v float32) // 单滑块 / range 的左端
	OnChange2 func(v float32) // range 的右端
}

// NewSlider 创建单滑块。
func NewSlider(min, max, value float32) *Slider {
	return &Slider{Min: min, Max: max, Value: value, Width: 160}
}

// NewSliderRange 创建双滑块。
func NewSliderRange(min, max, from, to float32) *Slider {
	return &Slider{Min: min, Max: max, Value: from, Value2: to, Range: true, Width: 160}
}

// snapStep 把值吸附到 Step 档（Step>0 时）。
func (b *Slider) snapStep(v float32) float32 {
	if b.Step <= 0 {
		return v
	}
	return b.Min + float32(round32((v-b.Min)/b.Step))*b.Step
}

func round32(v float32) int32 {
	if v >= 0 {
		return int32(v + 0.5)
	}
	return int32(v - 0.5)
}

// Build 实现 ui.Widget。
func (b *Slider) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	w := b.Width
	if w <= 0 {
		w = 160
	}
	const (
		trackH = float32(4)
		knob   = float32(14)
	)
	hasMarks := len(b.Marks) > 0
	bodyH := float32(24)
	markH := float32(0)
	if hasMarks {
		markH = 20
	}
	body := scene.Box(scene.BoxSpec{}).SetSize2(w, bodyH+markH)
	// 命中区覆盖整个 body（含 marks 区）：拖动不必精确落在 4px 轨道上。
	body.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "slider", Cursor: scene.CursorPointer, Disabled: b.Disabled})
	track := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceSunken), Radius: trackH / 2}).
		SetSize2(w, trackH)
	track.SetPositionKeepSize(true).SetStackPosition(geom.Pt(0, bodyH/2-trackH/2))
	body.Add(track)

	lo, hi := b.Min, b.Max
	if hi <= lo {
		hi = lo + 1
	}
	clamp := func(v float32) float32 {
		if v < lo {
			return lo
		}
		if v > hi {
			return hi
		}
		return v
	}
	// 轨道填充与滑块。禁用（antd disabled）：填充消失（轨道整体灰）、
	// 滑块描边转 colorBorder——整条滑块无主色。
	fillColor, knobBorder := th.Color.Primary, th.Color.Primary
	if b.Disabled {
		fillColor, knobBorder = th.Color.SurfaceDisabled, th.Color.Border
	}
	ratio := func(v float32) float32 { return (clamp(v) - lo) / (hi - lo) }
	place := func(v float32) float32 {
		x := w*ratio(v) - knob/2
		if x < 0 {
			return 0
		}
		if x > w-knob {
			return w - knob
		}
		return x
	}
	fill := func(from, to float32) *scene.Node {
		x0, x1 := ratio(from)*w, ratio(to)*w
		f := scene.Box(scene.BoxSpec{Fill: scene.Solid(fillColor), Radius: trackH / 2}).
			SetSize2(x1-x0, trackH)
		f.SetPositionKeepSize(true).SetStackPosition(geom.Pt(x0, 0))
		return f
	}
	if b.Range {
		track.Add(fill(b.Value, b.Value2))
		for i, v := range []float32{b.Value, b.Value2} {
			k := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface), Border: scene.Stroke(knobBorder, 2), Radius: knob / 2}).SetSize2(knob, knob)
			k.SetPositionKeepSize(true).SetStackPosition(geom.Pt(place(v), (trackH-knob)/2))
			_ = i
			track.Add(k)
		}
	} else {
		track.Add(fill(lo, b.Value))
		k := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface), Border: scene.Stroke(knobBorder, 2), Radius: knob / 2}).SetSize2(knob, knob)
		k.SetPositionKeepSize(true).SetStackPosition(geom.Pt(place(b.Value), (trackH-knob)/2))
		track.Add(k)
	}

	// marks：值→刻度圆点+文字（视觉对齐 antd 官方 slider-demo-mark：
	// 每个 mark 在轨道上是一个小圆点——选中范围内主色实心、范围外白底灰框；
	// 落在选中范围内的文字变主色）。
	if hasMarks {
		inSelected := func(mv float32) bool {
			if b.Range {
				return mv >= math32min(b.Value, b.Value2) && mv <= math32max(b.Value, b.Value2)
			}
			return mv <= b.Value
		}
		for mv, label := range b.Marks {
			x := w * ratio(mv)
			if x < 0 || x > w {
				continue
			}
			dot := float32(8)
			var dotNode *scene.Node
			if inSelected(mv) && !b.Disabled {
				dotNode = scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary), Radius: dot / 2}).SetSize2(dot, dot)
			} else {
				dotNode = scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface), Border: scene.Stroke(th.Color.Border, 1), Radius: dot / 2}).SetSize2(dot, dot)
			}
			dotNode.SetPositionKeepSize(true).SetStackPosition(geom.Pt(x-dot/2, bodyH/2-dot/2))
			body.Add(dotNode)
			lx := x - 14
			if lx < 0 {
				lx = 0
			}
			if lx > w-28 {
				lx = w - 28
			}
			txtColor := th.Color.TextMuted
			if inSelected(mv) {
				txtColor = th.Color.Primary
			}
			lbl := (&Label{Text: label, Size: th.Font.SizeXS, Color: txtColor}).Build(ctx)
			lbl.SetPositionKeepSize(true)
			lbl.SetStackPosition(geom.Pt(lx, bodyH/2+8))
			body.Add(lbl)
		}
	}

	// 取值：按点击/拖动位置换算。range 模式在**按下瞬间锁定手柄**
	// （antd 语义：拖哪端就动哪端，中途不翻转）。
	activeEnd := 0 // 0=未锁定 1=左端 2=右端
	apply := func(pos geom.Point, lock bool) {
		if b.Disabled {
			return
		}
		r := pos.X / w
		if r < 0 {
			r = 0
		}
		if r > 1 {
			r = 1
		}
		raw := b.snapStep(lo + r*(hi-lo))
		if !b.Range {
			if raw == b.Value {
				return
			}
			b.Value = raw
			b.MarkDirty()
			if b.OnChange != nil {
				b.OnChange(raw)
			}
			return
		}
		if lock && activeEnd == 0 {
			d1 := raw - b.Value
			if d1 < 0 {
				d1 = -d1
			}
			d2 := raw - b.Value2
			if d2 < 0 {
				d2 = -d2
			}
			if d1 <= d2 {
				activeEnd = 1
			} else {
				activeEnd = 2
			}
		}
		if activeEnd == 2 {
			if raw < b.Value {
				raw = b.Value
			}
			if raw == b.Value2 {
				return
			}
			b.Value2 = raw
			b.MarkDirty()
			if b.OnChange2 != nil {
				b.OnChange2(raw)
			}
			return
		}
		if raw > b.Value2 {
			raw = b.Value2
		}
		if raw == b.Value {
			return
		}
		b.Value = raw
		b.MarkDirty()
		if b.OnChange != nil {
			b.OnChange(raw)
		}
	}
	g := &ui.GestureDetector{
		// antd 语义：点击轨道/滑块直接跳到该位置（不必拖动）。
		OnTap:      func(pos geom.Point) { activeEnd = 0; apply(pos, true) },
		OnPanStart: func(pos geom.Point) { activeEnd = 0; apply(pos, true) },
		OnPan:      func(pos, _ geom.Point) { apply(pos, false) },
	}
	b.tap = g
	h := g.Handler()
	h.Key = func(e *input.KeyEvent) bool {
		if e.Kind == input.KeyUp || b.Disabled {
			return false
		}
		step := (hi - lo) / 20
		if b.Step > 0 {
			step = b.Step
		}
		switch e.Code {
		case input.CodeLeft:
			apply(geom.Pt((b.Value-step-lo)/(hi-lo)*w, 0), false)
			return true
		case input.CodeRight:
			apply(geom.Pt((b.Value+step-lo)/(hi-lo)*w, 0), false)
			return true
		}
		return ctx.Focus != nil && ctx.Focus.HandleKey(e)
	}

	b.attach(ctx, body, h, !b.Disabled)
	_ = fmt.Sprintf
	return body
}

// math32min/math32max 是 float32 的极值助手（避免引入 math 包的 float64 转换噪音）。
func math32min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func math32max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
