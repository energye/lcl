package host

import (
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/wg/geom"
)

// Host 把 IEngForm 绑定成一个画布宿主：持有根 Surface，跟踪客户区尺寸与 DPI。
//
// M0 阶段只落地"绑定 + 客户区 + 根画布"三件事；OnPaint 接管、WndProc 策略、
// 输入归一化（RawSink）、IME 与缓冲池在 M1/M2 接入（见 02-CANVAS.md）。
//
// 关键平台事实（继承自 examples/wg 的实证）：
//   - ICustomGraphicControl 无 HWND，其 Canvas 只在 OnPaint 内有效，
//     因此宿主必须在 OnPaint 里把像素一次性提交（单次不透明 CopyRect）。
//   - 窗口句柄只有 IEngForm 一个；Popup 不新增 HWND，由 Popup 画布实现。
type Host struct {
	form lcl.IEngForm

	root *Surface

	// client 是客户区像素尺寸。
	client geom.Size
	// scale 是画布缩放因子（96 DPI 为 1.0），用于 DIP ↔ 像素换算。
	scale float32
}

// New 创建宿主并建立根画布（Inline 种类，与窗口客户区同尺寸）。
//
// form 为 nil 时返回一个未绑定的 Host，Root 仍可用（尺寸为 0），
// 便于 headless 自检与单测：这类 Host 不会访问任何 LCL 符号。
func New(form lcl.IEngForm) *Host {
	h := &Host{form: form, scale: 1}
	if form != nil {
		h.client = geom.Sz(float32(form.ClientWidth()), float32(form.ClientHeight()))
		if s := form.GetCanvasScaleFactor(); s > 0 {
			h.scale = float32(s)
		}
	}
	h.root = NewSurface(SurfaceInline, geom.RectXYWH(0, 0, h.client.W, h.client.H))
	return h
}

// Form 返回绑定的窗体；headless 模式下为 nil。
func (h *Host) Form() lcl.IEngForm { return h.form }

// Bound 表示是否绑定了真实窗体。
func (h *Host) Bound() bool { return h.form != nil }

// Root 返回根画布。根画布不可 Detach，z 序为最底。
func (h *Host) Root() *Surface { return h.root }

// ClientSize 返回客户区尺寸（像素）。
func (h *Host) ClientSize() geom.Size { return h.client }

// Scale 返回画布缩放因子（96 DPI = 1）。
func (h *Host) Scale() float32 {
	if h.scale <= 0 {
		return 1
	}
	return h.scale
}

// DIPToPixel 把 DIP 换算为像素。
func (h *Host) DIPToPixel(v float32) float32 { return v * h.Scale() }

// PixelToDIP 把像素换算为 DIP。
func (h *Host) PixelToDIP(v float32) float32 { return v / h.Scale() }

// Resize 更新客户区尺寸并同步根画布（resize 去抖由 M1 的 dpi.go 负责，
// 这里只保证几何一致，不做防抖）。
func (h *Host) Resize(w, hgt int) {
	if w < 0 {
		w = 0
	}
	if hgt < 0 {
		hgt = 0
	}
	size := geom.Sz(float32(w), float32(hgt))
	if size == h.client {
		return
	}
	old := h.client
	h.client = size
	h.root.SetBounds(geom.RectXYWH(0, 0, size.W, size.H))
	// 契约 C20：窗口变大时新增区域必须重绘，否则残留旧像素。
	if size.W > old.W || size.H > old.H {
		h.root.Invalidate(h.root.LocalBounds())
	} else {
		h.root.Invalidate(geom.RectXYWH(0, 0, size.W, size.H))
	}
}

// RefreshScale 重新读取 DPI 缩放（DPI 变化时由 WndProc 调用）。
func (h *Host) RefreshScale() {
	if h.form == nil {
		return
	}
	if s := h.form.GetCanvasScaleFactor(); s > 0 {
		h.scale = float32(s)
	}
	h.root.Invalidate(h.root.LocalBounds())
}
