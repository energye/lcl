package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "tag", Title: "Tag 标签", Group: doc.GroupDisplay,
		What: []string{
			"轻量标签：主色底 40% 透明 + 同色文字，按文本 intrinsic 宽度自适应",
			"计划（M9）：closable 可关闭、checkable 可选中、preset 预设色",
		},
		Demos: []doc.Demo{
			{Title: "标签组", Note: "宽度随文本自适应", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 8)
				for _, t := range []string{"Go", "LCL", "无句柄画布", "渲染引擎"} {
					row.Add((&widgets.Tag{Text: t}).Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
		},
		API: []doc.APIRow{
			{Name: "Text", Type: "string", Desc: "标签文字"},
			{Name: "Color", Type: "geom.Color", Default: "Color.Primary", Desc: "底/文字基色"},
		},
	})
}
