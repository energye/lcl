package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var (
	cpBase   = widgets.NewColorPicker()
	cpAlpha  = widgets.NewColorPicker()
	cpEcho   = widgets.NewButton("当前：#1677FF")
	cpEchoA  = widgets.NewButton("当前：#1677FF80")
	cpInit   bool
	cpAInit  bool
)

func init() {
	cpBase.Presets = []string{
		"#F5222D", "#FA541C", "#FA8C16", "#FAAD14", "#FADB14",
		"#A0D911", "#52C41A", "#13C2C2", "#1677FF", "#2F54EB",
		"#722ED1", "#EB2F96", "#8C8C8C", "#595959", "#262626",
		"#FFFFFF", "#D9D9D9", "#8C8C8C", "#434343", "#000000",
	}
	cpBase.OnChange = func(hex string) {
		cpEcho.Text = "当前：" + hex
		cpEcho.MarkDirty()
	}
	// alpha 卡（antd showAlpha）：alpha 条+初值 #1677FF80（半透明）。
	cpAlpha.ShowAlpha = true
	cpAlpha.Presets = cpBase.Presets
	cpAlpha.OnChange = func(hex string) {
		cpEchoA.Text = "当前：" + hex
		cpEchoA.MarkDirty()
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "colorpicker", Title: "ColorPicker 颜色选择", Group: doc.GroupControl,
		What: []string{
			"触发框（圆角色块+HEX 文本）→ 点击弹取色面板（SV 双层渐变 render.NewLinearGradient+色相条+预设色板）",
			"HSV 状态机：SV 面板 x→s/y→v、色相条 x→h·360、点击取色 down 即执行；**面板保持打开**（antd 连续调色），外点/Esc/再点触发框关闭",
			"**showAlpha 透明度通道**：alpha 条（底渐变+当前色透明→不透明横渐变，点击 x→alpha）+Value 输出 #RRGGBBAA（alpha=1 时 6 位）+触发框色块半透明合成",
			"点击取色后触发框色块/HEX 随根重建刷新+OnChange",
		},
		Demos: []doc.Demo{
			{Title: "基础", Note: "默认 #1677FF；点 SV 面板/色相条/预设色块取色；预设首格 #F5222D；回显按钮看 OnChange 链路", Build: func(ctx *ui.BuildContext) ui.Widget {
				if !cpInit {
					cpInit = true
				}
				col := scene.Flex(scene.Vertical, 8)
				col.Add(cpBase.Build(ctx))
				col.Add(cpEcho.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "透明度", Note: "ShowAlpha：初值 #1677FF80（半透明蓝）；alpha 条点击调透明度，触发框色块白底格上叠色、HEX 显示 8 位", Build: func(ctx *ui.BuildContext) ui.Widget {
				if !cpAInit {
					cpAlpha.Value = "#1677FF80"
					cpAInit = true
				}
				col := scene.Flex(scene.Vertical, 8)
				col.Add(cpAlpha.Build(ctx))
				col.Add(cpEchoA.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewColorPicker", Type: "func() *ColorPicker", Desc: "构造（默认 #1677FF）"},
			{Name: "Value", Type: "string", Desc: "当前 hex（#RRGGBB；ShowAlpha 且 alpha<1 时 #RRGGBBAA）"},
			{Name: "ShowAlpha", Type: "bool", Default: "false", Desc: "透明度通道（antd showAlpha——alpha 条+8 位输出）"},
			{Name: "Defaults", Type: "string", Desc: "无值时初始色（antd defaultValue）"},
			{Name: "Presets", Type: "[]string", Desc: "预设色板（每行 10 格）"},
			{Name: "OnChange", Type: "func(hex string)", Desc: "取色回调（SV/色相/alpha/预设点击都触发）"},
			{Name: "Show / Close / IsOpen", Type: "func(...)", Desc: "程序化开关面板 / 状态查询"},
		},
	})
}
