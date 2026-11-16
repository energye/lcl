package launcher

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/host"
)

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
//
// --shot 不是把 headless 缓冲偷换成 PNG：它跑的是同一个真窗 + 同一个
// Presenter，从合成缓冲导出的同时打印 Presented/Painted 计数——后者证明
// 像素真的经 GDI 提交到了窗口 DC（而不是只合成在内存里）。
func startShot(f *featureForm, path string) {
	go func() {
		// 静止界面按 C14 不会持续出帧，等到"至少 2 帧"或 2s 超时即可。
		deadline := time.Now().Add(2 * time.Second)
		for f.container.Frames() < 2 && time.Now().Before(deadline) {
			time.Sleep(5 * time.Millisecond)
		}
		// 同步补一次呈现：刚显示的窗口可能还没收到 WM_PAINT，
		// 只等消息会得到 painted=0 的假阴性。
		lcl.RunOnMainThreadAsync(func(id uint32) {
			f.container.PresentNow()
		})
		time.Sleep(120 * time.Millisecond)
		// 真窗呈现器必须真的提交过帧：这是"窗口不是空白"的直接证据。
		ps := f.container.PresentStats()
		fmt.Printf("[WG-SHOT] frames=%d painted=%d buffered=%v damage=%v\n",
			ps.Frames, ps.Painted, ps.Buffered, ps.LastDamage)
		if ps.Frames < 1 {
			fmt.Println("[WG-SHOT] FAIL: 没有帧提交到呈现器")
			requestExit(f, 1)
			return
		}
		// 屏幕级证据：--shot 读的是内存合成缓冲，内存有内容 ≠ 窗口有像素。
		// 这里采样真实窗口 DC，全白直接判失败（M6 空白 bug 的回归门禁）。
		nonWhiteWin, sampledWin, cw, ch, okWin := host.WindowPixelStats(
			uintptr(f.Handle()), 8)
		if !okWin {
			fmt.Println("[WG-SHOT] FAIL: 无法采样窗口客户区")
			requestExit(f, 1)
			return
		}
		fmt.Printf("[WG-WINDOW] client=%dx%d 窗口非白=%d/%d\n", cw, ch, nonWhiteWin, sampledWin)
		if nonWhiteWin < 50 {
			fmt.Println("[WG-SHOT] FAIL: 真实窗口客户区几乎全白（内存缓冲有内容也不算通过）")
			requestExit(f, 1)
			return
		}
		lcl.RunOnMainThreadAsync(func(id uint32) {
			img, err := backingToImage(f.container)
			if err != nil {
				fmt.Println("[WG-SHOT] error:", err)
				requestExit(f, 1)
				return
			}
			// 统计非背景像素，证明画布里真的画了东西。
			nonBg := countNonWhiteNRGBA(img, 4)
			if err := writePNG(img, path); err != nil {
				fmt.Println("[WG-SHOT] error:", err)
				requestExit(f, 1)
				return
			}
			fmt.Printf("[WG-SHOT] %dx%d -> %s  非背景采样点=%d\n",
				img.Bounds().Dx(), img.Bounds().Dy(), path, nonBg)
			requestExit(f, 0)
		})
	}()
}

// countNonWhiteNRGBA 按 step 步长采样非白像素数。
func countNonWhiteNRGBA(img *image.NRGBA, step int) int {
	nonBg := 0
	for y := 0; y < img.Bounds().Dy(); y += step {
		for x := 0; x < img.Bounds().Dx(); x += step {
			r, g, b, _ := img.NRGBAAt(x, y).RGBA()
			if r>>8 != 255 || g>>8 != 255 || b>>8 != 255 {
				nonBg++
			}
		}
	}
	return nonBg
}

func writePNG(img *image.NRGBA, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
