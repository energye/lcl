package wg

import (
	"testing"
)

// TestDPIChangedAppliesInContainer 锁定 DPI 变更的容器路径：
// 缩放记录更新、整树重建、以及非法值安全回落。
func TestDPIChangedAppliesInContainer(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(400, 300)
	p := &vpPage{}
	c.MountWidget(p)
	if p.builds != 1 {
		t.Fatalf("首次挂载 Build=%d", p.builds)
	}
	if got := c.Host().Scale(); got != 1 {
		t.Fatalf("初始 scale=%v，应为 1", got)
	}

	c.DPIChanged(2)
	if got := c.Host().Scale(); got != 2 {
		t.Fatalf("DPIChanged 后 scale=%v，应为 2", got)
	}
	if _, err := c.FrameOnce(); err != nil {
		t.Fatal(err)
	}
	if p.builds != 2 {
		t.Fatalf("DPI 变更应触发一次整树重建，Build=%d", p.builds)
	}

	// 非法值回落 1（NewDPI 语义），不崩。
	c.DPIChanged(0)
	if got := c.Host().Scale(); got != 1 {
		t.Fatalf("非法 scale 应回落 1，实际 %v", got)
	}

	// 关闭后调用是安全的 no-op。
	c.Close()
	c.DPIChanged(1.5)
}
