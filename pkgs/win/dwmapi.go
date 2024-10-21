//go:build windows
// +build windows

package win

import (
	"github.com/energye/lcl/types"
	"syscall"
	"unsafe"
)

var (
	dwmAPI                        = syscall.NewLazyDLL("dwmapi.dll")
	_DwmSetWindowAttribute        = dwmAPI.NewProc("DwmSetWindowAttribute")
	_DwmGetWindowAttribute        = dwmAPI.NewProc("DwmGetWindowAttribute")
	_DwmExtendFrameIntoClientArea = dwmAPI.NewProc("DwmExtendFrameIntoClientArea")
)

type DWMWINDOWATTRIBUTE int32

const DwmwaUseImmersiveDarkModeBefore20h1 DWMWINDOWATTRIBUTE = 19
const DwmwaUseImmersiveDarkMode DWMWINDOWATTRIBUTE = 20
const DwmwaBorderColor DWMWINDOWATTRIBUTE = 34
const DwmwaCaptionColor DWMWINDOWATTRIBUTE = 35
const DwmwaTextColor DWMWINDOWATTRIBUTE = 36
const DwmwaSystemBackdropType DWMWINDOWATTRIBUTE = 38

type Margins struct {
	CxLeftWidth, CxRightWidth, CyTopHeight, CyBottomHeight int32
}

func ExtendFrameIntoClientArea(hwnd types.HWND, margins Margins) {
	_, _, _ = _DwmExtendFrameIntoClientArea.Call(hwnd, uintptr(unsafe.Pointer(&margins)))
}

func DwmSetWindowAttribute(hwnd types.HWND, dwAttribute DWMWINDOWATTRIBUTE, pvAttribute unsafe.Pointer, cbAttribute uintptr) types.HRESULT {
	r1, _, _ := _DwmSetWindowAttribute.Call(hwnd, uintptr(dwAttribute), uintptr(pvAttribute), cbAttribute)
	return types.HRESULT(r1)
}

func DwmGetWindowAttribute(hwnd types.HWND, dwAttribute DWMWINDOWATTRIBUTE, pvAttribute unsafe.Pointer, cbAttribute uintptr) types.HRESULT {
	r1, _, _ := _DwmGetWindowAttribute.Call(hwnd, uintptr(dwAttribute), uintptr(pvAttribute), cbAttribute)
	return types.HRESULT(r1)
}
