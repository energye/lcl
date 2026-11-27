package diag

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/frame"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/text"
)

// ── --bench ─────────────────────────────────────────────────────────

// BenchOptions 是性能基准选项。
type BenchOptions struct {
	Width, Height int
	Frames        int
	// Components 是场景里组件数量的下限（每 4 个图元算 1 个组件）。
	Components int
	// DisableAA 关闭抗锯齿（A/B 对照，对应 WG_AA=0）。
	DisableAA bool
}

// BenchResult 是 --bench 的结果。
type BenchResult struct {
	Frames    int
	Median    time.Duration
	P95       time.Duration
	Best      time.Duration
	FPS       float64
	RasterOps int
	Damaged   float32 // 平均每帧提交面积
	Samples   []time.Duration

	// 分段耗时中位数（§6A 热点定位）：清屏 / 形状光栅 / 文字通道。
	MedianClear  time.Duration
	MedianRaster time.Duration
	MedianText   time.Duration
}

// OK 判定是否达到门禁：1440×880 / 20+ 组件 / AA on ≥ 59 FPS。
func (r BenchResult) OK() bool { return r.FPS >= 59 }

// String 返回可读摘要。
func (r BenchResult) String() string {
	return fmt.Sprintf("bench: %d 帧, 中位数 %v, P95 %v, 最好 %v, %.1f FPS, 平均提交 %.0f px², 分段 clear=%v raster=%v text=%v",
		r.Frames, r.Median, r.P95, r.Best, r.FPS, r.Damaged,
		r.MedianClear, r.MedianRaster, r.MedianText)
}

// Bench 在无窗口环境下测"一帧从录制到提交"的完整开销。
//
// 为什么用 headless 而不是真窗口：真窗口能测到 BitBlt 与 GDI 提交的真实成本，
// 但会把 GDI 驱动、桌面合成、垂直同步全部卷进来，导致数据不可比。M1 先测
// 引擎自身开销；真实窗口的数字在 M2 的 --bench（窗口模式）里补。
func Bench(opts BenchOptions) (BenchResult, error) {
	if opts.Width <= 0 {
		opts.Width = 1440
	}
	if opts.Height <= 0 {
		opts.Height = 880
	}
	if opts.Frames <= 0 {
		opts.Frames = 120
	}
	if opts.Components <= 0 {
		opts.Components = 24
	}

	wo := wg.DefaultOptions()
	wo.AA = !opts.DisableAA
	c := wg.NewHeadless(wo)
	defer c.Close()
	c.Resize(opts.Width, opts.Height)
	clk := frame.NewManualClock(0, frame.DefaultStep)
	c.SetClock(clk)

	paint := BenchScene(opts.Width, opts.Height, opts.Components)
	c.SetPainter(paint)

	var res BenchResult
	res.Samples = make([]time.Duration, 0, opts.Frames)
	clearSamples := make([]time.Duration, 0, opts.Frames)
	rasterSamples := make([]time.Duration, 0, opts.Frames)
	textSamples := make([]time.Duration, 0, opts.Frames)

	// 首帧包含分配与首次字体度量，单独计入但不参与统计
	if _, err := c.FrameOnce(); err != nil {
		return res, err
	}
	clk.Advance()

	for i := 0; i < opts.Frames; i++ {
		// 每帧都标脏，模拟持续动画（最坏情况）
		c.Root().Invalidate(c.Root().LocalBounds())
		st, err := c.FrameOnce()
		if err != nil {
			return res, err
		}
		res.Samples = append(res.Samples, st.FrameDuration)
		clearSamples = append(clearSamples, st.ClearDuration)
		rasterSamples = append(rasterSamples, st.RasterDuration)
		textSamples = append(textSamples, st.TextDuration)
		res.RasterOps += st.RasterOps
		res.Damaged += st.Damage.W() * st.Damage.H()
		clk.Advance()
	}
	if len(res.Samples) == 0 {
		return res, fmt.Errorf("没有采到样本")
	}
	res.Frames = len(res.Samples)
	res.Median = frame.Median(res.Samples)
	res.P95 = frame.Percentile(res.Samples, 95)
	res.MedianClear = frame.Median(clearSamples)
	res.MedianRaster = frame.Median(rasterSamples)
	res.MedianText = frame.Median(textSamples)
	res.Best = res.Samples[0]
	for _, d := range res.Samples {
		if d < res.Best {
			res.Best = d
		}
	}
	if m := float64(res.Median.Seconds()); m > 0 {
		res.FPS = 1 / m
	}
	res.Damaged /= float32(res.Frames)
	return res, nil
}

// BenchScene 生成一个 n 组件的绘制回调：每组含卡片底、圆角、圆点、
// 两条分隔线与一段文字，覆盖 UI 里最典型的图元组合。
func BenchScene(w, h, components int) wg.Painter {
	label := text.LayoutText("WG Engine 渲染基准 12.5ms", text.DefaultFont(13), text.Options{})
	return func(c render.Canvas) {
		c.FillRect(geom.RectXYWH(0, 0, float32(w), float32(h)), render.Fill(geom.RGBA(250, 250, 252)))
		cols := 4
		cellW := float32(w) / float32(cols)
		cellH := float32(96)
		for i := 0; i < components; i++ {
			col := i % cols
			row := i / cols
			x := float32(col)*cellW + 12
			y := float32(row)*cellH + 12
			cw := cellW - 24
			ch := cellH - 24
			if y+ch > float32(h) {
				break
			}
			// 卡片
			c.FillRRect(geom.RectXYWH(x, y, cw, ch), 8, render.Fill(geom.RGBA(255, 255, 255)))
			c.StrokePath(rounded(geom.RectXYWH(x, y, cw, ch), 8), render.Stroke(geom.RGBA(226, 232, 240), 1))
			// 状态点
			c.FillOval(geom.RectXYWH(x+12, y+12, 12, 12), render.Fill(geom.RGBA(59, 130, 246)))
			// 分隔线
			c.DrawLine(x+12, y+40, x+cw-12, y+40, render.Stroke(geom.RGBA(235, 238, 245), 1))
			// 文本
			c.DrawText(label, geom.Pt(x+12, y+46), geom.RGBA(51, 65, 85))
		}
	}
}

func rounded(r geom.Rect, radius float32) *render.Path {
	return render.NewPath().AddRRect(r, radius)
}

// ── --damage ────────────────────────────────────────────────────────

// DamageResult 是 --damage 的结果。
type DamageResult struct {
	Frames      int
	FullDamage  float32
	HoverDamage float32
	HoverBox    geom.Rect
	Ratio       float32
}

// OK 判定 hover 帧提交面积 ≤ 按钮包围盒 × 1.5。
func (r DamageResult) OK() bool { return r.Ratio <= 1.5 }

// String 返回可读摘要。
func (r DamageResult) String() string {
	return fmt.Sprintf("damage: 全帧 %.0f px², hover 帧 %.0f px²（按钮 %.0f px², 比值 %.2f）",
		r.FullDamage, r.HoverDamage, r.HoverBox.W()*r.HoverBox.H(), r.Ratio)
}

// Damage 验证"局部重绘只画局部"：首帧全屏，随后只让一个按钮变脏，
// 提交面积必须接近该按钮而不是整屏。
func Damage() (DamageResult, error) {
	var res DamageResult
	const w, h = 960, 640
	btn := geom.RectXYWH(120, 96, 160, 40)
	res.HoverBox = btn

	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(w, h)
	clk := frame.NewManualClock(0, frame.DefaultStep)
	c.SetClock(clk)

	hover := false
	c.SetPainter(func(cv render.Canvas) {
		cv.FillRect(geom.RectXYWH(0, 0, w, h), render.Fill(geom.RGBA(250, 250, 252)))
		bg := geom.RGBA(59, 130, 246)
		if hover {
			bg = geom.RGBA(37, 99, 235)
		}
		cv.FillRRect(btn, 6, render.Fill(bg))
		cv.FillRRect(geom.RectXYWH(btn.Min.X, btn.Max.Y+16, 200, 24), 4, render.Fill(geom.RGBA(226, 232, 240)))
	})

	if _, err := c.FrameOnce(); err != nil {
		return res, err
	}
	res.FullDamage = float32(w * h)
	clk.Advance()

	// 只让按钮脏（hover 变色）
	hover = true
	c.Root().Invalidate(btn)
	st, err := c.FrameOnce()
	if err != nil {
		return res, err
	}
	if !st.Presented {
		return res, fmt.Errorf("hover 帧未提交")
	}
	res.Frames = 2
	res.HoverDamage = st.Damage.W() * st.Damage.H()
	box := btn.W() * btn.H()
	if box > 0 {
		res.Ratio = res.HoverDamage / box
	}
	return res, nil
}

// ── --golden ────────────────────────────────────────────────────────

// GoldenResult 是 --golden 的结果。
type GoldenResult struct {
	Cases   int
	Failed  int
	MaxDiff float64
	Updated bool
	Elapsed time.Duration
}

// OK 判定全部通过（最大平均差 < 2%）。
func (r GoldenResult) OK() bool { return r.Failed == 0 }

// String 返回可读摘要。
func (r GoldenResult) String() string {
	mode := "比对"
	if r.Updated {
		mode = "更新参考图"
	}
	return fmt.Sprintf("golden: %s %d 张，失败 %d，最大平均差 %.3f%%，用时 %v",
		mode, r.Cases, r.Failed, r.MaxDiff*100, r.Elapsed)
}

// goldenScene 是参考图的绘制内容：图元 + 文字 + 渐变，覆盖典型组合。
func goldenScene(w, h int) wg.Painter {
	title := text.LayoutText("WG Engine Golden", text.DefaultFont(18).WithWeight(text.Bold), text.Options{})
	body := text.LayoutText("抗锯齿图元与文字渲染的回归基线。", text.DefaultFont(13), text.Options{MaxWidth: 220, Ellipsis: "…"})
	return func(c render.Canvas) {
		c.FillRect(geom.RectXYWH(0, 0, float32(w), float32(h)), render.Fill(geom.RGBA(252, 252, 254)))
		c.FillRRect(geom.RectXYWH(16, 16, 248, 148), 10, render.Fill(geom.RGBA(255, 255, 255)))
		c.StrokePath(rounded(geom.RectXYWH(16, 16, 248, 148), 10), render.Stroke(geom.RGBA(226, 232, 240), 1))
		c.DrawText(title, geom.Pt(32, 32), geom.RGBA(15, 23, 42))
		c.DrawText(body, geom.Pt(32, 62), geom.RGBA(71, 85, 105))
		g := render.NewLinearGradient(geom.Pt(32, 110), geom.Pt(240, 110),
			render.Stop{Offset: 0, Color: geom.RGBA(59, 130, 246)},
			render.Stop{Offset: 1, Color: geom.RGBA(16, 185, 129)})
		c.FillRRect(geom.RectXYWH(32, 110, 208, 24), 6, render.FillWith(g))
		c.DrawLine(32, 150, 240, 150, render.Stroke(geom.RGBA(203, 213, 225), 1))
		c.FillOval(geom.RectXYWH(290, 40, 80, 80), render.Fill(geom.RGBA(59, 130, 246)))
		c.DrawArc(geom.RectXYWH(290, 140, 80, 80), 0, 3.14159, render.Stroke(geom.RGBA(249, 115, 22), 6))
	}
}

// Golden 渲染固定场景并与参考 PNG 比对（平均绝对差 < 2%）。
//
// update=true 时写入参考图并跳过比对——参考图必须由人 review 后提交，
// 否则"更新基线"会变成"掩盖回归"。
func Golden(dir string, update bool) (GoldenResult, error) {
	began := time.Now()
	const w, h = 400, 200
	var res GoldenResult
	res.Updated = update

	if err := os.MkdirAll(dir, 0o755); err != nil {
		return res, err
	}

	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(w, h)
	c.SetClock(frame.NewManualClock(0, frame.DefaultStep))
	c.SetPainter(goldenScene(w, h))
	if _, err := c.FrameOnce(); err != nil {
		return res, err
	}
	b := c.Backing()
	if b == nil {
		return res, fmt.Errorf("没有合成结果")
	}

	got := toNRGBA(b)
	path := filepath.Join(dir, "golden-basic.png")
	res.Cases = 1

	if update {
		if err := savePNG(path, got); err != nil {
			return res, err
		}
		res.Elapsed = time.Since(began)
		return res, nil
	}

	ref, err := loadPNG(path)
	if err != nil {
		res.Failed++
		res.Elapsed = time.Since(began)
		return res, fmt.Errorf("缺少参考图 %s（首次请用 --golden -update 生成）: %w", path, err)
	}
	if ref.Bounds().Dx() != w || ref.Bounds().Dy() != h {
		res.Failed++
		res.Elapsed = time.Since(began)
		return res, fmt.Errorf("参考图尺寸 %v 与当前 %dx%d 不符", ref.Bounds(), w, h)
	}
	diff := meanAbsDiff(got, ref)
	res.MaxDiff = diff
	if diff > 0.02 {
		res.Failed++
	}
	res.Elapsed = time.Since(began)
	return res, nil
}

func toNRGBA(b *raster.Backing) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, b.W(), b.H()))
	for y := 0; y < b.H(); y++ {
		row := b.Row(y)
		for x := 0; x < b.W(); x++ {
			px := row[x]
			a := px >> 24
			r := (px >> 16) & 0xFF
			g := (px >> 8) & 0xFF
			bl := px & 0xFF
			if a == 0 {
				img.SetNRGBA(x, y, color.NRGBA{255, 255, 255, 255})
				continue
			}
			img.SetNRGBA(x, y, color.NRGBA{
				R: uint8(min32(255, r*255/a)),
				G: uint8(min32(255, g*255/a)),
				B: uint8(min32(255, bl*255/a)),
				A: 255,
			})
		}
	}
	return img
}

func savePNG(path string, img image.Image) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, img)
}

func loadPNG(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

// meanAbsDiff 计算两张图在 0..1 区间的平均绝对差（只看 RGB，忽略 alpha）。
func meanAbsDiff(a, b image.Image) float64 {
	ba, bb := a.Bounds(), b.Bounds()
	w, h := minInt(ba.Dx(), bb.Dx()), minInt(ba.Dy(), bb.Dy())
	if w == 0 || h == 0 {
		return 1
	}
	var sum float64
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r1, g1, b1, _ := a.At(ba.Min.X+x, ba.Min.Y+y).RGBA()
			r2, g2, b2, _ := b.At(bb.Min.X+x, bb.Min.Y+y).RGBA()
			sum += float64(absDiff16(r1, r2)+absDiff16(g1, g2)+absDiff16(b1, b2)) / 3 / 65535
		}
	}
	return sum / float64(w*h)
}

func absDiff16(a, b uint32) uint32 {
	if a > b {
		return a - b
	}
	return b - a
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
