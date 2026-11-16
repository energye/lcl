package raster

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
)

// Caps 是后端能力声明。
//
// 后端必须诚实上报：报 false 的指令由上层计入 degraded 并显式报出来（C22），
// 不允许"报了 true 但悄悄不画"。
type Caps struct {
	// AA 支持抗锯齿。
	AA bool
	// Stroke 支持描边（否则只能填充）。
	Stroke bool
	// Dash 支持虚线。
	Dash bool
	// LinearGradient 支持线性渐变。
	LinearGradient bool
	// RadialGradient 支持径向渐变。
	RadialGradient bool
	// ClipRect 支持矩形裁剪。
	ClipRect bool
	// ClipPath 支持路径裁剪（精确）。
	ClipPath bool
	// Text 支持文字回放。
	Text bool
	// Image 支持绘制位图。
	Image bool
	// Transform 支持任意仿射变换（否则只支持平移/整数缩放）。
	Transform bool
}

// Stats 是一次绘制的统计。
type Stats struct {
	// Ops 处理的指令数。
	Ops int
	// Pixels 实际触碰的像素数（覆盖率的粗略度量）。
	Pixels int
	// Degraded 因能力缺失而跳过的指令数。
	Degraded int
	// Spans 生成的扫描线跨度数。
	Spans int
}

// Add 累加另一份统计。
func (s *Stats) Add(o Stats) {
	s.Ops += o.Ops
	s.Pixels += o.Pixels
	s.Degraded += o.Degraded
	s.Spans += o.Spans
}

// Backend 是光栅后端。
//
// 入参是"已经录制好的不可变指令表"，出参是统计——后端不持有对象生命周期，
// 也不缓存指令表（缓存归 composite 的 RasterCache 管，见 C7）。
type Backend interface {
	// Name 是后端名（software / gdi）。
	Name() string
	// Caps 返回能力声明。
	Caps() Caps
	// Draw 把指令表绘制到 dst 上，只处理与 clip 相交的部分。
	//
	// clip 为设备空间矩形；dst 与 clip 使用同一坐标系（原点在 dst 左上角）。
	Draw(dst *Backing, dl *render.DisplayList, clip geom.Rect) Stats
}

// Hit 判断矩形区域是否含有非透明像素（golden/矩阵断言用）。
//
// 只扫 alpha 通道，遇到第一个非零即返回。
func Hit(b *Backing, r geom.Rect) bool {
	if b == nil {
		return false
	}
	x0, y0, x1, y1, ok := b.clipRect(r)
	if !ok {
		return false
	}
	for y := y0; y < y1; y++ {
		row := b.pix[y*b.w : y*b.w+b.w]
		for x := x0; x < x1; x++ {
			if row[x]>>24 != 0 {
				return true
			}
		}
	}
	return false
}

// AlphaBounds 返回非透明像素的包围盒（空图返回 ZeroRect）。
//
// 这是 --render-matrix 的核心断言工具：比较两个后端画出的形状**位置**，
// 而不是逐像素相等（两个后端的 AA 算法不同，逐像素永远不可能相等）。
func AlphaBounds(b *Backing) geom.Rect {
	if b == nil {
		return geom.ZeroRect
	}
	minX, minY := b.w, b.h
	maxX, maxY := -1, -1
	for y := 0; y < b.h; y++ {
		row := b.pix[y*b.w : y*b.w+b.w]
		for x := 0; x < b.w; x++ {
			if row[x]>>24 == 0 {
				continue
			}
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
		}
	}
	if maxX < 0 {
		return geom.ZeroRect
	}
	return geom.RectLTRB(float32(minX), float32(minY), float32(maxX+1), float32(maxY+1))
}

// AlphaCount 返回非透明像素数量（覆盖率度量）。
func AlphaCount(b *Backing) int {
	if b == nil {
		return 0
	}
	n := 0
	for _, px := range b.pix {
		if px>>24 != 0 {
			n++
		}
	}
	return n
}

// AlphaCentroid 返回非透明像素的质心（亚像素精度），用于位置比对。
func AlphaCentroid(b *Backing) (x, y float32, ok bool) {
	if b == nil {
		return 0, 0, false
	}
	var sx, sy, sw float64
	for py := 0; py < b.h; py++ {
		row := b.pix[py*b.w : py*b.w+b.w]
		for px := 0; px < b.w; px++ {
			a := row[px] >> 24
			if a == 0 {
				continue
			}
			w := float64(a)
			sx += float64(px) * w
			sy += float64(py) * w
			sw += w
		}
	}
	if sw == 0 {
		return 0, 0, false
	}
	return float32(sx/sw) + 0.5, float32(sy/sw) + 0.5, true
}
