package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var ly = &widgets.Layout{
	Header:  &widgets.Label{Text: "顶栏 · 应用标题"},
	Sider:   &widgets.Label{Text: "导航\n菜单"},
	Content: &widgets.Label{Text: "主内容区——Layout 为语义化布局包装（Sider/Header/Content/Footer 分区锚）"},
	Footer:  &widgets.Label{Text: "底栏 · 版权信息"},
	Width:   640, Height: 320,
}

func main() {
	launcher.Main(doc.Meta{
		ID: "layout", Title: "Layout 布局容器", Group: doc.GroupControl,
		What: []string{
			"语义化布局包装：**Flex Vertical [Header, Flex Horizontal [Sider, Content], Footer]**——各子区语义锚（layout-sider/header/content/footer）",
			"antd 默认尺寸：Header 64 / Sider 200 / Footer 48；Sider 灰底近似（dark token 未做，如实标注）",
			"未实现：Sider 折叠动画/collapsedWidth/breakpoint/trigger/theme=dark",
		},
		Demos: []doc.Demo{
			{Title: "经典布局", Note: "Header 64 + Sider 200 + Content + Footer 48（总 640×320）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return ly
			}},
		},
		API: []doc.APIRow{
			{Name: "NewLayout", Type: "func(content ui.Widget) *Layout", Desc: "构造（主内容）"},
			{Name: "Header / Sider / Content / Footer", Type: "ui.Widget", Desc: "四区内容（语义锚 layout-header 等）"},
			{Name: "SiderWidth", Type: "float32", Default: "200", Desc: "侧边栏宽"},
			{Name: "HeaderHeight / FooterHeight", Type: "float32", Default: "64 / 48", Desc: "顶/底栏高"},
			{Name: "Width / Height", Type: "float32", Desc: "总尺寸（Height<=0 自适应）"},
		},
	})
}
