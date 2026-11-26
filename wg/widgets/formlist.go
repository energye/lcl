package widgets

// ── FormList（antd Form.List）───────────────────────────────────────
//
// 动态字段行容器：每行由 RowBuilder 生成一组 FormItem，支持 Add/Remove
// 增删行（antd Form.List operation.add / operation.remove）。
//
// 字段键按 `name[i]` 命名（antd name path 数组语义）：第 i 行第 k 个字段
// 键为 `Name[i].FieldK`，校验/取值按行独立。
//
// 未实现（如实标注）：move 上移/下移、行内动态子字段（Form.List 嵌套）、
// ErrorList（行级校验汇总展示）。

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// FormListRow 是一行动态字段的数据（键值对，键为行内字段名）。
type FormListRow struct {
	Values map[string]string
}

// FormList 是动态字段行容器。
type FormList struct {
	controlBase
	// Name 列表字段名（字段键前缀，antd Form.List name）。
	Name string
	// RowFields 行内字段名顺序（每行渲染的字段，antd children 里的 Form.Item name）。
	RowFields []string
	// RowLabels 行内字段标签（与 RowFields 一一对应；空 = 不显示标签）。
	RowLabels []string
	// Rows 当前行数据。
	Rows []*FormListRow
	// AddText 新增按钮文案（默认「+ 添加」）。
	AddText string
	// DeleteText 行删除按钮文案（默认「删除」）。
	DeleteText string

	// inputs 池化输入实例（跨重建保留编辑态/焦点，教训⑨）。
	inputs map[string]*Input
	// deletedTaps 池化删除手势。
	deletedTaps []*ui.GestureDetector
}

// NewFormList 创建动态字段行容器（初始 1 行）。
func NewFormList(name string, fields ...string) *FormList {
	l := &FormList{Name: name, RowFields: fields, inputs: map[string]*Input{},
		AddText: "+ 添加", DeleteText: "删除"}
	l.Add()
	return l
}

// Add 新增一行（antd operation.add）。
func (l *FormList) Add() {
	row := &FormListRow{Values: map[string]string{}}
	l.Rows = append(l.Rows, row)
	l.MarkDirty()
}

// Remove 删除第 i 行（antd operation.remove）；至少保留 0 行。
func (l *FormList) Remove(i int) {
	if i < 0 || i >= len(l.Rows) {
		return
	}
	l.Rows = append(l.Rows[:i], l.Rows[i+1:]...)
	l.MarkDirty()
}

// Len 行数。
func (l *FormList) Len() int { return len(l.Rows) }

// fieldKey 第 i 行第 field 字段的键（antd name path：Name[i].Field）。
func (l *FormList) fieldKey(i int, field string) string {
	return fmt.Sprintf("%s[%d].%s", l.Name, i, field)
}

// Values 汇总所有行字段值（键 = Name[i].Field）。
func (l *FormList) Values() map[string]string {
	out := map[string]string{}
	for i, row := range l.Rows {
		for _, f := range l.RowFields {
			out[l.fieldKey(i, f)] = row.Values[f]
		}
	}
	return out
}

// Build 实现 ui.Widget：逐行（字段 + 删除）+ 底部添加按钮。
func (l *FormList) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(l)
	th := ctx.ThemeOr()
	col := scene.Flex(scene.Vertical, th.Space.SM)
	col.SetHit(scene.HitSpec{Label: "form-list"})
	l.deletedTaps = l.deletedTaps[:0]

	for i, row := range l.Rows {
		i, row := i, row
		rowNode := scene.Flex(scene.Horizontal, th.Space.SM).SetCrossAlign(scene.CrossCenter)
		for k, f := range l.RowFields {
			key := l.fieldKey(i, f)
			in, ok := l.inputs[key]
			if !ok {
				in = NewInput("")
				in.Width = 140
				l.inputs[key] = in
			}
			in.Value = row.Values[f]
			field := f
			in.OnChanged = func(v string) {
				row.Values[field] = v
			}
			// 标签（可空）。
			if k < len(l.RowLabels) && l.RowLabels[k] != "" {
				lb := text.LayoutText(l.RowLabels[k], text.DefaultFont(th.Font.SizeBase), text.Options{})
				rowNode.Add(scene.Text(scene.TextSpec{Src: l.RowLabels[k], Layout: lb, Color: th.Color.TextMuted}))
			}
			rowNode.Add(in.Build(ctx))
		}
		// 行删除按钮（文本按钮）。
		del := scene.Text(scene.TextSpec{
			Src: l.DeleteText, Layout: text.LayoutText(l.DeleteText, text.DefaultFont(th.Font.SizeBase), text.Options{}),
			Color: th.Color.Danger,
		})
		del.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "form-list-remove", Cursor: scene.CursorPointer})
		if ctx.Input != nil {
			g := &ui.GestureDetector{OnTap: func(geom.Point) { l.Remove(i) }}
			l.deletedTaps = append(l.deletedTaps, g)
			ctx.Input.Set(del, g.Handler())
		}
		rowNode.Add(del)
		rowNode.SetHit(scene.HitSpec{Label: "form-list-row"})
		col.Add(rowNode)
	}

	// 添加按钮。
	add := scene.Text(scene.TextSpec{
		Src: l.AddText, Layout: text.LayoutText(l.AddText, text.DefaultFont(th.Font.SizeBase), text.Options{}),
		Color: th.Color.Primary,
	})
	add.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "form-list-add", Cursor: scene.CursorPointer})
	if ctx.Input != nil {
		g := &ui.GestureDetector{OnTap: func(geom.Point) { l.Add() }}
		l.deletedTaps = append(l.deletedTaps, g)
		ctx.Input.Set(add, g.Handler())
	}
	col.Add(add)
	return col
}

// 防止未使用导入告警（input 包在池化输入实例类型上使用）。
var _ input.Handler
