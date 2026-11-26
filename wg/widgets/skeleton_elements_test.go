package widgets

import (
	"testing"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

type skCap struct {
	child ui.Widget
	ctx   *ui.BuildContext
}

func (c *skCap) Build(ctx *ui.BuildContext) *scene.Node { c.ctx = ctx; return c.child.Build(ctx) }

// TestSkeletonElements 三元素形态尺寸与圆角。
func TestSkeletonElements(t *testing.T) {
	// Button：三档高度 24/32/40。
	for _, tc := range []struct {
		size ui.Size
		want float32
	}{{ui.SizeSmall, 24}, {ui.SizeMedium, 32}, {ui.SizeLarge, 40}} {
		if got := skelHeight(tc.size); got != tc.want {
			t.Errorf("skelHeight(%v)=%v，want %v", tc.size, got, tc.want)
		}
	}
	// Shape 圆角：default 4、square 0、round/circle h/2。
	if skelRadius(SkelShapeDefault, 32) != 4 {
		t.Fatal("default 圆角应 4")
	}
	if skelRadius(SkelShapeSquare, 32) != 0 {
		t.Fatal("square 圆角应 0")
	}
	if skelRadius(SkelShapeRound, 32) != 16 {
		t.Fatal("round 圆角应 h/2")
	}
	if skelRadius(SkelShapeCircle, 32) != 16 {
		t.Fatal("circle 圆角应 h/2")
	}
	// 真实挂载后测尺寸（Size() 需布局才有效）。
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(600, 400)
	var root *scene.Node
	cap := &skCap{child: nodeFn(func(ctx *ui.BuildContext) *scene.Node {
		col := scene.Flex(scene.Vertical, 8)
		col.Add(SkeletonButton{Size: ui.SizeMedium, Shape: SkelShapeCircle}.Build(ctx))
		col.Add(SkeletonImage{}.Build(ctx))
		col.Add(SkeletonInput{}.Build(ctx))
		root = col
		return col
	})}
	c.MountWidget(cap)
	c.FrameOnce()
	bt := root.Children()[0]
	if bt.Size().W != bt.Size().H {
		t.Fatalf("circle 按钮应宽=高，got %v×%v", bt.Size().W, bt.Size().H)
	}
	im := root.Children()[1]
	if im.Size().W != 96 || im.Size().H != 96 {
		t.Fatalf("image 默认应 96×96，got %v×%v", im.Size().W, im.Size().H)
	}
	in := root.Children()[2]
	if in.Size().W != 200 {
		t.Fatalf("input 默认宽应 200，got %v", in.Size().W)
	}
}
