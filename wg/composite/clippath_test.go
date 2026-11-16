package composite

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
)

// TestClipPathLayerIsExact 验证：节点级路径裁剪是靠蒙版精确裁的，
// 不是退化成包围盒。做法是让子节点画一块比路径更大的矩形，
// 只检查路径外的像素是否干净。
func TestClipPathLayerIsExact(t *testing.T) {
	tri := render.NewPath().MoveTo(50, 10).LineTo(90, 90).LineTo(10, 90).Close()
	child := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		cv.FillRect(geom.RectXYWH(0, 0, 100, 100), render.Fill(geom.RGBA(255, 0, 0)))
	}).SetSize2(100, 100)
	root := scene.Stack(child)
	root.SetSize2(100, 100)
	root.SetClip(&scene.ClipSpec{Path: tri, AA: true})
	root.Layout(scene.TightConstraints(geom.Sz(100, 100)))

	c := newComp(t, 120, 120)
	if root.Clip() == nil {
		t.Fatal("SetClip 未生效")
	}
	if root.Clip().Path == nil {
		t.Fatal("ClipSpec.Path 丢失")
	}
	tree := Build(root, BuildOptions{Cache: false})
	if tree.Stats().Clip == 0 {
		t.Fatalf("路径裁剪应建 Clip 层，实际 ring=%v stats=%+v layers=%d", tree.Ring(), tree.Stats(), tree.Layers())
	}
	c.DrawTree(tree)

	// 三角形外的左上角必须还是背景白；三角形内部必须被染成红。
	b := c.Backing()
	if p := b.Pix()[5*120+5]; p != 0xFFFFFFFF {
		t.Fatalf("路径外像素未裁掉: %08X", p)
	}
	if p := b.Pix()[80*120+50]; (p>>16)&0xFF < 200 {
		t.Fatalf("路径内像素被误裁: %08X", p)
	}
}
