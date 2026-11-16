package input

import (
	"math"
	"time"

	"github.com/energye/lcl/wg/geom"
)

// ── 速度估计 ────────────────────────────────────────────────────────
//
// 为什么不用"最后两点位移 / 时间差"：单次采样的抖动会被 1/dt 放大，
// 120Hz 鼠标最后两点可能是 1px/8ms → 125 DIP/s，实际上指针是匀速 300 DIP/s。
// 最小二乘把所有落在窗口里的采样一起用，抖动互相抵消。
//
// 为什么窗口是 100ms 而不是"最近 N 个点"：N 个点在低速拖动时覆盖很长时间，
// 松手前 200ms 的停顿会拖低速度；高速甩动时又覆盖太短。窗口是时间量纲，
// 与采样率无关（04-UI-INTERACTION.md §B5.3）。

const (
	// velocityWindow 是速度估计的采样窗口。
	velocityWindow = 100 * time.Millisecond
	// velocityStale 是"最后一次移动距今多久就不再算甩动"。
	// 按住不动时宿主不会送来 Move 事件，不处理的话"按住 1 秒再松手"
	// 会带着 1 秒前那次移动的速度飞出去。
	velocityStale = 50 * time.Millisecond
	// velocityRing 是环形缓冲容量（100ms 窗口下 120Hz 也不超过 16 个）。
	velocityRing = 16
	// MaxVelocity 是速度上限（DIP/s），与 Flutter kMaxFlingVelocity 同量级。
	MaxVelocity = 8000
)

// velocitySample 是一次采样。
type velocitySample struct {
	pos geom.Point
	t   time.Time
}

// VelocityTracker 估计指针速度（DIP/s）。
//
// 用法：AddPointer 时 Reset().Add(...)，每个 Move 都 Add，松手时 Velocity()。
// 零值不可用，必须用 NewVelocityTracker。
type VelocityTracker struct {
	buf  [velocityRing]velocitySample
	head int // 下一个写入位置
	n    int // 有效样本数
}

// NewVelocityTracker 创建速度跟踪器。
func NewVelocityTracker() *VelocityTracker { return &VelocityTracker{} }

// Reset 清空采样。
func (v *VelocityTracker) Reset() *VelocityTracker {
	v.head = 0
	v.n = 0
	return v
}

// Len 返回当前窗口内的采样数。
func (v *VelocityTracker) Len() int { return v.n }

// Add 记一次采样。
//
// 时间必须严格递增：同一时刻或回退的采样被忽略（合并事件、
// 宿主时间戳精度不足都会产生重复时间戳，它们对斜率没有信息量）。
// 零时间表示"宿主没提供时间"，同样忽略——用假时间算出来的速度
// 比没有速度更糟。
func (v *VelocityTracker) Add(p geom.Point, t time.Time) *VelocityTracker {
	if t.IsZero() {
		return v
	}
	if v.n > 0 && !t.After(v.newestTime()) {
		return v
	}
	v.prune(t)
	v.buf[v.head] = velocitySample{pos: p, t: t}
	v.head++
	if v.head == velocityRing {
		v.head = 0
	}
	if v.n < velocityRing {
		v.n++
	}
	return v
}

// at 返回第 i 个样本（0 = 最旧）。
func (v *VelocityTracker) at(i int) velocitySample {
	return v.buf[(v.head-v.n+i+velocityRing*2)%velocityRing]
}

func (v *VelocityTracker) newest() velocitySample { return v.at(v.n - 1) }

func (v *VelocityTracker) newestTime() time.Time { return v.newest().t }

// prune 丢弃窗口外的旧样本（相对 now，即最新样本时间）。
func (v *VelocityTracker) prune(now time.Time) {
	for v.n > 0 && now.Sub(v.at(0).t) > velocityWindow {
		v.n--
	}
}

// Velocity 返回估计速度（DIP/s）。不做陈旧检查，见 VelocityAt。
func (v *VelocityTracker) Velocity() geom.Point { return v.VelocityAt(time.Time{}) }

// VelocityAt 返回 now 时刻的速度估计（DIP/s）。
//
// now 晚于最后一次采样超过 velocityStale 时返回零：指针已经停住了，
// 只是没有 Move 事件告诉我们。now 为零表示"以最后一次采样为准"。
func (v *VelocityTracker) VelocityAt(now time.Time) geom.Point {
	if v.n < 2 {
		return geom.Point{}
	}
	newest := v.newest()
	if !now.IsZero() && now.Sub(newest.t) > velocityStale {
		return geom.Point{}
	}
	n := float64(v.n)
	var sumT, sumX, sumY float64
	for i := 0; i < v.n; i++ {
		s := v.at(i)
		t := s.t.Sub(newest.t).Seconds()
		sumT += t
		sumX += float64(s.pos.X)
		sumY += float64(s.pos.Y)
	}
	meanT, meanX, meanY := sumT/n, sumX/n, sumY/n
	var stt, stx, sty float64
	for i := 0; i < v.n; i++ {
		s := v.at(i)
		dt := s.t.Sub(newest.t).Seconds() - meanT
		stt += dt * dt
		stx += dt * (float64(s.pos.X) - meanX)
		sty += dt * (float64(s.pos.Y) - meanY)
	}
	// 采样挤在同一时刻（dt 全等）时 stt 为 0 或极小，斜率会炸成天文数字。
	if stt < 1e-7 {
		return geom.Point{}
	}
	return clampVelocity(geom.Point{X: float32(stx / stt), Y: float32(sty / stt)})
}

// clampVelocity 限制速度幅值。
func clampVelocity(v geom.Point) geom.Point {
	l := v.Len()
	if l <= MaxVelocity || l == 0 {
		return v
	}
	return v.Mul(MaxVelocity / l)
}

// ── 惯性 ────────────────────────────────────────────────────────────

const (
	// DefaultFlingDecay 是惯性的墙钟衰减系数 k（1/s）。
	DefaultFlingDecay = 5.0
	// DefaultMinFlingVelocity 是触发惯性所需的最小速度（DIP/s）。
	DefaultMinFlingVelocity = 50
	// DefaultStopVelocity 是惯性停止阈值（DIP/s）。
	DefaultStopVelocity = 20
	// DefaultMaxFlingTime 是惯性最长运行时间（防止极端参数下永不停止）。
	DefaultMaxFlingTime = 6 * time.Second
	// maxFlingStep 是单步 dt 上限。与 frame.DefaultMaxDelta 同值，
	// 但这里不 import frame：input 是叶子包，不该为常量反向依赖调度器。
	maxFlingStep = 250 * time.Millisecond
)

// Fling 是墙钟指数衰减的惯性运动。
//
// v(t) = v0·e^(-k·t)，位移按区间精确积分：
//
//	Δs = (v/k)·(1 - e^(-k·dt))
//
// 因此"一帧 16ms"与"四帧 4ms"走出的总距离完全一致——这是不用
// "每帧 ×0.95" 的原因：后者在 30FPS 下惯性距离会变成 120FPS 的一半，
// 同一手势在不同机器上手感不同（§B5.3）。
type Fling struct {
	vel   geom.Point
	k     float64
	stop  float32
	max   time.Duration
	elaps time.Duration
	alive bool
}

// NewFling 创建惯性运动。速度小于 DefaultMinFlingVelocity 时直接死掉。
func NewFling(v geom.Point) *Fling {
	f := &Fling{vel: v, k: DefaultFlingDecay, stop: DefaultStopVelocity,
		max: DefaultMaxFlingTime, alive: v.Len() >= DefaultMinFlingVelocity}
	return f
}

// SetDecay 设置衰减系数（默认 DefaultFlingDecay）。
func (f *Fling) SetDecay(k float64) *Fling {
	if k <= 0 {
		k = DefaultFlingDecay
	}
	f.k = k
	return f
}

// SetStopVelocity 设置停止阈值。
func (f *Fling) SetStopVelocity(v float32) *Fling {
	if v > 0 {
		f.stop = v
	}
	return f
}

// Alive 表示惯性是否仍在运动。
func (f *Fling) Alive() bool { return f.alive }

// Velocity 返回当前速度（已衰减）。
func (f *Fling) Velocity() geom.Point { return f.vel }

// Elapsed 返回已运行时长。
func (f *Fling) Elapsed() time.Duration { return f.elaps }

// Stop 立即停止。
func (f *Fling) Stop() { f.alive = false; f.vel = geom.Point{} }

// Step 推进 dt 并返回这段时间的位移。
//
// dt 会被夹到 DefaultMaxDelta：窗口被拖动/断点调试后 dt 可能是几秒，
// 一次走完会让内容瞬移出屏幕。
func (f *Fling) Step(dt time.Duration) geom.Point {
	if !f.alive || dt <= 0 {
		return geom.Point{}
	}
	if dt > maxFlingStep {
		dt = maxFlingStep
	}
	f.elaps += dt
	sec := dt.Seconds()
	decay := math.Exp(-f.k * sec)
	// 区间位移 = ∫v dt = v/k·(1 - e^(-k·dt))
	delta := f.vel.Mul(float32((1 - decay) / f.k))
	f.vel = f.vel.Mul(float32(decay))
	if f.vel.Len() < f.stop || f.elaps >= f.max {
		f.alive = false
	}
	return delta
}
