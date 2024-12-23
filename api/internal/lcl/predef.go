//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import "github.com/energye/lcl/api/imports"

// --------------
// --------------
// --------------
// --------------
// --------------
// --------------

var preImportDefs = []*imports.Table{
	imports.NewTable("Application_Instance", 0), //起始位置 按行号计算 20 = 0
	imports.NewTable("Application_CreateForm", 0),
	imports.NewTable("Application_Run", 0),
	imports.NewTable("Application_Initialize", 0),
	imports.NewTable("Application_SetRunLoopReceived", 0),
	imports.NewTable("Form_InheritedWndProc", 0),
	imports.NewTable("Form_Create2", 0),
	imports.NewTable("Form_SetOnWndProc", 0),
	imports.NewTable("Form_SetGoPtr", 0),
	imports.NewTable("SetEventCallback", 0),
	imports.NewTable("SetMessageCallback", 0),
	imports.NewTable("SetThreadSyncCallback", 0),
	imports.NewTable("SetRequestCallCreateParamsCallback", 0),
	imports.NewTable("SetRequestCallFormCreateCallback", 0),
	imports.NewTable("SetRemoveEventCallback", 0),
	imports.NewTable("SetThreadAsyncCallback", 0),
	imports.NewTable("DRunMainAsyncCall", 0),
	imports.NewTable("DGetStringArrOf", 0),
	imports.NewTable("DStrLen", 0),
	imports.NewTable("DMove", 0),
	imports.NewTable("DShowMessage", 0),
	imports.NewTable("DMessageDlg", 0),
	imports.NewTable("Mouse_Instance", 0),
	imports.NewTable("Screen_Instance", 0),
	imports.NewTable("DRunMainSyncCall", 0),
	imports.NewTable("DTextToShortCut", 0),
	imports.NewTable("DShortCutToText", 0),
	imports.NewTable("Clipboard_Instance", 0),
	imports.NewTable("DSetClipboard", 0),
	imports.NewTable("DRegisterClipboardFormat", 0),
	imports.NewTable("DPredefinedClipboardFormat", 0),
	imports.NewTable("DSysOpen", 0),
	imports.NewTable("DExtractFilePath", 0),
	imports.NewTable("DFileExists", 0),
	imports.NewTable("DSelectDirectory1", 0),
	imports.NewTable("DSelectDirectory2", 0),
	imports.NewTable("DInputBox", 0),
	imports.NewTable("DInputQuery", 0),
	imports.NewTable("DPasswordBox", 0),
	imports.NewTable("DInputCombo", 0),
	imports.NewTable("DInputComboEx", 0),
	imports.NewTable("DSysLocale", 0),
	imports.NewTable("DCreateURLShortCut", 0),
	imports.NewTable("DCreateShortCut", 0),
	imports.NewTable("DSetPropertyValue", 0),
	imports.NewTable("DSetPropertySecValue", 0),
	imports.NewTable("Printer_Instance", 0),
	imports.NewTable("Printer_SetPrinter", 0),
	imports.NewTable("DGUIDToString", 0),
	imports.NewTable("DStringToGUID", 0),
	imports.NewTable("DCreateGUID", 0),
	imports.NewTable("DGetLibResourceCount", 0),
	imports.NewTable("DGetLibResourceItem", 0),
	imports.NewTable("DModifyLibResource", 0),
	imports.NewTable("DLibAbout", 0),
	imports.NewTable("DLibStringEncoding", 0),
	imports.NewTable("GetLCLVersion", 0),
	imports.NewTable("DMainThreadId", 0),
	imports.NewTable("DCurrentThreadId", 0),
	imports.NewTable("DInitGoDll", 0),
	imports.NewTable("DFindControl", 0),
	imports.NewTable("DFindLCLControl", 0),
	imports.NewTable("DFindOwnerControl", 0),
	imports.NewTable("DFindControlAtPosition", 0),
	imports.NewTable("DFindLCLWindow", 0),
	imports.NewTable("DFindDragTarget", 0),
	imports.NewTable("ResFormLoadFromStream", 0),
	imports.NewTable("ResFormLoadFromFile", 0),
	imports.NewTable("ResFormLoadFromResourceName", 0),
	imports.NewTable("NSWindow_FromForm", 0),
	imports.NewTable("NSWindow_titleVisibility", 0),
	imports.NewTable("NSWindow_setTitleVisibility", 0),
	imports.NewTable("NSWindow_titlebarAppearsTransparent", 0),
	imports.NewTable("NSWindow_setTitlebarAppearsTransparent", 0),
	imports.NewTable("NSWindow_styleMask", 0),
	imports.NewTable("NSWindow_setStyleMask", 0),
	imports.NewTable("NSWindow_setRepresentedURL", 0),
	imports.NewTable("GdkWindow_GetXId", 0),
	imports.NewTable("GdkWindow_FromForm", 0),
	imports.NewTable("GtkWidget_GetGtkFixed", 0),
	imports.NewTable("GtkWidget_Window", 0),
	imports.NewTable("DSendMessage", 0),
	imports.NewTable("DPostMessage", 0),
	imports.NewTable("DIsIconic", 0),
	imports.NewTable("DIsWindow", 0),
	imports.NewTable("DIsZoomed", 0),
	imports.NewTable("DIsWindowVisible", 0),
	imports.NewTable("DGetDC", 0),
	imports.NewTable("DReleaseDC", 0),
	imports.NewTable("DSetForegroundWindow", 0),
	imports.NewTable("DWindowFromPoint", 0),
	imports.NewTable("SetExceptionHandlerCallback", 0),
	imports.NewTable("COMSequentialStream_Read", 0), // ISequentialStream
	imports.NewTable("COMSequentialStream_Write", 0),
	imports.NewTable("COMStream_Seek", 0), //SequentialStream
	imports.NewTable("COMStream_SetSize", 0),
	imports.NewTable("COMStream_CopyTo", 0),
	imports.NewTable("COMStream_Commit", 0),
	imports.NewTable("COMStream_Revert", 0),
	imports.NewTable("COMStream_LockRegion", 0),
	imports.NewTable("COMStream_UnlockRegion", 0),
	imports.NewTable("COMStream_Stat", 0),
	imports.NewTable("COMStream_Clone", 0),
	imports.NewTable("StreamAdapter_Clone", 0), // TStreamAdapter
	imports.NewTable("StreamAdapter_Commit", 0),
	imports.NewTable("StreamAdapter_CopyTo", 0),
	imports.NewTable("StreamAdapter_Create", 0),
	imports.NewTable("StreamAdapter_LockRegion", 0),
	imports.NewTable("StreamAdapter_Read", 0),
	imports.NewTable("StreamAdapter_Revert", 0),
	imports.NewTable("StreamAdapter_Seek", 0),
	imports.NewTable("StreamAdapter_SetSize", 0),
	imports.NewTable("StreamAdapter_Stat", 0),
	imports.NewTable("StreamAdapter_Stream", 0),
	imports.NewTable("StreamAdapter_StreamOwnership", 0),
	imports.NewTable("StreamAdapter_UnlockRegion", 0),
	imports.NewTable("StreamAdapter_Write", 0),
	imports.NewTable("InterfacedObject_Create", 0), // TInterfacedObject
	imports.NewTable("InterfacedObject_RefCount", 0),
	imports.NewTable("DFreeAndNil", 0),                          // FreeAndNil
	imports.NewTable("DSetNil", 0),                              // Obj := nil
	imports.NewTable("DToUnixTime", 0),                          // DateTime
	imports.NewTable("DUnixToTime", 0),                          // DateTime
	imports.NewTable("SetCEFEventCallback", 0),                  // CEF
	imports.NewTable("SetCEFRemoveEventCallback", 0),            // CEF
	imports.NewTable("SetWVEventCallback", 0),                   // WV
	imports.NewTable("SetWVRemoveEventCallback", 0),             // WV
	imports.NewTable("SetWKEventCallback", 0),                   // WK
	imports.NewTable("SetWKRemoveEventCallback", 0),             // WK
	imports.NewTable("VirtualTreeColumns_GetVisibleColumns", 0), // TVirtualTreeColumns
	imports.NewTable("VTImageInfo_Get", 0),                      // TVirtualTreeColumns
	imports.NewTable("VirtualNode_Create", 0),                   // PVirtualNode
	imports.NewTable("VirtualNode_Index", 0),                    // PVirtualNode
	imports.NewTable("VirtualNode_ChildCount", 0),               // PVirtualNode
	imports.NewTable("VirtualNode_NodeHeight", 0),               // PVirtualNode
	imports.NewTable("VirtualNode_States", 0),                   // PVirtualNode
	imports.NewTable("VirtualNode_Align", 0),                    // PVirtualNode
	imports.NewTable("VirtualNode_CheckState", 0),               // PVirtualNode
	imports.NewTable("VirtualNode_CheckType", 0),                // PVirtualNode
	imports.NewTable("VirtualNode_Dummy", 0),                    // PVirtualNode
	imports.NewTable("VirtualNode_TotalCount", 0),               // PVirtualNode
	imports.NewTable("VirtualNode_TotalHeight", 0),              // PVirtualNode
	imports.NewTable("VirtualNode_Parent", 0),                   // PVirtualNode
	imports.NewTable("VirtualNode_PrevSibling", 0),              // PVirtualNode
	imports.NewTable("VirtualNode_NextSibling", 0),              // PVirtualNode
	imports.NewTable("VirtualNode_FirstChild", 0),               // PVirtualNode
	imports.NewTable("VirtualNode_LastChild", 0),                // PVirtualNode
	imports.NewTable("VirtualNode_Data", 0),                     // PVirtualNode
}

// InitPreDefsImport 初始化预定义api
func InitPreDefsImport(imp *imports.Imports) {
	imp.SetImportTable(preImportDefs)
	imp.SetOk(true)
}

const (
	Application_Instance = iota
	Application_CreateForm
	Application_Run
	Application_Initialize
	Application_SetRunLoopReceived
	Form_InheritedWndProc
	Form_Create2
	Form_SetOnWndProc
	Form_SetGoPtr
	SetEventCallback
	SetMessageCallback
	SetThreadSyncCallback
	SetRequestCallCreateParamsCallback
	SetRequestCallFormCreateCallback
	SetRemoveEventCallback
	SetThreadAsyncCallback
	DRunMainAsyncCall
	DGETSTRINGARROF
	DSTRLEN
	DMOVE
	DSHOWMESSAGE
	DMESSAGEDLG
	MOUSE_INSTANCE
	SCREEN_INSTANCE
	DRunMainSyncCall
	DTEXTTOSHORTCUT
	DSHORTCUTTOTEXT
	CLIPBOARD_INSTANCE
	DSETCLIPBOARD
	DREGISTERCLIPBOARDFORMAT
	DPREDEFINEDCLIPBOARDFORMAT
	DSYSOPEN
	DEXTRACTFILEPATH
	DFILEEXISTS
	DSELECTDIRECTORY1
	DSELECTDIRECTORY2
	DINPUTBOX
	DINPUTQUERY
	DPASSWORDBOX
	DINPUTCOMBO
	DINPUTCOMBOEX
	DSYSLOCALE
	DCREATEURLSHORTCUT
	DCREATESHORTCUT
	DSETPROPERTYVALUE
	DSETPROPERTYSECVALUE
	PRINTER_INSTANCE
	PRINTER_SETPRINTER
	DGUIDTOSTRING
	DSTRINGTOGUID
	DCREATEGUID
	DGETLIBRESOURCECOUNT
	DGETLIBRESOURCEITEM
	DMODIFYLIBRESOURCE
	DLIBABOUT
	DLIBSTRINGENCODING
	GetLCLVersion
	DMAINTHREADID
	DCURRENTTHREADID
	DINITGODLL
	DFINDCONTROL
	DFINDLCLCONTROL
	DFINDOWNERCONTROL
	DFINDCONTROLATPOSITION
	DFINDLCLWINDOW
	DFINDDRAGTARGET
	ResFormLoadFromStream
	ResFormLoadFromFile
	ResFormLoadFromResourceName
	NSWINDOW_FROMFORM
	NSWINDOW_TITLEVISIBILITY
	NSWINDOW_SETTITLEVISIBILITY
	NSWINDOW_TITLEBARAPPEARSTRANSPARENT
	NSWINDOW_SETTITLEBARAPPEARSTRANSPARENT
	NSWINDOW_STYLEMASK
	NSWINDOW_SETSTYLEMASK
	NSWINDOW_SETREPRESENTEDURL
	GDKWINDOW_GETXID
	GDKWINDOW_FROMFORM
	GTKWIDGET_GETGTKFIXED
	GTKWIDGET_WINDOW
	DSENDMESSAGE
	DPOSTMESSAGE
	DISICONIC
	DISWINDOW
	DISZOOMED
	DISWINDOWVISIBLE
	DGETDC
	DRELEASEDC
	DSETFOREGROUNDWINDOW
	DWINDOWFROMPOINT
	SetExceptionHandlerCallback
	COMSequentialStream_Read
	COMSequentialStream_Write
	COMStream_Seek
	COMStream_SetSize
	COMStream_CopyTo
	COMStream_Commit
	COMStream_Revert
	COMStream_LockRegion
	COMStream_UnlockRegion
	COMStream_Stat
	COMStream_Clone
	StreamAdapter_Clone
	StreamAdapter_Commit
	StreamAdapter_CopyTo
	StreamAdapter_Create
	StreamAdapter_LockRegion
	StreamAdapter_Read
	StreamAdapter_Revert
	StreamAdapter_Seek
	StreamAdapter_SetSize
	StreamAdapter_Stat
	StreamAdapter_Stream
	StreamAdapter_StreamOwnership
	StreamAdapter_UnlockRegion
	StreamAdapter_Write
	InterfacedObject_Create
	InterfacedObject_RefCount
	DFreeAndNil
	DSetNil
	DToUnixTime
	DUnixToTime
	SetCEFEventCallback
	SetCEFRemoveEventCallback
	SetWVEventCallback
	SetWVRemoveEventCallback
	SetWKEventCallback
	SetWKRemoveEventCallback
	VirtualTreeColumns_GetVisibleColumns
	VTImageInfo_Get
	VirtualNode_Create
	VirtualNode_Index
	VirtualNode_ChildCount
	VirtualNode_NodeHeight
	VirtualNode_States
	VirtualNode_Align
	VirtualNode_CheckState
	VirtualNode_CheckType
	VirtualNode_Dummy
	VirtualNode_TotalCount
	VirtualNode_TotalHeight
	VirtualNode_Parent
	VirtualNode_PrevSibling
	VirtualNode_NextSibling
	VirtualNode_FirstChild
	VirtualNode_LastChild
	VirtualNode_Data
)
