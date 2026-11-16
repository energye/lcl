package diag

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── --overlaytest ───────────────────────────────────────────────────
//
// 浮层体系（§A5）+ 组件基座（§A4/A7）的门禁。
//
// 为什么浮层要单独一条命令：它的失败形态是**交互上不可见**的——
// 下拉框贴底边跑出窗口、模态遮罩没吞掉点击、Esc 关掉了底下那层、
// 锚点行被删了浮层还悬着。这些在截图里都正常，只有真的点一遍才发现。
// 所以这里按"用户动作序列"断言，而不是按函数逐个测。

// OverlayCase 是单条浮层/组件断言的结果。
type OverlayCase struct {
	Name string
	OK   bool
	Got  string
	Note string
}

// OverlayResult 是 --overlaytest 的汇总。
type OverlayResult struct {
	Cases   int
	Failed  int
	Elapsed time.Duration
	Details []OverlayCase
}

// OK 判定是否全绿（门禁 ≥14 组）。
func (r OverlayResult) OK() bool { return r.Failed == 0 && r.Cases >= 14 }

// String 返回可读摘要。
func (r OverlayResult) String() string {
	return fmt.Sprintf("overlaytest: %d 组，失败 %d，用时 %v", r.Cases, r.Failed, r.Elapsed)
}

// obase 是用例基准时刻（固定时钟，避免依赖真实时间）。
var obase = time.Date(2026, 11, 16, 10, 0, 0, 0, time.UTC)

// oViewport 是 800×600 的窗口客户区。
func oViewport() geom.Rect { return geom.RectXYWH(0, 0, 800, 600) }

// oDown 造按下事件。
func oDown(x, y float32) *input.PointerEvent {
	p := geom.Pt(x, y)
	return &input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer,
		Device: input.DeviceMouse, Pos: p, PosRoot: p, Time: obase}
}

// oEsc 造 Esc 按下事件。
func oEsc() *input.KeyEvent {
	return &input.KeyEvent{Kind: input.KeyDown, Code: input.CodeEscape, Time: obase}
}

// oContent 造一个固定尺寸的浮层内容。
func oContent(w, h float32) func() *scene.Node {
	return func() *scene.Node {
		n := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(0xFF, 0xFF, 0xFF))})
		n.SetSize2(w, h)
		return n
	}
}

// oclock 是可推进的固定时钟。
//
// 为什么不用恒定时钟：Timeout 的"恢复计时"读的是 Now()，时钟不动的话
// 恢复时刻与原弹出时刻相同，测出来的永远是"立刻消失"。
type oclock struct{ t time.Time }

// newOClock 以基准后 ms 毫秒为起点。
func newOClock(ms int) *oclock {
	return &oclock{t: obase.Add(time.Duration(ms) * time.Millisecond)}
}

// Now 返回当前时刻。
func (c *oclock) Now() time.Time { return c.t }

// Advance 前进 ms 毫秒并返回自身。
func (c *oclock) Advance(ms int) *oclock {
	c.t = c.t.Add(time.Duration(ms) * time.Millisecond)
	return c
}

// oHost 造一个带固定视口与固定时钟的宿主。
func oHost(ms int) *overlay.Host { return oHostAt(newOClock(ms)) }

// oHostAt 造一个用指定时钟的宿主。
func oHostAt(c *oclock) *overlay.Host {
	return overlay.New(overlay.Options{
		Viewport: func() geom.Rect { return oViewport() },
		Now:      c.Now,
	})
}

// OverlayTest 跑浮层与组件基座门禁。
func OverlayTest() (OverlayResult, error) {
	began := time.Now()
	var res OverlayResult
	cases := []struct {
		name string
		run  func() (string, error)
	}{
		// 定位：首选 → 翻转 → 挤压
		{"place-below-anchor", placeBelowAnchor},
		{"place-flips-above-at-bottom-edge", placeFlipsAbove},
		{"place-shifts-inside-viewport", placeShiftsInside},
		{"place-opposite-is-symmetric", placeOppositeSymmetric},
		// 关闭策略
		{"outside-click-dismisses", outsideClickDismisses},
		{"outside-click-opt-out", outsideClickOptOut},
		{"click-inside-does-not-dismiss", clickInsideKeeps},
		{"esc-closes-topmost-only", escClosesTopmost},
		{"timeout-dismisses-toast", timeoutDismisses},
		{"paused-timeout-survives", pausedTimeoutSurvives},
		{"anchor-gone-closes", anchorGoneCloses},
		// 遮罩与 z 序
		{"modal-mask-absorbs-click", modalMaskAbsorbs},
		{"z-order-by-band", zOrderByBand},
		// 组件基座
		{"ui-style-follows-status", uiStyleFollowsStatus},
		{"ui-anim-wallclock-not-frames", uiAnimWallClock},
		{"ui-keyed-list-reuses-nodes", uiKeyedReuses},
	}
	for _, c := range cases {
		got, err := c.run()
		d := OverlayCase{Name: c.name, OK: err == nil, Got: got}
		if err != nil {
			d.Note = err.Error()
			res.Failed++
		}
		res.Cases++
		res.Details = append(res.Details, d)
	}
	res.Elapsed = time.Since(began)
	return res, nil
}

// ── 定位 ────────────────────────────────────────────────────────────

// placeBelowAnchor：下拉框默认贴在锚点下方、左对齐，并带 2px 间隙。
func placeBelowAnchor() (string, error) {
	anchor := geom.RectXYWH(100, 100, 80, 24)
	p := overlay.Resolve(overlay.PlacementBottomStart, anchor,
		geom.Size{W: 160, H: 120}, geom.Pt(0, 2), oViewport())
	got := fmt.Sprintf("rect=%v flipped=%v shifted=%v", p.Rect, p.Flipped, p.Shifted)
	if p.Rect.Min != geom.Pt(100, 126) {
		return got, fmt.Errorf("应在锚点下方间隙 2px")
	}
	if p.Flipped || p.Shifted {
		return got, fmt.Errorf("空间充足时不该翻转或位移")
	}
	return got, nil
}

// placeFlipsAbove：锚点贴着底边 —— 内容翻到上方，且箭头要跟着换边。
func placeFlipsAbove() (string, error) {
	anchor := geom.RectXYWH(100, 560, 80, 24) // 560+24=584，下方只剩 16px
	p := overlay.Resolve(overlay.PlacementBottomStart, anchor,
		geom.Size{W: 160, H: 120}, geom.Pt(0, 2), oViewport())
	got := fmt.Sprintf("rect=%v placement=%v flipped=%v", p.Rect, p.Placement, p.Flipped)
	if !p.Flipped || p.Placement != overlay.PlacementTopStart {
		return got, fmt.Errorf("应翻转到上方")
	}
	if p.Rect.Max.Y > anchor.Min.Y {
		return got, fmt.Errorf("翻转后底边压住了锚点")
	}
	return got, nil
}

// placeShiftsInside：两侧都放不下（内容比可用空间高）—— 挤压进视口而不是反复翻转。
func placeShiftsInside() (string, error) {
	anchor := geom.RectXYWH(700, 300, 80, 24)
	p := overlay.Resolve(overlay.PlacementBottomEnd, anchor,
		geom.Size{W: 300, H: 120}, geom.Pt(0, 2), oViewport())
	got := fmt.Sprintf("rect=%v placement=%v flipped=%v shifted=%v",
		p.Rect, p.Placement, p.Flipped, p.Shifted)
	if p.Shifted && p.Flipped {
		return got, fmt.Errorf("挤压成功时不该同时翻转")
	}
	if p.Rect.Max.X > oViewport().Max.X {
		return got, fmt.Errorf("右边界越界 %.1f", p.Rect.Max.X)
	}
	return got, nil
}

// placeOppositeSymmetric：Opposite 必须成对——翻转一次能回来，否则"翻到对侧"
// 会在边界抖动（A→B→A 每帧跳）。
func placeOppositeSymmetric() (string, error) {
	all := []overlay.Placement{
		overlay.PlacementBottomStart, overlay.PlacementBottom, overlay.PlacementBottomEnd,
		overlay.PlacementTopStart, overlay.PlacementTop, overlay.PlacementTopEnd,
		overlay.PlacementRightStart, overlay.PlacementLeftStart,
	}
	for _, p := range all {
		o := p.Opposite()
		if o == p {
			return fmt.Sprintf("%v", p), fmt.Errorf("%v 没有对侧", p)
		}
		if o.Opposite() != p {
			return fmt.Sprintf("%v→%v→%v", p, o, o.Opposite()), fmt.Errorf("对侧不闭合")
		}
	}
	// Center / Point 没有对侧：必须返回自身（调用方据此跳过翻转）。
	for _, p := range []overlay.Placement{overlay.PlacementCenter, overlay.PlacementPoint} {
		if p.Opposite() != p {
			return fmt.Sprintf("%v→%v", p, p.Opposite()), fmt.Errorf("%v 不该有对侧", p)
		}
	}
	return fmt.Sprintf("%d 组闭合", len(all)), nil
}

// ── 关闭策略 ────────────────────────────────────────────────────────

// outsideClickDismisses：点击浮层外 → 关闭。
func outsideClickDismisses() (string, error) {
	h := oHost(0)
	id := h.Show(overlay.Entry{
		Band: theme.BandPopup, Anchor: overlay.RectAnchor(geom.RectXYWH(100, 100, 80, 24)),
		Placement: overlay.PlacementBottomStart, Size: geom.Size{W: 160, H: 120},
		Dismiss: overlay.DismissOutsideClick, Build: oContent(160, 120), Tag: "menu",
	})
	h.Relayout()
	inside := h.Entries()[0].Placed().Rect.Min
	pass := h.HandlePointer(oDown(10, 10), geom.Pt(10, 10))
	got := fmt.Sprintf("id=%d 内容左上=%v 透传=%v 剩余=%d", id, inside, pass, h.Count())
	if h.Count() != 0 {
		return got, fmt.Errorf("点击外部应关闭")
	}
	if pass {
		return got, fmt.Errorf("关闭浮层的那次点击不该继续下传（会误触下面的按钮）")
	}
	return got, nil
}

// outsideClickOptOut：没声明 OutsideClick 的浮层不因外部点击关闭。
func outsideClickOptOut() (string, error) {
	h := oHost(0)
	h.Show(overlay.Entry{
		Band: theme.BandPopup, Anchor: overlay.RectAnchor(geom.RectXYWH(100, 100, 80, 24)),
		Placement: overlay.PlacementBottomStart, Size: geom.Size{W: 160, H: 120},
		Dismiss: overlay.DismissEsc, Build: oContent(160, 120), Tag: "pinned",
	})
	h.Relayout()
	h.HandlePointer(oDown(10, 10), geom.Pt(10, 10))
	got := fmt.Sprintf("剩余=%d", h.Count())
	if h.Count() != 1 {
		return got, fmt.Errorf("未声明 OutsideClick 却被关闭")
	}
	return got, nil
}

// clickInsideKeeps：点在浮层内部 → 不关闭，且不透传给下层。
func clickInsideKeeps() (string, error) {
	h := oHost(0)
	h.Show(overlay.Entry{
		Band: theme.BandPopup, Anchor: overlay.RectAnchor(geom.RectXYWH(100, 100, 80, 24)),
		Placement: overlay.PlacementBottomStart, Size: geom.Size{W: 160, H: 120},
		Dismiss: overlay.DismissOutsideClick, Build: oContent(160, 120), Tag: "menu",
	})
	h.Relayout()
	r := h.Entries()[0].Placed().Rect
	pt := geom.Pt(r.Min.X+r.W()/2, r.Min.Y+r.H()/2)
	pass := h.HandlePointer(oDown(pt.X, pt.Y), pt)
	got := fmt.Sprintf("点=%v 透传=%v 剩余=%d", pt, pass, h.Count())
	if h.Count() != 1 {
		return got, fmt.Errorf("点在浮层内部却关闭了")
	}
	if pass {
		return got, fmt.Errorf("浮层内的点击不该下传到内容树")
	}
	return got, nil
}

// escClosesTopmost：多层都能被 Esc 关闭时，只关最上层（再按一次关下一层）。
func escClosesTopmost() (string, error) {
	h := oHost(0)
	h.Show(overlay.Entry{
		Band: theme.BandPopup, Anchor: overlay.RectAnchor(geom.RectXYWH(10, 10, 40, 20)),
		Placement: overlay.PlacementBottomStart, Size: geom.Size{W: 100, H: 60},
		Dismiss: overlay.DismissEsc, Build: oContent(100, 60), Tag: "low",
	})
	h.Show(overlay.Entry{
		Band: theme.BandModal, Anchor: overlay.RectAnchor(geom.RectXYWH(10, 10, 40, 20)),
		Placement: overlay.PlacementCenter, Size: geom.Size{W: 200, H: 120},
		Dismiss: overlay.DismissEsc, Build: oContent(200, 120), Tag: "high",
	})
	h.Relayout()
	first := h.HandleKey(oEsc())
	afterFirst := h.Count()
	h2 := h.HandleKey(oEsc())
	got := fmt.Sprintf("第一次=%v 剩 %d，第二次=%v 剩 %d", first, afterFirst, h2, h.Count())
	if !first || afterFirst != 1 {
		return got, fmt.Errorf("第一次 Esc 应只关最上层")
	}
	if !h2 || h.Count() != 0 {
		return got, fmt.Errorf("第二次 Esc 应关掉剩余那层")
	}
	return got, nil
}

// timeoutDismisses：Toast 到时自动消失（时间是推进的，不是靠帧数）。
func timeoutDismisses() (string, error) {
	h := oHost(0)
	h.Show(overlay.Entry{
		Band: theme.BandToast, Anchor: overlay.PointAnchor(geom.Pt(400, 560)),
		Placement: overlay.PlacementPoint, Size: geom.Size{W: 240, H: 40},
		Dismiss: overlay.DismissTimeout, Duration: 2 * time.Second,
		Build: oContent(240, 40), Tag: "toast",
	})
	h.Relayout()
	early := h.Tick(obase.Add(1500 * time.Millisecond))
	mid := h.Count()
	late := h.Tick(obase.Add(2100 * time.Millisecond))
	got := fmt.Sprintf("1.5s=%d(剩%d) 2.1s=%d(剩%d)", early, mid, late, h.Count())
	if early != 0 || mid != 1 {
		return got, fmt.Errorf("未到时长不该消失")
	}
	if late != 1 || h.Count() != 0 {
		return got, fmt.Errorf("超过时长应消失")
	}
	return got, nil
}

// pausedTimeoutSurvives：hover 暂停期间不消失；恢复后重新计时（不是立刻消失）。
func pausedTimeoutSurvives() (string, error) {
	clk := newOClock(0)
	h := oHostAt(clk)
	id := h.Show(overlay.Entry{
		Band: theme.BandToast, Anchor: overlay.PointAnchor(geom.Pt(400, 560)),
		Placement: overlay.PlacementPoint, Size: geom.Size{W: 240, H: 40},
		Dismiss: overlay.DismissTimeout, Duration: time.Second,
		Build: oContent(240, 40), Tag: "toast",
	})
	h.Relayout()
	h.SetPaused(id, true)
	clk.Advance(5000) // 指针悬停了 5 秒
	duringPause := h.Tick(clk.Now())
	alive := h.Count()
	// 恢复：从此刻（基准 + 5s）重新计时。
	h.SetPaused(id, false)
	clk.Advance(500)
	half := h.Tick(clk.Now())
	clk.Advance(700)
	done := h.Tick(clk.Now())
	got := fmt.Sprintf("暂停 5s 后关闭=%d 剩=%d / 恢复+0.5s=%d / 恢复+1.2s=%d 剩=%d",
		duringPause, alive, half, done, h.Count())
	if duringPause != 0 || alive != 1 {
		return got, fmt.Errorf("暂停期间不该消失")
	}
	if half != 0 {
		return got, fmt.Errorf("恢复后应重新计时，而不是按原始弹出时刻算")
	}
	if done != 1 || h.Count() != 0 {
		return got, fmt.Errorf("恢复后满时长应消失")
	}
	return got, nil
}

// anchorGoneCloses：锚点所在行被移除（不可见）后，浮层必须跟着关掉。
func anchorGoneCloses() (string, error) {
	h := oHost(0)
	row := scene.Box(scene.BoxSpec{}).SetSize2(200, 24)
	root := scene.Stack().SetSize2(800, 600)
	root.Add(row)
	h.Show(overlay.Entry{
		Band: theme.BandPopup, Anchor: overlay.NodeAnchor(row),
		Placement: overlay.PlacementBottomStart, Size: geom.Size{W: 160, H: 120},
		Dismiss: overlay.DismissAnchorGone | overlay.DismissOutsideClick,
		Build:   oContent(160, 120), Tag: "popover",
	})
	h.Relayout()
	alive := h.Count()
	row.SetVisible(false)
	closed := h.Tick(obase.Add(16 * time.Millisecond))
	got := fmt.Sprintf("可见时剩余=%d，隐藏后关闭=%d 剩=%d", alive, closed, h.Count())
	if alive != 1 {
		return got, fmt.Errorf("弹出失败")
	}
	if closed != 1 || h.Count() != 0 {
		return got, fmt.Errorf("锚点消失后浮层未关闭")
	}
	return got, nil
}

// ── 遮罩与 z 序 ─────────────────────────────────────────────────────

// modalMaskAbsorbs：模态遮罩吞掉点击——既不透传到内容树，也不关掉对话框。
func modalMaskAbsorbs() (string, error) {
	h := oHost(0)
	h.Show(overlay.Entry{
		Band: theme.BandModal, Placement: overlay.PlacementCenter,
		Size: geom.Size{W: 400, H: 240}, Modal: true, TrapFocus: true,
		Dismiss: overlay.DismissEsc, Build: oContent(400, 240), Tag: "dialog",
	})
	h.Relayout()
	e := h.Entries()[0]
	if e.Mask() == nil {
		return "mask=nil", fmt.Errorf("模态浮层缺遮罩")
	}
	// 点遮罩角落（对话框外）。
	pass := h.HandlePointer(oDown(5, 5), geom.Pt(5, 5))
	got := fmt.Sprintf("mask=%v 透传=%v 剩余=%d", e.Mask().HitSpec().Absorb, pass, h.Count())
	if pass {
		return got, fmt.Errorf("遮罩未吞掉点击（会误触下面的内容）")
	}
	if h.Count() != 1 {
		return got, fmt.Errorf("未声明 OutsideClick 的对话框不该被遮罩点击关掉")
	}
	return got, nil
}

// zOrderByBand：Toast 晚弹出但要画在 Modal 之上——z 由层带决定，不由调用顺序。
func zOrderByBand() (string, error) {
	h := oHost(0)
	anchor := overlay.RectAnchor(geom.RectXYWH(100, 100, 40, 20))
	// 先弹 Toast（低创建序），再弹 Modal（高创建序），最后弹 Popup。
	h.Show(overlay.Entry{Band: theme.BandToast, Anchor: anchor, Placement: overlay.PlacementBottomStart,
		Size: geom.Size{W: 100, H: 40}, Build: oContent(100, 40), Tag: "toast"})
	h.Show(overlay.Entry{Band: theme.BandModal, Anchor: anchor, Placement: overlay.PlacementCenter,
		Size: geom.Size{W: 200, H: 120}, Build: oContent(200, 120), Tag: "modal"})
	h.Show(overlay.Entry{Band: theme.BandPopup, Anchor: anchor, Placement: overlay.PlacementBottomStart,
		Size: geom.Size{W: 120, H: 80}, Build: oContent(120, 80), Tag: "popup"})
	h.Relayout()

	var order []string
	for _, e := range h.Entries() {
		order = append(order, e.TagOr("?"))
	}
	got := fmt.Sprintf("%v", order)
	want := []string{"popup", "modal", "toast"}
	for i := range want {
		if i >= len(order) || order[i] != want[i] {
			return got, fmt.Errorf("绘制顺序应为 popup→modal→toast（按层带，不按弹出顺序）")
		}
	}
	return got, nil
}

// ── 组件基座 ────────────────────────────────────────────────────────

// uiStyleFollowsStatus：样式由 (Variant, Status) 解析——禁用优先于 hover。
func uiStyleFollowsStatus() (string, error) {
	th := theme.Default()
	solid := ui.ResolveStyle(th, ui.VariantSolid, ui.StatusNormal)
	hover := ui.ResolveStyle(th, ui.VariantSolid, ui.StatusHover)
	disabled := ui.ResolveStyle(th, ui.VariantSolid, ui.StatusDisabled)
	got := fmt.Sprintf("normal=%v hover=%v disabled(op=%v fg=%v)",
		solid.Background, hover.Background, disabled.Opacity, disabled.Foreground)
	if hover.Background == solid.Background {
		return got, fmt.Errorf("hover 与常态同色")
	}
	if disabled.Opacity >= 1 {
		return got, fmt.Errorf("禁用未降不透明度")
	}
	// 位集折叠：禁用 + hover 必须显示禁用。
	states := ui.StateDisabled | ui.StateHover
	if states.Status() != ui.StatusDisabled {
		return got, fmt.Errorf("禁用的悬停项显示成了 hover——用户会以为能点")
	}
	return got, nil
}

// uiAnimWallClock：动画按墙钟推进——掉帧不掉速，且 0 时长也要走完状态机。
func uiAnimWallClock() (string, error) {
	lin := func(p float32) float32 { return p }
	c := ui.NewController(100*time.Millisecond, lin)
	c.Forward(0)
	mid := c.Step(50 * time.Millisecond)
	half := c.Value()
	// 模拟掉帧：直接跳到 150ms。
	more := c.Step(150 * time.Millisecond)
	got := fmt.Sprintf("50ms 值=%.2f 仍在动=%v，150ms 值=%.2f 仍在动=%v 状态=%v",
		half, mid, c.Value(), more, c.Status())
	if half < 0.49 || half > 0.51 {
		return got, fmt.Errorf("半程值应为 0.5")
	}
	if more || c.Value() != 1 || c.Status() != ui.AnimCompleted {
		return got, fmt.Errorf("掉帧后应直接落到终值并结束")
	}
	// "减少动效"：时长归零，值立刻到位但状态必须是完成。
	z := ui.NewController(0, lin)
	z.Forward(0)
	if z.Value() != 1 || z.Status() != ui.AnimCompleted {
		return got, fmt.Errorf("0 时长未走到 AnimCompleted")
	}
	return got, nil
}

// uiKeyedReuses：KeyedList 插入到中间时，同 Key 节点必须复用（否则输入框丢焦点/丢输入）。
func uiKeyedReuses() (string, error) {
	mk := func() *scene.Node { return scene.Box(scene.BoxSpec{}).SetSize2(10, 10) }
	old := []*scene.Node{mk(), mk(), mk()}
	keys := []ui.Key{"a", "b", "c"}
	items := []ui.Key{"a", "x", "b", "c"}
	built := 0
	nodes, _, d := ui.KeyedList(items, func(k ui.Key) ui.Key { return k }, old, keys,
		func(k ui.Key, i int, reuse *scene.Node) *scene.Node {
			if reuse != nil {
				return reuse
			}
			built++
			return mk()
		})
	got := fmt.Sprintf("%s 新建=%d", d, built)
	if built != 1 {
		return got, fmt.Errorf("应只新建 1 个节点")
	}
	if nodes[0] != old[0] || nodes[2] != old[1] || nodes[3] != old[2] {
		return got, fmt.Errorf("同 Key 节点未复用")
	}
	return got, nil
}
