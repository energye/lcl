package raster

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
)

func TestNewBacking(t *testing.T) {
	if b := NewBacking(0, 10); b != nil {
		t.Fatal("宽度 0 应返回 nil")
	}
	if b := NewBacking(10, -1); b != nil {
		t.Fatal("高度负数应返回 nil")
	}
	b := NewBacking(3, 2)
	if b == nil {
		t.Fatal("正常分配失败")
	}
	if b.W() != 3 || b.H() != 2 || b.Stride() != 12 {
		t.Fatalf("尺寸 = %dx%d stride=%d", b.W(), b.H(), b.Stride())
	}
	if len(b.Pix()) != 6 {
		t.Fatalf("像素数 = %d", len(b.Pix()))
	}
	if n := len(b.Bytes()); n != 24 {
		t.Fatalf("字节数 = %d", n)
	}
	if b.Opaque() {
		t.Fatal("新缓冲应为不透明黑之外的全透明，Opaque 应为 false")
	}
	if b.Bounds() != geom.RectXYWH(0, 0, 3, 2) {
		t.Fatalf("Bounds = %v", b.Bounds())
	}
	if b.Row(-1) != nil || b.Row(2) != nil {
		t.Fatal("越界行应返回 nil")
	}
	if len(b.Row(1)) != 3 {
		t.Fatalf("行长度 = %d", len(b.Row(1)))
	}
}

func TestFillAndOpaqueFlag(t *testing.T) {
	b := NewBacking(4, 4)
	b.Fill(geom.RGBA(1, 2, 3))
	if !b.Opaque() {
		t.Fatal("不透明满幅填充后 Opaque 应为 true")
	}
	// 局部不透明填充不应破坏整体不透明
	b.FillRect(geom.RectXYWH(0, 0, 1, 1), geom.RGBA(9, 9, 9))
	if !b.Opaque() {
		t.Fatal("局部不透明填充后 Opaque 应保持 true")
	}
	// 半透明填充后不再保证不透明
	b.BlendRect(geom.RectXYWH(1, 1, 2, 2), geom.RGBAa(255, 255, 255, 100))
	if b.Opaque() {
		t.Fatal("半透明混合后 Opaque 应为 false")
	}
	b.Clear()
	if b.Opaque() {
		t.Fatal("Clear 后 Opaque 应为 false")
	}
	for i, px := range b.Pix() {
		if px != 0 {
			t.Fatalf("Clear 后第 %d 个像素 = %#08x", i, px)
		}
	}
}

func TestFillRectClips(t *testing.T) {
	b := NewBacking(4, 4)
	// 超出边界的矩形应被裁剪，不 panic 且只填可见部分
	b.FillRect(geom.RectXYWH(-2, -2, 4, 4), geom.RGBA(255, 255, 255))
	if got := b.Pix()[0]; got != 0xFFFFFFFF {
		t.Fatalf("(0,0) = %#08x", got)
	}
	if got := b.Pix()[1*4+1]; got != 0xFFFFFFFF {
		t.Fatalf("(1,1) = %#08x", got)
	}
	if got := b.Pix()[2*4+2]; got != 0 {
		t.Fatalf("(2,2) 不该被填 = %#08x", got)
	}
	// 完全在外的矩形不改变任何像素
	before := b.Pix()[0]
	b.FillRect(geom.RectXYWH(100, 100, 10, 10), geom.RGBA(1, 1, 1))
	if b.Pix()[0] != before {
		t.Fatal("越界填充污染了缓冲")
	}
}

func TestBlendSourceOver(t *testing.T) {
	b := NewBacking(1, 1)
	b.Fill(geom.RGBA(0, 0, 255))                        // 不透明蓝
	b.BlendRect(b.Bounds(), geom.RGBAa(255, 0, 0, 128)) // 半透明红
	v := b.Pix()[0]
	r, g, bl, a := (v>>16)&0xFF, (v>>8)&0xFF, v&0xFF, v>>24
	if a != 255 {
		t.Fatalf("alpha = %d", a)
	}
	if r < 120 || r > 135 || bl < 120 || bl > 135 || g != 0 {
		t.Fatalf("混合结果 R=%d G=%d B=%d", r, g, bl)
	}
	// 全透明混合应无副作用
	b.Fill(geom.RGBA(0, 255, 0))
	b.BlendRect(b.Bounds(), geom.RGBAa(0, 0, 0, 0))
	if got := b.Pix()[0]; got != 0xFF00FF00 {
		t.Fatalf("全透明混合改变了像素: %#08x", got)
	}
}

func TestCopyRect(t *testing.T) {
	src := NewBacking(4, 4)
	src.Fill(geom.RGBA(0, 255, 0))
	dst := NewBacking(8, 8)
	dst.Fill(geom.RGBA(0, 0, 0))

	if !dst.CopyRect(src, src.Bounds(), 2, 3) {
		t.Fatal("拷贝失败")
	}
	if got := dst.Row(3)[2]; got != 0xFF00FF00 {
		t.Fatalf("目标像素 = %#08x", got)
	}
	if got := dst.Row(3)[1]; got != 0xFF000000 {
		t.Fatalf("左邻被污染: %#08x", got)
	}
	// 目标越界：只拷贝可见部分，不 panic
	if !dst.CopyRect(src, src.Bounds(), 6, 6) {
		t.Fatal("部分越界拷贝应成功")
	}
	if got := dst.Row(7)[7]; got != 0xFF00FF00 {
		t.Fatalf("(7,7) = %#08x", got)
	}
	// 完全越界返回 false
	if dst.CopyRect(src, src.Bounds(), 100, 100) {
		t.Fatal("完全越界应返回 false")
	}
	if dst.CopyRect(nil, geom.RectXYWH(0, 0, 1, 1), 0, 0) {
		t.Fatal("nil 源应返回 false")
	}
}

func TestBackingFromReuses(t *testing.T) {
	pix := make([]uint32, 4)
	b := NewBackingFrom(2, 2, pix)
	if b == nil {
		t.Fatal("复用构造失败")
	}
	b.Fill(geom.RGBA(255, 0, 0))
	if pix[0] != 0xFFFF0000 {
		t.Fatalf("未复用底层切片: %#08x", pix[0])
	}
	if NewBackingFrom(2, 2, make([]uint32, 3)) != nil {
		t.Fatal("长度不符应返回 nil")
	}
}

func TestFormatString(t *testing.T) {
	if FormatBGRA32Premul.BytesPerPixel() != 4 || FormatBGR24.BytesPerPixel() != 3 {
		t.Fatal("每像素字节数错误")
	}
	if FormatBGRA32Premul.String() != "bgra32-premul" || FormatBGR24.String() != "bgr24" {
		t.Fatal("格式名错误")
	}
}

func TestScrollBlitVertical(t *testing.T) {
	// 10x10，每行填行号，向上滚动 3 行。
	b := NewBacking(10, 10)
	for y := 0; y < 10; y++ {
		row := b.Row(y)
		for x := 0; x < 10; x++ {
			row[x] = uint32(y) << 24 // A=行号，便于断言
		}
	}
	exposed := b.ScrollBlit(geom.RectXYWH(0, 0, 10, 10), 0, -3)
	// 滚动后第 y 行应等于原第 y+3 行。
	for y := 0; y < 7; y++ {
		if got := b.Row(y)[0] >> 24; got != uint32(y+3) {
			t.Fatalf("行 %d 期望 %d 实得 %d", y, y+3, got)
		}
	}
	// 新暴露区应是底部 3 行。
	if exposed.IsEmpty() || exposed.Min.Y != 7 || exposed.Max.Y != 10 {
		t.Fatalf("新暴露区错误: %+v", exposed)
	}
}

func TestScrollBlitDown(t *testing.T) {
	b := NewBacking(10, 10)
	for y := 0; y < 10; y++ {
		row := b.Row(y)
		for x := 0; x < 10; x++ {
			row[x] = uint32(y) << 24
		}
	}
	// 向下滚动 2 行：重叠区必须从底部搬，避免覆盖。
	b.ScrollBlit(geom.RectXYWH(0, 0, 10, 10), 0, 2)
	for y := 2; y < 10; y++ {
		if got := b.Row(y)[0] >> 24; got != uint32(y-2) {
			t.Fatalf("行 %d 期望 %d 实得 %d", y, y-2, got)
		}
	}
}

func TestScrollBlitHorizontal(t *testing.T) {
	b := NewBacking(10, 4)
	for y := 0; y < 4; y++ {
		row := b.Row(y)
		for x := 0; x < 10; x++ {
			row[x] = uint32(x) << 24
		}
	}
	// 向右滚动 3 列：重叠区从右往左搬。
	b.ScrollBlit(geom.RectXYWH(0, 0, 10, 4), 3, 0)
	for x := 3; x < 10; x++ {
		if got := b.Row(0)[x] >> 24; got != uint32(x-3) {
			t.Fatalf("列 %d 期望 %d 实得 %d", x, x-3, got)
		}
	}
}
