package main

import (
	"math"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "scene-nodes", Title: "场景节点构造器", Group: doc.GroupLayout,
		What: []string{
			"scene 包的特殊节点：PathNode/ImageNode/Clip/Transform/Opacity/Effect/SurfaceNode/CustomPaintNode",
			"每个特殊节点触发合成层树里的一种组语义（裁剪/透明/滤镜/隔离）",
		},
		Demos: []doc.Demo{
			{Title: "七种特殊节点", Note: "PathNode 三角、ImageNode 渐变图、Clip 圆角、Transform 旋转、Opacity 45%、Effect 阴影、SurfaceNode 隔离组", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				tri := render.NewPath().MoveTo(0, 46).LineTo(46, 0).LineTo(92, 46).Close()
				pathN := scene.PathNode(scene.PathSpec{Path: tri, Paint: render.Fill(th.Color.Primary)}).SetSize2(100, 50)
				img := render.NewImageRef(24, 24)
				for y := 0; y < 24; y++ {
					for x := 0; x < 24; x++ {
						img.Pix[y*img.Stride+x] = geom.RGBA(uint8(x*10), uint8(y*10), 200).PremulBGRA()
					}
				}
				imageN := scene.ImageNode(scene.ImageSpec{Image: img, Fit: scene.FitFill}).SetSize2(80, 50)
				clipN := scene.Clip(scene.ClipSpec{Radius: 14}, doc.Box(ctx, "Clip 圆角", th.Color.Info, 140, 60))
				transN := scene.Transform(geom.Rotate(math.Pi/12), doc.Box(ctx, "Transform", th.Color.Warning, 130, 50))
				opN := scene.Opacity(0.45, doc.Box(ctx, "Opacity 45%", th.Color.Success, 120, 50))
				effN := scene.Effect(scene.EffectSpec{Shadow: &scene.ShadowSpec{Dx: 3, Dy: 4, Blur: 6, Color: geom.RGBAa(0, 0, 0, 90)}},
					doc.Box(ctx, "Effect 阴影", th.Color.Surface, 130, 50))
				surfN := scene.SurfaceNode(doc.Box(ctx, "SurfaceNode", th.Color.SurfaceActive, 130, 50))

				grid := scene.Wrap(scene.Horizontal, 12, 12).SetWidth(680)
				for _, n := range []*scene.Node{pathN, imageN, clipN, transN, opN, effN, surfN} {
					grid.Add(n)
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return grid })
			}},
		},
		API: []doc.APIRow{
			{Name: "scene.PathNode", Type: "func(PathSpec) *Node", Desc: "矢量路径节点（Path+Paint）"},
			{Name: "scene.ImageNode", Type: "func(ImageSpec) *Node", Desc: "位图节点；Fit 填充模式"},
			{Name: "scene.Clip", Type: "func(ClipSpec, child) *Node", Desc: "裁剪组（Rect/Radius/Path）"},
			{Name: "scene.Transform", Type: "func(Matrix, child) *Node", Desc: "仿射变换（不影响布局尺寸）"},
			{Name: "scene.Opacity", Type: "func(alpha float32, child) *Node", Desc: "组透明度（离屏合成）"},
			{Name: "scene.Effect", Type: "func(EffectSpec, child) *Node", Desc: "组滤镜（阴影/模糊/发光）"},
			{Name: "scene.SurfaceNode", Type: "func(child) *Node", Desc: "隔离组（独立离屏+合成）"},
			{Name: "scene.CustomPaintNode", Type: "func(fn(render.Canvas, Size)) *Node", Desc: "自绘节点（C17 允许的录制通道）"},
		},
	})
}
