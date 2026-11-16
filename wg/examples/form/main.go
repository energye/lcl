package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	form   = widgets.NewForm()
	result = widgets.NewButton("点提交查看校验结果")
)

func init() {
	form.Items = []widgets.FormItem{
		{Label: "用户名", Child: widgets.NewInput("必填"), Required: true},
		{Label: "邮箱", Child: widgets.NewInput("name@example.com")},
	}
	form.Rule("用户名", widgets.RequiredRule("用户名不能为空"))
	result.OnClick = func() {
		if form.Validate() {
			result.Text = "校验通过"
		} else {
			result.Text = fmt.Sprintf("校验失败：%v", form.Errors)
		}
		result.MarkDirty()
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "form", Title: "Form 表单", Group: doc.GroupData,
		What: []string{
			"表单容器：FormItem{Label,Child,Required} 组装字段行，Validate 汇总校验",
			"校验器按字段名注册（Rule），失败信息收集到 Errors 展示在字段下方",
			"计划（M12）：校验规则集（email/len/pattern）、横向布局、动态增删字段",
		},
		Demos: []doc.Demo{
			{Title: "必填校验", Note: "用户名留空提交 → 报错；填任意值 → 通过", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 12)
				col.Add(form.Build(ctx))
				col.Add(result.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewForm", Type: "func() *Form", Desc: "构造"},
			{Name: "Items", Type: "[]FormItem{Label,Child,Required,Error,LabelW}", Desc: "字段行"},
			{Name: "Values", Type: "map[string]string", Desc: "字段名 → 值"},
			{Name: "Rule", Type: "func(field string, v ...Validator)", Desc: "按字段注册校验器"},
			{Name: "Validate", Type: "func() bool", Desc: "全部校验；失败信息进 Errors"},
			{Name: "Errors", Type: "map[string]string", Desc: "校验失败信息"},
			{Name: "RequiredRule", Type: "func(msg string) Validator", Desc: "非空校验器（当前唯一内置）"},
		},
	})
}
