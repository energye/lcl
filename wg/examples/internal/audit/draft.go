// draft 生成器：逐屏枚举可交互元素，按类型映射出用例初稿（规范 §9）。
//
// 用途：为尚未写 L2/L3 用例的窗快速产出 spec 初稿，人工校准（补 demo 分组、
// 文本锚、特有断言）后并入 wg/examples/<id>/spec.json。draft 只覆盖有
// 标准模板的类型；button 等无模板元素继续由启发式兜底。
package audit

import (
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/scene"
)

// templateFor 返回标签的初稿模板（空 = 无模板）。
func templateFor(label string, disabled bool) string {
	switch label {
	case "switch", "checkbox":
		if disabled {
			return "disabled-no-op"
		}
		return "toggle"
	case "input", "input-number":
		if disabled {
			return "disabled-no-op"
		}
		return "input-focus"
	case "select":
		if disabled {
			return "disabled-no-op"
		}
		return "select-open-close"
	case "table-row":
		return "table-row-select"
	case "slider":
		return "slider-jump"
	case "rate-star":
		return "rate-set"
	case "segment", "tab", "page":
		return "segment-select"
	case "tree-node":
		return "tree-node-select"
	case "collapse":
		return "collapse-toggle"
	}
	return ""
}

// DraftSpec 逐屏枚举并生成用例初稿。每类元素最多两条（非禁用 + 禁用各一）。
func DraftSpec(c *wg.Container) *Spec {
	spec := &Spec{
		Sections: []SpecSection{{
			Name: "draft（校准后按 demo 标题分组）",
		}},
	}
	seen := map[string]bool{} // label|disabled
	labelIdx := map[string]int{}

	scrollPos := func() float32 {
		var off, area float32
		c.Scene().Walk(func(n *scene.Node) bool {
			o := n.ScrollOffset()
			if o == 0 {
				return true
			}
			if b := n.BoundsInRoot(); b.W()*b.H() > area {
				area = b.W() * b.H()
				off = o
			}
			return true
		})
		return off
	}

	// 回顶。
	prev := float32(-1)
	for i := 0; i < 40; i++ {
		off := scrollPos()
		if off == prev {
			break
		}
		prev = off
		injectWheel(c, 1035, 400, 3)
		time.Sleep(110 * time.Millisecond)
	}
	time.Sleep(200 * time.Millisecond)

	// 逐屏枚举。
	prev = -1
	for page := 0; page < 30; page++ {
		off := scrollPos()
		if off == prev {
			break
		}
		prev = off
		vh := c.Scene().Size().H
		rowSeen := map[uint64]bool{}
		c.Scene().Walk(func(n *scene.Node) bool {
			h := n.HitSpec()
			if _, ok := interactiveLabels[h.Label]; !ok {
				return true
			}
			b := n.BoundsInRoot()
			if b.IsEmpty() || b.Max.Y <= 48 || b.Min.Y >= vh-48 {
				return true
			}
			if h.Label == "table-row" && n.Parent() != nil {
				pid := n.Parent().ID()
				if rowSeen[pid] {
					return true
				}
				rowSeen[pid] = true
			}
			k := h.Label + "|" + map[bool]string{true: "1", false: "0"}[h.Disabled]
			if seen[k] {
				return true
			}
			tpl := templateFor(h.Label, h.Disabled)
			if tpl == "" {
				return true
			}
			seen[k] = true
			idx := labelIdx[h.Label]
			labelIdx[h.Label]++
			spec.Sections[0].Cases = append(spec.Sections[0].Cases, SpecCase{
				Name:   h.Label + "（" + tpl + "，待校准）",
				Target: &SpecTarget{Label: h.Label, Disabled: boolPtr(h.Disabled), Index: idx},
				Then:   []SpecAssertion{{Kind: "template", Template: tpl}},
			})
			return true
		})
		injectWheel(c, 1035, 400, -3)
		time.Sleep(220 * time.Millisecond)
	}
	return spec
}

func boolPtr(v bool) *bool { return &v }