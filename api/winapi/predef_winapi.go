//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package winapi LCL winAPI and RTL commonly used function APIs
package winapi

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/types"
	"unsafe"
)

func LOBYTE(w uint16) byte {
	return byte(w)
}

func HIBYTE(w uint16) byte {
	return byte(w >> 8 & 0xff)
}

func LOWORD(dw uint32) uint16 {
	return uint16(dw & 0xFFFF)
}

func HIWORD(dw uint32) uint16 {
	return uint16(dw >> 16 & 0xffff)
}

func GET_X_LPARAM(lp uintptr) int32 {
	return int32(int16(LOWORD(uint32(lp))))
}

func GET_Y_LPARAM(lp uintptr) int32 {
	return int32(int16(HIWORD(uint32(lp))))
}

type HCursor struct {
	instance unsafe.Pointer
}

func (m *HCursor) Free() {
	m.instance = nil
}

func CreateRectRgn(X1, Y1, X2, Y2 int32) types.HRGN {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateRectRgn).Call(uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return types.HRGN(r1)
}

func SetRectRgn(aRGN types.HRGN, X1, Y1, X2, Y2 int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetRectRgn).Call(aRGN, uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return api.GoBool(r1)
}

func DeleteObject(aRGN types.HRGN) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_DeleteObject).Call(aRGN)
	return api.GoBool(r1)
}

func CombineRgn(dest, src1, src2 types.HRGN, fnCombineMode types.RNGFnCombineMode) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CombineRgn).Call(dest, src1, src2, uintptr(fnCombineMode))
	return int32(r1)
}

func PtInRegion(RGN types.HRGN, X, Y int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_PtInRegion).Call(RGN, uintptr(X), uintptr(Y))
	return api.GoBool(r1)
}

func ScreenToClient(handle types.HWND, p *types.TPoint) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ScreenToClient).Call(uintptr(handle), uintptr(unsafe.Pointer(p)))
	return int32(r1)
}

func ClientToScreen(handle types.HWND, p *types.TPoint) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ClientToScreen).Call(uintptr(handle), uintptr(unsafe.Pointer(p)))
	return api.GoBool(r1)
}

func DefWindowProc(handle types.HWND, msg uint32, wParam types.WPARAM, lParam types.LPARAM) types.LResult {
	r1, _, _ := winapiImportAPI().Proc(winAPI_DefWindowProc).Call(uintptr(handle), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return types.LResult(r1)
}

func DefSubclassProc(handle types.HWND, msg uint32, wParam types.WPARAM, lParam types.LPARAM) types.LResult {
	r1, _, _ := winapiImportAPI().Proc(winAPI_DefSubclassProc).Call(uintptr(handle), uintptr(msg), uintptr(wParam), uintptr(lParam))
	return types.LResult(r1)
}

func CreateRoundRectRgn(_para1, _para2, _para3, _para4, _para5, _para6 int32) types.HRGN {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateRoundRectRgn).Call(uintptr(_para1), uintptr(_para2), uintptr(_para3), uintptr(_para4), uintptr(_para5), uintptr(_para6))
	return types.HRGN(r1)
}

func SetWindowRgn(handle types.HWND, hRgn types.HRGN, bRedraw bool) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetWindowRgn).Call(uintptr(handle), hRgn, api.PascalBool(bRedraw))
	return int32(r1)
}

func SetCursor(hCursor *HCursor) *HCursor {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetCursor).Call(uintptr(hCursor.instance))
	return &HCursor{
		instance: unsafe.Pointer(r1),
	}
}

func LoadCursor(handle types.HWND, lpCursorName int32) *HCursor {
	r1, _, _ := winapiImportAPI().Proc(winAPI_LoadCursor).Call(uintptr(handle), uintptr(lpCursorName))
	return &HCursor{
		instance: unsafe.Pointer(r1),
	}
}

func OnPaint(handle types.HWND) {
	winapiImportAPI().Proc(winAPI_OnPaint).Call(uintptr(handle))
}

//func SetDraggableRegions(aRGN types.HRGN, regions []cef.TCefDraggableRegion) {
//	//SetDraggableRegions 代码实现
//	draggableRegion := WinCreateRectRgn(0, 0, 0, 0)
//	WinSetRectRgn(draggableRegion, 0, 0, 0, 0)
//	for i := 0; i < regions.RegionsCount(); i++ {
//		region := regions.Region(i)
//		creRGN := WinCreateRectRgn(region.Bounds.X, region.Bounds.Y, region.Bounds.X+region.Bounds.Width, region.Bounds.Y+region.Bounds.Height)
//		if region.Draggable {
//			WinCombineRgn(draggableRegion, draggableRegion, creRGN, consts.RGN_OR)
//		} else {
//			WinCombineRgn(draggableRegion, draggableRegion, creRGN, consts.RGN_DIFF)
//		}
//		WinDeleteObject(creRGN)
//	}
//	fmt.Println("Check PtInRegion：", WinPtInRegion(draggableRegion, 50, 50))
//	winapiImportAPI().Proc(winAPI_SetDraggableRegions).Call(aRGN, uintptr(int32(len(regions))), uintptr(unsafe.Pointer(&regions[0])), uintptr(int32(len(regions))))
//}

func EndPaint(Handle types.HWND, PS types.TagPaintStruct) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_EndPaint).Call(uintptr(Handle), uintptr(unsafe.Pointer(&PS)))
	return int32(r1)
}

func BeginPaint(Handle types.HWND, PS types.TagPaintStruct) types.HDC {
	r1, _, _ := winapiImportAPI().Proc(winAPI_BeginPaint).Call(uintptr(Handle), uintptr(unsafe.Pointer(&PS)))
	return types.HDC(r1)
}

func Arc(DC types.HDC, Left, Top, Right, Bottom, Angle16Deg, Angle16DegLength int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_Arc).Call(uintptr(DC), uintptr(Left), uintptr(Top), uintptr(Right), uintptr(Bottom), uintptr(Angle16Deg), uintptr(Angle16DegLength))
	return bool(api.GoBool(r1))
}

func AngleChord(DC types.HDC, x1, y1, x2, y2, angle1, angle2 int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_AngleChord).Call(uintptr(DC), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(angle1), uintptr(angle2))
	return bool(api.GoBool(r1))
}

func CallNextHookEx(hhk types.HOOK, ncode int32, WParam types.WPARAM, LParam types.LPARAM) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CallNextHookEx).Call(uintptr(hhk), uintptr(ncode), uintptr(WParam), uintptr(LParam))
	return int32(r1)
}

func CallWindowProc(lpPrevWndFunc types.TFarProc, Handle types.HWND, Msg uint32, WParam types.WPARAM, LParam types.LPARAM) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CallWindowProc).Call(uintptr(lpPrevWndFunc), uintptr(Handle), uintptr(Msg), uintptr(WParam), uintptr(LParam))
	return int32(r1)
}

func BitBlt(DestDC types.HDC, X, Y, Width, Height int32, SrcDC types.HDC, XSrc, YSrc int32, Rop types.DWORD) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_BitBlt).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(Rop))
	return bool(api.GoBool(r1))
}

func CreateBitmap(Width, Height int32, Planes, BitCount int32, BitmapBits types.Pointer) types.HBITMAP {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateBitmap).Call(uintptr(Width), uintptr(Height), uintptr(Planes), uintptr(BitCount), uintptr(BitmapBits))
	return types.HBITMAP(r1)
}

func CreateBrushIndirect(LogBrush types.TagLogBrush) types.HBRUSH {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateBrushIndirect).Call(uintptr(unsafe.Pointer(&LogBrush)))
	return types.HBITMAP(r1)
}

func CreateBrushWithRadialGradient(LogBrush types.TLogRadialGradient) types.HBRUSH {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateBrushWithRadialGradient).Call(uintptr(unsafe.Pointer(&LogBrush)))
	return types.HBITMAP(r1)
}

func CreateCaret(Handle types.HWND, Bitmap types.HBITMAP, width, Height int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateCaret).Call(uintptr(Handle), uintptr(Bitmap), uintptr(width), uintptr(Height))
	return bool(api.GoBool(r1))
}

func CreateCompatibleBitmap(DC types.HDC, Width, Height int32) types.HBITMAP {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateCompatibleBitmap).Call(uintptr(DC), uintptr(Width), uintptr(Height))
	return types.HBITMAP(r1)
}

func CreateCompatibleDC(DC types.HDC) types.HDC {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateCompatibleDC).Call(uintptr(DC))
	return types.HDC(r1)
}

func CreateDIBitmap(DC types.HDC, InfoHeader types.TagBitmapInfoHeader, dwUsage types.DWORD, InitBits types.PChar, InitInfo types.TagBitmapInfo, wUsage uint32) types.HBITMAP {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateDIBitmap).Call(uintptr(DC), uintptr(unsafe.Pointer(&InfoHeader)), uintptr(dwUsage), api.PascalStr(InitBits), uintptr(unsafe.Pointer(&InitInfo)), uintptr(wUsage))
	return types.HBITMAP(r1)
}

func CreateDIBSection(DC types.HDC, BitmapInfo types.TagBitmapInfo, Usage uint32, Bits types.Pointer, SectionHandle types.THandle, Offset types.DWORD) types.HBITMAP {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateDIBSection).Call(uintptr(DC), uintptr(unsafe.Pointer(&BitmapInfo)), uintptr(Usage), uintptr(Bits), uintptr(SectionHandle), uintptr(Offset))
	return types.HBITMAP(r1)
}

func CreateEllipticRgn(X1, Y1, X2, Y2 int32) types.HRGN {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateEllipticRgn).Call(uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return types.HRGN(r1)
}

func CreateFontIndirect(LogFont types.LogFontA) types.HFONT {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateFontIndirect).Call(uintptr(unsafe.Pointer(&LogFont)))
	return types.HFONT(r1)
}

func CreateFontIndirectEx(LogFont types.LogFontA, LongFontName types.PChar) types.HFONT {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateFontIndirectEx).Call(uintptr(unsafe.Pointer(&LogFont)), api.PascalStr(LongFontName))
	return types.HFONT(r1)
}

func CreateIconIndirect(IconInfo types.ICONINFO) types.HICON {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreateIconIndirect).Call(uintptr(unsafe.Pointer(&IconInfo)))
	return types.HFONT(r1)
}

func CreatePalette(LogPalette types.TagLogPalette) types.HPALETTE {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreatePalette).Call(uintptr(unsafe.Pointer(&LogPalette)))
	return types.HPALETTE(r1)
}

func CreatePatternBrush(ABitmap types.HBITMAP) types.HBRUSH {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreatePatternBrush).Call(uintptr(ABitmap))
	return types.HBRUSH(r1)
}

func CreatePenIndirect(LogPen types.TagLogPen) types.HPEN {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreatePenIndirect).Call(uintptr(unsafe.Pointer(&LogPen)))
	return types.HPEN(r1)
}

func CreatePolygonRgn(points []types.TPoint, NumPts int32, FillMode int32) types.HRGN {
	r1, _, _ := winapiImportAPI().Proc(winAPI_CreatePolygonRgn).Call(uintptr(unsafe.Pointer(&points[0])), uintptr(NumPts), uintptr(FillMode))
	return types.HRGN(r1)
}

func DeleteCriticalSection(CritSection types.TCriticalSection) {
	winapiImportAPI().Proc(winAPI_DeleteCriticalSection).Call(uintptr(CritSection))
}

func DeleteDC(hDC types.HDC) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_DeleteDC).Call(uintptr(hDC))
	return bool(api.GoBool(r1))
}

func DestroyCaret(Handle types.HWND) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_DestroyCaret).Call(uintptr(Handle))
	return bool(api.GoBool(r1))
}

func DestroyCursor(Handle types.HCURSOR) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_DestroyCursor).Call(uintptr(Handle))
	return bool(api.GoBool(r1))
}

func DestroyIcon(Handle types.HICON) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_DestroyIcon).Call(uintptr(Handle))
	return bool(api.GoBool(r1))
}

func DrawFrameControl(DC types.HDC, Rect types.TRect, uType, uState types.Cardinal) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_DrawFrameControl).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(uType), uintptr(uState))
	return bool(api.GoBool(r1))
}

func DrawFocusRect(DC types.HDC, Rect types.TRect) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_DrawFocusRect).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)))
	return bool(api.GoBool(r1))
}

func DrawEdge(DC types.HDC, Rect types.TRect, edge types.Cardinal, grfFlags types.Cardinal) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_DrawEdge).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(edge), uintptr(grfFlags))
	return bool(api.GoBool(r1))
}

func DrawText(DC types.HDC, Str types.PChar, Count int32, Rect types.TRect, Flags types.Cardinal) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_DrawText).Call(uintptr(DC), api.PascalStr(Str), uintptr(Count), uintptr(unsafe.Pointer(&Rect)), uintptr(Flags))
	return int32(r1)
}

func EnableScrollBar(Wnd types.HWND, wSBflags, wArrows types.Cardinal) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_EnableScrollBar).Call(uintptr(Wnd), uintptr(wSBflags), uintptr(wArrows))
	return bool(api.GoBool(r1))
}

func EnableWindow(hWnd types.HWND, bEnable bool) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_EnableWindow).Call(uintptr(hWnd), api.PascalBool(bool(bEnable)))
	return bool(api.GoBool(r1))
}

func EnterCriticalSection(CritSection types.TCriticalSection) {
	winapiImportAPI().Proc(winAPI_EnterCriticalSection).Call(CritSection)
}

func Ellipse(DC types.HDC, x1, y1, x2, y2 int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_Ellipse).Call(uintptr(DC), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	return bool(api.GoBool(r1))
}

func EqualRgn(Rgn1 types.HRGN, Rgn2 types.HRGN) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_EqualRgn).Call(Rgn1, Rgn2)
	return bool(api.GoBool(r1))
}

func ExcludeClipRect(dc types.HDC, Left, Top, Right, Bottom int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ExcludeClipRect).Call(uintptr(dc), uintptr(Left), uintptr(Top), uintptr(Right), uintptr(Bottom))
	return int32(r1)
}

func ExtCreatePen(dwPenStyle, dwWidth types.DWORD, lplb types.TagLogBrush, dwStyleCount types.DWORD, lpStyle types.DWORD) types.HPEN {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ExtCreatePen).Call(uintptr(dwPenStyle), uintptr(dwWidth), uintptr(unsafe.Pointer(&lplb)), uintptr(dwStyleCount), uintptr(lpStyle))
	return types.HPEN(r1)
}

func ExtTextOut(DC types.HDC, X, Y int32, Options int32, Rect types.TRect, Str types.PChar, Count int32, Dx int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ExtTextOut).Call(uintptr(DC), uintptr(X), uintptr(Y), uintptr(Options), uintptr(unsafe.Pointer(&Rect)), api.PascalStr(Str), uintptr(Count), uintptr(Dx))
	return bool(api.GoBool(r1))
}

func ExtSelectClipRGN(dc types.HDC, rgn types.HRGN, Mode int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ExtSelectClipRGN).Call(uintptr(dc), rgn, uintptr(Mode))
	return int32(r1)
}

func FillRect(DC types.HDC, Rect types.TRect, Brush types.HBRUSH) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_FillRect).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(Brush))
	return bool(api.GoBool(r1))
}

func FillRgn(DC types.HDC, RegionHnd types.HRGN, hbr types.HBRUSH) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_FillRgn).Call(uintptr(DC), RegionHnd, uintptr(hbr))
	return bool(api.GoBool(r1))
}

func FloodFill(DC types.HDC, X, Y int32, Color types.TGraphicsColor, FillStyle types.TGraphicsFillStyle, Brush types.HBRUSH) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_FloodFill).Call(uintptr(DC), uintptr(X), uintptr(Y), uintptr(Color), uintptr(FillStyle), uintptr(Brush))
	return bool(api.GoBool(r1))
}

func FrameRect(DC types.HDC, Rect types.TRect, hBr types.HBRUSH) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_FrameRect).Call(uintptr(DC), uintptr(unsafe.Pointer(&Rect)), uintptr(hBr))
	return int32(r1)
}

func GetActiveWindow() types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetActiveWindow).Call()
	return types.HWND(r1)
}

func GetBitmapBits(Bitmap types.HBITMAP, Count int32, Bits types.Pointer) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetBitmapBits).Call(uintptr(Bitmap), uintptr(Count), uintptr(Bits))
	return int32(r1)
}

func GetBkColor(DC types.HDC) types.TColorRef {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetBkColor).Call(uintptr(DC))
	return types.TColorRef(r1)
}

func GetCapture() types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetCapture).Call()
	return types.HWND(r1)
}

func GetCaretPos(lpPoint types.TPoint) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetCaretPos).Call(uintptr(unsafe.Pointer(&lpPoint)))
	return bool(api.GoBool(r1))
}

func GetClientRect(handle types.HWND, Rect types.TRect) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetClientRect).Call(uintptr(handle), uintptr(unsafe.Pointer(&Rect)))
	return bool(api.GoBool(r1))
}

func GetClipBox(DC types.HDC, lpRect types.TRect) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetClipBox).Call(uintptr(DC), uintptr(unsafe.Pointer(&lpRect)))
	return int32(r1)
}

func GetClipRGN(DC types.HDC, RGN types.HRGN) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetClipRGN).Call(uintptr(DC), RGN)
	return int32(r1)
}

func GetCurrentObject(DC types.HDC, uObjectType uint32) types.HGDIOBJ {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetCurrentObject).Call(uintptr(DC), uintptr(uObjectType))
	return types.HGDIOBJ(r1)
}

func GetCursorPos(lpPoint types.TPoint) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetCursorPos).Call(uintptr(unsafe.Pointer(&lpPoint)))
	return bool(api.GoBool(r1))
}

func GetDC(hWnd types.HWND) types.HDC {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetDC).Call(uintptr(hWnd))
	return types.HDC(r1)
}

func GetDeviceCaps(DC types.HDC, Index int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetDeviceCaps).Call(uintptr(DC), uintptr(Index))
	return int32(r1)
}

func GetDIBits(DC types.HDC, Bitmap types.HBITMAP, StartScan, NumScans uint32, Bits types.Pointer, BitInfo types.TagBitmapInfo, Usage uint32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetDIBits).Call(uintptr(DC), uintptr(Bitmap), uintptr(StartScan), uintptr(NumScans), uintptr(Bits), uintptr(unsafe.Pointer(&BitInfo)), uintptr(Usage))
	return int32(r1)
}

func GetDoubleClickTime() uint32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetDoubleClickTime).Call()
	return uint32(r1)
}

func GetFocus() types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetFocus).Call()
	return types.HWND(r1)
}

func GetFontLanguageInfo(DC types.HDC) types.DWORD {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetFontLanguageInfo).Call(uintptr(DC))
	return types.DWORD(r1)
}

func GetForegroundWindow() types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetForegroundWindow).Call()
	return types.HWND(r1)
}

func GetIconInfo(AIcon types.HICON, AIconInfo types.ICONINFO) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetIconInfo).Call(uintptr(AIcon), uintptr(unsafe.Pointer(&AIconInfo)))
	return bool(api.GoBool(r1))
}

func GetKeyState(nVirtKey int32) int16 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetKeyState).Call(uintptr(nVirtKey))
	return int16(r1)
}

func GetMapMode(DC types.HDC) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetMapMode).Call(uintptr(DC))
	return int32(r1)
}

func GetMonitorInfo(hMonitor types.HMONITOR, lpmi types.TagMonitorInfo) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetMonitorInfo).Call(uintptr(hMonitor), uintptr(unsafe.Pointer(&lpmi)))
	return bool(api.GoBool(r1))
}

func GetDpiForMonitor(hmonitor types.HMONITOR, dpiType MONITOR_DPI_TYPE, dpiX *uint32, dpiY *uint32) types.HRESULT { // out
	var (
		outDpiX uint32
		outDpiY uint32
	)
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetDpiForMonitor).Call(uintptr(hmonitor), uintptr(dpiType), uintptr(unsafe.Pointer(&outDpiX)), uintptr(unsafe.Pointer(&outDpiY)))
	*dpiX = uint32(outDpiX)
	*dpiY = uint32(outDpiY)
	return types.HRESULT(r1)
}

func GetObject(GDIObject types.HGDIOBJ, BufSize int32, Buf types.Pointer) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetObject).Call(uintptr(GDIObject), uintptr(BufSize), uintptr(Buf))
	return int32(r1)
}

func GetParent(Handle types.HWND) types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetParent).Call(uintptr(Handle))
	return types.HWND(r1)
}

func GetProp(Handle types.HWND, Str types.PChar) types.Pointer {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetProp).Call(uintptr(Handle), api.PascalStr(Str))
	return types.HWND(r1)
}

func GetRgnBox(RGN types.HRGN, lpRect types.TRect) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetRgnBox).Call(RGN, uintptr(unsafe.Pointer(&lpRect)))
	return int32(r1)
}

func GetROP2(DC types.HDC) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetROP2).Call(uintptr(DC))
	return int32(r1)
}

func GetScrollInfo(Handle types.HWND, SBStyle int32, ScrollInfo types.TagScrollInfo) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetScrollInfo).Call(uintptr(Handle), uintptr(SBStyle), uintptr(unsafe.Pointer(&ScrollInfo)))
	return bool(api.GoBool(r1))
}

func GetStockObject(Value int32) types.THandle {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetStockObject).Call(uintptr(Value))
	return types.THandle(r1)
}

func GetSysColor(nIndex int32) types.DWORD {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetSysColor).Call(uintptr(nIndex))
	return types.DWORD(r1)
}

func GetSysColorBrush(nIndex int32) types.HBRUSH {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetSysColorBrush).Call(uintptr(nIndex))
	return types.HBRUSH(r1)
}

func GetSystemMetrics(nIndex int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetSystemMetrics).Call(uintptr(nIndex))
	return int32(r1)
}

func GetTextColor(DC types.HDC) types.TColorRef {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetTextColor).Call(uintptr(DC))
	return types.TColorRef(r1)
}

func GetTextExtentExPoint(DC types.HDC, Str types.PChar, Count, MaxWidth int32, MaxCount, PartialWidths int32, Size types.TSize) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetTextExtentExPoint).Call(uintptr(DC), api.PascalStr(Str), uintptr(Count), uintptr(MaxWidth), uintptr(MaxCount), uintptr(PartialWidths), uintptr(unsafe.Pointer(&Size)))
	return bool(api.GoBool(r1))
}

func GetTextExtentPoint(DC types.HDC, Str types.PChar, Count int32, Size types.TSize) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetTextExtentPoint).Call(uintptr(DC), api.PascalStr(Str), uintptr(Count), uintptr(unsafe.Pointer(&Size)))
	return bool(api.GoBool(r1))
}

func GetTextExtentPoint32(DC types.HDC, Str types.PChar, Count int32, Size types.TSize) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetTextExtentPoint32).Call(uintptr(DC), api.PascalStr(Str), uintptr(Count), uintptr(unsafe.Pointer(&Size)))
	return bool(api.GoBool(r1))
}

func GetTextMetrics(DC types.HDC, TM types.TagTextMetricA) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetTextMetrics).Call(uintptr(DC), uintptr(unsafe.Pointer(&TM)))
	return bool(api.GoBool(r1))
}

func GetViewPortExtEx(DC types.HDC, Size types.TSize) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetViewPortExtEx).Call(uintptr(DC), uintptr(unsafe.Pointer(&Size)))
	return int32(r1)
}

func GetViewPortOrgEx(DC types.HDC, P types.TPoint) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetViewPortOrgEx).Call(uintptr(DC), uintptr(unsafe.Pointer(&P)))
	return int32(r1)
}

func GetWindowExtEx(DC types.HDC, Size types.TSize) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetWindowExtEx).Call(uintptr(DC), uintptr(unsafe.Pointer(&Size)))
	return int32(r1)
}

func GetWindowLong(Handle types.HWND, int int32) types.PtrInt {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetWindowLong).Call(uintptr(Handle), uintptr(int))
	return types.PtrInt(r1)
}

func GetWindowRect(Handle types.HWND, Rect types.TRect) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetWindowRect).Call(uintptr(Handle), uintptr(unsafe.Pointer(&Rect)))
	return int32(r1)
}

func GetWindowSize(Handle types.HWND, Width, Height int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetWindowSize).Call(uintptr(Handle), uintptr(Width), uintptr(Height))
	return bool(api.GoBool(r1))
}

func GetWindowOrgEx(dc types.HDC, P types.TPoint) int32 { // because of delphi compatibility
	r1, _, _ := winapiImportAPI().Proc(winAPI_GetWindowOrgEx).Call(uintptr(dc), uintptr(unsafe.Pointer(&P)))
	return int32(r1)
}

func GradientFill(DC types.HDC, Vertices types.TagTriVertex, NumVertices int32, Meshes types.Pointer, NumMeshes int32, Mode int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_GradientFill).Call(uintptr(DC), uintptr(unsafe.Pointer(&Vertices)), uintptr(NumVertices), uintptr(Meshes), uintptr(NumMeshes), uintptr(Mode))
	return bool(api.GoBool(r1))
}

func HideCaret(hWnd types.HWND) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_HideCaret).Call(uintptr(hWnd))
	return bool(api.GoBool(r1))
}

func InitializeCriticalSection(CritSection types.TCriticalSection) {
	winapiImportAPI().Proc(winAPI_InitializeCriticalSection).Call(uintptr(CritSection))
}

func IntersectClipRect(dc types.HDC, Left, Top, Right, Bottom int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_IntersectClipRect).Call(uintptr(dc), uintptr(Left), uintptr(Top), uintptr(Right), uintptr(Bottom))
	return int32(r1)
}

func InvalidateRect(aHandle types.HWND, ARect types.TRect, bErase bool) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_InvalidateRect).Call(uintptr(aHandle), uintptr(unsafe.Pointer(&ARect)), api.PascalBool(bool(bErase)))
	return bool(api.GoBool(r1))
}

func InvalidateRgn(Handle types.HWND, Rgn types.HRGN, Erase bool) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_InvalidateRgn).Call(uintptr(Handle), Rgn, api.PascalBool(bool(Erase)))
	return bool(api.GoBool(r1))
}

func IsDBCSLeadByte(TestChar byte) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_IsDBCSLeadByte).Call(uintptr(TestChar))
	return bool(api.GoBool(r1))
}

func IsIconic(handle types.HWND) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_IsIconic).Call(uintptr(handle))
	return bool(api.GoBool(r1))
}

func IsWindow(handle types.HWND) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_IsWindow).Call(uintptr(handle))
	return bool(api.GoBool(r1))
}

func IsWindowEnabled(handle types.HWND) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_IsWindowEnabled).Call(uintptr(handle))
	return bool(api.GoBool(r1))
}

func IsWindowVisible(handle types.HWND) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_IsWindowVisible).Call(uintptr(handle))
	return bool(api.GoBool(r1))
}

func IsZoomed(handle types.HWND) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_IsZoomed).Call(uintptr(handle))
	return bool(api.GoBool(r1))
}

func LeaveCriticalSection(CritSection types.TCriticalSection) {
	winapiImportAPI().Proc(winAPI_LeaveCriticalSection).Call(uintptr(CritSection))
}

func LineTo(DC types.HDC, X, Y int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_LineTo).Call(uintptr(DC), uintptr(X), uintptr(Y))
	return bool(api.GoBool(r1))
}

func LoadBitmap(hInstance types.THandle, lpBitmapName types.PChar) types.HBITMAP {
	r1, _, _ := winapiImportAPI().Proc(winAPI_LoadBitmap).Call(uintptr(hInstance), api.PascalStr(lpBitmapName))
	return types.HBITMAP(r1)
}

func LoadIcon(hInstance types.THandle, lpIconName types.PChar) types.HICON {
	r1, _, _ := winapiImportAPI().Proc(winAPI_LoadIcon).Call(uintptr(hInstance), api.PascalStr(lpIconName))
	return types.HICON(r1)
}

func MaskBltRop(DestDC types.HDC, X, Y, Width, Height int32, SrcDC types.HDC, XSrc, YSrc int32, Mask types.HBITMAP, XMask, YMask int32, Rop types.DWORD) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_MaskBltRop).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(Mask), uintptr(XMask), uintptr(YMask), uintptr(Rop))
	return bool(api.GoBool(r1))
}

func MaskBlt(DestDC types.HDC, X, Y, Width, Height int32, SrcDC types.HDC, XSrc, YSrc int32, Mask types.HBITMAP, XMask, YMask int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_MaskBlt).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(Mask), uintptr(XMask), uintptr(YMask))
	return bool(api.GoBool(r1))
}

func MessageBox(hWnd types.HWND, lpText, lpCaption types.PChar, uType types.Cardinal) int32 { //def MB_OK
	r1, _, _ := winapiImportAPI().Proc(winAPI_MessageBox).Call(uintptr(hWnd), api.PascalStr(lpText), api.PascalStr(lpCaption), uintptr(uType))
	return int32(r1)
}

func MonitorFromPoint(ptScreenCoords types.TPoint, dwFlags types.DWORD) types.HMONITOR {
	r1, _, _ := winapiImportAPI().Proc(winAPI_MonitorFromPoint).Call(uintptr(unsafe.Pointer(&ptScreenCoords)), uintptr(dwFlags))
	return types.HMONITOR(r1)
}

func MonitorFromRect(lprcScreenCoords types.TRect, dwFlags types.DWORD) types.HMONITOR {
	r1, _, _ := winapiImportAPI().Proc(winAPI_MonitorFromRect).Call(uintptr(unsafe.Pointer(&lprcScreenCoords)), uintptr(dwFlags))
	return types.HMONITOR(r1)
}

func MonitorFromWindow(hWnd types.HWND, dwFlags types.DWORD) types.HMONITOR {
	r1, _, _ := winapiImportAPI().Proc(winAPI_MonitorFromWindow).Call(uintptr(hWnd), uintptr(dwFlags))
	return types.HMONITOR(r1)
}

func MoveToEx(DC types.HDC, X, Y int32, OldPoint types.TPoint) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_MoveToEx).Call(uintptr(DC), uintptr(X), uintptr(Y), uintptr(unsafe.Pointer(&OldPoint)))
	return bool(api.GoBool(r1))
}

func OffsetRgn(RGN types.HRGN, nXOffset, nYOffset int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_OffsetRgn).Call(RGN, uintptr(nXOffset), uintptr(nYOffset))
	return int32(r1)
}

func PaintRgn(DC types.HDC, RGN types.HRGN) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_PaintRgn).Call(uintptr(DC), RGN)
	return bool(api.GoBool(r1))
}

func Pie(DC types.HDC, x1, y1, x2, y2, sx, sy, ex, ey int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_Pie).Call(uintptr(DC), uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2), uintptr(sx), uintptr(sy), uintptr(ex), uintptr(ey))
	return bool(api.GoBool(r1))
}

func PolyBezier(DC types.HDC, Points types.TPoint, NumPts int32, Filled, Continuous bool) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_PolyBezier).Call(uintptr(DC), uintptr(unsafe.Pointer(&Points)), uintptr(NumPts), api.PascalBool(bool(Filled)), api.PascalBool(bool(Continuous)))
	return bool(api.GoBool(r1))
}

func Polygon(DC types.HDC, Points types.TPoint, NumPts int32, Winding bool) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_Polygon).Call(uintptr(DC), uintptr(unsafe.Pointer(&Points)), uintptr(NumPts), api.PascalBool(bool(Winding)))
	return bool(api.GoBool(r1))
}

func Polyline(DC types.HDC, Points types.TPoint, NumPts int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_Polyline).Call(uintptr(DC), uintptr(unsafe.Pointer(&Points)), uintptr(NumPts))
	return bool(api.GoBool(r1))
}

func PostMessage(Handle types.HWND, Msg types.Cardinal, WParam types.WPARAM, LParam types.LPARAM) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_PostMessage).Call(uintptr(Handle), uintptr(Msg), uintptr(WParam), uintptr(LParam))
	return bool(api.GoBool(r1))
}

func RealizePalette(DC types.HDC) types.Cardinal {
	r1, _, _ := winapiImportAPI().Proc(winAPI_RealizePalette).Call(uintptr(DC))
	return types.Cardinal(r1)
}

func Rectangle(DC types.HDC, X1, Y1, X2, Y2 int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_Rectangle).Call(uintptr(DC), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2))
	return bool(api.GoBool(r1))
}

func RectInRegion(RGN types.HRGN, ARect types.TRect) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_RectInRegion).Call(RGN, uintptr(unsafe.Pointer(&ARect)))
	return bool(api.GoBool(r1))
}

func RectVisible(DC types.HDC, ARect types.TRect) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_RectVisible).Call(uintptr(DC), uintptr(unsafe.Pointer(&ARect)))
	return bool(api.GoBool(r1))
}

func RedrawWindow(Wnd types.HWND, lprcUpdate types.TRect, hrgnUpdate types.HRGN, flags uint32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_RedrawWindow).Call(uintptr(Wnd), uintptr(unsafe.Pointer(&lprcUpdate)), hrgnUpdate, uintptr(flags))
	return bool(api.GoBool(r1))
}

func ReleaseCapture() bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ReleaseCapture).Call()
	return bool(api.GoBool(r1))
}

func ReleaseDC(hWnd types.HWND, DC types.HDC) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ReleaseDC).Call(uintptr(hWnd), uintptr(DC))
	return int32(r1)
}

func RemoveProp(Handle types.HWND, Str types.PChar) types.THandle {
	r1, _, _ := winapiImportAPI().Proc(winAPI_RemoveProp).Call(uintptr(Handle), api.PascalStr(Str))
	return types.THandle(r1)
}

func RestoreDC(DC types.HDC, SavedDC int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_RestoreDC).Call(uintptr(DC), uintptr(SavedDC))
	return bool(api.GoBool(r1))
}

func RoundRect(DC types.HDC, X1, Y1, X2, Y2 int32, RX, RY int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_RoundRect).Call(uintptr(DC), uintptr(X1), uintptr(Y1), uintptr(X2), uintptr(Y2), uintptr(RX), uintptr(RY))
	return bool(api.GoBool(r1))
}

func SaveDC(DC types.HDC) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SaveDC).Call(uintptr(DC))
	return int32(r1)
}

func ScrollWindowEx(hWnd types.HWND, dx, dy int32, prcScroll, prcClip types.TRect, hrgnUpdate types.HRGN, prcUpdate types.TRect, flags uint32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ScrollWindowEx).Call(uintptr(hWnd), uintptr(dx), uintptr(dy), uintptr(unsafe.Pointer(&prcScroll)), uintptr(unsafe.Pointer(&prcClip)), hrgnUpdate, uintptr(unsafe.Pointer(&prcUpdate)), uintptr(flags))
	return bool(api.GoBool(r1))
}

func SelectClipRGN(DC types.HDC, RGN types.HRGN) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SelectClipRGN).Call(uintptr(DC), RGN)
	return int32(r1)
}

func SelectObject(DC types.HDC, GDIObj types.HGDIOBJ) types.HGDIOBJ {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SelectObject).Call(uintptr(DC), uintptr(GDIObj))
	return types.HGDIOBJ(r1)
}

func SelectPalette(DC types.HDC, Palette types.HPALETTE, ForceBackground bool) types.HPALETTE {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SelectPalette).Call(uintptr(DC), uintptr(Palette), api.PascalBool(bool(ForceBackground)))
	return types.HPALETTE(r1)
}

func SendMessage(HandleWnd types.HWND, Msg types.Cardinal, WParam types.WPARAM, LParam types.LPARAM) types.LResult {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SendMessage).Call(uintptr(HandleWnd), uintptr(Msg), uintptr(WParam), uintptr(LParam))
	return types.LResult(r1)
}

func SetActiveWindow(Handle types.HWND) types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetActiveWindow).Call(uintptr(Handle))
	return types.HWND(r1)
}

func SetBkColor(DC types.HDC, Color types.TColorRef) types.TColorRef { //pbd
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetBkColor).Call(uintptr(DC), uintptr(Color))
	return types.TColorRef(r1)
}

func SetBkMode(DC types.HDC, bkMode int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetBkMode).Call(uintptr(DC), uintptr(bkMode))
	return int32(r1)
}

func SetCapture(AHandle types.HWND) types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetCapture).Call(uintptr(AHandle))
	return types.HWND(r1)
}

func SetCaretPos(X, Y int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetCaretPos).Call(uintptr(X), uintptr(Y))
	return bool(api.GoBool(r1))
}

func SetCaretPosEx(handle types.HWND, X, Y int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetCaretPosEx).Call(uintptr(handle), uintptr(X), uintptr(Y))
	return bool(api.GoBool(r1))
}

func SetCursorPos(X, Y int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetCursorPos).Call(uintptr(X), uintptr(Y))
	return bool(api.GoBool(r1))
}

func SetFocus(hWnd types.HWND) types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetFocus).Call(uintptr(hWnd))
	return types.HWND(r1)
}

func SetForegroundWindow(hWnd types.HWND) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetForegroundWindow).Call(uintptr(hWnd))
	return bool(api.GoBool(r1))
}

func SetMapMode(DC types.HDC, fnMapMode int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetMapMode).Call(uintptr(DC), uintptr(fnMapMode))
	return int32(r1)
}

func SetMenu(AWindowHandle types.HWND, AMenuHandle types.HMENU) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetMenu).Call(uintptr(AWindowHandle), uintptr(AMenuHandle))
	return bool(api.GoBool(r1))
}

func SetParent(hWndChild types.HWND, hWndParent types.HWND) types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetParent).Call(uintptr(hWndChild), uintptr(hWndParent))
	return types.HWND(r1)
}

func SetProp(Handle types.HWND, Str types.PChar, Data types.Pointer) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetProp).Call(uintptr(Handle), api.PascalStr(Str), uintptr(Data))
	return bool(api.GoBool(r1))
}

func SetROP2(DC types.HDC, Mode int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetROP2).Call(uintptr(DC), uintptr(Mode))
	return int32(r1)
}

func SetScrollInfo(Handle types.HWND, SBStyle int32, ScrollInfo types.TagScrollInfo, Redraw bool) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetScrollInfo).Call(uintptr(Handle), uintptr(SBStyle), uintptr(unsafe.Pointer(&ScrollInfo)), api.PascalBool(bool(Redraw)))
	return int32(r1)
}

func SetStretchBltMode(DC types.HDC, StretchMode int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetStretchBltMode).Call(uintptr(DC), uintptr(StretchMode))
	return int32(r1)
}

func SetTextCharacterExtra(_hdc types.HDC, nCharExtra int32) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetTextCharacterExtra).Call(uintptr(_hdc), uintptr(nCharExtra))
	return int32(r1)
}

func SetTextColor(DC types.HDC, Color types.TColorRef) types.TColorRef {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetTextColor).Call(uintptr(DC), uintptr(Color))
	return types.TColorRef(r1)
}

func SetWindowLong(Handle types.HWND, Idx int32, NewLong types.PtrInt) types.PtrInt {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetWindowLong).Call(uintptr(Handle), uintptr(Idx), uintptr(NewLong))
	return types.PtrInt(r1)
}

func SetViewPortExtEx(DC types.HDC, XExtent, YExtent int32, OldSize types.TSize) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetViewPortExtEx).Call(uintptr(DC), uintptr(XExtent), uintptr(YExtent), uintptr(unsafe.Pointer(&OldSize)))
	return bool(api.GoBool(r1))
}

func SetViewPortOrgEx(DC types.HDC, NewX, NewY int32, OldPoint types.TPoint) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetViewPortOrgEx).Call(uintptr(DC), uintptr(NewX), uintptr(NewY), uintptr(unsafe.Pointer(&OldPoint)))
	return bool(api.GoBool(r1))
}

func SetWindowExtEx(DC types.HDC, XExtent, YExtent int32, OldSize types.TSize) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetWindowExtEx).Call(uintptr(DC), uintptr(XExtent), uintptr(YExtent), uintptr(unsafe.Pointer(&OldSize)))
	return bool(api.GoBool(r1))
}

func SetWindowOrgEx(dc types.HDC, NewX, NewY int32, OldPoint types.TPoint) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetWindowOrgEx).Call(uintptr(dc), uintptr(NewX), uintptr(NewY), uintptr(unsafe.Pointer(&OldPoint)))
	return bool(api.GoBool(r1))
}

func SetWindowPos(hWnd types.HWND, hWndInsertAfter types.HWND, X, Y, cx, cy int32, uFlags uint32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SetWindowPos).Call(uintptr(hWnd), uintptr(hWndInsertAfter), uintptr(X), uintptr(Y), uintptr(cx), uintptr(cy), uintptr(uFlags))
	return bool(api.GoBool(r1))
}

func ShowCaret(hWnd types.HWND) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ShowCaret).Call(uintptr(hWnd))
	return bool(api.GoBool(r1))
}

func ShowScrollBar(Handle types.HWND, wBar int32, bShow bool) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ShowScrollBar).Call(uintptr(Handle), uintptr(wBar), api.PascalBool(bool(bShow)))
	return bool(api.GoBool(r1))
}

func ShowWindow(hWnd types.HWND, nCmdShow int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_ShowWindow).Call(uintptr(hWnd), uintptr(nCmdShow))
	return bool(api.GoBool(r1))
}

func StretchBlt(DestDC types.HDC, X, Y, Width, Height int32, SrcDC types.HDC, XSrc, YSrc, SrcWidth, SrcHeight int32, Rop types.Cardinal) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_StretchBlt).Call(uintptr(DestDC), uintptr(X), uintptr(Y), uintptr(Width), uintptr(Height), uintptr(SrcDC), uintptr(XSrc), uintptr(YSrc), uintptr(SrcWidth), uintptr(SrcHeight), uintptr(Rop))
	return bool(api.GoBool(r1))
}

func StretchDIBits(DC types.HDC, DestX, DestY, DestWidth, DestHeight, SrcX, SrcY, SrcWidth, SrcHeight int32, Bits types.Pointer, BitsInfo types.TagBitmapInfo, Usage uint32, Rop types.DWORD) int32 {
	r1, _, _ := winapiImportAPI().Proc(winAPI_StretchDIBits).Call(uintptr(DC), uintptr(DestX), uintptr(DestY), uintptr(DestWidth), uintptr(DestHeight), uintptr(SrcX), uintptr(SrcY), uintptr(SrcWidth), uintptr(SrcHeight), uintptr(Bits), uintptr(unsafe.Pointer(&BitsInfo)), uintptr(Usage), uintptr(Rop))
	return int32(r1)
}

func SystemParametersInfo(uiAction types.DWORD, uiParam types.DWORD, pvParam types.Pointer, fWinIni types.DWORD) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_SystemParametersInfo).Call(uintptr(uiAction), uintptr(uiParam), uintptr(pvParam), uintptr(fWinIni))
	return bool(api.GoBool(r1))
}

func TextOut(DC types.HDC, X, Y int32, Str types.PChar, Count int32) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_TextOut).Call(uintptr(DC), uintptr(X), uintptr(Y), api.PascalStr(Str), uintptr(Count))
	return bool(api.GoBool(r1))
}

func UpdateWindow(Handle types.HWND) bool {
	r1, _, _ := winapiImportAPI().Proc(winAPI_UpdateWindow).Call(uintptr(Handle))
	return bool(api.GoBool(r1))
}

func WindowFromPoint(Point types.TPoint) types.HWND {
	r1, _, _ := winapiImportAPI().Proc(winAPI_WindowFromPoint).Call(uintptr(unsafe.Pointer(&Point)))
	return types.HWND(r1)
}

func ScalePercent() float32 {
	dc := GetDC(0)
	dpiX := GetDeviceCaps(dc, LOGPIXELSX)
	ReleaseDC(0, dc)
	return float32(dpiX) / 96.0
}

var (
	winapiImport       *imports.Imports = nil
	winapiImportTables                  = []*imports.Table{
		imports.NewTable("WinAPI_CreateRectRgn", 0),
		imports.NewTable("WinAPI_SetRectRgn", 0),
		imports.NewTable("WinAPI_DeleteObject", 0),
		imports.NewTable("WinAPI_CombineRgn", 0),
		imports.NewTable("WinAPI_SetDraggableRegions", 0),
		imports.NewTable("WinAPI_PtInRegion", 0),
		imports.NewTable("WinAPI_ScreenToClient", 0),
		imports.NewTable("WinAPI_ClientToScreen", 0),
		imports.NewTable("WinAPI_DefWindowProc", 0),
		imports.NewTable("WinAPI_DefSubclassProc", 0),
		imports.NewTable("WinAPI_CreateRoundRectRgn", 0),
		imports.NewTable("WinAPI_SetWindowRgn", 0),
		imports.NewTable("WinAPI_SetCursor", 0),
		imports.NewTable("WinAPI_LoadCursor", 0),
		imports.NewTable("WinAPI_OnPaint", 0),
		imports.NewTable("WinAPI_EndPaint", 0),
		imports.NewTable("WinAPI_BeginPaint", 0),
		imports.NewTable("WinAPI_Arc", 0),
		imports.NewTable("WinAPI_AngleChord", 0),
		imports.NewTable("WinAPI_CallNextHookEx", 0),
		imports.NewTable("WinAPI_CallWindowProc", 0),
		imports.NewTable("WinAPI_BitBlt", 0),
		imports.NewTable("WinAPI_CreateBitmap", 0),
		imports.NewTable("WinAPI_CreateBrushIndirect", 0),
		imports.NewTable("WinAPI_CreateBrushWithRadialGradient", 0),
		imports.NewTable("WinAPI_CreateCaret", 0),
		imports.NewTable("WinAPI_CreateCompatibleBitmap", 0),
		imports.NewTable("WinAPI_CreateCompatibleDC", 0),
		imports.NewTable("WinAPI_CreateDIBitmap", 0),
		imports.NewTable("WinAPI_CreateDIBSection", 0),
		imports.NewTable("WinAPI_CreateEllipticRgn", 0),
		imports.NewTable("WinAPI_CreateFontIndirect", 0),
		imports.NewTable("WinAPI_CreateFontIndirectEx", 0),
		imports.NewTable("WinAPI_CreateIconIndirect", 0),
		imports.NewTable("WinAPI_CreatePalette", 0),
		imports.NewTable("WinAPI_CreatePatternBrush", 0),
		imports.NewTable("WinAPI_CreatePenIndirect", 0),
		imports.NewTable("WinAPI_CreatePolygonRgn", 0),
		imports.NewTable("WinAPI_DeleteCriticalSection", 0),
		imports.NewTable("WinAPI_DeleteDC", 0),
		imports.NewTable("WinAPI_DestroyCaret", 0),
		imports.NewTable("WinAPI_DestroyCursor", 0),
		imports.NewTable("WinAPI_DestroyIcon", 0),
		imports.NewTable("WinAPI_DrawFrameControl", 0),
		imports.NewTable("WinAPI_DrawFocusRect", 0),
		imports.NewTable("WinAPI_DrawEdge", 0),
		imports.NewTable("WinAPI_DrawText", 0),
		imports.NewTable("WinAPI_EnableScrollBar", 0),
		imports.NewTable("WinAPI_EnableWindow", 0),
		imports.NewTable("WinAPI_EnterCriticalSection", 0),
		imports.NewTable("WinAPI_Ellipse", 0),
		imports.NewTable("WinAPI_EqualRgn", 0),
		imports.NewTable("WinAPI_ExcludeClipRect", 0),
		imports.NewTable("WinAPI_ExtCreatePen", 0),
		imports.NewTable("WinAPI_ExtTextOut", 0),
		imports.NewTable("WinAPI_ExtSelectClipRGN", 0),
		imports.NewTable("WinAPI_FillRect", 0),
		imports.NewTable("WinAPI_FillRgn", 0),
		imports.NewTable("WinAPI_FloodFill", 0),
		imports.NewTable("WinAPI_FrameRect", 0),
		imports.NewTable("WinAPI_GetActiveWindow", 0),
		imports.NewTable("WinAPI_GetBitmapBits", 0),
		imports.NewTable("WinAPI_GetBkColor", 0),
		imports.NewTable("WinAPI_GetCapture", 0),
		imports.NewTable("WinAPI_GetCaretPos", 0),
		imports.NewTable("WinAPI_GetClientRect", 0),
		imports.NewTable("WinAPI_GetClipBox", 0),
		imports.NewTable("WinAPI_GetClipRGN", 0),
		imports.NewTable("WinAPI_GetCurrentObject", 0),
		imports.NewTable("WinAPI_GetCursorPos", 0),
		imports.NewTable("WinAPI_GetDC", 0),
		imports.NewTable("WinAPI_GetDeviceCaps", 0),
		imports.NewTable("WinAPI_GetDIBits", 0),
		imports.NewTable("WinAPI_GetDoubleClickTime", 0),
		imports.NewTable("WinAPI_GetFocus", 0),
		imports.NewTable("WinAPI_GetFontLanguageInfo", 0),
		imports.NewTable("WinAPI_GetForegroundWindow", 0),
		imports.NewTable("WinAPI_GetIconInfo", 0),
		imports.NewTable("WinAPI_GetKeyState", 0),
		imports.NewTable("WinAPI_GetMapMode", 0),
		imports.NewTable("WinAPI_GetMonitorInfo", 0),
		imports.NewTable("WinAPI_GetDpiForMonitor", 0),
		imports.NewTable("WinAPI_GetObject", 0),
		imports.NewTable("WinAPI_GetParent", 0),
		imports.NewTable("WinAPI_GetProp", 0),
		imports.NewTable("WinAPI_GetRgnBox", 0),
		imports.NewTable("WinAPI_GetROP2", 0),
		imports.NewTable("WinAPI_GetScrollInfo", 0),
		imports.NewTable("WinAPI_GetStockObject", 0),
		imports.NewTable("WinAPI_GetSysColor", 0),
		imports.NewTable("WinAPI_GetSysColorBrush", 0),
		imports.NewTable("WinAPI_GetSystemMetrics", 0),
		imports.NewTable("WinAPI_GetTextColor", 0),
		imports.NewTable("WinAPI_GetTextExtentExPoint", 0),
		imports.NewTable("WinAPI_GetTextExtentPoint", 0),
		imports.NewTable("WinAPI_GetTextExtentPoint32", 0),
		imports.NewTable("WinAPI_GetTextMetrics", 0),
		imports.NewTable("WinAPI_GetViewPortExtEx", 0),
		imports.NewTable("WinAPI_GetViewPortOrgEx", 0),
		imports.NewTable("WinAPI_GetWindowExtEx", 0),
		imports.NewTable("WinAPI_GetWindowLong", 0),
		imports.NewTable("WinAPI_GetWindowRect", 0),
		imports.NewTable("WinAPI_GetWindowSize", 0),
		imports.NewTable("WinAPI_GetWindowOrgEx", 0),
		imports.NewTable("WinAPI_GradientFill", 0),
		imports.NewTable("WinAPI_HideCaret", 0),
		imports.NewTable("WinAPI_InitializeCriticalSection", 0),
		imports.NewTable("WinAPI_IntersectClipRect", 0),
		imports.NewTable("WinAPI_InvalidateRect", 0),
		imports.NewTable("WinAPI_InvalidateRgn", 0),
		imports.NewTable("WinAPI_IsDBCSLeadByte", 0),
		imports.NewTable("WinAPI_IsIconic", 0),
		imports.NewTable("WinAPI_IsWindow", 0),
		imports.NewTable("WinAPI_IsWindowEnabled", 0),
		imports.NewTable("WinAPI_IsWindowVisible", 0),
		imports.NewTable("WinAPI_IsZoomed", 0),
		imports.NewTable("WinAPI_LeaveCriticalSection", 0),
		imports.NewTable("WinAPI_LineTo", 0),
		imports.NewTable("WinAPI_LoadBitmap", 0),
		imports.NewTable("WinAPI_LoadIcon", 0),
		imports.NewTable("WinAPI_MaskBltRop", 0),
		imports.NewTable("WinAPI_MaskBlt", 0),
		imports.NewTable("WinAPI_MessageBox", 0),
		imports.NewTable("WinAPI_MonitorFromPoint", 0),
		imports.NewTable("WinAPI_MonitorFromRect", 0),
		imports.NewTable("WinAPI_MonitorFromWindow", 0),
		imports.NewTable("WinAPI_MoveToEx", 0),
		imports.NewTable("WinAPI_OffsetRgn", 0),
		imports.NewTable("WinAPI_PaintRgn", 0),
		imports.NewTable("WinAPI_Pie", 0),
		imports.NewTable("WinAPI_PolyBezier", 0),
		imports.NewTable("WinAPI_Polygon", 0),
		imports.NewTable("WinAPI_Polyline", 0),
		imports.NewTable("WinAPI_PostMessage", 0),
		imports.NewTable("WinAPI_RealizePalette", 0),
		imports.NewTable("WinAPI_Rectangle", 0),
		imports.NewTable("WinAPI_RectInRegion", 0),
		imports.NewTable("WinAPI_RectVisible", 0),
		imports.NewTable("WinAPI_RedrawWindow", 0),
		imports.NewTable("WinAPI_ReleaseCapture", 0),
		imports.NewTable("WinAPI_ReleaseDC", 0),
		imports.NewTable("WinAPI_RemoveProp", 0),
		imports.NewTable("WinAPI_RestoreDC", 0),
		imports.NewTable("WinAPI_RoundRect", 0),
		imports.NewTable("WinAPI_SaveDC", 0),
		imports.NewTable("WinAPI_ScrollWindowEx", 0),
		imports.NewTable("WinAPI_SelectClipRGN", 0),
		imports.NewTable("WinAPI_SelectObject", 0),
		imports.NewTable("WinAPI_SelectPalette", 0),
		imports.NewTable("WinAPI_SendMessage", 0),
		imports.NewTable("WinAPI_SetActiveWindow", 0),
		imports.NewTable("WinAPI_SetBkColor", 0),
		imports.NewTable("WinAPI_SetBkMode", 0),
		imports.NewTable("WinAPI_SetCapture", 0),
		imports.NewTable("WinAPI_SetCaretPos", 0),
		imports.NewTable("WinAPI_SetCaretPosEx", 0),
		imports.NewTable("WinAPI_SetCursorPos", 0),
		imports.NewTable("WinAPI_SetFocus", 0),
		imports.NewTable("WinAPI_SetForegroundWindow", 0),
		imports.NewTable("WinAPI_SetMapMode", 0),
		imports.NewTable("WinAPI_SetMenu", 0),
		imports.NewTable("WinAPI_SetParent", 0),
		imports.NewTable("WinAPI_SetProp", 0),
		imports.NewTable("WinAPI_SetROP2", 0),
		imports.NewTable("WinAPI_SetScrollInfo", 0),
		imports.NewTable("WinAPI_SetStretchBltMode", 0),
		imports.NewTable("WinAPI_SetTextCharacterExtra", 0),
		imports.NewTable("WinAPI_SetTextColor", 0),
		imports.NewTable("WinAPI_SetWindowLong", 0),
		imports.NewTable("WinAPI_SetViewPortExtEx", 0),
		imports.NewTable("WinAPI_SetViewPortOrgEx", 0),
		imports.NewTable("WinAPI_SetWindowExtEx", 0),
		imports.NewTable("WinAPI_SetWindowOrgEx", 0),
		imports.NewTable("WinAPI_SetWindowPos", 0),
		imports.NewTable("WinAPI_ShowCaret", 0),
		imports.NewTable("WinAPI_ShowScrollBar", 0),
		imports.NewTable("WinAPI_ShowWindow", 0),
		imports.NewTable("WinAPI_StretchBlt", 0),
		imports.NewTable("WinAPI_StretchDIBits", 0),
		imports.NewTable("WinAPI_SystemParametersInfo", 0),
		imports.NewTable("WinAPI_TextOut", 0),
		imports.NewTable("WinAPI_UpdateWindow", 0),
		imports.NewTable("WinAPI_WindowFromPoint", 0),
		imports.NewTable("WinAPI_GetWindowLongPtr", 0),
		imports.NewTable("WinAPI_SetWindowLongPtr", 0),
		imports.NewTable("WinAPI_GetClassLongPtr", 0),
		imports.NewTable("WinAPI_SetClassLongPtr", 0),
		imports.NewTable("WinAPI_FindWindow", 0),
		imports.NewTable("WinAPI_FindWindowEx", 0),
		imports.NewTable("WinAPI_SetWindowText", 0),
		imports.NewTable("WinAPI_GetWindowText", 0),
		imports.NewTable("WinAPI_GetWindowTextLength", 0),
	}
)

func winapiImportAPI() *imports.Imports {
	if winapiImport == nil {
		winapiImport = api.NewDefaultImports()
		winapiImport.SetImportTable(winapiImportTables)
		winapiImportTables = nil
	}
	return winapiImport
}
