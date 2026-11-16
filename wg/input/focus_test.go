package input

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
)

// fn 造一个可聚焦节点并打标签（标签用于断言顺序）。
func fn(tag string) *FocusNode {
	n := NewFocusNode(true)
	n.Tag = tag
	return n
}

// tags 把焦点链转成标签切片，便于比较。
func tags(ns []*FocusNode) []string {
	out := make([]string, len(ns))
	for i, n := range ns {
		out[i] = n.Tag
	}
	return out
}

func eqTags(t *testing.T, got []*FocusNode, want ...string) {
	t.Helper()
	g := tags(got)
	if len(g) != len(want) {
		t.Fatalf("链 = %v，期望 %v", g, want)
	}
	for i := range want {
		if g[i] != want[i] {
			t.Fatalf("链 = %v，期望 %v", g, want)
		}
	}
}

func TestFocusOrderThenCreation(t *testing.T) {
	root := NewFocusNode(false)
	a := fn("a")
	b := fn("b")
	c := fn("c")
	c.Order = -1 // 显式排到最前
	root.AddChild(a).AddChild(b).AddChild(c)

	m := NewFocusManager(root)
	eqTags(t, m.List(), "c", "a", "b")
}

func TestFocusDFSOrder(t *testing.T) {
	// 根 → 组(容器) → [x,y] 与 z 并列：文档序是 DFS，不是层序。
	root := NewFocusNode(false)
	grp := NewFocusNode(false)
	x, y, z := fn("x"), fn("y"), fn("z")
	grp.AddChild(x).AddChild(y)
	root.AddChild(grp).AddChild(z)

	eqTags(t, NewFocusManager(root).List(), "x", "y", "z")
}

func TestFocusSkipsUnfocusable(t *testing.T) {
	root := NewFocusNode(false)
	dis := fn("disabled")
	dis.CanFocus = false
	root.AddChild(dis).AddChild(fn("ok"))
	eqTags(t, NewFocusManager(root).List(), "ok")
}

func TestTabWrapsInRootDomain(t *testing.T) {
	root := NewFocusNode(false)
	root.AddChild(fn("a")).AddChild(fn("b"))
	m := NewFocusManager(root)

	if !m.MoveNext(false) || m.Current().Tag != "a" {
		t.Fatalf("首次 Tab = %v", m.Current())
	}
	m.MoveNext(false)
	if m.Current().Tag != "b" {
		t.Fatalf("第二次 Tab = %v", m.Current())
	}
	m.MoveNext(false)
	if m.Current().Tag != "a" {
		t.Fatalf("根域应循环，得到 %v", m.Current())
	}
	m.MoveNext(true)
	if m.Current().Tag != "b" {
		t.Fatalf("Shift+Tab 应回到 b，得到 %v", m.Current())
	}
}

func TestTabFocusEventsOrder(t *testing.T) {
	root := NewFocusNode(false)
	a, b := fn("a"), fn("b")
	root.AddChild(a).AddChild(b)
	var log []string
	a.OnFocus = func() { log = append(log, "a.in") }
	a.OnBlur = func() { log = append(log, "a.out") }
	b.OnFocus = func() { log = append(log, "b.in") }
	b.OnBlur = func() { log = append(log, "b.out") }

	m := NewFocusManager(root)
	m.MoveNext(false)                         // a
	m.MoveNext(false)                         // b
	want := []string{"a.in", "a.out", "b.in"} // 旧节点先 Blur
	if len(log) != len(want) {
		t.Fatalf("事件 = %v", log)
	}
	for i := range want {
		if log[i] != want[i] {
			t.Fatalf("事件 = %v，期望 %v", log, want)
		}
	}
	if a.Focused() || !b.Focused() {
		t.Fatalf("聚焦标志错位: a=%v b=%v", a.Focused(), b.Focused())
	}
}

func TestTabJumpsDisabledCurrent(t *testing.T) {
	// 焦点节点被动态禁用后，下次 Tab 不能停在它上面。
	root := NewFocusNode(false)
	a, b := fn("a"), fn("b")
	root.AddChild(a).AddChild(b)
	m := NewFocusManager(root)
	m.MoveNext(false) // a
	a.CanFocus = false
	m.MoveNext(false)
	if m.Current() != b {
		t.Fatalf("应跳到 b，得到 %v", m.Current().Tag)
	}
}

// modalTree 造 外层 o1,o2 + 对话框(内含 d1,d2) + o3，返回根与域。
//
// 注意创建顺序即文档序（Children 按 (Order, 创建序) 排序，与插入顺序无关），
// 所以对话框容器必须先于 o3 创建，否则 o3 会跑到对话框子树前面。
func modalTree(t *testing.T) (root *FocusNode, dlg *FocusScope, o1, o2, d1, d2, o3 *FocusNode) {
	t.Helper()
	root = NewFocusNode(false)
	o1, o2 = fn("o1"), fn("o2")
	box := NewFocusNode(false)
	d1, d2 = fn("d1"), fn("d2")
	box.AddChild(d1).AddChild(d2)
	o3 = fn("o3")
	root.AddChild(o1).AddChild(o2).AddChild(box).AddChild(o3)
	dlg = NewFocusScope("dialog", box)
	return
}

func TestTrapScopeCycles(t *testing.T) {
	root, dlg, _, _, d1, d2, _ := modalTree(t)
	m := NewFocusManager(root)
	m.EnterScope(dlg)

	if m.Current() != d1 {
		t.Fatalf("进入域后焦点 = %v", m.Current().Tag)
	}
	m.MoveNext(false)
	if m.Current() != d2 {
		t.Fatalf("域内 Tab = %v", m.Current().Tag)
	}
	m.MoveNext(false) // Trap：绕回域内第一个，不能跑到 o3
	if m.Current() != d1 {
		t.Fatalf("模态域应循环，得到 %v", m.Current().Tag)
	}
	m.MoveNext(true)
	if m.Current() != d2 {
		t.Fatalf("Shift+Tab 绕回末位，得到 %v", m.Current().Tag)
	}
}

func TestNonTrapScopeMigratesOutward(t *testing.T) {
	root, dlg, _, _, d1, _, o3 := modalTree(t)
	dlg.SetTrap(false)
	m := NewFocusManager(root)
	m.EnterScope(dlg)
	if m.Current() != d1 {
		t.Fatalf("进入域后焦点 = %v", m.Current().Tag)
	}
	m.MoveNext(false) // d1 → d2
	if m.Current().Tag != "d2" {
		t.Fatalf("域内 = %v", m.Current().Tag)
	}
	m.MoveNext(false) // 出界 → 外层，且应从 o3 继续（不是从头 o1）
	if m.Current() != o3 {
		t.Fatalf("出域后 = %v，期望 o3", m.Current().Tag)
	}
}

func TestScopeSkippedWhenInactive(t *testing.T) {
	root, _, _, _, _, _, o3 := modalTree(t)
	m := NewFocusManager(root) // 未进入对话框
	seq := []string{}
	for i := 0; i < 4; i++ {
		m.MoveNext(false)
		seq = append(seq, m.Current().Tag)
	}
	// 未激活的域内控件不该出现在 Tab 链上
	for _, s := range seq {
		if s == "d1" || s == "d2" {
			t.Fatalf("Tab 链串进了未激活域: %v", seq)
		}
	}
	if seq[2] != o3.Tag {
		t.Fatalf("外层顺序 = %v", seq)
	}
}

func TestEnterExitScopeRestoresFocus(t *testing.T) {
	root, dlg, o1, _, _, _, _ := modalTree(t)
	m := NewFocusManager(root)
	m.Focus(o1)
	m.EnterScope(dlg)
	if m.Current().Tag == "o1" {
		t.Fatal("进入模态后不该还停在背景控件上")
	}
	m.ExitScope()
	if m.Current() != o1 {
		t.Fatalf("退出模态应还原焦点，得到 %v", m.Current().Tag)
	}
	if m.ActiveScope() != nil {
		t.Fatalf("活动域 = %v", m.ActiveScope())
	}
}

func TestTrapScopeRejectsOutsideFocus(t *testing.T) {
	root, dlg, o1, _, d1, _, _ := modalTree(t)
	m := NewFocusManager(root)
	m.EnterScope(dlg)
	if m.Focus(o1) {
		t.Fatal("模态外的节点不该被聚焦")
	}
	if m.Current() != d1 {
		t.Fatalf("焦点被抢走: %v", m.Current().Tag)
	}
}

func TestSuspendResumeRestoresFocus(t *testing.T) {
	root := NewFocusNode(false)
	a := fn("a")
	root.AddChild(a)
	m := NewFocusManager(root)
	m.Focus(a)

	m.Suspend()
	if m.Current() != nil || a.Focused() {
		t.Fatal("失活后应清空焦点")
	}
	if !m.Suspended() {
		t.Fatal("Suspended = false")
	}
	if m.MoveNext(false) {
		t.Fatal("失活期间 Tab 不该生效")
	}

	m.Resume()
	if m.Current() != a || !a.Focused() {
		t.Fatalf("恢复后焦点 = %v", m.Current())
	}
	if m.Stats().Suspends != 1 {
		t.Fatalf("统计 = %+v", m.Stats())
	}
}

func TestHighlightModeFollowsInputDevice(t *testing.T) {
	root := NewFocusNode(false)
	a, b := fn("a"), fn("b")
	root.AddChild(a).AddChild(b)
	m := NewFocusManager(root)

	m.NotePointer()
	if m.Mode() != HighlightPointer {
		t.Fatal("指针输入后应为 Pointer 模式")
	}
	m.MoveNext(false) // 键盘 Tab
	if m.Mode() != HighlightKeyboard {
		t.Fatal("Tab 后应切到 Keyboard 模式（焦点环要出现）")
	}
	m.NotePointer()
	if m.Mode() != HighlightPointer {
		t.Fatal("鼠标按下后应切回 Pointer 模式（焦点环消失）")
	}
}

func TestHandleKeyTabOnly(t *testing.T) {
	root := NewFocusNode(false)
	root.AddChild(fn("a")).AddChild(fn("b"))
	m := NewFocusManager(root)

	if m.HandleKey(&KeyEvent{Kind: KeyDown, Code: CodeEnter}) {
		t.Fatal("Enter 不该被焦点管理器吃掉")
	}
	if !m.HandleKey(&KeyEvent{Kind: KeyDown, Code: CodeTab}) {
		t.Fatal("Tab 应被处理")
	}
	if m.Current().Tag != "a" {
		t.Fatalf("焦点 = %v", m.Current().Tag)
	}
	if !m.HandleKey(&KeyEvent{Kind: KeyDown, Code: CodeTab, Mods: ModShift}) {
		t.Fatal("Shift+Tab 应被处理")
	}
	if m.Current().Tag != "b" {
		t.Fatalf("Shift+Tab 后 = %v", m.Current().Tag)
	}
	if m.HandleKey(&KeyEvent{Kind: KeyUp, Code: CodeTab}) {
		t.Fatal("抬起不该再次触发迁移")
	}
}

func TestFocusSyncsDispatcherForKeys(t *testing.T) {
	// 焦点节点的 Key 回调必须真的收到事件（这是"焦点"存在的意义）。
	rootScene := scene.Stack(
		scene.Box(scene.BoxSpec{Fill: render.Fill(geom.RGBA(200, 200, 200))}).SetSize2(40, 40),
		scene.Box(scene.BoxSpec{Fill: render.Fill(geom.RGBA(180, 180, 180))}).SetSize2(40, 40),
	).SetSize2(100, 40)
	rootScene.Layout(scene.Unbounded())
	kids := rootScene.Children()

	d := NewDispatcher(rootScene)
	var got []KeyCode
	for _, k := range kids {
		node := k
		d.Set(node, Handler{Key: func(e *KeyEvent) bool {
			got = append(got, e.Code)
			return true
		}})
	}

	fnRoot := NewFocusNode(false)
	f1, f2 := NewFocusNode(true), NewFocusNode(true)
	f1.Node, f2.Node = kids[0], kids[1]
	fnRoot.AddChild(f1).AddChild(f2)

	m := NewFocusManager(fnRoot)
	m.Attach(d)
	m.Focus(f2)

	if d.Focus() != kids[1] {
		t.Fatal("分发器焦点未同步")
	}
	if !d.Key(&KeyEvent{Kind: KeyDown, Code: CodeEnter}) {
		t.Fatal("Key 未投递")
	}
	if len(got) != 1 || got[0] != CodeEnter {
		t.Fatalf("收到 = %v", got)
	}

	m.Blur()
	if d.Focus() != nil {
		t.Fatal("失焦后分发器仍持有焦点节点")
	}
}

func TestFocusRejectsUnfocusable(t *testing.T) {
	root := NewFocusNode(false)
	n := NewFocusNode(false)
	root.AddChild(n)
	m := NewFocusManager(root)
	if m.Focus(n) {
		t.Fatal("不可聚焦节点不该被聚焦")
	}
	if m.Current() != nil {
		t.Fatal("焦点被错误设置")
	}
}
