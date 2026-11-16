package ui

import (
	"time"

	"github.com/energye/lcl/wg/theme"
)

// ── 动画 ────────────────────────────────────────────────────────────
//
// 三条硬规则（§A4）：
//   - 只由帧时钟驱动，组件**禁止**自开 time.Ticker/TTimer（C15）；
//   - 相位来自**墙钟**差值，不是帧计数——否则掉帧时动画变慢；
//   - 尊重"减少动效"：时长归零但状态变更照走。

// AnimStatus 是动画状态。
type AnimStatus uint8

const (
	// AnimIdle 静止。
	AnimIdle AnimStatus = iota
	// AnimForward 正向播放。
	AnimForward
	// AnimReverse 反向播放。
	AnimReverse
	// AnimCompleted 已完成（停在终值）。
	AnimCompleted
)

// String 返回可读名。
func (s AnimStatus) String() string {
	switch s {
	case AnimIdle:
		return "idle"
	case AnimForward:
		return "forward"
	case AnimReverse:
		return "reverse"
	case AnimCompleted:
		return "completed"
	}
	return "unknown"
}

// Controller 是 0..1 的动画控制器。
type Controller struct {
	from, to float32
	value    float32
	dur      time.Duration
	curve    theme.Curve
	status   AnimStatus

	// start 是本段动画的起始墙钟（毫秒）。
	start  time.Duration
	last   time.Duration
	inited bool

	// OnUpdate 在值变化时调用。
	OnUpdate func(float32)
}

// NewController 创建控制器（初值 0，时长与曲线默认取主题中值）。
func NewController(dur time.Duration, curve theme.Curve) *Controller {
	if curve == nil {
		curve = theme.Default().Motion.Easing
	}
	return &Controller{dur: dur, curve: curve, value: 0, status: AnimIdle}
}

// SetDuration 设置时长（0 表示瞬间完成）。
func (c *Controller) SetDuration(d time.Duration) *Controller { c.dur = d; return c }

// SetCurve 设置曲线。
func (c *Controller) SetCurve(curve theme.Curve) *Controller {
	if curve != nil {
		c.curve = curve
	}
	return c
}

// Status 返回状态。
func (c *Controller) Status() AnimStatus { return c.status }

// Value 返回当前值（0..1）。
func (c *Controller) Value() float32 { return c.value }

// Animating 表示是否还在动。
func (c *Controller) Animating() bool {
	return c.status == AnimForward || c.status == AnimReverse
}

// Forward 正向播放（从当前值到 1）。
func (c *Controller) Forward(now time.Duration) { c.play(1, AnimForward, now) }

// Reverse 反向播放（从当前值到 0）。
func (c *Controller) Reverse(now time.Duration) { c.play(0, AnimReverse, now) }

// To 播放到指定值。
func (c *Controller) To(target float32, now time.Duration) {
	st := AnimForward
	if target < c.value {
		st = AnimReverse
	}
	c.play(target, st, now)
}

// play 开始一段动画。
func (c *Controller) play(target float32, st AnimStatus, now time.Duration) {
	c.from = c.value
	c.to = target
	c.start = now
	c.last = now
	c.inited = true
	if c.dur <= 0 || c.from == target {
		// 时长为 0（含"减少动效"）：直接落在终值，状态仍然完成——
		// 只是没有中间帧，不能让依赖"动画结束"的逻辑永远等不到。
		c.value = target
		c.status = AnimCompleted
		if c.OnUpdate != nil {
			c.OnUpdate(c.value)
		}
		return
	}
	c.status = st
}

// Stop 停在当前位置。
func (c *Controller) Stop() {
	c.status = AnimIdle
	c.inited = true
}

// SetValue 直接设值（不带动画）。
func (c *Controller) SetValue(v float32) {
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	c.value = v
	c.status = AnimIdle
	if c.OnUpdate != nil {
		c.OnUpdate(v)
	}
}

// Step 推进到墙钟时刻 now（毫秒），返回是否仍在动。
//
// 用墙钟而不是"每帧加固定步长"：后者在掉帧时会变慢，在 120Hz 屏上会变快，
// 同一个 200ms 动画在不同机器上不一致（这也是 C15 的内容）。
func (c *Controller) Step(now time.Duration) bool {
	if !c.Animating() {
		return false
	}
	if !c.inited {
		c.start, c.last = now, now
		c.inited = true
	}
	c.last = now
	elapsed := now - c.start
	var p float32
	if c.dur <= 0 {
		p = 1
	} else {
		p = float32(elapsed) / float32(c.dur)
		if p < 0 {
			p = 0
		}
		if p > 1 {
			p = 1
		}
	}
	eased := c.curve(p)
	c.value = c.from + (c.to-c.from)*eased
	if p >= 1 {
		c.value = c.to
		c.status = AnimCompleted
	}
	if c.OnUpdate != nil {
		c.OnUpdate(c.value)
	}
	return c.Animating()
}

// Tween 是区间插值。
type Tween struct {
	From, To float32
	Curve    theme.Curve
}

// At 按进度 p（0..1）返回插值结果。
func (t Tween) At(p float32) float32 {
	if t.Curve != nil {
		p = t.Curve(p)
	}
	return t.From + (t.To-t.From)*p
}

// Lerp 是线性插值（无曲线）。
func Lerp(a, b, p float32) float32 { return a + (b-a)*p }

// Implicit 是隐式动画：目标值变了就朝新目标走一段。
//
// 用途：hover 高亮从 0 到 1、进度条从 30% 到 45%。手写这些很啰嗦，
// 而且每次都要重复"目标变了就重启动画"的判断。
type Implicit struct {
	cur    *Controller
	target float32
	dur    time.Duration
	curve  theme.Curve
}

// NewImplicit 创建隐式动画。
func NewImplicit(dur time.Duration, curve theme.Curve, initial float32) *Implicit {
	im := &Implicit{dur: dur, curve: curve, target: initial}
	im.cur = NewController(dur, curve)
	im.cur.SetValue(initial)
	return im
}

// To 设置目标值。值未变时不重启动画（避免每帧重置）。
func (im *Implicit) To(target float32, now time.Duration) *Implicit {
	if target == im.target && im.cur.Animating() {
		return im
	}
	if target == im.target && im.cur.Value() == target {
		return im
	}
	im.target = target
	im.cur.To(target, now)
	return im
}

// Step 推进。
func (im *Implicit) Step(now time.Duration) bool { return im.cur.Step(now) }

// Value 返回当前值。
func (im *Implicit) Value() float32 { return im.cur.Value() }

// Target 返回目标值。
func (im *Implicit) Target() float32 { return im.target }

// Controller 返回内部控制器。
func (im *Implicit) Controller() *Controller { return im.cur }
