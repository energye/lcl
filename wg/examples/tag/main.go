package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "tag", Title: "Tag 标签", Group: doc.GroupDisplay,
		What: []string{
			"默认（无 Color）= 灰底 #FAFAFA + colorBorder 描边 + 深色文字（antd 默认 Tag）",
			"显式 Color = 浅底（基色 40% 透明）+ 同色文字 + 基色淡描边",
			"计划（M9）：closable 可关闭、checkable 可选中、preset 预设色名",
		},
		Demos: []doc.Demo{
			{Title: "标签组", Note: "宽度随文本自适应", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 8)
				for _, t := range []string{"Go", "LCL", "无句柄画布", "渲染引擎"} {
					row.Add((&widgets.Tag{Text: t}).Build(ctx))
				}
				row.Add((&widgets.Tag{Text: "可配置色", Color: geom.RGBA(22, 119, 255)}).Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
		},
		API: []doc.APIRow{
			{Name: "Text", Type: "string", Desc: "标签文字"},
			{Name: "Color", Type: "geom.Color", Default: "零值=默认灰底", Desc: "显式时按预设色渲染（浅底+同色文字）"},
		},
	})
}
