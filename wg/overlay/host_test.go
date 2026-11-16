package overlay

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
)

// vp 是 400×300 的视口。
func vp() geom.Rect { return geom.RectXYWH(0, 0, 400, 300) }

// near 判定矩形四角近似相等。
func near(a, b geom.Rect) bool {
	const eps = 0.5
	return abs(a.Min.X-b.Min.X) < eps && abs(a.Min.Y-b.Min.Y) < eps &&
		abs(a.Max.X-b.Max.X) < eps && abs(a.Max.Y-b.Max.Y) < eps
}

func abs(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}

// ── 定位 ────────────────────────────────────────────────────────────

func TestPlaceBottomStart(t *testing.T) {
	anchor := geom.RectXYWH(100, 50, 80, 20)
	got := Place(PlacementBottomStart, anchor, geom.Sz(120, 40), geom.Point{})
	want := geom.RectXYWH(100, 70, 120, 40)
	if !near(got, want) {
		t.Fatalf("得到 %v 期望 %v", got, want)
	}
}

func TestPlaceCenter(t *testing.T) {
	anchor := geom.RectXYWH(100, 50, 80, 20) // 中心 (140,60)
	got := Place(PlacementCenter, anchor, geom.Sz(60, 30), geom.Point{})
	want := geom.RectXYWH(110, 45, 60, 30)
	if !near(got, want) {
		t.Fatalf("得到 %v 期望 %v", got, want)
	}
}

func TestResolveFlipsAtBottomEdge(t *testing.T) {
	// 锚点贴近底边：向下放不下，应翻到上方。
	anchor := geom.RectXYWH(100, 270, 80, 20)
	p := Resolve(PlacementBottomStart, anchor, geom.Sz(120, 40), geom.Point{}, vp())
	if !p.Flipped || p.Placement != PlacementTopStart {
		t.Fatalf("未翻转: %+v", p)
	}
	if p.Rect.Max.Y > anchor.Min.Y+0.5 {
		t.Fatalf("翻转后仍压在锚点上: %v", p.Rect)
	}
}

func TestResolveShiftsWhenBothSidesFail(t *testing.T) {
	// 视口只有 300 高、内容 240 高：锚点居中时两侧都放不下 → 挤压进视口。
	anchor := geom.RectXYWH(100, 140, 80, 20)
	p := Resolve(PlacementBottomStart, anchor, geom.Sz(120, 240), geom.Point{}, vp())
	if !p.Shifted {
		t.Fatalf("未挤压: %+v", p)
	}
	if p.Rect.Min.Y < -0.5 || p.Rect.Max.Y > 300.5 {
		t.Fatalf("挤压后仍越界: %v", p.Rect)
	}
}

func TestResolveClampsHorizontally(t *testing.T) {
	// 右对齐锚点在窗口右侧收缩时会跑出右边 → 必须左移进视口。
	anchor := geom.RectXYWH(360, 100, 40, 20)
	p := Resolve(PlacementBottomStart, anchor, geom.Sz(120, 40), geom.Point{}, vp())
	if p.Rect.Max.X > 400.5 {
		t.Fatalf("右侧越界: %v", p.Rect)
	}
}

func TestResolveNoViewportKeepsPreferred(t *testing.T) {
	anchor := geom.RectXYWH(100, 270, 80, 20)
	p := Resolve(PlacementBottomStart, anchor, geom.Sz(120, 40), geom.Point{}, geom.ZeroRect)
	if p.Flipped || p.Shifted {
		t.Fatalf("无边界时不该修正: %+v", p)
	}
}

// ── 宿主 ────────────────────────────────────────────────────────────

// mkNode 造一个内容节点。
func mkNode(w, h float32) *scene.Node {
	return scene.Box(scene.BoxSpec{Fill: render.Fill(geom.RGBA(30, 30, 30))}).SetSize2(w, h)
}

// newHost 造宿主（视口 400×300）。
func newHost() *Host {
	return New(Options{Viewport: vp})
}

func TestShowAddsAndCloseRemoves(t *testing.T) {
	h := newHost()
	id := h.Show(Entry{
		Band: theme.BandPopup, Anchor: RectAnchor(geom.RectXYWH(10, 10, 40, 20)),
		Placement: PlacementBottomStart, Size: geom.Sz(100, 60),
		Build: func() *scene.Node { return mkNode(100, 60) },
		Tag:   "menu",
	})
	if id == 0 || h.Count() != 1 {
		t.Fatalf("Show 失败: id=%d count=%d", id, h.Count())
	}
	// 根节点要挂上内容，且位置已按定位算好。
	if len(h.Root().Children()) != 1 {
		t.Fatalf("根节点子数 = %d", len(h.Root().Children()))
	}
	placed := h.Entries()[0].Placed().Rect
	if placed.Min.Y < 29 || placed.Min.Y > 31 {
		t.Fatalf("浮层位置 y = %v，期望 30", placed.Min.Y)
	}
	// 布局之后节点的实际偏移必须与定位一致（否则画在一处、命中在另一处）。
	h.Root().Layout(scene.Unbounded())
	if off := h.Root().Children()[0].Offset(); !near(geom.RectXYWH(off.X, off.Y, 1, 1), geom.RectXYWH(placed.Min.X, placed.Min.Y, 1, 1)) {
		t.Fatalf("节点偏移 %v 与定位 %v 不一致", off, placed.Min)
	}
	if !h.Close(id) {
		t.Fatal("Close 失败")
	}
	if h.Count() != 0 || len(h.Root().Children()) != 0 {
		t.Fatalf("关闭后残留: count=%d children=%d", h.Count(), len(h.Root().Children()))
	}
}

func TestBandOrdering(t *testing.T) {
	h := newHost()
	tooltip := h.Show(Entry{Band: theme.BandTooltip, Size: geom.Sz(50, 20),
		Build: func() *scene.Node { return mkNode(50, 20) }, Tag: "tip"})
	popup := h.Show(Entry{Band: theme.BandPopup, Size: geom.Sz(50, 20),
		Build: func() *scene.Node { return mkNode(50, 20) }, Tag: "popup"})
	es := h.Entries()
	if len(es) != 2 || es[0].id != popup || es[1].id != tooltip {
		t.Fatalf("层带顺序错误: %v/%v", es[0].Tag, es[1].Tag)
	}
}

func TestOutsideClickCloses(t *testing.T) {
	h := newHost()
	h.Show(Entry{
		Band: theme.BandPopup, Anchor: RectAnchor(geom.RectXYWH(100, 100, 40, 20)),
		Placement: PlacementBottomStart, Size: geom.Sz(100, 60),
		Dismiss: DismissOutsideClick,
		Build:   func() *scene.Node { return mkNode(100, 60) }, Tag: "menu",
	})
	// 点击浮层内部：不关闭，且不透传到下层。
	if h.HandlePointer(&input.PointerEvent{Kind: input.PointerDown}, geom.Pt(120, 130)) {
		t.Fatal("浮层内的点击不该透传")
	}
	if h.Count() != 1 {
		t.Fatal("内部点击误关闭")
	}
	// 点击外部：关闭，且这次点击不透传到下层（弹窗外的点击不该穿透，§B12）。
	if h.HandlePointer(&input.PointerEvent{Kind: input.PointerDown}, geom.Pt(5, 5)) {
		t.Fatal("关闭浮层的点击不该透传")
	}
	if h.Count() != 0 {
		t.Fatal("外部点击未关闭")
	}
}

func TestOutsideClickIgnoresNonDown(t *testing.T) {
	// 移动经过浮层外部不应关闭（否则拖动划过去就没了）。
	h := newHost()
	h.Show(Entry{Band: theme.BandPopup, Size: geom.Sz(50, 20), Dismiss: DismissOutsideClick,
		Build: func() *scene.Node { return mkNode(50, 20) }})
	h.HandlePointer(&input.PointerEvent{Kind: input.PointerHover}, geom.Pt(5, 5))
	if h.Count() != 1 {
		t.Fatal("hover 误关闭")
	}
}

func TestEscClosesTopmost(t *testing.T) {
	h := newHost()
	low := h.Show(Entry{Band: theme.BandPopup, Size: geom.Sz(40, 20), Dismiss: DismissEsc,
		Build: func() *scene.Node { return mkNode(40, 20) }, Tag: "low"})
	high := h.Show(Entry{Band: theme.BandModal, Size: geom.Sz(40, 20), Dismiss: DismissEsc,
		Build: func() *scene.Node { return mkNode(40, 20) }, Tag: "high"})

	if !h.HandleKey(&input.KeyEvent{Kind: input.KeyDown, Code: input.CodeEscape}) {
		t.Fatal("Esc 未处理")
	}
	if h.Count() != 1 {
		t.Fatalf("应只剩一层，实际 %d", h.Count())
	}
	if _, ok := h.entries[high]; ok {
		t.Fatal("被关闭的不是最上层")
	}
	if _, ok := h.entries[low]; !ok {
		t.Fatal("下层被误关")
	}
}

func TestEscKeyUpIgnored(t *testing.T) {
	h := newHost()
	h.Show(Entry{Band: theme.BandPopup, Size: geom.Sz(40, 20), Dismiss: DismissEsc,
		Build: func() *scene.Node { return mkNode(40, 20) }})
	if h.HandleKey(&input.KeyEvent{Kind: input.KeyUp, Code: input.CodeEscape}) {
		t.Fatal("抬起不该触发关闭")
	}
	if h.Count() != 1 {
		t.Fatal("抬起误关闭")
	}
}

func TestTimeoutDismisses(t *testing.T) {
	now := time.Date(2026, 11, 16, 10, 0, 0, 0, time.UTC)
	h := New(Options{Viewport: vp, Now: func() time.Time { return now }})
	h.Show(Entry{Band: theme.BandToast, Size: geom.Sz(200, 40), Dismiss: DismissTimeout,
		Duration: 3 * time.Second, Build: func() *scene.Node { return mkNode(200, 40) }, Tag: "toast"})

	now = now.Add(2 * time.Second)
	if n := h.Tick(now); n != 0 || h.Count() != 1 {
		t.Fatal("未到时长就消失了")
	}
	now = now.Add(1500 * time.Millisecond)
	if n := h.Tick(now); n != 1 || h.Count() != 0 {
		t.Fatal("超时未关闭")
	}
}

func TestTimeoutPausedOnHover(t *testing.T) {
	now := time.Date(2026, 11, 16, 10, 0, 0, 0, time.UTC)
	h := New(Options{Viewport: vp, Now: func() time.Time { return now }})
	id := h.Show(Entry{Band: theme.BandToast, Size: geom.Sz(200, 40), Dismiss: DismissTimeout,
		Duration: time.Second, Build: func() *scene.Node { return mkNode(200, 40) }})

	h.SetPaused(id, true)
	now = now.Add(5 * time.Second)
	if n := h.Tick(now); n != 0 {
		t.Fatal("暂停期间不该消失")
	}
	h.SetPaused(id, false) // 恢复后重新计时
	now = now.Add(1200 * time.Millisecond)
	if n := h.Tick(now); n != 1 {
		t.Fatal("恢复后应继续计时")
	}
}

func TestModalMaskBlocksUnderlying(t *testing.T) {
	h := newHost()
	h.Show(Entry{Band: theme.BandModal, Modal: true, Size: geom.Sz(200, 100),
		Placement: PlacementCenter, Build: func() *scene.Node { return mkNode(200, 100) }, Tag: "dialog"})

	if len(h.Root().Children()) != 2 {
		t.Fatalf("模态应有遮罩+内容两个节点，实际 %d", len(h.Root().Children()))
	}
	// 遮罩必须吞掉点击：返回 false 表示不继续往内容树分发。
	if h.HandlePointer(&input.PointerEvent{Kind: input.PointerDown}, geom.Pt(10, 10)) {
		t.Fatal("模态外的点击穿透了")
	}
	if h.Count() != 1 {
		t.Fatal("模态外的点击把对话框关了")
	}
}

func TestTrapFocusEntersAndRestores(t *testing.T) {
	// 焦点域进入/退出由宿主负责，关闭后焦点应回到打开前的节点。
	root := input.NewFocusNode(false)
	o1 := input.NewFocusNode(true)
	o1.Tag = "page"
	root.AddChild(o1)
	fm := input.NewFocusManager(root)
	fm.Focus(o1)

	h := New(Options{Viewport: vp, Focus: fm})
	id := h.Show(Entry{Band: theme.BandModal, Modal: true, TrapFocus: true,
		Size: geom.Sz(100, 60), Build: func() *scene.Node { return mkNode(100, 60) }, Tag: "dlg"})
	if fm.ActiveScope() == nil {
		t.Fatal("未进入焦点域")
	}
	h.Close(id)
	if fm.ActiveScope() != nil {
		t.Fatal("关闭后未退出焦点域")
	}
	if fm.Current() != o1 {
		t.Fatalf("焦点未还原，当前 = %v", fm.Current())
	}
}

func TestNodeAnchorFollowsLayout(t *testing.T) {
	// 锚点在滚动容器里移动 → Relayout 后浮层跟随。
	// 用 Flex 而不是 Stack：Stack 的绝对定位子节点会被紧约束拉满，
	// 锚点尺寸就变成整块内容的尺寸，测不出"跟随某个小节点"。
	box := mkNode(0, 0)
	box.SetSize2(40, 20)
	content := scene.Flex(scene.Vertical, 0, box, mkNode(40, 300)).SetSize2(40, 320)
	sc := scene.Scroll(scene.Vertical, content)
	sc.SetSize2(40, 40)
	sc.Layout(scene.Unbounded())

	h := newHost()
	h.Show(Entry{Band: theme.BandPopup, Anchor: NodeAnchor(box), Placement: PlacementBottomStart,
		Size: geom.Sz(60, 30), Build: func() *scene.Node { return mkNode(60, 30) }})
	y0 := h.Entries()[0].Placed().Rect.Min.Y

	sc.SetScrollOffset(15)
	h.Relayout()
	y1 := h.Entries()[0].Placed().Rect.Min.Y
	if !(y1 < y0) {
		t.Fatalf("浮层未跟随锚点: %v → %v", y0, y1)
	}
}

func TestCloseAllByBand(t *testing.T) {
	h := newHost()
	h.Show(Entry{Band: theme.BandPopup, Size: geom.Sz(20, 20), Build: func() *scene.Node { return mkNode(20, 20) }})
	h.Show(Entry{Band: theme.BandPopup, Size: geom.Sz(20, 20), Build: func() *scene.Node { return mkNode(20, 20) }})
	h.Show(Entry{Band: theme.BandToast, Size: geom.Sz(20, 20), Build: func() *scene.Node { return mkNode(20, 20) }})

	if n := h.CloseAll(theme.BandPopup); n != 2 {
		t.Fatalf("关闭 %d 个，期望 2", n)
	}
	if h.Count() != 1 {
		t.Fatalf("剩余 %d 个", h.Count())
	}
	if n := h.CloseAll(0); n != 1 {
		t.Fatalf("CloseAll(0) 关 %d 个", n)
	}
}

func TestShowNilBuildIgnored(t *testing.T) {
	h := newHost()
	if id := h.Show(Entry{Band: theme.BandPopup}); id != 0 {
		t.Fatal("空 Build 不该登记")
	}
	if h.Count() != 0 {
		t.Fatal("空 Build 被登记了")
	}
}

func TestAnchorGoneClosesOnTick(t *testing.T) {
	// 锚点所在行被隐藏（列表刷新/行被删）→ 浮层必须跟着消失，
	// 否则它会悬在一块空白上，而且再也关不掉（点外部也点不到它）。
	h := newHost()
	row := mkNode(80, 20)
	h.Show(Entry{Band: theme.BandPopup, Anchor: NodeAnchor(row), Placement: PlacementBottomStart,
		Size: geom.Sz(60, 30), Dismiss: DismissAnchorGone,
		Build: func() *scene.Node { return mkNode(60, 30) }})
	if n := h.Tick(time.Now()); n != 0 {
		t.Fatal("锚点可见时不该关闭")
	}
	row.SetVisible(false)
	if n := h.Tick(time.Now()); n != 1 {
		t.Fatal("锚点不可见后浮层未关闭")
	}
	if h.Count() != 0 {
		t.Fatalf("剩余 %d 个浮层", h.Count())
	}
	if h.Stats().AnchorGone != 1 {
		t.Fatalf("AnchorGone 计数 = %d", h.Stats().AnchorGone)
	}
}

func TestAnchorGoneIgnoresRectAnchor(t *testing.T) {
	// 固定矩形锚点不会自己消失，不该被 Tick 误关。
	h := newHost()
	h.Show(Entry{Band: theme.BandPopup, Anchor: RectAnchor(geom.RectXYWH(10, 10, 40, 20)),
		Placement: PlacementBottomStart, Size: geom.Sz(60, 30), Dismiss: DismissAnchorGone,
		Build: func() *scene.Node { return mkNode(60, 30) }})
	if n := h.Tick(time.Now()); n != 0 {
		t.Fatal("固定矩形锚点被当成已消失")
	}
	if h.Count() != 1 {
		t.Fatal("浮层被误关")
	}
}
