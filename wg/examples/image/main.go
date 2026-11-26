package main

import (
	"image"
	"image/color"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
// 渐变合成图（audit 确定性——无网络/文件 IO）。
var gradientImg = func() image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, 120, 90))
	for y := 0; y < 90; y++ {
		for x := 0; x < 120; x++ {
			img.SetNRGBA(x, y, color.NRGBA{R: uint8(30 + x), G: uint8(90 + y), B: uint8(200 - x/2), A: 255})
		}
	}
	return img
}()

var (
	imgBase   = widgets.NewImage(gradientImg, 200, 150)
	imgBroken = widgets.NewImage(nil, 200, 150)

	// 多图组（PreviewGroup）：三张缩略图，点击任一张开预览可左右切换。
	imgAlt1  = widgets.NewImage(gradientImg, 120, 90)
	imgAlt2  = widgets.NewImage(reverseImg, 120, 90)
	imgAlt3  = widgets.NewImage(gradientImg, 120, 90)
	imgGroup = widgets.NewImageGroup(imgAlt1, imgAlt2, imgAlt3)
)

// reverseImg 反向渐变（与 gradientImg 区分）。
var reverseImg = func() image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, 120, 90))
	for y := 0; y < 90; y++ {
		for x := 0; x < 120; x++ {
			img.SetNRGBA(x, y, color.NRGBA{R: uint8(230 - x), G: uint8(60 + y/2), B: uint8(80 + x/3), A: 255})
		}
	}
	return img
}()

func main() {
	launcher.Main(doc.Meta{
		ID: "image", Title: "Image 图片", Group: doc.GroupData,
		What: []string{
			"内存位图显示（render.ImageFrom+CustomPaint 双线性拉伸；引擎无网络/文件 IO——Src 由宿主注入）",
			"**点击预览**：BandModal 遮罩+居中大图+右上关闭+Esc（Modal 同 overlay 模式）",
			"Src 为 nil 渲染加载失败占位（antd fallback：灰底+提示文本）",
			"预览工具条：放大/缩小/左转/右转（antd preview toolbar；缩放限 [1,50]、旋转 90°/次）",
		},
		Demos: []doc.Demo{
			{Title: "基础与预览", Note: "渐变合成图 200×150：点击弹预览，底部工具条放大/缩小/左转/右转（antd preview toolbar）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return imgBase
			}},
			{Title: "多图切换", Note: "ImageGroup：三张缩略图，点任一张开预览，左右按钮循环切图+计数 current/total", Build: func(ctx *ui.BuildContext) ui.Widget {
				return imgGroup
			}},
			{Title: "加载失败占位", Note: "Src=nil：灰底+「图片加载失败」文本（antd fallback）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return imgBroken
			}},
		},
		API: []doc.APIRow{
			{Name: "NewImage", Type: "func(src image.Image, w, h float32) *Image", Desc: "构造（内存位图）"},
			{Name: "Src", Type: "image.Image", Desc: "位图源（nil=加载失败占位；宿主注入）"},
			{Name: "Width / Height", Type: "float32", Desc: "显示尺寸"},
			{Name: "Radius", Type: "float32", Default: "4", Desc: "圆角"},
			{Name: "Preview", Type: "bool", Default: "true", Desc: "点击 BandModal 大图预览（Esc/关闭钮退出）"},
			{Name: "NewImageGroup", Type: "func(thumbs ...*Image) *ImageGroup", Desc: "多图预览组：缩略图并排，预览内左右切换+计数+OnChange"},
			{Name: "预览工具条", Type: "内置", Desc: "放大/缩小（1.5× 步进，限 [1,50]）/左转/右转（90°/次）"},
		},
	})
}
