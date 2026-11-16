package widgets

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

type controlRig struct {
	ctx  *ui.BuildContext
	root *scene.Node
	disp *input.Dispatcher
	fm   *input.FocusManager
	ime  *host.IMEBridge
}

func newControlRig() *controlRig {
	root := scene.Stack().SetSize2(400, 300)
	disp := input.NewDispatcher(root)
	fm := input.NewFocusManager(input.NewFocusNode(false))
	fm.Attach(disp)
	ime := host.NewIMEBridge(nil)
	r := &controlRig{
		root: root, disp: disp, fm: fm, ime: ime,
		ctx: &ui.BuildContext{
			Theme: theme.Default(), Focus: fm, Input: disp,
			SetTextEditor: ime.SetEditor,
		},
	}
	return r
}

func (r *controlRig) mount(w ui.Widget) *scene.Node {
	n := w.Build(r.ctx)
	r.root.Add(n)
	r.root.Layout(scene.TightConstraints(geom.Size{W: 400, H: 300}))
	return n
}

func (r *controlRig) click(x, y float32) {
	p := geom.Pt(x, y)
	t := time.Now()
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer,
		Pos: p, PosRoot: p, Time: t})
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerUp, Pointer: input.MousePointer,
		Pos: p, PosRoot: p, Time: t.Add(16 * time.Millisecond)})
}

func (r *controlRig) drag(x0, x1 float32) {
	t := time.Now()
	p0 := geom.Pt(x0, 2)
	p1 := geom.Pt(x1, 2)
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer,
		Pos: p0, PosRoot: p0, Time: t})
	pm := geom.Pt((x0+x1)/2, 2)
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerMove, Pointer: input.MousePointer,
		Buttons: input.ButtonLeft, Pos: pm, PosRoot: pm, Time: t.Add(16 * time.Millisecond)})
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerMove, Pointer: input.MousePointer,
		Buttons: input.ButtonLeft, Pos: p1, PosRoot: p1, Time: t.Add(32 * time.Millisecond)})
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerUp, Pointer: input.MousePointer,
		Pos: p1, PosRoot: p1, Time: t.Add(48 * time.Millisecond)})
}

func TestButtonClickAndKeyboard(t *testing.T) {
	r := newControlRig()
	btn := NewButton("确定")
	clicks := 0
	btn.OnClick = func() { clicks++ }
	r.mount(btn)
	r.click(20, 10)
	if clicks != 1 {
		t.Fatalf("点击次数 = %d", clicks)
	}
	r.disp.Key(&input.KeyEvent{Code: input.CodeEnter})
	if clicks != 1 {
		t.Fatal("未聚焦的按钮不该响应键盘")
	}
	if btn.focus != nil {
		r.fm.Focus(btn.focus)
	}
	r.disp.Key(&input.KeyEvent{Code: input.CodeEnter})
	if clicks != 2 {
		t.Fatalf("聚焦后 Enter 点击次数 = %d", clicks)
	}
}

func TestButtonDisabledIgnoresClick(t *testing.T) {
	r := newControlRig()
	btn := NewButton("确定")
	btn.Disabled = true
	clicks := 0
	btn.OnClick = func() { clicks++ }
	r.mount(btn)
	r.click(20, 10)
	if clicks != 0 {
		t.Fatal("禁用按钮仍响应了点击")
	}
}

func TestCheckboxAndSwitchToggle(t *testing.T) {
	r := newControlRig()
	cb := NewCheckbox("选项", false)
	changed := 0
	cb.OnChange = func(bool) { changed++ }
	r.mount(cb)
	r.click(8, 8)
	if !cb.Checked || changed != 1 {
		t.Fatalf("复选框 = %v 回调 = %d", cb.Checked, changed)
	}

	r2 := newControlRig()
	sw := NewSwitch(false)
	r2.mount(sw)
	r2.click(8, 8)
	if !sw.On {
		t.Fatal("开关未切换")
	}
}

func TestRadioStaysSelected(t *testing.T) {
	r := newControlRig()
	radio := NewRadio("A", true)
	calls := 0
	radio.OnChange = func(bool) { calls++ }
	r.mount(radio)
	r.click(8, 8)
	if !radio.Checked || calls != 0 {
		t.Fatalf("已选中项不该再回调: checked=%v calls=%d", radio.Checked, calls)
	}
}

func TestInputKeyboardAndIME(t *testing.T) {
	r := newControlRig()
	w := NewInput("姓名")
	r.mount(w)
	if w.focus != nil {
		r.fm.Focus(w.focus)
	}
	h := w.Editor()
	h.InsertText("中")
	r.disp.Key(&input.KeyEvent{Code: input.CodeBackspace})
	if h.Value().Text != "" {
		t.Fatalf("退格后 = %q", h.Value().Text)
	}
	r.ime.Handle(host.IMFEvent{Kind: host.IMFStartComposition})
	r.ime.Handle(host.IMFEvent{Kind: host.IMFUpdateComposition, Text: "zhong", Cursor: 5})
	r.ime.Handle(host.IMFEvent{Kind: host.IMFResult, Text: "中"})
	if got := h.Value().Text; got != "中" {
		t.Fatalf("中文上屏 = %q", got)
	}
	if w.Value != "中" {
		t.Fatalf("控件 Value 未跟上编辑模型: %q", w.Value)
	}
}

func TestInputFocusConnectsIME(t *testing.T) {
	r := newControlRig()
	w := NewInput("姓名")
	r.mount(w)
	if w.focus == nil {
		t.Fatal("输入框没有焦点节点")
	}
	r.fm.Focus(w.focus)
	if r.ime.Active() || r.ime.Stats().Starts != 0 {
		t.Fatal("未输入时不该开始组合")
	}
	r.ime.Handle(host.IMFEvent{Kind: host.IMFStartComposition})
	if !r.ime.Active() {
		t.Fatal("IME 未接到聚焦的输入框（SetTextEditor 没生效）")
	}
	r.fm.Blur()
	if r.ime.Active() {
		t.Fatal("失焦后组合未撤销")
	}
}

func TestSliderDragChangesValue(t *testing.T) {
	r := newControlRig()
	s := NewSlider(0, 100, 10)
	changed := 0
	s.OnChange = func(float32) { changed++ }
	r.mount(s)
	r.drag(16, 120)
	if s.Value <= 10 || changed == 0 {
		t.Fatalf("拖动后 = %v 回调 = %d", s.Value, changed)
	}
}
