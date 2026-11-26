package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var (
	wmContent = scene.Flex(scene.Vertical, 8)
	wm        *widgets.Watermark
	wmMulti   *widgets.Watermark
)

func init() {
	wmContent.SetWidth(480)
	for _, ln := range []string{
		"季度经营摘要（内部资料）",
		"本季度营收同比增长 24%，毛利率稳定在 38%。",
		"重点客户续约率 92%，新增客户 37 家。",
		"下季度重点：渠道下沉与产品矩阵扩展。",
	} {
		lb := &widgets.Label{Text: ln}
		wmContent.Add(lb.Build(nil))
	}
	wm = widgets.NewWatermark(
		doc.Node(func(c *ui.BuildContext) *scene.Node { return wmContent }),
		"内部资料 · 请勿外传",
	)
	wm.Width = 480
	wm.Height = 240

	// 多行水印卡（antd content 数组）：三行纵向堆叠。
	wmMulti = widgets.NewWatermark(
		doc.Node(func(c *ui.BuildContext) *scene.Node { return wmContent }),
		"",
	)
	wmMulti.Contents = []string{"机密文件", "2026 年度", "请勿外传"}
	wmMulti.Width = 480
	wmMulti.Height = 240
}

func main() {
	launcher.Main(doc.Meta{
		ID: "watermark", Title: "Watermark 水印", Group: doc.GroupData,
		What: []string{
			"Stack 底层 Child + **水印层 CustomPaint 单节点**（Save→Rotate→网格 DrawText→Restore——canvas 原生旋转 antd 默认 -22°）",
			"alpha 语义：antd rgba(0,0,0,0.15) = geom.RGBAa(0,0,0,38)；覆盖范围按对角线扩展（旋转后盖满）",
			"**多行水印**（Contents 数组逐行纵向堆叠、中心对齐）；未实现：图片水印/防篡改/zIndex",
		},
		Demos: []doc.Demo{
			{Title: "多行水印", Note: "Contents 数组：三行纵向堆叠（机密文件 / 2026 年度 / 请勿外传），antd content 数组", Build: func(ctx *ui.BuildContext) ui.Widget {
				return wmMulti
			}},
			{Title: "文本平铺水印", Note: "-22° 旋转灰字平铺（\"内部资料 · 请勿外传\"）盖在经营摘要上", Build: func(ctx *ui.BuildContext) ui.Widget {
				return wm
			}},
		},
		API: []doc.APIRow{
			{Name: "NewWatermark", Type: "func(child ui.Widget, content string) *Watermark", Desc: "构造（-22° 默认旋转）"},
			{Name: "Child", Type: "ui.Widget", Desc: "被覆盖内容"},
			{Name: "Content", Type: "string", Desc: "水印文本（单行）"},
			{Name: "Contents", Type: "[]string", Desc: "多行水印文本（antd content 数组）；非空优先于 Content，逐行堆叠"},
			{Name: "GapX / GapY", Type: "float32", Default: "100 / 80", Desc: "平铺间距"},
			{Name: "FontSize / RotateDeg", Type: "float32", Default: "14 / -22", Desc: "字号/旋转角"},
			{Name: "Width / Height", Type: "float32", Desc: "显式尺寸（Stack 必需）"},
		},
	})
}
