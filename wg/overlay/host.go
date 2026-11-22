package overlay

import (
	"sort"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
)

// ── 浮层宿主 ────────────────────────────────────────────────────────
//
// 浮层不新增 HWND：内容是一棵 scene 子树，宿主把它挂在层带对应的位置，
// 由 Popup 画布绘制。这样弹窗不会脱离父窗口裁剪，也不会出现
// "子窗口挡不住主窗口绘制"的一类经典问题。
//
// 宿主独占三件事（组件不该自己实现）：
//   - **定位**：首选 → 翻转 → 挤压（place.go）；
//   - **z 序**：按 (Band, 创建序) 排，不靠调用方传 z；
//   - **关闭**：OutsideClick/Esc/Timeout/AnchorGone 由引擎判定，
//     组件写自己的"全局点击监听"必然互相打架（§A5.3）。

// Handle 是浮层句柄。
type Handle uint32

// Dismiss 是关闭策略位集。
type Dismiss uint8

const (
	// DismissOutsideClick 点击浮层外部关闭（由引擎统一判定）。
	DismissOutsideClick Dismiss = 1 << iota
	// DismissEsc Esc 关闭（先给最上层，未消费再下传）。
	DismissEsc
	// DismissTimeout 超时关闭（Toast）。
	DismissTimeout
	// DismissAnchorGone 锚点销毁后关闭。
	DismissAnchorGone
)

// Has 判定是否包含某策略。
func (d Dismiss) Has(v Dismiss) bool { return d&v != 0 }

// AnchorKind 是锚点类型。
type AnchorKind uint8

const (
	// AnchorNone 无锚点（居中或纯坐标）。
	AnchorNone AnchorKind = iota
	// AnchorNode 跟随节点（布局变化时重新定位）。
	AnchorNode
	// AnchorRect 固定矩形。
	AnchorRect
	// AnchorPoint 固定点（右键菜单）。
	AnchorPoint
)

// Anchor 是定位锚点。
type Anchor struct {
	Kind AnchorKind
	Node *scene.Node
	Rect geom.Rect
	Pt   geom.Point
}

// NodeAnchor 用节点做锚点。
func NodeAnchor(n *scene.Node) Anchor { return Anchor{Kind: AnchorNode, Node: n} }

// RectAnchor 用矩形做锚点。
func RectAnchor(r geom.Rect) Anchor { return Anchor{Kind: AnchorRect, Rect: r} }

// PointAnchor 用点做锚点。
func PointAnchor(p geom.Point) Anchor { return Anchor{Kind: AnchorPoint, Pt: p} }

// Rect 返回锚点在根空间的矩形。
func (a Anchor) Resolve() geom.Rect {
	switch a.Kind {
	case AnchorNode:
		if a.Node == nil {
			return geom.ZeroRect
		}
		return a.Node.BoundsInRoot()
	case AnchorRect:
		return a.Rect
	case AnchorPoint:
		return geom.RectXYWH(a.Pt.X, a.Pt.Y, 0, 0)
	}
	return geom.ZeroRect
}

// Entry 是一个浮层的描述。
type Entry struct {
	// Band 决定 z 序（theme.BandPopup/Tooltip/Modal/Toast）。
	Band theme.Band
	// Anchor 是定位锚点。
	Anchor Anchor
	// Placement 是首选方位。
	Placement Placement
	// Offset 是相对首选位置的偏移。
	Offset geom.Point
	// Size 是期望尺寸（用于定位；内容实际尺寸可能不同）。
	Size geom.Size
	// Dismiss 是关闭策略。
	Dismiss Dismiss
	// Modal 表示带遮罩并阻断下层。
	Modal bool
	// TrapFocus 表示键鼠焦点限制在浮层内（Modal 通常为 true）。
	TrapFocus bool
	// Duration 是 DismissTimeout 的时长。
	Duration time.Duration
	// Build 构造内容节点。调用方在更高层构造，overlay 不依赖 ui。
	Build func() *scene.Node
	// Tag 是诊断标签。
	Tag string
	// OnClose 在浮层被任何路径关闭后调用一次（组件自关 / Esc / 外部点击 /
	// 超时 / 锚点消失 / 强制 CloseAll）。组件用它复位自己的 handle——
	// 只在自关路径复位的话，Esc 关掉后再点打不开（--audit 实测）。
	OnClose func()

	// mask 是模态遮罩节点（Modal 时创建）。
	mask *scene.Node
	// content 是内容节点。
	content *scene.Node
	// shownAt 是弹出时刻（Timeout 计时起点）。
	shownAt time.Time
	// paused 表示计时暂停（hover 时 Toast 不消失）。
	paused bool
	// scope 是焦点域（TrapFocus 时创建）。
	scope *input.FocusScope
	// placed 是最近一次定位结果。
	placed Placed
	// id 由宿主在 Show 时写入（调用方靠它 Close）。
	id Handle
}

// ID 返回句柄（未登记时为零）。
func (e *Entry) ID() Handle { return e.id }

// Placed 返回最近一次定位结果。
func (e *Entry) Placed() Placed { return e.placed }

// Place 返回内容节点。
func (e *Entry) Content() *scene.Node { return e.content }

// Mask 返回模态遮罩节点（非模态为 nil）。
func (e *Entry) Mask() *scene.Node { return e.mask }

// setID 由宿主填写句柄。
func (e *Entry) setID(h Handle) { e.id = h }

// Options 是宿主参数。
type Options struct {
	// Viewport 返回当前可用区域（根空间）。
	Viewport func() geom.Rect
	// Focus 是焦点管理器（TrapFocus 用）。
	Focus *input.FocusManager
	// OnChange 在浮层集合变化时调用（宿主据此请求重排/重绘）。
	OnChange func()
	// Now 注入时钟（门禁用固定时钟）。
	Now func() time.Time
}

// Host 管理浮层集合。
type Host struct {
	opts    Options
	entries map[Handle]*Entry
	order   []Handle
	next    Handle
	root    *scene.Node
	stats   Stats
}

// Stats 是宿主统计。
type Stats struct {
	Shown      int
	Closed     int
	Dismissed  int
	Escapes    int
	Timeouts   int
	Outside    int
	AnchorGone int
	Flips      int
	Shifts     int
}

// New 创建宿主。
func New(opts Options) *Host {
	h := &Host{opts: opts, entries: make(map[Handle]*Entry, 8), next: 1}
	h.root = scene.Stack().SetSize2(0, 0)
	if opts.Now == nil {
		h.opts.Now = time.Now
	}
	return h
}

// Root 返回浮层根节点（宿主把它挂在内容之上）。
func (h *Host) Root() *scene.Node { return h.root }

// Stats 返回统计。
func (h *Host) Stats() Stats { return h.stats }

// Count 返回当前浮层数。
func (h *Host) Count() int { return len(h.entries) }

// Entries 返回按 (Band, 创建序) 排序的浮层快照（低层带在前 = 先绘）。
//
// 排序而不是只看插入顺序：Toast 可能在 Modal 之前弹出，但必须画在它上面。
// 调用方不需要传 z —— 层带就是唯一的 z 语义。
func (h *Host) Entries() []*Entry {
	out := make([]*Entry, 0, len(h.order))
	for _, id := range h.order {
		if e, ok := h.entries[id]; ok {
			out = append(out, e)
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Band != out[j].Band {
			return out[i].Band < out[j].Band
		}
		return out[i].id < out[j].id // 同层带按创建序，稳定可预测
	})
	return out
}

// Show 弹出一个浮层，返回句柄。
//
// Build 在 Show 内立即调用：浮层内容必须先建出来才能量尺寸定位。
// 若 Build 返回 nil，则不登记（调用方逻辑错误，但不要让引擎崩）。
func (h *Host) Show(e Entry) Handle {
	if e.Build == nil {
		return 0
	}
	content := e.Build()
	if content == nil {
		return 0
	}
	e.content = content
	e.shownAt = h.opts.Now()
	if e.Modal {
		e.mask = h.buildMask()
	}
	h.next++
	id := h.next
	e.setID(id)
	// 焦点域必须在登记**之前**建好：登记的是值副本，
	// 之后再改 e.scope 只会改到栈上的临时变量，关闭时找不到要退出的域。
	if e.TrapFocus && h.opts.Focus != nil {
		e.scope = input.NewFocusScope(e.TagOr("overlay"), focusRoot(&e))
	}
	h.entries[id] = eptr(e)
	h.order = append(h.order, id)
	h.stats.Shown++

	h.rebuild()
	h.relayout()
	if e.scope != nil {
		h.opts.Focus.EnterScope(e.scope)
	}
	h.changed()
	return id
}

// eptr 把值转成指针（Entry 登记后由宿主独占）。
func eptr(e Entry) *Entry {
	p := e
	return &p
}

// buildMask 构造模态遮罩：吞掉所有指针事件（HitSpec.Absorb）。
//
// 遮罩色是**不透明黑**，半透明靠节点 opacity——组透明度由层合成统一处理，
// 颜色里再写一次 alpha 会让遮罩比预期更淡（双重淡化）。
func (h *Host) buildMask() *scene.Node {
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(0, 0, 0))})
	n.SetSize2(0, 0)
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "overlay.mask", Absorb: true})
	n.SetOpacity(0.45)
	return n
}

// TagOr 返回诊断标签（空则用默认名）。
func (e *Entry) TagOr(def string) string {
	if e.Tag != "" {
		return e.Tag
	}
	return def
}

// Close 关闭指定浮层。
func (h *Host) Close(id Handle) bool {
	e, ok := h.entries[id]
	if !ok {
		return false
	}
	h.closeEntry(e)
	return true
}

// CloseAll 关闭某层带的全部浮层（band 为零值表示全部）。
func (h *Host) CloseAll(band theme.Band) int {
	n := 0
	for _, id := range append([]Handle(nil), h.order...) {
		e, ok := h.entries[id]
		if !ok {
			continue
		}
		if band != 0 && e.Band != band {
			continue
		}
		h.closeEntry(e)
		n++
	}
	return n
}

// closeEntry 真正移除并做清理。
func (h *Host) closeEntry(e *Entry) {
	if e.scope != nil && h.opts.Focus != nil {
		// 退出焦点域会还原打开前的焦点（§A5.4）。
		if h.opts.Focus.ActiveScope() == e.scope {
			h.opts.Focus.ExitScope()
		}
		e.scope = nil
	}
	delete(h.entries, e.id)
	for i, id := range h.order {
		if id == e.id {
			h.order = append(h.order[:i], h.order[i+1:]...)
			break
		}
	}
	if e.OnClose != nil {
		e.OnClose()
	}
	h.stats.Closed++
	h.rebuild()
	h.changed()
}

// changed 通知外部。
func (h *Host) changed() {
	if h.opts.OnChange != nil {
		h.opts.OnChange()
	}
}

// rebuild 按 (Band, 创建序) 重建根节点子树。
//
// 每次都重建而不是增量插入：浮层数量是个位数，重建的代价远小于
// "插入位置算错导致弹窗被别的弹窗压在下面"的排查成本。
func (h *Host) rebuild() {
	h.root.RemoveAll()
	for _, e := range h.Entries() {
		if e.mask != nil {
			h.root.Add(e.mask)
		}
		h.root.Add(e.content)
	}
}

// relayout 重新定位全部浮层（viewport/anchor 变化时调用）。
func (h *Host) relayout() {
	vp := geom.ZeroRect
	if h.opts.Viewport != nil {
		vp = h.opts.Viewport()
	}
	h.root.SetSize2(vp.W(), vp.H())
	for _, e := range h.Entries() {
		anchor := e.Anchor.Resolve()
		// AnchorNone 的居中语义以**视口中心**为锚（antd Modal 居中相对
		// 视口）：AnchorNone 解析为零矩形，Place(Center) 会算出负坐标、
		// 再被 Resolve 挤压拉回左上角（实测 modal 贴左上角而非居中）。
		if e.Anchor.Kind == AnchorNone && e.Placement == PlacementCenter && !vp.IsEmpty() {
			anchor = geom.RectXYWH(vp.W()/2, vp.H()/2, 0, 0)
		}
		// 高度 0 = 高度自适应：按内容根的自然高度定位（antd 对话框形态——
		// 高度随内容；固定 220 会截断多内容）。宽度仍用 Entry.Size。
		// content 若还没走过布局（Show 当帧），先按视口宽松约束测一次。
		size := e.Size
		if size.H <= 0 && e.content != nil {
			if e.content.Size().H <= 0 {
				e.content.Layout(scene.LooseConstraints(geom.Sz(e.Size.W, vp.H())))
			}
			if e.content.Size().H > 0 {
				size.H = e.content.Size().H
			}
		}
		p := Resolve(e.Placement, anchor, size, e.Offset, vp)
		if p.Flipped {
			h.stats.Flips++
		}
		if p.Shifted {
			h.stats.Shifts++
		}
		e.placed = p
		if e.mask != nil {
			e.mask.SetSize2(vp.W(), vp.H())
			e.mask.SetStackPosition(geom.Pt(0, 0))
		}
		e.content.SetStackPosition(p.Rect.Min)
		// 浮层由定位决定位置、由内容决定尺寸，不能按 Stack 默认的紧约束拉满。
		e.content.SetPositionKeepSize(true)
		// SetStackPosition 走 MarkNeedsLayout 异步折算 offset——审计在
		// Show 后 300ms 内 locate 时布局可能尚未跑完，浮层内节点的
		// BoundsInRoot 读到旧值/零（实测 select-option 点击坐标打不到
		// 浮层的根因）。SetStackPositionPaint **同步**写 offset（含
		// parent.ContentOrigin），定位当帧坐标即刻正确。
		e.content.SetStackPositionPaint(p.Rect.Min)
		// content 子树的 positioned 节点（modal 关闭钮等 Stack 内绝对
		// 定位）offset 仍依赖一次布局折算：若 Show→rebuild 后恰好有帧
		// 在 relayout 前跑了布局，子节点按 content 旧原点 (0,0) 折算，
		// 之后 MarkNeedsPaint 不触发布局——静窗（审计）下 stale offset
		// 永存（实测 image-close bounds y=10 误判越界）。强制重排整
		// 子树让折算与新原点一致。
		e.content.Relayout(scene.LooseConstraints(geom.Sz(size.W, size.H)))
	}
}

// Relayout 供宿主在 resize/滚动后调用。
func (h *Host) Relayout() {
	h.relayout()
	h.changed()
}

// Tick 推进时间（Timeout / AnchorGone 策略）。
//
// 单一入口而不是每个 Toast 自己开定时器：定时器要额外线程与取消逻辑，
// 而调度器本来就每帧推进时间（与长按识别器同样的理由）。
//
// AnchorGone 也走这里而不是监听节点销毁：scene.Node 没有"析构回调"，
// 挂在上面的监听在节点被摘下的那一刻没人调，浮层就悬空留在原地。
func (h *Host) Tick(now time.Time) int {
	n := 0
	for _, e := range h.Entries() {
		if e.Dismiss.Has(DismissAnchorGone) && h.anchorGone(e) {
			h.stats.AnchorGone++
			h.stats.Dismissed++
			h.closeEntry(e)
			n++
			continue
		}
		if !e.Dismiss.Has(DismissTimeout) || e.Duration <= 0 || e.paused {
			continue
		}
		if now.Sub(e.shownAt) >= e.Duration {
			h.stats.Timeouts++
			h.stats.Dismissed++
			h.closeEntry(e)
			n++
		}
	}
	return n
}

// anchorGone 判定锚点是否已不可用。
//
// 只对 AnchorNode 有意义（固定矩形/点不会自己消失）。判据是"整条祖先链可见"：
// 行被移出列表时组件会 SetVisible(false)，这条链就断了。
// 已知边界：节点被原地摘下但未置不可见时察觉不到——那种情况下浮层会留在
// 旧位置，所以移除锚点前必须先隐藏它（widget 卸载路径统一这么做）。
func (h *Host) anchorGone(e *Entry) bool {
	if e.Anchor.Kind != AnchorNode {
		return false
	}
	if e.Anchor.Node == nil {
		return true
	}
	return !e.Anchor.Node.VisibleInTree()
}

// SetPaused 设置某个浮层的计时暂停（hover 暂停 Toast）。
func (h *Host) SetPaused(id Handle, paused bool) {
	if e, ok := h.entries[id]; ok {
		e.paused = paused
		if !paused {
			// 恢复时从当前时刻重新计时：否则 hover 结束立刻消失。
			e.shownAt = h.opts.Now()
		}
	}
}

// HitEntry 返回命中点所在的浮层（不含遮罩），无命中返回 0。
//
// 用于 OutsideClick 判定：引擎统一算，组件不自己监听全局点击。
func (h *Host) HitEntry(p geom.Point) Handle {
	es := h.Entries()
	for i := len(es) - 1; i >= 0; i-- {
		e := es[i]
		if e.placed.Rect.Contains(p) {
			return e.id
		}
	}
	return 0
}

// ViewportRect 返回当前可用区域（根空间）；未配置视口时返回零矩形。
// 浮层组件（Drawer 等通栏形态）据此定自己的通栏尺寸。
func (h *Host) ViewportRect() geom.Rect {
	if h.opts.Viewport == nil {
		return geom.Rect{}
	}
	return h.opts.Viewport()
}

// HitContent 返回命中点所在浮层的**内容根节点**（事件只分发给该子树，
// 不透传下层内容树），无命中返回 nil。
//
// 容器据此把 PointerEvent 路由进浮层：命中浮层 ≠ 吞掉事件——浮层内的
// 可点项（autocomplete 建议、select 选项等）必须收到点击。
func (h *Host) HitContent(p geom.Point) *scene.Node {
	es := h.Entries()
	for i := len(es) - 1; i >= 0; i-- {
		e := es[i]
		if e.placed.Rect.Contains(p) {
			return e.content
		}
	}
	return nil
}

// HandlePointer 处理指针事件，返回是否应继续发给下层。
//
// 语义：
//   - 命中某个浮层 → 只发给它（返回 false，调用方不要再往内容树分发）；
//   - 命中模态遮罩 → 吞掉（返回 false）；
//   - 什么都没命中且最上层有 OutsideClick → 关闭它（返回 false，点击不透传）。
func (h *Host) HandlePointer(ev *input.PointerEvent, root geom.Point) bool {
	if ev == nil || len(h.order) == 0 {
		return true
	}
	if ev.Kind != input.PointerDown {
		// 非按下事件不参与关闭判定（否则拖动经过浮层外就会关掉）。
		return h.HitEntry(root) == 0
	}
	hit := h.HitEntry(root)
	if hit != 0 {
		// 命中浮层：点击不透传到下层（浮层是"最上面的一层纸"）。
		return false
	}
	// 未命中任何浮层：从最上层往下找第一个接受 OutsideClick 的。
	es := h.Entries()
	for i := len(es) - 1; i >= 0; i-- {
		if es[i].Dismiss.Has(DismissOutsideClick) {
			h.stats.Outside++
			h.stats.Dismissed++
			h.closeEntry(es[i])
			return false
		}
		if es[i].Modal {
			// 模态遮罩挡住一切：点击落在弹窗外也不该穿透。
			return false
		}
	}
	return true
}

// HandleKey 处理键盘事件（Esc），返回是否消费。
func (h *Host) HandleKey(ev *input.KeyEvent) bool {
	if ev == nil || ev.Kind == input.KeyUp || len(h.order) == 0 {
		return false
	}
	if ev.Code != input.CodeEscape {
		return false
	}
	es := h.Entries()
	for i := len(es) - 1; i >= 0; i-- {
		if es[i].Dismiss.Has(DismissEsc) {
			h.stats.Escapes++
			h.stats.Dismissed++
			h.closeEntry(es[i])
			return true
		}
	}
	return false
}

// focusRoot 构造浮层的焦点域根节点。
func focusRoot(e *Entry) *input.FocusNode {
	n := input.NewFocusNode(false)
	n.Tag = e.TagOr("overlay")
	return n
}
