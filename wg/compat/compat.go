// Package compat 提供旧 wg 风格的 API 门面，转发到新组件。
//
// 旧示例（energy/lcl/wg）用 TButton/TTab/TListBox 这类命名；新引擎用
// Button/Tabs/List。这一层只做**旧签名 → 新组件**的单向适配，让已有页面
// 能逐步迁移，不反向污染新 API。
//
// 门面刻意做得薄：不引入新概念、不持有像素，返回的就是新组件实例。
// 行为对照由 compat_test.go 保证（旧式调用产出的节点与新式一致）。
package compat

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func paddingInsets(w, h float32) geom.Insets {
	return geom.InLTRB(w, h, w, h)
}

// TButton 是旧式按钮构造器，转发到 widgets.NewButton。
func TButton(text string) *widgets.Button { return widgets.NewButton(text) }

// TCheckbox 是旧式复选框构造器。
func TCheckbox(label string, checked bool) *widgets.Checkbox {
	return widgets.NewCheckbox(label, checked)
}

// TRadio 是旧式单选框构造器。
func TRadio(label string, checked bool) *widgets.Radio {
	return widgets.NewRadio(label, checked)
}

// TSwitch 是旧式开关构造器。
func TSwitch(on bool) *widgets.Switch { return widgets.NewSwitch(on) }

// TInput 是旧式输入框构造器。
func TInput(placeholder string) *widgets.Input { return widgets.NewInput(placeholder) }

// TProgress 是旧式进度条构造器。
func TProgress(value float32) widgets.Progress {
	return widgets.Progress{Value: value}
}

// TLabel 是旧式标签构造器。
func TLabel(text string) widgets.Label { return widgets.Label{Text: text} }

// TCard 是旧式卡片构造器。
func TCard(content ui.Widget) *widgets.Card { return widgets.NewCard(content) }

// TList 是旧式列表构造器（转发到 KeyedList 语义的 List）。
func TList(items ...widgets.ListItem) *widgets.List { return widgets.NewList(items...) }

// TSelect 是旧式下拉框构造器。
func TSelect(options []string) *widgets.Select { return widgets.NewSelect(options) }

// TTable 是旧式表格构造器。
func TTable(cols []widgets.TableColumn, rows [][]any) *widgets.Table {
	return widgets.NewTable(cols, rows)
}

// TTree 是旧式树构造器。
func TTree(nodes []widgets.TreeNode) *widgets.Tree { return widgets.NewTree(nodes) }

// TModal 是旧式模态框构造器。
func TModal(title string, content ui.Widget) *widgets.Modal {
	return widgets.NewModal(title, content)
}

// TDrawer 是旧式抽屉构造器。
func TDrawer(title string, content ui.Widget) *widgets.Drawer {
	return widgets.NewDrawer(title, content)
}

// TNotification 是旧式通知构造器。
func TNotification(title, body string) *widgets.Notification {
	return widgets.NewNotification(title, body)
}

// TPopconfirm 是旧式确认气泡构造器。
func TPopconfirm(title string, child ui.Widget) *widgets.Popconfirm {
	return widgets.NewPopconfirm(title, child)
}

// TPagination 是旧式分页器构造器。
func TPagination(page, total int) *widgets.Pagination {
	return widgets.NewPagination(page, total)
}

// TCollapse 是旧式折叠面板构造器。
func TCollapse(panels ...widgets.CollapsePanel) *widgets.Collapse {
	return widgets.NewCollapse(panels...)
}

// TSegmented 是旧式分段选择器构造器。
func TSegmented(options ...string) *widgets.Segmented {
	return widgets.NewSegmented(options...)
}

// TRow / TColumn 是旧式布局构造器。
func TRow(children ...ui.Widget) widgets.Row  { return widgets.Row{Children: children} }
func TColumn(children ...ui.Widget) widgets.Column { return widgets.Column{Children: children} }

// TSpacer 是旧式占位构造器。
func TSpacer(w, h float32) widgets.Spacer { return widgets.Spacer{Width: w, Height: h} }

// TPadding 是旧式内边距构造器。
func TPadding(w, h float32, child ui.Widget) widgets.Padding {
	return widgets.Padding{Insets: paddingInsets(w, h), Child: child}
}

// AsNode 把任意组件按给定 context 构建成场景节点（旧式页面直接挂载用）。
func AsNode(ctx *ui.BuildContext, w ui.Widget) *scene.Node { return w.Build(ctx) }
