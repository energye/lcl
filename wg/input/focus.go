package input

import (
	"sort"
	"sync/atomic"

	"github.com/energye/lcl/wg/scene"
)

// ── 焦点 ────────────────────────────────────────────────────────────
//
// 系统焦点恒在 Form（那一个 HWND）：Windows 不会把焦点给一个没有句柄的
// 子区域。所以"哪个组件在接收键盘"完全是 Go 侧的概念——FocusNode 树。
// 好处是无窗口自检也能完整跑一遍 Tab 链与模态陷阱（04-UI-INTERACTION.md §B6）。

// focusSeq 是创建序号，用于 (Order, 创建序) 的稳定排序。
//
// 为什么需要第二关键字：同 Order（默认都是 0）时若按 map/指针顺序排，
// Tab 顺序会在不同运行间漂移——测试挂一次再跑就绿，是最难查的一类 bug。
var focusSeq int64

// FocusNode 是焦点树节点。
//
// 与场景树分离而不是复用 scene.Node：焦点顺序、scope、可聚焦性都是
// "键盘交互"的语义，塞进渲染树会让布局/绘制代码被迫理解 Tab 顺序。
// 两者通过 Node 字段关联（通常是同一个节点的指针）。
type FocusNode struct {
	// Parent 是父节点。
	Parent *FocusNode
	// Order 是显式排序权重（小的在前），相同时按创建序。
	Order int
	// CanFocus 表示可以接收焦点（禁用/只读时为 false）。
	CanFocus bool
	// Node 是关联的场景节点（Key/IME 的投递目标，可为 nil）。
	Node *scene.Node
	// Tag 是诊断标签。
	Tag string
	// OnFocus / OnBlur 在焦点迁移时调用（顺序：旧 Blur → 新 Focus）。
	OnFocus func()
	OnBlur  func()

	children []*FocusNode
	scope    *FocusScope
	seq      int64
	focused  bool
}

// NewFocusNode 创建焦点节点。canFocus 为 false 表示只是分组容器。
func NewFocusNode(canFocus bool) *FocusNode {
	return &FocusNode{CanFocus: canFocus, seq: atomic.AddInt64(&focusSeq, 1)}
}

// AddChild 追加子节点。
func (n *FocusNode) AddChild(c *FocusNode) *FocusNode {
	if c == nil || c == n {
		return n
	}
	c.detach()
	c.Parent = n
	n.children = append(n.children, c)
	return n
}

// Add 是 AddChild 的语法糖，便于链式构造。
func (n *FocusNode) Add(canFocus bool) *FocusNode {
	c := NewFocusNode(canFocus)
	n.AddChild(c)
	return c
}

// RemoveChild 摘除子节点。
func (n *FocusNode) RemoveChild(c *FocusNode) bool {
	for i, x := range n.children {
		if x == c {
			n.children = append(n.children[:i], n.children[i+1:]...)
			c.Parent = nil
			return true
		}
	}
	return false
}

// detach 从当前父节点摘下自己。
func (n *FocusNode) detach() {
	if n.Parent != nil {
		n.Parent.RemoveChild(n)
	}
}

// ChildCount 返回子节点数。
func (n *FocusNode) ChildCount() int { return len(n.children) }

// Children 返回按 (Order, 创建序) 排序的子节点副本。
//
// 返回副本而不是内部切片：调用方（遍历、诊断）不该有能力改结构，
// 排序也不能就地改内部顺序（否则创建序被 Order 覆盖，第二关键字失效）。
func (n *FocusNode) Children() []*FocusNode {
	out := make([]*FocusNode, len(n.children))
	copy(out, n.children)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Order != out[j].Order {
			return out[i].Order < out[j].Order
		}
		return out[i].seq < out[j].seq
	})
	return out
}

// Seq 返回创建序号（诊断用）。
func (n *FocusNode) Seq() int64 { return n.seq }

// Scope 返回本节点自己声明的焦点域（未声明为 nil）。
func (n *FocusNode) Scope() *FocusScope { return n.scope }

// EffectiveScope 返回生效的焦点域（向上找最近的声明）。
func (n *FocusNode) EffectiveScope() *FocusScope {
	for x := n; x != nil; x = x.Parent {
		if x.scope != nil {
			return x.scope
		}
	}
	return nil
}

// Focused 表示当前持有焦点。
func (n *FocusNode) Focused() bool { return n.focused }

// Contains 判断 c 是否在本子树内（含自身）。
func (n *FocusNode) Contains(c *FocusNode) bool {
	for x := c; x != nil; x = x.Parent {
		if x == n {
			return true
		}
	}
	return false
}

// Root 返回所在焦点树的根。
func (n *FocusNode) Root() *FocusNode {
	x := n
	for x.Parent != nil {
		x = x.Parent
	}
	return x
}

// Depth 返回层级（根为 0）。
func (n *FocusNode) Depth() int {
	d := 0
	for x := n.Parent; x != nil; x = x.Parent {
		d++
	}
	return d
}

// ── 焦点域 ──────────────────────────────────────────────────────────

// FocusScope 是焦点域（Modal / Drawer / 弹出层各自一个）。
//
// 域的意义是"Tab 链的边界"：对话框打开后 Tab 不该跑到底下的页面上去。
// Trap 为 true 时边界变循环（模态必须如此，否则焦点会"消失"在背景里）。
type FocusScope struct {
	// Name 是诊断名。
	Name string
	// Parent 是外层域（nil 表示根域）。
	Parent *FocusScope
	// Trap 表示越界时循环而不是退到外层。
	Trap bool

	// root 是声明本域的节点：它属于外层域，它的子树属于本域。
	root    *FocusNode
	restore *FocusNode
	owner   *FocusManager
}

// NewFocusScope 在 node 上声明一个焦点域。
func NewFocusScope(name string, node *FocusNode) *FocusScope {
	s := &FocusScope{Name: name, Trap: true, root: node}
	if node != nil {
		s.Parent = node.EffectiveScope()
		node.scope = s
	}
	return s
}

// Root 返回域根节点。
func (s *FocusScope) Root() *FocusNode { return s.root }

// SetTrap 设置是否循环。
func (s *FocusScope) SetTrap(v bool) *FocusScope { s.Trap = v; return s }

// ── 高亮模式 ────────────────────────────────────────────────────────

// HighlightMode 表示焦点环的显示策略。
//
// 为什么要分模式：鼠标点击按钮后按钮留下蓝色焦点环，是桌面端最典型的
// "像网页一样脏"的观感（§B6）。由最近一次输入设备决定，不由代码决定。
type HighlightMode uint8

const (
	// HighlightKeyboard 显示焦点环（最近一次输入来自键盘）。
	HighlightKeyboard HighlightMode = iota
	// HighlightPointer 不显示焦点环（最近一次输入来自指针）。
	HighlightPointer
)

// String 返回可读名。
func (m HighlightMode) String() string {
	if m == HighlightPointer {
		return "pointer"
	}
	return "keyboard"
}

// FocusStats 是焦点统计（诊断与门禁用）。
type FocusStats struct {
	Changes  int
	Tabs     int
	Blurs    int
	Suspends int
	Scopes   int
	Migrates int
}

// FocusManager 管理焦点树、Tab 链与焦点域。
//
// 焦点迁移只有一个入口（SetCurrent）：多入口必然出现"两个地方都以为自己
// 是焦点"的错位，而这类错位只在用户连续点击时暴露。
type FocusManager struct {
	root  *FocusNode
	cur   *FocusNode
	scope *FocusScope
	mode  HighlightMode
	disp  *Dispatcher

	saved     *FocusNode
	suspended bool
	stats     FocusStats
	onChange  func(cur, old *FocusNode)
}

// NewFocusManager 创建焦点管理器。root 可为 nil（稍后 SetRoot）。
func NewFocusManager(root *FocusNode) *FocusManager {
	m := &FocusManager{root: root}
	m.scope = ownerScope(root)
	return m
}

// SetRoot 更换根节点（重挂载）。焦点被清空。
func (m *FocusManager) SetRoot(root *FocusNode) {
	if m.cur != nil {
		m.cur.focused = false
		m.cur = nil
	}
	m.root = root
	m.scope = ownerScope(root)
	m.syncDispatcher()
}

// Root 返回根节点。
func (m *FocusManager) Root() *FocusNode { return m.root }

// Current 返回当前焦点节点。
func (m *FocusManager) Current() *FocusNode { return m.cur }

// Attach 绑定分发器：焦点变化会同步过去，Key/IME 才能找到投递目标。
func (m *FocusManager) Attach(d *Dispatcher) { m.disp = d }

// OnChanged 注册焦点变化回调。
func (m *FocusManager) OnChanged(fn func(cur, old *FocusNode)) { m.onChange = fn }

// Mode 返回高亮模式。
func (m *FocusManager) Mode() HighlightMode { return m.mode }

// NoteKeyboard 记录"最近一次输入来自键盘"（host 在按键时调用）。
func (m *FocusManager) NoteKeyboard() { m.mode = HighlightKeyboard }

// NotePointer 记录"最近一次输入来自指针"（host 在按下时调用）。
func (m *FocusManager) NotePointer() { m.mode = HighlightPointer }

// Suspended 表示焦点因窗口失活被临时清空。
func (m *FocusManager) Suspended() bool { return m.suspended }

// Stats 返回统计。
func (m *FocusManager) Stats() FocusStats { return m.stats }

// ActiveScope 返回当前生效的焦点域。
func (m *FocusManager) ActiveScope() *FocusScope { return m.scope }

// SetCurrent 迁移焦点（唯一入口）。
//
// 顺序：旧节点 Blur → 新节点 Focus → 同步分发器。旧节点先收到通知，
// 是因为它可能要提交输入框内容，而提交逻辑依赖"我确实还是焦点"这一事实。
func (m *FocusManager) SetCurrent(n *FocusNode) bool {
	if m.suspended || n == m.cur {
		return false
	}
	if n != nil && !n.CanFocus {
		return false
	}
	old := m.cur
	m.cur = n
	if old != nil {
		old.focused = false
		if old.OnBlur != nil {
			old.OnBlur()
		}
	}
	if n != nil {
		n.focused = true
		if n.OnFocus != nil {
			n.OnFocus()
		}
	}
	m.stats.Changes++
	m.syncDispatcher()
	if m.onChange != nil {
		m.onChange(n, old)
	}
	return true
}

// Focus 请求把焦点给 n。
//
// 模态陷阱外的节点直接拒绝：点击背景不该把焦点从对话框里偷走，
// 否则 Tab 链会从一个"看不见的当前位置"继续，用户按 Tab 毫无反应。
func (m *FocusManager) Focus(n *FocusNode) bool {
	if n == nil {
		return false
	}
	if m.scope != nil && m.scope.Trap && ownerScope(n) != m.scope {
		return false
	}
	return m.SetCurrent(n)
}

// Blur 清空焦点。
func (m *FocusManager) Blur() bool {
	if m.cur == nil {
		return false
	}
	m.stats.Blurs++
	return m.SetCurrent(nil)
}

// syncDispatcher 把焦点同步给分发器（Key/IME 的投递目标）。
func (m *FocusManager) syncDispatcher() {
	if m.disp == nil {
		return
	}
	var n *scene.Node
	if m.cur != nil {
		n = m.cur.Node
	}
	m.disp.SetFocus(n)
}

// ── Tab 链 ──────────────────────────────────────────────────────────

// ownerScope 返回节点归属的域（域根本身属于外层域）。
//
// 域根通常是对话框容器，它自己不是 Tab 停靠点；把它算进自己的域会让
// "离开内层回到外层"少一个落点。
func ownerScope(n *FocusNode) *FocusScope {
	if n == nil {
		return nil
	}
	if n.scope != nil {
		return n.scope.Parent
	}
	return n.EffectiveScope()
}

// flat 返回整棵树的文档序（DFS、子节点按 (Order, 创建序)），只含可聚焦节点。
//
// 不分域地铺平：Tab 链是"在这条序列上走，遇到不属于当前域的节点就跳过"，
// 域边界等价于序列遇到别的域的节点。这样"出域回外层"只是换一个过滤条件
// 继续向后走，不需要重新定位锚点——按锚点定位在域根不可聚焦时必然错位。
func (m *FocusManager) flat() []*FocusNode {
	var out []*FocusNode
	var walk func(n *FocusNode)
	walk = func(n *FocusNode) {
		if n.CanFocus {
			out = append(out, n)
		}
		for _, c := range n.Children() {
			walk(c)
		}
	}
	if m.root != nil {
		walk(m.root)
	}
	return out
}

// listIn 返回某个域内的 Tab 链（按文档序）。
func (m *FocusManager) listIn(s *FocusScope) []*FocusNode {
	var out []*FocusNode
	for _, n := range m.flat() {
		if ownerScope(n) == s {
			out = append(out, n)
		}
	}
	return out
}

// List 返回当前域的 Tab 链（诊断用）。
func (m *FocusManager) List() []*FocusNode { return m.listIn(m.scope) }

// MoveNext 按 Tab 链迁移焦点（back 表示 Shift+Tab）。
//
// 三种边界：
//   - 域内还有下一个 → 过去；
//   - 已到域边界且域是 Trap（模态）→ 绕回域内第一个；
//   - 已到域边界且域非 Trap → 退到外层域，从文档序的同一位置继续往后走
//     （不是从头开始，否则外层已经 Tab 过的控件会被重走一遍）。
func (m *FocusManager) MoveNext(back bool) bool {
	m.stats.Tabs++
	m.mode = HighlightKeyboard

	flat := m.flat()
	if len(flat) == 0 {
		return false
	}
	from := -1 // 虚拟起点：forward 从 0 开始、back 从末尾开始
	for i, n := range flat {
		if n == m.cur {
			from = i
			break
		}
	}
	if from < 0 && back {
		from = len(flat)
	}
	dir := 1
	if back {
		dir = -1
	}

	// scan 从 from 出发沿 dir 走，返回第一个属于 s 的节点；
	// 走出序列（= 域边界）返回 nil。
	scan := func(s *FocusScope) *FocusNode {
		for j := from + dir; j >= 0 && j < len(flat); j += dir {
			if ownerScope(flat[j]) == s {
				return flat[j]
			}
		}
		return nil
	}
	// scanWrap 绕回序列另一端继续找（Trap 域用）。
	scanWrap := func(s *FocusScope) *FocusNode {
		for k := 1; k <= len(flat); k++ {
			j := ((from+dir*k)%len(flat) + len(flat)) % len(flat)
			if j == from {
				break
			}
			if ownerScope(flat[j]) == s {
				return flat[j]
			}
		}
		return nil
	}

	s := m.scope
	for {
		if n := scan(s); n != nil {
			return m.SetCurrent(n)
		}
		if s == nil || s.Trap {
			if n := scanWrap(s); n != nil {
				return m.SetCurrent(n)
			}
			return false
		}
		// 非 Trap 出界 → 外层继续
		m.stats.Migrates++
		s = s.Parent
		m.scope = s
	}
}

// HandleKey 处理导航键（Tab / Shift+Tab）。
//
// 只吃 Down/Repeat：Windows 在按住 Tab 时发 Repeat，抬起时的 Up 不该
// 再触发一次迁移。
func (m *FocusManager) HandleKey(ev *KeyEvent) bool {
	if ev == nil || ev.Kind == KeyUp {
		return false
	}
	if ev.Code != CodeTab {
		return false
	}
	m.NoteKeyboard()
	return m.MoveNext(ev.Mods.Has(ModShift))
}

// ── 焦点域切换 ──────────────────────────────────────────────────────

// EnterScope 进入焦点域（打开模态）。当前焦点被记住，退出时还原。
func (m *FocusManager) EnterScope(s *FocusScope) {
	if s == nil {
		return
	}
	s.restore = m.cur
	m.stats.Scopes++
	m.scope = s
	if list := m.listIn(s); len(list) > 0 {
		m.SetCurrent(list[0])
	} else {
		m.SetCurrent(nil)
	}
}

// ExitScope 退出当前焦点域（关闭模态），还原进入前的焦点。
func (m *FocusManager) ExitScope() {
	s := m.scope
	if s == nil {
		return
	}
	m.stats.Scopes++
	m.scope = s.Parent
	if s.restore != nil && s.restore.CanFocus {
		m.SetCurrent(s.restore)
		s.restore = nil
		return
	}
	s.restore = nil
	if list := m.listIn(m.scope); len(list) > 0 {
		m.SetCurrent(list[0])
	} else {
		m.SetCurrent(nil)
	}
}

// ── 窗口失活 ────────────────────────────────────────────────────────

// Suspend 在窗口失活时清空键盘焦点（Alt+Tab 切走再切回来，焦点要还原）。
func (m *FocusManager) Suspend() {
	if m.suspended {
		return
	}
	m.suspended = true
	m.saved = m.cur
	m.stats.Suspends++
	if m.cur != nil {
		m.cur.focused = false
		if m.cur.OnBlur != nil {
			m.cur.OnBlur()
		}
	}
	m.cur = nil
	m.syncDispatcher()
}

// Resume 在窗口重新激活时还原焦点。
func (m *FocusManager) Resume() {
	if !m.suspended {
		return
	}
	m.suspended = false
	n := m.saved
	m.saved = nil
	if n != nil && n.CanFocus {
		m.cur = n
		n.focused = true
		if n.OnFocus != nil {
			n.OnFocus()
		}
		m.syncDispatcher()
		if m.onChange != nil {
			m.onChange(n, nil)
		}
		return
	}
	m.syncDispatcher()
}
