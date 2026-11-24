package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var ctxRef *ui.BuildContext

// open 包级持有：闭包内新建会让 hover 态随重建丢失。
var (
	open        = widgets.NewButton("从左侧滑出抽屉")
	largeBn     = widgets.NewButton("宽抽屉（736）")
	footerBtn   = widgets.NewButton("自定义 footer")
	noFooterBtn = widgets.NewButton("无 footer")
	resetB      = widgets.NewButton("重置")
	helpB       = widgets.NewButton("帮助")
	submitB     = widgets.NewButton("提交")
)

func init() {
	open.OnClick = func() {
		if ctxRef == nil || ctxRef.Overlay == nil {
			return
		}
		widgets.NewDrawer("侧边抽屉", doc.Node(func(ctx *ui.BuildContext) *scene.Node {
			return (&widgets.Label{Text: "Drawer 从边缘滑入，锚定视口一侧"}).Build(ctx)
		})).Show(ctxRef)
	}
	// 预设尺寸卡（antd size="large" 对应——PresetSize 736 优先于 Width）。
	largeBn.OnClick = func() {
		if ctxRef == nil || ctxRef.Overlay == nil {
			return
		}
		d := widgets.NewDrawer("宽抽屉", doc.Node(func(ctx *ui.BuildContext) *scene.Node {
			return (&widgets.Label{Text: "PresetSize=736（antd size=large），覆盖 Width"}).Build(ctx)
		}))
		d.PresetSize = widgets.DrawerSizeLarge
		d.Show(ctxRef)
	}
	footerBtn.OnClick = func() {
		if ctxRef == nil || ctxRef.Overlay == nil {
			return
		}
		d := widgets.NewDrawer("自定义 footer", doc.Node(func(ctx *ui.BuildContext) *scene.Node {
			return (&widgets.Label{Text: "底部内容由 Footer 属性提供，提交按钮关闭抽屉"}).Build(ctx)
		}))
		submitB.OnClick = func() { d.Close() }
		d.Footer = doc.Node(func(ctx *ui.BuildContext) *scene.Node {
			row := scene.Flex(scene.Horizontal, 8)
			row.SetMainAlign(scene.MainEnd)
			row.Add(resetB.Build(ctx))
			row.Add(helpB.Build(ctx))
			row.Add(submitB.Build(ctx))
			return row
		})
		d.Show(ctxRef)
	}
	noFooterBtn.OnClick = func() {
		if ctxRef == nil || ctxRef.Overlay == nil {
			return
		}
		d := widgets.NewDrawer("无 footer", doc.Node(func(ctx *ui.BuildContext) *scene.Node {
			return (&widgets.Label{Text: "FooterNull=true：显式隐藏 Footer（antd footer={null}）"}).Build(ctx)
		}))
		d.Footer = widgets.NewButton("隐藏按钮")
		d.FooterNull = true
		d.Show(ctxRef)
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "drawer", Title: "Drawer 抽屉", Group: doc.GroupFeedback,
		What: []string{
			"边缘滑出的面板：左/右/上/下四向锚定视口，BandModal 层带",
			"关闭策略同 Modal（Esc / 点遮罩 / OnClose）",
			"PresetSize 预设宽度（378 默认 / 736 宽——antd size 语义），优先于 Width",
			"Footer 自定义底部内容；默认无 footer，FooterNull 显式隐藏（antd footer/null）",
		},
		Demos: []doc.Demo{
			{Title: "左侧抽屉", Note: "Edge=Left，宽 320", Build: func(ctx *ui.BuildContext) ui.Widget {
				doc.CaptureCtx(&ctxRef, ctx)
				return open
			}},
			{Title: "预设尺寸", Note: "PresetSize=736 宽抽屉（antd size=large）", Build: func(ctx *ui.BuildContext) ui.Widget {
				doc.CaptureCtx(&ctxRef, ctx)
				return largeBn
			}},
			{Title: "自定义 footer", Note: "Footer ui.Widget 自定义底部（重置/帮助/提交，提交关框——antd footer ReactNode）", Build: func(ctx *ui.BuildContext) ui.Widget {
				doc.CaptureCtx(&ctxRef, ctx)
				return footerBtn
			}},
			{Title: "无 footer", Note: "FooterNull=true 显式隐藏 Footer（antd footer={null}，Esc 关闭）", Build: func(ctx *ui.BuildContext) ui.Widget {
				doc.CaptureCtx(&ctxRef, ctx)
				return noFooterBtn
			}},
		},
		API: []doc.APIRow{
			{Name: "NewDrawer", Type: "func(title string, content ui.Widget) *Drawer", Desc: "构造"},
			{Name: "Edge", Type: "枚举", Default: "Left", Desc: "Left/Right/Top/Bottom"},
			{Name: "Width", Type: "float32", Default: "320", Desc: "左右向宽度（上下向为高度 400）"},
			{Name: "PresetSize", Type: "float32", Desc: "预设宽度（DrawerSizeDefault 378 / DrawerSizeLarge 736），优先于 Width"},
			{Name: "Footer", Type: "ui.Widget", Desc: "自定义底部内容（antd footer ReactNode）；默认无 footer"},
			{Name: "FooterNull", Type: "bool", Default: "false", Desc: "显式隐藏 footer（antd footer={null}，优先于 Footer）"},
			{Name: "OnClose", Type: "func()", Desc: "关闭回调"},
			{Name: "Show/Close/IsOpen", Type: "同 Modal", Desc: "打开/关闭/状态"},
		},
	})
}
