//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package winapi

import "github.com/energye/lcl/api/imports"

var winapiImportDefs = []*imports.Table{
	imports.NewTable("CEF_Win_CreateRectRgn", 0),
	imports.NewTable("CEF_Win_SetRectRgn", 0),
	imports.NewTable("CEF_Win_DeleteObject", 0),
	imports.NewTable("CEF_Win_CombineRgn", 0),
	imports.NewTable("CEF_Win_SetDraggableRegions", 0),
	imports.NewTable("CEF_Win_PtInRegion", 0),
	imports.NewTable("CEF_Win_ScreenToClient", 0),
	imports.NewTable("CEF_Win_ClientToScreen", 0),
	imports.NewTable("CEF_Win_DefWindowProc", 0),
	imports.NewTable("CEF_Win_DefSubclassProc", 0),
	imports.NewTable("CEF_Win_CreateRoundRectRgn", 0),
	imports.NewTable("CEF_Win_SetWindowRgn", 0),
	imports.NewTable("CEF_Win_SetCursor", 0),
	imports.NewTable("CEF_Win_LoadCursor", 0),
	imports.NewTable("CEF_Win_OnPaint", 0),
	imports.NewTable("CEF_Win_EndPaint", 0),
	imports.NewTable("CEF_Win_BeginPaint", 0),
	imports.NewTable("CEF_Win_Arc", 0),
	imports.NewTable("CEF_Win_AngleChord", 0),
	imports.NewTable("CEF_Win_CallNextHookEx", 0),
	imports.NewTable("CEF_Win_CallWindowProc", 0),
	imports.NewTable("CEF_Win_BitBlt", 0),
	imports.NewTable("CEF_Win_CreateBitmap", 0),
	imports.NewTable("CEF_Win_CreateBrushIndirect", 0),
	imports.NewTable("CEF_Win_CreateBrushWithRadialGradient", 0),
	imports.NewTable("CEF_Win_CreateCaret", 0),
	imports.NewTable("CEF_Win_CreateCompatibleBitmap", 0),
	imports.NewTable("CEF_Win_CreateCompatibleDC", 0),
	imports.NewTable("CEF_Win_CreateDIBitmap", 0),
	imports.NewTable("CEF_Win_CreateDIBSection", 0),
	imports.NewTable("CEF_Win_CreateEllipticRgn", 0),
	imports.NewTable("CEF_Win_CreateFontIndirect", 0),
	imports.NewTable("CEF_Win_CreateFontIndirectEx", 0),
	imports.NewTable("CEF_Win_CreateIconIndirect", 0),
	imports.NewTable("CEF_Win_CreatePalette", 0),
	imports.NewTable("CEF_Win_CreatePatternBrush", 0),
	imports.NewTable("CEF_Win_CreatePenIndirect", 0),
	imports.NewTable("CEF_Win_CreatePolygonRgn", 0),
	imports.NewTable("CEF_Win_DeleteCriticalSection", 0),
	imports.NewTable("CEF_Win_DeleteDC", 0),
	imports.NewTable("CEF_Win_DestroyCaret", 0),
	imports.NewTable("CEF_Win_DestroyCursor", 0),
	imports.NewTable("CEF_Win_DestroyIcon", 0),
	imports.NewTable("CEF_Win_DrawFrameControl", 0),
	imports.NewTable("CEF_Win_DrawFocusRect", 0),
	imports.NewTable("CEF_Win_DrawEdge", 0),
	imports.NewTable("CEF_Win_DrawText", 0),
	imports.NewTable("CEF_Win_EnableScrollBar", 0),
	imports.NewTable("CEF_Win_EnableWindow", 0),
	imports.NewTable("CEF_Win_EnterCriticalSection", 0),
	imports.NewTable("CEF_Win_EnumDisplayMonitors", 0),             // no impl
	imports.NewTable("CEF_Win_EnumFontFamilies", 0),                // no impl
	imports.NewTable("CEF_Win_EnumFontFamiliesEx", 0),              // no impl
	imports.NewTable("CEF_Win_EnumDisplayMonitorsCallbackFree", 0), // no impl
	imports.NewTable("CEF_Win_EnumFontFamiliesCallbackFree", 0),    // no impl
	imports.NewTable("CEF_Win_EnumFontFamiliesExCallbackFree", 0),  // no impl
	imports.NewTable("CEF_Win_Ellipse", 0),
	imports.NewTable("CEF_Win_EqualRgn", 0),
	imports.NewTable("CEF_Win_ExcludeClipRect", 0),
	imports.NewTable("CEF_Win_ExtCreatePen", 0),
	imports.NewTable("CEF_Win_ExtTextOut", 0),
	imports.NewTable("CEF_Win_ExtSelectClipRGN", 0),
	imports.NewTable("CEF_Win_FillRect", 0),
	imports.NewTable("CEF_Win_FillRgn", 0),
	imports.NewTable("CEF_Win_FloodFill", 0),
	imports.NewTable("CEF_Win_FrameRect", 0),
	imports.NewTable("CEF_Win_GetActiveWindow", 0),
	imports.NewTable("CEF_Win_GetBitmapBits", 0),
	imports.NewTable("CEF_Win_GetBkColor", 0),
	imports.NewTable("CEF_Win_GetCapture", 0),
	imports.NewTable("CEF_Win_GetCaretPos", 0),
	imports.NewTable("CEF_Win_GetClientRect", 0),
	imports.NewTable("CEF_Win_GetClipBox", 0),
	imports.NewTable("CEF_Win_GetClipRGN", 0),
	imports.NewTable("CEF_Win_GetCurrentObject", 0),
	imports.NewTable("CEF_Win_GetCursorPos", 0),
	imports.NewTable("CEF_Win_GetDC", 0),
	imports.NewTable("CEF_Win_GetDeviceCaps", 0),
	imports.NewTable("CEF_Win_GetDIBits", 0),
	imports.NewTable("CEF_Win_GetDoubleClickTime", 0),
	imports.NewTable("CEF_Win_GetFocus", 0),
	imports.NewTable("CEF_Win_GetFontLanguageInfo", 0),
	imports.NewTable("CEF_Win_GetForegroundWindow", 0),
	imports.NewTable("CEF_Win_GetIconInfo", 0),
	imports.NewTable("CEF_Win_GetKeyState", 0),
	imports.NewTable("CEF_Win_GetMapMode", 0),
	imports.NewTable("CEF_Win_GetMonitorInfo", 0),
	imports.NewTable("CEF_Win_GetDpiForMonitor", 0),
	imports.NewTable("CEF_Win_GetObject", 0),
	imports.NewTable("CEF_Win_GetParent", 0),
	imports.NewTable("CEF_Win_GetProp", 0),
	imports.NewTable("CEF_Win_GetRgnBox", 0),
	imports.NewTable("CEF_Win_GetROP2", 0),
	imports.NewTable("CEF_Win_GetScrollInfo", 0),
	imports.NewTable("CEF_Win_GetStockObject", 0),
	imports.NewTable("CEF_Win_GetSysColor", 0),
	imports.NewTable("CEF_Win_GetSysColorBrush", 0),
	imports.NewTable("CEF_Win_GetSystemMetrics", 0),
	imports.NewTable("CEF_Win_GetTextColor", 0),
	imports.NewTable("CEF_Win_GetTextExtentExPoint", 0),
	imports.NewTable("CEF_Win_GetTextExtentPoint", 0),
	imports.NewTable("CEF_Win_GetTextExtentPoint32", 0),
	imports.NewTable("CEF_Win_GetTextMetrics", 0),
	imports.NewTable("CEF_Win_GetViewPortExtEx", 0),
	imports.NewTable("CEF_Win_GetViewPortOrgEx", 0),
	imports.NewTable("CEF_Win_GetWindowExtEx", 0),
	imports.NewTable("CEF_Win_GetWindowLong", 0),
	imports.NewTable("CEF_Win_GetWindowRect", 0),
	imports.NewTable("CEF_Win_GetWindowSize", 0),
	imports.NewTable("CEF_Win_GetWindowOrgEx", 0),
	imports.NewTable("CEF_Win_GradientFill", 0),
	imports.NewTable("CEF_Win_HideCaret", 0),
	imports.NewTable("CEF_Win_InitializeCriticalSection", 0),
	imports.NewTable("CEF_Win_IntersectClipRect", 0),
	imports.NewTable("CEF_Win_InvalidateRect", 0),
	imports.NewTable("CEF_Win_InvalidateRgn", 0),
	imports.NewTable("CEF_Win_IsDBCSLeadByte", 0),
	imports.NewTable("CEF_Win_IsIconic", 0),
	imports.NewTable("CEF_Win_IsWindow", 0),
	imports.NewTable("CEF_Win_IsWindowEnabled", 0),
	imports.NewTable("CEF_Win_IsWindowVisible", 0),
	imports.NewTable("CEF_Win_IsZoomed", 0),
	imports.NewTable("CEF_Win_LeaveCriticalSection", 0),
	imports.NewTable("CEF_Win_LineTo", 0),
	imports.NewTable("CEF_Win_LoadBitmap", 0),
	imports.NewTable("CEF_Win_LoadIcon", 0),
	imports.NewTable("CEF_Win_MaskBltRop", 0),
	imports.NewTable("CEF_Win_MaskBlt", 0),
	imports.NewTable("CEF_Win_MessageBox", 0),
	imports.NewTable("CEF_Win_MonitorFromPoint", 0),
	imports.NewTable("CEF_Win_MonitorFromRect", 0),
	imports.NewTable("CEF_Win_MonitorFromWindow", 0),
	imports.NewTable("CEF_Win_MoveToEx", 0),
	imports.NewTable("CEF_Win_OffsetRgn", 0),
	imports.NewTable("CEF_Win_PaintRgn", 0),
	imports.NewTable("CEF_Win_Pie", 0),
	imports.NewTable("CEF_Win_PolyBezier", 0),
	imports.NewTable("CEF_Win_Polygon", 0),
	imports.NewTable("CEF_Win_Polyline", 0),
	imports.NewTable("CEF_Win_PostMessage", 0),
	imports.NewTable("CEF_Win_RealizePalette", 0),
	imports.NewTable("CEF_Win_Rectangle", 0),
	imports.NewTable("CEF_Win_RectInRegion", 0),
	imports.NewTable("CEF_Win_RectVisible", 0),
	imports.NewTable("CEF_Win_RedrawWindow", 0),
	imports.NewTable("CEF_Win_ReleaseCapture", 0),
	imports.NewTable("CEF_Win_ReleaseDC", 0),
	imports.NewTable("CEF_Win_RemoveProp", 0),
	imports.NewTable("CEF_Win_RestoreDC", 0),
	imports.NewTable("CEF_Win_RoundRect", 0),
	imports.NewTable("CEF_Win_SaveDC", 0),
	imports.NewTable("CEF_Win_ScrollWindowEx", 0),
	imports.NewTable("CEF_Win_SelectClipRGN", 0),
	imports.NewTable("CEF_Win_SelectObject", 0),
	imports.NewTable("CEF_Win_SelectPalette", 0),
	imports.NewTable("CEF_Win_SendMessage", 0),
	imports.NewTable("CEF_Win_SetActiveWindow", 0),
	imports.NewTable("CEF_Win_SetBkColor", 0),
	imports.NewTable("CEF_Win_SetBkMode", 0),
	imports.NewTable("CEF_Win_SetCapture", 0),
	imports.NewTable("CEF_Win_SetCaretPos", 0),
	imports.NewTable("CEF_Win_SetCaretPosEx", 0),
	imports.NewTable("CEF_Win_SetCursorPos", 0),
	imports.NewTable("CEF_Win_SetFocus", 0),
	imports.NewTable("CEF_Win_SetForegroundWindow", 0),
	imports.NewTable("CEF_Win_SetMapMode", 0),
	imports.NewTable("CEF_Win_SetMenu", 0),
	imports.NewTable("CEF_Win_SetParent", 0),
	imports.NewTable("CEF_Win_SetProp", 0),
	imports.NewTable("CEF_Win_SetROP2", 0),
	imports.NewTable("CEF_Win_SetScrollInfo", 0),
	imports.NewTable("CEF_Win_SetStretchBltMode", 0),
	imports.NewTable("CEF_Win_SetTextCharacterExtra", 0),
	imports.NewTable("CEF_Win_SetTextColor", 0),
	imports.NewTable("CEF_Win_SetWindowLong", 0),
	imports.NewTable("CEF_Win_SetViewPortExtEx", 0),
	imports.NewTable("CEF_Win_SetViewPortOrgEx", 0),
	imports.NewTable("CEF_Win_SetWindowExtEx", 0),
	imports.NewTable("CEF_Win_SetWindowOrgEx", 0),
	imports.NewTable("CEF_Win_SetWindowPos", 0),
	imports.NewTable("CEF_Win_ShowCaret", 0),
	imports.NewTable("CEF_Win_ShowScrollBar", 0),
	imports.NewTable("CEF_Win_ShowWindow", 0),
	imports.NewTable("CEF_Win_StretchBlt", 0),
	imports.NewTable("CEF_Win_StretchDIBits", 0),
	imports.NewTable("CEF_Win_SystemParametersInfo", 0),
	imports.NewTable("CEF_Win_TextOut", 0),
	imports.NewTable("CEF_Win_UpdateWindow", 0),
	imports.NewTable("CEF_Win_WindowFromPoint", 0),
	imports.NewTable("CEF_Win_GetWindowLongPtr", 0),
	imports.NewTable("CEF_Win_SetWindowLongPtr", 0),
	imports.NewTable("CEF_Win_GetClassLongPtr", 0),
	imports.NewTable("CEF_Win_SetClassLongPtr", 0),
	imports.NewTable("CEF_Win_FindWindow", 0),
	imports.NewTable("CEF_Win_FindWindowEx", 0),
	imports.NewTable("CEF_Win_SetWindowText", 0),
	imports.NewTable("CEF_Win_GetWindowText", 0),
	imports.NewTable("CEF_Win_GetWindowTextLength", 0),
}

// InitWinAPIPreDefsImport 初始化winapi预定义api
func InitWinAPIPreDefsImport(imp *imports.Imports) {
	imp.SetImportTable(winapiImportDefs)
	imp.SetOk(true)
}
