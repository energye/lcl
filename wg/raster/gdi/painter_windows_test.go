//go:build windows

package gdi

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/text"
)

// textDL 造一份含一条文字指令的显示列表（identity 矩阵，落在 (6,6)）。
// 第二个返回值是指令矩形的近似面积来源（布局包围盒），供"不写成整块"断言用。
func textDL(s string, size float32, c geom.Color) (*render.DisplayList, geom.Rect) {
	dl := render.NewDisplayList(16)
	rec := render.NewRecorderAt(dl, geom.Identity, geom.RectXYWH(0, 0, 200, 60), true)
	lay := text.LayoutText(s, text.DefaultFont(size), text.Options{})
	rec.DrawText(lay, geom.Pt(6, 6), c)
	rect := geom.RectXYWH(6, 6, lay.Size.W, lay.Size.H)
	return rec.Finish(), rect
}

// premulBad 数出违反预乘不变式（RGB ≤ A）的像素数。
func premulBad(pix []uint32) int {
	bad := 0
	for _, px := range pix {
		a := int(px >> 24)
		if int((px>>16)&0xFF) > a || int((px>>8)&0xFF) > a || int(px&0xFF) > a {
			bad++
		}
	}
	return bad
}

// TestPaintTextAlphaPathTransparent 验证透明目标上的保 alpha 写入：
// 字形按覆盖率落进缓冲，而不是把整个文字矩形写成不透明（TextFlat 时代的行为）。
func TestPaintTextAlphaPathTransparent(t *testing.T) {
	const W, H = 120, 40
	dst := raster.NewBacking(W, H) // 零填充 = 全透明
	p := NewPainter()
	defer p.Close()
	dl, rect := textDL("wg", 20, geom.RGBA(255, 255, 255))
	if n := p.PaintText(dst, dl, geom.RectXYWH(0, 0, W, H)); n == 0 {
		t.Fatal("文字通道应画出内容")
	}

	area := int(rect.W()) * int(rect.H())
	opaque, covered := 0, 0
	pix := dst.Pix()
	for _, px := range pix {
		if a := px >> 24; a > 0 {
			covered++
			if a == 255 {
				opaque++
			}
		}
	}
	if covered == 0 || opaque == 0 {
		t.Fatalf("应有字形像素落进缓冲：覆盖 %d、不透明 %d", covered, opaque)
	}
	// 旧行为会把整个 cmd.Rect 写成不透明（≈area 像素）。字形墨迹通常只占
	// 矩形的一小部分，给一半的余量仍然能把"整块白盒"抓住。
	if opaque*2 > area {
		t.Fatalf("不透明像素 %d 超过文字矩形面积 %d 的一半：疑似整块写入", opaque, area)
	}
	if bad := premulBad(pix); bad != 0 {
		t.Fatalf("预乘不变式被破坏：%d 像素 RGB>A", bad)
	}
	// 文字矩形之外必须一个像素都没碰。
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if x >= int(rect.Min.X) && x < int(rect.Max.X) && y >= int(rect.Min.Y) && y < int(rect.Max.Y) {
				continue
			}
			if px := pix[y*W+x]; px != 0 {
				t.Fatalf("文字矩形外 (%d,%d) 被写入 %#08x", x, y, px)
			}
		}
	}
}

// TestPaintTextAlphaPathPreservesContent 验证保 alpha 写入不污染已有内容：
// 红条穿过文字区域，无字形的红条像素必须原样保留（旧行为会整条盖白）。
func TestPaintTextAlphaPathPreservesContent(t *testing.T) {
	const W, H = 120, 40
	dst := raster.NewBacking(W, H)
	dst.FillRect(geom.RectXYWH(0, 16, W, 8), geom.RGBA(255, 0, 0))
	p := NewPainter()
	defer p.Close()
	dl, rect := textDL("wg", 20, geom.RGBA(255, 255, 255))
	if n := p.PaintText(dst, dl, geom.RectXYWH(0, 0, W, H)); n == 0 {
		t.Fatal("文字通道应画出内容")
	}

	pix := dst.Pix()
	// 关键判定：文字矩形**外**的红条（条带贯穿整幅宽）必须原样保留。
	// 白盒式整块盖写只会污染 cmd.Rect 内部，这条还能抓出"矩形外没擦干净"的变体。
	for y := 16; y < 24; y++ {
		for x := 0; x < W; x++ {
			if x >= int(rect.Min.X) && x < int(rect.Max.X) {
				continue
			}
			if px := pix[y*W+x]; px != 0xFFFF0000 {
				t.Fatalf("文字矩形外红条 (%d,%d) 被改写为 %#08x", x, y, px)
			}
		}
	}
	// 矩形内、条带行：字形混合与保留并存，至少应有一个纯红像素幸存。
	survived := 0
	for y := 16; y < 24; y++ {
		for x := int(rect.Min.X); x < int(rect.Max.X); x++ {
			if pix[y*W+x] == 0xFFFF0000 {
				survived++
			}
		}
	}
	if survived == 0 {
		t.Fatal("文字矩形内条带应保留无字形像素（覆盖率写入不是整块覆盖）")
	}
	if bad := premulBad(pix); bad != 0 {
		t.Fatalf("预乘不变式被破坏：%d 像素 RGB>A", bad)
	}
}

// TestCompositeTextAlphaMath 用合成数据直接验证覆盖率提取的数学：
// 黑底趟 P_b = a·C（预乘文字色），白底趟 P_w = P_b + (1-a)·255。
func TestCompositeTextAlphaMath(t *testing.T) {
	dst := raster.NewBacking(3, 2)
	pix := dst.Pix()
	pix[0] = 0x00000000 // 透明
	pix[1] = 0xFFFFFFFF // 白
	pix[2] = 0xFFFF0000 // 红
	pix[3] = 0x00000000
	pix[4] = 0xFFFFFFFF
	pix[5] = 0xFFFF0000

	d := NewDIB(3, 2, 24)
	defer d.Delete()
	// 白字 a≈0.5：P_b=128，P_w=255
	for y := 0; y < 2; y++ {
		row := d.Row(y)
		for x := 0; x < 3; x++ {
			row[x*3], row[x*3+1], row[x*3+2] = 255, 255, 255
		}
	}
	saved := make([]byte, 3*2*3)
	for i := range saved {
		saved[i] = 128
	}
	compositeTextAlpha(dst, d, saved, geom.RectXYWH(0, 0, 3, 2))
	pix = dst.Pix()
	want := []uint32{
		0x80808080, // 透明底：预乘文字色 128，a=128
		0xFFFFFFFF, // 白底：全白（白字盖白底不可见）
		0xFFFF8080, // 红底：R 满、G/B 是文字贡献 128，a=255
		0x80808080,
		0xFFFFFFFF,
		0xFFFF8080,
	}
	for i, w := range want {
		if pix[i] != w {
			t.Fatalf("像素 %d = %#08x，应为 %#08x", i, pix[i], w)
		}
	}

	// 黑字 a≈0.5：黑底趟不可见（P_b=0），白底趟 P_w=255·(1-a)=128。
	for y := 0; y < 2; y++ {
		row := d.Row(y)
		for x := 0; x < 3; x++ {
			row[x*3], row[x*3+1], row[x*3+2] = 128, 128, 128
		}
	}
	for i := range saved {
		saved[i] = 0
	}
	dst2 := raster.NewBacking(3, 2)
	p2 := dst2.Pix()
	p2[0], p2[1], p2[2] = 0x00000000, 0xFFFFFFFF, 0xFFFF0000
	p2[3], p2[4], p2[5] = 0x00000000, 0xFFFFFFFF, 0xFFFF0000
	compositeTextAlpha(dst2, d, saved, geom.RectXYWH(0, 0, 3, 2))
	p2 = dst2.Pix()
	// 黑字 a=127：out = 0 + 128·dst/255。
	// 透明底 → 0x7F000000；白底 → 0xFF808080（被压暗一半）；红底 → 0xFF800000。
	wantBlack := []uint32{
		0x7F000000,
		0xFF808080,
		0xFF800000,
		0x7F000000,
		0xFF808080,
		0xFF800000,
	}
	for i, w := range wantBlack {
		if p2[i] != w {
			t.Fatalf("黑字像素 %d = %#08x，应为 %#08x", i, p2[i], w)
		}
	}
}
