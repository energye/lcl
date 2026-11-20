package widgets

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
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

// ProgressShape 是进度条形态（antd Progress.type）。
type ProgressShape uint8

const (
	// ProgressLine 线形（默认）。
	ProgressLine ProgressShape = iota
	// ProgressCircle 圆环（antd type="circle"，默认直径 120）。
	ProgressCircle
	// ProgressDashboard 仪表盘（antd type="dashboard"，缺口朝下 270°）。
	ProgressDashboard
)

// Progress 是进度条。Value 取值 0…1，越界自动夹住。
// Shape 决定线形/圆环/仪表盘三种形态（circle/dashboard 走
// CustomPaintNode 弧线绘制，与 Spin 同通道）。
type Progress struct {
	Value  float32
	Width  float32
	Height float32
	Status ProgressStatus
	Shape  ProgressShape
	// StrokeWidth 圆环描边宽（antd circle strokeWidth 默认 6）。
	StrokeWidth float32
	// ShowInfo 在右侧（线形）/中心（圆环）显示百分比文字（antd showInfo）。
	ShowInfo bool
}

// Build 实现 ui.Widget。
func (p Progress) Build(ctx *ui.BuildContext) *scene.Node {
	if p.Shape == ProgressCircle || p.Shape == ProgressDashboard {
		return p.buildRing(ctx)
	}
	return p.buildLine(ctx)
}

// ringGeometry 折算圆环几何：直径/描边宽/扫掠角（dashboard 缺口朝下
// 270°；circle 整圈 360°）。
func (p Progress) ringGeometry() (size, stroke, sweepDeg float32) {
	size = p.Width
	if size <= 0 {
		size = 120
	}
	stroke = p.StrokeWidth
	if stroke <= 0 {
		stroke = 6
	}
	if p.Shape == ProgressDashboard {
		return size, stroke, 270
	}
	return size, stroke, 360
}

// buildRing 构建圆环/仪表盘进度：trail 灰环 + 进度弧（CustomPaintNode
// 录制期按 Value 绘制——重建/断言读同一真相）+ 中心百分比文字。
// AddArc 第二参是扫掠角（Spin/Button loading 轮已验证），顺时针从
// 12 点方向起（AddRect 0° 起点+render 的角度约定，环形起点由 AddArc
// 内部固定，dashboard 缺口由 trail 起点偏移表现）。
func (p Progress) buildRing(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	size, stroke, sweepDeg := p.ringGeometry()
	v := p.Value
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	fillColor := th.Color.Primary
	if v >= 1 && p.Status == ProgressNormal {
		fillColor = th.Color.Success
	}
	switch p.Status {
	case ProgressSuccess:
		fillColor = th.Color.Success
	case ProgressException:
		fillColor = th.Color.Danger
	}
	// 环带用**填充扇区**而非描边弧：外弧正向+内弧反向闭合的环带多边形
	// 点数 = 内外两层弧各 ~20 点（Flatten 宽容差）≈ 40 点；等价 6px 描边
	// 弧的 stroke-to-fill 展开是 1.65 万点（每顶点段四边形+双 join 三角），
	// 填充活跃边表成本差 ~400 倍——progress 圆环首帧 4s（规则 §6A 案例）
	// 的根治手段。AddArc 第二参是扫掠角（Spin 轮已验证）。
	inner := geom.RectXYWH(stroke, stroke, size-2*stroke, size-2*stroke)
	outer := geom.RectXYWH(0, 0, size, size)
	arcBand := func(sweep float32) *render.Path {
		if sweep <= 0 {
			return nil
		}
		if sweep >= 360 {
			// 整环 = 外圆 + 内圆（evenodd 挖洞）；AddOval 4 段贝塞尔，
			// 描边/填充均几十点级。
			return render.NewPath().AddOval(outer).AddOval(inner)
		}
		band := render.NewPath()
		// 外弧正向（起 0°）。
		band.AddArc(outer, 0, sweep)
		// 内弧反向闭合：从外弧终点连到内弧同角点，内弧按 -sweep 回走。
		band.AddArc(inner, sweep, -sweep)
		band.Close()
		return band
	}
	ring := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		fillPaint := render.Fill(th.Color.SurfaceActive)
		// trail 灰环。
		if band := arcBand(sweepDeg); band != nil {
			cv.FillPath(band, fillPaint)
		}
		// 进度弧（v*sweep 扫掠）。
		if v > 0 {
			if band := arcBand(sweepDeg * v); band != nil {
				cv.FillPath(band, render.Fill(fillColor))
			}
		}
	}).SetSize2(size, size)
	ring.SetHit(scene.HitSpec{Label: "progress-circle"})
	if !p.ShowInfo {
		return ring
	}
	stack := scene.Stack().SetSize2(size, size)
	stack.Add(ring)
	row := scene.Flex(scene.Horizontal, 0).SetMainAlign(scene.MainCenter)
	pct := (&Label{Text: pctText(v), Size: th.Font.SizeBase, Color: fillColor}).Build(ctx)
	pct.SetHit(scene.HitSpec{Label: "progress-text"})
	row.Add(pct)
	row.SetPositionKeepSize(true)
	row.SetStackPosition(geom.Pt(0, (size-20)/2))
	stack.Add(row)
	stack.SetHit(scene.HitSpec{Label: "progress"})
	return stack
}

// buildLine 构建线形进度（原实现，行为不变）。
func (p Progress) buildLine(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	w, h := p.Width, p.Height
	if w <= 0 {
		w = 160
	}
	if h <= 0 {
		h = 8 // antd Progress 线形 strokeWidth 默认 8（此前 6 偏细）
	}
	// trail 用 SurfaceActive（#F0F0F0，antd 线形进度条默认灰轨）——
	// SurfaceSunken 与演示舞台同色，底轨会隐形。默认高 8 = antd Progress
	// lineHeight（antd 默认轨道 8px，此前 6 偏细）。
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
		track.SetHit(scene.HitSpec{Label: "progress"})
		return track
	}
	row := scene.Flex(scene.Horizontal, 8).SetCrossAlign(scene.CrossCenter)
	row.Add(track)
	txt := (&Label{Text: pctText(v), Size: th.Font.SizeSM, Color: fillColor}).Build(ctx)
	txt.SetHit(scene.HitSpec{Label: "progress-text"}) // 被动锚：文字色断言
	row.Add(txt)
	row.SetHit(scene.HitSpec{Label: "progress"})
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
