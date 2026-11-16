package wg

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/contracts"
	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/raster"
)

// CheckResult 是自检中的一项结果。
type CheckResult struct {
	// Group 是分组名（contracts / geom / raster / host / frame / container）。
	Group string
	// Name 是本项检查名。
	Name string
	// OK 表示通过。
	OK bool
	// Detail 是失败原因或补充信息。
	Detail string
	// Cost 是本项耗时。
	Cost time.Duration
}

// String 返回定宽单行。
func (r CheckResult) String() string {
	mark := "ok  "
	if !r.OK {
		mark = "FAIL"
	}
	line := fmt.Sprintf("[%s] %-10s %-34s %6.2fms", mark, r.Group, r.Name, float64(r.Cost.Microseconds())/1000)
	if r.Detail != "" {
		line += "  " + r.Detail
	}
	return line
}

// SelfTest 在无窗口环境下跑完整的 M0 自检，返回逐项结果。
//
// 覆盖：
//   - 契约表一致性（Validate）与当期断言（CheckAll，M0 只要求 C17）；
//   - 几何与矩阵基本正确性（旋转/求逆/裁剪/并集）；
//   - 帧缓冲的填充、混合、拷贝与字节视图；
//   - 画布树的成环拒绝、damage 上卷、detach；
//   - 手动时钟的确定性；
//   - headless 容器的空帧：帧序号、Delta、静止判定。
func SelfTest() []CheckResult {
	var out []CheckResult

	add := func(group, name string, fn func() error) {
		began := time.Now()
		err := fn()
		r := CheckResult{Group: group, Name: name, OK: err == nil, Cost: time.Since(began)}
		if err != nil {
			r.Detail = err.Error()
		}
		out = append(out, r)
	}

	// ── 契约 ─────────────────────────────────────────────────────────
	add("contracts", "Validate（编号/归属/断言齐全）", func() error {
		return contracts.Validate()
	})
	add("contracts", "CheckAll（当期里程碑全绿）", func() error {
		rs := contracts.CheckAll()
		s := contracts.Summarize(rs)
		if s.OK() {
			return nil
		}
		for _, r := range rs {
			if r.Status == contracts.Fail || r.Status == contracts.Missing {
				return fmt.Errorf("%s %s: %s", r.ID, r.Status, r.Detail)
			}
		}
		return fmt.Errorf("契约检查失败: %s", s)
	})

	// ── 几何 ─────────────────────────────────────────────────────────
	add("geom", "矩阵旋转与逆变换往返", func() error {
		m := geom.Translate(10, 20).Concat(geom.Rotate(0.5)).Concat(geom.Scale(2, 2))
		inv, ok := m.Invert()
		if !ok {
			return fmt.Errorf("矩阵不可逆")
		}
		p := geom.Pt(3, 7)
		back := inv.MapPoint(m.MapPoint(p))
		if d := back.Sub(p).Len(); d > 0.01 {
			return fmt.Errorf("往返误差 %.4f", d)
		}
		return nil
	})
	add("geom", "矩形交集/并集/包含", func() error {
		a := geom.RectXYWH(0, 0, 100, 100)
		b := geom.RectXYWH(50, 50, 100, 100)
		if got := a.Intersect(b); got != geom.RectLTRB(50, 50, 100, 100) {
			return fmt.Errorf("交集 = %v", got)
		}
		if got := a.Union(b); got != geom.RectLTRB(0, 0, 150, 150) {
			return fmt.Errorf("并集 = %v", got)
		}
		if a.Contains(geom.Pt(99.9, 99.9)) != true || a.Contains(geom.Pt(100, 100)) != false {
			return fmt.Errorf("半开区间判定错误")
		}
		return nil
	})
	add("geom", "颜色预乘打包（BGRA 小端）", func() error {
		c := geom.RGBAa(255, 128, 0, 128)
		v := c.PremulBGRA()
		if a := v >> 24; a != 128 {
			return fmt.Errorf("alpha = %d", a)
		}
		if r := (v >> 16) & 0xFF; r != 128 {
			return fmt.Errorf("预乘 R 应为 128，实际 %d", r)
		}
		return nil
	})

	// ── 帧缓冲 ───────────────────────────────────────────────────────
	add("raster", "backing 分配与不透明填充", func() error {
		b := raster.NewBacking(8, 4)
		if b == nil {
			return fmt.Errorf("分配失败")
		}
		b.Fill(geom.RGBA(255, 0, 0))
		// 打包值 0xFFFF0000 在小端内存里是 B=00 G=00 R=FF A=FF，即不透明红。
		if got := b.Pix()[0]; got != 0xFFFF0000 {
			return fmt.Errorf("像素 = %#08x，期望 0xffff0000", got)
		}
		if !b.Opaque() {
			return fmt.Errorf("不透明填充后 Opaque 应为 true")
		}
		if n := len(b.Bytes()); n != 8*4*4 {
			return fmt.Errorf("字节视图长度 = %d", n)
		}
		return nil
	})
	add("raster", "alpha 混合（source-over）", func() error {
		b := raster.NewBacking(2, 2)
		b.Fill(geom.RGBA(0, 0, 255))                  // 不透明蓝
		b.BlendRect(b.Bounds(), geom.RGBAa(255, 0, 0, 128)) // 半透明红
		v := b.Pix()[0]
		r, g, bl, a := (v>>16)&0xFF, (v>>8)&0xFF, v&0xFF, v>>24
		if a != 255 {
			return fmt.Errorf("alpha = %d", a)
		}
		if r < 120 || r > 135 || bl < 120 || bl > 135 || g != 0 {
			return fmt.Errorf("混合结果 R=%d G=%d B=%d", r, g, bl)
		}
		return nil
	})
	add("raster", "CopyRect 不透明拷贝（帧槽提交语义）", func() error {
		src := raster.NewBacking(4, 4)
		src.Fill(geom.RGBA(0, 255, 0))
		dst := raster.NewBacking(16, 16)
		dst.Fill(geom.RGBA(0, 0, 0))
		if !dst.CopyRect(src, src.Bounds(), 2, 3) {
			return fmt.Errorf("拷贝返回 false")
		}
		if got := dst.Row(3)[2]; got != 0xFF00FF00 {
			return fmt.Errorf("目标像素 = %#08x", got)
		}
		if got := dst.Row(3)[1]; got != 0xFF000000 {
			return fmt.Errorf("拷贝越界污染了左邻像素: %#08x", got)
		}
		return nil
	})

	// ── 画布树 ───────────────────────────────────────────────────────
	add("host", "Surface 成环拒绝", func() error {
		root := host.NewSurface(host.SurfaceInline, geom.RectXYWH(0, 0, 100, 100))
		child := host.NewSurface(host.SurfaceInline, geom.RectXYWH(0, 0, 10, 10))
		if err := root.Append(child); err != nil {
			return fmt.Errorf("正常挂接失败: %w", err)
		}
		if err := root.Append(root); err == nil {
			return fmt.Errorf("自己挂自己未被拒绝")
		}
		if err := child.Append(root); err == nil {
			return fmt.Errorf("祖先挂到后代未被拒绝")
		}
		if err := root.Append(child); err == nil {
			return fmt.Errorf("重复挂接未被拒绝")
		}
		return nil
	})
	add("host", "子 Surface damage 上卷到根", func() error {
		root := host.NewSurface(host.SurfaceInline, geom.RectXYWH(0, 0, 100, 100))
		child := host.NewSurface(host.SurfaceInline, geom.RectXYWH(10, 20, 30, 30))
		if err := root.Append(child); err != nil {
			return err
		}
		root.ClearDamageTree()
		child.Invalidate(geom.RectXYWH(0, 0, 5, 5))
		want := geom.RectLTRB(10, 20, 15, 25)
		if got := root.Damage(); got != want {
			return fmt.Errorf("根脏区 = %v，期望 %v", got, want)
		}
		// 越界部分被裁剪，不应扩大根脏区
		child.Invalidate(geom.RectXYWH(-50, -50, 10, 10))
		if got := root.Damage(); got != want {
			return fmt.Errorf("裁剪失效，根脏区 = %v", got)
		}
		return nil
	})
	add("host", "结构脏标记向上传播", func() error {
		root := host.NewSurface(host.SurfaceInline, geom.RectXYWH(0, 0, 100, 100))
		child := host.NewSurface(host.SurfaceLayer, geom.RectXYWH(0, 0, 10, 10))
		if err := root.Append(child); err != nil {
			return err
		}
		root.ClearStructureDirty()
		child.SetBounds(geom.RectXYWH(0, 0, 20, 20))
		if !root.StructureDirty() {
			return fmt.Errorf("子几何变化未传播到根")
		}
		return nil
	})

	// ── 时钟 ─────────────────────────────────────────────────────────
	add("frame", "ManualClock 确定性", func() error {
		a := frame.NewManualClock(0, 10*time.Millisecond)
		b := frame.NewManualClock(0, 10*time.Millisecond)
		for i := 0; i < 5; i++ {
			if x, y := a.Advance(), b.Advance(); x != y {
				return fmt.Errorf("第 %d 步不一致: %v vs %v", i, x, y)
			}
		}
		if got := a.Now(); got != 50*time.Millisecond {
			return fmt.Errorf("当前时刻 = %v", got)
		}
		if a.Ticks() != 5 {
			return fmt.Errorf("tick 计数 = %d", a.Ticks())
		}
		return nil
	})
	add("frame", "帧间隔中位数", func() error {
		d := []time.Duration{16, 17, 16, 300, 16}
		if got := frame.Median(d); got != 16 {
			return fmt.Errorf("中位数 = %v", got)
		}
		return nil
	})

	// ── 容器空帧 ─────────────────────────────────────────────────────
	add("container", "headless 空帧：静止判定与序号", func() error {
		c := NewHeadless(DefaultOptions())
		defer c.Close()
		clk := frame.NewManualClock(0, frame.DefaultStep)
		c.SetClock(clk)

		st0, err := c.FrameOnce()
		if err != nil {
			return err
		}
		if st0.Index != 0 || st0.Delta != 0 {
			return fmt.Errorf("首帧 Index=%d Delta=%v", st0.Index, st0.Delta)
		}
		if st0.Dirty {
			return fmt.Errorf("空容器的首帧不应为脏（damage=%v）", st0.Damage)
		}
		if st0.Surfaces != 1 {
			return fmt.Errorf("空容器应有 1 个根 Surface，实际 %d", st0.Surfaces)
		}
		clk.Advance()
		st1, err := c.FrameOnce()
		if err != nil {
			return err
		}
		if st1.Index != 1 {
			return fmt.Errorf("第二帧 Index = %d", st1.Index)
		}
		if st1.Delta != frame.DefaultStep {
			return fmt.Errorf("第二帧 Delta = %v，期望 %v", st1.Delta, frame.DefaultStep)
		}
		if !st1.Idle() {
			return fmt.Errorf("静止帧应满足 Idle()")
		}
		return nil
	})
	add("container", "resize 产生脏区且尺寸同步", func() error {
		c := NewHeadless(DefaultOptions())
		defer c.Close()
		c.Resize(120, 80)
		root := c.Root()
		if root == nil {
			return fmt.Errorf("Root 为 nil")
		}
		if got := root.Bounds().Size(); got != geom.Sz(120, 80) {
			return fmt.Errorf("根尺寸 = %v", got)
		}
		st, err := c.FrameOnce()
		if err != nil {
			return err
		}
		if !st.Dirty {
			return fmt.Errorf("resize 后应产生脏区")
		}
		if st.Damage.IsEmpty() {
			return fmt.Errorf("damage 为空")
		}
		return nil
	})
	add("container", "关闭后拒绝出帧", func() error {
		c := NewHeadless(DefaultOptions())
		c.Close()
		if !c.Closed() {
			return fmt.Errorf("Closed() = false")
		}
		if c.Root() != nil {
			return fmt.Errorf("关闭后 Root 应为 nil")
		}
		if _, err := c.FrameOnce(); err == nil {
			return fmt.Errorf("关闭后 FrameOnce 应返回错误")
		}
		c.Close() // 幂等
		return nil
	})

	return out
}

// SelfTestOK 表示全部自检通过。
func SelfTestOK(rs []CheckResult) bool {
	for _, r := range rs {
		if !r.OK {
			return false
		}
	}
	return true
}

// FirstFailure 返回第一条失败项，便于命令行直接报错退出。
func FirstFailure(rs []CheckResult) (CheckResult, bool) {
	for _, r := range rs {
		if !r.OK {
			return r, true
		}
	}
	return CheckResult{}, false
}
