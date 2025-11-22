//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

//go:build linux
// +build linux

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/types"
)

type GdkWindow uintptr

type Gtk3Window uintptr

type GtkWidget uintptr

type GtkFixed uintptr

type XID uintptr

// PlatformHandle gtk hand
func PlatformHandle(h types.HWND) GtkWidget {
	return GtkWidget(h)
}

// PlatformWindow gtk2
func PlatformWindow(windowInstance uintptr) GdkWindow {
	if api.Widget().IsGTK2() {
		return GdkWindow(api.GdkWindow_FromForm(windowInstance))
	}
	println("[WARNING] PlatformWindow Current ENV Widget Not GTK2")
	return 0
}

func (m GdkWindow) XID() (xid XID) {
	return XID(api.GdkWindow_GetXId(uintptr(m)))
}

// FixedWidget
//
// gtk2 中首先是一个widget，然后上面用了一个fixedWidget来处理的。
func (m GtkWidget) FixedWidget() GtkFixed {
	if api.Widget().IsGTK2() {
		return GtkFixed(api.GtkWidget_GetGtkFixed(uintptr(m)))
	}
	println("[WARNING] FixedWidget Current ENV Widget Not GTK2")
	return 0
}

// Window gtk2
func (m GtkWidget) Window() GdkWindow {
	if api.Widget().IsGTK2() {
		return GdkWindow(api.GtkWidget_Window(uintptr(m)))
	}
	println("[WARNING] Window Current ENV Widget Not GTK2")
	return 0
}

// Gtk3Window 获取 gtk3 窗口
func (m GtkWidget) Gtk3Window() Gtk3Window {
	if api.Widget().IsGTK3() {
		return Gtk3Window(api.Gtk3Widget_GtkWindow(uintptr(m)))
	}
	println("[WARNING] Gtk3Window Current ENV Widget Not GTK3")
	return 0
}

// Gtk3Widget 获取 gtk3 控件
func (m GtkWidget) Gtk3Widget() GtkWidget {
	if api.Widget().IsGTK3() {
		return GtkWidget(api.Gtk3Widget_Widget(uintptr(m)))
	}
	println("[WARNING] Gtk3Widget Current ENV Widget Not GTK3")
	return 0
}

// DragWindow
//
//	button 1 鼠标左键
func DragWindow(handle types.HWND, windowLeft, windowTop, button int32, edge api.TGdkWindowEdge) {
	api.GtkWindow_Drag(handle, windowLeft, windowTop, button, edge)
}
