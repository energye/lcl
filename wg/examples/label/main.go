package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// demoLabel 五档字号各一行。组件无状态，直接在 Build 里构造。
func demoLabel(ctx *ui.BuildContext) ui.Widget {
	th := ctx.ThemeOr()
	col := func() ui.Widget {
		rows := []struct {
			text string
			size float32
		}{
			{"标题 24px", th.Font.SizeTitle},
			{"大号 20px", th.Font.SizeXL},
			{"正文 14px · 中文排版走 rune 偏移", th.Font.SizeBase},
			{"小号 12px", th.Font.SizeSM},
			{"极小 11px", th.Font.SizeXS},
		}
		w := widgets.Column{Gap: 8}
		for _, s := range rows {
			w.Children = append(w.Children, &widgets.Label{Text: s.text, Size: s.size})
		}
		return &w
	}()
	return col
}

func main() {
	launcher.Main(doc.Meta{
		ID: "label", Title: "Label 文本", Group: doc.GroupBasic,
		What: []string{
			"显示一行或多行静态文本，是最基础的呈现组件",
			"排版按 rune 偏移计算（UTF-16 只在 host 层换算），中文/emoji 混排安全",
			"需要长文本截断/多行时用 MaxWidth/MaxLines/Ellipsis",
		},
		Demos: []doc.Demo{
			{Title: "字号阶梯", Note: "主题 Font token 的五档字号；颜色未设置时取 Text", Build: demoLabel},
		},
		API: []doc.APIRow{
			{Name: "Text", Type: "string", Desc: "显示的文本（支持 \\n 多行）"},
			{Name: "Size", Type: "float32", Default: "Font.SizeBase", Desc: "字号，<=0 时取主题正文字号"},
			{Name: "Color", Type: "geom.Color", Default: "Color.Text", Desc: "文字颜色，零值时取主题"},
			{Name: "MaxWidth", Type: "float32", Default: "0（不限）", Desc: "排版最大宽度，超出换行"},
			{Name: "MaxLines", Type: "int", Default: "0（不限）", Desc: "最大行数"},
			{Name: "Ellipsis", Type: "string", Default: "\"\"", Desc: "截断省略符，配合 MaxLines/MaxWidth 使用"},
			{Name: "Build", Type: "func(*ui.BuildContext) *scene.Node", Desc: "ui.Widget 接口实现"},
		},
	})
}
