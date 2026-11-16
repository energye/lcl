package main

import (
	"fmt"
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
		ID: "host-dpi", Title: "DPI 与坐标换算", Group: doc.GroupVisual,
		What: []string{
			"MediaQuery.Scale 是唯一的 DPI 入口：DIP↔像素换算、对齐到设备像素网格（SnapRect）",
			"真窗里跟随窗体 DPI（WM_DPICHANGED 归一化重算）；headless 下 Scale=1",
		},
		Demos: []doc.Demo{
			{Title: "DIP 与像素对齐", Note: "半透明矩形按 DIP 绘制，描边是 SnapRect 对齐后的设备像素边界", Build: func(ctx *ui.BuildContext) ui.Widget {
				scale := ctx.Media.Scale
				if scale <= 0 {
					scale = 1
				}
				dip := func(px float32) float32 { return px / scale }
				pixels := func(v float32) float32 { return v * scale }
				snapRect := func(r geom.Rect) geom.Rect {
					return geom.RectLTRB(
						float32(math.Round(float64(r.Min.X*scale)))/scale,
						float32(math.Round(float64(r.Min.Y*scale)))/scale,
						float32(math.Round(float64(r.Max.X*scale)))/scale,
						float32(math.Round(float64(r.Max.Y*scale)))/scale,
					)
				}
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(geom.RGBA(248, 249, 252)))
					r := geom.RectXYWH(20, 20, 160, 90)
					cv.FillRect(r, render.Fill(geom.RGBAa(70, 130, 220, 90)))
					cv.StrokePath(render.NewPath().AddRect(snapRect(r)), render.Stroke(geom.RGBA(230, 80, 60), 2))
					for x := float32(0); x < sz.W; x += 8 {
						cv.DrawLine(x, 120, x, 140, render.Stroke(geom.RGBA(180, 185, 195), 1))
					}
				}, 240, 160)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("Scale=%.2f Percent=%d%% Identity=%v DIP(24)=%.1f Pixels(24)=%.1f",
					scale, int(scale*100+0.5), scale == 1, dip(24), pixels(24))))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "MediaQuery.Scale", Type: "float32", Default: "1（headless）", Desc: "DPI 缩放系数（96 基准）"},
			{Name: "MediaQuery.DPI/TextScale/ReduceMotion", Type: "字段", Desc: "设备 DPI / 文字缩放 / 动效偏好"},
			{Name: "host.NewDPI/DPIFor", Type: "host 层", Desc: "DIP↔像素换算（真窗归一化器使用）"},
			{Name: "SnapRect", Type: "对齐", Desc: "矩形对齐到设备像素网格（本窗演示同语义算法）"},
		},
	})
}
