// TestDocPageScrollInteract 锁定窗体级滚动（doc 页 widgets.Scroll Both）与
// 交互的三条回归：
//
//	症状1：聚焦触发整页重建后，滚动条/输入框/卡片必须完整重画（合成器
//	       继承裁剪链曾被 walk 的 restore 清零，层 blend 被裁成空）；
//	症状2：demo 控件实例跨重建复用，点击聚焦→重建→焦点不丢、键入生效；
//	症状3：prefix 图标贴在输入框内（重建后几何不漂移）。
package wg_test

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/geom"

	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 高保真复现：真实 doc.Page（examples/doc 框架）+ input 示例同款 demo 结构，
// headless 像素断言三症状。

func frame3(t *testing.T, c *wg.Container) {
	t.Helper()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
}

func clickPt(c *wg.Container, pt geom.Point) {
	now := time.Now()
	c.PushEvent(&input.PointerEvent{Kind: input.PointerDown, Pointer: input.MousePointer, Pos: pt, PosRoot: pt, Time: now})
	c.PushEvent(&input.PointerEvent{Kind: input.PointerUp, Pointer: input.MousePointer, Pos: pt, PosRoot: pt, Time: now})
}

func hoverPt(c *wg.Container, pt geom.Point) {
	c.PushEvent(&input.PointerEvent{Kind: input.PointerHover, Pointer: input.MousePointer, Pos: pt, PosRoot: pt, Time: time.Now()})
}

func nonBgRect(c *wg.Container, r geom.Rect) int {
	b := c.Backing()
	if b == nil {
		return -1
	}
	px := b.Pix()
	w := b.W()
	n := 0
	for y := int(r.Min.Y); y < int(r.Max.Y); y += 2 {
		if y < 0 || y >= b.H() {
			continue
		}
		for x := int(r.Min.X); x < int(r.Max.X); x += 2 {
			if x < 0 || x >= w {
				continue
			}
			if px[y*w+x] != 0xFFFFFFFF {
				n++
			}
		}
	}
	return n
}

func center(n *scene.Node) geom.Point {
	b := n.BoundsInRoot()
	return geom.Pt((b.Min.X+b.Max.X)/2, (b.Min.Y+b.Max.Y)/2)
}

func findNode(root *scene.Node, pred func(*scene.Node) bool) *scene.Node {
	var out *scene.Node
	root.Walk(func(n *scene.Node) bool {
		if out != nil {
			return false // 已命中：剪掉后续所有分支（Walk 只剪子树，不终止遍历）
		}
		if pred(n) {
			out = n
			return false
		}
		return true
	})
	return out
}

func TestDocPageScrollInteract(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1080, 760)

	// 与 examples/input 相同的 demo 结构：交互控件实例在闭包外持有，
	// 闭包内只做 instance.Build(ctx)（重建产新节点、状态挂实例）。
	var (
		inp         = widgets.NewInput("试试输入中文，或用输入法组合")
		status      = widgets.NewTextView("当前内容：")
		edit        = widgets.NewInput("试试 Backspace / Delete / Ctrl+A")
		statusInput = widgets.NewInput("error 状态")
		prefixed    = widgets.NewInput("站内搜索")
	)
	inp.Width = 420
	status.Width = 520
	edit.Width = 460
	statusInput.Width, statusInput.Status = 260, widgets.InputStatusError
	prefixed.Width, prefixed.AllowClear = 300, true
	prefixed.Prefix = &widgets.Icon{Name: "search", Size: 14}

	c.MountWidget(doc.Page(doc.Meta{
		ID: "input", Title: "Input 输入框", Group: doc.GroupControl,
		What: []string{"复现用页面"},
		Demos: []doc.Demo{
			{Title: "受控输入与 IME", Build: func(ctx *ui.BuildContext) ui.Widget {
				clear := widgets.NewButton("清空")
				col := scene.Flex(scene.Vertical, 10)
				col.Add(inp.Build(ctx))
				row := scene.Flex(scene.Horizontal, 8)
				row.Add(status.Build(ctx))
				row.Add(clear.Build(ctx))
				col.Add(row)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "三种 variant", Build: func(ctx *ui.BuildContext) ui.Widget {
				o := widgets.NewInput("outlined")
				o.Width = 260
				f := widgets.NewInput("filled")
				f.Width, f.Variant = 260, widgets.InputFilled
				bl := widgets.NewInput("borderless")
				bl.Width, bl.Variant = 260, widgets.InputBorderless
				return widgets.Column{Gap: 8, Children: []ui.Widget{o, f, bl}}
			}},
			{Title: "status 与尺寸", Build: func(ctx *ui.BuildContext) ui.Widget {
				warn := widgets.NewInput("warning 状态")
				warn.Width, warn.Status = 260, widgets.InputStatusWarning
				sm := widgets.NewInput("small")
				sm.Width, sm.Size = 260, ui.SizeSmall
				lg := widgets.NewInput("large")
				lg.Width, lg.Size = 260, ui.SizeLarge
				return widgets.Column{Gap: 8, Children: []ui.Widget{statusInput, warn, sm, lg}}
			}},
			{Title: "前后缀与清除", Build: func(ctx *ui.BuildContext) ui.Widget {
				suffixed := widgets.NewInput("金额")
				suffixed.Width = 300
				suffixed.Suffix = &widgets.Icon{Name: "info-circle", Size: 14}
				return widgets.Column{Gap: 8, Children: []ui.Widget{prefixed, suffixed}}
			}},
			{Title: "文本编辑按键", Build: func(ctx *ui.BuildContext) ui.Widget {
				return edit
			}},
		},
		API: []doc.APIRow{
			{Name: "NewInput", Type: "func(string) *Input", Desc: "构造"},
			{Name: "Variant", Type: "widgets.InputVariant", Default: "InputOutlined", Desc: "三态"},
			{Name: "Status", Type: "widgets.InputStatus", Default: "StatusOK", Desc: "错误/警告"},
			{Name: "Size", Type: "ui.Size", Default: "SizeMedium", Desc: "24/32/40"},
			{Name: "Prefix/Suffix", Type: "ui.Widget", Desc: "框内装饰"},
			{Name: "AllowClear", Type: "bool", Default: "false", Desc: "清除图标"},
		},
	}))
	frame3(t, c)

	inputNodes := func() []*scene.Node {
		var out []*scene.Node
		c.Scene().Walk(func(n *scene.Node) bool {
			if n.HitSpec().Label == "input" {
				out = append(out, n)
			}
			return true
		})
		return out
	}
	barRect := geom.RectXYWH(1030, 24, 10, 400)

	// ── 阶段A：hover → 点击第一个（受控输入）输入框（症状1） ──
	nodes := inputNodes()
	if len(nodes) == 0 {
		t.Fatal("找不到输入框")
	}
	a := nodes[0]
	ac := center(a)
	scanThumb := func(tag string) {
		b := c.Backing()
		px, st := b.Bytes(), b.Stride()
		start, end, found := -1, -1, false
		for y := 26; y < 726; y++ {
			r := px[y*st+1034*4+2]
			dark := r < 200
			if dark && !found {
				start, found = y, true
			}
			if dark {
				end = y
			}
		}
		t.Logf("[thumb %s] %v..%v", tag, start, end)
	}
	scanThumb("initial")
	hoverPt(c, ac)
	frame3(t, c)
	scanThumb("after-hover")
	clickPt(c, ac)
	frame3(t, c)
	scanThumb("after-focus")
	frame3(t, c)
	scanThumb("focus+1")
	// 聚焦触发整页重建：重建帧必须完整重画（聚焦环/滚动条/卡片一个不少）。
	if n := nonBgRect(c, barRect); n == 0 {
		t.Fatalf("[症状1] 聚焦重建后滚动条区域全白")
	}
	if n := nonBgRect(c, geom.RectXYWH(40, 24, 500, 120)); n == 0 {
		t.Fatalf("[症状1] 聚焦重建后页面顶部全白")
	}
	// thumb 像素完整性：track(1030..1038) 的 x=1034 列上，thumb 段必须从
	// 轨道顶(y≈26)开始（聚焦重建不得裁掉 thumb 顶部）。
	if img := c.Backing(); img != nil {
		px := img.Bytes()

		darkStart, darkEnd, found := -1, -1, false
		for y := 26; y < 726; y++ {
			r := px[y*img.Stride()+1034*4+2] // R 通道（BGRA）
			dark := r < 200
			if dark && !found {
				darkStart, found = y, true
			}
			if dark {
				darkEnd = y
			}
		}
		t.Logf("[thumb] dark=%v..%v", darkStart, darkEnd)
		if found && darkStart > 40 {
			t.Fatalf("[症状1] thumb 顶部被裁：dark 从 %d 开始（应≈26）", darkStart)
		}
	}
	a2 := findNode(c.Scene(), func(n *scene.Node) bool { return n.HitSpec().Label == "input" })
	if a2 == nil {
		t.Fatal("[症状1] 聚焦重建后场景里没有输入框")
	}
	if n := nonBgRect(c, a2.BoundsInRoot()); n == 0 {
		t.Fatalf("[症状1] 聚焦重建后输入框区域全白：bounds=%v", a2.BoundsInRoot())
	}

	// ── 阶段B：滚到 status/尺寸 demo，点击其输入框并键入（症状2） ──
	c.PushEvent(&input.WheelEvent{Pos: geom.Pt(500, 400), Delta: geom.Pt(0, 500), Time: time.Now()})
	frame3(t, c)
	if statusInput == nil {
		t.Fatal("[B] demo 输入实例未被捕获（框架缓存未生效）")
	}
	sb := statusInput.Node().BoundsInRoot()
	if sb.Min.Y < 0 || sb.Max.Y > 700 {
		t.Fatalf("[B] status 输入框不在可视区：%v", sb)
	}
	clickPt(c, center(statusInput.Node()))
	frame3(t, c)
	frame3(t, c)
	if !statusInput.FocusNode().Focused() {
		t.Fatalf("[症状2] 点击 demo 输入框并重建后焦点丢失")
	}
	c.PushEvent(&input.KeyEvent{Kind: input.KeyDown, Text: "x", Time: time.Now()})
	frame3(t, c)
	if got := statusInput.Editor().Value().Text; got != "x" {
		t.Fatalf("[症状2] 键入结果 = %q（应 x）", got)
	}

	// ── 阶段C：滚到前后缀 demo，检查图标相对输入框的位置（症状3） ──
	// 前后缀 demo 在内容 y≈900；累计滚 500+420 到可视区。
	c.PushEvent(&input.WheelEvent{Pos: geom.Pt(500, 400), Delta: geom.Pt(0, 420), Time: time.Now()})
	frame3(t, c)
	title := findNode(c.Scene(), func(n *scene.Node) bool {
		return n.Text() != nil && n.Text().Src == "站内搜索"
	})
	if title == nil {
		t.Fatal("[C] 找不到前后缀 demo 的输入框（站内搜索）")
	}
	inBox := title.Parent() // 占位文本的父 = 输入框 Box
	if inBox == nil {
		t.Fatal("[C] 占位文本没有父节点")
	}
	ib := inBox.BoundsInRoot()
	var iconN *scene.Node
	inBox.Walk(func(n *scene.Node) bool {
		if n != inBox && n.Kind() == scene.KindCustomPaint {
			iconN = n
			return false
		}
		return true
	})
	if iconN == nil {
		t.Fatal("[症状3] 输入框内找不到 prefix 图标节点")
	}
	gb := iconN.BoundsInRoot()
	t.Logf("[C] 输入框 bounds=%v prefix 图标 bounds=%v", ib, gb)
	if gb.Min.X < ib.Min.X-2 || gb.Max.X > ib.Max.X+2 || gb.Min.Y < ib.Min.Y-2 || gb.Max.Y > ib.Max.Y+2 {
		t.Fatalf("[症状3] prefix 图标跑到输入框外：icon=%v input=%v", gb, ib)
	}
}
