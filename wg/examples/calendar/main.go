package main

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var cal = widgets.NewCalendar()
var calLog = ""

// calEvents 固定事件表（DateCellRender 演示：antd Notice Calendar 形态
// ——日期数字仍在，自定义内容追加在数字下方）。
var calEvents = map[int]string{
	11: "里程碑评审",
	15: "发布窗口",
	27: "维护停机",
}

func init() {
	cal.Now = func() time.Time {
		return time.Date(2026, 11, 21, 12, 0, 0, 0, time.Local)
	}
	cal.OnSelect = func(date time.Time) {
		calLog = fmt.Sprintf("select %04d-%02d-%02d", date.Year(), date.Month(), date.Day())
	}
	cal.DateCellRender = func(date time.Time) ui.Widget {
		if date.Month() != time.November || date.Year() != 2026 {
			return nil
		}
		ev, ok := calEvents[date.Day()]
		if !ok {
			return nil
		}
		return &widgets.Badge{Text: ev, Color: geom.RGBA(82, 196, 26)}
	}
}

// calYear 年面板演示（Mode="year"：4×3 月格，点选切月）。
var calYear = func() *widgets.Calendar {
	c := widgets.NewCalendar()
	c.Mode = "year"
	c.Now = cal.Now
	c.Value = time.Date(2026, 11, 21, 12, 0, 0, 0, time.Local)
	c.SetView(2026, 11)
	return c
}()

// calRange validRange 演示：仅 2026-11-05 ~ 2026-11-25 可选（范围外
// 灰置；上月/下月切换钳制）。
var calRange = func() *widgets.Calendar {
	c := widgets.NewCalendar()
	c.Now = cal.Now
	c.SetView(2026, 11)
	c.ValidRange = [2]time.Time{
		time.Date(2026, 11, 5, 0, 0, 0, 0, time.Local),
		time.Date(2026, 11, 25, 0, 0, 0, 0, time.Local),
	}
	return c
}()

func main() {
	launcher.Main(doc.Meta{
		ID: "calendar", Title: "Calendar 日历", Group: doc.GroupData,
		What: []string{
			"整页日历：页头（年月大字+上月/今天/下月 Button 组合复用）+ 星期表头 + 6×7 日期格",
			"日期区 **Flex 行方案**（DatePicker 轮实证的零歧义通道）；today 主色描边、选中主色实底",
			"点击格 onSelect（down 即执行）；今天按钮回当月并选中今天",
			"DateCellRender：日期数字下方追加自定义内容（antd cellRender 对应，Notice 形态）",
			"Mode=year 年面板（4×3 月格点选切月，页头按钮切换）；ValidRange 范围外灰置+切换钳制",
		},
		Demos: []doc.Demo{
			{Title: "全屏日历", Note: "点击日期格选中（实底迁移）；上月/下月切换；今天回当月；11/15/27 日挂事件 Badge（DateCellRender）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return cal
			}},
			{Title: "年面板", Note: "Mode=year：4×3 月格，点选切到对应月（页头「月面板」按钮切回）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return calYear
			}},
			{Title: "有效范围", Note: "ValidRange=[11-05, 11-25]：范围外灰置不可选，上月/下月切换钳制", Build: func(ctx *ui.BuildContext) ui.Widget {
				return calRange
			}},
		},
		API: []doc.APIRow{
			{Name: "NewCalendar", Type: "func() *Calendar", Desc: "构造（默认当前月）"},
			{Name: "Value", Type: "time.Time", Desc: "选中日期（零值=未选）"},
			{Name: "OnSelect", Type: "func(date time.Time)", Desc: "点击日期格回调"},
			{Name: "OnPanelChange", Type: "func(month time.Time)", Desc: "面板月份切换回调"},
			{Name: "DateCellRender", Type: "func(date time.Time) ui.Widget", Desc: "自定义日期格内容（数字下方追加，antd cellRender 对应）"},
			{Name: "Mode", Type: "string", Default: "month", Desc: "面板模式（month 日历 / year 年面板）"},
			{Name: "ValidRange", Type: "[2]time.Time", Desc: "有效范围（范围外灰置不可选，切换钳制）"},
			{Name: "Width", Type: "float32", Default: "640", Desc: "总宽"},
		},
	})
}
