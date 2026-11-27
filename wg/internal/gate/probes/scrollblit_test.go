//go:build gateprobe

package probes

import (
	"fmt"
	"testing"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// scrollProbeContent 造一个文字+色块混合的可滚动内容（文字是滚动帧主要成本）。
type scrollProbeContent struct{ rows int }

func (r scrollProbeContent) Build(ctx *ui.BuildContext) *scene.Node {
	col := scene.Flex(scene.Vertical, 6)
	for i := 0; i < r.rows; i++ {
		box := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(uint8(40*i), 120, 200))})
		box.SetPadding(geom.InLTRB(8, 6, 8, 6))
		box.Add((&widgets.Label{Text: fmt.Sprintf("第 %02d 行 内容文字 ABC", i + 1)}).Build(ctx))
		col.Add(box)
	}
	return col
}

// TestProbeScrollFrame 验证滚动帧间隔 ≤16ms（§6A.4 滚动行）与 blit 生效。
//
// 两个场景各用独立容器：滚轮滚动、拖动滑块（用户体感卡顿的直接来源）。
// 每帧都必须 ≤16ms 不掉帧；拖动帧的实际提交面积必须远小于整视口
// （证明 blit 生效，不是退回整视口重画）。
func TestProbeScrollFrame(t *testing.T) {
	const w, h = 1440, 880

	newScroll := func() (*wg.Container, *widgets.Scroll, *frame.ManualClock) {
		c := wg.NewHeadless(wg.DefaultOptions())
		c.Resize(w, h)
		clk := frame.NewManualClock(0, frame.DefaultStep)
		c.SetClock(clk)
		s := &widgets.Scroll{Axis: scene.Vertical, Width: 1200, Height: 800}
		s.Child = scrollProbeContent{rows: 40}
		c.MountWidget(s)
		if _, err := c.FrameOnce(); err != nil {
			t.Fatal(err)
		}
		clk.Advance()
		return c, s, clk
	}

	// 1) 滚轮滚动：逐帧推进、交替方向避免触底静止。
	c1, _, clk1 := newScroll()
	defer c1.Close()
	pt := geom.Pt(600, 400)
	for i := 0; i < 10; i++ { // 预热：首几帧含字体/缓存冷启动
		c1.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(0, 16), Time: time.Now()})
		if _, err := c1.FrameOnce(); err != nil {
			t.Fatal(err)
		}
		clk1.Advance()
	}
	var wheelSamples []time.Duration
	dir := 1
	for i := 0; i < 120; i++ {
		if i%40 == 39 {
			dir = -dir
		}
		c1.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(0, float32(dir)*16), Time: time.Now()})
		st, err := c1.FrameOnce()
		if err != nil {
			t.Fatal(err)
		}
		wheelSamples = append(wheelSamples, st.FrameDuration)
		clk1.Advance()
	}
	wheelP95 := frame.Percentile(wheelSamples, 95)
	if wheelP95 > 16*time.Millisecond {
		t.Fatalf("滚轮滚动帧 P95=%v，阈值 16ms（§6A 滚动不掉帧）", wheelP95)
	}

	// 2) 拖动滑块：独立容器，PointerDown 在滑块上，逐帧 PointerMove 下移。
	c2, _, clk2 := newScroll()
	defer c2.Close()
	var thumb *scene.Node
	c2.Walk(func(n *scene.Node) bool {
		if n.HitSpec().Label == "scrollbar-thumb" {
			thumb = n
			return false
		}
		return true
	})
	if thumb == nil {
		t.Fatal("未找到滑块节点 scrollbar-thumb")
	}
	tb := thumb.BoundsInRoot()
	start := geom.Pt((tb.Min.X+tb.Max.X)/2, (tb.Min.Y+tb.Max.Y)/2)
	c2.PushEvent(&input.PointerEvent{Kind: input.PointerDown, Pos: start, PosRoot: start, Time: time.Now()})
	if _, err := c2.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	clk2.Advance()
	var dragSamples []time.Duration
	var maxDamage float32
	for i := 0; i < 100; i++ {
		p := geom.Pt(start.X, start.Y+float32(i+1)*3)
		c2.PushEvent(&input.PointerEvent{Kind: input.PointerMove, Pos: p, PosRoot: p, Time: time.Now()})
		st, err := c2.FrameOnce()
		if err != nil {
			t.Fatal(err)
		}
		dragSamples = append(dragSamples, st.FrameDuration)
		if st.DamageArea > maxDamage {
			maxDamage = st.DamageArea
		}
		clk2.Advance()
	}
	dragP95 := frame.Percentile(dragSamples, 95)
	if dragP95 > 16*time.Millisecond {
		t.Fatalf("拖动滑块帧 P95=%v，阈值 16ms（§6A 滚动不掉帧）", dragP95)
	}
	// blit 生效证明：拖动帧实际提交面积远小于整视口（1200×800=96 万）。
	if maxDamage > 0.5*1200*800 {
		t.Fatalf("拖动帧提交面积 %.0f px² 超过视口一半——blit 未生效", maxDamage)
	}
	t.Logf("scroll: 滚轮 P95=%v 拖动 P95=%v 最大拖动提交面积=%.0f px²", wheelP95, dragP95, maxDamage)
}
