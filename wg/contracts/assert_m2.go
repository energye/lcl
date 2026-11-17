package contracts

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/composite"
	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/raster/software"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
)

// 本文件是 M2 到期契约的断言实现。
//
// M2 的契约几乎全是"交互状态机"性质的：捕获会不会漏 Up、竞技场会不会
// 出现双胜者、焦点会不会跑到模态外面。这类断言不需要像素，只需要
// 把事件序列喂进去看状态机是否守约——所以它们能跑在无窗口自检里。

// m2Clock 是 M2 断言用的基准时刻（非零，避免被当成"没有时间戳"）。
func m2Clock(ms int) time.Time {
	return time.Date(2026, 11, 16, 10, 0, 0, 0, time.UTC).Add(time.Duration(ms) * time.Millisecond)
}

// m2Down / m2Up 造按下/抬起事件。
func m2Down(x, y float32, ms int) *input.PointerEvent {
	p := geom.Pt(x, y)
	return &input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer,
		Device: input.DeviceMouse, Pos: p, PosRoot: p, Time: m2Clock(ms)}
}

func m2Up(x, y float32, ms int) *input.PointerEvent {
	e := m2Down(x, y, ms)
	e.Kind = input.PointerUp
	return e
}

// arenaProbe 是只记账不宣布的竞技场成员。
type arenaProbe struct {
	accept int
	reject int
	events int
	state  int // 模拟组件视觉状态：1 = 按下高亮
}

func (p *arenaProbe) AddPointer(*input.PointerEvent) { p.state = 1 }
func (p *arenaProbe) Handle(*input.PointerEvent)     { p.events++ }
func (p *arenaProbe) Accept()                        { p.accept++ }
func (p *arenaProbe) Reject()                        { p.reject++; p.state = 0 }

// ── C27 单指针单胜者，败者 Reject 且不改状态 ─────────────────────

func assertSingleWinner() error {
	a := input.NewArena()
	first, second, third := &arenaProbe{}, &arenaProbe{}, &arenaProbe{}
	a.AddNamed(first, "first")
	a.AddNamed(second, "second")
	a.AddNamed(third, "third")
	a.Close()

	// 第二个成员抢先宣布：它必须是唯一胜者，其余全部被拒。
	if a.Resolve(second) != true {
		return fmt.Errorf("首次宣布未被接受")
	}
	if a.Winner() != input.Recognizer(second) {
		return fmt.Errorf("胜者不是宣布者")
	}
	if second.accept != 1 {
		return fmt.Errorf("胜者 Accept 次数 = %d", second.accept)
	}
	if second.reject != 0 {
		return fmt.Errorf("胜者被拒绝 %d 次", second.reject)
	}

	// 败者：各被拒一次，且视觉状态被回滚（state 必须归零）。
	for i, p := range []*arenaProbe{first, third} {
		if p.reject != 1 {
			return fmt.Errorf("败者 #%d Reject 次数 = %d", i, p.reject)
		}
		if p.state != 0 {
			return fmt.Errorf("败者 #%d 未回滚视觉状态", i)
		}
		if p.accept != 0 {
			return fmt.Errorf("败者 #%d 收到 Accept", i)
		}
	}

	// 再宣布一次必须失败，且胜者不被改写（单指针单胜者）。
	if a.Resolve(third) {
		return fmt.Errorf("第二个宣布者被接受，出现双胜者")
	}
	if a.Winner() != input.Recognizer(second) {
		return fmt.Errorf("重复宣布改写了胜者")
	}

	// 已裁决后再 sweep 不得重开、不得二次通知。
	if got := a.Sweep(); got != input.Recognizer(second) {
		return fmt.Errorf("sweep 返回了别的成员")
	}
	if third.reject != 1 || second.accept != 1 {
		return fmt.Errorf("重复裁决：third.reject=%d second.accept=%d", third.reject, second.accept)
	}
	return nil
}

// ── C3 ERASEBKGND 必须吃掉 ──────────────────────────────────────

func assertEraseEaten() error {
	sink := &nullSink{}
	norm := host.NewNormalizer(host.NewDPI(1))
	r := host.NewRouter(norm, sink)
	if !r.Handle(host.Message{Msg: 0x0014}) { // WM_ERASEBKGND
		return fmt.Errorf("ERASEBKGND 未被吃掉（会与整窗提交叠加，肉眼见闪烁）")
	}
	// 反向对照：WM_SIZE 必须透传，否则窗口类拿不到尺寸变化。
	if r.Handle(host.Message{Msg: 0x0005, LParam: uintptr(800) | uintptr(600)<<16}) {
		return fmt.Errorf("WM_SIZE 被吞掉，默认处理会失效")
	}
	return nil
}

// nullSink 是丢弃事件的 sink（契约断言只要路由行为）。
type nullSink struct{ frames int }

func (s *nullSink) PushEvent(input.Event) bool { return false }
func (s *nullSink) RequestFrame()              { s.frames++ }

// ── C4 静态内容只烘焙一次 ───────────────────────────────────────

func assertBakeOnce() error {
	inner := scene.Stack(
		m2Box(20, 20, geom.RGBA(255, 0, 0)).SetStackPosition(geom.Pt(0, 0)),
		m2Box(20, 20, geom.RGBA(0, 0, 255)).SetStackPosition(geom.Pt(10, 0)),
	).SetSize2(30, 20)
	root := scene.Opacity(0.5, inner)
	root.Layout(scene.Unbounded())

	c := m2Comp(60, 60)
	tr := composite.Build(root, composite.BuildOptions{Cache: true})
	if _, err := c.DrawTree(tr); err != nil {
		return fmt.Errorf("首帧: %w", err)
	}
	// 脏区为空即静止帧、不提交（C14），所以第二帧必须显式标脏才走合成。
	c.InvalidateAll()
	st, err := c.DrawTree(tr)
	if err != nil {
		return fmt.Errorf("次帧: %w", err)
	}
	// 第二帧必须命中缓存：命中说明层没有重新光栅化。
	if st.LayerCacheHits < 1 {
		return fmt.Errorf("第二帧缓存命中 %d（静态层应只烘焙一次）", st.LayerCacheHits)
	}
	return nil
}

// ── C7 光栅缓存归属唯一 ────────────────────────────────────────

func assertCacheOwnership() error {
	inner := m2Box(20, 20, geom.RGBA(255, 0, 0))
	root := scene.Opacity(0.5, inner)
	root.Layout(scene.Unbounded())

	c := m2Comp(40, 40)
	// 重新 Build 一棵树再用同一合成器画：缓存必须仍在（说明它挂在
	// 合成器上，而不是挂在被重建的层树或节点上）。
	if _, err := c.DrawTree(composite.Build(root, composite.BuildOptions{Cache: true})); err != nil {
		return err
	}
	c.InvalidateAll()
	st, err := c.DrawTree(composite.Build(root, composite.BuildOptions{Cache: true}))
	if err != nil {
		return err
	}
	if st.LayerCacheHits < 1 {
		return fmt.Errorf("层树重建后缓存未命中，缓存归属不在合成器上")
	}
	// 两个合成器不得共享同一缓存实例（否则一处失效会污染另一处）。
	other := m2Comp(40, 40)
	if c.Cache() == other.Cache() {
		return fmt.Errorf("两个合成器共享了同一缓存实例")
	}
	return nil
}

// ── C10 resize 节流+尾随 ──────────────────────────────────────

func assertResizeThrottle() error {
	clock := frame.NewManualClock(0, frame.DefaultStep)
	ft := frame.NewManualTicker(clock)
	rt := frame.NewManualTicker(clock)
	var sizes []string
	s, err := frame.NewScheduler(frame.Options{
		Clock: clock, Ticker: ft, ResizeTicker: rt,
		OnFrame:  func(now, delta time.Duration) bool { return true },
		OnResize: func(w, h int) { sizes = append(sizes, fmt.Sprintf("%dx%d", w, h)) },
		WantFrame: func(now time.Duration) (bool, time.Duration) {
			return false, 0
		},
	})
	if err != nil {
		return err
	}
	for i := 0; i < 7; i++ {
		s.RequestResize(800+i*10, 600+i*10)
		rt.Advance(20 * time.Millisecond)
	}
	rt.Advance(frame.DefaultResizeDebounce + time.Millisecond)
	// 节流前沿：首次立即；拖动中每 ~60ms 跟随一次；尾随保证最后尺寸落地。
	if len(sizes) < 2 {
		return fmt.Errorf("7 次突发 resize 只触发 %d 次（拖动中应节流跟随）", len(sizes))
	}
	if sizes[0] != "800x600" {
		return fmt.Errorf("首次尺寸 = %s（前沿应立即执行）", sizes[0])
	}
	if sizes[len(sizes)-1] != "860x660" {
		return fmt.Errorf("尾随尺寸 = %s（应取最后一次）", sizes[len(sizes)-1])
	}
	return nil
}

// ── C12 capture 自愈 ───────────────────────────────────────────

func assertCaptureHeal() error {
	root, target := m2Tree()
	d := input.NewDispatcher(root)
	state := &m2Widget{}
	d.Set(target, input.Handler{Pointer: state.handle})

	feed := func(e *input.PointerEvent) { d.Pointer(e) }
	feed(m2Down(10, 10, 0))
	if d.Captured(input.MousePointer) == nil {
		return fmt.Errorf("按下后没有捕获")
	}
	// 模态循环吞掉 Up：引擎必须能自己补 Cancel。
	if !d.Cancel(input.MousePointer, nil) {
		return fmt.Errorf("自愈 Cancel 失败")
	}
	if state.cancels != 1 || state.pressed != 0 {
		return fmt.Errorf("Cancel 未送达或状态未回滚: cancels=%d pressed=%d",
			state.cancels, state.pressed)
	}
	if d.Captured(input.MousePointer) != nil {
		return fmt.Errorf("Cancel 后捕获未释放")
	}
	return nil
}

// ── C21 释放顺序：视图 → 位图 → 池 ─────────────────────────────

func assertPoolReleaseOrder() error {
	p := host.NewPool(4, 8<<20)
	b := p.Acquire(64, 64)
	if b == nil {
		return fmt.Errorf("取缓冲失败")
	}
	// 视图还在就归池：必须被拒，否则下一帧会拿到仍在被画的缓冲。
	if p.Release(b) {
		return fmt.Errorf("有活视图时仍然归池（会串帧）")
	}
	if p.Stats().WrongOrder != 1 {
		return fmt.Errorf("顺序违规未被记录")
	}
	if !p.ReleaseView(b) {
		return fmt.Errorf("解视图失败")
	}
	if !p.Release(b) {
		return fmt.Errorf("解视图后应可归池")
	}
	if p.Leaked() != 0 {
		return fmt.Errorf("存在未解视图 %d 个", p.Leaked())
	}
	return nil
}

// ── C26 捕获指针必有 Up 或 Cancel ─────────────────────────────

func assertCaptureTerminated() error {
	// 路径一：正常抬起。
	root, target := m2Tree()
	d := input.NewDispatcher(root)
	w := &m2Widget{}
	d.Set(target, input.Handler{Pointer: w.handle})
	d.Pointer(m2Down(10, 10, 0))
	d.Pointer(m2Up(10, 10, 10))
	if d.Captured(input.MousePointer) != nil {
		return fmt.Errorf("Up 之后捕获仍存在")
	}
	if w.ups != 1 {
		return fmt.Errorf("Up 未送达")
	}

	// 路径二：捕获被系统夺走（拖出窗口/模态循环）→ 必须补 Cancel。
	root2, target2 := m2Tree()
	d2 := input.NewDispatcher(root2)
	w2 := &m2Widget{}
	d2.Set(target2, input.Handler{Pointer: w2.handle})
	d2.Pointer(m2Down(10, 10, 0))

	norm := host.NewNormalizer(host.NewDPI(1))
	sink := &pipeSink{d: d2}
	r := host.NewRouter(norm, sink)
	r.Handle(host.Message{Msg: 0x0215, Time: m2Clock(20)}) // WM_CAPTURECHANGED

	if w2.cancels != 1 {
		return fmt.Errorf("捕获丢失后收到 %d 次 Cancel（应为 1）", w2.cancels)
	}
	if w2.pressed != 0 {
		return fmt.Errorf("Cancel 后状态未回滚")
	}
	if d2.Captured(input.MousePointer) != nil {
		return fmt.Errorf("Cancel 之后捕获仍存在")
	}
	return nil
}

// pipeSink 把事件直接转给分发器。
type pipeSink struct{ d *input.Dispatcher }

func (s *pipeSink) PushEvent(e input.Event) bool {
	switch ev := e.(type) {
	case *input.PointerEvent:
		return s.d.Pointer(ev)
	case *input.WheelEvent:
		return s.d.Wheel(ev)
	case *input.KeyEvent:
		return s.d.Key(ev)
	}
	return false
}

func (s *pipeSink) RequestFrame() {}

// ── 小工具 ──────────────────────────────────────────────────────

// m2Box 造一个实心节点。
func m2Box(w, h float32, c geom.Color) *scene.Node {
	return scene.Box(scene.BoxSpec{Fill: render.Fill(c)}).SetSize2(w, h)
}

// m2Comp 造一个 headless 合成器。
func m2Comp(w, h int) *composite.Compositor {
	c, err := composite.New(composite.Options{
		Backend: software.New(software.Options{}),
		Width:   w, Height: h,
		Background: geom.RGBA(255, 255, 255),
	})
	if err != nil {
		return nil
	}
	return c
}

// m2Tree 造 根(60×60) → 目标(30×30 @10,10) 的两层树。
func m2Tree() (root, target *scene.Node) {
	target = m2Box(30, 30, geom.RGBA(200, 200, 200)).
		SetStackPosition(geom.Pt(10, 10)).
		SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "target"})
	root = scene.Stack(target).SetSize2(60, 60)
	root.SetHit(scene.HitSpec{Behavior: scene.HitTranslucent, Label: "root"})
	root.Layout(scene.Unbounded())
	return root, target
}

// m2Widget 是记录状态的假组件。
type m2Widget struct {
	pressed int
	ups     int
	cancels int
}

func (w *m2Widget) handle(e *input.PointerEvent) bool {
	switch e.Kind {
	case input.PointerDown:
		w.pressed = 1
	case input.PointerUp:
		w.ups++
	case input.PointerCancel:
		w.cancels++
		w.pressed = 0
	}
	return true
}

// ── C11 单 outstanding 帧槽 ────────────────────────────────────

func assertSingleOutstanding() error {
	clock := frame.NewManualClock(0, frame.DefaultStep)
	ft := frame.NewManualTicker(clock)
	rt := frame.NewManualTicker(clock)
	frames, skipped := 0, 0
	s, err := frame.NewScheduler(frame.Options{
		Clock: clock, Ticker: ft, ResizeTicker: rt,
		AsyncPresent: true,
		OnFrame:      func(now, delta time.Duration) bool { frames++; return true },
		WantFrame:    func(now time.Duration) (bool, time.Duration) { return false, 0 },
		OnSkip:       func() { skipped++ },
	})
	if err != nil {
		return err
	}
	defer s.Stop()

	s.RequestFrame()
	ft.Advance(frame.DefaultStep)
	if !s.Awaiting() {
		return fmt.Errorf("提交后帧槽应处于未回填状态")
	}
	// 帧槽未回填时多给一拍：必须被丢弃，不能排队等下一帧。
	s.Tick()
	if frames != 1 {
		return fmt.Errorf("帧槽未回填仍然出帧（frames=%d）", frames)
	}
	if skipped == 0 && s.Stats().Skipped == 0 {
		return fmt.Errorf("丢弃的唤醒未被记录")
	}
	// 回填后应能继续出帧。
	s.Present()
	if s.Awaiting() {
		return fmt.Errorf("Present 之后帧槽仍占用")
	}
	return nil
}

// ── C15 墙钟相位 ──────────────────────────────────────────────

func assertWallClockPhase() error {
	clock := frame.NewManualClock(0, frame.DefaultStep)
	ft := frame.NewManualTicker(clock)
	rt := frame.NewManualTicker(clock)
	var deltas []time.Duration
	s, err := frame.NewScheduler(frame.Options{
		Clock: clock, Ticker: ft, ResizeTicker: rt,
		OnFrame: func(now, delta time.Duration) bool {
			deltas = append(deltas, delta)
			return true
		},
		WantFrame: func(now time.Duration) (bool, time.Duration) { return true, 0 },
	})
	if err != nil {
		return err
	}
	defer s.Stop()

	s.RequestFrame()
	ft.Advance(frame.DefaultStep)
	// 上一帧之后过了 3 步（模拟掉帧/卡顿）：delta 必须是墙钟差，
	// 而不是"一帧 = 16ms"。
	s.RequestFrame()
	ft.Advance(3 * frame.DefaultStep)
	if len(deltas) < 2 {
		return fmt.Errorf("只出了 %d 帧", len(deltas))
	}
	if deltas[1] != 3*frame.DefaultStep {
		return fmt.Errorf("第二帧 delta = %v，应为墙钟差 %v（按帧计数会得到 %v）",
			deltas[1], 3*frame.DefaultStep, frame.DefaultStep)
	}
	return nil
}

// ── C19 Layer 必须声明 Reason 且稳态命中缓存 ───────────────────

func assertLayerReason() error {
	inner := scene.Stack(
		m2Box(20, 20, geom.RGBA(255, 0, 0)).SetStackPosition(geom.Pt(0, 0)),
		m2Box(20, 20, geom.RGBA(0, 0, 255)).SetStackPosition(geom.Pt(10, 0)),
	).SetSize2(30, 20)
	root := scene.Opacity(0.5, inner)
	root.Layout(scene.Unbounded())

	tr := composite.Build(root, composite.BuildOptions{Cache: true})
	if tr == nil {
		return fmt.Errorf("层树为空")
	}
	// 除根层外，每层都必须有非 None 的合成理由（C19 的一半）。
	bad := 0
	tr.Walk(func(l *composite.Layer) bool {
		if l.Depth() > 0 && l.Reason() == composite.ReasonNone {
			bad++
		}
		return true
	})
	if bad > 0 {
		return fmt.Errorf("%d 个层没有声明合成理由", bad)
	}

	c := m2Comp(60, 60)
	if _, err := c.DrawTree(tr); err != nil {
		return err
	}
	c.InvalidateAll()
	st, err := c.DrawTree(tr)
	if err != nil {
		return err
	}
	// 另一半：稳态必须命中缓存（每次重画的层说明理由声明形同虚设）。
	if st.LayerCacheHits < 1 {
		return fmt.Errorf("稳态未命中缓存（命中 %d，栅格化 %d）", st.LayerCacheHits, st.LayerRasters)
	}
	return nil
}
