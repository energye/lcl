package host

import (
	"sync"

	"github.com/energye/lcl/lcl"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
)

// ── 呈现链路（最后一公里） ──────────────────────────────────────────
//
// composite 把一帧合成为 BGRA32 预乘像素（raster.Backing），但"合成完"
// 不等于"画到屏幕"：Backing 只是一块 Go 内存。真正让像素出现在窗口上的
// 是 Presenter —— 它在帧泵线程记下最新帧，在主线程 WM_PAINT 里把像素
// 提交到窗口 DC（Windows 走 GDI StretchDIBits，单次不透明拷贝）。
//
// 为什么必须有它：Router 接管了全部消息，WM_PAINT 被吞掉后 LCL 不会画，
// 于是窗口永远是空白的（只看"进程没退出"会误判为真窗可用）。
//
// 线程模型：Frame 在帧泵 goroutine 调用，Paint 在窗口主线程调用，
// 二者只交换"缓冲指针 + 脏区"，不拷贝像素；读写由 mu 保护。合成器的 dst
// 在 Resize 前是同一块缓冲，因此 Paint 期间指针稳定（C21 视图约束）。
type Presenter struct {
	hwnd uintptr
	host *Host // 用于校验帧尺寸与当前客户区一致（同包，不构成外部循环）
	mu   sync.RWMutex

	// plat 是平台私有状态（Windows：LCL 24bpp 离屏位图）。
	plat presentPlatform

	// canvas 返回当前可绘制的窗口画布（LCL TCanvas）。
	// 只在 OnPaint 期间有效，所以由 Host 在 paint 回调里设置。
	canvas func() lcl.ICanvas
	// repaint 请求窗口重绘（LCL Repaint/Invalidate）。
	repaint func()

	b       *raster.Backing
	// own 是 Presenter 自己持有的稳定副本：合成器的 Backing 在下一帧会被
	// 覆写，若 Paint 直接读它，主线程与帧泵线程会竞争同一块像素（撕裂），
	// 且在某些驱动上 StretchDIBits 会因源缓冲正被写入而失败。
	own     *raster.Backing
	damage  geom.Rect
	frames  int
	painted int
}

// PresentStats 是呈现统计（诊断与单测用）。
type PresentStats struct {
	// Frames 是收到并记录的合成帧数。
	Frames int
	// Painted 是成功提交到窗口 DC 的次数。
	Painted int
	// LastDamage 是最近一帧提交的脏区。
	LastDamage geom.Rect
	// Buffered 表示当前是否缓存了一帧可提交的像素。
	Buffered bool
}

// NewPresenter 创建呈现器。host 为 nil 或 hwnd 为 0 时得到无窗口呈现器（headless）。
func NewPresenter(hwnd uintptr, h *Host) *Presenter {
	return &Presenter{hwnd: hwnd, host: h}
}

// SetCanvas 设置窗口画布提供者（LCL TCanvas，仅在 OnPaint 内有效）。
func (p *Presenter) SetCanvas(fn func() lcl.ICanvas) { p.canvas = fn }

// SetRepaint 设置重绘请求函数（LCL Repaint）。
func (p *Presenter) SetRepaint(fn func()) { p.repaint = fn }

// Frame 记录一帧合成结果，并让窗口无效化以便在 OnPaint 中提交。
// damage 为空表示整帧（headless 出图时常见）。
//
// 这里**拷贝**像素到自有缓冲：合成器的 Backing 是复用的，下一帧会原地覆写；
// 若 Paint 直接读它，主线程与帧泵线程竞争同一块像素会撕裂，实测还会让
// StretchDIBits 返回 0（窗口全白）。拷贝只发生在 Presented 帧，静止界面
// 一帧都不拷，符合 C14 的"按需装填"。
func (p *Presenter) Frame(b *raster.Backing, damage geom.Rect) {
	if b == nil {
		return
	}
	w, h := b.W(), b.H()
	if w <= 0 || h <= 0 {
		return
	}

	p.mu.Lock()
	if p.own == nil || p.own.W() != w || p.own.H() != h {
		p.own = raster.NewBacking(w, h)
	}
	// 只拷脏区（damage 为空则整帧）；逐行 memcpy 保持 stride 语义。
	src := b.Bytes()
	dst := p.own.Bytes()
	if len(src) != len(dst) {
		p.mu.Unlock()
		return
	}
	if damage.IsEmpty() {
		copy(dst, src)
	} else {
		r := damage.Intersect(b.Bounds())
		if r.IsEmpty() {
			p.mu.Unlock()
			return
		}
		x0, y0 := int(r.Min.X), int(r.Min.Y)
		x1, y1 := int(r.Max.X), int(r.Max.Y)
		for y := y0; y < y1; y++ {
			row := y * p.own.Stride()
			copy(dst[row+x0*4:row+x1*4], src[row+x0*4:row+x1*4])
		}
	}
	p.damage = damage
	p.frames++
	p.mu.Unlock()
	if p.hwnd != 0 {
		p.invalidate(damage)
	}
}

// Paint 把缓存帧提交到窗口 DC；没有缓存或提交失败返回 false。
// 必须在窗口主线程（WM_PAINT 处理中）调用。
func (p *Presenter) Paint() bool {
	p.mu.RLock()
	b := p.own
	d := p.damage
	p.mu.RUnlock()
	if b == nil || p.hwnd == 0 {
		return false
	}
	if !p.submit(b, d) {
		return false
	}
	p.mu.Lock()
	p.painted++
	p.mu.Unlock()
	return true
}

// Stats 返回呈现统计。
func (p *Presenter) Stats() PresentStats {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return PresentStats{
		Frames:     p.frames,
		Painted:    p.painted,
		LastDamage: p.damage,
		Buffered:   p.own != nil,
	}
}

// Close 释放平台资源。
func (p *Presenter) Close() {
	p.mu.Lock()
	p.b = nil
	p.own = nil
	p.mu.Unlock()
	p.release()
}
