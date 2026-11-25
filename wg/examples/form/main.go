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
	form     = widgets.NewForm()
	formRule = widgets.NewForm()
	demoForm = widgets.NewDemoForm()
	result   = widgets.NewButton("点提交查看校验结果")
	// 规则集卡的提交结果按钮（与必填卡按钮同款文本协议，spec 文本断言用）。
	resultRule = widgets.NewButton("点提交查看校验结果")
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

	// 规则集卡（antd rules 直译）：required/min/email/url/pattern/
	// validator 全路径演示。
	formRule.Items = []widgets.FormItem{
		{Label: "用户名", Child: widgets.NewInput("2-10 字"), Required: true},
		{Label: "邮箱", Child: widgets.NewInput("name@example.com")},
		{Label: "主页", Child: widgets.NewInput("https://")},
		{Label: "验证码", Child: widgets.NewInput("4 位数字")},
		{Label: "密码", Child: widgets.NewInput("至少 6 位")},
	}
	formRule.Rules("用户名", widgets.FormRule{Required: true, Min: 2, Max: 10})
	formRule.Rules("邮箱", widgets.FormRule{Required: true, Email: true})
	formRule.Rules("主页", widgets.FormRule{URL: true})
	formRule.Rules("验证码", widgets.FormRule{Pattern: `^\d{4}$`, Message: "验证码为 4 位数字"})
	formRule.Rules("密码", widgets.FormRule{Validator: func(v string) string {
		if len([]rune(v)) < 6 {
			return "密码至少 6 位"
		}
		return ""
	}})
	resultRule.OnClick = func() {
		if formRule.Validate() {
			resultRule.Text = "校验通过"
		} else {
			resultRule.Text = fmt.Sprintf("校验失败：%v", formRule.Errors)
		}
		resultRule.MarkDirty()
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "form", Title: "Form 表单", Group: doc.GroupData,
		What: []string{
			"表单容器：FormItem{Label,Child,Required} 组装字段行，Validate 汇总校验",
			"过程式校验器（Rule/Validator）+ 声明式规则集（Rules/FormRule，antd rules 直译）",
			"FormRule：required/whitespace/len/min/max/pattern/email/url/validator+message",
			"集成表单页：NewDemoForm 组合 Input/Radio/Switch/Slider/Tabs/Tooltip/提交链路",
			"计划（M12 尾）：横向布局、动态增删字段",
		},
		Demos: []doc.Demo{
			{Title: "必填校验", Note: "用户名留空提交 → 报错；填任意值 → 通过", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 12)
				col.Add(form.Build(ctx))
				col.Add(result.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "规则集", Note: "声明式 FormRule：required+min/max、email、url、pattern（4 位数字）、validator 自定义（密码≥6 位）", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 12)
				col.Add(formRule.Build(ctx))
				col.Add(resultRule.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "集成表单页", Note: "NewDemoForm：中文输入、Radio、Switch、Slider、Tabs、Tooltip 与提交状态联动", Build: func(ctx *ui.BuildContext) ui.Widget {
				return demoForm
			}},
		},
		API: []doc.APIRow{
			{Name: "NewForm", Type: "func() *Form", Desc: "构造"},
			{Name: "Items", Type: "[]FormItem{Label,Child,Required,Error,LabelW}", Desc: "字段行"},
			{Name: "Values", Type: "map[string]string", Desc: "字段名 → 值"},
			{Name: "Rule", Type: "func(field string, v ...Validator)", Desc: "按字段注册过程式校验器"},
			{Name: "Rules", Type: "func(field string, r ...FormRule)", Desc: "按字段注册声明式规则集（antd rules 直译）"},
			{Name: "FormRule", Type: "struct{Required,Whitespace,Len,Min,Max,Pattern,Email,URL,Validator,Message}", Desc: "声明式规则（首个失败即停）"},
			{Name: "Validate", Type: "func() bool", Desc: "全部校验；失败信息进 Errors"},
			{Name: "Errors", Type: "map[string]string", Desc: "校验失败信息"},
			{Name: "RequiredRule", Type: "func(msg string) Validator", Desc: "非空校验器"},
		},
	})
}
