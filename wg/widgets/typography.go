package widgets

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ── Typography（M8 对齐 antd）───────────────────────────────────────
//
// antd：<Typography.Title level={1..5}> / <Paragraph> / <Text
//
//	type mark code strong underline delete disabled>
//
// 视觉对齐：标题 38/30/24/20/16 加粗 600；正文 14/行高 1.57；
// mark 黄底 #FEFCE6 近似 antd #FFE58F 淡化；code 等宽灰底。
// 未实现（如实标注）：copyable（无剪贴板通道）、editable、ellipsis 展开。

// typoWeight 是 antd 标题字重（600）。
const typoWeight = 600

// Title 是标题（antd Typography.Title，level 1..5）。
type Title struct {
	Text  string
	Level int // 1..5；超出范围按 1 处理
}

// Build 实现 ui.Widget。
func (t Title) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	sizes := []float32{38, 30, 24, 20, 16}
	lvl := t.Level
	if lvl < 1 || lvl > 5 {
		lvl = 1
	}
	size := sizes[lvl-1]
	font := text.DefaultFont(size).WithWeight(text.Weight(typoWeight))
	// 标题自带上下间距（antd h* margin-bottom 0.5em）。
	n := scene.Text(scene.TextSpec{
		Src: t.Text, Layout: text.LayoutText(t.Text, font, text.Options{LineHeight: 1.3}), Color: th.Color.Text,
	})
	n.SetMargin(geom.InLTRB(0, 0, 0, size*0.4))
	n.SetHit(scene.HitSpec{Label: "typography"}) // 定位锚（DeferToChild 不影响交互）
	return n
}

// Paragraph 是段落（antd Typography.Paragraph）。
type Paragraph struct {
	Text     string
	MaxWidth float32
	Color    geom.Color // 零值取 Text
}

// Build 实现 ui.Widget。
func (p Paragraph) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	c := p.Color
	if c == (geom.Color{}) {
		c = th.Color.Text
	}
	font := text.DefaultFont(th.Font.SizeBase)
	n := scene.Text(scene.TextSpec{
		Src: p.Text, Layout: text.LayoutText(p.Text, font, text.Options{
			MaxWidth: p.MaxWidth, LineHeight: 1.57,
		}), Color: c,
	})
	n.SetHit(scene.HitSpec{Label: "typography"})
	return n
}

// TextType 是 Text 的语义类型（antd Typography.Text.type）。
type TextType uint8

const (
	// TypePlain 默认色。
	TypePlain TextType = iota
	// TypeSecondary 次要色（rgba(0,0,0,.65) 等价）。
	TypeSecondary
	// TypeSuccess 成功色。
	TypeSuccess
	// TypeWarning 警告色。
	TypeWarning
	// TypeDanger 危险色。
	TypeDanger
)

// Text 是内联文本（antd Typography.Text）。
type Text struct {
	Body      string
	Type      TextType
	Strong    bool // 加粗
	Disabled  bool // 禁用色
	Mark      bool // 黄底
	Code      bool // 等宽灰底
	Underline bool // 下划线
	Delete    bool // 删除线
}

// Build 实现 ui.Widget。
func (t Text) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	color := t.color(th)
	font := text.DefaultFont(th.Font.SizeBase)
	if t.Strong {
		font = font.WithWeight(text.Bold)
	}
	if t.Code {
		font = font.WithFamily("Consolas")
	}
	layout := text.LayoutText(t.Body, font, text.Options{LineHeight: 1.57})

	var inner *scene.Node
	if t.Underline || t.Delete {
		// 下划线/删除线：CustomPaint 画字后再补线（TextSpec 无装饰通道）。
		inner = scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
			cv.DrawText(layout, geom.Pt(0, 0), color)
			y := layout.Size.H - 3
			if t.Underline {
				cv.DrawLine(0, y, layout.Size.W, y, render.Stroke(color, 1))
			}
			if t.Delete {
				cv.DrawLine(0, layout.Size.H/2, layout.Size.W, layout.Size.H/2, render.Stroke(color, 1))
			}
		}).SetSize2(layout.Size.W, layout.Size.H)
	} else {
		inner = scene.Text(scene.TextSpec{Src: t.Body, Layout: layout, Color: color})
	}
	inner.SetHit(scene.HitSpec{Label: "typography"})

	switch {
	case t.Mark:
		box := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(255, 229, 143))}, inner)
		box.SetPadding(geom.InLTRB(2, 0, 2, 0))
		return box
	case t.Code:
		box := scene.Box(scene.BoxSpec{
			Fill:   scene.Solid(th.Color.SurfaceSunken),
			Border: scene.Stroke(th.Color.Border, 1),
			Radius: th.Radius.SM,
		}, inner)
		box.SetPadding(geom.InLTRB(4, 1, 4, 1))
		return box
	}
	return inner
}

// color 折算语义色。
func (t Text) color(th *theme.Data) geom.Color {
	c := th.Color
	if t.Disabled {
		return c.BorderStrong // colorTextDisabled 实色等价
	}
	switch t.Type {
	case TypeSecondary:
		return c.TextMuted
	case TypeSuccess:
		return c.Success
	case TypeWarning:
		return c.Warning
	case TypeDanger:
		return c.Danger
	}
	return c.Text
}
