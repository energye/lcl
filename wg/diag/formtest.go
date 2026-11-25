package diag

import (
	"errors"
	"fmt"
	"go/build"
	"go/parser"
	"go/token"
	"strings"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/widgets"
)

// ── --formtest：表单页 demo 门禁（M3）───────────────────────────────
//
// 与单测的区别：这里的入口是 Container 的公共 API——MountWidget/PushEvent/
// IME/Overlay——和真实程序走同一条链路。组件内部改了但接线断了，单测可能
// 仍然全绿，这里一定会红。

// FormCase 是单条断言结果。
type FormCase struct {
	Name string
	OK   bool
	Got  string
	Note string
}

// FormResult 是汇总。
type FormResult struct {
	Cases   int
	Failed  int
	Elapsed time.Duration
	Details []FormCase
}

// OK 判定是否全绿（门禁 7 组）。
func (r FormResult) OK() bool { return r.Failed == 0 && r.Cases >= 7 }

func (r FormResult) String() string {
	return fmt.Sprintf("formtest: %d 项，失败 %d，用时 %v", r.Cases, r.Failed, r.Elapsed)
}

// formRig 是一套已挂载、已布局的表单页。
type formRig struct {
	c    *wg.Container
	form *widgets.DemoForm
}

func newFormRig() *formRig {
	c := wg.NewHeadless(wg.DefaultOptions())
	c.Resize(900, 620)
	form := widgets.NewDemoForm()
	c.MountWidget(form)
	_, _ = c.FrameOnce()
	return &formRig{c: c, form: form}
}

func (r *formRig) close() { r.c.Close() }

// focus 把焦点交给控件（与点击得到的焦点路径一致）。
func (r *formRig) focus(n *input.FocusNode) {
	r.c.Focus().SetCurrent(n)
}

func (r *formRig) ime(text string) {
	r.c.IME().Handle(host.IMFEvent{Kind: host.IMFStartComposition})
	r.c.IME().Handle(host.IMFEvent{Kind: host.IMFUpdateComposition, Text: text, Cursor: int32(len([]rune(text)))})
	r.c.IME().Handle(host.IMFEvent{Kind: host.IMFResult, Text: text})
	r.c.IME().Handle(host.IMFEvent{Kind: host.IMFEndComposition})
}

// clickNode 按节点的根空间包围盒中心投递一次真实点击。
func (r *formRig) clickNode(n *scene.Node) {
	if n == nil {
		return
	}
	b := n.BoundsInRoot()
	p := geom.Pt(b.Min.X+b.W()/2, b.Min.Y+b.H()/2)
	t := time.Now()
	r.c.PushEvent(&input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer,
		Pos: p, PosRoot: p, Time: t})
	r.c.PushEvent(&input.PointerEvent{Kind: input.PointerUp, Pointer: input.MousePointer,
		Pos: p, PosRoot: p, Time: t.Add(16 * time.Millisecond)})
}

func (r *formRig) hoverNode(n *scene.Node, kind input.PointerKind) {
	if n == nil {
		return
	}
	b := n.BoundsInRoot()
	p := geom.Pt(b.Min.X+b.W()/2, b.Min.Y+b.H()/2)
	r.c.PushEvent(&input.PointerEvent{Kind: kind, Pointer: input.MousePointer,
		Pos: p, PosRoot: p, Time: time.Now()})
}

// FormTest 跑表单页门禁。
func FormTest() (FormResult, error) {
	began := time.Now()
	var res FormResult
	cases := []struct {
		name string
		run  func() (string, error)
	}{
		{"form-builds", caseFormBuilds},
		{"ime-chinese-name", caseFormIME},
		{"tab-moves-focus", caseFormTab},
		{"submit-validates", caseFormSubmit},
		{"tabs-switch", caseFormTabsSwitch},
		{"tooltip-show-hide", caseFormTooltip},
		{"zero-lcl", caseFormZeroLCL},
	}
	for _, c := range cases {
		got, err := c.run()
		d := FormCase{Name: c.name, OK: err == nil, Got: got}
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

func caseFormBuilds() (string, error) {
	r := newFormRig()
	defer r.close()
	nodes := 0
	var walk func(*scene.Node)
	walk = func(n *scene.Node) {
		nodes++
		for _, c := range n.Children() {
			walk(c)
		}
	}
	walk(r.c.Scene())
	focusables := len(r.c.Focus().List())
	got := fmt.Sprintf("节点=%d 可聚焦=%d", nodes, focusables)
	if nodes < 20 || focusables < 8 {
		return got, errors.New("表单页没搭起来（节点/焦点链太少）")
	}
	return got, nil
}

func caseFormIME() (string, error) {
	r := newFormRig()
	defer r.close()
	r.focus(r.form.Name.FocusNode())
	r.ime("张伟")
	got := fmt.Sprintf("name=%q", r.form.Name.Value)
	if r.form.Name.Value != "张伟" {
		return got, errors.New("中文没有上屏到姓名字段")
	}
	return got, nil
}

func caseFormTab() (string, error) {
	r := newFormRig()
	defer r.close()
	r.focus(r.form.Name.FocusNode())
	r.c.PushEvent(&input.KeyEvent{Code: input.CodeTab})
	cur := r.c.Focus().Current()
	got := fmt.Sprintf("focus=%v", cur != nil && cur == r.form.City.FocusNode())
	if cur != r.form.City.FocusNode() {
		return got, errors.New("输入框里按 Tab 没有移到下一格")
	}
	return got, nil
}

func caseFormSubmit() (string, error) {
	r := newFormRig()
	defer r.close()
	// 空表单提交：应提示填姓名。
	r.clickNode(r.form.Submit.Node())
	if r.form.Status != "请填写姓名" {
		return fmt.Sprintf("status=%q", r.form.Status), errors.New("空表单未校验")
	}
	r.focus(r.form.Name.FocusNode())
	r.ime("张伟")
	r.focus(r.form.City.FocusNode())
	r.ime("上海")
	r.clickNode(r.form.Submit.Node())
	got := fmt.Sprintf("status=%q", r.form.Status)
	if !strings.Contains(r.form.Status, "张伟") || !strings.Contains(r.form.Status, "上海") {
		return got, errors.New("提交结果没有带上输入内容")
	}
	return got, nil
}

func caseFormTabsSwitch() (string, error) {
	r := newFormRig()
	defer r.close()
	root := r.form.Tabs.Node()
	if root == nil || len(root.Children()) == 0 {
		return "找不到标签条", errors.New("Tabs 未构建")
	}
	// line 型页签的指示条已提升到 bar 层 Stack，children 顺序不保证；
	// 直接按 tab 命中锚收集页签表头。
	headers := collectHitLabel(root, "tab")
	if len(headers) < 2 {
		return fmt.Sprintf("页签数=%d", len(headers)), errors.New("页签不足两个")
	}
	r.clickNode(headers[1])
	got := fmt.Sprintf("active=%d", r.form.Tabs.Active)
	if r.form.Tabs.Active != 1 {
		return got, errors.New("点击第二个页签没有切换")
	}
	return got, nil
}

// collectHitLabel 深度优先按文档序收集命中锚为 label 的节点。
func collectHitLabel(n *scene.Node, label string) []*scene.Node {
	if n == nil {
		return nil
	}
	var out []*scene.Node
	if n.HitSpec().Label == label {
		out = append(out, n)
	}
	for _, c := range n.Children() {
		out = append(out, collectHitLabel(c, label)...)
	}
	return out
}

func caseFormTooltip() (string, error) {
	r := newFormRig()
	defer r.close()
	before := r.c.Overlay().Count()
	r.hoverNode(r.form.Submit.Node(), input.PointerHover)
	during := r.c.Overlay().Count()
	r.hoverNode(r.form.Submit.Node(), input.PointerLeave)
	after := r.c.Overlay().Count()
	got := fmt.Sprintf("%d → %d → %d", before, during, after)
	if before != 0 || during != 1 || after != 0 {
		return got, errors.New("Tooltip 没有随悬停显示/离开关闭")
	}
	return got, nil
}

func caseFormZeroLCL() (string, error) {
	pkg, err := build.Import("github.com/energye/lcl/wg/widgets", "", build.FindOnly)
	if err != nil {
		return err.Error(), err
	}
	fset := token.NewFileSet()
	parsed, err := parser.ParseDir(fset, pkg.Dir, nil, parser.ImportsOnly)
	if err != nil {
		return err.Error(), err
	}
	files, imports := 0, 0
	for _, p := range parsed {
		for _, f := range p.Files {
			files++
			for _, spec := range f.Imports {
				path := strings.Trim(spec.Path.Value, `"`)
				if path == "github.com/energye/lcl" || strings.HasPrefix(path, "github.com/energye/lcl/lcl") {
					imports++
				}
			}
		}
	}
	got := fmt.Sprintf("文件=%d 违规 import=%d", files, imports)
	if files < 5 || imports != 0 {
		return got, errors.New("组件库出现 lcl 依赖（表单页就不算零 lcl 了）")
	}
	return got, nil
}
