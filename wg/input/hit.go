package input

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
)

// Path 是命中路径（根 → 目标），并携带每一层对应的**局部坐标**。
//
// 为什么返回整条路径而不是一个节点：分发要在它上面做捕获（根→叶）与
// 冒泡（叶→根），hover 差分要知道"从哪个共同祖先开始分叉"，
// 光标要沿路径向上找第一个非默认值。只返回目标节点的话，这些都做不了。
type Path struct {
	nodes  []*scene.Node
	locals []geom.Point
}

// Empty 表示没有命中。
func (p Path) Empty() bool { return len(p.nodes) == 0 }

// Len 返回路径长度。
func (p Path) Len() int { return len(p.nodes) }

// Nodes 返回路径节点（根在前）。
func (p Path) Nodes() []*scene.Node { return p.nodes }

// Target 返回最深节点（真正被点中的那个）。
func (p Path) Target() *scene.Node {
	if len(p.nodes) == 0 {
		return nil
	}
	return p.nodes[len(p.nodes)-1]
}

// Local 返回目标节点局部坐标下的点。
func (p Path) Local() geom.Point {
	if len(p.locals) == 0 {
		return geom.Point{}
	}
	return p.locals[len(p.locals)-1]
}

// At 返回第 i 层节点的局部坐标。
func (p Path) At(i int) (n *scene.Node, local geom.Point) {
	if i < 0 || i >= len(p.nodes) {
		return nil, geom.Point{}
	}
	return p.nodes[i], p.locals[i]
}

// Contains 判定节点是否在路径上。
func (p Path) Contains(n *scene.Node) bool {
	for _, v := range p.nodes {
		if v == n {
			return true
		}
	}
	return false
}

// IndexOf 返回节点在路径中的下标，-1 表示不在。
func (p Path) IndexOf(n *scene.Node) int {
	for i, v := range p.nodes {
		if v == n {
			return i
		}
	}
	return -1
}

// Absorb 表示路径上是否有节点要求吞掉事件（模态遮罩）。
func (p Path) Absorb() bool {
	for _, n := range p.nodes {
		if n.HitSpec().Absorb {
			return true
		}
	}
	return false
}

// Cursor 返回路径上最深的显式光标（默认返回 CursorDefault）。
func (p Path) Cursor() scene.Cursor {
	for i := len(p.nodes) - 1; i >= 0; i-- {
		if c := p.nodes[i].HitSpec().Cursor; c != scene.CursorDefault {
			return c
		}
	}
	return scene.CursorDefault
}

// Label 返回路径上最深节点的语义标签（自动化测试选择器用）。
func (p Path) Label() string {
	for i := len(p.nodes) - 1; i >= 0; i-- {
		if l := p.nodes[i].HitSpec().Label; l != "" {
			return l
		}
	}
	return ""
}

// label 是 Label 的包内别名（诊断里更短）。
func (p Path) label() string { return p.Label() }

// HitTest 在场景树上做命中测试，p 是**根空间**（DIP）的坐标。
//
// 算法（04-UI-INTERACTION.md §B3.1）：
//
//	逐层求逆映射把点带到节点局部空间 → 测本层裁剪 → 逆绘制序找子节点
//	→ 都没有才测自身形状。
//
// 三个必须这样做的理由：
//   - 逆映射而不是正向变换包围盒：旋转/缩放过的节点用包围盒判会明显跑偏；
//   - 逆绘制序：后画的在上，第一个命中的就是视觉最上层；
//   - 形状命中复用 scene.LocalHit（圆角/椭圆/路径），与绘制同一份几何，
//     否则会出现"点方角也响应"。
func HitTest(root *scene.Node, p geom.Point) Path {
	if root == nil {
		return Path{}
	}
	var nodes []*scene.Node
	var locals []geom.Point
	if hitNode(root, p, &nodes, &locals) {
		return Path{nodes: nodes, locals: locals}
	}
	return Path{}
}

// hitNode 在 n 的**父内容空间**里接收点 p。
func hitNode(n *scene.Node, p geom.Point, nodes *[]*scene.Node, locals *[]geom.Point) bool {
	if n == nil || !n.Visible() || !n.Enabled() {
		return false
	}
	inv, ok := n.FrameMatrix(geom.Identity).Invert()
	if !ok {
		return false // 退化矩阵（缩放到 0）：不参与命中，也不可能看清
	}
	q := inv.MapPoint(p)

	// 本层裁剪：ClipSpec 是节点局部坐标，逐层测试等价于沿路径求交，
	// 且天然支持圆角/路径裁剪（累加一个矩形反而只能近似）。
	if spec := n.Clip(); spec != nil {
		r := spec.Rect
		if r.IsEmpty() {
			r = geom.RectXYWH(0, 0, n.Size().W, n.Size().H)
		}
		switch {
		case spec.Path != nil:
			if !spec.Path.Bounds().Contains(q) {
				return false
			}
		case spec.Radius > 0:
			if !scene.PointInRRect(q.Sub(r.Min), r.Size(), spec.Radius) {
				return false
			}
		default:
			if !r.Contains(q) {
				return false
			}
		}
	}

	ch := n.Children()
	for i := len(ch) - 1; i >= 0; i-- {
		if hitNode(ch[i], q, nodes, locals) {
			prepend(nodes, locals, n, q)
			return true
		}
	}

	spec := n.HitSpec()
	if spec.Behavior == scene.HitDeferToChild {
		return false
	}
	if !n.LocalHit(q) {
		return false
	}
	prepend(nodes, locals, n, q)
	return true
}

// prepend 把 (n, q) 插到路径头部（路径是根→叶，回溯时自然后插先得）。
func prepend(nodes *[]*scene.Node, locals *[]geom.Point, n *scene.Node, q geom.Point) {
	*nodes = append(*nodes, nil)
	copy((*nodes)[1:], *nodes)
	(*nodes)[0] = n
	*locals = append(*locals, geom.Point{})
	copy((*locals)[1:], *locals)
	(*locals)[0] = q
}

// mapDown 把根空间点沿既定路径逐层映射回各层局部坐标。
//
// 拖动时需要它：路径是按下那一刻快照的（捕获语义要求"不重新命中"），
// 但节点可能已经移动，所以局部坐标必须按**当前**矩阵重算。
func mapDown(root geom.Point, path []*scene.Node) []geom.Point {
	out := make([]geom.Point, 0, len(path))
	p := root
	for _, n := range path {
		inv, ok := n.FrameMatrix(geom.Identity).Invert()
		if !ok {
			break
		}
		p = inv.MapPoint(p)
		out = append(out, p)
	}
	return out
}
