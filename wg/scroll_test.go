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
