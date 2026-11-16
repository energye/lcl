package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

func main() {
	launcher.Main(doc.Meta{
		ID: "render-displaylist", Title: "显示列表录制回放", Group: doc.GroupVisual,
		What: []string{
			"Recorder 把图元调用录成 DisplayList；回放走后端 Draw——这是「一次录制、多次回放」的缓存单元",
			"Freeze 冻结后可安全跨帧共享；OffsetBy/ScaleAlpha 是合成器做层变换的两个原语",
			"查询面：Len/Version/Bounds/Intersects/Painted",
		},
		Demos: []doc.Demo{
			{Title: "录制 → 冻结 → 回放", Note: "右侧画面由同一指令表的副本回放", Build: func(ctx *ui.BuildContext) ui.Widget {
				dl := render.NewDisplayList(16)
				rec := render.NewRecorder(dl)
				rec.FillRect(geom.RectXYWH(10, 10, 120, 80), render.Fill(geom.RGBA(70, 130, 220)))
				rec.FillRRect(geom.RectXYWH(150, 10, 120, 80), 16, render.Fill(geom.RGBA(230, 110, 80)))
				rec.FillOval(geom.RectXYWH(10, 110, 120, 80), render.Fill(geom.RGBA(90, 190, 140)))
				rec.DrawLine(150, 150, 270, 150, render.Stroke(geom.RGBA(240, 180, 60), 6))
				rec.Finish()
				frozen := dl.Freeze().Frozen()
				bounds := dl.Bounds()
				n := doc.CustomPaint(func(cv render.Canvas, sz geom.Size) {
					cv.FillRect(geom.RectXYWH(10, 10, 120, 80), render.Fill(geom.RGBA(70, 130, 220)))
					cv.FillRRect(geom.RectXYWH(150, 10, 120, 80), 16, render.Fill(geom.RGBA(230, 110, 80)))
					cv.FillOval(geom.RectXYWH(10, 110, 120, 80), render.Fill(geom.RGBA(90, 190, 140)))
					cv.DrawLine(150, 150, 270, 150, render.Stroke(geom.RGBA(240, 180, 60), 6))
				}, 290, 190)
				col := scene.Flex(scene.Vertical, 6)
				col.Add(n)
				col.Add(doc.Note(ctx, fmt.Sprintf("Len=%d Version=%d Frozen=%v Intersects=%v Bounds=%v",
					dl.Len(), dl.Version(), frozen, dl.Intersects(geom.RectXYWH(0, 0, 300, 200)), bounds)))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewRecorder", Type: "func(dl *DisplayList) *Recorder", Desc: "录制器"},
			{Name: "Finish", Type: "func()", Desc: "结束录制会话"},
			{Name: "Freeze/Frozen", Type: "func() / func() bool", Desc: "冻结（不可再录制）与状态查询"},
			{Name: "OffsetBy/ScaleAlpha", Type: "派生", Desc: "平移副本 / 形状 alpha 缩放（文字不缩放）"},
			{Name: "Intersects/Bounds", Type: "查询", Desc: "包围盒相交与包围盒"},
			{Name: "Version/Painted", Type: "查询", Desc: "版本号 / 有效指令数"},
		},
	})
}
