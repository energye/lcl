package widgets

import (
	"testing"
)

func splitterPanels() []SplitterPanel {
	return []SplitterPanel{
		{Content: nil, DefaultSize: 200},
		{Content: nil},
	}
}

// TestSplitterInitSizes 初始尺寸：DefaultSize + 平分剩余。
func TestSplitterInitSizes(t *testing.T) {
	s := NewSplitter(splitterPanels()...)
	s.Width = 600
	n := s.Build(dataCtx())
	if n == nil {
		t.Fatal("Build 失败")
	}
	// 总 600 - 1 个把手×2 = 598；左 200 → 右 398。
	if s.sizes[0] != 200 || s.sizes[1] != 398 {
		t.Fatalf("sizes = %v，应 [200 398]", s.sizes)
	}
}

// TestSplitterEvenSplit 无 DefaultSize 平分。
func TestSplitterEvenSplit(t *testing.T) {
	s := NewSplitter(SplitterPanel{}, SplitterPanel{}, SplitterPanel{})
	s.Width = 600
	s.Build(dataCtx())
	// 600 - 2×2 = 596 / 3 ≈ 198.67。
	if s.sizes[0] <= 0 || s.sizes[1] <= 0 || s.sizes[2] <= 0 {
		t.Fatalf("平分尺寸异常: %v", s.sizes)
	}
	sum := s.sizes[0] + s.sizes[1] + s.sizes[2]
	if sum < 595 || sum > 598 {
		t.Fatalf("总和不等于可用长: %v", sum)
	}
}

// TestSplitterClamp 拖拽钳制 min/max。
func TestSplitterClamp(t *testing.T) {
	s := NewSplitter(
		SplitterPanel{DefaultSize: 300, Min: 100, Max: 400},
		SplitterPanel{DefaultSize: 300, Min: 100},
	)
	s.Width = 600
	s.Build(dataCtx())
	s.dragResize(0, -500) // 左拽到 300-500=-200 → 钳到 Min 100；右相应 +? ——右侧 298
	if s.sizes[0] != 100 {
		t.Fatalf("左面板应钳到 Min 100，实际 %v", s.sizes[0])
	}
	s.dragResize(0, 500) // 右拽：左 600？→ Max 400；右 198（Min 100 内）
	if s.sizes[0] != 400 {
		t.Fatalf("左面板应钳到 Max 400，实际 %v", s.sizes[0])
	}
}

// TestSplitterSetSizes SetSizes 受控驱动（OnResize 回调+钳制）。
func TestSplitterSetSizes(t *testing.T) {
	s := NewSplitter(
		SplitterPanel{DefaultSize: 300, Min: 150},
		SplitterPanel{DefaultSize: 300},
	)
	s.Width = 600
	s.Build(dataCtx())
	called := 0
	var got []float32
	s.OnResize = func(sz []float32) { called++; got = sz }
	s.SetSizes([]float32{100, 500}) // 100 < Min 150 → 钳到 150；末面板吃余差 448
	if called != 1 {
		t.Fatalf("OnResize 未回调")
	}
	if got[0] != 150 || got[1] != 448 {
		t.Fatalf("SetSizes 钳制后 = %v，应 [150 448]", got)
	}
	// Sizes() 返回副本。
	sz := s.Sizes()
	sz[0] = 999
	if s.sizes[0] == 999 {
		t.Fatal("Sizes() 应返回副本")
	}
}

// TestSplitterVertical 垂直分栏主轴长取 Height。
func TestSplitterVertical(t *testing.T) {
	s := NewSplitter(SplitterPanel{}, SplitterPanel{})
	s.Orientation = SplitVertical
	s.Height = 400
	s.Build(dataCtx())
	// 400 - 2 = 398 / 2 = 199。
	if s.sizes[0] < 198 || s.sizes[0] > 200 {
		t.Fatalf("垂直平分 = %v，应 ≈199", s.sizes[0])
	}
}
