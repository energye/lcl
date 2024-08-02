package win

import (
	"syscall"
	"unsafe"
)

var (
	dwmAPI                        = syscall.NewLazyDLL("dwmapi.dll")
	_DwmExtendFrameIntoClientArea = dwmAPI.NewProc("DwmExtendFrameIntoClientArea")
)

type Margins struct {
	CxLeftWidth, CxRightWidth, CyTopHeight, CyBottomHeight int32
}

func ExtendFrameIntoClientArea(hwnd uintptr, margins Margins) {
	_, _, _ = _DwmExtendFrameIntoClientArea.Call(hwnd, uintptr(unsafe.Pointer(&margins)))
}
