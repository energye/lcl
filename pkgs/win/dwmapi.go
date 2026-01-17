//go:build windows
// +build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/energye/lcl/types"
)

var (
	dwmAPI                        = syscall.NewLazyDLL("dwmapi.dll")
	_DwmSetWindowAttribute        = dwmAPI.NewProc("DwmSetWindowAttribute")
	_DwmGetWindowAttribute        = dwmAPI.NewProc("DwmGetWindowAttribute")
	_DwmExtendFrameIntoClientArea = dwmAPI.NewProc("DwmExtendFrameIntoClientArea")
	_DwmEnableBlurBehindWindow    = dwmAPI.NewProc("DwmEnableBlurBehindWindow")
)

type DWMWINDOWATTRIBUTE int32

const DwmwaUseImmersiveDarkModeBefore20h1 DWMWINDOWATTRIBUTE = 19
const DwmwaUseImmersiveDarkMode DWMWINDOWATTRIBUTE = 20
const DwmwaBorderColor DWMWINDOWATTRIBUTE = 34
const DwmwaCaptionColor DWMWINDOWATTRIBUTE = 35
const DwmwaTextColor DWMWINDOWATTRIBUTE = 36
const DwmwaSystemBackdropType DWMWINDOWATTRIBUTE = 38

const DWMSBT_ACRYLIC = 2         // 亚克力
const DWMSBT_MAINWINDOW = 1      // 云母（主窗口）
const DWMSBT_TRANSIENTWINDOW = 3 // 云母（临时窗口）
const DWMWCP_ROUND = 2           // 强制圆角
const DWMWCP_DONOTROUND = 1      // 禁止圆角
const DWMWCP_DEFAULT = 0         // 系统默认圆角

const (
	DWM_BB_ENABLE                = 0x00000001
	DWM_BB_BLURREGION            = 0x00000002
	DWM_BB_TRANSITIONONMAXIMIZED = 0x00000004
)
const DWMWA_SYSTEMBACKDROP_TYPE = 38
const DWMWA_MICA_EFFECT = 1029
const DWMWA_WINDOW_CORNER_PREFERENCE = 33
const DWMWA_USE_IMMERSIVE_DARK_MODE = 20
const DWMWA_ALLOW_NCPAINT = 4

type Margins struct {
	CxLeftWidth, CxRightWidth, CyTopHeight, CyBottomHeight int32
}

type DWMBlurBehind struct {
	DwFlags                uint32         // 配置标识（启用哪些选项）
	FEnable                bool           // 是否启用模糊背景
	HRgnBlur               syscall.Handle // 模糊区域句柄（NULL 或 全屏区域）
	FTransitionOnMaximized bool           // 窗口最大化时是否保留模糊效果
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

func DwmEnableBlurBehindWindow(hwnd types.HWND, blurBehind DWMBlurBehind) error {
	_, _, err := _DwmEnableBlurBehindWindow.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&blurBehind)))
	if err != syscall.Errno(0) {
		return err
	}
	return nil
}
