package input

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/geom"
)

// t0 是测试基准时刻（非零，避免被当成"没有时间戳"）。
var t0 = time.Date(2026, 11, 16, 10, 0, 0, 0, time.UTC)

// ms 返回相对基准的偏移时刻。
func at(ms int) time.Time { return t0.Add(time.Duration(ms) * time.Millisecond) }

// ── 速度 ────────────────────────────────────────────────────────────

func TestVelocityLinearMotion(t *testing.T) {
	// 每 10ms 右移 3 DIP → 300 DIP/s。
	v := NewVelocityTracker()
	for i := 0; i <= 10; i++ {
		v.Add(geom.Pt(float32(i*3), 0), at(i*10))
	}
	got := v.Velocity()
	if !approxPt(got, geom.Pt(300, 0)) {
		t.Fatalf("速度 = %v，期望 (300,0)", got)
	}
}

func TestVelocityAveragesJitter(t *testing.T) {
	// 匀速 200 DIP/s 上叠加 ±1 DIP 抖动：单点差分会被放大，最小二乘应稳住。
	off := []float32{0, 1, -1, 1, -1, 0}
	v := NewVelocityTracker()
	for i := 0; i < 6; i++ {
		v.Add(geom.Pt(2*float32(i)+off[i], 0), at(i*10))
	}
	got := v.Velocity().X
	if got < 150 || got > 250 {
		t.Fatalf("速度 X = %v，期望在 200 附近", got)
	}
}

func TestVelocityWindowDropsOldSamples(t *testing.T) {
	// 先慢慢走，再快速走：窗口只保留最近 100ms，速度应接近快速段。
	v := NewVelocityTracker()
	for i := 0; i <= 10; i++ { // 0–1000ms，慢速 10 DIP/s
		v.Add(geom.Pt(float32(i), 0), at(i*100))
	}
	for i := 1; i <= 5; i++ { // 1000ms 起，快速 500 DIP/s
		v.Add(geom.Pt(float32(10+i*5), 0), at(1000+i*10))
	}
	if v.Len() > velocityRing {
		t.Fatalf("环形缓冲溢出: %d", v.Len())
	}
	got := v.Velocity().X
	if got < 400 {
		t.Fatalf("速度 X = %v，慢速样本没被窗口剔除", got)
	}
}

func TestVelocityZeroWhenStale(t *testing.T) {
	// 甩动后按住不动 200ms 再松手：不应带着旧速度飞出去。
	v := NewVelocityTracker()
	for i := 0; i <= 5; i++ {
		v.Add(geom.Pt(float32(i*4), 0), at(i*10))
	}
	if got := v.VelocityAt(at(300)); got.Len() != 0 {
		t.Fatalf("陈旧速度 = %v，期望 0", got)
	}
	if got := v.VelocityAt(at(55)); got.Len() == 0 {
		t.Fatalf("50ms 内不应判陈旧")
	}
}

func TestVelocityClamped(t *testing.T) {
	v := NewVelocityTracker()
	v.Add(geom.Pt(0, 0), at(0))
	v.Add(geom.Pt(5000, 0), at(1)) // 1ms 走 5000 DIP → 5e6 DIP/s
	if got := v.Velocity(); got.X > MaxVelocity+1 {
		t.Fatalf("速度未夹到上限: %v", got)
	}
}

func TestVelocityIgnoresNonIncreasingTime(t *testing.T) {
	v := NewVelocityTracker()
	v.Add(geom.Pt(0, 0), at(0))
	v.Add(geom.Pt(10, 0), at(0)) // 同刻
	v.Add(geom.Pt(20, 0), time.Time{})
	if v.Len() != 1 {
		t.Fatalf("应只保留 1 个采样，得到 %d", v.Len())
	}
	if got := v.Velocity(); got.Len() != 0 {
		t.Fatalf("单采样速度应为 0，得到 %v", got)
	}
}

// ── 惯性 ────────────────────────────────────────────────────────────

// runFling 以固定步长推进直到停止，返回总位移与用时。
func runFling(f *Fling, step time.Duration, limit time.Duration) (geom.Point, time.Duration) {
	total := geom.Point{}
	var used time.Duration
	for f.Alive() && used < limit {
		total = total.Add(f.Step(step))
		used += step
	}
	return total, used
}

func TestFlingFrameRateIndependent(t *testing.T) {
	// 同一初速、同一段墙钟时长，16ms 一帧与 4ms 一帧走出的距离必须一致
	// （§B5.3 的核心：距离只取决于时间，不取决于帧数）。
	const span = 400 * time.Millisecond
	coarse := NewFling(geom.Pt(1200, 0))
	fine := NewFling(geom.Pt(1200, 0))
	dc, df := geom.Point{}, geom.Point{}
	for i := 0; i < int(span/(16*time.Millisecond)); i++ {
		dc = dc.Add(coarse.Step(16 * time.Millisecond))
	}
	for i := 0; i < int(span/(4*time.Millisecond)); i++ {
		df = df.Add(fine.Step(4 * time.Millisecond))
	}
	if !approxPt(dc, df) {
		t.Fatalf("帧率影响了惯性距离: 16ms=%v 4ms=%v", dc, df)
	}
	if dc.Len() <= 0 {
		t.Fatal("惯性没有位移")
	}
}

func TestFlingDecaysMonotonically(t *testing.T) {
	f := NewFling(geom.Pt(0, 1000))
	prev := f.Velocity().Len()
	for i := 0; i < 10; i++ {
		f.Step(16 * time.Millisecond)
		v := f.Velocity().Len()
		if v > prev {
			t.Fatalf("速度反弹: %v → %v", prev, v)
		}
		prev = v
	}
}

func TestFlingStops(t *testing.T) {
	f := NewFling(geom.Pt(0, 200))
	_, used := runFling(f, 16*time.Millisecond, 5*time.Second)
	if f.Alive() {
		t.Fatal("惯性未停止")
	}
	if used > 3*time.Second {
		t.Fatalf("停止太慢: %v", used)
	}
}

func TestFlingBelowThresholdNeverStarts(t *testing.T) {
	f := NewFling(geom.Pt(30, 0)) // < DefaultMinFlingVelocity
	if f.Alive() {
		t.Fatal("低于阈值的速度不应触发惯性")
	}
	if d := f.Step(16 * time.Millisecond); d.Len() != 0 {
		t.Fatalf("死掉的惯性不该产生位移: %v", d)
	}
}

func TestFlingClampsHugeDelta(t *testing.T) {
	f := NewFling(geom.Pt(5000, 0))
	d := f.Step(10 * time.Second) // 断点调试后的大 dt
	maxDist := float32(5000) * float32(maxFlingStep.Seconds())
	if d.Len() > maxDist+1 {
		t.Fatalf("单步位移未夹到上限: %v", d.Len())
	}
}

// ── 竞技场 ──────────────────────────────────────────────────────────

// pDown / pMove / pUp 造指针事件。
func pDown(x, y float32, at time.Time) *PointerEvent {
	p := geom.Pt(x, y)
	return &PointerEvent{Kind: PointerDown, Pointer: MousePointer, Device: DeviceMouse,
		Pos: p, PosRoot: p, Time: at}
}

func pMove(x, y float32, at time.Time) *PointerEvent {
	e := pDown(x, y, at)
	e.Kind = PointerMove
	return e
}

func pUp(x, y float32, at time.Time) *PointerEvent {
	e := pDown(x, y, at)
	e.Kind = PointerUp
	return e
}

// drive 把事件依次转发给所有成员；抬起后若无人宣布则 sweep（模拟真实顺序：
// 先让所有成员看 Up，再裁决）。
func drive(a *Arena, members []Recognizer, evs ...*PointerEvent) {
	for i, e := range evs {
		if i == 0 {
			for _, m := range members {
				m.AddPointer(e)
			}
			continue
		}
		for _, m := range members {
			m.Handle(e)
		}
		if e.Kind == PointerUp && a.State() != ArenaResolved {
			a.Sweep()
		}
	}
}

// tab 记录 Tap 的结果。
type tapLog struct {
	taps int
	end  []bool
}

func TestArenaSliderBeatsPanOnHorizontalDrag(t *testing.T) {
	a := NewArena()
	cfg := DefaultGestureConfig()
	sliderCfg := cfg
	sliderCfg.AxisLock = AxisHorizontal

	var started int
	slider := Drag(a, sliderCfg,
		func(geom.Point) { started++ },
		func(pos, d geom.Point) {},
		func(pos, v geom.Point) {})
	pan := Drag(a, cfg, nil, nil, nil)
	a.AddNamed(slider, "slider")
	a.AddNamed(pan, "pan")

	drive(a, []Recognizer{slider, pan},
		pDown(10, 10, at(0)),
		pMove(24, 12, at(16)),
		pUp(24, 12, at(32)))

	if a.WinnerName() != "slider" {
		t.Fatalf("胜者 = %q，losers = %v", a.WinnerName(), a.Losers())
	}
	if started != 1 {
		t.Fatalf("拖动开始回调 %d 次", started)
	}
	if len(a.Losers()) != 1 || a.Losers()[0] != "pan" {
		t.Fatalf("losers = %v", a.Losers())
	}
}

func TestArenaSliderWithdrawsOnVerticalDrag(t *testing.T) {
	a := NewArena()
	cfg := DefaultGestureConfig()
	sc := cfg
	sc.AxisLock = AxisHorizontal

	var sliderStarts int
	slider := Drag(a, sc, func(geom.Point) { sliderStarts++ }, nil, nil)
	pan := Drag(a, cfg, nil, nil, nil)
	a.AddNamed(slider, "slider")
	a.AddNamed(pan, "pan")

	drive(a, []Recognizer{slider, pan},
		pDown(10, 10, at(0)),
		pMove(12, 40, at(16)),
		pUp(12, 40, at(32)))

	if a.WinnerName() != "pan" {
		t.Fatalf("胜者 = %q", a.WinnerName())
	}
	if sliderStarts != 0 {
		t.Fatal("方向锁不成立的识别器不该启动")
	}
	// 自愿退出不算被拒：它自己已经回滚过，不该再收到 Reject。
	if len(a.Losers()) != 0 {
		t.Fatalf("退出的成员不该出现在 losers 里: %v", a.Losers())
	}
}

func TestArenaTapBeatsLongPressOnQuickRelease(t *testing.T) {
	a := NewArena()
	cfg := DefaultGestureConfig()
	tl := &tapLog{}
	tap := Tap(a, cfg, func(geom.Point) { tl.taps++ }, func(ok bool) { tl.end = append(tl.end, ok) })
	lp := LongPress(a, cfg, func(geom.Point) { t.Fatal("不该长按") })
	a.AddNamed(tap, "tap")
	a.AddNamed(lp, "longpress")

	drive(a, []Recognizer{tap, lp},
		pDown(50, 50, at(0)),
		pUp(50, 50, at(80)))

	if a.WinnerName() != "tap" || tl.taps != 1 {
		t.Fatalf("胜者 = %q taps = %d", a.WinnerName(), tl.taps)
	}
	if len(a.Losers()) != 1 || a.Losers()[0] != "longpress" {
		t.Fatalf("losers = %v", a.Losers())
	}
}

func TestArenaLangPressWinsAfterDelay(t *testing.T) {
	a := NewArena()
	cfg := DefaultGestureConfig()
	tl := &tapLog{}
	var longs int
	tap := Tap(a, cfg, func(geom.Point) { tl.taps++ }, func(ok bool) { tl.end = append(tl.end, ok) })
	lp := LongPress(a, cfg, func(geom.Point) { longs++ })
	a.AddNamed(tap, "tap")
	a.AddNamed(lp, "longpress")

	tap.AddPointer(pDown(50, 50, at(0)))
	lp.AddPointer(pDown(50, 50, at(0)))
	a.Close()

	if lp.(*longPressRec).Tick(at(400)) {
		t.Fatal("400ms 时不该判定长按")
	}
	if !lp.(*longPressRec).Tick(at(600)) {
		t.Fatal("600ms 时应判定长按")
	}
	if a.WinnerName() != "longpress" || longs != 1 {
		t.Fatalf("胜者 = %q longs = %d", a.WinnerName(), longs)
	}
	// Tap 落败：必须收到一次 accepted=false 以便回滚"按下高亮"。
	if len(tl.end) != 1 || tl.end[0] {
		t.Fatalf("Tap 落败通知 = %v", tl.end)
	}
	if tl.taps != 0 {
		t.Fatal("落败的 Tap 不该触发 onTap")
	}
}

// inertRec 是永不宣布的识别器（模拟"只看不抢"的观察者）。
type inertRec struct {
	accept, reject int
	handled        int
}

func (r *inertRec) AddPointer(*PointerEvent) {}
func (r *inertRec) Handle(*PointerEvent)     { r.handled++ }
func (r *inertRec) Accept()                  { r.accept++ }
func (r *inertRec) Reject()                  { r.reject++ }

func TestArenaFirstMemberWinsWhenNobodyDeclares(t *testing.T) {
	a := NewArena()
	a1, a2 := &inertRec{}, &inertRec{}
	a.AddNamed(a1, "first")
	a.AddNamed(a2, "second")
	a.Close()

	if got := a.Sweep(); got != Recognizer(a1) {
		t.Fatalf("sweep 胜者 = %v", got)
	}
	if a.WinnerName() != "first" {
		t.Fatalf("胜者名 = %q", a.WinnerName())
	}
	if a1.accept != 1 || a2.reject != 1 || a2.accept != 0 {
		t.Fatalf("accept/reject = %d/%d, 败者 accept = %d", a1.accept, a2.reject, a2.accept)
	}
}

func TestArenaSecondResolveFails(t *testing.T) {
	a := NewArena()
	a1, a2 := &inertRec{}, &inertRec{}
	a.AddNamed(a1, "first")
	a.AddNamed(a2, "second")
	a.Close()

	if !a.Resolve(a1) {
		t.Fatal("第一次宣布应成功")
	}
	if a.Resolve(a2) {
		t.Fatal("单指针只允许一个胜者")
	}
	if a2.reject != 1 {
		t.Fatalf("败者应恰好被拒一次，得到 %d", a2.reject)
	}
	if a.WinnerName() != "first" {
		t.Fatalf("胜者被改写: %q", a.WinnerName())
	}
	// 已裁决后再 sweep 不得重开
	if got := a.Sweep(); got != Recognizer(a1) {
		t.Fatalf("sweep 返回 %v", got)
	}
	if a1.accept != 1 || a2.reject != 1 {
		t.Fatalf("重复裁决: accept=%d reject=%d", a1.accept, a2.reject)
	}
}

func TestArenaAddRejectedWhenClosed(t *testing.T) {
	a := NewArena()
	a.Close()
	if a.Add(&inertRec{}) {
		t.Fatal("封闭后不该接受新成员")
	}
	if a.Members() != 0 {
		t.Fatalf("成员数 = %d", a.Members())
	}
}

func TestArenaRejectAllOnCancel(t *testing.T) {
	a := NewArena()
	a1, a2 := &inertRec{}, &inertRec{}
	a.Add(a1)
	a.Add(a2)
	a.Close()
	a.RejectAll()
	if a1.reject != 1 || a2.reject != 1 {
		t.Fatalf("reject = %d/%d", a1.reject, a2.reject)
	}
	if a.Winner() != nil {
		t.Fatal("取消后不该有胜者")
	}
	if st := a.Stats(); st.Rejected != 2 || st.Added != 2 {
		t.Fatalf("统计 = %+v", st)
	}
}

func TestDragEndCarriesVelocity(t *testing.T) {
	a := NewArena()
	cfg := DefaultGestureConfig()
	var endVel geom.Point
	got := false
	d := Drag(a, cfg, nil, nil, func(pos, v geom.Point) { endVel, got = v, true })
	a.AddNamed(d, "drag")

	d.AddPointer(pDown(0, 0, at(0)))
	a.Close()
	for i := 1; i <= 5; i++ {
		d.Handle(pMove(float32(i*10), 0, at(i*10)))
	}
	d.Handle(pUp(50, 0, at(60)))

	if !got {
		t.Fatal("拖动结束未回调")
	}
	if endVel.X < 500 {
		t.Fatalf("结束速度 = %v，期望接近 1000 DIP/s", endVel)
	}
}
