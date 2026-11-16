package diag

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
)

// ── --inputtest ─────────────────────────────────────────────────────
//
// 覆盖"消息 → 归一化 → 分发 → 组件状态"整条链路，22 项。
// 重点在 C26（捕获必有 Up 或 Cancel）与 C12（捕获自愈）：
// 这两条错了的表现是"偶尔卡在按下状态"，靠手点几乎复现不了。

// InputCase 是单条输入断言的结果。
type InputCase struct {
	Name string
	OK   bool
	Got  string
	Note string
}

// InputResult 是 --inputtest 的汇总。
type InputResult struct {
	Cases   int
	Failed  int
	Elapsed time.Duration
	Details []InputCase
}

// OK 判定是否全绿（门禁 ≥22 项）。
func (r InputResult) OK() bool { return r.Failed == 0 && r.Cases >= 22 }

// String 返回可读摘要。
func (r InputResult) String() string {
	return fmt.Sprintf("inputtest: %d 项，失败 %d，用时 %v", r.Cases, r.Failed, r.Elapsed)
}

// itBase 是输入用例的基准时刻。
var itBase = time.Date(2026, 11, 16, 10, 0, 0, 0, time.UTC)

// itAt 返回基准后第 ms 毫秒。
func itAt(ms int) time.Time { return itBase.Add(time.Duration(ms) * time.Millisecond) }

// probeState 是组件在测试里的状态记录。
type probeState struct {
	pressed  bool
	entered  bool
	leaves   int
	downs    int
	ups      int
	cancels  int
	clicks   int
	moves    int
	wheels   int
	keys     int
	cursor   scene.Cursor
	rollback int
}

// testSink 把归一化事件送进分发器，并记账。
type testSink struct {
	disp   *input.Dispatcher
	events []input.Event
	frames int
	consum int
}

func (s *testSink) PushEvent(ev input.Event) bool {
	s.events = append(s.events, ev)
	var ok bool
	switch e := ev.(type) {
	case *input.PointerEvent:
		ok = s.disp.Pointer(e)
	case *input.WheelEvent:
		ok = s.disp.Wheel(e)
	case *input.KeyEvent:
		ok = s.disp.Key(e)
	case *input.IMEEvent:
		ok = s.disp.IME(e)
	}
	if ok {
		s.consum++
	}
	return ok
}

func (s *testSink) RequestFrame() { s.frames++ }

// inputRig 是一套"归一化器 + 路由 + 分发器 + 场景树"。
type inputRig struct {
	dpi    host.DPI
	norm   *host.Normalizer
	router *host.Router
	sink   *testSink
	disp   *input.Dispatcher
	root   *scene.Node
	target *scene.Node
	st     *probeState
	fm     *input.FocusManager
}

// newRig 造一套 100×100 的画布，内含一个 80×80 的按钮节点。
func newRig(scale float32) *inputRig {
	r := &inputRig{dpi: host.NewDPI(scale), st: &probeState{}}
	btn := scene.Box(scene.BoxSpec{Fill: render.Fill(geom.RGBA(200, 200, 200))}).
		SetSize2(80, 80).
		SetStackPosition(geom.Pt(10, 10)).
		SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "btn", Cursor: scene.CursorPointer})
	r.target = btn
	r.root = scene.Stack(btn).SetSize2(100, 100)
	r.root.SetHit(scene.HitSpec{Behavior: scene.HitTranslucent, Label: "root"})
	r.root.Layout(scene.Unbounded())

	r.disp = input.NewDispatcher(r.root)
	r.sink = &testSink{disp: r.disp}
	r.norm = host.NewNormalizer(r.dpi)
	r.router = host.NewRouter(r.norm, r.sink)

	st := r.st
	r.disp.Set(btn, input.Handler{
		Pointer: func(e *input.PointerEvent) bool {
			switch e.Kind {
			case input.PointerDown:
				st.pressed = true
				st.downs++
				st.clicks = e.ClickCount
			case input.PointerUp:
				st.ups++
			case input.PointerMove:
				st.moves++
			case input.PointerCancel:
				st.cancels++
				if st.pressed {
					st.pressed = false
					st.rollback++
				}
			}
			return true
		},
		Enter: func(e *input.PointerEvent) { st.entered = true },
		Leave: func(e *input.PointerEvent) { st.leaves++; st.pressed = false },
		Wheel: func(e *input.WheelEvent) bool { st.wheels++; return true },
	})
	r.disp.OnCursorChanged(func(c scene.Cursor) { st.cursor = c })
	return r
}

// msg 造一条消息（坐标为**客户区像素**）。
func itMsg(kind uint32, x, y int32, wparam uintptr, ms int) host.Message {
	l := uintptr(uint32(uint16(x))) | uintptr(uint32(uint16(y)))<<16
	return host.Message{Msg: kind, WParam: wparam, LParam: l, Time: itAt(ms)}
}

// 消息编号（与 host 内部常量同值；测试里显式写出便于阅读）。
const (
	mMove   = 0x0200
	mLDown  = 0x0201
	mLUp    = 0x0202
	mWheel  = 0x020A
	mLeave  = 0x02A3
	mCapEnd = 0x0215
	mKeyDn  = 0x0100
	mChar   = 0x0102
)

// InputTest 跑输入链路门禁。
func InputTest() (InputResult, error) {
	began := time.Now()
	var res InputResult
	cases := []struct {
		name string
		run  func() (string, error)
	}{
		{"dip-scaling", caseDIPScaling},
		{"hover-vs-drag", caseHoverVsDrag},
		{"enter-before-move", caseEnterBeforeMove},
		{"leave-synth", caseLeaveSynth},
		{"wheel-notch", caseWheelNotch},
		{"double-click", caseDoubleClick},
		{"char-control-dropped", caseCharControlDropped},
		{"char-ctrl-modifier-dropped", caseCharCtrlDropped},
		{"insert-replace-selection", caseInsertReplace},
		{"delete-backward-cjk", caseDeleteBackwardCJK},
		{"composing-replace", caseComposingReplace},
		{"composing-commit", caseComposingCommit},
		{"ime-utf16-cursor", caseIMEUTF16},
		{"ime-cancel-on-end", caseIMECancelOnEnd},
		{"ime-result-commits", caseIMEResult},
		{"capture-cancel-on-lost", caseCaptureLost},
		{"capture-cancel-carries-button", caseCaptureCarriesButton},
		{"capture-up-releases", caseCaptureReleases},
		{"modal-swallow-cancels", caseModalSwallow},
		{"tab-navigates", caseTabNavigates},
		{"key-bubbles-from-focus", caseKeyBubbles},
		{"cursor-follows-hit", caseCursorFollows},
	}
	for _, c := range cases {
		got, err := c.run()
		d := InputCase{Name: c.name, OK: err == nil, Got: got}
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

// ── 指针归一化 ──────────────────────────────────────────────────────

func caseDIPScaling() (string, error) {
	r := newRig(1.5) // 144 DPI
	r.router.Handle(itMsg(mMove, 192, 96, 0, 0))
	last := r.sink.events[len(r.sink.events)-1].(*input.PointerEvent)
	got := fmt.Sprintf("(%d,%d)px → (%.1f,%.1f)dip", 192, 96, last.Pos.X, last.Pos.Y)
	if !approxP(last.Pos, geom.Pt(128, 64)) {
		return got, fmt.Errorf("DIP 换算错误")
	}
	return got, nil
}

func caseHoverVsDrag() (string, error) {
	r := newRig(1)
	r.router.Handle(itMsg(mMove, 20, 20, 0, 0))
	first := r.sink.events[len(r.sink.events)-1].(*input.PointerEvent)
	r.router.Handle(itMsg(mLDown, 20, 20, 0, 16))
	r.router.Handle(itMsg(mMove, 24, 20, 0x0001, 32)) // 按住左键移动
	last := r.sink.events[len(r.sink.events)-1].(*input.PointerEvent)
	got := fmt.Sprintf("无键=%s 按住=%s", first.Kind, last.Kind)
	if first.Kind != input.PointerHover {
		return got, fmt.Errorf("无按键时不该是 Move")
	}
	if last.Kind != input.PointerMove || !last.Buttons.Has(input.ButtonLeft) {
		return got, fmt.Errorf("按住时应是带左键的 Move")
	}
	return got, nil
}

func caseEnterBeforeMove() (string, error) {
	r := newRig(1)
	r.router.Handle(itMsg(mMove, 20, 20, 0, 0))
	if len(r.sink.events) != 2 {
		return fmt.Sprintf("%d 个事件", len(r.sink.events)), fmt.Errorf("首条移动应补 Enter")
	}
	k0 := r.sink.events[0].(*input.PointerEvent).Kind
	k1 := r.sink.events[1].(*input.PointerEvent).Kind
	got := fmt.Sprintf("%s,%s", k0, k1)
	if k0 != input.PointerEnter || k1 != input.PointerHover {
		return got, fmt.Errorf("顺序应为 Enter→Hover")
	}
	if !r.st.entered {
		return got, fmt.Errorf("组件未收到 Enter")
	}
	return got, nil
}

func caseLeaveSynth() (string, error) {
	r := newRig(1)
	r.router.Handle(itMsg(mMove, 20, 20, 0, 0))
	r.router.Handle(itMsg(mLeave, 20, 20, 0, 16))
	n := len(r.sink.events)
	// 重复 Leave 必须被吞：系统可能连发两条，重发会让 hover 动画重启。
	r.router.Handle(itMsg(mLeave, 20, 20, 0, 32))
	got := fmt.Sprintf("leave后事件=%d 二次leave新增=%d st.leaves=%d",
		n, len(r.sink.events)-n, r.st.leaves)
	if r.st.leaves != 1 {
		return got, fmt.Errorf("Leave 应恰好一次")
	}
	if len(r.sink.events) != n {
		return got, fmt.Errorf("重复 Leave 产生了多余事件")
	}
	return got, nil
}

func caseWheelNotch() (string, error) {
	r := newRig(1)
	r.router.Handle(itMsg(mMove, 20, 20, 0, 0))
	// 一格向上：wParam 高位是 +120
	r.router.Handle(itMsg(mWheel, 20, 20, uintptr(uint32(120)<<16), 16))
	last := r.sink.events[len(r.sink.events)-1].(*input.WheelEvent)
	got := fmt.Sprintf("Δ=%.1f", last.Delta.Y)
	if last.Delta.Y != -host.WheelStep {
		return got, fmt.Errorf("一格应为 -%v", host.WheelStep)
	}
	if r.st.wheels != 1 {
		return got, fmt.Errorf("组件未收到滚轮")
	}
	return got, nil
}

func caseDoubleClick() (string, error) {
	r := newRig(1)
	r.router.Handle(itMsg(mLDown, 20, 20, 0, 0))
	r.router.Handle(itMsg(mLUp, 20, 20, 0, 40))
	r.router.Handle(itMsg(mLDown, 21, 21, 0, 120))
	r.router.Handle(itMsg(mLUp, 21, 21, 0, 160))
	first := r.st.clicks
	// 超过双击间隔再来一次：应回到 1
	r.router.Handle(itMsg(mLDown, 21, 21, 0, 900))
	third := r.st.clicks
	got := fmt.Sprintf("二次=%d 三次=%d", first, third)
	if first != 2 {
		return got, fmt.Errorf("第二次按下应为双击")
	}
	if third != 1 {
		return got, fmt.Errorf("超间隔应重新计数")
	}
	return got, nil
}

// ── 文本编辑 ────────────────────────────────────────────────────────

func caseCharControlDropped() (string, error) {
	r := newRig(1)
	for _, ch := range []uintptr{0x0D, 0x01} {
		m := host.Message{Msg: mChar, WParam: ch, Time: itAt(0)}
		r.router.Handle(m)
	}
	st := r.norm.Stats()
	got := fmt.Sprintf("事件=%d 丢弃=%d", len(r.sink.events), st.Dropped)
	if st.Dropped != 2 || len(r.sink.events) != 0 {
		return got, fmt.Errorf("控制字符应被丢弃")
	}
	return got, nil
}

func caseCharCtrlDropped() (string, error) {
	r := newRig(1)
	r.router.SetMods(input.ModControl)
	m := host.Message{Msg: mChar, WParam: uintptr('a'), Time: itAt(0)}
	r.router.Handle(m)
	got := fmt.Sprintf("事件=%d 丢弃=%d", len(r.sink.events), r.norm.Stats().Dropped)
	if len(r.sink.events) != 0 {
		return got, fmt.Errorf("Ctrl+A 不该产生文本")
	}
	return got, nil
}

func caseInsertReplace() (string, error) {
	e := input.NewTextEditor(false)
	e.SetValue(input.EditingValue{Text: "hello world", Selection: input.RangeOf(6, 11)})
	e.InsertText("there")
	got := fmt.Sprintf("%q", e.Value().Text)
	if e.Value().Text != "hello there" {
		return got, fmt.Errorf("选区替换失败")
	}
	if e.Value().Cursor() != 11 {
		return got, fmt.Errorf("光标位置 = %d", e.Value().Cursor())
	}
	return got, nil
}

func caseDeleteBackwardCJK() (string, error) {
	e := input.NewTextEditor(false)
	e.SetValue(input.EditingValue{Text: "中文", Selection: input.RangeAt(2)})
	e.DeleteBackward()
	got := fmt.Sprintf("%q", e.Value().Text)
	if e.Value().Text != "中" {
		return got, fmt.Errorf("应按 rune 删除")
	}
	e.DeleteBackward()
	if e.Value().Text != "" {
		return got, fmt.Errorf("再删一次后 = %q", e.Value().Text)
	}
	return got, nil
}

func caseComposingReplace() (string, error) {
	e := input.NewTextEditor(false)
	e.SetValue(input.EditingValue{Text: "", Selection: input.RangeAt(0)})
	e.SetComposingRegion(0, 0, "z")
	e.SetComposingRegion(0, 1, "zh")
	e.SetComposingRegion(0, 2, "zhong")
	got := fmt.Sprintf("text=%q comp=%v", e.Value().Text, e.Value().Composing)
	if e.Value().Text != "zhong" {
		return got, fmt.Errorf("组合串被追加而不是替换")
	}
	if e.Value().ComposingText() != "zhong" {
		return got, fmt.Errorf("组合区间 = %v", e.Value().Composing)
	}
	return got, nil
}

func caseComposingCommit() (string, error) {
	e := input.NewTextEditor(false)
	e.SetValue(input.EditingValue{Text: "ab", Selection: input.RangeAt(2)})
	e.SetComposingRegion(2, 2, "中")
	e.CommitComposing()
	got := fmt.Sprintf("text=%q comp=%v commits=%d",
		e.Value().Text, e.Value().Composing, e.Stats().Commits)
	if e.Value().Text != "ab中" {
		return got, fmt.Errorf("提交后文本错误")
	}
	if !e.Value().Composing.IsEmpty() {
		return got, fmt.Errorf("提交后组合区间未清空")
	}
	return got, nil
}

// ── IME ─────────────────────────────────────────────────────────────

func caseIMEUTF16() (string, error) {
	// 组合串 "😀中"：emoji 占 2 个 UTF-16 码元，光标在 3 表示"emoji+中"之后。
	e := input.NewTextEditor(false)
	e.SetValue(input.EditingValue{Text: "", Selection: input.RangeAt(0)})
	br := host.NewIMEBridge(e)
	br.SetClock(func() time.Time { return itAt(0) })
	br.Handle(host.IMFEvent{Kind: host.IMFStartComposition})
	br.Handle(host.IMFEvent{Kind: host.IMFUpdateComposition, Text: "😀中", Cursor: 3})
	comp := br.Composition()
	got := fmt.Sprintf("text=%q cursor=%d(rune) utf16=%d",
		e.Value().Text, comp.Cursor, host.UTF16Len("😀中"))
	if e.Value().Text != "😀中" {
		return got, fmt.Errorf("组合串写入错误")
	}
	if comp.Cursor != 2 {
		return got, fmt.Errorf("UTF-16 光标 3 应换算成 rune 2")
	}
	return got, nil
}

func caseIMECancelOnEnd() (string, error) {
	e := input.NewTextEditor(false)
	e.SetValue(input.EditingValue{Text: "x", Selection: input.RangeAt(1)})
	br := host.NewIMEBridge(e)
	br.Handle(host.IMFEvent{Kind: host.IMFStartComposition})
	br.Handle(host.IMFEvent{Kind: host.IMFUpdateComposition, Text: "ni"})
	br.Handle(host.IMFEvent{Kind: host.IMFEndComposition})
	got := fmt.Sprintf("text=%q cancels=%d", e.Value().Text, br.Stats().Cancels)
	if e.Value().Text != "x" {
		return got, fmt.Errorf("取消后不该留下拼音")
	}
	if br.Active() {
		return got, fmt.Errorf("组合状态未复位")
	}
	return got, nil
}

func caseIMEResult() (string, error) {
	e := input.NewTextEditor(false)
	e.SetValue(input.EditingValue{Text: "", Selection: input.RangeAt(0)})
	br := host.NewIMEBridge(e)
	br.Handle(host.IMFEvent{Kind: host.IMFStartComposition})
	br.Handle(host.IMFEvent{Kind: host.IMFUpdateComposition, Text: "zhong"})
	br.Handle(host.IMFEvent{Kind: host.IMFResult, Text: "中"})
	got := fmt.Sprintf("text=%q commits=%d", e.Value().Text, br.Stats().Commits)
	if e.Value().Text != "中" {
		return got, fmt.Errorf("上屏结果错误（可能重复或丢拼音）")
	}
	return got, nil
}

// ── 捕获自愈（C26 / C12）───────────────────────────────────────────

func caseCaptureLost() (string, error) {
	r := newRig(1)
	r.router.Handle(itMsg(mMove, 20, 20, 0, 0))
	r.router.Handle(itMsg(mLDown, 20, 20, 0, 16))
	if r.disp.Captured(input.MousePointer) == nil {
		return "无捕获", fmt.Errorf("按下后应有捕获")
	}
	r.router.Handle(host.Message{Msg: mCapEnd, Time: itAt(32)})
	got := fmt.Sprintf("cancels=%d pressed=%v captured=%v",
		r.st.cancels, r.st.pressed, r.disp.Captured(input.MousePointer) != nil)
	if r.st.cancels != 1 {
		return got, fmt.Errorf("捕获丢失必须补 Cancel")
	}
	if r.st.pressed {
		return got, fmt.Errorf("组件未回滚按下状态")
	}
	if r.disp.Captured(input.MousePointer) != nil {
		return got, fmt.Errorf("捕获未释放")
	}
	return got, nil
}

func caseCaptureCarriesButton() (string, error) {
	r := newRig(1)
	var gotBtn input.Buttons
	r.disp.Set(r.target, input.Handler{
		Pointer: func(e *input.PointerEvent) bool {
			if e.Kind == input.PointerCancel {
				gotBtn = e.Buttons
			}
			return true
		},
	})
	r.router.Handle(itMsg(mMove, 20, 20, 0, 0))
	r.router.Handle(itMsg(mLDown, 20, 20, 0, 16))
	r.router.Handle(host.Message{Msg: mCapEnd, Time: itAt(32)})
	got := fmt.Sprintf("cancel.buttons=%v", gotBtn)
	if !gotBtn.Has(input.ButtonLeft) {
		return got, fmt.Errorf("Cancel 应带上被取消的键")
	}
	return got, nil
}

func caseCaptureReleases() (string, error) {
	r := newRig(1)
	r.router.Handle(itMsg(mMove, 20, 20, 0, 0))
	r.router.Handle(itMsg(mLDown, 20, 20, 0, 16))
	r.router.Handle(itMsg(mLUp, 20, 20, 0, 32))
	got := fmt.Sprintf("ups=%d captured=%v presses=%v",
		r.st.ups, r.disp.Captured(input.MousePointer) != nil, r.norm.Buttons())
	if r.st.ups != 1 {
		return got, fmt.Errorf("Up 未送达")
	}
	if r.disp.Captured(input.MousePointer) != nil {
		return got, fmt.Errorf("Up 后捕获未释放")
	}
	return got, nil
}

func caseModalSwallow() (string, error) {
	// 模态循环吞掉 Up：应用自己发现"捕获还在但键已松"，调 Cancel 自愈。
	r := newRig(1)
	r.router.Handle(itMsg(mMove, 20, 20, 0, 0))
	r.router.Handle(itMsg(mLDown, 20, 20, 0, 16))
	r.disp.Set(r.target, input.Handler{
		Pointer: func(e *input.PointerEvent) bool {
			if e.Kind == input.PointerCancel {
				r.st.cancels++
			}
			return true
		},
	})
	if !r.disp.Cancel(input.MousePointer, nil) {
		return "Cancel 返回 false", fmt.Errorf("自愈 Cancel 应成功")
	}
	got := fmt.Sprintf("cancels=%d captured=%v", r.st.cancels, r.disp.Captured(input.MousePointer) != nil)
	if r.st.cancels != 1 || r.disp.Captured(input.MousePointer) != nil {
		return got, fmt.Errorf("Cancel 未生效")
	}
	return got, nil
}

// ── 焦点与键盘 ──────────────────────────────────────────────────────

func caseTabNavigates() (string, error) {
	r := newRig(1)
	second := scene.Box(scene.BoxSpec{Fill: render.Fill(geom.RGBA(180, 180, 180))}).
		SetSize2(20, 20).SetStackPosition(geom.Pt(60, 60))
	r.root.Add(second)
	r.root.Layout(scene.Unbounded())

	f1, f2 := input.NewFocusNode(true), input.NewFocusNode(true)
	f1.Node, f2.Node = r.target, second
	fr := input.NewFocusNode(false)
	fr.AddChild(f1).AddChild(f2)
	fm := input.NewFocusManager(fr)
	fm.Attach(r.disp)

	var hits []string
	r.disp.Set(second, input.Handler{Key: func(e *input.KeyEvent) bool {
		hits = append(hits, "second")
		return true
	}})
	r.disp.Set(r.target, input.Handler{Key: func(e *input.KeyEvent) bool {
		hits = append(hits, "btn")
		return true
	}})

	fm.MoveNext(false) // → btn
	fm.MoveNext(false) // → second
	if r.disp.Focus() != second {
		return "焦点未同步", fmt.Errorf("Dispatcher 焦点与 FocusNode 不一致")
	}
	r.router.Handle(host.Message{Msg: mKeyDn, WParam: uintptr('A'), Time: itAt(0)})
	got := fmt.Sprintf("focus=%d hits=%v", fm.Current().Seq(), hits)
	if len(hits) != 1 || hits[0] != "second" {
		return got, fmt.Errorf("按键未投递到新焦点")
	}
	return got, nil
}

func caseKeyBubbles() (string, error) {
	r := newRig(1)
	var order []string
	r.disp.Set(r.target, input.Handler{Key: func(e *input.KeyEvent) bool {
		order = append(order, "btn")
		return false // 不消费 → 冒泡
	}})
	r.disp.Set(r.root, input.Handler{Key: func(e *input.KeyEvent) bool {
		order = append(order, "root")
		return true
	}})
	f := input.NewFocusNode(true)
	f.Node = r.target
	fm := input.NewFocusManager(f)
	fm.Attach(r.disp)
	fm.Focus(f)

	r.router.Handle(host.Message{Msg: mKeyDn, WParam: uintptr('B'), Time: itAt(0)})
	got := fmt.Sprintf("%v", order)
	if len(order) != 2 || order[0] != "btn" || order[1] != "root" {
		return got, fmt.Errorf("冒泡链错误")
	}
	return got, nil
}

func caseCursorFollows() (string, error) {
	r := newRig(1)
	r.router.Handle(itMsg(mMove, 20, 20, 0, 0)) // 命中按钮
	first := r.st.cursor
	r.router.Handle(itMsg(mMove, 2, 2, 0, 16)) // 移出按钮（命中 root）
	second := r.st.cursor
	got := fmt.Sprintf("%v → %v", first, second)
	if first != scene.CursorPointer {
		return got, fmt.Errorf("按钮上应为手型")
	}
	if second != scene.CursorDefault {
		return got, fmt.Errorf("移出后应恢复默认")
	}
	return got, nil
}

// approxP 判定两点误差 < 0.01。
func approxP(a, b geom.Point) bool {
	const eps = 0.01
	d := a.Sub(b)
	return d.X > -eps && d.X < eps && d.Y > -eps && d.Y < eps
}
