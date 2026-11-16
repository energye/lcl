package diag

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
)

// ── --hit ───────────────────────────────────────────────────────────

// HitCase 是单条命中断言的结果。
type HitCase struct {
	Name string
	OK   bool
	Got  string
	Note string
}

// HitResult 是 --hit 的汇总。
type HitResult struct {
	Cases   int
	Failed  int
	Elapsed time.Duration
	Details []HitCase
}

// OK 判定是否全绿（门禁 ≥12 组）。
func (r HitResult) OK() bool { return r.Failed == 0 && r.Cases >= 12 }

// String 返回可读摘要。
func (r HitResult) String() string {
	return fmt.Sprintf("hit: %d 组，失败 %d，用时 %v", r.Cases, r.Failed, r.Elapsed)
}

// hitCase 是一条断言：给定根节点与点，比较命中目标标签。
type hitCase struct {
	name string
	at   geom.Point
	want string // 期望命中的 Label（"" 表示期望不命中）
	run  func(at geom.Point) input.Path
	// check 可选：额外校验（顺序、局部坐标等）
	check func(input.Path) error
}

// Hit 跑命中测试门禁：覆盖变换、裁剪、重叠、行为、禁用/不可见、滚动。
func Hit() (HitResult, error) {
	began := time.Now()
	var res HitResult
	for _, c := range hitCases() {
		p := c.run(c.at)
		got := p.Label()
		d := HitCase{Name: c.name, OK: got == c.want, Got: displayLabel(got)}
		if !d.OK {
			d.Note = "期望 " + displayLabel(c.want)
		}
		if d.OK && c.check != nil {
			if err := c.check(p); err != nil {
				d.OK = false
				d.Note = err.Error()
			}
		}
		if !d.OK {
			res.Failed++
		}
		res.Cases++
		res.Details = append(res.Details, d)
	}
	res.Elapsed = time.Since(began)
	return res, nil
}

// displayLabel 让空标签在输出里可读。
func displayLabel(s string) string {
	if s == "" {
		return "(miss)"
	}
	return s
}

func hitCases() []hitCase {
	// hitBox 造一个可命中矩形。
	hitBox := func(w, h float32, label string) *scene.Node {
		return scene.Box(scene.BoxSpec{Fill: render.Fill(geom.RGBA(200, 200, 200))}).
			SetSize2(w, h).
			SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: label})
	}
	lay := func(n *scene.Node) *scene.Node {
		n.Layout(scene.Unbounded())
		return n
	}

	return []hitCase{
		{
			name: "rect-inside",
			at:   geom.Pt(10, 10),
			want: "box",
			run: func(at geom.Point) input.Path {
				return input.HitTest(lay(hitBox(40, 40, "box")), at)
			},
		},
		{
			name: "rect-outside",
			at:   geom.Pt(90, 10),
			want: "",
			run: func(at geom.Point) input.Path {
				return input.HitTest(lay(hitBox(40, 40, "box")), at)
			},
		},
		{
			name: "z-order-top-wins",
			at:   geom.Pt(15, 15),
			want: "top",
			run: func(at geom.Point) input.Path {
				root := scene.Stack(
					hitBox(40, 40, "bottom"),
					hitBox(40, 40, "top").SetStackPosition(geom.Pt(10, 10)),
				).SetSize2(80, 80)
				return input.HitTest(lay(root), at)
			},
		},
		{
			name: "defer-to-child-container",
			at:   geom.Pt(35, 35),
			want: "",
			run: func(at geom.Point) input.Path {
				// 容器 DeferToChild：自身不命中，空白处应当穿透
				root := scene.Stack(hitBox(20, 20, "child")).SetSize2(60, 60)
				root.SetHit(scene.HitSpec{Behavior: scene.HitDeferToChild, Label: "container"})
				return input.HitTest(lay(root), at)
			},
		},
		{
			name: "container-hit-inside-child-area",
			at:   geom.Pt(5, 5),
			want: "child",
			run: func(at geom.Point) input.Path {
				root := scene.Stack(hitBox(20, 20, "child")).SetSize2(60, 60)
				root.SetHit(scene.HitSpec{Behavior: scene.HitDeferToChild})
				return input.HitTest(lay(root), at)
			},
		},
		{
			name: "translucent-parent-fills-gap",
			at:   geom.Pt(35, 35),
			want: "parent",
			run: func(at geom.Point) input.Path {
				root := scene.Stack(hitBox(20, 20, "child")).SetSize2(60, 60)
				root.SetHit(scene.HitSpec{Behavior: scene.HitTranslucent, Label: "parent"})
				return input.HitTest(lay(root), at)
			},
		},
		{
			name: "rotated-45deg-corner-miss",
			at:   geom.Pt(1, 1),
			want: "",
			run: func(at geom.Point) input.Path {
				// 20×20 绕中心 (10,10) 转 45° 成菱形：沿对角线方向的半宽仍是 10，
				// 而 (1,1) 到中心距离 12.7 > 10，已在菱形外——虽然它落在原包围盒里。
				n := hitBox(20, 20, "box").SetTransform(geom.Rotate(3.14159265 / 4))
				return input.HitTest(lay(n), at)
			},
		},
		{
			name: "rotated-45deg-center",
			at:   geom.Pt(10, 10),
			want: "box",
			run: func(at geom.Point) input.Path {
				n := hitBox(20, 20, "box").SetTransform(geom.Rotate(3.14159265 / 4))
				return input.HitTest(lay(n), at)
			},
		},
		{
			name: "scaled-parent-child-at-edge",
			at:   geom.Pt(38, 38),
			want: "child",
			run: func(at geom.Point) input.Path {
				// 父节点绕中心缩放 2 倍（支点 = 中心，见 scene.FrameMatrix）：
				//   (38,38) → 根局部 (24,24) → 子局部 (14,14)
				// 若关掉缩放，同样的点会落在节点外——这条断言同时守住了
				// "缩放参与命中"和"支点是中心"两件事。
				child := hitBox(10, 10, "child").SetStackPosition(geom.Pt(10, 10))
				root := scene.Stack(child).SetSize2(20, 20)
				root.SetTransform(geom.ScaleUniform(2))
				return input.HitTest(lay(root), at)
			},
			check: func(p input.Path) error {
				if p.Len() != 2 {
					return fmt.Errorf("路径长度 %d，应为 2（根 + 子）", p.Len())
				}
				if got := p.Local(); !approxPt(got, geom.Pt(14, 14)) {
					return fmt.Errorf("目标局部坐标 %v，应为 (14,14)", got)
				}
				return nil
			},
		},
		{
			name: "rrect-clip-corner-miss",
			at:   geom.Pt(2, 2),
			want: "",
			run: func(at geom.Point) input.Path {
				// 圆角 16 的 40×40：左上角 (2,2) 在圆外
				root := scene.Clip(scene.ClipSpec{Radius: 16}, hitBox(40, 40, "inner"))
				root.SetHit(scene.HitSpec{Behavior: scene.HitDeferToChild})
				return input.HitTest(lay(root), at)
			},
		},
		{
			name: "rrect-clip-center-hit",
			at:   geom.Pt(20, 20),
			want: "inner",
			run: func(at geom.Point) input.Path {
				root := scene.Clip(scene.ClipSpec{Radius: 16}, hitBox(40, 40, "inner"))
				return input.HitTest(lay(root), at)
			},
		},
		{
			name: "rect-clip-excludes-child",
			at:   geom.Pt(30, 30),
			want: "",
			run: func(at geom.Point) input.Path {
				root := scene.Clip(scene.ClipSpec{Rect: geom.RectXYWH(0, 0, 20, 20)}, hitBox(40, 40, "inner"))
				return input.HitTest(lay(root), at)
			},
		},
		{
			name: "scroll-offset-shifts-hit",
			at:   geom.Pt(5, 5),
			want: "",
			run: func(at geom.Point) input.Path {
				// 内容高 100，可命中块只占顶部 20：滚 20 之后 (5,5) 已是空白
				content := scene.Stack(hitBox(40, 20, "item")).SetSize2(40, 100)
				sc := scene.Scroll(scene.Vertical, content)
				sc.SetSize2(40, 40)
				sc.Layout(scene.Unbounded())
				sc.SetScrollOffset(20)
				return input.HitTest(sc, at)
			},
		},
		{
			name: "scroll-offset-hit-lower",
			at:   geom.Pt(5, 5),
			want: "item",
			run: func(at geom.Point) input.Path {
				content := scene.Stack(hitBox(40, 20, "item")).SetSize2(40, 100)
				sc := scene.Scroll(scene.Vertical, content)
				sc.SetSize2(40, 40)
				sc.Layout(scene.Unbounded())
				sc.SetScrollOffset(0)
				return input.HitTest(sc, at)
			},
		},
		{
			name: "circle-shape-corner-miss",
			at:   geom.Pt(2, 2),
			want: "",
			run: func(at geom.Point) input.Path {
				n := hitBox(40, 40, "dot").SetHit(scene.HitSpec{
					Behavior: scene.HitOpaque, Shape: scene.ShapeCircle, Label: "dot",
				})
				return input.HitTest(lay(n), at)
			},
		},
		{
			name: "invisible-and-disabled-miss",
			at:   geom.Pt(10, 10),
			want: "",
			run: func(at geom.Point) input.Path {
				a := hitBox(40, 40, "hidden").SetVisible(false)
				b := hitBox(40, 40, "disabled").SetEnabled(false)
				root := scene.Stack(a, b).SetSize2(40, 40)
				return input.HitTest(lay(root), at)
			},
		},
	}
}

// approxPt 判定两点是否近似相等（容差 0.01）。
func approxPt(a, b geom.Point) bool {
	const eps = 0.01
	d := a.Sub(b)
	return d.X > -eps && d.X < eps && d.Y > -eps && d.Y < eps
}
