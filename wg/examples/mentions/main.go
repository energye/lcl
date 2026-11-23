package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var mn = widgets.NewMentions([]string{"afc", "antd", "uikit", "bluetooth"})

func main() {
	launcher.Main(doc.Meta{
		ID: "mentions", Title: "Mentions 提及", Group: doc.GroupControl,
		What: []string{
			"**纯组合复用 Input**（持久实例+HitLabel 嵌入锚）+ 键入 @ 弹选项浮层（BandPopup，Select 同通道）",
			"选项 **down 即执行** → 插入 \"@label \" 到文本（Input 重建同步 editor）+ OnChange",
			"**键盘上下选**（Input.OnKey 外置钩子）：面板打开时 ↑↓ 移动高亮（InfoBg 底+主色文字原地迁移）、Enter 确认插入；Esc 关闭浮层",
			"Esc/外部点击关闭浮层",
		},
		Demos: []doc.Demo{
			{Title: "基础", Note: "键入 @ 弹出候选 → 点选 afc → 文本插入 \"@afc \"；↑↓ 移动高亮、Enter 确认", Build: func(ctx *ui.BuildContext) ui.Widget {
				return mn
			}},
		},
		API: []doc.APIRow{
			{Name: "NewMentions", Type: "func(options []string) *Mentions", Desc: "构造（候选列表）"},
			{Name: "Options", Type: "[]string", Desc: "@ 候选（antd options 的 value/label 合一）"},
			{Name: "Value", Type: "string", Desc: "文本内容"},
			{Name: "OnChange", Type: "func(text string)", Desc: "文本变化回调（键入与插入都触发）"},
			{Name: "Placeholder", Type: "string", Default: "输入 @ 唤起提及", Desc: "占位文本"},
			{Name: "Width", Type: "float32", Default: "260", Desc: "宽度"},
			{Name: "Show / Close / IsOpen", Type: "func(...)", Desc: "程序化开关选项面板 / 状态查询"},
		},
	})
}
