//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build windows
// +build windows

package win

import (
	"github.com/energye/lcl/types"
)

var (
	_GetWindowLongPtr = user32dll.NewProc("GetWindowLongW")
	_SetWindowLongPtr = user32dll.NewProc("SetWindowLongW")
	_SetClassLongPtr  = user32dll.NewProc("SetClassLongW")
	_GetClassLongPtr  = user32dll.NewProc("GetClassLongW")
)

func WindowFromPoint(point types.TPoint) types.HWND {
	r, _, _ := _WindowFromPoint.Call(uintptr(point.X), uintptr(point.Y))
	return types.HWND(r)
}

func WindowFromPhysicalPoint(point types.TPoint) types.HWND {
	r, _, _ := _WindowFromPhysicalPoint.Call(uintptr(point.Y), uintptr(point.Y))
	return types.HWND(r)
}

func ChildWindowFromPoint(hWndParent types.HWND, point types.TPoint) types.HWND {
	r, _, _ := _ChildWindowFromPoint.Call(uintptr(hWndParent), uintptr(point.Y), uintptr(point.Y))
	return types.HWND(r)
}
