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
	tpBase   = widgets.NewTimePicker()
	tpInit   = widgets.NewTimePicker()
	tp12     = widgets.NewTimePicker()
	tpStep   = widgets.NewTimePicker()
	tpInit12 = widgets.NewTimePicker()
	tpDis    = widgets.NewTimePicker()
	tpFmt    = widgets.NewTimePicker()
	tpRange  = widgets.NewTimeRangePicker()
)

func init() {
	tpInit.Value = time.Date(0, 1, 1, 9, 30, 0, 0, time.Local)
	// 12 小时制卡（antd use12Hours）：4 列面板+上午/下午，触发框
	// hh:mm:ss 上午/下午。初值 15:20:00=下午 03:20:00。
	tp12.Use12Hours = true
	tpInit12.Use12Hours = true
	tpInit12.Width = 180
	tpInit12.Value = time.Date(0, 1, 1, 15, 20, 0, 0, time.Local)
	// 区间时间卡（antd TimePicker.RangePicker）：双触发框+双面板并排。
	tpRange.Start.Value = time.Date(0, 1, 1, 9, 0, 0, 0, time.Local)
	tpRange.End.Value = time.Date(0, 1, 1, 18, 30, 0, 0, time.Local)

	// 列裁剪卡（antd format hide-column）：HH:mm 省略秒→秒列消失。
	tpFmt.Format = "15:04"
	tpFmt.Value = time.Date(0, 1, 1, 14, 25, 0, 0, time.Local)

	// 禁用时段卡（antd disabledTime）：夜间 22:00-05:59 不可选，且选
	// 09 点时 30-59 分不可选（联动 disabledMinutes）。
	tpDis.Value = time.Date(0, 1, 1, 9, 10, 0, 0, time.Local)
	tpDis.DisabledTime = func(time.Time) widgets.DisabledTimeSpec {
		return widgets.DisabledTimeSpec{
			DisabledHours: func() []int {
				return []int{22, 23, 0, 1, 2, 3, 4, 5}
			},
			DisabledMinutes: func(h int) []int {
				if h == 9 {
					return []int{30, 45}
				}
				return nil
			},
		}
	}

	// 步进卡（antd hourStep/minuteStep/secondStep）：列只渲染步长倍数。
	tpStep.HourStep, tpStep.MinuteStep, tpStep.SecondStep = 2, 10, 15
}

func main() {
	launcher.Main(doc.Meta{
		ID: "timepicker", Title: "TimePicker 时间选择", Group: doc.GroupControl,
		What: []string{
			"触发框（Input 形态+clock 图标）→ 点击弹三列面板（时/分/秒，列 56px、cell 28px、Scroll 组合复用）",
			"选择高亮：选中主色实底白字、hover 浅底（SetBox 原地迁移零重建）；底部「现在/确定」",
			"antd needConfirm 默认形态：选择不关面板（onChange 已触发），「确定」/Esc/外部点击才关闭",
			"**use12Hours** 12 小时制：上午/下午第 4 列+小时列 12 格（12,01..11），触发框 hh:mm:ss 上午/下午",
			"hourStep/minuteStep/secondStep：列只渲染步长倍数（antd hourStep 等）",
		},
		Demos: []doc.Demo{
			{Title: "基础", Note: "点击触发框弹面板；点选时刻高亮迁移；「现在」填当前时刻；「确定」关闭", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tpBase
			}},
			{Title: "带初值", Note: "Value=09:30:00：打开面板时三列滚动定位到选中项", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tpInit
			}},
			{Title: "12 小时制", Note: "Use12Hours：初值 15:20=下午 03:20:00；面板 4 列（时/分/秒/上午下午）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tpInit12
			}},
			{Title: "区间时间", Note: "RangePicker：起~止双触发框，点开双面板并排，确定提交区间（起>止自动交换）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tpRange
			}},
			{Title: "列裁剪", Note: "Format=15:04（HH:mm）：秒列消失；触发框只显示 时:分", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tpFmt
			}},
			{Title: "禁用时段", Note: "DisabledTime：22-05 时不可选、9 点的 30/45 分不可选（灰置 time-cell-disabled）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tpDis
			}},
			{Title: "步进", Note: "HourStep=2/MinuteStep=10/SecondStep=15：列只渲染步长倍数（时列 24→12 格）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return tpStep
			}},
		},
		API: []doc.APIRow{
			{Name: "NewTimeRangePicker", Type: "func() *TimeRangePicker", Desc: "区间时间选择：Start/End 双端 + OnChangeRange(start,end)"},
			{Name: "Format", Type: "string", Default: "15:04:05", Desc: "时间格式（Go 布局）；省略分/秒则对应列消失（antd hide-column）"},
			{Name: "DisabledTime", Type: "func(now time.Time) DisabledTimeSpec", Desc: "禁用时段（disabledHours/disabledMinutes(h)/disabledSeconds(h,m)），禁用格灰置不可点"},
			{Name: "NewTimePicker", Type: "func() *TimePicker", Desc: "构造（默认无值显示占位）"},
			{Name: "Value", Type: "time.Time", Desc: "选中时刻（仅时分秒参与显示与选择）"},
			{Name: "Use12Hours", Type: "bool", Default: "false", Desc: "12 小时制（上午/下午第 4 列，触发框 hh:mm:ss 上午/下午）"},
			{Name: "HourStep/MinuteStep/SecondStep", Type: "int", Default: "1", Desc: "列值步长（列只渲染步长倍数）"},
			{Name: "OnChange", Type: "func(t time.Time)", Desc: "选择/「现在」回调（antd 选择即触发，面板保持打开）"},
			{Name: "Placeholder", Type: "string", Default: "请选择时间", Desc: "占位文本"},
			{Name: "Width", Type: "float32", Default: "120", Desc: "触发框宽度"},
			{Name: "Show / Close / IsOpen", Type: "func(...)", Desc: "程序化开关面板 / 状态查询"},
		},
	})
}
