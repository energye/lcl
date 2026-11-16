package wg

import (
	"fmt"
	"testing"

	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// countingWidget 记录 Build 次数，用来验证"SetState 后下一帧自动重建"。
type countingWidget struct {
	ui.Base
	label string
}

func (w *countingWidget) Build(*ui.BuildContext) *scene.Node {
	n := scene.Box(scene.BoxSpec{}).SetSize2(80, 24)
	n.SetKey(fmt.Sprintf("%s#%d", w.label, w.Builds()))
	return n
}

func newWidgetContainer(t *testing.T) *Container {
	t.Helper()
	c := NewHeadless(DefaultOptions())
	c.Resize(320, 200)
	return c
}

func TestMountWidgetBuildsAndRebuilds(t *testing.T) {
	c := newWidgetContainer(t)
	defer c.Close()
	w := &countingWidget{label: "w"}
	c.MountWidget(w)
	if c.Widget() != w {
		t.Fatal("Widget() 未返回挂载的组件")
	}
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if w.Builds() != 1 {
		t.Fatalf("首次 Build 次数 = %d", w.Builds())
	}
	if c.Content() == nil {
		t.Fatal("构建结果未成为内容树")
	}

	w.SetState(func() {})
	if !w.Dirty() {
		t.Fatal("SetState 未标脏")
	}
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if w.Builds() != 2 || w.Dirty() {
		t.Fatalf("重建后 Build=%d dirty=%v", w.Builds(), w.Dirty())
	}
}

func TestMountReplacesWidget(t *testing.T) {
	c := newWidgetContainer(t)
	defer c.Close()
	c.MountWidget(&countingWidget{label: "w"})
	c.Mount(scene.Box(scene.BoxSpec{}).SetSize2(10, 10))
	if c.Widget() != nil {
		t.Fatal("直接 Mount 应替换声明式组件")
	}
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
}

func TestSetThemeRebuildsWidget(t *testing.T) {
	c := newWidgetContainer(t)
	defer c.Close()
	w := &countingWidget{label: "w"}
	c.MountWidget(w)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	th := theme.Default()
	th.Font.SizeBase = 20
	c.SetTheme(th)
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if w.Builds() != 2 {
		t.Fatalf("换主题后 Build 次数 = %d，应重建（字号是度量）", w.Builds())
	}
}

func TestInheritedReachesBuildContext(t *testing.T) {
	c := newWidgetContainer(t)
	defer c.Close()
	type marker struct{}
	c.SetInherited(marker{}, "value")
	got := ""
	c.MountWidget(&probeWidget{read: func(ctx *ui.BuildContext) {
		if v, ok := ctx.Inherited(marker{}).(string); ok {
			got = v
		}
	}})
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if got != "value" {
		t.Fatalf("Inherited 未传到组件: %q", got)
	}
}

type probeWidget struct {
	ui.Base
	read func(*ui.BuildContext)
}

func (w *probeWidget) Build(ctx *ui.BuildContext) *scene.Node {
	w.read(ctx)
	return scene.Empty().SetSize2(10, 10)
}
