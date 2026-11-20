package widgets

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// ── CheckboxGroup / RadioGroup（M8 对齐 antd）───────────────────────
//
// antd：<Checkbox.Group options value onChange> /
//
//	<Radio.Group options value onChange>
//
// 组持有选中状态并为每个选项缓存持久子控件实例（跨重建保状态）。

// CheckboxGroup 是复选框组。
type CheckboxGroup struct {
	Options  []string
	Values   []string // 选中的选项
	Disabled bool
	OnChange func([]string)

	boxes map[string]*Checkbox
}

func (g *CheckboxGroup) has(v string) bool {
	for _, s := range g.Values {
		if s == v {
			return true
		}
	}
	return false
}

func (g *CheckboxGroup) toggle(v string, on bool) {
	out := make([]string, 0, len(g.Values)+1)
	for _, s := range g.Values {
		if s != v {
			out = append(out, s)
		}
	}
	if on {
		out = append(out, v)
	}
	g.Values = out
	if g.OnChange != nil {
		g.OnChange(out)
	}
}

// Build 实现 ui.Widget。
func (g *CheckboxGroup) Build(ctx *ui.BuildContext) *scene.Node {
	if g.boxes == nil {
		g.boxes = map[string]*Checkbox{}
	}
	row := scene.Flex(scene.Horizontal, 12).SetCrossAlign(scene.CrossCenter)
	for _, opt := range g.Options {
		cb, ok := g.boxes[opt]
		if !ok {
			cb = NewCheckbox(opt, g.has(opt))
			g.boxes[opt] = cb
			checked := opt
			cb.OnChange = func(on bool) { g.toggle(checked, on) }
		}
		cb.Label = opt
		cb.Disabled = g.Disabled
		row.Add(cb.Build(ctx))
	}
	return row
}

// RadioGroup 是单选组（互斥由组管理）。
// RadioOptionType 是选项呈现形态（antd optionType）。
type RadioOptionType uint8

const (
	// RadioOptionDefault 默认（圆点+文字）。
	RadioOptionDefault RadioOptionType = iota
	// RadioOptionButton 按钮组（antd optionType="button"）：相连方框、
	// 选中主色实底白字。
	RadioOptionButton
)

type RadioGroup struct {
	controlBase
	Options    []string
	Value      string
	Disabled   bool
	// OptionType 选项形态（默认圆点组）。
	OptionType RadioOptionType
	OnChange   func(string)

	radios map[string]*Radio
}

// Build 实现 ui.Widget。
func (g *RadioGroup) Build(ctx *ui.BuildContext) *scene.Node {
	if g.OptionType == RadioOptionButton {
		return g.buildButtons(ctx)
	}
	if g.radios == nil {
		g.radios = map[string]*Radio{}
	}
	row := scene.Flex(scene.Horizontal, 12).SetCrossAlign(scene.CrossCenter)
	for _, opt := range g.Options {
		r, ok := g.radios[opt]
		if !ok {
			r = NewRadio(opt, g.Value == opt)
			g.radios[opt] = r
			choice := opt
			r.OnChange = func(bool) {
				g.Value = choice
				for k, rr := range g.radios {
					rr.Checked = k == choice
					rr.MarkDirty()
				}
				if g.OnChange != nil {
					g.OnChange(choice)
				}
			}
		}
		r.Label = opt
		r.Checked = g.Value == opt
		r.Disabled = g.Disabled
		row.Add(r.Build(ctx))
	}
	return row
}

// buildButtons 构建按钮组形态（antd optionType="button"）：相连方框
// （负 margin 抵消双边框）、选中主色实底白字。整框可点（HitOpaque+
// 手势按索引分流——Tag 轮结论复用）。
func (g *RadioGroup) BuildButtons(ctx *ui.BuildContext) *scene.Node { return g.buildButtons(ctx) }

func (g *RadioGroup) buildButtons(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	row := scene.Flex(scene.Horizontal, 0)
	for i, opt := range g.Options {
		selected := g.Value == opt
		fill := th.Color.Surface
		fg := th.Color.Text
		if selected {
			fill, fg = th.Color.Primary, th.Color.OnPrimary
		} else if i > 0 {
			// 相邻框左边框去重（相连成组）。
			_ = i
		}
		font := text.DefaultFont(th.Font.SizeBase)
		layout := text.LayoutText(opt, font, text.Options{})
		w := layout.Size.W + th.Space.MD*2
		h := ui.MetricsFor(th, ui.SizeMedium).Height
		n := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(fill),
			Border: scene.Stroke(th.Color.Border, 1),
		}).SetSize2(w, h)
		t := scene.Text(scene.TextSpec{Src: opt, Layout: layout, Color: fg})
		t.SetMargin(geom.InLTRB((w-layout.Size.W)/2, (h-layout.Size.H)/2, 0, 0))
		n.Add(t)
		anchor := opt // 锚=选项文本（可定位）
		_ = anchor
		n.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "radio-button", Cursor: scene.CursorPointer})
		n.SetKey("radio-btn-" + opt)
		if ctx.Input != nil {
			choice := opt
			g2 := &ui.GestureDetector{OnTap: func(geom.Point) {
				if g.Disabled || g.Value == choice {
					return
				}
				g.Value = choice
				g.MarkDirty() // controlBase 标准重建链（InvalidateRoot 注入）
				if g.OnChange != nil {
					g.OnChange(choice)
				}
			}}
			ctx.Input.Set(n, g2.Handler())
			g.attach(ctx, n, g2.Handler(), false)
		}
		row.Add(n)
	}
	return row
}
