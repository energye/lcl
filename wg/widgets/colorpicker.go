package widgets

// ── ColorPicker 颜色选择（M12 高级录入）──────────────────────────
//
// antd：<ColorPicker value onChange presets> 默认触发器：
// 圆角色块 + HEX 文本；面板：SV 饱和度/亮度渐变面板 + 色相条 +
// 预设色板 + HEX 显示。
// 渐变直接用 render.NewLinearGradient（BoxSpec.Fill 持 Shader——
// scene.Solid 即 render.Fill，同族）；点击取色经事件节点坐标换算
// （x→s、y→1-v；hue 条 x→h·360）。down 即执行（浮层结论复用）。
//
// 未实现（如实标注）：透明度通道/alpha 条、RGB/HSB 输入切换、
// format 定制、disabledAlpha、自定义触发器/panelRender、
// 拖拽连续取色（仅点击单点取色）、child 触发器。

import (
	"fmt"
	"math"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// ColorPicker 是颜色选择器（HSV 状态机）。
type ColorPicker struct {
	controlBase
	Value    string // 当前 hex（如 #1677FF）
	Defaults string // 无值时初始色（antd defaultValue）
	Presets  []string
	OnChange func(hex string)
	Width    float32 // 触发框宽（默认 0=仅色块形态 42）

	handle    overlay.Handle
	fb        feedbackHost
	childNode *scene.Node
	panelOpen bool

	h, s, v float32 // HSV：h 0-360、s/v 0-1
	curTh   *theme.Data
}

// NewColorPicker 创建颜色选择器（默认 #1677FF）。
func NewColorPicker() *ColorPicker {
	return &ColorPicker{Value: "#1677FF"}
}

// hsvToHex HSV(0-360,0-1,0-1) → #RRGGBB。
func hsvToHex(h, s, v float32) string {
	c := v * s
	x := c * (1 - float32(math.Abs(float64(math.Mod(float64(h/60), 2)-1))))
	m := v - c
	var r, g, b float32
	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}
	return fmt.Sprintf("#%02X%02X%02X",
		uint8((r+m)*255+0.5), uint8((g+m)*255+0.5), uint8((b+m)*255+0.5))
}

// hueColor hue(0-360) 的纯色（s=v=1）。
func hueColor(h float32) geom.Color {
	hex := hsvToHex(h, 1, 1)
	r, g, b := parseHex(hex)
	return geom.RGBA(r, g, b)
}

// parseHex #RRGGBB → (r,g,b)；失败返回黑。
func parseHex(hex string) (uint8, uint8, uint8) {
	if len(hex) != 7 || hex[0] != '#' {
		return 0, 0, 0
	}
	v := uint32(0)
	for i := 1; i < 7; i++ {
		c := hex[i]
		var d uint32
		switch {
		case c >= '0' && c <= '9':
			d = uint32(c - '0')
		case c >= 'A' && c <= 'F':
			d = uint32(c-'A') + 10
		case c >= 'a' && c <= 'f':
			d = uint32(c-'a') + 10
		default:
			return 0, 0, 0
		}
		v = v<<4 | d
	}
	return uint8(v >> 16), uint8(v >> 8), uint8(v)
}

// Build 实现 ui.Widget：触发框（圆角色块+HEX 文本）。
func (c *ColorPicker) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(c)
	th := ctx.ThemeOr()
	c.curTh = th
	w := c.Width
	if w <= 0 {
		w = 42
	}

	// 无值时用 Defaults 初始化 HSV。
	if c.Value == "" && c.Defaults != "" {
		c.Value = c.Defaults
	}
	if c.Value == "" {
		c.Value = "#1677FF"
	}
	hr, hg, hb := parseHex(c.Value)
	blockColor := geom.RGBA(hr, hg, hb)

	row := scene.Flex(scene.Horizontal, th.Space.XS).SetCrossAlign(scene.CrossCenter)
	block := scene.Box(scene.BoxSpec{
		Fill:   scene.Solid(blockColor),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.SM,
	}).SetSize2(24, 24)
	block.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "color-swatch", Cursor: scene.CursorPointer})
	row.Add(block)
	hex := c.Value
	hl := text.LayoutText(hex, text.DefaultFont(th.Font.SizeBase), text.Options{})
	hexText := scene.Text(scene.TextSpec{Src: hex, Layout: hl, Color: th.Color.Text})
	hexText.SetHit(scene.HitSpec{Label: "color-hex"})
	row.Add(hexText)

	row.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "color-picker", Cursor: scene.CursorPointer})
	if ctx.Input != nil {
		anchor := row
		g := &ui.GestureDetector{OnTap: func(geom.Point) {
			if c.panelOpen {
				c.closePanel(ctx)
				return
			}
			c.openPanel(ctx, anchor, th)
		}}
		ctx.Input.Set(row, g.Handler())
	}
	c.childNode = row
	return row
}

// Show 程序化打开面板（测试/自动化用）。
func (c *ColorPicker) Show(ctx *ui.BuildContext) {
	if c.childNode != nil {
		c.openPanel(ctx, c.childNode, ctx.ThemeOr())
	}
}

// IsOpen 返回面板是否打开。
func (c *ColorPicker) IsOpen() bool { return c.panelOpen }

// Close 关闭面板。
func (c *ColorPicker) Close() {
	if c.fb.host != nil {
		c.closePanel(nil)
	}
}

// openPanel 打开取色面板。
func (c *ColorPicker) openPanel(ctx *ui.BuildContext, anchor *scene.Node, th *theme.Data) {
	if c.handle != 0 {
		return
	}
	c.curTh = th
	c.panelOpen = true
	if c.h == 0 && c.s == 0 && c.v == 0 {
		r, g, b := parseHex(c.Value)
		c.setFromRGB(r, g, b)
	}
	panel := scene.Flex(scene.Vertical, th.Space.SM)
	panel.SetBox(scene.BoxSpec{
		Fill:   scene.Solid(th.Color.Surface),
		Border: scene.Stroke(th.Color.Border, 1),
		Radius: th.Radius.Base,
	})
	panel.SetPadding(geom.InLTRB(th.Space.SM, th.Space.SM, th.Space.SM, th.Space.SM))
	panel.SetWidth(236)
	panel.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "color-panel"})

	// SV 面板：底层 hue 纯色 + 白横渐变 + 黑纵渐变（antd 双层叠加）。
	hueBg := hueColor(c.h)
	sv := scene.Stack()
	sv.SetSize2(220, 160)
	sv.SetBox(scene.BoxSpec{Fill: scene.Solid(hueBg)})
	sv.Add(scene.Box(scene.BoxSpec{
		Fill: render.FillWith(render.NewLinearGradient(
			geom.Pt(0, 0), geom.Pt(220, 0),
			render.Stop{Offset: 0, Color: geom.RGBAa(255, 255, 255, 255)},
			render.Stop{Offset: 1, Color: geom.RGBAa(255, 255, 255, 0)},
		)),
	}).SetSize2(220, 160))
	sv.Add(scene.Box(scene.BoxSpec{
		Fill: render.FillWith(render.NewLinearGradient(
			geom.Pt(0, 0), geom.Pt(0, 160),
			render.Stop{Offset: 0, Color: geom.RGBAa(0, 0, 0, 0)},
			render.Stop{Offset: 1, Color: geom.RGBAa(0, 0, 0, 255)},
		)),
	}).SetSize2(220, 160))
	sv.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "color-sv", Cursor: scene.CursorPointer})
	if ctx.Input != nil {
		ctx.Input.Set(sv, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
			if !ev.Down() {
				return false
			}
			// down 即执行；Pos 已是节点空间：x→s、y→1-v。
			s := clamp01(ev.Pos.X / 220)
			v := clamp01(1 - ev.Pos.Y/160)
			c.s, c.v = s, v
			c.commit(ctx)
			return true
		}})
	}
	panel.Add(sv)

	// 色相条：红→黄→绿→青→蓝→紫→红 渐变。
	hue := scene.Box(scene.BoxSpec{
		Fill: render.FillWith(render.NewLinearGradient(
			geom.Pt(0, 0), geom.Pt(220, 0),
			render.Stop{Offset: 0, Color: geom.RGBA(255, 0, 0)},
			render.Stop{Offset: 1.0 / 6, Color: geom.RGBA(255, 255, 0)},
			render.Stop{Offset: 2.0 / 6, Color: geom.RGBA(0, 255, 0)},
			render.Stop{Offset: 3.0 / 6, Color: geom.RGBA(0, 255, 255)},
			render.Stop{Offset: 4.0 / 6, Color: geom.RGBA(0, 0, 255)},
			render.Stop{Offset: 5.0 / 6, Color: geom.RGBA(255, 0, 255)},
			render.Stop{Offset: 1, Color: geom.RGBA(255, 0, 0)},
		)),
		Radius: th.Radius.SM,
	}).SetSize2(220, 14)
	hue.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "color-hue", Cursor: scene.CursorPointer})
	if ctx.Input != nil {
		ctx.Input.Set(hue, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
			if !ev.Down() {
				return false
			}
			c.h = clamp01(ev.Pos.X/220) * 360
			c.commit(ctx)
			return true
		}})
	}
	panel.Add(hue)

	// 预设色板（网格）。
	if len(c.Presets) > 0 {
		grid := scene.Flex(scene.Vertical, 4)
		const perRow = 10
		for i := 0; i < len(c.Presets); i += perRow {
			grow := scene.Flex(scene.Horizontal, 4)
			for j := i; j < i+perRow && j < len(c.Presets); j++ {
				hexP := c.Presets[j]
				pr, pg, pb := parseHex(hexP)
				swatch := scene.Box(scene.BoxSpec{
					Fill:   scene.Solid(geom.RGBA(pr, pg, pb)),
					Radius: th.Radius.SM,
				}).SetSize2(20, 20)
				swatch.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "color-preset", Cursor: scene.CursorPointer})
				if ctx.Input != nil {
					ctx.Input.Set(swatch, input.Handler{Pointer: func(ev *input.PointerEvent) bool {
						if !ev.Down() {
							return false
						}
						r2, g2, b2 := parseHex(hexP)
						c.setFromRGB(r2, g2, b2)
						c.commit(ctx)
						return true
					}})
				}
				grow.Add(swatch)
			}
			grid.Add(grow)
		}
		panel.Add(grid)
	}

	c.handle = ctx.Overlay.Show(overlay.Entry{
		Band:      theme.BandPopup,
		Anchor:    overlay.NodeAnchor(anchor),
		Placement: overlay.PlacementBottomStart,
		Size:      geom.Size{W: 236, H: 0}, // 高自适应
		Dismiss:   overlay.DismissOutsideClick | overlay.DismissEsc,
		Tag:       "colorpicker",
		OnClose: func() {
			c.handle = 0
			c.panelOpen = false
		},
		Build: func() *scene.Node { return panel },
	})
}

// setFromRGB RGB → HSV。
func (c *ColorPicker) setFromRGB(r, g, b uint8) {
	rf, gf, bf := float32(r)/255, float32(g)/255, float32(b)/255
	max, min := rf, rf
	if gf > max {
		max = gf
	}
	if bf > max {
		max = bf
	}
	if gf < min {
		min = gf
	}
	if bf < min {
		min = bf
	}
	if max == min {
		c.h, c.s, c.v = 0, 0, max
		return
	}
	d := max - min
	c.s = d / max
	c.v = max
	var h float32
	switch max {
	case rf:
		h = 60 * float32(math.Mod(float64((gf-bf)/d), 6))
	case gf:
		h = 60 * ((bf-rf)/d + 2)
	default:
		h = 60 * ((rf-gf)/d + 4)
	}
	if h < 0 {
		h += 360
	}
	c.h = h
}

// commit HSV → hex：更新 Value、回调。**不关面板**（antd 连续调色
// 语义——SV/色相/预设点击后面板保持，外点/Esc/再点触发框关闭）。
func (c *ColorPicker) commit(ctx *ui.BuildContext) {
	c.Value = hsvToHex(c.h, c.s, c.v)
	c.MarkDirty() // 触发框色块/HEX 随根重建刷新
	if c.OnChange != nil {
		c.OnChange(c.Value)
	}
}

func clamp01(v float32) float32 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

// closePanel 关闭面板。
func (c *ColorPicker) closePanel(ctx *ui.BuildContext) {
	if c.handle != 0 && c.fb.host != nil {
		c.fb.host.Close(c.handle)
	}
	c.handle = 0
	c.panelOpen = false
}
