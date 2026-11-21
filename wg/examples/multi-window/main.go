package main

// multi-window 是多窗口能力的真窗示例：**一个进程、多个 wg 窗口**，
// 每窗独立的帧泵/事件分发/呈现（Container.Start），共用一条平台消息循环
// （主窗 Container.Run 里的 Application.Run）。
//
// 窗口构成：
//   - 主窗：说明页 + 「打开交互副窗」按钮（动态建窗演示）
//   - 动画副窗（启动即开）：持续帧窗（60Hz 单 outstanding），验证多窗并发渲染
//   - 交互副窗（按钮动态开）：可点组件页，验证每窗事件隔离
//
// 生命周期：副窗关闭 → OnClose 停本窗帧泵并释放呈现层，其余窗不受影响；
// 主窗关闭 → Application.Run 返回 → 程序退出（未关副窗随进程回收）。
//
// 性能语义（C14/C11 每窗独立）：静止窗零帧零闹钟（帧泵按需武装，空闲成本
// 为零）；动画窗 60Hz；N 窗的 FrameOnce 各自 goroutine 可并行（无全局锁）。
//
// 真窗验证记录（WG_DEBUG 时代的一次探针实测，供手工回归对照）：
// 主窗按钮 tap → openInteractive 动态开窗 ✓；交互副窗计数按钮 tap 且
// OnClick 只落本窗 ✓；动画窗 painted 持续增长、主窗静止零帧 ✓；
// resize 后 backing 重建 ✓。当时"点击无反应"的根因是**自动化点击选错了
// 节点**：Node.Walk 的 return false 只剪枝当前子树、父层循环会继续走
// 兄弟，"找第一个按钮"写成提前 return false 拿到的是最后一个按钮——
// 恰好是旁边那颗禁用按钮，而禁用控件的 GestureDetector 没有 OnTap，
// down 消费、up 拒绝，tap 永不触发（这是设计行为，不是引擎缺陷）。

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

type mwForm struct {
	lcl.TEngForm
	container *wg.Container
}

var openCount int

func main() {
	if hasArg("--selfcheck") {
		selfcheck()
		return
	}
	// ── 审计模式：跳过副窗启动 + 定时自动退出（回归不挂——主窗
	// Run 永久阻塞消息循环且此前无 --audit 解析，全量回归会卡死
	// 在本窗；窗口自动关闭为用户明确要求）。
	auditMode := hasArg("--audit")
	if auditMode {
		time.AfterFunc(10*time.Second, func() {
			lcl.RunOnMainThreadAsync(func(uint32) {
				lcl.Application.Terminate()
			})
		})
	}
	lcl.Init()
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	// ── 主窗：说明 + 动态开窗按钮 ────────────────────────────────
	var main mwForm
	lcl.Application.NewForm(&main)
	setupWindow(&main, "wg 多窗口 · 主窗（关我 = 退出）", 1080, 760)
	main.container = wg.NewContainer(&main, wg.DefaultOptions())
	main.container.MountWidget(mainPage())

	// ── 动画副窗（启动即开）：持续帧窗（audit 模式跳过——副窗
	// 生命周期不进启发式审计，留着即窗残留）─────────────────────
	var anim mwForm
	skipAnim := auditMode
	if !skipAnim {
		lcl.Application.NewForm(&anim)
		setupWindow(&anim, "wg 多窗口 · 动画副窗（60Hz）", 700, 520)
		anim.container = wg.NewContainer(&anim, wg.DefaultOptions())
		anim.container.SetAnimating(true) // 动态组件驱动渲染：60Hz 标脏+续帧（按需渲染）
		// 副窗顺序：**先 Start 后 MountWidget**——自驱动动画页在 Build 里
		// 用 sched.RequestFrameIn 续帧，Build 必须发生在帧泵就位之后
		// （Mount 先于 Start 时首帧 Build 的 Scheduler 为 nil，续帧被帧尾
		// deferred 吞掉，动画只出 1 帧就停）。
		if err := anim.container.Start(); err != nil {
			fmt.Fprintln(os.Stderr, "动画副窗启动失败:", err)
		}
		anim.container.MountWidget(animPage())
		anim.Show()
	}

	// ── 主窗进消息循环（副窗的消息经各自的 binder 到达）──────────
	if err := main.container.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "主窗启动失败:", err)
	}
}

// openInteractive 动态打开一个交互副窗（主窗按钮回调，已在主线程）。
func openInteractive() {
	if hasArg("--audit") {
		return // 审计模式：开窗按钮点击不真开副窗（防窗残留）
	}
	var f mwForm
	lcl.Application.NewForm(&f)
	openCount++
	setupWindow(&f, fmt.Sprintf("wg 多窗口 · 交互副窗 #%d", openCount), 720, 560)
	f.container = wg.NewContainer(&f, wg.DefaultOptions())
	// 同动画副窗：Start 先于 MountWidget（顺序约定见 main）。
	if err := f.container.Start(); err != nil {
		fmt.Fprintln(os.Stderr, "交互副窗启动失败:", err)
	}
	f.container.MountWidget(interactivePage())
	f.Show()
}

func setupWindow(f *mwForm, title string, w, h int32) {
	f.SetCaption(title)
	f.SetPosition(types.PoScreenCenter)
	f.SetWidth(w)
	f.SetHeight(h)
	// 场景自持离屏缓冲，关 LCL 层二次缓冲（避免整窗多余拷贝）。
	f.SetDoubleBuffered(false)
}

// ── 页面 ────────────────────────────────────────────────────────────

// mainPage 主窗内容：多窗口说明 + 动态开窗按钮。
func mainPage() ui.Widget {
	open := widgets.NewButton("打开交互副窗（动态建窗）")
	open.Type = widgets.TypePrimary
	open.OnClick = openInteractive
	hint := widgets.NewButton("动画副窗已随程序启动（60Hz 持续帧）")
	hint.Type = widgets.TypeText
	hint.Disabled = true

	return page(widgets.Column{Gap: 14, Children: []ui.Widget{
		widgets.Title{Text: "wg 多窗口", Level: 2},
		&widgets.Label{Text: "一个进程 · 多个 wg 窗口 · 每窗独立的帧泵/事件分发/呈现"},
		&widgets.Divider{Length: 560},
		widgets.Title{Text: "窗口构成", Level: 4},
		&widgets.Label{Text: "· 主窗（本窗）：关我 = 程序退出"},
		&widgets.Label{Text: "· 动画副窗：60Hz 持续帧窗，验证多窗并发渲染与空闲隔离"},
		&widgets.Label{Text: "· 交互副窗：点下面的按钮动态创建，每窗事件互不干扰"},
		widgets.Title{Text: "性能语义", Level: 4},
		&widgets.Label{Text: "· 帧泵按需武装：静止窗零帧零闹钟（C14），空闲成本为零"},
		&widgets.Label{Text: "· 动画窗 60Hz 单 outstanding（C11）；N 窗的合成可并行（无全局锁）"},
		&widgets.Label{Text: "· 关闭副窗即停该窗帧泵并释放呈现层"},
		&widgets.Row{Gap: 12, Children: []ui.Widget{open, hint}},
	}})
}

// interactivePage 交互副窗内容：可点组件，验证每窗事件独立。
func interactivePage() ui.Widget {
	counter := widgets.NewButton("已点击 0 次")
	sw := widgets.NewSwitch(true)
	sl := widgets.NewSlider(0, 100, 40)
	sl.Width = 320
	n := 0
	counter.OnClick = func() {
		n++
		counter.Text = fmt.Sprintf("已点击 %d 次", n)
		counter.MarkDirty()
	}
	return page(widgets.Column{Gap: 12, Children: []ui.Widget{
		widgets.Title{Text: "交互副窗", Level: 3},
		&widgets.Label{Text: "本窗的点击/切换/拖动只影响本窗；主窗与动画窗不受干扰"},
		&widgets.Divider{Length: 480},
		counter,
		&widgets.Row{Gap: 16, Children: []ui.Widget{sw, &widgets.Label{Text: "开关"}}},
		sl,
	}})
}

// ── 动画页：持续帧窗（60Hz），多窗并发渲染的压测样张 ────────────────

// animPage 是动画副窗的整页（页面 widget 自驱动续帧）。
func animPage() ui.Widget {
	return &animBoard{start: time.Now()}
}

// animBoard 持续重绘：旋转弧 + 滚动波形 + 帧号，验证多窗下帧泵独立性。
type animBoard struct {
	ui.Base
	start time.Time
}

func (a *animBoard) Build(ctx *ui.BuildContext) *scene.Node {
	th := ctx.ThemeOr()
	board := scene.CustomPaintNode(func(cv render.Canvas, sz geom.Size) {
		elapsed := time.Since(a.start)
		phase := float32(math.Mod(elapsed.Seconds()*0.8, 1))
		// 旋转弧
		cv.FillRect(geom.RectXYWH(0, 0, sz.W, sz.H), render.Fill(th.Color.SurfaceSunken))
		arc := render.NewPath().AddArc(
			geom.RectXYWH(sz.W/2-60, 30, 120, 120), phase*6.283, phase*6.283+4.4)
		cv.StrokePath(arc, render.Stroke(th.Color.Primary, 10).WithCap(render.CapRound))
		// 滚动波形
		for i := 0; i < 60; i++ {
			t := float32(i)/6 + float32(elapsed.Seconds())*3
			h := 40 + 30*float32(math.Sin(float64(t)))
			cv.FillRect(geom.RectXYWH(float32(i)*float32(sz.W-60)/60+30, sz.H-40-h, 6, h),
				render.Fill(th.Color.Success))
		}
		// 帧号文字
		lay := text.LayoutText(fmt.Sprintf("动画副窗 · 已运行 %v · 帧持续更新", elapsed.Round(time.Second)),
			text.DefaultFont(th.Font.SizeSM), text.Options{})
		cv.DrawText(lay, geom.Pt(16, 12), th.Color.Text)
	}).SetSize2(640, 420)
	col := scene.Flex(scene.Vertical, 12).SetCrossAlign(scene.CrossCenter)
	col.Add(board)
	return col
}

// ── 工具 ────────────────────────────────────────────────────────────

// selfcheck 以 headless 方式验证多窗口：两个容器各自出帧、各自有像素。
func selfcheck() {
	// 多窗口自检：headless 双容器各自出帧（交互页 + 动画页），
	// 输出**单行聚合 JSON**（与枚举门禁的解析约定一致）。
	totalNonBg, totalDegraded, allOK := 0, 0, true
	for i, w := range []int{720, 640} {
		c := wg.NewHeadless(wg.DefaultOptions())
		c.Resize(w, 520)
		if i == 0 {
			c.MountWidget(interactivePage())
		} else {
			c.MountWidget(animPage())
		}
		st, err := c.FrameOnce()
		nonBg := 0
		if b := c.Backing(); b != nil {
			px := b.Pix()
			for j := 0; j < len(px); j += 16 {
				if px[j] != 0xFFFFFFFF {
					nonBg++
				}
			}
		}
		totalNonBg += nonBg
		totalDegraded += st.Degraded
		if err != nil || !st.Presented || nonBg < 100 {
			allOK = false
		}
	}
	fmt.Printf("{\"id\":\"multi-window\",\"title\":\"多窗口\",\"group\":\"%s\",\"demos\":1,\"apiRows\":0,"+
		"\"presented\":%v,\"nodes\":0,\"degraded\":%d,\"nonBg\":%d,\"ok\":%v}\n",
		doc.GroupLayout, allOK, totalDegraded, totalNonBg, allOK)
	if !allOK {
		os.Exit(1)
	}
}

// hasArg 判断命令行参数（--selfcheck 用）。
func hasArg(name string) bool {
	for _, a := range os.Args[1:] {
		if a == name {
			return true
		}
	}
	return false
}

// ── 根包装 ──────────────────────────────────────────────────────────

// mwPage 是多窗页面的根 widget 包装：嵌 ui.Base 使根具备 Dirty/MarkDirty
// 能力（rebuildDirtyWidget 依赖它重建；widgets.Column 等值类型根没有）。
type mwPage struct {
	ui.Base
	child ui.Widget
}

func page(w ui.Widget) ui.Widget { return &mwPage{child: w} }

// Build 实现 ui.Widget。
func (m *mwPage) Build(ctx *ui.BuildContext) *scene.Node {
	ctx.BindState(m)
	return m.child.Build(ctx)
}
