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
	qr    = widgets.NewQRCode("https://example.com/join?token=abc123")
	qrStd = widgets.NewQRCode("https://example.com/join?token=abc123")
	qrLow = widgets.NewQRCode("https://example.com/join?token=abc123")
)

func init() {
	// 默认卡已是标准编码（QREncode 默认 M 级）；纠错级别对比卡：
	// 同一 Value 不同级别产出不同模块分布（EC 码字越多数据占比越低）。
	qrStd.ErrorLevel = widgets.QRHigh
	qrLow.ErrorLevel = widgets.QRLow
}

func main() {
	launcher.Main(doc.Meta{
		ID: "qrcode", Title: "QRCode 二维码", Group: doc.GroupData,
		What: []string{
			"**标准编码引擎**（QREncode 纯 Go）：byte 模式 v1-10 四级别——Reed-Solomon GF(256) 纠错+8 掩码 penalty 择优+结构交织",
			"白底+黑模块+三定位角+对齐图案+timing+format 信息（ISO 18004 布置）",
			"ErrorLevel 纠错级别（L/M/Q/H，默认 M——antd 最常用档）；Matrix 注入通道保留（v11+/其他模式）",
			"Value 下方灰字展示（antd status 提示形态）",
		},
		Demos: []doc.Demo{
			{Title: "标准编码", Note: "QREncode 默认 M 级：https://example.com/join?token=abc123（v2，36 字节容量）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return qr
			}},
			{Title: "纠错级别对比", Note: "同 Value 不同 ErrorLevel：H（左，30% 纠错）与 L（右，7%）模块分布不同", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					row := scene.Flex(scene.Horizontal, 24)
					row.Add(qrStd.Build(c))
					row.Add(qrLow.Build(c))
					return row
				})
			}},
		},
		API: []doc.APIRow{
			{Name: "NewQRCode", Type: "func(value string) *QRCode", Desc: "构造"},
			{Name: "Value", Type: "string", Desc: "二维码文本（下方展示+编码源）"},
			{Name: "ErrorLevel", Type: "QRErrorLevel", Default: "QRMedium", Desc: "纠错级别 L/M/Q/H（byte 模式容量 v1-10）"},
			{Name: "Matrix", Type: "[][]bool", Desc: "模块矩阵注入（外部库通道，v11+/其他模式；nil=内置 QREncode）"},
			{Name: "Size", Type: "float32", Default: "160", Desc: "显示尺寸"},
		},
	})
}
