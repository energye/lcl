package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var (
	anV = widgets.NewAnchor(
		widgets.AnchorItem{Key: "intro", Title: "基本介绍"},
		widgets.AnchorItem{Key: "usage", Title: "何时使用"},
		widgets.AnchorItem{Key: "api", Title: "API 说明"},
	)
	anH = widgets.NewAnchor(
		widgets.AnchorItem{Key: "basic", Title: "Basic"},
		widgets.AnchorItem{Key: "static", Title: "Static"},
		widgets.AnchorItem{Key: "api", Title: "API"},
	)
	anLog  string
	clickN = widgets.NewButton("模拟滚动到 API 区块")
)

func init() {
	anH.Direction = widgets.AnchorHorizontal
	anH.Width = 360
	anV.Width = 140
	anV.OnClick = func(key string) {
		anLog = "点击: " + key
	}
	anV.OnChange = func(key string) {
		anLog += fmt.Sprintf(" → 高亮 %s", key)
	}
	// 模拟滚动按钮：注入 SectionTop 均分假设下滚动到 420（API 区块）。
	clickN.OnClick = func() {
		anV.ScrollPos = func() float32 { return 420 }
		anV.SectionTop = func(key string) float32 {
			switch key {
			case "intro":
				return 0
			case "usage":
				return 200
			default:
				return 400
			}
		}
		anV.MarkDirty() // 重建时 syncScroll 翻转到 api
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "anchor", Title: "Anchor 锚点", Group: doc.GroupControl,
		What: []string{
			"页面锚点导航（antd Anchor 对应）：左侧 2px 主色墨条（ink bar）+ 锚项列表",
			"点击高亮迁移（宿主在 OnClick 里驱动自家滚动——wg 无浏览器滚动宿主）",
			"滚动位置监听经 ScrollPos/SectionTop 查询注入（Affix ScrollOffset 同源可测性设计）",
		},
		Demos: []doc.Demo{
			{Title: "竖排锚点", Note: "墨条随高亮迁移；点锚项看 OnClick/OnChange 日志", Build: func(ctx *ui.BuildContext) ui.Widget {
				return anV
			}},
			{Title: "横排锚点", Note: "Direction=horizontal：底部墨条", Build: func(ctx *ui.BuildContext) ui.Widget {
				return anH
			}},
			{Title: "滚动监听", Note: "ScrollPos/SectionTop 注入：点按钮模拟滚动到 API 区块——高亮自动翻转（OnChange 触发）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					col := scene.Flex(scene.Vertical, 12)
					col.Add(clickN.Build(c))
					if anLog != "" {
						lb := (&widgets.Label{Text: anLog}).Build(c)
						lb.SetHit(scene.HitSpec{Label: "anchor-log"})
						col.Add(lb)
					}
					return col
				})
			}},
		},
		API: []doc.APIRow{
			{Name: "NewAnchor", Type: "func(items ...AnchorItem) *Anchor", Desc: "构造"},
			{Name: "Items", Type: "[]AnchorItem{Key, Title}", Desc: "锚项（antd items 简化：无嵌套）"},
			{Name: "Direction", Type: "AnchorDirection", Default: "AnchorVertical", Desc: "竖排/横排"},
			{Name: "ActiveKey", Type: "string", Desc: "当前高亮锚（可受控赋值）"},
			{Name: "OnClick / OnChange", Type: "func(key string)", Desc: "点击 / 高亮变化回调"},
			{Name: "ScrollPos / SectionTop", Type: "func() float32 / func(key) float32", Desc: "滚动位置与区块坐标查询注入（滚动监听）"},
			{Name: "Bounds", Type: "float32", Default: "5", Desc: "高亮判定提前量（antd bounds）"},
			{Name: "Width", Type: "float32", Default: "120", Desc: "宽度"},
		},
	})
}
