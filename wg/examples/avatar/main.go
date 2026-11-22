package main

import (
	"image"
	"image/color"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// demoImage 程序生成 64×64 渐变位图（图片头像 Src 演示——引擎无文件
// IO，Src 由宿主注入；程序生成等价于宿主解码结果）。
func demoImage() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 64, 64))
	for y := 0; y < 64; y++ {
		for x := 0; x < 64; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(80 + x*2),
				G: uint8(100 + y),
				B: uint8(200 - x),
				A: 255,
			})
		}
	}
	return img
}

func main() {
	launcher.Main(doc.Meta{
		ID: "avatar", Title: "Avatar 头像", Group: doc.GroupDisplay,
		What: []string{
			"纯画布圆形头像：取文本首 rune 显示，无位图依赖",
			"默认形态=灰底 #F0F0F0 + 深色文字（antd 默认 Avatar）；显式 Color 时主色底+白字",
			"图标头像：Icon 组合复用（user 图标，尺寸 ≈ 直径×0.5）",
			"图片头像：Src image.Image 圆形裁剪（render.ImageFrom 双线性——宿主解码位图注入）",
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
			{Title: "图片头像", Note: "Src image.Image 圆形裁剪（程序生成渐变位图演示——宿主解码后注入同通道）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					row := scene.Flex(scene.Horizontal, 12)
					row.Add((&widgets.Avatar{Src: demoImage(), Size: 36}).Build(c))
					row.Add((&widgets.Avatar{Src: demoImage(), Size: 48}).Build(c))
					return row
				})
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
			{Name: "Src", Type: "image.Image", Desc: "图片头像（圆形裁剪位图，优先于 Icon/Text）"},
			{Name: "Color", Type: "geom.Color", Default: "零值=灰底 #F0F0F0", Desc: "底色（显式时主色语义）"},
			{Name: "TextColor", Type: "geom.Color", Default: "随底色（灰底深字/主色底白字）", Desc: "文字色"},
			{Name: "Icon", Type: "string", Desc: "图标名（Icon 组件组合复用，优先于 Text）"},
			{Name: "AvatarGroup", Type: "struct{Avatars []*Avatar; Max int}", Desc: "头像组（Max 超出合并 +N 尾项）"},
		},
	})
}
