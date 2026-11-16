// Package host 是唯一接触 LCL 的宿主层：把 IEngForm + ICustomGraphicControl
// 绑定成一个无句柄的画布宿主，并维护 Surface 树。
//
// 本包是 wg 中仅有的三个允许 import lcl 的地方之一（另两个是 raster/gdi 与
// raster/software）。上层任何包出现 lcl 依赖都视为 C17 违规。
//
// 线程模型：Surface 树不是线程安全的。所有结构性变更（Append/Remove/Resize）
// 必须发生在帧边界之前，由 wg.Container 串行化（契约 C18：帧内禁止资源拓扑变更）。
// 跨线程只允许调用 wg.Container.RequestFrame 投递"下一帧再来"。
package host

import "github.com/energye/lcl/wg/geom"

// SurfaceKind 是画布种类。全部无句柄，层级由容器统一管理。
type SurfaceKind uint8

const (
	// SurfaceInline 与父共享像素，不额外分配（面板、列表、滚动视口）。
	SurfaceInline SurfaceKind = iota
	// SurfaceLayer 持有离屏缓冲，用于组透明度/模糊/蒙版/缓存（升级有代价）。
	SurfaceLayer
	// SurfacePopup 是顶层 z 的画布：下拉、菜单、Tooltip、Modal（不新增 HWND）。
	SurfacePopup
	// SurfaceOffscreen 是完全离屏的画布：导出、预合成、缓存烘焙。
	SurfaceOffscreen
)

// String 返回可读名。
func (k SurfaceKind) String() string {
	switch k {
	case SurfaceInline:
		return "inline"
	case SurfaceLayer:
		return "layer"
	case SurfacePopup:
		return "popup"
	case SurfaceOffscreen:
		return "offscreen"
	default:
		return "unknown"
	}
}

// Offscreen 表示该种类是否需要独立像素缓冲。
func (k SurfaceKind) Offscreen() bool {
	return k == SurfaceLayer || k == SurfaceOffscreen
}

// Surface 是画布树的一个节点。
//
// 坐标为相对父的 DIP；裁剪默认开启（子内容裁到 bounds 内）。
// z 序即 children 顺序（后追加者在上），不做单独的 z 值——这样"插入顺序 =
// 绘制顺序 = 命中顺序"三者一致，避免示例里 z 与命中不一致的老问题。
type Surface struct {
	kind     SurfaceKind
	parent   *Surface
	children []*Surface

	bounds geom.Rect   // 相对父
	matrix geom.Matrix // 相对父（默认 Identity）
	clip   bool

	damage geom.Rect // 本帧待重绘区域（相对本 Surface）
	// surfaceDirty 表示子树结构或几何发生变化，需要重新走布局/重建。
	surfaceDirty bool
}

// NewSurface 创建一个画布节点；bounds 为相对父的矩形（DIP）。
func NewSurface(kind SurfaceKind, bounds geom.Rect) *Surface {
	return &Surface{
		kind:   kind,
		bounds: bounds,
		matrix: geom.Identity,
		clip:   true,
		damage: geom.ZeroRect,
	}
}

// Kind 返回画布种类。
func (s *Surface) Kind() SurfaceKind { return s.kind }

// Parent 返回父画布；根画布返回 nil。
func (s *Surface) Parent() *Surface { return s.parent }

// Children 返回子画布（只读视图，勿修改）。
func (s *Surface) Children() []*Surface { return s.children }

// Bounds 返回相对父的矩形。
func (s *Surface) Bounds() geom.Rect { return s.bounds }

// SetBounds 修改相对父的矩形并标记结构性脏。
func (s *Surface) SetBounds(r geom.Rect) {
	if s.bounds == r {
		return
	}
	s.bounds = r
	s.MarkStructureDirty()
}

// Matrix 返回相对父的变换矩阵。
func (s *Surface) Matrix() geom.Matrix { return s.matrix }

// SetMatrix 设置相对父的变换矩阵并标记结构性脏。
func (s *Surface) SetMatrix(m geom.Matrix) {
	if s.matrix == m {
		return
	}
	s.matrix = m
	s.MarkStructureDirty()
}

// SetTranslate 只改平移的快捷方式（最常见的变换）。
func (s *Surface) SetTranslate(dx, dy float32) {
	s.SetMatrix(geom.Translate(dx, dy))
}

// Clip 返回是否裁剪到 bounds。
func (s *Surface) Clip() bool { return s.clip }

// SetClip 设置裁剪开关。
func (s *Surface) SetClip(v bool) {
	if s.clip == v {
		return
	}
	s.clip = v
	s.MarkStructureDirty()
}

// Append 把 child 追加为最后一个子（即最上层）。
//
// 拒绝：nil、self、会形成环的祖先、已经挂在别处的画布。
// 契约 C18 的前置保护：成环会让帧遍历无限递归，必须在入口挡住。
func (s *Surface) Append(child *Surface) error {
	if child == nil {
		return ErrNilSurface
	}
	if child == s {
		return ErrSurfaceCycle
	}
	if child.parent != nil {
		return ErrSurfaceAttached
	}
	for p := s; p != nil; p = p.parent {
		if p == child {
			return ErrSurfaceCycle
		}
	}
	child.parent = s
	s.children = append(s.children, child)
	child.markAllDirty()
	s.MarkStructureDirty()
	return nil
}

// InsertAt 把 child 插到 index 位置（z 序调整），index 越界时追加到末尾。
func (s *Surface) InsertAt(index int, child *Surface) error {
	if err := s.Append(child); err != nil {
		return err
	}
	n := len(s.children)
	if index < 0 || index >= n-1 {
		return nil // Append 已放到末尾，等价
	}
	// 摘出末尾插到 index
	c := s.children[n-1]
	copy(s.children[index+1:], s.children[index:n-1])
	s.children[index] = c
	s.MarkStructureDirty()
	return nil
}

// Remove 把 child 从子列表移除。
func (s *Surface) Remove(child *Surface) bool {
	for i, c := range s.children {
		if c == child {
			copy(s.children[i:], s.children[i+1:])
			s.children[len(s.children)-1] = nil
			s.children = s.children[:len(s.children)-1]
			child.parent = nil
			s.MarkStructureDirty()
			return true
		}
	}
	return false
}

// Detach 把自己从父上摘下。
func (s *Surface) Detach() bool {
	if s.parent == nil {
		return false
	}
	return s.parent.Remove(s)
}

// Index 返回自己在父的 children 中的位置；根返回 -1。
func (s *Surface) Index() int {
	if s.parent == nil {
		return -1
	}
	for i, c := range s.parent.children {
		if c == s {
			return i
		}
	}
	return -1
}

// Depth 返回层级深度（根为 0）。
func (s *Surface) Depth() int {
	d := 0
	for p := s.parent; p != nil; p = p.parent {
		d++
	}
	return d
}

// IsAncestorOf 判定 s 是否为 o 的祖先（含自身）。
func (s *Surface) IsAncestorOf(o *Surface) bool {
	for p := o; p != nil; p = p.parent {
		if p == s {
			return true
		}
	}
	return false
}

// LocalToParent 返回本画布**局部坐标**到父坐标的变换。
//
// 局部坐标的原点是本画布的左上角（bounds.Min）。因此这个变换是
// "先应用 matrix，再平移到 bounds.Min"——顺序不能反：
// matrix 表达的是画布自身内容相对自身左上角的变形。
func (s *Surface) LocalToParent() geom.Matrix {
	return geom.Translate(s.bounds.Min.X, s.bounds.Min.Y).Concat(s.matrix)
}

// LocalToRoot 返回局部坐标到根坐标的变换。
func (s *Surface) LocalToRoot() geom.Matrix {
	if s.parent == nil {
		return s.LocalToParent()
	}
	return s.parent.LocalToRoot().Concat(s.LocalToParent())
}

// AbsoluteMatrix 是 LocalToRoot 的旧名（保留以对齐文档术语）。
//
// Deprecated: 用 LocalToRoot。
func (s *Surface) AbsoluteMatrix() geom.Matrix { return s.LocalToRoot() }

// AbsoluteOrigin 返回本画布局部原点在根坐标系中的位置。
func (s *Surface) AbsoluteOrigin() geom.Point {
	return s.LocalToRoot().MapPoint(geom.Point{})
}

// BoundsInRoot 返回本画布在根坐标系中的外包矩形。
func (s *Surface) BoundsInRoot() geom.Rect {
	return s.LocalToRoot().MapRect(s.LocalBounds())
}

// Walk 前序遍历子树；fn 返回 false 时终止遍历。返回遍历到的节点数。
func (s *Surface) Walk(fn func(*Surface) bool) int {
	if s == nil {
		return 0
	}
	n := 0
	var walk func(*Surface) bool
	walk = func(cur *Surface) bool {
		n++
		if !fn(cur) {
			return false
		}
		for _, c := range cur.children {
			if !walk(c) {
				return false
			}
		}
		return true
	}
	walk(s)
	return n
}

// Count 返回子树节点总数（含自身）。
func (s *Surface) Count() int {
	return s.Walk(func(*Surface) bool { return true })
}

// Invalidate 标记本画布的一块区域需要重绘，并把脏区上卷到整棵祖先链（契约 C20）。
//
// r 使用**本画布坐标**：原点在本画布的左上角（即 bounds.Min），不是父坐标。
// 开启裁剪时先裁到 (0,0,W,H)；上卷到父时先按 matrix 变换，再平移 bounds.Min。
func (s *Surface) Invalidate(r geom.Rect) {
	if r.IsEmpty() {
		return
	}
	if s.clip {
		r = r.Intersect(s.LocalBounds())
		if r.IsEmpty() {
			return
		}
	}
	s.damage = s.damage.Union(r)
	if s.parent == nil {
		return
	}
	// 上卷：局部 → 父，再交给父继续上卷（父会做自己的裁剪）。
	s.parent.Invalidate(s.LocalToParent().MapRect(r))
}

// LocalBounds 返回以自身左上角为原点的本地矩形 (0,0,W,H)。
func (s *Surface) LocalBounds() geom.Rect {
	return geom.RectXYWH(0, 0, s.bounds.W(), s.bounds.H())
}

// Damage 返回本帧脏区（相对本画布）。
func (s *Surface) Damage() geom.Rect { return s.damage }

// HasDamage 表示本帧有脏区。
func (s *Surface) HasDamage() bool { return !s.damage.IsEmpty() }

// ClearDamage 清空本画布脏区（不递归）。
func (s *Surface) ClearDamage() { s.damage = geom.ZeroRect }

// ClearDamageTree 清空整棵子树的脏区（帧末提交后调用）。
func (s *Surface) ClearDamageTree() {
	s.Walk(func(c *Surface) bool {
		c.damage = geom.ZeroRect
		return true
	})
}

// MarkStructureDirty 标记结构/几何发生变化。
func (s *Surface) MarkStructureDirty() {
	s.surfaceDirty = true
	if s.parent != nil {
		s.parent.MarkStructureDirty()
	}
}

// StructureDirty 表示子树需要重新布局或重建合成层。
func (s *Surface) StructureDirty() bool { return s.surfaceDirty }

// ClearStructureDirty 清除结构性脏标记（整棵子树）。
func (s *Surface) ClearStructureDirty() {
	s.Walk(func(c *Surface) bool {
		c.surfaceDirty = false
		return true
	})
}

// markAllDirty 在挂接时把整棵子树标脏，避免复用旧 damage。
func (s *Surface) markAllDirty() {
	s.Walk(func(c *Surface) bool {
		c.surfaceDirty = true
		c.damage = c.LocalBounds()
		return true
	})
}
