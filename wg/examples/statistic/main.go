package main

import (
	"time"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var (
	stAmount = func() *widgets.Statistic {
		s := widgets.NewStatistic("账户总余额", 189345.25)
		s.Precision = 2
		s.Prefix = "¥"
		s.Width = 160
		return s
	}()
	stUsers = func() *widgets.Statistic {
		s := widgets.NewStatistic("活跃用户", 112893)
		s.Suffix = "人"
		s.Width = 160
		return s
	}()
	// cdDay / cdNow 注入固定基准（跨天确定性）：目标=基准+2h30m，初始
	// 显示恒 02:30:00；cdReal 真实时钟正计时（Timer countup 形态）。
	cdBase = time.Date(2026, 11, 22, 12, 0, 0, 0, time.Local)
	cdDay  = func() *widgets.Countdown {
		c := widgets.NewCountdown("限时抢购", cdBase.Add(26*time.Hour+5*time.Second))
		c.Now = func() time.Time { return cdBase }
		c.Format = "D 天 HH:mm:ss"
		c.Suffix = "截止"
		c.Width = 220
		return c
	}()
	cdNow = func() *widgets.Countdown {
		c := widgets.NewCountdown("距开售", cdBase.Add(2*time.Hour+30*time.Minute))
		c.Now = func() time.Time { return cdBase }
		c.Width = 220
		return c
	}()
	cdReal = func() *widgets.Countdown {
		c := widgets.NewCountdown("运行时长", time.Now())
		c.CountUp = true
		c.Width = 220
		return c
	}()
)

func main() {
	launcher.Main(doc.Meta{
		ID: "statistic", Title: "Statistic 统计数值", Group: doc.GroupData,
		What: []string{
			"标题灰字 + **大号粗体数值**（SemiBold 600/24px）+ 前后缀小字",
			"千分位分隔（Separator 默认 \",\"）+ Precision 小数精度",
			"Countdown 倒计时（Statistic.Countdown 对应）：秒级自排程 tick+OnFinish/OnChange+Now 注入",
		},
		Demos: []doc.Demo{
			{Title: "金额", Note: "Precision=2 + 前缀 ¥：¥189,345.25", Build: func(ctx *ui.BuildContext) ui.Widget {
				return stAmount
			}},
			{Title: "用户数", Note: "千分位 + 后缀 人：112,893 人", Build: func(ctx *ui.BuildContext) ui.Widget {
				return stUsers
			}},
			{Title: "倒计时", Note: "Now 注入固定基准：距开售恒 02:30:00、大促 1 天 02:00:05（跨天确定性）；秒级 tick 自排程", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					row := scene.Flex(scene.Horizontal, 24)
					row.Add(cdNow.Build(c))
					row.Add(cdDay.Build(c))
					return row
				})
			}},
			{Title: "正计时", Note: "CountUp=true（antd Timer type=countup）：真实时钟运行时长", Build: func(ctx *ui.BuildContext) ui.Widget {
				return cdReal
			}},
		},
		API: []doc.APIRow{
			{Name: "NewStatistic", Type: "func(title string, value float64) *Statistic", Desc: "构造"},
			{Name: "Precision", Type: "int", Default: "0", Desc: "小数位数"},
			{Name: "Separator", Type: "string", Default: "\",\"", Desc: "千分位分隔符（空=不分隔）"},
			{Name: "Prefix / Suffix", Type: "string", Desc: "前缀/后缀小字（¥/元/人）"},
			{Name: "NewCountdown", Type: "func(title string, target time.Time) *Countdown", Desc: "倒计时构造（Statistic.Countdown 对应）"},
			{Name: "Countdown.Target / CountUp", Type: "time.Time / bool", Desc: "截止或起点时刻 / 正计时（Timer type=countup）"},
			{Name: "Countdown.Format", Type: "string", Default: "HH:mm:ss", Desc: "dayjs token 子集 D/HH/mm/ss/SSS"},
			{Name: "OnFinish / OnChange", Type: "func() / func(d)", Desc: "时间到 / 每 tick 回调"},
			{Name: "Countdown.Now", Type: "func() time.Time", Desc: "时刻注入点（审计跨天确定性）"},
			{Name: "Width", Type: "float32", Default: "140", Desc: "宽度"},
		},
	})
}
