package widgets

// ── Typography editable（antd Typography editable）──────────────────
//
// 文本右侧显示编辑图标（hover 可见），点击切换为 Input 编辑态：
// Enter 提交（OnChange）、Esc 取消（还原）、失焦提交。编辑态复用 Input
// 组件（持久实例，焦点/编辑器跨重建保留）。
//
// 未实现（如实标注）：editable.icon 自定义、tooltip 文案、triggerType
// (icon/text/both)、enterIcon、maxLength 截断（Input 有 Count 但不限长）。

import (
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// EditableText 是可编辑的内联文本（antd Typography editable）。
type EditableText struct {
	controlBase
	Body string
	// Type 语义色（同 Text.Type）。
	Type TextType
	// MaxLength 最大长度（0 = 不限；编辑态超出不截断，仅语义标注）。
	MaxLength int
	// OnChange 提交回调（Enter/失焦触发）。
	OnChange func(text string)
	// OnCancel 取消回调（Esc 触发，值还原）。
	OnCancel func()

	editing bool
	editor  *Input // 编辑态 Input（持久实例）
}

// NewEditableText 创建可编辑文本。
func NewEditableText(body string) *EditableText { return &EditableText{Body: body} }

// Editing 返回是否处于编辑态（诊断/断言用）。
func (e *EditableText) Editing() bool { return e.editing }

// BeginEdit 进入编辑态。
func (e *EditableText) BeginEdit() {
	e.editing = true
	e.MarkDirty()
}

// Commit 提交编辑（值来自编辑框）。
func (e *EditableText) Commit() {
	if e.editor != nil {
		e.Body = e.editor.Value
	}
	e.editing = false
	e.MarkDirty()
	if e.OnChange != nil {
		e.OnChange(e.Body)
	}
}

// Cancel 取消编辑（值还原）。
func (e *EditableText) Cancel() {
	e.editing = false
	e.MarkDirty()
	if e.OnCancel != nil {
		e.OnCancel()
	}
}

// Build 实现 ui.Widget：非编辑态=文本+edit 图标；编辑态=Input。
func (e *EditableText) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(e)
	th := ctx.ThemeOr()

	if e.editing {
		if e.editor == nil {
			e.editor = NewInput("")
		}
		e.editor.Value = e.Body
		e.editor.Width = 200
		e.editor.OnSubmit = func(string) { e.Commit() }
		// Esc 取消（OnKey 优先消费，不落编辑模型）。
		e.editor.OnKey = func(ev *input.KeyEvent) bool {
			if ev.Code == input.CodeEscape {
				e.Cancel()
				return true
			}
			return false
		}
		n := e.editor.Build(ctx)
		n.SetHit(scene.HitSpec{Label: "typography-edit-input"})
		return n
	}

	color := Text{Type: e.Type}.color(th)
	font := text.DefaultFont(th.Font.SizeBase)
	layout := text.LayoutText(e.Body, font, text.Options{LineHeight: 1.57})
	txt := scene.Text(scene.TextSpec{Src: e.Body, Layout: layout, Color: color})
	txt.SetHit(scene.HitSpec{Label: "typography"})

	row := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
	row.Add(txt)
	ic := (&Icon{Name: "edit", Size: 14, Color: th.Color.TextMuted}).Build(ctx)
	ic.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "typography-edit", Cursor: scene.CursorPointer})
	if ctx.Input != nil {
		ctx.Input.Set(ic, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
			if !ev.Down() {
				return false
			}
			e.BeginEdit()
			return true
		}})
	}
	row.Add(ic)
	row.SetHit(scene.HitSpec{Label: "typography-editable"})
	return row
}

var _ = theme.Default
