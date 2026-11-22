package widgets

// ── Splitter 分栏（第一百零七轮，低频项收敛）─────────────────────
//
// antd：<Splitter orientation onResize lazy><Splitter.Panel
// defaultSize min max resizable>。可拖拽分栏布局：面板间 2px 拖拽
// 把手（splitBarSize），6px 触发区（splitTriggerSize），hover/拖拽
// 中主色高亮。
//
// 布局：sizes[] 为各面板主轴像素（把手占 2×(n-1)）；拖拽沿主轴
// delta 直接改相邻两面板（min/max 钳制，超出时向另一侧传导到极限
// 为止——简化为单对钳制）。SetSizes 编程驱动入口（受控语义，走同
// 款回调——审计/自动化用；真拖拽 Move 注入序列在 headless 下不稳）。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// SplitterOrientation 是分栏方向（antd orientation/layout）。
type SplitterOrientation uint8

const (
	// SplitHorizontal 水平分栏（默认）。
	SplitHorizontal SplitterOrientation = iota
	// SplitVertical 垂直分栏。
	SplitVertical
)

// SplitterPanel 是分栏面板定义。
type SplitterPanel struct {
	Content ui.Widget
	// DefaultSize 初始主轴像素（<=0 时平分剩余）。
	DefaultSize float32
	// Min/Max 主轴像素钳制（<=0 不限）。
	Min, Max float32
	// Resizable 是否可拖拽（默认 true）。
	Resizable bool
}

// Splitter 是分栏容器。
type Splitter struct {
	controlBase
	Orientation SplitterOrientation
	Panels      []SplitterPanel
	// Width/Height 总尺寸（水平用 Width，垂直用 Height；<=0 时
	// 水平 600 / 垂直 300）。
	Width  float32
	Height float32

	OnResize      func(sizes []float32)
	OnResizeStart func(sizes []float32)
	OnResizeEnd   func(sizes []float32)

	sizes    []float32 // 各面板主轴像素（inited 后有效）
	inited   bool
	dragging bool // 拖拽中（把手主色高亮）
	dragIdx  int  // 拖拽中的把手序（面板 i 与 i+1 之间）
}

// NewSplitter 创建分栏。
func NewSplitter(panels ...SplitterPanel) *Splitter {
	return &Splitter{Panels: panels}
}

const (
	splitBarSize     = 2 // 视觉线宽（antd splitBarSize）
	splitTriggerSize = 6 // 命中区宽（antd splitTriggerSize）
)

// initSizes 首次布局：DefaultSize 或平分剩余（总长 - 把手）。
func (s *Splitter) initSizes() {
	n := len(s.Panels)
	if n == 0 {
		return
	}
	total := s.mainSize()
	total -= splitBarSize * float32(n-1)
	fixed := float32(0)
	free := 0
	for i := range s.Panels {
		if s.Panels[i].DefaultSize > 0 {
			s.sizes[i] = s.Panels[i].DefaultSize
			fixed += s.sizes[i]
		} else {
			free++
		}
	}
	if free > 0 && total > fixed {
		each := (total - fixed) / float32(free)
		for i := range s.Panels {
			if s.Panels[i].DefaultSize <= 0 {
				s.sizes[i] = each
			}
		}
	} else if free > 0 {
		each := total / float32(n)
		for i := range s.Panels {
			if s.Panels[i].DefaultSize <= 0 {
				s.sizes[i] = each
			}
		}
	}
	s.inited = true
}

// mainSize 返回主轴总长。
func (s *Splitter) mainSize() float32 {
	if s.Orientation == SplitVertical {
		if s.Height > 0 {
			return s.Height
		}
		return 300
	}
	if s.Width > 0 {
		return s.Width
	}
	return 600
}

// SetSizes 编程设置各面板尺寸（受控语义同 antd size 属性：min/max
// 钳制后末面板吃掉余差守恒总长；走 OnResize 回调——审计/自动化
// 驱动入口）。
func (s *Splitter) SetSizes(sizes []float32) {
	if len(sizes) != len(s.Panels) {
		return
	}
	if !s.inited {
		s.sizes = make([]float32, len(s.Panels))
		s.initSizes()
	}
	n := len(s.sizes)
	// 先赋值+钳制非末面板。
	for i := 0; i < n-1; i++ {
		s.sizes[i] = s.clamp(i, sizes[i])
	}
	// 末面板吃余差（守恒总长：可用长 = mainSize - 把手），再钳 Min。
	avail := s.mainSize() - splitBarSize*float32(n-1)
	used := float32(0)
	for i := 0; i < n-1; i++ {
		used += s.sizes[i]
	}
	last := avail - used
	if p := s.Panels[n-1]; p.Min > 0 && last < p.Min {
		last = p.Min
	}
	if p := s.Panels[n-1]; p.Max > 0 && last > p.Max {
		last = p.Max
	}
	s.sizes[n-1] = last
	s.MarkDirty()
	if s.OnResize != nil {
		s.OnResize(append([]float32{}, s.sizes...))
	}
}

// Sizes 返回当前面板尺寸副本。
func (s *Splitter) Sizes() []float32 {
	out := append([]float32{}, s.sizes...)
	return out
}

// clamp 面板 i 尺寸钳制。
func (s *Splitter) clamp(i int, v float32) float32 {
	if p := s.Panels[i]; p.Min > 0 && v < p.Min {
		v = p.Min
	}
	if p := s.Panels[i]; p.Max > 0 && v > p.Max {
		v = p.Max
	}
	if v < 0 {
		v = 0
	}
	return v
}

// buildPanel 构建面板内容节点。
func (s *Splitter) buildPanel(ctx *ui.BuildContext, i int) *scene.Node {
	p := s.Panels[i]
	node := scene.Box(scene.BoxSpec{}).SetSize2(s.sizes[i], s.mainSize())
	if s.Orientation == SplitVertical {
		node.SetSize2(s.Width, s.sizes[i])
	}
	if p.Content != nil {
		if ch := buildChild(ctx, p.Content); ch != nil {
			node.Add(ch)
		}
	}
	node.SetHit(scene.HitSpec{Label: "split-panel"})
	return node
}

// buildDragger 构建把手（2px 视觉线+6px 命中区+pan 手势）。
func (s *Splitter) buildDragger(ctx *ui.BuildContext, th *theme.Data, i int) *scene.Node {
	h := s.mainSize()
	hot := s.dragging && s.dragIdx == i
	// 视觉线 2px（水平=竖线），居中于 6px 命中区。
	line := scene.Box(scene.BoxSpec{Fill: scene.Solid(map[bool]geom.Color{
		true: th.Color.Primary, false: th.Color.Border,
	}[hot]), Radius: th.Radius.SM})
	if s.Orientation == SplitVertical {
		line.SetSize2(h, splitBarSize)
	} else {
		line.SetSize2(splitBarSize, h)
	}
	trigger := scene.Stack().SetSize2(map[bool]float32{true: splitTriggerSize, false: h}[s.Orientation == SplitVertical],
		map[bool]float32{true: h, false: splitTriggerSize}[s.Orientation == SplitVertical])
	if s.Orientation == SplitVertical {
		trigger.SetSize2(h, splitTriggerSize)
	} else {
		trigger.SetSize2(splitTriggerSize, h)
	}
	line.SetPositionKeepSize(true)
	if s.Orientation == SplitVertical {
		line.SetStackPosition(geom.Pt(0, (splitTriggerSize-splitBarSize)/2))
	} else {
		line.SetStackPosition(geom.Pt((splitTriggerSize-splitBarSize)/2, 0))
	}
	trigger.Add(line)
	trigger.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "split-dragger", Cursor: scene.CursorResizeEW})
	if s.Panels[i].Resizable && ctx.Input != nil {
		g := &ui.GestureDetector{
			OnPanStart: func(pos geom.Point) {
				s.dragging, s.dragIdx = true, i
				if s.OnResizeStart != nil {
					s.OnResizeStart(append([]float32{}, s.sizes...))
				}
				s.MarkDirty()
			},
			OnPan: func(pos, delta geom.Point) {
				d := delta.X
				if s.Orientation == SplitVertical {
					d = delta.Y
				}
				s.dragResize(i, d)
			},
			OnPanEnd: func(pos, velocity geom.Point) {
				s.dragging = false
				s.MarkDirty()
				if s.OnResizeEnd != nil {
					s.OnResizeEnd(append([]float32{}, s.sizes...))
				}
			},
		}
		ctx.Input.Set(trigger, g.Handler())
	}
	return trigger
}

// dragResize 把手 i 拖拽：左面板吃 delta（钳制后实际变化量），右
// 面板吃余差守恒总长（min/max 钳制）。
func (s *Splitter) dragResize(i int, delta float32) {
	if !s.inited {
		return
	}
	left := s.clamp(i, s.sizes[i]+delta)
	s.sizes[i] = left
	// 右面板吃余差。
	avail := s.mainSize() - splitBarSize*float32(len(s.sizes)-1)
	used := float32(0)
	for k := 0; k < len(s.sizes); k++ {
		if k != i+1 {
			used += s.sizes[k]
		}
	}
	right := s.clamp(i+1, avail-used)
	s.sizes[i+1] = right
	s.MarkDirty()
	if s.OnResize != nil {
		s.OnResize(append([]float32{}, s.sizes...))
	}
}

// Build 实现 ui.Widget。
func (s *Splitter) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(s)
	th := ctx.ThemeOr()
	if !s.inited {
		s.sizes = make([]float32, len(s.Panels))
		s.initSizes()
	}
	axis := scene.Horizontal
	if s.Orientation == SplitVertical {
		axis = scene.Vertical
	}
	row := scene.Flex(axis, 0)
	row.SetSize2(s.Width, s.Height)
	for i := range s.Panels {
		row.Add(s.buildPanel(ctx, i))
		if i < len(s.Panels)-1 {
			row.Add(s.buildDragger(ctx, th, i))
		}
	}
	_ = text.DefaultFont // 保留 text 引用（拖拽把手无文字；此行防未用导入）
	row.SetHit(scene.HitSpec{Label: "splitter"})
	return row
}

var _ = input.PointerDown // 保留 input 引用
