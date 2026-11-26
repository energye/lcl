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
	editText = widgets.NewEditableText("点右侧铅笔编辑")
	ellText  = widgets.NewEllipsisText("这是一段很长的说明文字，用来演示多行省略与展开收起。Ant Design 的 Typography ellipsis 会在超过指定行数时截断并显示省略号，用户点击「展开」可查看全文，点击「收起」回到截断状态。这段文字故意写得比较长，以确保在较窄的容器宽度下一定会超过两行，从而触发省略效果。")
)

func init() {
	copyText.OnCopy = func(t string) {
		copyLog.Body = "已复制：" + t
	}
	editText.OnChange = func(t string) {
		copyLog.Body = "已编辑：" + t
	}
	ellText.Rows = 2
	ellText.MaxWidth = 480
	ellText.Expandable = true
	ellText.OnExpand = func(exp bool) {
		if exp {
			copyLog.Body = "已展开全文"
		} else {
			copyLog.Body = "已收起"
		}
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "typography", Title: "Typography 排版", Group: doc.GroupBasic,
		What: []string{
			"文本呈现三件套：Title（标题 1-5 级）/ Paragraph（段落）/ Text（内联语义文本）",
			"对齐 antd：标题 38/30/24/20/16 字重 600；正文 14 行高 1.57；Text 支持 secondary/success/warning/danger 语义色与 mark/code/strong/underline/delete/disabled",
			"CopyableText：右侧复制图标（OnCopy 注入宿主剪贴板）、EditableText 点击编辑（Enter 提交/Esc 取消）；EllipsisText 多行省略+展开收起",
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
			{Title: "编辑", Note: "EditableText：点铅笔切 Input，Enter 提交/Esc 取消；OnChange 回显", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Column{Gap: 8, Children: []ui.Widget{editText, copyLog}}
			}},
			{Title: "省略展开", Note: "EllipsisText：Rows=2 截断+省略号，点「展开/收起」切换；OnExpand 回显", Build: func(ctx *ui.BuildContext) ui.Widget {
				return widgets.Column{Gap: 8, Children: []ui.Widget{ellText, copyLog}}
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
			{Name: "NewEllipsisText", Type: "func(string) *EllipsisText", Desc: "多行省略：Rows 限行+Expandable 展开/收起（symbol 自定义）"},
			{Name: "NewEditableText/OnChange", Type: "func(string) *EditableText / func(string)", Desc: "可编辑文本；Enter/失焦提交，Esc 取消"},
			{Name: "CopyableText.CopyText", Type: "string", Desc: "复制内容（空 = Body）"},
			{Name: "Text.Disabled", Type: "bool", Desc: "禁用色"},
		},
	})
}
