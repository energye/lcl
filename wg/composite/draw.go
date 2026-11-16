package composite

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
)

// drawLayer 把一个层画进 dst，ox/oy 是本层局部原点在 dst 中的像素位置。
//
// alpha 是"尚未施加到内容上的整体不透明度"：
//   - 走离屏的层自己处理透明度（混合时乘），所以给子层传 1；
//   - 扁平化的层（含文字的降级层、根层）把 alpha 乘给子层，
//     由 drawPicture 逐指令缩放——近似但不会丢内容。
//
// 这条区分很重要：如果两处都乘，透明度会被平方（0.5 变 0.25）。
func (c *Compositor) drawLayer(l *Layer, dst *raster.Backing, ox, oy int, clip geom.Rect, alpha float32) {
	if l == nil {
		return
	}
	switch l.kind {
	case LayerPicture:
		st := c.drawPicture(dst, l.picture, ox, oy, clip, alpha)
		c.st.ops += st.Ops
		c.st.degraded += st.Degraded
	case LayerGroup:
		if !l.offscreen {
			for _, ch := range l.children {
				c.drawLayer(ch, dst, ox+roundf(ch.origin.X), oy+roundf(ch.origin.Y), clip, alpha*l.opacity)
			}
			return
		}
		c.drawGroup(l, dst, ox, oy, clip)
	}
}

// drawPicture 回放一段指令（形状走光栅后端，文字走设备空间通道）。
func (c *Compositor) drawPicture(dst *raster.Backing, dl *render.DisplayList, dx, dy int, clip geom.Rect, alpha float32) raster.Stats {
	if dl == nil || dl.Empty() || dst == nil {
		return raster.Stats{}
	}
	if alpha < 0.999 {
		dl = dl.ScaleAlpha(alpha)
	}
	if dx != 0 || dy != 0 {
		dl = dl.OffsetBy(float32(dx), float32(dy))
	}
	st := c.backend.Draw(dst, dl, clip)
	if dl.HasText() {
		if c.text != nil {
			c.st.textLines += c.text.PaintText(dst, dl, clip)
		} else {
			c.st.textSkipped = true
		}
	}
	return st
}

// drawGroup 把需要离屏的组渲染到一块缓冲里再整体混合。
func (c *Compositor) drawGroup(l *Layer, dst *raster.Backing, ox, oy int, clip geom.Rect) {
	w, h := l.backingSize()
	if w <= 0 || h <= 0 {
		return
	}
	base := l.bounds.RoundOut().Min
	dx := ox + int(base.X)
	dy := oy + int(base.Y)

	// 脏区外整层跳过：这是 damage 的第二个杠杆（第一个是清屏面积）。
	if !geom.RectXYWH(float32(dx), float32(dy), float32(w), float32(h)).Intersects(clip) {
		c.st.layerSkip++
		return
	}

	key := CacheKey{ID: l.node.ID(), Version: l.version, W: int32(w), H: int32(h)}
	if c.cacheOn && l.cacheable && c.cache != nil {
		if e := c.cache.Lookup(key); e != nil {
			l.cached = true
			c.st.layerHits++
			c.st.blended += blendBacking(dst, e.backing, dx, dy, l.opacity, e.mask)
			return
		}
	}

	off := c.cache.Lease(w, h)
	if off == nil {
		return
	}
	off.Clear()
	for _, ch := range l.children {
		// 子层的 origin 在**本层局部空间**，离屏的 (0,0) 是 bounds.Min，
		// 所以减掉 base 才是离屏像素坐标。
		cx := roundf(ch.origin.X - base.X)
		cy := roundf(ch.origin.Y - base.Y)
		c.drawLayer(ch, off, cx, cy, off.Bounds(), 1)
	}

	var mask []uint8
	if l.clip != ClipNone && l.clip != ClipRect {
		mask = buildMask(l, w, h)
	}
	c.st.layerRaster++
	c.st.blended += blendBacking(dst, off, dx, dy, l.opacity, mask)

	if c.cacheOn && l.cacheable && c.cache != nil {
		c.cache.Store(key, off, mask)
	} else {
		c.cache.Release(off)
	}
}

// roundf 四舍五入到整数像素。
//
// 为什么允许四舍五入而不是保留亚像素：离屏缓冲只能落在整数像素上，
// 非整数的落点要么插值（软件后端没做）要么错半像素。层内容本身
// 在层局部空间里已经是亚像素精度，整体落点取整后视觉上是"整层移动"，
// 这和 Skia 对 raster cache 的处理一致（缓存层按整像素贴图）。
func roundf(v float32) int {
	if v >= 0 {
		return int(v + 0.5)
	}
	return int(v - 0.5)
}
