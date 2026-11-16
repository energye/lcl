package input

import (
	"time"

	"github.com/energye/lcl/wg/geom"
)

// Recognizer 是手势识别器（竞技场成员）。
//
// 生命周期由竞技场驱动：AddPointer 之后所有事件都通过 Handle 转发，
// 最终必收到一次 Accept 或 Reject（二选一，且只一次）——这条是 C27 的核心：
// 单指针只有一个胜者，败者必须被明确告知以便回滚视觉状态。
//
// **Accept 是唯一"开始执行手势"的入口**：识别器在 Handle 里只做判断并
// arena.Resolve(r)，绝不自己调用业务回调。否则"Sweep 选出来的胜者"这条
// 路径（长按的 Tick 没跑、抬起时无人宣布）就永远不会触发回调——同一个
// 手势会因为"谁先发现"而时灵时不灵。
type Recognizer interface {
	// AddPointer 接收新指针（按下那一刻）。
	AddPointer(ev *PointerEvent)
	// Handle 处理该指针的后续事件（Move/Up/Cancel）。
	Handle(ev *PointerEvent)
	// Accept 被裁决为胜者：现在执行手势并通知业务。
	Accept()
	// Reject 被裁决为败者（必须回滚视觉状态）。
	Reject()
}

// ArenaState 是竞技场状态。
type ArenaState uint8

const (
	// ArenaOpen 还能接纳新成员（按下事件分发期间）。
	ArenaOpen ArenaState = iota
	// ArenaClosed 已封闭，等待裁决。
	ArenaClosed
	// ArenaResolved 已有胜者。
	ArenaResolved
)

// member 是竞技场成员及其"是否已退出竞争"状态。
type member struct {
	rec Recognizer
	// out 表示该成员已退出竞争（自己放弃或已被拒），sweep 时不再考虑。
	out bool
}

// Arena 是单指针的手势竞技场。
//
// 为什么必须有它：同一个按下常常被多个角色认领（表格里拖滑块 = Slider.Drag
// vs 表格横向 Pan；列表项 = Tap vs LongPress）。让每个组件自己写互斥逻辑
// 必然打架，而且"谁先反应"取决于事件顺序这种不可控因素。竞技场把裁决
// 集中成一条规则：谁先明确宣布谁赢；没人宣布则指针抬起时第一个**仍在场**的
// 成员赢（C27）。
type Arena struct {
	members []member
	state   ArenaState
	winner  Recognizer
	won     bool

	// names 是成员诊断名。没有它，门禁与现场日志只能打印 "某个识别器赢了"，
	// 而"表格里横向滚动失效"这类问题正需要知道到底是谁抢走了手势。
	names  map[Recognizer]string
	losers []string

	// 统计与诊断
	added    int
	rejected int
	sweeps   int
	// OnResolved 在裁决完成时调用（诊断用）。
	OnResolved func(winner Recognizer, members int)
}

// NewArena 创建竞技场。
func NewArena() *Arena { return &Arena{state: ArenaOpen} }

// Add 加入一个成员。只有在开放状态下才能加入。
func (a *Arena) Add(r Recognizer) bool { return a.AddNamed(r, "") }

// AddNamed 加入成员并设诊断名。
func (a *Arena) AddNamed(r Recognizer, name string) bool {
	if r == nil || a.state != ArenaOpen {
		return false
	}
	a.members = append(a.members, member{rec: r})
	a.added++
	if name != "" {
		if a.names == nil {
			a.names = make(map[Recognizer]string, 4)
		}
		a.names[r] = name
	}
	return true
}

// SetName 为已加入的成员设诊断名。
func (a *Arena) SetName(r Recognizer, name string) {
	if a.names == nil {
		a.names = make(map[Recognizer]string, 4)
	}
	a.names[r] = name
}

// nameOf 返回成员诊断名。
func (a *Arena) nameOf(r Recognizer) string { return a.names[r] }

// WinnerName 返回胜者诊断名。
func (a *Arena) WinnerName() string { return a.nameOf(a.Winner()) }

// Losers 返回被拒成员的诊断名（按被拒顺序）。
func (a *Arena) Losers() []string { return a.losers }

// Close 封闭竞技场（该指针的按下事件已分发完）。
func (a *Arena) Close() {
	if a.state == ArenaOpen {
		a.state = ArenaClosed
	}
}

// State 返回当前状态。
func (a *Arena) State() ArenaState { return a.state }

// Members 返回成员数。
func (a *Arena) Members() int { return len(a.members) }

// Winner 返回胜者（未裁决或无人胜出时为 nil）。
func (a *Arena) Winner() Recognizer {
	if !a.won {
		return nil
	}
	return a.winner
}

// Resolve 由成员宣布胜出。
//
// 只有"还没人赢"的竞技场接受宣布：这保证了单指针单胜者——
// 第二个宣布的成员会得到 false，它必须当作失败处理（C27 的另一半：
// 败者不得改变状态）。
func (a *Arena) Resolve(r Recognizer) bool {
	if r == nil || a.state == ArenaResolved || a.won || len(a.members) == 0 {
		return false
	}
	i := a.indexOf(r)
	if i < 0 {
		return false
	}
	a.winner = r
	a.won = true
	a.state = ArenaResolved
	a.members[i].out = false
	a.rejectOthers(r)
	if a.OnResolved != nil {
		a.OnResolved(r, len(a.members))
	}
	// 先拒后接受：败者先回滚视觉，胜者再开始，避免两个"激活态"同时存在。
	r.Accept()
	return true
}

// Withdraw 让成员退出竞争（它自己判断手势不成立）。
//
// 与 Reject 的区别：Withdraw 是自愿退出，不调用 Reject 回调（调用方已经
// 自己回滚过了）；Reject 是裁决结果，必须通知。
func (a *Arena) Withdraw(r Recognizer) {
	if i := a.indexOf(r); i >= 0 {
		a.members[i].out = true
	}
}

// Sweep 在指针抬起（或超时）时裁决：没人宣布则第一个仍在场的成员胜出。
func (a *Arena) Sweep() Recognizer {
	a.sweeps++
	if a.state == ArenaResolved {
		return a.Winner()
	}
	a.state = ArenaResolved
	for i := range a.members {
		if a.members[i].out {
			continue
		}
		a.winner = a.members[i].rec
		a.won = true
		a.members[i].out = false
		a.rejectOthers(a.winner)
		if a.OnResolved != nil {
			a.OnResolved(a.winner, len(a.members))
		}
		a.winner.Accept()
		return a.winner
	}
	return nil
}

// RejectAll 全部拒绝（指针被取消、或事件流异常）。
func (a *Arena) RejectAll() {
	if a.state == ArenaResolved {
		return
	}
	a.state = ArenaResolved
	for i := range a.members {
		if a.members[i].out {
			continue
		}
		a.members[i].out = true
		a.rejected++
		a.losers = append(a.losers, a.nameOf(a.members[i].rec))
		a.members[i].rec.Reject()
	}
}

// indexOf 返回成员下标。
func (a *Arena) indexOf(r Recognizer) int {
	for i, m := range a.members {
		if m.rec == r {
			return i
		}
	}
	return -1
}

// rejectOthers 拒绝除胜者外仍在场的成员。
func (a *Arena) rejectOthers(winner Recognizer) {
	for i := range a.members {
		if a.members[i].rec == winner || a.members[i].out {
			continue
		}
		a.members[i].out = true
		a.rejected++
		a.losers = append(a.losers, a.nameOf(a.members[i].rec))
		a.members[i].rec.Reject()
	}
}

// ArenaStats 是竞技场统计。
type ArenaStats struct {
	Added    int
	Rejected int
	Sweeps   int
}

// Stats 返回统计。
func (a *Arena) Stats() ArenaStats {
	return ArenaStats{Added: a.added, Rejected: a.rejected, Sweeps: a.sweeps}
}

// ── 识别器 ──────────────────────────────────────────────────────────

// GestureConfig 是识别器参数（默认值见 DefaultGestureConfig）。
type GestureConfig struct {
	// Slop 是"还算点击"的最大位移（DIP）。
	Slop float32
	// LongPress 是长按判定时长。
	LongPress time.Duration
	// MaxTapDelay 是点击的最长按压时长。
	MaxTapDelay time.Duration
	// AxisLock 表示只认主轴方向（横向拖 vs 纵向拖靠它区分）。
	AxisLock Axis
}

// Axis 是手势主轴。
type Axis uint8

const (
	// AxisNone 不锁定方向。
	AxisNone Axis = iota
	// AxisHorizontal 只认横向。
	AxisHorizontal
	// AxisVertical 只认纵向。
	AxisVertical
)

// DefaultGestureConfig 返回默认参数（04-UI-INTERACTION.md §B5.3）。
func DefaultGestureConfig() GestureConfig {
	return GestureConfig{
		Slop:        4,
		LongPress:   500 * time.Millisecond,
		MaxTapDelay: 500 * time.Millisecond,
		AxisLock:    AxisNone,
	}
}

// tapRec 是点击识别器。
type tapRec struct {
	arena *Arena
	cfg   GestureConfig
	onTap func(geom.Point)
	onEnd func(accepted bool)

	down    geom.Point
	downAt  time.Time
	started bool
	done    bool
}

// Tap 创建点击识别器：按下到抬起位移 < Slop 且时长受限即算点击。
//
// onEnd 在裁决后调用一次（accepted=false 表示落败，调用方据此回滚
// "按下高亮"之类的视觉状态）。
func Tap(a *Arena, cfg GestureConfig, onTap func(geom.Point), onEnd func(accepted bool)) Recognizer {
	return &tapRec{arena: a, cfg: cfg, onTap: onTap, onEnd: onEnd}
}

func (r *tapRec) AddPointer(ev *PointerEvent) {
	r.down = ev.Pos
	r.downAt = ev.Time
	r.started = true
}

func (r *tapRec) Handle(ev *PointerEvent) {
	if !r.started || r.done {
		return
	}
	switch ev.Kind {
	case PointerMove:
		if ev.Pos.Sub(r.down).Len() > r.cfg.Slop {
			// 超出 slop：自己退出竞争（不宣布，别人还有机会）
			r.arena.Withdraw(r)
			r.finish(false)
		}
	case PointerUp:
		if ev.Pos.Sub(r.down).Len() > r.cfg.Slop || r.tooLong(ev) {
			// 不成立：主动退出，别让 Sweep 把自己选成胜者
			r.arena.Withdraw(r)
			r.finish(false)
			return
		}
		// 宣布胜出后由 Accept 统一触发 onTap；宣布失败说明别人已经赢了。
		if !r.arena.Resolve(r) {
			r.finish(false)
		}
	case PointerCancel:
		r.arena.Withdraw(r)
		r.finish(false)
	}
}

// tooLong 判断按压时长是否超过点击上限。
func (r *tapRec) tooLong(ev *PointerEvent) bool {
	if ev.Time.IsZero() || r.downAt.IsZero() {
		return false
	}
	return ev.Time.Sub(r.downAt) > r.cfg.MaxTapDelay
}

func (r *tapRec) Accept() {
	if r.done {
		return
	}
	r.finish(true)
	if r.onTap != nil {
		// 用落点而不是抬起点：双击判定比较的是两次落点距离。
		r.onTap(r.down)
	}
}

func (r *tapRec) Reject() { r.finish(false) }

func (r *tapRec) finish(accepted bool) {
	if r.done {
		return
	}
	r.done = true
	if r.onEnd != nil {
		r.onEnd(accepted)
	}
}

// Tickable 是需要"时间推进"才能判定的识别器（长按类）。
//
// 单独拆出接口而不是把 Tick 塞进 Recognizer：绝大多数识别器只关心事件，
// 平白多一个空方法会让实现者困惑"我该在这里做什么"。
type Tickable interface {
	Recognizer
	// Tick 推进到 now，返回是否在本次判定中胜出。
	Tick(now time.Time) bool
}

// longPressRec 是长按识别器。
type longPressRec struct {
	arena *Arena
	cfg   GestureConfig
	onLP  func(geom.Point)
	down  geom.Point
	downT time.Time
	live  bool
	done  bool
}

// LongPress 创建长按识别器。
//
// 时长判定不靠自己的定时器，而靠 Tick(now)：定时器要额外一条线程与
// 一整套取消逻辑，而调度器本来每帧都会推进时间。门禁里 Tick 由用例显式调用。
func LongPress(a *Arena, cfg GestureConfig, onLongPress func(geom.Point)) Tickable {
	return &longPressRec{arena: a, cfg: cfg, onLP: onLongPress}
}

func (r *longPressRec) AddPointer(ev *PointerEvent) {
	r.down = ev.Pos
	r.downT = ev.Time
	r.live = true
}

func (r *longPressRec) Handle(ev *PointerEvent) {
	if !r.live || r.done {
		return
	}
	switch ev.Kind {
	case PointerMove:
		if ev.Pos.Sub(r.down).Len() > r.cfg.Slop {
			r.arena.Withdraw(r)
			r.live = false
			r.done = true
		}
	case PointerUp, PointerCancel:
		r.live = false
		r.done = true
	}
}

// Tick 推进时间：到点且指针仍按住则宣布胜出。
//
// 返回 true 表示本次判定赢了（onLongPress 已由 Accept 触发）。
func (r *longPressRec) Tick(now time.Time) bool {
	if !r.live || r.done {
		return false
	}
	if r.downT.IsZero() || now.Sub(r.downT) < r.cfg.LongPress {
		return false
	}
	if !r.arena.Resolve(r) {
		r.done = true
		return false
	}
	return true
}

// Accept 在宣布成功或被 Sweep 选中时触发长按回调。
//
// Sweep 选中这条路径是真实存在的：宿主没跑 Tick（例如空闲时不出帧），
// 用户按住 800ms 再松手，谁也不宣布，Sweep 会挑第一个仍在场的成员——
// 如果那时不触发回调，长按就"按了没反应"。
func (r *longPressRec) Accept() {
	if r.done {
		return
	}
	r.done = true
	r.live = false
	if r.onLP != nil {
		r.onLP(r.down)
	}
}

func (r *longPressRec) Reject() { r.done = true; r.live = false }

// dragRec 是拖动识别器（Pan / 单轴 Drag 共用一个实现，靠 AxisLock 区分）。
type dragRec struct {
	arena   *Arena
	cfg     GestureConfig
	onStart func(geom.Point)
	onMove  func(pos, delta geom.Point)
	onEnd   func(pos, velocity geom.Point)

	tracker *VelocityTracker
	down    geom.Point
	last    geom.Point
	live    bool
	won     bool
	done    bool
}

// Drag 创建拖动识别器。
func Drag(a *Arena, cfg GestureConfig,
	onStart func(geom.Point), onMove func(pos, delta geom.Point), onEnd func(pos, velocity geom.Point),
) Recognizer {
	return &dragRec{arena: a, cfg: cfg, onStart: onStart, onMove: onMove, onEnd: onEnd,
		tracker: NewVelocityTracker()}
}

func (r *dragRec) AddPointer(ev *PointerEvent) {
	r.down = ev.Pos
	r.last = ev.Pos
	r.live = true
	r.tracker.Reset().Add(ev.Pos, ev.Time)
}

func (r *dragRec) Handle(ev *PointerEvent) {
	if !r.live || r.done {
		return
	}
	switch ev.Kind {
	case PointerMove:
		r.tracker.Add(ev.Pos, ev.Time)
		d := ev.Pos.Sub(r.down)
		if !r.won {
			if d.Len() <= r.cfg.Slop {
				return
			}
			// 方向锁：主轴不占优就自己退出，别去抢别人的手势
			if !r.axisOK(d) {
				r.arena.Withdraw(r)
				r.live = false
				r.done = true
				return
			}
			if !r.arena.Resolve(r) {
				r.done = true
				return
			}
			// Accept 已把 onStart 发出去，之后走正常移动分支。
			r.last = ev.Pos
			return
		}
		if r.onMove != nil {
			r.onMove(ev.Pos, ev.Pos.Sub(r.last))
		}
		r.last = ev.Pos
	case PointerUp, PointerCancel:
		r.live = false
		r.done = true
		if r.won && r.onEnd != nil {
			v := geom.Point{}
			if ev.Kind == PointerUp {
				v = r.tracker.VelocityAt(ev.Time)
			}
			r.onEnd(ev.Pos, v)
		}
	}
}

// axisOK 判定位移是否落在锁定的主轴上。
func (r *dragRec) axisOK(d geom.Point) bool {
	ax, ay := d.X, d.Y
	if ax < 0 {
		ax = -ax
	}
	if ay < 0 {
		ay = -ay
	}
	switch r.cfg.AxisLock {
	case AxisHorizontal:
		return ax >= ay
	case AxisVertical:
		return ay > ax
	default:
		return true
	}
}

// Accept 开始拖动：onStart 只在这里发一次。
func (r *dragRec) Accept() {
	if r.won {
		return
	}
	r.won = true
	if r.onStart != nil {
		r.onStart(r.down)
	}
}

func (r *dragRec) Reject() {
	r.won = false
	r.done = true
	r.live = false
}
