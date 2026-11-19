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
	WidthBetween  []float32   `json:"widthBetween,omitempty"`
	HeightBetween []float32   `json:"heightBetween,omitempty"`
	CenterYWithin []float32   `json:"centerYWithin,omitempty"`
	InViewport    bool        `json:"inViewport,omitempty"`

	// overlay / scroll / text
	Band        string   `json:"band,omitempty"`
	Delta       *float32 `json:"delta,omitempty"`
	MinDelta    *float32 `json:"minDelta,omitempty"`
	Text        string   `json:"text,omitempty"`
	NotContains bool     `json:"notContains,omitempty"` // text 断言的反向模式：断言节点子树不含该文本

	// borderWidth
	ExpectWidth int `json:"expectWidth,omitempty"`
	// centeredIn/sideIn 的参照锚
	Ref *SpecTarget `json:"ref,omitempty"`
	// Side 是 sideIn 的方位：right/left（被测元素中心相对参照中心的水平侧）
	Side string `json:"side,omitempty"`
	// GapExpect 是 gapBetween 的期望水平间距（默认 8）
	GapExpect *float64 `json:"gapExpect,omitempty"`
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
		r.c.Walk(func(n *scene.Node) bool {
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
		r.c.Walk(func(n *scene.Node) bool {
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
	r.c.Walk(func(n *scene.Node) bool {
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
	r.c.Walk(func(nd *scene.Node) bool {
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
	r.c.Walk(func(n *scene.Node) bool {
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
	// 拼接子树全部文本再判断：两字中文按钮被拆成两个 Text 节点
	//（antd letter-spacing 的逐字排版），单节点匹配会漏。
	var sb strings.Builder
	n.Walk(func(ch *scene.Node) bool {
		if sp := ch.Text(); sp != nil {
			sb.WriteString(sp.Src)
		}
		return true
	})
	return strings.Contains(sb.String(), sub)
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

// allAssertionsNotContains 判断一个用例的全部断言是否都是 text 的
// notContains 模式——这类用例的目标节点不存在本身就是通过态。
func allAssertionsNotContains(then []SpecAssertion) bool {
	for _, a := range then {
		if a.Kind != "text" || !a.NotContains {
			return false
		}
	}
	return len(then) > 0
}

func (r *specRunner) runCase(cs SpecCase) {
	name := cs.Name
	if cs.Target != nil {
		r.covered[cs.Target.Label] = true
	}
	r.closeOverlays()
	injectLeave(r.c)
	time.Sleep(120 * time.Millisecond)

	// 目标默认滚入可视域；记录初始位置（点击可能改变文本，文本锚会失效）。
	// 纯 notContains 用例外（折叠收起后内容节点不存在本身就是通过态，
	// ensureVisible 的硬失败会把它误判为定位错误）。
	allNotContains := cs.Target != nil && cs.Target.Label != "" &&
		len(cs.Then) > 0 && allAssertionsNotContains(cs.Then)
	r.caseBounds = geom.Rect{}
	if cs.Target != nil {
		if !allNotContains && !r.ensureVisible(*cs.Target) {
			r.add(name, "定位", false, "目标不可见："+targetDesc(cs.Target))
			return
		}
		if n, ok := r.locate(*cs.Target); ok {
			r.caseBounds = n.BoundsInRoot()
		}
		// 落定重验：滚轮惯性/延迟重建会让"bounds 判定可见"与"实际呈现的
		// 像素"短暂错位（采样矩形对上一帧像素 → 像素断言随机 0.000）。
		// bounds 连续两次一致才算落定；漂移就重滚。
		for i := 0; i < 4; i++ {
			time.Sleep(150 * time.Millisecond)
			n2, ok2 := r.locate(*cs.Target)
			if !ok2 {
				if allNotContains {
					break // 目标不存在=notContains 的通过态，交断言处理
				}
				return // 重建后目标消失，走常规失败路径
			}
			b2 := n2.BoundsInRoot()
			if b2 == r.caseBounds {
				break
			}
			r.caseBounds = b2
			if !r.ensureVisible(*cs.Target) {
				r.add(name, "定位", false, "落定后目标不可见："+targetDesc(cs.Target))
				return
			}
			if n3, ok3 := r.locate(*cs.Target); ok3 {
				r.caseBounds = n3.BoundsInRoot()
			}
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
		// hover 重建落帧等待：高负载（全量串行回归）下 hover 置位→重建→
		// 呈现的链路偶发超过 200ms，固定 wait 后采样仍读到旧帧（borderColor
		// 读到常态色 / colorDelta 读到 0.000，三次全量三个窗轮番 flake）。
		// 等 PresentStats 帧号增长（hover 高亮必然触发一次重绘）再返回。
		f0 := r.c.PresentStats().Frames
		deadline := time.Now().Add(600 * time.Millisecond)
		for r.c.PresentStats().Frames <= f0 && time.Now().Before(deadline) {
			time.Sleep(20 * time.Millisecond)
		}
		time.Sleep(250 * time.Millisecond) // 帧呈现到栅格化完成的余量（高负载下 80ms 偶发不够）
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
	time.Sleep(150 * time.Millisecond)
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
	case "borderWidth":
		img := grab(r.c)
		rect := r.sampleRect(target, as.Frac)
		got := borderWidthOf(img, rect)
		tol := 1
		if as.Tol != nil {
			tol = int(*as.Tol)
		}
		pass := got >= as.ExpectWidth-tol && got <= as.ExpectWidth+tol
		r.add(caseName, "borderWidth", pass,
			fmt.Sprintf("%s 边框 %dpx（期望 %d±%d）", targetDesc(target), got, as.ExpectWidth, tol))
	case "borderColor":
		img := grab(r.c)
		at := target
		if as.Target != nil {
			at = as.Target // 断言级覆盖：交互目标与采样目标不同（hover 打在行上、采盒子）
		}
		rect := r.sampleRect(at, as.Frac)
		expect, note, ok := parseExpectedColor(as.Expect)
		if !ok {
			r.add(caseName, "borderColor", false, "无法解析 expect")
			return
		}
		tol := 24.0
		if as.Tol != nil {
			tol = *as.Tol
		}
		got, found := borderColorOf(img, rect)
		// frac 带采样对 1px 边框不可靠：边框行随设备像素对齐落在带内不同
		// 行，整数截断还会把 crisp 行裁掉。扫描线找"与内部参考色差异最大
		// 的行"=边框行（光晕/阴影行差异小），取该行横向前 10%..90% 均色。
		pass := found && abs3Diff(got, [3]float64{float64(expect.R), float64(expect.G), float64(expect.B)}) <= tol
		r.add(caseName, "borderColor", pass,
			fmt.Sprintf("%s@rect(%.1f,%.1f %.1fx%.1f) 边框行色 ≈ %s（实测 %s Δ%.0f ≤ %.0f）", targetDesc(target),
				rect.Min.X, rect.Min.Y, rect.W(), rect.H(), note, fmtHex(got),
				abs3Diff(got, [3]float64{float64(expect.R), float64(expect.G), float64(expect.B)}), tol))
	case "centeredIn":
		a, ok := r.locate(*target)
		if !ok {
			r.add(caseName, "centeredIn", false, "被测元素不存在："+targetDesc(target))
			return
		}
		if as.Ref == nil {
			r.add(caseName, "centeredIn", false, "缺少 ref 参照")
			return
		}
		b, ok := r.locate(*as.Ref)
		if !ok {
			r.add(caseName, "centeredIn", false, "参照元素不存在："+targetDesc(as.Ref))
			return
		}
		ba, bb := a.BoundsInRoot(), b.BoundsInRoot()
		tol := float32(2)
		if as.Tol != nil {
			tol = float32(*as.Tol)
		}
		dy := ba.Center().Y - bb.Center().Y
		if dy < 0 {
			dy = -dy
		}
		insideX := ba.Min.X >= bb.Min.X-1 && ba.Max.X <= bb.Max.X+1
		pass := dy <= tol && insideX
		r.add(caseName, "centeredIn", pass,
			fmt.Sprintf("%s 垂直中心差 %.1f ≤ %.1f 且水平在 %s 内=%v",
				targetDesc(target), dy, tol, targetDesc(as.Ref), insideX))
	case "sideIn":
		a, ok := r.locate(*target)
		if !ok || as.Ref == nil {
			r.add(caseName, "sideIn", false, "被测/参照元素缺失")
			return
		}
		bb, ok := r.locate(*as.Ref)
		if !ok {
			r.add(caseName, "sideIn", false, "参照元素不存在："+targetDesc(as.Ref))
			return
		}
		ax := a.BoundsInRoot().Center().X
		bx := bb.BoundsInRoot().Center().X
		pass := (as.Side == "right" && ax > bx) || (as.Side == "left" && ax < bx)
		r.add(caseName, "sideIn", pass,
			fmt.Sprintf("%s 中心 x=%.0f，%s %s 中心 x=%.0f", targetDesc(target), ax,
				as.Side, targetDesc(as.Ref), bx))
	case "vAlign":
		a, ok := r.locate(*target)
		if !ok || as.Ref == nil {
			r.add(caseName, "vAlign", false, "被测/参照元素缺失")
			return
		}
		bb, ok := r.locate(*as.Ref)
		if !ok {
			r.add(caseName, "vAlign", false, "参照元素不存在："+targetDesc(as.Ref))
			return
		}
		dy := a.BoundsInRoot().Center().Y - bb.BoundsInRoot().Center().Y
		if dy < 0 {
			dy = -dy
		}
		tol := float32(2)
		if as.Tol != nil {
			tol = float32(*as.Tol)
		}
		r.add(caseName, "vAlign", dy <= tol,
			fmt.Sprintf("%s 与 %s 垂直中心差 %.1f ≤ %.1f", targetDesc(target), targetDesc(as.Ref), dy, tol))
	case "gapBetween":
		a, ok := r.locate(*target)
		if !ok || as.Ref == nil {
			r.add(caseName, "gapBetween", false, "被测/参照元素缺失")
			return
		}
		bb, ok := r.locate(*as.Ref)
		if !ok {
			r.add(caseName, "gapBetween", false, "参照元素不存在："+targetDesc(as.Ref))
			return
		}
		ba, bbb := a.BoundsInRoot(), bb.BoundsInRoot()
		gap := bbb.Min.X - ba.Max.X // 被测在左
		if gap < 0 {
			gap = ba.Min.X - bbb.Max.X // 被测在右
		}
		want := 8.0
		if as.GapExpect != nil {
			want = *as.GapExpect
		}
		tol := 1.5
		if as.Tol != nil {
			tol = *as.Tol
		}
		d := float64(gap) - want
		if d < 0 {
			d = -d
		}
		r.add(caseName, "gapBetween", d <= tol,
			fmt.Sprintf("%s 与 %s 间距 %.1f（期望 %.1f±%.1f）", targetDesc(target), targetDesc(as.Ref), gap, want, tol))
	case "colorRatio":
		expect, note, ok := parseExpectedColor(as.Expect)
		if !ok {
			r.add(caseName, "colorRatio", false, "无法解析 expect")
			return
		}
		minRatio := 0.05
		if as.MinRatio != nil {
			minRatio = *as.MinRatio
		}
		tol := 24.0
		if as.Tol != nil {
			tol = *as.Tol
		}
		at := target
		if as.Target != nil {
			at = as.Target // 断言级覆盖：case 锚保证可见性、采样打别的节点
		}
		img := grab(r.c)
		rect := r.sampleRect(at, as.Frac)
		ratio := colorRatioOf(img, rect, expect, tol)
		r.add(caseName, "colorRatio", ratio >= minRatio,
			fmt.Sprintf("%s %s 占比 %.3f ≥ %.3f", targetDesc(target), note, ratio, minRatio))
	case "text":
		n, ok := r.locate(*target)
		if !ok && !r.caseBounds.IsEmpty() {
			// 点击后文案变了，文本锚失效：按初始位置就近匹配同 label 节点。
			n, ok = r.locateNear(*target, r.caseBounds.Center())
		}
		contains := ok && nodeTextContains(n, as.Text)
		if as.NotContains {
			// 反向断言：子树必须**不含**该文本（折叠收起后内容节点不存在）。
			// 节点定位不到本身就含"内容缺失"的语义（内容已随收起消失），
			// 不算定位错误——ensureVisible 已在采样前跑过。
			pass := !contains
			detail := fmt.Sprintf("文本不含：%s", as.Text)
			if contains {
				detail = fmt.Sprintf("文本仍包含：%s", as.Text)
			}
			r.add(caseName, "text", pass, fmt.Sprintf("%s %s", targetDesc(target), detail))
			return
		}
		pass := contains
		detail := "文本未包含：" + as.Text
		if pass {
			detail = "文本包含：" + as.Text
		}
		r.add(caseName, "text", pass, fmt.Sprintf("%s %s", targetDesc(target), detail))
	case "template":
		r.runTemplate(caseName, def, as)
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
	if len(as.WidthBetween) == 2 {
		checks = append(checks, struct {
			ok     bool
			detail string
		}{b.W() >= as.WidthBetween[0] && b.W() <= as.WidthBetween[1],
			fmt.Sprintf("宽 %.0f ∈ [%.0f,%.0f]", b.W(), as.WidthBetween[0], as.WidthBetween[1])})
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

func (r *specRunner) runTemplate(caseName string, def *SpecTarget, as SpecAssertion) {
	if def == nil {
		r.add(caseName, "template:"+as.Template, false, "模板缺少 target")
		return
	}
	switch as.Template {
	case "toggle":
		r.tplToggle(caseName, *def)
	case "disabled-no-op":
		r.tplDisabledNoOp(caseName, *def)
	case "input-focus":
		r.tplInputFocus(caseName, *def)
	case "input-type":
		r.tplInputType(caseName, *def)
	case "autocomplete-select":
		r.tplAutocompleteSelect(caseName, *def, as)
	case "tooltip-dark":
		r.tplTooltipDark(caseName, *def)
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
	case "button-press":
		r.tplButtonPress(caseName, *def)
	case "spinner-spin":
		r.tplSpinnerSpin(caseName, *def, as)
	default:
		r.add(caseName, "template:"+as.Template, false, "未实现的模板")
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
	// 主判据=目标区域零变化；全帧阈值 0.002 容纳滚动条 thumb/页面级微噪
	//（带滚动条的页面整帧 3px 步进采样会扫到非零差异，实测 0.0015）。
	pass := region == 0 && frame < 0.002
	r.add(caseName, "template:disabled-no-op", pass,
		fmt.Sprintf("区域变化 %.3f==0 全帧 %.4f<0.002", region, frame))
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

// tplTooltipDark：悬停触发元素后，弹层（BandTooltip）应为深色底
// （antd colorBgSpotlight）——在锚点周边 ±160px 区域统计深色像素占比。
func (r *specRunner) tplTooltipDark(caseName string, t SpecTarget) {
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:tooltip-dark", false, "目标不存在")
		return
	}
	b := n.BoundsInRoot()
	p := clickPoint(b, nil)
	hoverAt(r.c, p.X, p.Y)
	time.Sleep(900 * time.Millisecond)
	// 弹层横向不超出锚点太多（默认 Top 居中/Bottom）：只统计锚点正上
	// 40px + 正下 40px 两条带（弹层必然落在其中之一），避免大面积白底
	// 页面稀释占比。
	above := geom.RectXYWH(b.Min.X, b.Min.Y-40, b.W(), 40)
	below := geom.RectXYWH(b.Min.X, b.Max.Y, b.W(), 40)
	img := grab(r.c)
	dark, total := 0, 0
	for _, region := range []geom.Rect{above, below} {
		x0, y0 := int(region.Min.X), int(region.Min.Y)
		x1, y1 := int(region.Max.X), int(region.Max.Y)
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
		for y := y0; y < y1; y++ {
			i := y*img.Stride + x0*4
			for x := x0; x < x1; x++ {
				lum := (int(img.Pix[i])*299 + int(img.Pix[i+1])*587 + int(img.Pix[i+2])*114) / 1000
				total++
				if lum < 80 {
					dark++
				}
				i += 4
			}
		}
	}
	ratio := 0.0
	if total > 0 {
		ratio = float64(dark) / float64(total)
	}
	pass := ratio >= 0.15
	r.add(caseName, "template:tooltip-dark", pass,
		fmt.Sprintf("锚点周边深色像素占比 %.3f ≥ 0.15（colorBgSpotlight）", ratio))
	injectLeave(r.c)
	time.Sleep(200 * time.Millisecond)
}

// tplAutocompleteSelect：聚焦输入过滤词 → 建议弹层出现 → 点第一条建议 →
// 输入框回填完整建议文本（antd AutoComplete 核心交互链）。
// 过滤词与期望回填取自断言的 text 字段（格式 "过滤词=期望回填"；
// 未带 = 时过滤词=期望文本）——target.text 只作定位锚。
func (r *specRunner) tplAutocompleteSelect(caseName string, t SpecTarget, as SpecAssertion) {
	keyword, expect := as.Text, as.Text
	if i := strings.Index(as.Text, "="); i > 0 {
		keyword, expect = as.Text[:i], as.Text[i+1:]
	}
	if keyword == "" {
		r.add(caseName, "template:autocomplete-select", false, "缺少过滤词（断言 text 字段，格式 过滤词=期望回填）")
		return
	}
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:autocomplete-select", false, "输入框不存在")
		return
	}
	p := clickPoint(n.BoundsInRoot(), nil)
	injectDown(r.c, p.X, p.Y)
	injectUp(r.c, p.X, p.Y)
	time.Sleep(200 * time.Millisecond)
	for _, ch := range keyword {
		injectChar(r.c, ch)
		time.Sleep(20 * time.Millisecond)
	}
	time.Sleep(350 * time.Millisecond)
	base := r.overlayCount("")
	r.add(caseName, "template:autocomplete-select", base >= 1,
		fmt.Sprintf("键入「%s」后建议弹层出现（浮层 %d）", keyword, base))
	// 点弹层第一条建议（弹层锚定输入框下方，取输入框左缘 + 一行高的一半）。
	// 回填验证不能用原锚（占位文字在键入后已消失），按 label 不带文本重找。
	if base >= 1 {
		b := n.BoundsInRoot()
		item := geom.Pt(b.Min.X+30, b.Max.Y+15)
		injectDown(r.c, item.X, item.Y)
		injectUp(r.c, item.X, item.Y)
		time.Sleep(350 * time.Millisecond)
		n2, ok2 := r.locate(SpecTarget{Label: t.Label})
		got := ""
		if ok2 {
			var sb strings.Builder
			n2.Walk(func(ch *scene.Node) bool {
				if sp := ch.Text(); sp != nil {
					sb.WriteString(sp.Src)
				}
				return true
			})
			got = sb.String()
		}
		pass := ok2 && nodeTextContains(n2, expect)
		r.add(caseName, "template:autocomplete-select", pass,
			fmt.Sprintf("点建议后输入框回填「%s」（实测 %q）", expect, got))
	}
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
	r.c.Walk(func(n *scene.Node) bool {
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

// borderWidthOf 沿元素中心列自上而下扫描：与内部色差 > 24 的连续像素段
// 长度即顶部边框宽度（抗锯齿边缘含在内，调用方按 ±1 容差判定）。
// 内部参考色取"边框之后、文字之前"的窄带（中心列 y∈[Min+5,Min+8]）：
// 直接取元素中心会把居中的文字当内部色，导致整段被误判成边框（实测 7px）。
func borderWidthOf(img *image.RGBA, rect geom.Rect) int {
	if img == nil || rect.IsEmpty() {
		return 0
	}
	cx := int((rect.Min.X + rect.Max.X) / 2)
	var sum [3]int
	n := 0
	for y := int(rect.Min.Y) + 5; y <= int(rect.Min.Y)+8 && y < img.Bounds().Dy(); y++ {
		px := img.RGBAAt(cx, y)
		sum[0] += int(px.R)
		sum[1] += int(px.G)
		sum[2] += int(px.B)
		n++
	}
	if n == 0 {
		return 0
	}
	inner := [3]float64{float64(sum[0]) / float64(n), float64(sum[1]) / float64(n), float64(sum[2]) / float64(n)}
	differs := func(y int) bool {
		px := img.RGBAAt(cx, y)
		v := [3]float64{float64(px.R), float64(px.G), float64(px.B)}
		return abs3Diff(v, inner) > 24
	}
	count, started := 0, false
	for y := int(rect.Min.Y); y < int(rect.Min.Y)+8 && y < img.Bounds().Dy(); y++ {
		if differs(y) {
			count++
			started = true
		} else if started {
			break
		}
	}
	return count
}

// borderColorOf 返回区域**左边缘**（垂直边中点行）的边框色。改为左缘扫描
// 的原因：顶行扫描有三类污染——①小宽盒（tag "Go"）正中列撞文字墨迹，GDI
// ClearType 亚像素 AA 会产生彩色边缘（实测 #B9FAFA 青边）；②16px 圆钮
// （radius=半宽）整条顶边都是弧，行均色必被稀释；③内部参考固定窄带在
// 16px 小盒上会撞勾/横杠。左缘中点行：盒边是几何描边（无文字）、圆左极点
// 切线垂直（该列恰好铺满）、内部参考取扫描窗内**众数列**（16 级量化分桶
// ——固定 +2..+4 列在窄盒上会撞文字彩边，实测 #DAE9FF）。found=false 表示
// 前 4 列与参考色无明显差异（无边框）。
func borderColorOf(img *image.RGBA, rect geom.Rect) ([3]float64, bool) {
	if img == nil || rect.IsEmpty() {
		return [3]float64{}, false
	}
	cy := int(rect.Min.Y + rect.H()/2)
	if cy >= img.Bounds().Dy() {
		return [3]float64{}, false
	}
	px := func(x int) [3]float64 {
		i := cy*img.Stride + x*4
		return [3]float64{float64(img.Pix[i]), float64(img.Pix[i+1]), float64(img.Pix[i+2])}
	}
	cols := make([][3]float64, 0, 5)
	for x := int(rect.Min.X); x < int(rect.Min.X)+4 && x < img.Bounds().Dx(); x++ {
		cols = append(cols, px(x))
	}
	// 众数列：16 级量化分桶——边框 1 列 vs 背景 3+ 列，众数=内部参考色。
	buckets := map[[3]int]*struct {
		v   [3]float64
		cnt int
	}{}
	for _, v := range cols {
		var key [3]int
		for c := 0; c < 3; c++ {
			key[c] = int(v[c]) / 16
		}
		b, ok := buckets[key]
		if !ok {
			b = &struct {
				v   [3]float64
				cnt int
			}{}
			buckets[key] = b
		}
		b.v[0] += v[0]
		b.v[1] += v[1]
		b.v[2] += v[2]
		b.cnt++
	}
	var inner [3]float64
	maxCnt := 0
	for _, b := range buckets {
		if b.cnt > maxCnt {
			maxCnt = b.cnt
			inner = [3]float64{b.v[0] / float64(b.cnt), b.v[1] / float64(b.cnt), b.v[2] / float64(b.cnt)}
		}
	}
	best, bestD := [3]float64{}, -1.0
	for _, v := range cols {
		if d := abs3Diff(v, inner); d > bestD {
			bestD, best = d, v
		}
	}
	return best, bestD > 12
}

// colorRatioOf 返回区域内与目标色每通道差 <= tol 的像素占比。
func colorRatioOf(img *image.RGBA, rect geom.Rect, want geom.Color, tol float64) float64 {
	if img == nil || rect.IsEmpty() {
		return 0
	}
	w := [3]float64{float64(want.R), float64(want.G), float64(want.B)}
	x0, y0 := int(rect.Min.X), int(rect.Min.Y)
	x1, y1 := int(rect.Max.X), int(rect.Max.Y)
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
		return 0
	}
	total, hit := 0, 0
	for y := y0; y < y1; y++ {
		i := y*img.Stride + x0*4
		for x := x0; x < x1; x++ {
			v := [3]float64{float64(img.Pix[i]), float64(img.Pix[i+1]), float64(img.Pix[i+2])}
			if abs3Diff(v, w) <= tol {
				hit++
			}
			total++
			i += 4
		}
	}
	return float64(hit) / float64(total)
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

// tplButtonPress：按下期间出现 :active 色（colorPrimaryActive），
// 且占比必须显著高于 hover 基线——这是"点击有压下效果"的判据。
func (r *specRunner) tplButtonPress(caseName string, t SpecTarget) {
	n, ok := r.locate(t)
	if !ok {
		r.add(caseName, "template:button-press", false, "目标不存在")
		return
	}
	rect := n.BoundsInRoot()
	p := clickPoint(rect, nil)
	hoverAt(r.c, p.X, p.Y)
	time.Sleep(240 * time.Millisecond)
	before := grab(r.c)
	injectDown(r.c, p.X, p.Y)
	time.Sleep(150 * time.Millisecond)
	mid := grab(r.c)
	injectUp(r.c, p.X, p.Y)
	injectLeave(r.c)
	time.Sleep(200 * time.Millisecond)

	active, ok := tokenColor("PrimaryActive")
	if !ok {
		r.add(caseName, "template:button-press", false, "主题 token 缺失")
		return
	}
	rBefore := colorRatioOf(before, rect, active, 40)
	rMid := colorRatioOf(mid, rect, active, 40)
	pass := rMid-rBefore >= 0.02
	r.add(caseName, "template:button-press", pass,
		fmt.Sprintf("按下 active 色占比 %.3f（hover 基线 %.3f，需 +0.02）", rMid, rBefore))
}

// tplSpinnerSpin：点击触发 loading 后间隔抓两帧，图标区必须持续变化（在转）。
// saveCropRGBA 裁剪并放大导出 RGBA 区域（spinner 诊断用，WG_SPIN_DEBUG 门控）。
func saveCropRGBA(img *image.RGBA, r geom.Rect, path string) {
	if img == nil {
		return
	}
	x0, y0 := int(r.Min.X), int(r.Min.Y)
	w, h := int(r.Max.X-r.Min.X), int(r.Max.Y-r.Min.Y)
	if x0 < 0 || y0 < 0 || w <= 0 || h <= 0 || x0+w > img.Bounds().Dx() || y0+h > img.Bounds().Dy() {
		return
	}
	const sc = 4
	out := image.NewRGBA(image.Rect(0, 0, w*sc, h*sc))
	for y := 0; y < h*sc; y++ {
		for x := 0; x < w*sc; x++ {
			out.Set(x, y, img.RGBAAt(x0+x/sc, y0+y/sc))
		}
	}
	savePNG(out, path)
}

// 参数由断言 text 字段携带："no-still" = 持续动画目标（Spin 组件等，
// 无 loading 结束概念），跳过点击触发与静止对判定，只验证帧差非零。
func (r *specRunner) tplSpinnerSpin(caseName string, t SpecTarget, as SpecAssertion) {
	stillCheck := as.Text != "no-still"
	var n *scene.Node
	var ok bool
	var rect geom.Rect
	if stillCheck {
		n, ok = r.locate(t)
		if !ok {
			r.add(caseName, "template:spinner-spin", false, "触发按钮不存在")
			return
		}
		rect = n.BoundsInRoot()
		p := clickPoint(rect, nil)
		injectDown(r.c, p.X, p.Y)
		time.Sleep(60 * time.Millisecond)
		injectUp(r.c, p.X, p.Y)
		// 立即移出指针：hover 清除是一次性跳变，若留在"静止对"采样前才做，
		// 悬停样式恢复会整段计入静止帧差（实测 0.288，把判据淹了）。
		injectLeave(r.c)
		time.Sleep(300 * time.Millisecond)
	} else {
		n, ok = r.locate(t)
		if !ok {
			r.add(caseName, "template:spinner-spin", false, "目标不存在")
			return
		}
		rect = n.BoundsInRoot()
	}

	// 采样框稍大于图标槽（0.06..0.34 × 0.20..0.80）：弧线扫过整槽，
	// 采样框只框弧线本体会把"扫出框"的变化漏掉。
	iconRect := nodeSampleRect(n, []float32{0.06, 0.20, 0.28, 0.60})

	// 相对判据：loading 中两帧必须有差异，loading 结束后的静止两帧必须一致。
	// 弧线细（1.2px）且半透明，逐像素阈值计数会漏掉小相位差（实测 0.000），
	// 改用平均绝对差 + 500ms 观察窗 × 3 采样取最大：225° 相位差下弧线
	// 扫过整槽，meanAbsDiff 显著非零；静止帧 ≈ 0。
	f0 := r.c.PresentStats().Frames
	a := grab(r.c)
	time.Sleep(500 * time.Millisecond)
	b := grab(r.c)
	time.Sleep(500 * time.Millisecond)
	c2 := grab(r.c)
	f1 := r.c.PresentStats().Frames
	movingAB := meanAbsDiff(a, b, iconRect)
	movingBC := meanAbsDiff(b, c2, iconRect)
	movingAC := meanAbsDiff(a, c2, iconRect)
	if os.Getenv("WG_SPIN_DEBUG") != "" {
		dbg := geom.RectXYWH(rect.Min.X-4, rect.Min.Y-4, rect.W()+8, rect.H()+8)
		saveCropRGBA(a, dbg, ".shots/audit/button/audit-spin-a.png")
		saveCropRGBA(b, dbg, ".shots/audit/button/audit-spin-b.png")
		saveCropRGBA(c2, dbg, ".shots/audit/button/audit-spin-c.png")
	}
	moving := movingAB
	if movingBC > moving {
		moving = movingBC
	}
	if movingAC > moving {
		moving = movingAC
	}

	still := 0.0
	if stillCheck {
		time.Sleep(1200 * time.Millisecond) // 等 demo 的 2s loading 结束并重建完
		d2 := grab(r.c)
		time.Sleep(500 * time.Millisecond)
		e2 := grab(r.c)
		still = meanAbsDiff(d2, e2, iconRect)
	}
	pass := moving >= 0.5 && f1 > f0 && (!stillCheck || moving > still*2+0.2)
	r.add(caseName, "template:spinner-spin", pass,
		fmt.Sprintf("loading meanAbsDiff %.2f/%.2f/%.2f vs 静止 %.2f（帧 %d→%d，需 moving≥0.5 且 >2×静止+0.2）",
			movingAB, movingBC, movingAC, still, f0, f1))
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
