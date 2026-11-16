package diag

import (
	"testing"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
)

// BenchmarkFrame 用 go test -bench 测单帧全链路开销，
// 并可用 -cpuprofile 定位热点（--bench 只给聚合数字，不给调用栈）。
//
//	go test ./wg/diag -run XXX -bench Frame$ -cpuprofile cpu.out -benchtime 200x
//	go tool pprof -top -nodecount=25 cpu.out
func BenchmarkFrame(b *testing.B) {
	const w, h = 1440, 880
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(w, h)
	clk := frame.NewManualClock(0, frame.DefaultStep)
	c.SetClock(clk)
	c.SetPainter(BenchScene(w, h, 24))

	if _, err := c.FrameOnce(); err != nil {
		b.Fatal(err)
	}
	clk.Advance()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Root().Invalidate(c.Root().LocalBounds())
		if _, err := c.FrameOnce(); err != nil {
			b.Fatal(err)
		}
		clk.Advance()
	}
}

// BenchmarkClear 只测"全屏不透明矩形填充"这一项，用来回答
// "逐像素解析覆盖率是否值得走快速路径"。
func BenchmarkClear(b *testing.B) {
	const w, h = 1440, 880
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(w, h)
	c.SetClock(frame.NewManualClock(0, frame.DefaultStep))
	c.SetPainter(func(cv render.Canvas) {
		cv.FillRect(geom.RectXYWH(0, 0, w, h), render.Fill(geom.RGBA(250, 250, 252)))
	})
	if _, err := c.FrameOnce(); err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Root().Invalidate(c.Root().LocalBounds())
		if _, err := c.FrameOnce(); err != nil {
			b.Fatal(err)
		}
	}
}
