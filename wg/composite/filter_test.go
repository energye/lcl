package composite

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/scene"
)

// TestBlurBackingSoftensEdge：box 模糊必须把一条硬边变成渐变。
func TestBlurBackingSoftensEdge(t *testing.T) {
	b := raster.NewBacking(40, 8)
	for y := 0; y < 8; y++ {
		for x := 0; x < 20; x++ {
			b.Pix()[y*b.W()+x] = geom.RGBA(0, 0, 0).PremulBGRA()
		}
	}
	blurBacking(b, 3)
	// 边界附近应出现中间值；远处保持两端。
	mid := 0
	for y := 0; y < 8; y++ {
		for x := 15; x < 25; x++ {
			a := b.Pix()[y*b.W()+x] >> 24
			if a > 10 && a < 245 {
				mid++
			}
		}
	}
	if mid == 0 {
		t.Fatal("模糊后没有过渡像素，仍是硬边")
	}
	if a := b.Pix()[0] >> 24; a < 200 {
		t.Fatalf("远离边界的实心区被过度模糊: %d", a)
	}
}

// TestSoftShadowRendersGradient：Blur>0 的盒阴影必须产生半透明过渡，
// 且阴影确实落在内容下方（内容颜色不被污染）。
func TestSoftShadowRendersGradient(t *testing.T) {
	card := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(geom.RGBA(255, 255, 255)),
		Radius: 8,
		Shadow: &scene.ShadowSpec{Dx: 0, Dy: 6, Blur: 6, Color: geom.RGBAa(0, 0, 0, 120)},
	}).SetSize2(60, 40)

	root := scene.Stack(card)
	root.SetSize2(120, 120)
	root.Layout(scene.TightConstraints(geom.Sz(120, 120)))

	tree := Build(root, BuildOptions{Cache: false})
	if tree.Stats().Filter == 0 {
		t.Fatalf("软阴影应建滤镜层，实际 %+v", tree.Stats())
	}
	if err := tree.Validate(); err != nil {
		t.Fatal(err)
	}
	c := newComp(t, 120, 120)
	c.DrawTree(tree)
	b := c.Backing()

	// 卡片下方（阴影区）应有半透明过渡：既不是纯白，也不是纯黑。
	grad := 0
	for y := 46; y < 66; y++ {
		for x := 10; x < 50; x++ {
			p := b.Pix()[y*b.W()+x]
			r := (p >> 16) & 0xFF
			if r > 60 && r < 245 {
				grad++
			}
		}
	}
	if grad < 50 {
		t.Fatalf("软阴影没有渐变过渡，过渡像素=%d", grad)
	}
	// 卡片内部必须仍是纯白（阴影不能画在内容上面）。
	p := b.Pix()[20*b.W()+30]
	if r := (p >> 16) & 0xFF; r != 255 {
		t.Fatalf("阴影污染了卡片内容: %08X", p)
	}
}

// TestGlowExtendsOutsideContent：发光必须在内容之外还有像素。
func TestGlowExtendsOutsideContent(t *testing.T) {
	dot := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(80, 160, 255)), Radius: 6}).SetSize2(20, 20)
	wrapped := scene.Effect(scene.EffectSpec{
		GlowRadius: 8,
		GlowColor:  geom.RGBAa(80, 160, 255, 255),
	}, dot)
	root := scene.Stack(wrapped)
	root.SetSize2(100, 100)
	root.Layout(scene.TightConstraints(geom.Sz(100, 100)))

	tree := Build(root, BuildOptions{Cache: false})
	c := newComp(t, 100, 100)
	c.DrawTree(tree)
	b := c.Backing()

	// 内容在 (0,0)-(20,20)，发光应在 (25..35) 附近有可观测颜色。
	hit := 0
	for y := 5; y < 20; y++ {
		for x := 24; x < 34; x++ {
			p := b.Pix()[y*b.W()+x]
			if (p>>16)&0xFF < 250 || (p>>8)&0xFF < 250 || p&0xFF < 250 {
				hit++
			}
		}
	}
	if hit == 0 {
		t.Fatal("发光没有扩散到内容之外")
	}
}

// TestGroupBlurSoftensContent：BlurSigma 必须真的模糊内容。
func TestGroupBlurSoftensContent(t *testing.T) {
	sq := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(0, 0, 0))}).SetSize2(20, 20)
	blurred := scene.Effect(scene.EffectSpec{BlurSigma: 4}, sq)
	root := scene.Stack(blurred)
	root.SetSize2(80, 80)
	root.Layout(scene.TightConstraints(geom.Sz(80, 80)))

	tree := Build(root, BuildOptions{Cache: false})
	c := newComp(t, 80, 80)
	c.DrawTree(tree)
	b := c.Backing()
	mid := 0
	for y := 5; y < 15; y++ {
		for x := 18; x < 26; x++ {
			r := (b.Pix()[y*b.W()+x] >> 16) & 0xFF
			if r > 20 && r < 235 {
				mid++
			}
		}
	}
	if mid == 0 {
		t.Fatal("组模糊没有产生过渡")
	}
}

