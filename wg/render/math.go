package render

import "math"

// 三角函数集中在这里：render 需要它做圆弧近似，但不值得为几个函数
// 让调用方背 math 的 float64 转换开销（编译器会把常量折叠掉一部分）。

func cosf(v float32) float32 { return float32(math.Cos(float64(v))) }
func sinf(v float32) float32 { return float32(math.Sin(float64(v))) }
func tanf(v float32) float32 { return float32(math.Tan(float64(v))) }
