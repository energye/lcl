package examples

import (
	"fmt"

	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func init() {
	{
		cols := []widgets.TableColumn{
			{Title: "字段", Mode: widgets.ColumnAuto, Sortable: true},
			{Title: "类型", Mode: widgets.ColumnAuto},
			{Title: "主键", Mode: widgets.ColumnFixed, Width: 60},
			{Title: "说明", Mode: widgets.ColumnFlex, Flex: 1},
		}
		rows := [][]any{
			{"id", "INTEGER", "是", "主键"},
			{"name", "TEXT", "否", "名称"},
			{"created_at", "DATETIME", "否", "创建时间"},
			{"owner_id", "INTEGER", "否", "外键 → user.id"},
		}
		tbl := widgets.NewTable(cols, rows)
		tbl.Width, tbl.Height = 620, 180
		Register(Example{ID: "table", Name: "Table 表格", Group: GroupData,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				return Section(ctx, "Table", "Q9 三档列宽混用 + 点击表头排序 + 行虚拟化", tbl.Build(ctx))
			}})
	}

	{
		tree := widgets.NewTree([]widgets.TreeNode{
			{Key: "src", Label: "src", Children: []widgets.TreeNode{
				{Key: "render", Label: "render", Children: []widgets.TreeNode{
					{Key: "canvas", Label: "canvas.go"},
				}},
				{Key: "scene", Label: "scene"},
			}},
			{Key: "widgets", Label: "widgets"},
		})
		Register(Example{ID: "tree", Name: "Tree 树", Group: GroupData,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				return Section(ctx, "Tree", "点击箭头展开/收起，点击标签选中", tree.Build(ctx))
			}})
	}

	{
		sel := widgets.NewSelect([]string{"软件光栅化", "GDI 混合", "自动选择"})
		sel.Width = 260
		status := widgets.NewButton("尚未选择")
		sel.OnChange = func(i int) {
			status.Text = fmt.Sprintf("已选：%d · %s", i, sel.Options[i])
			status.MarkDirty()
		}
		Register(Example{ID: "select", Name: "Select 下拉", Group: GroupData,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				col := scene.Flex(scene.Vertical, 10)
				col.Add(sel.Build(ctx))
				col.Add(status.Build(ctx))
				return Section(ctx, "Select", "下拉列表走 overlay.Host，OutsideClick / Esc 关闭", col)
			}})
	}

	{
		form := widgets.NewForm()
		form.Items = []widgets.FormItem{
			{Label: "用户名", Child: widgets.NewInput("必填"), Required: true},
			{Label: "邮箱", Child: widgets.NewInput("name@example.com")},
		}
		form.Rule("用户名", widgets.RequiredRule("用户名不能为空"))
		result := widgets.NewButton("点提交查看校验结果")
		Register(Example{ID: "form", Name: "Form 表单校验", Group: GroupData,
			Build: func(ctx *ui.BuildContext) *scene.Node {
				result.OnClick = func() {
					ok := form.Validate()
					if ok {
						result.Text = "校验通过"
					} else {
						result.Text = fmt.Sprintf("校验失败：%v", form.Errors)
					}
					result.MarkDirty()
				}
				col := scene.Flex(scene.Vertical, 12)
				col.Add(form.Build(ctx))
				col.Add(result.Build(ctx))
				return Section(ctx, "Form / FormItem", "Validator 汇总到 Errors；Required 项为空时拒绝提交", col)
			}})
	}
}
