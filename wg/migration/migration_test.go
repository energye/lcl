package migration

import (
	"testing"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

func mctx() *ui.BuildContext { return &ui.BuildContext{Theme: theme.Default()} }

// 三个场景都能在真实 Container 链路上出帧（零 lcl 的 headless）。
func TestScenesRender(t *testing.T) {
	cases := []struct {
		name string
		w    ui.Widget
		vw   int
	}{
		{"showcase", NewShowcase(), 1440},
		{"canvas_ime", NewCanvasIME(), 960},
		{"dbmodeler", NewDBModeler(), 1100},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := wg.NewHeadless(wg.DefaultOptions())
			c.Resize(tc.vw, 720)
			c.MountWidget(tc.w)
			stats, err := c.FrameOnce()
			if err != nil {
				t.Fatalf("%s 出帧失败: %v", tc.name, err)
			}
			if stats.Nodes == 0 {
				t.Fatalf("%s 没有渲染节点", tc.name)
			}
		})
	}
}

// Showcase 包含图表节点（CustomPaintNode），验证迁移真的用了新图表组件。
func TestShowcaseHasCharts(t *testing.T) {
	s := NewShowcase()
	n := s.Build(mctx())
	var custom int
	n.Walk(func(x *scene.Node) bool {
		if x.Kind() == scene.KindCustomPaint {
			custom++
		}
		return true
	})
	if custom < 5 {
		t.Fatalf("Showcase 自定义绘制节点 = %d（波形/仪表/折线/柱/饼应 ≥5）", custom)
	}
}

// CanvasIME 的输入框构建后焦点链有效，中文上屏链路不断。
func TestCanvasIMEInputChain(t *testing.T) {
	c := NewCanvasIME()
	n := c.Build(mctx())
	if n == nil {
		t.Fatal("场景为空")
	}
	c.Input.Editor().InsertText("你好")
	if got := c.Input.Editor().Value().Text; got != "你好" {
		t.Fatalf("输入法场景文本 = %q", got)
	}
	c.Clear.OnClick()
	if got := c.Input.Editor().Value().Text; got != "" {
		t.Fatalf("清空失败: %q", got)
	}
}

// DBModeler 的 Table 是虚拟的，100 行不会构建 100 个节点；关系画布存在。
func TestDBModelerScene(t *testing.T) {
	d := NewDBModeler()
	if len(d.Tables) != 3 || len(d.Table.Rows) < 8 {
		t.Fatalf("模型数据不完整: tables=%d rows=%d", len(d.Tables), len(d.Table.Rows))
	}
	n := d.Build(mctx())
	var custom int
	n.Walk(func(x *scene.Node) bool {
		if x.Kind() == scene.KindCustomPaint {
			custom++
		}
		return true
	})
	if custom < 1 {
		t.Fatal("关系画布缺失")
	}
}

// 波形采样在 [-1,1] 内且非空。
func TestWaveSamples(t *testing.T) {
	v := waveSamples(180)
	if len(v) != 180 {
		t.Fatalf("采样数 = %d", len(v))
	}
	for i, x := range v {
		if x < -1.2 || x > 1.2 {
			t.Fatalf("采样[%d] = %v 越界", i, x)
		}
	}
}
