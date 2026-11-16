package examples

import (
	"fmt"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// Gallery 是真窗示例集的主页面：左侧是示例导航，右侧是当前示例。
//
// 它自身就是一个 ui.Widget，点击导航走的是完整的真窗链路：
// 平台鼠标消息 → host.Binder → Dispatcher → GestureDetector → SetState
// → 容器下一帧重建 → 新节点树参与布局/合成/Present。
// 因此画廊不只是一个展示壳，它本身就是事件与重建的真窗回归样例。
type Gallery struct {
	ui.Base
	selected int
}

// NewGallery 创建示例画廊，默认展示第一条。
func NewGallery() *Gallery { return &Gallery{} }

// Current 返回当前示例（诊断/自动化用）。
func (g *Gallery) Current() Example {
	if len(All) == 0 {
		return Example{}
	}
	if g.selected < 0 || g.selected >= len(All) {
		return All[0]
	}
	return All[g.selected]
}

// Select 切到指定 ID，找不到返回 false。调用方在真窗线程调用即可。
func (g *Gallery) Select(id string) bool {
	for i, e := range All {
		if e.ID == id {
			i := i
			g.SetState(func() { g.selected = i })
			return true
		}
	}
	return false
}

// Next 切到下一条示例，便于在真窗里逐项走查。
func (g *Gallery) Next() {
	if len(All) == 0 {
		return
	}
	g.SetState(func() { g.selected = (g.selected + 1) % len(All) })
}

// Build 实现 ui.Widget。
func (g *Gallery) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	root := scene.Flex(scene.Horizontal, 1)

	// ── 左侧导航 ────────────────────────────────────────────────
	nav := scene.Flex(scene.Vertical, 3).SetWidth(226)
	nav.SetPadding(geom.InLTRB(8, 8, 8, 8))
	nav.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Surface)})
	nav.Add((&widgets.Label{Text: "wg 示例集", Size: th.Font.SizeTitle}).Build(ctx))
	nav.Add((&widgets.Label{Text: fmt.Sprintf("共 %d 项 · 真窗运行", len(All)), Size: th.Font.SizeSM}).Build(ctx))
	nav.Add((&widgets.Divider{}).Build(ctx))

	lastGroup := ""
	for i, ex := range All {
		if ex.Group != lastGroup {
			lastGroup = ex.Group
			nav.Add((&widgets.Label{Text: "— " + lastGroup, Size: th.Font.SizeSM}).Build(ctx))
		}
		b := widgets.NewButton(ex.Name)
		i := i
		b.OnClick = func() {
			idx := i
			g.SetState(func() { g.selected = idx })
		}
		nav.Add(b.Build(ctx))
	}

	next := widgets.NewButton("→ 下一项")
	next.OnClick = g.Next
	nav.Add(next.Build(ctx))
	root.Add(nav)

	// ── 右侧内容 ────────────────────────────────────────────────
	body := scene.Flex(scene.Vertical, 8).SetFlex(1, scene.FlexTight)
	body.SetPadding(geom.InLTRB(16, 12, 16, 16))
	body.SetBox(scene.BoxSpec{Fill: scene.Solid(th.Color.Background)})

	cur := g.Current()
	body.Add((&widgets.Label{
		Text: fmt.Sprintf("%s / %s", cur.Group, cur.Name),
		Size: th.Font.SizeTitle,
	}).Build(ctx))
	body.Add((&widgets.Divider{}).Build(ctx))
	if cur.Build != nil {
		body.Add(cur.Build(ctx))
	} else {
		body.Add(Note(ctx, "此示例尚未实现"))
	}
	root.Add(body)
	return root
}
