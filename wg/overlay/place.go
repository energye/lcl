package overlay

import "github.com/energye/lcl/wg/geom"

// ── 定位 ────────────────────────────────────────────────────────────
//
// 浮层定位看起来只是"算个坐标"，但真正难的是**边界行为**：
// 下拉框在屏幕底部打开时，内容朝下会跑出窗口；朝上会盖住触发按钮；
// 两边都放不下时还得挤压到窗口里。这三步（首选 → 翻转 → 位移）
// 是 dbmodeler.PlaceAt 里被反复重写过的逻辑，这里一次做对并单测。

// Placement 是首选方位。
type Placement uint8

const (
	// PlacementBottomStart 锚点下方，左对齐（下拉框默认）。
	PlacementBottomStart Placement = iota
	// PlacementBottom 锚点下方居中。
	PlacementBottom
	// PlacementBottomEnd 锚点下方，右对齐。
	PlacementBottomEnd
	// PlacementTopStart 锚点上方，左对齐。
	PlacementTopStart
	// PlacementTop 锚点上方居中。
	PlacementTop
	// PlacementTopEnd 锚点上方，右对齐。
	PlacementTopEnd
	// PlacementRightStart 锚点右侧，上对齐。
	PlacementRightStart
	// PlacementLeftStart 锚点左侧，上对齐。
	PlacementLeftStart
	// PlacementCenter 覆盖在锚点上居中（Modal/Dialog）。
	PlacementCenter
	// PlacementPoint 以鼠标点为左上角（ContextMenu）。
	PlacementPoint
)

// String 返回可读名。
func (p Placement) String() string {
	switch p {
	case PlacementBottomStart:
		return "bottom-start"
	case PlacementBottom:
		return "bottom"
	case PlacementBottomEnd:
		return "bottom-end"
	case PlacementTopStart:
		return "top-start"
	case PlacementTop:
		return "top"
	case PlacementTopEnd:
		return "top-end"
	case PlacementRightStart:
		return "right-start"
	case PlacementLeftStart:
		return "left-start"
	case PlacementCenter:
		return "center"
	case PlacementPoint:
		return "point"
	}
	return "unknown"
}

// Along 表示主轴是纵向（上下）还是横向（左右）。
func (p Placement) Along() (vertical, horizontal bool) {
	switch p {
	case PlacementBottomStart, PlacementBottom, PlacementBottomEnd,
		PlacementTopStart, PlacementTop, PlacementTopEnd:
		return true, false
	case PlacementRightStart, PlacementLeftStart:
		return false, true
	}
	return false, false
}

// Opposite 返回对侧方位（翻转用）。
func (p Placement) Opposite() Placement {
	switch p {
	case PlacementBottomStart:
		return PlacementTopStart
	case PlacementBottom:
		return PlacementTop
	case PlacementBottomEnd:
		return PlacementTopEnd
	case PlacementTopStart:
		return PlacementBottomStart
	case PlacementTop:
		return PlacementBottom
	case PlacementTopEnd:
		return PlacementBottomEnd
	case PlacementRightStart:
		return PlacementLeftStart
	case PlacementLeftStart:
		return PlacementRightStart
	}
	return p
}

// Placed 是一次定位的结果。
type Placed struct {
	// Rect 是最终矩形（根空间 DIP）。
	Rect geom.Rect
	// Placement 是**实际**使用的方位（可能已翻转）。
	Placement Placement
	// Flipped 表示发生过翻转（箭头要跟着换边）。
	Flipped bool
	// Shifted 表示发生过位移（贴边挤压）。
	Shifted bool
}

// Place 按方位算出首选矩形（不含边界修正）。
func Place(p Placement, anchor geom.Rect, size geom.Size, offset geom.Point) geom.Rect {
	ax, ay := anchor.Min.X, anchor.Min.Y
	aw, ah := anchor.W(), anchor.H()
	cx, cy := ax+aw/2, ay+ah/2
	switch p {
	case PlacementBottomStart:
		return geom.RectXYWH(ax+offset.X, anchor.Max.Y+offset.Y, size.W, size.H)
	case PlacementBottom:
		return geom.RectXYWH(cx-size.W/2+offset.X, anchor.Max.Y+offset.Y, size.W, size.H)
	case PlacementBottomEnd:
		return geom.RectXYWH(anchor.Max.X-size.W+offset.X, anchor.Max.Y+offset.Y, size.W, size.H)
	case PlacementTopStart:
		return geom.RectXYWH(ax+offset.X, ay-size.H-offset.Y, size.W, size.H)
	case PlacementTop:
		return geom.RectXYWH(cx-size.W/2+offset.X, ay-size.H-offset.Y, size.W, size.H)
	case PlacementTopEnd:
		return geom.RectXYWH(anchor.Max.X-size.W+offset.X, ay-size.H-offset.Y, size.W, size.H)
	case PlacementRightStart:
		return geom.RectXYWH(anchor.Max.X+offset.X, ay+offset.Y, size.W, size.H)
	case PlacementLeftStart:
		return geom.RectXYWH(ax-size.W-offset.X, ay+offset.Y, size.W, size.H)
	case PlacementCenter:
		return geom.RectXYWH(cx-size.W/2, cy-size.H/2, size.W, size.H)
	case PlacementPoint:
		return geom.RectXYWH(ax+offset.X, ay+offset.Y, size.W, size.H)
	}
	return geom.RectXYWH(ax, ay, size.W, size.H)
}

// Resolve 计算浮层最终矩形：首选 → 翻转 → 挤压。
//
// viewport 是可用区域（根空间，通常是窗口客户区）。零矩形表示不限边界。
func Resolve(p Placement, anchor geom.Rect, size geom.Size, offset geom.Point, viewport geom.Rect) Placed {
	out := Placed{Placement: p}
	out.Rect = Place(p, anchor, size, offset)
	if viewport.IsEmpty() {
		return out
	}
	// 第一步：放不下就翻到对侧。
	if !fits(out.Rect, viewport, p) {
		alt := p.Opposite()
		if alt != p {
			altRect := Place(alt, anchor, size, offset)
			if fits(altRect, viewport, alt) {
				out.Rect = altRect
				out.Placement = alt
				out.Flipped = true
				return out
			}
		}
	}
	// 第二步：两侧都放不下 → 沿主轴挤压到边界内（尽量不翻转）。
	if shifted := shift(out.Rect, viewport); shifted != out.Rect {
		out.Rect = shifted
		out.Shifted = true
	}
	return out
}

// fits 判断矩形在该方位下是否完全落在视口内。
//
// 只检查"该方位涉及的边"：Bottom 系列只关心下边界——内容比视口高
// 时无论怎么翻都放不下，那种情况交给挤压而不是反复翻转。
func fits(r, v geom.Rect, p Placement) bool {
	vertical, horizontal := p.Along()
	switch {
	case p == PlacementPoint || p == PlacementCenter:
		return v.ContainsRect(r)
	case vertical && (p == PlacementBottomStart || p == PlacementBottom || p == PlacementBottomEnd):
		return r.Max.Y <= v.Max.Y
	case vertical:
		return r.Min.Y >= v.Min.Y
	case horizontal && p == PlacementRightStart:
		return r.Max.X <= v.Max.X
	case horizontal:
		return r.Min.X >= v.Min.X
	}
	return v.ContainsRect(r)
}

// shift 把矩形平移进视口（超出右侧时左移，超出上侧时下移）。
func shift(r, v geom.Rect) geom.Rect {
	dx, dy := float32(0), float32(0)
	if r.Max.X > v.Max.X {
		dx = v.Max.X - r.Max.X
	}
	if r.Min.X+dx < v.Min.X {
		dx = v.Min.X - r.Min.X
	}
	if r.Max.Y > v.Max.Y {
		dy = v.Max.Y - r.Max.Y
	}
	if r.Min.Y+dy < v.Min.Y {
		dy = v.Min.Y - r.Min.Y
	}
	if r.W() > v.W() {
		dx = v.Min.X - r.Min.X
	}
	if r.H() > v.H() {
		dy = v.Min.Y - r.Min.Y
	}
	return r.Offset(dx, dy)
}
