package audit

import (
	"image"
	"testing"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
)

// TestGrabChannelOrder 钉死合成缓冲 BGRA→RGBA 的通道序。
// 通道颠倒时"变化量"判据无感（对称），但绝对色断言整片错色
// （试点实测 #BAE0FF 被读成 #FFE0BA）——这条契约必须在单测里锁住。
func TestGrabChannelOrder(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(64, 64)
	box := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(255, 0, 0))}).SetSize2(32, 32)
	c.Mount(box)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal("首帧失败:", err)
	}
	img := grab(c)
	if img == nil {
		t.Fatal("grab 返回 nil")
	}
	px := img.RGBAAt(4, 4)
	if px.R != 255 || px.G != 0 || px.B != 0 {
		t.Fatalf("通道序错误：got R=%d G=%d B=%d，want 255,0,0", px.R, px.G, px.B)
	}
}

func TestHexColor(t *testing.T) {
	col, note, ok := hexColor("#BAE0FF")
	if !ok || note != "#BAE0FF" {
		t.Fatalf("解析失败: ok=%v note=%q", ok, note)
	}
	if col.R != 0xBA || col.G != 0xE0 || col.B != 0xFF || col.A != 0xFF {
		t.Fatalf("色值错误: %+v", col)
	}
	if _, _, ok := hexColor("BAD"); ok {
		t.Fatal("非法 hex 应失败")
	}
}

func TestParseExpectedColor(t *testing.T) {
	col, note, ok := parseExpectedColor(map[string]any{"token": "Selection"})
	if !ok || note != "token:Selection" {
		t.Fatalf("token 解析失败: ok=%v note=%q", ok, note)
	}
	if col.R != 0xBA || col.G != 0xE0 || col.B != 0xFF {
		t.Fatalf("Selection token 值错误: %+v", col)
	}
	col, note, ok = parseExpectedColor(map[string]any{"hex": "#1677FF", "note": "Primary"})
	if !ok || note != "Primary" || col.R != 0x16 || col.G != 0x77 || col.B != 0xFF {
		t.Fatalf("hex+note 解析失败: ok=%v note=%q col=%+v", ok, note, col)
	}
	if _, _, ok := parseExpectedColor(42); ok {
		t.Fatal("未知类型应失败")
	}
}

func TestClickPointFrac(t *testing.T) {
	b := geom.RectXYWH(10, 20, 100, 40)
	c := clickPoint(b, nil)
	if c.X != 60 || c.Y != 40 {
		t.Fatalf("默认中心应为 (60,40)，got %v", c)
	}
	c = clickPoint(b, []float32{0.25, 0.75})
	if c.X != 35 || c.Y != 50 {
		t.Fatalf("frac 点应为 (35,50)，got %v", c)
	}
}

func TestRegionChangedThresh(t *testing.T) {
	a := image.NewRGBA(image.Rect(0, 0, 8, 8))
	b := image.NewRGBA(image.Rect(0, 0, 8, 8))
	for i := 3; i < len(a.Pix); i += 4 {
		a.Pix[i-3], a.Pix[i-2], a.Pix[i-1], a.Pix[i] = 250, 250, 250, 255
		b.Pix[i-3], b.Pix[i-2], b.Pix[i-1], b.Pix[i] = 245, 245, 245, 255
	}
	if got := regionChangedThresh(a, b, geom.RectXYWH(0, 0, 8, 8), 3); got != 1 {
		t.Fatalf("5 色阶差应被低阈值检出，got %.3f", got)
	}
	if got := regionChanged(a, b, geom.RectXYWH(0, 0, 8, 8)); got != 0 {
		t.Fatalf("5 色阶差应被默认阈值忽略，got %.3f", got)
	}
}
