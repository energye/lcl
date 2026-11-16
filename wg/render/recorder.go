package render

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/text"
)

// recState 是 Save/Restore 保存的状态。
type recState struct {
	matrix    geom.Matrix
	clip      geom.Rect
	clipValid bool
}

// Recorder 把 Canvas 调用录制成 DisplayList。
//
// 三个关键决定：
//  1. **变换在录制期就累积进指令**，光栅期没有矩阵栈。好处是单条指令可独立
//     重放（缓存、命中、局部重绘都要这个），代价是 Save/Restore 只影响后续录制。
//  2. **路径保持局部坐标**，随指令带绝对矩阵。这样曲线扁平化发生在光栅期，
//     精度按最终缩放到设备空间后决定，不会因为提前变换而在高 DPI 下失真。
//  3. **轴对齐矩形走 FillRect 快速路径**（设备空间矩形），旋转/斜切时自动
//     退化为路径填充——而不是用包围盒近似（那会画错形状）。
type Recorder struct {
	dl *DisplayList

	matrix    geom.Matrix
	stack     []recState
	clip      geom.Rect
	clipValid bool

	// scratch 复用一条路径，避免 FillOval/DrawArc 每次分配。
	scratch *Path
	// tolerance 是扁平化容差（DIP）。
	tolerance float32
}

// NewRecorder 创建录制器；dl 为 nil 时内部新建。
func NewRecorder(dl *DisplayList) *Recorder {
	if dl == nil {
		dl = NewDisplayList(64)
	}
	r := &Recorder{
		dl:        dl,
		matrix:    geom.Identity,
		tolerance: 0.25,
		scratch:   NewPath(),
	}
	return r
}

// DisplayList 返回正在录制的指令表。
func (r *Recorder) DisplayList() *DisplayList { return r.dl }

// Finish 冻结并返回指令表。
func (r *Recorder) Finish() *DisplayList {
	// 未配对的 Save 要在末尾补齐时不能改指令（会破坏已录制内容语义），
	// 这里只清空栈：多余的 Save 不会影响后续重放（裁剪在重放期自会展开）。
	r.stack = r.stack[:0]
	r.matrix = geom.Identity
	r.clip, r.clipValid = geom.ZeroRect, false
	return r.dl.Freeze()
}

// Reset 复用指令表继续录制（version 递增）。
func (r *Recorder) Reset() {
	r.dl.Reset()
	r.stack = r.stack[:0]
	r.matrix = geom.Identity
	r.clip, r.clipValid = geom.ZeroRect, false
}

// NewRecorderAt 创建录制器并把初始状态设为 (m, clip)，**不产生任何指令**。
//
// 为什么需要它：层树构建时要把一条录制流在层边界处切成两条，新的那条
// 必须从分叉点的矩阵与裁剪继续，否则层内指令整体偏移。用 Concat 设置
// 矩阵会多出一条 OpConcat（虽然当前后端忽略它，但那是实现细节，
// 不能拿来兜底）。
func NewRecorderAt(dl *DisplayList, m geom.Matrix, clip geom.Rect, hasClip bool) *Recorder {
	r := NewRecorder(dl)
	r.matrix = m
	r.clip, r.clipValid = clip, hasClip
	return r
}

// Fork 以当前状态（矩阵、裁剪、Save 深度）开一条新的录制流。
//
// 复制 Save 深度是关键：分叉点常常位于若干层祖先的 Save 之内，
// 缺了栈之后第一个 Restore 就会变成空操作，后续兄弟节点会带着
// 错误矩阵继续录。
func (r *Recorder) Fork(dl *DisplayList) *Recorder {
	nr := NewRecorderAt(dl, r.matrix, r.clip, r.clipValid)
	if len(r.stack) > 0 {
		nr.stack = append(make([]recState, 0, len(r.stack)), r.stack...)
	}
	return nr
}

// Restart 交出已录内容并**原地**从同一状态继续录制，返回冻结后的指令表。
//
// 与"Finish 再新建"的区别只有一点，但那一点是致命的：Finish 会清空
// 矩阵/裁剪/Save 栈，之后新建的录制器等于从零开始，Restore 就配不上对。
// 层树切分录制流时（父层 picture 关闭、插入子层、再开新 picture）
// 走的就是这条路径，所以需要"交出内容但保留状态"这一原子操作。
//
// 当前指令表为空时返回 (nil, false)，调用方据此跳过空 picture 层。
func (r *Recorder) Restart() (*DisplayList, bool) {
	if r.dl.Len() == 0 {
		return nil, false
	}
	m, clip, hasClip := r.matrix, r.clip, r.clipValid
	stack := r.stack[:len(r.stack):len(r.stack)]
	frozen := r.dl.Freeze()
	r.dl = NewDisplayList(cap(frozen.cmds))
	r.matrix, r.clip, r.clipValid = m, clip, hasClip
	r.stack = stack
	return frozen, true
}

// Rebase 把当前矩阵的平移分量并进裁剪后清零，返回被清零的**设备原点**。
//
// 这是"层局部坐标"的开关：进入一个合成层时调用，之后录进该层的指令
// 原点是层自己的 (0,0)。带来两个直接好处：
//   - 层内容签名与"层在屏幕上的位置"解耦，滚动/拖拽整层移动时缓存照样命中；
//   - 离屏缓冲只需层自身尺寸，不需要把屏幕坐标搬来搬去。
//
// 旋转/缩放分量保留：层空间是"节点自己的空间"，斜着的内容进去还是斜的。
func (r *Recorder) Rebase() geom.Point {
	origin := geom.Pt(r.matrix.E, r.matrix.F)
	r.matrix.E, r.matrix.F = 0, 0
	if r.clipValid {
		r.clip = r.clip.Offset(-origin.X, -origin.Y)
	}
	return origin
}

// Tolerance 设置扁平化容差（DIP）。
func (r *Recorder) Tolerance(v float32) { r.tolerance = v }

// Matrix 返回当前累积矩阵。
func (r *Recorder) Matrix() geom.Matrix { return r.matrix }

// Clip 返回当前设备空间裁剪；第二返回值为 false 表示无裁剪。
func (r *Recorder) Clip() (geom.Rect, bool) { return r.clip, r.clipValid }

// ── Canvas 实现 ─────────────────────────────────────────────────────

// Save 压栈。
func (r *Recorder) Save() {
	r.stack = append(r.stack, recState{matrix: r.matrix, clip: r.clip, clipValid: r.clipValid})
	r.dl.append(Command{Op: OpSave, Matrix: r.matrix, Path: -1, Paint: -1, Text: -1, Image: -1})
}

// Restore 出栈。空栈时忽略，避免一个组件的多余 Restore 破坏整帧。
func (r *Recorder) Restore() {
	if len(r.stack) == 0 {
		return
	}
	st := r.stack[len(r.stack)-1]
	r.stack = r.stack[:len(r.stack)-1]
	r.matrix = st.matrix
	r.clip, r.clipValid = st.clip, st.clipValid
	r.dl.append(Command{Op: OpRestore, Matrix: r.matrix, Path: -1, Paint: -1, Text: -1, Image: -1})
}

// Translate 平移。
func (r *Recorder) Translate(dx, dy float32) { r.Concat(geom.Translate(dx, dy)) }

// Scale 缩放。
func (r *Recorder) Scale(sx, sy float32) { r.Concat(geom.Scale(sx, sy)) }

// Rotate 旋转。
func (r *Recorder) Rotate(rad float32) { r.Concat(geom.Rotate(rad)) }

// Concat 左乘变换。
func (r *Recorder) Concat(m geom.Matrix) {
	r.matrix = r.matrix.Concat(m)
	r.dl.append(Command{Op: OpConcat, Matrix: r.matrix, Path: -1, Paint: -1, Text: -1, Image: -1})
}

// ClipRect 与当前裁剪求交。
func (r *Recorder) ClipRect(rc geom.Rect) {
	dev := r.matrix.MapRect(rc)
	if r.clipValid {
		dev = dev.Intersect(r.clip)
	}
	r.clip, r.clipValid = dev, true
	r.dl.append(Command{Op: OpClipRect, Rect: dev, Matrix: r.matrix, Path: -1, Paint: -1, Text: -1, Image: -1})
}

// ClipPath 按路径裁剪。
//
// 指令同时携带路径与设备空间包围盒：包围盒先收 scissor（省像素），
// 路径由后端做精确蒙版。早期只记包围盒并计 degraded，异形裁剪会画出去。
func (r *Recorder) ClipPath(p *Path) {
	if p == nil || p.IsEmpty() {
		return
	}
	dev := r.matrix.MapRect(p.Bounds())
	if r.clipValid {
		dev = dev.Intersect(r.clip)
	}
	r.clip, r.clipValid = dev, true
	r.dl.paths = append(r.dl.paths, p)
	r.dl.append(Command{Op: OpClipPath, Rect: dev, Matrix: r.matrix,
		Path: int32(len(r.dl.paths) - 1), Paint: -1, Text: -1, Image: -1})
}

// FillRect 填充矩形。旋转/斜切时退化为路径填充，保证形状正确。
func (r *Recorder) FillRect(rc geom.Rect, paint Paint) {
	if rc.IsEmpty() {
		return
	}
	if r.matrix.IsAxisAligned() {
		dev := r.matrix.MapRect(rc)
		if !r.visible(dev) {
			return
		}
		r.dl.append(Command{Op: OpFillRect, Rect: dev, Matrix: r.matrix, Paint: r.addPaint(paint), Path: -1, Text: -1, Image: -1})
		return
	}
	r.scratch.Reset().AddRect(rc)
	r.FillPath(r.scratch, paint)
}

// FillRRect 填充圆角矩形。
func (r *Recorder) FillRRect(rc geom.Rect, radius float32, paint Paint) {
	if rc.IsEmpty() {
		return
	}
	r.scratch.Reset().AddRRect(rc, radius)
	r.FillPath(r.scratch, paint)
}

// FillOval 填充椭圆。
func (r *Recorder) FillOval(rc geom.Rect, paint Paint) {
	if rc.IsEmpty() {
		return
	}
	r.scratch.Reset().AddOval(rc)
	r.FillPath(r.scratch, paint)
}

// DrawLine 画线。线条按描边语义处理，忽略 paint.Style。
func (r *Recorder) DrawLine(x0, y0, x1, y1 float32, paint Paint) {
	if x0 == x1 && y0 == y1 {
		return
	}
	r.scratch.Reset().MoveTo(x0, y0).LineTo(x1, y1)
	paint.Style = StyleStroke
	r.StrokePath(r.scratch, paint)
}

// DrawArc 画圆弧。填充时闭合为弓形。
func (r *Recorder) DrawArc(rc geom.Rect, start, sweep float32, paint Paint) {
	if rc.IsEmpty() || sweep == 0 {
		return
	}
	r.scratch.Reset().AddArc(rc, start, sweep)
	if paint.Style == StyleFill {
		r.scratch.Close()
		r.FillPath(r.scratch, paint)
		return
	}
	r.StrokePath(r.scratch, paint)
}

// FillPath 填充路径。
func (r *Recorder) FillPath(p *Path, paint Paint) {
	if p == nil || p.IsEmpty() {
		return
	}
	dev := r.matrix.MapRect(p.Bounds())
	if !r.visible(dev) {
		return
	}
	r.dl.paths = append(r.dl.paths, p.Clone())
	r.dl.append(Command{
		Op: OpFillPath, Rect: dev, Matrix: r.matrix,
		Path: int32(len(r.dl.paths) - 1), Paint: r.addPaint(paint), Text: -1, Image: -1,
	})
}

// StrokePath 描边路径。
func (r *Recorder) StrokePath(p *Path, paint Paint) {
	if p == nil || p.IsEmpty() {
		return
	}
	w := paint.StrokeWidthOr()
	// 描边包围盒 = 路径包围盒外扩半个线宽（斜接时略小，保守即可）。
	dev := r.matrix.MapRect(p.Bounds()).Expand(w * 0.5 * r.matrix.ScaleFactor())
	if !r.visible(dev) {
		return
	}
	paint.Style = StyleStroke
	r.dl.paths = append(r.dl.paths, p.Clone())
	r.dl.append(Command{
		Op: OpStrokePath, Rect: dev, Matrix: r.matrix,
		Path: int32(len(r.dl.paths) - 1), Paint: r.addPaint(paint), Text: -1, Image: -1,
	})
}

// DrawPath 按 Style 分派。
func (r *Recorder) DrawPath(p *Path, paint Paint) {
	if paint.Style == StyleStroke {
		r.StrokePath(p, paint)
		return
	}
	r.FillPath(p, paint)
}

// DrawText 记录文字。origin 为用户空间左上角，此处立即换算为设备空间（C24）。
func (r *Recorder) DrawText(l *text.Layout, origin geom.Point, color geom.Color) {
	if l.Empty() || color.IsTransparent() {
		return
	}
	dev := r.matrix.MapPoint(origin)
	size := l.Size
	rect := geom.RectXYWH(dev.X, dev.Y, size.W*r.matrix.ScaleFactor(), size.H*r.matrix.ScaleFactor())
	if r.clipValid {
		// 文字命令的 Rect 同时是"文字通道的裁剪矩形"：文字不在光栅后端回放，
		// 而是由 GDI 在像素写回后按这条矩形裁剪（见 03-ENGINE.md §9）。
		// 不在这里与录制期裁剪求交，ClipRect 之后的文字就会溢出裁剪区。
		rect = rect.Intersect(r.clip)
	}
	if !r.visible(rect) {
		return
	}
	r.dl.texts = append(r.dl.texts, TextItem{Layout: l, Origin: dev, Color: color})
	r.dl.append(Command{
		Op: OpDrawText, Rect: rect, Matrix: r.matrix,
		Path: -1, Paint: -1, Text: int32(len(r.dl.texts) - 1), Image: -1,
	})
	r.dl.hasText = true
}

// DrawImage 记录位图指令（M1 后端只计数，不绘制）。
func (r *Recorder) DrawImage(img *ImageRef, dst geom.Rect, q ImageQuality) {
	if img == nil || dst.IsEmpty() {
		return
	}
	dev := r.matrix.MapRect(dst)
	if !r.visible(dev) {
		return
	}
	r.dl.images = append(r.dl.images, img)
	r.dl.append(Command{
		Op: OpDrawImage, Rect: dev, Matrix: r.matrix, Quality: q,
		Path: -1, Paint: -1, Text: -1, Image: int32(len(r.dl.images) - 1),
	})
}

// ── 内部 ────────────────────────────────────────────────────────────

// addPaint 把画笔放进 side table，返回索引。
//
// 关闭 AA（p.AntiAlias=false）是 A/B 对照用的显式开关，不算能力缺失，
// 因此不计入 degraded——degraded 只统计"这台后端做不到"的指令（C22）。
func (r *Recorder) addPaint(p Paint) int32 {
	r.dl.paints = append(r.dl.paints, p)
	return int32(len(r.dl.paints) - 1)
}

// visible 判定设备矩形是否与当前裁剪相交。
func (r *Recorder) visible(dev geom.Rect) bool {
	if dev.IsEmpty() {
		return false
	}
	if !r.clipValid {
		return true
	}
	return dev.Intersects(r.clip)
}
