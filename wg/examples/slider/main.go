package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 展示实例包级持有：闭包内新建会让拖动值随重建丢失。
var (
	slider = widgets.NewSlider(0, 100, 35)
	marked = widgets.NewSlider(0, 100, 60)
	rg     = widgets.NewSliderRange(0, 100, 20, 70)
	st     = widgets.NewSlider(0, 100, 50)
	dis    = widgets.NewSlider(0, 100, 35)
)

func init() {
	slider.Width = 420
	marked.Width = 420
	marked.Marks = widgets.SliderMarks{0: "0℃", 50: "50℃", 100: "100℃"}
	rg.Width = 420
	st.Width = 420
	st.Step = 25
	dis.Width = 420
	dis.Disabled = true
}

func main() {
	launcher.Main(doc.Meta{
		ID: "slider", Title: "Slider 滑块", Group: doc.GroupControl,
		What: []string{
			"取值滑块：M8 对齐 antd——marks 刻度、range 双滑块、step 分档吸附、键盘步进",
			"命中区覆盖整个滑块带（不必精确落在 4px 轨道上）；拖动走 GestureArena",
			"未实现：tooltip 取值气泡、垂直方向",
		},
		Demos: []doc.Demo{
			{Title: "基本滑动", Note: "Min=0 Max=100 初始 35；键盘 ←/→ 步进", Build: func(ctx *ui.BuildContext) ui.Widget {
				return slider
			}},
			{Title: "marks 刻度", Note: "0/50/100 三个刻度线与文字", Build: func(ctx *ui.BuildContext) ui.Widget {
				return marked
			}},
			{Title: "range 双滑块", Note: "就近端拖动；[20, 70] 初始区间", Build: func(ctx *ui.BuildContext) ui.Widget {
				return rg
			}},
			{Title: "step 分档", Note: "Step=25 只取 0/25/50/75/100", Build: func(ctx *ui.BuildContext) ui.Widget {
				return st
			}},
			{Title: "禁用", Note: "antd disabled：填充消失、滑块描边转灰、无主色", Build: func(ctx *ui.BuildContext) ui.Widget {
				return dis
			}},
		},
		API: []doc.APIRow{
			{Name: "NewSlider", Type: "func(min, max, value float32) *Slider", Desc: "单滑块构造"},
			{Name: "NewSliderRange", Type: "func(min, max, from, to float32) *Slider", Desc: "双滑块构造（antd range）"},
			{Name: "Step", Type: "float32", Default: "0（连续）", Desc: "分档吸附"},
			{Name: "Marks", Type: "widgets.SliderMarks", Desc: "map[值]文字 的刻度表（antd marks）"},
			{Name: "Value/Value2", Type: "float32", Desc: "单滑块值 / range 第二端"},
			{Name: "Width", Type: "float32", Desc: "轨道宽度"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnChange/OnChange2", Type: "func(float32)", Desc: "值变化回调（range 两端）"},
		},
	})
}
