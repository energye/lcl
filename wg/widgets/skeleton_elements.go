package widgets

// ── Skeleton 元素形态（antd Skeleton.Button/Input/Image）────────────
//
// 三种独立占位元素，复用 Skeleton 的呼吸底（Animated 时按墙钟实时相位）：
//   - SkeletonButton：高度随 Size 三档（24/32/40），Shape 控制圆角，
//     Block 撑满父宽；默认宽 64。
//   - SkeletonInput：高度三档，默认宽 200。
//   - SkeletonImage：固定方形占位（默认 96×96），圆角 8。
//
// 未实现（如实标注）：Skeleton.Avatar 已由 Skeleton.Avatar 开关覆盖（圆形
// 头像条）；Skeleton.Node（任意尺寸盒）未做。

import (
	"time"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// SkeletonShape 是骨架元素的形状（antd Skeleton.Button shape）。
type SkeletonShape uint8

const (
	// SkelShapeDefault 默认（圆角 4）。
	SkelShapeDefault SkeletonShape = iota
	// SkelShapeSquare 直角。
	SkelShapeSquare
	// SkelShapeRound 胶囊圆角。
	SkelShapeRound
	// SkelShapeCircle 正圆（宽高相等）。
	SkelShapeCircle
)

// skelHeight 按 Size 档取高度（antd controlHeight 24/32/40）。
func skelHeight(s ui.Size) float32 {
	switch s {
	case ui.SizeSmall:
		return 24
	case ui.SizeLarge:
		return 40
	}
	return 32
}

// skelRadius 按 Shape 取圆角。
func skelRadius(sh SkeletonShape, h float32) float32 {
	switch sh {
	case SkelShapeSquare:
		return 0
	case SkelShapeRound:
		return h / 2
	case SkelShapeCircle:
		return h / 2
	}
	return 4
}

// skelBar 构建一条呼吸底（Animated 时实时相位，同 Skeleton.bar）。
func skelBar(ctx *ui.BuildContext, th *theme.Data, w, h, round float32, animated bool) *scene.Node {
	if !animated {
		return scene.Box(scene.BoxSpec{
			Fill: scene.Solid(th.Color.SurfaceActive), Radius: round,
		}).SetSize2(w, h)
	}
	cp := render.NewPath().AddRRect(geom.RectXYWH(0, 0, w, h), round)
	return scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		const breathePeriod = 1400 * time.Millisecond
		p := float32(time.Since(skeletonBase)%breathePeriod) / float32(breathePeriod)
		op := 1.0 - 0.55*(0.5-0.5*cos32(p*2*3.14159265))
		col := th.Color.SurfaceActive
		cv.FillPath(cp, render.Fill(geom.RGBAa(col.R, col.G, col.B, uint8(op*255))))
	}).SetSize2(w, h)
}

// SkeletonButton 是按钮骨架（antd Skeleton.Button）。
type SkeletonButton struct {
	Size     ui.Size
	Shape    SkeletonShape
	Block    bool
	Width    float32
	Animated bool
}

// Build 实现 ui.Widget。
func (b SkeletonButton) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	h := skelHeight(b.Size)
	w := b.Width
	if b.Shape == SkelShapeCircle {
		w = h
	} else if w <= 0 {
		w = 64
	}
	if b.Block {
		if vp := ctx.ViewportSize(); vp.W > 0 {
			w = vp.W - 80
		}
	}
	n := skelBar(ctx, th, w, h, skelRadius(b.Shape, h), b.Animated)
	n.SetHit(scene.HitSpec{Label: "skeleton-button"})
	return n
}

// SkeletonInput 是输入框骨架（antd Skeleton.Input）。
type SkeletonInput struct {
	Size     ui.Size
	Width    float32
	Animated bool
}

// Build 实现 ui.Widget。
func (s SkeletonInput) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	h := skelHeight(s.Size)
	w := s.Width
	if w <= 0 {
		w = 200
	}
	n := skelBar(ctx, th, w, h, 4, s.Animated)
	n.SetHit(scene.HitSpec{Label: "skeleton-input"})
	return n
}

// SkeletonImage 是图片骨架（antd Skeleton.Image，固定方形）。
type SkeletonImage struct {
	Width, Height float32
	Animated      bool
}

// Build 实现 ui.Widget。
func (s SkeletonImage) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	w, h := s.Width, s.Height
	if w <= 0 {
		w = 96
	}
	if h <= 0 {
		h = 96
	}
	n := skelBar(ctx, th, w, h, 8, s.Animated)
	n.SetHit(scene.HitSpec{Label: "skeleton-image"})
	return n
}
