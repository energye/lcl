package widgets

import (
	"fmt"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
)

// ── Progress（M8 对齐 antd 线形部分）────────────────────────────────
//
// antd：<Progress percent status="success|exception" showInfo>
// 已对齐：线形、status 语义色；未实现（M10）：circle/dashboard、trail 色。

// ProgressStatus 是进度条语义（antd Progress.status）。
type ProgressStatus uint8

const (
	// ProgressNormal 主色。
	ProgressNormal ProgressStatus = iota
	// ProgressSuccess 成功色（percent=100 时 antd 自动切绿）。
	ProgressSuccess
	// ProgressException 异常红。
	ProgressException
)

// Progress 是线形进度条。Value 取值 0…1，越界自动夹住。
type Progress struct {
	Value  float32
	Width  float32
	Height float32
	Status ProgressStatus
	// ShowInfo 在右侧显示百分比文字（antd showInfo）。
	ShowInfo bool
}

// Build 实现 ui.Widget。
func (p Progress) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	w, h := p.Width, p.Height
	if w <= 0 {
		w = 160
	}
	if h <= 0 {
		h = 6
	}
	// trail 用 SurfaceActive（#F0F0F0，antd 线形进度条默认灰轨）——
	// SurfaceSunken 与演示舞台同色，底轨会隐形。
	track := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceActive), Radius: h / 2}).SetSize2(w, h)
	v := p.Value
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	fillColor := th.Color.Primary
	if v >= 1 && p.Status == ProgressNormal {
		fillColor = th.Color.Success // antd：满进度自动转 success
	}
	switch p.Status {
	case ProgressSuccess:
		fillColor = th.Color.Success
	case ProgressException:
		fillColor = th.Color.Danger
	}
	if v > 0 {
		fill := scene.Box(scene.BoxSpec{Fill: scene.Solid(fillColor), Radius: h / 2}).
			SetSize2(w*v, h)
		track.Add(fill)
	}
	if !p.ShowInfo {
		return track
	}
	row := scene.Flex(scene.Horizontal, 8).SetCrossAlign(scene.CrossCenter)
	row.Add(track)
	row.Add((&Label{Text: pctText(v), Size: th.Font.SizeSM, Color: fillColor}).Build(ctx))
	return row
}

func pctText(v float32) string {
	switch {
	case v >= 1:
		return "100%"
	case v <= 0:
		return "0%"
	}
	return fmtPct(v)
}

// fmtPct 是百分比文本（整数化）。
func fmtPct(v float32) string {
	return fmt.Sprintf("%d%%", int(v*100+0.5))
}
