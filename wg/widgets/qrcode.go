package widgets

// ── QRCode 二维码（M12 长尾；第一百零五轮接标准编码）─────────────
//
// antd：<QRCode value size status errorLevel> 二维码图。
// **标准编码引擎已集成**（QREncode——纯 Go byte 模式 v1-10 四级别：
// RS 纠错+掩码择优，见 qrencode.go）：Value 默认走标准编码；Matrix
// 注入通道保留（v11+/其他模式仍由外部库算好传入）。视觉形态：白底
// +黑模块+三定位角+对齐图案+timing+format 信息。
// 未实现（如实标注）：v11-40、kanji/numeric 模式、status 扫描动画、
// icon 叠加。

import (
	"fmt"
	"hash/fnv"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// QRCode 是二维码组件（标准编码；Matrix 注入优先）。
type QRCode struct {
	// Value 二维码文本（antd value；Matrix 为空时经 QREncode 标准编码）。
	Value string
	// Matrix 模块矩阵（true=黑块；显式注入优先于内置编码——外部库
	// 通道，v11+/kanji 等内置引擎未覆盖的形态）。
	Matrix [][]bool
	// ErrorLevel 纠错级别（antd errorLevel；默认 QRMedium——antd 同款
	// 最常用档）。
	ErrorLevel QRErrorLevel
	// Size 显示尺寸（默认 160，antd 默认 160）。
	Size float32
	// OnClick 点击回调（antd 无；预览场景用）。
	OnClick func()
}

// NewQRCode 创建二维码。
func NewQRCode(value string) *QRCode {
	return &QRCode{Value: value}
}

// findPattern 标准定位角（7×7：黑框+白环+3×3 黑心）。
func findPattern() [][]bool {
	p := make([][]bool, 7)
	for y := 0; y < 7; y++ {
		p[y] = make([]bool, 7)
		for x := 0; x < 7; x++ {
			edge := x == 0 || x == 6 || y == 0 || y == 6
			inner := x >= 2 && x <= 4 && y >= 2 && y <= 4
			p[y][x] = edge || inner
		}
	}
	return p
}

// demoMatrix 按 Value 哈希生成确定性演示矩阵（21×21 版本 1 形态；
// **非标准 QR 编码**——只保证视觉形态与确定性）。
func demoMatrix(value string) [][]bool {
	const n = 21
	h := fnv.New32a()
	fmt.Fprintf(h, "%s", value)
	m := make([][]bool, n)
	for y := range m {
		m[y] = make([]bool, n)
		for x := 0; x < n; x++ {
			h2 := fnv.New32a()
			fmt.Fprintf(h2, "%s:%d:%d", value, x, y)
			m[y][x] = h2.Sum32()&1 == 1
		}
	}
	// 三定位角盖写 + 分隔白环。
	pat := findPattern()
	for _, origin := range [][2]int{{0, 0}, {0, n - 7}, {n - 7, 0}} {
		cx, cy := origin[0], origin[1]
		for y := 0; y < 7; y++ {
			for x := 0; x < 7; x++ {
				m[cy+y][cx+x] = pat[y][x]
			}
		}
		// 周圈白环（8 宽边界留白）。
		for y := -1; y <= 7; y++ {
			for x := -1; x <= 7; x++ {
				if y < 0 || y > 6 || x < 0 || x > 6 {
					yy, xx := cy+y, cx+x
					if yy >= 0 && yy < n && xx >= 0 && xx < n {
						m[yy][xx] = false
					}
				}
			}
		}
	}
	return m
}

// Build 实现 ui.Widget。
func (q *QRCode) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	size := q.Size
	if size <= 0 {
		size = 160
	}

	m := q.Matrix
	if m == nil {
		m = QREncode(q.Value, q.ErrorLevel)
	}
	if m == nil {
		// 超出内置引擎容量（v10 上限）：回退演示矩阵（视觉形态兜底，
		// 不再是标准编码——调用方应自行注入大容量 Matrix）。
		m = demoMatrix(q.Value)
	}
	n := len(m)
	mod := size / float32(n)

	// **CustomPaint 单节点绘制全部模块**（antd 用 canvas/svg——21×21=300+
	// Box 会节点爆炸；单节点零层级开销，§6A 极致）。
	root := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(geom.Color{R: 255, G: 255, B: 255, A: 255}))
		black := render.Fill(geom.Color{R: 0, G: 0, B: 0, A: 255})
		for y := 0; y < n; y++ {
			for x := 0; x < n; x++ {
				if m[y][x] {
					cv.FillRect(geom.RectXYWH(float32(x)*mod, float32(y)*mod, mod, mod), black)
				}
			}
		}
	}).SetSize2(size, size)
	root.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "qrcode", Cursor: scene.CursorPointer})
	if ctx.Input != nil && q.OnClick != nil {
		fn := q.OnClick
		ctx.Input.Set(root, inputHandlerTap(fn))
	}

	col := scene.Flex(scene.Vertical, th.Space.XS)
	col.SetWidth(size)
	col.Add(root)
	vl := text.LayoutText(q.Value, text.DefaultFont(th.Font.SizeXS), text.Options{})
	vt := scene.Text(scene.TextSpec{Src: q.Value, Layout: vl, Color: th.Color.TextFaint})
	vt.SetHit(scene.HitSpec{Label: "qrcode-value"})
	col.Add(vt)
	return col
}

// inputHandlerTap tap gesture wrapper (instance escapes to heap to stay alive).
func inputHandlerTap(fn func()) input.Handler {
	g := &ui.GestureDetector{OnTap: func(_ geom.Point) { fn() }}
	return g.Handler()
}
