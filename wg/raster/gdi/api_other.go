//go:build !windows

package gdi

import (
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/raster"
	"github.com/energye/lcl/wg/render"
	"github.com/energye/lcl/wg/text"
)

// 非 Windows 平台没有 GDI：本包保留同一套导出符号，让上层代码可以
// 无条件引用（composite 只按 Capabilities 决定是否走文字通道）。
//
// 行为约定：NewDIB 返回 nil，TextPass 返回 0，Backend.Draw 把整表计为
// degraded —— 明确"做不到"，而不是假装画了（C22）。

// DIB 是不可用占位类型。
type DIB struct {
	W, H, Stride, Bpp int
}

// NewDIB 在非 Windows 平台始终返回 nil。
func NewDIB(w, h, bpp int) *DIB { return nil }

// Bits 返回 nil。
func (d *DIB) Bits() []byte { return nil }

// Row 返回 nil。
func (d *DIB) Row(y int) []byte { return nil }

// Flush 无操作。
func (d *DIB) Flush() {}

// Clear 无操作。
func (d *DIB) Clear(color uint32) {}

// Delete 无操作。
func (d *DIB) Delete() {}

// TextPass 在非 Windows 平台不执行任何绘制。
func TextPass(d *DIB, dl *render.DisplayList, clip geom.Rect) int { return 0 }

// MeasureText 退回 text 包的估算度量。
func MeasureText(s string, f text.Font) geom.Size { return text.Measure(s, f) }

// Backend 是非 Windows 的占位后端：全部指令降级。
type Backend struct {
	caps raster.Caps
}

// New 创建占位后端。
func New() *Backend {
	return &Backend{caps: raster.Caps{ClipRect: true}}
}

// Name 实现 raster.Backend。
func (b *Backend) Name() string { return "gdi(unavailable)" }

// Caps 实现 raster.Backend。
func (b *Backend) Caps() raster.Caps { return b.caps }

// TextOps 返回 0。
func (b *Backend) TextOps() int { return 0 }

// Draw 把整表计为降级。
func (b *Backend) Draw(dst *raster.Backing, dl *render.DisplayList, clip geom.Rect) raster.Stats {
	if dl == nil {
		return raster.Stats{}
	}
	return raster.Stats{Degraded: dl.Len()}
}

// Close 无操作。
func (b *Backend) Close() {}
