package input

import (
	"sync"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
)

// Handler 是节点的事件回调集合。
//
// 用"函数字段的结构体"而不是一个带十个方法的接口：
// 绝大多数节点只关心其中一两个回调，接口会逼着它们写一堆空方法，
// 也让"有没有挂处理器"变成一个需要反射/类型断言的问题。
// 零值 Handler 表示全不关心，等价于没挂。
type Handler struct {
	// Capture 在**捕获阶段**（根 → 叶）调用，返回 true 表示已处理并终止传播。
	Capture func(*PointerEvent) bool
	// Pointer 是**目标阶段**回调。
	Pointer func(*PointerEvent) bool
	// Bubble 在**冒泡阶段**（叶 → 根）调用。
	Bubble func(*PointerEvent) bool
	// Enter/Leave 是 hover 差分回调（不参与传播链）。
	Enter func(*PointerEvent)
	Leave func(*PointerEvent)
	// Wheel 处理滚轮；返回 false 会继续向上冒泡（嵌套滚动就是这样串起来的）。
	Wheel func(*WheelEvent) bool
	// Key 处理键盘（目标是焦点节点，未处理则沿父链冒泡）。
	Key func(*KeyEvent) bool
	// IME 处理输入法组合。
	IME func(*IMEEvent) bool
	// Focus 处理焦点变化。
	Focus func(*FocusEvent) bool
}

// empty 表示没有任何回调。
func (h Handler) empty() bool {
	return h.Capture == nil && h.Pointer == nil && h.Bubble == nil &&
		h.Enter == nil && h.Leave == nil && h.Wheel == nil &&
		h.Key == nil && h.IME == nil && h.Focus == nil
}

// capture 是某个指针的捕获快照。
//
// 存路径而不是只存目标：捕获期间指针可能移出原节点，但冒泡顺序仍应是
// "按下那一刻的路径"，否则拖动到别的节点上会让父级跟着换人响应。
type capture struct {
	target *scene.Node
	path   []*scene.Node
}

// DispStats 是分发统计（门禁与诊断用）。
//
// 计数只在主线程的输入路径上累加（Pointer/Key/IME/Wheel/Cancel 都从
// 主线程进入），诊断方也在主线程读——无锁。
type DispStats struct {
	Pointers      int
	Handled       int
	Enters        int
	Leaves        int
	Captures      int
	Cancels       int
	CursorChanges int
}

// Dispatcher 分发归一化事件到场景树。
//
// 三件事必须由它独占（不能下放给组件）：
//   - **指针捕获**：按下后后续事件不重新命中（滑块拖出轨道仍生效，C26）；
//   - **hover 差分**：Enter/Leave 由路径差分合成，不要求每个组件自己记账；
//   - **Cancel 自愈**：模态循环吞掉 Up 时，由它合成 Cancel 并清捕获（C12）。
//
// 线程模型：输入事件只从主线程进入（Pointer/Key/IME/Wheel/Cancel）；
// 帧协程会在重建时并发调用 Set/SetRoot/ForgetTree（swapContent）与
// SetFocus（attach 重定向焦点）。mu 只保护这些共享结构，且**所有用户
// 回调都在锁外触发**——回调会再入分发器（键盘路由回 SetFocus），持锁
// 回调就是自己等自己。cursor/stats 只由主线程写，留在锁外。
type Dispatcher struct {
	mu       sync.Mutex
	root     *scene.Node
	handlers map[uint64]Handler
	caps     map[PointerID]capture
	hover    []*scene.Node
	focus    *scene.Node
	cursor   scene.Cursor
	onCursor func(scene.Cursor)
	stats    DispStats
}

// NewDispatcher 创建分发器。
func NewDispatcher(root *scene.Node) *Dispatcher {
	return &Dispatcher{
		root:     root,
		handlers: make(map[uint64]Handler, 32),
		caps:     make(map[PointerID]capture, 4),
	}
}

// SetRoot 更换根节点（重建换树/重挂载时用）。
//
// 与旧语义的关键差别：**不清指针捕获、不清悬停链**。重建换掉的是节点，
// 不是正在进行的手势——按下到抬起之间整树重建是常态（悬停高亮就靠
// rebuild 换色），捕获若被清掉，抬起就走重新命中：新树新节点新竞技场，
// tap 永远不触发（多窗探针实测的"点击无反应"根因）。捕获沿旧路径完成
// 手势，回调仍落在同一个控件实例上；悬停链交给 updateHover 的 Key 差分：
// 同一组件换节点不算状态迁移，不重复发 Enter/Leave。
func (d *Dispatcher) SetRoot(root *scene.Node) {
	d.mu.Lock()
	d.root = root
	d.mu.Unlock()
}

// Root 返回根节点。
func (d *Dispatcher) Root() *scene.Node { return d.root }

// ForgetTree 清理刚被卸载的子树的 handler 登记（重建换树后调用）。
//
// 重建每次产出全新节点：旧登记不清，handlers 表随重建无限增长
// （悬停切换、逐键输入都是重建）。两类例外必须保留——
//   - 活跃捕获路径上的节点：手势要在"按下那一刻的路径"上完成（C26）；
//   - 悬停链上的节点：指针移开后还要能给旧节点发 Leave（否则控件的
//     hover 位卡在 true，重绘高亮不消）。
func (d *Dispatcher) ForgetTree(root *scene.Node) {
	if root == nil {
		return
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	keep := make(map[uint64]bool, 8)
	for _, c := range d.caps {
		for _, n := range c.path {
			keep[n.ID()] = true
		}
	}
	for _, n := range d.hover {
		keep[n.ID()] = true
	}
	var walk func(*scene.Node)
	walk = func(n *scene.Node) {
		if !keep[n.ID()] {
			delete(d.handlers, n.ID())
		}
		for _, ch := range n.Children() {
			walk(ch)
		}
	}
	walk(root)
}

// Set 挂处理器；零值 Handler 等价于移除。
func (d *Dispatcher) Set(n *scene.Node, h Handler) {
	if n == nil {
		return
	}
	d.mu.Lock()
	if h.empty() {
		delete(d.handlers, n.ID())
	} else {
		d.handlers[n.ID()] = h
	}
	d.mu.Unlock()
}

// Handler 返回节点的处理器。
func (d *Dispatcher) Handler(n *scene.Node) (Handler, bool) {
	if n == nil {
		return Handler{}, false
	}
	d.mu.Lock()
	h, ok := d.handlers[n.ID()]
	d.mu.Unlock()
	return h, ok
}

// SetFocus 设置焦点节点（Key/IME 的投递目标）。
func (d *Dispatcher) SetFocus(n *scene.Node) {
	d.mu.Lock()
	if d.focus == n {
		d.mu.Unlock()
		return
	}
	old := d.focus
	d.focus = n
	var hOld, hNew Handler
	if old != nil {
		hOld = d.handlers[old.ID()]
	}
	if n != nil {
		hNew = d.handlers[n.ID()]
	}
	d.mu.Unlock()
	// FocusOut 先于 FocusIn：让旧节点有机会提交/取消编辑状态。
	if hOld.Focus != nil {
		hOld.Focus(&FocusEvent{Kind: FocusOut, Node: old})
	}
	if hNew.Focus != nil {
		hNew.Focus(&FocusEvent{Kind: FocusIn, Node: n})
	}
}

// Focus 返回焦点节点。
func (d *Dispatcher) Focus() *scene.Node {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.focus
}

// Cursor 返回当前光标形状。
func (d *Dispatcher) Cursor() scene.Cursor { return d.cursor }

// OnCursorChanged 注册光标变化回调（host 用它 SetCursor）。
func (d *Dispatcher) OnCursorChanged(fn func(scene.Cursor)) { d.onCursor = fn }

// Stats 返回统计快照。
func (d *Dispatcher) Stats() DispStats { return d.stats }

// Captured 返回某指针当前捕获的目标（nil 表示未捕获）。
func (d *Dispatcher) Captured(id PointerID) *scene.Node {
	d.mu.Lock()
	defer d.mu.Unlock()
	if c, ok := d.caps[id]; ok {
		return c.target
	}
	return nil
}

// ── 指针 ────────────────────────────────────────────────────────────

// hoverPair 是"待触发的悬停迁移"快照：节点与它的处理器一起取，
// 触发在锁外（Enter/Leave 会 MarkDirty → 帧协程重建 → ForgetTree）。
type hoverPair struct {
	n *scene.Node
	h Handler
}

// sameTarget 判断两个节点是否同一逻辑悬停目标：同指针，或同一 Key 身份
// （SetKey 非空且相等——重建换节点但组件没换的情形）。
func sameTarget(a, b *scene.Node) bool {
	if a == nil || b == nil {
		return false
	}
	if a == b {
		return true
	}
	if k := a.Key(); k != "" {
		return k == b.Key()
	}
	return false
}

// stageHoverLocked 计算悬停差分、更新 hover 链并快照待触发项（调用方持锁）。
//
// 差分按"逻辑目标"比较而不是节点指针，是悬停高亮不抖动的关键：重建后的
// 新路径与旧悬停链在指针层面完全不同，但组件身份相同——不该有任何
// Enter/Leave，否则"Enter→重建→指针差异→再 Enter"每移动一次鼠标就
// 重建整树一次（hover 风暴），而且重建恰好落在 down/up 之间时 tap 必丢。
func (d *Dispatcher) stageHoverLocked(nu []*scene.Node) (enters, leaves []hoverPair) {
	old := d.hover
	common := 0
	for common < len(old) && common < len(nu) && sameTarget(old[common], nu[common]) {
		common++
	}
	oldSuf, newSuf := old[common:], nu[common:]
	// 分歧后缀里按身份配对：同一组件换了节点（含深度/祖先变化）不算迁移。
	survives := func(n *scene.Node, list []*scene.Node) bool {
		for _, m := range list {
			if sameTarget(n, m) {
				return true
			}
		}
		return false
	}
	for i := len(oldSuf) - 1; i >= 0; i-- {
		n := oldSuf[i]
		if survives(n, newSuf) {
			continue
		}
		if h, ok := d.handlers[n.ID()]; ok && h.Leave != nil {
			leaves = append(leaves, hoverPair{n, h})
		}
		d.stats.Leaves++
	}
	for i := 0; i < len(newSuf); i++ {
		n := newSuf[i]
		if survives(n, oldSuf) {
			continue
		}
		if h, ok := d.handlers[n.ID()]; ok && h.Enter != nil {
			enters = append(enters, hoverPair{n, h})
		}
		d.stats.Enters++
	}
	d.hover = append(d.hover[:0], nu...)
	return enters, leaves
}

// stageCursorLocked 更新光标形状（调用方持锁），返回新光标与是否变化。
func (d *Dispatcher) stageCursorLocked(path Path) (scene.Cursor, bool) {
	c := path.Cursor()
	if c == d.cursor {
		return c, false
	}
	d.cursor = c
	d.stats.CursorChanges++
	return c, true
}

// snapshotLocked 取路径上各节点的处理器快照（调用方持锁）。
// 回调在锁外触发，分发期间用快照而不是每次查表。
func (d *Dispatcher) snapshotLocked(nodes []*scene.Node) []Handler {
	if len(nodes) == 0 {
		return nil
	}
	snap := make([]Handler, len(nodes))
	for i, n := range nodes {
		snap[i] = d.handlers[n.ID()]
	}
	return snap
}

// applyHover 触发悬停迁移（锁外）：先全部 Leave（深 → 浅），再 Enter。
func (d *Dispatcher) applyHover(enters, leaves []hoverPair, ev *PointerEvent) {
	for _, p := range leaves {
		p.h.Leave(ev)
	}
	for _, p := range enters {
		p.h.Enter(ev)
	}
}

// fireCursor 光标变化落地（锁外；onCursor 在装配期注册一次）。
func (d *Dispatcher) fireCursor(cur scene.Cursor, changed bool) {
	if changed && d.onCursor != nil {
		d.onCursor(cur)
	}
}

// PointerIn 把指针事件限定在指定子树内分发（浮层路由用）。
//
// 与 Pointer 的差别：HitTest 以 root（浮层内容根）为起点——事件只能命中
// 该子树内的节点，模态遮罩后面的内容树天然不可达。down/up 两次调用传同一
// 个 root，捕获路径（caps）内的节点都在浮层子树内，语义自洽。
// 调用方（容器）必须保证 root 是浮层 content：这些节点的 handler 由
// ctx.Input.Set 登记，与全局表共用。
func (d *Dispatcher) PointerIn(root *scene.Node, ev *PointerEvent) bool {
	if root == nil || ev == nil {
		return false
	}
	d.mu.Lock()
	saved := d.root
	d.root = root
	d.mu.Unlock()
	defer func() {
		d.mu.Lock()
		d.root = saved
		d.mu.Unlock()
	}()
	return d.Pointer(ev)
}

// Pointer 分发一个指针事件，返回是否被处理。
func (d *Dispatcher) Pointer(ev *PointerEvent) bool {
	if ev == nil {
		return false
	}
	d.stats.Pointers++
	switch ev.Kind {
	case PointerDown:
		d.mu.Lock()
		path := HitTest(d.root, ev.PosRoot)
		if path.Empty() {
			delete(d.caps, ev.Pointer)
			_, leaves := d.stageHoverLocked(nil)
			d.stageCursorLocked(Path{})
			d.mu.Unlock()
			d.applyHover(nil, leaves, ev)
			return false
		}
		d.caps[ev.Pointer] = capture{target: path.Target(), path: path.Nodes()}
		d.stats.Captures++
		enters, leaves := d.stageHoverLocked(path.nodes)
		cur, curChanged := d.stageCursorLocked(path)
		snap := d.snapshotLocked(path.nodes)
		d.mu.Unlock()
		d.applyHover(enters, leaves, ev)
		d.fireCursor(cur, curChanged)
		return d.fire(snap, path.nodes, path.locals, ev)

	case PointerMove:
		var (
			snap           []Handler
			nodes          []*scene.Node
			locals         []geom.Point
			enters, leaves []hoverPair
			cur            scene.Cursor
			curChanged     bool
		)
		d.mu.Lock()
		if c, ok := d.caps[ev.Pointer]; ok {
			// 捕获：直接路由，**不重新命中**（否则拖动会中途换响应者）
			nodes, locals = c.path, mapDown(ev.PosRoot, c.path)
			snap = d.snapshotLocked(nodes)
		} else {
			path := HitTest(d.root, ev.PosRoot)
			nodes, locals = path.nodes, path.locals
			snap = d.snapshotLocked(nodes)
			enters, leaves = d.stageHoverLocked(path.nodes)
			cur, curChanged = d.stageCursorLocked(path)
		}
		d.mu.Unlock()
		d.applyHover(enters, leaves, ev)
		d.fireCursor(cur, curChanged)
		return d.fire(snap, nodes, locals, ev)

	case PointerUp:
		return d.finish(ev)

	case PointerCancel:
		return d.Cancel(ev.Pointer, ev)

	case PointerEnter, PointerHover:
		// 无按键移动也必须算 hover 差分：宿主把"没按键的移动"归成 Hover，
		// 如果这里不做差分，鼠标悬停按钮永远不会高亮（因为整个事件流里
		// 根本不会出现 Down）。光标同理——手型光标就是靠 hover 更新出来的。
		d.mu.Lock()
		path := HitTest(d.root, ev.PosRoot)
		snap := d.snapshotLocked(path.nodes)
		enters, leaves := d.stageHoverLocked(path.nodes)
		cur, curChanged := d.stageCursorLocked(path)
		d.mu.Unlock()
		d.applyHover(enters, leaves, ev)
		d.fireCursor(cur, curChanged)
		return d.fire(snap, path.nodes, path.locals, ev)

	case PointerLeave:
		// 指针离开窗口：没有任何节点处于 hover，路径必须清空而不是
		// "在当前坐标再命中一次"——那会让离开前最后一个节点一直保持高亮。
		d.mu.Lock()
		_, leaves := d.stageHoverLocked(nil)
		d.stageCursorLocked(Path{})
		d.mu.Unlock()
		d.applyHover(nil, leaves, ev)
		return false
	}
	return false
}

// finish 处理抬起：沿捕获路径分发，然后释放捕获并重算 hover。
func (d *Dispatcher) finish(ev *PointerEvent) bool {
	var (
		snap   []Handler
		nodes  []*scene.Node
		locals []geom.Point
	)
	d.mu.Lock()
	c, hasCap := d.caps[ev.Pointer]
	if hasCap {
		nodes, locals = c.path, mapDown(ev.PosRoot, c.path)
	} else {
		path := HitTest(d.root, ev.PosRoot)
		nodes, locals = path.nodes, path.locals
	}
	snap = d.snapshotLocked(nodes)
	d.mu.Unlock()

	handled := d.fire(snap, nodes, locals, ev)

	d.mu.Lock()
	delete(d.caps, ev.Pointer)
	// 抬起后指针已经不在按下位置：hover 必须重算，否则高亮会卡住。
	path := HitTest(d.root, ev.PosRoot)
	enters, leaves := d.stageHoverLocked(path.nodes)
	cur, curChanged := d.stageCursorLocked(path)
	d.mu.Unlock()
	d.applyHover(enters, leaves, ev)
	d.fireCursor(cur, curChanged)
	return handled
}

// Cancel 合成一次取消（C12：丢失捕获必须补 Cancel，不能等组件自己发现）。
//
// ev 为 nil 表示由 host 触发（焦点被夺、进入模态循环）。取消事件先发给
// 捕获目标，再释放捕获——顺序反了的话，目标会在"已经不是捕获者"的状态下
// 收到 Cancel，回滚逻辑可能拒绝执行。
func (d *Dispatcher) Cancel(id PointerID, ev *PointerEvent) bool {
	d.mu.Lock()
	c, ok := d.caps[id]
	if !ok {
		d.mu.Unlock()
		return false
	}
	if ev == nil {
		ev = &PointerEvent{Kind: PointerCancel, Pointer: id, Device: DeviceMouse}
	}
	ev.Kind = PointerCancel
	snap := d.snapshotLocked(c.path)
	d.mu.Unlock()

	handled := d.fire(snap, c.path, mapDown(ev.PosRoot, c.path), ev)

	d.mu.Lock()
	delete(d.caps, id)
	_, leaves := d.stageHoverLocked(nil)
	d.stageCursorLocked(Path{})
	d.stats.Cancels++
	d.mu.Unlock()
	d.applyHover(nil, leaves, ev)
	return handled
}

// ── 滚轮 / 键盘 / 输入法 ────────────────────────────────────────────

// Wheel 分发滚轮事件：目标是命中节点，未处理则沿父链冒泡（嵌套滚动链）。
func (d *Dispatcher) Wheel(ev *WheelEvent) bool {
	if ev == nil {
		return false
	}
	path := HitTest(d.root, ev.Pos)
	d.mu.Lock()
	snap := d.snapshotLocked(path.nodes)
	d.mu.Unlock()
	for i := len(path.nodes) - 1; i >= 0; i-- {
		if h := snap[i]; h.Wheel != nil {
			if h.Wheel(ev) {
				d.stats.Handled++
				return true
			}
		}
	}
	return false
}

// Key 分发键盘事件到焦点节点，未处理则沿父链冒泡。
//
// 焦点链在事件开始时快照：回调（h.Key → 焦点路由 → SetFocus）在锁外，
// 可能改写 d.focus，但本次事件的传播链不应被中途改变。
func (d *Dispatcher) Key(ev *KeyEvent) bool {
	if ev == nil {
		return false
	}
	d.mu.Lock()
	chain := make([]Handler, 0, 8)
	for n := d.focus; n != nil; n = n.Parent() {
		chain = append(chain, d.handlers[n.ID()])
	}
	d.mu.Unlock()
	for _, h := range chain {
		if h.Key != nil {
			if h.Key(ev) {
				d.stats.Handled++
				return true
			}
		}
	}
	return false
}

// IME 分发输入法事件到焦点节点，未处理则沿父链冒泡。
func (d *Dispatcher) IME(ev *IMEEvent) bool {
	if ev == nil {
		return false
	}
	d.mu.Lock()
	chain := make([]Handler, 0, 8)
	for n := d.focus; n != nil; n = n.Parent() {
		chain = append(chain, d.handlers[n.ID()])
	}
	d.mu.Unlock()
	for _, h := range chain {
		if h.IME != nil {
			if h.IME(ev) {
				d.stats.Handled++
				return true
			}
		}
	}
	return false
}

// ── 内部 ────────────────────────────────────────────────────────────

// fire 走一遍捕获 → 目标 → 冒泡（快照驱动，回调在锁外）。
//
// 每次回调前把 ev.Pos 换成该节点局部坐标：组件不该自己做逆变换，
// 那会把"命中用的矩阵"和"组件以为的矩阵"分成两份，迟早对不上。
func (d *Dispatcher) fire(snap []Handler, nodes []*scene.Node, locals []geom.Point, ev *PointerEvent) bool {
	if len(nodes) == 0 {
		return false
	}
	orig := ev.Pos
	defer func() { ev.Pos = orig }()

	for i := range nodes {
		h := snap[i]
		if h.Capture == nil {
			continue
		}
		if i < len(locals) {
			ev.Pos = locals[i]
		}
		if h.Capture(ev) {
			d.stats.Handled++
			return true
		}
	}

	ti := len(nodes) - 1
	if h := snap[ti]; h.Pointer != nil {
		if ti < len(locals) {
			ev.Pos = locals[ti]
		}
		if h.Pointer(ev) {
			d.stats.Handled++
			return true
		}
	}

	for i := ti - 1; i >= 0; i-- {
		h := snap[i]
		if h.Bubble == nil {
			continue
		}
		if i < len(locals) {
			ev.Pos = locals[i]
		}
		if h.Bubble(ev) {
			d.stats.Handled++
			return true
		}
	}
	return false
}
