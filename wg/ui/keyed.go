package ui

import (
	"github.com/energye/lcl/wg/scene"
)

// ── Key 与最小 diff ─────────────────────────────────────────────────
//
// 列表是"节点复用"唯一必须显式处理的场景：其他地方的子节点集合是稳定的，
// 只有列表会插入/删除/重排。没有 Key 的时候按下标复用，插入一项就会让
// 后面所有项的状态（输入内容、滚动位置、动画相位）错位一格。

// Diff 是一次 diff 的结果。
type Diff struct {
	Reused  int
	Created int
	Removed int
	Moved   int
}

// Empty 表示没有任何变化。
func (d Diff) Empty() bool { return d.Created == 0 && d.Removed == 0 && d.Moved == 0 }

// String 返回可读摘要（诊断用）。
func (d Diff) String() string {
	return "reuse=" + itoa(d.Reused) + " create=" + itoa(d.Created) +
		" remove=" + itoa(d.Removed) + " move=" + itoa(d.Moved)
}

// itoa 是极小整数转字符串（避免为诊断引入 strconv）。
func itoa(v int) string {
	if v == 0 {
		return "0"
	}
	neg := v < 0
	if neg {
		v = -v
	}
	var b [12]byte
	i := len(b)
	for v > 0 {
		i--
		b[i] = byte('0' + v%10)
		v /= 10
	}
	if neg {
		i--
		b[i] = '-'
	}
	return string(b[i:])
}

// DiffKeys 比较前后两次的 Key 序列。
//
// 只关心"哪些能复用、哪些新增、哪些删除、哪些换了位置"——具体的节点
// 挂接由调用方做（它才知道节点的父子关系）。
func DiffKeys(old, next []Key) Diff {
	var d Diff
	if len(old) == 0 {
		d.Created = len(next)
		return d
	}
	index := make(map[Key]int, len(old))
	for i, k := range old {
		// 重复 Key 只记第一次：重复本身是调用方的 bug，不该让 diff 失真。
		if _, ok := index[k]; !ok {
			index[k] = i
		}
	}
	seen := make(map[Key]bool, len(next))
	lastPos := -1
	for _, k := range next {
		if k == "" {
			d.Created++ // 无 Key 的项一律当新增（按位置复用由调用方决定）
			continue
		}
		pos, ok := index[k]
		if !ok || seen[k] {
			d.Created++
			continue
		}
		seen[k] = true
		d.Reused++
		// 位置相对前一个复用项倒退 → 算一次移动（用于判断是否需要重排）。
		if pos < lastPos {
			d.Moved++
		} else {
			lastPos = pos
		}
	}
	for k := range index {
		if !seen[k] {
			d.Removed++
		}
	}
	return d
}

// KeyedList 按 Key 复用构建节点列表。
//
// reuse 传上一轮的节点（可为 nil）；同 Key 的节点会被就地复用，
// 调用方可在 build 里检查 old != nil 决定"原地改属性"还是"新建"。
//
// 返回新列表与 diff。节点被删除时由调用方负责从父节点摘除（本函数
// 不持有父引用，猜错会摘到别的树上）。
func KeyedList[T any](
	items []T,
	key func(T) Key,
	old []*scene.Node,
	oldKeys []Key,
	build func(item T, idx int, reuse *scene.Node) *scene.Node,
) (nodes []*scene.Node, keys []Key, d Diff) {
	d = DiffKeys(oldKeys, keysOf(items, key))
	byKey := make(map[Key]*scene.Node, len(old))
	for i, k := range oldKeys {
		if i < len(old) && k != "" {
			if _, ok := byKey[k]; !ok {
				byKey[k] = old[i]
			}
		}
	}
	nodes = make([]*scene.Node, 0, len(items))
	keys = make([]Key, 0, len(items))
	for i, it := range items {
		k := key(it)
		var reuse *scene.Node
		if k != "" {
			reuse = byKey[k]
		} else if i < len(old) {
			// 无 Key：按位置复用（长度变化时尾部增删）。
			reuse = old[i]
		}
		n := build(it, i, reuse)
		if n == nil {
			continue
		}
		nodes = append(nodes, n)
		keys = append(keys, k)
	}
	return nodes, keys, d
}

// keysOf 提取 Key 序列。
func keysOf[T any](items []T, key func(T) Key) []Key {
	out := make([]Key, len(items))
	for i, it := range items {
		out[i] = key(it)
	}
	return out
}
