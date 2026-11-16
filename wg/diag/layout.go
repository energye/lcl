package diag

import (
	"fmt"
	"math"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
)

// ── --layout ────────────────────────────────────────────────────────

// LayoutCase 是单条布局断言的结果。
type LayoutCase struct {
	Name string
	OK   bool
	Got  string
	Want string
	Note string
}

// LayoutResult 是 --layout 的汇总。
type LayoutResult struct {
	Cases   int
	Failed  int
	Elapsed time.Duration
	Details []LayoutCase
}

// OK 判定是否全绿（门禁 ≥20 组）。
func (r LayoutResult) OK() bool { return r.Failed == 0 && r.Cases >= 20 }

// String 返回可读摘要。
func (r LayoutResult) String() string {
	return fmt.Sprintf("layout: %d 组，失败 %d，用时 %v", r.Cases, r.Failed, r.Elapsed)
}

// layoutCase 是一条断言：给约束、跑布局、比尺寸并可选校验子节点位置。
type layoutCase struct {
	name  string
	build func() *scene.Node
	cons  scene.Constraints
	want  geom.Size
	// skipSize 表示只校验 check（尺寸依赖字体度量时用）。
	skipSize bool
	check    func(*scene.Node) error
}

// Layout 跑约束布局矩阵（门禁：20 组，含 flex/wrap/grid/scroll/stack/text）。
func Layout() (LayoutResult, error) {
	began := time.Now()
	var res LayoutResult
	for _, c := range layoutCases() {
		node := c.build()
		got := node.Layout(c.cons)
		ok := true
		if !c.skipSize {
			ok = approx(got.W, c.want.W) && approx(got.H, c.want.H)
		}
		detail := LayoutCase{
			Name: c.name,
			OK:   ok,
			Got:  sizeStr(got),
			Want: sizeStr(c.want),
		}
		if c.skipSize {
			detail.Want = "—"
		}
		if ok && c.check != nil {
			if err := c.check(node); err != nil {
				detail.OK = false
				detail.Note = err.Error()
			}
		}
		if !detail.OK {
			res.Failed++
		}
		res.Cases++
		res.Details = append(res.Details, detail)
	}
	res.Elapsed = time.Since(began)
	return res, nil
}

func layoutCases() []layoutCase {
	// 一个 40×20 的叶子（Box + 固定尺寸），用作 flex/grid 的稳定砖块。
	leaf := func(w, h float32) *scene.Node {
		return scene.Box(scene.BoxSpec{Fill: render.Fill(geom.RGBA(200, 200, 200))}).
			SetSize2(w, h)
	}
	return []layoutCase{
		{
			name:  "tight-constraints-win",
			build: func() *scene.Node { return leaf(10, 10) },
			cons:  scene.TightConstraints(geom.Sz(120, 80)),
			want:  geom.Sz(120, 80),
		},
		{
			name:  "loose-shrink-wrap",
			build: func() *scene.Node { return leaf(30, 12) },
			cons:  scene.LooseConstraints(geom.Sz(100, 100)),
			want:  geom.Sz(30, 12),
		},
		{
			name: "unbounded-uses-content",
			build: func() *scene.Node {
				return scene.Flex(scene.Vertical, 4, leaf(50, 10), leaf(30, 20))
			},
			cons: scene.Unbounded(),
			want: geom.Sz(50, 34), // 10 + 4 + 20
		},
		{
			name:  "min-max-clamp",
			build: func() *scene.Node { return leaf(10, 10).SetMinSize(40, 30) },
			cons:  scene.LooseConstraints(geom.Sz(1000, 1000)),
			want:  geom.Sz(40, 30),
		},
		{
			name:  "fixed-width-only",
			build: func() *scene.Node { return leaf(nan32(), nan32()).SetWidth(72) },
			cons:  scene.LooseConstraints(geom.Sz(500, 500)),
			want:  geom.Sz(72, 0),
		},
		{
			name: "aspect-ratio-from-width",
			build: func() *scene.Node {
				return scene.Box(scene.BoxSpec{}).SetWidth(160).SetAspectRatio(2)
			},
			cons: scene.LooseConstraints(geom.Sz(500, 500)),
			want: geom.Sz(160, 80),
		},
		{
			name: "padding-adds-to-content",
			build: func() *scene.Node {
				return scene.Box(scene.BoxSpec{}).SetPadding(geom.InLTRB(10, 6, 10, 6)).Add(leaf(40, 20))
			},
			cons: scene.LooseConstraints(geom.Sz(500, 500)),
			want: geom.Sz(60, 32), // 40+20, 20+12
		},
		{
			name: "margin-adds-to-outer",
			build: func() *scene.Node {
				child := leaf(40, 20).SetMargin(geom.In(8))
				return scene.Box(scene.BoxSpec{}).Add(child)
			},
			cons: scene.LooseConstraints(geom.Sz(500, 500)),
			want: geom.Sz(56, 36), // 40+16, 20+16
			check: func(n *scene.Node) error {
				return expectOffset("子节点应偏移 margin", n.Children()[0], geom.Pt(8, 8))
			},
		},
		{
			name: "flex-row-equal-weights",
			build: func() *scene.Node {
				a := scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight)
				b := scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight)
				row := scene.Flex(scene.Horizontal, 0, a, b)
				row.SetCrossAlign(scene.CrossStretch).SetSize2(200, 40)
				return row
			},
			cons: scene.TightConstraints(geom.Sz(200, 40)),
			want: geom.Sz(200, 40),
			check: func(n *scene.Node) error {
				if err := expectSize("a", n.Children()[0], geom.Sz(100, 40)); err != nil {
					return err
				}
				if err := expectOffset("b", n.Children()[1], geom.Pt(100, 0)); err != nil {
					return err
				}
				return nil
			},
		},
		{
			name: "flex-column-with-gap",
			build: func() *scene.Node {
				return scene.Flex(scene.Vertical, 8, leaf(20, 10), leaf(20, 10), leaf(20, 10))
			},
			cons: scene.LooseConstraints(geom.Sz(100, 200)),
			want: geom.Sz(20, 46), // 30 + 16
			check: func(n *scene.Node) error {
				return expectOffset("第二个应在 y=18", n.Children()[1], geom.Pt(0, 18))
			},
		},
		{
			name: "flex-fixed-plus-flexible",
			build: func() *scene.Node {
				fixed := leaf(60, 20)
				grow := scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight)
				row := scene.Flex(scene.Horizontal, 10, fixed, grow)
				row.SetCrossAlign(scene.CrossStretch).SetSize2(200, 20)
				return row
			},
			cons: scene.TightConstraints(geom.Sz(200, 20)),
			want: geom.Sz(200, 20),
			check: func(n *scene.Node) error {
				// 200 - 60 - 10 = 130
				if err := expectSize("grow", n.Children()[1], geom.Sz(130, 20)); err != nil {
					return err
				}
				return expectOffset("grow 起点", n.Children()[1], geom.Pt(70, 0))
			},
		},
		{
			name: "flex-space-between",
			build: func() *scene.Node {
				row := scene.Flex(scene.Horizontal, 0, leaf(20, 10), leaf(20, 10), leaf(20, 10))
				row.SetMainAlign(scene.MainSpaceBetween).SetSize2(300, 10)
				return row
			},
			cons: scene.TightConstraints(geom.Sz(300, 10)),
			want: geom.Sz(300, 10),
			check: func(n *scene.Node) error {
				// 剩余 240 均分成 2 个间隙
				if err := expectOffset("中间项", n.Children()[1], geom.Pt(140, 0)); err != nil {
					return err
				}
				return expectOffset("末项", n.Children()[2], geom.Pt(280, 0))
			},
		},
		{
			name: "flex-cross-stretch",
			build: func() *scene.Node {
				// 只有固定宽、没有固定高 → 可被拉伸（对应 CSS align-items:stretch）
				item := scene.Box(scene.BoxSpec{}).SetWidth(30)
				row := scene.Flex(scene.Horizontal, 0, item)
				row.SetCrossAlign(scene.CrossStretch).SetSize2(100, 44)
				return row
			},
			cons: scene.TightConstraints(geom.Sz(100, 44)),
			want: geom.Sz(100, 44),
			check: func(n *scene.Node) error {
				// 有显式高度的子节点不该被拉伸，这条一起断言掉
				if err := expectSize("子节点被拉伸", n.Children()[0], geom.Sz(30, 44)); err != nil {
					return err
				}
				other := scene.Box(scene.BoxSpec{}).SetSize2(30, 10)
				row2 := scene.Flex(scene.Horizontal, 0, other)
				row2.SetCrossAlign(scene.CrossStretch).SetSize2(100, 44)
				row2.Layout(scene.TightConstraints(geom.Sz(100, 44)))
				return expectSize("显式高度不被拉伸", other, geom.Sz(30, 10))
			},
		},
		{
			name: "wrap-three-runs",
			build: func() *scene.Node {
				return scene.Wrap(scene.Horizontal, 4, 4,
					leaf(50, 10), leaf(50, 10), leaf(50, 10), leaf(50, 10))
			},
			cons: scene.LooseConstraints(geom.Sz(110, 500)),
			want: geom.Sz(104, 24), // 每行 2 个（50+4+50=104），2 行（10+4+10）
			check: func(n *scene.Node) error {
				return expectOffset("第三个应换行", n.Children()[2], geom.Pt(0, 14))
			},
		},
		{
			name: "stack-center-align",
			build: func() *scene.Node {
				s := scene.Stack(leaf(40, 20))
				s.SetSize2(200, 100).SetAlign(scene.Center)
				return s
			},
			cons: scene.TightConstraints(geom.Sz(200, 100)),
			want: geom.Sz(200, 100),
			check: func(n *scene.Node) error {
				return expectOffset("居中", n.Children()[0], geom.Pt(80, 40))
			},
		},
		{
			name: "stack-positioned-child",
			build: func() *scene.Node {
				s := scene.Stack(leaf(10, 10).SetStackPosition(geom.Pt(30, 15)))
				s.SetSize2(100, 80)
				return s
			},
			cons: scene.TightConstraints(geom.Sz(100, 80)),
			want: geom.Sz(100, 80),
			check: func(n *scene.Node) error {
				return expectOffset("绝对定位", n.Children()[0], geom.Pt(30, 15))
			},
		},
		{
			name: "grid-three-columns",
			build: func() *scene.Node {
				// 4 个子节点 / 3 列 → 2 行；宽度 (300-2*10)/3
				g := scene.Grid(3, 10,
					leaf(0, 20), leaf(0, 20), leaf(0, 20), leaf(0, 20))
				g.SetSize2(300, 50)
				return g
			},
			cons: scene.TightConstraints(geom.Sz(300, 50)),
			want: geom.Sz(300, 50),
			check: func(n *scene.Node) error {
				// 单元格宽 (300-20)/3 = 93.33
				if err := expectSize("单元格", n.Children()[0], geom.Sz(93.3333, 20)); err != nil {
					return err
				}
				// 第二个单元格 x = 93.33 + 10
				return expectOffset("第二列", n.Children()[1], geom.Pt(103.3333, 0))
			},
		},
		{
			name: "grid-span-two",
			build: func() *scene.Node {
				wide := leaf(0, 20).SetSpan(2)
				g := scene.Grid(3, 10, wide, leaf(0, 20), leaf(0, 20))
				g.SetSize2(300, 44)
				return g
			},
			cons: scene.TightConstraints(geom.Sz(300, 44)),
			want: geom.Sz(300, 44),
			check: func(n *scene.Node) error {
				// 跨 2 列 = 93.33*2 + 10 = 196.67
				return expectSize("跨列", n.Children()[0], geom.Sz(196.6667, 20))
			},
		},
		{
			name: "scroll-content-exceeds-viewport",
			build: func() *scene.Node {
				content := scene.Flex(scene.Vertical, 0, leaf(10, 100), leaf(10, 100))
				s := scene.Scroll(scene.Vertical, content)
				s.SetSize2(60, 50)
				return s
			},
			cons: scene.TightConstraints(geom.Sz(60, 50)),
			want: geom.Sz(60, 50),
			check: func(n *scene.Node) error {
				if got := n.ScrollContentSize(); !approx(got.H, 200) {
					return fmt.Errorf("内容高应为 200，实际 %v", got.H)
				}
				// 越界偏移必须被 clamp
				n.SetScrollOffset(1000)
				if got := n.ScrollOffset(); !approx(got, 150) {
					return fmt.Errorf("偏移应 clamp 到 150，实际 %v", got)
				}
				return nil
			},
		},
		{
			name: "text-wraps-at-max-width",
			build: func() *scene.Node {
				l := text.LayoutText("WG Engine 布局", text.DefaultFont(13), text.Options{})
				return scene.Text(scene.TextSpec{
					Src: "WG Engine 布局", Layout: l, LayoutFont: text.DefaultFont(13),
					Color: geom.RGBA(0, 0, 0), MaxWidth: 40,
				})
			},
			cons:     scene.LooseConstraints(geom.Sz(40, 200)),
			skipSize: true, // 高度由字体度量决定，只校验"确实换行了"
			check: func(n *scene.Node) error {
				if n.Size().H <= 0 {
					return fmt.Errorf("换行后高度应 > 0，实际 %v", n.Size().H)
				}
				if n.Text().Layout == nil || len(n.Text().Layout.Lines) < 2 {
					return fmt.Errorf("40dip 宽应换行成多行")
				}
				return nil
			},
		},
	}
}

// approx 是布局比较的容差：0.01dip。浮点累加误差远小于它，
// 而任何真实的布局错误都远大于它。
func approx(a, b float32) bool {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d <= 0.01
}

func expectSize(what string, n *scene.Node, want geom.Size) error {
	if !approx(n.Size().W, want.W) || !approx(n.Size().H, want.H) {
		return fmt.Errorf("%s 尺寸应为 %v，实际 %v", what, want, n.Size())
	}
	return nil
}

func expectOffset(what string, n *scene.Node, want geom.Point) error {
	if !approx(n.Offset().X, want.X) || !approx(n.Offset().Y, want.Y) {
		return fmt.Errorf("%s 位置应为 %v，实际 %v", what, want, n.Offset())
	}
	return nil
}

func sizeStr(s geom.Size) string { return fmt.Sprintf("%.2f×%.2f", s.W, s.H) }

// nan32 返回"未设置"的尺寸哨兵。scene 用 NaN 表示"由约束与内容决定"，
// 用例里写 `.SetWidth(nan32())` 就是"清掉固定宽度、回到自动"。
func nan32() float32 { return float32(math.NaN()) }
