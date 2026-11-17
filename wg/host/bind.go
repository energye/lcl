//go:build windows

package host

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/wg/geom"
	"github.com/energye/lcl/wg/input"
	"github.com/energye/lcl/wg/scene"
)

// ── LCL 事件绑定 ────────────────────────────────────────────────────
//
// 这一层只做平台动作，不含任何决策：
//   - 接管 SetOnWndProc 并把消息交给 Router；
//   - 注册 TrackMouseEvent（否则永远收不到 WM_MOUSELEAVE）；
//   - 维护 OS 级捕获（SetCapture）：不设捕获的话"拖出窗口"就收不到移动，
//     逻辑层的指针捕获再正确也没用——事件根本没进来（C26 的前提）；
//   - GetKeyState 取修饰键（wParam 里的修饰位只在鼠标消息里有）。
//
// 所有判断都在 Router/Normalizer 里，因此无窗口自检能覆盖行为，
// 这层剩下的风险只有"平台 API 用错"。

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	procGetMessageTime  = user32.NewProc("GetMessageTime")
	procGetKeyState     = user32.NewProc("GetKeyState")
	procSetCapture      = user32.NewProc("SetCapture")
	procReleaseCapture  = user32.NewProc("ReleaseCapture")
	procTrackMouseEvent = user32.NewProc("TrackMouseEvent")
	procClientToScreen  = user32.NewProc("ClientToScreen")
	procScreenToClient  = user32.NewProc("ScreenToClient")
	procGetDC           = user32.NewProc("GetDC")
	procReleaseDC       = user32.NewProc("ReleaseDC")
	procInvalidateRect  = user32.NewProc("InvalidateRect")
	procPostMessage     = user32.NewProc("PostMessageW")
	procGetUpdateRect   = user32.NewProc("GetUpdateRect")
)

// clientUpdateRect 返回窗口当前无效区（客户区像素坐标；空 = 无）。
func clientUpdateRect(hwnd uintptr) geom.Rect {
	var r winRect
	ok, _, _ := procGetUpdateRect.Call(hwnd, uintptr(unsafe.Pointer(&r)), 0)
	if ok == 0 || r.Right <= r.Left || r.Bottom <= r.Top {
		return geom.Rect{}
	}
	return geom.RectXYWH(float32(r.Left), float32(r.Top),
		float32(r.Right-r.Left), float32(r.Bottom-r.Top))
}

// wmRequestPresent 是私有呈现消息（WM_APP 区间）。
//
// 为什么必须：GetMessage 的取消息顺序是 sent > posted > 硬件输入 >
// WM_PAINT。帧泵在 goroutine 里合成，靠 InvalidateRect 等 WM_PAINT 会在
// 拖动/连续鼠标消息下被饿死（内容不跟窗）。posted 消息在模态尺寸循环里
// 也会被及时派发（showcase 的实证做法）。
const wmRequestPresent = 0x8000 + 0x21

// PostPresent 请求主线程尽快呈现（在帧泵 goroutine 调用；Presenter 自带
// 合并标记，一条消息对应最新像素）。
func PostPresent(hwnd uintptr) {
	if hwnd != 0 {
		procPostMessage.Call(hwnd, wmRequestPresent, 0, 0)
	}
}

// trackMouseEvent 是 TRACKMOUSEEVENT。
type trackMouseEvent struct {
	size      uint32
	flags     uint32
	hwnd      uintptr
	hoverTime uint32
}

const tmeLeave = 0x00000002

// Binder 把 LCL 窗体接到输入链路。
type Binder struct {
	host   *Host
	norm   *Normalizer
	router *Router
	form   lcl.IEngForm
	hwnd   uintptr
	ime    *IMEBridge
	// presenter 在 WM_PAINT 中把缓存帧提交到窗口 DC（呈现链路最后一公里）。
	presenter *Presenter

	installed bool
	tracking  bool
	captured  bool
	lastCur   scene.Cursor
	cursorSet bool
}

// NewBinder 创建绑定器。
func NewBinder(h *Host, sink EventSink) *Binder {
	norm := NewNormalizer(NewDPI(h.Scale()))
	r := NewRouter(norm, sink)
	b := &Binder{host: h, norm: norm, router: r, form: h.Form()}
	r.SetScreenToClient(b.screenToClient)
	// 只记录客户区尺寸；Surface 几何由容器在 c.mu 下应用（见 Host.SetClientSize）。
	r.OnResize = func(w, hgt int) { b.host.SetClientSize(w, hgt) }
	r.OnDPI = func(d DPI) {
		b.norm.SetDPI(d)
		b.host.RefreshScale()
	}
	b.ime = NewIMEBridge(nil)
	if b.form != nil {
		b.hwnd = uintptr(b.form.Handle())
		r.SetIME(newPlatformIME(b.hwnd, b.ime), b.ime)
	}
	return b
}

// Router 返回消息路由。
func (b *Binder) Router() *Router { return b.router }

// IME 返回桥实例；容器把当前聚焦输入框的 TextEditor 交给它。
func (b *Binder) IME() *IMEBridge { return b.ime }

// SetIMEBridge 替换桥实例（容器持有长期实例，Binder 只在安装时借用）。
func (b *Binder) SetIMEBridge(br *IMEBridge) {
	if br == nil {
		return
	}
	b.ime = br
	if b.form != nil {
		b.router.SetIME(newPlatformIME(b.hwnd, br), br)
	}
}

// SetPresenter 绑定呈现器：WM_PAINT 时由它把缓存帧提交到窗口 DC。
func (b *Binder) SetPresenter(p *Presenter) { b.presenter = p }

// Normalizer 返回归一化器。
func (b *Binder) Normalizer() *Normalizer { return b.norm }

// Install 接管窗体消息。
func (b *Binder) Install() error {
	if b.form == nil {
		return ErrNoForm
	}
	if b.installed {
		return nil
	}
	b.form.SetOnWndProc(b.WndProc)
	b.installed = true
	return nil
}

// Uninstall 释放接管资源。
//
// 刻意不调用 SetOnWndProc(nil)：Uninstall 发生在 Application.Run 返回之后，
// 此时窗口可能已被 Terminate 销毁，回调对象也已失效；再经 FFI 清空事件会
// 触发访问违例（真窗 --shot 退出时实测 0xc0000005）。进程即将退出，LCL
// 会随窗口一起回收回调，注销没有收益只有风险。
func (b *Binder) Uninstall() {
	if b.form == nil || !b.installed {
		return
	}
	if b.captured {
		procReleaseCapture.Call()
		b.captured = false
	}
	b.installed = false
}

// WndProc 是接管后的窗口过程。
//
// 导出是为了让宿主自己做消息循环时也能直接调（例如嵌到别的框架里）。
func (b *Binder) WndProc(msg *types.TLMessage) {
	if msg == nil || b.form == nil {
		return
	}
	m := Message{
		Msg:    uint32(msg.Msg),
		WParam: uintptr(msg.WParam),
		LParam: uintptr(msg.LParam),
		Time:   messageTime(),
	}

	// 私有呈现消息：闭环呈现（画完才允许投下一帧，慢帧后留输入空档）。
	if m.Msg == wmRequestPresent {
		if b.presenter != nil {
			b.presenter.PresentDone()
		}
		msg.Result = 0
		return
	}

	// 修饰键用 GetKeyState 取，而不是靠 wParam 猜：键盘消息的 wParam
	// 不含修饰位，鼠标消息的修饰位又可能与真实状态不同步。
	b.router.SetMods(b.currentMods())

	consumed := b.router.Handle(m)

	// WM_PAINT：Router 只请求一帧并吞掉消息（不让 LCL 画背景）。
	// 真正的像素提交在这里完成 —— 把上一帧缓存的 Backing 贴到窗口 DC。
	// 必须在主线程同步执行：GDI DC 只在消息处理期间稳定。
	if m.Msg == wmPaint && b.presenter != nil {
		b.presenter.PaintOnMessage(clientUpdateRect(b.hwnd))
	}

	// OS 级捕获：按下时抓住，全部抬起时释放。逻辑捕获（Dispatcher）
	// 决定"谁收事件"，OS 捕获决定"事件能不能到窗口"。
	switch m.Msg {
	case wmLButtonDown, wmRButtonDown, wmMButtonDown,
		wmLButtonDblClk, wmRButtonDblClk, wmMButtonDblClk:
		if !b.captured {
			procSetCapture.Call(b.hwnd)
			b.captured = true
		}
	case wmLButtonUp, wmRButtonUp, wmMButtonUp:
		if b.captured && b.norm.Buttons() == 0 {
			procReleaseCapture.Call()
			b.captured = false
		}
	case wmMouseMove:
		// 只需注册一次；注册后系统在指针离开客户区时发 WM_MOUSELEAVE，
		// 之后自动失效（要再注册），所以 tracking 标记在 Leave 时清掉。
		if !b.tracking {
			b.trackLeave()
		}
	case wmMouseLeave:
		b.tracking = false
	}

	if !consumed {
		b.form.InheritedWndProc(msg)
	}
}

// trackLeave 注册离开通知。
func (b *Binder) trackLeave() {
	tme := trackMouseEvent{
		size:  uint32(unsafe.Sizeof(trackMouseEvent{})) - 4, // cbSize 不含自身字段
		flags: tmeLeave,
		hwnd:  b.hwnd,
	}
	// TrackMouseEvent 的 cbSize 是结构体大小（含 dwHoverTime）。
	tme.size = uint32(unsafe.Sizeof(trackMouseEvent{}))
	if r, _, _ := procTrackMouseEvent.Call(uintptr(unsafe.Pointer(&tme))); r != 0 {
		b.tracking = true
	}
}

// currentMods 由 GetKeyState 读取修饰键。
func (b *Binder) currentMods() input.Modifiers {
	sh, _, _ := procGetKeyState.Call(vkShift)
	ct, _, _ := procGetKeyState.Call(vkControl)
	al, _, _ := procGetKeyState.Call(vkMenu)
	return modsFrom(sh&vkStateDown != 0, ct&vkStateDown != 0, al&vkStateDown != 0)
}

// SetCursor 设置窗口光标（只在变化时调用，避免每帧 FFI，§B9）。
func (b *Binder) SetCursor(c scene.Cursor) {
	if b.form == nil || (b.cursorSet && b.lastCur == c) {
		return
	}
	b.lastCur = c
	b.cursorSet = true
	b.form.SetCursor(cursorOf(c))
}

// cursorOf 把场景光标映射到 Windows 光标常量。
func cursorOf(c scene.Cursor) types.TCursor {
	switch c {
	case scene.CursorPointer:
		return 32649 // IDC_HAND
	case scene.CursorText:
		return 32513 // IDC_IBEAM
	case scene.CursorCrosshair:
		return 32515 // IDC_CROSS
	case scene.CursorMove:
		return 32646 // IDC_SIZEALL
	case scene.CursorResizeEW:
		return 32644 // IDC_SIZEWE
	case scene.CursorResizeNS:
		return 32645 // IDC_SIZENS
	case scene.CursorNotAllowed:
		return 32648 // IDC_NO
	}
	return 32512 // IDC_ARROW
}

// screenToClient 把屏幕坐标换成客户区坐标（滚轮消息用）。
func (b *Binder) screenToClient(x, y int32) (int32, int32) {
	if b.hwnd == 0 {
		return x, y
	}
	pt := struct{ X, Y int32 }{x, y}
	procScreenToClient.Call(b.hwnd, uintptr(unsafe.Pointer(&pt)))
	return pt.X, pt.Y
}

// messageTime 返回当前消息的原生时间戳。
//
// 必须在 WndProc 内调用（消息队列里才有效）；取不到时回落到 0，
// Normalizer 会把它当"无时间戳"处理而不是当成 1970 年。
func messageTime() time.Time {
	ms, _, _ := procGetMessageTime.Call()
	if ms == 0 {
		return time.Time{}
	}
	return time.UnixMilli(int64(int32(ms)))
}
