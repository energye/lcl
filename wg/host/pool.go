package host

import (
	"sort"

	"github.com/energye/lcl/wg/raster"
)

// ── 后备存储池 ──────────────────────────────────────────────────────
//
// 为什么必须有池：拖拽窗口边缘时每秒几十次尺寸切换，每次 1440×880×4 ≈ 5MB。
// 不池化的话 GC 每秒要收几十 MB 大对象，STW 会直接表现为"拖窗口卡顿"
// （示例实测：池化后拖拽卡顿消失）。
//
// 契约 C21 的释放顺序在这里被强制：**视图 → 位图 → 池**。
// 顺序错了会有两类后果，且都是偶发：
//   - 视图未解就归池：池把缓冲交给下一帧，旧视图仍在往上画 → 撕裂/串帧；
//   - 位图未归池就删视图：视图持有的 DIB 已失效 → GDI 句柄泄漏。
// 因此 Release 会拒绝"还有活视图"的缓冲并计入 WrongOrder，
// 而不是默默接受——静默接受会让问题漂到最后才炸。

// PoolStats 是池统计（C21 断言与 --resizetest 用）。
type PoolStats struct {
	Acquires   int
	Releases   int
	Hits       int
	Misses     int
	Evictions  int
	WrongOrder int
	InUse      int
	Free       int
	Bytes      int64
}

// Reused 返回命中率（0–1）。
func (s PoolStats) Reused() float64 {
	if s.Hits+s.Misses == 0 {
		return 0
	}
	return float64(s.Hits) / float64(s.Hits+s.Misses)
}

// Pool 复用帧缓冲。
//
// 只做**同尺寸**复用：不同尺寸的缓冲借出去之后调用方拿到的是
// (w0,h0) 的子视图，还得为它维护 stride/offset —— 那点内存省下来的价值
// 远不如下一轮 7 次尺寸切换时"每次都是新缓冲"的确定性。
type Pool struct {
	maxFree  int
	maxBytes int64

	free  []*raster.Backing // 按面积升序，便于淘汰最大的
	views map[*raster.Backing]int
	bytes int64
	seed  bool // 是否已清零（新缓冲也需要清）

	stats PoolStats
}

// NewPool 创建池。maxFree 是空闲缓冲数上限，maxBytes 是空闲总字节预算。
func NewPool(maxFree int, maxBytes int64) *Pool {
	if maxFree <= 0 {
		maxFree = 4
	}
	if maxBytes <= 0 {
		maxBytes = 48 << 20 // 48MiB
	}
	return &Pool{maxFree: maxFree, maxBytes: maxBytes, views: make(map[*raster.Backing]int, 8)}
}

// Acquire 取一块 w×h 的缓冲（内容已清零）。
//
// 清零不是礼貌而是正确性：复用缓冲里的旧像素会当作新内容被合成进去，
// 表现为"窗口里闪过上一帧的界面"。
func (p *Pool) Acquire(w, h int) *raster.Backing {
	if w <= 0 || h <= 0 {
		return nil
	}
	p.stats.Acquires++
	for i, b := range p.free {
		if b.W() == w && b.H() == h {
			p.free = append(p.free[:i], p.free[i+1:]...)
			p.stats.Hits++
			b.Clear()
			p.use(b)
			return b
		}
	}
	p.stats.Misses++
	b := raster.NewBacking(w, h)
	if b == nil {
		return nil
	}
	p.use(b)
	return b
}

// use 记账：借出的缓冲带一个"视图"（调用方必须显式 ReleaseView）。
func (p *Pool) use(b *raster.Backing) {
	if p.views == nil {
		p.views = make(map[*raster.Backing]int, 8)
	}
	p.views[b] = p.views[b] + 1
	p.stats.InUse = len(p.views)
	p.stats.Bytes = p.idleBytes()
}

// AcquireView 在已借出的缓冲上再挂一个视图（多路引用同一块像素）。
func (p *Pool) AcquireView(b *raster.Backing) bool {
	if b == nil {
		return false
	}
	if _, ok := p.views[b]; !ok {
		// 不是池借出去的缓冲：允许登记，但归池时会被当外部缓冲拒绝。
		p.views[b] = 0
	}
	p.views[b]++
	return true
}

// ReleaseView 解一个视图。返回 false 表示该缓冲没有活视图。
func (p *Pool) ReleaseView(b *raster.Backing) bool {
	if b == nil {
		return false
	}
	n, ok := p.views[b]
	if !ok || n <= 0 {
		p.stats.WrongOrder++
		return false
	}
	if n == 1 {
		delete(p.views, b)
	} else {
		p.views[b] = n - 1
	}
	p.stats.InUse = len(p.views)
	return true
}

// Views 返回某缓冲的视图数。
func (p *Pool) Views(b *raster.Backing) int { return p.views[b] }

// Release 归还缓冲。仍有视图时拒绝并记 WrongOrder（C21）。
func (p *Pool) Release(b *raster.Backing) bool {
	if b == nil {
		return false
	}
	p.stats.Releases++
	if n := p.views[b]; n > 0 {
		p.stats.WrongOrder++
		return false
	}
	delete(p.views, b)
	p.stats.InUse = len(p.views)

	size := int64(b.W()) * int64(b.H()) * 4
	// 超过预算的缓冲不归池：它一来就把预算吃满，后面所有请求都 miss。
	if size > p.maxBytes/2 {
		p.stats.Evictions++
		p.stats.Bytes = p.idleBytes()
		return false
	}
	p.free = append(p.free, b)
	p.bytes += size
	p.trim()
	p.stats.Bytes = p.idleBytes()
	return true
}

// trim 按 (数量上限, 字节预算) 淘汰空闲缓冲（优先淘汰最大的）。
func (p *Pool) trim() {
	if len(p.free) > p.maxFree || p.bytes > p.maxBytes {
		sort.SliceStable(p.free, func(i, j int) bool {
			return p.free[i].W()*p.free[i].H() > p.free[j].W()*p.free[j].H()
		})
		for len(p.free) > p.maxFree || (p.bytes > p.maxBytes && len(p.free) > 0) {
			b := p.free[0]
			p.bytes -= int64(b.W()) * int64(b.H()) * 4
			p.free = p.free[1:]
			p.stats.Evictions++
		}
	}
}

// idleBytes 返回空闲缓冲总字节。
func (p *Pool) idleBytes() int64 {
	var n int64
	for _, b := range p.free {
		n += int64(b.W()) * int64(b.H()) * 4
	}
	return n
}

// Trim 清空所有空闲缓冲（失活/最小化时调用）。
func (p *Pool) Trim() {
	p.stats.Evictions += len(p.free)
	p.free = p.free[:0]
	p.bytes = 0
}

// Stats 返回统计快照。
func (p *Pool) Stats() PoolStats {
	s := p.stats
	s.InUse = len(p.views)
	s.Free = len(p.free)
	s.Bytes = p.idleBytes()
	return s
}

// Leaked 返回仍有视图未解的缓冲数（>0 表示泄漏）。
func (p *Pool) Leaked() int {
	n := 0
	for _, v := range p.views {
		if v > 0 {
			n++
		}
	}
	return n
}
