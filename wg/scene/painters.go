package scene

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
)

// ── 画笔便捷构造 ────────────────────────────────────────────────────
//
// BoxSpec.Fill / Border 的类型就是 render.Paint，但**调用方不该为了填个色
// 去 import render**：
//   - 依赖收窄（C17）：overlay、ui、业务组件只需要 scene + theme；
//   - 心智：组件作者说的是"这框填这个色"，不是"给我一个 Paint 结构"。
//
// 这里只做直通转发，不引入任何语义。真要渐变/画刷时再直接用 render.Paint
// （scene 已经允许被更高层依赖，无需绕路）。

// Solid 返回实心填充画笔。
func Solid(c geom.Color) render.Paint { return render.Fill(c) }

// Stroke 返回描边画笔（width<=0 时退化为 1px，避免"设了颜色却看不见"）。
func Stroke(c geom.Color, width float32) render.Paint {
	if width <= 0 {
		width = 1
	}
	return render.Stroke(c, width)
}

// Clear 返回"什么都不画"的画笔。
//
// 用途：BoxSpec.Fill 的零值本身就是不画，但表达"显式清空"时不希望
// 调用方写 `render.Paint{}`（那要 import render）。
func Clear() render.Paint { return render.Paint{} }
