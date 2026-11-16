package diag

import (
	"errors"
	"fmt"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
)

// ── --idle ──────────────────────────────────────────────────────────
//
// 静止零帧（C14）的门禁：1000 个节点、跑满一轮交互、然后什么都不做，
// 墙钟推进 2 秒必须**一帧都不出**。
//
// 为什么单独立一条命令：这是"笔记本风扇转不转"的差别，也是最容易被
// 后续改动破坏的一条——只要有人在某处加了一个"每帧检查一下 X"，
// 空转就回来了，而所有单测仍然全绿（它们只跑有限帧）。
// 这里的断言是**在无外部刺激下连续推进时钟**，只有真正做到按需装填
// 才可能为 0。

// IdleResult 是 --idle 的汇总。
type IdleResult struct {
	Nodes int
	// SettleFrames 是首次进入静止所花的帧数。
	SettleFrames int
	// IdleTicks 是静止后推进的时钟片数（每片 16ms）。
	IdleTicks int
	// IdleFrames 是静止后又出的帧数（必须为 0）。
	IdleFrames int
	// TouchFrames 是"改一个节点"之后出的帧数（应为 1）。
	TouchFrames int
	// TouchDamage 是那次改动的脏区。
	TouchDamage geom.Rect
	// Elapsed 是总耗时。
	Elapsed time.Duration
	Detail  string
}

// OK 判定是否全绿。
func (r IdleResult) OK() bool {
	return r.IdleFrames == 0 && r.TouchFrames == 1 && r.Nodes >= 1000 && r.IdleTicks >= 100
}

// String 返回可读摘要。
func (r IdleResult) String() string {
	s := fmt.Sprintf("idle: %d 节点 / 静止用 %d 帧 / 静止后 %d 个时钟片出了 %d 帧 / 改动 1 个节点 %d 帧 脏区 %.0fx%.0f / 用时 %v",
		r.Nodes, r.SettleFrames, r.IdleTicks, r.IdleFrames, r.TouchFrames,
		r.TouchDamage.W(), r.TouchDamage.H(), r.Elapsed)
	if r.Detail != "" {
		s += "  ← " + r.Detail
	}
	return s
}

// IdleTest 跑静止零帧门禁。
func IdleTest() (IdleResult, error) {
	began := time.Now()
	var res IdleResult
	defer func() { res.Elapsed = time.Since(began) }()

	const w, h = 900, 600
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(w, h)

	// 32×32 = 1024 个方块，每格 12px。
	grid := m2Grid(32, 12)
	c.Mount(grid)
	res.Nodes = countTree(grid)

	clk := frame.NewManualClock(0, 16*time.Millisecond)
	tk := c.RunManual(clk)
	if tk == nil {
		return res, fmt.Errorf("手动帧泵未启动")
	}

	// ① 跑到静止。
	for i := 0; i < 32 && tk.Armed(); i++ {
		tk.Advance(16 * time.Millisecond)
		res.SettleFrames++
	}
	if tk.Armed() {
		res.Detail = "32 帧后仍未静止"
		return res, errors.New(res.Detail)
	}

	// ② 什么都不做，推进 2 秒墙钟（125 片）。
	before := c.Frames()
	for i := 0; i < 125; i++ {
		tk.Advance(16 * time.Millisecond)
		res.IdleTicks++
	}
	res.IdleFrames = c.Frames() - before
	if res.IdleFrames != 0 {
		res.Detail = fmt.Sprintf("静止后又出了 %d 帧", res.IdleFrames)
		return res, errors.New(res.Detail)
	}

	// ③ 改一个节点：只该出一帧，且脏区只是它自己。
	before = c.Frames()
	target := grid.Children()[100]
	target.SetOpacity(0.5)
	for i := 0; i < 4 && tk.Armed(); i++ {
		tk.Advance(16 * time.Millisecond)
	}
	res.TouchFrames = c.Frames() - before
	if res.TouchFrames != 1 {
		res.Detail = fmt.Sprintf("改动后出了 %d 帧，应为 1", res.TouchFrames)
		return res, errors.New(res.Detail)
	}
	res.TouchDamage = c.LastDamage()

	// ④ 再次静止。
	before = c.Frames()
	for i := 0; i < 125; i++ {
		tk.Advance(16 * time.Millisecond)
	}
	if extra := c.Frames() - before; extra != 0 {
		res.Detail = fmt.Sprintf("再次静止后又出了 %d 帧", extra)
		return res, errors.New(res.Detail)
	}
	if res.TouchDamage.W() > 12*4 || res.TouchDamage.H() > 12*4 {
		res.Detail = fmt.Sprintf("脏区 %.0fx%.0f 超过节点尺寸", res.TouchDamage.W(), res.TouchDamage.H())
		return res, errors.New(res.Detail)
	}
	return res, nil
}

// m2Grid 造 cols×cols 个 size 方块。
func m2Grid(cols int, size float32) *scene.Node {
	kids := make([]*scene.Node, 0, cols*cols)
	for i := 0; i < cols*cols; i++ {
		b := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(180, 190, 200))})
		b.SetSize2(size, size)
		kids = append(kids, b)
	}
	return scene.Grid(cols, 1, kids...).SetSize2(float32(cols)*(size+1), float32(cols)*(size+1))
}

// countTree 统计节点数。
func countTree(n *scene.Node) int {
	if n == nil {
		return 0
	}
	c := 1
	for _, ch := range n.Children() {
		c += countTree(ch)
	}
	return c
}
