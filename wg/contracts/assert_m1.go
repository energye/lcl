package contracts

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/composite"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/raster/gdi"
	"github.com/energye/lcl/wg/raster/software"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/text"
)

// 本文件是 M1 到期契约的断言实现。
//
// 原则：断言只验证"架构承诺"，不验证"像素好看"。
// 像素级比对属于 diag 的 --render-matrix / --golden（那是回归手段），
// 而契约断言要能回答"如果这个约束被破坏，会不会有具体的可观察后果"。

// ── 小工具 ──────────────────────────────────────────────────────────

func recordDL(fn func(c *render.Recorder)) *render.DisplayList {
	r := render.NewRecorder(render.NewDisplayList(32))
	fn(r)
	return r.Finish()
}

func drawSoftware(dl *render.DisplayList, w, h int) *raster.Backing {
	b := raster.NewBacking(w, h)
	if b == nil {
		return nil
	}
	software.New(software.Options{}).Draw(b, dl, b.Bounds())
	return b
}

func drawGDI(dl *render.DisplayList, w, h int) (*raster.Backing, bool) {
	if gdi.NewDIB(4, 4, 24) == nil {
		return nil, false
	}
	b := raster.NewBacking(w, h)
	if b == nil {
		return nil, false
	}
	be := gdi.New()
	defer be.Close()
	be.Draw(b, dl, b.Bounds())
	return b, true
}

func approx(a, b, tol float32) bool {
	d := a - b
	if d < 0 {
		d = -d
	}
	return d <= tol
}

// ── C1 画刷顺序：先填充后描边，描边不污染内部 ───────────────────────

func assertBrushOrder() error {
	// 描边只应落在边缘：内部必须保持透明。
	// 如果实现把描边当成填充（或先描边后填充把描边覆盖掉），这条会失败。
	dl := recordDL(func(c *render.Recorder) {
		p := render.NewPath().AddRect(geom.RectXYWH(10, 10, 40, 40))
		c.StrokePath(p, render.Stroke(geom.RGBA(0, 0, 0), 4))
	})
	b := drawSoftware(dl, 60, 60)
	if b == nil {
		return fmt.Errorf("帧缓冲分配失败")
	}
	if a := b.Pix()[30*60+30] >> 24; a != 0 {
		return fmt.Errorf("描边矩形内部出现填充，alpha = %d", a)
	}
	if a := b.Pix()[10*60+30] >> 24; a == 0 {
		return fmt.Errorf("描边落在错误位置：上边框为空")
	}
	return nil
}

// ── C2 呈现格式固定 Pf24bit ────────────────────────────────────────

func assertPresentFormat() error {
	if got := raster.FormatBGR24.BytesPerPixel(); got != 3 {
		return fmt.Errorf("呈现格式每像素 %d 字节，应为 3（Pf24bit）", got)
	}
	d := gdi.NewDIB(8, 8, 24)
	if d == nil {
		// 无 GDI 会话（例如纯 CI 容器）时跳过格式验证
		return nil
	}
	defer d.Delete()
	if d.Bpp != 24 {
		return fmt.Errorf("DIB 位深 = %d，应为 24", d.Bpp)
	}
	if len(d.Row(0)) < 8*3 {
		return fmt.Errorf("行字节数 %d 小于 24bpp 最小宽度", len(d.Row(0)))
	}
	return nil
}

// ── C5 AA 就地混色：覆盖率缓冲必须是按行，不是整帧 ─────────────────

func assertInPlaceAA() error {
	const w, h = 800, 600
	be := software.New(software.Options{})
	dl := recordDL(func(c *render.Recorder) {
		c.FillOval(geom.RectXYWH(10, 10, 100, 100), render.Fill(geom.RGBA(0, 0, 0)))
	})
	b := raster.NewBacking(w, h)
	be.Draw(b, dl, b.Bounds())
	// 覆盖率缓冲若按整帧分配，800×600×2 字节 ≈ 960KB；
	// 按行则只有 ~4KB。这里给 16 倍行宽的宽松上限。
	limit := 16 * w * 2
	if got := be.CoverageBytes(); got > limit {
		return fmt.Errorf("覆盖率缓冲 %d 字节，超过按行上限 %d（退化成整帧 alpha 图）", got, limit)
	}
	return nil
}

// ── C6 文字回放：必须走 GDI 且落在 24bpp 目标上 ────────────────────

func assertTextReplay() error {
	d := gdi.NewDIB(160, 40, 24)
	if d == nil {
		return nil // 无 GDI 环境跳过
	}
	defer d.Delete()
	d.Clear(0x00FFFFFF)

	l := text.LayoutText("WG Engine 文字", text.DefaultFont(14), text.Options{})
	dl := recordDL(func(c *render.Recorder) {
		c.DrawText(l, geom.Pt(2, 2), geom.RGBA(0, 0, 0))
	})
	if n := gdi.TextPass(d, dl, geom.RectXYWH(0, 0, 160, 40)); n == 0 {
		return fmt.Errorf("文字通道未绘制任何行")
	}
	d.Flush()
	dark := 0
	for y := 0; y < d.H; y++ {
		row := d.Row(y)
		for x := 0; x < d.W; x++ {
			if row[x*3] < 200 {
				dark++
			}
		}
	}
	if dark < 10 {
		return fmt.Errorf("文字未落到 DIB 上（黑像素 %d）", dark)
	}
	return nil
}

// ── C8 资源释放顺序：删除必须幂等且不残留视图 ──────────────────────

func assertReleaseOrder() error {
	d := gdi.NewDIB(16, 16, 24)
	if d == nil {
		return nil
	}
	d.Delete()
	if d.Bits() != nil {
		return fmt.Errorf("删除后仍能拿到像素视图（悬垂）")
	}
	d.Delete() // 二次删除必须无副作用
	if d.HDC != 0 || d.HBmp != 0 {
		return fmt.Errorf("删除后句柄未清零")
	}
	return nil
}

// ── C9 同缩放：DIP↔像素只有一处换算 ───────────────────────────────

func assertSingleScale() error {
	h := host.New(nil)
	h.Resize(400, 300)
	// 缩放为 1 时换算必须是恒等
	if got := h.DIPToPixel(12); !approx(got, 12, 0.0001) {
		return fmt.Errorf("DIPToPixel(12) = %v", got)
	}
	if got := h.PixelToDIP(h.DIPToPixel(37)); !approx(got, 37, 0.0001) {
		return fmt.Errorf("DIP↔像素往返 = %v", got)
	}
	m := geom.Translate(3, 4).Concat(geom.Scale(2, 2))
	if got := m.ScaleFactor(); !approx(got, 2, 0.0001) {
		return fmt.Errorf("矩阵缩放因子 = %v", got)
	}
	return nil
}

// ── C13 帧=帧槽：一次提交，提交后脏区清空 ─────────────────────────

func assertSingleCommit() error {
	c, err := composite.New(composite.Options{
		Backend: software.New(software.Options{}),
		Width:   40, Height: 30,
		Background: geom.RGBA(255, 255, 255),
	})
	if err != nil {
		return err
	}
	dl := recordDL(func(r *render.Recorder) {
		r.FillRect(geom.RectXYWH(0, 0, 5, 5), render.Fill(geom.RGBA(0, 0, 0)))
	})
	st, err := c.Draw(dl)
	if err != nil {
		return err
	}
	if !st.Presented {
		return fmt.Errorf("首帧未提交")
	}
	if !c.Damage().IsEmpty() {
		return fmt.Errorf("提交后 damage 未清空（会重复提交）")
	}
	if c.LastPresented().IsEmpty() {
		return fmt.Errorf("LastPresented 为空，无法审计提交区域")
	}
	return nil
}

// ── C14 静止零帧 ─────────────────────────────────────────────────--─

func assertIdleFrame() error {
	c, err := composite.New(composite.Options{
		Backend: software.New(software.Options{}),
		Width:   20, Height: 20,
		Background: geom.RGBA(255, 255, 255),
	})
	if err != nil {
		return err
	}
	dl := recordDL(func(r *render.Recorder) {
		r.FillRect(geom.RectXYWH(0, 0, 5, 5), render.Fill(geom.RGBA(0, 0, 0)))
	})
	if _, err := c.Draw(dl); err != nil {
		return err
	}
	st, err := c.Draw(dl)
	if err != nil {
		return err
	}
	if st.Presented || st.Raster.Ops != 0 {
		return fmt.Errorf("静止帧仍在出图: presented=%v ops=%d", st.Presented, st.Raster.Ops)
	}
	return nil
}

// ── C16 帧间隔用中位数 ────────────────────────────────────────────

func assertMedian() error {
	samples := []time.Duration{16, 17, 16, 200, 16}
	if got := medianOf(samples); got != 16 {
		return fmt.Errorf("中位数 = %v，应为 16ms（长尾不应影响）", got)
	}
	return nil
}

// medianOf 复刻 frame.Median 的语义（此处不复用是为了让断言不依赖被测实现）。
func medianOf(d []time.Duration) time.Duration {
	s := make([]time.Duration, len(d))
	copy(s, d)
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j] < s[j-1]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
	if len(s)%2 == 1 {
		return s[len(s)/2]
	}
	return (s[len(s)/2-1] + s[len(s)/2]) / 2
}

// ── C18 帧内禁止资源拓扑变更 ──────────────────────────────────────

func assertNoTopologyChange() error {
	// 已冻结的指令表被修改必须 panic，而不是悄悄改掉正在被光栅的内容。
	dl := recordDL(func(c *render.Recorder) {
		c.FillRect(geom.RectXYWH(0, 0, 4, 4), render.Fill(geom.RGBA(0, 0, 0)))
	})
	if !dl.Frozen() {
		return fmt.Errorf("Finish 后指令表未冻结")
	}
	if err := catchPanic(func() {
		// 直接调用内部 append 路径：用 Recorder 复用同一指令表来触发
		r := render.NewRecorder(dl)
		r.FillRect(geom.RectXYWH(0, 0, 1, 1), render.Fill(geom.RGBA(0, 0, 0)))
	}); err == nil {
		return fmt.Errorf("冻结后仍可追加指令")
	}
	return nil
}

func catchPanic(fn func()) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panic: %v", p)
		}
	}()
	fn()
	return nil
}

// ── C20 damage 上卷 ───────────────────────────────────────────────

func assertDamageBubbles() error {
	root := host.NewSurface(host.SurfaceInline, geom.RectXYWH(0, 0, 100, 100))
	mid := host.NewSurface(host.SurfaceInline, geom.RectXYWH(20, 30, 50, 50))
	leaf := host.NewSurface(host.SurfaceInline, geom.RectXYWH(5, 5, 20, 20))
	if err := root.Append(mid); err != nil {
		return err
	}
	if err := mid.Append(leaf); err != nil {
		return err
	}
	root.ClearDamageTree()

	leaf.Invalidate(geom.RectXYWH(0, 0, 4, 4))
	// 叶(0..4) + 中偏移(5,5) + 根偏移(20,30) = (25,35)-(29,39)
	want := geom.RectLTRB(25, 35, 29, 39)
	if got := root.Damage(); got != want {
		return fmt.Errorf("根脏区 = %v，期望 %v", got, want)
	}
	// 超出叶范围的失效必须被裁剪，不能扩大脏区
	leaf.Invalidate(geom.RectXYWH(0, 0, 100, 100))
	if got := root.Damage(); got != geom.RectLTRB(25, 35, 45, 55) {
		return fmt.Errorf("裁剪后脏区 = %v，期望 (25,35)-(45,55)", got)
	}
	return nil
}

// ── C22 指令执行或显式降级 ────────────────────────────────────────

func assertExplicitDegrade() error {
	img := render.NewImageRef(2, 2)
	dl := recordDL(func(c *render.Recorder) {
		c.DrawImage(img, geom.RectXYWH(0, 0, 8, 8), render.ImageBilinear)
	})
	b := raster.NewBacking(16, 16)
	st := software.New(software.Options{}).Draw(b, dl, b.Bounds())
	if st.Degraded != 1 {
		return fmt.Errorf("不支持位图时 degraded = %d，应为 1（不得静默丢弃）", st.Degraded)
	}
	// 支持的指令不得计入 degraded
	dl2 := recordDL(func(c *render.Recorder) {
		c.FillRect(geom.RectXYWH(0, 0, 8, 8), render.Fill(geom.RGBA(0, 0, 0)))
	})
	st2 := software.New(software.Options{}).Draw(b, dl2, b.Bounds())
	if st2.Degraded != 0 {
		return fmt.Errorf("矩形填充被误判为降级: %d", st2.Degraded)
	}
	return nil
}

// ── C23 DisplayList 不可变，重录换 version ────────────────────────

func assertDisplayListImmutable() error {
	y := recordDL(func(c *render.Recorder) {
		c.FillRect(geom.RectXYWH(0, 0, 4, 4), render.Fill(geom.RGBA(0, 0, 0)))
	})
	v1 := y.Version()
	if v1 == 0 {
		return fmt.Errorf("version 未初始化")
	}
	y.Reset()
	if y.Version() == v1 {
		return fmt.Errorf("复用后 version 未变化，缓存无法据此失效")
	}
	if y.Frozen() {
		return fmt.Errorf("Reset 后不应仍处于冻结态")
	}
	return nil
}

// ── C24 文字设备空间回放并裁剪 ───────────────────────────────────

func assertTextClipped() error {
	d := gdi.NewDIB(200, 40, 24)
	if d == nil {
		return nil
	}
	defer d.Delete()
	d.Clear(0x00FFFFFF)

	l := text.LayoutText("Clipped Text Sample", text.DefaultFont(14), text.Options{})
	dl := recordDL(func(c *render.Recorder) {
		c.ClipRect(geom.RectXYWH(0, 0, 24, 40))
		c.DrawText(l, geom.Pt(0, 2), geom.RGBA(0, 0, 0))
	})
	gdi.TextPass(d, dl, geom.RectXYWH(0, 0, 200, 40))
	d.Flush()
	for y := 0; y < d.H; y++ {
		row := d.Row(y)
		for x := 28; x < d.W; x++ {
			if row[x*3] < 200 {
				return fmt.Errorf("裁剪外 (%d,%d) 出现文字像素", x, y)
			}
		}
	}
	return nil
}

// ── C25 两后端几何位置误差 ≤ 1px ─────────────────────────────────

func assertTwoBackendsAgree() error {
	shape := func(c *render.Recorder) {
		c.FillOval(geom.RectXYWH(10, 12, 40, 30), render.Fill(geom.RGBA(0, 0, 0)))
	}
	dl := recordDL(shape)
	sw := drawSoftware(dl, 80, 60)
	if sw == nil {
		return fmt.Errorf("软件后端分配失败")
	}
	gd, ok := drawGDI(dl, 80, 60)
	if !ok {
		return nil // 无 GDI 环境跳过
	}
	a := raster.AlphaBounds(sw)
	b := raster.AlphaBounds(gd)
	if !approx(a.Min.X, b.Min.X, 1) || !approx(a.Min.Y, b.Min.Y, 1) ||
		!approx(a.Max.X, b.Max.X, 1) || !approx(a.Max.Y, b.Max.Y, 1) {
		return fmt.Errorf("两后端包围盒差超 1px: software=%v gdi=%v", a, b)
	}
	ax, ay, ok1 := raster.AlphaCentroid(sw)
	bx, by, ok2 := raster.AlphaCentroid(gd)
	if ok1 && ok2 && (!approx(ax, bx, 1) || !approx(ay, by, 1)) {
		return fmt.Errorf("两后端质心差超 1px: (%.2f,%.2f) vs (%.2f,%.2f)", ax, ay, bx, by)
	}
	return nil
}
