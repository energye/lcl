package software

import (
	"math"
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
)

// rec 是测试辅助：录一条指令表。
func rec(fn func(c *render.Recorder)) *render.DisplayList {
	r := render.NewRecorder(render.NewDisplayList(16))
	fn(r)
	return r.Finish()
}

func newDst(w, h int) *raster.Backing {
	b := raster.NewBacking(w, h)
	if b == nil {
		panic("backing alloc failed")
	}
	return b
}

func draw(dl *render.DisplayList, b *raster.Backing) raster.Stats {
	be := New(Options{})
	return be.Draw(b, dl, b.Bounds())
}

func TestFillRectExact(t *testing.T) {
	b := newDst(40, 40)
	dl := rec(func(c *render.Recorder) {
		c.FillRect(geom.RectXYWH(10, 10, 20, 20), render.Fill(geom.RGBA(255, 0, 0)))
	})
	st := draw(dl, b)
	if st.Degraded != 0 {
		t.Fatalf("不应有降级: %+v", st)
	}
	if got := raster.AlphaBounds(b); got != geom.RectLTRB(10, 10, 30, 30) {
		t.Fatalf("包围盒 = %v", got)
	}
	if px := b.Pix()[15*40+15]; px>>24 != 255 {
		t.Fatalf("内部像素 alpha = %d", px>>24)
	}
	if px := b.Pix()[5*40+5]; px != 0 {
		t.Fatalf("外部像素 = %#08x", px)
	}
	if !b.Opaque() {
		// 局部填充不会让整块变成不透明
		t.Log("整块仍标记为非不透明（正确）")
	}
}

func TestFillRectFractionalAA(t *testing.T) {
	b := newDst(20, 20)
	dl := rec(func(c *render.Recorder) {
		c.FillRect(geom.RectXYWH(4.5, 4.5, 5, 5), render.Fill(geom.RGBA(0, 0, 0)))
	})
	draw(dl, b)
	// 左边缘半个像素覆盖
	edge := b.Pix()[6*20+4] >> 24
	if edge < 100 || edge > 155 {
		t.Fatalf("边缘覆盖率 = %d，期望约 128", edge)
	}
	// 内部全不透明
	if a := b.Pix()[7*20+7] >> 24; a != 255 {
		t.Fatalf("内部 alpha = %d", a)
	}
}

func TestFillRectClips(t *testing.T) {
	b := newDst(20, 20)
	dl := rec(func(c *render.Recorder) {
		c.Save()
		c.ClipRect(geom.RectXYWH(0, 0, 10, 20))
		c.FillRect(geom.RectXYWH(0, 0, 20, 20), render.Fill(geom.RGBA(255, 0, 0)))
		c.Restore()
	})
	draw(dl, b)
	if got := raster.AlphaBounds(b); got != geom.RectLTRB(0, 0, 10, 20) {
		t.Fatalf("裁剪后包围盒 = %v", got)
	}
}

func TestFillOval(t *testing.T) {
	b := newDst(60, 60)
	dl := rec(func(c *render.Recorder) {
		c.FillOval(geom.RectXYWH(10, 10, 40, 40), render.Fill(geom.RGBA(0, 128, 255)))
	})
	draw(dl, b)
	got := raster.AlphaBounds(b)
	// 圆的包围盒应贴合 10..50（±1px）
	if math.Abs(float64(got.Min.X)-10) > 1 || math.Abs(float64(got.Max.X)-50) > 1 {
		t.Fatalf("椭圆包围盒 = %v", got)
	}
	// 四角必须空
	if b.Pix()[10*60+10]>>24 != 0 {
		t.Fatal("左上角应透明（椭圆不覆盖角）")
	}
	// 圆心必须实
	if b.Pix()[30*60+30]>>24 != 255 {
		t.Fatal("圆心应为实心")
	}
}

func TestEvenOddDonut(t *testing.T) {
	b := newDst(60, 60)
	p := render.NewPath().AddCircle(30, 30, 20).AddCircle(30, 30, 10)
	dl := rec(func(c *render.Recorder) {
		paint := render.Fill(geom.RGBA(0, 0, 0)).WithFillRule(render.EvenOdd)
		c.FillPath(p, paint)
	})
	draw(dl, b)
	if a := b.Pix()[30*60+30] >> 24; a != 0 {
		t.Fatalf("圆心应镂空，alpha = %d", a)
	}
	if a := b.Pix()[30*60+15] >> 24; a != 255 {
		t.Fatalf("环体应实心，alpha = %d", a)
	}
}

func TestStrokeLineWidth(t *testing.T) {
	b := newDst(40, 40)
	dl := rec(func(c *render.Recorder) {
		c.DrawLine(5, 20, 35, 20, render.Stroke(geom.RGBA(0, 0, 0), 4))
	})
	draw(dl, b)
	got := raster.AlphaBounds(b)
	// 线宽 4 → 高 4（18..22），长 30（5..35）
	if math.Abs(float64(got.H())-4) > 1 {
		t.Fatalf("线高 = %v，期望约 4", got.H())
	}
	if math.Abs(float64(got.W())-30) > 1 {
		t.Fatalf("线长 = %v，期望约 30", got.W())
	}
}

func TestStrokeRectBorder(t *testing.T) {
	b := newDst(60, 60)
	dl := rec(func(c *render.Recorder) {
		p := render.NewPath().AddRect(geom.RectXYWH(10, 10, 40, 40))
		c.StrokePath(p, render.Stroke(geom.RGBA(0, 0, 0), 2))
	})
	draw(dl, b)
	// 中心必须是空的（没有填充）
	if a := b.Pix()[30*60+30] >> 24; a != 0 {
		t.Fatalf("描边矩形中心应透明，alpha = %d", a)
	}
	// 边框上必须是实的
	if a := b.Pix()[10*60+30] >> 24; a != 255 {
		t.Fatalf("上边框 alpha = %d", a)
	}
}

func TestStrokeRoundCapExtends(t *testing.T) {
	b := newDst(40, 40)
	dl := rec(func(c *render.Recorder) {
		c.DrawLine(15, 20, 25, 20, render.Stroke(geom.RGBA(0, 0, 0), 6).WithCap(render.CapRound))
	})
	draw(dl, b)
	got := raster.AlphaBounds(b)
	// 圆头会向两端各延伸 3，总长 10+6=16
	if math.Abs(float64(got.W())-16) > 1.5 {
		t.Fatalf("圆头长度 = %v，期望约 16", got.W())
	}
}

func TestDashCreatesGaps(t *testing.T) {
	b := newDst(60, 20)
	dl := rec(func(c *render.Recorder) {
		c.DrawLine(5, 10, 55, 10, render.Dashed(geom.RGBA(0, 0, 0), 2, 5, 5))
	})
	draw(dl, b)
	// 沿中线统计非透明像素，虚线应明显少于实线
	n := 0
	for x := 0; x < 60; x++ {
		if b.Pix()[10*60+x]>>24 != 0 {
			n++
		}
	}
	if n < 20 || n > 35 {
		t.Fatalf("虚线覆盖像素 = %d，期望约 25（50%% 占空比）", n)
	}
}

func TestLinearGradient(t *testing.T) {
	b := newDst(40, 10)
	g := render.NewLinearGradient(geom.Pt(0, 0), geom.Pt(40, 0),
		render.Stop{Offset: 0, Color: geom.RGBA(255, 0, 0)},
		render.Stop{Offset: 1, Color: geom.RGBA(0, 0, 255)},
	)
	dl := rec(func(c *render.Recorder) {
		c.FillRect(geom.RectXYWH(0, 0, 40, 10), render.FillWith(g))
	})
	draw(dl, b)
	left := b.Pix()[5*40+0]
	right := b.Pix()[5*40+39]
	if (left>>16)&0xFF < 200 {
		t.Fatalf("左端应偏红: %#08x", left)
	}
	if right&0xFF < 200 {
		t.Fatalf("右端应偏蓝: %#08x", right)
	}
}

func TestRadialGradient(t *testing.T) {
	b := newDst(40, 40)
	g := render.NewRadialGradient(geom.Pt(20, 20), 20,
		render.Stop{Offset: 0, Color: geom.RGBA(255, 255, 255)},
		render.Stop{Offset: 1, Color: geom.RGBA(0, 0, 0)},
	)
	dl := rec(func(c *render.Recorder) {
		c.FillRect(geom.RectXYWH(0, 0, 40, 40), render.FillWith(g))
	})
	draw(dl, b)
	center := (b.Pix()[20*40+20] >> 16) & 0xFF
	corner := (b.Pix()[0] >> 16) & 0xFF
	if center <= corner {
		t.Fatalf("圆心应比角落亮: center=%d corner=%d", center, corner)
	}
}

func TestTransformAndRotation(t *testing.T) {
	b := newDst(40, 40)
	dl := rec(func(c *render.Recorder) {
		c.Save()
		c.Translate(20, 20)
		c.Rotate(math.Pi / 4)
		c.FillRect(geom.RectXYWH(-5, -5, 10, 10), render.Fill(geom.RGBA(0, 200, 0)))
		c.Restore()
	})
	draw(dl, b)
	got := raster.AlphaBounds(b)
	// 旋转 45° 的 10×10 方块，包围盒约 14.1×14.1
	if got.W() < 12 || got.W() > 16 {
		t.Fatalf("旋转后宽度 = %v，期望约 14", got.W())
	}
	if math.Abs(float64(got.Center().X)-20) > 1 {
		t.Fatalf("旋转中心偏移: %v", got.Center())
	}
}

func TestNoAAGivesHardEdges(t *testing.T) {
	b := newDst(20, 20)
	be := New(Options{DisableAA: true})
	dl := rec(func(c *render.Recorder) {
		c.FillRect(geom.RectXYWH(4.5, 4.5, 5, 5), render.Fill(geom.RGBA(0, 0, 0)))
	})
	be.Draw(b, dl, b.Bounds())
	for y := 0; y < 20; y++ {
		for x := 0; x < 20; x++ {
			a := b.Pix()[y*20+x] >> 24
			if a != 0 && a != 255 {
				t.Fatalf("关闭 AA 后不应出现半透明像素: (%d,%d) alpha=%d", x, y, a)
			}
		}
	}
}

// TestDrawImagePaints：位图绘制必须真的落像素，且不是降级。
// 曾经后端把 OpDrawImage 一律计 degraded，导致只画文字的场景"看起来通过"，
// 位图本身完全不可见（真窗表现为空白框）。
func TestDrawImagePaints(t *testing.T) {
	b := newDst(20, 20)
	img := render.NewImageRef(4, 4)
	for i := range img.Pix {
		img.Pix[i] = geom.RGBA(220, 40, 40).PremulBGRA()
	}
	dl := rec(func(c *render.Recorder) {
		c.DrawImage(img, geom.RectXYWH(2, 2, 10, 10), render.ImageBilinear)
	})
	st := draw(dl, b)
	if st.Degraded != 0 {
		t.Fatalf("位图绘制不应降级，实际 degraded=%d", st.Degraded)
	}
	if st.Pixels < 90 {
		t.Fatalf("位图覆盖像素过少: %d", st.Pixels)
	}
	// 采样中心应为不透明红。
	p := b.Pix()[7*20+7]
	if r := (p >> 16) & 0xFF; r < 200 {
		t.Fatalf("位图像素颜色不对: %08X", p)
	}
	if a := p >> 24; a != 255 {
		t.Fatalf("不透明位图 alpha 应为 255，实际 %d", a)
	}
	// 目标矩形外必须干净。
	if p := b.Pix()[0]; p != 0 {
		t.Fatalf("位图溢出到目标矩形外: %08X", p)
	}
}

// TestDrawImageNearestVsBilinear：三档采样都要能画，且 Nearest 保持硬边。
func TestDrawImageNearestVsBilinear(t *testing.T) {
	img := render.NewImageRef(2, 1)
	img.Pix[0] = geom.RGBA(255, 0, 0).PremulBGRA()
	img.Pix[1] = geom.RGBA(0, 0, 255).PremulBGRA()

	for _, q := range []render.ImageQuality{render.ImageNearest, render.ImageBilinear, render.ImageMipmap} {
		b := newDst(16, 8)
		dl := rec(func(c *render.Recorder) {
			c.DrawImage(img, geom.RectXYWH(0, 0, 16, 8), q)
		})
		st := draw(dl, b)
		if st.Degraded != 0 || st.Pixels == 0 {
			t.Fatalf("quality=%d degraded=%d pixels=%d", q, st.Degraded, st.Pixels)
		}
	}
}

func TestEmptyDisplayList(t *testing.T) {
	b := newDst(10, 10)
	st := draw(render.NewDisplayList(0).Freeze(), b)
	if st.Ops != 0 {
		t.Fatalf("空指令表 Ops = %d", st.Ops)
	}
	if got := raster.AlphaBounds(b); !got.IsEmpty() {
		t.Fatalf("空指令表不应画出内容: %v", got)
	}
}
