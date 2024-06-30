//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package winapi

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/types"
	"syscall"
)

var (
	user32dll            = syscall.NewLazyDLL("user32.dll")
	_BeginDeferWindowPos = user32dll.NewProc("BeginDeferWindowPos")
	_DeferWindowPos      = user32dll.NewProc("DeferWindowPos")
	_EndDeferWindowPos   = user32dll.NewProc("EndDeferWindowPos")
	_GetDpiForWindow     = user32dll.NewProc("GetDpiForWindow")
)
var (
	gdi32dll  = syscall.NewLazyDLL("gdi32.dll")
	_FrameRgn = gdi32dll.NewProc("FrameRgn")
)

func GetDpiForWindow(hwnd types.HWND) (types.UINT, error) {
	if err := _GetDpiForWindow.Find(); err == nil {
		dpi, _, _ := _GetDpiForWindow.Call(uintptr(hwnd))
		return types.UINT(dpi), nil
	} else {
		return 0, err
	}
}

func FrameRgn(hdc types.HDC, hrgn types.HRGN, hbr types.HBRUSH, w, h int) bool {
	r1, _, _ := _FrameRgn.Call(hdc, hrgn, hbr, uintptr(w), uintptr(h))
	return r1 > 0
}

func BeginDeferWindowPos(nNumWindows int) types.HDWP {
	r1, _, _ := _BeginDeferWindowPos.Call(uintptr(nNumWindows))
	return types.HDWP(r1)
}

func DeferWindowPos(hWinPosInfo types.HDWP, hWnd types.HWND, hWndInsertAfter types.HWND, x, y, cx, cy int, uFlags uint) types.HDWP {
	r1, _, _ := _DeferWindowPos.Call(uintptr(hWinPosInfo), hWnd, hWndInsertAfter, uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	return types.HDWP(r1)
}

func EndDeferWindowPos(hWinPosInfo types.HDWP) bool {
	r1, _, _ := _EndDeferWindowPos.Call(uintptr(hWinPosInfo))
	return r1 > 0
}

func GetWindowLongPtr(hWnd types.HWND, nIndex int32) uintptr {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetWindowLongPtr).Call(uintptr(hWnd), uintptr(nIndex))
	return uintptr(r1)
}

func SetWindowLongPtr(hWnd types.HWND, nIndex int32, dwNewLong uintptr) uintptr {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetWindowLongPtr).Call(uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	return uintptr(r1)
}

func GetClassLongPtr(hWnd types.HWND, nIndex int32) uintptr {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetClassLongPtr).Call(uintptr(hWnd), uintptr(nIndex))
	return uintptr(r1)
}

func SetClassLongPtr(hWnd types.HWND, nIndex int32, dwNewLong uintptr) uintptr {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetClassLongPtr).Call(uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	return uintptr(r1)
}

func FindWindow(lpClassName string, lpWindowName string) types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_FindWindow).Call(api.PascalStr(lpClassName), api.PascalStr(lpWindowName))
	return types.HWND(r1)
}

func FindWindowEx(_para1 types.HWND, _para2 types.HWND, _para3 string, _para4 string) types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_FindWindowEx).Call(uintptr(_para1), uintptr(_para2), api.PascalStr(_para3), api.PascalStr(_para4))
	return types.HWND(r1)
}

func SetWindowText(hWnd types.HWND, lpString string) types.LongBool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetWindowText).Call(uintptr(hWnd), api.PascalStr(lpString))
	return types.LongBool(api.GoBool(r1))
}

func GetWindowText(hWnd types.HWND, lpString string, nMaxCount int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetWindowText).Call(uintptr(hWnd), api.PascalStr(lpString), uintptr(nMaxCount))
	return int32(r1)
}

func GetWindowTextLength(hWnd types.HWND) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetWindowTextLength).Call(uintptr(hWnd))
	return int32(r1)
}
