// wg-examples 是 wg 引擎的**真窗**示例入口，也是真窗呈现的自检工具。
//
// 与 wg-formdemo / wg-showcase / wg-imedemo / wg-dbmodeler 的 headless 口径
// 不同，这里创建真实的 lcl.TEngForm 并进入 lcl.Application 消息循环：
// 鼠标、键盘、输入法、DPI、重绘都由平台消息驱动，走的是完整真窗链路。
//
//	go run ./wg/cmd/wg-examples                 # 交互浏览 65 条示例
//	go run ./wg/cmd/wg-examples --shot=out.png  # 出图并打印呈现统计后退出
//
// --shot 不是把 headless 缓冲偷换成 PNG：它跑的是同一个真窗 + 同一个
// Presenter，等 30 帧稳定后从合成缓冲导出，同时打印 Presented/Painted
// 计数——后者证明像素真的经 GDI 提交到了窗口 DC（而不是只合成在内存里）。
package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"
	"time"

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

var (
	shotPath   string
	selectID   string
	exitCode   = -1 // -1 = 正常窗口运行；>=0 = 自检结束后退出
	galleryRef *examples.Gallery
)

func requestExit(code int) {
	lcl.RunOnMainThreadAsync(func(id uint32) {
		exitCode = code
		lcl.Application.Terminate()
	})
}

func main() {
	lcl.Init()
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.NewForm(&mainForm)

	for _, a := range os.Args[1:] {
		if strings.HasPrefix(a, "--shot=") {
			shotPath = strings.TrimPrefix(a, "--shot=")
		}
		if strings.HasPrefix(a, "--select=") {
			selectID = strings.TrimPrefix(a, "--select=")
		}
	}

	mainForm.SetCaption("wg 示例集 · 真窗（无句柄画布）")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(1360)
	mainForm.SetHeight(860)
	// 场景自己持有离屏缓冲，关掉 LCL 层的二次缓冲，避免整窗多余拷贝。
	mainForm.SetDoubleBuffered(false)

	mainForm.container = wg.NewContainer(&mainForm, wg.DefaultOptions())
	galleryRef = examples.NewGallery()
	mainForm.container.MountWidget(galleryRef)
	if selectID != "" && !galleryRef.Select(selectID) {
		fmt.Fprintf(os.Stderr, "wg 示例启动失败: 未找到示例 %q\n", selectID)
		os.Exit(2)
	}

	if shotPath != "" {
		startShot(shotPath)
	}

	// Container.Run 内部已调用 lcl.Application.Run()，此处不要再调一次。
	if err := mainForm.container.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "wg 示例启动失败:", err)
		os.Exit(1)
	}
	if exitCode >= 0 {
		os.Exit(exitCode)
	}
}

// backingToImage 把合成器的 BGRA32 预乘缓冲读成 PNG。
// 这是 Container.Backing() 暴露的真实帧内容，不是另画一遍。
func backingToImage(c *wg.Container) (*image.NRGBA, error) {
	b := c.Backing()
	if b == nil {
		return nil, fmt.Errorf("合成缓冲不可用（未出帧）")
	}
	w, h := b.W(), b.H()
	if w < 1 || h < 1 {
		return nil, fmt.Errorf("合成缓冲尺寸非法 %dx%d", w, h)
	}
	stride := b.Stride()
	px := b.Bytes()
	out := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		src := y * stride
		dst := y * out.Stride
		for x := 0; x < w; x++ {
			// premul BGRA: 字节序 B,G,R,A
			s := src + x*4
			a := px[s+3]
			out.Pix[dst+x*4+0] = px[s+2] // R
			out.Pix[dst+x*4+1] = px[s+1] // G
			out.Pix[dst+x*4+2] = px[s+0] // B
			out.Pix[dst+x*4+3] = a
		}
	}
	return out, nil
}

// startShot 等帧泵稳定后导出 PNG，并打印呈现统计与"非背景像素"证据。
func startShot(path string) {
	go func() {
		deadline := time.Now().Add(20 * time.Second)
		for mainForm.container.Frames() < 30 && time.Now().Before(deadline) {
			time.Sleep(10 * time.Millisecond)
		}
		// 真窗呈现器必须真的提交过帧：这是"窗口不是空白"的直接证据。
		ps := mainForm.container.PresentStats()
		fmt.Printf("[WG-SHOT] frames=%d painted=%d buffered=%v damage=%v\n",
			ps.Frames, ps.Painted, ps.Buffered, ps.LastDamage)
		if ps.Frames < 1 {
			fmt.Println("[WG-SHOT] FAIL: 没有帧提交到呈现器")
			requestExit(1)
			return
		}
		lcl.RunOnMainThreadAsync(func(id uint32) {
			img, err := backingToImage(mainForm.container)
			if err != nil {
				fmt.Println("[WG-SHOT] error:", err)
				requestExit(1)
				return
			}
			// 统计非背景像素，证明画布里真的画了东西。
			nonBg := 0
			for y := 0; y < img.Bounds().Dy(); y += 4 {
				for x := 0; x < img.Bounds().Dx(); x += 4 {
					r, g, b, _ := img.NRGBAAt(x, y).RGBA()
					if r>>8 != 255 || g>>8 != 255 || b>>8 != 255 {
						nonBg++
					}
				}
			}
			f, err := os.Create(path)
			if err != nil {
				fmt.Println("[WG-SHOT] create:", err)
				requestExit(1)
				return
			}
			if err := png.Encode(f, img); err != nil {
				f.Close()
				fmt.Println("[WG-SHOT] encode:", err)
				requestExit(1)
				return
			}
			f.Close()
			fmt.Printf("[WG-SHOT] %dx%d -> %s  非背景采样点=%d\n",
				img.Bounds().Dx(), img.Bounds().Dy(), path, nonBg)
			requestExit(0)
		})
	}()
}
