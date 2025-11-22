//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

//go:build windows
// +build windows

package lcl

import "github.com/energye/lcl/types"

type Window types.HWND

func PlatformHandle(h types.HWND) Window {
	return Window(h)
}

func PlatformWindow(windowInstance uintptr) Window {
	return Window(windowInstance)
}
