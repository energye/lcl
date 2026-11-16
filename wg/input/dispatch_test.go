package input

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
)

// box 造一个可命中的矩形（Opaque + 标签）。
func box(w, h float32, label string) *scene.Node {
	return scene.Box(scene.BoxSpec{Fill: render.Fill(geom.RGBA(200, 200, 200))}).
		SetSize2(w, h).
		SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: label})
}

// tree 造 父(60×60) → 子(20×20 @0,0) 的两层树。
func tree() (root, child *scene.Node) {
	child = box(20, 20, "child")
	root = scene.Stack(child).SetSize2(60, 60)
	root.SetHit(scene.HitSpec{Behavior: scene.HitTranslucent, Label: "root"})
	root.Layout(scene.Unbounded())
	return root, child
}

// ev 造一个指针事件。
func ev(k PointerKind, x, y float32) *PointerEvent {
	p := geom.Pt(x, y)
	return &PointerEvent{Kind: k, Pointer: MousePointer, Device: DeviceMouse, Pos: p, PosRoot: p}
}

func TestHitPathIncludesAncestors(t *testing.T) {
	root, _ := tree()
	p := HitTest(root, geom.Pt(5, 5))
	if p.Len() != 2 || p.Nodes()[0] != root || p.Target().HitSpec().Label != "child" {
		t.Fatalf("路径错误: len=%d label=%s", p.Len(), p.Target().HitSpec().Label)
	}
	if !approxPt(p.Local(), geom.Pt(5, 5)) {
		t.Fatalf("局部坐标 = %v", p.Local())
	}
}

func TestDispatchOrderCaptureTargetBubble(t *testing.T) {
	root, child := tree()
	d := NewDispatcher(root)
	var order []string
	d.Set(root, Handler{
		Capture: func(e *PointerEvent) bool { order = append(order, "root.capture"); return false },
		Bubble:  func(e *PointerEvent) bool { order = append(order, "root.bubble"); return false },
	})
	d.Set(child, Handler{
		Capture: func(e *PointerEvent) bool { order = append(order, "child.capture"); return false },
		Pointer: func(e *PointerEvent) bool { order = append(order, "child.target"); return false },
		Bubble:  func(e *PointerEvent) bool { order = append(order, "child.bubble"); return false },
	})
	d.Pointer(ev(PointerDown, 5, 5))
	want := []string{"root.capture", "child.capture", "child.target", "root.bubble"}
	if len(order) != len(want) {
		t.Fatalf("顺序 = %v，期望 %v", order, want)
	}
	for i := range want {
		if order[i] != want[i] {
			t.Fatalf("顺序 = %v，期望 %v", order, want)
		}
	}
}

func TestCaptureStopsPropagation(t *testing.T) {
	root, child := tree()
	d := NewDispatcher(root)
	hit := false
	d.Set(root, Handler{Capture: func(e *PointerEvent) bool { return true }})
	d.Set(child, Handler{Pointer: func(e *PointerEvent) bool { hit = true; return false }})
	d.Pointer(ev(PointerDown, 5, 5))
	if hit {
		t.Fatal("捕获阶段返回 true 后不该再进目标阶段")
	}
}

func TestPointerCaptureRoutesOutside(t *testing.T) {
	root, child := tree()
	d := NewDispatcher(root)
	moves := 0
	d.Set(child, Handler{
		Pointer: func(e *PointerEvent) bool {
			if e.Kind == PointerMove {
				moves++
			}
			return false
		},
	})
	d.Pointer(ev(PointerDown, 5, 5))
	// 移到远处（落在 root 而不是 child 上）
	d.Pointer(ev(PointerMove, 55, 55))
	if moves != 1 {
		t.Fatalf("捕获期间移出节点仍应收到 move，实际 %d 次", moves)
	}
	if d.Captured(MousePointer) != child {
		t.Fatal("捕获目标应保持为按下时的节点")
	}
	// 抬起后释放
	d.Pointer(ev(PointerUp, 55, 55))
	if d.Captured(MousePointer) != nil {
		t.Fatal("抬起后应释放捕获")
	}
	d.Pointer(ev(PointerMove, 55, 55))
	if moves != 1 {
		t.Fatalf("释放后不该再收到 move，实际 %d 次", moves)
	}
}

func TestHoverDiffOncePerTransition(t *testing.T) {
	root, child := tree()
	d := NewDispatcher(root)
	enters, leaves := 0, 0
	d.Set(child, Handler{
		Enter: func(e *PointerEvent) { enters++ },
		Leave: func(e *PointerEvent) { leaves++ },
	})
	d.Pointer(ev(PointerMove, 5, 5))
	d.Pointer(ev(PointerMove, 6, 6)) // 同一节点内移动
	if enters != 1 || leaves != 0 {
		t.Fatalf("同一节点内移动不该重复 Enter: enter=%d leave=%d", enters, leaves)
	}
	d.Pointer(ev(PointerMove, 50, 50)) // 移出子树
	if leaves != 1 {
		t.Fatalf("移出应发 Leave: %d", leaves)
	}
	d.Pointer(ev(PointerMove, 7, 7))
	if enters != 2 {
		t.Fatalf("再次进入应发 Enter: %d", enters)
	}
}

func TestCancelReleasesAndRolls(t *testing.T) {
	root, child := tree()
	d := NewDispatcher(root)
	cancelled := 0
	var downPos geom.Point
	d.Set(child, Handler{
		Pointer: func(e *PointerEvent) bool {
			if e.Kind == PointerDown {
				downPos = e.Pos
			}
			if e.Kind == PointerCancel {
				cancelled++
				if e.Pos != downPos {
					t.Errorf("Cancel 的局部坐标应回到按下位置: %v vs %v", e.Pos, downPos)
				}
			}
			return false
		},
	})
	d.Pointer(ev(PointerDown, 5, 5))
	d.Pointer(ev(PointerMove, 15, 15))
	// host 发现按键已被系统吞掉：合成 Cancel
	if !d.Cancel(MousePointer, ev(PointerCancel, 5, 5)) {
		// 目标处理器返回 false，所以 Cancel 的返回值是 false；这里只校验语义
	}
	if cancelled != 1 {
		t.Fatalf("应收到一次 Cancel，实际 %d", cancelled)
	}
	if d.Captured(MousePointer) != nil {
		t.Fatal("Cancel 后应释放捕获")
	}
	if d.Stats().Cancels != 1 {
		t.Fatalf("应计一次 Cancel: %+v", d.Stats())
	}
}

func TestWheelBubblesToParent(t *testing.T) {
	root, child := tree()
	d := NewDispatcher(root)
	var got []string
	d.Set(child, Handler{Wheel: func(e *WheelEvent) bool {
		got = append(got, "child")
		return false // 不处理 → 继续冒泡（滚动链）
	}})
	d.Set(root, Handler{Wheel: func(e *WheelEvent) bool {
		got = append(got, "root")
		return true
	}})
	if !d.Wheel(&WheelEvent{Pos: geom.Pt(5, 5), Delta: geom.Pt(0, 40)}) {
		t.Fatal("父节点应处理滚轮")
	}
	if len(got) != 2 || got[0] != "child" || got[1] != "root" {
		t.Fatalf("滚轮顺序 = %v", got)
	}
}

func TestKeyGoesToFocusThenBubbles(t *testing.T) {
	root, child := tree()
	d := NewDispatcher(root)
	var got []string
	d.Set(child, Handler{Key: func(e *KeyEvent) bool {
		got = append(got, "child")
		return false
	}})
	d.Set(root, Handler{Key: func(e *KeyEvent) bool {
		got = append(got, "root")
		return true
	}})
	d.SetFocus(child)
	if !d.Key(&KeyEvent{Kind: KeyDown, Code: CodeEnter}) {
		t.Fatal("应被父节点处理")
	}
	if len(got) != 2 || got[0] != "child" || got[1] != "root" {
		t.Fatalf("键盘顺序 = %v", got)
	}
}

func TestCursorFollowsDeepestExplicit(t *testing.T) {
	root, child := tree()
	child.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "child", Cursor: scene.CursorPointer})
	d := NewDispatcher(root)
	var seen []scene.Cursor
	d.OnCursorChanged(func(c scene.Cursor) { seen = append(seen, c) })
	d.Pointer(ev(PointerMove, 5, 5))
	if d.Cursor() != scene.CursorPointer {
		t.Fatalf("光标 = %v，应为 pointer", d.Cursor())
	}
	d.Pointer(ev(PointerMove, 50, 50))
	if d.Cursor() != scene.CursorDefault {
		t.Fatalf("移出后光标应为默认，实际 %v", d.Cursor())
	}
	if len(seen) != 2 {
		t.Fatalf("光标变化通知 %d 次，应为 2", len(seen))
	}
}

func TestFocusSwitchSendsOutThenIn(t *testing.T) {
	root, child := tree()
	d := NewDispatcher(root)
	var got []FocusKind
	d.Set(root, Handler{Focus: func(e *FocusEvent) bool { got = append(got, e.Kind); return true }})
	d.Set(child, Handler{Focus: func(e *FocusEvent) bool { got = append(got, e.Kind); return true }})
	d.SetFocus(child)
	d.SetFocus(root)
	// 首次 SetFocus 没有旧节点，只发 FocusIn：共 3 个事件
	if len(got) != 3 {
		t.Fatalf("焦点事件 = %v", got)
	}
	if got[0] != FocusIn || got[1] != FocusOut || got[2] != FocusIn {
		t.Fatalf("焦点顺序 = %v", got)
	}
}

// approxPt 判定两点近似相等。
func approxPt(a, b geom.Point) bool {
	const eps = 0.01
	d := a.Sub(b)
	return d.X > -eps && d.X < eps && d.Y > -eps && d.Y < eps
}
