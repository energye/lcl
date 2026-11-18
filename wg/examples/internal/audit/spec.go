// spec 执行器：解释声明式测试计划（spec.json）。
//
// 规范见 docs/wg-engine/06-COMPONENT-TEST-SPEC.md。设计边界：
//   - 只用可观测信号：命中树（label/文本/bounds）、合成缓冲像素、焦点、
//     浮层计数、页面滚动偏移；
//   - 禁止绝对坐标：定位靠 label+文本锚点，元素内位置用 frac（0–1）；
//   - 模板优先：交互元素的标准断言集（toggle/disabled-no-op/…）在这里
//     展开，窗级 spec 只写本窗特有预期。
package audit

import (
	"encoding/json"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
)

// ── 模型 ────────────────────────────────────────────────────────────

// Spec 是一个功能窗的声明式测试计划。
type Spec struct {
	ID         string          `json:"id,omitempty"`
	Structure  *SpecStructure  `json:"structure,omitempty"`
	Sections   []SpecSection   `json:"sections,omitempty"`
	Exemptions []SpecExemption `json:"exemptions,omitempty"`
}

// SpecStructure 是页面骨架断言（L1）。
type SpecStructure struct {
	Sections    []string `json:"sections,omitempty"`    // demo 卡标题必须存在
	APIRows     int      `json:"apiRows,omitempty"`     // API 表行数下限（封顶 12，虚拟化窗口）
	MinElements int      `json:"minElements,omitempty"` // 全树可交互节点计数下限
}

// SpecSection 对应一张 demo 卡。
type SpecSection struct {
	Name  string     `json:"name"`
	Cases []SpecCase `json:"cases,omitempty"`
}

// SpecCase 是一条用例。
type SpecCase struct {
	Name   string          `json:"name"`
	Target *SpecTarget     `json:"target,omitempty"`
	Given  []SpecGiven     `json:"given,omitempty"`
	When   []SpecAction    `json:"when,omitempty"`
	Then   []SpecAssertion `json:"then,omitempty"`
}

// SpecTarget 是定位锚。
type SpecTarget struct {
	Label    string `json:"label"`
	Text     string `json:"text,omitempty"` // 节点子树内文本子串
	Disabled *bool  `json:"disabled,omitempty"`
	Index    int    `json:"index,omitempty"` // 同锚第 n 个（树序）
}

// SpecGiven 是前置步骤。
type SpecGiven struct {
	ScrollTo      string      `json:"scrollTo,omitempty"`
	CloseOverlays *bool       `json:"closeOverlays,omitempty"`
	Focus         *SpecTarget `json:"focus,omitempty"`
}

// SpecAction 是一次操作。
type SpecAction struct {
	Click    *SpecClick `json:"click,omitempty"`
	Hover    *SpecClick `json:"hover,omitempty"`
	TypeText string     `json:"typeText,omitempty"`
	Key      string     `json:"key,omitempty"`
	Wheel    *float32   `json:"wheel,omitempty"`
	MoveOut  bool       `json:"moveOut,omitempty"`
	Wait     int        `json:"wait,omitempty"` // 毫秒（Tooltip 悬停延迟等）
}

// SpecClick 是点击/悬停位置（目标锚内的 frac，默认中心；Target 可覆盖锚）。
type SpecClick struct {
	Frac   []float32   `json:"frac,omitempty"`
	Target *SpecTarget `json:"target,omitempty"`
}

// SpecAssertion 是一条断言。
type SpecAssertion struct {
	Kind     string      `json:"kind"`
	Target   *SpecTarget `json:"target,omitempty"`
	Of       *SpecTarget `json:"of,omitempty"`
	Count    int         `json:"count,omitempty"`
	Scope    string      `json:"scope,omitempty"`
	Expect   any         `json:"expect,omitempty"`
	Tol      *float64    `json:"tol,omitempty"`
	MinRatio *float64    `json:"minRatio,omitempty"`
	TolPeak  *float64    `json:"tolPeak,omitempty"`
	Frac     []float32   `json:"frac,omitempty"`
	Template string      `json:"template,omitempty"`

	// geometry
	RightOf       *SpecTarget `json:"rightOf,omitempty"`
	Below         *SpecTarget `json:"below,omitempty"`
	WidthAtLeast  float32     `json:"widthAtLeast,omitempty"`
	HeightBetween []float32   `json:"heightBetween,omitempty"`
	CenterYWithin []float32   `json:"centerYWithin,omitempty"`
	InViewport    bool        `json:"inViewport,omitempty"`

	// overlay / scroll / text
	Band     string   `json:"band,omitempty"`
	Delta    *float32 `json:"delta,omitempty"`
	MinDelta *float32 `json:"minDelta,omitempty"`
	Text     string   `json:"text,omitempty"`
}

// SpecExemption 是显式豁免（antd 语义下正确但无视觉变化的行为）。
type SpecExemption struct {
	Case   string `json:"case"`
	Reason string `json:"reason,omitempty"`
}

// ParseSpec 解析 spec.json。
func ParseSpec(data []byte) (*Spec, error) {
	var s Spec
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

// ── 执行器 ──────────────────────────────────────────────────────────

type specRunner struct {
	c       *wg.Container
	outDir  string
	rep     *Report
	covered map[string]bool

	start      *image.RGBA // case 起始帧（colorDelta / 往返对照）
	startOff   float32     // case 起始页面滚动偏移
	caseBounds geom.Rect   // case target 的初始位置（点击改变文本后的回退锚）
}

// runSpecModel 执行整份 spec，返回被用例覆盖的标签集合（供启发式兜底跳过）。
func runSpecModel(c *wg.Container, rep *Report, outDir string, spec *Spec) map[string]bool {
	r := &specRunner{c: c, outDir: outDir, rep: rep, covered: map[string]bool{}}
	r.runStructure(spec.Structure)
	exempt := map[string]bool{}
	for _, ex := range spec.Exemptions {
		exempt[ex.Case] = true
	}
	for _, sec := range spec.Sections {
		for _, cs := range sec.Cases {
			if exempt[cs.Name] {
				rep.Findings = append(rep.Findings, Finding{Element: cs.Name, Check: "豁免",
					OK: true, Detail: "spec exemptions（正确但无视觉变化）"})
				continue
			}
			r.runCase(cs)
		}
	}
	return r.covered
}

func (r *specRunner) add(caseName, check string, ok bool, detail string) {
	r.rep.Findings = append(r.rep.Findings, Finding{Element: caseName, Check: check, OK: ok, Detail: detail})
	if !ok {
		r.rep.Failed++
		r.rep.SpecFailed++
		if os.Getenv("AUDIT_SHOTS") != "" {
			savePNG(grab(r.c), filepath.Join(r.outDir, fmt.Sprintf("fail-spec-%03d.png", len(r.rep.Findings))))
		}
	}
}

// ── 结构断言（L1）────────────────────────────────────────────────────

func (r *specRunner) runStructure(st *SpecStructure) {
	if st == nil {
		return
	}
	for _, title := range st.Sections {
		found := false
		r.c.Scene().Walk(func(n *scene.Node) bool {
			if n.Kind() == scene.KindText {
				if sp := n.Text(); sp != nil && sp.Src == title {
					found = true
					return false
				}
			}
			return true
		})
		r.add("structure", "demo 卡", found, "标题："+title)
	}
	if st.APIRows > 0 {
		want := st.APIRows
		if want > 12 { // API 表虚拟化，可见窗口约 12 行
			want = 12
		}
		got := r.countNodes(SpecTarget{Label: "api-row"})
		r.add("structure", "API 表", got >= want, fmt.Sprintf("api-row %d ≥ %d", got, want))
	}
	if st.MinElements > 0 {
		got := 0
		r.c.Scene().Walk(func(n *scene.Node) bool {
			if _, ok := interactiveLabels[n.HitSpec().Label]; ok {
				got++
			}
			return true
		})
		r.add("structure", "元素枚举", got >= st.MinElements,
			fmt.Sprintf("可交互节点 %d ≥ %d", got, st.MinElements))
	}
}

// ── 定位 ────────────────────────────────────────────────────────────

func (r *specRunner) locate(t SpecTarget) (*scene.Node, bool) {
	var matches []*scene.Node
	r.c.Scene().Walk(func(n *scene.Node) bool {
		h := n.HitSpec()
		if h.Label != t.Label {
			return true
		}
		if t.Disabled != nil && h.Disabled != *t.Disabled {
			return true
		}
		if t.Text != "" && !nodeTextContains(n, t.Text) {
			return true
		}
		matches = append(matches, n)
		return true
	})
	if len(matches) <= t.Index {
		return nil, false
	}
	return matches[t.Index], true
}

func (r *specRunner) countNodes(t SpecTarget) int {
	n := 0
	r.c.Scene().Walk(func(nd *scene.Node) bool {
		h := nd.HitSpec()
		if h.Label == t.Label {
			if t.Disabled != nil && h.Disabled != *t.Disabled {
				return true
			}
			if t.Text != "" && !nodeTextContains(nd, t.Text) {
				return true
			}
			n++
		}
		return true
	})
	return n
}

// locateNear 按中心距离就近匹配（点击改变文本后的回退定位）。
func (r *specRunner) locateNear(t SpecTarget, center geom.Point) (*scene.Node, bool) {
	var best *scene.Node
	var bestD float32
	r.c.Scene().Walk(func(n *scene.Node) bool {
		h := n.HitSpec()
		if h.Label != t.Label {
			return true
		}
		if t.Disabled != nil && h.Disabled != *t.Disabled {
			return true
		}
		d := dist(n.BoundsInRoot().Center(), center)
		if best == nil || d < bestD {
			best, bestD = n, d
		}
		return true
	})
	if best != nil && bestD < 400 { // 中心距 < 20px
		return best, true
	}
	return nil, false
}

func nodeTextContains(n *scene.Node, sub string) bool {
	found := false
	n.Walk(func(ch *scene.Node) bool {
		if found {
			return false
		}
		if sp := ch.Text(); sp != nil && strings.Contains(sp.Src, sub) {
			found = true
			return false
		}
		return true
	})
	return found
}

func targetDesc(t *SpecTarget) string {
	if t == nil {
		return "nil"
	}
	s := t.Label
	if t.Text != "" {
		s += "[" + t.Text + "]"
	}
	if t.Index > 0 {
		s += fmt.Sprintf("#%d", t.Index)
	}
	return s
}

// ensureVisible 滚动页面直到目标元素落在可视域（visMargin 内）。
func (r *specRunner) ensureVisible(t SpecTarget) bool {
	vh := r.c.Scene().Size().H
	for i := 0; i < 40; i++ {
		n, ok := r.locate(t)
		if !ok {
			return false
		}
		b := n.BoundsInRoot()
		if b.Max.Y > 48 && b.Min.Y < vh-48 {
			return true
		}
		if b.Min.Y >= vh-48 {
			injectWheel(r.c, 1035, 400, -3)
		} else {
			injectWheel(r.c, 1035, 400, 3)
		}
		time.Sleep(170 * time.Millisecond)
	}
	n, ok := r.locate(t)
	if !ok {
		return false
	}
	b := n.BoundsInRoot()
	return b.Max.Y > 48 && b.Min.Y < vh-48
}

// clickPoint 计算点击位置（目标锚 bounds 内 frac，默认中心）。
func clickPoint(b geom.Rect, frac []float32) geom.Point {
	if len(frac) != 2 {
		return geom.Point{X: (b.Min.X + b.Max.X) / 2, Y: (b.Min.Y + b.Max.Y) / 2}
	}
	return geom.Point{X: b.Min.X + b.W()*frac[0], Y: b.Min.Y + b.H()*frac[1]}
}

// ── 用例执行 ────────────────────────────────────────────────────────

func (r *specRunner) runCase(cs SpecCase) {
	name := cs.Name
	if cs.Target != nil {
		r.covered[cs.Target.Label] = true
	}
	r.closeOverlays()
	injectLeave(r.c)
	time.Sleep(120 * time.Millisecond)

	// 目标默认滚入可视域；记录初始位置（点击可能改变文本，文本锚会失效）。
	r.caseBounds = geom.Rect{}
	if cs.Target != nil {
		if !r.ensureVisible(*cs.Target) {
			r.add(name, "定位", false, "目标不可见："+targetDesc(cs.Target))
			return
		}
		if n, ok := r.locate(*cs.Target); ok {
			r.caseBounds = n.BoundsInRoot()
		}
	}
	// 前置。
	for _, g := range cs.Given {
		if g.CloseOverlays != nil && *g.CloseOverlays {
			r.closeOverlays()
		}
		if g.ScrollTo != "" {
			if !r.ensureVisible(SpecTarget{Label: g.ScrollTo}) {
				r.add(name, "定位", false, "scrollTo 不可见："+g.ScrollTo)
				return
			}
		}
		if g.Focus != nil {
			n, ok := r.locate(*g.Focus)
			if !ok {
				r.add(name, "定位", false, "focus 目标不存在："+targetDesc(g.Focus))
				return
			}
			b := n.BoundsInRoot()
			p := clickPoint(b, nil)
			injectDown(r.c, p.X, p.Y)
			injectUp(r.c, p.X, p.Y)
			time.Sleep(200 * time.Millisecond)
		}
	}
	// 基线（colorDelta / 往返对照）。
	r.start = grab(r.c)
	r.startOff = r.scrollOffset()

	// 操作。
	for _, act := range cs.When {
		if !r.doAction(name, cs.Target, act) {
			return
		}
	}
	time.Sleep(250 * time.Millisecond)

	// 断言。
	for _, as := range cs.Then {
		r.doAssert(name, cs.Target, as)
	}
	r.closeOverlays()
	injectLeave(r.c)
	time.Sleep(80 * time.Millisecond)
}

// doAction 执行一个操作；返回 false 表示前置失败（已记录）。
func (r *specRunner) doAction(caseName string, def *SpecTarget, a SpecAction) bool {
	switch {
	case a.Click != nil:
		t := def
		if a.Click.Target != nil {
			t = a.Click.Target
		}
		if t == nil {
			r.add(caseName, "操作", false, "click 缺少目标")
			return false
		}
		n, ok := r.locate(*t)
		if !ok {
			r.add(caseName, "操作", false, "click 目标不存在："+targetDesc(t))
			return false
		}
		p := clickPoint(n.BoundsInRoot(), a.Click.Frac)
		injectDown(r.c, p.X, p.Y)
		time.Sleep(60 * time.Millisecond)
		injectUp(r.c, p.X, p.Y)
	case a.Hover != nil:
		t := def
		if a.Hover.Target != nil {
			t = a.Hover.Target
		}
		if t == nil {
			r.add(caseName, "操作", false, "hover 缺少目标")
			return false
		}
		n, ok := r.locate(*t)
		if !ok {
			r.add(caseName, "操作", false, "hover 目标不存在："+targetDesc(t))
			return false
		}
		p := clickPoint(n.BoundsInRoot(), a.Hover.Frac)
		hoverAt(r.c, p.X, p.Y)
	case a.TypeText != "":
		for _, ch := range a.TypeText {
			injectChar(r.c, ch)
			time.Sleep(15 * time.Millisecond)
		}
	case a.Key != "":
		code, ok := keyCodeOf(a.Key)
		if !ok {
			r.add(caseName, "操作", false, "未知键："+a.Key)
			return false
		}
		r.c.PushEvent(&input.KeyEvent{Kind: input.KeyDown, Code: code, Time: time.Now()})
		time.Sleep(20 * time.Millisecond)
	case a.Wheel != nil:
		injectWheel(r.c, 1035, 400, *a.Wheel)
	case a.MoveOut:
		injectLeave(r.c)
	case a.Wait > 0:
		time.Sleep(time.Duration(a.Wait) * time.Millisecond)
	default:
		r.add(caseName, "操作", false, "空操作")
		return false
	}
	time.Sleep(120 * time.Millisecond)
	return true
}

func keyCodeOf(name string) (input.KeyCode, bool) {
	switch name {
	case "escape":
		return input.CodeEscape, true
	case "enter":
		return input.CodeEnter, true
	case "tab":
		return input.CodeTab, true
	case "space":
		return input.CodeSpace, true
	case "backspace":
		return input.CodeBackspace, true
	case "delete":
		return input.CodeDelete, true
	case "left":
		return input.CodeLeft, true
	case "right":
		return input.CodeRight, true
	case "up":
		return input.CodeUp, true
	case "down":
		return input.CodeDown, true
	}
	return input.CodeUnknown, false
}

// ── 断言 ────────────────────────────────────────────────────────────

func (r *specRunner) doAssert(caseName string, def *SpecTarget, as SpecAssertion) {
	target := as.Target
	if target == nil && as.Of != nil {
		target = as.Of
	}
	if target == nil {
		target = def
	}
	switch as.Kind {
	case "exists":
		need := as.Count
		if need <= 0 {
			need = 1
		}
		got := r.countNodes(*target)
		r.add(caseName, "exists", got >= need, fmt.Sprintf("%s ×%d ≥ %d", targetDesc(target), got, need))
	case "absent":
		got := r.countNodes(*target)
		r.add(caseName, "absent", got == 0, fmt.Sprintf("%s ×%d", targetDesc(target), got))
	case "geometry":
		r.assertGeometry(caseName, target, as)
	case "color":
		expect, note, ok := parseExpectedColor(as.Expect)
		if !ok {
			r.add(caseName, "color", false, "无法解析 expect")
			return
		}
		tol := 10.0
		if as.Tol != nil {
			tol = *as.Tol
		}
		img := grab(r.c)
		rect := r.sampleRect(target, as.Frac)
		avg := sampleRect(img, rect)
		want := [3]float64{float64(expect.R), float64(expect.G), float64(expect.B)}
		pass := abs3Diff(avg, want) <= tol
		r.add(caseName, "color", pass,
			fmt.Sprintf("%s 实测 %s 期望 %s(%s) Δ≤%.0f", targetDesc(target), fmtHex(avg), fmtHex(want), note, tol))
	case "colorDelta":
		minRatio := 0.02
		if as.MinRatio != nil {
			minRatio = *as.MinRatio
		}
		// tolPeak 控制通道阈值：hover 底色变化只有几个色阶（#FFF→#FAFAFA），
		// 用默认 12 会漏检，需要 3。
		thresh := uint8(12)
		if as.TolPeak != nil {
			thresh = uint8(*as.TolPeak)
		}
		img := grab(r.c)
		rect := r.sampleRect(target, as.Frac)
		ratio := regionChangedThresh(r.start, img, rect, thresh)
		r.add(caseName, "colorDelta", ratio >= minRatio,
			fmt.Sprintf("%s 变化 %.3f ≥ %.3f（阈值 %d）", targetDesc(target), ratio, minRatio, thresh))
	case "focus":
		label := "nil"
		if d := r.c.Dispatcher(); d != nil {
			if fn := d.Focus(); fn != nil {
				label = fn.HitSpec().Label
			}
		}
		want := ""
		if target != nil {
			want = target.Label
		}
		pass := label == want
		r.add(caseName, "focus", pass, fmt.Sprintf("焦点 %s == %s", label, want))
	case "overlay":
		got := r.overlayCount(as.Band)
		pass := got >= as.Count
		if as.Count == 0 {
			pass = got == 0
		}
		r.add(caseName, "overlay", pass, fmt.Sprintf("%s 浮层 %d（期望 %d）", as.Band, got, as.Count))
	case "scroll":
		got := r.scrollOffset() - r.startOff
		pass := true
		detail := fmt.Sprintf("滚动偏移变化 %.0f", got)
		if as.MinDelta != nil {
			if *as.MinDelta >= 0 {
				pass = got >= *as.MinDelta
			} else {
				pass = got <= *as.MinDelta
			}
			detail = fmt.Sprintf("滚动偏移变化 %.0f，期望 <=%.0f", got, *as.MinDelta)
		}
		r.add(caseName, "scroll", pass, detail)
	case "text":
		n, ok := r.locate(*target)
		if !ok && !r.caseBounds.IsEmpty() {
			// 点击后文案变了，文本锚失效：按初始位置就近匹配同 label 节点。
			n, ok = r.locateNear(*target, r.caseBounds.Center())
		}
		pass := ok && nodeTextContains(n, as.Text)
		detail := "文本未包含：" + as.Text
		if pass {
			detail = "文本包含：" + as.Text
		}
		r.add(caseName, "text", pass, fmt.Sprintf("%s %s", targetDesc(target), detail))
	case "template":
		r.runTemplate(caseName, def, as.Template)
	default:
		r.add(caseName, "断言", false, "未知 kind："+as.Kind)
	}
}

func (r *specRunner) assertGeometry(caseName string, target *SpecTarget, as SpecAssertion) {
	n, ok := r.locate(*target)
	if !ok {
		r.add(caseName, "geometry", false, "目标不存在："+targetDesc(target))
		return
	}
	b := n.BoundsInRoot()
	vh := r.c.Scene().Size().H
	checks := []struct {
		ok     bool
		detail string
	}{}
	if as.RightOf != nil {
		if o, ok2 := r.locate(*as.RightOf); ok2 {
			ob := o.BoundsInRoot()
			checks = append(checks, struct {
				ok     bool
				detail string
			}{b.Min.X >= ob.Max.X, fmt.Sprintf("x %.0f ≥ 右邻右缘 %.0f", b.Min.X, ob.Max.X)})
		}
	}
	if as.Below != nil {
		if o, ok2 := r.locate(*as.Below); ok2 {
			ob := o.BoundsInRoot()
			checks = append(checks, struct {
				ok     bool
				detail string
			}{b.Min.Y >= ob.Max.Y-1, fmt.Sprintf("y %.0f ≥ 上邻底缘 %.0f", b.Min.Y, ob.Max.Y)})
		}
	}
	if as.WidthAtLeast > 0 {
		checks = append(checks, struct {
			ok     bool
			detail string
		}{b.W() >= as.WidthAtLeast, fmt.Sprintf("宽 %.0f ≥ %.0f", b.W(), as.WidthAtLeast)})
	}
	if len(as.HeightBetween) == 2 {
		checks = append(checks, struct {
			ok     bool
			detail string
		}{b.H() >= as.HeightBetween[0] && b.H() <= as.HeightBetween[1],
			fmt.Sprintf("高 %.0f ∈ [%.0f,%.0f]", b.H(), as.HeightBetween[0], as.HeightBetween[1])})
	}
	if len(as.CenterYWithin) == 2 {
		cy := (b.Min.Y + b.Max.Y) / 2
		checks = append(checks, struct {
			ok     bool
			detail string
		}{cy >= as.CenterYWithin[0] && cy <= as.CenterYWithin[1],
			fmt.Sprintf("中心 y %.0f ∈ [%.0f,%.0f]", cy, as.CenterYWithin[0], as.CenterYWithin[1])})
	}
	if as.InViewport {
		checks = append(checks, struct {
			ok     bool
			detail string
		}{b.Max.Y > 0 && b.Min.Y < vh, fmt.Sprintf("bounds %v 在视口 [0,%.0f]", b, vh)})
	}
	if len(checks) == 0 {
		r.add(caseName, "geometry", false, "geometry 无可判定项")
		return
	}
	for _, c := range checks {
		r.add(caseName, "geometry", c.ok, c.detail)
	}
}

// ── 交互模板（§5）──────────────────────────────────────────────────

func (r *specRunner) runTemplate(caseName string, def *SpecTarget, tpl string) {
	if def == nil {
		r.add(caseName, "template:"+tpl, false, "模板缺少 target")
		return
	}
	switch tpl {
	case "toggle":
		r.tplToggle(caseName, *def)
	case "disabled-no-op":
		r.tplDisabledNoOp(caseName, *def)
	case "input-focus":
		r.tplInputFocus(caseName, *def)
	case "input-type":
		r.tplInputType(caseName, *def)
	case "table-row-select":
		r.tplTableRowSelect(caseName, *def)
	case "select-open-close":
		r.tplSelectOpenClose(caseName, *def)
	case "slider-jump":
		r.tplSliderJump(caseName, *def)
	case "rate-set":
		r.tplRateSet(caseName, *def)
	case "segment-select":
		r.tplSegmentSelect(caseName, *def)
	case "tree-node-select":
		r.tplTreeNodeSelect(caseName, *def)
	case "collapse-toggle":
		r.tplCollapseToggle(caseName, *def)
	default:
		r.add(caseName, "template:"+tpl, false, "未实现的模板")
	}
}

// tplToggle：切换后视觉变化，再点同一位置切回后画面复原（验证 toggle 语义）。
// 第二次点击用首次的位置：文案开关切换后文本变了（"启用"→"停用"），
// 按文本重新定位会失败，但位置是稳定的。
func (r *specRunner) tplToggle(caseName string, t SpecTarget) {
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:toggle", false, "目标不存在")
		return
	}
	rect := n.BoundsInRoot()
	before := grab(r.c)
	avgBefore := sampleRect(before, rect)

	p := clickPoint(rect, nil)
	injectDown(r.c, p.X, p.Y)
	time.Sleep(60 * time.Millisecond)
	injectUp(r.c, p.X, p.Y)
	time.Sleep(280 * time.Millisecond)
	mid := grab(r.c)
	ratio := regionChanged(before, mid, rect)
	r.add(caseName, "template:toggle", ratio >= 0.02,
		fmt.Sprintf("切换变化 %.3f ≥ 0.020", ratio))

	// 再点一次：状态回到原样。指针先移出再比：悬停态会让"复原"差出色差。
	injectDown(r.c, p.X, p.Y)
	time.Sleep(60 * time.Millisecond)
	injectUp(r.c, p.X, p.Y)
	injectLeave(r.c)
	time.Sleep(300 * time.Millisecond)
	back := grab(r.c)
	avgBack := sampleRect(back, rect)
	pass := abs3Diff(avgBefore, avgBack) <= 12
	r.add(caseName, "template:toggle", pass,
		fmt.Sprintf("往返复原 Δ=%.1f ≤ 12", abs3Diff(avgBefore, avgBack)))
}

// tplDisabledNoOp：禁区点击前后无任何变化（反向断言）。
func (r *specRunner) tplDisabledNoOp(caseName string, t SpecTarget) {
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:disabled-no-op", false, "目标不存在")
		return
	}
	rect := n.BoundsInRoot()
	before := grab(r.c)
	p := clickPoint(rect, nil)
	injectDown(r.c, p.X, p.Y)
	time.Sleep(60 * time.Millisecond)
	injectUp(r.c, p.X, p.Y)
	time.Sleep(280 * time.Millisecond)
	after := grab(r.c)
	region := regionChangedThresh(before, after, rect, 3)
	frame := frameChanged(before, after)
	pass := region == 0 && frame < 0.0005
	r.add(caseName, "template:disabled-no-op", pass,
		fmt.Sprintf("区域变化 %.3f==0 全帧 %.4f<0.0005", region, frame))
}

// tplInputFocus：点击后焦点归属该输入框。
func (r *specRunner) tplInputFocus(caseName string, t SpecTarget) {
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:input-focus", false, "目标不存在")
		return
	}
	p := clickPoint(n.BoundsInRoot(), nil)
	injectDown(r.c, p.X, p.Y)
	time.Sleep(60 * time.Millisecond)
	injectUp(r.c, p.X, p.Y)
	time.Sleep(300 * time.Millisecond)
	label := "nil"
	if d := r.c.Dispatcher(); d != nil {
		if fn := d.Focus(); fn != nil {
			label = fn.HitSpec().Label
		}
	}
	r.add(caseName, "template:input-focus", label == t.Label,
		fmt.Sprintf("焦点 %s == %s", label, t.Label))
}

// tplInputType：聚焦后键入，文本出现在输入框内。
func (r *specRunner) tplInputType(caseName string, t SpecTarget) {
	r.tplInputFocus(caseName, t)
	for _, ch := range "wg" {
		injectChar(r.c, ch)
		time.Sleep(20 * time.Millisecond)
	}
	time.Sleep(300 * time.Millisecond)
	n, ok := r.locate(t)
	pass := ok && nodeTextContains(n, "wg")
	r.add(caseName, "template:input-type", pass, "输入框文本包含 wg")
}

// tplTableRowSelect：点击行后行背景 ≈ token Selection。
func (r *specRunner) tplTableRowSelect(caseName string, t SpecTarget) {
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:table-row-select", false, "目标不存在")
		return
	}
	p := clickPoint(n.BoundsInRoot(), nil)
	injectDown(r.c, p.X, p.Y)
	time.Sleep(60 * time.Millisecond)
	injectUp(r.c, p.X, p.Y)
	time.Sleep(300 * time.Millisecond)
	n2, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:table-row-select", false, "重建后目标不存在")
		return
	}
	// 采样行顶部条带避开行内文字（文字垂直居中；纯背景才可比 token）。
	avg := sampleRect(grab(r.c), nodeSampleRect(n2, []float32{0.10, 0.05, 0.80, 0.12}))
	want, ok := tokenColor("Selection")
	tol := 18.0
	pass := ok && abs3Diff(avg, [3]float64{float64(want.R), float64(want.G), float64(want.B)}) <= tol
	r.add(caseName, "template:table-row-select", pass,
		fmt.Sprintf("行背景 %s ≈ token Selection", fmtHex(avg)))
}

// tplSelectOpenClose：点击开浮层、Esc 关、再次点击能重新打开（覆盖句柄复位）。
func (r *specRunner) tplSelectOpenClose(caseName string, t SpecTarget) {
	open := func() int {
		n, ok := r.locate(t)
		if !ok {
			return -1
		}
		p := clickPoint(n.BoundsInRoot(), nil)
		injectDown(r.c, p.X, p.Y)
		injectUp(r.c, p.X, p.Y)
		time.Sleep(300 * time.Millisecond)
		return r.overlayCount("")
	}
	base := r.overlayCount("")
	afterOpen := open()
	r.add(caseName, "template:select-open-close", afterOpen == base+1,
		fmt.Sprintf("点击后浮层 %d → %d", base, afterOpen))
	injectEsc(r.c)
	time.Sleep(250 * time.Millisecond)
	afterEsc := r.overlayCount("")
	r.add(caseName, "template:select-open-close", afterEsc == base,
		fmt.Sprintf("Esc 后浮层 %d → %d", afterOpen, afterEsc))
	afterReopen := open()
	r.add(caseName, "template:select-open-close", afterReopen == base+1,
		fmt.Sprintf("再次点击浮层 → %d（句柄复位）", afterReopen))
	injectEsc(r.c)
	time.Sleep(150 * time.Millisecond)
}

// ── 工具 ────────────────────────────────────────────────────────────

func (r *specRunner) closeOverlays() {
	ov := r.c.Overlay()
	if ov == nil || ov.Count() == 0 {
		return
	}
	injectEsc(r.c)
	time.Sleep(120 * time.Millisecond)
	if ov.Count() > 0 {
		ov.CloseAll(theme.BandPopup)
		ov.CloseAll(theme.BandTooltip)
		ov.CloseAll(theme.BandModal)
		ov.CloseAll(theme.BandToast)
		time.Sleep(120 * time.Millisecond)
	}
}

func (r *specRunner) overlayCount(band string) int {
	ov := r.c.Overlay()
	if ov == nil {
		return 0
	}
	if band == "" {
		return ov.Count()
	}
	want := bandOf(band)
	n := 0
	for _, e := range ov.Entries() {
		if e.Band == want {
			n++
		}
	}
	return n
}

func bandOf(name string) theme.Band {
	switch name {
	case "popup":
		return theme.BandPopup
	case "tooltip":
		return theme.BandTooltip
	case "modal":
		return theme.BandModal
	case "toast":
		return theme.BandToast
	}
	return theme.BandPopup
}

func (r *specRunner) scrollOffset() float32 {
	var off, area float32
	r.c.Scene().Walk(func(n *scene.Node) bool {
		o := n.ScrollOffset()
		if o == 0 {
			return true
		}
		if b := n.BoundsInRoot(); b.W()*b.H() > area {
			area = b.W() * b.H()
			off = o
		}
		return true
	})
	return off
}

// sampleRect 计算元素内采样区域（frac 为该元素内的相对矩形，默认整元素）。
func (r *specRunner) sampleRect(t *SpecTarget, frac []float32) geom.Rect {
	n, ok := r.locate(*t)
	if !ok {
		return geom.Rect{}
	}
	return nodeSampleRect(n, frac)
}

func nodeSampleRect(n *scene.Node, frac []float32) geom.Rect {
	b := n.BoundsInRoot()
	if len(frac) != 4 {
		return b
	}
	return geom.RectXYWH(b.Min.X+b.W()*frac[0], b.Min.Y+b.H()*frac[1], b.W()*frac[2], b.H()*frac[3])
}

func sampleRect(img *image.RGBA, r geom.Rect) [3]float64 {
	if img == nil || r.IsEmpty() {
		return [3]float64{}
	}
	x0, y0 := int(r.Min.X), int(r.Min.Y)
	x1, y1 := int(r.Max.X), int(r.Max.Y)
	bounds := img.Bounds()
	if x0 < bounds.Min.X {
		x0 = bounds.Min.X
	}
	if y0 < bounds.Min.Y {
		y0 = bounds.Min.Y
	}
	if x1 > bounds.Max.X {
		x1 = bounds.Max.X
	}
	if y1 > bounds.Max.Y {
		y1 = bounds.Max.Y
	}
	if x1 <= x0 || y1 <= y0 {
		return [3]float64{}
	}
	var sum [3]int
	n := 0
	for y := y0; y < y1; y++ {
		i := y*img.Stride + x0*4
		for x := x0; x < x1; x++ {
			sum[0] += int(img.Pix[i])
			sum[1] += int(img.Pix[i+1])
			sum[2] += int(img.Pix[i+2])
			n++
			i += 4
		}
	}
	return [3]float64{float64(sum[0]) / float64(n), float64(sum[1]) / float64(n), float64(sum[2]) / float64(n)}
}

func abs3Diff(a, b [3]float64) float64 {
	peak := 0.0
	for i := 0; i < 3; i++ {
		d := a[i] - b[i]
		if d < 0 {
			d = -d
		}
		if d > peak {
			peak = d
		}
	}
	return peak
}

func fmtHex(c [3]float64) string {
	return fmt.Sprintf("#%02X%02X%02X", int(c[0]+0.5), int(c[1]+0.5), int(c[2]+0.5))
}

// parseExpectedColor 解析 {"token": name} / {"hex": "#RRGGBB", "note": ...} / "#RRGGBB"。
func parseExpectedColor(v any) (geom.Color, string, bool) {
	switch e := v.(type) {
	case string:
		return hexColor(e)
	case map[string]any:
		if tok, ok := e["token"].(string); ok {
			col, ok2 := tokenColor(tok)
			return col, "token:" + tok, ok2
		}
		if hex, ok := e["hex"].(string); ok {
			col, hnote, ok2 := hexColor(hex)
			note, _ := e["note"].(string)
			if note == "" {
				note = hnote
			}
			return col, note, ok2
		}
	}
	return geom.Color{}, "", false
}

func hexColor(s string) (geom.Color, string, bool) {
	s = strings.TrimPrefix(s, "#")
	if len(s) != 6 {
		return geom.Color{}, "", false
	}
	var r, g, b uint8
	_, err := fmt.Sscanf(s, "%02X%02X%02X", &r, &g, &b)
	if err != nil {
		return geom.Color{}, "", false
	}
	return geom.RGBA(r, g, b), "#" + s, true
}

// tokenColor 按字段名解析主题色板（定制主题的窗应改用 hex + note）。
func tokenColor(name string) (geom.Color, bool) {
	tokens := theme.Default().Color
	f := reflect.ValueOf(tokens).FieldByName(name)
	if !f.IsValid() {
		return geom.Color{}, false
	}
	col, ok := f.Interface().(geom.Color)
	return col, ok
}

// hoverAt 把指针移入指定位置（无按键移动）。
func hoverAt(c *wg.Container, x, y float32) {
	p := geom.Pt(x, y)
	c.PushEvent(&input.PointerEvent{
		Kind: input.PointerHover, Pointer: input.MousePointer, Device: input.DeviceMouse,
		Pos: p, PosRoot: p, Time: time.Now(),
	})
}

// tplSliderJump：点击轨道取值（antd 语义），值已等于点击位置时换位置再验。
func (r *specRunner) tplSliderJump(caseName string, t SpecTarget) {
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:slider-jump", false, "目标不存在")
		return
	}
	rect := n.BoundsInRoot()
	before := grab(r.c)
	try := func(frac float64) float64 {
		p := clickPoint(rect, []float32{float32(frac), 0.5})
		injectDown(r.c, p.X, p.Y)
		time.Sleep(60 * time.Millisecond)
		injectUp(r.c, p.X, p.Y)
		time.Sleep(300 * time.Millisecond)
		return regionChanged(before, grab(r.c), rect)
	}
	ratio := try(0.25)
	if ratio < 0.02 {
		ratio = try(0.90)
	}
	r.add(caseName, "template:slider-jump", ratio >= 0.02,
		fmt.Sprintf("取值变化 %.3f ≥ 0.020", ratio))
}

// tplRateSet：点击第 k 星后该星为 Warning 色、下一颗为 SurfaceActive。
func (r *specRunner) tplRateSet(caseName string, t SpecTarget) {
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:rate-set", false, "目标不存在")
		return
	}
	p := clickPoint(n.BoundsInRoot(), nil)
	injectDown(r.c, p.X, p.Y)
	time.Sleep(60 * time.Millisecond)
	injectUp(r.c, p.X, p.Y)
	injectLeave(r.c)
	time.Sleep(300 * time.Millisecond)

	img := grab(r.c)
	picked := func(n2 *scene.Node) [3]float64 {
		return sampleRect(img, nodeSampleRect(n2, []float32{0.35, 0.35, 0.30, 0.30}))
	}
	n2, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:rate-set", false, "重建后目标不存在")
		return
	}
	warn, ok1 := tokenColor("Warning")
	active, ok2 := tokenColor("SurfaceActive")
	if !ok1 || !ok2 {
		r.add(caseName, "template:rate-set", false, "主题 token 缺失")
		return
	}
	avg := picked(n2)
	pass := abs3Diff(avg, [3]float64{float64(warn.R), float64(warn.G), float64(warn.B)}) <= 48
	r.add(caseName, "template:rate-set", pass,
		fmt.Sprintf("被点星 %s ≈ token Warning", fmtHex(avg)))

	next := SpecTarget{Label: t.Label, Disabled: t.Disabled, Index: t.Index + 1}
	if n3, ok := r.locate(next); ok {
		avg2 := picked(n3)
		pass2 := abs3Diff(avg2, [3]float64{float64(active.R), float64(active.G), float64(active.B)}) <= 48
		r.add(caseName, "template:rate-set", pass2,
			fmt.Sprintf("下一颗星 %s ≈ token SurfaceActive", fmtHex(avg2)))
	}
}

// tplSegmentSelect：点击未选中项选中迁移；点已选中项则改点同组邻居。
func (r *specRunner) tplSegmentSelect(caseName string, t SpecTarget) {
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:segment-select", false, "目标不存在")
		return
	}
	rect := n.BoundsInRoot()
	before := grab(r.c)
	p := clickPoint(rect, nil)
	injectDown(r.c, p.X, p.Y)
	time.Sleep(60 * time.Millisecond)
	injectUp(r.c, p.X, p.Y)
	time.Sleep(300 * time.Millisecond)
	ratio := regionChanged(before, grab(r.c), rect)
	if ratio >= 0.02 {
		r.add(caseName, "template:segment-select", true, fmt.Sprintf("点击选中变化 %.3f", ratio))
		return
	}
	// 已选中项点击无变化属正确：改点同组邻居，验证指示迁移。
	off := 0
	if t.Index == 0 {
		off = 1
	}
	neighbor := SpecTarget{Label: t.Label, Disabled: t.Disabled, Index: off}
	if _, ok := r.locate(neighbor); ok {
		// 邻居可能在另一屏：先滚入可视域再点（跨 demo 的同类元素）。
		r.ensureVisible(neighbor)
		if n2, ok2 := r.locate(neighbor); ok2 {
			p2 := clickPoint(n2.BoundsInRoot(), nil)
			injectDown(r.c, p2.X, p2.Y)
			time.Sleep(60 * time.Millisecond)
			injectUp(r.c, p2.X, p2.Y)
			time.Sleep(300 * time.Millisecond)
			frame := frameChanged(before, grab(r.c))
			r.add(caseName, "template:segment-select", frame >= 0.0005,
				fmt.Sprintf("邻居切换全帧变化 %.4f ≥ 0.0005", frame))
			return
		}
	}
	r.add(caseName, "template:segment-select", false, "无可用邻居")
}

// tplTreeNodeSelect：点击节点后节点背景 ≈ token Selection（同表格行语义）。
func (r *specRunner) tplTreeNodeSelect(caseName string, t SpecTarget) {
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:tree-node-select", false, "目标不存在")
		return
	}
	p := clickPoint(n.BoundsInRoot(), nil)
	injectDown(r.c, p.X, p.Y)
	time.Sleep(60 * time.Millisecond)
	injectUp(r.c, p.X, p.Y)
	injectLeave(r.c)
	time.Sleep(300 * time.Millisecond)
	n2, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:tree-node-select", false, "重建后目标不存在")
		return
	}
	avg := sampleRect(grab(r.c), nodeSampleRect(n2, []float32{0.05, 0.05, 0.90, 0.15}))
	want, ok := tokenColor("Selection")
	pass := ok && abs3Diff(avg, [3]float64{float64(want.R), float64(want.G), float64(want.B)}) <= 20
	r.add(caseName, "template:tree-node-select", pass,
		fmt.Sprintf("节点背景 %s ≈ token Selection", fmtHex(avg)))
}

// tplCollapseToggle：点击表头展开/收起，内容区大面积变化。
func (r *specRunner) tplCollapseToggle(caseName string, t SpecTarget) {
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:collapse-toggle", false, "目标不存在")
		return
	}
	before := grab(r.c)
	p := clickPoint(n.BoundsInRoot(), nil)
	injectDown(r.c, p.X, p.Y)
	time.Sleep(60 * time.Millisecond)
	injectUp(r.c, p.X, p.Y)
	injectLeave(r.c)
	time.Sleep(400 * time.Millisecond)
	frame := frameChanged(before, grab(r.c))
	r.add(caseName, "template:collapse-toggle", frame >= 0.001,
		fmt.Sprintf("展开/收起全帧变化 %.4f ≥ 0.001", frame))
}
