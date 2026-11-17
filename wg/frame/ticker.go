package frame

import (
	"sync"
	"time"
)

// Ticker 是"到点叫我"的抽象。
//
// 为什么不让调度器直接用 time.Ticker：固定 60Hz 的心跳会把"静止零帧"（C14）
// 变成"每秒 60 次空转"，也会让动画节奏被心跳周期量化。调度器要的是
// **按需装填的闹钟**：有活就装，没活就拆。
//
// Arm 的语义是"替换"而不是"排队"：同一时刻只有一个待触发的安排，
// 重复装填以最后一次为准（帧调度本来就不允许积压，见 C11）。
type Ticker interface {
	// Arm 安排在 d 之后触发 f（d <= 0 表示尽快）。会取消上一次未触发的安排。
	Arm(d time.Duration, f func())
	// Disarm 取消未触发的安排。
	Disarm()
	// Stop 释放底层资源（之后 Arm 无效）。
	Stop()
}

// timerTicker 是基于 time.AfterFunc 的真实实现。
type timerTicker struct {
	mu      sync.Mutex
	timer   *time.Timer
	gen     uint64
	stopped bool
}

// NewTicker 创建真实定时器。
func NewTicker() Ticker { return &timerTicker{} }

// Arm 实现 Ticker。
func (t *timerTicker) Arm(d time.Duration, f func()) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.stopped || f == nil {
		return
	}
	if t.timer != nil {
		t.timer.Stop()
	}
	if d < 0 {
		d = 0
	}
	t.gen++
	gen := t.gen
	// AfterFunc 的 Stop 与已触发之间存在竞态，用代次号丢弃过期触发。
	t.timer = time.AfterFunc(d, func() {
		t.mu.Lock()
		dropped := t.stopped || gen != t.gen
		t.mu.Unlock()
		if dropped {
			return
		}
		f()
	})
}

// Disarm 实现 Ticker。
func (t *timerTicker) Disarm() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.gen++
	if t.timer != nil {
		t.timer.Stop()
		t.timer = nil
	}
}

// Stop 实现 Ticker。
func (t *timerTicker) Stop() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.stopped = true
	t.gen++
	if t.timer != nil {
		t.timer.Stop()
		t.timer = nil
	}
}

// ManualTicker 是确定性 ticker：时间由调用方推进，用于无窗口门禁与单测。
//
// 它绑到 ManualClock 上，保证"时钟"与"闹钟"永远同步——门禁里最怕的就是
// 两者各走各的，断言出来的相位差全是测试自身造成的。
type ManualTicker struct {
	mu       sync.Mutex
	clock    *ManualClock
	now      time.Duration
	deadline time.Duration
	armed    bool
	f        func()
	fires    int
	stopped  bool
}

// NewManualTicker 创建手动 ticker；clock 为 nil 时用内部时间轴。
func NewManualTicker(clock *ManualClock) *ManualTicker {
	return &ManualTicker{clock: clock}
}

// Arm 实现 Ticker。
func (t *ManualTicker) Arm(d time.Duration, f func()) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.stopped || f == nil {
		return
	}
	if d < 0 {
		d = 0
	}
	if t.clock != nil {
		t.now = t.clock.Now()
	}
	t.deadline = t.now + d
	t.armed = true
	t.f = f
}

// Disarm 实现 Ticker。
func (t *ManualTicker) Disarm() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.armed = false
	t.f = nil
}

// Stop 实现 Ticker。
func (t *ManualTicker) Stop() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.stopped = true
	t.armed = false
	t.f = nil
}

// Armed 表示是否有待触发的安排。
func (t *ManualTicker) Armed() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.armed
}

// Deadline 返回待触发的绝对时刻（未装填时为零值）。
func (t *ManualTicker) Deadline() time.Duration {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.deadline
}

// Fires 返回累计触发次数。
func (t *ManualTicker) Fires() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.fires
}

// Advance 推进 d 并触发一次到期的回调，返回是否触发。
//
// 一次只触发一次：如果回调里又 Arm 了一个"立刻"的安排，也要等下一次
// Advance。这避免了"回调内装填 → 立刻再触发"的递归，也让门禁里的
// 帧序列严格可控（一推进 = 至多一帧）。
func (t *ManualTicker) Advance(d time.Duration) bool {
	if t.clock != nil {
		t.clock.AdvanceBy(d)
	}
	t.mu.Lock()
	if t.stopped {
		t.mu.Unlock()
		return false
	}
	if t.clock != nil {
		t.now = t.clock.Now()
	} else if d > 0 {
		t.now += d
	}
	due := t.armed && t.now >= t.deadline
	var fn func()
	if due {
		fn = t.f
		t.armed = false
		t.f = nil
		t.fires++
	}
	t.mu.Unlock()
	if fn != nil {
		fn()
		return true
	}
	return false
}

// Now 返回手动 ticker 的时间轴当前值。
func (t *ManualTicker) Now() time.Duration {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.now
}
