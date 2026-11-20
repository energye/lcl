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

// 临时 dump 探针（提交前删除）：点击按钮后 dump 浮层 message 节点子树。

func TestMessageDumpProbe(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1080, 760)

	var ctxRef *ui.BuildContext
	var infoBtn *widgets.Button
	msg := widgets.NewMessage(widgets.MessageInfo, "这是一条信息提示")
	c.MountWidget(doc.Page(doc.Meta{
		ID: "message", Title: "Message", Group: doc.GroupFeedback,
		What: []string{"dump probe"},
		Demos: []doc.Demo{
			{Title: "触发", Build: func(ctx *ui.BuildContext) ui.Widget {
				ctxRef = ctx
				infoBtn = widgets.NewButton("显示信息")
				infoBtn.OnClick = func() {
					if ctxRef != nil && ctxRef.Overlay != nil {
						msg.Show(ctxRef)
					}
				}
				return infoBtn
			}},
		},
	}))
	c.FrameOnce()

	// 直接调 Show（等价点击效果），然后 dump
	if ctxRef != nil && ctxRef.Overlay != nil {
		msg.Show(ctxRef)
	}
	time.Sleep(100 * time.Millisecond)
	c.FrameOnce()

	found := 0
	var dump func(nd *scene.Node, depth int)
	dump = func(nd *scene.Node, depth int) {
		h := nd.HitSpec()
		label := h.Label
		if label == "" && nd.Kind() == scene.KindText {
			if sp := nd.Text(); sp != nil {
				label = "TEXT:" + sp.Src
			}
		}
		if label != "" || depth < 12 {
			prefix := ""
			for i := 0; i < depth; i++ {
				prefix += "  "
			}
			if label != "" {
				t.Logf("%s%s [%s]", prefix, nd.Kind(), label)
				found++
			}
		}
		for _, ch := range nd.Children() {
			dump(ch, depth+1)
		}
	}
	c.Walk(func(nd *scene.Node) bool {
		if nd.HitSpec().Label == "message" {
			t.Logf("=== message 节点 dump ===")
			dump(nd, 0)
			return false
		}
		return true
	})
	t.Logf("found=%d overlayCount=%d", found, func() int {
		if ov := c.Overlay(); ov != nil {
			return ov.Count()
		}
		return -1
	}())
	_ = infoBtn
}
