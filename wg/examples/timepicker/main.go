package main

import (
	"time"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有：demo Build 闭包里新建实例会在每次根重建时被
// 换成全新实例（选择状态归零、触发框永远占位）——教训⑨变体。
var (
	tpBase = widgets.NewTimePicker()
	tpInit = widgets.NewTimePicker()
)

func init() {
	tpInit.Value = time.Date(0, 1, 1, 9, 30, 0, 0, time.Local)
}

func main() {
	launcher.Main(doc.Meta{
		ID: "timepicker", Title: "TimePicker 时间选择", Group: doc.GroupControl,
		What: []string{
			"触发框（Input 形态+clock 图标）→ 点击弹三列面板（时/分/秒，列 56px、cell 28px、Scroll 组合复用）",
			"选择高亮：选中主色实底白字、hover 浅底（SetBox 原地迁移零重建）；底部「现在/确定」",
			"antd needConfirm 默认形态：选择不关面板（onChange 已触发），「确定」/Esc/外部点击才关闭",
		},
		Demos: []doc.Demo{
			{Title: "基础", Note: "点击触发框弹面板；点选时刻高亮迁移；「现在」填当前时刻；「确定」关闭", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tpBase
			}},
			{Title: "带初值", Note: "Value=09:30:00：打开面板时三列滚动定位到选中项", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tpInit
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTimePicker", Type: "func() *TimePicker", Desc: "构造（默认无值显示占位）"},
			{Name: "Value", Type: "time.Time", Desc: "选中时刻（仅时分秒参与显示与选择）"},
			{Name: "OnChange", Type: "func(t time.Time)", Desc: "选择/「现在」回调（antd 选择即触发，面板保持打开）"},
			{Name: "Placeholder", Type: "string", Default: "请选择时间", Desc: "占位文本"},
			{Name: "Width", Type: "float32", Default: "120", Desc: "触发框宽度"},
			{Name: "Show / Close / IsOpen", Type: "func(...)", Desc: "程序化开关面板 / 状态查询"},
		},
	})
}
