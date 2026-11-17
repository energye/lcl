package wg

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// ipPage 包装一个 widget 并计数 Build（验证"编辑不重建"）。
type ipPage struct {
	ui.Base
	child  ui.Widget
	builds int
}

func (p *ipPage) Build(ctx *ui.BuildContext) *scene.Node {
	p.builds++
	return p.child.Build(ctx)
}

// TestInputEditDoesNotRebuild 锁定"打字只更新节点、不整页重建"。
func TestInputEditDoesNotRebuild(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 260)

	in := widgets.NewInput("请输入")
	in.Width = 200
	page := &ipPage{child: in}
	c.MountWidget(page)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}

	// 聚焦 → Focus 事件 MarkDirty → 下一帧重建（建立光标节点）。
	c.focus.Focus(in.FocusNode())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	base := page.builds
	if base < 2 {
		t.Fatalf("聚焦应重建一次，builds=%d", base)
	}

	// 编辑（等价于键入）：只应原地刷新，不应重建整页。
	in.Editor().InsertText("a")
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if got := in.Editor().Value().Text; got != "a" {
		t.Fatalf("编辑器文本 = %q", got)
	}
	if page.builds != base {
		t.Fatalf("键入触发了整页重建：%d → %d", base, page.builds)
	}
}

// TestInputTypedKeyDoesNotRebuild 走真实分发路径（KeyEvent.Text）验证不重建。
func TestInputTypedKeyDoesNotRebuild(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 260)

	in := widgets.NewInput("请输入")
	in.Width = 200
	page := &ipPage{child: in}
	c.MountWidget(page)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.focus.Focus(in.FocusNode())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	base := page.builds
	for _, ch := range "abc" {
		c.PushEvent(&input.KeyEvent{Kind: input.KeyDown, Text: string(ch), Time: time.Now()})
	}
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if got := in.Editor().Value().Text; got != "abc" {
		t.Fatalf("文本 = %q", got)
	}
	if page.builds != base {
		t.Fatalf("键入触发了整页重建：%d → %d", base, page.builds)
	}
}

// TestInputBlurOnOutsideClick 锁定"点击非焦点区域失焦"。
func TestInputBlurOnOutsideClick(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 260)

	in := widgets.NewInput("请输入")
	in.Width = 200
	page := &ipPage{child: widgets.Column{Gap: 12, Children: []ui.Widget{
		in,
		&widgets.Label{Text: "其它内容"},
	}}}
	c.MountWidget(page)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.focus.Focus(in.FocusNode())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if c.focus.Current() == nil {
		t.Fatal("聚焦失败")
	}

	now := time.Now()
	// 点击空白区（右下方，不在输入框子树内）。
	pt := geom.Pt(380, 240)
	c.PushEvent(&input.PointerEvent{
		Kind: input.PointerDown, Pointer: input.MousePointer,
		Pos: pt, PosRoot: pt, Time: now,
	})
	if c.focus.Current() != nil {
		t.Fatal("点击其它区域后输入框仍有焦点")
	}
}

// TestIMECaretAnchor 锁定 IME 光标锚点：布局后换算成客户区坐标。
func TestIMECaretAnchor(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 260)

	in := widgets.NewInput("")
	in.Width = 200
	c.MountWidget(&ipPage{child: widgets.Column{Children: []ui.Widget{in}}})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.focus.Focus(in.FocusNode())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.mu.Lock()
	ok, rect := c.imeCaretOK, c.imeCaret
	c.mu.Unlock()
	if !ok {
		t.Fatal("聚焦输入框后应有 IME 光标锚点")
	}
	// 输入框内边距 11，行高内垂直居中：x 应约为 11。
	if rect.Min.X < 10 || rect.Min.X > 13 {
		t.Fatalf("锚点 x=%v（应 ≈11）", rect.Min.X)
	}
	if rect.Max.Y <= rect.Min.Y {
		t.Fatalf("锚点高度非法: %v", rect)
	}
}

// TestInputTypingDamageIsLocal 锁定"键入只重绘输入框附近"。
//
// 每次按键触发全窗重排/重绘是输入延迟的主因：整页 1080×760 的
// 重新布局 + 全窗光栅在真窗里是几十毫秒，表现为"打字发粘"。
func TestInputTypingDamageIsLocal(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 260)

	in := widgets.NewInput("")
	in.Width = 200
	c.MountWidget(&ipPage{child: widgets.Column{Children: []ui.Widget{in}}})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.focus.Focus(in.FocusNode())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.PushEvent(&input.KeyEvent{Kind: input.KeyDown, Text: "a", Time: time.Now()})
	st, err := c.FrameOnce()
	if err != nil {
		t.Fatal(err)
	}
	if !st.Presented {
		t.Fatal("键入帧未提交")
	}
	// 输入框 200 宽 + 光标；留一点余量，仍应远小于 400×260 的整窗。
	if st.Damage.W() > 260 || st.Damage.H() > 80 {
		t.Fatalf("键入帧脏区过大：%v（应≈输入框）", st.Damage)
	}
}

// TestInputCaretBlinks 锁定插入符闪烁：530ms 明暗切换，编辑时重置相位。
func TestInputCaretBlinks(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 260)
	clk := frame.NewManualClock(0, 530*time.Millisecond)
	c.SetClock(clk)

	in := widgets.NewInput("")
	in.Width = 200
	c.MountWidget(&ipPage{child: in})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.focus.Focus(in.FocusNode())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	caret := in.CaretNode()
	if caret == nil {
		t.Fatal("聚焦后应有插入符节点")
	}
	if !caret.Visible() {
		t.Fatal("聚焦初始插入符应可见")
	}

	clk.Advance() // +530ms
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if caret.Visible() {
		t.Fatal("一个闪烁周期后插入符应变暗")
	}

	clk.Advance() // +530ms
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if !caret.Visible() {
		t.Fatal("第二个周期后插入符应恢复可见")
	}

	// 变暗期间敲字：编辑必须立即恢复可见并重置相位。
	clk.Advance()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if caret.Visible() {
		t.Fatal("编辑前应处于变暗相位")
	}
	c.PushEvent(&input.KeyEvent{Kind: input.KeyDown, Text: "a", Time: time.Now()})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if !caret.Visible() {
		t.Fatal("编辑后插入符应立刻恢复可见")
	}
}

// TestInputCaretBlinkSchedulesFrames 锁定"闪烁由帧泵自驱动"：
// 无输入时也应每 530ms 自动出一帧切换插入符可见性（静止窗 C14 的低频例外）。
func TestInputCaretBlinkSchedulesFrames(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 260)

	in := widgets.NewInput("")
	in.Width = 200
	c.MountWidget(&ipPage{child: in})

	clk := frame.NewManualClock(0, frame.DefaultStep)
	tk := c.RunManual(clk)
	if tk == nil {
		t.Fatal("RunManual 失败")
	}
	tk.Advance(frame.DefaultStep) // 首帧
	c.focus.Focus(in.FocusNode())
	tk.Advance(frame.DefaultStep) // 聚焦帧（建光标并装闪烁闹钟）
	caret := in.CaretNode()
	if caret == nil || !caret.Visible() {
		t.Fatal("聚焦后插入符应可见")
	}
	// 帧内的 deferred 续帧可能先占一个 60Hz 槽；多推几拍，等到闪烁闹钟生效。
	armed := func() time.Duration {
		for i := 0; i < 6; i++ {
			if tk.Armed() {
				if d := tk.Deadline() - clk.Now(); d > 100*time.Millisecond {
					return d
				}
			}
			tk.Advance(frame.DefaultStep)
		}
		return 0
	}
	d := armed()
	if d == 0 {
		t.Fatal("帧末应装填闪烁闹钟（>100ms）")
	}
	if !caret.Visible() {
		t.Fatal("闪烁前插入符应可见")
	}
	tk.Advance(d)
	if caret.Visible() {
		t.Fatal("到点后插入符应变暗")
	}
	d = armed()
	if d == 0 {
		t.Fatal("闪烁后应继续装填下一拍")
	}
	tk.Advance(d)
	if !caret.Visible() {
		t.Fatal("第二拍后插入符应恢复")
	}
}

// TestFocusedInputOwnsCaretAcrossSiblings 锁定"多输入框时上报不互相清掉"。
//
// 页面上每个未聚焦 Input 也会上报 nil；若宿主不按 owner 归属处理，
// 后构建的未聚焦输入框会清掉已聚焦输入框的插入符与 IME 锚点
// （真窗表现：刚点进输入框/重新聚焦时光标不闪、候选框定位丢失）。
func TestFocusedInputOwnsCaretAcrossSiblings(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 260)

	a := widgets.NewInput("a")
	a.Width = 200
	b := widgets.NewInput("b")
	b.Width = 200
	// 聚焦的 a 先构建、未聚焦的 b 后构建——旧实现里 b 的 nil 会清掉 a。
	c.MountWidget(&ipPage{child: widgets.Column{Children: []ui.Widget{a, b}}})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.focus.Focus(a.FocusNode())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	caret := a.CaretNode()
	if caret == nil || !caret.Visible() {
		t.Fatal("聚焦输入框应有可见插入符")
	}
	c.mu.Lock()
	owner, armed, ok := c.caretNode, c.caretArmed, c.imeCaretOK
	c.mu.Unlock()
	if owner != caret || !armed {
		t.Fatal("未聚焦输入框的上报清掉了插入符登记")
	}
	if !ok {
		t.Fatal("未聚焦输入框的上报清掉了 IME 锚点")
	}

	// 焦点迁移到 b：登记归属必须跟着走。
	c.focus.Focus(b.FocusNode())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.mu.Lock()
	owner = c.caretNode
	c.mu.Unlock()
	if owner != b.CaretNode() {
		t.Fatal("焦点迁移后插入符登记未切换到新输入框")
	}
}

// TestInputCountUpdatesInPlace 锁定"带 Count 的单行输入也不整页重建"。
//
// 字数行参与原地刷新：每键只更新文本/光标/字数节点，脏区局部。
func TestInputCountUpdatesInPlace(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 260)

	in := widgets.NewInput("")
	in.Width = 200
	in.Count = true
	// 外层 Column 让 Count 行按自然宽度布局（直接作为根会被视口紧约束拉满）。
	page := &ipPage{child: widgets.Column{Children: []ui.Widget{in}}}
	c.MountWidget(page)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	c.focus.Focus(in.FocusNode())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	base := page.builds

	now := time.Now()
	for _, ch := range "ab" {
		c.PushEvent(&input.KeyEvent{Kind: input.KeyDown, Text: string(ch), Time: now})
	}
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if page.builds != base {
		t.Fatalf("带 Count 的键入触发了整页重建：%d → %d", base, page.builds)
	}
	if got := in.Editor().Value().Text; got != "ab" {
		t.Fatalf("编辑器文本 = %q", got)
	}
	cnt := findTextNode(c.Scene(), "2")
	if cnt == nil {
		t.Fatal("字数行未更新为 2")
	}
	if fs := cnt.Text().Layout.Font.Size; fs <= 0 {
		t.Fatalf("字数行字体被零值重排：%v", fs)
	}

	c.PushEvent(&input.KeyEvent{Kind: input.KeyDown, Text: "c", Time: now})
	st, err := c.FrameOnce()
	if err != nil {
		t.Fatal(err)
	}
	if st.Damage.W() > 260 || st.Damage.H() > 100 {
		t.Fatalf("带 Count 的键入脏区过大：%v", st.Damage)
	}
	if findTextNode(c.Scene(), "3") == nil {
		t.Fatal("字数行未更新为 3")
	}
}
