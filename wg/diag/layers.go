package diag

import (
	"fmt"
	"time"

	"github.com/energye/lcl/wg/composite"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster/software"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
)

// ── --layers ────────────────────────────────────────────────────────

// LayerCase 是单条层树/合成断言的结果。
type LayerCase struct {
	Name string
	OK   bool
	Got  string
	Note string
}

// LayerResult 是 --layers 的汇总。
type LayerResult struct {
	Cases   int
	Failed  int
	Elapsed time.Duration
	Details []LayerCase
}

// OK 判定是否全绿（门禁 ≥16 组）。
func (r LayerResult) OK() bool { return r.Failed == 0 && r.Cases >= 16 }

// String 返回可读摘要。
func (r LayerResult) String() string {
	return fmt.Sprintf("layers: %d 组，失败 %d，用时 %v", r.Cases, r.Failed, r.Elapsed)
}

// layerCase 是一条断言：造一棵场景树，构建层树（必要时合成），校验。
type layerCase struct {
	name string
	// want 是"期望看到的东西"，出错时打印。
	want string
	run  func() (got string, err error)
}

// Layers 跑层树门禁：覆盖 C19（层必须声明理由）、C7（缓存归属）、
// 组透明语义、圆角裁剪、缓存稳态与失效、降级计数。
func Layers() (LayerResult, error) {
	began := time.Now()
	var res LayerResult
	for _, c := range layerCases() {
		got, err := c.run()
		d := LayerCase{Name: c.name, OK: err == nil, Got: got}
		if err != nil {
			d.Note = err.Error()
		}
		if !d.OK {
			res.Failed++
		}
		if d.Got == "" {
			d.Got = "—"
		}
		res.Cases++
		res.Details = append(res.Details, d)
	}
	res.Elapsed = time.Since(began)
	return res, nil
}

// ── 用例 ────────────────────────────────────────────────────────────

func layerCases() []layerCase {
	return []layerCase{
		{
			name: "flatten-default",
			want: "200 节点 0 特效 → 1 组层 + 1 切片",
			run: func() (string, error) {
				root := scene.Stack().SetSize2(200, 200)
				for i := 0; i < 200; i++ {
					root.Add(box(8, 8, geom.RGBA(30, 60, 90)).
						SetStackPosition(geom.Pt(float32(i%20)*10, float32(i/20)*10)))
				}
				root.Layout(scene.Unbounded())
				tr := composite.Build(root, composite.BuildOptions{Cache: true})
				got := fmt.Sprintf("组层 %d / 切片 %d / 节点 %d", tr.Groups(), tr.Pictures(), tr.Stats().Nodes)
				if err := tr.Validate(); err != nil {
					return got, err
				}
				if tr.Groups() != 1 || tr.Pictures() != 1 {
					return got, fmt.Errorf("扁平化失效：应 1 组 1 切片")
				}
				return got, nil
			},
		},
		{
			name: "every-group-declares-reason",
			want: "所有组层的 reason != none（C19）",
			run: func() (string, error) {
				root := scene.Stack(
					scene.Opacity(0.5, box(20, 20, geom.RGBA(1, 2, 3))),
					scene.Clip(scene.ClipSpec{Radius: 6}, box(20, 20, geom.RGBA(4, 5, 6))),
					box(10, 10, geom.RGBA(7, 8, 9)).SetRepaintBoundary(true),
				).SetSize2(80, 80)
				root.Layout(scene.Unbounded())
				tr := composite.Build(root, composite.BuildOptions{Cache: true})
				n := 0
				var bad []string
				tr.Walk(func(l *composite.Layer) bool {
					if l.Kind() == composite.LayerGroup {
						n++
						if l.Reason() == composite.ReasonNone {
							bad = append(bad, "无理由组层")
						}
					}
					return true
				})
				got := fmt.Sprintf("%d 组层", n)
				if len(bad) > 0 {
					return got, fmt.Errorf("有组层未声明理由")
				}
				if n < 4 {
					return got, fmt.Errorf("期望 4 个组层（根+3）")
				}
				return got, nil
			},
		},
		{
			name: "opacity-goes-offscreen",
			want: "opacity<1 的层必须离屏且可缓存",
			run: func() (string, error) {
				root := scene.Opacity(0.5, box(20, 20, geom.RGBA(0, 0, 255)))
				root.Layout(scene.Unbounded())
				l := findReason(root, composite.ReasonOpacity)
				if l == nil {
					return "未找到", fmt.Errorf("未建 opacity 层")
				}
				got := fmt.Sprintf("offscreen=%v cacheable=%v", l.Offscreen(), l.Cacheable())
				if !l.Offscreen() || !l.Cacheable() {
					return got, fmt.Errorf("opacity 层必须离屏且可缓存")
				}
				return got, nil
			},
		},
		{
			name: "rect-clip-stays-flat",
			want: "矩形裁剪交给 scissor，不建层（层数=1）",
			run: func() (string, error) {
				root := scene.Clip(scene.ClipSpec{Rect: geom.RectXYWH(0, 0, 10, 10)},
					box(20, 20, geom.RGBA(1, 2, 3)))
				root.Layout(scene.Unbounded())
				tr := composite.Build(root, composite.BuildOptions{})
				got := fmt.Sprintf("组层 %d", tr.Groups())
				if tr.Groups() != 1 {
					return got, fmt.Errorf("矩形裁剪不该建层")
				}
				return got, nil
			},
		},
		{
			name: "rrect-clip-cuts-corners",
			want: "半径 20 的 40×40 圆角裁剪：角落露出背景",
			run: func() (string, error) {
				c := comp(60, 60)
				root := scene.Clip(scene.ClipSpec{Radius: 20}, box(40, 40, geom.RGBA(255, 0, 0)))
				root.Layout(scene.Unbounded())
				if _, err := c.DrawTree(composite.Build(root, composite.BuildOptions{})); err != nil {
					return "", err
				}
				corner := c.Backing().Pix()[1*60+1]
				center := c.Backing().Pix()[20*60+20]
				got := fmt.Sprintf("角 %#08x 中心 %#08x", corner, center)
				if corner != 0xFFFFFFFF {
					return got, fmt.Errorf("角落未被裁掉")
				}
				if center>>24 != 255 || center&0xFF != 0 {
					return got, fmt.Errorf("中心像素错")
				}
				return got, nil
			},
		},
		{
			name: "group-opacity-not-double-darkened",
			want: "两个 50% 方块套 50% 组透明：重叠区绿通道 ≈159（逐指令乘 alpha 会 ≈143）",
			run: func() (string, error) {
				c := comp(60, 60)
				inner := scene.Stack(
					box(20, 20, geom.RGBAa(255, 0, 0, 128)).SetStackPosition(geom.Pt(0, 0)),
					box(20, 20, geom.RGBAa(255, 0, 0, 128)).SetStackPosition(geom.Pt(10, 0)),
				).SetSize2(30, 20)
				root := scene.Opacity(0.5, inner)
				root.Layout(scene.Unbounded())
				if _, err := c.DrawTree(composite.Build(root, composite.BuildOptions{Cache: true})); err != nil {
					return "", err
				}
				one := int(c.Backing().Pix()[5*60+5] >> 8 & 0xFF)
				two := int(c.Backing().Pix()[5*60+15] >> 8 & 0xFF)
				got := fmt.Sprintf("单层 %d 重叠 %d", one, two)
				if one < 186 || one > 196 {
					return got, fmt.Errorf("单层区应 ≈191")
				}
				if two < 154 || two > 165 {
					return got, fmt.Errorf("重叠区应 ≈159")
				}
				return got, nil
			},
		},
		{
			name: "filter-downgrade-counted",
			want: "组模糊 M2 未实现 → 不建层但降级计数=1（C22）",
			run: func() (string, error) {
				root := scene.Effect(scene.EffectSpec{BlurSigma: 4}, box(20, 20, geom.RGBA(1, 2, 3)))
				root.Layout(scene.Unbounded())
				tr := composite.Build(root, composite.BuildOptions{})
				got := fmt.Sprintf("组层 %d 降级 %d", tr.Groups(), tr.Stats().Downgraded)
				if tr.Groups() != 1 {
					return got, fmt.Errorf("不该建层")
				}
				if tr.Stats().Downgraded != 1 {
					return got, fmt.Errorf("降级必须计数")
				}
				return got, nil
			},
		},
		{
			name: "text-on-layer-flattens",
			want: "含文字的 opacity 层不离屏（GDI 文字在透明底上会丢 alpha），计 TextFlat",
			run: func() (string, error) {
				font := text.DefaultFont(14)
				txt := scene.Text(scene.TextSpec{
					Src:        "wg",
					Layout:     text.LayoutText("wg", font, text.Options{}),
					LayoutFont: font,
					Color:      geom.RGBA(0, 0, 0),
				})
				root := scene.Opacity(0.5, txt)
				root.Layout(scene.Unbounded())
				tr := composite.Build(root, composite.BuildOptions{Cache: true})
				l := findReason(root, composite.ReasonOpacity)
				if l == nil {
					return "未找到层", fmt.Errorf("未建层")
				}
				got := fmt.Sprintf("offscreen=%v TextFlat=%d", l.Offscreen(), tr.Stats().TextFlat)
				if l.Offscreen() {
					return got, fmt.Errorf("含文字层不该离屏")
				}
				if tr.Stats().TextFlat != 1 {
					return got, fmt.Errorf("应计一次 TextFlat")
				}
				return got, nil
			},
		},
		{
			name: "cache-steady-state",
			want: "同一棵树第二帧：命中>0 且栅格化=0",
			run: func() (string, error) {
				c := comp(80, 80)
				root := scene.Stack(
					scene.Opacity(0.6, box(30, 30, geom.RGBA(0, 128, 255))).SetStackPosition(geom.Pt(4, 4)),
				).SetSize2(80, 80)
				root.Layout(scene.Unbounded())
				tr := composite.Build(root, composite.BuildOptions{Cache: true})
				if _, err := c.DrawTree(tr); err != nil {
					return "", err
				}
				c.InvalidateAll()
				st, err := c.DrawTree(tr)
				if err != nil {
					return "", err
				}
				got := fmt.Sprintf("命中 %d 栅格化 %d", st.LayerCacheHits, st.LayerRasters)
				if st.LayerCacheHits == 0 || st.LayerRasters != 0 {
					return got, fmt.Errorf("稳态未命中缓存")
				}
				return got, nil
			},
		},
		{
			name: "cache-invalidates-on-change",
			want: "内容变了（重录+签名变）→ 必须重新栅格化",
			run: func() (string, error) {
				c := comp(80, 80)
				root := scene.Opacity(0.6, box(30, 30, geom.RGBA(0, 128, 255)))
				root.Layout(scene.Unbounded())
				if _, err := c.DrawTree(composite.Build(root, composite.BuildOptions{Cache: true})); err != nil {
					return "", err
				}
				root.SetOpacity(0.3)
				root.Layout(scene.Unbounded())
				c.InvalidateAll()
				st, err := c.DrawTree(composite.Build(root, composite.BuildOptions{Cache: true}))
				if err != nil {
					return "", err
				}
				got := fmt.Sprintf("栅格化 %d", st.LayerRasters)
				if st.LayerRasters == 0 {
					return got, fmt.Errorf("内容变化后应重新栅格化")
				}
				return got, nil
			},
		},
		{
			name: "cache-survives-translation",
			want: "滚动（位置变了、内容没变）→ 仍然命中",
			run: func() (string, error) {
				c := comp(60, 60)
				content := scene.Opacity(0.6, scene.Stack(box(40, 20, geom.RGBA(0, 128, 255))).SetSize2(40, 200))
				sc := scene.Scroll(scene.Vertical, content)
				sc.SetSize2(40, 40)
				sc.Layout(scene.Unbounded())
				if _, err := c.DrawTree(composite.Build(sc, composite.BuildOptions{Cache: true})); err != nil {
					return "", err
				}
				sc.SetScrollOffset(5)
				c.InvalidateAll()
				st, err := c.DrawTree(composite.Build(sc, composite.BuildOptions{Cache: true}))
				if err != nil {
					return "", err
				}
				got := fmt.Sprintf("命中 %d 栅格化 %d", st.LayerCacheHits, st.LayerRasters)
				if st.LayerCacheHits == 0 || st.LayerRasters != 0 {
					return got, fmt.Errorf("整层平移不该失效缓存")
				}
				return got, nil
			},
		},
		{
			name: "cache-toggle-pixel-identical",
			want: "开/关缓存必须像素完全一致",
			run: func() (string, error) {
				mk := func() *scene.Node {
					r := scene.Stack(
						scene.Opacity(0.55, box(20, 20, geom.RGBA(220, 40, 40))).SetStackPosition(geom.Pt(6, 6)),
						scene.Clip(scene.ClipSpec{Radius: 8}, box(20, 20, geom.RGBA(40, 200, 120))).
							SetStackPosition(geom.Pt(20, 20)),
					).SetSize2(60, 60)
					r.Layout(scene.Unbounded())
					return r
				}
				a := comp(60, 60)
				if _, err := a.DrawTree(composite.Build(mk(), composite.BuildOptions{Cache: true})); err != nil {
					return "", err
				}
				b := comp(60, 60)
				b.SetRasterCache(false)
				if _, err := b.DrawTree(composite.Build(mk(), composite.BuildOptions{Cache: false})); err != nil {
					return "", err
				}
				diff := 0
				pa, pb := a.Backing().Pix(), b.Backing().Pix()
				for i := range pa {
					if pa[i] != pb[i] {
						diff++
					}
				}
				got := fmt.Sprintf("差异 %d 像素", diff)
				if diff != 0 {
					return got, fmt.Errorf("缓存开关改变了像素")
				}
				return got, nil
			},
		},
		{
			name: "damage-skips-far-layers",
			want: "脏区远离层 → 整层跳过（不清也不画）",
			run: func() (string, error) {
				c := comp(120, 120)
				root := scene.Opacity(0.6, box(20, 20, geom.RGBA(0, 0, 255)))
				root.Layout(scene.Unbounded())
				tr := composite.Build(root, composite.BuildOptions{Cache: true})
				if _, err := c.DrawTree(tr); err != nil {
					return "", err
				}
				c.Invalidate(geom.RectXYWH(80, 80, 10, 10))
				st, err := c.DrawTree(tr)
				if err != nil {
					return "", err
				}
				got := fmt.Sprintf("跳过 %d 栅格化 %d", st.LayerSkipped, st.LayerRasters)
				if st.LayerSkipped == 0 {
					return got, fmt.Errorf("远离脏区的层应被跳过")
				}
				return got, nil
			},
		},
		{
			name: "layer-count-tracks-effects",
			want: "10 个 opacity 节点 → 11 个组层（根+10），与节点数无关",
			run: func() (string, error) {
				root := scene.Stack().SetSize2(200, 200)
				for i := 0; i < 10; i++ {
					root.Add(scene.Opacity(0.9, box(10, 10, geom.RGBA(1, 1, 1))).
						SetStackPosition(geom.Pt(float32(i)*12, 0)))
				}
				root.Layout(scene.Unbounded())
				tr := composite.Build(root, composite.BuildOptions{Cache: true})
				got := fmt.Sprintf("组层 %d / 节点 %d", tr.Groups(), tr.Stats().Nodes)
				if tr.Groups() != 11 {
					return got, fmt.Errorf("组层数应等于特效数+1")
				}
				return got, nil
			},
		},
		{
			name: "depth-limit-flattens",
			want: "嵌套超过上限 → 停止建层并降级计数",
			run: func() (string, error) {
				var n = box(10, 10, geom.RGBA(1, 2, 3))
				for i := 0; i < 12; i++ {
					n = scene.Opacity(0.9, n)
				}
				n.Layout(scene.Unbounded())
				tr := composite.Build(n, composite.BuildOptions{Cache: true, MaxDepth: 4})
				got := fmt.Sprintf("组层 %d 最大深度 %d", tr.Groups(), tr.Stats().MaxDepth)
				if tr.Groups() > 5 {
					return got, fmt.Errorf("超出上限的嵌套应停止建层")
				}
				if tr.Stats().MaxDepth > 4 {
					return got, fmt.Errorf("层深超过上限")
				}
				return got, nil
			},
		},
		{
			name: "empty-layer-not-kept",
			want: "内容被裁空 → 不留空层",
			run: func() (string, error) {
				// 裁剪矩形落在内容之外：交集为空 → 该层不该被保留
				root := scene.Clip(scene.ClipSpec{Rect: geom.RectXYWH(100, 100, 10, 10), Radius: 4},
					box(50, 50, geom.RGBA(1, 2, 3)))
				root.Layout(scene.Unbounded())
				tr := composite.Build(root, composite.BuildOptions{Cache: true})
				if err := tr.Validate(); err != nil {
					return fmt.Sprintf("组层 %d", tr.Groups()), err
				}
				got := fmt.Sprintf("组层 %d", tr.Groups())
				if tr.Groups() != 1 {
					return got, fmt.Errorf("应只剩根层")
				}
				return got, nil
			},
		},
	}
}

// ── 辅助 ────────────────────────────────────────────────────────────

// box 造一个实心矩形节点。
func box(w, h float32, c geom.Color) *scene.Node {
	return scene.Box(scene.BoxSpec{Fill: render.Fill(c)}).SetSize2(w, h)
}

// comp 造一个 headless 合成器。
func comp(w, h int) *composite.Compositor {
	c, err := composite.New(composite.Options{
		Backend:    software.New(software.Options{}),
		Width:      w,
		Height:     h,
		Background: geom.RGBA(255, 255, 255),
	})
	if err != nil {
		panic("wg-diag: 创建合成器失败: " + err.Error())
	}
	return c
}

// findReason 在层树里按理由找第一个组层。
func findReason(root *scene.Node, r composite.Reason) *composite.Layer {
	tr := composite.Build(root, composite.BuildOptions{Cache: true})
	var found *composite.Layer
	tr.Walk(func(l *composite.Layer) bool {
		if found == nil && l.Kind() == composite.LayerGroup && l.Reason() == r {
			found = l
		}
		return true
	})
	return found
}
