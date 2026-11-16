package composite

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
)

// Reason 说明"这一层为什么存在"（契约 C19：每一层都必须给出理由）。
//
// 这不是装饰性字段。层是有代价的（一次离屏分配、一次整体混合、一份缓存），
// 没有理由的层就是纯损失。因此构建期的每一条分支都必须能回答"不建层会怎样"：
//
//   - 不建 Opacity 层 → 子树逐指令乘 alpha，重叠处重复变暗（错）；
//   - 不建 Clip 层 → 光栅层只会矩形裁剪，圆角处画出去（错）；
//   - 不建 RepaintBoundary 层 → 无法把一块静止内容拦成可缓存的图（慢）。
//
// 反过来，"有理由也不该建层"的情况同样要写清楚（见 reasonOf 的 ReasonTransform）。
type Reason uint8

const (
	// ReasonNone 表示未声明理由——**非法状态**，Validate 会报出来。
	// 用零值占位是为了让"忘了设置"变成可检测的，而不是悄悄变成第一种理由。
	ReasonNone Reason = iota

	// ReasonRoot 是合成器插入的根层（全树唯一）。
	ReasonRoot

	// ReasonContent 是一条指令切片：它的语义由所属的组层决定，本身不独立。
	// 这样"层数"仍然可以按组层数统计，而 picture 切片不参与缓存与理由断言。
	ReasonContent

	// ReasonOpacity 是组透明度（opacity < 1）。
	ReasonOpacity

	// ReasonClip 是非矩形裁剪（圆角/路径）：光栅层只能盒裁剪，圆角必须走蒙版。
	ReasonClip

	// ReasonFilter 是组滤镜（模糊/发光/组阴影）。**M2 未实现**，见 reasonOf。
	ReasonFilter

	// ReasonTransform 是非轴对齐变换。
	//
	// M2 **不为此建层**：变换已经固化进每条指令的 Matrix（见 render.Command），
	// 建层的唯一收益是"纹理只变换一次"，而软件后端没有纹理管线，
	// 建层只会多一次离屏 + 一次混合。理由仍然声明，留给 GPU 后端与 M4。
	ReasonTransform

	// ReasonRepaintBoundary 是显式重绘边界：为了把静止子树拦成一张可复用的图。
	ReasonRepaintBoundary

	// ReasonSurface 是子画布边界（视频/原生窗口/独立 DPI）。
	// M2 降级为扁平化，真正的子表面在 M3 随 host 落地。
	ReasonSurface
)

// String 返回可读名（诊断输出用）。
func (r Reason) String() string {
	switch r {
	case ReasonNone:
		return "none"
	case ReasonRoot:
		return "root"
	case ReasonContent:
		return "content"
	case ReasonOpacity:
		return "opacity"
	case ReasonClip:
		return "clip"
	case ReasonFilter:
		return "filter"
	case ReasonTransform:
		return "transform"
	case ReasonRepaintBoundary:
		return "repaint-boundary"
	case ReasonSurface:
		return "surface"
	default:
		return "unknown"
	}
}

// LayerKind 是层种类（封闭集合，和 scene.Kind 一样用具体 struct 而不是接口）。
type LayerKind uint8

const (
	// LayerPicture 只有一段指令，没有子层。
	LayerPicture LayerKind = iota
	// LayerGroup 有子层，可能需要整体离屏合成。
	LayerGroup
)

// String 返回可读名。
func (k LayerKind) String() string {
	if k == LayerGroup {
		return "group"
	}
	return "picture"
}

// ClipKind 是层的裁剪方式。
type ClipKind uint8

const (
	// ClipNone 不裁剪。
	ClipNone ClipKind = iota
	// ClipRect 矩形裁剪（交给指令流的 scissor，不需要蒙版）。
	ClipRect
	// ClipRRect 圆角矩形裁剪（需要蒙版，AA 由蒙版给）。
	ClipRRect
	// ClipPath 路径裁剪（需要蒙版）。
	ClipPath
)

// String 返回可读名。
func (c ClipKind) String() string {
	switch c {
	case ClipRect:
		return "rect"
	case ClipRRect:
		return "rrect"
	case ClipPath:
		return "path"
	default:
		return "none"
	}
}

// Layer 是合成层。
//
// 坐标系（唯一一条，写在这里免得各处再猜）：
//   - 层的 bounds/picture/children 的 origin 全部在**本层局部空间**；
//   - 单位是**设备像素**（录制时矩阵已经把 DIP 乘进去了）；
//   - l.origin 是本层局部原点在**父层局部空间**里的位置；
//   - l.absOrigin 是本层局部原点在**设备空间**里的位置（跳过剔除/遮挡用）。
//
// 层局部空间而不是设备空间，是为了让"整层平移"不改变层内容：
// 缓存键里因此没有位置，只有尺寸 + 内容签名（滚动列表就靠这条）。
type Layer struct {
	kind     LayerKind
	reason   Reason
	node     *scene.Node // nil 表示合成器插入的层（根层）
	parent   *Layer
	children []*Layer

	bounds    geom.Rect // 本层内容的包围盒（局部空间）
	origin    geom.Point
	absOrigin geom.Point

	picture *render.DisplayList // 局部空间

	// 组属性（只有 LayerGroup 有意义）
	opacity float32
	clip    ClipKind
	// clipRect 是局部空间的裁剪矩形；radius/path 二选一。
	clipRect   geom.Rect
	clipRadius float32
	clipPath   *render.Path

	// version 是内容签名：由子树里所有节点的 (id, signature) 混合而成。
	// 它刻意**不**包含位置——位置变了内容没变，缓存应当继续命中。
	version uint64
	depth   int

	// offscreen 表示本层必须经过离屏缓冲才能正确合成。
	offscreen bool
	// cacheable 表示离屏缓冲可以留在 RasterCache 里复用（C7：只有合成器持有）。
	cacheable bool
	// cached 表示本帧的内容直接来自缓存（没有重新栅格化）。
	cached bool
	// filter 是组滤镜（模糊/发光/组阴影）；nil 表示无。
	filter *scene.EffectSpec
	// boxShadow 是节点自身盒阴影；仅当 Blur>0 时存在（硬阴影仍在
	// scene 的指令流里直接画，不需要开层）。
	boxShadow *scene.ShadowSpec
	// boxRadius 是盒阴影对应的圆角半径。
	boxRadius float32
	// nodeSize 是节点尺寸（盒阴影的形状来源）。
	nodeSize geom.Size
}

// Kind 返回层种类。
func (l *Layer) Kind() LayerKind { return l.kind }

// Reason 返回建层理由。
func (l *Layer) Reason() Reason { return l.reason }

// Node 返回来源节点（根层为 nil）。
func (l *Layer) Node() *scene.Node { return l.node }

// Parent 返回父层。
func (l *Layer) Parent() *Layer { return l.parent }

// Children 返回子层（按绘制顺序）。
func (l *Layer) Children() []*Layer { return l.children }

// Bounds 返回局部空间的内容包围盒。
func (l *Layer) Bounds() geom.Rect { return l.bounds }

// Origin 返回局部原点在父层局部空间中的位置。
func (l *Layer) Origin() geom.Point { return l.origin }

// AbsOrigin 返回局部原点在设备空间中的位置。
func (l *Layer) AbsOrigin() geom.Point { return l.absOrigin }

// AbsBounds 返回设备空间的包围盒。
func (l *Layer) AbsBounds() geom.Rect { return l.bounds.Offset(l.absOrigin.X, l.absOrigin.Y) }

// Picture 返回本层指令（只对 LayerPicture 非 nil）。
func (l *Layer) Picture() *render.DisplayList { return l.picture }

// Opacity 返回组透明度。
func (l *Layer) Opacity() float32 { return l.opacity }

// Clip 返回裁剪方式。
func (l *Layer) Clip() ClipKind { return l.clip }

// ClipRect 返回局部空间的裁剪矩形。
func (l *Layer) ClipRect() geom.Rect { return l.clipRect }

// ClipRadius 返回圆角半径。
func (l *Layer) ClipRadius() float32 { return l.clipRadius }

// ClipPath 返回裁剪路径。
func (l *Layer) ClipPath() *render.Path { return l.clipPath }

// Version 返回内容签名。
func (l *Layer) Version() uint64 { return l.version }

// Depth 返回层深度（根为 0）。
func (l *Layer) Depth() int { return l.depth }

// Offscreen 表示本层需要离屏缓冲。
func (l *Layer) Offscreen() bool { return l.offscreen }

// Cacheable 表示离屏缓冲可进 RasterCache。
func (l *Layer) Cacheable() bool { return l.cacheable }

// Cached 表示本帧内容来自缓存。
func (l *Layer) Cached() bool { return l.cached }

// Count 返回子树层数（含自身）。
func (l *Layer) Count() int {
	n := 1
	for _, c := range l.children {
		n += c.Count()
	}
	return n
}

// Walk 前序遍历子树，fn 返回 false 时不再深入该分支。
func (l *Layer) Walk(fn func(*Layer) bool) {
	if l == nil || !fn(l) {
		return
	}
	for _, c := range l.children {
		c.Walk(fn)
	}
}

// computeBounds 自底向上求出组层的内容包围盒（局部空间）。
//
// 组层不会在绘制时自动裁剪到节点矩形（否则阴影/溢出内容会被切），
// 所以包围盒取"子层包围盒并集"，而不是节点尺寸——这一点上
// 离屏分配大小与"节点画了多大"是两件事。
func (l *Layer) computeBounds() {
	if l.kind == LayerPicture {
		if l.picture != nil {
			l.bounds = l.picture.Bounds()
		}
		return
	}
	var u geom.Rect
	for _, c := range l.children {
		c.computeBounds()
		if c.bounds.IsEmpty() {
			continue
		}
		u = u.Union(c.bounds.Offset(c.origin.X, c.origin.Y))
	}
	if l.clip != ClipNone && !l.clipRect.IsEmpty() {
		// 裁剪到可视区：蒙版外的像素反正不参与混合，不必为它们分配离屏。
		u = u.Intersect(l.clipRect.RoundOut())
	}
	l.bounds = u
}

// setAbsOrigin 自顶向下设置设备空间原点。
func (l *Layer) setAbsOrigin() {
	for _, c := range l.children {
		c.absOrigin = l.absOrigin.Add(c.origin)
		c.setAbsOrigin()
	}
}

// backingSize 返回本层离屏缓冲应有的像素尺寸。
//
// 用 RoundOut 再取整而不是 ceil：AA 边缘的半像素覆盖落在像素格内，
// 少取一个像素就会把边缘切掉一条（golden 比对能看出来）。
func (l *Layer) backingSize() (int, int) {
	if l.bounds.IsEmpty() {
		return 0, 0
	}
	r := l.bounds.RoundOut()
	return int(r.W()), int(r.H())
}
