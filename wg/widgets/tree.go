package widgets

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// ── Tree：可展开树 ────────────────────────────────────────────────

// TreeNode 是树节点。Children 为空的节点是叶子。
type TreeNode struct {
	Key      string
	Label    string
	Children []TreeNode
}

// Tree 是可展开树。Expanded 存展开节点的 Key 集合；Selected 为选中 Key。
type Tree struct {
	controlBase
	Nodes    []TreeNode
	Expanded map[string]bool
	Selected string
	Indent   float32
	RowH     float32
	Width    float32
	OnSelect func(key string)
	OnToggle func(key string, expanded bool)

	taps []*ui.GestureDetector
}

// NewTree 创建树。
func NewTree(nodes []TreeNode) *Tree {
	return &Tree{
		Nodes: nodes, Expanded: make(map[string]bool),
		Indent: 18, RowH: 28, Width: 240,
	}
}

// Toggle 展开/收起节点。
func (t *Tree) Toggle(key string) {
	next := !t.Expanded[key]
	t.Expanded[key] = next
	t.MarkDirty()
	if t.OnToggle != nil {
		t.OnToggle(key, next)
	}
}

// Build 实现 ui.Widget。
func (t *Tree) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(t)
	th := ctx.ThemeOr()
	col := scene.Flex(scene.Vertical, 0).SetWidth(t.Width)
	t.taps = t.taps[:0]

	var addNodes func(nodes []TreeNode, depth int)
	addNodes = func(nodes []TreeNode, depth int) {
		for i := range nodes {
			n := nodes[i]
			row := scene.Flex(scene.Horizontal, 4).SetHeight(t.RowH).SetWidth(t.Width)
			hasKids := len(n.Children) > 0
			arrow := ""
			if hasKids {
				if t.Expanded[n.Key] {
					arrow = "▾"
				} else {
					arrow = "▸"
				}
			}
			// 缩进 + 展开箭头
			prefix := ""
			for k := 0; k < depth; k++ {
				prefix += "  "
			}
			prefix += arrow
			pl := text.LayoutText(prefix, text.DefaultFont(th.Font.SizeSM), text.Options{})
			pt := scene.Text(scene.TextSpec{Src: prefix, Layout: pl, Color: th.Color.TextMuted})
			pt.SetMargin(geom.InLTRB(float32(depth)*t.Indent+2, (t.RowH-pl.Size.H)/2, 0, 0))
			row.Add(pt)

			ll := text.LayoutText(n.Label, text.DefaultFont(th.Font.SizeSM), text.Options{})
			fg := th.Color.Text
			if n.Key == t.Selected {
				fg = th.Color.Primary
				row.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Selection)})
			}
			lt := scene.Text(scene.TextSpec{Src: n.Label, Layout: ll, Color: fg})
			lt.SetMargin(geom.InLTRB(0, (t.RowH-ll.Size.H)/2, 0, 0))
			row.Add(lt)

			if ctx.Input != nil && (hasKids || t.OnSelect != nil) {
				row.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "tree-node", Cursor: scene.CursorPointer})
				key, kids := n.Key, hasKids
				g := &ui.GestureDetector{OnTap: func(geom.Point) {
					if kids {
						t.Toggle(key)
					}
					if t.OnSelect != nil {
						t.Selected = key
						t.MarkDirty()
						t.OnSelect(key)
					}
				}}
				t.taps = append(t.taps, g)
				ctx.Input.Set(row, g.Handler())
			}
			col.Add(row)
			if hasKids && t.Expanded[n.Key] {
				addNodes(n.Children, depth+1)
			}
		}
	}
	addNodes(t.Nodes, 0)
	t.node = col
	return col
}

// ── Form / FormItem ───────────────────────────────────────────────

// FormItem 是表单项：标签 + 控件 + 校验提示。
type FormItem struct {
	Label    string
	Child    ui.Widget
	Required bool
	Error    string
	LabelW   float32
}

// Build 实现 ui.Widget。
func (f FormItem) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	row := scene.Flex(scene.Horizontal, th.Space.SM)
	labelW := f.LabelW
	if labelW <= 0 {
		labelW = 80
	}
	label := f.Label
	if f.Required {
		label = "* " + label
	}
	ll := text.LayoutText(label, text.DefaultFont(th.Font.SizeSM), text.Options{})
	lt := scene.Text(scene.TextSpec{Src: label, Layout: ll, Color: th.Color.TextMuted})
	lt.SetWidth(labelW).SetMargin(geom.InLTRB(0, 6, 0, 0))
	row.Add(lt)
	if f.Child != nil {
		row.Add(buildChild(ctx, f.Child))
	}
	col := scene.Flex(scene.Vertical, 2, row)
	if f.Error != "" {
		el := text.LayoutText(f.Error, text.DefaultFont(th.Font.SizeXS), text.Options{})
		et := scene.Text(scene.TextSpec{Src: f.Error, Layout: el, Color: th.Color.Danger})
		et.SetMargin(geom.InLTRB(labelW+th.Space.SM, 0, 0, 0))
		col.Add(et)
	}
	return col
}

// Validator 校验一个值，返回错误信息（空串表示通过）。
type Validator func(value string) string

// RequiredRule 返回必填校验器。
func RequiredRule(msg string) Validator {
	if msg == "" {
		msg = "此项必填"
	}
	return func(v string) string {
		if v == "" {
			return msg
		}
		return ""
	}
}

// Form 是表单容器。校验器按字段名注册，Validate 汇总错误。
type Form struct {
	ui.Base
	Items     []FormItem
	Values    map[string]string
	Validators map[string][]Validator
	Errors    map[string]string
}

// NewForm 创建表单。
func NewForm() *Form {
	return &Form{
		Values: make(map[string]string),
		Validators: make(map[string][]Validator),
		Errors: make(map[string]string),
	}
}

// Rule 为字段注册校验器。
func (f *Form) Rule(field string, v ...Validator) *Form {
	f.Validators[field] = append(f.Validators[field], v...)
	return f
}

// Validate 校验所有字段，返回是否通过（同时写入 Errors）。
func (f *Form) Validate() bool {
	f.Errors = make(map[string]string)
	ok := true
	for field, rules := range f.Validators {
		for _, v := range rules {
			if msg := v(f.Values[field]); msg != "" {
				f.Errors[field] = msg
				ok = false
				break
			}
		}
	}
	f.MarkDirty()
	return ok
}

// Build 实现 ui.Widget。
func (f *Form) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(f)
	th := ctx.ThemeOr()
	col := scene.Flex(scene.Vertical, th.Space.MD)
	for i := range f.Items {
		item := f.Items[i]
		if msg := f.Errors[item.Label]; msg != "" {
			item.Error = msg
		}
		col.Add(item.Build(ctx))
	}
	return col
}
