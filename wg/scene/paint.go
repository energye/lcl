package scene

import (
	"math"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
)

// Record 把整棵子树录进 Canvas（只录制，不碰像素）。
//
// 坐标约定：子节点的 offset 是相对**父节点原点**的（已含父的 padding），
// 因此这里只叠加自身 offset 与 transform，不需要再平移 padding。
// 这条约定必须与 layout.go 的 place* 保持一致，否则整棵树会整体偏移一个
// 内边距——`TestLayoutPadding` 守着它。
//
// 变换以**节点中心**为支点（旋转/缩放符合直觉）；要改支点用 SetTransformOrigin。
func (n *Node) Record(c render.Canvas) {
	if !n.PaintFrame(c) {
		return
	}
	n.PaintSelf(c)
	for _, ch := range n.children {
		ch.Record(c)
	}
	n.EndFrame(c)
}

// PaintFrame 应用节点的边帧（offset、以支点为中心的 transform、clip）。
//
// 返回 false 表示该子树整体不可见，调用方必须**跳过整棵子树**且不要调
// EndFrame。这是给外部合成器（composite 的层树构建）准备的入口：它需要
// 自己递归子节点，以便在遇到层边界时把录制流切开，而 Record 会一口气
// 录完整个子树、没有切分点。
func (n *Node) PaintFrame(c render.Canvas) bool {
	if c == nil || !n.visible || n.opacity <= 0 {
		return false
	}
	c.Save()
	if n.offset.X != 0 || n.offset.Y != 0 {
		c.Translate(n.offset.X, n.offset.Y)
	}
	if !n.transform.IsIdentity() {
		px, py := n.transformOrigin()
		c.Translate(px, py)
		c.Concat(n.transform)
		c.Translate(-px, -py)
	}
	if n.clip != nil {
		n.recordClip(c)
	}
	return true
}

// EndFrame 与 PaintFrame 配对出栈。
func (n *Node) EndFrame(c render.Canvas) {
	if c != nil {
		c.Restore()
	}
}

// PaintSelf 只画节点自身的视觉，**不**递归子节点、不应用边帧。
func (n *Node) PaintSelf(c render.Canvas) {
	if c == nil || !n.visible {
		return
	}
	n.recordSelf(c)
}

// FrameMatrix 返回 parent ∘ 节点边帧 的矩阵（offset + 以支点为中心的 transform）。
//
// 与 PaintFrame 是同一份数学，只是以矩阵形式返回：合成层需要"先知道
// 节点在设备空间的落点，再决定离屏缓冲开多大"，而 Canvas 的 Save/Translate
// 只能顺序施加、拿不到中间结果。
func (n *Node) FrameMatrix(parent geom.Matrix) geom.Matrix {
	m := parent.Concat(geom.Translate(n.offset.X, n.offset.Y))
	if n.transform.IsIdentity() {
		return m
	}
	px, py := n.transformOrigin()
	return m.Concat(geom.Translate(px, py)).
		Concat(n.transform).
		Concat(geom.Translate(-px, -py))
}

// 组透明度（opacity<1）**不在这里处理**：把子树逐指令乘 alpha 会在
// 重叠处重复变暗，是错的。正确做法是离屏合成——composite 会为 opacity<1
// 的节点升级成 OpacityLayer（ReasonOpacity），由它整体混合。
// 因此直接调 Record 得到的 DisplayList 是"不透明版本"，这是刻意的：
// Record 只负责指令，层语义属于合成阶段。

// transformOrigin 返回变换支点（默认节点中心，可用 SetTransformOrigin 改）。
func (n *Node) transformOrigin() (float32, float32) {
	a := n.transformOriginAlign()
	p := a.AlignChild(n.size.W, n.size.H, 0, 0)
	return p.X, p.Y
}

// recordClip 应用裁剪。矩形走 scissor，圆角/路径走 ClipPath。
//
// M2 的 ClipPath 在光栅层降级为包围盒（并计入 degraded）；
// 需要真正圆角裁剪时由 composite 把它升级成 ClipLayer 走蒙版合成。
func (n *Node) recordClip(c render.Canvas) {
	s := n.clip
	r := s.Rect
	if r.IsEmpty() {
		r = geom.RectXYWH(0, 0, n.size.W, n.size.H)
	}
	switch {
	case s.Path != nil:
		c.ClipPath(s.Path)
	case s.Radius > 0:
		c.ClipPath(render.NewPath().AddRRect(r, s.Radius))
	default:
		c.ClipRect(r)
	}
}

// recordSelf 画节点自身的视觉（不含子树）。
func (n *Node) recordSelf(c render.Canvas) {
	switch n.kind {
	case KindBox, KindStack, KindFlex:
		// 容器也画自己的 box：表格行斑马纹/选中态、树节点选中、下拉项
		// 高亮、弹层面板底色都写在 Flex/Stack 的 SetBox 上。以前只认
		// KindBox，这些填充从未上屏（--audit 的点击零像素变化暴露）。
		n.recordBox(c)
	case KindText:
		n.recordText(c)
	case KindPath:
		n.recordPath(c)
	case KindImage:
		n.recordImage(c)
	case KindCustomPaint:
		if n.custom != nil {
			n.custom(c, n.size)
		}
	}
}

func (n *Node) recordBox(c render.Canvas) {
	b := n.box
	if b == nil {
		return
	}
	r := geom.RectXYWH(0, 0, n.size.W, n.size.H)

	// 阴影：Blur<=0 走硬阴影（偏移 + 半透明，零成本）；
	// Blur>0 交给 composite 的滤镜层做离屏卷积——这里**不能**同时画，
	// 否则硬影和软影会叠成两层。
	if b.Shadow != nil && b.Shadow.Color.A > 0 && b.Shadow.Blur <= 0 {
		c.FillRRect(r.Offset(b.Shadow.Dx, b.Shadow.Dy), b.Radius, render.Fill(b.Shadow.Color))
	}
	if b.Fill.Shader != nil || !b.Fill.Color.IsTransparent() {
		if b.Radius > 0 {
			c.FillRRect(r, b.Radius, b.Fill)
		} else {
			c.FillRect(r, b.Fill)
		}
	}
	if b.Border.StrokeWidth > 0 && !b.Border.Color.IsTransparent() &&
		b.Border.Color.A > 0 && b.Fill.Color != b.Border.Color {
		// CSS border-box 语义：边框画在盒**内**。描边路径内缩半线宽——
		// 否则描边骑在盒边上，内外各半：内侧墨量只剩 50%（行色=Mix(边框，
		// 底色)），边框颜色断言不可靠、视觉上边框发虚且向外溢出半像素。
		bw := b.Border.StrokeWidth
		inset := bw * 0.5
		r2 := geom.RectXYWH(inset, inset, n.size.W-bw, n.size.H-bw)
		// 设备像素对齐（CSS border 渲染语义：1px 就是 1px 实色）：布局是
		// 分数坐标（行高 1.5×14=21、gap 8 累加出 .2/.6 尾数），不取整时
		// 1px 描边跨两个像素行各 AA 一半——浅色边框（#D9D9D9）还能混过
		// tol，饱和色（主色/红/黄）墨量被白底稀释到断言必失败。
		// 描边以路径为中心覆盖 [p-0.5, p+0.5]：路径对齐到**半整数**设备
		// 像素，描边才正好铺满一个像素行（整数路径仍会跨行各半）。
		// 仅在纯平移矩阵（scale=1）下取整；带缩放的 HiDPI 路径不动。
		if rc, ok := c.(*render.Recorder); ok {
			if m := rc.Matrix(); m.A == 1 && m.B == 0 && m.C == 0 && m.D == 1 {
				snap := func(v, off float32) float32 {
					return float32(math.Floor(float64(v+off)+0.5)) + 0.5 - off
				}
				r2 = geom.RectXYWH(
					snap(r2.Min.X, m.E), snap(r2.Min.Y, m.F),
					r2.W(), r2.H())
			}
		}
		if b.Radius > 0 {
			radius := b.Radius - inset
			if radius < 0 {
				radius = 0
			}
			c.StrokePath(render.NewPath().AddRRect(r2, radius), b.Border)
		} else {
			c.StrokePath(render.NewPath().AddRect(r2), b.Border)
		}
	}
}

func (n *Node) recordText(c render.Canvas) {
	t := n.text
	if t == nil || t.Layout == nil {
		return
	}
	c.DrawText(t.Layout, geom.Pt(0, 0), t.Color)
}

func (n *Node) recordPath(c render.Canvas) {
	p := n.path
	if p == nil || p.Path == nil {
		return
	}
	b := p.Path.Bounds()
	// 路径坐标是"路径自己的"空间；节点尺寸由 bounds（含半线宽）得到，
	// 因此把路径左上角对齐到节点左上角，描边自动居中在边界内。
	var pad float32
	if p.Paint.StrokeWidth > 0 {
		pad = p.Paint.StrokeWidth * 0.5
	}
	c.Save()
	c.Translate(pad-b.Min.X, pad-b.Min.Y)
	c.DrawPath(p.Path, p.Paint)
	c.Restore()
}

func (n *Node) recordImage(c render.Canvas) {
	im := n.image
	if im == nil || im.Image == nil {
		return
	}
	c.DrawImage(im.Image, n.imageFitRect(), render.ImageBilinear)
}

// imageFitRect 按 Fit 计算位图在节点内的目标矩形。
func (n *Node) imageFitRect() geom.Rect {
	dst := geom.RectXYWH(0, 0, n.size.W, n.size.H)
	im := n.image
	if im == nil || im.Image == nil || im.Fit == FitFill || im.Fit == FitNone {
		return dst
	}
	srcW, srcH := float32(im.Image.W), float32(im.Image.H)
	if srcW <= 0 || srcH <= 0 {
		return dst
	}
	sx := dst.W() / srcW
	sy := dst.H() / srcH
	s := sx
	if im.Fit == FitContain {
		s = minf(sx, sy)
	} else { // FitCover
		s = maxf(sx, sy)
	}
	w, h := srcW*s, srcH*s
	return geom.RectXYWH(
		dst.Min.X+(dst.W()-w)*0.5,
		dst.Min.Y+(dst.H()-h)*0.5,
		w, h,
	)
}

// LocalHit 判定点是否命中本节点的形状（纯几何，不含遍历与裁剪栈）。
//
// 形状与绘制用**同一份几何**：圆角矩形按圆角判、路径按路径判，
// 避免"点方角也响应"这类不一致。
func (n *Node) LocalHit(p geom.Point) bool {
	if n.size.W <= 0 || n.size.H <= 0 {
		// 零尺寸节点（间距、热区）只在显式声明为矩形命中时按 0 尺寸处理
		return n.hit.Shape == ShapeRect && p == geom.Pt(0, 0)
	}
	switch n.hit.Shape {
	case ShapeRRect:
		return pointInRRect(p, n.size, n.hit.Radius)
	case ShapeCircle:
		cx, cy := n.size.W*0.5, n.size.H*0.5
		rx, ry := cx, cy
		if rx <= 0 || ry <= 0 {
			return false
		}
		dx, dy := (p.X-cx)/rx, (p.Y-cy)/ry
		return dx*dx+dy*dy <= 1
	case ShapePath:
		if n.hit.Path != nil {
			return hitPath(n.hit.Path, p)
		}
		return geom.RectXYWH(0, 0, n.size.W, n.size.H).Contains(p)
	default:
		return geom.RectXYWH(0, 0, n.size.W, n.size.H).Contains(p)
	}
}

// PointInRRect 判定点是否落在 0..size 的圆角矩形内。
//
// 导出是为了让合成层的离屏蒙版与命中测试共用**同一份**圆角几何：
// 蒙版自己算一遍圆角，迟早会出现"点不进的地方画出来了"。
func PointInRRect(p geom.Point, s geom.Size, radius float32) bool {
	return pointInRRect(p, s, radius)
}

// pointInRRect 判定点是否落在圆角矩形内（四角用圆判定）。
func pointInRRect(p geom.Point, s geom.Size, radius float32) bool {
	r := geom.RectXYWH(0, 0, s.W, s.H)
	if !r.Contains(p) {
		return false
	}
	radius = minf(radius, minf(s.W, s.H)*0.5)
	if radius <= 0 {
		return true
	}
	cx := clampf(p.X, radius, s.W-radius)
	cy := clampf(p.Y, radius, s.H-radius)
	dx, dy := p.X-cx, p.Y-cy
	return dx*dx+dy*dy <= radius*radius
}

// hitPath 用路径包围盒做近似命中（精确的点在多边形内判定留给命中缓存优化，
// M2 先用包围盒 + 调用方显式声明 ShapePath 的场景已知形状）。
func hitPath(path *render.Path, p geom.Point) bool {
	return path.Bounds().Contains(p)
}

// RecordTree 录制整棵树，返回一个 DisplayList。
//
// 这是"没有合成层"时的扁平化路径：整棵树进一个列表。
// 有 Reason 的子树由 composite 单独录制（见 composite/tree.go）。
func (n *Node) RecordTree(buf *render.DisplayList) *render.DisplayList {
	rec := render.NewRecorder(buf)
	n.Record(rec)
	return rec.Finish()
}
