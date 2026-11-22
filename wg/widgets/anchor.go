package widgets

// ── Anchor 锚点导航（第一百零八轮，缺失组件清零）─────────────────
//
// antd：<Anchor items direction onChange onClick affix>。
// 页面锚点导航：左侧 2px 主色墨条（ink bar）+ 锚项列表，点击切换
// 高亮（宿主可在 OnClick 里驱动自家滚动）；滚动位置监听经
// **ScrollPos/SectionTop 查询注入**（Affix ScrollOffset 同源可测性
// 设计——nil=纯点击高亮）。
//
// wg 无浏览器滚动宿主：点击不真滚（SetScrollOffset 先例——宿主
// 组合驱动）；滚动 Build 内查询→最近区块→ActiveKey 翻转只同步+
// OnChange 不自标脏（Affix 防翻转循环同款纪律）。

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// AnchorItem 是锚项定义（antd AnchorItem 简化：无嵌套 children）。
type AnchorItem struct {
	Key   string
	Title string
}

// AnchorDirection 是锚点方向（antd direction）。
type AnchorDirection uint8

const (
	// AnchorVertical 竖排（默认，带墨条）。
	AnchorVertical AnchorDirection = iota
	// AnchorHorizontal 横排（底部墨条）。
	AnchorHorizontal
)

// Anchor 是锚点导航。
type Anchor struct {
	controlBase
	Items     []AnchorItem
	Direction AnchorDirection
	// ActiveKey 当前高亮锚（点击或滚动监听更新；可外部受控赋值）。
	ActiveKey string
	// OnChange 高亮锚变化回调（含点击与滚动监听触发）。
	OnChange func(key string)
	// OnClick 锚项点击回调（宿主据此驱动自家滚动）。
	OnClick func(key string)
	// ScrollPos 滚动容器位置查询注入（nil=不滚动监听）。
	ScrollPos func() float32
	// SectionTop 各锚区块在滚动内容中的 y 坐标（nil=按 Items 序号
	// ×200 均分假设）。
	SectionTop func(key string) float32
	// Bounds 高亮判定距离（antd bounds 默认 5；wg 用最近区块语义，
	// 该值作为区块顶判定提前量）。
	Bounds float32

	Width float32

	taps []*ui.GestureDetector
}

// NewAnchor 创建锚点导航。
func NewAnchor(items ...AnchorItem) *Anchor {
	return &Anchor{Items: items, Bounds: 5}
}

// buildItem 构建锚项行（文字+手势）。
func (a *Anchor) buildItem(ctx *ui.BuildContext, th *theme.Data, it AnchorItem) *scene.Node {
	active := it.Key == a.ActiveKey
	color := th.Color.TextMuted
	if active {
		color = th.Color.Primary
	}
	l := text.LayoutText(it.Title, text.DefaultFont(th.Font.SizeBase), text.Options{})
	tv := scene.Text(scene.TextSpec{Src: it.Title, Layout: l, Color: color})
	tv.SetMargin(geom.InLTRB(12, 4, 4, 4))
	tv.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "anchor-item", Cursor: scene.CursorPointer})
	if ctx.Input != nil {
		key := it.Key
		g := &ui.GestureDetector{OnTap: func(geom.Point) {
			a.activate(key)
			if a.OnClick != nil {
				a.OnClick(key)
			}
		}}
		a.taps = append(a.taps, g)
		ctx.Input.Set(tv, g.Handler())
	}
	return tv
}

// activate 切换高亮并重建。
func (a *Anchor) activate(key string) {
	if a.ActiveKey == key {
		return
	}
	a.ActiveKey = key
	a.MarkDirty()
	if a.OnChange != nil {
		a.OnChange(key)
	}
}

// syncScroll 滚动位置→最近区块（只同步+OnChange，不自标脏——宿主
// 滚动帧重跑 Build 时查询；Affix 防翻转循环同款纪律）。
func (a *Anchor) syncScroll() {
	if a.ScrollPos == nil || len(a.Items) == 0 {
		return
	}
	pos := a.ScrollPos()
	best, bestKey := -1, a.ActiveKey
	for _, it := range a.Items {
		top := float32(0)
		if a.SectionTop != nil {
			top = a.SectionTop(it.Key)
		} else {
			top = float32(indexOf(a.Items, it.Key)) * 200
		}
		if pos+a.Bounds >= top {
			best++
			bestKey = it.Key
		}
	}
	if best >= 0 && bestKey != a.ActiveKey {
		a.ActiveKey = bestKey
		if a.OnChange != nil {
			a.OnChange(bestKey)
		}
	}
}

func indexOf(items []AnchorItem, key string) int {
	for i, it := range items {
		if it.Key == key {
			return i
		}
	}
	return 0
}

// inkBar 构建墨条（竖向 2px 主色竖线 / 横向底部 2px 线）。
func (a *Anchor) inkBar(th *theme.Data, height float32) *scene.Node {
	if a.Direction == AnchorHorizontal {
		return scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary), Radius: th.Radius.Pill}).
			SetHeight(2).SetFlex(1, scene.FlexTight)
	}
	bar := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary), Radius: th.Radius.Pill})
	bar.SetSize2(2, height)
	return bar
}

// Build 实现 ui.Widget。
func (a *Anchor) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(a)
	th := ctx.ThemeOr()
	a.taps = a.taps[:0]
	if a.ActiveKey == "" && len(a.Items) > 0 {
		a.ActiveKey = a.Items[0].Key
	}
	a.syncScroll()
	w := a.Width
	if w <= 0 {
		w = 120
	}
	itemH := float32(28)
	if a.Direction == AnchorHorizontal {
		// 横排：行内锚项 + 底部墨条（简化为通栏主色线）。
		col := scene.Flex(scene.Vertical, 0)
		col.SetWidth(w)
		col.SetHit(scene.HitSpec{Label: "anchor"})
		row := scene.Flex(scene.Horizontal, th.Space.SM)
		for _, it := range a.Items {
			row.Add(a.buildItem(ctx, th, it))
		}
		col.Add(row)
		col.Add(a.inkBar(th, 0))
		return col
	}
	// 竖排：左墨条列 + 锚项列（antd root Flex Horizontal）。
	row := scene.Flex(scene.Horizontal, 0)
	row.SetWidth(w)
	row.SetHit(scene.HitSpec{Label: "anchor"})
	// 墨条轨：全高灰轨不可见（antd rail 是 split 色淡线）+ 主色墨条
	// 贴 active 项位置。
	rail := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Border)})
	rail.SetSize2(2, itemH*float32(len(a.Items)))
	ink := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.Primary), Radius: th.Radius.Pill})
	activeIdx := 0
	for i, it := range a.Items {
		if it.Key == a.ActiveKey {
			activeIdx = i
			break
		}
	}
	inkY := float32(activeIdx)*itemH + 4
	ink.SetSize2(2, 18)
	inkStack := scene.Stack().SetSize2(2, rail.Size().H)
	rail.SetPositionKeepSize(true)
	rail.SetStackPosition(geom.Pt(0, 0))
	ink.SetPositionKeepSize(true)
	ink.SetStackPosition(geom.Pt(0, inkY))
	inkStack.Add(rail)
	inkStack.Add(ink)
	inkStack.SetHit(scene.HitSpec{Label: "anchor-ink"})
	row.Add(inkStack)
	itemsCol := scene.Flex(scene.Vertical, 0)
	for _, it := range a.Items {
		itemsCol.Add(a.buildItem(ctx, th, it))
	}
	row.Add(itemsCol)
	return row
}

var _ = input.PointerDown // 保留 input 引用
