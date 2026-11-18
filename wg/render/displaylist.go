package render

import (
	"sync/atomic"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/text"
)

// TextItem 是一条文字指令的载荷。
//
// Origin 是**设备空间**左上角：录制时已把用户空间坐标乘以当时矩阵，
// 因此光栅阶段不再需要矩阵栈（C24：文字在设备空间回放）。
type TextItem struct {
	Layout *text.Layout
	Origin geom.Point
	Color  geom.Color
}

// DisplayList 是不可变的绘制指令序列（契约 C23）。
//
// 不可变的含义是：Frozen 之后任何 Append 都会被忽略（debug 下 panic）。
// 需要改内容就重录并递增 version —— 这让"缓存是否仍有效"可以只看 version，
// 不需要深比较指令。
type DisplayList struct {
	cmds   []Command
	paths  []*Path
	paints []Paint
	texts  []TextItem
	images []*ImageRef

	bounds   geom.Rect
	version  uint32
	frozen   bool
	degraded int
	painted  int
	hasText  bool
	opaque   bool // 所有指令都覆盖整块且有透明？用于快速判断，默认 false
}

// NewDisplayList 创建空指令表。
func NewDisplayList(capacity int) *DisplayList {
	if capacity < 0 {
		capacity = 0
	}
	return &DisplayList{
		cmds:    make([]Command, 0, capacity),
		bounds:  geom.ZeroRect,
		version: nextVersion(),
	}
}

// Commands 返回指令序列。调用方不得修改。
func (d *DisplayList) Commands() []Command { return d.cmds }

// Len 返回指令条数。
func (d *DisplayList) Len() int {
	if d == nil {
		return 0
	}
	return len(d.cmds)
}

// Path 按索引取路径；越界返回 nil。
func (d *DisplayList) Path(i int32) *Path {
	if d == nil || i < 0 || int(i) >= len(d.paths) {
		return nil
	}
	return d.paths[i]
}

// Paint 按索引取画笔。
func (d *DisplayList) Paint(i int32) *Paint {
	if d == nil || i < 0 || int(i) >= len(d.paints) {
		return nil
	}
	return &d.paints[i]
}

// Text 按索引取文字载荷。
func (d *DisplayList) Text(i int32) *TextItem {
	if d == nil || i < 0 || int(i) >= len(d.texts) {
		return nil
	}
	return &d.texts[i]
}

// Image 按索引取图片。
func (d *DisplayList) Image(i int32) *ImageRef {
	if d == nil || i < 0 || int(i) >= len(d.images) {
		return nil
	}
	return d.images[i]
}

// Bounds 返回所有**绘制**指令在设备空间的并集包围盒（裁剪/状态指令
// 不画像素、不占包围盒，理由见 append）。
func (d *DisplayList) Bounds() geom.Rect {
	if d == nil {
		return geom.ZeroRect
	}
	return d.bounds
}

// Version 返回录制版本号（内容变化必然换号）。
func (d *DisplayList) Version() uint32 {
	if d == nil {
		return 0
	}
	return d.version
}

// Degraded 返回录制阶段就被判定为"无法精确执行"的指令数（C22）。
func (d *DisplayList) Degraded() int {
	if d == nil {
		return 0
	}
	return d.degraded
}

// HasText 表示是否含文字指令（决定是否需要设备空间的文字回放阶段）。
func (d *DisplayList) HasText() bool { return d != nil && d.hasText }

// Frozen 表示已冻结。
func (d *DisplayList) Frozen() bool { return d != nil && d.frozen }

// Intersects 判定指令表是否与矩形相交（damage 局部重绘用）。
func (d *DisplayList) Intersects(r geom.Rect) bool {
	if d == nil || r.IsEmpty() {
		return false
	}
	return d.bounds.Intersects(r)
}

// Empty 表示没有任何指令。
func (d *DisplayList) Empty() bool { return d.Len() == 0 }

// Freeze 冻结指令表并返回自身（供链式使用）。
//
// 冻结之后 Recorder 必须 Reset 才能继续录制；若复用同一个 DisplayList
// 的底层切片，version 会继续递增，缓存据此失效（C23）。
func (d *DisplayList) Freeze() *DisplayList {
	if d == nil {
		return d
	}
	d.frozen = true
	return d
}

// Reset 清空指令表以便复用（会递增 version）。
func (d *DisplayList) Reset() {
	if d == nil {
		return
	}
	d.cmds = d.cmds[:0]
	d.paths = d.paths[:0]
	d.paints = d.paints[:0]
	d.texts = d.texts[:0]
	d.images = d.images[:0]
	d.bounds = geom.ZeroRect
	d.frozen = false
	d.degraded = 0
	d.painted = 0
	d.hasText = false
	d.version = nextVersion()
}

// OffsetBy 返回平移后的指令表视图（共享 side table，只复制指令数组）。
//
// 用途只有一个：**离屏合成时的落位**。层的内容录在"层局部坐标"里，
// 画进离屏缓冲时要把它挪到离屏的 (0,0)，画完再整体混到目标缓冲。
// 指令的 Matrix 也要跟着平移——路径几何是 local + Matrix 形式，
// 只改 Rect 会让 OpFillPath 画在原地。
//
// 返回的是视图而不是副本：paths/paints/texts/images 与本体共享，因此
// 调用方不得对它 Freeze/Reset/追加。version 保持不变是刻意的——
// 平移不改变"内容"，缓存键因此不该因为平移而失效（滚动优化依赖这点）。
func (d *DisplayList) OffsetBy(dx, dy float32) *DisplayList {
	if d == nil || (dx == 0 && dy == 0) {
		return d
	}
	out := &DisplayList{
		cmds:     make([]Command, len(d.cmds)),
		paths:    d.paths,
		paints:   d.paints,
		images:   d.images,
		bounds:   d.bounds.Offset(dx, dy),
		version:  d.version,
		frozen:   true,
		degraded: d.degraded,
		painted:  d.painted,
		hasText:  d.hasText,
		opaque:   d.opaque,
	}
	for i, c := range d.cmds {
		c.Rect = c.Rect.Offset(dx, dy)
		c.Matrix.E += dx
		c.Matrix.F += dy
		out.cmds[i] = c
	}
	if len(d.texts) > 0 {
		out.texts = make([]TextItem, len(d.texts))
		copy(out.texts, d.texts)
		for i := range out.texts {
			out.texts[i].Origin = out.texts[i].Origin.Add(geom.Pt(dx, dy))
		}
	}
	return out
}

// ScaleAlpha 返回一条"整体调过透明度"的视图（画笔颜色与渐变色标按 f 缩放）。
//
// 用途只有一个：**组透明度的降级路径**。正确的组透明是离屏整体混合
// （composite 的 OpacityLayer），但含 GDI 文字的层无法离屏（文字通道只写 RGB、
// 不写 alpha，画在透明底上会消失）。这种层退化成"逐指令乘 alpha"：
// 重叠处会重复变暗，但至少不是"要么不变、要么消失"。
//
// 这是一条明面上的近似，调用方必须计入降级统计（composite 的 TextFlat）。
// 与 OffsetBy 一样是视图：paths/images/texts 共享，paints 是副本（要改值）。
func (d *DisplayList) ScaleAlpha(f float32) *DisplayList {
	if d == nil || f >= 1 {
		return d
	}
	if f < 0 {
		f = 0
	}
	out := &DisplayList{
		cmds:     make([]Command, len(d.cmds)),
		paths:    d.paths,
		images:   d.images,
		texts:    d.texts,
		bounds:   d.bounds,
		version:  d.version,
		frozen:   true,
		degraded: d.degraded,
		painted:  d.painted,
		hasText:  d.hasText,
		opaque:   false,
	}
	copy(out.cmds, d.cmds)
	out.paints = make([]Paint, len(d.paints))
	for i, p := range d.paints {
		out.paints[i] = scalePaintAlpha(p, f)
	}
	return out
}

// scalePaintAlpha 返回按 f 缩小 alpha 的画笔副本（含渐变色标）。
func scalePaintAlpha(p Paint, f float32) Paint {
	p.Color = scaleColorAlpha(p.Color, f)
	if p.Shader != nil {
		s := *p.Shader
		s.Stops = make([]Stop, len(p.Shader.Stops))
		for i, st := range p.Shader.Stops {
			s.Stops[i] = Stop{Offset: st.Offset, Color: scaleColorAlpha(st.Color, f)}
		}
		p.Shader = &s
	}
	return p
}

// scaleColorAlpha 按 f 缩放颜色的 alpha 分量（四舍五入）。
func scaleColorAlpha(c geom.Color, f float32) geom.Color {
	a := float32(c.A) * f
	if a <= 0 {
		c.A = 0
		return c
	}
	if a >= 255 {
		c.A = 255
		return c
	}
	c.A = uint8(a + 0.5)
	return c
}

// Painted 返回真正产生像素的指令数（填充/描边/文字/位图）。
//
// 与 Len 的区别：Len 把 Save/Restore/Concat/Clip 这些"状态指令"也算一条。
// 层树切分录制流时会切出一段只含状态指令的碎片（例如一次 Save），
// 它不产生任何像素，不该成为一层——用 Painted 判定，而不是 Len。
func (d *DisplayList) Painted() int {
	if d == nil {
		return 0
	}
	return d.painted
}

// append 追加一条指令并更新包围盒（仅 Recorder 调用）。
//
// bounds 只统计**绘制**指令：裁剪/状态指令不画像素，不占包围盒。
// 这一条是层缓存平移不变性的前提——随滚动平移的视口裁剪窗如果进了
// bounds，层缓存的离屏尺寸就会随滚动变化，缓存键永远命中不了。
func (d *DisplayList) append(c Command) {
	if d.frozen {
		panic("wg/render: 冻结后的 DisplayList 被修改（违反 C23）")
	}
	d.cmds = append(d.cmds, c)
	switch c.Op {
	case OpFillRect, OpFillPath, OpStrokePath, OpDrawText, OpDrawImage:
		d.painted++
		if !c.Rect.IsEmpty() {
			d.bounds = d.bounds.Union(c.Rect)
		}
	}
}

// versionSeq 保证跨 DisplayList 的 version 全局唯一递增。
//
// 用全局计数器而不是每实例计数器：缓存键是 (pointer, version)，
// 复用同一个 DisplayList 对象时，实例内计数也能工作，但全局递增更省心，
// 且便于日志里认出"这是第几版"。
var versionSeq uint32

func nextVersion() uint32 { return atomic.AddUint32(&versionSeq, 1) }
