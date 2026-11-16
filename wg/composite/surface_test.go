package composite

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
)

// TestSurfaceIsIsolatedGroup：SurfaceNode 不再扁平化，而是作为隔离组离屏合成。
// 隔离的意义：子树的重叠在半透明下按整组处理，不会出现"逐指令叠两次"的加深。
func TestSurfaceIsIsolatedGroup(t *testing.T) {
	inner := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(200, 40, 40))}).SetSize2(40, 40)
	surf := scene.SurfaceNode(inner).SetSize2(40, 40)
	root := scene.Stack(surf)
	root.SetSize2(60, 60)
	root.Layout(scene.TightConstraints(geom.Sz(60, 60)))

	tr := Build(root, BuildOptions{Cache: false})
	if tr.Stats().Surface != 1 {
		t.Fatalf("SurfaceNode 应建隔离组层: %+v", tr.Stats())
	}
	if tr.Stats().Downgraded != 0 {
		t.Fatalf("SurfaceNode 不应再降级: %+v", tr.Stats())
	}
	if err := tr.Validate(); err != nil {
		t.Fatal(err)
	}

	c := newComp(t, 60, 60)
	c.DrawTree(tr)
	if p := c.Backing().Pix()[20*60+20]; (p>>16)&0xFF < 180 {
		t.Fatalf("Surface 内容未合成: %08X", p)
	}
}

// TestSurfaceIsolatesGroupOpacity：组透明度必须整体生效，
// 内部重叠区不能因为逐指令乘 alpha 而变暗。
func TestSurfaceIsolatesGroupOpacity(t *testing.T) {
	a := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(255, 0, 0))}).SetSize2(40, 40)
	b := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(255, 0, 0))}).SetSize2(40, 40)
	stack := scene.Stack(a, b).SetSize2(40, 40)
	surf := scene.SurfaceNode(stack).SetSize2(40, 40)
	surf.SetOpacity(0.5)
	root := scene.Stack(surf)
	root.SetSize2(50, 50)
	root.Layout(scene.TightConstraints(geom.Sz(50, 50)))

	c := newComp(t, 50, 50)
	c.DrawTree(Build(root, BuildOptions{Cache: false}))
	// 红底叠白背：50% 时绿/蓝通道约 127；若逐指令叠两次会掉到 ~64。
	p := c.Backing().Pix()[10*50+10]
	g := (p >> 8) & 0xFF
	if g < 115 || g > 140 {
		t.Fatalf("组透明度未整体生效，重叠区绿=%d（期望 ~127）", g)
	}
}
