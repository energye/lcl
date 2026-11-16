// Package frame 提供帧调度原语：帧时钟、帧统计与（后续）单 outstanding 调度器。
//
// 本包不依赖渲染，也不 import lcl：它只回答"现在第几帧、该在什么时刻画"。
//
// 施工纪律（05-API-ROADMAP.md §7.1）：不许在这里引入多 outstanding、
// 短超时丢包看门狗或用 TTimer 当帧时钟——那会直接复现已修掉的输入饿死
// 与 39 FPS 问题。调度只有一个尚未应答的帧槽，槽回填才允许申请下一帧。
package frame

import (
	"sync"
	"time"
)

// ClockKind 是帧时钟来源。
type ClockKind uint8

const (
	// Wall 是墙钟（默认）：单调时钟驱动，配合 ticker 产生 vsync 相位。
	Wall ClockKind = iota
	// Compositor 由合成器/系统提供回调节奏（未来 GPU 后端）。
	Compositor
	// Manual 由调用方推进，用于无窗口自检与 golden 出帧（确定性）。
	Manual
)

// String 返回可读名。
func (k ClockKind) String() string {
	switch k {
	case Wall:
		return "wall"
	case Compositor:
		return "compositor"
	case Manual:
		return "manual"
	default:
		return "unknown"
	}
}

// FrameClock 是帧时间源。
//
// 所有实现都必须满足：Now() 单调不减；同一帧内多次调用返回同一值
// （由调度器在帧开始时取一次快照，不依赖实现保证）。
type FrameClock interface {
	// Kind 返回时钟来源。
	Kind() ClockKind
	// Now 返回自时钟起点以来的时间，用于动画求值与帧间隔统计。
	Now() time.Duration
}

// WallClock 是基于单调时钟的实现，起点为构造时刻。
type WallClock struct {
	origin time.Time
}

// NewWallClock 创建墙钟时钟。
func NewWallClock() *WallClock {
	return &WallClock{origin: time.Now()}
}

// Kind 实现 FrameClock。
func (c *WallClock) Kind() ClockKind { return Wall }

// Now 实现 FrameClock。
func (c *WallClock) Now() time.Duration { return time.Since(c.origin) }

// ManualClock 是手动推进的确定性时钟，用于 headless 自检与 golden 测试。
//
// 它保证同样的推进序列产生完全一致的帧时间，因此出图可复现。
type ManualClock struct {
	mu    sync.Mutex
	now   time.Duration
	step  time.Duration
	ticks int
}

// NewManualClock 创建手动时钟。start 为初始时刻，step 为每次 Advance 的步长；
// step <= 0 时按 60Hz（约 16.667ms）取步长。
func NewManualClock(start, step time.Duration) *ManualClock {
	if step <= 0 {
		step = DefaultStep
	}
	return &ManualClock{now: start, step: step}
}

// DefaultStep 是 60Hz 下的帧间隔（16.667ms，向上取整到微秒）。
const DefaultStep = 16667 * time.Microsecond

// Kind 实现 FrameClock。
func (c *ManualClock) Kind() ClockKind { return Manual }

// Now 实现 FrameClock。
func (c *ManualClock) Now() time.Duration {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

// Step 返回当前步长。
func (c *ManualClock) Step() time.Duration {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.step
}

// SetStep 修改步长（必须 > 0）。
func (c *ManualClock) SetStep(d time.Duration) {
	if d <= 0 {
		return
	}
	c.mu.Lock()
	c.step = d
	c.mu.Unlock()
}

// Advance 前进一个步长并返回新的当前时刻。
func (c *ManualClock) Advance() time.Duration {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.now += c.step
	c.ticks++
	return c.now
}

// AdvanceBy 前进指定时长并返回新的当前时刻；d <= 0 时不前进。
func (c *ManualClock) AdvanceBy(d time.Duration) time.Duration {
	c.mu.Lock()
	defer c.mu.Unlock()
	if d > 0 {
		c.now += d
		c.ticks++
	}
	return c.now
}

// Ticks 返回已推进次数。
func (c *ManualClock) Ticks() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ticks
}

// Reset 重置到 start 并清零计数。
func (c *ManualClock) Reset(start time.Duration) {
	c.mu.Lock()
	c.now = start
	c.ticks = 0
	c.mu.Unlock()
}
