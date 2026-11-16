package ui

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
)

// ── Base ────────────────────────────────────────────────────────────

func TestSetStateMarksDirtyAndInvalidates(t *testing.T) {
	b := &Base{}
	frames := 0
	b.Attach(func() { frames++ })
	if b.Dirty() {
		t.Fatal("新组件不该是脏的")
	}
	b.SetState(func() {})
	if !b.Dirty() {
		t.Fatal("SetState 未标脏")
	}
	if frames != 1 {
		t.Fatalf("请求帧 %d 次", frames)
	}
	b.Rebuild()
	if b.Dirty() {
		t.Fatal("Rebuild 后仍是脏的")
	}
	if b.Builds() != 1 {
		t.Fatalf("重建计数 = %d", b.Builds())
	}
}

func TestSetStateWithoutAttachDoesNotPanic(t *testing.T) {
	b := &Base{}
	b.SetState(func() {})
	if !b.Dirty() {
		t.Fatal("未挂载时也应标脏")
	}
}

func TestBaseDetachStopsInvalidate(t *testing.T) {
	b := &Base{}
	frames := 0
	b.Attach(func() { frames++ })
	b.Detach()
	b.SetState(func() {})
	if frames != 0 {
		t.Fatal("解挂后仍请求帧")
	}
}

func TestBuildContextDefaults(t *testing.T) {
	var c *BuildContext
	if c.ThemeOr() == nil {
		t.Fatal("nil ctx 应返回默认主题")
	}
	if got := c.Motion(200); got != 200 {
		t.Fatalf("Motion = %v", got)
	}
	c = &BuildContext{Media: MediaQuery{ReduceMotion: true}}
	if got := c.Motion(200); got != 0 {
		t.Fatal("减少动效时时长应归零")
	}
}

// ── 样式 ────────────────────────────────────────────────────────────

func TestStateSetStatusPriority(t *testing.T) {
	// 禁用的选中项被 hover：应显示禁用，否则用户以为能点。
	s := StateDisabled | StateSelected | StateHover
	if s.Status() != StatusDisabled {
		t.Fatalf("Status = %v", s.Status())
	}
	if (StateHover | StateSelected).Status() != StatusHover {
		t.Fatal("hover 应优先于 selected")
	}
	if StateSet(0).Status() != StatusNormal {
		t.Fatal("空集应为常态")
	}
}

func TestResolveStyleVariants(t *testing.T) {
	th := theme.Default()
	c := th.Color

	solid := ResolveStyle(th, VariantSolid, StatusNormal)
	if solid.Background != c.Primary || solid.Foreground != c.OnPrimary {
		t.Fatalf("实心配色错误: %+v", solid)
	}
	if solid.Opacity != 1 {
		t.Fatalf("常态不透明度 = %v", solid.Opacity)
	}

	hover := ResolveStyle(th, VariantSolid, StatusHover)
	if hover.Background == solid.Background {
		t.Fatal("hover 背景与常态相同")
	}

	// M8 对齐 antd：禁用不再用透明度降维，改用 disabled 专色
	//（bg=fillQuaternary 等价、文字 rgba(0,0,0,.25) 实色等价）。
	disabled := ResolveStyle(th, VariantSolid, StatusDisabled)
	if disabled.Opacity != 1 {
		t.Fatal("antd 禁用态不降不透明度（用专色）")
	}
	if disabled.Background != c.SurfaceHover || disabled.Foreground != c.BorderStrong {
		t.Fatalf("禁用专色错误: %+v", disabled)
	}

	outline := ResolveStyle(th, VariantOutline, StatusNormal)
	if outline.BorderWidth == 0 {
		t.Fatal("描边变体必须有边框宽度")
	}

	ghost := ResolveStyle(th, VariantGhost, StatusNormal)
	if ghost.BorderWidth != 0 {
		t.Fatal("幽灵变体不该有边框")
	}
}

func TestMetricsForSizes(t *testing.T) {
	th := theme.Default()
	s, m, l := MetricsFor(th, SizeSmall), MetricsFor(th, SizeMedium), MetricsFor(th, SizeLarge)
	if !(s.Height < m.Height && m.Height < l.Height) {
		t.Fatalf("高度不递增: %v %v %v", s.Height, m.Height, l.Height)
	}
	if !(s.Font < m.Font && m.Font < l.Font) {
		t.Fatalf("字号不递增: %v %v %v", s.Font, m.Font, l.Font)
	}
}

func TestStyleOfCombines(t *testing.T) {
	th := theme.Default()
	st := StyleOf(th, VariantOutline, SizeSmall, StateHover|StateFocus)
	if st.Padding.Left() <= 0 {
		t.Fatal("缺少内边距")
	}
	if st.Radius != th.Radius.SM {
		t.Fatalf("小尺寸圆角 = %v", st.Radius)
	}
	// M8 对齐 antd：default 按钮 hover 是文字/边框转主色，底色不变。
	if st.Border != th.Color.Primary || st.Foreground != th.Color.Primary {
		t.Fatalf("outline hover 应转主色: %+v", st)
	}
}

// ── 动画 ────────────────────────────────────────────────────────────

func TestControllerWallClock(t *testing.T) {
	c := NewController(100*time.Millisecond, func(p float32) float32 { return p }) // 线性
	c.Forward(0)
	if !c.Step(50 * time.Millisecond) {
		t.Fatal("半程应仍在动")
	}
	if v := c.Value(); v < 0.49 || v > 0.51 {
		t.Fatalf("50ms 时值 = %v", v)
	}
	// 掉帧：直接跳到 150ms（超出时长）→ 必须落到终值，而不是"按帧计数走一半"。
	if c.Step(150 * time.Millisecond) {
		t.Fatal("超时长应结束")
	}
	if c.Value() != 1 || c.Status() != AnimCompleted {
		t.Fatalf("终值 = %v 状态 = %v", c.Value(), c.Status())
	}
}

func TestControllerZeroDurationStillCompletes(t *testing.T) {
	// "减少动效"把时长归零：值立刻到位，但状态必须是完成——
	// 否则依赖"动画结束"的逻辑永远等不到。
	c := NewController(0, nil)
	c.Forward(0)
	if c.Value() != 1 {
		t.Fatalf("值 = %v", c.Value())
	}
	if c.Status() != AnimCompleted {
		t.Fatalf("状态 = %v", c.Status())
	}
}

func TestControllerReverse(t *testing.T) {
	c := NewController(100*time.Millisecond, nil)
	c.SetValue(1)
	c.Reverse(0)
	c.Step(100 * time.Millisecond)
	if c.Value() != 0 {
		t.Fatalf("反向终值 = %v", c.Value())
	}
}

func TestControllerCurveApplied(t *testing.T) {
	ease := func(p float32) float32 { return p * p } // 慢进
	c := NewController(100*time.Millisecond, ease)
	c.Forward(0)
	c.Step(50 * time.Millisecond)
	if v := c.Value(); v < 0.24 || v > 0.26 {
		t.Fatalf("曲线后值 = %v（线性会是 0.5）", v)
	}
}

func TestTweenAt(t *testing.T) {
	tw := Tween{From: 10, To: 20}
	if got := tw.At(0); got != 10 {
		t.Fatalf("起点 = %v", got)
	}
	if got := tw.At(1); got != 20 {
		t.Fatalf("终点 = %v", got)
	}
	if got := tw.At(0.5); got != 15 {
		t.Fatalf("中点 = %v", got)
	}
}

func TestImplicitDoesNotRestartOnSameTarget(t *testing.T) {
	im := NewImplicit(100*time.Millisecond, nil, 0)
	im.To(1, 0)
	im.Step(50 * time.Millisecond)
	mid := im.Value()
	im.To(1, 50*time.Millisecond) // 目标没变：不该重启
	if im.Controller().Status() == AnimCompleted {
		t.Fatal("目标未变却重启了动画")
	}
	if im.Value() != mid {
		t.Fatalf("值被重置: %v → %v", mid, im.Value())
	}
	im.Step(100 * time.Millisecond)
	if im.Value() != 1 {
		t.Fatalf("终值 = %v", im.Value())
	}
}

// ── KeyedList ───────────────────────────────────────────────────────

func TestDiffKeys(t *testing.T) {
	old := []Key{"a", "b", "c"}
	next := []Key{"a", "c", "d"}
	d := DiffKeys(old, next)
	if d.Reused != 2 || d.Created != 1 || d.Removed != 1 {
		t.Fatalf("%s", d)
	}
}

func TestDiffKeysReorder(t *testing.T) {
	d := DiffKeys([]Key{"a", "b", "c"}, []Key{"c", "a", "b"})
	if d.Reused != 3 || d.Moved == 0 {
		t.Fatalf("重排未识别: %s", d)
	}
}

func TestDiffKeysFromEmpty(t *testing.T) {
	d := DiffKeys(nil, []Key{"a", "b"})
	if d.Created != 2 || d.Reused != 0 {
		t.Fatalf("%s", d)
	}
}

func TestKeyedListReusesByKey(t *testing.T) {
	type item struct {
		ID   Key
		Name string
	}
	items := []item{{"a", "A"}, {"b", "B"}}
	var built []*scene.Node
	nodes, keys, _ := KeyedList(items, func(it item) Key { return it.ID }, nil, nil,
		func(it item, i int, reuse *scene.Node) *scene.Node {
			n := scene.Box(scene.BoxSpec{}).SetSize2(10, 10)
			built = append(built, n)
			return n
		})
	if len(nodes) != 2 || len(keys) != 2 {
		t.Fatalf("产出 %d 节点 %d key", len(nodes), len(keys))
	}

	// 插入一项到中间：首尾节点必须被复用（同指针），只有新项是新的。
	items2 := []item{{"a", "A"}, {"x", "X"}, {"b", "B"}}
	built = nil
	nodes2, keys2, d := KeyedList(items2, func(it item) Key { return it.ID }, nodes, keys,
		func(it item, i int, reuse *scene.Node) *scene.Node {
			if reuse != nil {
				return reuse
			}
			n := scene.Box(scene.BoxSpec{}).SetSize2(10, 10)
			built = append(built, n)
			return n
		})
	if len(built) != 1 {
		t.Fatalf("应只新建 1 个节点，实际 %d", len(built))
	}
	if nodes2[0] != nodes[0] || nodes2[2] != nodes[1] {
		t.Fatal("同 Key 节点未被复用")
	}
	if d.Created != 1 || d.Reused != 2 || d.Removed != 0 {
		t.Fatalf("diff = %s", d)
	}
	if keys2[1] != "x" {
		t.Fatalf("keys = %v", keys2)
	}
}

func TestKeyedListPositionalWhenNoKey(t *testing.T) {
	// 无 Key 时按位置复用：长度变化只在尾部增删。
	old := []*scene.Node{scene.Box(scene.BoxSpec{}), scene.Box(scene.BoxSpec{})}
	items := []int{1, 2, 3}
	nodes, _, _ := KeyedList(items, func(int) Key { return "" }, old, []Key{"", ""},
		func(v int, i int, reuse *scene.Node) *scene.Node {
			if reuse != nil {
				return reuse
			}
			return scene.Box(scene.BoxSpec{})
		})
	if len(nodes) != 3 || nodes[0] != old[0] || nodes[1] != old[1] {
		t.Fatal("按位置复用失败")
	}
}

// ── GestureDetector ─────────────────────────────────────────────────

func TestGestureDetectorTap(t *testing.T) {
	var taps int
	var ended []bool
	g := &GestureDetector{OnTap: func(geom.Point) { taps++ },
		OnTapEnd: func(ok bool) { ended = append(ended, ok) }}
	h := g.Handler()

	h.Pointer(down(10, 10))
	h.Pointer(up(10, 10))
	if taps != 1 {
		t.Fatalf("tap 次数 = %d", taps)
	}
	if len(ended) != 1 || !ended[0] {
		t.Fatalf("结束回调 = %v", ended)
	}
}

func TestGestureDetectorPan(t *testing.T) {
	var starts, moves int
	var lastVel geom.Point
	g := &GestureDetector{
		OnPanStart: func(geom.Point) { starts++ },
		OnPan:      func(pos, d geom.Point) { moves++ },
		OnPanEnd:   func(pos, v geom.Point) { lastVel = v },
	}
	h := g.Handler()
	h.Pointer(down(10, 10))
	h.Pointer(move(30, 10))
	h.Pointer(move(60, 10))
	h.Pointer(up(60, 10))
	if starts != 1 || moves == 0 {
		t.Fatalf("start=%d move=%d", starts, moves)
	}
	if lastVel.Len() == 0 {
		t.Fatal("结束速度为空")
	}
}

func TestGestureDetectorTapLosesToPan(t *testing.T) {
	var taps int
	var tapEnd []bool
	g := &GestureDetector{
		OnTap:    func(geom.Point) { taps++ },
		OnTapEnd: func(ok bool) { tapEnd = append(tapEnd, ok) },
		OnPan:    func(pos, d geom.Point) {},
	}
	h := g.Handler()
	h.Pointer(down(10, 10))
	h.Pointer(move(40, 10)) // 超过 slop → pan 胜
	h.Pointer(up(40, 10))
	if taps != 0 {
		t.Fatal("拖动后不该触发点击")
	}
	if len(tapEnd) != 1 || tapEnd[0] {
		t.Fatalf("点击应收到 accepted=false 以便回滚: %v", tapEnd)
	}
}

func TestGestureDetectorLongPressNeedsTick(t *testing.T) {
	var lp int
	g := &GestureDetector{OnLongPress: func(geom.Point) { lp++ }}
	h := g.Handler()
	h.Pointer(down(10, 10))
	if g.Tick(t0(400)) {
		t.Fatal("400ms 不该判定")
	}
	if !g.Tick(t0(600)) {
		t.Fatal("600ms 应判定")
	}
	if lp != 1 {
		t.Fatalf("长按回调 %d 次", lp)
	}
}

func TestGestureDetectorRegistersOnlyConfigured(t *testing.T) {
	g := &GestureDetector{OnTap: func(geom.Point) {}}
	g.Handler().Pointer(down(1, 1))
	if len(g.members) != 1 {
		t.Fatalf("只配了 tap，成员数 = %d", len(g.members))
	}
	g2 := &GestureDetector{}
	g2.Handler().Pointer(down(1, 1))
	if g2.Active() {
		t.Fatal("什么都没配却有活动手势")
	}
}

// ── 测试辅助 ────────────────────────────────────────────────────────

// t0 是固定时钟（非零，避免被当作"无时间戳"）。
func t0(ms int) time.Time {
	return time.Date(2026, 11, 16, 10, 0, 0, 0, time.UTC).Add(time.Duration(ms) * time.Millisecond)
}

// down / move / up 造指针事件。
func down(x, y float32) *input.PointerEvent {
	p := geom.Pt(x, y)
	return &input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer,
		Device: input.DeviceMouse, Pos: p, PosRoot: p, Time: t0(0)}
}

func move(x, y float32) *input.PointerEvent {
	e := down(x, y)
	e.Kind = input.PointerMove
	e.Time = t0(16)
	return e
}

func up(x, y float32) *input.PointerEvent {
	e := down(x, y)
	e.Kind = input.PointerUp
	e.Time = t0(32)
	return e
}
