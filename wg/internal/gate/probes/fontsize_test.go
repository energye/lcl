//go:build gateprobe

package probes

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// TestProbeFontSizeAlign 覆盖 §6A：正文字号对齐后组件的 mount/首帧/
// 稳态帧，以及 100 实例批量 mount。字号本身是度量输入，改动会影响
// 布局与文字光栅，故按组件级探针验证不出现明显回退。
func TestProbeFontSizeAlign(t *testing.T) {
	th := theme.Default()
	ctx := &ui.BuildContext{Theme: th}

	// 单实例首帧：Table 是字号改动最集中的组件（列宽按字号计算）。
	mkTable := func() *widgets.Table {
		return &widgets.Table{
			Columns: []widgets.TableColumn{{Title: "名称"}, {Title: "数量"}},
			Rows:    [][]any{{"苹果", "10"}, {"香蕉", "20"}},
			Width:   360,
		}
	}
	// 预热：首次文字度量含字体缓存冷启动，不计入失效重绘计时。
	mkTable().Build(ctx)
	tbl := mkTable()
	began := time.Now()
	if n := tbl.Build(ctx); n == nil {
		t.Fatal("Table 未构建")
	}
	if elapsed := time.Since(began); elapsed > 16*time.Millisecond {
		t.Fatalf("Table 首帧=%v，阈值 16ms", elapsed)
	}

	// 100 实例批量 mount：同类控件密集页面的构建开销。
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1200, 900)
	root := scene.Flex(scene.Vertical, 4)
	for i := 0; i < 100; i++ {
		root.Add((&widgets.Label{Text: "字号对齐探针", Size: th.Font.SizeBase}).Build(ctx))
	}
	c.MountWidget(nodeWidget{root})
	began = time.Now()
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if elapsed := time.Since(began); elapsed > 50*time.Millisecond {
		t.Fatalf("100 实例 mount=%v，阈值 50ms", elapsed)
	}
	// 稳态帧：无失效时不应持续重绘。
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
}

type nodeWidget struct{ n *scene.Node }

func (w nodeWidget) Build(*ui.BuildContext) *scene.Node { return w.n }
