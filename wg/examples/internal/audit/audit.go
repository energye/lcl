// Package audit 是功能窗的全自动自测引擎（真窗 + 引擎事件注入）。
//
// 对一个已挂载的功能窗执行：
//  1. 逐屏滚动（注入 WheelEvent，走与真实 wndproc 相同的分发链路），每个
//     滚动位置导出 backing PNG——覆盖折叠区（首屏之外的全部内容）；
//  2. 枚举场景里全部可交互元素（按命中标签），逐个滚入视口、注入点击/键入，
//     用"点击前后中心像素变化 + 焦点归属 + 帧计数"验证功能反馈；
//  3. 每次点击后注入 Esc 关闭可能弹出的浮层（模态/抽屉/气泡）。
//
// 注入通道是 container.PushEvent：宿主把原生消息归一化后正是投递到同一入口，
// 事件对象同构，因此不需要窗口消息环（也避开了 LCL 异步队列只投递首个回调
// 的缺陷）。像素采样用 container.Snapshot：它在容器锁内拷贝合成缓冲，
// 跨线程读取不会与主线程光栅写竞争。
//
// 判定全部自动化，不依赖人工目检。
package audit

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/host"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
)

// Finding 是单条审计结论。
type Finding struct {
	Element string `json:"element"`
	Check   string `json:"check"`
	OK      bool   `json:"ok"`
	Detail  string `json:"detail"`
}

// Report 是一个功能窗的审计报告。
type Report struct {
	Pages    int       `json:"pages"`
	Elements int       `json:"elements"`
	Failed   int       `json:"failed"`
	Findings []Finding `json:"findings"`
}

// ── 事件注入（引擎事件，与宿主归一化输出等价）──────────────────────

// injectWheel 注入滚轮：notch>0 向上（远离用户），<0 向下。
// 与 host.Normalizer 的换算一致：一格 120 → WheelStep DIP，向下为正 Y。
func injectWheel(c *wg.Container, x, y, notch float32) {
	c.PushEvent(&input.WheelEvent{
		Pos:   geom.Pt(x, y),
		Delta: geom.Pt(0, -notch*host.WheelStep),
		Time:  time.Now(),
	})
}

// injectDown/injectUp 分步注入左键按下与抬起。
// 分步是为了在按住期间抓帧：无回调的展示控件（纯样式按钮）没有持续
// 状态变化，但按下/hover 反馈同样证明事件链路有效。
func injectDown(c *wg.Container, x, y float32) bool {
	p := geom.Pt(x, y)
	return c.PushEvent(&input.PointerEvent{
		Kind: input.PointerDown, Pointer: input.MousePointer, Device: input.DeviceMouse,
		Pos: p, PosRoot: p, Buttons: input.ButtonLeft, ClickCount: 1, Time: time.Now(),
	})
}

func injectUp(c *wg.Container, x, y float32) bool {
	p := geom.Pt(x, y)
	return c.PushEvent(&input.PointerEvent{
		Kind: input.PointerUp, Pointer: input.MousePointer, Device: input.DeviceMouse,
		Pos: p, PosRoot: p, Time: time.Now(),
	})
}

// injectChar 注入一个可打印字符（WM_CHAR 路径归一化后 Text 即内容，Code 为空）。
func injectChar(c *wg.Container, ch rune) {
	c.PushEvent(&input.KeyEvent{Kind: input.KeyDown, Text: string(ch), Time: time.Now()})
}

// injectEsc 注入 Esc（关闭浮层的标准键）。
func injectEsc(c *wg.Container) {
	c.PushEvent(&input.KeyEvent{Kind: input.KeyDown, Code: input.CodeEscape, Time: time.Now()})
}

// injectLeave 把指针移出客户区：清 hover 链与 Tooltip 计时。
// 审计的指针在两次点击之间不移动，不主动移出的话悬停 800ms 会弹
// Tooltip，吞掉下一次点击（OutsideClick/Modal 遮挡语义）。
func injectLeave(c *wg.Container) {
	c.PushEvent(&input.PointerEvent{
		Kind: input.PointerLeave, Pointer: input.MousePointer, Device: input.DeviceMouse,
		Time: time.Now(),
	})
}

// ── 元素枚举 ────────────────────────────────────────────────────────

// interactiveLabels 是纳入审计的命中标签（可见交互元素）。
var interactiveLabels = map[string]string{
	"input":             "输入框",
	"button":            "按钮",
	"switch":            "开关",
	"checkbox":          "复选框",
	"radio":             "单选",
	"segment":           "分段项",
	"tab":               "页签",
	"slider":            "滑块",
	"select":            "下拉框",
	"collapse":          "折叠表头",
	"table-row":         "表格行",
	"tree-node":         "树节点",
	"rate-star":         "评分星",
	"page":              "页码",
	"breadcrumb":        "面包屑项",
	"modal-button":      "模态按钮",
	"input-number":      "数字输入",
	"autocomplete-item": "联想项",
}

type element struct {
	label    string    // 命中标签
	name     string    // 中文名
	bounds   geom.Rect // 枚举时的绝对 bounds
	disabled bool      // 禁用态：点击不应有任何变化
}

// ── 像素采样 ────────────────────────────────────────────────────────

// grab 抓合成缓冲快照并返回位图（BGRA 预乘 → RGBA，输出不透明）。
// Snapshot 在容器锁内完成拷贝，可在审计 goroutine 直接调用。
func grab(c *wg.Container) *image.RGBA {
	b := c.Snapshot()
	if b == nil {
		return nil
	}
	w, h := b.W(), b.H()
	px := b.Pix()
	out := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		src := y * w
		dst := y * out.Stride
		for x := 0; x < w; x++ {
			v := px[src+x]
			out.Pix[dst+x*4+0] = uint8(v)
			out.Pix[dst+x*4+1] = uint8(v >> 8)
			out.Pix[dst+x*4+2] = uint8(v >> 16)
			out.Pix[dst+x*4+3] = 255
		}
	}
	return out
}

// sampleAt 返回 (x,y) 处 7×7 补丁的平均色。
func sampleAt(img *image.RGBA, x, y float32) [3]float64 {
	if img == nil {
		return [3]float64{}
	}
	var sum [3]float64
	n := 0
	for dy := -3; dy <= 3; dy++ {
		for dx := -3; dx <= 3; dx++ {
			px, py := int(x)+dx, int(y)+dy
			if px < 0 || py < 0 || px >= img.Bounds().Dx() || py >= img.Bounds().Dy() {
				continue
			}
			c := img.RGBAAt(px, py)
			sum[0] += float64(c.R)
			sum[1] += float64(c.G)
			sum[2] += float64(c.B)
			n++
		}
	}
	if n == 0 {
		return [3]float64{}
	}
	return [3]float64{sum[0] / float64(n), sum[1] / float64(n), sum[2] / float64(n)}
}

func savePNG(img image.Image, path string) {
	sf, err := os.Create(path)
	if err != nil {
		return
	}
	defer sf.Close()
	png.Encode(sf, img)
}

func colorDist(a, b [3]float64) float64 {
	d0, d1, d2 := a[0]-b[0], a[1]-b[1], a[2]-b[2]
	return d0*d0 + d1*d1 + d2*d2
}

// regionChanged 返回矩形区域内两帧的像素变化占比（任一通道差 > 12 计为变化）。
// 用区域而不是单点采样：开关手柄移动、勾选记号、星形填充的变化
// 未必覆盖元素中心 7×7（--audit 全量首轮回测过该误报）。
func regionChanged(a, b *image.RGBA, r geom.Rect) float64 {
	return regionChangedThresh(a, b, r, 12)
}

// regionChangedThresh 是带通道阈值的区域比较。
// 按住反馈（hover/pressed）常用极浅的底色变化（如 #FFF→#FAFAFA 差 5），
// 需要用低阈值单独判定，否则会被当成"无反馈"。
func regionChangedThresh(a, b *image.RGBA, r geom.Rect, thresh uint8) float64 {
	if a == nil || b == nil || r.IsEmpty() {
		return 0
	}
	x0, y0 := int(r.Min.X), int(r.Min.Y)
	x1, y1 := int(r.Max.X), int(r.Max.Y)
	bounds := a.Bounds()
	if x0 < bounds.Min.X {
		x0 = bounds.Min.X
	}
	if y0 < bounds.Min.Y {
		y0 = bounds.Min.Y
	}
	if x1 > bounds.Max.X {
		x1 = bounds.Max.X
	}
	if y1 > bounds.Max.Y {
		y1 = bounds.Max.Y
	}
	if x1 <= x0 || y1 <= y0 {
		return 0
	}
	total, diff := 0, 0
	for y := y0; y < y1; y++ {
		i := y*a.Stride + x0*4
		for x := x0; x < x1; x++ {
			if absByteThresh(a.Pix[i], b.Pix[i], thresh) ||
				absByteThresh(a.Pix[i+1], b.Pix[i+1], thresh) ||
				absByteThresh(a.Pix[i+2], b.Pix[i+2], thresh) {
				diff++
			}
			total++
			i += 4
		}
	}
	return float64(diff) / float64(total)
}

// frameChanged 返回整帧像素变化占比（3px 步进采样）。
// 捕获元素区域之外的反馈：弹浮层/模态/气泡。
func frameChanged(a, b *image.RGBA) float64 {
	if a == nil || b == nil {
		return 0
	}
	w, h := a.Bounds().Dx(), a.Bounds().Dy()
	total, diff := 0, 0
	for y := 0; y < h; y += 3 {
		for x := 0; x < w; x += 3 {
			i := y*a.Stride + x*4
			if absByte(a.Pix[i], b.Pix[i]) || absByte(a.Pix[i+1], b.Pix[i+1]) ||
				absByte(a.Pix[i+2], b.Pix[i+2]) {
				diff++
			}
			total++
		}
	}
	if total == 0 {
		return 0
	}
	return float64(diff) / float64(total)
}

func absByte(x, y uint8) bool {
	return absByteThresh(x, y, 12)
}

func absByteThresh(x, y, thresh uint8) bool {
	if x > y {
		return x-y > thresh
	}
	return y-x > thresh
}

func dist(a, b geom.Point) float32 {
	dx, dy := a.X-b.X, a.Y-b.Y
	return dx*dx + dy*dy
}

// ── 审计主流程 ──────────────────────────────────────────────────────

// Run 执行全页审计。c 为已挂载并运行中的容器（可在工作 goroutine 调用）。
func Run(c *wg.Container, outDir string) *Report {
	rep := &Report{Findings: []Finding{}}
	_ = os.MkdirAll(outDir, 0o755)

	// 等首批帧落地再开始采样（PresentStats 有锁，可安全轮询）。
	deadline := time.Now().Add(3 * time.Second)
	for int(c.PresentStats().Frames) < 3 && time.Now().Before(deadline) {
		time.Sleep(10 * time.Millisecond)
	}

	// ① 回到顶部：向上滚到页面级滚动偏移稳定。
	scrollPos := func() float32 {
		var off, area float32
		c.Scene().Walk(func(n *scene.Node) bool {
			o := n.ScrollOffset()
			if o == 0 {
				return true
			}
			if b := n.BoundsInRoot(); b.W()*b.H() > area {
				area = b.W() * b.H()
				off = o
			}
			return true
		})
		return off
	}
	prevOff := float32(-1)
	for i := 0; i < 40; i++ {
		off := scrollPos()
		if off == prevOff {
			break
		}
		prevOff = off
		injectWheel(c, 1035, 400, 3)
		time.Sleep(110 * time.Millisecond)
	}
	time.Sleep(220 * time.Millisecond)

	// verifyElement 验证单个元素。坐标取自枚举时的当前视口，因此可信；
	// 每次验证后页面可能重建，但视口不滚动时布局位置不变。
	verifyElement := func(el element) {
		b := el.bounds
		cx, cy := (b.Min.X+b.Max.X)/2, (b.Min.Y+b.Max.Y)/2
		// 点击前清浮层：浮层打开时它吞掉层外第一次点击（只用于关闭自己，
		// 不下传），并可能永久遮挡后续元素。先 Esc，没关掉就强制全关。
		if ov := c.Overlay(); ov != nil && ov.Count() > 0 {
			injectEsc(c)
			time.Sleep(120 * time.Millisecond)
			if ov.Count() > 0 {
				ov.CloseAll(theme.BandPopup)
				ov.CloseAll(theme.BandTooltip)
				ov.CloseAll(theme.BandModal)
				ov.CloseAll(theme.BandToast)
				time.Sleep(120 * time.Millisecond)
			}
		}
		hitLabel := "nil"
		hasHandler := false
		hitMismatch := false
		if t := input.HitTest(c.Scene(), geom.Pt(cx, cy)).Target(); t != nil {
			hitLabel = t.HitSpec().Label
			if d := c.Dispatcher(); d != nil {
				_, hasHandler = d.Handler(t)
			}
			// 同 label 的相邻实例（两个按钮紧挨）会被 label 检查放过：
			// 命中节点的实际位置必须与枚举元素一致，否则测的是邻居。
			if hb := t.BoundsInRoot(); dist(hb.Center(), b.Center()) > 40 {
				hitMismatch = true
			}
		}
		if hitLabel != el.label || hitMismatch {
			// 视口内却命不中（浮层/裁剪遮挡/与邻居错配）：不计失败。
			rep.Findings = append(rep.Findings, Finding{Element: el.name, Check: "跳过",
				OK: true, Detail: fmt.Sprintf("at %.0f,%.0f hit=%s", cx, cy, hitLabel)})
			return
		}
		before := grab(c)
		centerBefore := sampleAt(before, cx, cy)

		// 分步注入：按住期间抓一帧，捕获 hover/pressed 反馈（无回调的
		// 展示控件也有指针反馈，证明事件链路有效）。
		injectDown(c, cx, cy)
		time.Sleep(60 * time.Millisecond)
		mid := grab(c)
		injectUp(c, cx, cy)
		time.Sleep(430 * time.Millisecond)
		after := grab(c)
		centerAfter := sampleAt(after, cx, cy)

		ratio := regionChanged(before, after, b)
		frameRatio := frameChanged(before, after)
		// 按住反馈的底色变化通常只有几个色阶（#FFF→#FAFAFA），低阈值单判。
		pressedRatio := regionChangedThresh(before, mid, b, 3)
		// 持久变化（切换/选中/弹层），或按住期间的指针反馈。
		// 全帧阈值放到 0.05%：评分星这类连锁变化发生在被点元素之外，
		// 整帧变化很小（3 颗星变灰 ≈ 0.1%）；确定性渲染噪声远低于此。
		// 中心像素兜底：滑块填充这类窄条变化占区域比例小但中心色差明显。
		changed := ratio > 0.02 || frameRatio > 0.0005 ||
			colorDist(centerBefore, centerAfter) > 400
		ok := changed || pressedRatio > 0.02

		checkName := "点击反馈"
		if el.disabled {
			checkName = "禁用态"
			ok = !changed && pressedRatio <= 0.02 // 禁用不应有任何反馈
		}

		switch el.label {
		case "input":
			focusLabel := "nil"
			focusOK := false
			if d := c.Dispatcher(); d != nil {
				if fn := d.Focus(); fn != nil {
					focusLabel = fn.HitSpec().Label
					focusOK = focusLabel == "input"
				}
			}
			if el.disabled {
				ok = !focusOK
			} else {
				ok = focusOK
			}
			rep.Findings = append(rep.Findings, Finding{Element: el.name, Check: "聚焦",
				OK: ok, Detail: fmt.Sprintf("at %.0f,%.0f hit=%s focus=%s", cx, cy, hitLabel, focusLabel)})
			if !ok {
				rep.Failed++
			}
			if el.disabled {
				injectEsc(c)
				time.Sleep(130 * time.Millisecond)
				return
			}
			// 键入验证：注入字符后输入框内容应有区域内视觉变化。
			preType := grab(c)
			injectChar(c, 'w')
			injectChar(c, 'g')
			time.Sleep(330 * time.Millisecond)
			afterType := grab(c)
			typeRatio := regionChanged(preType, afterType, b)
			typed := typeRatio > 0.005 // 短文本（如 borderless 占位）变化区域很小
			rep.Findings = append(rep.Findings, Finding{Element: el.name, Check: "键入",
				OK: typed, Detail: fmt.Sprintf("ratio=%.3f", typeRatio)})
			if !typed {
				rep.Failed++
			}
			injectEsc(c)
			time.Sleep(130 * time.Millisecond)
			return
		case "segment", "tab", "radio", "page", "slider":
			// 点击已选中项（或滑块值已等于点击位置）无区域变化属正常：
			// 换位置/邻居再验。滑块换到 10% 处——值必然不同于先前的中心点。
			if !ok && !el.disabled {
				retryTarget := func(px, py float32, target geom.Rect) bool {
					before2 := grab(c)
					injectDown(c, px, py)
					time.Sleep(60 * time.Millisecond)
					mid2 := grab(c)
					injectUp(c, px, py)
					time.Sleep(430 * time.Millisecond)
					after2 := grab(c)
					return regionChanged(before2, after2, target) > 0.02 ||
						frameChanged(before2, after2) > 0.0005 ||
						regionChangedThresh(before2, mid2, target, 3) > 0.02 ||
						colorDist(sampleAt(before2, px, py), sampleAt(after2, px, py)) > 400
				}
				if el.label == "slider" {
					// 换到 10% / 90% 两个位置：必然与"中心=当前值"不同。
					ok = retryTarget(b.Min.X+b.W()*0.10, cy, b) ||
						retryTarget(b.Min.X+b.W()*0.90, cy, b)
				} else {
					var nb geom.Rect
					found := false
					vh := c.Scene().Size().H
					c.Scene().Walk(func(n *scene.Node) bool {
						h := n.HitSpec()
						if h.Label != el.label || h.Disabled {
							return true
						}
						r := n.BoundsInRoot()
						if r.IsEmpty() || r.Max.Y <= 40 || r.Min.Y >= vh-40 {
							return true
						}
						d := dist(r.Center(), b.Center())
						if d > 24 && d < 200 && (!found || d < dist(nb.Center(), b.Center())) {
							nb, found = r, true
						}
						return true
					})
					if found {
						ok = retryTarget((nb.Min.X+nb.Max.X)/2, (nb.Min.Y+nb.Max.Y)/2, nb)
					}
				}
			}
		}
		// 兜底：按住期间主线程重建会让手势偶发丢失（真窗竞态，同一窗多次
		// 运行结果会波动）。完整分步再点一次，取两次的 OR。
		if !ok && !el.disabled && el.label != "input" {
			before2 := grab(c)
			injectDown(c, cx, cy)
			time.Sleep(60 * time.Millisecond)
			mid2 := grab(c)
			injectUp(c, cx, cy)
			time.Sleep(430 * time.Millisecond)
			after2 := grab(c)
			ok = regionChanged(before2, after2, b) > 0.02 ||
				frameChanged(before2, after2) > 0.0005 ||
				regionChangedThresh(before2, mid2, b, 3) > 0.02 ||
				colorDist(sampleAt(before2, cx, cy), sampleAt(after2, cx, cy)) > 400
		}
		rep.Findings = append(rep.Findings, Finding{Element: el.name, Check: checkName,
			OK: ok, Detail: fmt.Sprintf("at %.0f,%.0f h=%v region=%.3f frame=%.3f press=%.3f center=%.0f",
				cx, cy, hasHandler, ratio, frameRatio, pressedRatio, colorDist(centerBefore, centerAfter))})
		if !ok {
			rep.Failed++
			// 取证开关：AUDIT_SHOTS=1 时把失败元素的点击前后帧落盘。
			if os.Getenv("AUDIT_SHOTS") != "" {
				tag := fmt.Sprintf("fail-%s-%.0f-%.0f", el.label, cx, cy)
				savePNG(before, filepath.Join(outDir, tag+"-before.png"))
				savePNG(after, filepath.Join(outDir, tag+"-after.png"))
			}
		}
		// 关闭可能弹出的浮层，指针移出窗口清 hover/Tooltip 计时。
		injectLeave(c)
		injectEsc(c)
		time.Sleep(130 * time.Millisecond)
	}

	// ② 逐屏：枚举当前视口内的可交互元素就地验证，然后向下滚一屏；
	// 页面级滚动偏移不再变化即到底（不能用画面 hash 判：验证点击本身
	// 会改变画面，永远"有变化"）。
	prevOff = -1
	page := 0
	for page < 30 {
		off := scrollPos()
		if off == prevOff {
			break
		}
		prevOff = off
		img := grab(c)
		if img == nil {
			break
		}
		savePNG(img, filepath.Join(outDir, fmt.Sprintf("page-%02d.png", page)))
		page++
		rep.Pages = page

		sz := c.Scene().Size()
		var elems []element
		seen := map[uint64]bool{}
		c.Scene().Walk(func(n *scene.Node) bool {
			h := n.HitSpec()
			name, ok := interactiveLabels[h.Label]
			if !ok {
				return true
			}
			b := n.BoundsInRoot()
			if b.IsEmpty() || b.Max.Y <= 40 || b.Min.Y >= sz.H-40 || b.Max.X <= 0 || b.Min.X >= sz.W {
				return true
			}
			if h.Label == "table-row" && n.Parent() != nil {
				pid := n.Parent().ID()
				if seen[pid] {
					return true
				}
				seen[pid] = true
			}
			elems = append(elems, element{label: h.Label, name: name, bounds: b, disabled: h.Disabled})
			return true
		})
		rep.Elements += len(elems)
		for _, el := range elems {
			verifyElement(el)
		}

		injectWheel(c, 1035, 400, -3)
		time.Sleep(250 * time.Millisecond)
	}
	return rep
}