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
	num     = widgets.NewInputNumber(0, 100, 20)
	numDis  = widgets.NewInputNumber(0, 100, 42)
	numPrec = widgets.NewInputNumber(0, 100, 12.3456)
	numStat = widgets.NewButton("值 = 20")
)

func init() {
	num.Width = 140
	numDis.Width = 140
	numDis.Disabled = true
	num.OnChange = func(v float32) {
		numStat.Text = fmt.Sprintf("值 = %v", v)
		numStat.MarkDirty()
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "input-number", Title: "InputNumber 数字输入框", Group: doc.GroupControl,
		What: []string{
			"只接受数字：右侧步进器（±Step）、聚焦后 ↑/↓ 键步进、min/max 钳制",
			"非法输入自动回退上一个合法值",
			"未实现：precision 小数位格式化、长按连发",
		},
		Demos: []doc.Demo{
			{Title: "小数位 precision", Note: "Precision=2：显示两位小数（12.3456→12.35）；失焦后格式化", Build: func(ctx *ui.BuildContext) ui.Widget {
				numPrec.Precision = 2
				return numPrec
			}},
			{Title: "0..100 步进 1", Note: "点右侧上下箭头或聚焦后按 ↑/↓；直接输入非数字会回退", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 10).SetCrossAlign(scene.CrossCenter)
				row.Add(num.Build(ctx))
				row.Add(numDis.Build(ctx))
				row.Add(numStat.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewInputNumber", Type: "func(min, max, value float32) *InputNumber", Desc: "构造"},
			{Name: "Min/Max", Type: "float32", Desc: "值域钳制"},
			{Name: "Step", Type: "float32", Default: "1", Desc: "步进幅度（按钮与键盘）"},
			{Name: "Value", Type: "float32", Desc: "当前数值"},
			{Name: "Width", Type: "float32", Default: "120", Desc: "框宽"},
			{Name: "Precision", Type: "int", Default: "-1", Desc: "小数位（>=0 固定位数四舍五入显示；-1 不限制）"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnChange", Type: "func(float32)", Desc: "数值变化回调"},
		},
	})
}
