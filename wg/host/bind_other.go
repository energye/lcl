//go:build !windows

package host

// Binder 在非 Windows 平台只提供与 Windows 同名的空实现。
//
// 为什么要有它：Container 在 Run/RunManual 里统一安装 Binder，若平台
// 差异散到容器代码里就会出现 build tag 的分支污染。真正的平台消息
// 循环本来就不存在于 headless/CI，stub 让同一份容器代码可跨平台编译。
type Binder struct {
	ime *IMEBridge
}

// NewBinder 创建空绑定器。
func NewBinder(*Host, EventSink) *Binder { return &Binder{ime: NewIMEBridge(nil)} }

// Router 返回 nil：无窗口平台没有消息路由。
func (*Binder) Router() *Router { return nil }

// Normalizer 返回 nil。
func (*Binder) Normalizer() *Normalizer { return nil }

// IME 返回桥（无平台实现，但编辑模型仍可测）。
func (b *Binder) IME() *IMEBridge { return b.ime }

// SetIMEBridge 替换桥。
func (b *Binder) SetIMEBridge(br *IMEBridge) {
	if br != nil {
		b.ime = br
	}
}

// Install 在无窗口平台无可安装对象。
func (*Binder) Install() error { return ErrNoForm }

// Uninstall 空操作。
func (*Binder) Uninstall() {}
