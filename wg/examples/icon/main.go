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
		ID: "icon", Title: "Icon 图标", Group: doc.GroupBasic,
		What: []string{
			"语义图标：内置 84 个常用线性图标（对齐 @ant-design/icons 的 outline 风格命名——方向/数据/操作/编辑/状态/其他六类）",
			"实现是 wg 自绘的 render.Path（24×24 视口、2px 圆帽描边）——**近似**而非官方 SVG 逐点复刻；需要精确图标时用自定义路径扩展",
			"Color 零值取主题文字色；放按钮里传 Foreground 即实现 currentColor 继承",
		},
		Demos: []doc.Demo{
			{Title: "全部内置图标", Note: "线性 2px 圆帽描边，语义色", Build: func(ctx *ui.BuildContext) ui.Widget {
				wrap := scene.Wrap(scene.Horizontal, 10, 10).SetWidth(620)
				for _, name := range widgets.IconNames() {
					wrap.Add((&widgets.Icon{Name: name, Size: 20}).Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return wrap })
			}},
			{Title: "语义色与旋转", Note: "Success/Warning/Danger 着色；loading 旋转", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				row := scene.Flex(scene.Horizontal, 16).SetCrossAlign(scene.CrossCenter)
				ok := widgets.Icon{Name: "check-circle", Size: 20, Color: th.Color.Success}
				warn := widgets.Icon{Name: "warning-circle", Size: 20, Color: th.Color.Warning}
				err := widgets.Icon{Name: "close-circle", Size: 20, Color: th.Color.Danger}
				load := widgets.Icon{Name: "loading", Size: 20, Color: th.Color.Primary, Spin: true}
				for _, ic := range []*widgets.Icon{&ok, &warn, &err, &load} {
					row.Add(ic.Build(ctx))
				}
				row.Add(doc.Note(ctx, "loading 以 800ms 周期旋转（真窗可见）"))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return row })
			}},
		},
		API: []doc.APIRow{
			{Name: "Name", Type: "string", Desc: "图标名 84 个：方向（up/down/left/right + circle 系 + double/arrow）/数据（database/folder/file/image/mail/link/printer）/操作（lock/key/reload/undo/zoom/fullscreen/drag）/编辑（bold/italic/underline/align）/状态（play/pause/sound/wifi/cloud/thunderbolt/smile）/其他（shopping-cart/car/phone/bell/book/global/team 等）"},
			{Name: "Size", Type: "float32", Default: "16", Desc: "边长（正方形视口）"},
			{Name: "Color", Type: "geom.Color", Default: "Color.Text", Desc: "着色；零值取主题"},
			{Name: "Spin", Type: "bool", Default: "false", Desc: "旋转动画（用于 loading）"},
			{Name: "IconNames", Type: "func() []string", Desc: "内置图标名清单（排序稳定）"},
		},
	})
}
