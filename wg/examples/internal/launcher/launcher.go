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
	"path/filepath"
	"strings"
	"time"

	"github.com/energye/lcl/wg/examples/internal/audit"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/examples/doc"
)

// 默认窗口尺寸（--width/--height 可覆盖）。
const (
	defaultW = 1080
	// defaultH 1100：4 卡 demo 页总高常超 760 视口的滚动上限，页尾卡
	// 底部元素在 ensureVisible 40 次滚动后仍不可达（pagination/simple
	// 轮反复踩）。1100 让内容高 ~1400px 的滚动上限覆盖页尾卡。
	// --height 可覆盖。
	defaultH = 1100
)

type options struct {
	shot      string
	selfcheck bool
	audit     bool
	auditMode string
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
	if opts.audit {
		os.Exit(auditWindow(meta, opts))
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
		case a == "--audit":
			o.audit = true
		case strings.HasPrefix(a, "--audit="):
			o.audit = true
			o.auditMode = strings.TrimPrefix(a, "--audit=")
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

// dumpMeta 打印功能元数据（枚举门禁收集 ID/分组用；specgen 用 demos 标题生成 L1 用例）。
func dumpMeta(meta doc.Meta) {
	titles := make([]string, len(meta.Demos))
	for i, d := range meta.Demos {
		titles[i] = d.Title
	}
	out, err := json.Marshal(map[string]any{
		"id": meta.ID, "title": meta.Title, "group": meta.Group,
		"demos": titles, "apiRows": len(meta.API),
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

// auditWindow 打开真窗后运行全自动审计（滚动截屏 + 逐元素点击/键入），
// 报告与全页截图写入 --audit 指定目录（默认 .shots/audit/<id>）。
// 退出码 = 审计失败条数（上限 255）。
func auditWindow(meta doc.Meta, o options) int {
	outDir := ".shots/audit/" + meta.ID
	lcl.Init()
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	var form featureForm
	lcl.Application.NewForm(&form)
	form.SetCaption(meta.Title + " · 审计")
	form.SetPosition(types.PoScreenCenter)
	form.SetWidth(int32(o.w))
	form.SetHeight(int32(o.h))
	form.SetDoubleBuffered(false)

	form.container = wg.NewContainer(&form, wg.DefaultOptions())
	form.container.MountWidget(doc.Page(meta))

	go func() {
		deadline := time.Now().Add(3 * time.Second)
		for form.container.PresentStats().Frames < 2 && time.Now().Before(deadline) {
			time.Sleep(10 * time.Millisecond)
		}
		time.Sleep(300 * time.Millisecond)

		// draft 模式：逐屏枚举生成用例初稿，不跑断言。
		if o.auditMode == "draft" {
			spec := audit.DraftSpec(form.container)
			data, merr := json.MarshalIndent(spec, "", "  ")
			if merr == nil {
				rf, ferr := os.Create(filepath.Join(outDir, "spec-draft.json"))
				if ferr == nil {
					rf.WriteString(string(data) + "\n")
					rf.Close()
				}
			}
			fmt.Printf("[audit] %s draft cases=%d → %s\n",
				meta.ID, len(spec.Sections[0].Cases), filepath.Join(outDir, "spec-draft.json"))
			requestExit(&form, 0)
			return
		}

		// 声明式用例：wg/examples/<id>/spec.json（规范见 docs §06）。
		var spec *audit.Spec
		if data, err := os.ReadFile(filepath.Join("wg", "examples", meta.ID, "spec.json")); err == nil {
			if s, perr := audit.ParseSpec(data); perr != nil {
				fmt.Fprintln(os.Stderr, "spec 解析失败:", perr)
			} else {
				spec = s
			}
		}
		rep := audit.Run(form.container, outDir, spec)
		data, _ := json.MarshalIndent(rep, "", "  ")
		rf, err := os.Create(filepath.Join(outDir, "report.json"))
		if err == nil {
			rf.WriteString(string(data))
			rf.Close()
		}
		fmt.Printf("[audit] %s pages=%d elements=%d failed=%d\n",
			meta.ID, rep.Pages, rep.Elements, rep.Failed)
		for _, fd := range rep.Findings {
			if !fd.OK {
				fmt.Printf("[audit]   FAIL %-8s %-8s %s\n", fd.Element, fd.Check, fd.Detail)
			}
		}
		// 门禁退出码：spec 模式取声明式用例失败数（启发式偶发不阻塞）。
		code := rep.Failed
		if rep.Spec != "" {
			code = rep.SpecFailed
		}
		requestExit(&form, minInt(code, 255))
	}()

	if err := form.container.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "审计窗启动失败:", err)
		return 1
	}
	return exitCode
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func requestExit(f *featureForm, code int) {
	lcl.RunOnMainThreadAsync(func(id uint32) {
		exitCode = code
		if f.container != nil {
			f.container.ReleasePresent()
		}
		lcl.Application.Terminate()
	})
}
