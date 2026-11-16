package examples

import (
	"testing"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
)

// TestEveryExampleInContainer 是逐项容器测试的主体：每条示例独占一个
// headless Container，依次验证「首帧出图 → 注入按下/抬起 → 再出帧」。
//
// 断言的不是某个组件的业务结果，而是每条示例都真实通过了容器管线：
// BuildContext → 布局 → 录制 → 合成 → Present，并且在输入分发后仍健康。
// 这样新增任何组件或引擎功能示例时，只要忘了接进容器，这里立刻红灯。
func TestEveryExampleInContainer(t *testing.T) {
	if len(All) == 0 {
		t.Fatal("示例注册表为空")
	}
	seen := make(map[string]bool, len(All))

	for _, ex := range All {
		if seen[ex.ID] {
			t.Fatalf("重复示例 ID: %s", ex.ID)
		}
		seen[ex.ID] = true

		ex := ex
		t.Run(ex.ID, func(t *testing.T) {
			c := wg.NewHeadless(wg.DefaultOptions())
			c.Resize(1280, 800)
			c.MountWidget(ex.Widget())

			st, err := c.FrameOnce()
			if err != nil {
				t.Fatalf("首帧失败: %v", err)
			}
			if !st.Presented {
				t.Fatal("首帧未 Present")
			}
			if st.Nodes < 1 {
				t.Fatalf("节点数 = %d，应至少 1", st.Nodes)
			}
			if c.Content() == nil {
				t.Fatal("Content 为空：示例未被挂载")
			}
			if c.Frames() < 1 {
				t.Fatalf("帧计数 = %d", c.Frames())
			}

			// 内容区中点一下：覆盖命中路径/手势/浮层关闭策略。
			p := geom.Pt(640, 400)
			c.PushEvent(&input.PointerEvent{
				Kind: input.PointerDown, Pos: p, PosRoot: p,
				Buttons: input.ButtonLeft, ClickCount: 1,
			})
			c.PushEvent(&input.PointerEvent{
				Kind: input.PointerUp, Pos: p, PosRoot: p,
				ClickCount: 1,
			})
			if _, err := c.FrameOnce(); err != nil {
				t.Fatalf("输入后出帧失败: %v", err)
			}
		})
	}
}

// TestEveryExamplePaintsPixels 逐条验证示例的首帧真的画出了非背景像素。
//
// TestEveryExampleInContainer 只断言 Presented/Nodes，会把"节点全在零宽区域"
// 这类真窗空白 bug 误判为通过（M6 的教训）。这里直接采样合成缓冲：每条示例
// 都必须有足够多的非背景像素，才能证明该能力在容器上真实可见。
func TestEveryExamplePaintsPixels(t *testing.T) {
	for _, ex := range All {
		ex := ex
		t.Run(ex.ID, func(t *testing.T) {
			c := wg.NewHeadless(wg.DefaultOptions())
			c.Resize(1024, 768)
			c.MountWidget(ex.Widget())
			if _, err := c.FrameOnce(); err != nil {
				t.Fatalf("首帧失败: %v", err)
			}
			b := c.Backing()
			if b == nil {
				t.Fatal("合成缓冲为空")
			}
			nonBg := 0
			for y := 0; y < b.H(); y += 4 {
				for x := 0; x < b.W(); x += 4 {
					p := b.Pix()[y*b.W()+x]
					r, g, bl := byte(p>>16), byte(p>>8), byte(p)
					if r != 255 || g != 255 || bl != 255 {
						nonBg++
					}
				}
			}
			if nonBg < 100 {
				t.Fatalf("非背景采样点=%d，画布疑似空白", nonBg)
			}
		})
	}
}

// TestGallerySelectsEveryExample 验证画廊能逐项切换并出帧。
// 这是对真窗里"点左侧导航 → 右侧换页"这条链路的 headless 回归。
func TestGallerySelectsEveryExample(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	c.Resize(1280, 800)
	g := NewGallery()
	c.MountWidget(g)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatalf("画廊首帧失败: %v", err)
	}

	for _, ex := range All {
		if !g.Select(ex.ID) {
			t.Fatalf("Select(%s) 返回 false", ex.ID)
		}
		if got := g.Current().ID; got != ex.ID {
			t.Fatalf("Current() = %s，期望 %s", got, ex.ID)
		}
		if _, err := c.FrameOnce(); err != nil {
			t.Fatalf("切到 %s 后出帧失败: %v", ex.ID, err)
		}
	}

	g.Next()
	if len(All) > 0 && g.Current().ID == "" {
		t.Fatal("Next 后当前示例为空")
	}
}

// TestGalleryContentFillsViewport 是这次真窗空白 bug 的回归门禁。
//
// 症状：内容节点被 Stack 用 LooseConstraints 测量，只按内容自然宽高布局，
// 于是左侧导航 226px + 右侧内容 0px，图表画在零宽区域里——进程活着、
// Presented=true、单测全绿，但真窗里什么都看不见。
func TestGalleryContentFillsViewport(t *testing.T) {
	const vw, vh = 1360, 860
	c := wg.NewHeadless(wg.DefaultOptions())
	c.Resize(vw, vh)
	c.MountWidget(NewGallery())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatalf("首帧失败: %v", err)
	}
	root := c.Content()
	if root == nil {
		t.Fatal("Content 为空")
	}
	if got := root.Size(); int(got.W) != vw || int(got.H) != vh {
		t.Fatalf("画廊根节点 = %v，应铺满 %dx%d", got, vw, vh)
	}
	if len(root.Children()) < 2 {
		t.Fatalf("画廊子节点数 = %d，应有导航+内容两栏", len(root.Children()))
	}
	body := root.Children()[1]
	if got := body.Size(); got.W < float32(vw)/2 || got.H < float32(vh)/2 {
		t.Fatalf("内容栏 = %v，未铺满视口过半", got)
	}
	// Section 的列也必须横向拉伸，否则图表仍落在零宽区域。
	if len(body.Children()) > 2 {
		sec := body.Children()[2]
		if got := sec.Size(); got.W < 400 || got.H < 200 {
			t.Fatalf("示例内容 Section = %v，尺寸不足", got)
		}
	}
}

// TestGalleryFrameHasPixels 验证首帧合成缓冲里真的有内容（像素级）。
// 只看节点数/Presented 会把"所有节点都在零宽区域"误判为通过。
func TestGalleryFrameHasPixels(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	c.Resize(1360, 860)
	c.MountWidget(NewGallery())
	if _, err := c.FrameOnce(); err != nil {
		t.Fatalf("首帧失败: %v", err)
	}
	b := c.Backing()
	if b == nil {
		t.Fatal("合成缓冲为空")
	}
	// 右侧内容区采样：默认首条是折线图，应有坐标轴/曲线等非白像素。
	nonBg := 0
	for y := 120; y < 500; y += 3 {
		for x := 300; x < 1200; x += 3 {
			if x >= b.W() || y >= b.H() {
				continue
			}
			p := b.Pix()[y*b.W()+x]
			r, g, bl := byte(p>>16), byte(p>>8), byte(p)
			if r != 255 || g != 255 || bl != 255 {
				nonBg++
			}
		}
	}
	if nonBg < 200 {
		t.Fatalf("内容区非背景采样点 = %d，画布疑似空白", nonBg)
	}
}

// TestAllGroupsCovered 防止新增分组却忘了配示例。
func TestAllGroupsCovered(t *testing.T) {
	groups := []string{
		GroupBasic, GroupControl, GroupDisplay, GroupData,
		GroupFeedback, GroupChart, GroupLayout, GroupVisual, GroupInput,
	}
	for _, g := range groups {
		n := 0
		for _, ex := range All {
			if ex.Group == g {
				n++
			}
		}
		if n == 0 {
			t.Errorf("分组 %q 没有任何示例", g)
		}
	}
}
