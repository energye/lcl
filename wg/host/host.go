package host

import (
	"sync"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
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

	// canvasHost 是真正承载像素的无 HWND 控件（TCustomGraphicControl）。
	//
	// 为什么需要它：直接对窗口 DC 调 GDI（GetDC + StretchDIBits）在 DWM
	// 重定向表面上会间歇性失败；LCL 的 TGraphicControl 本身没有句柄，
	// 它的 Canvas 只在 OnPaint 内有效，把像素提交放在这里既稳定又符合
	// "无句柄画布"的设计前提（见 examples/wg/showcase 的实证）。
	canvasHost lcl.ICustomGraphicControl

	// OnPaint 由容器注入：在控件重绘时提交最近一帧。
	OnPaint func(canvas lcl.ICanvas)

	// mu 保护 client/scale：WM_SIZE 在消息线程更新客户区，而容器的
	// resize 应用（c.Resize）在帧定时器 goroutine 上执行——两条路径都会
	// 调 Resize/ClientSize，必须互斥。
	mu sync.Mutex
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

// CanvasControl 返回承载像素的无句柄控件；headless 时为 nil。
func (h *Host) CanvasControl() lcl.ICustomGraphicControl { return h.canvasHost }

// BindCanvas 在窗体上创建无 HWND 的画布控件，并把 OnPaint 接到 h.OnPaint。
//
// 必须在窗体句柄存在后调用（RunOnMainThreadAsync 里）。它替代"裸 GDI 写
// 窗口 DC"：像素提交发生在控件自己的 OnPaint 内，Canvas 在那时保证有效。
func (h *Host) BindCanvas() error {
	if h.form == nil {
		return ErrNoForm
	}
	if h.canvasHost != nil {
		return nil
	}
	c := lcl.NewCustomGraphicControl(h.form)
	if c == nil || !c.IsValid() {
		return ErrNoForm
	}
	c.SetAlign(types.AlClient)
	c.SetParent(h.form)
	c.SetParentBackground(false)
	c.SetControlStyle(c.ControlStyle().Include(types.CsOpaque))
	c.SetOnPaint(func(sender lcl.IObject) {
		if h.OnPaint == nil {
			return
		}
		if cv := c.Canvas(); cv != nil && cv.IsValid() {
			h.OnPaint(cv)
		}
	})
	h.canvasHost = c
	return nil
}

// Repaint 请求画布控件重绘（下一轮 OnPaint 会提交最新帧）。
func (h *Host) Repaint() {
	if h.canvasHost != nil {
		h.canvasHost.Repaint()
	}
}

// Form 返回绑定的窗体；headless 模式下为 nil。
func (h *Host) Form() lcl.IEngForm { return h.form }

// Bound 表示是否绑定了真实窗体。
func (h *Host) Bound() bool { return h.form != nil }

// Root 返回根画布。根画布不可 Detach，z 序为最底。
func (h *Host) Root() *Surface { return h.root }

// ClientSize 返回客户区尺寸（像素）。
func (h *Host) ClientSize() geom.Size {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.client
}

// Scale 返回画布缩放因子（96 DPI = 1）。
func (h *Host) Scale() float32 {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.scale <= 0 {
		return 1
	}
	return h.scale
}

// DIPToPixel 把 DIP 换算为像素。
func (h *Host) DIPToPixel(v float32) float32 { return v * h.Scale() }

// PixelToDIP 把像素换算为 DIP。
func (h *Host) PixelToDIP(v float32) float32 { return v / h.Scale() }

// SetClientSize 只更新客户区尺寸记录，不触碰 Surface 树。
//
// 平台回调（WM_SIZE）在消息线程上调用它：Surface 树不是线程安全的
// （见包注释），结构性变更必须由 wg.Container 在 c.mu 下串行化；消息线程
// 写 Surface 会与帧线程的整帧合成竞争。根画布几何与失效由容器在
// Container.Resize 里统一应用。
func (h *Host) SetClientSize(w, hgt int) {
	if w < 0 {
		w = 0
	}
	if hgt < 0 {
		hgt = 0
	}
	h.mu.Lock()
	h.client = geom.Sz(float32(w), float32(hgt))
	h.mu.Unlock()
}

// Resize 更新客户区尺寸并同步根画布（resize 节流由 frame.Scheduler 负责，
// 这里只保证几何一致，不做防抖）。
//
// 必须在容器锁（c.mu）下调用：它修改 Surface 树（SetBounds/Invalidate），
// 与整帧合成互斥。
//
// 判定基准必须是**根画布当前 bounds**，不是 client：平台回调已经在
// WM_SIZE 里把 client 更新到新值（SetClientSize），拿 client 比较会让这里
// 永远早退——根画布 bounds 停在初始尺寸（根失效/裁剪全错）。
func (h *Host) Resize(w, hgt int) {
	if w < 0 {
		w = 0
	}
	if hgt < 0 {
		hgt = 0
	}
	size := geom.Sz(float32(w), float32(hgt))
	h.mu.Lock()
	h.client = size
	h.mu.Unlock()

	b := h.root.LocalBounds()
	if b.W() == size.W && b.H() == size.H {
		return
	}
	oldW, oldH := b.W(), b.H()
	h.root.SetBounds(geom.RectXYWH(0, 0, size.W, size.H))
	// 契约 C20：窗口变大时新增区域必须重绘，否则残留旧像素。
	if size.W > oldW || size.H > oldH {
		h.root.Invalidate(h.root.LocalBounds())
	} else {
		h.root.Invalidate(geom.RectXYWH(0, 0, size.W, size.H))
	}
}

// SetScale 更新画布缩放因子（线程安全，只写记录）。
func (h *Host) SetScale(scale float32) {
	d := NewDPI(scale)
	h.mu.Lock()
	h.scale = d.Scale()
	h.mu.Unlock()
}

// RefreshScale 重新读取窗体 DPI 缩放（兼容入口）。
//
// **不再触碰 Surface 树**：DPI 变更的布局/重绘由容器在 c.mu 下统一应用
// （见 Container.DPIChanged）；在消息线程直接失效根画布会与帧线程竞争。
func (h *Host) RefreshScale() {
	if h.form == nil {
		return
	}
	if s := h.form.GetCanvasScaleFactor(); s > 0 {
		h.SetScale(float32(s))
	}
}
