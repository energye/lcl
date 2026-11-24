package widgets

import (
	"testing"
	"time"

	"github.com/energye/lcl/wg/overlay"
	"github.com/energye/lcl/wg/scene"
	"github.com/energye/lcl/wg/theme"
	"github.com/energye/lcl/wg/ui"
)

// 用真实的 overlay.Host + theme 构造 BuildContext，验证浮层控制器接线。
func feedbackCtx() (*ui.BuildContext, *overlay.Host) {
	h := overlay.New(overlay.Options{})
	return &ui.BuildContext{Theme: theme.Default(), Overlay: h}, h
}

func TestModalShowClose(t *testing.T) {
	ctx, host := feedbackCtx()
	m := NewModal("确认删除", Label{Text: "此操作不可撤销"})
	confirmed := false
	m.OnConfirm = func() { confirmed = true }
	m.Show(ctx)
	if !m.IsOpen() || host.Count() != 1 {
		t.Fatalf("Modal 未显示: open=%v count=%d", m.IsOpen(), host.Count())
	}
	if host.Entries()[0].Band != theme.BandModal || !host.Entries()[0].Modal || !host.Entries()[0].TrapFocus {
		t.Fatalf("Modal 应带 BandModal/遮罩/焦点域: %+v", host.Entries()[0])
	}
	m.Close()
	if m.IsOpen() || host.Count() != 0 {
		t.Fatal("Modal 未关闭")
	}
	// 确认按钮的回调在下一帧的内容里；这里只验证状态机
	_ = confirmed
}

func TestModalShowTwiceIgnored(t *testing.T) {
	ctx, host := feedbackCtx()
	m := NewModal("标题", nil)
	m.Show(ctx)
	m.Show(ctx)
	if host.Count() != 1 {
		t.Fatalf("重复 Show 产生了 %d 条浮层", host.Count())
	}
}

func TestDrawerEdges(t *testing.T) {
	for _, edge := range []DrawerEdge{DrawerLeft, DrawerRight, DrawerTop, DrawerBottom} {
		ctx, host := feedbackCtx()
		d := NewDrawer("设置", Label{Text: "内容"})
		d.Edge = edge
		d.Show(ctx)
		if !d.IsOpen() || host.Count() != 1 {
			t.Fatalf("Drawer(edge=%d) 未显示", edge)
		}
		e := host.Entries()[0]
		if e.Band != theme.BandModal || !e.Modal {
			t.Fatalf("Drawer(edge=%d) 应有模态遮罩", edge)
		}
		if e.Dismiss.Has(overlay.DismissEsc) == false {
			t.Fatalf("Drawer(edge=%d) 应支持 Esc 关闭", edge)
		}
		d.Close()
	}
}

func TestDrawerFooter(t *testing.T) {
	tests := []struct {
		name       string
		footer     ui.Widget
		footerNull bool
		wantFooter bool
	}{
		{name: "默认无 footer"},
		{name: "自定义 footer", footer: NewButton("自定义"), wantFooter: true},
		{name: "FooterNull 优先", footer: NewButton("自定义"), footerNull: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, host := feedbackCtx()
			d := NewDrawer("设置", Label{Text: "内容"})
			d.Footer = tt.footer
			d.FooterNull = tt.footerNull
			d.Show(ctx)
			root := host.Entries()[0].Build()
			var hasFooter, hasCustom bool
			root.Walk(func(n *scene.Node) bool {
				if n.HitSpec().Label == "drawer-footer" {
					hasFooter = true
				}
				if spec := n.Text(); spec != nil && spec.Src == "自定义" {
					hasCustom = true
				}
				return true
			})
			if hasFooter != tt.wantFooter {
				t.Fatalf("drawer-footer=%v, want %v", hasFooter, tt.wantFooter)
			}
			if hasCustom != tt.wantFooter {
				t.Fatalf("自定义 footer 内容=%v, want %v", hasCustom, tt.wantFooter)
			}
		})
	}
}

func TestNotificationTimeout(t *testing.T) {
	ctx, host := feedbackCtx()
	n := NewNotification("保存成功", "数据已写入")
	n.Duration = 2 * time.Second
	n.Show(ctx)
	if !n.IsOpen() || host.Count() != 1 {
		t.Fatal("Notification 未显示")
	}
	e := host.Entries()[0]
	if e.Band != theme.BandToast || !e.Dismiss.Has(overlay.DismissTimeout) || e.Duration != 2*time.Second {
		t.Fatalf("通知应有 BandToast/超时关闭/时长: %+v", e)
	}
	n.Close()
	if n.IsOpen() {
		t.Fatal("通知未关闭")
	}
}

func TestPopconfirmAnchor(t *testing.T) {
	ctx, host := feedbackCtx()
	btn := NewButton("删除")
	pc := NewPopconfirm("确定删除？", btn)
	n := pc.Build(ctx)
	if n == nil {
		t.Fatal("Popconfirm 构建失败")
	}
	// 直接触发 show（点击路径由手势处理器覆盖，这里验证浮层声明）
	pc.show(ctx, n)
	if !pc.IsOpen() || host.Count() != 1 {
		t.Fatal("Popconfirm 未显示")
	}
	e := host.Entries()[0]
	if e.Band != theme.BandPopup || e.Anchor.Kind != overlay.AnchorNode || !e.Dismiss.Has(overlay.DismissOutsideClick) {
		t.Fatalf("Popconfirm 应锚定节点且外部点击关闭: %+v", e)
	}
	pc.Close()
	if pc.IsOpen() {
		t.Fatal("Popconfirm 未关闭")
	}
}

// 无 overlay 时 Show 必须安全 no-op（组件能在纯 headless 单测里用）。
func TestFeedbackWithoutOverlayNoPanic(t *testing.T) {
	ctx := &ui.BuildContext{Theme: theme.Default()}
	NewModal("t", nil).Show(ctx)
	NewDrawer("t", nil).Show(ctx)
	NewNotification("t", "b").Show(ctx)
	pc := NewPopconfirm("t", Label{Text: "x"})
	pc.Build(ctx)
	pc.show(ctx, nil)
}

// TestModalLoading 验证 loading 语义：内容区替换为骨架屏、footer 按钮隐藏、
// 标题保留（antd loading：title 与 body 独立，仅替换 body 并隐藏 footer）。
func TestModalLoading(t *testing.T) {
	tests := []struct {
		name     string
		newModal func() *Modal
	}{
		{
			name: "普通",
			newModal: func() *Modal {
				return NewModal("加载中", Label{Text: "内容加载中"})
			},
		},
		{
			name: "预设",
			newModal: func() *Modal {
				return NewModalPreset(ModalInfo, "加载中", Label{Text: "内容加载中"}, nil)
			},
		},
		{
			name: "自定义 footer",
			newModal: func() *Modal {
				m := NewModal("加载中", Label{Text: "内容加载中"})
				m.Footer = NewButton("自定义")
				return m
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, host := feedbackCtx()
			m := tt.newModal()
			m.Loading = true
			m.Show(ctx)
			if !m.IsOpen() || host.Count() != 1 {
				t.Fatalf("Modal 未显示: open=%v count=%d", m.IsOpen(), host.Count())
			}
			root := host.Entries()[0].Build()
			if root == nil {
				t.Fatal("Modal 内容未构建")
			}
			var hasSkeleton, hasTitle bool
			buttonCount := 0
			root.Walk(func(n *scene.Node) bool {
				switch n.HitSpec().Label {
				case "skeleton":
					hasSkeleton = true
				case "modal-title":
					hasTitle = true
				case "button":
					buttonCount++
				}
				return true
			})
			if !hasSkeleton {
				t.Error("loading=true 时正文应出现骨架屏")
			}
			if !hasTitle {
				t.Error("loading=true 时标题应保留")
			}
			if buttonCount != 0 {
				t.Errorf("loading=true 时 footer 按钮应隐藏，实际 %d 个", buttonCount)
			}
		})
	}
}
