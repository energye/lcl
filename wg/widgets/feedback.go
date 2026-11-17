package widgets

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── 反馈浮层：Modal / Drawer / Notification / Popconfirm ──────────
//
// 它们都是**控制器**而非静态 Widget：调用 Show 时向 overlay.Host 登记
// 一条 Entry，浮层的定位/关闭/z 序/遮罩/焦点域全部交给浮层体系。
// 内容节点用 BuildContext 的 theme 构建，保持零 lcl。

// feedbackHost 保存显示所需的 overlay 上下文。
type feedbackHost struct {
	host *overlay.Host
	th   *theme.Data
}

func captureHost(ctx *ui.BuildContext) feedbackHost {
	return feedbackHost{host: ctx.Overlay, th: ctx.ThemeOr()}
}

// ── Modal ─────────────────────────────────────────────────────────

// Modal 是模态对话框。Content 为自定义内容，按钮为确定/取消。
type Modal struct {
	Title     string
	Content   ui.Widget
	Confirm   string
	Cancel    string
	Width     float32
	OnConfirm func()
	OnCancel  func()

	handle overlay.Handle
	fb     feedbackHost
}

// NewModal 创建模态框。
func NewModal(title string, content ui.Widget) *Modal {
	return &Modal{Title: title, Content: content, Confirm: "确定", Cancel: "取消", Width: 400}
}

// Show 显示模态框。无 overlay 宿主时安全 no-op。
func (m *Modal) Show(ctx *ui.BuildContext) {
	if ctx == nil || ctx.Overlay == nil || m.handle != 0 {
		return
	}
	m.fb = captureHost(ctx)
	th := m.fb.th
	title, confirm, cancel := m.Title, m.Confirm, m.Cancel
	content := m.Content
	w := m.Width
	if w <= 0 {
		w = 400
	}
	onConfirm, onCancel := m.OnConfirm, m.OnCancel
	m.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandModal,
		Anchor:    overlay.Anchor{Kind: overlay.AnchorNone},
		Placement: overlay.PlacementCenter,
		Size:      geom.Size{W: w, H: 220},
		Modal:     true,
		TrapFocus: true,
		Dismiss:   overlay.DismissEsc,
		Tag:       "modal",
		Build: func() *scene.Node {
			col := scene.Flex(scene.Vertical, th.Space.MD).SetWidth(w)
			col.SetBox(scene.BoxSpec{
				Fill:   scene.Solid(th.Color.Surface),
				Border: scene.Stroke(th.Color.BorderStrong, 1),
				Radius: th.Radius.LG,
			})
			col.SetPadding(geom.InLTRB(th.Space.LG, th.Space.LG, th.Space.LG, th.Space.LG))
			// 标题
			if title != "" {
				tl := text.LayoutText(title, text.DefaultFont(th.Font.SizeLG), text.Options{})
				col.Add(scene.Text(scene.TextSpec{Src: title, Layout: tl, Color: th.Color.Text}))
			}
			if content != nil {
				col.Add(buildChild(ctx, content))
			}
			// 按钮行
			row := scene.Flex(scene.Horizontal, th.Space.SM)
			if cancel != "" {
				row.Add(modalButton(ctx, cancel, false, th, func() {
					m.close()
					if onCancel != nil {
						onCancel()
					}
				}))
			}
			if confirm != "" {
				row.Add(modalButton(ctx, confirm, true, th, func() {
					m.close()
					if onConfirm != nil {
						onConfirm()
					}
				}))
			}
			col.Add(row)
			return col
		},
	})
}

func (m *Modal) close() {
	if m.handle != 0 && m.fb.host != nil {
		m.fb.host.Close(m.handle)
	}
	m.handle = 0
}

// Close 关闭模态框。
func (m *Modal) Close() { m.close() }

// IsOpen 返回是否正在显示。
func (m *Modal) IsOpen() bool { return m.handle != 0 }

// modalButton 构造模态框里的一个按钮（内部函数）。
func modalButton(ctx *ui.BuildContext, label string, primary bool, th *theme.Data, onClick func()) *scene.Node {
	w, h := float32(80), float32(34)
	fill, fg := geom.Color{}, th.Color.Text
	if primary {
		fill, fg = th.Color.Primary, th.Color.OnPrimary
	}
	n := scene.Box(scene.BoxSpec{
		Fill: scene.Solid(fill), Border: scene.Stroke(th.Color.Border, 1), Radius: th.Radius.Base,
	}).SetSize2(w, h)
	l := text.LayoutText(label, text.DefaultFont(th.Font.SizeSM), text.Options{})
	tv := scene.Text(scene.TextSpec{Src: label, Layout: l, Color: fg})
	tv.SetMargin(geom.InLTRB((w-l.Size.W)/2, (h-l.Size.H)/2, 0, 0))
	n.Add(tv)
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "modal-button", Cursor: scene.CursorPointer})
	g := &ui.GestureDetector{OnTap: func(geom.Point) { onClick() }}
	if ctx.Input != nil {
		ctx.Input.Set(n, g.Handler())
	}
	// g 由浮层内容节点的处理器闭包持有，浮层销毁前一直可达。
	return n
}

// ── Drawer ────────────────────────────────────────────────────────

// DrawerEdge 是抽屉弹出的边。
type DrawerEdge uint8

const (
	DrawerLeft DrawerEdge = iota
	DrawerRight
	DrawerTop
	DrawerBottom
)

// Drawer 是抽屉浮层。
type Drawer struct {
	Title   string
	Content ui.Widget
	Edge    DrawerEdge
	Width   float32
	OnClose func()

	handle overlay.Handle
	fb     feedbackHost
}

// NewDrawer 创建抽屉。
func NewDrawer(title string, content ui.Widget) *Drawer {
	return &Drawer{Title: title, Content: content, Edge: DrawerLeft, Width: 320}
}

// Show 显示抽屉。
func (d *Drawer) Show(ctx *ui.BuildContext) {
	if ctx == nil || ctx.Overlay == nil || d.handle != 0 {
		return
	}
	d.fb = captureHost(ctx)
	th := d.fb.th
	placement := map[DrawerEdge]overlay.Placement{
		DrawerLeft: overlay.PlacementLeftStart, DrawerRight: overlay.PlacementRightStart,
		DrawerTop: overlay.PlacementTopStart, DrawerBottom: overlay.PlacementBottomStart,
	}[d.Edge]
	w := d.Width
	if w <= 0 {
		w = 320
	}
	title, content := d.Title, d.Content
	onClose := d.OnClose
	d.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandModal,
		Anchor:    overlay.Anchor{Kind: overlay.AnchorNone},
		Placement: placement,
		Size:      geom.Size{W: w, H: 400},
		Modal:     true,
		TrapFocus: true,
		Dismiss:   overlay.DismissEsc | overlay.DismissOutsideClick,
		Tag:       "drawer",
		Build: func() *scene.Node {
			col := scene.Flex(scene.Vertical, th.Space.MD).SetWidth(w).SetHeight(400)
			col.SetBox(scene.BoxSpec{
				Fill:   scene.Solid(th.Color.Surface),
				Border: scene.Stroke(th.Color.Border, 1),
			})
			col.SetPadding(geom.InLTRB(th.Space.LG, th.Space.LG, th.Space.LG, th.Space.LG))
			if title != "" {
				tl := text.LayoutText(title, text.DefaultFont(th.Font.SizeLG), text.Options{})
				col.Add(scene.Text(scene.TextSpec{Src: title, Layout: tl, Color: th.Color.Text}))
			}
			if content != nil {
				col.Add(buildChild(ctx, content))
			}
			closeBtn := modalButton(ctx, "关闭", false, th, func() {
				d.close()
				if onClose != nil {
					onClose()
				}
			})
			closeBtn.SetStackPosition(geom.Pt(w-90, 8)).SetPositionKeepSize(true)
			col.Add(closeBtn)
			return col
		},
	})
}

func (d *Drawer) close() {
	if d.handle != 0 && d.fb.host != nil {
		d.fb.host.Close(d.handle)
	}
	d.handle = 0
}

// Close 关闭抽屉。
func (d *Drawer) Close() { d.close() }

// IsOpen 返回是否正在显示。
func (d *Drawer) IsOpen() bool { return d.handle != 0 }

// ── Notification ─────────────────────────────────────────────────

// Notification 是右上角通知（超时自动消失）。
type Notification struct {
	Title    string
	Body     string
	Duration time.Duration
	handle   overlay.Handle
	fb       feedbackHost
}

// NewNotification 创建通知。
func NewNotification(title, body string) *Notification {
	return &Notification{Title: title, Body: body, Duration: 3 * time.Second}
}

// Show 显示通知。
func (n *Notification) Show(ctx *ui.BuildContext) {
	if ctx == nil || ctx.Overlay == nil || n.handle != 0 {
		return
	}
	n.fb = captureHost(ctx)
	th := n.fb.th
	title, body, dur := n.Title, n.Body, n.Duration
	if dur <= 0 {
		dur = 3 * time.Second
	}
	n.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandToast,
		Anchor:    overlay.Anchor{Kind: overlay.AnchorNone},
		Placement: overlay.PlacementTopEnd,
		Size:      geom.Size{W: 280, H: 72},
		Dismiss:   overlay.DismissTimeout | overlay.DismissOutsideClick,
		Duration:  dur,
		Tag:       "notification",
		Build: func() *scene.Node {
			col := scene.Flex(scene.Vertical, 4).SetWidth(280).SetHeight(72)
			col.SetBox(scene.BoxSpec{
				Fill:   scene.Solid(th.Color.Surface),
				Border: scene.Stroke(th.Color.Border, 1), Radius: th.Radius.Base,
			})
			col.SetPadding(geom.InLTRB(12, 10, 12, 10))
			if title != "" {
				tl := text.LayoutText(title, text.DefaultFont(th.Font.SizeBase), text.Options{})
				col.Add(scene.Text(scene.TextSpec{Src: title, Layout: tl, Color: th.Color.Text}))
			}
			if body != "" {
				bl := text.LayoutText(body, text.DefaultFont(th.Font.SizeSM), text.Options{})
				col.Add(scene.Text(scene.TextSpec{Src: body, Layout: bl, Color: th.Color.TextMuted}))
			}
			return col
		},
	})
}

// Close 关闭通知。
func (n *Notification) Close() {
	if n.handle != 0 && n.fb.host != nil {
		n.fb.host.Close(n.handle)
	}
	n.handle = 0
}

// IsOpen 返回是否正在显示。
func (n *Notification) IsOpen() bool { return n.handle != 0 }

// ── Popconfirm ───────────────────────────────────────────────────

// Popconfirm 是锚定在控件旁的确认气泡。
type Popconfirm struct {
	controlBase
	Title     string
	Content   ui.Widget
	Child     ui.Widget
	Placement overlay.Placement
	Confirm   string
	Cancel    string
	OnConfirm func()
	OnCancel  func()

	handle overlay.Handle
	fb     feedbackHost
	taps   []*ui.GestureDetector
}

// NewPopconfirm 创建确认气泡。
func NewPopconfirm(title string, child ui.Widget) *Popconfirm {
	return &Popconfirm{
		Title: title, Child: child, Placement: overlay.PlacementTop,
		Confirm: "确定", Cancel: "取消",
	}
}

// Build 实现 ui.Widget：Child 点击时弹出气泡。
func (p *Popconfirm) Build(ctx *ui.BuildContext) *scene.Node {
	if p.Child == nil {
		return scene.Box(scene.BoxSpec{})
	}
	anchor := buildChild(ctx, p.Child)
	g := &ui.GestureDetector{OnTap: func(geom.Point) { p.show(ctx, anchor) }}
	p.taps = append(p.taps[:0], g)
	if ctx.Input != nil {
		ctx.Input.Set(anchor, g.Handler())
	}
	p.node = anchor
	return anchor
}

func (p *Popconfirm) show(ctx *ui.BuildContext, anchor *scene.Node) {
	if ctx == nil || ctx.Overlay == nil || p.handle != 0 {
		return
	}
	p.fb = captureHost(ctx)
	th := p.fb.th
	title, confirm, cancel := p.Title, p.Confirm, p.Cancel
	content := p.Content
	onConfirm, onCancel := p.OnConfirm, p.OnCancel
	p.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: p.Placement,
		Size:      geom.Size{W: 240, H: 110},
		Dismiss:   overlay.DismissOutsideClick | overlay.DismissEsc,
		Tag:       "popconfirm",
		Build: func() *scene.Node {
			col := scene.Flex(scene.Vertical, th.Space.SM).SetWidth(240)
			col.SetBox(scene.BoxSpec{
				Fill:   scene.Solid(th.Color.Surface),
				Border: scene.Stroke(th.Color.Border, 1), Radius: th.Radius.Base,
			})
			col.SetPadding(geom.InLTRB(12, 10, 12, 10))
			if title != "" {
				tl := text.LayoutText(title, text.DefaultFont(th.Font.SizeBase), text.Options{})
				col.Add(scene.Text(scene.TextSpec{Src: title, Layout: tl, Color: th.Color.Text}))
			}
			if content != nil {
				col.Add(buildChild(ctx, content))
			}
			row := scene.Flex(scene.Horizontal, th.Space.SM)
			if cancel != "" {
				row.Add(modalButton(ctx, cancel, false, th, func() {
					p.close()
					if onCancel != nil {
						onCancel()
					}
				}))
			}
			if confirm != "" {
				row.Add(modalButton(ctx, confirm, true, th, func() {
					p.close()
					if onConfirm != nil {
						onConfirm()
					}
				}))
			}
			col.Add(row)
			return col
		},
	})
}

func (p *Popconfirm) close() {
	if p.handle != 0 && p.fb.host != nil {
		p.fb.host.Close(p.handle)
	}
	p.handle = 0
}

// Close 关闭气泡。
func (p *Popconfirm) Close() { p.close() }

// IsOpen 返回是否正在显示。
func (p *Popconfirm) IsOpen() bool { return p.handle != 0 }
