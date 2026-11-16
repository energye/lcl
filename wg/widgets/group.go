package widgets

import (
	"github.com/energye/lcl/wg/scene"
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
type RadioGroup struct {
	Options  []string
	Value    string
	Disabled bool
	OnChange func(string)

	radios map[string]*Radio
}

// Build 实现 ui.Widget。
func (g *RadioGroup) Build(ctx *ui.BuildContext) *scene.Node {
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
