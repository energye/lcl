//go:build windows

// Package gdi 是 GDI 光栅后端：几何对照（A/B）与文字回放的通道。
//
// 两个职责要分清：
//   - **几何对照**：软件后端是主路径，GDI 提供"无抗锯齿"的硬边参照，
//     用来验证形状位置（--render-matrix 的 ≤1px 断言）与做性能对照（WG_AA=0）。
//   - **文字回放**：字形必须由 GDI 在设备空间画到最终 DIB 上，才能保留
//     ClearType（C6/C24）。这条通道与形状光栅无关，两个后端都走它。
//
// 能力边界：GDI 没有抗锯齿，Polygon 使用奇偶填充，不支持渐变与任意变换矩阵
// （本包在 Go 侧先把路径扁平化并变换，再交给 GDI，因此变换是支持的，
// 但填充规则只能是奇偶）。
package gdi

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/text"
)

// Backend 是 GDI 形状后端。
//
// 它的输出是"黑色遮罩"：把形状画成纯黑，再按亮度反推覆盖率写入目标缓冲。
// 这样同一份指令表可以在两个后端上比较几何位置，而不被抗锯齿与颜色差异干扰。
type Backend struct {
	aa   bool
	tol  float32
	dib  *DIB
	polys []render.Polyline
	pts  []point

	textOps int
	caps    raster.Caps
}

// New 创建 GDI 后端。
func New() *Backend {
	return &Backend{
		tol: 0.25,
		caps: raster.Caps{
			AA:             false, // GDI 无抗锯齿
			Stroke:         true,
			Dash:           false, // 虚线在软件后端实现；GDI 对照不做
			LinearGradient: false,
			RadialGradient: false,
			ClipRect:       true,
			ClipPath:       false,
			Text:           true,
			Image:          false,
			Transform:      true, // 在 Go 侧变换后交给 GDI
		},
	}
}

// Name 实现 raster.Backend。
func (b *Backend) Name() string { return "gdi" }

// Caps 实现 raster.Backend。
func (b *Backend) Caps() raster.Caps { return b.caps }

// TextOps 返回上次 Draw 遇到的文字指令数（这些指令由 TextPass 执行）。
func (b *Backend) TextOps() int { return b.textOps }

// Draw 实现 raster.Backend：把形状画进内部 DIB，再转成像素写入 dst。
func (b *Backend) Draw(dst *raster.Backing, dl *render.DisplayList, clip geom.Rect) raster.Stats {
	var st raster.Stats
	if dst == nil || dl == nil || dl.Empty() {
		return st
	}
	b.textOps = 0
	if b.dib == nil || b.dib.W != dst.W() || b.dib.H != dst.H() {
		if b.dib != nil {
			b.dib.Delete()
		}
		b.dib = NewDIB(dst.W(), dst.H(), 32)
		if b.dib == nil {
			// 没有 GDI 会话：整表计为降级，而不是静默不出图（C22）
			return raster.Stats{Degraded: dl.Len()}
		}
	}
	b.dib.Clear(0x00FFFFFF) // 白底

	var clips []geom.Rect
	cur := clip
	if cur.IsEmpty() {
		cur = dst.Bounds()
	}
	black := uint32(0x00000000)

	for _, cmd := range dl.Commands() {
		st.Ops++
		switch cmd.Op {
		case render.OpSave:
			clips = append(clips, cur)
		case render.OpRestore:
			if n := len(clips); n > 0 {
				cur = clips[n-1]
				clips = clips[:n-1]
			}
		case render.OpClipRect, render.OpClipPath:
			cur = cur.Intersect(cmd.Rect)

		case render.OpFillRect:
			r := cmd.Rect.Intersect(cur)
			if r.IsEmpty() {
				continue
			}
			b.dib.fillRect(int(r.Min.X+0.5), int(r.Min.Y+0.5), int(r.Max.X+0.5), int(r.Max.Y+0.5), black)
			st.Spans++

		case render.OpFillPath, render.OpStrokePath:
			path := dl.Path(cmd.Path)
			paint := dl.Paint(cmd.Paint)
			if path == nil || paint == nil {
				st.Degraded++
				continue
			}
			if paint.Shader != nil {
				st.Degraded++ // GDI 不支持渐变
				continue
			}
			saved := b.dib.clipRect([4]float32{cur.Min.X, cur.Min.Y, cur.Max.X, cur.Max.Y})
			b.polys = path.Flatten(b.tol, b.polys)
			if cmd.Op == render.OpFillPath {
				if paint.FillRule == render.EvenOdd || paint.FillRule == render.NonZero {
					// GDI Polygon 是奇偶填充：即使请求 nonzero 也照奇偶画，
					// 这是已知差异，--render-matrix 只比位置不比像素。
					for _, pl := range b.polys {
						b.pts = b.pts[:0]
						for _, p := range pl.Pts {
							q := cmd.Matrix.MapPoint(p)
							b.pts = append(b.pts, point{int32(q.X + 0.5), int32(q.Y + 0.5)})
						}
						if b.dib.fillPolygon(b.pts, black) {
							st.Spans++
						}
					}
				}
			} else {
				w := paint.StrokeWidthOr() * cmd.Matrix.ScaleFactor()
				width := int(w + 0.5)
				if width < 1 {
					width = 1
				}
				for _, pl := range b.polys {
					b.pts = b.pts[:0]
					for _, p := range pl.Pts {
						q := cmd.Matrix.MapPoint(p)
						b.pts = append(b.pts, point{int32(q.X + 0.5), int32(q.Y + 0.5)})
					}
					if b.dib.strokePolyline(b.pts, width, black) {
						st.Spans++
					}
				}
			}
			b.dib.restore(saved)

		case render.OpDrawText:
			b.textOps++

		case render.OpDrawImage:
			st.Degraded++
		}
	}

	b.dib.Flush()
	b.writeMask(dst)
	return st
}

// writeMask 把 DIB 的"黑形状/白底"转成覆盖率写入 dst。
func (b *Backend) writeMask(dst *raster.Backing) {
	for y := 0; y < dst.H() && y < b.dib.H; y++ {
		src := b.dib.Row(y)
		dstRow := dst.Row(y)
		for x := 0; x < dst.W() && x < b.dib.W; x++ {
			px := src[x*4:]
			// BGRX：亮度越低覆盖率越高
			lum := (int(px[0])*77 + int(px[1])*150 + int(px[2])*29) >> 8
			a := uint32(255 - lum)
			if a == 0 {
				continue
			}
			dstRow[x] = blendBlack(dstRow[x], a)
		}
	}
}

// blendBlack 把黑色按 alpha 混合到目标（预乘 BGRA）。
func blendBlack(dst, a uint32) uint32 {
	if a >= 255 {
		return 0xFF000000
	}
	inv := 255 - a
	da := dst >> 24
	dr := (dst >> 16) & 0xFF
	dg := (dst >> 8) & 0xFF
	db := dst & 0xFF
	oa := a + da*inv/255
	orr := dr * inv / 255
	og := dg * inv / 255
	ob := db * inv / 255
	return oa<<24 | orr<<16 | og<<8 | ob
}

// FillRuleSupport 返回 GDI 后端支持的填充规则（诊断用）。
func (b *Backend) FillRuleSupport() []string { return []string{"even-odd", "non-zero(按奇偶近似)"} }

// Close 释放内部 DIB。
func (b *Backend) Close() {
	if b.dib != nil {
		b.dib.Delete()
		b.dib = nil
	}
}

// 确保 text 包被引用（FontHandle 在 text.go 中使用；这里显式声明依赖方向）。
var _ = text.DefaultFont
