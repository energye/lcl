package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。autoplay 关闭保证审计确定性。
var cr = widgets.NewCarousel([]ui.Widget{
	&widgets.Label{Text: "第一页：产品发布", Color: geom.RGBA(22, 119, 255)},
	&widgets.Label{Text: "第二页：版本更新", Color: geom.RGBA(82, 196, 26)},
	&widgets.Label{Text: "第三页：活动预告", Color: geom.RGBA(250, 84, 28)},
})

func main() {
	launcher.Main(doc.Meta{
		ID: "carousel", Title: "Carousel 走马灯", Group: doc.GroupData,
		What: []string{
			"**Stack 叠层切换**（SetVisible 翻页——零成本切换）+ 底部圆点指示（当前点主色放大）",
			"**点击圆点切换**（down 即执行）；Autoplay 定时轮播（默认 0 关闭——审计确定性，演示可开）",
			"未实现：滑动过渡动画/arrows/vertical/fade/Swipe 触摸",
		},
		Demos: []doc.Demo{
			{Title: "三页走马灯", Note: "点击底部圆点切换页面（当前点主色放大）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return cr
			}},
		},
		API: []doc.APIRow{
			{Name: "NewCarousel", Type: "func(slides []ui.Widget) *Carousel", Desc: "构造（页 Widget 列表）"},
			{Name: "Slides", Type: "[]ui.Widget", Desc: "每页内容"},
			{Name: "Autoplay", Type: "int", Default: "0（关）", Desc: "自动轮播间隔毫秒"},
			{Name: "Width / Height", Type: "float32", Default: "480 / 180", Desc: "页尺寸"},
		},
	})
}
