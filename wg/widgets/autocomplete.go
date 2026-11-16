package widgets

import (
	"strings"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── AutoComplete（M8 对齐 antd）─────────────────────────────────────
//
// antd：<AutoComplete options onSelect onChange>
// 已对齐：输入过滤（contains，大小写不敏感）、下拉建议、选中回填。
// 未实现：自定义 option 渲染、分类建议。

// AutoComplete 是自动完成输入框。
type AutoComplete struct {
	controlBase
	Suggestions []string
	Value       string
	Width       float32
	Disabled    bool
	Placeholder string
	OnSelect    func(string) // 选中建议
	OnChanged   func(string) // 文本变化

	handle overlay.Handle
	input  *Input
	taps   []*ui.GestureDetector
}

// NewAutoComplete 创建自动完成。
func NewAutoComplete(suggestions []string) *AutoComplete {
	in := NewInput("输入以过滤建议")
	in.Width = 240
	return &AutoComplete{Suggestions: suggestions, input: in, Width: 240}
}

// filtered 过滤后的建议（contains，大小写不敏感；空词显示全部）。
func (a *AutoComplete) filtered() []string {
	needle := strings.ToLower(a.input.Value)
	out := make([]string, 0, len(a.Suggestions))
	for _, s := range a.Suggestions {
		if needle == "" || strings.Contains(strings.ToLower(s), needle) {
			out = append(out, s)
		}
	}
	return out
}

// Build 实现 ui.Widget。
func (a *AutoComplete) Build(ctx *ui.BuildContext) *scene.Node {

	a.input.Width = a.Width
	a.input.Placeholder = a.Placeholder
	a.input.Disabled = a.Disabled
	a.input.Value = a.Value
	a.input.OnChanged = func(s string) {
		a.Value = s
		a.MarkDirty()
		if a.OnChanged != nil {
			a.OnChanged(s)
		}
		a.suggest(ctx)
	}
	n := a.input.Build(ctx)
	a.attach(ctx, n, input.Handler{}, false)
	// 输入框的处理器已被 attach 覆盖？不——Input 自己 attach 了自己的
	// 处理器；AutoComplete 只需要在文本变化时弹层，OnChanged 已接好。
	return n
}

// suggest 展示过滤建议弹层。
func (a *AutoComplete) suggest(ctx *ui.BuildContext) {
	th := ctx.ThemeOr()
	if ctx.Overlay == nil || a.Disabled || a.input.node == nil {
		return
	}
	// 关掉上一轮
	if a.handle != 0 {
		ctx.Overlay.Close(a.handle)
		a.handle = 0
	}
	items := a.filtered()
	if len(items) == 0 {
		return
	}
	const rowH = float32(30)
	h := float32(len(items)) * rowH
	if h > 220 {
		h = 220
	}
	a.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(a.input.node),
		Placement: overlay.PlacementBottomStart,
		Size:      geom.Size{W: a.input.Width, H: h},
		Dismiss:   overlay.Dismiss(0),
		Tag:       "autocomplete",
		Build: func() *scene.Node {
			col := scene.Flex(scene.Vertical, 0).SetWidth(a.input.Width)
			col.SetOverflow(true)
			for _, item := range items {
				row := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface)}).
					SetWidth(a.input.Width).SetHeight(rowH)
				l := text.LayoutText(item, text.DefaultFont(th.Font.SizeSM), text.Options{})
				t := scene.Text(scene.TextSpec{Src: item, Layout: l, Color: th.Color.Text})
				t.SetMargin(geom.InLTRB(11, (rowH-l.Size.H)/2, 0, 0))
				row.Add(t)
				row.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "autocomplete-item", Cursor: scene.CursorPointer})
				choice := item
				og := &ui.GestureDetector{OnTap: func(geom.Point) {
					a.Value = choice
					a.input.Value = choice
					a.input.Editor().SetValue(inputEditing(choice))
					a.MarkDirty()
					a.input.MarkDirty()
					if a.OnSelect != nil {
						a.OnSelect(choice)
					}
					if ctx.Overlay != nil && a.handle != 0 {
						ctx.Overlay.Close(a.handle)
					}
					a.handle = 0
				}}
				a.taps = append(a.taps, og)
				if ctx.Input != nil {
					ctx.Input.Set(row, og.Handler())
				}
				col.Add(row)
			}
			return col
		},
	})
}

// inputEditing 构造编辑值（光标在末尾）。
func inputEditing(s string) input.EditingValue {
	return input.EditingValue{Text: s, Selection: input.RangeAt(input.RuneLen(s))}
}
