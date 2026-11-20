// Package widgets 的 P2 展示型组件。
//
// 这一组都是"纯呈现"组件：读主题 token、把 props 折成 scene 节点，
// 不持有像素、不碰 lcl。需要交互的组件挂 Handler，图形（Spin 等）
// 走 CustomPaintNode + render.Canvas——这是 C17 允许的录制通道。
package widgets

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── Card ───────────────────────────────────────────────────────────

// Card 是一张带标题与可选底部操作区的卡片。
//
// 宽高由父级约束决定：放进 Flex 时给 flex，卡片自己只按内容增长。
type Card struct {
	Title   ui.Widget
	Content ui.Widget
	Footer  ui.Widget
	// Cover 封面区（antd Card.cover）：渲染在卡片最顶部、**通栏无内边距**
	// （封面图元直贴卡片边缘，与 antd cover 语义一致）。
	Cover   ui.Widget
	// Actions 操作区（antd Card actions）：渲染在卡片底部，顶部 hairline
	// 分隔、横排均分（MainSpaceEvenly）、无左右内边距。
	Actions  []ui.Widget
	Padding  geom.Insets
	Elevated bool
}

// NewCard 创建卡片。
func NewCard(content ui.Widget) *Card { return &Card{Content: content, Elevated: true} }

// Build 实现 ui.Widget。
func (c Card) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	pad := c.Padding
	if pad == (geom.Insets{}) {
		pad = geom.InLTRB(th.Space.LG, th.Space.LG, th.Space.LG, th.Space.LG)
	}
	spec := scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.LG,
	}
	if c.Elevated && len(th.Elevation) > 0 {
		e := th.Elevation[0]
		spec.Shadow = &scene.ShadowSpec{Dx: e.Dx, Dy: e.Dy, Blur: e.Blur, Color: e.Color}
	}
	col := scene.Flex(scene.Vertical, th.Space.MD)
	// 交叉轴拉伸：分隔线要通栏（宽度=卡片内容宽）——固定 1px 高的
	// Box 在 ShrinkWrap 列里没有宽度来源，FlexTight 主轴拉伸也不适用
	// （主轴是纵向）。CrossStretch 拉伸所有子节点宽度，文字节点不受影响。
	col.SetCrossAlign(scene.CrossStretch)
	// Cover 封面：直贴卡片顶缘（无内边距——负 margin 抵消，与 actions
	// 同款通栏手法）；卡片圆角上缘由 Card Box 的 Radius+Fill 承担，
	// Cover 内容自身若需圆角由调用方控制。
	if c.Cover != nil {
		cover := buildChild(ctx, c.Cover)
		cover.SetMargin(geom.InLTRB(-pad.Left, -pad.Top, -pad.Left, 0))
		cover.SetHit(scene.HitSpec{Label: "card-cover"}) // 被动定位锚
		col.Add(cover)
	}
	if c.Title != nil {
		col.Add(buildChild(ctx, c.Title))
		divider := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetHeight(1)
		// 被动定位锚：标题分隔线断言用（标题高度随内容变化，frac 采样
		// 带不可靠）。
		divider.SetHit(scene.HitSpec{Label: "card-divider"})
		col.Add(divider)
	}
	if c.Content != nil {
		col.Add(buildChild(ctx, c.Content))
	}
	if c.Footer != nil {
		col.Add(buildChild(ctx, c.Footer))
	}
	// Actions 操作区：顶部 hairline 分隔 + 横排均分（antd actions 形态——
	// 操作项直接贴卡片底缘，无左右内边距）。
	if len(c.Actions) > 0 {
		col.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)}).SetHeight(1))
		arow := scene.Flex(scene.Horizontal, 0).SetHeight(40).SetCrossAlign(scene.CrossCenter).SetMainAlign(scene.MainSpaceEvenly)
		for _, act := range c.Actions {
			arow.Add(buildChild(ctx, act))
		}
		// actions 区通栏：负 margin 抵消卡片内边距（直贴卡片底缘）。
		arow.SetMargin(geom.InLTRB(-pad.Left, 0, -pad.Left, 0))
		col.Add(arow)
	}
	n := scene.Box(spec, col)
	n.SetPadding(pad)
	// 定位锚（自动化测试选择器；DeferToChild 不参与命中，不影响交互）。
	n.SetHit(scene.HitSpec{Label: "card"})
	return n
}

// ── Badge / Tag ────────────────────────────────────────────────────

// Badge 是角标（小圆点或数字），常作为其它组件的装饰子组件。
type Badge struct {
	Text      string
	Count     int
	Dot       bool
	Color     geom.Color
	TextColor geom.Color
	// Status 语义色（antd Badge.status）：Success 绿 / Processing 主色 /
	// Error 红。显式 Color 优先于 Status。
	Status BadgeStatus
	// Child 被包裹组件（antd Badge 核心用法）：角标叠加在其右上角
	// （Stack 布局，角标偏移 (childW-6, -6) 出血）。nil 时独立渲染。
	Child ui.Widget
	// OverflowCount 显示上限（antd overflowCount，默认 99）：超过显示
	// "N+"。Count>上限时钳制。
	OverflowCount int
}

// BadgeStatus 是徽标语义（antd Badge.status）。
type BadgeStatus uint8

const (
	// BadgeStatusDefault 无语义（Danger 红底，与 Count/Text 一致）。
	BadgeStatusDefault BadgeStatus = iota
	// BadgeStatusSuccess 成功绿。
	BadgeStatusSuccess
	// BadgeStatusProcessing 处理中主色。
	BadgeStatusProcessing
	// BadgeStatusError 错误红（与默认同色，显式声明用）。
	BadgeStatusError
)

// statusFill 折算 Status 语义底色（Color 显式时优先）。
func (b Badge) statusFill(th *theme.Data) geom.Color {
	if b.Color != (geom.Color{}) {
		return b.Color
	}
	switch b.Status {
	case BadgeStatusSuccess:
		return th.Color.Success
	case BadgeStatusProcessing:
		return th.Color.Primary
	case BadgeStatusError:
		return th.Color.Danger
	}
	return th.Color.Danger // Default：红底与 Count/Text 默认一致
}

// Build 实现 ui.Widget。
func (b Badge) Build(ctx *ui.BuildContext) *scene.Node {
	badge := b.badgeNode(ctx)
	if b.Child == nil {
		return badge
	}
	// 包裹形态（antd Badge 核心用法）：Stack = Child + 角标叠右上角。
	child := buildChild(ctx, b.Child)
	cw, ch := child.Size().W, child.Size().H
	if cw < 16 {
		cw = 16
	}
	if ch < 16 {
		ch = 16
	}
	stack := scene.Stack().SetSize2(cw, ch)
	child.SetPositionKeepSize(true)
	child.SetStackPosition(geom.Pt(0, 0))
	stack.Add(child)
	badge.SetPositionKeepSize(true)
	// 角标出血：右上角 (childW-badgeW*0.4, -badgeH*0.4)（badge 自身尺寸
	// 在布局后才知，Dot 8/文本条按实测 size 折算偏移）。
	bw, bh := badge.Size().W, badge.Size().H
	badge.SetStackPosition(geom.Pt(cw-bw*0.4, -bh*0.4))
	stack.Add(badge)
	stack.SetHit(scene.HitSpec{Label: "badge-wrap", Behavior: scene.HitDeferToChild})
	return stack
}

// badgeNode 构建角标本体（Dot 圆点 / 文本条）。
func (b Badge) badgeNode(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	fill := b.statusFill(th)
	if b.Dot {
		dot := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: th.Radius.Pill}).SetSize2(8, 8)
		dot.SetHit(scene.HitSpec{Label: "badge"})
		return dot
	}
	label := b.Text
	if label == "" && b.Count > 0 {
		// antd overflowCount 语义：超过上限显示 "N+"（默认 99）。
		limit := b.OverflowCount
		if limit <= 0 {
			limit = 99
		}
		if b.Count > limit {
			label = fmt.Sprintf("%d+", limit)
		} else {
			label = fmt.Sprintf("%d", b.Count)
		}
	}
	font := text.DefaultFont(th.Font.SizeXS)
	layout := text.LayoutText(label, font, text.Options{})
	w := layout.Size.W + th.Space.MD
	h := layout.Size.H + th.Space.XS*2
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: h / 2}).SetSize2(w, h)
	fg := b.TextColor
	if fg == (geom.Color{}) {
		fg = th.Color.OnDanger
	}
	t := scene.Text(scene.TextSpec{Src: label, Layout: layout, Color: fg})
	t.SetMargin(geom.InLTRB(th.Space.SM, (h-layout.Size.H)/2, 0, 0))
	n.Add(t)
	n.SetHit(scene.HitSpec{Label: "badge"})
	return n
}

// Tag 是标签（浅底圆角文字）。指针接收器：closable/checkable 需要跨
// Build 持有手势与选中态（值接收器的字段写入不落回源实例）。
type Tag struct {
	controlBase
	Text  string
	Color geom.Color // 零值时取主题默认灰底
	// Closable 可关闭：右侧 × 图标（复用 Icon 组件 close 图元），
	// 点击触发 OnClose（antd closable）。
	Closable bool
	// Checkable 可选中：点击切换选中态（选中=主色实底+白字，
	// antd checkable 选中态）；未选保持默认灰底。OnCheck 回调切换值。
	Checkable bool
	Checked   bool
	// Preset 预设色名（antd color 预设值）："success"/"error"/
	// "warning"/"processing"/"default"；设置后按语义色渲染（浅底+
	// 同色文字+淡描边），Preset 解析优先于显式 Color。
	Preset string
	OnClose func()
	OnCheck func(bool)

	taps []*ui.GestureDetector // 跨 Build 持有（实例池，防 hover/状态丢失）
}

// presetColor 把预设色名解析为语义色（零值=未匹配）。
func presetColor(th *theme.Data, name string) geom.Color {
	switch name {
	case "success":
		return th.Color.Success
	case "error":
		return th.Color.Danger
	case "warning":
		return th.Color.Warning
	case "processing":
		return th.Color.Primary
	}
	return geom.Color{}
}

// Build 实现 ui.Widget。
//
// antd Tag 两态：**默认**（无 color）= 灰底 colorFillQuaternary + colorBorder
// 描边 + colorText 文字；**彩色**（显式 Color 或 Preset 预设色）= 浅底
// （基色 40% 透明）+ 同色文字 + 基色淡描边。checkable 选中态=主色实底+
// 白字；closable=右侧 × 图标。
func (t *Tag) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	base := t.Color
	if t.Preset != "" {
		base = presetColor(th, t.Preset)
	}
	textColor, fillColor, borderColor := th.Color.Text, th.Color.SurfaceSunken, th.Color.Border
	if base != (geom.Color{}) {
		textColor = base
		fillColor = base.WithAlpha(40)
		borderColor = base.WithAlpha(80)
	}
	// checkable 选中态：antd 主色实底+白字。
	selected := t.Checkable && t.Checked
	if selected {
		textColor, fillColor = th.Color.OnPrimary, th.Color.Primary
		borderColor = th.Color.Primary
	}
	font := text.DefaultFont(th.Font.SizeSM)
	layout := text.LayoutText(t.Text, font, text.Options{})
	contentW := layout.Size.W + th.Space.MD
	if t.Closable {
		contentW += 12 // × 图标
	}
	h := layout.Size.H + th.Space.SM
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(fillColor), Border: scene.Stroke(borderColor, 1), Radius: th.Radius.SM}).SetSize2(contentW, h)
	label := scene.Text(scene.TextSpec{Src: t.Text, Layout: layout, Color: textColor})
	label.SetMargin(geom.InLTRB(th.Space.SM, (h-layout.Size.H)/2, 0, 0))
	n.Add(label)
	if t.Closable {
		// × 图标：复用 Icon close 图元（组合复用），紧贴文字。
		ix := (&Icon{Name: "close", Size: 12, Color: th.Color.TextFaint}).Build(ctx)
		ix.SetMargin(geom.InLTRB(th.Space.SM, (h-12)/2, th.Space.SM, 0))
		n.Add(ix)
		ix.SetHit(scene.HitSpec{Label: "tag-close", Cursor: scene.CursorPointer})
	}
	// 命中与手势：**根节点统一 HitOpaque，单手势按点击坐标分流**——
	// closable 的 × 是框内热区（右侧 16px），其余区域归 checkable 切换
	// （antd 交互：整框可点，× 独立热区）。此前 × 手势绑在子 Icon 节点 +
	// 根 DeferToChild 的方案实测命中不到（--audit 点击 × 无效），改热区
	// 分流后根节点直接吃事件，无子节点命中穿透问题。
	n.SetHit(scene.HitSpec{Label: "tag", Behavior: scene.HitOpaque})
	// 走标准 attach 接线：InvalidateRoot 帧请求 + ctlKey 钉实例身份
	// （缺 ctlKey 时重建被 hover 差分误判为悬停迁移，事件路由不稳定）。
	hnd := t.tapHandler(ctx, false, func() {})
	if ctx.Input != nil && (t.Closable || t.Checkable) {
		t.taps = t.taps[:0]
		w := contentW // 根框宽（× 热区从右缘向左 16px）
		gc := &ui.GestureDetector{OnTap: func(p geom.Point) {
			if t.Closable && p.X > w-16 {
				if t.OnClose != nil {
					t.OnClose()
				}
				return
			}
			if t.Checkable {
				t.Checked = !t.Checked
				t.MarkDirty()
				if t.OnCheck != nil {
					t.OnCheck(t.Checked)
				}
			}
		}}
		t.taps = append(t.taps, gc)
		hnd = gc.Handler()
	}
	t.attach(ctx, n, hnd, false)
	return n
}

// ── Avatar ─────────────────────────────────────────────────────────

// Avatar 是圆形头像：文字取首字符，或 Icon 图标（组合复用）。
type Avatar struct {
	Text  string
	Size  float32
	Color geom.Color
	// TextColor 头像文字色（零值按底色取反色：默认灰底配深字，显式主色
	// 底配 OnPrimary 白字）。
	TextColor geom.Color
	// Icon 图标头像（antd Avatar.icon）：设置后优先于 Text（组合复用
	// Icon 组件，颜色随文字语义）。
	Icon string
	// FullText 整体显示 Text 不截首字符（"+N" 尾项等短语义文本用）。
	FullText bool
}

// Build 实现 ui.Widget。
//
// antd Avatar 默认形态：**灰底 #F0F0F0 + 深色文字**（colorText）；显式
// Color 时按预设色语义（主色底配 OnPrimary 白字）。文字字号 ≈ size×0.375
// （antd 32px 头像字号 14）。此前默认主色实底+白字，与官方对照不符。
func (a *Avatar) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	s := a.Size
	if s <= 0 {
		s = 32
	}
	fill := a.Color
	fg := a.TextColor
	if fill == (geom.Color{}) {
		// 默认：灰底 + 深字（colorFillTertiary 实色 + colorText）。
		fill = th.Color.SurfaceActive
		if fg == (geom.Color{}) {
			fg = th.Color.Text
		}
	} else if fg == (geom.Color{}) {
		fg = th.Color.OnPrimary
	}
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: s / 2}).SetSize2(s, s)
	switch {
	case a.Icon != "":
		// 图标头像：Icon 组件居中（尺寸 ≈ size×0.5，antd 语义）。
		ic := (&Icon{Name: a.Icon, Size: s * 0.5, Color: fg}).Build(ctx)
		ic.SetMargin(geom.InLTRB((s-s*0.5)/2, (s-s*0.5)/2, 0, 0))
		n.Add(ic)
	case a.Text != "":
		label := a.Text
		if !a.FullText {
			r := []rune(a.Text)
			label = string(r[:1])
		}
		font := text.DefaultFont(s * 0.375)
		layout := text.LayoutText(label, font, text.Options{})
		t := scene.Text(scene.TextSpec{Src: label, Layout: layout, Color: fg})
		t.SetMargin(geom.InLTRB((s-layout.Size.W)/2, (s-layout.Size.H)/2, 0, 0))
		n.Add(t)
	}
	n.SetHit(scene.HitSpec{Label: "avatar"})
	return n
}

// AvatarGroup 是头像组（antd Avatar.Group）：横向排布，max 控制显示
// 个数，超出部分合并为 "+N" 尾项（灰底深字）。
type AvatarGroup struct {
	Avatars []*Avatar
	// Max 最多显示个数（超出部分合并为 "+N" 尾项；0=全部显示）。
	Max int
}

// Build 实现 ui.Widget。
func (g *AvatarGroup) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	row := scene.Flex(scene.Horizontal, th.Space.XS)
	n := len(g.Avatars)
	if g.Max <= 0 || g.Max >= n {
		for _, a := range g.Avatars {
			row.Add(a.Build(ctx))
		}
		return row
	}
	for _, a := range g.Avatars[:g.Max] {
		row.Add(a.Build(ctx))
	}
	size := g.Avatars[0].Size
	if size <= 0 {
		size = 32
	}
	tail := &Avatar{Text: fmt.Sprintf("+%d", n-g.Max), Size: size, FullText: true}
	tail.TextColor = th.Color.Text
	row.Add(tail.Build(ctx))
	return row
}

// ── Empty / Skeleton / Spin ────────────────────────────────────────

// Empty 是空状态占位（图标圆 + 主文案 + 可选副文案）。
type Empty struct {
	Text    string
	Subtext string
	Width   float32
}

// Build 实现 ui.Widget。
func (e Empty) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	w := e.Width
	if w <= 0 {
		w = 220
	}
	n := scene.Box(scene.BoxSpec{}).SetSize2(w, 96)
	icon := scene.Box(scene.BoxSpec{
		Fill: scene.Solid(th.Color.SurfaceSunken), Radius: 20,
	}).SetSize2(40, 40)
	icon.SetStackPosition(geom.Pt((w-40)/2, 0)).SetPositionKeepSize(true)
	n.Add(icon)
	label := e.Text
	if label == "" {
		label = "暂无数据"
	}
	font := text.DefaultFont(th.Font.SizeBase)
	layout := text.LayoutText(label, font, text.Options{})
	t := scene.Text(scene.TextSpec{Src: label, Layout: layout, Color: th.Color.TextMuted})
	t.SetStackPosition(geom.Pt((w-layout.Size.W)/2, 48)).SetPositionKeepSize(true)
	n.Add(t)
	if e.Subtext != "" {
		sub := text.LayoutText(e.Subtext, text.DefaultFont(th.Font.SizeSM), text.Options{})
		st := scene.Text(scene.TextSpec{Src: e.Subtext, Layout: sub, Color: th.Color.TextFaint})
		st.SetStackPosition(geom.Pt((w-sub.Size.W)/2, 70)).SetPositionKeepSize(true)
		n.Add(st)
	}
	n.SetHit(scene.HitSpec{Label: "empty"})
	return n
}

// Skeleton 是加载骨架。
//
// 基础形态：n 条灰色圆角条，最后一条较短（Rows/Width/Height）。
// 组合形态（antd 默认骨架，Avatar/Title 任一开启时启用）：圆形头像 +
// 标题条 + 段落条组水平排布（组合复用同一条形图元，无新绘制通道）。
type Skeleton struct {
	Rows   int
	Width  float32
	Height float32
	// Avatar 头像条（32×32 圆，组合形态开关之一）。
	Avatar bool
	// Title 标题条（宽 40%，组合形态开关之一）。
	Title bool
	// Phase 呼吸相位 0..1（透明度调制 0.45..1），由调用方推进；Animated=true
	// 时相位按墙钟自动推进（1.4s/圈，antd active 呼吸节奏）并声明组件级
	// 动画——外部不必推帧。
	Phase    float32
	Animated bool
}

// skeletonBase 是 Skeleton 呼吸动画的墙钟起点。
var skeletonBase = time.Now()

// Build 实现 ui.Widget。
func (s Skeleton) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	rows := s.Rows
	if rows <= 0 {
		rows = 3
	}
	w, h := s.Width, s.Height
	if w <= 0 {
		w = 240
	}
	if h <= 0 {
		h = 12
	}
	if s.Animated && ctx.SetAnimating != nil {
		ctx.SetAnimating(&s, true)
	}
	phase := s.Phase
	if s.Animated {
		const breathePeriod = 1400 * time.Millisecond
		phase = float32(time.Since(skeletonBase)%breathePeriod) / float32(breathePeriod)
	}
	// antd active 呼吸：透明度在 1 与 0.45 间摆动（正弦半周期）。
	//
	// Animated 时**相位必须在渲染闭包内按墙钟实时算**：此前在 Build 里算
	// 一次、烘焙进节点 opacity 属性——帧泵虽按 claim 每拍标脏重录，但记录
	// 的属性是常量，画面零变化（--audit animate 模板实测逐帧差恒 0.000，
	// "呼吸态"实为死动画）。Spin 的 loading 旋转即此模式（phase 在闭包里算）。
	// 非 Animated：保留 Phase 外驱（调用方每帧改 Phase 并标脏，属性随之更新）。
	bar := func(bw, bh float32, round float32) *scene.Node {
		if !s.Animated {
			b := scene.Box(scene.BoxSpec{
				// SurfaceActive(#F0F0F0)：SurfaceSunken 与演示舞台同色，灰条会隐形。
				Fill: scene.Solid(th.Color.SurfaceActive), Radius: round,
			}).SetSize2(bw, bh)
			return b
		}
		// 渲染时实时相位：闭包每次录制重执行，alpha 随墙钟摆动。
		// Fill 走几何颜色带 alpha（FillHover 等已验证该通道），
		// 填充路径 AddRRect 保证圆角与 Box 形态一致。
		cp := render.NewPath().AddRRect(geom.RectXYWH(0, 0, bw, bh), round)
		n := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
			const breathePeriod = 1400 * time.Millisecond
			p := float32(time.Since(skeletonBase)%breathePeriod) / float32(breathePeriod)
			op := 1.0 - 0.55*(0.5-0.5*cos32(p*2*3.14159265))
			col := th.Color.SurfaceActive
			cv.FillPath(cp, render.Fill(geom.RGBAa(col.R, col.G, col.B, uint8(op*255))))
		}).SetSize2(bw, bh)
		return n
	}
	// 外驱模式（非 Animated）：Phase 由调用方推进，opacity 烘焙进节点属性。
	opacity := 1.0
	if phase > 0 {
		opacity = 1.0 - 0.55*(0.5-0.5*cos32(phase*2*3.14159265))
	}
	// 组合形态（Avatar/Title 任一开启）：antd 默认骨架布局。
	if s.Avatar || s.Title {
		root := scene.Flex(scene.Horizontal, th.Space.MD).SetCrossAlign(scene.CrossCenter)
		// Animated 时各条内部实时 alpha，root 不再叠加（否则二次相乘变暗）。
		if !s.Animated {
			root.SetOpacity(float32(opacity))
		}
		if s.Avatar {
			av := bar(32, 32, th.Radius.Pill)
			// 被动定位锚：头像圆独立采样（几何/颜色断言）。
			av.SetHit(scene.HitSpec{Label: "skeleton-avatar"})
			root.Add(av)
		}
		col := scene.Flex(scene.Vertical, th.Space.SM)
		if s.Title {
			t := bar(w*0.4, 16, th.Radius.SM)
			t.SetHit(scene.HitSpec{Label: "skeleton-title"}) // 被动定位锚
			col.Add(t)
		}
		for i := 0; i < rows; i++ {
			rw := w
			if i == rows-1 {
				rw = w * 0.6
			}
			col.Add(bar(rw, h, th.Radius.SM))
		}
		root.Add(col)
		root.SetHit(scene.HitSpec{Label: "skeleton"})
		return root
	}
	n := scene.Flex(scene.Vertical, th.Space.MD)
	if !s.Animated {
		n.SetOpacity(float32(opacity))
	}
	for i := 0; i < rows; i++ {
		rw := w
		if i == rows-1 {
			rw = w * 0.6
		}
		n.Add(bar(rw, h, th.Radius.SM))
	}
	n.SetHit(scene.HitSpec{Label: "skeleton"})
	return n
}

func cos32(v float32) float64 { return math.Cos(float64(v)) }

// Spin 是加载指示器（环形缺口；缺口位置由 Phase 决定，外部动画驱动）。
//
// Tip 非空时进入**容器包裹形态**（antd Spin tip）：指示器下方显示 Tip
// 文案（TextMuted），整体垂直居中排布。
type Spin struct {
	Size  float32
	Phase float32 // 0..1
	// Spin 自动旋转：相位在录制闭包内按墙钟计算（800ms/圈），并声明
	// 组件级动画（容器 60Hz 标脏重录）——外部不需要自己推帧。
	Spin bool
	// Tip 加载文案（非空时容器包裹形态：指示器+文案纵排）。
	Tip string
}

// spinBase 是 Spin 自动旋转的墙钟起点（与 icon.Spin 同约定）。
var spinBase = time.Now()

// Build 实现 ui.Widget。
func (s Spin) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	size := s.Size
	if size <= 0 {
		size = 20
	}
	if s.Spin && ctx.SetAnimating != nil {
		ctx.SetAnimating(&s, true)
	}
	n := scene.CustomPaintNode(func(c render.Canvas, sz geom.Size) {
		phase := s.Phase
		if s.Spin {
			const spinPeriod = 800 * time.Millisecond
			phase = float32(time.Since(spinBase)%spinPeriod) / float32(spinPeriod)
		}
		p := render.NewPath()
		p.AddArc(geom.RectXYWH(2, 2, sz.W-4, sz.H-4), phase*6.2831853, 4.7)
		c.StrokePath(p, render.Stroke(th.Color.Primary, 2).WithCap(render.CapRound))
	})
	n.SetSize2(size, size)
	// 被动定位锚：指示器几何/色断言用。
	n.SetHit(scene.HitSpec{Label: "spin"})
	if s.Tip != "" {
		// 容器包裹形态（antd Spin tip）：指示器下方显示文案。
		col := scene.Flex(scene.Vertical, th.Space.SM).SetCrossAlign(scene.CrossCenter)
		col.Add(n)
		tl := text.LayoutText(s.Tip, text.DefaultFont(th.Font.SizeSM), text.Options{})
		tip := scene.Text(scene.TextSpec{Src: s.Tip, Layout: tl, Color: th.Color.TextMuted})
		tip.SetHit(scene.HitSpec{Label: "spin-tip"}) // 被动锚：tip 文案断言用
		col.Add(tip)
		return col
	}
	return n
}

// ── Steps ───────────────────────────────────────────────────────────

// Steps 是步骤条。Current 为当前步（0 基），之前的步骤显示对勾。
type Steps struct {
	Labels  []string
	Current int
	// ErrorStep 错误步下标（-1 无，antd status="error"）：该步圆点转红、
	// 符号转 ×（优先于 Current 推导的状态）。
	ErrorStep int
}

// NewSteps 创建步骤条（ErrorStep=-1 无错误步）。
func NewSteps(labels []string, current int) *Steps {
	return &Steps{Labels: labels, Current: current, ErrorStep: -1}
}

// Build 实现 ui.Widget。
func (s Steps) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	// antd 步骤条：圆点一行、步骤间有连接线（已完成段主色、待办段灰），
	// 标签在圆点下方居中。
	n := scene.Flex(scene.Horizontal, 0)
	n.SetCrossAlign(scene.CrossStart)
	for i, label := range s.Labels {
		step := scene.Flex(scene.Vertical, th.Space.XS)
		step.SetCrossAlign(scene.CrossCenter)
		fill := th.Color.SurfaceActive
		fg := th.Color.TextMuted
		// antd wait 步的圆点是**空心描边圆**（透明底 + 灰边），完成/进行步
		// 是主色实心圆。此前 wait 用实心灰底，与 antd 官方对照不符。
		wait := i > s.Current
		errorStep := i == s.ErrorStep
		if errorStep {
			// antd status="error"：圆点与符号转红（colorError）。
			fill, fg = th.Color.Danger, th.Color.OnDanger
		} else if i <= s.Current {
			fill, fg = th.Color.Primary, th.Color.OnPrimary
		}
		circle := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: 10}).SetSize2(20, 20)
		if wait && !errorStep {
			circle = scene.Box(scene.BoxSpec{Border: scene.Stroke(th.Color.BorderStrong, 2), Radius: 10}).SetSize2(20, 20)
		}
		circle.SetHit(scene.HitSpec{Label: "steps-circle"}) // 被动锚：步骤圆点断言
		cap := fmt.Sprintf("%d", i+1)
		if errorStep {
			cap = "×"
		} else if i < s.Current {
			cap = "✓"
		}
		fl := text.LayoutText(cap, text.DefaultFont(th.Font.SizeXS), text.Options{})
		ct := scene.Text(scene.TextSpec{Src: cap, Layout: fl, Color: fg})
		ct.SetMargin(geom.InLTRB((20-fl.Size.W)/2, (20-fl.Size.H)/2, 0, 0))
		circle.Add(ct)
		step.Add(circle)
		ll := text.LayoutText(label, text.DefaultFont(th.Font.SizeSM), text.Options{})
		step.Add(scene.Text(scene.TextSpec{Src: label, Layout: ll, Color: th.Color.Text}))
		n.Add(step)
		if i < len(s.Labels)-1 {
			line := th.Color.Border
			if i < s.Current {
				line = th.Color.Primary
			}
			n.Add(scene.Box(scene.BoxSpec{Fill: scene.Solid(line), Radius: 1}).
				SetSize2(40, 2).SetMargin(geom.InLTRB(4, 9, 4, 0)))
		}
	}
	n.SetHit(scene.HitSpec{Label: "steps"})
	return n
}

// ── Segmented ──────────────────────────────────────────────────────

// Segmented 是分段选择器。选中段用 Primary 底色。
type Segmented struct {
	controlBase
	Options  []string
	Selected int
	OnChange func(int)

	// items 跨重建持有：分段项的 hover 态与手势实例必须稳定。
	items []*segmentItem
}

// NewSegmented 创建分段选择器。
func NewSegmented(options ...string) *Segmented { return &Segmented{Options: options} }

// Build 实现 ui.Widget。
func (s *Segmented) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(s)
	th := ctx.ThemeOr()
	if len(s.items) != len(s.Options) {
		s.items = make([]*segmentItem, len(s.Options))
		for i := range s.items {
			s.items[i] = &segmentItem{index: i, parent: s}
		}
	}
	children := make([]*scene.Node, 0, len(s.Options))
	for i, opt := range s.Options {
		s.items[i].label = opt
		children = append(children, s.items[i].Build(ctx))
	}
	row := scene.Flex(scene.Horizontal, 0, children...)
	// 轨道用 SurfaceActive（#F0F0F0，与 antd 分段容器一致）——
	// SurfaceSunken(#FAFAFA) 与演示舞台同色，轨道会隐形。
	n := scene.Box(scene.BoxSpec{
		Fill: scene.Solid(th.Color.SurfaceActive), Radius: th.Radius.Base,
	}, row)
	n.SetPadding(geom.InLTRB(2, 2, 2, 2))
	// 被动定位锚（HitDeferToChild 不影响交互）：容器轨道色断言用。
	n.SetHit(scene.HitSpec{Label: "segmented"})
	s.node = n
	return n
}

// segmentItem 是单个分段（内部组件）。
type segmentItem struct {
	controlBase
	index  int
	label  string
	parent *Segmented
}

func (s *segmentItem) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	selected := s.index == s.parent.Selected
	// antd Segmented：灰轨道上的选中项是**白色滑块**（带轻投影）+ 深色
	// 文字，不是主色实心块；未选项 hover 轻微加深（colorFill 系反馈）。
	fill := geom.Color{}
	fg := th.Color.Text
	shadow := (*scene.ShadowSpec)(nil)
	if selected {
		fill = th.Color.Surface
		shadow = &scene.ShadowSpec{Dx: 0, Dy: 1, Blur: 3, Color: th.Color.Shadow}
	} else if s.hover {
		fill = th.Color.FillHover
	}
	font := text.DefaultFont(th.Font.SizeSM)
	layout := text.LayoutText(s.label, font, text.Options{})
	w := layout.Size.W + th.Space.LG
	const h = float32(26)
	item := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: th.Radius.SM, Shadow: shadow}).SetSize2(w, h)
	label := scene.Text(scene.TextSpec{Src: s.label, Layout: layout, Color: fg})
	label.SetMargin(geom.InLTRB(th.Space.SM, (h-layout.Size.H)/2, 0, 0))
	item.Add(label)
	item.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "segment", Cursor: scene.CursorPointer})
	parent, index := s.parent, s.index
	hnd := s.tapHandler(ctx, false, func() {
		if parent.Selected == index {
			return
		}
		parent.Selected = index
		parent.MarkDirty()
		if parent.OnChange != nil {
			parent.OnChange(index)
		}
	})
	s.attach(ctx, item, hnd, false)
	return item
}

// ── Pagination ─────────────────────────────────────────────────────

// Pagination 是分页器。Page 从 1 开始。
// Total 语义随 PageSize 变化：PageSize<=0 时 Total 即总页数（简单模式）；
// PageSize>0 时按 antd 语义解释为总条目数，总页数=ceil(Total/PageSize)。
type Pagination struct {
	controlBase
	Page  int
	Total int
	// PageSize 每页条数（antd pageSize）。>0 时 Total 是总条目数。
	PageSize int
	// ShowTotal 显示"共 X 条"文字（需 PageSize>0 才有意义）。
	ShowTotal bool
	// ShowSizeChanger 显示"X 条/页"切换器（需 PageSize>0）。
	ShowSizeChanger bool
	// ShowQuickJumper 显示"跳至 X 页"快速跳页输入框（回车提交，超范围钳制）。
	ShowQuickJumper bool
	OnChange        func(int)
	// OnPageSize 每页条数变化回调（切换后 Page 已钳制到新总页数范围）。
	OnPageSize func(int)

	// buttons 跨重建持有：页码按钮的 hover 态与手势实例必须稳定——
	// Build 内新建实例会让 hover 高亮随重建丢失（--audit 实测）。
	buttons []*pageButton
	// sizer/jumper 同理跨重建持有（内嵌子组件实例池）。
	sizer  *Select
	jumper *Input
}

// sizeChangerOptions 是每页条数三档（与 antd 常用档位对齐）。
var sizeChangerOptions = []string{"10 条/页", "20 条/页", "50 条/页"}

// NewPagination 创建分页器。
func NewPagination(page, total int) *Pagination {
	if page < 1 {
		page = 1
	}
	return &Pagination{Page: page, Total: total}
}

// totalPages 返回总页数：PageSize>0 时按条目数换算，否则 Total 即页数。
func (p *Pagination) totalPages() int {
	if p.PageSize > 0 {
		n := (p.Total + p.PageSize - 1) / p.PageSize
		if n < 1 {
			n = 1
		}
		return n
	}
	if p.Total < 1 {
		return 1
	}
	return p.Total
}

// SetPage 跳页（钳制到 [1,总页数]），触发 OnChange。
func (p *Pagination) SetPage(n int) {
	if n < 1 {
		n = 1
	}
	if max := p.totalPages(); n > max {
		n = max
	}
	if p.Page == n {
		return
	}
	p.Page = n
	p.MarkDirty()
	if p.OnChange != nil {
		p.OnChange(n)
	}
}

// setPageSize 切换每页条数并钳制当前页，触发 OnPageSize。
func (p *Pagination) setPageSize(size int) {
	if size < 1 || size == p.PageSize {
		return
	}
	p.PageSize = size
	if max := p.totalPages(); p.Page > max {
		p.Page = max
	}
	p.MarkDirty()
	if p.OnPageSize != nil {
		p.OnPageSize(size)
	}
}

// Build 实现 ui.Widget。
func (p *Pagination) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(p)
	th := ctx.ThemeOr()
	row := p.buildPages(ctx, th)

	// 附加件全关：保持单行页码（与旧版完全一致，索引锚不位移）。
	if !p.ShowTotal && !p.ShowSizeChanger && !p.ShowQuickJumper {
		p.node = row
		return row
	}
	wrap := scene.Flex(scene.Horizontal, th.Space.MD).SetAlign(scene.Center)
	if p.ShowTotal && p.PageSize > 0 {
		total := fmt.Sprintf("共 %d 条", p.Total)
		l := text.LayoutText(total, text.DefaultFont(th.Font.SizeSM), text.Options{})
		t := scene.Text(scene.TextSpec{Src: total, Layout: l, Color: th.Color.Text})
		// 被动定位锚：total 文字无 hit，文本断言无法定位（breadcrumb 先例）。
		t.SetHit(scene.HitSpec{Label: "page-total"})
		wrap.Add(t)
	}
	wrap.Add(row)
	if p.ShowSizeChanger && p.PageSize > 0 {
		wrap.Add(p.buildSizer(ctx, th))
	}
	if p.ShowQuickJumper {
		wrap.Add(p.buildJumper(ctx, th))
	}
	p.node = wrap
	return wrap
}

// buildPages 构建页码行（±2 窗口 + 实例池）。
func (p *Pagination) buildPages(ctx *ui.BuildContext, th *theme.Data) *scene.Node {
	pages := p.visiblePages()
	// 实例池与可见页对齐：页码窗口滑动时增删实例（保留旧实例的 hover 态）。
	if len(p.buttons) != len(pages) {
		p.buttons = make([]*pageButton, len(pages))
		for i := range p.buttons {
			p.buttons[i] = &pageButton{parent: p}
		}
	}
	row := scene.Flex(scene.Horizontal, th.Space.XS)
	for i, page := range pages {
		b := p.buttons[i]
		b.page, b.current, b.label = page, page == p.Page, fmt.Sprintf("%d", page)
		row.Add(b.Build(ctx))
	}
	return row
}

// buildSizer 构建每页条数切换器（Select 内嵌实例，跨重建持有）。
func (p *Pagination) buildSizer(ctx *ui.BuildContext, th *theme.Data) *scene.Node {
	if p.sizer == nil {
		p.sizer = &Select{Width: 88, MaxHeight: 240, Placeholder: ""}
		p.sizer.Options = append([]string{}, sizeChangerOptions...)
	}
	// Selected 与 PageSize 同步（外部直改 PageSize 也要反映，不触发回调）。
	sel := -1
	for i, size := range []int{10, 20, 50} {
		if size == p.PageSize {
			sel = i
			break
		}
	}
	p.sizer.Selected = sel
	p.sizer.OnChange = func(i int) {
		if i < 0 || i >= len(sizeChangerOptions) {
			return
		}
		p.setPageSize([]int{10, 20, 50}[i])
	}
	return p.sizer.Build(ctx)
}

// buildJumper 构建快速跳页："跳至 [input] 页"，回车提交（超范围钳制）。
func (p *Pagination) buildJumper(ctx *ui.BuildContext, th *theme.Data) *scene.Node {
	inner := scene.Flex(scene.Horizontal, th.Space.XS).SetAlign(scene.Center)
	jump := "跳至"
	l := text.LayoutText(jump, text.DefaultFont(th.Font.SizeSM), text.Options{})
	hint := scene.Text(scene.TextSpec{Src: jump, Layout: l, Color: th.Color.Text})
	hint.SetHit(scene.HitSpec{Label: "page-jump-hint"}) // 被动定位锚
	inner.Add(hint)
	if p.jumper == nil {
		p.jumper = &Input{Placeholder: "", Width: 48, Height: 26, HitLabel: "page-jumper"}
	}
	p.jumper.OnSubmit = func(v string) {
		if n, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
			p.SetPage(n)
		}
	}
	inner.Add(p.jumper.Build(ctx))
	to := "页"
	l2 := text.LayoutText(to, text.DefaultFont(th.Font.SizeSM), text.Options{})
	inner.Add(scene.Text(scene.TextSpec{Src: to, Layout: l2, Color: th.Color.Text}))
	return inner
}

func (p *Pagination) visiblePages() []int {
	total := p.totalPages()
	if total < 1 {
		return []int{1}
	}
	lo, hi := 1, total
	if total > 7 {
		lo, hi = p.Page-2, p.Page+2
		if lo < 1 {
			lo, hi = 1, 5
		}
		if hi > total {
			hi, lo = total, total-4
		}
	}
	out := make([]int, 0, hi-lo+1)
	for i := lo; i <= hi; i++ {
		out = append(out, i)
	}
	return out
}

// pageButton 是单个页码按钮（内部组件）。
type pageButton struct {
	controlBase
	page    int
	current bool
	label   string
	parent  *Pagination
}

func (b *pageButton) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	fill := geom.Color{}
	fg := th.Color.Text
	// antd：未选页 hover 显示浅灰底（colorFillTertiary 系）；当前页主色实底+白字。
	if b.current {
		fill, fg = th.Color.Primary, th.Color.OnPrimary
	} else if b.hover {
		fill = th.Color.FillHover
	}
	layout := text.LayoutText(b.label, text.DefaultFont(th.Font.SizeSM), text.Options{})
	w := layout.Size.W + th.Space.MD
	if w < 26 {
		w = 26
	}
	n := scene.Box(scene.BoxSpec{Fill: scene.Solid(fill), Radius: th.Radius.SM}).SetSize2(w, 26)
	t := scene.Text(scene.TextSpec{Src: b.label, Layout: layout, Color: fg})
	t.SetMargin(geom.InLTRB((w-layout.Size.W)/2, (26-layout.Size.H)/2, 0, 0))
	n.Add(t)
	n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "page", Cursor: scene.CursorPointer})
	parent, page := b.parent, b.page
	hnd := b.tapHandler(ctx, false, func() {
		parent.SetPage(page)
	})
	b.attach(ctx, n, hnd, false)
	return n
}

// ── Breadcrumb ─────────────────────────────────────────────────────

// Breadcrumb 是面包屑导航。最后一项为当前项（高亮且不可点）。
type Breadcrumb struct {
	Items   []string
	OnClick func(int)
	node    *scene.Node
	taps    []*ui.GestureDetector // 跨 Build 存活，避免回调闭包被 GC
}

// Node 返回最近一次构建出的场景节点。
func (b *Breadcrumb) Node() *scene.Node { return b.node }

// Build 实现 ui.Widget。
func (b *Breadcrumb) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	row := scene.Flex(scene.Horizontal, th.Space.SM)
	b.taps = b.taps[:0]
	for i, label := range b.Items {
		if i > 0 {
			sep := text.LayoutText("›", text.DefaultFont(th.Font.SizeBase), text.Options{})
			row.Add(scene.Text(scene.TextSpec{Src: "›", Layout: sep, Color: th.Color.TextFaint}))
		}
		color := th.Color.TextMuted
		last := i == len(b.Items)-1
		if last {
			color = th.Color.Text
		}
		layout := text.LayoutText(label, text.DefaultFont(th.Font.SizeBase), text.Options{})
		item := scene.Text(scene.TextSpec{Src: label, Layout: layout, Color: color})
		if last {
			// 末项（当前页）被动定位锚：末项不可点、原本无 hit spec，
			// 颜色断言永远无法采样——加锚只此一处，不碰可点项。
			item.SetHit(scene.HitSpec{Label: "breadcrumb-item"})
		}
		if !last && b.OnClick != nil {
			idx := i
			g := &ui.GestureDetector{OnTap: func(geom.Point) { b.OnClick(idx) }}
			b.taps = append(b.taps, g)
			item.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "breadcrumb", Cursor: scene.CursorPointer})
			if ctx.Input != nil {
				ctx.Input.Set(item, g.Handler())
			}
		}
		row.Add(item)
	}
	b.node = row
	return row
}

// ── Collapse ───────────────────────────────────────────────────────

// CollapsePanel 是折叠面板的一项。
type CollapsePanel struct {
	Title   string
	Content ui.Widget
}

// Collapse 是可折叠面板组。Open 为展开项索引，-1 表示全部收起。
type Collapse struct {
	controlBase
	Panels   []CollapsePanel
	Open     int    // 单开（默认）：展开的面板下标，-1=全收
	Multiple bool   // 多开模式（antd 默认）：Multiple 模式下各面板独立开关
	OpenSet  []bool // Multiple 模式下各面板的展开状态（与 Panels 对齐）
	OnChange func(int)
}

// NewCollapse 创建折叠面板组（默认全部收起）。
func NewCollapse(panels ...CollapsePanel) *Collapse {
	return &Collapse{Panels: panels, Open: -1}
}

// Build 实现 ui.Widget。
func (c *Collapse) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(c)
	th := ctx.ThemeOr()
	// Multiple 模式：OpenSet 与 Panels 对齐（Panels 可变长）。
	if c.Multiple && len(c.OpenSet) != len(c.Panels) {
		c.OpenSet = make([]bool, len(c.Panels))
	}
	col := scene.Flex(scene.Vertical, th.Space.XS)
	// 面板通栏铺满父级宽度（antd 折叠面板是全宽边框容器，不是紧凑按钮盒）。
	col.SetCrossAlign(scene.CrossStretch)
	for i := range c.Panels {
		col.Add((&collapseItem{index: i, parent: c}).Build(ctx))
	}
	c.node = col
	return col
}

// collapseItem 是单个折叠项（内部组件）。
type collapseItem struct {
	controlBase
	index  int
	parent *Collapse
}

func (c *collapseItem) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	panel := c.parent.Panels[c.index]
	open := c.index == c.parent.Open
	if c.parent.Multiple {
		open = c.parent.OpenSet[c.index]
	}
	// antd 折叠面板：面板=通栏边框容器；表头一行"标题居左、箭头居右"
	// （SpaceBetween，宽度随父级自适应，箭头不再写死 x=320 飘在面板外）。
	layout := text.LayoutText(panel.Title, text.DefaultFont(th.Font.SizeBase), text.Options{})
	label := scene.Text(scene.TextSpec{Src: panel.Title, Layout: layout, Color: th.Color.Text})
	label.SetMargin(geom.InLTRB(th.Space.MD, 0, 0, 0))
	arrow := "▸"
	if open {
		arrow = "▾"
	}
	al := text.LayoutText(arrow, text.DefaultFont(th.Font.SizeSM), text.Options{})
	at := scene.Text(scene.TextSpec{Src: arrow, Layout: al, Color: th.Color.TextMuted})
	at.SetMargin(geom.InLTRB(0, 0, th.Space.MD, 0))
	header := scene.Flex(scene.Horizontal, 0, label, at)
	header.SetCrossAlign(scene.CrossCenter).SetMainAlign(scene.MainSpaceBetween)
	header.SetHeight(36)
	header.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "collapse", Cursor: scene.CursorPointer})
	parent, index := c.parent, c.index
	hnd := c.tapHandler(ctx, false, func() {
		if parent.Multiple {
			parent.OpenSet[index] = !parent.OpenSet[index]
			// 兼容 Open 单开语义：多开模式下把最后一个展开项同步给 Open。
			parent.Open = -1
			for i, v := range parent.OpenSet {
				if v {
					parent.Open = i
				}
			}
			parent.MarkDirty()
			if parent.OnChange != nil {
				parent.OnChange(parent.Open)
			}
			return
		}
		next := -1
		if parent.Open != index {
			next = index
		}
		parent.Open = next
		parent.MarkDirty()
		if parent.OnChange != nil {
			parent.OnChange(next)
		}
	})
	c.attach(ctx, header, hnd, false)
	// 内容区：表头下方带分隔线与内边距（antd 形态）。
	rootChildren := []*scene.Node{header}
	if open && panel.Content != nil {
		wrap := scene.Box(scene.BoxSpec{})
		wrap.SetPadding(geom.InLTRB(th.Space.MD, th.Space.SM, th.Space.MD, th.Space.MD))
		wrap.Add(buildChild(ctx, panel.Content))
		rootChildren = append(rootChildren, wrap)
	}
	root := scene.Flex(scene.Vertical, 0, rootChildren...)
	root.SetCrossAlign(scene.CrossStretch)
	// 定位 + 紧约束：内容根被拉满面板宽度（表头 SpaceBetween 才有全宽）。
	root.SetPositionKeepSize(false)
	root.SetStackPosition(geom.Pt(0, 0))
	panelBox := scene.Box(scene.BoxSpec{
		Fill: scene.Solid(th.Color.Surface), Border: scene.Stroke(th.Color.Border, 1), Radius: th.Radius.Base,
	}, root)
	// 被动定位锚（HitDeferToChild 不影响交互）：面板级内容存在性断言用
	// （表头锚的子树不含内容文字——内容挂在面板容器下）。
	panelBox.SetHit(scene.HitSpec{Label: "collapse-panel"})
	return panelBox
}
