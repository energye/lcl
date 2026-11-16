package diag

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
)

// ── --schedtest ─────────────────────────────────────────────────────

// SchedCase 是单条调度断言的结果。
type SchedCase struct {
	Name string
	OK   bool
	Got  string
	Note string
}

// SchedResult 是 --schedtest 的汇总。
type SchedResult struct {
	Cases   int
	Failed  int
	Elapsed time.Duration
	Details []SchedCase
}

// OK 判定是否全绿（门禁 ≥14 组）。
func (r SchedResult) OK() bool { return r.Failed == 0 && r.Cases >= 14 }

// String 返回可读摘要。
func (r SchedResult) String() string {
	return fmt.Sprintf("schedtest: %d 组，失败 %d，用时 %v", r.Cases, r.Failed, r.Elapsed)
}

// schedRig 是确定性调度环境。
type schedRig struct {
	clock  *frame.ManualClock
	frame  *frame.ManualTicker
	resize *frame.ManualTicker
	s      *frame.Scheduler

	frames   int
	presents int
	delta    time.Duration
	now      time.Duration
	want     bool
	in       time.Duration
	sizes    []string
}

// newSchedRig 造一套手动时钟 + 手动闹钟。
func newSchedRig(async bool) *schedRig {
	r := &schedRig{clock: frame.NewManualClock(0, frame.DefaultStep)}
	r.frame = frame.NewManualTicker(r.clock)
	r.resize = frame.NewManualTicker(r.clock)
	s, err := frame.NewScheduler(frame.Options{
		Clock:        r.clock,
		Ticker:       r.frame,
		ResizeTicker: r.resize,
		AsyncPresent: async,
		OnFrame: func(now, delta time.Duration) bool {
			r.frames++
			r.now, r.delta = now, delta
			r.presents++
			return true
		},
		OnResize: func(w, h int) {
			r.sizes = append(r.sizes, fmt.Sprintf("%dx%d", w, h))
		},
		WantFrame: func(now time.Duration) (bool, time.Duration) {
			return r.want, r.in
		},
	})
	if err != nil {
		panic("wg-diag: 创建调度器失败: " + err.Error())
	}
	r.s = s
	return r
}

// ── 用例 ────────────────────────────────────────────────────────────

// schedCase 是一条断言。
type schedCase struct {
	name string
	want string
	run  func() (string, error)
}

// SchedTest 跑帧调度门禁（C11/C13/C14/C15/C10）。
func SchedTest() (SchedResult, error) {
	began := time.Now()
	var res SchedResult
	for _, c := range schedCases() {
		got, err := c.run()
		d := SchedCase{Name: c.name, OK: err == nil, Got: got}
		if err != nil {
			d.Note = err.Error()
		}
		if d.Got == "" {
			d.Got = "—"
		}
		if !d.OK {
			res.Failed++
		}
		res.Cases++
		res.Details = append(res.Details, d)
	}
	res.Elapsed = time.Since(began)
	return res, nil
}

func schedCases() []schedCase {
	return []schedCase{
		{
			name: "request-only-arms",
			want: "RequestFrame 不直接出帧，只装闹钟",
			run: func() (string, error) {
				r := newSchedRig(false)
				defer r.s.Stop()
				r.s.RequestFrame()
				got := fmt.Sprintf("frames=%d armed=%v", r.frames, r.s.Armed())
				if r.frames != 0 {
					return got, fmt.Errorf("申请帧时就同步画了")
				}
				if !r.s.Armed() {
					return got, fmt.Errorf("应装填闹钟")
				}
				return got, nil
			},
		},
		{
			name: "one-present-per-frame",
			want: "一次唤醒 = 至多一帧 = 至多一次提交（C13）",
			run: func() (string, error) {
				r := newSchedRig(false)
				defer r.s.Stop()
				r.s.RequestFrame()
				r.frame.Advance(frame.DefaultStep)
				got := fmt.Sprintf("frames=%d presents=%d", r.frames, r.presents)
				if r.frames != 1 || r.presents != 1 {
					return got, fmt.Errorf("应恰好 1 帧 1 提交")
				}
				return got, nil
			},
		},
		{
			name: "single-outstanding-skip",
			want: "异步呈现：帧槽未回填时唤醒被丢弃（C11）",
			run: func() (string, error) {
				r := newSchedRig(true)
				defer r.s.Stop()
				r.s.RequestFrame()
				r.frame.Advance(frame.DefaultStep)
				if !r.s.Awaiting() {
					return "awaiting=false", fmt.Errorf("提交后帧槽应未回填")
				}
				r.s.Tick() // 模拟驱动层多给了一拍
				got := fmt.Sprintf("frames=%d skipped=%d", r.frames, r.s.Stats().Skipped)
				if r.frames != 1 {
					return got, fmt.Errorf("槽未回填不该出帧")
				}
				if r.s.Stats().Skipped != 1 {
					return got, fmt.Errorf("应计一次 skipped")
				}
				return got, nil
			},
		},
		{
			name: "present-refills-slot",
			want: "Present 回填后立即恢复出帧",
			run: func() (string, error) {
				r := newSchedRig(true)
				defer r.s.Stop()
				r.s.RequestFrame()
				r.frame.Advance(frame.DefaultStep)
				r.s.Present()
				r.s.RequestFrame()
				r.frame.Advance(0)
				got := fmt.Sprintf("frames=%d awaiting=%v", r.frames, r.s.Awaiting())
				if r.frames != 2 {
					return got, fmt.Errorf("回填后应能出第二帧")
				}
				return got, nil
			},
		},
		{
			name: "idle-zero-frames",
			want: "静止 30 拍：0 帧（C14）",
			run: func() (string, error) {
				r := newSchedRig(false)
				defer r.s.Stop()
				r.s.RequestFrame()
				r.frame.Advance(frame.DefaultStep)
				base := r.frames
				for i := 0; i < 30; i++ {
					r.frame.Advance(frame.DefaultStep)
				}
				got := fmt.Sprintf("静止期间新增 %d 帧", r.frames-base)
				if r.frames != base {
					return got, fmt.Errorf("静止期间不该出帧")
				}
				if r.s.Armed() {
					return got, fmt.Errorf("静止后应拆掉闹钟")
				}
				return got, nil
			},
		},
		{
			name: "wall-clock-delta",
			want: "帧增量取墙钟差值而不是固定步长（C15）",
			run: func() (string, error) {
				r := newSchedRig(false)
				defer r.s.Stop()
				r.s.RequestFrame()
				r.frame.Advance(0)
				r.s.RequestFrame()
				r.clock.AdvanceBy(40 * time.Millisecond)
				r.frame.Advance(0)
				got := fmt.Sprintf("delta=%v", r.delta)
				if r.delta != 40*time.Millisecond {
					return got, fmt.Errorf("应取 40ms")
				}
				return got, nil
			},
		},
		{
			name: "delta-clamped-after-sleep",
			want: "睡眠 5s：增量截到 250ms，且计一次截断",
			run: func() (string, error) {
				r := newSchedRig(false)
				defer r.s.Stop()
				r.s.RequestFrame()
				r.frame.Advance(0)
				r.s.RequestFrame()
				r.clock.AdvanceBy(5 * time.Second)
				r.frame.Advance(0)
				got := fmt.Sprintf("delta=%v clamped=%d", r.delta, r.s.Stats().Clamped)
				if r.delta != frame.DefaultMaxDelta {
					return got, fmt.Errorf("应截到 %v", frame.DefaultMaxDelta)
				}
				return got, nil
			},
		},
		{
			name: "animation-cadence",
			want: "动画需求 10ms → 闹钟装 10ms；结束后拆掉",
			run: func() (string, error) {
				r := newSchedRig(false)
				defer r.s.Stop()
				r.want, r.in = false, 10*time.Millisecond
				r.s.RequestFrame()
				r.frame.Advance(frame.DefaultStep)
				d := r.frame.Deadline() - r.clock.Now()
				if d != 10*time.Millisecond {
					return fmt.Sprintf("下次 %v", d), fmt.Errorf("应装 10ms")
				}
				r.frame.Advance(10 * time.Millisecond)
				if r.frames != 2 {
					return fmt.Sprintf("frames=%d", r.frames), fmt.Errorf("到点应出帧")
				}
				r.want, r.in = false, 0
				r.frame.Advance(10 * time.Millisecond)
				if r.frame.Armed() {
					return "armed=true", fmt.Errorf("动画结束应拆闹钟")
				}
				return "ok", nil
			},
		},
		{
			name: "requests-coalesce",
			want: "20 次申请合并成 1 帧",
			run: func() (string, error) {
				r := newSchedRig(false)
				defer r.s.Stop()
				for i := 0; i < 20; i++ {
					r.s.RequestFrame()
				}
				r.frame.Advance(frame.DefaultStep)
				got := fmt.Sprintf("frames=%d requested=%d", r.frames, r.s.Stats().Requested)
				if r.frames != 1 {
					return got, fmt.Errorf("应合并成 1 帧")
				}
				return got, nil
			},
		},
		{
			name: "resize-debounced-once",
			want: "连续 resize 只重排一次，用最后的尺寸（C10）",
			run: func() (string, error) {
				r := newSchedRig(false)
				defer r.s.Stop()
				for i := 1; i <= 10; i++ {
					r.s.RequestResize(100+i*10, 200+i*10)
					r.resize.Advance(8 * time.Millisecond)
				}
				r.resize.Advance(frame.DefaultResizeDebounce)
				got := fmt.Sprintf("fires=%d sizes=%v", len(r.sizes), r.sizes)
				if len(r.sizes) != 1 || r.sizes[0] != "200x300" {
					return got, fmt.Errorf("应只执行一次且用最后尺寸 200x300")
				}
				return got, nil
			},
		},
		{
			name: "resize-restarts-window",
			want: "新请求重置 180ms 去抖窗口",
			run: func() (string, error) {
				r := newSchedRig(false)
				defer r.s.Stop()
				r.s.RequestResize(10, 10)
				r.resize.Advance(170 * time.Millisecond)
				r.s.RequestResize(20, 20)
				r.resize.Advance(170 * time.Millisecond)
				if len(r.sizes) != 0 {
					return fmt.Sprintf("提前触发 %v", r.sizes), fmt.Errorf("窗口应被重置")
				}
				r.resize.Advance(20 * time.Millisecond)
				if len(r.sizes) != 1 {
					return fmt.Sprintf("%v", r.sizes), fmt.Errorf("应执行一次")
				}
				return "ok", nil
			},
		},
		{
			name: "resize-requests-frame",
			want: "resize 执行后立刻申请一帧（否则会看到拉伸的旧内容）",
			run: func() (string, error) {
				r := newSchedRig(false)
				defer r.s.Stop()
				r.s.RequestResize(80, 60)
				r.resize.Advance(frame.DefaultResizeDebounce)
				got := fmt.Sprintf("armed=%v", r.s.Armed())
				if !r.s.Armed() {
					return got, fmt.Errorf("resize 后应申请帧")
				}
				return got, nil
			},
		},
		{
			name: "static-scene-never-requests",
			want: "1000 节点静止：整轮只出 1 帧（脏门控 + C14）",
			run: func() (string, error) {
				root := scene.Stack().SetSize2(600, 600)
				for i := 0; i < 1000; i++ {
					root.Add(box(6, 6, geom.RGBA(20, 40, 80)).
						SetStackPosition(geom.Pt(float32(i%50)*8, float32(i/50)*8)))
				}
				root.Layout(scene.Unbounded())

				clock := frame.NewManualClock(0, frame.DefaultStep)
				ft := frame.NewManualTicker(clock)
				rt := frame.NewManualTicker(clock)
				frames := 0
				sc, err := frame.NewScheduler(frame.Options{
					Clock:        clock,
					Ticker:       ft,
					ResizeTicker: rt,
					OnFrame: func(now, delta time.Duration) bool {
						frames++
						root.ClearDirtyTree() // 帧末清脏（Container 的职责）
						return root.Count() > 0
					},
					// 真实接线：只要还有脏节点就在下一拍继续
					WantFrame: func(now time.Duration) (bool, time.Duration) {
						return root.NeedsPaint(), 0
					},
				})
				if err != nil {
					return "", err
				}
				defer sc.Stop()

				sc.RequestFrame()
				for i := 0; i < 40; i++ {
					ft.Advance(frame.DefaultStep)
				}
				got := fmt.Sprintf("frames=%d armed=%v", frames, sc.Armed())
				if frames != 1 {
					return got, fmt.Errorf("静止场景应只出首帧")
				}
				return got, nil
			},
		},
		{
			name: "animation-keeps-cadence",
			want: "动画持续 10 帧：每帧提交一次且节奏稳定",
			run: func() (string, error) {
				clock := frame.NewManualClock(0, frame.DefaultStep)
				ft := frame.NewManualTicker(clock)
				rt := frame.NewManualTicker(clock)
				left := 10
				frames := 0
				sc, err := frame.NewScheduler(frame.Options{
					Clock:        clock,
					Ticker:       ft,
					ResizeTicker: rt,
					OnFrame: func(now, delta time.Duration) bool {
						frames++
						left--
						return true
					},
					WantFrame: func(now time.Duration) (bool, time.Duration) {
						if left > 0 {
							return false, frame.DefaultStep
						}
						return false, 0
					},
				})
				if err != nil {
					return "", err
				}
				defer sc.Stop()
				sc.RequestFrame()
				for i := 0; i < 20; i++ {
					ft.Advance(frame.DefaultStep)
				}
				got := fmt.Sprintf("frames=%d", frames)
				if frames != 10 {
					return got, fmt.Errorf("应连出 10 帧后停")
				}
				if sc.Armed() {
					return got, fmt.Errorf("动画结束后应拆闹钟")
				}
				return got, nil
			},
		},
	}
}
