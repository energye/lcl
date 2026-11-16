// wg-examples 是 wg 引擎的**真窗**示例入口。
//
// 与 wg-formdemo / wg-showcase / wg-imedemo / wg-dbmodeler 的 headless 口径
// 不同，这里创建真实的 lcl.TEngForm 并进入 lcl.Application 消息循环：
// 鼠标、键盘、输入法、DPI、重绘都由平台消息驱动，走的是完整真窗链路。
//
//	go run ./wg/cmd/wg-examples
//
// 左侧点选示例，右侧即时切换；「下一项」可顺序走查全部示例。
package main

import (
	"fmt"
	"os"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/examples"
)

type TExamplesForm struct {
	lcl.TEngForm
	container *wg.Container
}

var mainForm TExamplesForm

func main() {
	lcl.Init()
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.NewForm(&mainForm)

	mainForm.SetCaption("wg 示例集 · 真窗（无句柄画布）")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(1360)
	mainForm.SetHeight(860)
	// 场景自己持有离屏缓冲，关掉 LCL 层的二次缓冲，避免整窗多余拷贝。
	mainForm.SetDoubleBuffered(false)

	mainForm.container = wg.NewContainer(&mainForm, wg.DefaultOptions())
	mainForm.container.MountWidget(examples.NewGallery())

	// Container.Run 内部已调用 lcl.Application.Run()，此处不要再调一次。
	if err := mainForm.container.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "wg 示例启动失败:", err)
		os.Exit(1)
	}
}
