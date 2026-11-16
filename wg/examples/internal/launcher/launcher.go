// Package launcher 是 examples/<feature>/main.go 的共享运行器。
//
// M7 起每个功能一个独立 main 包（package main，单独 go run），窗口自举、
// 真窗截图自检（--shot）、headless 自检（--selfcheck）、元数据导出（--meta）
// 都在这里统一实现，各功能 main 只组装 doc.Meta 并调用 Main。
//
//	launcher.Main(meta)               # 交互浏览（真窗）
//	go run ./wg/examples/button --shot=out.png    # 真窗出图自检
//	go run ./wg/examples/button --selfcheck       # headless 像素/降级判定（供 go test 枚举）
//	go run ./wg/examples/button --meta            # 打印功能元数据 JSON
//
// 依赖说明（C17）：examples 规则已放行组合根（wg.Container）、host 与 lcl，
// 本包是 examples 层里唯一被允许触碰它们的封装；功能 main 不得绕过这里
// 直接创建窗口。
package launcher

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/examples/doc"
)

// 默认窗口尺寸（--width/--height 可覆盖）。
const (
	defaultW = 1080
	defaultH = 760
)

type options struct {
	shot      string
	selfcheck bool
	meta      bool
	w, h      int
}

// Main 是每个功能 main 的唯一入口：解析参数 → 分派到对应运行模式。
// 返回的退出码通过 os.Exit 结束进程（headless 模式）或在消息循环结束后
// 返回（真窗模式）。
func Main(meta doc.Meta) {
	opts := parseArgs()
	if opts.w <= 0 {
		opts.w = defaultW
	}
	if opts.h <= 0 {
		opts.h = defaultH
	}
	if opts.meta {
		dumpMeta(meta)
		return
	}
	if opts.selfcheck {
		os.Exit(selfcheck(meta, opts))
		return
	}
	os.Exit(realWindow(meta, opts))
}

func parseArgs() options {
	var o options
	o.w, o.h = defaultW, defaultH
	for _, a := range os.Args[1:] {
		switch {
		case strings.HasPrefix(a, "--shot="):
			o.shot = strings.TrimPrefix(a, "--shot=")
		case a == "--selfcheck":
			o.selfcheck = true
		case a == "--meta":
			o.meta = true
		case strings.HasPrefix(a, "--width="):
			fmt.Sscanf(a, "--width=%d", &o.w)
		case strings.HasPrefix(a, "--height="):
			fmt.Sscanf(a, "--height=%d", &o.h)
		}
	}
	return o
}

// dumpMeta 打印功能元数据（枚举门禁收集 ID/分组用）。
func dumpMeta(meta doc.Meta) {
	out, err := json.Marshal(map[string]any{
		"id": meta.ID, "title": meta.Title, "group": meta.Group,
		"demos": len(meta.Demos), "apiRows": len(meta.API),
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "meta 序列化失败:", err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}

// featureForm 是功能窗宿主。每个功能一个独立进程，天然一窗一种。
type featureForm struct {
	lcl.TEngForm
	container *wg.Container
}

// realWindow 创建真窗并进入消息循环。返回进程退出码。
// 退出协议与旧 cmd/wg-examples 一致：先 ReleasePresent 再 Terminate
// （Terminate 之后运行时开始销毁，再 Free 呈现位图会 0xc0000005）。
func realWindow(meta doc.Meta, o options) int {
	lcl.Init()
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	var form featureForm
	lcl.Application.NewForm(&form)
	form.SetCaption(meta.Title + " · wg 功能窗")
	form.SetPosition(types.PoScreenCenter)
	form.SetWidth(int32(o.w))
	form.SetHeight(int32(o.h))
	// 场景自己持有离屏缓冲，关掉 LCL 层的二次缓冲，避免整窗多余拷贝。
	form.SetDoubleBuffered(false)

	form.container = wg.NewContainer(&form, wg.DefaultOptions())
	form.container.MountWidget(doc.Page(meta))

	if o.shot != "" {
		startShot(&form, o.shot)
	}
	// Container.Run 内部已调用 lcl.Application.Run()，此处不要再调一次。
	if err := form.container.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "功能窗启动失败:", err)
		return 1
	}
	return exitCode
}

var exitCode = -1 // -1 = 正常窗口运行；>=0 = 自检结束后退出

func requestExit(f *featureForm, code int) {
	lcl.RunOnMainThreadAsync(func(id uint32) {
		exitCode = code
		if f.container != nil {
			f.container.ReleasePresent()
		}
		lcl.Application.Terminate()
	})
}
