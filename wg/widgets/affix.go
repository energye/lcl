package widgets

// ── Affix 固钉（M12 收尾）────────────────────────────────────────
//
// antd：<Affix offsetTop offsetBottom onChange> 滚动固定：页面滚动超
// 阈值（offsetTop）或距底部不足（offsetBottom，吸底场景）时元素转为
// 固定态。**滚动上下文注入**（Calendar.Now 同款可测性设计）：
// ScrollOffset func() float32 由宿主提供（如组合在 Scroll 内时闭
// 包捕获 Scroll.Offset）；BottomOffset func() float32 提供距底查询
// （吸底场景——宿主闭包捕获内容高-可见高等折算）；宿主滚动触发外层
// 重建 → Affix.Build 重跑 → 查询切换 affixed 态。视觉：固定态加主色
// 边框+"已固定"角标（antd 固定态语义视觉简化）。
// 未实现（如实标注）：真实 fixed 定位（引擎无视口固定层级——固定
// 态为语义视觉）、target 容器、刷新时序。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
)

// Affix 是固钉（滚动超阈值切换固定态）。
type Affix struct {
	controlBase
	Child ui.Widget
	// OffsetTop 触发阈值（滚动 offset ≥ 此值进入固定态）。
	OffsetTop float32
	// OffsetBottom 距底触发阈值（antd offsetBottom——距底部 ≤ 此值进入
	// 固定态，吸底场景；0=不启用）。
	OffsetBottom float32
	// ScrollOffset 滚动查询注入（宿主闭包；nil=顶部永不触发）。
	ScrollOffset func() float32
	// BottomOffset 距底查询注入（宿主闭包——内容高-可见高-滚动 offset
	// 等折算；nil=底部永不触发）。
	BottomOffset func() float32
	// OnAffixChange 固定态变化回调。
	OnAffixChange func(affixed bool)
	Width         float32
	Height        float32

	affixed bool
}

// NewAffix 创建固钉。
func NewAffix(child ui.Widget) *Affix {
	return &Affix{Child: child}
}

// Affixed 返回当前是否固定态。
func (a *Affix) Affixed() bool { return a.affixed }

// Build 实现 ui.Widget。
func (a *Affix) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(a)
	th := ctx.ThemeOr()
	w := a.Width
	if w <= 0 {
		w = 200
	}
	h := a.Height
	if h <= 0 {
		h = 48
	}

	// 查询滚动 → 切换 affixed（翻转只同步+通知，不自标脏防循环——
	// 本次 Build 已是新状态，宿主滚动帧会再次进来）。双阈值独立启用：
	// 顶部判定要求 ScrollOffset 注入、底部判定要求 BottomOffset 注入
	// 且 OffsetBottom>0（antd offsetTop/offsetBottom undefined=不参与
	// 的语义——吸底卡不注入 ScrollOffset 时顶部判定不误触发）。
	affixed := false
	if a.ScrollOffset != nil && a.ScrollOffset() >= a.OffsetTop {
		affixed = true
	}
	if !affixed && a.OffsetBottom > 0 && a.BottomOffset != nil && a.BottomOffset() <= a.OffsetBottom {
		affixed = true
	}
	if affixed != a.affixed {
		a.affixed = affixed
		if a.OnAffixChange != nil {
			a.OnAffixChange(affixed)
		}
	}

	var inner *scene.Node
	if a.Child != nil {
		inner = a.Child.Build(ctx)
	}
	if inner == nil {
		inner = scene.Empty()
	}

	box := scene.Box(scene.BoxSpec{
		Fill: func() render.Paint {
			if affixed {
				return scene.Solid(th.Color.InfoBg)
			}
			return scene.Solid(th.Color.Surface)
		}(),
		Border: func() render.Paint {
			if affixed {
				return scene.Stroke(th.Color.Primary, 1)
			}
			return scene.Stroke(th.Color.Border, 1)
		}(),
	}).SetSize2(w, h)

	inner.SetMargin(geom.InLTRB(10, (h-inner.Size().H)/2, 10, 0))
	box.Add(inner)
	if affixed {
		l := text.LayoutText("已固定", text.DefaultFont(th.Font.SizeXS), text.Options{})
		badge := scene.Text(scene.TextSpec{Src: "已固定", Layout: l, Color: th.Color.Primary})
		badge.SetMargin(geom.InLTRB(0, (h-l.Size.H)/2, 10, 0))
		badge.SetPositionKeepSize(true)
		badge.SetStackPosition(geom.Pt(w-l.Size.W-10, 0))
		badge.SetHit(scene.HitSpec{Label: "affix-badge"})
		box.Add(badge)
	}
	box.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "affix"})
	return box
}
