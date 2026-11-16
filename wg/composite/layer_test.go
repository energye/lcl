package composite

import (
	"testing"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/text"
)

// ── 辅助 ────────────────────────────────────────────────────────────

// solid 造一个实心矩形节点。
func solid(w, h float32, c geom.Color) *scene.Node {
	return scene.Box(scene.BoxSpec{Fill: render.Fill(c)}).SetSize2(w, h)
}

// half 造一个半透明矩形节点。
func half(w, h float32, c geom.Color) *scene.Node {
	return solid(w, h, c.WithAlpha(128))
}

// buildDraw 构建层树并画一帧（首帧整屏脏，保证真的提交）。
func buildDraw(t *testing.T, root *scene.Node, opts BuildOptions) (*Tree, Stats) {
	t.Helper()
	c := newComp(t, 60, 60)
	tr := Build(root, opts)
	if err := tr.Validate(); err != nil {
		t.Fatalf("层树校验失败: %v", err)
	}
	st, err := c.DrawTree(tr)
	if err != nil {
		t.Fatalf("合成失败: %v", err)
	}
	return tr, st
}

// buildOnly 只构建不画（结构断言用）。
func buildOnly(root *scene.Node) *Tree { return Build(root, BuildOptions{Cache: true}) }

// ── 结构：默认扁平化 ────────────────────────────────────────────────

func TestFlattenByDefault(t *testing.T) {
	// 100 个普通节点、0 个特效：整棵树只应有 1 个组层 + 1 段指令。
	root := scene.Stack().SetSize2(200, 200)
	for i := 0; i < 100; i++ {
		root.Add(solid(10, 10, geom.RGBA(10, 20, 30)).SetStackPosition(geom.Pt(float32(i), float32(i))))
	}
	root.Layout(scene.Unbounded())
	tr := buildOnly(root)
	if tr.Groups() != 1 {
		t.Fatalf("组层数 = %d，应只有根层（扁平化失效）", tr.Groups())
	}
	if tr.Pictures() != 1 {
		t.Fatalf("指令切片数 = %d，应为 1", tr.Pictures())
	}
	if tr.Stats().Nodes != 101 {
		t.Fatalf("节点数 = %d，应为 101", tr.Stats().Nodes)
	}
	if tr.Stats().Downgraded != 0 {
		t.Fatalf("无特效时不该有降级: %+v", tr.Stats())
	}
}

func TestReasonPerEffect(t *testing.T) {
	cases := []struct {
		name string
		root *scene.Node
		want Reason
	}{
		{"opacity", scene.Opacity(0.5, solid(10, 10, geom.RGBA(1, 2, 3))), ReasonOpacity},
		{"rrect-clip", scene.Clip(scene.ClipSpec{Radius: 4}, solid(10, 10, geom.RGBA(1, 2, 3))), ReasonClip},
		{"boundary", solid(10, 10, geom.RGBA(1, 2, 3)).SetRepaintBoundary(true), ReasonRepaintBoundary},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.root.Layout(scene.Unbounded())
			tr := buildOnly(tc.root)
			var got Reason
			tr.Walk(func(l *Layer) bool {
				if l.Kind() == LayerGroup && l.Reason() != ReasonRoot {
					got = l.Reason()
				}
				return true
			})
			if got != tc.want {
				t.Fatalf("理由 = %v，应为 %v", got, tc.want)
			}
		})
	}
}

func TestRectClipDoesNotCreateLayer(t *testing.T) {
	// 矩形裁剪由 scissor 精确完成，建层是纯损失。
	root := scene.Clip(scene.ClipSpec{Rect: geom.RectXYWH(0, 0, 5, 5)}, solid(10, 10, geom.RGBA(1, 2, 3)))
	root.Layout(scene.Unbounded())
	if tr := buildOnly(root); tr.Groups() != 1 {
		t.Fatalf("矩形裁剪不该建层，组层数 = %d", tr.Groups())
	}
}

// 组滤镜（模糊/发光）现在真的建层，不再降级。
func TestFilterCreatesLayer(t *testing.T) {
	root := scene.Effect(scene.EffectSpec{BlurSigma: 3}, solid(10, 10, geom.RGBA(1, 2, 3)))
	root.Layout(scene.Unbounded())
	tr := buildOnly(root)
	if tr.Groups() != 2 {
		t.Fatalf("组模糊应建滤镜层（根 + filter）: %d", tr.Groups())
	}
	if tr.Stats().Filter != 1 {
		t.Fatalf("滤镜层统计未记录: %+v", tr.Stats())
	}
	if tr.Stats().Downgraded != 0 {
		t.Fatalf("组模糊不应再降级: %+v", tr.Stats())
	}
}

// ── 正确性：组透明不是逐指令乘 alpha ────────────────────────────────

func TestGroupOpacityDoesNotDoubleDarken(t *testing.T) {
	// 两个 50% 红方块部分重叠，整体再套 50% 透明度。
	//
	// 正确结果（组透明 = 先合并再整体混合）：
	//   单层区绿通道 ≈ 255×(1−0.25) ≈ 191
	//   重叠区绿通道 ≈ 255×(1−0.375) ≈ 159
	// 错误的做法（逐指令乘 alpha）重叠区会是 255×(1−0.25)² ≈ 143——
	// 这就是"重叠处重复变暗"的经典 bug，断言窗口把 143 排除在外。
	inner := scene.Stack(
		half(20, 20, geom.RGBA(255, 0, 0)).SetStackPosition(geom.Pt(0, 0)),
		half(20, 20, geom.RGBA(255, 0, 0)).SetStackPosition(geom.Pt(10, 0)),
	).SetSize2(30, 20)
	root := scene.Opacity(0.5, inner)
	root.Layout(scene.Unbounded())

	c := newComp(t, 60, 60)
	tr := Build(root, BuildOptions{Cache: true})
	if _, err := c.DrawTree(tr); err != nil {
		t.Fatal(err)
	}

	b := c.Backing()
	oneG := int(b.Pix()[5*60+5] >> 8 & 0xFF)   // 只有一个方块覆盖
	twoG := int(b.Pix()[5*60+15] >> 8 & 0xFF)  // 两个方块重叠

	if oneG < 186 || oneG > 196 {
		t.Fatalf("单层区绿通道 = %d，应约 191", oneG)
	}
	if twoG < 154 || twoG > 165 {
		t.Fatalf("重叠区绿通道 = %d，应约 159（逐指令乘 alpha 会得到 ≈143）", twoG)
	}
	if twoG >= oneG {
		t.Fatalf("重叠区(%d) 应比单层区(%d) 更不透明", twoG, oneG)
	}
}

func TestOpacityLayerIsOffscreen(t *testing.T) {
	root := scene.Opacity(0.5, solid(10, 10, geom.RGBA(0, 0, 255)))
	root.Layout(scene.Unbounded())
	tr := buildOnly(root)
	var found *Layer
	tr.Walk(func(l *Layer) bool {
		if l.Reason() == ReasonOpacity {
			found = l
		}
		return true
	})
	if found == nil || !found.Offscreen() {
		t.Fatalf("opacity 层必须走离屏: %+v", found)
	}
	if !found.Cacheable() {
		t.Fatal("opacity 层应可缓存")
	}
}

// ── 正确性：圆角裁剪真的裁掉了角 ────────────────────────────────────

func TestRRectClipCutsCorners(t *testing.T) {
	root := scene.Clip(scene.ClipSpec{Radius: 20}, solid(40, 40, geom.RGBA(255, 0, 0)))
	root.Layout(scene.Unbounded())

	c := newComp(t, 60, 60)
	tr := Build(root, BuildOptions{})
	if _, err := c.DrawTree(tr); err != nil {
		t.Fatal(err)
	}
	b := c.Backing()
	// 角上应该是背景（白），中心应该是不透明红
	corner := b.Pix()[1*60+1]
	center := b.Pix()[20*60+20]
	if corner != 0xFFFFFFFF {
		t.Fatalf("左上角应露出背景白（被圆角裁掉），实际 %#08x", corner)
	}
	if center&0xFF != 0 || (center>>16)&0xFF != 255 || center>>24 != 255 {
		t.Fatalf("中心像素错: %#08x", center)
	}
	// 角内侧（4,4）：到圆心距离 22.6 > 20，必须被裁掉
	if got := b.Pix()[4*60+4]; got>>16&0xFF == 255 && got&0xFF == 0 && got>>24 == 255 {
		t.Fatalf("角内侧未被裁掉: %#08x", got)
	}
}

func TestRRectClipOutsideIsBackground(t *testing.T) {
	// 圆角裁剪外必须是**背景原样**：如果蒙版写错（比如漏乘），
	// 会出现"裁剪区外有半透明残影"，这在深色主题上非常明显。
	root := scene.Clip(scene.ClipSpec{Radius: 10}, solid(20, 20, geom.RGBA(0, 0, 0)))
	root.Layout(scene.Unbounded())
	c := newComp(t, 30, 30)
	if _, err := c.DrawTree(Build(root, BuildOptions{})); err != nil {
		t.Fatal(err)
	}
	if got := c.Backing().Pix()[1*30+1]; got != 0xFFFFFFFF {
		t.Fatalf("裁剪外应为纯背景白，实际 %#08x", got)
	}
}

// ── 缓存 ────────────────────────────────────────────────────────────

func TestCacheSteadyStateHits(t *testing.T) {
	root := scene.Stack(
		scene.Opacity(0.6, solid(20, 20, geom.RGBA(0, 128, 255))).SetStackPosition(geom.Pt(5, 5)),
	).SetSize2(60, 60)
	root.Layout(scene.Unbounded())

	c := newComp(t, 60, 60)
	tr := Build(root, BuildOptions{Cache: true})

	// 首帧：层被栅格化
	c.InvalidateAll()
	if _, err := c.DrawTree(tr); err != nil {
		t.Fatal(err)
	}
	// 第二帧：同一棵树（内容没变），脏区只覆盖该层
	c.Invalidate(geom.RectXYWH(0, 0, 60, 60))
	st, err := c.DrawTree(tr)
	if err != nil {
		t.Fatal(err)
	}
	if st.LayerCacheHits == 0 {
		t.Fatalf("稳态应命中缓存: %+v", st)
	}
	if st.LayerRasters != 0 {
		t.Fatalf("稳态不该重新栅格化: %+v", st)
	}
}

func TestCacheMissAfterRebuild(t *testing.T) {
	mk := func() *scene.Node {
		n := scene.Opacity(0.6, solid(20, 20, geom.RGBA(0, 128, 255)))
		n.Layout(scene.Unbounded())
		return n
	}
	root := mk()
	c := newComp(t, 60, 60)
	c.InvalidateAll()
	if _, err := c.DrawTree(Build(root, BuildOptions{Cache: true})); err != nil {
		t.Fatal(err)
	}

	// 内容变了：重录 + 换签名 → 版本换号 → 必须重新栅格化
	root.Children()[0].SetOpacity(0.4)
	root.Layout(scene.Unbounded())
	c.InvalidateAll()
	st, err := c.DrawTree(Build(root, BuildOptions{Cache: true}))
	if err != nil {
		t.Fatal(err)
	}
	if st.LayerRasters == 0 {
		t.Fatalf("内容变化后应重新栅格化: %+v", st)
	}
}

func TestCacheTranslationInvariant(t *testing.T) {
	// 层内容录在层局部空间，整层平移不该让缓存失效——这是 RasterCache 存在的理由。
	// 用滚动来制造平移：placeScroll 直接改写子节点 offset，**不**会换签名，
	// 正是"内容没变、位置变了"的真实场景。
	content := scene.Opacity(0.6, scene.Stack(solid(40, 20, geom.RGBA(0, 128, 255))).SetSize2(40, 200))
	sc := scene.Scroll(scene.Vertical, content)
	sc.SetSize2(40, 40)
	sc.Layout(scene.Unbounded())

	c := newComp(t, 60, 60)
	tr := Build(sc, BuildOptions{Cache: true})
	c.InvalidateAll()
	if _, err := c.DrawTree(tr); err != nil {
		t.Fatal(err)
	}
	// 60% 不透明蓝叠在白底上：蓝通道满、红通道 ≈102
	if got := c.Backing().Pix()[5*60+5]; got&0xFF != 0xFF || (got>>16)&0xFF > 140 {
		t.Fatalf("滚动前 (5,5) 应是（半透明的）蓝色，实际 %#08x", got)
	}
	if got := c.Backing().Pix()[18*60+5]; got&0xFF != 0xFF {
		t.Fatalf("滚动前 (5,18) 应是蓝色，实际 %#08x", got)
	}

	sc.SetScrollOffset(5) // 内容上移 5：下半部滚出视口顶边
	tr2 := Build(sc, BuildOptions{Cache: true})
	c.InvalidateAll()
	st, err := c.DrawTree(tr2)
	if err != nil {
		t.Fatal(err)
	}
	if st.LayerCacheHits == 0 {
		t.Fatalf("整层平移应命中缓存（版本不随位置变化）: %+v", st)
	}
	if st.LayerRasters != 0 {
		t.Fatalf("平移不该重新栅格化: %+v", st)
	}
	if got := c.Backing().Pix()[18*60+5]; got != 0xFFFFFFFF {
		t.Fatalf("滚动后 (5,18) 应是背景白，实际 %#08x", got)
	}
	if got := c.Backing().Pix()[5*60+5]; got&0xFF != 0xFF {
		t.Fatalf("滚动后 (5,5) 仍应是蓝色，实际 %#08x", got)
	}
}

func TestCacheDisabledStillCorrect(t *testing.T) {
	root := scene.Opacity(0.5, solid(20, 20, geom.RGBA(255, 0, 0)))
	root.Layout(scene.Unbounded())

	c1 := newComp(t, 60, 60)
	if _, err := c1.DrawTree(Build(root, BuildOptions{Cache: true})); err != nil {
		t.Fatal(err)
	}
	c2 := newComp(t, 60, 60)
	c2.SetRasterCache(false)
	if _, err := c2.DrawTree(Build(root, BuildOptions{Cache: false})); err != nil {
		t.Fatal(err)
	}
	if diff := diffBacking(c1.Backing(), c2.Backing()); diff != 0 {
		t.Fatalf("开/关缓存像素必须一致，差异像素 %d", diff)
	}
}

func TestCacheBudgetEvicts(t *testing.T) {
	cache := NewRasterCache(64 * 64 * 4) // 只够一块
	cache.BeginFrame()
	a := cache.Lease(64, 64)
	cache.Store(CacheKey{ID: 1, Version: 1, W: 64, H: 64}, a, nil)
	b := cache.Lease(64, 64)
	cache.Store(CacheKey{ID: 2, Version: 1, W: 64, H: 64}, b, nil)
	if got := cache.Stats().Entries; got != 1 {
		t.Fatalf("超预算应淘汰到 1 块，实际 %d", got)
	}
	if cache.Lookup(CacheKey{ID: 1, Version: 1, W: 64, H: 64}) != nil {
		t.Fatal("最久未用的那块应被淘汰")
	}
}

// ── 降级路径：文字 ──────────────────────────────────────────────────

func TestTextOnOpacityLayerFlattens(t *testing.T) {
	// GDI 文字画在透明离屏上会丢 alpha，所以含文字的层不走向量离屏，
	// 而是降级为扁平化。这条测试守着"降级必须被计数且不建离屏层"。
	txt := scene.Text(scene.TextSpec{
		Src:        "hello",
		Layout:     testLayout(),
		LayoutFont: text.DefaultFont(14),
		Color:      geom.RGBA(0, 0, 0),
	})
	root := scene.Opacity(0.5, txt)
	root.Layout(scene.Unbounded())
	tr := buildOnly(root)
	if tr.Stats().TextFlat != 1 {
		t.Fatalf("含文字层应计一次 TextFlat: %+v", tr.Stats())
	}
	var off bool
	tr.Walk(func(l *Layer) bool {
		if l.Reason() == ReasonOpacity && l.Offscreen() {
			off = true
		}
		return true
	})
	if off {
		t.Fatal("含文字的 opacity 层不该离屏")
	}
}

// ── Validate ────────────────────────────────────────────────────────

func TestValidateCatchesMissingReason(t *testing.T) {
	tr := &Tree{root: &Layer{kind: LayerGroup, reason: ReasonNone}}
	if err := tr.Validate(); err == nil {
		t.Fatal("组层缺理由应被 Validate 抓住（C19）")
	}
}

func TestValidateCatchesPictureWithChildren(t *testing.T) {
	tr := &Tree{root: &Layer{
		kind: LayerGroup, reason: ReasonRoot,
		children: []*Layer{{
			kind: LayerPicture, reason: ReasonContent,
			children: []*Layer{{kind: LayerPicture, reason: ReasonContent}},
		}},
	}}
	if err := tr.Validate(); err == nil {
		t.Fatal("picture 层带子层应被抓住")
	}
}

// diffBacking 返回两块缓冲的不同像素数（尺寸不同则返回 -1）。
func diffBacking(a, b *raster.Backing) int {
	if a.W() != b.W() || a.H() != b.H() {
		return -1
	}
	pa, pb := a.Pix(), b.Pix()
	n := 0
	for i := range pa {
		if pa[i] != pb[i] {
			n++
		}
	}
	return n
}
