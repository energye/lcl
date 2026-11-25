package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var (
	copyText = widgets.NewCopyableText("点右侧图标复制这段文字")
	copyLog  = widgets.NewCopyableText("")
)

func init() {
	copyText.OnCopy = func(t string) {
		copyLog.Body = "已复制：" + t
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "typography", Title: "Typography 排版", Group: doc.GroupBasic,
		What: []string{
			"文本呈现三件套：Title（标题 1-5 级）/ Paragraph（段落）/ Text（内联语义文本）",
			"对齐 antd：标题 38/30/24/20/16 字重 600；正文 14 行高 1.57；Text 支持 secondary/success/warning/danger 语义色与 mark/code/strong/underline/delete/disabled",
			"CopyableText：右侧复制图标（OnCopy 注入宿主剪贴板）、editable/ellipsis 展开未实现",
		},
		Demos: []doc.Demo{
			{Title: "标题层级", Note: "antd h1-h5 的字号与字重", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Column{Gap: 6, Children: []ui.Widget{
					widgets.Title{Text: "h1. 一级标题", Level: 1},
					widgets.Title{Text: "h2. 二级标题", Level: 2},
					widgets.Title{Text: "h3. 三级标题", Level: 3},
					widgets.Title{Text: "h4. 四级标题", Level: 4},
					widgets.Title{Text: "h5. 五级标题", Level: 5},
				}}
			}},
			{Title: "Text 语义", Note: "语义色 + strong/mark/code/underline/delete/disabled", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Column{Gap: 8, Children: []ui.Widget{
					widgets.Text{Body: "默认文本"},
					widgets.Text{Body: "次要文本 secondary", Type: widgets.TypeSecondary},
					widgets.Text{Body: "成功 success", Type: widgets.TypeSuccess},
					widgets.Text{Body: "警告 warning", Type: widgets.TypeWarning},
					widgets.Text{Body: "危险 danger", Type: widgets.TypeDanger},
					widgets.Text{Body: "加粗 strong", Strong: true},
					widgets.Text{Body: "标记 mark", Mark: true},
					widgets.Text{Body: "CODE_CODE", Code: true},
					widgets.Text{Body: "下划线 underline", Underline: true},
					widgets.Text{Body: "删除线 delete", Delete: true},
					widgets.Text{Body: "禁用 disabled", Disabled: true},
				}}
			}},
			{Title: "复制", Note: "CopyableText：点右侧图标复制，图标短暂变勾；OnCopy 回显（宿主写剪贴板）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Column{Gap: 8, Children: []ui.Widget{copyText, copyLog}}
			}},
			{Title: "段落", Note: "Paragraph 按宽度换行", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Paragraph{MaxWidth: 560, Text: "段落组件按容器宽度换行，行高 1.57 对齐 antd 正文排印。中文按 rune 断行，英文按词断行；长段落适合承载说明性内容。"}
			}},
		},
		API: []doc.APIRow{
			{Name: "Title.Level", Type: "int 1..5", Desc: "标题级别；字号 38/30/24/20/16、字重 600"},
			{Name: "Paragraph.MaxWidth", Type: "float32", Desc: "换行宽度"},
			{Name: "Text.Type", Type: "widgets.TextType", Desc: "Secondary/Success/Warning/Danger"},
			{Name: "Text.Strong/Mark/Code", Type: "bool", Desc: "加粗 / 黄底标记 / 等宽灰底代码"},
			{Name: "Text.Underline/Delete", Type: "bool", Desc: "下划线 / 删除线（自绘横线）"},
			{Name: "NewCopyableText/OnCopy", Type: "func(string) *CopyableText / func(string)", Desc: "可复制文本；OnCopy 注入宿主剪贴板（wg 无剪贴板通道）"},
			{Name: "CopyableText.CopyText", Type: "string", Desc: "复制内容（空 = Body）"},
			{Name: "Text.Disabled", Type: "bool", Desc: "禁用色"},
		},
	})
}
