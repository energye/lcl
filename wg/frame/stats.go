package frame

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/geom"
)

// Stats 描述一帧的度量结果，供 wg-diag 与性能门禁使用。
//
// M0 阶段只填充结构信息（Surface/Node 计数、damage、耗时）；
// 光栅与合成分段耗时（PREPARE/RECORD/RASTER/PRESENT）在 M1 接入。
type Stats struct {
	// Index 是帧序号，从 0 开始。
	Index int
	// Now 是帧开始时刻的时钟快照。
	Now time.Duration
	// Delta 是距上一帧的间隔（首帧为 0）。
	Delta time.Duration

	// Surfaces 是本帧遍历到的 Surface 数。
	Surfaces int
	// Nodes 是本帧遍历到的渲染节点数（M2 接入）。
	Nodes int
	// Layers 是本帧的合成层数（M2 接入）。
	Layers int

	// Damage 是本帧提交的脏区（DIP，已上卷并合并且裁到根）。
	Damage geom.Rect
	// Dirty 表示本帧是否有实际绘制内容。静止帧为 false，此时不提交。
	Dirty bool
	// Presented 表示本帧是否真的走到提交（有脏区且完成光栅）。
	Presented bool
	// RasterOps 是光栅后端处理的指令数。
	RasterOps int
	// Degraded 是本帧因后端能力缺失而降级的指令数（C22）。
	Degraded int
	// TextLines 是文字通道绘制的行数。
	TextLines int

	// FrameDuration 是本帧总耗时。
	FrameDuration time.Duration
}

// String 返回单行可读摘要。
func (s Stats) String() string {
	state := "idle"
	if s.Dirty {
		state = "dirty"
	}
	return fmt.Sprintf("frame=%d t=%s d=%s surfaces=%d nodes=%d layers=%d damage=[%.1f,%.1f %.1fx%.1f] %s %s",
		s.Index, s.Now, s.Delta, s.Surfaces, s.Nodes, s.Layers,
		s.Damage.Min.X, s.Damage.Min.Y, s.Damage.W(), s.Damage.H(),
		state, s.FrameDuration)
}

// Idle 表示这是一帧无内容变更的静止帧（契约 C14：静止零帧）。
func (s Stats) Idle() bool { return !s.Dirty }

// Median 返回帧间隔中位数（契约 C16：用中位数而非平均值，避免长尾污染）。
//
// 输入为空时返回 0。会就地排序传入切片。
func Median(deltas []time.Duration) time.Duration {
	if len(deltas) == 0 {
		return 0
	}
	// 手写插入排序：样本量小（几百），避免引入 sort 接口装箱。
	for i := 1; i < len(deltas); i++ {
		v := deltas[i]
		j := i - 1
		for j >= 0 && deltas[j] > v {
			deltas[j+1] = deltas[j]
			j--
		}
		deltas[j+1] = v
	}
	n := len(deltas)
	if n%2 == 1 {
		return deltas[n/2]
	}
	return (deltas[n/2-1] + deltas[n/2]) / 2
}

// Percentile 返回 p 分位（p ∈ [0,100]）帧间隔，用于定位卡顿。
func Percentile(deltas []time.Duration, p int) time.Duration {
	if len(deltas) == 0 {
		return 0
	}
	if p <= 0 {
		p = 0
	}
	if p >= 100 {
		p = 100
	}
	s := make([]time.Duration, len(deltas))
	copy(s, deltas)
	Median(s) // 复用排序
	idx := (len(s) - 1) * p / 100
	return s[idx]
}
