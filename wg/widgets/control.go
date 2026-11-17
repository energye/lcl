package widgets

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── 控件公共基座 ───────────────────────────────────────────────────

// ctlSeq 是控件实例序号：生成跨重建稳定的节点身份（node.SetKey）。
var ctlSeq atomic.Uint64

// controlBase 管理控件的焦点节点与交互状态。
//
// 焦点节点必须跨 Build 存活：每次重建都新建 FocusNode 会让 Tab 链
// 长度随着重建次数增长，而且焦点会莫名其妙丢失。
type controlBase struct {
	ui.Base
	node    *scene.Node
	focus   *input.FocusNode
	hover   bool
	pressed bool
	tap     *ui.GestureDetector
	ctlKey  string // 跨重建稳定的实例身份（悬停差分认"换节点不换控件"用）
}

// Node 返回最近一次构建出的场景节点（诊断/自动化点击用）。
func (b *controlBase) Node() *scene.Node { return b.node }

func (b *controlBase) states(disabled bool) ui.StateSet {
	var s ui.StateSet
	if disabled {
		s |= ui.StateDisabled
	}
	if b.hover {
		s |= ui.StateHover
	}
	if b.pressed {
		s |= ui.StateActive
	}
	if b.focus != nil && b.focus.Focused() {
		s |= ui.StateFocus
	}
	return s
}

// attach 注册事件处理器并把持久焦点节点指向新构建的节点。
//
// 重建后 Dispatcher 的焦点仍指向**旧**节点，必须手动重定向：
// FocusManager.SetCurrent 看到"还是同一个 FocusNode"会直接返回，
// 不补这一步，输入框重建一次后就再也收不到按键。
func (b *controlBase) attach(ctx *ui.BuildContext, node *scene.Node, h input.Handler, canFocus bool) {
	// 状态回调（MarkDirty/SetState）必须能叫醒帧泵并把根 widget 置脏：
	// 容器只重建根 widget 的 Build 结果，深层组件不接线的话，静止窗口
	// （C14）点击后永远不会重绘——真窗表现为"点了没反应"。
	if ctx != nil {
		if ctx.InvalidateRoot != nil {
			b.Base.Attach(ctx.InvalidateRoot)
		} else if ctx.RequestFrame != nil {
			b.Base.Attach(ctx.RequestFrame)
		}
	}
	b.node = node
	// 实例身份钉到节点上：重建产出全新节点，但控件没换。输入层的
	// hover 差分按这个身份比较，"重建"不再被误判成"悬停迁移"
	// ——否则每次鼠标移动都触发一次整树重建（hover 风暴），而且
	// 重建落在 down/up 之间时指针捕获被破坏、tap 丢失。
	if b.ctlKey == "" {
		b.ctlKey = fmt.Sprintf("ctl-%d", ctlSeq.Add(1))
	}
	node.SetKey(b.ctlKey)
	if ctx.Input != nil {
		ctx.Input.Set(node, h)
	}
	if !canFocus {
		return
	}
	if b.focus == nil {
		b.focus = input.NewFocusNode(true)
	}
	old := b.focus.Node
	b.focus.Node = node
	if ctx.Focus == nil {
		return
	}
	root := ctx.Focus.Root()
	if root == nil {
		root = input.NewFocusNode(false)
		ctx.Focus.SetRoot(root)
	}
	if b.focus.Parent == nil {
		root.AddChild(b.focus)
	}
	if b.focus.Focused() && old != node && ctx.Input != nil {
		ctx.Input.SetFocus(node)
	}
}

// tapHandler 生成"点一下就回调"的手势处理器。
func (b *controlBase) tapHandler(ctx *ui.BuildContext, disabled bool, onTap func()) input.Handler {
	g := &ui.GestureDetector{}
	if !disabled && onTap != nil {
		g.OnTap = func(geom.Point) { onTap() }
	}
	h := g.Handler()
	h.Enter = func(*input.PointerEvent) {
		// 状态真正变化才置脏：身份差分之外的第二道防线，重复 Enter
		//（路径里同一控件出现两次等边角）不再触发重建。
		if !disabled && !b.hover {
			b.hover = true
			b.MarkDirty()
		}
	}
	h.Leave = func(*input.PointerEvent) {
		if b.hover || b.pressed {
			b.hover, b.pressed = false, false
			b.MarkDirty()
		}
	}
	// GestureDetector 每次按下都会替换成员，闭包需要长期存活；
	// 存在控件上而不是局部变量，避免 GC 后回调消失。
	b.tap = g
	return h
}

// ── Button：见 button.go（M8 对齐 antd type 体系） ─────────────────

// ── Checkbox / Radio / Switch ──────────────────────────────────────

// Checkbox 是复选框（M8 对齐 antd：indeterminate 半选态）。
type Checkbox struct {
	controlBase
	Checked       bool
	Indeterminate bool // 半选：横杠代替勾（父级聚合状态用）
	Disabled      bool
	Label         string
	OnChange      func(bool)
}

// NewCheckbox 创建复选框。
func NewCheckbox(label string, checked bool) *Checkbox {
	return &Checkbox{Label: label, Checked: checked}
}

// Build 实现 ui.Widget。
//
// 视觉对齐 antd（官方 checkbox-demo-check-all 对照）：
//   - 勾选：主色实底 + 白对勾；
//   - 半选 indeterminate：**白底 + 主色框 + 主色横杠**（不是主色实底）；
//   - 禁用：底/记号整体转 disabled 专色。
func (b *Checkbox) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	st := ui.StyleOf(th, ui.VariantOutline, ui.SizeSmall, b.states(b.Disabled))
	var boxFill, markColor, borderColor geom.Color
	switch {
	case b.Indeterminate && b.Disabled:
		boxFill, markColor, borderColor = th.Color.SurfaceHover, th.Color.BorderStrong, th.Color.Border
	case b.Indeterminate:
		boxFill, markColor, borderColor = th.Color.Surface, th.Color.Primary, th.Color.Primary
	case b.Checked && b.Disabled:
		boxFill, markColor, borderColor = th.Color.SurfaceHover, th.Color.BorderStrong, th.Color.Border
	case b.Checked:
		boxFill, markColor, borderColor = th.Color.Primary, th.Color.OnPrimary, th.Color.Primary
	default:
		boxFill, markColor, borderColor = th.Color.Surface, th.Color.Primary, st.Border
	}
	box := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(boxFill),
		Border: scene.Stroke(borderColor, 1),
		Radius: th.Radius.SM,
	}).SetSize2(16, 16)
	if b.Indeterminate {
		dash := scene.Box(scene.BoxSpec{Fill: scene.Solid(markColor)}).SetSize2(8, 2)
		dash.SetPositionKeepSize(true).SetStackPosition(geom.Pt(4, 7))
		box.Add(dash)
	} else if b.Checked {
		// 自绘白对勾（两条线段）：Icon 缩放到 12px 后笔画只剩 1px 看不清。
		check := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
			p := render.NewPath()
			p.MoveTo(2.5, 8).LineTo(6.5, 12).LineTo(13.5, 3.5)
			cv.StrokePath(p, render.Stroke(markColor, 2).WithCap(render.CapRound).WithJoin(render.JoinRound))
		}).SetSize2(16, 16)
		check.SetPositionKeepSize(true).SetStackPosition(geom.Pt(0, 0))
		box.Add(check)
	}
	box.SetOpacity(st.Opacity)

	// Label 文字必须渲染（修复：Label 字段此前从未画出来）；整行可点。
	root := scene.Flex(scene.Horizontal, 8).SetCrossAlign(scene.CrossCenter)
	root.Add(box)
	if b.Label != "" {
		lc := th.Color.Text
		if b.Disabled {
			lc = th.Color.TextFaint
		}
		root.Add((&Label{Text: b.Label, Size: th.Font.SizeBase, Color: lc}).Build(ctx))
	}
	root.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "checkbox", Cursor: scene.CursorPointer})

	toggle := func() {
		if b.Disabled {
			return
		}
		b.Checked = !b.Checked
		b.MarkDirty()
		if b.OnChange != nil {
			b.OnChange(b.Checked)
		}
	}
	h := b.tapHandler(ctx, b.Disabled, toggle)
	h.Key = func(e *input.KeyEvent) bool {
		if e.Code == input.CodeSpace && e.Kind != input.KeyUp && !b.Disabled {
			toggle()
			return true
		}
		return ctx.Focus != nil && ctx.Focus.HandleKey(e)
	}
	b.attach(ctx, root, h, !b.Disabled)
	return root
}

// Radio 是单选钮（同一组的互斥由调用方的 OnChange 负责）。
type Radio struct {
	controlBase
	Checked  bool
	Disabled bool
	Label    string
	OnChange func(bool)
}

// NewRadio 创建单选钮。
func NewRadio(label string, checked bool) *Radio {
	return &Radio{Label: label, Checked: checked}
}

// Build 实现 ui.Widget。
func (b *Radio) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	st := ui.StyleOf(th, ui.VariantOutline, ui.SizeSmall, b.states(b.Disabled))
	// 选中/悬停时圆点与外框转主色（antd 视觉）。
	border, dotColor := st.Border, th.Color.Primary
	if b.Checked {
		border = th.Color.Primary
	}
	if b.Disabled {
		border, dotColor = th.Color.Border, th.Color.BorderStrong
	}
	n := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(border, 1),
		Radius: 8,
	}).SetSize2(16, 16)
	if b.Checked {
		dot := scene.Box(scene.BoxSpec{Fill: scene.Solid(dotColor), Radius: 4}).SetSize2(8, 8)
		dot.SetPositionKeepSize(true).SetStackPosition(geom.Pt(4, 4))
		n.Add(dot)
	}
	n.SetOpacity(st.Opacity)

	// Label 文字（修复同 Checkbox 的丢字 bug）；整行可点。
	root := scene.Flex(scene.Horizontal, 8).SetCrossAlign(scene.CrossCenter)
	root.Add(n)
	if b.Label != "" {
		lc := th.Color.Text
		if b.Disabled {
			lc = th.Color.TextFaint
		}
		root.Add((&Label{Text: b.Label, Size: th.Font.SizeBase, Color: lc}).Build(ctx))
	}
	root.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "radio", Cursor: scene.CursorPointer})

	selectIt := func() {
		if b.Disabled || b.Checked {
			return
		}
		b.Checked = true
		b.MarkDirty()
		if b.OnChange != nil {
			b.OnChange(true)
		}
	}
	h := b.tapHandler(ctx, b.Disabled, selectIt)
	b.attach(ctx, root, h, !b.Disabled)
	return root
}

// Switch 是开关（M8 对齐 antd：checkedChildren/unCheckedChildren 文案、
// loading 态、small 尺寸）。
type Switch struct {
	controlBase
	On                bool
	CheckedChildren   string // 开态文案（显示在滑块外侧）
	UnCheckedChildren string // 关态文案
	Loading           bool   // 加载中：旋转指示并忽略点击
	Small             bool   // antd size="small"（24×12）
	Disabled          bool
	OnChange          func(bool)
}

// NewSwitch 创建开关。
func NewSwitch(on bool) *Switch { return &Switch{On: on} }

// Build 实现 ui.Widget。
func (b *Switch) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	st := ui.StyleOf(th, ui.VariantSoft, ui.SizeSmall, b.states(b.Disabled || b.Loading))
	trackW, trackH, knob := float32(36), float32(20), float32(16)
	if b.Small {
		trackW, trackH, knob = 24, 12, 8
	}
	track := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(bgOf(th, st, b.On)),
		Radius: trackH / 2,
	}).SetSize2(trackW, trackH)

	// 文案：显示在滑块对侧（antd checkedChildren/unCheckedChildren）。
	labelText := ""
	if b.On && b.CheckedChildren != "" {
		labelText = b.CheckedChildren
	} else if !b.On && b.UnCheckedChildren != "" {
		labelText = b.UnCheckedChildren
	}
	if labelText != "" && !b.Small {
		font := text.DefaultFont(11)
		lay := text.LayoutText(labelText, font, text.Options{})
		t := scene.Text(scene.TextSpec{Src: labelText, Layout: lay, Color: th.Color.OnPrimary})
		tx := float32(4)
		if b.On {
			tx = 2 // 滑块在右，文字靠左
		} else {
			tx = trackW - 2 - lay.Size.W - 2 // 滑块在左，文字靠右
		}
		t.SetPositionKeepSize(true).SetStackPosition(geom.Pt(tx, (trackH-lay.Size.H)/2))
		track.Add(t)
	}

	knobX := float32(2)
	if b.On {
		knobX = trackW - knob - 2
	}
	if b.Loading {
		ic := (&Icon{Name: "loading", Size: knob - 4, Color: th.Color.Primary, Spin: true}).Build(ctx)
		ic.SetPositionKeepSize(true).SetStackPosition(geom.Pt(knobX+2, 2))
		track.Add(ic)
		if ctx.Scheduler != nil {
			ctx.Scheduler.RequestFrameIn(33 * time.Millisecond)
		}
	} else {
		k := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface), Radius: knob / 2}).SetSize2(knob, knob)
		k.SetPositionKeepSize(true).SetStackPosition(geom.Pt(knobX, (trackH-knob)/2))
		track.Add(k)
	}
	track.SetOpacity(st.Opacity)
	track.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "switch", Cursor: scene.CursorPointer})

	toggle := func() {
		if b.Disabled || b.Loading {
			return
		}
		b.On = !b.On
		b.MarkDirty()
		if b.OnChange != nil {
			b.OnChange(b.On)
		}
	}
	h := b.tapHandler(ctx, b.Disabled || b.Loading, toggle)
	b.attach(ctx, track, h, !b.Disabled)
	return track
}

// bgOf 根据布尔状态取语义色（不用组件自己的色值）。
func bgOf(th *theme.Data, st ui.Style, on bool) geom.Color {
	if !on {
		return th.Color.SurfaceSunken
	}
	if st.Background == (geom.Color{}) {
		return th.Color.Primary
	}
	return st.Background
}
