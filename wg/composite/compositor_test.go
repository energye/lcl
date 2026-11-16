package composite

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/raster/software"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/text"
)

func newComp(t *testing.T, w, h int) *Compositor {
	t.Helper()
	c, err := New(Options{
		Backend:       software.New(software.Options{}),
		Width:         w,
		Height:        h,
		Background:    geom.RGBA(255, 255, 255),
	})
	if err != nil {
		t.Fatalf("创建合成器失败: %v", err)
	}
	return c
}

func dlOf(fn func(c *render.Recorder)) *render.DisplayList {
	r := render.NewRecorder(render.NewDisplayList(16))
	fn(r)
	return r.Finish()
}

func TestFirstFramePresents(t *testing.T) {
	c := newComp(t, 40, 40)
	dl := dlOf(func(r *render.Recorder) {
		r.FillRect(geom.RectXYWH(5, 5, 10, 10), render.Fill(geom.RGBA(255, 0, 0)))
	})
	st, err := c.Draw(dl)
	if err != nil {
		t.Fatal(err)
	}
	if !st.Presented {
		t.Fatal("首帧应提交（resize 已把整屏标脏）")
	}
	if st.DamageArea != 40*40 {
		t.Fatalf("首帧脏区面积 = %v，应为整屏", st.DamageArea)
	}
	// 背景应是白色，矩形内是红色
	if px := c.Backing().Pix()[0]; px>>24 != 255 || (px>>16)&0xFF != 255 {
		t.Fatalf("背景像素 = %#08x", px)
	}
	if px := c.Backing().Pix()[10*40+10]; (px>>16)&0xFF != 255 || px&0xFF != 0 {
		t.Fatalf("矩形像素 = %#08x", px)
	}
}

func TestIdleFrameDoesNotPresent(t *testing.T) {
	c := newComp(t, 40, 40)
	dl := dlOf(func(r *render.Recorder) {
		r.FillRect(geom.RectXYWH(5, 5, 10, 10), render.Fill(geom.RGBA(255, 0, 0)))
	})
	if _, err := c.Draw(dl); err != nil {
		t.Fatal(err)
	}
	// 第二次：没有新的 damage → 静止帧（C14）
	st, err := c.Draw(dl)
	if err != nil {
		t.Fatal(err)
	}
	if st.Presented {
		t.Fatalf("静止帧不该提交: %+v", st)
	}
	if st.Raster.Ops != 0 {
		t.Fatalf("静止帧不该跑指令: %+v", st.Raster)
	}
}

func TestDamageLimitedRepaint(t *testing.T) {
	c := newComp(t, 100, 100)
	dl := dlOf(func(r *render.Recorder) {
		r.FillRect(geom.RectXYWH(10, 10, 20, 20), render.Fill(geom.RGBA(0, 0, 255)))
	})
	if _, err := c.Draw(dl); err != nil {
		t.Fatal(err)
	}

	// 只让按钮区域脏
	btn := geom.RectXYWH(10, 10, 20, 20)
	c.Invalidate(btn)
	st, err := c.Draw(dl)
	if err != nil {
		t.Fatal(err)
	}
	if !st.Presented {
		t.Fatal("有 damage 应提交")
	}
	// 提交面积必须接近按钮面积，不能是整屏
	if st.DamageArea > btn.W()*btn.H()*1.5 {
		t.Fatalf("脏区面积 = %v，应 ≤ %v", st.DamageArea, btn.W()*btn.H()*1.5)
	}
	if st.Raster.Pixels <= 0 {
		t.Fatal("应实际触碰像素")
	}
}

func TestDamageIsClearedAfterPresent(t *testing.T) {
	c := newComp(t, 20, 20)
	dl := dlOf(func(r *render.Recorder) {
		r.FillRect(geom.RectXYWH(0, 0, 5, 5), render.Fill(geom.RGBA(0, 0, 0)))
	})
	if _, err := c.Draw(dl); err != nil {
		t.Fatal(err)
	}
	if !c.Damage().IsEmpty() {
		t.Fatalf("提交后 damage 应清空: %v", c.Damage().Rects())
	}
	if c.LastPresented().IsEmpty() {
		t.Fatal("LastPresented 应记录上一帧")
	}
}

func TestResizeInvalidatesAll(t *testing.T) {
	c := newComp(t, 20, 20)
	dl := dlOf(func(r *render.Recorder) {})
	if _, err := c.Draw(dl); err != nil {
		t.Fatal(err)
	}
	if err := c.Resize(40, 30); err != nil {
		t.Fatal(err)
	}
	st, err := c.Draw(dl)
	if err != nil {
		t.Fatal(err)
	}
	if !st.Presented || st.DamageArea != 40*30 {
		t.Fatalf("resize 后应整屏提交: %+v", st)
	}
}

func TestTextWithoutPainterIsReported(t *testing.T) {
	c := newComp(t, 100, 30)
	dl := dlOf(func(r *render.Recorder) {
		r.DrawText(testLayout(), geom.Pt(2, 2), geom.RGBA(0, 0, 0))
	})
	st, err := c.Draw(dl)
	if err != nil {
		t.Fatal(err)
	}
	if !st.TextSkipped {
		t.Fatalf("没有文字通道时应报告 TextSkipped: %+v", st)
	}
}

// degradeBackend 是"永远降级"的后端桩：用来验证 Compositor 把后端的
// degraded 计数如实透传，而不是验证某个具体能力缺失。
type degradeBackend struct{}

func (degradeBackend) Name() string          { return "stub-degrade" }
func (degradeBackend) Caps() raster.Caps     { return raster.Caps{} }
func (degradeBackend) Draw(*raster.Backing, *render.DisplayList, geom.Rect) raster.Stats {
	return raster.Stats{Degraded: 1}
}

func TestDegradedIsPropagated(t *testing.T) {
	c, err := New(Options{
		Backend:    degradeBackend{},
		Width:      40,
		Height:     40,
		Background: geom.RGBA(255, 255, 255),
	})
	if err != nil {
		t.Fatal(err)
	}
	dl := dlOf(func(r *render.Recorder) {
		r.FillRect(geom.RectXYWH(0, 0, 10, 10), render.Fill(geom.RGBA(0, 0, 0)))
	})
	st, err := c.Draw(dl)
	if err != nil {
		t.Fatal(err)
	}
	if st.Raster.Degraded != 1 {
		t.Fatalf("degraded 未透传: %+v", st.Raster)
	}
}

func TestZeroSizeRejected(t *testing.T) {
	if _, err := New(Options{Backend: software.New(software.Options{})}); err == nil {
		t.Fatal("尺寸为 0 应报错")
	}
	if _, err := New(Options{Width: 10, Height: 10}); err == nil {
		t.Fatal("缺少后端应报错")
	}
}

// testLayout 造一份文本布局（度量精度不影响本测试，只验证通道是否被调用）。
func testLayout() *text.Layout {
	return text.LayoutText("hello", text.DefaultFont(14), text.Options{})
}

var _ = raster.Stats{}
