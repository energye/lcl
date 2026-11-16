package main

import (
	"fmt"

	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/examples/internal/launcher"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

var status = widgets.NewButton("按任意键（先点一下此区域）")

func main() {
	launcher.Main(doc.Meta{
		ID: "keyboard", Title: "键盘事件", Group: doc.GroupInput,
		What: []string{
			"键盘事件按焦点域分发：VK → Code 归一化；焦点节点未处理时沿父链冒泡",
			"编辑类按键（Input 内）由 TextKeyHandler 统一处理（见 input 窗）",
		},
		Demos: []doc.Demo{
			{Title: "按键回显", Note: "点击区域抢焦点后按任意键，回显 Kind/Code/Mods", Build: func(ctx *ui.BuildContext) ui.Widget {
				th := ctx.ThemeOr()
				area := scene.Box(scene.BoxSpec{Fill: scene.Solid(th.Color.SurfaceActive)}).SetSize2(420, 90)
				area.Add((&widgets.Label{Text: "键盘目标区域"}).Build(ctx))
				if ctx.Input != nil {
					ctx.Input.Set(area, input.Handler{
						Key: func(ev *input.KeyEvent) bool {
							status.Text = fmt.Sprintf("Kind=%d code=%d text=%q mods=%v", ev.Kind, ev.Code, ev.Text, ev.Mods)
							status.MarkDirty()
							return true
						},
					})
				}
				col := scene.Flex(scene.Vertical, 10)
				col.Add(area)
				col.Add(status.Build(ctx))
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
		API: []doc.APIRow{
			{Name: "Handler.Key", Type: "func(*KeyEvent) bool", Desc: "键盘回调；返回 true 表示已处理（停止冒泡）"},
			{Name: "KeyEvent", Type: "{Kind,Code,Text,Mods}", Desc: "归一化键盘事件"},
			{Name: "冒泡", Type: "契约", Desc: "焦点节点未处理时沿父链上抛"},
		},
	})
}
