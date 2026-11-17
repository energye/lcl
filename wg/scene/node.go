// Package scene 是渲染树与布局：Node、Constraints、measure/arrange、录制。
//
// 它只描述"多大、放哪、谁在上、画什么"，**不持有像素**，也不碰 lcl。
//
// # 为什么 Node 是具体 struct 而不是 interface
//
// 设计文档 03-ENGINE.md §2.1 规定节点是**封闭集合**（15 种）。封闭集合一旦
// 承认，interface 的唯一作用就消失了：不需要用类型断言分派、不需要为每种
// 节点分配不同的堆对象、字段访问也不需要经过接口间接层。整棵树因此可以是
// 扁平的对象图，遍历成本与 `[]*Node` 一致。
//
// 代价是 `Node` 是个大 struct（约 40 个字段）。这是**刻意的**：UI 节点的
// 数量级是 10³–10⁵，而节点类型的数量级是 10¹。为前者优化访问局部性，
// 比后者那点内存浪费重要得多。
//
// Kind 决定**布局语义**，paint spec（Box/Text/Path/Image）决定**画什么**。
// 两者必须一致，`Validate` 负责抓不一致——见 node_test.go。
//
// 依赖：geom / render / text / theme（见 05-API-ROADMAP §1 的 C17 表）。
package scene

import (
	"math"
	"sync/atomic"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/text"
)

// Kind 是节点类型（封闭集合，15 种）。
type Kind uint8

const (
	// KindEmpty 只有布局与命中，不画（间距、点击热区、焦点环承载）。
	KindEmpty Kind = iota
	// KindBox 背景/边框/圆角/阴影 + 子节点按对齐叠放。
	KindBox
	// KindText 单/多行文本。
	KindText
	// KindPath 任意路径填充/描边（图标、图表、连线）。
	KindPath
	// KindImage 位图（含 fit 缩放）。
	KindImage
	// KindStack 层叠定位（z 序 = 子节点顺序）。
	KindStack
	// KindFlex 主轴排列（Wrap=true 时换行）。
	KindFlex
	// KindGrid 行列 + span + gutter。
	KindGrid
	// KindScroll 视口 + 内容偏移（单子节点）。
	KindScroll
	// KindClip 裁剪（矩形/圆角/路径），单子节点。
	KindClip
	// KindTransform 2D 仿射，单子节点；不影响布局尺寸。
	KindTransform
	// KindOpacity 组透明度，单子节点。
	KindOpacity
	// KindEffect 模糊/阴影/发光作用于整组，单子节点。
	KindEffect
	// KindSurface 子画布边界，单子节点。
	KindSurface
	// KindCustomPaint 逃生舱：回调里用 Canvas 录制。
	KindCustomPaint
)

// KindCount 是 Kind 的总数（断言与遍历用）。
const KindCount = 15

// String 返回可读名。
func (k Kind) String() string {
	switch k {
	case KindEmpty:
		return "empty"
	case KindBox:
		return "box"
	case KindText:
		return "text"
	case KindPath:
		return "path"
	case KindImage:
		return "image"
	case KindStack:
		return "stack"
	case KindFlex:
		return "flex"
	case KindGrid:
		return "grid"
	case KindScroll:
		return "scroll"
	case KindClip:
		return "clip"
	case KindTransform:
		return "transform"
	case KindOpacity:
		return "opacity"
	case KindEffect:
		return "effect"
	case KindSurface:
		return "surface"
	case KindCustomPaint:
		return "custom-paint"
	default:
		return "?"
	}
}

// ClipOp 是裁剪运算。
type ClipOp uint8

const (
	// ClipIntersect 与已有裁剪求交（默认）。
	ClipIntersect ClipOp = iota
	// ClipDifference 从已有裁剪中减去（挖洞）。
	ClipDifference
)

// ClipSpec 是节点的裁剪描述。
//
// 矩形裁剪走 scissor（快）；圆角/路径裁剪走覆盖率（慢但正确）。
// 裁剪矩形默认是"节点自身内容区"，即 (0,0,size)。
type ClipSpec struct {
	// Rect 是相对节点内容原点的裁剪矩形；零值表示用节点尺寸。
	Rect geom.Rect
	// Radius > 0 时按圆角裁剪。
	Radius float32
	// Path 非 nil 时按路径裁剪（优先级高于 Radius）。
	Path *render.Path
	// Op 是裁剪运算。
	Op ClipOp
	// AA 是否抗锯齿（圆角/路径裁剪有效）。
	AA bool
}

// BoxSpec 是 KindBox 的绘制描述。
//
// Fill 与 Border 都是 `render.Paint` 值（可比较、可批处理），
// 因此"hover 换个颜色"只是换一个字段值，不需要重建节点。
type BoxSpec struct {
	Fill   render.Paint
	Border render.Paint
	Radius float32
	// Shadow 是投影参数；nil 表示不投影。
	Shadow *ShadowSpec
}

// ShadowSpec 是阴影（dx,dy,blur,color）。
type ShadowSpec struct {
	Dx, Dy float32
	Blur   float32
	Color  geom.Color
}

// TextSpec 是 KindText 的绘制描述。
type TextSpec struct {
	// Src 是原始文本（重排时需要）。
	Src    string
	Layout *text.Layout
	Color  geom.Color
	// Composing 是 IME 预编辑区（rune 区间），用于画下划线；[0,0) 表示无。
	ComposingStart, ComposingEnd int32
	// Selection 是选区（rune 区间）与选区色。
	SelStart, SelEnd int32
	SelectionColor   geom.Color
	// MaxWidth/MaxLines/Ellipsis 参与重新排版（Resize 或主题变化时用）。
	MaxWidth float32
	MaxLines int
	Ellipsis string
	// LayoutFont 是排版用的字体。
	LayoutFont text.Font
	// LayoutW 是当前 Layout 是按多宽排的（0 = 不换行），用于判断要不要重排。
	LayoutW float32
}

// PathSpec 是 KindPath 的绘制描述。
type PathSpec struct {
	Path  *render.Path
	Paint render.Paint
}

// ImageSpec 是 KindImage 的绘制描述。
type ImageSpec struct {
	Image *render.ImageRef
	Fit   Fit
	Paint render.Paint
}

// EffectSpec 是 KindEffect 的效果描述（M2 只落阴影与不透明度近似）。
type EffectSpec struct {
	// BlurSigma > 0 时做组模糊。
	BlurSigma float32
	// Shadow 非 nil 时在组下方投影。
	Shadow *ShadowSpec
	// GlowRadius > 0 时在组外发光。
	GlowRadius float32
	GlowColor  geom.Color
}

// CustomPaint 是逃生舱回调：只允许用 render.Canvas 录制，不能碰像素。
type CustomPaint func(c render.Canvas, size geom.Size)

// HitBehavior 是命中穿透行为。
type HitBehavior uint8

const (
	// HitDeferToChild 自身不命中，只作为路径的一环（纯布局容器默认）。
	HitDeferToChild HitBehavior = iota
	// HitOpaque 命中即终止穿透。
	HitOpaque
	// HitTranslucent 自身可命中，同时允许继续向下（悬浮按钮、角标）。
	HitTranslucent
)

// HitShape 是命中形状。
type HitShape uint8

const (
	// ShapeRect 矩形（默认）。
	ShapeRect HitShape = iota
	// ShapeRRect 圆角矩形。
	ShapeRRect
	// ShapeCircle 椭圆/圆。
	ShapeCircle
	// ShapePath 任意路径。
	ShapePath
)

// Cursor 是鼠标指针形状（与平台无关的枚举）。
type Cursor uint8

const (
	// CursorDefault 默认箭头。
	CursorDefault Cursor = iota
	// CursorPointer 手型（可点击）。
	CursorPointer
	// CursorText 文本插入符。
	CursorText
	// CursorCrosshair 十字。
	CursorCrosshair
	// CursorMove 移动。
	CursorMove
	// CursorResizeEW 水平调整。
	CursorResizeEW
	// CursorResizeNS 垂直调整。
	CursorResizeNS
	// CursorNotAllowed 禁止。
	CursorNotAllowed
)

// HitSpec 是节点的命中描述。
//
// 形状命中与绘制用**同一份几何**（圆角半径、路径），避免"点方角也响应"。
type HitSpec struct {
	Behavior HitBehavior
	Shape    HitShape
	// Radius 用于 ShapeRRect。
	Radius float32
	// Path 用于 ShapePath。
	Path *render.Path
	// Absorb 吞掉事件且不让下层收到（模态遮罩）。
	Absorb bool
	Cursor Cursor
	// Tooltip 非空时，hover 800ms 后在 Tooltip 层带弹出。
	Tooltip string
	// Label 语义标签（自动化测试选择器用）。
	Label string
}

// idSeq 是全局节点 id 分配器。节点 id 参与缓存键与命中缓存，"重排不换 id"很重要。
var idSeq atomic.Uint64

// Node 是渲染树节点（见包注释：具体 struct 而非 interface）。
type Node struct {
	id     uint64
	kind   Kind
	key    string
	parent *Node
	// children 的顺序就是绘制顺序（后面的画在上面）；命中按逆序。
	children []*Node

	// ── 布局 ─────────────────────────────────────────────────────────
	constraints Constraints
	size        geom.Size
	// offset 是相对**父节点内容原点**的左上角位置，由 arrange 写入。
	offset geom.Point

	margin  geom.Insets
	padding geom.Insets
	// align 是 Box/Stack 内子节点的默认对齐。
	align Alignment

	// 尺寸覆盖：NaN 表示"由约束与内容决定"。
	fixedW, fixedH float32
	minW, maxW     float32
	minH, maxH     float32
	// ratio > 0 时按宽高比收口（宽定高、或高定宽）。
	ratio float32

	// flex 是 Flex 主轴权重（0 = 不参与分配）。
	flex    float32
	flexFit FlexFit
	// position 是 Stack 内的绝对定位偏移（positioned=true 时生效）。
	position   geom.Point
	positioned bool
	// keepSize 表示绝对定位子节点保持自身尺寸（不拉满父内容区）。
	keepSize bool
	// span 是 Grid 跨列数（0/1 = 一列）。
	span int

	// Flex 参数
	axis       Axis
	gap        float32
	runGap     float32
	mainAlign  MainAlign
	crossAlign CrossAlign
	wrap       bool

	// Grid 参数
	columns int

	// Scroll 参数
	contentSize  geom.Size
	scrollOffset float32

	// ── 绘制 ─────────────────────────────────────────────────────────
	transform geom.Matrix
	// originAlign 是变换支点（归一化对齐；零值表示 Center）。
	originAlign Alignment
	opacity     float32
	clip        *ClipSpec
	box         *BoxSpec
	text        *TextSpec
	path        *PathSpec
	image       *ImageSpec
	effect      *EffectSpec
	custom      CustomPaint

	// paint 是本节点（含子树）的录制结果；nil 表示还没录。
	paint *render.DisplayList
	// signature 是"视觉签名"，重录时递增，参与缓存键。
	signature uint64

	// ── 合成 ─────────────────────────────────────────────────────────
	// repaintBoundary 显式声明"这里是可以整体复用的重绘边界"。
	repaintBoundary bool
	// overflow 表示子内容允许逃逸父裁剪（Popup 用）。
	overflow bool

	// ── 命中 ─────────────────────────────────────────────────────────
	hit     HitSpec
	visible bool
	enabled bool

	// ── 脏标记（三档，必须独立） ──────────────────────────────────────
	//
	// 常见错误是把三者合并成一个 dirty：改个 hover 颜色会触发整棵树布局，
	// 60 FPS 立刻掉到 30。
	needsLayout    bool
	needsPaint     bool
	needsComposite bool
	// selfDirty 表示"**这个节点自己的**视觉变了"（不含被子节点传染）。
	//
	// 为什么必须与 needsPaint 分开：needsPaint 会向上冒泡（父节点要重录，
	// 因为录制流是一段），但如果脏区也按冒泡算，改一个 8px 方块就会得到
	// 整屏脏区——局部重绘等于没做。DirtyRect 只认 selfDirty。
	selfDirty bool
	// onInvalidate 只在根节点上生效：子树任何层级变脏都会调到它。
	onInvalidate func()
	// layoutBoundary 表示本节点是布局边界：约束不变时其子树不重排。
	layoutBoundary bool
	// laidOut 表示布局已算过；lastConstraints 是上次的输入，用于命中缓存。
	laidOut         bool
	lastConstraints Constraints
}

// ── 构造 ────────────────────────────────────────────────────────────

func newNode(kind Kind) *Node {
	return &Node{
		id:         idSeq.Add(1),
		kind:       kind,
		visible:    true,
		enabled:    true,
		opacity:    1,
		transform:  geom.Identity,
		fixedW:     nan,
		fixedH:     nan,
		minW:       nan,
		maxW:       nan,
		minH:       nan,
		maxH:       nan,
		align:      TopLeft,
		mainAlign:  MainStart,
		crossAlign: CrossStart,
		axis:       Vertical,
		columns:    1,
		flexFit:    FlexTight,
	}
}

// Empty 创建不绘制、只占位的节点。
func Empty() *Node { return newNode(KindEmpty) }

// Box 创建背景/边框/圆角节点。children 按 align 叠放在 padding 内。
func Box(box BoxSpec, children ...*Node) *Node {
	n := newNode(KindBox)
	spec := box
	n.box = &spec
	n.Add(children...)
	return n
}

// Text 创建文本节点。layout 为 nil 时按空文本处理（尺寸 0）。
func Text(spec TextSpec) *Node {
	n := newNode(KindText)
	s := spec
	n.text = &s
	if s.Layout != nil {
		n.size = s.Layout.Size
	}
	return n
}

// Path 创建路径节点。
func PathNode(spec PathSpec) *Node {
	n := newNode(KindPath)
	s := spec
	n.path = &s
	return n
}

// Image 创建位图节点。
func ImageNode(spec ImageSpec) *Node {
	n := newNode(KindImage)
	s := spec
	n.image = &s
	return n
}

// Stack 创建层叠节点。
func Stack(children ...*Node) *Node {
	n := newNode(KindStack)
	n.Add(children...)
	return n
}

// Flex 创建主轴排列节点（Row/Column）。
func Flex(axis Axis, gap float32, children ...*Node) *Node {
	n := newNode(KindFlex)
	n.axis = axis
	n.gap = gap
	n.Add(children...)
	return n
}

// Wrap 创建换行排列节点。
func Wrap(axis Axis, gap, runGap float32, children ...*Node) *Node {
	n := Flex(axis, gap, children...)
	n.wrap = true
	n.runGap = runGap
	return n
}

// Grid 创建行列节点。
func Grid(columns int, gap float32, children ...*Node) *Node {
	n := newNode(KindGrid)
	if columns < 1 {
		columns = 1
	}
	n.columns = columns
	n.gap = gap
	n.Add(children...)
	return n
}

// Scroll 创建滚动视口（单子节点）。
func Scroll(axis Axis, child *Node) *Node {
	n := newNode(KindScroll)
	n.axis = axis
	if child != nil {
		n.Add(child)
	}
	return n
}

// Clip 创建裁剪节点（单子节点）。
func Clip(spec ClipSpec, child *Node) *Node {
	n := newNode(KindClip)
	s := spec
	n.clip = &s
	if child != nil {
		n.Add(child)
	}
	return n
}

// Transform 创建仿射变换节点（单子节点，不影响布局尺寸）。
func Transform(m geom.Matrix, child *Node) *Node {
	n := newNode(KindTransform)
	n.transform = m
	if child != nil {
		n.Add(child)
	}
	return n
}

// Opacity 创建组透明度节点（单子节点）。
func Opacity(alpha float32, child *Node) *Node {
	n := newNode(KindOpacity)
	n.opacity = clampf(alpha, 0, 1)
	if child != nil {
		n.Add(child)
	}
	return n
}

// Effect 创建组效果节点（单子节点）。
func Effect(spec EffectSpec, child *Node) *Node {
	n := newNode(KindEffect)
	s := spec
	n.effect = &s
	if child != nil {
		n.Add(child)
	}
	return n
}

// SurfaceNode 创建子画布边界节点（单子节点）。
func SurfaceNode(child *Node) *Node {
	n := newNode(KindSurface)
	if child != nil {
		n.Add(child)
	}
	return n
}

// CustomPaintNode 创建逃生舱节点。
func CustomPaintNode(fn CustomPaint) *Node {
	n := newNode(KindCustomPaint)
	n.custom = fn
	return n
}

// ── 树结构 ──────────────────────────────────────────────────────────

// ID 返回节点 id（稳定，重排不变）。
func (n *Node) ID() uint64 { return n.id }

// Kind 返回节点类型。
func (n *Node) Kind() Kind { return n.kind }

// Key 返回 diff 用的 Key。
//
// 语义：**跨重建稳定的组件身份**。重建每次产出全新节点（id 全换），
// 有状态控件把自己的实例身份（如 "ctl-12"）钉在节点上，输入层的
// hover 差分就能认出"换节点不换控件"，不把重建误判成悬停迁移。
// 留空的节点退回指针同一性。
func (n *Node) Key() string { return n.key }

// SetKey 设置 Key 并返回自身（链式）。
func (n *Node) SetKey(k string) *Node { n.key = k; return n }

// Parent 返回父节点。
func (n *Node) Parent() *Node { return n.parent }

// Children 返回子节点切片（只读视图，不要修改）。
func (n *Node) Children() []*Node { return n.children }

// Add 追加子节点（nil 会被忽略，用于"条件渲染不生成节点"的场景）。
//
// 追加会标 needsLayout + needsPaint：结构变了，尺寸与绘制都可能变。
func (n *Node) Add(children ...*Node) *Node {
	for _, c := range children {
		if c == nil {
			continue
		}
		if c.parent != nil {
			c.parent.Remove(c)
		}
		c.parent = n
		n.children = append(n.children, c)
	}
	if len(children) > 0 {
		n.MarkNeedsLayout()
	}
	return n
}

// Remove 移除直接子节点。
func (n *Node) Remove(c *Node) bool {
	for i, x := range n.children {
		if x == c {
			n.children = append(n.children[:i], n.children[i+1:]...)
			c.parent = nil
			n.MarkNeedsLayout()
			return true
		}
	}
	return false
}

// RemoveAll 清空子节点。
func (n *Node) RemoveAll() *Node {
	for _, c := range n.children {
		c.parent = nil
	}
	n.children = n.children[:0]
	n.MarkNeedsLayout()
	return n
}

// IndexInParent 返回在父 children 中的位置（无父返回 -1）。
func (n *Node) IndexInParent() int {
	if n.parent == nil {
		return -1
	}
	for i, c := range n.parent.children {
		if c == n {
			return i
		}
	}
	return -1
}

// Depth 返回从根算起的深度（根为 0）。
func (n *Node) Depth() int {
	d := 0
	for p := n.parent; p != nil; p = p.parent {
		d++
	}
	return d
}

// Count 返回子树节点数（含自身）。
func (n *Node) Count() int {
	c := 1
	for _, ch := range n.children {
		c += ch.Count()
	}
	return c
}

// ── 布局结果 ────────────────────────────────────────────────────────

// Size 返回布局得到的尺寸。
func (n *Node) Size() geom.Size { return n.size }

// SetSize 直接设置尺寸（仅"直接式"改属性时用，会标 needsPaint）。
func (n *Node) SetSize(s geom.Size) *Node {
	if n.size == s {
		return n
	}
	n.size = s
	n.MarkNeedsPaint()
	return n
}

// Offset 返回相对父内容原点的位置。
func (n *Node) Offset() geom.Point { return n.offset }

// Bounds 返回节点在**父内容坐标**下的矩形。
func (n *Node) Bounds() geom.Rect { return geom.RectXYWH(n.offset.X, n.offset.Y, n.size.W, n.size.H) }

// LocalMatrix 返回节点自身坐标系 → 父坐标系的矩阵。
func (n *Node) LocalMatrix() geom.Matrix {
	return geom.Translate(n.offset.X, n.offset.Y).Concat(n.transform)
}

// ContentOrigin 返回内容区（扣掉 padding）在节点坐标系下的原点。
func (n *Node) ContentOrigin() geom.Point {
	return geom.Pt(n.padding.Left, n.padding.Top)
}

// ContentSize 返回内容区尺寸。
func (n *Node) ContentSize() geom.Size {
	return geom.Sz(
		maxf(0, n.size.W-n.padding.Horizontal()),
		maxf(0, n.size.H-n.padding.Vertical()),
	)
}

// ── 脏标记 ──────────────────────────────────────────────────────────

// MarkNeedsLayout 标记布局脏，并向上冒泡到布局边界（含）。
//
// 之所以要冒泡：子节点变大后父节点的尺寸可能与内容不符（shrink-wrap），
// 但冒泡必须停在布局边界，否则"改一个叶子就重排整棵树"。
func (n *Node) MarkNeedsLayout() {
	n.signature++
	if n.needsLayout {
		return
	}
	n.needsLayout = true
	n.needsPaint = true
	n.needsComposite = true
	if n.layoutBoundary {
		return
	}
	if n.parent != nil {
		n.parent.MarkNeedsLayout()
	}
}

// MarkNeedsPaint 只标绘制脏（颜色、文本内容变化）。
//
// signature **无条件**递增：它是"视觉版本号"，参与合成层的缓存键
// （见 compositor 的 CacheKey）。脏标记的幂等去重只作用于标记与传播，
// 不能作用于版本号——否则"同一帧里改两次属性"会让缓存认为内容没变，
// 于是贴出上一帧的图。这一条在 --layers 的 cache-invalidates 上挂着。
func (n *Node) MarkNeedsPaint() {
	n.signature++
	// selfDirty 必须在去重之前置位：同一帧里改两次属性时 needsPaint 已经是
	// true，但"这个节点的像素变了"这个事实不能因此丢掉（脏区要用它）。
	n.selfDirty = true
	if n.needsPaint {
		return
	}
	n.needsPaint = true
	n.invalidate()
}

// MarkNeedsComposite 标记合成脏（opacity/transform/clip/层级变化）。
// signature 无条件递增，理由同 MarkNeedsPaint。
func (n *Node) MarkNeedsComposite() {
	n.signature++
	n.selfDirty = true
	if n.needsComposite {
		return
	}
	n.needsComposite = true
	n.invalidate()
}

// invalidate 把自身 bounds 上卷（渲染树没有 Surface 的 damage 机制时，
// 由 Container 在帧末统一收集——这里只标脏，不做几何）。
//
// 同时通知**根节点**上的变脏回调：树本身不知道帧循环在哪，而"有东西变了"
// 是帧泵唯一愿意被叫醒的理由。没有这层通知，改完节点必须由调用方手工
// RequestFrame——漏一处就是"界面不动，鼠标划过才刷新"。
func (n *Node) invalidate() {
	if n.parent != nil {
		n.parent.needsPaint = true
	}
	root := n
	for root.parent != nil {
		root = root.parent
	}
	if root.onInvalidate != nil {
		root.onInvalidate()
	}
}

// SetInvalidateHook 在根节点上安装"子树变脏"回调（Container 用它申请帧）。
//
// 只有挂在根上的回调会生效：节点的 invalidate 会向上找到根再调，
// 因此子树里的任何层级都能触发，不需要逐节点安装。
func (n *Node) SetInvalidateHook(fn func()) {
	n.onInvalidate = fn
}

// NeedsLayout 报告布局是否脏。
func (n *Node) NeedsLayout() bool { return n.needsLayout }

// NeedsPaint 报告绘制是否脏。
func (n *Node) NeedsPaint() bool { return n.needsPaint }

// NeedsComposite 报告合成是否脏。
func (n *Node) NeedsComposite() bool { return n.needsComposite }

// ClearDirty 清三档脏标记（帧末由 Container 调用）。
func (n *Node) ClearDirty() {
	n.needsLayout = false
	n.needsPaint = false
	n.needsComposite = false
	n.selfDirty = false
}

// ClearDirtyTree 递归清脏标记。
func (n *Node) ClearDirtyTree() {
	n.ClearDirty()
	for _, c := range n.children {
		c.ClearDirtyTree()
	}
}

// Dirty 是树级脏标记汇总。
type Dirty struct {
	Layout    bool
	Paint     bool
	Composite bool
}

// Any 表示是否有任何一档脏。
func (d Dirty) Any() bool { return d.Layout || d.Paint || d.Composite }

// String 返回可读名（诊断输出用）。
func (d Dirty) String() string {
	if !d.Any() {
		return "clean"
	}
	s := ""
	if d.Layout {
		s += "L"
	}
	if d.Paint {
		s += "P"
	}
	if d.Composite {
		s += "C"
	}
	return s
}

// AnyDirty 汇总整棵子树的脏标记。
//
// 三档分开返回而不是合成一个 bool：布局脏要重排 + 重录，绘制脏只要重录，
// 合成脏连重录都省了（只重建层树）。合成一个 bool 会让"只改了个透明度"
// 也走全量重排——那正是"改颜色导致整页闪烁"的来源。
//
// 复杂度 O(n)。1000 节点每帧走一遍约几微秒，相对录制/光栅可忽略；
// 换成"节点里存一个向上传递的计数"会引入"忘了减一就永远脏"的调试噩梦。
func (n *Node) AnyDirty() Dirty {
	d := Dirty{Layout: n.needsLayout, Paint: n.needsPaint, Composite: n.needsComposite}
	for _, c := range n.children {
		cd := c.AnyDirty()
		d.Layout = d.Layout || cd.Layout
		d.Paint = d.Paint || cd.Paint
		d.Composite = d.Composite || cd.Composite
	}
	return d
}

// DirtyRect 返回脏节点在根空间的包围盒并集（无脏时为零矩形）。
//
// 用途：把"哪些节点变了"换算成"要重画哪块像素"。没有它就只能整屏重画，
// 或者反过来只标脏不重画（画面停在上一帧，看起来像卡死）。
//
// 只统计 **selfDirty**（自己变了）的节点：needsPaint 会向上冒泡到祖先，
// 若按它算脏区，改一个 8px 方块会得到整屏脏区——局部重绘等于没做。
// 祖先的 needsPaint 只说明"要重录"，不代表它的像素变了。
//
// 前置条件：树已完成布局——BoundsInRoot 读的是当前 offset/size。
// 布局脏的帧调用方应该整屏重画，不要用这个并集。
func (n *Node) DirtyRect() geom.Rect {
	var out geom.Rect
	var walk func(*Node)
	walk = func(x *Node) {
		if x.selfDirty {
			r := x.BoundsInRoot()
			if !r.IsEmpty() {
				if out.IsEmpty() {
					out = r
				} else {
					out = out.Union(r)
				}
			}
		}
		for _, c := range x.children {
			walk(c)
		}
	}
	walk(n)
	return out
}

// Signature 返回视觉签名（重录/缓存键的一部分）。
func (n *Node) Signature() uint64 { return n.signature }

// DisplayList 返回最近一次录制结果（可能为 nil）。
func (n *Node) DisplayList() *render.DisplayList { return n.paint }

// SetDisplayList 写入录制结果（合成阶段用）。
func (n *Node) SetDisplayList(dl *render.DisplayList) { n.paint = dl }

// ── 属性读写（链式，便于声明式构造） ────────────────────────────────

// SetVisible 设置可见性（不可见节点不参与绘制与命中，但仍占布局空间）。
func (n *Node) SetVisible(v bool) *Node {
	if n.visible != v {
		n.visible = v
		n.MarkNeedsPaint()
	}
	return n
}

// Visible 报告是否可见。
func (n *Node) Visible() bool { return n.visible }

// SetEnabled 设置可用性（不可用不参与命中，视觉上通常叠加禁用色）。
func (n *Node) SetEnabled(v bool) *Node {
	if n.enabled != v {
		n.enabled = v
		n.MarkNeedsPaint()
	}
	return n
}

// Enabled 报告是否可用。
func (n *Node) Enabled() bool { return n.enabled }

// SetOpacity 设置组透明度（<1 时合成阶段会产生 OpacityLayer）。
func (n *Node) SetOpacity(a float32) *Node {
	a = clampf(a, 0, 1)
	if n.opacity != a {
		n.opacity = a
		n.MarkNeedsComposite()
	}
	return n
}

// Opacity 返回透明度。
func (n *Node) Opacity() float32 { return n.opacity }

// SetTransform 设置仿射变换（非轴对齐时合成阶段产生 TransformLayer）。
func (n *Node) SetTransform(m geom.Matrix) *Node {
	n.transform = m
	n.MarkNeedsComposite()
	return n
}

// SetClip 设置裁剪。
func (n *Node) SetClip(spec *ClipSpec) *Node {
	n.clip = spec
	n.MarkNeedsComposite()
	return n
}

// Transform 返回仿射变换。
func (n *Node) Transform() geom.Matrix { return n.transform }

// TransformOrigin 返回变换支点（节点局部坐标，默认节点中心）。
//
// 合成层构建要用它复现边帧；不在那里硬编码"sized/2"是因为
// SetTransformOrigin 可以改，两处各算一遍迟早会分叉。
func (n *Node) TransformOrigin() geom.Point {
	a := n.transformOriginAlign()
	return a.AlignChild(n.size.W, n.size.H, 0, 0)
}

// Clip 返回裁剪描述；nil 表示不裁剪。
func (n *Node) Clip() *ClipSpec { return n.clip }

// Effect 返回组效果描述；nil 表示无效果。
func (n *Node) Effect() *EffectSpec { return n.effect }

// RepaintBoundary 表示该节点要求独立重绘与独立缓存。
func (n *Node) RepaintBoundary() bool { return n.repaintBoundary }

// Overflow 表示允许子树超出自身边界绘制（不自动裁剪）。
func (n *Node) Overflow() bool { return n.overflow }

// SetAlign 设置子节点对齐（Box/Stack 生效）。
func (n *Node) SetAlign(a Alignment) *Node {
	n.align = a
	n.MarkNeedsLayout()
	return n
}

// SetPadding 设置内边距。
func (n *Node) SetPadding(ins geom.Insets) *Node {
	n.padding = ins
	n.MarkNeedsLayout()
	return n
}

// SetMargin 设置外边距。
func (n *Node) SetMargin(ins geom.Insets) *Node {
	n.margin = ins
	n.MarkNeedsLayout()
	return n
}

// SetWidth 固定宽度（NaN 恢复自动）。
func (n *Node) SetWidth(w float32) *Node { n.fixedW = w; n.MarkNeedsLayout(); return n }

// SetHeight 固定高度（NaN 恢复自动）。
func (n *Node) SetHeight(h float32) *Node { n.fixedH = h; n.MarkNeedsLayout(); return n }

// SetSize2 固定宽高。
func (n *Node) SetSize2(w, h float32) *Node {
	n.fixedW, n.fixedH = w, h
	n.MarkNeedsLayout()
	return n
}

// SetMinSize 设置最小尺寸。
func (n *Node) SetMinSize(w, h float32) *Node {
	n.minW, n.minH = w, h
	n.MarkNeedsLayout()
	return n
}

// SetMaxSize 设置最大尺寸。
func (n *Node) SetMaxSize(w, h float32) *Node {
	n.maxW, n.maxH = w, h
	n.MarkNeedsLayout()
	return n
}

// SetAspectRatio 设置宽高比（<=0 清除）。
func (n *Node) SetAspectRatio(r float32) *Node { n.ratio = r; n.MarkNeedsLayout(); return n }

// SetFlex 设置 flex 权重与 fit。
func (n *Node) SetFlex(f float32, fit FlexFit) *Node {
	n.flex, n.flexFit = f, fit
	n.MarkNeedsLayout()
	return n
}

// SetStackPosition 设置 Stack 内的绝对定位偏移。
func (n *Node) SetStackPosition(p geom.Point) *Node {
	n.position = p
	n.positioned = true
	n.MarkNeedsLayout()
	return n
}

// SetPositionKeepSize 让绝对定位的子节点保持自身尺寸。
//
// 默认行为（positioned → 紧约束拉满内容区）对"铺满的遮罩/背景"是对的，
// 但对浮层、Tooltip、拖拽影子是灾难：它们由浮动位置决定尺寸，
// 拉满父级会变成一个盖住整屏的方块。
func (n *Node) SetPositionKeepSize(v bool) *Node {
	n.keepSize = v
	n.MarkNeedsLayout()
	return n
}

// KeepSize 返回是否保持自身尺寸。
func (n *Node) KeepSize() bool { return n.keepSize }

// SetSpan 设置 Grid 跨列数。
func (n *Node) SetSpan(s int) *Node {
	if s < 1 {
		s = 1
	}
	n.span = s
	n.MarkNeedsLayout()
	return n
}

// SetGap 设置主轴间距。
func (n *Node) SetGap(g float32) *Node { n.gap = g; n.MarkNeedsLayout(); return n }

// SetMainAlign 设置主轴对齐。
func (n *Node) SetMainAlign(a MainAlign) *Node { n.mainAlign = a; n.MarkNeedsLayout(); return n }

// SetCrossAlign 设置交叉轴对齐。
func (n *Node) SetCrossAlign(a CrossAlign) *Node { n.crossAlign = a; n.MarkNeedsLayout(); return n }

// SetHit 设置命中描述。
func (n *Node) SetHit(h HitSpec) *Node { n.hit = h; return n }

// Hit 返回命中描述（值拷贝）。
func (n *Node) HitSpec() HitSpec { return n.hit }

// SetRepaintBoundary 显式声明重绘边界（对应 ReasonCache 层）。
func (n *Node) SetRepaintBoundary(v bool) *Node {
	n.repaintBoundary = v
	n.MarkNeedsComposite()
	return n
}

// SetLayoutBoundary 显式声明布局边界（约束不变时子树不重排）。
func (n *Node) SetLayoutBoundary(v bool) *Node { n.layoutBoundary = v; return n }

// SetOverflow 设置是否允许子内容逃逸父裁剪。
func (n *Node) SetOverflow(v bool) *Node { n.overflow = v; return n }

// SetScrollOffset 设置滚动偏移（会自动 clamp 到内容范围内）。
//
// 滚动**不触发重排**：只重摆自己的子节点（placeScroll），因为内容尺寸不随
// 偏移改变。子节点的 Layout 会命中约束缓存，所以这里只改写 offset，O(子树)。
// 如果走 MarkNeedsLayout，一次滚轮就会把整页重排一遍——这是滚动掉帧的经典原因。
func (n *Node) SetScrollOffset(v float32) *Node {
	if n.kind != KindScroll {
		return n
	}
	next := clampf(v, 0, n.maxScroll())
	if next == n.scrollOffset {
		return n
	}
	n.scrollOffset = next
	n.placeScroll()
	n.MarkNeedsPaint()
	return n
}

// ScrollOffset 返回滚动偏移。
func (n *Node) ScrollOffset() float32 { return n.scrollOffset }

// ContentSize 返回滚动内容尺寸。
func (n *Node) ScrollContentSize() geom.Size { return n.contentSize }

// maxScroll 返回可滚动距离（滚动轴）。
func (n *Node) maxScroll() float32 {
	var view, content float32
	if n.axis == Horizontal {
		view, content = n.size.W, n.contentSize.W
	} else {
		view, content = n.size.H, n.contentSize.H
	}
	if content <= view {
		return 0
	}
	return content - view
}

// Text 返回文本 spec（非文本节点为 nil）。
func (n *Node) Text() *TextSpec { return n.text }

// SetText 替换文本内容并重排（内容变化 → needsLayout）。
func (n *Node) SetText(spec TextSpec) *Node {
	if n.kind != KindText {
		return n
	}
	s := spec
	n.text = &s
	if s.Layout != nil {
		n.size = s.Layout.Size
	}
	n.MarkNeedsLayout()
	return n
}

// Box 返回 Box spec（非 Box 节点为 nil）。
func (n *Node) Box() *BoxSpec { return n.box }

// SetBox 替换 Box 样式。
func (n *Node) SetBox(spec BoxSpec) *Node {
	if n.kind != KindBox {
		return n
	}
	s := spec
	n.box = &s
	n.MarkNeedsPaint()
	return n
}

// Walk 深度优先遍历子树（draw 序：父先于子，子按 children 顺序）。
// fn 返回 false 时剪掉该子树。
func (n *Node) Walk(fn func(*Node) bool) {
	if n == nil {
		return
	}
	if !fn(n) {
		return
	}
	for _, c := range n.children {
		c.Walk(fn)
	}
}

// WalkHit 按命中序（逆绘制序）遍历：子节点从后往前，父在最后。
func (n *Node) WalkHit(fn func(*Node) bool) {
	if n == nil {
		return
	}
	if !fn(n) {
		return
	}
	for i := len(n.children) - 1; i >= 0; i-- {
		n.children[i].WalkHit(fn)
	}
}

// Validate 检查 Kind 与 paint spec 的一致性，以及显而易见的结构错误。
//
// 它只在测试与 --selftest 里跑：每帧跑一遍会把性能门禁打穿。
// 返回第一个问题，nil 表示没问题。
func (n *Node) Validate() error {
	return n.walkValidate(nil)
}

// nan 是"未设置"的尺寸哨兵。
//
// 用 NaN 而不是 0：0 是合法的固定尺寸（`SetWidth(0)` 有意义），
// 用 0 当哨兵会让"显式设成 0"与"没设"无法区分。
var nan = float32(math.NaN())

// isSet 报告一个尺寸覆盖值是否被设置过。
func isSet(v float32) bool { return !math.IsNaN(float64(v)) }

// FrameInRoot 返回本节点到根空间的变换（祖先链自根向下合成）。
//
// 为什么需要它：浮层定位、Tooltip、拖拽影子都要知道"某个节点在屏幕上的
// 实际位置"，而节点只存相对父的 offset。让每个调用方自己爬父链迟早
// 会漏掉 transform（旋转过的锚点定位会偏）。
func (n *Node) FrameInRoot() geom.Matrix {
	var chain []*Node
	for x := n; x != nil; x = x.Parent() {
		chain = append(chain, x)
	}
	m := geom.Identity
	for i := len(chain) - 1; i >= 0; i-- {
		m = chain[i].FrameMatrix(m)
	}
	return m
}

// AbsOrigin 返回节点左上角在根空间的坐标。
func (n *Node) AbsOrigin() geom.Point {
	return n.FrameInRoot().MapPoint(geom.Point{})
}

// BoundsInRoot 返回节点在根空间的轴对齐外接矩形。
//
// 旋转/缩放过的节点返回外接框（不是最小包围盒）：浮层定位要的是
// "能对齐的那个矩形"，紧贴旋转后的斜四边形既难算也没有视觉收益。
func (n *Node) BoundsInRoot() geom.Rect {
	return n.FrameInRoot().MapRect(geom.RectXYWH(0, 0, n.size.W, n.size.H))
}

// VisibleInTree 表示节点自身与所有祖先都可见。
func (n *Node) VisibleInTree() bool {
	for x := n; x != nil; x = x.Parent() {
		if !x.Visible() {
			return false
		}
	}
	return true
}
