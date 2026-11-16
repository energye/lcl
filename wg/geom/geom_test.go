package geom

import "testing"

func TestRectBasics(t *testing.T) {
	r := RectXYWH(10, 20, 30, 40)
	if r.Left() != 10 || r.Top() != 20 || r.Right() != 40 || r.Bottom() != 60 {
		t.Fatalf("四边 = %v", r)
	}
	if got := r.Size(); got != (Size{30, 40}) {
		t.Fatalf("Size = %v", got)
	}
	if got := r.Center(); got != (Point{25, 40}) {
		t.Fatalf("Center = %v", got)
	}
	if !r.Contains(Pt(10, 20)) || r.Contains(Pt(40, 60)) {
		t.Fatalf("Contains 应为半开区间")
	}
	if r.IsEmpty() {
		t.Fatalf("非空矩形被判为空")
	}
}

func TestRectIntersectUnion(t *testing.T) {
	a := RectXYWH(0, 0, 10, 10)
	b := RectXYWH(5, 5, 10, 10)
	if got := a.Intersect(b); got != RectLTRB(5, 5, 10, 10) {
		t.Fatalf("Intersect = %v", got)
	}
	if got := a.Union(b); got != RectLTRB(0, 0, 15, 15) {
		t.Fatalf("Union = %v", got)
	}
	c := RectXYWH(100, 100, 5, 5)
	if got := a.Intersect(c); got != ZeroRect {
		t.Fatalf("无交集应返回 ZeroRect，实际 %v", got)
	}
	// 相邻不算相交（半开区间）
	d := RectXYWH(10, 0, 5, 5)
	if a.Intersects(d) {
		t.Fatalf("相邻矩形不应相交")
	}
}

func TestRectRound(t *testing.T) {
	r := RectXYWH(1.2, 2.7, 3.3, 4.4)
	if got := r.RoundOut(); got != RectLTRB(1, 2, 5, 8) {
		t.Fatalf("RoundOut = %v", got)
	}
	if got := r.RoundIn(); got != RectLTRB(2, 3, 4, 7) {
		t.Fatalf("RoundIn = %v", got)
	}
	thin := RectXYWH(1.1, 1.1, 0.2, 0.2)
	if got := thin.RoundIn(); got != ZeroRect {
		t.Fatalf("亚像素矩形 RoundIn 应为 ZeroRect，实际 %v", got)
	}
}

func TestMatrixRoundTrip(t *testing.T) {
	m := Translate(10, 20).Concat(Rotate(0.3)).Concat(Scale(2, 3))
	inv, ok := m.Invert()
	if !ok {
		t.Fatal("矩阵应可逆")
	}
	p := Pt(7, -3)
	back := inv.MapPoint(m.MapPoint(p))
	if d := back.Sub(p).Len(); d > 0.01 {
		t.Fatalf("往返误差 %v", d)
	}
}

func TestMatrixMapRect(t *testing.T) {
	m := Translate(5, 5)
	if got := m.MapRect(RectXYWH(0, 0, 10, 10)); got != RectLTRB(5, 5, 15, 15) {
		t.Fatalf("平移后矩形 = %v", got)
	}
	// 旋转 90°：宽高互换
	r := Rotate(1.5707963).MapRect(RectXYWH(0, 0, 10, 20))
	if w, h := r.W(), r.H(); w < 19.9 || w > 20.1 || h < 9.9 || h > 10.1 {
		t.Fatalf("旋转后外包尺寸 = %vx%v", w, h)
	}
	if !Identity.IsIdentity() || Translate(1, 0).IsIdentity() {
		t.Fatal("IsIdentity 判定错误")
	}
	if !Translate(1, 2).IsAxisAligned() || !Scale(2, 2).IsAxisAligned() {
		t.Fatal("IsAxisAligned 判定错误")
	}
	if Rotate(0.1).IsAxisAligned() {
		t.Fatal("旋转不该被判为轴对齐")
	}
}

func TestColorPremul(t *testing.T) {
	// 完全不透明红：打包值 0xFFFF0000，小端内存为 B=0 G=0 R=255 A=255
	if got := RGBA(255, 0, 0).PremulBGRA(); got != 0xFFFF0000 {
		t.Fatalf("不透明红 = %#08x", got)
	}
	if got := (Color{}).PremulBGRA(); got != 0 {
		t.Fatalf("全透明 = %#08x", got)
	}
	// 半透明红：R=255*128/255=128
	got := RGBAa(255, 0, 0, 128).PremulBGRA()
	if a, r := got>>24, (got>>16)&0xFF; a != 128 || r != 128 {
		t.Fatalf("半透明红 A=%d R=%d", a, r)
	}
	if !RGBA(0, 0, 0).IsOpaque() || RGBAa(0, 0, 0, 0).IsOpaque() {
		t.Fatal("IsOpaque 判定错误")
	}
	if got := RGBA(0, 0, 0).WithAlpha(10).A; got != 10 {
		t.Fatalf("WithAlpha = %d", got)
	}
}

func TestColorLerp(t *testing.T) {
	a := RGBA(0, 0, 0)
	b := RGBA(255, 255, 255)
	if got := a.Lerp(b, 0.5); got != (Color{127, 127, 127, 255}) && got != (Color{128, 128, 128, 255}) {
		t.Fatalf("中点 = %v", got)
	}
	if got := a.Lerp(b, -1); got != a {
		t.Fatalf("t<0 应原样返回")
	}
	if got := a.Lerp(b, 2); got != b {
		t.Fatalf("t>1 应返回目标")
	}
}

func TestInsets(t *testing.T) {
	in := InLTRB(1, 2, 3, 4)
	if in.Horizontal() != 4 || in.Vertical() != 6 {
		t.Fatalf("水平和/垂直和错误")
	}
	r := RectXYWH(0, 0, 100, 100)
	if got := in.Deflate(r); got != RectLTRB(1, 2, 97, 96) {
		t.Fatalf("Deflate = %v", got)
	}
	if got := in.Inflate(r); got != RectLTRB(-1, -2, 103, 104) {
		t.Fatalf("Inflate = %v", got)
	}
}
