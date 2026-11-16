package geom

import "math"

// 本文件集中放置浮点小工具，避免在多个文件里重复。
// 特意不用 Go 1.21 的内建 min/max：lcl 模块的 go.mod 声明 go 1.20。

func minf(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func maxf(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func sqrt(v float32) float32 { return float32(math.Sqrt(float64(v))) }

func sinf(v float32) float32 { return float32(math.Sin(float64(v))) }

func cosf(v float32) float32 { return float32(math.Cos(float64(v))) }

func floorf(v float32) float32 { return float32(math.Floor(float64(v))) }

func ceilf(v float32) float32 { return float32(math.Ceil(float64(v))) }
