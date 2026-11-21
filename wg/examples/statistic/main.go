package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
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
)

func main() {
	launcher.Main(doc.Meta{
		ID: "statistic", Title: "Statistic 统计数值", Group: doc.GroupData,
		What: []string{
			"标题灰字 + **大号粗体数值**（SemiBold 600/24px）+ 前后缀小字",
			"千分位分隔（Separator 默认 \",\"）+ Precision 小数精度",
			"纯展示排版组件（Countdown 倒计时形态未实现）",
		},
		Demos: []doc.Demo{
			{Title: "金额", Note: "Precision=2 + 前缀 ¥：¥189,345.25", Build: func(ctx *ui.BuildContext) ui.Widget {
				return stAmount
			}},
			{Title: "用户数", Note: "千分位 + 后缀 人：112,893 人", Build: func(ctx *ui.BuildContext) ui.Widget {
				return stUsers
			}},
		},
		API: []doc.APIRow{
			{Name: "NewStatistic", Type: "func(title string, value float64) *Statistic", Desc: "构造"},
			{Name: "Precision", Type: "int", Default: "0", Desc: "小数位数"},
			{Name: "Separator", Type: "string", Default: "\",\"", Desc: "千分位分隔符（空=不分隔）"},
			{Name: "Prefix / Suffix", Type: "string", Desc: "前缀/后缀小字（¥/元/人）"},
			{Name: "Width", Type: "float32", Default: "140", Desc: "宽度"},
		},
	})
}
