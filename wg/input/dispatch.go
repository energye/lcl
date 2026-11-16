package input

import (
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
type Dispatcher struct {
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

// SetRoot 更换根节点（重挂载时用）。
func (d *Dispatcher) SetRoot(root *scene.Node) {
	d.root = root
	d.hover = nil
	d.caps = make(map[PointerID]capture, 4)
}

// Root 返回根节点。
func (d *Dispatcher) Root() *scene.Node { return d.root }

// Set 挂处理器；零值 Handler 等价于移除。
func (d *Dispatcher) Set(n *scene.Node, h Handler) {
	if n == nil {
		return
	}
	if h.empty() {
		delete(d.handlers, n.ID())
		return
	}
	d.handlers[n.ID()] = h
}

// Handler 返回节点的处理器。
func (d *Dispatcher) Handler(n *scene.Node) (Handler, bool) {
	if n == nil {
		return Handler{}, false
	}
	h, ok := d.handlers[n.ID()]
	return h, ok
}

// SetFocus 设置焦点节点（Key/IME 的投递目标）。
func (d *Dispatcher) SetFocus(n *scene.Node) {
	if d.focus == n {
		return
	}
	old := d.focus
	d.focus = n
	// FocusOut 先于 FocusIn：让旧节点有机会提交/取消编辑状态。
	d.sendFocus(old, FocusOut)
	d.sendFocus(n, FocusIn)
}

// Focus 返回焦点节点。
func (d *Dispatcher) Focus() *scene.Node { return d.focus }

// Cursor 返回当前光标形状。
func (d *Dispatcher) Cursor() scene.Cursor { return d.cursor }

// OnCursorChanged 注册光标变化回调（host 用它 SetCursor）。
func (d *Dispatcher) OnCursorChanged(fn func(scene.Cursor)) { d.onCursor = fn }

// Stats 返回统计快照。
func (d *Dispatcher) Stats() DispStats { return d.stats }

// Captured 返回某指针当前捕获的目标（nil 表示未捕获）。
func (d *Dispatcher) Captured(id PointerID) *scene.Node {
	if c, ok := d.caps[id]; ok {
		return c.target
	}
	return nil
}

// ── 指针 ────────────────────────────────────────────────────────────

// Pointer 分发一个指针事件，返回是否被处理。
func (d *Dispatcher) Pointer(ev *PointerEvent) bool {
	if ev == nil {
		return false
	}
	d.stats.Pointers++
	switch ev.Kind {
	case PointerDown:
		path := HitTest(d.root, ev.PosRoot)
		if path.Empty() {
			d.release(ev.Pointer)
			d.updateHover(Path{}, ev)
			d.updateCursor(Path{})
			return false
		}
		d.caps[ev.Pointer] = capture{target: path.Target(), path: path.Nodes()}
		d.stats.Captures++
		d.updateHover(path, ev)
		d.updateCursor(path)
		return d.dispatch(ev, path.Nodes(), path.locals)

	case PointerMove:
		if c, ok := d.caps[ev.Pointer]; ok {
			// 捕获：直接路由，**不重新命中**（否则拖动会中途换响应者）
			return d.dispatch(ev, c.path, mapDown(ev.PosRoot, c.path))
		}
		path := HitTest(d.root, ev.PosRoot)
		d.updateHover(path, ev)
		d.updateCursor(path)
		return d.dispatch(ev, path.Nodes(), path.locals)

	case PointerUp:
		return d.finish(ev)

	case PointerCancel:
		return d.Cancel(ev.Pointer, ev)

	case PointerEnter, PointerLeave, PointerHover:
		path := HitTest(d.root, ev.PosRoot)
		return d.dispatch(ev, path.Nodes(), path.locals)
	}
	return false
}

// finish 处理抬起：沿捕获路径分发，然后释放捕获并重算 hover。
func (d *Dispatcher) finish(ev *PointerEvent) bool {
	c, hasCap := d.caps[ev.Pointer]
	var handled bool
	if hasCap {
		handled = d.dispatch(ev, c.path, mapDown(ev.PosRoot, c.path))
	} else {
		path := HitTest(d.root, ev.PosRoot)
		handled = d.dispatch(ev, path.Nodes(), path.locals)
	}
	d.release(ev.Pointer)
	// 抬起后指针已经不在按下位置：hover 必须重算，否则高亮会卡住。
	path := HitTest(d.root, ev.PosRoot)
	d.updateHover(path, ev)
	d.updateCursor(path)
	return handled
}

// Cancel 合成一次取消（C12：丢失捕获必须补 Cancel，不能等组件自己发现）。
//
// ev 为 nil 表示由 host 触发（焦点被夺、进入模态循环）。取消事件先发给
// 捕获目标，再释放捕获——顺序反了的话，目标会在"已经不是捕获者"的状态下
// 收到 Cancel，回滚逻辑可能拒绝执行。
func (d *Dispatcher) Cancel(id PointerID, ev *PointerEvent) bool {
	c, ok := d.caps[id]
	if !ok {
		return false
	}
	if ev == nil {
		ev = &PointerEvent{Kind: PointerCancel, Pointer: id, Device: DeviceMouse}
	}
	ev.Kind = PointerCancel
	handled := d.dispatch(ev, c.path, mapDown(ev.PosRoot, c.path))
	d.release(id)
	d.updateHover(Path{}, ev)
	d.updateCursor(Path{})
	d.stats.Cancels++
	return handled
}

// release 释放指针捕获。
func (d *Dispatcher) release(id PointerID) {
	delete(d.caps, id)
}

// ── 滚轮 / 键盘 / 输入法 ────────────────────────────────────────────

// Wheel 分发滚轮事件：目标是命中节点，未处理则沿父链冒泡（嵌套滚动链）。
func (d *Dispatcher) Wheel(ev *WheelEvent) bool {
	if ev == nil {
		return false
	}
	path := HitTest(d.root, ev.Pos)
	for i := len(path.nodes) - 1; i >= 0; i-- {
		n := path.nodes[i]
		if h, ok := d.handlers[n.ID()]; ok && h.Wheel != nil {
			if h.Wheel(ev) {
				d.stats.Handled++
				return true
			}
		}
	}
	return false
}

// Key 分发键盘事件到焦点节点，未处理则沿父链冒泡。
func (d *Dispatcher) Key(ev *KeyEvent) bool {
	if ev == nil {
		return false
	}
	for n := d.focus; n != nil; n = n.Parent() {
		if h, ok := d.handlers[n.ID()]; ok && h.Key != nil {
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
	for n := d.focus; n != nil; n = n.Parent() {
		if h, ok := d.handlers[n.ID()]; ok && h.IME != nil {
			if h.IME(ev) {
				d.stats.Handled++
				return true
			}
		}
	}
	return false
}

// ── 内部 ────────────────────────────────────────────────────────────

// dispatch 走一遍捕获 → 目标 → 冒泡。
//
// 每次回调前把 ev.Pos 换成该节点局部坐标：组件不该自己做逆变换，
// 那会把"命中用的矩阵"和"组件以为的矩阵"分成两份，迟早对不上。
func (d *Dispatcher) dispatch(ev *PointerEvent, nodes []*scene.Node, locals []geom.Point) bool {
	if len(nodes) == 0 {
		return false
	}
	orig := ev.Pos
	defer func() { ev.Pos = orig }()

	for i, n := range nodes {
		h, ok := d.handlers[n.ID()]
		if !ok || h.Capture == nil {
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
	if h, ok := d.handlers[nodes[ti].ID()]; ok && h.Pointer != nil {
		if ti < len(locals) {
			ev.Pos = locals[ti]
		}
		if h.Pointer(ev) {
			d.stats.Handled++
			return true
		}
	}

	for i := ti - 1; i >= 0; i-- {
		h, ok := d.handlers[nodes[i].ID()]
		if !ok || h.Bubble == nil {
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

// updateHover 用路径差分合成 Enter/Leave。
//
// 差分而不是"每次全发"：Enter/Leave 是状态迁移语义，重复发送会让
// 悬停动画反复重启（示例里的按钮 hover 抖动正是这么来的）。
func (d *Dispatcher) updateHover(path Path, ev *PointerEvent) {
	old := d.hover
	nu := path.nodes
	common := 0
	for common < len(old) && common < len(nu) && old[common] == nu[common] {
		common++
	}
	for i := len(old) - 1; i >= common; i-- {
		if h, ok := d.handlers[old[i].ID()]; ok && h.Leave != nil {
			h.Leave(ev)
		}
		d.stats.Leaves++
	}
	for i := common; i < len(nu); i++ {
		if h, ok := d.handlers[nu[i].ID()]; ok && h.Enter != nil {
			h.Enter(ev)
		}
		d.stats.Enters++
	}
	d.hover = append(d.hover[:0], nu...)
}

// updateCursor 更新光标形状（宿主用它调 SetCursor）。
func (d *Dispatcher) updateCursor(path Path) {
	c := path.Cursor()
	if c == d.cursor {
		return
	}
	d.cursor = c
	d.stats.CursorChanges++
	if d.onCursor != nil {
		d.onCursor(c)
	}
}

// sendFocus 发送焦点事件。
func (d *Dispatcher) sendFocus(n *scene.Node, kind FocusKind) {
	if n == nil {
		return
	}
	if h, ok := d.handlers[n.ID()]; ok && h.Focus != nil {
		h.Focus(&FocusEvent{Kind: kind, Node: n})
	}
}
