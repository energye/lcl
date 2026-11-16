// Package composite 的层树构建。
//
// 一帧的完整路径是：
//
//	scene.Node 树 --Build--> Layer 树 --Compositor.DrawTree--> 像素
//
// Build 只做两件事：决定"哪里要开层"，以及把指令录进正确的层。
// 默认路径是**就地扁平化**：没有理由的节点不产生层，指令直接追加进
// 当前 picture。于是"层数 ≈ 特效数"而不是"节点数"——1000 个节点、
// 0 个特效时整棵树只有 1 个根层（`--layers` 断言这一条）。
package composite

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
)

// BuildOptions 是层树构建选项。
type BuildOptions struct {
	// Matrix 是设备变换（DPI 缩放等），零值表示单位矩阵。
	// 它只影响初始矩阵，不影响 DIP 布局——布局在 scene 阶段已经算完。
	Matrix geom.Matrix

	// Clip 是初始裁剪（设备空间）；零值表示不裁剪。
	Clip geom.Rect

	// MaxDepth 是层嵌套上限（0 = 默认 8）。
	//
	// 为什么要上限：层嵌套意味着离屏缓冲嵌套，内存是指数级的风险
	// （每层一块，最坏情况面积之和 = 层数 × 全屏）。超过上限后
	// 后续层降级为扁平化并计数，画面仍然正确，只是少一次隔离。
	MaxDepth int

	// Cache 控制 RepaintBoundary 是否升级为"离屏 + 可缓存"。
	//
	// 关掉它有两个用途：A/B 对照（验证层管线不改变像素），
	// 以及显存/内存紧张时退让。opacity/clip 层不受此开关影响——
	// 它们不走向量离屏就是错的，不是慢的问题。
	Cache bool
}

// 默认层嵌套上限。
const defaultMaxDepth = 8

// treeStats 是构建统计。
type treeStats struct {
	Nodes      int
	Groups     int
	Pictures   int
	Opacity    int
	Clip       int
	Boundary   int
	Filter     int
	Surface    int
	MaxDepth   int
	Downgraded int // 有理由但 M2 不建层的次数（Filter/Surface/超深度）
}

// Tree 是一棵构建好的合成层树。
//
// 生命周期：一次 Build 产出一个 Tree，Tree 在**内容不变期间被反复合成**。
// 这就是缓存能命中的前提：picture 的 version 随重录递增，
// 树被重建 = 所有 version 换号 = 缓存全冷。
// 帧调度器（frame/scheduler.go）因此只在 needsPaint 帧调 Build。
type Tree struct {
	root  *Layer
	stats treeStats
	frame int
}

// Root 返回根层（永不为 nil，除非 Build 传入 nil 根节点）。
func (t *Tree) Root() *Layer {
	if t == nil {
		return nil
	}
	return t.root
}

// Layers 返回总层数（含 picture 切片）。
func (t *Tree) Layers() int {
	if t == nil || t.root == nil {
		return 0
	}
	return t.root.Count()
}

// Groups 返回组层数（真正的"合并单元"）。
func (t *Tree) Groups() int { return t.stats.Groups }

// Pictures 返回指令切片数。
func (t *Tree) Pictures() int { return t.stats.Pictures }

// Stats 返回构建统计。
func (t *Tree) Stats() treeStats {
	if t == nil {
		return treeStats{}
	}
	return t.stats
}

// Ring 返回层数直方图（按 Reason），诊断输出用。
func (t *Tree) Ring() map[Reason]int {
	out := map[Reason]int{}
	if t == nil {
		return out
	}
	t.Walk(func(l *Layer) bool {
		out[l.reason]++
		return true
	})
	return out
}

// Walk 前序遍历所有层。
func (t *Tree) Walk(fn func(*Layer) bool) {
	if t == nil || t.root == nil {
		return
	}
	t.root.Walk(fn)
}

// Frame 返回构建时的帧序号。
func (t *Tree) Frame() int { return t.frame }

// SetFrame 记录构建时的帧序号（由帧调度器填写）。
func (t *Tree) SetFrame(f int) { t.frame = f }

// Validate 检查层树的一致性，返回第一条违规描述（nil 表示通过）。
//
// 校验的三条铁律：
//  1. 除根层外每层都必须有理由（C19）；
//  2. 只有组层能有子层，picture 层必须是叶子；
//  3. 需要离屏的组层必须有非空包围盒（否则缓存键里的尺寸为 0，永远不命中）。
func (t *Tree) Validate() error {
	if t == nil || t.root == nil {
		return nil
	}
	return t.root.validate()
}

func (l *Layer) validate() error {
	if l.kind == LayerPicture && len(l.children) > 0 {
		return fmt.Errorf("wg/composite: picture 层不能有子层（%s）", l.reason)
	}
	if l.kind == LayerGroup && l.reason == ReasonNone {
		return fmt.Errorf("wg/composite: 组层未声明理由（违反 C19）")
	}
	if l.offscreen && l.kind == LayerGroup && l.bounds.IsEmpty() {
		return fmt.Errorf("wg/composite: 离屏层包围盒为空（%s）", l.reason)
	}
	for _, c := range l.children {
		if err := c.validate(); err != nil {
			return err
		}
	}
	return nil
}

// ── 构建 ────────────────────────────────────────────────────────────

// builder 是构建期状态。
type builder struct {
	opts BuildOptions
	t    *Tree
	ctx  []*bctx
}

// bctx 是一条录制流上下文：一个 recorder + 层落点 + 深度。
type bctx struct {
	rec   *render.Recorder
	out   *[]*Layer
	group *Layer
	depth int
}

// Build 从场景树构建层树。
//
// root 为 nil 时返回一个只含根层的空树（不是 nil），
// 这样合成侧不需要到处判空。
func Build(root *scene.Node, opts BuildOptions) *Tree {
	if opts.MaxDepth <= 0 {
		opts.MaxDepth = defaultMaxDepth
	}
	m := opts.Matrix
	if m == (geom.Matrix{}) {
		m = geom.Identity
	}
	t := &Tree{}
	top := &Layer{kind: LayerGroup, reason: ReasonRoot, opacity: 1, cacheable: false}
	t.root = top

	rec := render.NewRecorderAt(render.NewDisplayList(64), m, opts.Clip, !opts.Clip.IsEmpty())
	b := &builder{opts: opts, t: t}
	b.push(&bctx{rec: rec, out: &top.children, group: top, depth: 0})

	if root != nil {
		b.walk(root)
	}
	b.closePicture()
	b.pop()

	top.computeBounds()
	top.setAbsOrigin()

	// 统计
	top.Walk(func(l *Layer) bool {
		if l.kind == LayerGroup {
			t.stats.Groups++
			switch l.reason {
			case ReasonOpacity:
				t.stats.Opacity++
			case ReasonClip:
				t.stats.Clip++
			case ReasonRepaintBoundary:
				t.stats.Boundary++
			case ReasonFilter:
				t.stats.Filter++
			case ReasonSurface:
				t.stats.Surface++
			}
		} else {
			t.stats.Pictures++
		}
		if l.depth > t.stats.MaxDepth {
			t.stats.MaxDepth = l.depth
		}
		return true
	})
	return t
}

func (b *builder) cur() *bctx { return b.ctx[len(b.ctx)-1] }

func (b *builder) push(c *bctx) { b.ctx = append(b.ctx, c) }

func (b *builder) pop() {
	if n := len(b.ctx); n > 0 {
		b.ctx = b.ctx[:n-1]
	}
}

func (b *builder) depth() int { return b.cur().depth }

// walk 递归处理一个节点，返回该子树的**内容签名**。
//
// 签名在构建时顺手累加，不额外遍历：缓存有效性必须能回答
// "这层内容自上次栅格化以来变过没有"，而 scene 的 signature 是
// "被标脏次数"，两者混合就是既保守又便宜的答案。
func (b *builder) walk(n *scene.Node) uint64 {
	if n == nil {
		return 0
	}
	b.t.stats.Nodes++
	h := hashNode(n)
	if !n.Visible() || n.Opacity() <= 0 {
		// 不可见子树不进指令，但仍要参与签名：它变可见时必须失效缓存。
		return h
	}
	if r := b.reasonOf(n); r != ReasonNone && b.depth() < b.opts.MaxDepth {
		return mix64(h, b.boundary(n, r))
	}
	c := b.cur()
	if !n.PaintFrame(c.rec) {
		return h
	}
	n.PaintSelf(c.rec)
	for _, ch := range n.Children() {
		h = mix64(h, b.walk(ch))
	}
	n.EndFrame(c.rec)
	return h
}

// reasonOf 判定节点是否要开层。
//
// 判定顺序就是优先级：一个节点只开**一层**，该层的属性表里叠加它的
// opacity 与 clip；reason 记录的是"第一个触发条件"，用于诊断与断言。
// 这样做的理由：把 opacity 与 clip 拆成两层会让离屏嵌套，
// 而它们作用在同一个内容上，一层就能同时表达，没有拆的必要。
func (b *builder) reasonOf(n *scene.Node) Reason {
	// 组滤镜：模糊/发光/组阴影需要离屏卷积，必须开层。
	if e := n.Effect(); e != nil && (e.BlurSigma > 0 || e.GlowRadius > 0 || e.Shadow != nil) {
		return ReasonFilter
	}
	// 盒阴影带模糊半径时同样要走离屏卷积（硬阴影仍在指令流里直接画）。
	if bs := n.Box(); bs != nil && bs.Shadow != nil && bs.Shadow.Blur > 0 && bs.Shadow.Color.A > 0 {
		return ReasonFilter
	}
	// 子表面：当作隔离组处理（独立离屏 + 独立合成），不再扁平化。
	if n.Kind() == scene.KindSurface {
		return ReasonSurface
	}
	if n.Opacity() < 1 {
		return ReasonOpacity
	}
	if spec := n.Clip(); spec != nil && (spec.Path != nil || spec.Radius > 0) {
		return ReasonClip
	}
	// 非轴对齐变换：见 ReasonTransform 的注释——M2 刻意不建层。
	if n.RepaintBoundary() {
		return ReasonRepaintBoundary
	}
	return ReasonNone
}

// boundary 为一个需要开层的节点建层，并递归它的子树。
//
// 关键点：**不调 PaintFrame**。节点边帧（offset/transform/clip）被算成矩阵
// 直接作为这一层的基准矩阵（scene.FrameMatrix），于是层内指令天生落在
// 层局部空间，不需要"先录进设备空间再平移回去"。
// 这同时避免了另一个坑：PaintFrame 会把 ClipRect 作为指令写进流里，
// 而那条指令携带的是设备坐标，重定基之后它就成了错坐标的裁剪。
func (b *builder) boundary(n *scene.Node, reason Reason) uint64 {
	c := b.cur()
	parentM := c.rec.Matrix()
	parentClip, hasClip := c.rec.Clip()

	// 1) 先关掉父层当前 picture：保证层按绘制顺序插入，而不是全部漂到父层尾部。
	b.closePicture()

	m := n.FrameMatrix(parentM)
	origin := geom.Pt(m.E, m.F)
	local := m
	local.E, local.F = 0, 0

	layer := &Layer{
		kind:      LayerGroup,
		reason:    reason,
		node:      n,
		origin:    origin,
		absOrigin: c.group.absOrigin.Add(origin),
		opacity:   clamp01(n.Opacity()),
		depth:     c.depth + 1,
		nodeSize:  n.Size(),
	}
	if e := n.Effect(); e != nil {
		layer.filter = e
	}
	if bs := n.Box(); bs != nil && bs.Shadow != nil && bs.Shadow.Blur > 0 && bs.Shadow.Color.A > 0 {
		layer.boxShadow = bs.Shadow
		layer.boxRadius = bs.Radius
	}

	// 2) 局部裁剪 = (父裁剪 - 层原点) ∩ 自身裁剪
	clipLocal := parentClip.Offset(-origin.X, -origin.Y)
	clipValid := hasClip
	if spec := n.Clip(); spec != nil {
		r := spec.Rect
		if r.IsEmpty() {
			r = geom.RectXYWH(0, 0, n.Size().W, n.Size().H)
		}
		switch {
		case spec.Path != nil:
			layer.clip = ClipPath
			layer.clipPath = spec.Path
			layer.clipRect = r
		case spec.Radius > 0:
			layer.clip = ClipRRect
			layer.clipRadius = spec.Radius
			layer.clipRect = r
		default:
			layer.clip = ClipRect
			layer.clipRect = r
		}
		if clipValid {
			clipLocal = clipLocal.Intersect(r)
		} else {
			clipLocal, clipValid = r, true
		}
	} else if !clipValid {
		clipLocal = geom.ZeroRect
	}

	inner := render.NewRecorderAt(render.NewDisplayList(32), local, clipLocal, clipValid)
	if clipValid && !clipLocal.IsEmpty() {
		// 矩形裁剪交给 scissor（精确且零成本）；圆角/路径先用包围盒兜一层，
		// 圆角部分由蒙版精确裁——蒙版负责 AA，scissor 负责别去碰外面的像素。
		inner.ClipRect(clipLocal.RoundOut())
	}

	// 3) 层内上下文
	b.push(&bctx{rec: inner, out: &layer.children, group: layer, depth: c.depth + 1})
	n.PaintSelf(inner)
	var h uint64
	for _, ch := range n.Children() {
		h = mix64(h, b.walk(ch))
	}
	b.closePicture()
	b.pop()

	layer.version = mix64(hashNode(n), h)
	layer.computeBounds()

	// 4) 决定合成方式
	//
	// 文字与离屏不再冲突：GDI 文字通道对非不透明目标走黑白双趟
	// 保 alpha 写入（raster/gdi 的 Painter），文字可以正常进离屏，
	// 组透明度/裁剪蒙版/滤镜/缓存对文字全部生效——没有 TextFlat 降级。
	needOff := layer.opacity < 1 || layer.clip != ClipNone ||
		reason == ReasonFilter || reason == ReasonSurface
	if reason == ReasonRepaintBoundary && b.opts.Cache {
		needOff = true
	}
	layer.offscreen = needOff
	layer.cacheable = needOff && b.opts.Cache

	if len(layer.children) == 0 || layer.bounds.IsEmpty() {
		// 空层：没有内容可合成（全被裁剪、或本来就没画）。不留垃圾层。
		b.t.stats.Downgraded++
		return layer.version
	}

	layer.parent = c.group
	*c.out = append(*c.out, layer)
	return layer.version
}

// closePicture 关闭当前录制流并把非空内容落成一个 picture 层。
//
// 走 Recorder.Restart 而不是 Finish+新建：Finish 会清空矩阵/裁剪/Save 栈，
// 重建出来的录制器就"忘了自己在谁的 Save 里面"，后续 Restore 全部失效。
//
// 两处细节都不能省：
//   - 只含状态指令的碎片（一次 Save、一条 scissor）不成层：用 Painted 判定，
//     否则每个层边界都会凭空多出两片空 picture，层数指标立刻失去意义；
//   - 新流必须**重放**当前 scissor：裁剪在旧流里是一条 OpClipRect 指令，
//     新流从空指令表开始，不重放就会漏裁剪（只有嵌套在矩形裁剪里的层才会现形）。
func (b *builder) closePicture() {
	c := b.cur()
	dl, ok := c.rec.Restart()
	if !ok {
		return
	}
	if clip, has := c.rec.Clip(); has && !clip.IsEmpty() {
		c.rec.ClipRect(clip)
	}
	if dl.Painted() == 0 {
		return
	}
	l := &Layer{
		kind:      LayerPicture,
		reason:    ReasonContent,
		parent:    c.group,
		origin:    geom.Point{},
		absOrigin: c.group.absOrigin,
		picture:   dl,
		bounds:    dl.Bounds(),
		depth:     c.depth,
	}
	*c.out = append(*c.out, l)
}

// ── 签名 ────────────────────────────────────────────────────────────

// hashNode 把节点的 (id, signature) 混成一个 64 位数。
func hashNode(n *scene.Node) uint64 {
	return mix64(n.ID()*0x9E3779B97F4A7C15, n.Signature())
}

// mix64 是 splitmix64 的混合函数：把两个数搅成一个，保证低位也参与。
func mix64(a, b uint64) uint64 {
	x := a ^ (b + 0x9E3779B97F4A7C15 + (a << 6) + (a >> 2))
	x ^= x >> 30
	x *= 0xBF58476D1CE4E5B9
	x ^= x >> 27
	x *= 0x94D049BB133111EB
	x ^= x >> 31
	return x
}

// clamp01 把值夹到 [0,1]。
func clamp01(v float32) float32 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}
