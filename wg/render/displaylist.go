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

// Bounds 返回所有指令在设备空间的并集包围盒。
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
	d.hasText = false
	d.version = nextVersion()
}

// append 追加一条指令并更新包围盒（仅 Recorder 调用）。
func (d *DisplayList) append(c Command) {
	if d.frozen {
		panic("wg/render: 冻结后的 DisplayList 被修改（违反 C23）")
	}
	d.cmds = append(d.cmds, c)
	if !c.Rect.IsEmpty() {
		d.bounds = d.bounds.Union(c.Rect)
	}
}

// versionSeq 保证跨 DisplayList 的 version 全局唯一递增。
//
// 用全局计数器而不是每实例计数器：缓存键是 (pointer, version)，
// 复用同一个 DisplayList 对象时，实例内计数也能工作，但全局递增更省心，
// 且便于日志里认出"这是第几版"。
var versionSeq uint32

func nextVersion() uint32 { return atomic.AddUint32(&versionSeq, 1) }
