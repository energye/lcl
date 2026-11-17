package frame

import (
	"testing"
	"time"
)

// rig 是一套确定性调度环境：手动时钟 + 手动帧闹钟 + 手动 resize 闹钟。
type rig struct {
	clock  *ManualClock
	tick   *ManualTicker
	resize *ManualTicker
	s      *Scheduler

	frames  int
	present int
	now     time.Duration
	delta   time.Duration
	sizes   []string
	// presenting 决定 OnFrame 的返回值（false = 静止帧）
	presenting bool
	// want 是动画需求
	want bool
	in   time.Duration
}

func newRig(t *testing.T, async bool) *rig {
	t.Helper()
	r := &rig{clock: NewManualClock(0, DefaultStep), presenting: true}
	r.tick = NewManualTicker(r.clock)
	r.resize = NewManualTicker(r.clock)
	s, err := NewScheduler(Options{
		Clock:        r.clock,
		Ticker:       r.tick,
		ResizeTicker: r.resize,
		AsyncPresent: async,
		OnFrame: func(now, delta time.Duration) bool {
			r.frames++
			r.now, r.delta = now, delta
			if r.presenting {
				r.present++
			}
			return r.presenting
		},
		OnResize: func(w, h int) {
			r.sizes = append(r.sizes, itoa(w)+"x"+itoa(h))
		},
		WantFrame: func(now time.Duration) (bool, time.Duration) {
			return r.want, r.in
		},
	})
	if err != nil {
		t.Fatalf("创建调度器失败: %v", err)
	}
	r.s = s
	return r
}

func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	var b []byte
	for v > 0 {
		b = append([]byte{byte('0' + v%10)}, b...)
		v /= 10
	}
	return string(b)
}

// step 推进一个 60Hz 帧并返回是否触发了帧回调。
func (r *rig) step() bool { return r.tick.Advance(DefaultStep) }

func TestRequestFrameOnlyArms(t *testing.T) {
	r := newRig(t, false)
	defer r.s.Stop()
	r.s.RequestFrame()
	if r.frames != 0 {
		t.Fatalf("RequestFrame 不该同步出帧，实际 %d 帧", r.frames)
	}
	if !r.s.Armed() || !r.tick.Armed() {
		t.Fatal("RequestFrame 应该装填闹钟")
	}
	if !r.step() {
		t.Fatal("推进后应触发一次帧")
	}
	if r.frames != 1 || r.present != 1 {
		t.Fatalf("应出 1 帧并提交: frames=%d present=%d", r.frames, r.present)
	}
}

func TestFirstFrameDeltaIsZero(t *testing.T) {
	r := newRig(t, false)
	defer r.s.Stop()
	r.s.RequestFrame()
	r.step()
	if r.delta != 0 {
		t.Fatalf("首帧 delta 应为 0（没有上一帧），实际 %v", r.delta)
	}
	r.s.RequestFrame()
	r.step()
	if r.delta != DefaultStep {
		t.Fatalf("第二帧 delta 应为 %v，实际 %v", DefaultStep, r.delta)
	}
}

func TestWallClockPhaseNotFrameCount(t *testing.T) {
	// C15：帧增量必须来自墙钟。这里让"时钟"比闹钟走得更快，
	// 如果实现按固定步长积分，delta 会恒等于 16.667ms。
	r := newRig(t, false)
	defer r.s.Stop()

	r.s.RequestFrame()
	r.tick.Advance(0) // 首帧，delta=0

	// 时钟先跳 40ms，再触发闹钟
	r.s.RequestFrame()
	r.clock.AdvanceBy(40 * time.Millisecond)
	r.tick.Advance(0)
	if r.delta != 40*time.Millisecond {
		t.Fatalf("帧增量应取墙钟差值 40ms，实际 %v", r.delta)
	}
}

func TestDeltaClampedAfterSleep(t *testing.T) {
	r := newRig(t, false)
	defer r.s.Stop()
	r.s.RequestFrame()
	r.tick.Advance(0)

	r.s.RequestFrame()
	r.clock.AdvanceBy(5 * time.Second)
	r.tick.Advance(0)
	if r.delta != DefaultMaxDelta {
		t.Fatalf("睡眠 5s 后增量应被截到 %v，实际 %v", DefaultMaxDelta, r.delta)
	}
	if got := r.s.Stats().Clamped; got != 1 {
		t.Fatalf("截断次数 = %d，应为 1", got)
	}
}

func TestSingleOutstandingSkipsTick(t *testing.T) {
	r := newRig(t, true) // 异步呈现
	defer r.s.Stop()

	r.s.RequestFrame()
	r.step()
	if r.frames != 1 {
		t.Fatalf("应出 1 帧，实际 %d", r.frames)
	}
	if !r.s.Awaiting() {
		t.Fatal("异步模式下提交后帧槽应处于未回填状态（C11）")
	}

	// 槽未回填：即使装填了闹钟，唤醒也必须被丢弃
	r.s.RequestFrame()
	if r.tick.Armed() {
		t.Fatal("槽未回填时不该装填闹钟")
	}
	r.tick.Advance(DefaultStep)
	if r.frames != 1 {
		t.Fatalf("槽未回填时不该出帧，实际 %d", r.frames)
	}
	if got := r.s.Stats().Skipped; got != 0 {
		// 这里没装闹钟所以走不到 Tick；跳过计数应由"直接 Tick"体现
		t.Fatalf("未装填时不应有 skipped，实际 %d", got)
	}

	// 直接 Tick 模拟驱动的唤醒
	r.s.Tick()
	if got := r.s.Stats().Skipped; got != 1 {
		t.Fatalf("直接唤醒在槽未回填时应计一次 skipped，实际 %d", got)
	}
	if r.frames != 1 {
		t.Fatalf("跳过的唤醒不该执行帧，实际 %d", r.frames)
	}

	// 回填后恢复
	r.s.Present()
	if r.s.Awaiting() {
		t.Fatal("Present 应回填帧槽")
	}
	if !r.step() {
		t.Fatal("回填后应能继续出帧")
	}
	if r.frames != 2 {
		t.Fatalf("回填后应出第 2 帧，实际 %d", r.frames)
	}
}

func TestIdleDisarmsTicker(t *testing.T) {
	// C14：没有挂起工作、没有动画 → 不装闹钟，推进时间一帧都不出。
	r := newRig(t, false)
	defer r.s.Stop()

	r.s.RequestFrame()
	r.step()
	if r.tick.Armed() {
		t.Fatal("静止后不该继续装填闹钟（静止零帧）")
	}
	for i := 0; i < 10; i++ {
		r.step()
	}
	if r.frames != 1 {
		t.Fatalf("静止期间不该出帧，实际 %d 帧", r.frames)
	}
}

func TestAnimationDrivesFrames(t *testing.T) {
	r := newRig(t, false)
	defer r.s.Stop()

	// 动画需要 10ms 后的一帧
	r.want, r.in = false, 10*time.Millisecond
	r.s.RequestFrame()
	r.step()
	if !r.tick.Armed() {
		t.Fatal("有动画需求时应装填闹钟")
	}
	if got := r.tick.Deadline() - r.clock.Now(); got != 10*time.Millisecond {
		t.Fatalf("下一次帧应在 10ms 后，实际 %v", got)
	}

	// 到点触发
	r.tick.Advance(10 * time.Millisecond)
	if r.frames != 2 {
		t.Fatalf("到点应出帧，实际 %d", r.frames)
	}

	// 动画结束：want=false, in=0 → 拆闹钟
	r.want, r.in = false, 0
	r.step()
	if r.tick.Armed() {
		t.Fatal("动画结束后应拆掉闹钟")
	}
}

func TestRequestsCoalesce(t *testing.T) {
	r := newRig(t, false)
	defer r.s.Stop()
	for i := 0; i < 20; i++ {
		r.s.RequestFrame()
	}
	r.step()
	if r.frames != 1 {
		t.Fatalf("20 次申请应合并成 1 帧，实际 %d", r.frames)
	}
	if got := r.s.Stats().Requested; got != 20 {
		t.Fatalf("申请次数统计错误: %d", got)
	}
}

func TestFrameRequestedInsideFrame(t *testing.T) {
	// 动画在帧回调里继续申请：下一帧要紧接着来（不等外部再推）。
	clock := NewManualClock(0, DefaultStep)
	tk := NewManualTicker(clock)
	n := 0
	var sc *Scheduler
	sc, err := NewScheduler(Options{
		Clock:  clock,
		Ticker: tk,
		OnFrame: func(now, delta time.Duration) bool {
			n++
			if n < 3 {
				sc.RequestFrame() // 动画继续
			}
			return true
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer sc.Stop()
	sc.RequestFrame()
	for i := 0; i < 5; i++ {
		tk.Advance(DefaultStep)
	}
	if n != 3 {
		t.Fatalf("应连出 3 帧后停下，实际 %d", n)
	}
	if tk.Armed() {
		t.Fatal("动画结束后应拆掉闹钟")
	}
}

func TestResizeThrottledDuringDrag(t *testing.T) {
	r := newRig(t, false)
	defer r.s.Stop()

	// 连续 10 次 resize，间隔 8ms（模拟拖动窗口边缘）。旧纯尾随语义下
	// 拖动全程一次都不触发（每次请求都重置闹钟），backing 停在旧尺寸、
	// 提交全程被拒——节流前沿必须让拖动中每 ~60ms 真正执行一次。
	for i := 1; i <= 10; i++ {
		r.s.RequestResize(100+i*10, 200+i*10)
		r.resize.Advance(8 * time.Millisecond)
	}
	if len(r.sizes) < 2 {
		t.Fatalf("拖动期间应有节流触发的 resize（≥2 次），实际 %v", r.sizes)
	}
	if got := r.s.Stats().Resizes; got != 10 {
		t.Fatalf("请求数 = %d，应为 10", got)
	}
	// 拖拽停下：尾随补偿保证最后一次尺寸落地。
	r.resize.Advance(DefaultResizeDebounce)
	if last := r.sizes[len(r.sizes)-1]; last != "200x300" {
		t.Fatalf("最后一次应使用 200x300，实际 %v", r.sizes)
	}
	// resize 后应立刻申请一帧
	if !r.tick.Armed() {
		t.Fatal("resize 后应申请一帧")
	}
}

func TestResizeTrailingPinnedToLastFire(t *testing.T) {
	r := newRig(t, false)
	defer r.s.Stop()

	r.s.RequestResize(10, 10) // 首次请求：前沿立即执行
	if len(r.sizes) != 1 || r.sizes[0] != "10x10" {
		t.Fatalf("首次 resize 应立即执行，实际 %v", r.sizes)
	}
	r.resize.Advance(50 * time.Millisecond)
	r.s.RequestResize(20, 20) // 窗口内：尾随钉在 lastFire+60ms，而不是"从现在起 60ms"
	r.resize.Advance(50 * time.Millisecond)
	if len(r.sizes) != 2 || r.sizes[1] != "20x20" {
		t.Fatalf("尾随应在 lastFire+60ms 触发 20x20，实际 %v", r.sizes)
	}
	r.resize.Advance(DefaultResizeDebounce)
	if len(r.sizes) != 2 {
		t.Fatalf("尾随只该执行一次，实际 %v", r.sizes)
	}
}

func TestResizePendingTracksRequestLifecycle(t *testing.T) {
	r := newRig(t, false)
	defer r.s.Stop()

	if r.s.ResizePending() {
		t.Fatal("未请求时不该挂起")
	}
	r.s.RequestResize(10, 10) // 前沿立即执行
	if r.s.ResizePending() {
		t.Fatal("前沿执行后不该挂起")
	}
	r.s.RequestResize(20, 20) // 窗口内，等尾随
	if !r.s.ResizePending() {
		t.Fatal("尾随期间应挂起")
	}
	r.resize.Advance(DefaultResizeDebounce)
	if r.s.ResizePending() {
		t.Fatal("尾随执行后不该挂起")
	}
}

func TestStopSilencesEverything(t *testing.T) {
	r := newRig(t, false)
	r.s.RequestFrame()
	r.s.Stop()
	if r.step() {
		t.Fatal("Stop 后不该再触发帧")
	}
	r.s.RequestFrame()
	if r.tick.Armed() {
		t.Fatal("Stop 后不该装填闹钟")
	}
}

func TestMissingDepsRejected(t *testing.T) {
	if _, err := NewScheduler(Options{}); err == nil {
		t.Fatal("缺少时钟应报错")
	}
	if _, err := NewScheduler(Options{Clock: NewWallClock()}); err == nil {
		t.Fatal("缺少闹钟应报错")
	}
}
