package widgets

import "testing"

// TestGridBreakpointSpan 断点解析：命中最大断点、未设回退。
func TestGridBreakpointSpan(t *testing.T) {
	i := func(v int) *int { return &v }
	c := GridCol{Span: 6, Md: i(12), Xs: i(24)}

	cases := []struct {
		vw       float32
		wantSpan int
	}{
		{1280, 12}, // ≥lg 未设，回退到 md 的 12（antd 断点延续）
		{1000, 12}, // lg 未设，回退 md → 12
		{800, 12},  // md 命中 → 12
		{600, 24},  // sm 未设，回退到 xs → 24
		{400, 24},  // xs 命中 → 24
	}
	for _, tc := range cases {
		span, _ := breakpointSpan(c, tc.vw)
		if span != tc.wantSpan {
			t.Errorf("vw=%v span=%d，want %d", tc.vw, span, tc.wantSpan)
		}
	}
	// 默认 Span（0）按 24。
	span, _ := breakpointSpan(GridCol{}, 1280)
	if span != 24 {
		t.Fatalf("空 span 应默认 24，got %d", span)
	}
	// Offset 透传。
	span, off := breakpointSpan(GridCol{Span: 8, Offset: 4}, 1280)
	if span != 8 || off != 4 {
		t.Fatalf("span=%d offset=%d，want 8/4", span, off)
	}
}

// TestGridBuild 栅格构建：行内 24 栅格换行。
func TestGridBuild(t *testing.T) {
	g := NewGrid(
		GridCol{Child: Label{Text: "a"}, Span: 12},
		GridCol{Child: Label{Text: "b"}, Span: 12},
		GridCol{Child: Label{Text: "c"}, Span: 12},
	)
	n := g.Build(dataCtx())
	if n == nil {
		t.Fatal("Grid 构建失败")
	}
	// 两行（12+12 换行 + 单 12）。
	if len(n.Children()) != 2 {
		t.Fatalf("行数=%d，want 2（12+12 换行 + 12）", len(n.Children()))
	}
}
