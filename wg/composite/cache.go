package composite

import (
	"container/list"

	"github.com/energye/lcl/wg/raster"
)

// RasterCache 缓存合成层的离屏栅格结果。
//
// 归属（契约 C7）：只有合成器持有一份，层树只写"我想要缓存"的意图，
// 不持有缓冲。后端（raster.Backend）**不**缓存，它每次只画到调用方给的
// 缓冲上——这条边界让"缓存有没有生效"永远是可控可测的。
//
// 键里没有位置，只有 (节点 id, 内容签名, 像素尺寸)：
// 层内容录在层局部空间，整层平移不改变内容，于是滚动/拖拽时缓存照样命中。
// 这条设计是 RasterCache 存在的意义所在——否则每次列表滚动都要重栅格化
// 整块内容，缓存只剩"静止时省一次"的边际收益。
type RasterCache struct {
	entries map[CacheKey]*list.Element
	lru     *list.List // 表头最近使用，表尾最久未用

	free   []*raster.Backing
	budget int64
	used   int64
	frame  int64

	hits      int
	misses    int
	evictions int
	leases    int

	// maxAge 是"多少帧没用过就回收"，0 表示只用字节预算淘汰。
	maxAge int64
}

// CacheKey 是离屏缓存的键。
type CacheKey struct {
	// ID 是来源节点 id（同一节点不同内容靠 Version 区分）。
	ID uint64
	// Version 是合并后的子树内容签名。
	Version uint64
	// W/H 是离屏像素尺寸：尺寸一变，图就不能用了。
	W, H int32
}

// cacheEntry 是一条缓存记录。
type cacheEntry struct {
	key     CacheKey
	backing *raster.Backing
	mask    []uint8 // 圆角/路径裁剪蒙版（nil 表示不需要）
	bytes   int64
	used    int64 // 最后一次命中的帧号（淘汰策略用）
}

// CacheStats 是缓存统计。
type CacheStats struct {
	Entries   int
	Bytes     int64
	Hits      int
	Misses    int
	Evictions int
	Leases    int
}

// hitRatio 返回命中率（无访问时返回 0）。
func (s CacheStats) hitRatio() float64 {
	total := s.Hits + s.Misses
	if total == 0 {
		return 0
	}
	return float64(s.Hits) / float64(total)
}

// NewRasterCache 创建缓存；budget <= 0 时用默认 8MiB。
func NewRasterCache(budget int64) *RasterCache {
	if budget <= 0 {
		budget = defaultCacheBytes
	}
	return &RasterCache{
		entries: make(map[CacheKey]*list.Element, 16),
		lru:     list.New(),
		budget:  budget,
		maxAge:  defaultCacheMaxAge,
	}
}

const (
	// defaultCacheBytes 是默认预算：8MiB ≈ 1080p 全屏一块。
	//
	// 为什么这么小：缓存是"可选的加速"，内存紧张时它必须最先退让。
	// 一块 4K 全屏（8.3M 像素）就 33MiB，设太大等于默认吃掉几十兆。
	defaultCacheBytes int64 = 8 << 20

	// defaultCacheMaxAge 是回收阈值（帧）。动画结束后的两秒内自然回收，
	// 不做"立即清空"——那会在滚动停下的瞬间把刚用过的图扔掉。
	defaultCacheMaxAge int64 = 120

	// maxFreeBackings 是回收池容量上限。
	maxFreeBackings = 4
)

// SetBudget 调整字节预算，超出时立即淘汰。
func (c *RasterCache) SetBudget(bytes int64) {
	if c == nil || bytes <= 0 {
		return
	}
	c.budget = bytes
	c.trim(0)
}

// BeginFrame 推进帧号（淘汰策略与 maxAge 依赖它）。
func (c *RasterCache) BeginFrame() {
	if c == nil {
		return
	}
	c.frame++
}

// Lookup 查缓存；命中时把它移到 LRU 表头并返回记录。
func (c *RasterCache) Lookup(k CacheKey) *cacheEntry {
	if c == nil {
		return nil
	}
	el, ok := c.entries[k]
	if !ok {
		c.misses++
		return nil
	}
	c.lru.MoveToFront(el)
	c.hits++
	e := el.Value.(*cacheEntry)
	e.used = c.frame
	return e
}

// Store 放入一条记录（接管 backing 的所有权）。
func (c *RasterCache) Store(k CacheKey, b *raster.Backing, mask []uint8) {
	if c == nil || b == nil {
		return
	}
	bytes := int64(b.W()) * int64(b.H()) * 4
	if bytes > c.budget {
		// 单块就超预算：不缓存，放进回收池（调用方本帧用完即弃）。
		c.Release(b)
		return
	}
	if el, ok := c.entries[k]; ok {
		c.removeElement(el)
	}
	e := &cacheEntry{key: k, backing: b, mask: mask, bytes: bytes, used: c.frame}
	c.entries[k] = c.lru.PushFront(e)
	c.used += bytes
	c.trim(0)
}

// Lease 取一块 W×H 的离屏缓冲（已清零）。
//
// 优先复用回收池里的同尺寸块：滚动/动画时尺寸稳定，复用能省掉大量分配。
// 返回 nil 表示尺寸非法。
func (c *RasterCache) Lease(w, h int) *raster.Backing {
	if w <= 0 || h <= 0 {
		return nil
	}
	if c != nil {
		c.leases++
		for i := len(c.free) - 1; i >= 0; i-- {
			b := c.free[i]
			if b.W() == w && b.H() == h {
				c.free = append(c.free[:i], c.free[i+1:]...)
				b.Clear()
				return b
			}
		}
	}
	return raster.NewBacking(w, h)
}

// Release 把不再使用的缓冲放回回收池。
//
// 只放少数几块：回收池是"省一次分配"，不是第二份缓存。
func (c *RasterCache) Release(b *raster.Backing) {
	if c == nil || b == nil {
		return
	}
	if len(c.free) >= maxFreeBackings {
		return // 交给 GC
	}
	c.free = append(c.free, b)
}

// Trim 回收长时间未命中的记录。
func (c *RasterCache) Trim() {
	if c == nil {
		return
	}
	c.trim(c.maxAge)
}

// trim 淘汰：先按 age（maxAge>0 时），再按字节预算。
func (c *RasterCache) trim(maxAge int64) {
	if c == nil {
		return
	}
	if maxAge > 0 {
		for el := c.lru.Back(); el != nil; {
			prev := el.Prev()
			if c.frame-el.Value.(*cacheEntry).used > maxAge {
				c.removeElement(el)
			}
			el = prev
		}
	}
	for c.used > c.budget {
		el := c.lru.Back()
		if el == nil {
			break
		}
		c.removeElement(el)
	}
}

// removeElement 移除一条记录并把缓冲（可能）回收进池。
func (c *RasterCache) removeElement(el *list.Element) {
	e := el.Value.(*cacheEntry)
	c.lru.Remove(el)
	delete(c.entries, e.key)
	c.used -= e.bytes
	c.evictions++
	c.Release(e.backing)
}

// Clear 清空全部记录（尺寸变化/换主题时调用）。
func (c *RasterCache) Clear() {
	if c == nil {
		return
	}
	for el := c.lru.Front(); el != nil; el = el.Next() {
		e := el.Value.(*cacheEntry)
		c.Release(e.backing)
	}
	c.entries = make(map[CacheKey]*list.Element, 16)
	c.lru.Init()
	c.used = 0
}

// Stats 返回统计快照。
func (c *RasterCache) Stats() CacheStats {
	if c == nil {
		return CacheStats{}
	}
	return CacheStats{
		Entries:   len(c.entries),
		Bytes:     c.used,
		Hits:      c.hits,
		Misses:    c.misses,
		Evictions: c.evictions,
		Leases:    c.leases,
	}
}
