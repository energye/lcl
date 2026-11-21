package widgets

// ── Dropdown（M11 导航）────────────────────────────────────────────
//
// antd：<Dropdown overlay menu trigger="hover|click">
// 已对齐：悬停/点击触发（Enter/Leave 管浮层，同 Tooltip 机制——
// Dismiss(0) 不吞后续点击）、BandPopup 层带+PlacementBottom、
// 菜单项列表（复用 MenuItem{Key,Label,Icon}）+ OnClick。
// 未实现（如实标注）：点击触发模式、禁用项、分割线。

import (
	"time"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// Dropdown 是下拉菜单（触发元素 + BandPopup 浮层菜单项列表）。
type Dropdown struct {
	// Child 触发元素（Button 等；Enter 显示浮层/Leave 隐藏）。
	Child ui.Widget
	// Items 菜单项（复用 MenuItem{Key,Label,Icon}——Menu 同结构）。
	Items []*MenuItem
	OnClick func(key string)
	handle    overlay.Handle
	fb        feedbackHost
	taps      []*ui.GestureDetector
	hideTimer *time.Timer // Child Leave 延迟隐藏（鼠标移入浮层则取消）
	childNode *scene.Node // 最近构建的触发元素节点（程序化 Show 锚）
}

// NewDropdown 创建下拉菜单。
func NewDropdown(child ui.Widget, items ...*MenuItem) *Dropdown {
	return &Dropdown{Child: child, Items: items}
}

// Build 实现 ui.Widget。
func (d *Dropdown) Build(ctx *ui.BuildContext) *scene.Node {
	child := buildChild(ctx, d.Child)
	d.childNode = child
	if child == nil {
		return scene.Empty()
	}
	// 合并 Child 已注册的 handler（Tooltip 同款——不覆盖按钮点击）。
	h := input.Handler{}
	if ctx.Input != nil {
		if old, ok := ctx.Input.Handler(child); ok {
			h = old
		}
	}
	show := func() {
		if d.hideTimer != nil {
			d.hideTimer.Stop()
			d.hideTimer = nil
		}
		d.show(ctx, child)
	}
	hide := func() {
		// 延迟 150ms 隐藏：鼠标从 Child 移向浮层的间隙 Leave 先触发——
		// 立即 hide 会销毁浮层导致菜单项点击落空（实测）。鼠标在延迟窗
		// 内进入浮层（菜单项 Enter 取消定时器）则保持显示。
		if d.hideTimer != nil {
			d.hideTimer.Stop()
		}
		d.hideTimer = time.AfterFunc(150*time.Millisecond, func() {
			d.hide(ctx)
		})
	}
	oldEnter, oldLeave := h.Enter, h.Leave
	h.Enter = func(e *input.PointerEvent) {
		if oldEnter != nil {
			oldEnter(e)
		}
		show()
	}
	h.Leave = func(e *input.PointerEvent) {
		if oldLeave != nil {
			oldLeave(e)
		}
		hide()
	}
	if ctx.Input != nil {
		ctx.Input.Set(child, h)
	}
	d.taps = d.taps[:0]
	_ = d.taps
	return child
}

// show 显示下拉浮层（BandPopup+PlacementBottom，锚=Child 节点）。
func (d *Dropdown) show(ctx *ui.BuildContext, anchor *scene.Node) {
	if ctx.Overlay == nil || d.handle != 0 || len(d.Items) == 0 {
		return
	}
	d.fb = captureHost(ctx)
	th := d.fb.th
	items := d.Items
	onClick := d.OnClick

	// 菜单项列表：白底圆角浮层+阴影，每行图标+文本（Menu 纵向项同款形态）。
	col := scene.Flex(scene.Vertical, 2)
	col.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Radius: th.Radius.Base,
		Shadow: func() *scene.ShadowSpec {
			if len(th.Elevation) > 1 {
				e := th.Elevation[1]
				return &scene.ShadowSpec{Dx: e.Dx, Dy: e.Dy, Blur: e.Blur, Color: e.Color}
			}
			return nil
		}(),
	})
	col.SetPadding(geom.InLTRB(4, 4, 4, 4))
	maxW := float32(0)
	for _, it := range items {
		font := text.DefaultFont(th.Font.SizeBase)
		layout := text.LayoutText(it.Label, font, text.Options{})
		w := layout.Size.W + 56 // 图标+内边距
		if w > maxW {
			maxW = w
		}
	}
	for _, it := range items {
		it2 := it
		row := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
		row.SetHeight(32).SetWidth(maxW)
		row.SetPadding(geom.InLTRB(10, 0, 10, 0))
		if it.Icon != "" {
			ic := (&Icon{Name: it.Icon, Size: 14, Color: th.Color.Text}).Build(ctx)
			row.Add(ic)
		}
		font := text.DefaultFont(th.Font.SizeBase)
		layout := text.LayoutText(it.Label, font, text.Options{})
		row.Add(scene.Text(scene.TextSpec{Src: it.Label, Layout: layout, Color: th.Color.Text}))
		row.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "dropdown-item", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			// **down 即执行**（不走 GestureDetector tap 配对）：浮层菜单项
			// 的 tap 配对在"Child Leave→hideTimer→重建/隐藏"时序下实测
			// 永不触发（DD-ONTAP trace 零输出）——菜单项就是点选语义，
			// down 直接触发回调并关浮层，最可靠。
			g := &ui.GestureDetector{}
			h2 := input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				if !ev.Down() {
					return false
				}
				if d.hideTimer != nil {
					d.hideTimer.Stop()
					d.hideTimer = nil
				}
				d.hide(ctx)
				if onClick != nil {
					onClick(it2.Key)
				}
				return true
			}, Enter: func(e *input.PointerEvent) {
				if d.hideTimer != nil {
					d.hideTimer.Stop()
					d.hideTimer = nil
				}
			}}
			_ = g
			d.taps = append(d.taps, g)
			ctx.Input.Set(row, h2)
		}
		col.Add(row)
	}
	d.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: overlay.PlacementBottomStart,
		// 高度自适应（[引擎] 浮层高度自适应）。
		Size:     geom.Size{W: maxW, H: 0},
		Dismiss:  overlay.Dismiss(0),
		Tag:      "dropdown",
		OnClose:  func() { d.handle = 0 },
		Build:    func() *scene.Node { return col },
	})
}

// Show 程序化显示下拉浮层（anchor 取 Child 最近构建节点；测试/自动化用）。
func (d *Dropdown) Show(ctx *ui.BuildContext) {
	child := d.childNode
	if child == nil {
		return
	}
	d.show(ctx, child)
}

// hide 隐藏下拉浮层。
func (d *Dropdown) hide(ctx *ui.BuildContext) {
	if ctx.Overlay == nil || d.handle == 0 {
		return
	}
	ctx.Overlay.Close(d.handle)
}

// Close 立即关闭浮层。
func (d *Dropdown) Close() { d.hide(nil) }

// IsOpen 返回浮层是否打开。
func (d *Dropdown) IsOpen() bool { return d.handle != 0 }
