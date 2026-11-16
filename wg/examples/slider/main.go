package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var slider = widgets.NewSlider(0, 100, 35)

func init() { slider.Width = 420 }

func main() {
	launcher.Main(doc.Meta{
		ID: "slider", Title: "Slider 滑块", Group: doc.GroupControl,
		What: []string{
			"单滑块取值控件：拖动走 GestureArena，越过 slop 的位置决定新值，避免首帧跳动",
			"获得焦点后 ←/→ 键步进",
			"计划（M8）：对齐 antd——marks 刻度、range 双滑块、tooltip 取值",
		},
		Demos: []doc.Demo{
			{Title: "基本滑动", Note: "Min=0 Max=100 初始 35", Build: func(ctx *ui.BuildContext) ui.Widget {
				return slider
			}},
		},
		API: []doc.APIRow{
			{Name: "NewSlider", Type: "func(min, max, value float32) *Slider", Desc: "构造"},
			{Name: "Min/Max", Type: "float32", Desc: "值域"},
			{Name: "Value", Type: "float32", Desc: "当前值"},
			{Name: "Width", Type: "float32", Desc: "轨道宽度"},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "禁用态"},
			{Name: "OnChange", Type: "func(float32)", Desc: "值变化回调"},
		},
	})
}
