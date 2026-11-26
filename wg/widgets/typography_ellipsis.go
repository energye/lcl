package widgets

// ── Typography ellipsis（antd Typography ellipsis）──────────────────
//
// 多行文本按 Rows 限制行数，超出以省略符结尾；Expandable 时末尾追加
// "展开/收起" 文本按钮，点击切换全量/截断显示（antd symbol 默认
// Expand/Collapse）。
//
// 未实现（如实标注）：tooltip 悬浮全文、suffix 中间省略、controlled
// expanded/onEllipsis、expandable='collapsible'（收起时不留省略号）。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// EllipsisText 是可多行省略/展开的文本（antd Typography ellipsis）。
type EllipsisText struct {
	controlBase
	Body string
	// Rows 限制行数（antd ellipsis.rows，0 = 不限即不省略）。
	Rows int
	// MaxWidth 换行宽度（antd 需显式宽——Text 无 100% 宽）。
	MaxWidth float32
	// Expandable 可展开/收起（antd ellipsis.expandable）。
	Expandable bool
	// Symbol 展开/收起按钮文案（antd ellipsis.symbol；默认 展开/收起）。
	Symbol string
	// DefaultExpanded 初始展开态（antd ellipsis.defaultExpanded）。
	DefaultExpanded bool
	// OnExpand 展开/收起回调（antd ellipsis.onExpand）。
	OnExpand func(expanded bool)

	expanded bool
}

// NewEllipsisText 创建可省略/展开的文本。
func NewEllipsisText(body string) *EllipsisText {
	return &EllipsisText{Body: body, Rows: 2, MaxWidth: 400}
}

// Expanded 返回当前展开态（诊断/断言用）。
func (e *EllipsisText) Expanded() bool { return e.expanded }

// Toggle 程序化切换展开/收起（测试/自动化用）。
func (e *EllipsisText) Toggle() {
	e.expanded = !e.expanded
	e.MarkDirty()
	if e.OnExpand != nil {
		e.OnExpand(e.expanded)
	}
}

// Build 实现 ui.Widget。
func (e *EllipsisText) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(e)
	th := ctx.ThemeOr()
	font := text.DefaultFont(th.Font.SizeBase)
	maxW := e.MaxWidth
	if maxW <= 0 {
		maxW = 400
	}

	rows := e.Rows
	if rows <= 0 {
		rows = 0 // 不限制
	}
	opts := text.Options{MaxWidth: maxW, LineHeight: 1.57}
	// 收起态：限行+省略符；展开态：不限行（但保留换行宽）。
	if !e.expanded && rows > 0 {
		opts.MaxLines = rows
		opts.Ellipsis = "…"
	}
	layout := text.LayoutText(e.Body, font, opts)
	tv := scene.Text(scene.TextSpec{
		Src: e.Body, Layout: layout, Color: th.Color.Text,
		MaxWidth: maxW, LayoutFont: font, LayoutW: maxW,
	})
	tv.SetHit(scene.HitSpec{Label: "typography"})

	if !e.Expandable {
		return tv
	}

	col := scene.Flex(scene.Vertical, th.Space.XS)
	col.SetWidth(maxW)
	col.Add(tv)

	// 展开/收起按钮（文本按钮，主色）。
	label := e.symbolLabel()
	l := text.LayoutText(label, text.DefaultFont(th.Font.SizeSM), text.Options{})
	btn := scene.Text(scene.TextSpec{Src: label, Layout: l, Color: th.Color.Primary})
	btn.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "typography-expand", Cursor: scene.CursorPointer})
	if ctx.Input != nil {
		ctx.Input.Set(btn, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
			if !ev.Down() {
				return false
			}
			e.expanded = !e.expanded
			e.MarkDirty()
			if e.OnExpand != nil {
				e.OnExpand(e.expanded)
			}
			return true
		}})
	}
	col.Add(btn)
	col.SetHit(scene.HitSpec{Label: "typography-ellipsis"})
	return col
}

// symbolLabel 展开/收起按钮文案：自定义 Symbol 优先，否则 展开/收起。
func (e *EllipsisText) symbolLabel() string {
	if e.Symbol != "" {
		if e.expanded {
			return "收起" + e.Symbol
		}
		return e.Symbol
	}
	if e.expanded {
		return "收起"
	}
	return "展开"
}

var _ = theme.Default
var _ = geom.Pt
