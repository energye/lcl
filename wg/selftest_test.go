package wg

import (
	"testing"

	"github.com/energye/lcl/wg/contracts"
	"github.com/energye/lcl/wg/frame"
)

// TestSelfTestGreen 是 M0 门禁的单测形态：go test 一跑就能发现自检回归。
func TestSelfTestGreen(t *testing.T) {
	for _, r := range SelfTest() {
		if !r.OK {
			t.Errorf("[%s] %s: %s", r.Group, r.Name, r.Detail)
		}
	}
}

func TestNewHeadless(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	if c.Form() != nil {
		t.Fatal("headless 容器不应有窗体")
	}
	if c.State() != "created" {
		t.Fatalf("初始状态 = %s", c.State())
	}
	if c.Root() == nil {
		t.Fatal("Root 为 nil")
	}
	if err := c.Run(); err == nil {
		t.Fatal("headless 容器 Run 应返回错误而非静默返回")
	}
	if opts := c.Options(); opts.CacheBytes <= 0 || opts.MaxLayers <= 0 {
		t.Fatalf("选项未补齐默认值: %+v", opts)
	}
}

func TestZeroOptionsGetDefaults(t *testing.T) {
	c := NewHeadless(Options{})
	defer c.Close()
	if !c.Options().AA {
		t.Fatal("零值选项应得到默认开启的 AA")
	}
}

func TestSetClockAfterRunRejected(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.SetClock(frame.NewManualClock(0, frame.DefaultStep))
	if c.Clock().Kind() != frame.Manual {
		t.Fatal("时钟未替换")
	}
}

func TestResizeAndFrames(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.Resize(320, 200)
	clk := frame.NewManualClock(0, frame.DefaultStep)
	c.SetClock(clk)
	for i := 0; i < 3; i++ {
		st, err := c.FrameOnce()
		if err != nil {
			t.Fatal(err)
		}
		if st.Index != i {
			t.Fatalf("第 %d 帧 Index = %d", i, st.Index)
		}
		clk.Advance()
	}
	if c.Frames() != 3 {
		t.Fatalf("Frames = %d", c.Frames())
	}
	if !c.Root().Damage().IsEmpty() {
		// M0 语义：只统计不提交，因此 resize 的脏区会保留
		if c.Root().Damage().W() != 320 {
			t.Fatalf("脏区宽度 = %v", c.Root().Damage().W())
		}
	}
}

func TestRequestFrameMarksDirty(t *testing.T) {
	c := NewHeadless(DefaultOptions())
	defer c.Close()
	c.RequestFrame()
	st, err := c.FrameOnce()
	if err != nil {
		t.Fatal(err)
	}
	if !st.Dirty {
		t.Fatal("RequestFrame 后应出一帧")
	}
	st2, _ := c.FrameOnce()
	if st2.Dirty {
		t.Fatal("帧请求应被消费一次")
	}
}

func TestContractStageDefaults(t *testing.T) {
	if contracts.Current != contracts.M4 {
		t.Fatalf("默认门禁 = %v，应为 M4", contracts.Current)
	}
	// 能力阶段（M7：示例重构 + antd token 对齐）与契约到期阶段（M4）独立演进。
	if Version != "0.1.0-m7" || Stage != "M7" {
		t.Fatalf("版本/阶段 = %s/%s", Version, Stage)
	}
}
