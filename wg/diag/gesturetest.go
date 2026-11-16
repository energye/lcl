package diag

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
)

// ── --gesturetest ───────────────────────────────────────────────────
//
// C27：单指针单胜者；败者 Reject 且不改状态。
// 六组用例覆盖"同一个按下被多个角色认领"的真实冲突（04-UI-INTERACTION.md §B5.1），
// 而不是抽象地测 Resolve 返回 bool——抽象测试全绿但表格里滑块和横滚打架，
// 是这类代码最常见的失败形态。

// GestureCase 是单条手势冲突断言的结果。
type GestureCase struct {
	Name string
	OK   bool
	Got  string
	Note string
}

// GestureResult 是 --gesturetest 的汇总。
type GestureResult struct {
	Cases   int
	Failed  int
	Elapsed time.Duration
	Details []GestureCase
}

// OK 判定是否全绿（门禁 ≥6 组）。
func (r GestureResult) OK() bool { return r.Failed == 0 && r.Cases >= 6 }

// String 返回可读摘要。
func (r GestureResult) String() string {
	return fmt.Sprintf("gesturetest: %d 组，失败 %d，用时 %v", r.Cases, r.Failed, r.Elapsed)
}

// gbase 是手势用例的基准时刻。
var gbase = time.Date(2026, 11, 16, 10, 0, 0, 0, time.UTC)

// gAt 返回基准后第 ms 毫秒。
func gAt(ms int) time.Time { return gbase.Add(time.Duration(ms) * time.Millisecond) }

// gDown / gMove / gUp 造指针事件（根空间坐标）。
func gDown(x, y float32, at time.Time) *input.PointerEvent {
	p := geom.Pt(x, y)
	return &input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer,
		Device: input.DeviceMouse, Pos: p, PosRoot: p, Time: at}
}

func gMove(x, y float32, at time.Time) *input.PointerEvent {
	e := gDown(x, y, at)
	e.Kind = input.PointerMove
	return e
}

func gUp(x, y float32, at time.Time) *input.PointerEvent {
	e := gDown(x, y, at)
	e.Kind = input.PointerUp
	return e
}

// feed 按真实顺序驱动竞技场：Down 分发给全部成员（开放期），之后每帧转发，
// Up 之后才 Sweep（先让所有成员看到抬起，再裁决）。
func feed(a *input.Arena, members []input.Recognizer, evs ...*input.PointerEvent) {
	for i, e := range evs {
		if i == 0 {
			for _, m := range members {
				m.AddPointer(e)
			}
			a.Close()
			continue
		}
		for _, m := range members {
			m.Handle(e)
		}
		if e.Kind == input.PointerUp && a.State() != input.ArenaResolved {
			a.Sweep()
		}
	}
}

// probe 是永不宣布的识别器，用来测 Sweep 的"第一个在场者胜"与重复裁决。
type probe struct {
	name   string
	accept int
	reject int
}

func (p *probe) AddPointer(*input.PointerEvent) {}
func (p *probe) Handle(*input.PointerEvent)     {}
func (p *probe) Accept()                        { p.accept++ }
func (p *probe) Reject()                        { p.reject++ }

// GestureTest 跑手势竞技场门禁。
func GestureTest() (GestureResult, error) {
	began := time.Now()
	var res GestureResult
	cases := []struct {
		name string
		run  func() (string, error)
	}{
		{"slider-beats-pan-on-horizontal", sliderVsPanHorizontal},
		{"slider-yields-on-vertical", sliderVsPanVertical},
		{"tap-beats-longpress-quick", tapVsLongPressQuick},
		{"longpress-wins-on-hold", tapVsLongPressHold},
		{"sweep-picks-first-member", sweepPicksFirst},
		{"second-declare-rejected", secondDeclareRejected},
	}
	for _, c := range cases {
		got, err := c.run()
		d := GestureCase{Name: c.name, OK: err == nil, Got: got}
		if err != nil {
			d.Note = err.Error()
			res.Failed++
		}
		res.Cases++
		res.Details = append(res.Details, d)
	}
	res.Elapsed = time.Since(began)
	return res, nil
}

// sliderVsPanHorizontal：表格里横向拖滑块 —— Slider.Drag 胜，Pan 被拒。
func sliderVsPanHorizontal() (string, error) {
	a := input.NewArena()
	cfg := input.DefaultGestureConfig()
	sliderCfg := cfg
	sliderCfg.AxisLock = input.AxisHorizontal

	var starts, moves, ends int
	slider := input.Drag(a, sliderCfg,
		func(geom.Point) { starts++ },
		func(pos, d geom.Point) { moves++ },
		func(pos, v geom.Point) { ends++ })
	pan := input.Drag(a, cfg, nil, nil, nil)
	a.AddNamed(slider, "slider")
	a.AddNamed(pan, "pan")

	feed(a, []input.Recognizer{slider, pan},
		gDown(10, 10, gAt(0)),
		gMove(24, 12, gAt(16)),
		gMove(40, 14, gAt(32)),
		gUp(40, 14, gAt(48)))

	got := fmt.Sprintf("%s 胜 / losers=%v / start=%d move=%d end=%d",
		a.WinnerName(), a.Losers(), starts, moves, ends)
	if a.WinnerName() != "slider" {
		return got, fmt.Errorf("胜者应为 slider")
	}
	if len(a.Losers()) != 1 || a.Losers()[0] != "pan" {
		return got, fmt.Errorf("pan 应被拒")
	}
	if starts != 1 || ends != 1 || moves != 1 {
		return got, fmt.Errorf("回调次数不对")
	}
	return got, nil
}

// sliderVsPanVertical：纵向拖 —— 滑块方向锁不成立，主动退出，Pan 胜。
func sliderVsPanVertical() (string, error) {
	a := input.NewArena()
	cfg := input.DefaultGestureConfig()
	sliderCfg := cfg
	sliderCfg.AxisLock = input.AxisHorizontal

	var sliderStarts int
	slider := input.Drag(a, sliderCfg, func(geom.Point) { sliderStarts++ }, nil, nil)
	pan := input.Drag(a, cfg, nil, nil, nil)
	a.AddNamed(slider, "slider")
	a.AddNamed(pan, "pan")

	feed(a, []input.Recognizer{slider, pan},
		gDown(10, 10, gAt(0)),
		gMove(12, 40, gAt(16)),
		gUp(12, 40, gAt(32)))

	got := fmt.Sprintf("%s 胜 / losers=%v / sliderStart=%d",
		a.WinnerName(), a.Losers(), sliderStarts)
	if a.WinnerName() != "pan" {
		return got, fmt.Errorf("胜者应为 pan")
	}
	if sliderStarts != 0 {
		return got, fmt.Errorf("方向锁不成立时不该启动")
	}
	// 自愿退出不是被拒：它自己已经回滚，重复 Reject 会让状态机二次回滚。
	if len(a.Losers()) != 0 {
		return got, fmt.Errorf("退出的成员不该收到 Reject")
	}
	return got, nil
}

// tapVsLongPressQuick：短按 —— Tap 胜，LongPress 被拒且不触发回调。
func tapVsLongPressQuick() (string, error) {
	a := input.NewArena()
	cfg := input.DefaultGestureConfig()

	taps, ends, longs := 0, 0, 0
	tap := input.Tap(a, cfg, func(geom.Point) { taps++ }, func(ok bool) { ends++ })
	lp := input.LongPress(a, cfg, func(geom.Point) { longs++ })
	a.AddNamed(tap, "tap")
	a.AddNamed(lp, "longpress")

	feed(a, []input.Recognizer{tap, lp},
		gDown(50, 50, gAt(0)),
		gUp(50, 50, gAt(80)))

	got := fmt.Sprintf("%s 胜 / losers=%v / taps=%d longs=%d ends=%d",
		a.WinnerName(), a.Losers(), taps, longs, ends)
	if a.WinnerName() != "tap" || taps != 1 || ends != 1 {
		return got, fmt.Errorf("短按应判为点击")
	}
	if longs != 0 {
		return got, fmt.Errorf("落败的长按不该触发")
	}
	if len(a.Losers()) != 1 || a.Losers()[0] != "longpress" {
		return got, fmt.Errorf("长按应被拒")
	}
	return got, nil
}

// tapVsLongPressHold：按住 600ms —— LongPress 胜，Tap 必须收到一次"落败"通知
// 以便回滚按下高亮（C27 的另一半）。
func tapVsLongPressHold() (string, error) {
	a := input.NewArena()
	cfg := input.DefaultGestureConfig()

	taps, longs := 0, 0
	var tapEnds []bool
	tap := input.Tap(a, cfg, func(geom.Point) { taps++ }, func(ok bool) { tapEnds = append(tapEnds, ok) })
	lp := input.LongPress(a, cfg, func(geom.Point) { longs++ })
	a.AddNamed(tap, "tap")
	a.AddNamed(lp, "longpress")

	down := gDown(50, 50, gAt(0))
	tap.AddPointer(down)
	lp.AddPointer(down)
	a.Close()

	// 长按靠 Tick 推进时长（不另开定时器，见内联注释）。
	t400 := lp.Tick(gAt(400))
	t600 := lp.Tick(gAt(600))

	got := fmt.Sprintf("%s 胜 / t400=%v t600=%v / longs=%d taps=%d tapEnds=%v",
		a.WinnerName(), t400, t600, longs, taps, tapEnds)
	if t400 {
		return got, fmt.Errorf("400ms 不该判定")
	}
	if !t600 {
		return got, fmt.Errorf("600ms 应判定长按")
	}
	if a.WinnerName() != "longpress" || longs != 1 || taps != 0 {
		return got, fmt.Errorf("长按应胜出且不触发点击")
	}
	if len(tapEnds) != 1 || tapEnds[0] {
		return got, fmt.Errorf("Tap 应收到恰好一次 accepted=false")
	}
	return got, nil
}

// sweepPicksFirst：没有任何成员宣布时，抬起后第一个仍在场者胜。
func sweepPicksFirst() (string, error) {
	a := input.NewArena()
	a1, a2 := &probe{name: "first"}, &probe{name: "second"}
	a.AddNamed(a1, a1.name)
	a.AddNamed(a2, a2.name)

	feed(a, []input.Recognizer{a1, a2},
		gDown(10, 10, gAt(0)),
		gMove(12, 12, gAt(16)),
		gUp(12, 12, gAt(32)))

	got := fmt.Sprintf("%s 胜 / accept=%d,%d reject=%d,%d",
		a.WinnerName(), a1.accept, a2.accept, a1.reject, a2.reject)
	if a.WinnerName() != "first" {
		return got, fmt.Errorf("应由第一个在场成员胜出")
	}
	if a1.accept != 1 || a2.accept != 0 || a2.reject != 1 {
		return got, fmt.Errorf("裁决通知次数不对")
	}
	return got, nil
}

// secondDeclareRejected：已裁决后第二个宣布必须失败，且胜者不被改写。
func secondDeclareRejected() (string, error) {
	a := input.NewArena()
	a1, a2 := &probe{name: "first"}, &probe{name: "second"}
	a.AddNamed(a1, a1.name)
	a.AddNamed(a2, a2.name)
	a.Close()

	first := a.Resolve(a1)
	second := a.Resolve(a2)
	resweep := a.Sweep()

	got := fmt.Sprintf("resolve1=%v resolve2=%v sweep=%s / accept=%d,%d reject=%d,%d",
		first, second, a.WinnerName(), a1.accept, a2.accept, a1.reject, a2.reject)
	if !first {
		return got, fmt.Errorf("首次宣布应成功")
	}
	if second {
		return got, fmt.Errorf("单指针只允许一个胜者")
	}
	if a.WinnerName() != "first" {
		return got, fmt.Errorf("胜者被改写")
	}
	if resweep != input.Recognizer(a1) {
		return got, fmt.Errorf("重复 sweep 返回了别的成员")
	}
	if a1.accept != 1 || a2.reject != 1 {
		return got, fmt.Errorf("重复通知：败者被拒%d次、胜者接受%d次", a2.reject, a1.accept)
	}
	return got, nil
}
