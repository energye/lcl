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
	rate     = widgets.NewRate()
	rateStat = widgets.NewButton("评分：0")
)

func init() {
	rate.OnChange = func(v float32) {
		rateStat.Text = fmt.Sprintf("评分：%v", v)
		rateStat.MarkDirty()
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "rate", Title: "Rate 评分", Group: doc.GroupControl,
		What: []string{
			"星级评分：整星步进、悬停预览、再点当前值清除（antd 同语义）",
			"选中色对齐 antd（warning 黄）；未选为灰",
			"未实现：allowHalf 半星、tooltips 文案",
		},
		Demos: []doc.Demo{
			{Title: "五星评分", Note: "悬停预览、点击定格；真窗里移动鼠标看预览", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Vertical, 8)
				row.Add(rate.Build(ctx))
				row.Add(rateStat.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "三星只读", Note: "Disabled 后悬停与点击无效", Build: func(ctx *ui.BuildContext) ui.Widget {
				ro := widgets.NewRate()
				ro.Count = 3
				ro.Value = 2
				ro.Disabled = true
				return ro
			}},
		},
		API: []doc.APIRow{
			{Name: "NewRate", Type: "func() *Rate", Desc: "构造（默认 5 星）"},
			{Name: "Value", Type: "float32", Default: "0", Desc: "当前评分"},
			{Name: "Count", Type: "int", Default: "5", Desc: "星数"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "只读"},
			{Name: "OnChange", Type: "func(float32)", Desc: "评分回调"},
		},
	})
}
