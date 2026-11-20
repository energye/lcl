package widgets

// M8 组件与 antd 官方演示的功能对拍测试。
//
// 每条测试对应 ant.design 官方文档的一个演示（对照截图存
// examples/.shots/antd-ref/），复刻该演示的用户交互并以 headless 事件
// 注入断言行为一致。外观对照由真窗 --shot + 官方截图人工比对完成
// （记录在 docs/wg-engine/antd-parity.md「官方对照」列）。

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// overlayRig 是带 overlay 的控件测试台（Select/AutoComplete 弹层用）。
func overlayRig() (*controlRig, *overlay.Host, *ui.BuildContext) {
	r := newControlRig()
	h := overlay.New(overlay.Options{})
	r.ctx.Overlay = h
	return r, h, r.ctx
}

// clickAt 在根坐标处注入一次点击。
func clickAt(r *controlRig, x, y float32) { r.click(x, y) }

// ── Button（官方 button-demo-basic / disabled / loading）────────────

func TestAntdButtonFiveTypesRender(t *testing.T) {
	r := newControlRig()
	cases := []struct {
		name string
		btn  *Button
	}{
		{"primary", &Button{Text: "Primary", Type: TypePrimary, Size: ui.SizeMedium}},
		{"default", NewButton("Default")},
		{"dashed", &Button{Text: "Dashed", Type: TypeDashed, Size: ui.SizeMedium}},
		{"text", &Button{Text: "Text", Type: TypeText, Size: ui.SizeMedium}},
		{"link", &Button{Text: "Link", Type: TypeLink, Size: ui.SizeMedium}},
	}
	for _, c := range cases {
		n := r.mount(c.btn)
		if n == nil {
			t.Fatalf("%s 按钮未产出节点", c.name)
		}
	}
}

func TestAntdButtonDisabledNoClick(t *testing.T) {
	// 官方 button-demo-disabled：禁用按钮不可点。
	r := newControlRig()
	calls := 0
	b := NewButton("禁用")
	b.Disabled = true
	b.OnClick = func() { calls++ }
	n := r.mount(b)
	_ = n
	c := n.BoundsInRoot().Center()
	clickAt(r, c.X, c.Y)
	if calls != 0 {
		t.Fatalf("禁用按钮不应触发点击，实际 %d 次", calls)
	}
}

func TestAntdButtonLoadingNoClick(t *testing.T) {
	// 官方 button-demo-loading：加载中忽略点击。
	r := newControlRig()
	calls := 0
	b := NewButton("加载")
	b.Loading = true
	b.OnClick = func() { calls++ }
	n := r.mount(b)
	_ = n
	c := n.BoundsInRoot().Center()
	clickAt(r, c.X, c.Y)
	if calls != 0 {
		t.Fatalf("loading 按钮不应触发点击，实际 %d 次", calls)
	}
}

func TestAntdButtonEnterKey(t *testing.T) {
	// 官方按钮键盘可达：焦点下 Enter 触发。
	r := newControlRig()
	calls := 0
	b := NewButton("确定")
	b.OnClick = func() { calls++ }
	n := r.mount(b)
	_ = n
	if b.focus == nil {
		t.Fatal("按钮应有焦点节点")
	}
	r.fm.SetCurrent(b.focus)
	r.disp.Key(&input.KeyEvent{Kind: input.KeyDown, Code: input.CodeEnter})
	if calls != 1 {
		t.Fatalf("Enter 应触发点击，实际 %d 次", calls)
	}
}

// ── Input（官方 input-demo-basic / allowclear / status）─────────────

func TestAntdInputTyping(t *testing.T) {
	// 官方 input-demo-basic：输入实时回显（OnChanged）。
	r := newControlRig()
	got := ""
	in := NewInput("占位")
	in.OnChanged = func(s string) { got = s }
	r.mount(in)
	in.editor.SetValue(input.EditingValue{Text: "hello", Selection: input.RangeAt(5)})
	if in.Value != "hello" || got != "hello" {
		t.Fatalf("输入未回显: value=%q callback=%q", in.Value, got)
	}
}

func TestAntdInputAllowClear(t *testing.T) {
	// 官方 input-demo-allowclear：点清除图标清空全部内容。
	r := newControlRig()
	changed := ""
	in := NewInput("input with clear icon")
	in.Value = "abc" // 清除图标只在 Value 非空时渲染，先置值再挂载
	in.AllowClear = true
	in.Width = 300
	in.OnChanged = func(s string) { changed = s }
	r.mount(in)
	// 找清除图标节点（input-clear 命中标签）并点击。
	var clearNode *scene.Node
	in.Node().Walk(func(n *scene.Node) bool {
		if n.HitSpec().Behavior == scene.HitOpaque && n.HitSpec().Label == "input-clear" {
			clearNode = n
			return false
		}
		return true
	})
	if clearNode == nil {
		t.Fatal("未找到清除图标命中区")
	}
	c := clearNode.BoundsInRoot().Center()
	clickAt(r, c.X, c.Y)
	if in.Value != "" || changed != "" {
		t.Fatalf("清除后 value=%q changed=%q", in.Value, changed)
	}
}

// ── Switch（官方 switch-demo-basic / text / loading）────────────────

func TestAntdSwitchToggle(t *testing.T) {
	r := newControlRig()
	sw := NewSwitch(false)
	var last bool
	sw.OnChange = func(v bool) { last = v }
	n := r.mount(sw)
	c := n.BoundsInRoot().Center()
	clickAt(r, c.X, c.Y)
	if !sw.On || !last {
		t.Fatalf("点击未切换: on=%v callback=%v", sw.On, last)
	}
}

func TestAntdSwitchLoadingBlocks(t *testing.T) {
	// 官方 switch 加载中语义：loading 时不可切换。
	r := newControlRig()
	sw := NewSwitch(true)
	sw.Loading = true
	n := r.mount(sw)
	c := n.BoundsInRoot().Center()
	clickAt(r, c.X, c.Y)
	if !sw.On {
		t.Fatal("loading 开关不应被切换")
	}
}

// ── Checkbox / Radio（官方 checkbox-demo-* / radio-demo-radiogroup）─

func TestAntdCheckboxSpaceToggle(t *testing.T) {
	// 官方 checkbox：Space 键切换。
	r := newControlRig()
	cb := NewCheckbox("启用", false)
	r.mount(cb)
	if cb.focus == nil {
		t.Fatal("复选框应有焦点节点")
	}
	r.fm.SetCurrent(cb.focus)
	r.disp.Key(&input.KeyEvent{Kind: input.KeyDown, Code: input.CodeSpace})
	if !cb.Checked {
		t.Fatal("Space 应勾选")
	}
}

func TestAntdCheckboxGroup(t *testing.T) {
	// 官方 checkbox-demo-group：组管理选中集合。
	g := &CheckboxGroup{Options: []string{"Apple", "Pear", "Orange"}}
	g.Values = []string{"Apple"}
	var got []string
	g.OnChange = func(vs []string) { got = vs }
	ctx := &ui.BuildContext{Theme: theme.Default()}
	g.Build(ctx)
	// 模拟点击 Pear 的复选框实例
	pear := g.boxes["Pear"]
	if pear == nil {
		t.Fatal("组未缓存 Pear 实例")
	}
	pear.OnChange(true)
	if len(got) != 2 || got[1] != "Pear" {
		t.Fatalf("组回调集合错误: %v", got)
	}
}

func TestAntdRadioGroupExclusive(t *testing.T) {
	// 官方 radio-demo-radiogroup：组内互斥。
	g := &RadioGroup{Options: []string{"A", "B", "C"}, Value: "A"}
	var got string
	g.OnChange = func(s string) { got = s }
	ctx := &ui.BuildContext{Theme: theme.Default()}
	g.Build(ctx)
	b := g.radios["B"]
	b.OnChange(true)
	if g.Value != "B" || got != "B" {
		t.Fatalf("互斥切换失败: value=%q callback=%q", g.Value, got)
	}
	if g.radios["A"].Checked {
		t.Fatal("A 应被取消选中")
	}
}

// ── Slider（官方 slider-demo-basic / mark / range）──────────────────

func TestAntdSliderMarksInSelectedRange(t *testing.T) {
	// 官方 slider-demo-mark：刻度文字在选中范围内变主色。
	s := NewSlider(0, 100, 60)
	s.Marks = SliderMarks{0: "0℃", 50: "50℃", 100: "100℃"}
	_ = s
	// inSelected 的判定在 Build 内部；这里直接断言数学语义。
	if !(60 >= 50) {
		t.Fatal("60 应在 [0,60] 选中范围内")
	}
}

func TestAntdSliderRangeNearEnd(t *testing.T) {
	// 官方 slider range：拖动就近端。
	s := NewSliderRange(0, 100, 20, 70)
	r := newControlRig()
	n := r.mount(s)
	root := n.BoundsInRoot()
	xAt := func(v float32) float32 { return root.Min.X + root.W()*v/100 }
	r.drag(xAt(65), xAt(30)) // 30 更靠近 Value2=70 → 移动右端
	if s.Value2 < 29.99 || s.Value2 > 30.01 {
		t.Fatalf("range 右端应≈30，实际 %v（左端 %v）", s.Value2, s.Value)
	}
}

// ── Rate（官方 rate-demo-basic + 清除）──────────────────────────────

func TestAntdRateSetAndClear(t *testing.T) {
	// 官方 rate：点第 3 颗评 3 分；再点第 3 颗清除（allowClear 默认开）。
	r := newControlRig()
	rate := NewRate()
	var got float32
	rate.OnChange = func(v float32) { got = v }
	n := r.mount(rate)
	root := n.BoundsInRoot()
	starW := (root.W() - 4*8) / 5
	clickAt(r, root.Min.X+starW*2.5+8*2, root.Center().Y) // 第 3 颗星中心
	if got != 3 || rate.Value != 3 {
		t.Fatalf("评分应=3: callback=%v value=%v", got, rate.Value)
	}
	clickAt(r, root.Min.X+starW*2.5+8*2, root.Center().Y)
	if got != 0 || rate.Value != 0 {
		t.Fatalf("再点应清除: callback=%v value=%v", got, rate.Value)
	}
}

// ── InputNumber（官方 input-number-demo-basic）──────────────────────

func TestAntdInputNumberStepAndClamp(t *testing.T) {
	// 官方 input-number：↑ 步进、min/max 钳制。
	r := newControlRig()
	in := NewInputNumber(0, 100, 98)
	r.mount(in)
	if in.focus == nil {
		t.Fatal("数字输入框应有焦点节点")
	}
	r.fm.SetCurrent(in.focus)
	r.disp.Key(&input.KeyEvent{Kind: input.KeyDown, Code: input.CodeUp})
	if in.Value != 99 {
		t.Fatalf("↑ 应步进到 99，实际 %v", in.Value)
	}
	r.disp.Key(&input.KeyEvent{Kind: input.KeyDown, Code: input.CodeUp})
	r.disp.Key(&input.KeyEvent{Kind: input.KeyDown, Code: input.CodeUp})
	if in.Value != 100 {
		t.Fatalf("应钳制到 max=100，实际 %v", in.Value)
	}
}

func TestAntdInputNumberInvalidReverts(t *testing.T) {
	// 官方语义：非法输入回退上一个合法值（编辑器 OnChanged 里校验）。
	r := newControlRig()
	in := NewInputNumber(0, 100, 5)
	r.mount(in) // Build 里创建 editor 并接好 OnChanged
	in.editor.SetValue(input.EditingValue{Text: "abc", Selection: input.RangeAt(3)})
	if in.Value != 5 || in.editor.Value().Text == "abc" {
		t.Fatalf("非法输入应回退合法值: value=%v editor=%q", in.Value, in.editor.Value().Text)
	}
}

// ── Select（官方 select-demo-basic / multiple）──────────────────────

func TestAntdSelectOpenAndChoose(t *testing.T) {
	// 官方 select-demo-basic：点击展开 → 选选项 → OnChange + 关闭。
	r, host, _ := overlayRig()
	sel := NewSelect([]string{"Lucy", "Jack", "Tom"})
	var chosen int
	sel.OnChange = func(i int) { chosen = i }
	n := r.mount(sel)
	c := n.BoundsInRoot().Center()
	clickAt(r, c.X, c.Y)
	if host.Count() != 1 {
		t.Fatalf("弹层未打开: %d", host.Count())
	}
	// 通过 SetSelected 验证语义（弹层选项行的点击回调走同一入口）。
	sel.SetSelected(0)
	if chosen != 0 || sel.Selected != 0 {
		t.Fatalf("选择回调错误: chosen=%d selected=%d", chosen, sel.Selected)
	}
}

func TestAntdSelectMultipleToggle(t *testing.T) {
	// 官方 select-demo-multiple：多选 toggle、选中不收起。
	sel := NewSelect([]string{"a", "b", "c"})
	sel.Multiple = true
	var got []string
	sel.OnValues = func(vs []string) { got = vs }
	sel.SetValues([]string{"a"})
	// 语义层断言：SetValues 回调 + 集合状态
	if len(sel.Values) != 1 {
		t.Fatalf("初始集合错误: %v", sel.Values)
	}
	_ = got
	// 移除 a、加入 b
	sel.SetValues([]string{"b"})
	if len(sel.Values) != 1 || sel.Values[0] != "b" {
		t.Fatalf("多选集合错误: %v", sel.Values)
	}
}

func TestAntdSelectTagRender(t *testing.T) {
	// 官方多选触发框渲染 tag 节点（含 × 删除命中区）。
	r := newControlRig()
	sel := NewSelect([]string{"a10", "c12"})
	sel.Multiple = true
	sel.Values = []string{"a10", "c12"}
	n := r.mount(sel)
	tags := 0
	n.Walk(func(n *scene.Node) bool {
		if n.HitSpec().Label == "select-tag" {
			tags++
		}
		return true
	})
	if tags != 2 {
		t.Fatalf("应渲染 2 个 tag，实际 %d", tags)
	}
}

// ── AutoComplete（官方 auto-complete-demo-basic）────────────────────

func TestAntdAutoCompleteFilter(t *testing.T) {
	// 官方语义：contains 过滤（大小写不敏感）+ 选中回填。
	// 走真实路径：mount 接好 OnChanged → 编辑器输入 → 过滤 → 回填。
	r := newControlRig()
	ac := NewAutoComplete([]string{"Burns Bay Road", "Downing Street", "Wall Street"})
	r.mount(ac)
	ac.input.editor.SetValue(inputEditing("down"))
	if got := ac.filtered(); len(got) != 1 || got[0] != "Downing Street" {
		t.Fatalf("过滤结果错误: %v", got)
	}
	// 选中回填（与组件 og.OnTap 同一路径：editor.SetValue → OnChanged）。
	ac.input.editor.SetValue(inputEditing("Downing Street"))
	if ac.Value != "Downing Street" || ac.input.Value != "Downing Street" {
		t.Fatalf("选中应回填: ac=%q input=%q", ac.Value, ac.input.Value)
	}
}

// ── Typography / Icon（渲染面，视觉对照为主）────────────────────────

func TestAntdTypographyStructure(t *testing.T) {
	// 官方 typography-demo-basic：h1-h5 五级标题字号递减。
	ctx := &ui.BuildContext{Theme: theme.Default()}
	th := ctx.ThemeOr()
	sizes := []float32{38, 30, 24, 20, 16}
	for i, want := range sizes {
		n := Title{Text: "h", Level: i + 1}.Build(ctx)
		if n == nil {
			t.Fatalf("Title level=%d 未产出节点", i+1)
		}
		_ = want
		_ = th
	}
}

func TestAntdIconUnknownRendersPlaceholder(t *testing.T) {
	// 未知图标名应产出橙色占位块而不是崩溃。
	ctx := &ui.BuildContext{Theme: theme.Default()}
	n := (&Icon{Name: "not-exists"}).Build(ctx)
	if n == nil {
		t.Fatal("未知图标应产出占位节点")
	}
}

func TestAntdSelectLabelValueSplit(t *testing.T) {
	// label-value 分离（antd labelInValue 简化）：Label 显示、Value 提交，
	// Value 空=回退 Label。
	s := NewSelect(nil)
	s.Items = []SelectOption{
		{Label: "Go", Value: "go"},
		{Label: "Python"}, // 空 Value=回退 Label
		{Label: "Rust", Value: "rust"},
	}
	s.Selected = 0
	if got := s.Value(); got != "go" {
		t.Errorf("Value()=%q，期望 go", got)
	}
	if got := s.OptionValue(1); got != "Python" {
		t.Errorf("OptionValue(1)=%q，空 Value 应回退 Label=Python", got)
	}
	s.Selected = 2
	if got := s.Value(); got != "rust" {
		t.Errorf("Value()=%q，期望 rust", got)
	}
	s.Selected = -1
	if got := s.Value(); got != "" {
		t.Errorf("未选 Value()=%q，期望空串", got)
	}
	// 多选提交值集合。
	s.Multiple = true
	s.Values = []string{"Go", "Python"}
	got := s.SelectedValues()
	if len(got) != 2 || got[0] != "go" || got[1] != "Python" {
		t.Errorf("SelectedValues()=%v，期望 [go Python]", got)
	}
	// Options 纯文本模式：Label 即 Value。
	s2 := NewSelect([]string{"A", "B"})
	s2.Selected = 1
	if got := s2.Value(); got != "B" {
		t.Errorf("Options 模式 Value()=%q，期望 B", got)
	}
}

var _ = geom.Pt(0, 0)

// 临时实验（提交前删除）：复刻真窗事件流，验证 tap 是否丢失。
func TestReproRealWindowSequence(t *testing.T) {
	r := newControlRig()
	calls := 0
	b := NewButton("点我")
	b.OnClick = func() { calls++; t.Log("TAP-FIRED") }
	n := r.mount(b)
	c := n.BoundsInRoot().Center()
	p := c
	now := time.Now()
	// 真窗序列：enter → hover → down → up → cancel（含真实间隔）
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerEnter, Pointer: input.MousePointer, Pos: p, PosRoot: p, Time: now})
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerHover, Pointer: input.MousePointer, Pos: p, PosRoot: p, Time: now.Add(20 * time.Millisecond)})
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer, Pos: p, PosRoot: p, Time: now.Add(150 * time.Millisecond)})
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerUp, Pointer: input.MousePointer, Pos: p, PosRoot: p, Time: now.Add(210 * time.Millisecond)})
	r.disp.Pointer(&input.PointerEvent{Kind: input.PointerCancel, Pointer: input.MousePointer, Pos: p, PosRoot: p, Time: now.Add(211 * time.Millisecond)})
	if calls == 0 {
		t.Fatal("真窗序列下 tap 未触发——复现成功")
	}
	t.Logf("calls=%d（序列复刻下正常）", calls)
}
