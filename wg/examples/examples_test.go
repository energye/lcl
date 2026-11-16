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
