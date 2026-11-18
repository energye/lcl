package scene

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/text"
)

// 布局是两趟：**约束下传（measure）→ 尺寸上传（arrange）**。
//
// 实现上把"算自己的尺寸"和"放子节点"拆成 desiredSize / placeChildren 两步，
// 而不是 Flutter 那样的 measure 递归 + 独立 arrange 递归。原因：
//
//	容器的尺寸往往依赖子节点尺寸，而子节点的位置又依赖容器的最终尺寸
//	（CrossStretch、MainSpaceBetween、Grid 列宽都是这样）。拆成"先定尺寸
//	再摆放"会让这类容器必须两趟递归，代码量翻倍且容易写错顺序。
//
//	合成一个 Layout 方法后，每个容器只需要回答两个问题：我要多大？
//	拿到最终尺寸后子节点放哪？子节点可能被测量两次（第一次拿自然尺寸，
//	第二次拿 flex 分配或拉伸后的尺寸）——但那两次都是 Layout 调用，
//	有约束缓存兜底，代价可接受。
//
// 纪律（03-ENGINE.md §3）：measure 必须可重入、无副作用（除缓存）；
// 布局全程 float32，不做像素取整；禁止在布局里读窗口尺寸或调 TextOut。

// Layout 计算子树布局并返回自身尺寸。
//
// 约束与上次相同且不脏时直接返回缓存——这是"改一个 hover 颜色不重排整棵树"
// 的兜底（真正的保护是 MarkNeedsPaint 不冒泡 MarkNeedsLayout）。
func (n *Node) Layout(c Constraints) geom.Size {
	if n.laidOut && n.lastConstraints == c && !n.needsLayout {
		return n.size
	}
	n.lastConstraints = c
	n.constraints = c

	desired := n.desiredSize(c)
	n.size = c.Constrain(n.finishBox(c, desired))
	n.placeChildren()
	n.laidOut = true
	// 布局已经满足了这次的需要，就地消费掉 needsLayout。
	// needsPaint / needsComposite 不在这里清——它们由帧末统一消费
	//（Container 决定要不要重录/重建层），语义不同不能合并。
	n.needsLayout = false
	return n.size
}

// Relayout 强制重排子树（忽略缓存）。主题/DPI 变化或显式刷新时用。
func (n *Node) Relayout(c Constraints) geom.Size {
	n.invalidateLayoutTree()
	return n.Layout(c)
}

// invalidateLayoutTree 递归标脏布局。
func (n *Node) invalidateLayoutTree() {
	n.needsLayout = true
	n.laidOut = false
	for _, c := range n.children {
		c.invalidateLayoutTree()
	}
}

// finishBox 应用节点自身的尺寸覆盖（固定尺寸 / min-max / 宽高比）。
//
// 顺序：固定尺寸 → 节点 min/max → 宽高比。父约束最后再生效（在 Layout 里），
// 因此"父给了紧约束"永远赢——这是 ConstrainedBox 的语义，也是唯一不会
// 让子节点撑破父节点的顺序。
func (n *Node) finishBox(c Constraints, desired geom.Size) geom.Size {
	w, h := desired.W, desired.H
	if isSet(n.fixedW) {
		w = n.fixedW
	}
	if isSet(n.fixedH) {
		h = n.fixedH
	}
	if isSet(n.minW) {
		w = maxf(w, n.minW)
	}
	if isSet(n.maxW) {
		w = minf(w, n.maxW)
	}
	if isSet(n.minH) {
		h = maxf(h, n.minH)
	}
	if isSet(n.maxH) {
		h = minf(h, n.maxH)
	}
	if n.ratio > 0 {
		// 宽已定（固定或紧约束）→ 由宽推高；否则由高推宽。
		if isSet(n.fixedW) || c.IsTightWidth() {
			h = w / n.ratio
		} else {
			w = h * n.ratio
		}
	}
	return geom.Sz(maxf(0, w), maxf(0, h))
}

// desiredSize 计算"不考虑父约束收口"时的自然尺寸。
func (n *Node) desiredSize(c Constraints) geom.Size {
	switch n.kind {
	case KindText:
		return n.desiredText(c)
	case KindPath:
		return n.desiredPath(c)
	case KindImage:
		return n.desiredImage(c)
	case KindBox:
		return n.desiredBox(c)
	case KindStack:
		return n.desiredStack(c)
	case KindFlex:
		return n.desiredFlex(c)
	case KindGrid:
		return n.desiredGrid(c)
	case KindScroll:
		return n.desiredScroll(c)
	case KindClip, KindTransform, KindOpacity, KindEffect, KindSurface:
		return n.desiredSingle(c)
	case KindCustomPaint, KindEmpty:
		return c.MinSize()
	}
	return c.MinSize()
}

// placeChildren 在最终尺寸已知后定位子节点。
func (n *Node) placeChildren() {
	if len(n.children) == 0 {
		return
	}
	switch n.kind {
	case KindBox, KindStack:
		n.placeStackLike()
	case KindFlex:
		n.placeFlex()
	case KindGrid:
		n.placeGrid()
	case KindScroll:
		n.placeScroll()
	case KindClip, KindTransform, KindOpacity, KindEffect, KindSurface:
		n.placeSingle()
	}
}

// ── 叶子 ────────────────────────────────────────────────────────────

func (n *Node) desiredText(c Constraints) geom.Size {
	if n.text == nil || n.text.Layout == nil {
		return geom.Sz(0, 0)
	}
	// 目标排版宽度：节点显式 MaxWidth 优先；否则用父给的可用宽度。
	target := n.text.MaxWidth
	if target <= 0 && c.HasBoundedWidth() {
		target = c.MaxW
	}
	if target > 0 && absf(target-n.text.LayoutW) > 0.5 {
		// 宽度变了就重排。LayoutText 内部按 (字体, 文本) 缓存度量，
		// 所以这里不是"每帧调 GDI"，而是缓存命中。
		//
		// 字体优先取 TextSpec.LayoutFont；没填时退回 Layout 自带的字体。
		// 缺了这条回退会用零值字体重排——字号/行高全变，真窗表现是
		// "重建后文字整体变大/偏移"（输入框 hover 移出后右下偏移的根因）。
		font := n.text.LayoutFont
		if font.Size <= 0 && n.text.Layout != nil {
			font = n.text.Layout.Font
		}
		opts := text.Options{
			MaxWidth:   target,
			MaxLines:   n.text.MaxLines,
			Ellipsis:   n.text.Ellipsis,
			Align:      n.text.Layout.Align,
			LineHeight: 0,
		}
		n.text.Layout = text.LayoutText(n.text.Src, font, opts)
		n.text.LayoutW = target
	}
	sz := n.text.Layout.Size
	// 宽度取排版宽度（右对齐时布局框要撑满），高度取实际行高和。
	if target > 0 && target > sz.W {
		sz.W = target
	}
	return sz
}

func (n *Node) desiredPath(c Constraints) geom.Size {
	if n.path == nil || n.path.Path == nil {
		return geom.Sz(0, 0)
	}
	b := n.path.Path.Bounds()
	w, h := b.W(), b.H()
	if n.path.Paint.Style == 0 && n.path.Paint.StrokeWidth > 0 {
		// 描边向外扩半个线宽
		half := n.path.Paint.StrokeWidth * 0.5
		w += 2 * half
		h += 2 * half
	}
	return geom.Sz(w, h)
}

func (n *Node) desiredImage(c Constraints) geom.Size {
	if n.image == nil || n.image.Image == nil {
		return geom.Sz(0, 0)
	}
	src := geom.Sz(float32(n.image.Image.W), float32(n.image.Image.H))
	switch {
	case isSet(n.fixedW) && isSet(n.fixedH):
		return geom.Sz(n.fixedW, n.fixedH)
	case isSet(n.fixedW):
		return geom.Sz(n.fixedW, n.fixedW*src.H/maxf(1, src.W))
	case isSet(n.fixedH):
		return geom.Sz(n.fixedH*src.W/maxf(1, src.H), n.fixedH)
	}
	return src
}

// desiredBox 是"背景 + 子节点按对齐叠放"的自然尺寸：内容区取子节点包围盒，
// 加上 padding。
func (n *Node) desiredBox(c Constraints) geom.Size {
	inner := c.Deflate(n.padding)
	var content geom.Size
	for _, ch := range n.children {
		if !ch.visible {
			continue
		}
		cc := inner.Loosen()
		s := ch.Layout(cc)
		s = addMargin(ch, s)
		content.W = maxf(content.W, s.W)
		content.H = maxf(content.H, s.H)
	}
	if n.kind == KindBox && len(n.children) == 0 {
		content = geom.Sz(0, 0)
	}
	return geom.Sz(content.W+n.padding.Horizontal(), content.H+n.padding.Vertical())
}

func (n *Node) desiredSingle(c Constraints) geom.Size {
	if len(n.children) == 0 {
		return c.MinSize()
	}
	inner := c.Deflate(n.padding)
	child := n.children[0]
	want := inner.Loosen()
	if n.kind == KindTransform || n.kind == KindOpacity || n.kind == KindEffect {
		// 这三种不改变布局尺寸：给子节点紧约束，尺寸随子节点。
		want = inner
	}
	s := child.Layout(want)
	s = addMargin(child, s)
	return geom.Sz(s.W+n.padding.Horizontal(), s.H+n.padding.Vertical())
}

// ── Stack / Box ─────────────────────────────────────────────────────

func (n *Node) desiredStack(c Constraints) geom.Size {
	inner := c.Deflate(n.padding)
	var content geom.Size
	for _, ch := range n.children {
		if !ch.visible {
			continue
		}
		if ch.positioned {
			// 绝对定位子节点不参与自然尺寸
			continue
		}
		s := ch.Layout(inner.Loosen())
		s = addMargin(ch, s)
		content.W = maxf(content.W, s.W)
		content.H = maxf(content.H, s.H)
	}
	return geom.Sz(content.W+n.padding.Horizontal(), content.H+n.padding.Vertical())
}

// placeStackLike 把子节点按 align 叠放在 padding 内容区。
func (n *Node) placeStackLike() {
	inner := n.ContentSize()
	ox, oy := n.ContentOrigin().X, n.ContentOrigin().Y
	for _, ch := range n.children {
		if ch.positioned {
			continue // placeChildren 末尾统一处理
		}
		s := ch.Layout(LooseConstraints(inner))
		s = addMargin(ch, s)
		p := n.align.AlignChild(inner.W, inner.H, s.W, s.H)
		ch.offset = geom.Pt(ox+p.X+ch.margin.Left, oy+p.Y+ch.margin.Top)
	}
	for _, ch := range n.children {
		if !ch.positioned {
			continue
		}
		// 绝对定位：偏移相对内容区原点。尺寸默认按内容区紧约束（铺满）；
		// keepSize=true 的（浮层/影子/滑块手柄）按自身期望尺寸，不受父级
		// 内容区钳制——松约束仍会把超过父级的子节点压小（4px 滑轨上的
		// 14px 手柄被压成 8px 椭圆），与"保持自身尺寸"的语义相悖。
		if ch.keepSize {
			ch.Layout(Unbounded())
		} else {
			ch.Layout(TightConstraints(inner))
		}
		ch.offset = geom.Pt(ox+ch.position.X+ch.margin.Left, oy+ch.position.Y+ch.margin.Top)
	}
}

// ── Flex / Wrap ─────────────────────────────────────────────────────

// mainOf / crossOf 取尺寸在主轴/交叉轴上的分量。
func mainOf(a Axis, s geom.Size) float32 {
	if a == Horizontal {
		return s.W
	}
	return s.H
}

func crossOf(a Axis, s geom.Size) float32 {
	if a == Horizontal {
		return s.H
	}
	return s.W
}

func sizeOf(a Axis, main, cross float32) geom.Size {
	if a == Horizontal {
		return geom.Sz(main, cross)
	}
	return geom.Sz(cross, main)
}

// ptOf 同 sizeOf，但产出 Point（放置子节点时用）。
func ptOf(a Axis, main, cross float32) geom.Point {
	if a == Horizontal {
		return geom.Pt(main, cross)
	}
	return geom.Pt(cross, main)
}

// desiredFlex：非 flex 子节点按自然尺寸累加，flex 子节点先记 0——
// 它们的尺寸必须等主轴可用空间确定后才知道，提前测只会浪费一次递归。
func (n *Node) desiredFlex(c Constraints) geom.Size {
	inner := c.Deflate(n.padding)
	innerSize := inner.MaxSize()
	if n.wrap {
		return n.desiredWrap(inner)
	}
	var mainSum, crossMax float32
	count := 0
	for _, ch := range n.children {
		if !ch.visible {
			continue
		}
		if ch.flex > 0 {
			count++
			continue
		}
		s := ch.Layout(n.childConstraints(innerSize, false, 0))
		s = addMargin(ch, s)
		mainSum += mainOf(n.axis, s)
		crossMax = maxf(crossMax, crossOf(n.axis, s))
		count++
	}
	if count > 1 {
		mainSum += n.gap * float32(count-1)
	}
	s := sizeOf(n.axis, mainSum, crossMax)
	return geom.Sz(s.W+n.padding.Horizontal(), s.H+n.padding.Vertical())
}

// desiredWrap 需要真正跑一遍换行，才知道有几行、最宽一行多宽。
func (n *Node) desiredWrap(inner Constraints) geom.Size {
	maxMain := mainOf(n.axis, inner.MaxSize())
	if !hasBoundedMain(n.axis, inner) {
		maxMain = Inf
	}
	var runMain, runCross, maxRunMain, totalCross float32
	runs := 1
	for _, ch := range n.children {
		if !ch.visible {
			continue
		}
		s := ch.Layout(inner.Loosen())
		s = addMargin(ch, s)
		m, cr := mainOf(n.axis, s), crossOf(n.axis, s)
		need := m
		if runMain > 0 {
			need += n.gap
		}
		if maxMain < Inf && runMain+need > maxMain && runMain > 0 {
			maxRunMain = maxf(maxRunMain, runMain)
			totalCross += runCross
			runs++
			runMain, runCross = m, cr
			continue
		}
		runMain += need
		runCross = maxf(runCross, cr)
	}
	maxRunMain = maxf(maxRunMain, runMain)
	totalCross += runCross
	if runs > 1 {
		totalCross += n.runGap * float32(runs-1)
	}
	w := maxRunMain
	h := totalCross
	if n.axis == Horizontal {
		return geom.Sz(w+n.padding.Horizontal(), h+n.padding.Vertical())
	}
	return geom.Sz(h+n.padding.Horizontal(), w+n.padding.Vertical())
}

func hasBoundedMain(a Axis, c Constraints) bool {
	if a == Horizontal {
		return c.HasBoundedWidth()
	}
	return c.HasBoundedHeight()
}

// childConstraints 给 flex 子节点造约束。
// flex=true 时主轴是"固定份额"，否则主轴是 0..Inf（自然尺寸）。
func (n *Node) childConstraints(inner geom.Size, flex bool, share float32) Constraints {
	if flex {
		if n.axis == Horizontal {
			return NewConstraints(share, share, 0, inner.H)
		}
		return NewConstraints(0, inner.W, share, share)
	}
	if n.axis == Horizontal {
		return NewConstraints(0, Inf, 0, inner.H)
	}
	return NewConstraints(0, inner.W, 0, Inf)
}

// placeFlex 按最终尺寸分配主轴空间并定位。
func (n *Node) placeFlex() {
	if n.wrap {
		n.placeWrap()
		return
	}
	inner := n.ContentSize()
	ox, oy := n.ContentOrigin().X, n.ContentOrigin().Y
	main := mainOf(n.axis, inner)
	cross := crossOf(n.axis, inner)

	// ① 非 flex 子节点先测（主轴拿自然尺寸）
	var used float32
	visible := 0
	totalFlex := float32(0)
	for _, ch := range n.children {
		if !ch.visible {
			continue
		}
		visible++
		if ch.flex > 0 {
			totalFlex += ch.flex
			continue
		}
		s := ch.Layout(n.childConstraints(inner, false, 0))
		used += mainOf(n.axis, addMargin(ch, s))
	}
	if visible > 1 {
		used += n.gap * float32(visible-1)
	}

	// ② 剩余空间按权重分配
	free := maxf(0, main-used)
	unit := float32(0)
	if totalFlex > 0 {
		unit = free / totalFlex
	}

	// ③ 测 flex 子节点，并把每项的"主轴份额区间"记下来。
	//
	// 为什么要记：交叉轴拉伸（CrossStretch）需要**重测**子节点（交叉轴变紧），
	// 重测时主轴约束必须保持第③步分配的结果，否则会把 flex 份额冲掉——
	// 这正是 `--layout` 第一轮把 100 宽打成 0 宽的原因。
	shareMin := make([]float32, len(n.children))
	shareMax := make([]float32, len(n.children))
	for i, ch := range n.children {
		shareMax[i] = Inf
		if !ch.visible || ch.flex <= 0 {
			continue
		}
		share := maxf(0, unit*ch.flex-marginMain(n.axis, ch))
		shareMax[i] = share
		if ch.flexFit == FlexLoose {
			shareMin[i] = 0
		} else {
			shareMin[i] = share
		}
		var cs Constraints
		if n.axis == Horizontal {
			cs = NewConstraints(shareMin[i], shareMax[i], 0, inner.H)
		} else {
			cs = NewConstraints(0, inner.W, shareMin[i], shareMax[i])
		}
		ch.Layout(cs)
	}

	// ④ 重算实际占用（flex 子节点的真实尺寸由它们自己决定）
	used = 0
	for _, ch := range n.children {
		if !ch.visible {
			continue
		}
		used += mainOf(n.axis, addMargin(ch, ch.Size()))
	}
	if visible > 1 {
		used += n.gap * float32(visible-1)
	}

	// ⑤ 主轴游标：MainSpaceXxx 决定起点与额外间隙
	cursor := float32(0)
	extra := float32(0)
	switch n.mainAlign {
	case MainStart:
		cursor = 0
	case MainCenter:
		cursor = maxf(0, (main-used)/2)
	case MainEnd:
		cursor = maxf(0, main-used)
	case MainSpaceBetween, MainSpaceAround, MainSpaceEvenly:
		gaps := float32(0)
		switch n.mainAlign {
		case MainSpaceBetween:
			if visible > 1 {
				gaps = float32(visible - 1)
			}
		case MainSpaceAround:
			gaps = float32(visible)
		case MainSpaceEvenly:
			gaps = float32(visible + 1)
		}
		if gaps > 0 {
			slack := maxf(0, main-used)
			if n.mainAlign == MainSpaceAround || n.mainAlign == MainSpaceEvenly {
				cursor = slack / gaps
			}
			extra = slack / gaps
		}
	}

	first := true
	for i, ch := range n.children {
		if !ch.visible {
			continue
		}
		if !first {
			cursor += n.gap + extra
		}
		first = false

		// 交叉轴拉伸必须**先于**读数：它会把交叉尺寸改掉。
		// 但主轴游标依赖的 m 是主轴分量，不受交叉轴改变影响，所以顺序安全。
		if n.crossAlign == CrossStretch && !isSet(n.fixedCross(ch)) {
			target := maxf(0, cross-marginCross(n.axis, ch))
			var cs Constraints
			if n.axis == Horizontal {
				cs = NewConstraints(shareMin[i], shareMax[i], target, target)
			} else {
				cs = NewConstraints(target, target, shareMin[i], shareMax[i])
			}
			ch.Layout(cs)
		}

		s := addMargin(ch, ch.Size())
		m, cr := mainOf(n.axis, s), crossOf(n.axis, s)

		var co float32
		switch n.crossAlign {
		case CrossStart:
			co = 0
		case CrossCenter:
			co = (cross - cr) / 2
		case CrossEnd:
			co = cross - cr
		case CrossStretch:
			co = 0
		}

		pt := ptOf(n.axis, cursor+marginMainStart(n.axis, ch), co+marginCrossStart(n.axis, ch))
		ch.offset = geom.Pt(ox+pt.X, oy+pt.Y)
		cursor += m
	}
}

// placeWrap 按行（run）摆放。与 placeFlex 的差别只在"换行"这一步：
// flex 权重在 wrap 里不参与（换行容器不支持 flex-grow）。
func (n *Node) placeWrap() {
	inner := n.ContentSize()
	ox, oy := n.ContentOrigin().X, n.ContentOrigin().Y
	maxMain := mainOf(n.axis, inner)

	type item struct {
		n       *Node
		main    float32
		cross   float32
		started bool
	}
	var runs [][]item
	var cur []item
	var runMain float32
	for _, ch := range n.children {
		if !ch.visible {
			continue
		}
		s := ch.Layout(LooseConstraints(inner))
		s = addMargin(ch, s)
		m, cr := mainOf(n.axis, s), crossOf(n.axis, s)
		need := m
		if len(cur) > 0 {
			need += n.gap
		}
		if len(cur) > 0 && runMain+need > maxMain {
			runs = append(runs, cur)
			cur = nil
			runMain = 0
			need = m
		}
		cur = append(cur, item{n: ch, main: m, cross: cr})
		runMain += need
	}
	if len(cur) > 0 {
		runs = append(runs, cur)
	}

	cursor := float32(0)
	for _, run := range runs {
		var runCross, runTotal float32
		for i, it := range run {
			runCross = maxf(runCross, it.cross)
			runTotal += it.main
			if i > 0 {
				runTotal += n.gap
			}
		}
		x := float32(0)
		for i, it := range run {
			if i > 0 {
				x += n.gap
			}
			co := float32(0)
			switch n.crossAlign {
			case CrossCenter:
				co = (runCross - it.cross) / 2
			case CrossEnd:
				co = runCross - it.cross
			}
			pt := ptOf(n.axis, x+marginMainStart(n.axis, it.n), co+marginCrossStart(n.axis, it.n))
			it.n.offset = geom.Pt(ox+pt.X, oy+cursor+pt.Y)
			x += it.main
		}
		cursor += runCross + n.runGap
	}
}

// margin* 取子节点外边距在主/交叉轴上的分量。
func marginMain(a Axis, ch *Node) float32 {
	if a == Horizontal {
		return ch.margin.Left + ch.margin.Right
	}
	return ch.margin.Top + ch.margin.Bottom
}

func marginCross(a Axis, ch *Node) float32 { return marginMain(a.Flip(), ch) }

func marginMainStart(a Axis, ch *Node) float32 {
	if a == Horizontal {
		return ch.margin.Left
	}
	return ch.margin.Top
}

func marginCrossStart(a Axis, ch *Node) float32 { return marginMainStart(a.Flip(), ch) }

func addMargin(ch *Node, s geom.Size) geom.Size {
	return geom.Sz(s.W+ch.margin.Horizontal(), s.H+ch.margin.Vertical())
}

// fixedCross 返回子节点交叉轴上的显式尺寸（用于判断能否拉伸）。
func (n *Node) fixedCross(ch *Node) float32 {
	if n.axis == Horizontal {
		return ch.fixedH
	}
	return ch.fixedW
}

// ── Grid ────────────────────────────────────────────────────────────

func (n *Node) desiredGrid(c Constraints) geom.Size {
	inner := c.Deflate(n.padding)
	cols := maxi(1, n.columns)
	cellW := float32(0)
	if inner.HasBoundedWidth() {
		cellW = (inner.MaxW - n.gap*float32(cols-1)) / float32(cols)
	}
	var col, row float32
	var rowH float32
	for _, ch := range n.children {
		if !ch.visible {
			continue
		}
		span := maxi(1, mini(ch.span, cols))
		var cs Constraints
		if cellW > 0 {
			w := cellW*float32(span) + n.gap*float32(span-1)
			cs = TightConstraints(geom.Sz(w, 0)).Loosen()
			cs = NewConstraints(0, w, 0, Inf)
		} else {
			cs = Unbounded()
		}
		s := ch.Layout(cs)
		s = addMargin(ch, s)
		rowH = maxf(rowH, s.H)
		row += s.H
		col += s.W
	}
	cols2 := cols
	if cellW <= 0 && len(n.children) > 0 {
		cellW = col / float32(cols2)
	}
	rows := (len(n.children) + cols - 1) / cols
	totalH := row + n.gap*float32(maxi(0, rows-1))
	totalW := cellW*float32(cols) + n.gap*float32(cols-1)
	return geom.Sz(totalW+n.padding.Horizontal(), totalH+n.padding.Vertical())
}

func (n *Node) placeGrid() {
	inner := n.ContentSize()
	ox, oy := n.ContentOrigin().X, n.ContentOrigin().Y
	cols := maxi(1, n.columns)
	cellW := (inner.W - n.gap*float32(cols-1)) / float32(cols)

	// 逐行算行高（跨列子节点按 span 展宽）
	type cell struct {
		ch   *Node
		span int
		col  int
		row  int
	}
	cells := make([]cell, 0, len(n.children))
	row := 0
	col := 0
	for _, ch := range n.children {
		if !ch.visible {
			continue
		}
		span := maxi(1, mini(ch.span, cols))
		if col+span > cols {
			row++
			col = 0
		}
		cells = append(cells, cell{ch: ch, span: span, col: col, row: row})
		col += span
		if col >= cols {
			row++
			col = 0
		}
	}

	rows := row
	if col > 0 {
		rows++
	}
	rowH := make([]float32, maxi(1, rows))
	for _, cl := range cells {
		w := cellW*float32(cl.span) + n.gap*float32(cl.span-1)
		s := cl.ch.Layout(NewConstraints(0, w, 0, Inf))
		s = addMargin(cl.ch, s)
		if cl.row < len(rowH) {
			rowH[cl.row] = maxf(rowH[cl.row], s.H)
		}
	}
	for _, cl := range cells {
		w := cellW*float32(cl.span) + n.gap*float32(cl.span-1)
		h := rowH[cl.row]
		target := geom.Sz(w, maxf(0, h-cl.ch.margin.Vertical()))
		cl.ch.Layout(TightConstraints(target))
		x := float32(cl.col)*(cellW+n.gap) + cl.ch.margin.Left
		var ry float32
		for k := 0; k < cl.row; k++ {
			ry += rowH[k] + n.gap
		}
		cl.ch.offset = geom.Pt(ox+x, oy+ry+cl.ch.margin.Top)
	}
}

// ── Scroll ──────────────────────────────────────────────────────────

func (n *Node) desiredScroll(c Constraints) geom.Size {
	// 视口在滚动轴上吃满约束，另一轴随内容（但有界时吃满）
	if n.axis == Vertical {
		h := c.MaxH
		if !c.HasBoundedHeight() {
			h = n.contentSize.H
		}
		w := c.MaxW
		if !c.HasBoundedWidth() {
			w = n.contentSize.W
		}
		return geom.Sz(w, h)
	}
	w := c.MaxW
	if !c.HasBoundedWidth() {
		w = n.contentSize.W
	}
	h := c.MaxH
	if !c.HasBoundedHeight() {
		h = n.contentSize.H
	}
	return geom.Sz(w, h)
}

// placeScroll 用无限约束测内容，记录内容尺寸，再按偏移平移。
func (n *Node) placeScroll() {
	if len(n.children) == 0 {
		return
	}
	inner := n.ContentSize()
	var cs Constraints
	if n.axis == Vertical {
		// 竖直滚动：宽度固定（内容不能横向溢出视口），高度无界
		cs = NewConstraints(inner.W, inner.W, 0, Inf)
	} else {
		cs = NewConstraints(0, Inf, inner.H, inner.H)
	}
	ch := n.children[0]
	s := ch.Layout(cs)
	s = addMargin(ch, s)
	n.contentSize = s
	n.scrollOffset = clampf(n.scrollOffset, 0, n.maxScroll())
	ch.offset = geom.Pt(n.padding.Left, n.padding.Top)
	if n.axis == Vertical {
		ch.offset.Y -= n.scrollOffset
	} else {
		ch.offset.X -= n.scrollOffset
	}
}

// ── 单子节点容器 ────────────────────────────────────────────────────

func (n *Node) placeSingle() {
	if len(n.children) == 0 {
		return
	}
	inner := n.ContentSize()
	ch := n.children[0]
	var cs Constraints
	switch n.kind {
	case KindTransform, KindOpacity, KindEffect:
		cs = TightConstraints(inner)
	default:
		cs = LooseConstraints(inner)
	}
	s := ch.Layout(cs)
	s = addMargin(ch, s)
	p := n.align.AlignChild(inner.W, inner.H, s.W, s.H)
	ch.offset = geom.Pt(n.ContentOrigin().X+p.X+ch.margin.Left, n.ContentOrigin().Y+p.Y+ch.margin.Top)
}

func absf(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}
