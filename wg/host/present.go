package host

import (
	"sync"
	"sync/atomic"
	"time"

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

	// staged 是 24bpp BGR 的最新帧副本（紧凑 stride = w*3）：帧泵 goroutine
	// 把合成缓冲（BGRA32）转换进来，主线程 Paint 只做行 memcpy + 位图装载。
	//
	// 为什么转换不能在主线程：全窗逐像素转换 + LoadFromIntfImage 实测
	// 20–37ms（2200×1200），正好落在 WM_SIZE/模态尺寸循环里，窗口边框都
	// 跟不上鼠标。转换下放到帧线程后，主线程每次提交只剩几 ms。
	staged []byte
	sw, sh int

	damage  geom.Rect
	frames  int
	painted int

	// post 请求主线程尽快呈现（真窗：PostMessage(WM_APP+n)）。
	post func()
	// presenting 是闭环槽：一条呈现消息在队列里或正在主线程画时保持 true，
	// 画完才允许投下一帧——否则 post 队列永远不空，硬件输入（鼠标/键盘）
	// 在 GetMessage 的优先级里永远排在 posted 之后 = 最大化后点不动。
	presenting atomic.Bool
	// missed 表示"上一帧还在画时又来了新帧"：画完立刻补投一次。
	missed atomic.Bool
	// paintCostUs 是最近一次主线程提交的耗时（微秒）。超过一帧预算时，
	// 帧泵在投递下一帧前留出 inputBreath 的空队列窗口，让输入先被取走
	// （showcase §9：慢帧后必须给 GetMessage 留机会，否则输入饿死）。
	paintCostUs atomic.Int64
	// stretchResize 为 true 时过渡帧整窗缩放；默认 1:1 + 背景填边。
	stretchResize bool
	background    geom.Color
	// dirty 表示 staged 里有尚未提交的新像素；没有新帧的 WM_PAINT
	// （被遮挡后重新露出）必须整帧重放，否则新暴露区是旧/残像素。
	dirty bool
	// allowPaint 表示当前这次 WM_PAINT 是 PresentDone 主动 Repaint 触发的：
	// 它就是要提交新帧，不能被下面的"去重"跳过。
	allowPaint atomic.Bool
	// needsFull 在"有 WM_PAINT 被让位给 posted 呈现"时置位：
	// 下一次提交整帧重放（覆盖系统要求重绘的区域），而不是只画脏区。
	needsFull atomic.Bool
	// suppressPost 表示渲染本身已在主线程（P0 方案①）：present 不需要
	// 再投递消息，渲染方会同步 Repaint。
	suppressPost atomic.Bool
}

// PendingSubmit 表示有尚未提交的新帧（dirty）或呈现消息在途（presenting）。
func (p *Presenter) PendingSubmit() bool {
	p.mu.RLock()
	d := p.dirty
	p.mu.RUnlock()
	return d || p.presenting.Load()
}

// PaintOnMessage 处理 WM_PAINT。
//
//   - 专用呈现（PresentDone 主动 Repaint 触发）：按帧脏区/整帧提交；
//   - 有在途新帧：让位给 posted 呈现（避免同一帧两条路径各画一遍），
//     记 needsFull 让下一次提交整帧覆盖；
//   - 其余系统重绘：只补 GetUpdateRect 的无效区（原生 rcPaint 语义）。
func (p *Presenter) PaintOnMessage(update geom.Rect) {
	if p.allowPaint.Load() {
		p.Paint()
		return
	}
	if p.PendingSubmit() {
		p.needsFull.Store(true)
		return
	}
	if update.IsEmpty() {
		return
	}
	p.paintRegion(update)
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
	// LastPaintCost 是最近一次主线程提交的耗时（诊断）。
	LastPaintCost time.Duration
}

// NewPresenter 创建呈现器。host 为 nil 或 hwnd 为 0 时得到无窗口呈现器（headless）。
func NewPresenter(hwnd uintptr, h *Host) *Presenter {
	return &Presenter{hwnd: hwnd, host: h}
}

// SetCanvas 设置窗口画布提供者（LCL TCanvas，仅在 OnPaint 内有效）。
func (p *Presenter) SetCanvas(fn func() lcl.ICanvas) { p.canvas = fn }

// SetRepaint 设置重绘请求函数（LCL Repaint）。
func (p *Presenter) SetRepaint(fn func()) {
	p.mu.Lock()
	p.repaint = fn
	p.mu.Unlock()
}

// SetResizePresent 设置过渡帧呈现方式与背景色（容器启动时注入）。
func (p *Presenter) SetResizePresent(stretch bool, bg geom.Color) {
	p.mu.Lock()
	p.stretchResize = stretch
	p.background = bg
	p.mu.Unlock()
}

// SetPost 设置"请求主线程尽快呈现"的投递函数（真窗为 PostMessage(WM_APP+n)）。
//
// 为什么要 posted 消息：GetMessage 的优先级是 sent > posted > 硬件输入 >
// WM_PAINT。帧泵在 goroutine 里合成，若只靠 InvalidateRect 等 WM_PAINT，
// 拖动/连续鼠标消息会把它饿死（内容不跟窗）。posted 消息在模态尺寸循环里
// 也会被及时派发，showcase 同款做法。
func (p *Presenter) SetPost(fn func()) {
	p.mu.Lock()
	p.post = fn
	p.mu.Unlock()
}

// SetSuppressPost 控制 present 是否还投递 posted 消息。
//
// 渲染在主线程进行时置 true：渲染方在 FrameOnce 之后同步 Repaint，
// 不再经过消息队列（也不在帧内 sleep inputBreath）。
func (p *Presenter) SetSuppressPost(v bool) { p.suppressPost.Store(v) }

// Present 在主线程同步完成一次呈现（渲染在主线程时由渲染方调用）。
func (p *Presenter) Present() {
	p.allowPaint.Store(true)
	p.RepaintNow()
	p.allowPaint.Store(false)
}

// PresentDone 在主线程处理完一条呈现消息后调用：同步提交 + 释放闭环槽。
//
// 画完才允许下一帧投递（showcase 的 FrameDone 语义）：否则慢帧下 post
// 队列永不空，硬件输入（鼠标/键盘）在 GetMessage 的优先级里排在 posted
// 之后，永远轮不到 = "内容还在动，但窗口点不动"。
//
// 画期间又来了新帧（missed）时**不能在主线程里直接补投**：主线程还没
// 回到 GetMessage，队列一空就被自己填上，输入照样没机会。慢帧用
// AfterFunc 把补投推迟 inputBreath，让主线程先取走挂起输入。
func (p *Presenter) PresentDone() {
	p.allowPaint.Store(true)
	p.RepaintNow()
	p.allowPaint.Store(false)
	p.presenting.Store(false)
	if !p.missed.Swap(false) {
		return
	}
	if p.paintCostUs.Load() > int64(frameBudget) {
		time.AfterFunc(inputBreath, p.postMissed)
	} else {
		p.postMissed()
	}
}

// postMissed 补投一帧（timer goroutine 或主线程调用）。
func (p *Presenter) postMissed() {
	if !p.presenting.CompareAndSwap(false, true) {
		return
	}
	p.mu.RLock()
	post := p.post
	p.mu.RUnlock()
	if post == nil {
		p.presenting.Store(false)
		return
	}
	post()
}

// RepaintNow 在主线程同步提交（LCL Repaint = Invalidate+Update → 立即
// OnPaint → Paint，不等低优先级 WM_PAINT）。
func (p *Presenter) RepaintNow() {
	p.mu.RLock()
	fn := p.repaint
	p.mu.RUnlock()
	if fn != nil {
		fn()
	}
}

// inputBreath 是慢帧后留给输入的取消息窗口。
//
// GetMessage 的优先级 posted > 硬件输入：只要 post 队列里一直有我们的
// 呈现消息，鼠标/键盘就永远排不上。慢帧（提交超过一帧预算）之后在投递
// 前空出 1.5ms，主线程会在这段时间进入 GetMessage 取走挂起输入
// （showcase §9 的实证值；正常 60Hz 帧不付这笔钱）。
const inputBreath = 1500 * time.Microsecond

// frameBudget 是 60Hz 的单帧预算；超过即视为慢帧。
const frameBudget = 16 * time.Millisecond

// breathIfSlow 在上一帧超预算时睡 inputBreath（帧泵 goroutine 调用）。
func (p *Presenter) breathIfSlow() {
	if p.paintCostUs.Load() > int64(frameBudget) {
		time.Sleep(inputBreath)
	}
}

// Frame 记录一帧合成结果，并让窗口无效化以便在 OnPaint 中提交。
// damage 为空表示整帧（headless 出图时常见）。
//
// 这里**拷贝并转换**像素到自有缓冲：合成器的 Backing 是复用的，下一帧会
// 原地覆写；若 Paint 直接读它，主线程与帧泵线程竞争同一块像素会撕裂。
// 转换（BGRA→24bpp BGR）放在帧线程完成，主线程只搬运。
func (p *Presenter) Frame(b *raster.Backing, damage geom.Rect) {
	if b == nil {
		return
	}
	w, h := b.W(), b.H()
	if w <= 0 || h <= 0 {
		return
	}

	p.mu.Lock()
	if p.staged == nil || p.sw != w || p.sh != h {
		p.staged = make([]byte, w*3*h)
		p.sw, p.sh = w, h
	}
	// 只转换脏区（damage 为空则整帧）。
	x0, y0, x1, y1 := 0, 0, w, h
	if !damage.IsEmpty() {
		r := damage.Intersect(b.Bounds())
		if r.IsEmpty() {
			p.mu.Unlock()
			return
		}
		x0, y0 = int(r.Min.X), int(r.Min.Y)
		x1, y1 = int(r.Max.X), int(r.Max.Y)
	}
	src := b.Bytes()
	sstride := b.Stride()
	dstride := w * 3
	for y := y0; y < y1; y++ {
		srow := y * sstride
		drow := y * dstride
		for x := x0; x < x1; x++ {
			s := srow + x*4
			d := drow + x*3
			p.staged[d+0] = src[s+0] // B
			p.staged[d+1] = src[s+1] // G
			p.staged[d+2] = src[s+2] // R
		}
	}
	p.damage = damage
	p.frames++
	p.dirty = true
	p.mu.Unlock()
	if p.hwnd != 0 {
		p.invalidate(damage)
	}
}

// Paint 把缓存帧提交到窗口 DC；没有缓存或提交失败返回 false。
// 必须在窗口主线程（WM_PAINT 处理中）调用。
func (p *Presenter) Paint() bool {
	return p.paintRegion(geom.Rect{})
}

// paintRegion 提交缓存帧；region 非空时只更新该区域（系统无效区）。
func (p *Presenter) paintRegion(region geom.Rect) bool {
	began := time.Now()
	if p.hwnd == 0 {
		return false
	}
	p.mu.RLock()
	buffered := p.staged != nil
	p.mu.RUnlock()
	if !buffered || !p.submitRegion(region) {
		return false
	}
	p.mu.Lock()
	p.painted++
	p.mu.Unlock()
	p.paintCostUs.Store(int64(time.Since(began)))
	return true
}

// Stats 返回呈现统计。
func (p *Presenter) Stats() PresentStats {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return PresentStats{
		Frames:        p.frames,
		Painted:       p.painted,
		LastDamage:    p.damage,
		Buffered:      p.staged != nil,
		LastPaintCost: time.Duration(p.paintCostUs.Load()),
	}
}

// Close 释放平台资源。
//
// 同时断开呈现通道：关闭后队列里可能还有一条已投递的呈现消息，WndProc
// 拿到的会是空回调（不能再去碰正在销毁的 LCL 画布）。
func (p *Presenter) Close() {
	p.mu.Lock()
	p.staged = nil
	p.sw, p.sh = 0, 0
	p.dirty = false
	p.post = nil
	p.repaint = nil
	p.mu.Unlock()
	p.presenting.Store(false)
	p.missed.Store(false)
	p.release()
}
