package main

import (
	"time"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	dp          = widgets.NewDatePicker()
	dpRange     = widgets.NewDatePicker()
	dpCell      = widgets.NewDatePicker()
	rangeInited bool
)

func init() {
	// 自定义单元格卡（antd cellRender 对应）：面板日期格数字下方追加
	// Badge（11/15/27 日挂事件）；Now 口径缺位——面板无 Now 注入，
	// 打开时按 Value 定月（Value 固定 2026-11-21 保证跨天确定性）。
	dpCell.Value = time.Date(2026, 11, 21, 12, 0, 0, 0, time.Local)
	dpCell.DateCellRender = func(date time.Time) ui.Widget {
		if date.Month() != time.November || date.Year() != 2026 {
			return nil
		}
		switch date.Day() {
		case 11:
			return &widgets.Badge{Text: "评审", Color: geom.RGBA(82, 196, 26)}
		case 15:
			return &widgets.Badge{Text: "发布", Color: geom.RGBA(250, 140, 22)}
		case 27:
			return &widgets.Badge{Text: "停机", Color: geom.RGBA(255, 77, 79)}
		}
		return nil
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "datepicker", Title: "DatePicker 日期选择", Group: doc.GroupControl,
		What: []string{
			"触发框（Input 形态+日历图标）→ 点击弹日历面板（BandPopup，Dropdown 同通道）",
			"面板：月标题+‹›切换（跨年自动进退位）+星期行+6×7 日期格（Grid 组合复用）",
			"日期格选择：主色实底选中态；today 主色文字；选择后自动关闭面板",
			"DateCellRender：面板日期格数字下方追加自定义内容（antd cellRender 对应）",
		},
		Demos: []doc.Demo{
			{Title: "日期选择", Note: "点击触发框弹面板；‹›切换月份；点击日期格选择并关闭", Build: func(ctx *ui.BuildContext) ui.Widget {
				return dp
			}},
			{Title: "范围选择", Note: "Range=true：双触发框，面板两击选起止（第二击早于第一击自动交换）", Build: func(ctx *ui.BuildContext) ui.Widget {
				if !rangeInited {
					dpRange.Range = true
					dpRange.Placeholder = "开始日期"
					dpRange.Placeholder2 = "结束日期"
					rangeInited = true
				}
				return dpRange
			}},
			{Title: "自定义单元格", Note: "DateCellRender：面板日期格数字下方追加事件 Badge（11 评审/15 发布/27 停机，antd cellRender 对应）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return dpCell
			}},
		},
		API: []doc.APIRow{
			{Name: "NewDatePicker", Type: "func() *DatePicker", Desc: "构造（默认当前日期）"},
			{Name: "Value", Type: "time.Time", Desc: "选中日期"},
			{Name: "OnChange", Type: "func(t time.Time)", Desc: "选择回调（选择后面板自动关闭）"},
			{Name: "Range", Type: "bool", Default: "false", Desc: "范围选择模式：双触发框，面板两击选起止"},
			{Name: "Value2 / OnChangeRange", Type: "time.Time / func(t1,t2)", Desc: "范围结束日期 / 范围选择回调（第二击早于第一击自动交换）"},
			{Name: "DateCellRender", Type: "func(date time.Time) ui.Widget", Desc: "自定义面板日期格内容（数字下方追加，antd cellRender 对应）"},
			{Name: "Placeholder / Placeholder2", Type: "string", Default: "请选择日期", Desc: "触发框占位文本（范围模式为起/止框）"},
			{Name: "Show / Close / IsOpen / RangeStage", Type: "func(...)", Desc: "程序化开关面板 / 状态查询（RangeStage：0=未开始 1=已选起待选止）"},
			{Name: "Width", Type: "float32", Default: "200", Desc: "触发框宽度（范围模式为单框宽）"},
		},
	})
}
