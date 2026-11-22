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
	open    = widgets.NewButton("从左侧滑出抽屉")
	largeBn = widgets.NewButton("宽抽屉（736）")
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
}

func main() {
	launcher.Main(doc.Meta{
		ID: "drawer", Title: "Drawer 抽屉", Group: doc.GroupFeedback,
		What: []string{
			"边缘滑出的面板：左/右/上/下四向锚定视口，BandModal 层带",
			"关闭策略同 Modal（Esc / 点遮罩 / OnClose）",
			"PresetSize 预设宽度（378 默认 / 736 宽——antd size 语义），优先于 Width",
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
		},
		API: []doc.APIRow{
			{Name: "NewDrawer", Type: "func(title string, content ui.Widget) *Drawer", Desc: "构造"},
			{Name: "Edge", Type: "枚举", Default: "Left", Desc: "Left/Right/Top/Bottom"},
			{Name: "Width", Type: "float32", Default: "320", Desc: "左右向宽度（上下向为高度 400）"},
			{Name: "PresetSize", Type: "float32", Desc: "预设宽度（DrawerSizeDefault 378 / DrawerSizeLarge 736），优先于 Width"},
			{Name: "OnClose", Type: "func()", Desc: "关闭回调"},
			{Name: "Show/Close/IsOpen", Type: "同 Modal", Desc: "打开/关闭/状态"},
		},
	})
}
