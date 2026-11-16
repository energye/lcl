// Command wg-formdemo 是 M3 的表单页示例：一张完全由 wg 组件搭出来的页面，
// 不直接依赖 lcl，也不需要显示器（走 NewHeadless + 手动事件注入）。
//
// 运行：go run ./wg/cmd/wg-formdemo
//
// 它演示四件真实程序里一定会用到的事：
//   - MountWidget 挂一棵有状态的组件树；
//   - 用焦点管理器与 IME 桥输入中文；
//   - 用指针事件点击按钮（不是直接调 OnClick）；
//   - 提交后由 Base.SetState 机制自动重建页面。
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/widgets"
)

func main() {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(900, 620)

	form := widgets.NewDemoForm()
	c.MountWidget(form)
	if _, err := c.FrameOnce(); err != nil {
		fmt.Fprintln(os.Stderr, "首帧失败:", err)
		os.Exit(1)
	}

	// 中文输入：焦点 → IME 组合 → 上屏。与真实窗口走的是同一个桥。
	c.Focus().SetCurrent(form.Name.FocusNode())
	typeChinese(c, "张伟")
	c.Focus().SetCurrent(form.City.FocusNode())
	typeChinese(c, "上海")

	// 选专业版。
	form.PlanB.Checked = true
	form.PlanB.SetState(func() {
		form.PlanA.Checked = false
	})

	// 真正的点击：按按钮当前布局位置投递 Down/Up。
	click(c, form.Submit.Node())
	_, _ = c.FrameOnce()

	fmt.Println("表单状态:", form.Status)
	fmt.Println("姓名:", form.Name.Value, "城市:", form.City.Value)
	fmt.Println("方案:", map[bool]string{true: "专业版", false: "标准版"}[form.PlanB.Checked])
	fmt.Println("音量:", int(form.Level.Value))
}

func typeChinese(c *wg.Container, s string) {
	c.IME().Handle(host.IMFEvent{Kind: host.IMFStartComposition})
	c.IME().Handle(host.IMFEvent{Kind: host.IMFUpdateComposition, Text: s, Cursor: int32(len([]rune(s)))})
	c.IME().Handle(host.IMFEvent{Kind: host.IMFResult, Text: s})
	c.IME().Handle(host.IMFEvent{Kind: host.IMFEndComposition})
}

// click 按节点当前布局位置投递一次真实指针点击。
func click(c *wg.Container, n *scene.Node) {
	if n == nil {
		return
	}
	b := n.BoundsInRoot()
	p := geom.Pt(b.Min.X+b.W()/2, b.Min.Y+b.H()/2)
	t := time.Now()
	c.PushEvent(&input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer,
		Pos: p, PosRoot: p, Time: t})
	c.PushEvent(&input.PointerEvent{Kind: input.PointerUp, Pointer: input.MousePointer,
		Pos: p, PosRoot: p, Time: t.Add(16 * time.Millisecond)})
}
