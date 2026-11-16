package chart

import (
	"math"
	"testing"

	"github.com/energye/lcl/wg"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/ui"
)

func TestValueMapping(t *testing.T) {
	lo, hi := niceRange(3, 9)
	if lo >= 3 || hi <= 9 || lo < 0 {
		t.Fatalf("全正区间应两端留白且不跨负轴: [%v,%v]", lo, hi)
	}
	if got := yOf(6, lo, hi, 10, 110); got <= 10 || got >= 110 {
		t.Fatalf("中间值 y 应在画布内: %v", got)
	}
	if got := yOf(hi, lo, hi, 10, 110); got != 10 {
		t.Fatalf("上界应贴顶: %v", got)
	}
	if got := yOf(lo, lo, hi, 10, 110); got != 110 {
		t.Fatalf("下界应贴底: %v", got)
	}
	if x := xOf(2, 5, 40, 300); x != 40+2*((300-40)/4) {
		t.Fatalf("x 均匀分布计算错: %v", x)
	}
	if x := xOf(0, 1, 40, 300); x != 170 {
		t.Fatalf("单点应居中: %v", x)
	}
}

func TestNiceRangeEdgeCases(t *testing.T) {
	if lo, hi := niceRange(math.NaN(), 5); lo != 0 || hi != 1 {
		t.Fatalf("NaN 输入兜底: [%v,%v]", lo, hi)
	}
	if lo, hi := niceRange(5, 5); lo != 0 || hi <= 5 {
		t.Fatalf("常量正序列兜底: [%v,%v]", lo, hi)
	}
	if lo, hi := niceRange(-5, -5); lo >= -5 || hi != 0 {
		t.Fatalf("常量负序列兜底: [%v,%v]", lo, hi)
	}
}

func TestColorAtWraps(t *testing.T) {
	a, b := geom.RGBA(1, 2, 3), geom.RGBA(4, 5, 6)
	p := Palette{Series: []geom.Color{a, b}}
	if colorAt(p, 0) != a || colorAt(p, 1) != b || colorAt(p, 2) != a || colorAt(p, 3) != b {
		t.Fatal("色板应按序列数循环")
	}
	if colorAt(Palette{}, 5) != (geom.Color{}) {
		t.Fatal("空色板应返回零值（由调用方兜底）")
	}
}

func TestFormatTicks(t *testing.T) {
	if got := formatTick(3); got != "3" {
		t.Fatalf("整数刻度 = %q", got)
	}
	if got := formatTick(3.14); got != "3.1" {
		t.Fatalf("小数刻度 = %q", got)
	}
	if got := formatPercent(0.456); got != "46%" {
		t.Fatalf("百分比 = %q", got)
	}
}

// 每一类图表都通过真实 Container 链路出帧：Build → Layout → 录制 → Present。
// 任何自定义绘制里的 panic / 非法矩形都在这里暴露。
func TestChartsRenderThroughContainer(t *testing.T) {
	cases := []struct {
		name string
		w    ui.Widget
	}{
		{"line", LineChart{
			Series: [][]float64{{1, 3, 2, 5, 4}, {2, 2, 4, 3, 6}},
			Labels: []string{"一", "二", "三", "四", "五"},
			FillArea: true, ShowDots: true, AxisLabel: "趋势",
		}},
		{"bar", BarChart{
			Groups: [][]float64{{3, 5}, {4, 2}, {6, 3}},
			Labels: []string{"A", "B", "C"},
		}},
		{"pie", PieChart{Values: []float64{4, 3, 2, 1}, Labels: []string{"甲", "乙", "丙", "丁"}}},
		{"donut", PieChart{Values: []float64{6, 4}, Donut: true, Labels: []string{"完成", "剩余"}}},
		{"gauge", Gauge{Value: 0.72, Label: "72%"}},
		{"waveform", Waveform{Values: []float64{0, 1, -1, 0.5, -0.5, 0.2, 0}}},
		{"scatter", Scatter{Points: []ScatterPoint{
			{X: 1, Y: 2}, {X: 3, Y: 5}, {X: 4, Y: 1, Size: 4}, {X: 6, Y: 7},
		}}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := wg.NewHeadless(wg.DefaultOptions())
			c.Resize(400, 260)
			c.MountWidget(tc.w)
			if _, err := c.FrameOnce(); err != nil {
				t.Fatalf("%s 出帧失败: %v", tc.name, err)
			}
		})
	}
}

// 默认尺寸 / 空数据兜底不出错。
func TestChartsEmptyData(t *testing.T) {
	empty := []ui.Widget{
		LineChart{},
		BarChart{},
		PieChart{},
		Gauge{},
		Waveform{},
		Scatter{},
	}
	for i, w := range empty {
		c := wg.NewHeadless(wg.DefaultOptions())
		c.Resize(200, 160)
		c.MountWidget(w)
		if _, err := c.FrameOnce(); err != nil {
			t.Fatalf("空数据图表[%d] 出帧失败: %v", i, err)
		}
	}
}
