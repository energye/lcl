package wg

import (
	"fmt"
	"testing"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// richScrollContent 造一个文字+色块混合的可滚动内容（文字是滚动帧的主要成本，
// 也最考验 blit 的像素等价性）。
type richScrollContent struct{ rows int }

func (r richScrollContent) Build(ctx *ui.BuildContext) *scene.Node {
	col := scene.Flex(scene.Vertical, 6)
	for i := 0; i < r.rows; i++ {
		box := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(uint8(40*i), 120, 200))})
		box.SetPadding(geom.InLTRB(8, 6, 8, 6))
		box.Add((&widgets.Label{Text: fmt.Sprintf("第 %02d 行 内容文字 ABC", i + 1)}).Build(ctx))
		col.Add(box)
	}
	return col
}

// backingBytes 复制 backing 的像素字节，用于逐字节比较。
func backingBytes(b interface {
	W() int
	H() int
	Row(int) []uint32
}) []byte {
	if b == nil {
		return nil
	}
	w, h := b.W(), b.H()
	out := make([]byte, 0, w*h*4)
	for y := 0; y < h; y++ {
		row := b.Row(y)
		for x := 0; x < w; x++ {
			p := row[x]
			out = append(out, byte(p), byte(p>>8), byte(p>>16), byte(p>>24))
		}
	}
	return out
}

// TestScrollBlitPixelEqual 验证滚动 blit 与全量重画**逐像素一致**。
//
// 同一段滚动序列跑两遍：一遍开 blit，一遍关 blit，每一步的 backing 必须
// 字节相同。任何花屏/残影/错位都会在这里暴露——这是 scroll-blit 上线的
// 硬门禁（软件渲染没有"看起来差不多"，必须完全相同）。
func TestScrollBlitPixelEqual(t *testing.T) {
	run := func(disableBlit bool) [][]byte {
		c := NewHeadless(DefaultOptions())
		defer c.Close()
		c.Resize(240, 200)
		c.disableScrollBlit = disableBlit

		s := &widgets.Scroll{Axis: scene.Vertical, Width: 200, Height: 160}
		s.Child = richScrollContent{rows: 30}
		c.MountWidget(&ipPage{child: s})
		if _, err := c.FrameOnce(); err != nil {
			t.Fatal(err)
		}
		var frames [][]byte
		pt := geom.Pt(50, 40)
		// 逐次滚轮向下滚动（走真实滚轮路径：Scroll.Offset + 视图偏移）。
		for step := 0; step < 10; step++ {
			c.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(0, 40), Time: time.Now()})
			if _, err := c.FrameOnce(); err != nil {
				t.Fatal(err)
			}
			frames = append(frames, backingBytes(c.Backing()))
		}
		return frames
	}

	withBlit := run(false)
	noBlit := run(true)
	if len(withBlit) != len(noBlit) || len(withBlit) == 0 {
		t.Fatalf("帧数不一致: blit=%d noBlit=%d", len(withBlit), len(noBlit))
	}
	for i := range withBlit {
		a, b := withBlit[i], noBlit[i]
		if len(a) != len(b) {
			t.Fatalf("第 %d 帧字节数不一致", i)
		}
		for j := range a {
			if a[j] != b[j] {
				t.Fatalf("第 %d 帧第 %d 字节不同: blit=%d noBlit=%d（blit 与全量重画不一致）", i, j, a[j], b[j])
			}
		}
	}
}

// wideContent 造一个横向很宽的内容（横向滚动测试用）。
type wideContent struct{ cols int }

func (r wideContent) Build(ctx *ui.BuildContext) *scene.Node {
	row := scene.Flex(scene.Horizontal, 6)
	for i := 0; i < r.cols; i++ {
		box := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(200, uint8(30*i), 90))})
		box.SetSize2(60, 120)
		box.SetPadding(geom.InLTRB(6, 6, 6, 6))
		box.Add((&widgets.Label{Text: fmt.Sprintf("列%d", i + 1)}).Build(ctx))
		row.Add(box)
	}
	return row
}

// TestScrollBlitHorizontalPixelEqual 横向滚动 blit 与全量重画逐像素一致。
func TestScrollBlitHorizontalPixelEqual(t *testing.T) {
	run := func(disableBlit bool) [][]byte {
		c := NewHeadless(DefaultOptions())
		defer c.Close()
		c.Resize(240, 200)
		c.disableScrollBlit = disableBlit
		s := &widgets.Scroll{Axis: scene.Horizontal, Width: 200, Height: 160}
		s.Child = wideContent{cols: 20}
		c.MountWidget(&ipPage{child: s})
		if _, err := c.FrameOnce(); err != nil {
			t.Fatal(err)
		}
		var frames [][]byte
		pt := geom.Pt(50, 40)
		for step := 0; step < 8; step++ {
			c.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(40, 0), Time: time.Now()})
			if _, err := c.FrameOnce(); err != nil {
				t.Fatal(err)
			}
			frames = append(frames, backingBytes(c.Backing()))
		}
		return frames
	}
	a, b := run(false), run(true)
	if len(a) != len(b) || len(a) == 0 {
		t.Fatalf("帧数不一致: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			t.Fatalf("第 %d 帧长度不一致", i)
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				t.Fatalf("第 %d 帧第 %d 字节不同: blit=%d noBlit=%d", i, j, a[i][j], b[i][j])
			}
		}
	}
}

// TestScrollBlitBothPixelEqual 双轴滚动（Both）逐像素一致。
func TestScrollBlitBothPixelEqual(t *testing.T) {
	run := func(disableBlit bool) [][]byte {
		c := NewHeadless(DefaultOptions())
		defer c.Close()
		c.Resize(240, 200)
		c.disableScrollBlit = disableBlit
		s := &widgets.Scroll{Axis: scene.Both, Width: 200, Height: 160}
		s.Child = bothRichContent{}
		c.MountWidget(&ipPage{child: s})
		if _, err := c.FrameOnce(); err != nil {
			t.Fatal(err)
		}
		var frames [][]byte
		pt := geom.Pt(50, 40)
		for step := 0; step < 8; step++ {
			// 交替纵向、横向滚。
			if step%2 == 0 {
				c.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(0, 30), Time: time.Now()})
			} else {
				c.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(30, 0), Time: time.Now()})
			}
			if _, err := c.FrameOnce(); err != nil {
				t.Fatal(err)
			}
			frames = append(frames, backingBytes(c.Backing()))
		}
		return frames
	}
	a, b := run(false), run(true)
	if len(a) != len(b) || len(a) == 0 {
		t.Fatalf("帧数不一致: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			t.Fatalf("第 %d 帧长度不一致", i)
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				t.Fatalf("第 %d 帧第 %d 字节不同: blit=%d noBlit=%d", i, j, a[i][j], b[i][j])
			}
		}
	}
}

// bothRichContent 双轴都溢出的内容（网格）。
type bothRichContent struct{}

func (bothRichContent) Build(ctx *ui.BuildContext) *scene.Node {
	col := scene.Flex(scene.Vertical, 6)
	for r := 0; r < 20; r++ {
		row := scene.Flex(scene.Horizontal, 6)
		for i := 0; i < 12; i++ {
			box := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(uint8(20*r), uint8(15*i), 150))})
			box.SetSize2(50, 26)
			row.Add(box)
		}
		col.Add(row)
	}
	return col
}

// TestScrollBlitDragPixelEqual 拖动滑块（paint-only 滑块位移）与全量重画逐像素一致。
//
// 拖动走 syncThumbPaint（只改 offset），与滚轮（syncThumb，走布局）是两条路径；
// 这里直接改 Scroll.Offset 并 SetScrollOffset，模拟拖动后的视图偏移，
// 校验 blit 帧与全量重画帧字节一致。
func TestScrollBlitDragPixelEqual(t *testing.T) {
	run := func(disableBlit bool) [][]byte {
		c := NewHeadless(DefaultOptions())
		defer c.Close()
		c.Resize(240, 200)
		c.disableScrollBlit = disableBlit
		s := &widgets.Scroll{Axis: scene.Vertical, Width: 200, Height: 160}
		s.Child = richScrollContent{rows: 30}
		c.MountWidget(&ipPage{child: s})
		if _, err := c.FrameOnce(); err != nil {
			t.Fatal(err)
		}
		vbar, _ := s.Bar()
		var frames [][]byte
		for step := 0; step < 10; step++ {
			v := vbar.Max * float32(step+1) / 10
			vbar.SetValue(v)
			if vbar.OnChange != nil {
				vbar.OnChange(v)
			}
			if _, err := c.FrameOnce(); err != nil {
				t.Fatal(err)
			}
			frames = append(frames, backingBytes(c.Backing()))
		}
		return frames
	}
	a, b := run(false), run(true)
	if len(a) != len(b) || len(a) == 0 {
		t.Fatalf("帧数不一致: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			t.Fatalf("第 %d 帧长度不一致", i)
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				t.Fatalf("第 %d 帧第 %d 字节不同: blit=%d noBlit=%d", i, j, a[i][j], b[i][j])
			}
		}
	}
}
