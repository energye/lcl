//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// TODO : All functions of this API have not been fully tested yet

package winapi

//
//import (
//	"github.com/energye/lcl/api"
//	"github.com/energye/lcl/lcl"
//	"github.com/energye/lcl/types"
//	"unsafe"
//)
//
//type enumDisplayMonitorsProc func(hMonitor types.HMONITOR, hdcMonitor types.HDC, lprcMonitor types.TRect, dwData types.LPARAM) types.LongBool
//type enumFontFamiliesProc func(ELogFont *types.TagEnumLogFontA, Metric *types.TNewTextMetric, FontType types.LongInt, Data types.LPARAM) types.LongInt
//type enumFontFamiliesExProc func(ELogFont *types.TagEnumLogFontExA, Metric *types.TNewTextMetricEx, FontType types.LongInt, Data types.LPARAM) types.LongInt
//
//type EnumDisplayMonitorsCallback struct {
//	instance uintptr
//}
//
//type EnumFontFamiliesCallback struct {
//	instance uintptr
//}
//
//type EnumFontFamiliesExCallback struct {
//	instance uintptr
//}
//
//func NewEnumDisplayMonitorsCallback() *EnumDisplayMonitorsCallback {
//	return &EnumDisplayMonitorsCallback{}
//}
//
//func (m *EnumDisplayMonitorsCallback) Callback(fn enumDisplayMonitorsProc) {
//	if m.instance == 0 {
//		m.instance = api.MakeEventDataPtr(fn)
//	}
//}
//
//// Free 使用完需要将释放掉
//func (m *EnumDisplayMonitorsCallback) Free() {
//	if m.instance != 0 {
//		api.RemoveEventElement(m.instance)
//		api.WinAPI().Proc(cef_win_EnumDisplayMonitorsCallbackFree).Call()
//	}
//}
//
//func NewEnumFontFamiliesCallback() *EnumFontFamiliesCallback {
//	return &EnumFontFamiliesCallback{}
//}
//
//func (m *EnumFontFamiliesCallback) Callback(fn enumFontFamiliesProc) {
//	if m.instance == 0 {
//		m.instance = api.MakeEventDataPtr(fn)
//	}
//}
//
//// Free 使用完需要将释放掉
//func (m *EnumFontFamiliesCallback) Free() {
//	if m.instance != 0 {
//		api.RemoveEventElement(m.instance)
//		api.WinAPI().Proc(cef_win_EnumFontFamiliesCallbackFree).Call()
//	}
//}
//
//func NewEnumFontFamiliesExCallback() *EnumFontFamiliesExCallback {
//	return &EnumFontFamiliesExCallback{}
//}
//
//func (m *EnumFontFamiliesExCallback) Callback(fn enumFontFamiliesExProc) {
//	if m.instance == 0 {
//		m.instance = api.MakeEventDataPtr(fn)
//		api.WinAPI().Proc(cef_win_EnumFontFamiliesExCallbackFree).Call()
//	}
//}
//
//// Free 使用完需要将释放掉
//func (m *EnumFontFamiliesExCallback) Free() {
//	if m.instance != 0 {
//		api.RemoveEventElement(m.instance)
//	}
//}
//
//func init() {
//	var readAnsiCharOffset = func(ptr uintptr, size int) (result []types.AnsiChar) {
//		result = make([]types.AnsiChar, size)
//		for i := 0; i < size; i++ {
//			result[i] = *(*types.AnsiChar)(unsafe.Pointer(ptr + uintptr(i)))
//		}
//		return result
//	}
//	var readDWORDOffset = func(ptr uintptr, size int) (result []types.DWORD) {
//		result = make([]types.DWORD, size)
//		for i := 0; i < size; i++ {
//			result[i] = *(*types.DWORD)(unsafe.Pointer(ptr + uintptr(i)))
//		}
//		return result
//	}
//	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
//		getPtr := func(i int) unsafe.Pointer {
//			return unsafe.Pointer(getVal(i))
//		}
//		switch fn.(type) {
//		case enumDisplayMonitorsProc:
//			var (
//				hMonitor    = types.HMONITOR(getVal(0))
//				hdcMonitor  = types.HDC(getVal(1))
//				lprcMonitor = *(*types.TRect)(getPtr(2))
//				dwData      = types.LPARAM(getVal(3))
//				resultPtr   = (*types.LongBool)(getPtr(4))
//			)
//			*resultPtr = fn.(enumDisplayMonitorsProc)(hMonitor, hdcMonitor, lprcMonitor, dwData)
//		case enumFontFamiliesProc:
//			var (
//				ELogFontPtr = (*types.TagEnumLogFontAPtr)(getPtr(0))
//				ELogFont    *types.TagEnumLogFontA
//				Metric      = (*types.TNewTextMetric)(getPtr(1))
//				FontType    = types.LongInt(getVal(2))
//				Data        = types.LPARAM(getVal(3))
//				resultPtr   = (*types.LongInt)(getPtr(4))
//			)
//			ELogFont = &types.TagEnumLogFontA{
//				ElfLogFont:  (*types.LogFontA)(unsafe.Pointer(ELogFontPtr.ElfLogFont)),
//				ElfFullName: readAnsiCharOffset(ELogFontPtr.ElfFullName, 64),
//				ElfStyle:    readAnsiCharOffset(ELogFontPtr.ElfStyle, 32),
//			}
//			ELogFont.ElfLogFont.LfFaceName = readAnsiCharOffset(ELogFontPtr.LfFaceName, 32)
//			*resultPtr = fn.(enumFontFamiliesProc)(ELogFont, Metric, FontType, Data)
//		case enumFontFamiliesExProc:
//			var (
//				ELogFontPtr          = (*types.TagEnumLogFontExAPtr)(getPtr(0))
//				ELogFont             *types.TagEnumLogFontExA
//				MetricPtr            = (*types.TNewTextMetricExPtr)(getPtr(1))
//				Metric               *types.TNewTextMetricEx
//				NtmeFontSignaturePtr *types.TFontSignaturePtr
//				FontType             = types.LongInt(getVal(2))
//				Data                 = types.LPARAM(getVal(3))
//				resultPtr            = (*types.LongInt)(getPtr(4))
//			)
//			ELogFont = &types.TagEnumLogFontExA{
//				ElfLogFont:  (*types.LogFontA)(unsafe.Pointer(ELogFontPtr.ElfLogFont)),
//				ElfFullName: readAnsiCharOffset(ELogFontPtr.ElfFullName, 64),
//				ElfStyle:    readAnsiCharOffset(ELogFontPtr.ElfStyle, 32),
//				ElfScript:   readAnsiCharOffset(ELogFontPtr.ElfScript, 32),
//			}
//			ELogFont.ElfLogFont.LfFaceName = readAnsiCharOffset(ELogFontPtr.LfFaceName, 32)
//			NtmeFontSignaturePtr = (*types.TFontSignaturePtr)(unsafe.Pointer(MetricPtr.NtmeFontSignature))
//			Metric = &types.TNewTextMetricEx{
//				Ntmentm: *(*types.TNewTextMetric)(unsafe.Pointer(MetricPtr.Ntmentm)),
//				NtmeFontSignature: types.TFontSignature{
//					FsUsb: readDWORDOffset(NtmeFontSignaturePtr.FsUsb, 4),
//					FsCsb: readDWORDOffset(NtmeFontSignaturePtr.FsCsb, 2),
//				},
//			}
//			*resultPtr = fn.(enumFontFamiliesExProc)(ELogFont, Metric, FontType, Data)
//		default:
//			return false
//		}
//		return true
//	})
//}
