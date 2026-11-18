package wg

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// TestScrollWheelScrolls 锁定滚动视口的滚轮行为与偏移持久化。
func TestScrollWheelScrolls(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 300)

	s := &widgets.Scroll{Axis: scene.Vertical, Width: 200, Height: 100}
	s.Child = scrollRows(20)
	c.MountWidget(&ipPage{child: s})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if s.Offset != 0 {
		t.Fatalf("初始偏移 = %v", s.Offset)
	}
	pt := geom.Pt(50, 40)
	c.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(0, 48), Time: time.Now()})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if s.Offset != 48 {
		t.Fatalf("滚轮后偏移 = %v（应 48）", s.Offset)
	}
	// 反向滚到顶部后继续滚：到边界不消费（交给外层）。
	c.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(0, -48), Time: time.Now()})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if s.Offset != 0 {
		t.Fatalf("反向滚后偏移 = %v（应 0）", s.Offset)
	}
	// 再次滚到 >0 后重建（模拟其它状态变化）：Offset 应保留。
	c.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(0, 48), Time: time.Now()})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.MountWidget(&ipPage{child: s})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if s.Offset != 48 {
		t.Fatalf("重建后偏移 = %v（应保留 48）", s.Offset)
	}
}

// TestScrollBothAxes 锁定双轴滚动：Delta 分轴派发、两轴独立钳制、
// Offset/OffsetX 跨重建保留、条按各自溢出情况出现。
func TestScrollBothAxes(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 300)

	s := &widgets.Scroll{Axis: scene.Both, Width: 200, Height: 100}
	s.Child = bothContent{}
	c.MountWidget(&ipPage{child: s})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	// 内容 400×300 vs 视口 200×100：两轴都溢出 200。
	barV, barH := s.Bar()
	if barV == nil || barV.Max != 200 || barH == nil || barH.Max != 200 {
		t.Fatalf("条范围 = 纵:%v 横:%v，应各 200", barV.Max, barH.Max)
	}
	pt := geom.Pt(50, 40)
	// 纵向滚轮只动 Offset。
	c.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(0, 48), Time: time.Now()})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if s.Offset != 48 || s.OffsetX != 0 {
		t.Fatalf("纵滚后 offset=(%v,%v)，应 (48,0)", s.Offset, s.OffsetX)
	}
	// 横向滚轮只动 OffsetX。
	c.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(48, 0), Time: time.Now()})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if s.OffsetX != 48 || s.Offset != 48 {
		t.Fatalf("横滚后 offset=(%v,%v)，应 (48,48)", s.Offset, s.OffsetX)
	}
	// 越过纵向边界继续滚：不消费、不再变化。
	c.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(0, 1000), Time: time.Now()})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if s.Offset != 200 {
		t.Fatalf("滚到底 Offset = %v，应钳到 200", s.Offset)
	}
	c.PushEvent(&input.WheelEvent{Pos: pt, Delta: geom.Pt(0, 48), Time: time.Now()})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if s.Offset != 200 {
		t.Fatalf("边界外滚轮 Offset = %v，应保持 200", s.Offset)
	}
	// 重建保留（组件实例在 Build 外持有）。
	c.MountWidget(&ipPage{child: s})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if s.Offset != 200 || s.OffsetX != 48 {
		t.Fatalf("重建后 offset=(%v,%v)，应 (200,48)", s.Offset, s.OffsetX)
	}
}

// bothContent 固定 400×300 的内容：纵向与横向都溢出 200×100 视口。
type bothContent struct{}

func (bothContent) Build(ctx *ui.BuildContext) *scene.Node {
	box := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(255, 0, 0))})
	box.SetSize2(400, 300)
	return box
}

// scrollRows 造一个可滚动的高内容。
func scrollRows(n int) ui.Widget {
	return docRows{n: n}
}

type docRows struct{ n int }

func (d docRows) Build(ctx *ui.BuildContext) *scene.Node {
	col := scene.Flex(scene.Vertical, 4)
	for i := 0; i < d.n; i++ {
		col.Add((&widgets.Label{Text: "行"}).Build(ctx))
	}
	return col
}
