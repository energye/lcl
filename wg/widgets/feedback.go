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

// ModalKind 是模态框预设类型（antd Modal.info/success/warning/error）。
type ModalKind uint8

const (
	// ModalNormal 普通确认对话框（确定/取消，无图标）。
	ModalNormal ModalKind = iota
	// ModalInfo 信息提示（蓝圆+i）。
	ModalInfo
	// ModalSuccess 成功（绿圆+勾）。
	ModalSuccess
	// ModalWarning 警告（黄圆+叹号）。
	ModalWarning
	// ModalError 错误（红圆+×）。
	ModalError
)

// Modal 是模态对话框。Content 为自定义内容，按钮为确定/取消。
// Kind 非 ModalNormal 时渲染预设确认形态：状态图标圆底 + 标题/内容 +
// 单"知道了"按钮（footer 按钮为 Button 组件）。
type Modal struct {
	Title     string
	Content   ui.Widget
	Confirm   string
	Cancel    string
	Width     float32
	OnConfirm func()
	OnCancel  func()
	// Kind 预设类型（默认 ModalNormal 保持普通形态）。
	Kind ModalKind
	// OKText 预设按钮文字（默认"知道了"）。
	OKText string

	handle overlay.Handle
	fb     feedbackHost
}

// NewModal 创建模态框。
func NewModal(title string, content ui.Widget) *Modal {
	return &Modal{Title: title, Content: content, Confirm: "确定", Cancel: "取消", Width: 400}
}

// NewModalPreset 创建预设确认框（antd Modal.info/success/warning/error）：
// 状态图标+标题+内容+单 OK 按钮（组合 Modal/Button/Icon，零新绘制通道）。
func NewModalPreset(kind ModalKind, title string, content ui.Widget, onOk func()) *Modal {
	return &Modal{Title: title, Content: content, Kind: kind, Width: 416, OnConfirm: onOk}
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
	preset := m.Kind != ModalNormal
	if preset {
		// antd 预设确认框：单 OK 按钮（OKText > "知道了"），无 Cancel，
		// 无标题分隔线。
		cancel = ""
		confirm = m.OKText
		if confirm == "" {
			confirm = "知道了"
		}
	}
	onConfirm, onCancel := m.OnConfirm, m.OnCancel
	m.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandModal,
		Anchor:    overlay.Anchor{Kind: overlay.AnchorNone},
		Placement: overlay.PlacementCenter,
		// 高度 0 = 由内容决定（Overlay 对 Center 放置取内容自然尺寸），
		// 固定 220 会在内容多时截断（antd 对话框高度自适应）。
		Size:      geom.Size{W: w},
		Modal:     true,
		TrapFocus: true,
		Dismiss:   overlay.DismissEsc,
		Tag:       "modal",
		OnClose:   func() { m.handle = 0 },
		Build: func() *scene.Node {
			col := scene.Flex(scene.Vertical, th.Space.MD).SetWidth(w)
			col.SetBox(scene.BoxSpec{
				Fill:   scene.Solid(th.Color.Surface),
				Border: scene.Stroke(th.Color.BorderStrong, 1),
				Radius: th.Radius.LG,
			})
			col.SetPadding(geom.InLTRB(th.Space.LG, th.Space.LG, th.Space.LG, th.Space.LG))
			if preset {
				// 预设确认形态：状态图标圆底 + 右列（标题+内容）。
				col.Add(presetBody(ctx, th, m.Kind, title, content))
			} else {
				// 标题 + 分隔线（antd 对话框标题区有 hairline 分隔）。
				if title != "" {
					tl := text.LayoutText(title, text.DefaultFont(th.Font.SizeLG), text.Options{})
					col.Add(scene.Text(scene.TextSpec{Src: title, Layout: tl, Color: th.Color.Text}))
					col.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).
						SetHeight(1).SetFlex(1, scene.FlexTight))
				}
				if content != nil {
					col.Add(buildChild(ctx, content))
				}
			}
			// 按钮行右对齐（antd footer 按钮靠右）。
			row := scene.Flex(scene.Horizontal, th.Space.SM)
			row.SetMainAlign(scene.MainEnd)
			row.SetCrossAlign(scene.CrossCenter)
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

// modalButton 构造模态框 footer 按钮：**复用 Button 组件**（Flutter 标准
// 哲学——AlertDialog.actions 就是调用方传 TextButton，框架不自绘按钮；
// 自绘盒钮只有 OnTap，hover/按压反馈全部缺失）。primary 走 TypePrimary，
// 其余 TypeDefault。
func modalButton(ctx *ui.BuildContext, label string, primary bool, th *theme.Data, onClick func()) *scene.Node {
	b := &Button{Text: label, OnClick: onClick}
	if primary {
		b.Type = TypePrimary
	}
	return b.Build(ctx)
}

// presetBody 构建预设确认框内容区：状态图标圆底（Stack：语义色 32×32 圆
// Box + 白图形 Icon 居中——antd 实底圆白图形）+ 右列（标题+内容）。
// 图标圆底加 modal-preset-icon 被动定位锚供断言采样。
func presetBody(ctx *ui.BuildContext, th *theme.Data, kind ModalKind, title string, content ui.Widget) *scene.Node {
	fill := th.Color.Primary
	name := "info-body"
	switch kind {
	case ModalSuccess:
		fill, name = th.Color.Success, "check"
	case ModalWarning:
		fill, name = th.Color.Warning, "exclamation"
	case ModalError:
		fill, name = th.Color.Danger, "close"
	}
	row := scene.Flex(scene.Horizontal, th.Space.MD).SetCrossAlign(scene.CrossStart)
	disc := scene.Stack().SetSize2(32, 32)
	disc.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: th.Radius.Pill}).SetSize2(32, 32))
	ic := (&Icon{Name: name, Size: 18, Color: th.Color.OnPrimary}).Build(ctx)
	ic.SetPositionKeepSize(true)
	ic.SetStackPosition(geom.Pt(7, 7)) // (32-18)/2 居中
	disc.Add(ic)
	disc.SetHit(scene.HitSpec{Label: "modal-preset-icon"})
	row.Add(disc)
	colR := scene.Flex(scene.Vertical, th.Space.SM)
	if title != "" {
		tl := text.LayoutText(title, text.DefaultFont(th.Font.SizeLG), text.Options{})
		colR.Add(scene.Text(scene.TextSpec{Src: title, Layout: tl, Color: th.Color.Text}))
	}
	if content != nil {
		colR.Add(buildChild(ctx, content))
	}
	row.Add(colR)
	return row
}

// ── Alert ─────────────────────────────────────────────────────────

// AlertType 是提示语义（antd Alert.type）。
type AlertType uint8

const (
	// AlertInfo 信息（主色）。
	AlertInfo AlertType = iota
	// AlertSuccess 成功（绿）。
	AlertSuccess
	// AlertWarning 警告（黄）。
	AlertWarning
	// AlertError 错误（红）。
	AlertError
)

// alertPalette 折算 Alert 类型的底色/图标色/图标名。
func alertPalette(th *theme.Data, typ AlertType) (bg, fg geom.Color, icon string) {
	switch typ {
	case AlertSuccess:
		return th.Color.SuccessBg, th.Color.Success, "check-circle"
	case AlertWarning:
		return th.Color.WarningBg, th.Color.Warning, "warning-circle"
	case AlertError:
		return th.Color.DangerBg, th.Color.Danger, "close-circle"
	default:
		return th.Color.InfoBg, th.Color.Primary, "info-circle"
	}
}

// Alert 是静态提示（antd Alert）：浅色语义底 + 图标（showIcon）+
// 标题/描述 + 可关闭（closable）。全部内容组合复用 Icon/Label 组件。
type Alert struct {
	controlBase
	Type AlertType
	// Title 主文案（加粗语义——用默认字重+正色即可，antd title 无特殊字重）。
	Title string
	// Description 副文案（TextMuted 弱化）。
	Description string
	// ShowIcon 显示左侧类型图标（antd showIcon）。
	ShowIcon bool
	// Closable 可关闭：右上 × 点击触发 OnClose（antd closable）。
	Closable bool
	OnClose  func()
	Width    float32

	taps []*ui.GestureDetector // × 手势实例池（跨重建持有）
}

// Build 实现 ui.Widget。
func (a *Alert) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	bg, fg, icon := alertPalette(th, a.Type)
	w := a.Width
	if w <= 0 {
		w = 420
	}
	col := scene.Flex(scene.Vertical, th.Space.XS).SetWidth(w)
	col.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(bg),
		Border: scene.Stroke(fg.WithAlpha(60), 1),
		Radius: th.Radius.Base,
	})
	col.SetPadding(geom.InLTRB(th.Space.MD, th.Space.SM, th.Space.MD, th.Space.SM))

	// 主行：图标（showIcon）+ 标题 + 右上 ×（closable）。
	mainRow := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	if a.ShowIcon {
		ic := (&Icon{Name: icon, Size: 16, Color: fg}).Build(ctx)
		mainRow.Add(ic)
	}
	if a.Title != "" {
		t := (&Label{Text: a.Title, Size: th.Font.SizeBase, Color: fg}).Build(ctx)
		mainRow.Add(t)
	}
	// × 用 MainEnd 推到行尾：加一个弹性占位把 × 推右。
	if a.Closable {
		sp := scene.Box(scene.BoxSpec{}).SetFlex(1, scene.FlexTight)
		mainRow.Add(sp)
		ix := (&Icon{Name: "close", Size: 12, Color: th.Color.TextFaint}).Build(ctx)
		ix.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "alert-close", Cursor: scene.CursorPointer})
		mainRow.Add(ix)
		// × 手势：标准 attach 接线（Input.Set + InvalidateRoot 帧请求 +
		// ctlKey 钉身份——closable 移除后的重绘链依赖它）。
		if ctx.Input != nil {
			g := &ui.GestureDetector{OnTap: func(geom.Point) {
				if a.OnClose != nil {
					a.OnClose()
				}
			}}
			a.taps = append(a.taps, g)
			a.attach(ctx, ix, g.Handler(), false)
		}
	}
	col.Add(mainRow)

	if a.Description != "" {
		d := (&Label{Text: a.Description, Size: th.Font.SizeSM, Color: th.Color.TextMuted}).Build(ctx)
		col.Add(d)
	}
	col.SetHit(scene.HitSpec{Label: "alert"})
	return col
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
	// 通栏尺寸（antd Drawer 全高/全宽）：左右抽屉高=视口高，上下抽屉宽=
	// 视口宽。ViewportRect 未就绪（零矩形）时退回默认 400。
	vp := d.fb.host.ViewportRect()
	vpH, vpW := float32(400), float32(400)
	if !vp.IsEmpty() {
		vpH, vpW = vp.H(), vp.W()
	}
	size := geom.Size{W: w, H: vpH}
	if d.Edge == DrawerTop || d.Edge == DrawerBottom {
		size = geom.Size{W: vpW, H: 400}
	}
	d.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandModal,
		Anchor:    overlay.Anchor{Kind: overlay.AnchorNone},
		Placement: placement,
		Size:      size,
		Modal:     true,
		TrapFocus: true,
		Dismiss:   overlay.DismissEsc | overlay.DismissOutsideClick,
		Tag:       "drawer",
		OnClose:   func() { d.handle = 0 },
		Build: func() *scene.Node {
			colH := size.H
			col := scene.Flex(scene.Vertical, th.Space.MD).SetWidth(size.W).SetHeight(colH)
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
			closeBtn.SetStackPosition(geom.Pt(size.W-90, 8)).SetPositionKeepSize(true)
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
		// 高度 0 = 按内容自适应（正文长不被截断，[引擎] 浮层高度自适应）。
		Size:      geom.Size{W: 280},
		Dismiss:   overlay.DismissTimeout | overlay.DismissOutsideClick,
		Duration:  dur,
		Tag:       "notification",
		OnClose:   func() { n.handle = 0 },
		Build: func() *scene.Node {
			col := scene.Flex(scene.Vertical, 4).SetWidth(280)
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
		OnClose:   func() { p.handle = 0 },
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
