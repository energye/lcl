package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var qr = widgets.NewQRCode("https://example.com/join?token=abc123")

func main() {
	launcher.Main(doc.Meta{
		ID: "qrcode", Title: "QRCode 二维码", Group: doc.GroupData,
		What: []string{
			"矩阵渲染：白底+黑模块+**三定位角**（finder pattern+白环分隔）",
			"**编码引擎未集成**（QR 需 Reed-Solomon+掩码，标准库无实现）——Matrix [][]bool 供外部库注入；未注入时按 Value 哈希生成**确定性演示矩阵**（非标准编码，如实标注）",
			"Value 下方灰字展示（antd status 提示形态）",
		},
		Demos: []doc.Demo{
			{Title: "演示矩阵", Note: "Value 哈希演示矩阵（三定位角正确）——标准编码待外部库", Build: func(ctx *ui.BuildContext) ui.Widget {
				return qr
			}},
		},
		API: []doc.APIRow{
			{Name: "NewQRCode", Type: "func(value string) *QRCode", Desc: "构造"},
			{Name: "Value", Type: "string", Desc: "二维码文本（下方展示+演示矩阵哈希源）"},
			{Name: "Matrix", Type: "[][]bool", Desc: "模块矩阵注入（外部库标准编码入口；nil=演示矩阵）"},
			{Name: "Size", Type: "float32", Default: "160", Desc: "显示尺寸"},
		},
	})
}
