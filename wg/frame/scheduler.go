package frame

import (
	"errors"
	"sync"
	"time"
)

// 调度器默认参数。
const (
	// DefaultMaxDelta 是单帧时间增量的上限。
	//
	// 为什么需要上限：窗口最小化、系统睡眠、调试断点都会让墙钟一次跳过几秒。
	// 按真实间隔积分动画会得到"恢复瞬间动画已经播完"的跳变，
	// 而按帧数积分又会慢放（这正是 C15 要禁的东西）。取 250ms 作为
	// "一次卡顿的上限"，超过就按上限积分——动画会慢一点，但不会跳。
	DefaultMaxDelta = 250 * time.Millisecond

	// DefaultResizeDebounce 是 resize 节流窗口（C10）：一个 60Hz 帧间隔。
	// 连续拖动窗口边缘时 WM_SIZE 以 ~10–15ms 的间隔涌来（比帧间隔还密）。
	// 纯尾随去抖在拖动中一次都不会触发（每次请求都把闹钟推后），backing
	// 停在旧尺寸，Presenter 全程拒提交、窗口冻结在旧帧 = "内容不跟窗"。
	// 节流窗口 = 一帧：拖动中**每帧跟随一次**（>1 帧的请求合并到下一拍），
	// 停下时尾随保证最后一次尺寸落地；间歇由 DWM 拉伸旧帧兜底。
	DefaultResizeDebounce = DefaultStep
)

// Options 是调度器配置。
type Options struct {
	// Clock 是帧时间源（必填）。
	Clock FrameClock
	// Ticker 是帧闹钟（必填）。
	Ticker Ticker
	// ResizeTicker 是 resize 节流闹钟（可选，默认用真实定时器）。
	// 单独一条是因为它和帧节奏无关：帧闹钟随时可能被拆掉，节流窗口不能被牵连。
	ResizeTicker Ticker
	// AsyncPresent 表示提交是异步的：一帧画完后必须等 Present()
	// 回填帧槽，才允许申请下一帧（C11）。
	//
	// 软件后端是同步呈现，槽在帧内就回填了，所以默认 false；
	// 异步模式存在是为了 GPU/合成器路径，也是 --schedtest 验证 C11 的场景。
	AsyncPresent bool
	// MaxDelta 是单帧增量上限（0 = DefaultMaxDelta）。
	MaxDelta time.Duration
	// ResizeDebounce 是 resize 节流窗口（0 = DefaultResizeDebounce）。
	ResizeDebounce time.Duration
	// OnFrame 执行一帧：布局 → 录制 → 合成 → 提交。返回 true 表示真的提交了。
	//
	// 回调在锁外调用：里面可以安全地再调 RequestFrame（动画继续）。
	OnFrame func(now, delta time.Duration) bool
	// OnResize 在节流前沿或尾随到期时被调用（尺寸为最新值）。
	OnResize func(w, h int)
	// WantFrame 回答"还需要帧吗、多久之后"。nil 表示没有动画系统。
	WantFrame func(now time.Duration) (need bool, in time.Duration)
	// OnSkip 在"帧槽未回填却收到唤醒"时调用（诊断/追踪用）。
	OnSkip func()
}

// SchedStats 是调度统计。
type SchedStats struct {
	// Frames 是已执行的帧数。
	Frames int
	// Presents 是真正提交的帧数（OnFrame 返回 true）。
	Presents int
	// Idle 是执行了但没提交的帧（静止帧）。
	Idle int
	// Skipped 是因帧槽未回填而被丢弃的唤醒次数（C11）。
	Skipped int
	// Requested 是 RequestFrame 的调用次数（用于观察合并率）。
	Requested int
	// Clamped 是 delta 被 MaxDelta 截断的次数。
	Clamped int
	// Resizes 是收到的 resize 请求数，ResizeFires 是节流后真正处理的次数。
	Resizes     int
	ResizeFires int
	// LastDelta 是最近一帧的时间增量。
	LastDelta time.Duration
}

// Scheduler 是单 outstanding 帧调度器。
//
// 它只做三件事，每件都对应一条契约：
//
//  1. **按需出帧**（C14）：没有挂起工作、没有动画需求时不装闹钟，
//     于是静止界面一帧都不出，而不是每秒 60 次空转。
//  2. **单帧槽**（C11/C13）：同一时刻只有一帧在飞；异步呈现时槽没回填
//     就拒绝下一次唤醒，宁可丢一拍也不排队（排队会在卡顿时变成滚雪球）。
//  3. **墙钟相位**（C15）：帧时间来自 Clock.Now()，增量是"这次与上次之差"，
//     不是固定步长。动画因此按真实时间推进，掉帧只会变粗不会变慢。
type Scheduler struct {
	mu           sync.Mutex
	opts         Options
	resizeTicker Ticker

	frame    int
	lastNow  time.Duration
	started  bool
	pending  bool
	awaiting bool
	armed    bool
	armedAt  time.Duration
	// inTick 表示帧回调正在执行：单 outstanding 的硬保证。
	//
	// 帧内任何 arm(0)（RequestFrame/平台 WM_PAINT 回调）都可能让定时器
	// 在 OnFrame 尚未返回时再次触发 Tick——没有这道闸，两帧会并发跑
	// 布局/合成（-race 实测双 Tick 竞争场景树）。重入拍直接丢弃，
	// 但记 wakePending：帧末 reschedule 补装，否则那一拍携带的"下一帧"
	// 意图（尤其是动画续帧的 RequestFrameIn）会随定时器一起消失，帧泵
	// 永久停摆（-race 下渲染 >16.7ms 时实测复现）。
	inTick bool
	// wakePending 表示有一拍唤醒在 inTick/awaiting 期间被丢弃，需要在
	// 当前帧结束（或 Present 回填）后补装。
	wakePending bool

	rw, rh        int
	resizePending bool
	// resizeFired 表示真正执行过 resize（此后按最后执行时刻做节流判断）。
	resizeFired  bool
	lastResizeAt time.Duration

	stats  SchedStats
	closed bool
}

// NewScheduler 创建调度器。
func NewScheduler(opts Options) (*Scheduler, error) {
	if opts.Clock == nil {
		return nil, errors.New("wg/frame: 缺少帧时钟")
	}
	if opts.Ticker == nil {
		return nil, errors.New("wg/frame: 缺少帧闹钟")
	}
	if opts.MaxDelta <= 0 {
		opts.MaxDelta = DefaultMaxDelta
	}
	if opts.ResizeDebounce <= 0 {
		opts.ResizeDebounce = DefaultResizeDebounce
	}
	rt := opts.ResizeTicker
	if rt == nil {
		rt = NewTicker()
	}
	// 帧闹钟的精度前提：Windows 默认 15.6ms 粒度会把 16ms 帧钉成 ~31ms。
	ensureTimerResolution()
	return &Scheduler{opts: opts, resizeTicker: rt}, nil
}

// Frame 返回已执行的帧数。
func (s *Scheduler) Frame() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.frame
}

// Stats 返回统计快照。
func (s *Scheduler) Stats() SchedStats {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stats
}

// Awaiting 表示帧槽是否还没回填。
func (s *Scheduler) Awaiting() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.awaiting
}

// Armed 表示是否已装填下一次帧闹钟。
func (s *Scheduler) Armed() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.armed
}

// RequestFrame 申请一帧（输入、属性变化、动画步进都走这里）。
//
// 只装闹钟、不直接出帧：帧的唯一入口是 Tick，这样"一帧"的边界永远是
// 明确的，不会因为某个 setter 在奇怪的地方同步画一次而破坏帧语义。
func (s *Scheduler) RequestFrame() {
	s.mu.Lock()
	s.pending = true
	s.stats.Requested++
	need := !s.closed && !s.awaiting
	s.mu.Unlock()
	if need {
		s.arm(0)
	}
}

// RequestFrameIn 申请"d 之后的一帧"（动画定时用）。
func (s *Scheduler) RequestFrameIn(d time.Duration) {
	s.mu.Lock()
	s.stats.Requested++
	need := !s.closed && !s.awaiting
	s.mu.Unlock()
	if need {
		s.arm(d)
	}
}

// Tick 执行一帧（由 Ticker 唤醒调用）。
func (s *Scheduler) Tick() {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return
	}
	if s.inTick {
		// 同帧重入：这一拍对应的闹钟已经触发（定时器已被消费），必须
		// 清掉 armed/armedAt；同时记 wakePending，帧末 reschedule 补装，
		// 否则这一拍携带的续帧意图会丢失（帧泵永久停摆）。
		s.armed = false
		s.armedAt = 0
		s.wakePending = true
		s.mu.Unlock()
		return
	}
	if s.awaiting {
		// C11：上一帧还没被呈现消费掉，这一拍直接丢弃，不入队。
		// 同上：唤醒已被消费，armed 清掉，意图记入 wakePending，
		// Present 回填后的 reschedule 补装。
		s.stats.Skipped++
		s.armed = false
		s.armedAt = 0
		s.wakePending = true
		skip := s.opts.OnSkip
		s.mu.Unlock()
		if skip != nil {
			skip()
		}
		return
	}
	s.armed = false
	s.inTick = true
	now := s.opts.Clock.Now()
	delta := time.Duration(0)
	if s.started {
		delta = now - s.lastNow
		if delta < 0 {
			delta = 0
		}
		if delta > s.opts.MaxDelta {
			delta = s.opts.MaxDelta
			s.stats.Clamped++
		}
	}
	s.pending = false
	onFrame := s.opts.OnFrame
	s.mu.Unlock()

	var presented bool
	if onFrame != nil {
		presented = onFrame(now, delta)
	}

	s.mu.Lock()
	s.started = true
	s.lastNow = now
	s.frame++
	s.stats.Frames++
	s.stats.LastDelta = delta
	if presented {
		s.stats.Presents++
		if s.opts.AsyncPresent {
			s.awaiting = true
		}
	} else {
		s.stats.Idle++
	}
	s.inTick = false
	s.mu.Unlock()

	s.reschedule()
}

// Present 回填帧槽（异步呈现时由 host 在"像素已上屏"后调用）。
func (s *Scheduler) Present() {
	s.mu.Lock()
	if s.closed || !s.awaiting {
		s.mu.Unlock()
		return
	}
	s.awaiting = false
	s.mu.Unlock()
	s.reschedule()
}

// RequestResize 申请改变合成目标尺寸。
//
// 节流 + 尾随语义（C10）：距上次真正执行满 ResizeDebounce（=一帧）就立即
// 执行（前沿，首次请求也走这里）；否则把执行钉在 lastFire+ResizeDebounce
// （尾随，多次请求只留最后一次尺寸）。连续拖动窗口 = 每帧跟随一次，
// 停下后必有一次收尾——纯尾随去抖在 WM_SIZE 间隔小于窗口时永远等不到
// 静默期，拖拽全程一次都不执行（backing 停在旧尺寸，提交全程被拒）。
func (s *Scheduler) RequestResize(w, h int) {
	if w <= 0 || h <= 0 {
		return
	}
	s.mu.Lock()
	s.rw, s.rh = w, h
	s.resizePending = true
	s.stats.Resizes++
	if s.closed {
		s.mu.Unlock()
		return
	}
	now := s.opts.Clock.Now()
	window := s.opts.ResizeDebounce
	if !s.resizeFired || now-s.lastResizeAt >= window {
		// 旧尾随闹钟在锁内拆掉：锁外拆会误杀并发请求刚装的新闹钟。
		// 即使它随后到期也会因 resizePending=false 空转。
		s.resizeTicker.Disarm()
		s.mu.Unlock()
		// 同步执行：帧渲染已在主线程（P0 方案①），消息线程不会再与
		// 帧线程的整帧合成互等；c.Resize 只剩 comp.Resize 重建（~0.5ms）。
		s.fireResize()
		return
	}
	delay := window - (now - s.lastResizeAt)
	s.mu.Unlock()
	s.resizeTicker.Arm(delay, s.fireResize)
}

// fireResize 真正执行一次 resize（前沿直达或尾随闹钟到期）。
func (s *Scheduler) fireResize() {
	s.mu.Lock()
	if s.closed || !s.resizePending {
		s.mu.Unlock()
		return
	}
	w, h := s.rw, s.rh
	s.resizePending = false
	s.stats.ResizeFires++
	s.resizeFired = true
	s.lastResizeAt = s.opts.Clock.Now()
	cb := s.opts.OnResize
	s.mu.Unlock()

	if cb != nil {
		cb(w, h)
	}
	// 尺寸变了整屏都脏：立刻申请一帧，否则用户会看到"旧内容被拉伸"。
	s.RequestFrame()
}

// Stop 停止调度（之后所有申请与唤醒都被忽略）。
func (s *Scheduler) Stop() {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return
	}
	s.closed = true
	s.armed = false
	tk := s.opts.Ticker
	s.mu.Unlock()
	tk.Stop()
	s.resizeTicker.Stop()
}

// ── 内部 ────────────────────────────────────────────────────────────

// arm 装填帧闹钟，采用"更早者胜"策略。
//
// 为什么不是"每次都重装"：一帧里可能有几十个 RequestFrame（每个脏节点一次），
// 每次都 Stop+AfterFunc 既浪费又可能把已经快到的闹钟推后。
// 只在"新 deadline 更早"时才换，语义正确且调用廉价。
func (s *Scheduler) arm(d time.Duration) {
	if d < 0 {
		d = 0
	}
	s.mu.Lock()
	if s.closed || s.awaiting {
		s.mu.Unlock()
		return
	}
	now := s.opts.Clock.Now()
	at := now + d
	if s.armed && s.armedAt <= at {
		s.mu.Unlock()
		return
	}
	tk := s.opts.Ticker
	s.armed = true
	s.armedAt = at
	s.mu.Unlock()
	tk.Arm(d, s.Tick)
}

// disarm 拆掉闹钟。
func (s *Scheduler) disarm() {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return
	}
	if !s.armed {
		s.mu.Unlock()
		return
	}
	s.armed = false
	tk := s.opts.Ticker
	s.mu.Unlock()
	tk.Disarm()
}

// reschedule 重新计算下一次帧的时刻。
//
// 决策顺序就是优先级：挂起工作 > 动画需求 > 不装闹钟（静止零帧，C14）。
// WantFrame 在锁外调用：它是上层代码，可能顺手 RequestFrame。
func (s *Scheduler) reschedule() {
	s.mu.Lock()
	if s.closed {
		s.mu.Unlock()
		return
	}
	if s.awaiting {
		// 槽未回填：不装闹钟；wakePending 保留到 Present 回填后再补装。
		s.mu.Unlock()
		s.disarm()
		return
	}
	if s.wakePending {
		// 有被丢弃的唤醒：补装一拍（调用方会按挂起/动画需求继续续期）。
		s.wakePending = false
		s.mu.Unlock()
		s.arm(0)
		return
	}
	if s.pending {
		s.mu.Unlock()
		s.arm(0)
		return
	}
	want := s.opts.WantFrame
	now := s.opts.Clock.Now()
	s.mu.Unlock()

	if want == nil {
		// 没有动画系统（WantFrame=nil）：帧的续期完全由显式的
		// RequestFrame/RequestFrameIn 管理——**不能 disarm**。帧内
		// RequestFrameIn(16ms) 刚武装的动画续帧闹钟会被这里拆掉，
		// 动画只出一帧就停（真窗多窗探针实测的停止根因）。
		// 已武装的闹钟到点自会 Tick；静止时本无武装，自然零帧（C14）。
		// （锁已在上方解锁；此处不得再 Unlock——双重解锁会 fatal。）
		return
	}
	need, in := want(now)
	if need {
		s.arm(0)
		return
	}
	if in > 0 {
		s.arm(in)
		return
	}
	s.disarm()
}
