//go:build windows

package gdi

import (
	"math"
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/text"
)

func recDL(fn func(c *render.Recorder)) *render.DisplayList {
	r := render.NewRecorder(render.NewDisplayList(16))
	fn(r)
	return r.Finish()
}

func TestDIBLifecycle(t *testing.T) {
	d := NewDIB(32, 16, 24)
	if d == nil {
		t.Skip("当前会话没有 GDI（无声显示环境）")
	}
	defer d.Delete()
	if d.W != 32 || d.H != 16 {
		t.Fatalf("尺寸 = %dx%d", d.W, d.H)
	}
	if d.Stride < 32*3 {
		t.Fatalf("stride = %d", d.Stride)
	}
	if len(d.Bits()) != d.Stride*d.H {
		t.Fatalf("字节数 = %d，期望 %d", len(d.Bits()), d.Stride*d.H)
	}
	if len(d.Row(0)) != d.Stride {
		t.Fatal("Row 长度不符")
	}
	d.Clear(0x00FFFFFF)
	if d.Row(0)[0] != 0xFF || d.Row(0)[1] != 0xFF || d.Row(0)[2] != 0xFF {
		t.Fatalf("Clear 后像素 = %v", d.Row(0)[:3])
	}
	d.Delete()
	d.Delete() // 幂等
	if d.Bits() != nil {
		t.Fatal("删除后 Bits 应为 nil")
	}
}

func TestBackendShapePositionMatchesSoftware(t *testing.T) {
	if NewDIB(8, 8, 24) == nil {
		t.Skip("当前会话没有 GDI")
	}
	dl := recDL(func(c *render.Recorder) {
		c.FillOval(geom.RectXYWH(10, 10, 40, 40), render.Fill(geom.RGBA(0, 0, 0)))
	})

	be := New()
	defer be.Close()
	dst := raster.NewBacking(60, 60)
	st := be.Draw(dst, dl, dst.Bounds())
	if st.Degraded != 0 {
		t.Fatalf("GDI 后端不应降级: %+v", st)
	}
	got := raster.AlphaBounds(dst)
	// 与软件后端相同容差：±1px
	wantMin, wantMax := float32(10), float32(50)
	if math.Abs(float64(got.Min.X)-float64(wantMin)) > 1 || math.Abs(float64(got.Max.X)-float64(wantMax)) > 1 {
		t.Fatalf("GDI 椭圆包围盒 = %v，期望约 10..50", got)
	}
	// GDI 无 AA：内圈不能出现半透明
	if a := dst.Pix()[30*60+30] >> 24; a != 255 {
		t.Fatalf("GDI 填充应为硬边，中心 alpha = %d", a)
	}
}

func TestBackendStrokeWidth(t *testing.T) {
	if NewDIB(8, 8, 24) == nil {
		t.Skip("当前会话没有 GDI")
	}
	dl := recDL(func(c *render.Recorder) {
		c.DrawLine(5, 20, 35, 20, render.Stroke(geom.RGBA(0, 0, 0), 4))
	})
	be := New()
	defer be.Close()
	dst := raster.NewBacking(40, 40)
	be.Draw(dst, dl, dst.Bounds())
	got := raster.AlphaBounds(dst)
	if math.Abs(float64(got.H())-4) > 1.5 {
		t.Fatalf("GDI 线高 = %v，期望约 4", got.H())
	}
}

func TestTextPassDrawsPixels(t *testing.T) {
	d := NewDIB(200, 60, 24)
	if d == nil {
		t.Skip("当前会话没有 GDI")
	}
	defer d.Delete()
	d.Clear(0x00FFFFFF)

	l := text.LayoutText("Hello 你好", text.DefaultFont(16), text.Options{})
	dl := recDL(func(c *render.Recorder) {
		c.DrawText(l, geom.Pt(4, 4), geom.RGBA(0, 0, 0))
	})
	n := TextPass(d, dl, geom.RectXYWH(0, 0, 200, 60))
	if n == 0 {
		t.Fatal("TextPass 未报告绘制行数")
	}
	d.Flush()

	// 统计非白像素：文字必须真的落到 DIB 上
	dark := 0
	for y := 0; y < d.H; y++ {
		row := d.Row(y)
		for x := 0; x < d.W; x++ {
			px := row[x*3:]
			if px[0] < 200 || px[1] < 200 || px[2] < 200 {
				dark++
			}
		}
	}
	if dark < 20 {
		t.Fatalf("文字像素过少: %d", dark)
	}
}

func TestTextPassRespectsClip(t *testing.T) {
	d := NewDIB(200, 60, 24)
	if d == nil {
		t.Skip("当前会话没有 GDI")
	}
	defer d.Delete()
	d.Clear(0x00FFFFFF)

	l := text.LayoutText("Clipped Text", text.DefaultFont(16), text.Options{})
	// 把裁剪限制在左半部分，只画第一行的一半
	dl := recDL(func(c *render.Recorder) {
		c.ClipRect(geom.RectXYWH(0, 0, 30, 60))
		c.DrawText(l, geom.Pt(0, 2), geom.RGBA(0, 0, 0))
	})
	TextPass(d, dl, geom.RectXYWH(0, 0, 200, 60))
	d.Flush()

	for y := 0; y < d.H; y++ {
		row := d.Row(y)
		for x := 35; x < d.W; x++ {
			px := row[x*3:]
			if px[0] < 200 {
				t.Fatalf("裁剪失效：(%d,%d) 出现文字像素", x, y)
			}
		}
	}
}
