package widgets

// ── Mentions 提及（M12 高级录入）─────────────────────────────────
//
// antd：<Mentions options defaultValue onChange placeholder prefix>
// 第一轮：**纯组合复用 Input**（Input 实例持久持有，HitLabel 自定义
// "mentions-input"——Input 注释明示此嵌入用法）+ 输入 "@" 弹选项
// 浮层（BandPopup+BottomStart，Select 同通道）+ 选项 **down 即执行**
// （浮层内 tap 配对不可靠结论）→ 插入 "@label " 到文本 + OnChange。
//
// 触发判定：OnChanged 文本末尾为 '@' 且面板未开 → 弹面板；
// 选项选中后文本末尾拼接 label+" "（antd 插入后跟分隔符）。
// 未实现（如实标注）：自动定位光标处的 @、多个 prefix（数组）、
// split 分隔、textarea 多行形态（Multiline 可经字段开启）、
// loading 选项、status/variant 定制（Input 字段透传可加）。
// 键盘上下选已交付（第一百三十二轮）：面板打开时 Input 的 Key 链
// 拦截 ↑↓Enter——↑↓ 移动高亮项（activeIdx+原地迁移渲染）、Enter
// 确认 pick；Esc 走浮层宿主关闭（DismissEsc 已有）。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// Mentions 是提及输入框（Input 组合 + @ 触发选项浮层）。
type Mentions struct {
	controlBase
	Options     []string // @ 候选（antd options 的 value/label 合一）
	Value       string
	Placeholder string
	Width       float32
	// OnChange 文本变化（键入与选项插入都触发）。
	OnChange func(text string)

	input     *Input // 组合复用（持久实例——教训⑨）
	handle    overlay.Handle
	fb        feedbackHost
	anchor    *scene.Node
	panelOpen bool
	// activeIdx 键盘高亮项下标（↑↓ 移动、Enter 确认；-1=无高亮）。
	activeIdx int
	// rows 面板打开期间的行节点与文案（高亮原地迁移——Select 选项
	// 同款 SetBox/SetTextPaint 手法）。
	rows []mentionsRow
	// curTh 面板主题快照（高亮迁移用——openPanel 时记）。
	curTh *theme.Data
	// inserting 程序化插入标志：插入后的重建会再次走 OnChanged 路径
	// 前 Input.Value 已是新值，无需去重——保留字段供后续多前缀扩展。
	inserting bool
}

// mentionsRow 一行的节点与排版（高亮迁移用）。
type mentionsRow struct {
	node   *scene.Node
	txt    *scene.Node
	src    string
	layout *text.Layout
}

// NewMentions 创建提及输入框。
func NewMentions(options []string) *Mentions {
	return &Mentions{Options: options, Placeholder: "输入 @ 唤起提及", Width: 260}
}

// Build 实现 ui.Widget：Input 组合（面板打开时 Key 链拦截 ↑↓Enter）。
func (m *Mentions) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(m)
	if m.input == nil {
		m.input = &Input{HitLabel: "mentions-input"}
	}
	// 字段同步（持久实例跨 Build 复用）。
	m.input.Placeholder = m.Placeholder
	m.input.Value = m.Value
	if m.Width > 0 {
		m.input.Width = m.Width
	}
	var inputNode *scene.Node
	m.input.OnChanged = func(s string) {
		m.Value = s
		if m.OnChange != nil {
			m.OnChange(s)
		}
		// 触发判定：文本末尾为 '@' → 弹选项浮层。
		if len(s) > 0 && s[len(s)-1] == '@' && !m.panelOpen {
			m.activeIdx = -1
			m.openPanel(ctx, inputNode)
		}
	}
	// 键盘上下选（antd 键盘交互）：面板打开时优先消费 ↑↓Enter——
	// Input 组合的 Key 链在 Build 后替换（白拿 Input 焦点/聚焦链）。
	m.input.OnKey = func(ev *input.KeyEvent) bool {
		if !m.panelOpen {
			return false
		}
		if ev.Kind == input.KeyUp {
			return false
		}
		switch ev.Code {
		case input.CodeDown:
			if m.activeIdx < len(m.Options)-1 {
				m.moveHighlight(m.activeIdx + 1)
			} else if m.activeIdx < 0 {
				m.moveHighlight(0)
			}
			return true
		case input.CodeUp:
			if m.activeIdx > 0 {
				m.moveHighlight(m.activeIdx - 1)
			} else if m.activeIdx < 0 {
				m.moveHighlight(len(m.Options) - 1)
			}
			return true
		case input.CodeEnter:
			if m.activeIdx >= 0 && m.activeIdx < len(m.Options) {
				m.pick(m.Options[m.activeIdx], ctx)
				return true
			}
		}
		return false
	}
	n := m.input.Build(ctx)
	m.anchor = n
	inputNode = n
	return n
}

// moveHighlight 高亮项迁移（原地——Select 选项 SetBox/SetTextPaint 同款；
// 越界钳制：无行时保持 -1）。
func (m *Mentions) moveHighlight(idx int) {
	if len(m.rows) == 0 {
		m.activeIdx = -1
		return
	}
	if idx < 0 {
		idx = -1
	}
	if idx > len(m.rows)-1 {
		idx = len(m.rows) - 1
	}
	if old := m.activeIdx; old >= 0 && old < len(m.rows) {
		m.paintRow(m.rows[old], false)
	}
	m.activeIdx = idx
	if idx >= 0 && idx < len(m.rows) {
		m.paintRow(m.rows[idx], true)
	}
}

// paintRow 原地刷行高亮（InfoBg 浅底——antd mentions 高亮视觉）。
func (m *Mentions) paintRow(r mentionsRow, on bool) {
	spec := scene.BoxSpec{}
	color := m.curTh.Color.Text
	if on {
		spec = scene.BoxSpec{Fill: scene.Solid(m.curTh.Color.InfoBg)}
		color = m.curTh.Color.Primary
	}
	r.node.SetBox(spec)
	r.txt.SetTextPaint(scene.TextSpec{Src: r.src, Layout: r.layout, Color: color})
	r.node.MarkNeedsPaint()
}

// Show 程序化弹出选项面板（测试/自动化用）。
func (m *Mentions) Show(ctx *ui.BuildContext) {
	if m.anchor != nil {
		m.openPanel(ctx, m.anchor)
	}
}

// IsOpen 返回选项面板是否打开。
func (m *Mentions) IsOpen() bool { return m.panelOpen }

// Close 关闭选项面板。
func (m *Mentions) Close() {
	if m.fb.host != nil {
		m.closePanel(nil)
	}
}

// openPanel 弹选项浮层（锚=Input 节点）。
func (m *Mentions) openPanel(ctx *ui.BuildContext, anchor *scene.Node) {
	if m.handle != 0 || ctx == nil || ctx.Overlay == nil {
		return
	}
	th := ctx.ThemeOr()
	m.curTh = th
	m.panelOpen = true
	m.activeIdx = -1
	m.rows = m.rows[:0]
	w := m.Width
	if w <= 0 {
		w = 260
	}
	panel := scene.Flex(scene.Vertical, 0)
	panel.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	panel.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "mentions-panel"})
	for _, opt := range m.Options {
		opt := opt
		row := scene.Box(scene.BoxSpec{})
		row.SetWidth(w)
		row.SetHeight(30)
		l := text.LayoutText(opt, text.DefaultFont(th.Font.SizeBase), text.Options{})
		txt := scene.Text(scene.TextSpec{Src: opt, Layout: l, Color: th.Color.Text})
		txt.SetMargin(geom.InLTRB(12, (30-l.Size.H)/2, 0, 0))
		row.Add(txt)
		row.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "mentions-option", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			ctx.Input.Set(row, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
				if !ev.Down() {
					return false
				}
				// down 即执行（浮层内 tap 配对不可靠结论复用）。
				m.pick(opt, ctx)
				return true
			}})
		}
		panel.Add(row)
		m.rows = append(m.rows, mentionsRow{node: row, txt: txt, src: opt, layout: l})
	}
	m.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: overlay.PlacementBottomStart,
		Size:      geom.Size{W: w, H: 0}, // 高自适应
		Dismiss:   overlay.DismissOutsideClick | overlay.DismissEsc,
		Tag:       "mentions",
		OnClose: func() {
			m.handle = 0
			m.panelOpen = false
		},
		Build: func() *scene.Node { return panel },
	})
}

// pick 选中选项：去掉触发 '@'、插入 "@label "，经 Input 重建同步文本。
func (m *Mentions) pick(opt string, ctx *ui.BuildContext) {
	base := m.Value
	if len(base) > 0 && base[len(base)-1] == '@' {
		base = base[:len(base)-1]
	}
	m.Value = base + "@" + opt + " "
	if m.input != nil {
		m.input.Value = m.Value
	}
	m.MarkDirty() // Input 重建同步 editor 文本 + 触发框刷新
	if ctx != nil && ctx.Overlay != nil && m.handle != 0 {
		ctx.Overlay.Close(m.handle)
	}
	m.handle = 0
	m.panelOpen = false
	if m.OnChange != nil {
		m.OnChange(m.Value)
	}
}

// closePanel 关闭面板。
func (m *Mentions) closePanel(ctx *ui.BuildContext) {
	if m.handle != 0 && m.fb.host != nil {
		m.fb.host.Close(m.handle)
	}
	m.handle = 0
	m.panelOpen = false
}
