package main

import (
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var (
	// scrollSim/bottomSim 模拟滚动源（审计 wheel 坐标硬编码局限——按钮驱动确定性）。
	scrollSim float32
	bottomSim = float32(1000) // 距底距离（初始远离底部）
	affix     = widgets.NewAffix(&widgets.Label{Text: "固钉内容 · 滚动超阈值进入固定态"})
	affixBot  = widgets.NewAffix(&widgets.Label{Text: "吸底固钉 · 距底部不足进入固定态"})
	affixLog  = ""
	botLog    = ""
)

func init() {
	affix.OffsetTop = 100
	affix.ScrollOffset = func() float32 { return scrollSim }
	affix.OnAffixChange = func(affixed bool) {
		if affixed {
			affixLog = "已固定"
		} else {
			affixLog = "未固定"
		}
	}
	// 吸底卡：OffsetBottom=120 距底阈值；BottomOffset 注入距底查询。
	affixBot.OffsetBottom = 120
	affixBot.BottomOffset = func() float32 { return bottomSim }
	affixBot.OnAffixChange = func(affixed bool) {
		if affixed {
			botLog = "已吸底"
		} else {
			botLog = "未吸底"
		}
	}
}

func main() {
	launcher.Main(doc.Meta{
		ID: "affix", Title: "Affix 固钉", Group: doc.GroupControl,
		What: []string{
			"滚动超阈值切换固定态：**ScrollOffset 查询注入**（Calendar.Now 同款可测性设计——宿主滚动帧重跑 Build 查询）",
			"吸底（offsetBottom）：**BottomOffset 距底查询注入**——距底部 ≤ OffsetBottom 进入固定态（antd 双阈值语义）",
			"固定态视觉：InfoBg 底+主色边框+「已固定」角标；OnAffixChange 回调",
			"demo 用按钮驱动模拟滚动（审计 wheel 注入坐标硬编码局限）；真实 fixed 定位未实现（如实标注）",
		},
		Demos: []doc.Demo{
			{Title: "滚动固定", Note: "点「模拟滚动」驱动 offset=150>100 → 固定态；点「模拟回滚」还原", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 8).SetCrossAlign(scene.CrossCenter)
				b1 := &widgets.Button{Text: "模拟滚动", Type: widgets.TypePrimary, Size: ui.SizeSmall, OnClick: func() {
					scrollSim = 150
					affix.MarkDirty()
				}}
				b2 := &widgets.Button{Text: "模拟回滚", Size: ui.SizeSmall, OnClick: func() {
					scrollSim = 0
					affix.MarkDirty()
				}}
				row.Add(b1.Build(ctx))
				row.Add(b2.Build(ctx))
				affixWrap := scene.Flex(scene.Vertical, 8)
				affixWrap.Add(row)
				affixWrap.Add(affix.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return affixWrap })
			}},
			{Title: "吸底", Note: "OffsetBottom=120：点「滚到底部」驱动距底=80<120 → 吸底态；点「滚回顶部」还原", Build: func(ctx *ui.BuildContext) ui.Widget {
				row := scene.Flex(scene.Horizontal, 8).SetCrossAlign(scene.CrossCenter)
				b1 := &widgets.Button{Text: "滚到底部", Type: widgets.TypePrimary, Size: ui.SizeSmall, OnClick: func() {
					bottomSim = 80
					affixBot.MarkDirty()
				}}
				b2 := &widgets.Button{Text: "滚回顶部", Size: ui.SizeSmall, OnClick: func() {
					bottomSim = 1000
					affixBot.MarkDirty()
				}}
				row.Add(b1.Build(ctx))
				row.Add(b2.Build(ctx))
				wrap := scene.Flex(scene.Vertical, 8)
				wrap.Add(row)
				wrap.Add(affixBot.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return wrap })
			}},
		},
		API: []doc.APIRow{
			{Name: "NewAffix", Type: "func(child ui.Widget) *Affix", Desc: "构造"},
			{Name: "OffsetTop", Type: "float32", Desc: "触发阈值（offset ≥ 此值进入固定态）"},
			{Name: "OffsetBottom", Type: "float32", Default: "0", Desc: "距底触发阈值（距底 ≤ 此值进入固定态，吸底；0=不启用）"},
			{Name: "ScrollOffset", Type: "func() float32", Desc: "滚动查询注入（宿主闭包）"},
			{Name: "BottomOffset", Type: "func() float32", Desc: "距底查询注入（宿主闭包——配合 OffsetBottom）"},
			{Name: "OnAffixChange", Type: "func(affixed bool)", Desc: "固定态变化回调"},
			{Name: "Affixed", Type: "func() bool", Desc: "当前态查询"},
		},
	})
}
