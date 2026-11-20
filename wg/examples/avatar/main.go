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
		ID: "avatar", Title: "Avatar 头像", Group: doc.GroupDisplay,
		What: []string{
			"纯画布圆形头像：取文本首 rune 显示，无位图依赖",
			"默认形态=灰底 #F0F0F0 + 深色文字（antd 默认 Avatar）；显式 Color 时主色底+白字",
			"图标头像：Icon 组合复用（user 图标，尺寸 ≈ 直径×0.5）",
			"AvatarGroup 组合：Max 控制显示个数，超出合并为「+N」尾项",
		},
		Demos: []doc.Demo{
			{Title: "三档尺寸", Note: "28 / 36 / 48，文字取首字符", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 12)
				for _, s := range []struct {
					text string
					size float32
				}{{"张", 28}, {"李", 36}, {"WG", 48}} {
					row.Add((&widgets.Avatar{Text: s.text, Size: s.size}).Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
			{Title: "图标头像", Note: "Icon=\"user\"（Icon 组件组合复用）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					row := scene.Flex(scene.Horizontal, 12)
					row.Add((&widgets.Avatar{Icon: "user", Size: 36}).Build(c))
					row.Add((&widgets.Avatar{Icon: "user", Size: 36, Color: geom.RGBA(22, 119, 255)}).Build(c))
					return row
				})
			}},
			{Title: "头像组", Note: "5 人 Max=3 → 3 头像 + 「+2」尾项", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					grp := &widgets.AvatarGroup{Max: 3}
					for _, name := range []string{"张", "李", "王", "赵", "钱"} {
						grp.Avatars = append(grp.Avatars, &widgets.Avatar{Text: name, Size: 32})
					}
					return grp.Build(c)
				})
			}},
		},
		API: []doc.APIRow{
			{Name: "Text", Type: "string", Desc: "头像文字（显示首 rune）"},
			{Name: "Size", Type: "float32", Default: "36", Desc: "直径"},
			{Name: "Color", Type: "geom.Color", Default: "零值=灰底 #F0F0F0", Desc: "底色（显式时主色语义）"},
			{Name: "TextColor", Type: "geom.Color", Default: "随底色（灰底深字/主色底白字）", Desc: "文字色"},
			{Name: "Icon", Type: "string", Desc: "图标名（Icon 组件组合复用，优先于 Text）"},
			{Name: "AvatarGroup", Type: "struct{Avatars []*Avatar; Max int}", Desc: "头像组（Max 超出合并 +N 尾项）"},
		},
	})
}
