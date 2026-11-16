package widgets

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// DemoForm 是一张只用 wg 组件搭出来的表单页。
//
// 它同时是 M3 的验收样本：12 个 P1 组件全部出现在一张真实可交互的页面里，
// 而且**零 lcl**——widgets 包的依赖由 checkdeps 静态守住。
// 页面不认识窗口、不认识 HWND，只认识 BuildContext 提供的能力。
type DemoForm struct {
	ui.Base

	Name   *Input
	City   *Input
	Agree  *Checkbox
	PlanA  *Radio
	PlanB  *Radio
	Dark   *Switch
	Level  *Slider
	List   *List
	Tabs   *Tabs
	Tip    *Tooltip
	Submit *Button

	Status string
}

// NewDemoForm 构造表单页。控件只在这里创建一次，重建时复用。
func NewDemoForm() *DemoForm {
	f := &DemoForm{}
	f.Name = NewInput("姓名")
	f.City = NewInput("城市")
	f.Agree = NewCheckbox("我已阅读并同意", false)
	f.PlanA = NewRadio("标准版", true)
	f.PlanB = NewRadio("专业版", false)
	f.Dark = NewSwitch(false)
	f.Level = NewSlider(0, 100, 30)
	f.List = NewList(
		ListItem{Key: "t1", Widget: Label{Text: "姓名用于账号显示"}},
		ListItem{Key: "t2", Widget: Label{Text: "城市用于本地化推荐"}},
	)
	f.Submit = NewButton("提交")
	f.Tip = &Tooltip{Text: "提交前请确认姓名与城市", Child: f.Submit, Placement: overlay.PlacementTop}

	f.PlanA.OnChange = func(on bool) {
		if on && f.PlanB.Checked {
			f.PlanB.Checked = false
			f.PlanB.MarkDirty()
		}
	}
	f.PlanB.OnChange = func(on bool) {
		if on && f.PlanA.Checked {
			f.PlanA.Checked = false
			f.PlanA.MarkDirty()
		}
	}
	f.Submit.OnClick = func() {
		name, city := f.Name.Value, f.City.Value
		switch {
		case name == "":
			f.Status = "请填写姓名"
		case city == "":
			f.Status = "请填写城市"
		default:
			plan := "标准版"
			if f.PlanB.Checked {
				plan = "专业版"
			}
			f.Status = fmt.Sprintf("已提交：%s / %s / %s / %d", name, city, plan, int(f.Level.Value))
		}
		f.MarkDirty()
	}

	f.Tabs = NewTabs(
		Tab{Key: "basic", Label: "基本信息", Content: Column{Gap: 8, Children: []ui.Widget{
			Padding{Insets: geom.InLTRB(16, 8, 16, 0), Child: Row{Gap: 8, Children: []ui.Widget{
				Label{Text: "姓名"}, f.Name,
			}}},
			Padding{Insets: geom.InLTRB(16, 0, 16, 0), Child: Row{Gap: 8, Children: []ui.Widget{
				Label{Text: "城市"}, f.City,
			}}},
			Padding{Insets: geom.InLTRB(16, 0, 16, 0), Child: Row{Gap: 16, Children: []ui.Widget{
				f.Agree, f.PlanA, f.PlanB,
			}}},
		}}},
		Tab{Key: "pref", Label: "偏好设置", Content: f.prefPage()},
	)
	return f
}

// Build 实现 ui.Widget。
func (f *DemoForm) Build(ctx *ui.BuildContext) *scene.Node {
	// 进度条必须跟上滑块当前值（它是无状态组件，每帧按当前值构建）。
	f.Tabs.Items[1].Content = f.prefPage()
	status := f.Status
	if status == "" {
		status = "填写表单后点击提交"
	}
	return Column{Gap: 12, Children: []ui.Widget{
		Padding{Insets: geom.InLTRB(16, 16, 16, 0), Child: Label{Text: "账号设置", Size: 22}},
		&Divider{Length: 320},
		f.Tabs,
		Row{Gap: 8, Children: []ui.Widget{
			Spacer{Width: 16, Height: 1},
			f.Tip,
			Spacer{Width: 8, Height: 1},
			Label{Text: status},
		}},
	}}.Build(ctx)
}

// prefPage 构造偏好页（滑块值每帧可能变，进度条按当前值重建）。
func (f *DemoForm) prefPage() ui.Widget {
	return Column{Gap: 8, Children: []ui.Widget{
		Padding{Insets: geom.InLTRB(16, 8, 16, 0), Child: Row{Gap: 8, Children: []ui.Widget{
			Label{Text: "深色模式"}, f.Dark,
		}}},
		Padding{Insets: geom.InLTRB(16, 0, 16, 0), Child: Row{Gap: 8, Children: []ui.Widget{
			Label{Text: "音量"}, f.Level, Progress{Value: f.Level.Value / 100, Width: 80},
		}}},
		Padding{Insets: geom.InLTRB(16, 0, 16, 0), Child: f.List},
	}}
}
