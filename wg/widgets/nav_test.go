package widgets

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

func TestTabsSwitch(t *testing.T) {
	r := newControlRig()
	tabs := NewTabs(
		Tab{Key: "a", Label: "A", Content: Label{Text: "page-a"}},
		Tab{Key: "b", Label: "B", Content: Label{Text: "page-b"}},
	)
	node := r.mount(tabs)
	headers := node.Children()[0].Children()
	if len(headers) != 2 {
		t.Fatalf("页签数 = %d", len(headers))
	}
	b := headers[1].BoundsInRoot()
	p := geom.Pt(b.Min.X+b.W()/2, b.Min.Y+b.H()/2)
	now := time.Now()
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer, Pos: p, PosRoot: p, Time: now})
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerUp, Pointer: input.MousePointer, Pos: p, PosRoot: p, Time: now.Add(16 * time.Millisecond)})
	if tabs.Active != 1 {
		t.Fatalf("切换后 Active = %d", tabs.Active)
	}
}

func TestTooltipShowsAndHides(t *testing.T) {
	root := scene.Stack().SetSize2(400, 300)
	disp := input.NewDispatcher(root)
	root.Layout(scene.TightConstraints(geom.Size{W: 400, H: 300}))
	ov := overlay.New(overlay.Options{
		Viewport: func() geom.Rect { return geom.RectXYWH(0, 0, 400, 300) },
		Focus:    input.NewFocusManager(input.NewFocusNode(false)),
		Now:      time.Now,
	})
	ctx := &ui.BuildContext{Theme: theme.Default(), Input: disp, Overlay: ov}
	tip := &Tooltip{Text: "提示文字", Child: NewButton("按钮")}
	n := tip.Build(ctx)
	root.Add(n)
	root.Layout(scene.TightConstraints(geom.Size{W: 400, H: 300}))

	p := geom.Pt(20, 10)
	disp.Pointer(&input.PointerEvent{Kind: input.PointerHover, Pointer: input.MousePointer, Pos: p, PosRoot: p})
	if ov.Count() != 1 {
		t.Fatalf("悬停后浮层数 = %d", ov.Count())
	}
	// 与真实链路一致：Tooltip 打开后仍要能点击锚点（不被 OutsideClick 吞掉）。
	disp.Pointer(&input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer, Pos: p, PosRoot: p})
	if ov.Count() != 1 {
		t.Fatal("点击锚点不该关掉 Tooltip（否则下一次点击会被吞）")
	}
	disp.Pointer(&input.PointerEvent{Kind: input.PointerLeave, Pointer: input.MousePointer, Pos: p, PosRoot: p})
	if ov.Count() != 0 {
		t.Fatalf("离开后浮层数 = %d", ov.Count())
	}
}

func TestTooltipKeepsChildHandler(t *testing.T) {
	root := scene.Stack().SetSize2(200, 100)
	disp := input.NewDispatcher(root)
	root.Layout(scene.TightConstraints(geom.Size{W: 200, H: 100}))
	ctx := &ui.BuildContext{Theme: theme.Default(), Input: disp}
	btn := NewButton("确定")
	// 先让子组件注册自己的处理器，再包 Tooltip。
	child := btn.Build(ctx)
	tip := &Tooltip{Text: "t", Child: &prebuiltWidget{node: child}}
	n := tip.Build(ctx)
	if n != child {
		t.Fatal("Tooltip 应直接返回子节点")
	}
	h, ok := disp.Handler(n)
	if !ok || h.Pointer == nil || h.Enter == nil {
		t.Fatal("Tooltip 覆盖了子按钮的处理器，而没有合并 Enter")
	}
}

// prebuiltWidget 让测试可以直接塞一个已构建的节点给 Tooltip。
type prebuiltWidget struct{ node *scene.Node }

func (w *prebuiltWidget) Build(*ui.BuildContext) *scene.Node { return w.node }
