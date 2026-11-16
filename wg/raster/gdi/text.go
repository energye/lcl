//go:build windows

package gdi

import (
	"syscall"
	"unsafe"

	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/text"
)

// TextPass 在**像素写回之后**于设备空间回放文字（契约 C6/C24）。
//
// 这是整个渲染管线里唯一"不是纯 Go"的一步，也是刻意为之：
//   - ClearType 依赖 GDI 对设备像素的次像素定位，软件实现拿不到同样的结果；
//   - 文字必须画在最终 DIB 上（而不是中间 ARGB 缓冲上），否则无法做
//     亚像素抗锯齿——24bpp 的 DIB 正是 ClearType 的工作格式。
//
// 调用顺序在 composite 中固定：形状光栅 → 拷贝到呈现 DIB → 本函数 → 提交。
//
// 性能要点（M1 实测）：
//   - **只在变化时**重设字体/颜色/对齐。反复 SelectObject(HFONT) 会让 GDI
//     反复丢弃并重建字体实现（字体缓存抖动），这是文字通道最大的坑；
//   - 裁剪交给 ETO_CLIPPED 的 lpRect，不再为每段文字 SaveDC+IntersectClipRect。
//
// 返回实际绘制的行数。
func TextPass(d *DIB, dl *render.DisplayList, clip geom.Rect) int {
	if d == nil || dl == nil || !dl.HasText() {
		return 0
	}
	drawn := 0

	// 全程不变的两个状态，设一次即可
	procSetBkMode.Call(d.HDC, transparent)
	procSetTextAlign.Call(d.HDC, taLeft|taBaseline)

	var (
		curFont  uintptr
		curColor uint32 = 1 << 31 // 不可能等于任何 COLORREF，保证首个颜色一定被设置
		origFont uintptr
	)

	for _, cmd := range dl.Commands() {
		if cmd.Op != render.OpDrawText {
			continue
		}
		item := dl.Text(cmd.Text)
		if item == nil || item.Layout.Empty() {
			continue
		}
		area := cmd.Rect.Intersect(clip)
		if area.IsEmpty() {
			continue
		}

		if hf := text.FontHandle(item.Layout.Font); hf != 0 && hf != curFont {
			prev, _, _ := procSelectObject.Call(d.HDC, hf)
			if origFont == 0 {
				origFont = prev // DC 原有字体，结束时换回去
			}
			curFont = hf
		}
		// COLORREF 是 0x00BBGGRR
		if cr := uint32(item.Color.B)<<16 | uint32(item.Color.G)<<8 | uint32(item.Color.R); cr != curColor {
			procSetTextColor.Call(d.HDC, uintptr(cr))
			curColor = cr
		}

		scale := cmd.Matrix.ScaleFactor()
		if scale <= 0 {
			scale = 1
		}
		rc := rect{
			Left: int32(area.Min.X), Top: int32(area.Min.Y),
			Right: int32(area.Max.X), Bottom: int32(area.Max.Y),
		}
		l := item.Layout
		for i, line := range l.Lines {
			if line.Text == "" {
				continue
			}
			x := int32(item.Origin.X + l.LineX(i)*scale)
			y := int32(item.Origin.Y + line.Baseline*scale)
			u := syscall.StringToUTF16(line.Text)
			n := len(u) - 1 // 去掉结尾的 0
			if n <= 0 {
				continue
			}
			ok, _, _ := procExtTextOutW.Call(
				d.HDC, uintptr(x), uintptr(y), etoClipped,
				uintptr(unsafe.Pointer(&rc)),
				uintptr(unsafe.Pointer(&u[0])),
				uintptr(n), 0,
			)
			if ok != 0 {
				drawn++
			}
		}
	}

	if origFont != 0 {
		procSelectObject.Call(d.HDC, origFont)
	}
	return drawn
}

// MeasureText 用 GDI 度量文本（供布局在无 LCL 环境下使用）。
func MeasureText(s string, f text.Font) geom.Size { return text.Measure(s, f) }
