package main

import (
	"testing"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/examples/doc"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/ui"
	"github.com/energye/lcl/wg/widgets"
)

// 临时 dump 探针（提交前删除）：浮层菜单项完整 bounds。

func TestDropdownDumpProbe(t *testing.T) {
	c := wg.NewHeadless(wg.DefaultOptions())
	defer c.Close()
	c.Resize(1080, 1100)

	var dd *widgets.Dropdown
	var childNode *scene.Node
	var ctxRef *ui.BuildContext
	c.MountWidget(doc.Page(doc.Meta{
		ID: "dropdown", Title: "Dropdown", Group: doc.GroupControl,
		What: []string{"dump probe"},
		Demos: []doc.Demo{
			{Title: "下拉", Build: func(ctx *ui.BuildContext) ui.Widget {
				ctxRef = ctx
				trigger := widgets.NewButton("悬停下拉")
				dd = widgets.NewDropdown(trigger,
					&widgets.MenuItem{Key: "copy", Label: "复制"},
					&widgets.MenuItem{Key: "edit", Label: "编辑"})
				childNode = dd.Build(ctx)
				return doc.Node(func(*ui.BuildContext) *scene.Node { return childNode })
			}},
		},
	}))
	c.FrameOnce()

	// 直调 Show（anchor=childNode）。
	dd.Show(ctxRef)
	c.FrameOnce()

	c.Walk(func(nd *scene.Node) bool {
		h := nd.HitSpec()
		if h.Label == "dropdown-item" || h.Label == "dropdown" {
			t.Logf("%q kind=%v bounds=%v", h.Label, nd.Kind(), nd.BoundsInRoot())
		}
		return true
	})
}
