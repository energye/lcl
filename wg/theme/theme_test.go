package theme

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
)

// 黄金值测试：**外观对齐的真相源（替代官方截图对照）**。
//
// antd v5 默认主题的关键颜色/圆角/字号/控件高值逐值锁定——截图对照
// 成本高且不可执行验证，token 值断言即改即测：任何 token 改动导致
// 与官方值或既定设计漂移，本测试先挂。规格来源=antd v5 默认算法
// 的输出值（light 注释标注官方十六进制；alpha 通道保留原始值）。
//
// 注意两类断言对象：
//  1. 纯色 token（Solid 语义）：断言 RGB 与官方值相等；
//  2. 半透明 token（FillHover/FillActive/Overlay/Shadow）：组件渲染
//     时与白底（light）或黑底（dark）混合，此处断言**原始 RGBA**，
//     混合等价色以注释标注。
//
// Space 阶梯是 wg 自身的紧凑设计（非官方 8/12/16/24/32），黄金值
// 测试冻结现状防回归，注释如实标注差异。

func c(col geom.Color) [3]int {
	return [3]int{int(col.R), int(col.G), int(col.B)}
}

func want(t *testing.T, got, want [3]int, name string) {
	t.Helper()
	d := 0
	for i := 0; i < 3; i++ {
		d += (got[i] - want[i]) * (got[i] - want[i])
	}
	if d != 0 {
		t.Errorf("%s = %v，与官方值 %v 漂移（分量差平方和 %d）", name, got, want, d)
	}
}

func wantAlpha(t *testing.T, col geom.Color, a uint8, name string) {
	t.Helper()
	if col.A != a {
		t.Errorf("%s alpha = %v，期望 %v", name, col.A, a)
	}
}

// TestDefaultTokenColors：light 色板与 antd 默认主题逐值相等。
func TestDefaultTokenColors(t *testing.T) {
	th := Default()
	col := th.Color
	// antd 官方值：primary #1677FF / hover #4096FF / active #0958D9。
	want(t, c(col.Primary), [3]int{22, 119, 255}, "Primary #1677FF")
	want(t, c(col.PrimaryHover), [3]int{64, 150, 255}, "PrimaryHover #4096FF")
	want(t, c(col.PrimaryActive), [3]int{9, 88, 217}, "PrimaryActive #0958D9")
	want(t, c(col.OnPrimary), [3]int{255, 255, 255}, "OnPrimary #FFFFFF")

	// 背景与中性阶：白底/容器 #FFFFFF、fill #F0F0F0、填充三级 #FAFAFA。
	want(t, c(col.Surface), [3]int{255, 255, 255}, "Surface #FFFFFF")
	want(t, c(col.SurfaceActive), [3]int{240, 240, 240}, "SurfaceActive #F0F0F0")
	want(t, c(col.SurfaceSunken), [3]int{250, 250, 250}, "SurfaceSunken #FAFAFA")
	want(t, c(col.SurfaceDisabled), [3]int{245, 245, 245}, "SurfaceDisabled #F5F5F5")
	want(t, c(col.SurfaceHover), [3]int{250, 250, 250}, "SurfaceHover #FAFAFA")

	// 填充与描边：rgba(0,0,0,.04)/.15 原值 + #D9D9D9/#BFBFBF。
	want(t, c(col.FillHover), [3]int{0, 0, 0}, "FillHover 原始 RGB（白底等价 #F5F5F5）")
	wantAlpha(t, col.FillHover, 10, "FillHover alpha（=rgba(0,0,0,.04)）")
	want(t, c(col.FillActive), [3]int{0, 0, 0}, "FillActive 原始 RGB")
	wantAlpha(t, col.FillActive, 38, "FillActive alpha（=rgba(0,0,0,.15)）")
	want(t, c(col.Border), [3]int{217, 217, 217}, "Border #D9D9D9")
	want(t, c(col.BorderStrong), [3]int{191, 191, 191}, "BorderStrong #BFBFBF")

	// 文字：colorText(.88)/secondary(.65)/tertiary(.45)/placeholder 实色等价。
	want(t, c(col.Text), [3]int{31, 31, 31}, "Text 实色等价 #1F1F1F（rgba(0,0,0,.88) 白底）")
	want(t, c(col.TextMuted), [3]int{89, 89, 89}, "TextMuted #595959")
	want(t, c(col.TextFaint), [3]int{140, 140, 140}, "TextFaint #8C8C8C")
	want(t, c(col.TextPlaceholder), [3]int{191, 191, 191}, "TextPlaceholder #BFBFBF")

	// 语义色：#52C41A / #FAAD14 / #FF4D4F 系（含 hover/active/底色）。
	want(t, c(col.Success), [3]int{82, 196, 26}, "Success #52C41A")
	want(t, c(col.Warning), [3]int{250, 173, 20}, "Warning #FAAD14")
	want(t, c(col.WarningHover), [3]int{255, 197, 61}, "WarningHover #FFC53D")
	want(t, c(col.Danger), [3]int{255, 77, 79}, "Danger #FF4D4F")
	want(t, c(col.DangerHover), [3]int{255, 120, 117}, "DangerHover #FF7875")
	want(t, c(col.DangerActive), [3]int{217, 54, 62}, "DangerActive #D9363E")
	want(t, c(col.DangerBg), [3]int{255, 241, 240}, "DangerBg #FFF1F0")
	want(t, c(col.OnDanger), [3]int{255, 255, 255}, "OnDanger #FFFFFF")
	want(t, c(col.Selection), [3]int{186, 224, 255}, "Selection #BAE0FF（primary 淡底）")
	want(t, c(col.Info), [3]int{22, 119, 255}, "Info = Primary")
	want(t, c(col.FocusRing), [3]int{22, 119, 255}, "FocusRing = Primary")

	// 浮层遮罩 rgba(0,0,0,.45) 原值。
	want(t, c(col.Overlay), [3]int{0, 0, 0}, "Overlay 原始 RGB（黑底等价 rgba(0,0,0,.45)）")
	wantAlpha(t, col.Overlay, 115, "Overlay alpha（≈.45）")
}

// TestDefaultTokenRadiusFont：圆角与字号与官方值逐值相等。
func TestDefaultTokenRadiusFont(t *testing.T) {
	th := Default()
	// antd borderRadius：SM 4 / Base 6 / LG 8 / pill 999。
	if th.Radius.SM != 4 || th.Radius.Base != 6 || th.Radius.LG != 8 || th.Radius.Pill != 999 {
		t.Errorf("Radius = SM%v Base%v LG%v Pill%v，与官方 4/6/8/999 不符",
			th.Radius.SM, th.Radius.Base, th.Radius.LG, th.Radius.Pill)
	}
	// antd 字号：xs 11 / sm 12 / base 14 / lg 16 / xl 20 / title 24。
	ft := th.Font
	if ft.SizeXS != 11 || ft.SizeSM != 12 || ft.SizeBase != 14 || ft.SizeLG != 16 || ft.SizeXL != 20 || ft.SizeTitle != 24 {
		t.Errorf("Font = %+v，与官方字号 11/12/14/16/20/24 不符", ft)
	}
	if ft.LineHeight != 1.5 {
		t.Errorf("LineHeight = %v，期望 1.5", ft.LineHeight)
	}
}

// TestDefaultTokenSpace：间距阶梯冻结 wg 紧凑设计（非官方 8/12/16/24/32），
// 防回归漂移；如实标注与官方的差异。
func TestDefaultTokenSpace(t *testing.T) {
	th := Default()
	sp := th.Space
	wantN(t, sp.XS, 2, "Space.XS（wg 紧凑阶梯，官方 xs=8）")
	wantN(t, sp.SM, 4, "Space.SM（wg 紧凑阶梯，官方 sm=12）")
	wantN(t, sp.MD, 8, "Space.MD（wg 紧凑阶梯，官方 md=16）")
	wantN(t, sp.LG, 12, "Space.LG（wg 紧凑阶梯，官方 lg=24）")
	wantN(t, sp.XL, 16, "Space.XL（wg 紧凑阶梯，官方 xl=32）")
	wantN(t, sp.XXL, 24, "Space.XXL（wg 紧凑阶梯）")
}

func wantN(t *testing.T, got, want float32, name string) {
	t.Helper()
	if got != want {
		t.Errorf("%s = %v，期望 %v", name, got, want)
	}
}

// TestDarkTokenColors：暗色主题同键换值（dark 算法基底），关键色锁定。
func TestDarkTokenColors(t *testing.T) {
	th := Dark()
	col := th.Color
	// dark 算法：primary #1668DC（hover #3C89E8 / active #1554AD）。
	want(t, c(col.Primary), [3]int{22, 104, 220}, "Dark Primary #1668DC")
	want(t, c(col.PrimaryHover), [3]int{60, 137, 232}, "Dark PrimaryHover #3C89E8")
	// dark 算法：#49AA19 / #D89614 / #DC4446 / #1668DC。
	want(t, c(col.Success), [3]int{73, 170, 25}, "Dark Success #49AA19")
	want(t, c(col.Warning), [3]int{216, 150, 20}, "Dark Warning #D89614")
	want(t, c(col.Danger), [3]int{220, 68, 70}, "Dark Danger #DC4446")
	// 背景基底 #141414；文字 rgba(255,255,255,.85) 白底混合等价。
	want(t, c(col.Surface), [3]int{20, 20, 20}, "Dark Surface #141414")
	want(t, c(col.Text), [3]int{219, 219, 219}, "Dark Text 实色等价 #DBDBDB")
	// dark 遮罩 rgba(0,0,0,.65) 原值。
	want(t, c(col.Overlay), [3]int{0, 0, 0}, "Dark Overlay 原始 RGB")
	wantAlpha(t, col.Overlay, 166, "Dark Overlay alpha（≈.65）")
}

// TestDarkSameRadiusFontSpace：几何 token 双主题一致（同键同值）。
func TestDarkSameRadiusFontSpace(t *testing.T) {
	l, d := Light(), Dark()
	if d.Radius != l.Radius {
		t.Errorf("dark/light Radius 不一致：%+v vs %+v", d.Radius, l.Radius)
	}
	if d.Space != l.Space {
		t.Errorf("dark/light Space 不一致：%+v vs %+v", d.Space, l.Space)
	}
	if d.Font != l.Font {
		t.Errorf("dark/light Font 不一致：%+v vs %+v", d.Font, l.Font)
	}
}
