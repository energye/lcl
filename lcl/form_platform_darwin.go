//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

//go:build darwin
// +build darwin

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/types"
)

/*

   #cgo CFLAGS: -x objective-c
   #cgo LDFLAGS: -framework Cocoa

   #include <Cocoa/Cocoa.h>

  static NSWindow* toNSWindow(void* ptr) {
      return ((__bridge NSWindow*)ptr); //
  }

  // test
  static void NSWindow_setTitleVisibility(void *ptr) {
     NSWindow *win = toNSWindow(ptr);
     win.TitleVisibility = NSWindowTitleHidden;
  }

*/
//import "C"

type (
	NSObject uintptr

	NSWindow uintptr

	NSURL uintptr
)

// NSObject
func PlatformHandle(h types.HWND) NSObject {
	return NSObject(h)
}

func PlatformWindow(windowInstance uintptr) NSWindow {
	return NSWindow(api.NSWindow_FromForm(windowInstance))
}

func (n NSWindow) TitleVisibility() types.NSWindowTitleVisibility {
	return types.NSWindowTitleVisibility(api.NSWindow_titleVisibility(uintptr(n)))
}

func (n NSWindow) SetTitleVisibility(flag types.NSWindowTitleVisibility) {
	//C.NSWindow_setTitleVisibility(unsafe.Pointer(n))
	api.NSWindow_setTitleVisibility(uintptr(n), uintptr(flag))
}

func (n NSWindow) TitleBarAppearsTransparent() bool {
	return api.NSWindow_titlebarAppearsTransparent(uintptr(n))
}

func (n NSWindow) SetTitleBarAppearsTransparent(flag bool) {
	api.NSWindow_setTitlebarAppearsTransparent(uintptr(n), flag)
}

func (n NSWindow) SetRepresentedURL(url NSURL) {
	api.NSWindow_setRepresentedURL(uintptr(n), uintptr(url))
}

func (n NSWindow) StyleMask() uint {
	return api.NSWindow_styleMask(uintptr(n))
}

func (n NSWindow) SetStyleMask(mask uint) {
	api.NSWindow_setStyleMask(uintptr(n), mask)
}
