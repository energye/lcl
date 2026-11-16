// Package diag 是诊断命令矩阵的实现：--render-matrix、--golden、--bench、
// --damage 等。唯一入口是 cmd/wg-diag（main 只做参数解析与退出码）。
//
// 为什么把逻辑放这里而不是 main：每个矩阵都要能被 go test 直接调用，
// 否则"门禁命令"就只能靠人记着跑，最终一定会荒废。
//
// 依赖：任意（诊断需要观测所有层）。
package diag

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/raster/gdi"
	"github.com/energye/lcl/wg/raster/software"
	"github.com/energye/lcl/wg/render"
)

// MatrixCase 是图元矩阵中的一行。
type MatrixCase struct {
	Name string
	// Software/GDI 是两个后端画出的非透明包围盒。
	Software geom.Rect
	GDI      geom.Rect
	// CenterDelta 是两后端质心距离（px）。
	CenterDelta float32
	// BoxDelta 是两后端包围盒四边的最大差（px）。
	BoxDelta float32
	// Degraded 是后端自报的降级指令数。
	Degraded int
	// Skipped 表示该图元不参与双后端比对（例如渐变只有软件后端支持）。
	Skipped bool
	// Note 是补充说明。
	Note string
}

// OK 判定单行是否通过：位置误差 ≤ 1px，且没有静默丢弃。
func (c MatrixCase) OK() bool {
	if c.Skipped {
		return true
	}
	return c.BoxDelta <= 1.0 && c.CenterDelta <= 1.0
}

// MatrixReport 是 --render-matrix 的结果。
type MatrixReport struct {
	Cases   []MatrixCase
	Failed  int
	Skipped int
	Elapsed time.Duration
}

// OK 表示全部通过。
func (r MatrixReport) OK() bool { return r.Failed == 0 }

// String 返回可读报告。
func (r MatrixReport) String() string {
	var b []byte
	for _, c := range r.Cases {
		status := "ok  "
		switch {
		case c.Skipped:
			status = "skip"
		case !c.OK():
			status = "FAIL"
		}
		b = append(b, []byte(fmt.Sprintf("%s %-22s boxΔ=%.2f centerΔ=%.2f degraded=%d %s\n",
			status, c.Name, c.BoxDelta, c.CenterDelta, c.Degraded, c.Note))...)
	}
	b = append(b, []byte(fmt.Sprintf("render-matrix: %d 项，失败 %d，跳过 %d，用时 %v\n",
		len(r.Cases), r.Failed, r.Skipped, r.Elapsed))...)
	return string(b)
}

// caseSize 是每个图元的画布尺寸。
const caseSize = 120

// primitives 返回 20 个图元的录制函数。
//
// 选择标准：覆盖 03-ENGINE.md §4 的能力矩阵，且每一条都能被"包围盒 + 质心"
// 客观度量。纯颜色差异（渐变）不参与双后端比对，因为 GDI 不支持渐变，
// 硬要比只会得到"永远失败"的噪声。
func primitives() []struct {
	Name string
	Draw func(c render.Canvas)
	// Compare 为 false 时只验证软件后端能画出内容
	Compare bool
} {
	black := geom.RGBA(0, 0, 0)
	blue := geom.RGBA(30, 90, 200)
	out := []struct {
		Name string
		Draw func(c render.Canvas)
		Compare bool
	}{
		{"rect", func(c render.Canvas) { c.FillRect(geom.RectXYWH(20, 30, 60, 40), render.Fill(black)) }, true},
		{"rect-fractional", func(c render.Canvas) { c.FillRect(geom.RectXYWH(20.4, 30.6, 60.3, 40.2), render.Fill(black)) }, true},
		{"rrect-r8", func(c render.Canvas) { c.FillRRect(geom.RectXYWH(20, 30, 60, 40), 8, render.Fill(black)) }, true},
		{"rrect-r0", func(c render.Canvas) { c.FillRRect(geom.RectXYWH(20, 30, 60, 40), 0, render.Fill(black)) }, true},
		{"oval", func(c render.Canvas) { c.FillOval(geom.RectXYWH(20, 30, 60, 40), render.Fill(black)) }, true},
		{"circle", func(c render.Canvas) { c.FillOval(geom.RectXYWH(30, 30, 40, 40), render.Fill(black)) }, true},
		{"arc-90", func(c render.Canvas) {
			c.DrawArc(geom.RectXYWH(20, 20, 60, 60), 0, 1.5708, render.Stroke(black, 4))
		}, true},
		{"arc-270", func(c render.Canvas) {
			c.DrawArc(geom.RectXYWH(20, 20, 60, 60), 0.5, 4.712, render.Stroke(black, 3))
		}, true},
		{"line-h", func(c render.Canvas) { c.DrawLine(10, 60, 110, 60, render.Stroke(black, 2)) }, true},
		{"line-diag", func(c render.Canvas) { c.DrawLine(10, 10, 110, 110, render.Stroke(black, 3)) }, true},
		{"line-steep", func(c render.Canvas) { c.DrawLine(60, 5, 66, 115, render.Stroke(black, 2)) }, true},
		{"polyline", func(c render.Canvas) {
			p := render.NewPath().MoveTo(10, 90).LineTo(40, 30).LineTo(70, 95).LineTo(105, 35)
			c.StrokePath(p, render.Stroke(black, 3).WithJoin(render.JoinRound))
		}, true},
		{"quad", func(c render.Canvas) {
			p := render.NewPath().MoveTo(10, 90).QuadTo(60, 0, 110, 90)
			c.StrokePath(p, render.Stroke(black, 3))
		}, true},
		{"cubic", func(c render.Canvas) {
			p := render.NewPath().MoveTo(10, 90).CubicTo(30, 0, 90, 120, 110, 30)
			c.StrokePath(p, render.Stroke(black, 3))
		}, true},
		{"dash", func(c render.Canvas) { c.DrawLine(5, 60, 115, 60, render.Dashed(black, 4, 10, 6)) }, false},
		{"clip", func(c render.Canvas) {
			c.Save()
			c.ClipRect(geom.RectXYWH(40, 40, 40, 40))
			c.FillRect(geom.RectXYWH(0, 0, 120, 120), render.Fill(black))
			c.Restore()
		}, true},
		{"rotate-30", func(c render.Canvas) {
			c.Save()
			c.Translate(60, 60)
			c.Rotate(0.5236)
			c.FillRect(geom.RectXYWH(-25, -15, 50, 30), render.Fill(black))
			c.Restore()
		}, true},
		{"scale-2x", func(c render.Canvas) {
			c.Save()
			c.Scale(2, 2)
			c.FillRect(geom.RectXYWH(10, 10, 25, 20), render.Fill(blue))
			c.Restore()
		}, true},
		{"gradient-linear", func(c render.Canvas) {
			g := render.NewLinearGradient(geom.Pt(10, 0), geom.Pt(110, 0),
				render.Stop{Offset: 0, Color: geom.RGBA(255, 0, 0)},
				render.Stop{Offset: 1, Color: geom.RGBA(0, 0, 255)})
			c.FillRect(geom.RectXYWH(10, 30, 100, 60), render.FillWith(g))
		}, false},
		{"gradient-radial", func(c render.Canvas) {
			g := render.NewRadialGradient(geom.Pt(60, 60), 50,
				render.Stop{Offset: 0, Color: geom.RGBA(255, 255, 255)},
				render.Stop{Offset: 1, Color: geom.RGBA(0, 0, 0)})
			c.FillOval(geom.RectXYWH(20, 20, 80, 80), render.FillWith(g))
		}, false},
	}
	return out
}

// RenderMatrix 跑 20 个图元 × 2 后端的位置比对。
func RenderMatrix() MatrixReport {
	began := time.Now()
	var rep MatrixReport
	hasGDI := gdi.NewDIB(4, 4, 24) != nil

	for _, prim := range primitives() {
		dl := renderOne(prim.Draw)
		sw := raster.NewBacking(caseSize, caseSize)
		swStats := software.New(software.Options{}).Draw(sw, dl, sw.Bounds())

		cse := MatrixCase{
			Name:     prim.Name,
			Software: raster.AlphaBounds(sw),
			Degraded: swStats.Degraded,
		}
		if cse.Software.IsEmpty() {
			cse.Note = "软件后端未画出内容"
			rep.Failed++
			rep.Cases = append(rep.Cases, cse)
			continue
		}
		if !prim.Compare || !hasGDI {
			cse.Skipped = true
			if !prim.Compare {
				cse.Note = "仅软件后端（GDI 不支持该能力）"
			} else {
				cse.Note = "无 GDI 会话"
			}
			rep.Skipped++
			rep.Cases = append(rep.Cases, cse)
			continue
		}

		gd := raster.NewBacking(caseSize, caseSize)
		be := gdi.New()
		gdStats := be.Draw(gd, dl, gd.Bounds())
		be.Close()
		cse.GDI = raster.AlphaBounds(gd)
		cse.Degraded += gdStats.Degraded

		if cse.GDI.IsEmpty() {
			cse.Note = "GDI 后端未画出内容"
			rep.Failed++
			rep.Cases = append(rep.Cases, cse)
			continue
		}
		cse.BoxDelta = max4(
			absf(cse.Software.Min.X-cse.GDI.Min.X),
			absf(cse.Software.Min.Y-cse.GDI.Min.Y),
			absf(cse.Software.Max.X-cse.GDI.Max.X),
			absf(cse.Software.Max.Y-cse.GDI.Max.Y),
		)
		if sx, sy, ok1 := raster.AlphaCentroid(sw); ok1 {
			if gx, gy, ok2 := raster.AlphaCentroid(gd); ok2 {
				cse.CenterDelta = geom.Pt(sx-gx, sy-gy).Len()
			}
		}
		if !cse.OK() {
			rep.Failed++
		}
		rep.Cases = append(rep.Cases, cse)
	}
	rep.Elapsed = time.Since(began)
	return rep
}

// renderOne 录制一个图元。
func renderOne(draw func(c render.Canvas)) *render.DisplayList {
	rec := render.NewRecorder(render.NewDisplayList(16))
	draw(rec)
	return rec.Finish()
}

// WriteMatrixImages 把每个图元的两后端输出写成 PNG，便于人工核对。
func WriteMatrixImages(dir string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	for _, prim := range primitives() {
		dl := renderOne(prim.Draw)
		sw := raster.NewBacking(caseSize, caseSize)
		software.New(software.Options{}).Draw(sw, dl, sw.Bounds())
		if err := writePNG(filepath.Join(dir, prim.Name+"-software.png"), sw); err != nil {
			return err
		}
		if gdi.NewDIB(4, 4, 24) != nil {
			gd := raster.NewBacking(caseSize, caseSize)
			be := gdi.New()
			be.Draw(gd, dl, gd.Bounds())
			be.Close()
			if err := writePNG(filepath.Join(dir, prim.Name+"-gdi.png"), gd); err != nil {
				return err
			}
		}
	}
	return nil
}

// writePNG 把预乘 BGRA 缓冲写成 PNG（反预乘到非预乘，白底）。
func writePNG(path string, b *raster.Backing) error {
	img := image.NewNRGBA(image.Rect(0, 0, b.W(), b.H()))
	for y := 0; y < b.H(); y++ {
		row := b.Row(y)
		for x := 0; x < b.W(); x++ {
			px := row[x]
			a := px >> 24
			r := (px >> 16) & 0xFF
			g := (px >> 8) & 0xFF
			bl := px & 0xFF
			if a == 0 {
				img.SetNRGBA(x, y, color.NRGBA{255, 255, 255, 255})
				continue
			}
			img.SetNRGBA(x, y, color.NRGBA{
				R: uint8(min32(255, r*255/a)),
				G: uint8(min32(255, g*255/a)),
				B: uint8(min32(255, bl*255/a)),
				A: uint8(a),
			})
		}
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

// SortCases 按名称排序（报告稳定输出）。
func SortCases(cs []MatrixCase) {
	sort.Slice(cs, func(i, j int) bool { return cs[i].Name < cs[j].Name })
}

func absf(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}

func max4(a, b, c, d float32) float32 {
	m := a
	if b > m {
		m = b
	}
	if c > m {
		m = c
	}
	if d > m {
		m = d
	}
	return m
}

func min32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}
