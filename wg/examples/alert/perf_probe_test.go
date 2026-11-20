package main

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 临时性能探针（§6A.1，提交前删除）：Alert 窗 mount/首帧/稳态阈值验证。
// 阈值：mount≤50ms、首帧≤50ms、稳态帧≤2ms。

func TestAlertPerfProbe(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1080, 760)

	t0 := time.Now()
	c.MountWidget(doc.Page(doc.Meta{
		ID: "alert", Title: "Alert", Group: doc.GroupFeedback,
		What: []string{"perf probe"},
		Demos: []doc.Demo{
			{Title: "四类型", Build: func(ctx *ui.BuildContext) ui.Widget {
				col := scene.Flex(scene.Vertical, 10)
				for _, a := range []widgets.AlertType{
					widgets.AlertInfo, widgets.AlertSuccess, widgets.AlertWarning, widgets.AlertError,
				} {
					col.Add((&widgets.Alert{Type: a, Title: "提示标题", ShowIcon: true, Width: 420}).Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
			{Title: "大批量", Build: func(ctx *ui.BuildContext) ui.Widget {
				// §6A.1-2：100 个实例渲染不降级。
				col := scene.Flex(scene.Vertical, 6)
				for i := 0; i < 100; i++ {
					col.Add((&widgets.Alert{Type: widgets.AlertInfo, Title: "批量", Width: 400}).Build(ctx))
				}
				return doc.Node(func(*ui.BuildContext) *scene.Node { return col })
			}},
		},
	}))
	mount := time.Since(t0)

	t1 := time.Now()
	c.FrameOnce()
	frame1 := time.Since(t1)

	t2 := time.Now()
	for i := 0; i < 30; i++ {
		c.FrameOnce()
	}
	steady := time.Since(t2) / 30

	t.Logf("mount=%v frame1=%v steady=%v", mount, frame1, steady)
	if mount > 50*time.Millisecond {
		t.Errorf("mount %v > 50ms", mount)
	}
	if frame1 > 50*time.Millisecond {
		t.Errorf("首帧 %v > 50ms", frame1)
	}
	if steady > 2*time.Millisecond {
		t.Errorf("稳态帧 %v > 2ms", steady)
	}
}
