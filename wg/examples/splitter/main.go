package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 实例必须包级持有（教训⑨）。
var (
	spH = func() *widgets.Splitter {
		s := widgets.NewSplitter(
			widgets.SplitterPanel{Content: panelCard("左面板", "默认 40%"),
				DefaultSize: 240, Min: 160, Max: 400},
			widgets.SplitterPanel{Content: panelCard("右面板", "平分剩余"),
				Min: 160},
		)
		s.Width, s.Height = 600, 160
		return s
	}()
	spV = func() *widgets.Splitter {
		s := widgets.NewSplitter(
			widgets.SplitterPanel{Content: panelCard("上面板", "40%")},
			widgets.SplitterPanel{Content: panelCard("下面板", "60%")},
		)
		s.Orientation = widgets.SplitVertical
		s.Width, s.Height = 600, 280
		return s
	}()
	spProg = func() *widgets.Splitter {
		s := widgets.NewSplitter(
			widgets.SplitterPanel{Content: panelCard("A", ""), Min: 120},
			widgets.SplitterPanel{Content: panelCard("B", ""), Min: 120},
		)
		s.Width, s.Height = 600, 140
		return s
	}()
	swapBtn   = widgets.NewButton("左右对调")
	resizeBtn = widgets.NewButton("右面板吃满")
	resizeLog = ""
)

// panelCard 面板内容（灰底圆角卡+标题）。
func panelCard(title, note string) ui.Widget {
	return doc.Node(func(c *ui.BuildContext) *scene.Node {
		box := scene.Box(scene.BoxSpec{Fill: scene.Solid(geom.RGBA(245, 245, 245)), Radius: 4})
		box.SetFlex(1, scene.FlexTight)
		inner := scene.Flex(scene.Vertical, 4)
		lb := &widgets.Label{Text: title}
		inner.Add(lb.Build(c))
		if note != "" {
			inner.Add((&widgets.Label{Text: note}).Build(c))
		}
		box.Add(inner)
		box.SetHit(scene.HitSpec{Behavior: scene.HitOpaque, Label: "split-card"})
		return box
	})
}

func init() {
	swapBtn.OnClick = func() {
		sz := spProg.Sizes()
		spProg.SetSizes([]float32{sz[1], sz[0]})
		resizeLog = fmt.Sprintf("对调后: [%.0f %.0f]", sz[1], sz[0])
	}
	resizeBtn.OnClick = func() {
		spProg.SetSizes([]float32{120, 0}) // A 钳 Min 120，B 吃余差
		resizeLog = "B 吃满余差: [120 478]"
	}
	spProg.OnResize = func(sz []float32) {
		resizeLog = fmt.Sprintf("OnResize: [%.0f %.0f]", sz[0], sz[1])
	}
}

// buildChildLog 日志文本行。
func buildChildLog(s string, c *ui.BuildContext) *scene.Node {
	n := (&widgets.Label{Text: s}).Build(c)
	n.SetHit(scene.HitSpec{Label: "splitter-log"})
	return n
}

func main() {
	launcher.Main(doc.Meta{
		ID: "splitter", Title: "Splitter 分栏", Group: doc.GroupControl,
		What: []string{
			"可拖拽分栏（antd Splitter/Panel 对应）：2px 把手+6px 触发区，hover/拖拽中主色",
			"Orientation 水平/垂直；Panel DefaultSize/Min/Max 钳制；Resizing 走 Pan 手势",
			"SetSizes 编程驱动（受控语义走 OnResize 回调——审计/自动化入口）；末面板吃余差守恒总长",
		},
		Demos: []doc.Demo{
			{Title: "水平分栏", Note: "左 240 (Min 160/Max 400) + 右平分；拖中间把手调整", Build: func(ctx *ui.BuildContext) ui.Widget {
				return spH
			}},
			{Title: "垂直分栏", Note: "Orientation=vertical：上下分栏", Build: func(ctx *ui.BuildContext) ui.Widget {
				return spV
			}},
			{Title: "编程驱动", Note: "SetSizes 受控：按钮左右对调/右吃满——OnResize 回调同步（真拖拽 Move 注入 headless 不稳的务实通道）", Build: func(ctx *ui.BuildContext) ui.Widget {
				return doc.Node(func(c *ui.BuildContext) *scene.Node {
					col := scene.Flex(scene.Vertical, 12)
					col.Add(spProg.Build(c))
					row := scene.Flex(scene.Horizontal, 8)
					row.Add(swapBtn.Build(c))
					row.Add(resizeBtn.Build(c))
					col.Add(row)
					if resizeLog != "" {
						col.Add(buildChildLog(resizeLog, c))
					}
					return col
				})
			}},
		},
		API: []doc.APIRow{
			{Name: "NewSplitter", Type: "func(panels ...SplitterPanel) *Splitter", Desc: "构造"},
			{Name: "Orientation", Type: "SplitterOrientation", Default: "SplitHorizontal", Desc: "分栏方向（水平/垂直）"},
			{Name: "Panel.DefaultSize", Type: "float32", Desc: "初始主轴像素（<=0 平分剩余）"},
			{Name: "Panel.Min / Max", Type: "float32", Desc: "主轴像素钳制（<=0 不限）"},
			{Name: "Panel.Resizable", Type: "bool", Default: "true", Desc: "是否可拖拽"},
			{Name: "OnResize / Start / End", Type: "func(sizes []float32)", Desc: "尺寸变化回调（拖拽中/开始/结束）"},
			{Name: "SetSizes", Type: "func(sizes []float32)", Desc: "编程设置（受控语义，末面板吃余差守恒）"},
			{Name: "Width / Height", Type: "float32", Default: "600 / 300", Desc: "总尺寸"},
		},
	})
}
