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

import "github.com/energye/lcl/types"

type TBlendFunction struct {
	BlendOp             uint8
	BlendFlags          uint8
	SourceConstantAlpha uint8
	AlphaFormat         uint8
}

type TSystemInfo struct {

	//0: (
	//dwOemId: DWORD);
	//1: (
	ProcessorArchitecture uint16
	Reserved              uint16

	PageSize                  uint32
	MinimumApplicationAddress uintptr
	MaximumApplicationAddress uintptr
	ActiveProcessorMask       uintptr
	NumberOfProcessors        uint32
	ProcessorType             uint32
	AllocationGranularity     uint32
	ProcessorLevel            uint16
	ProcessorRevision         uint16
}

type TSecurityAttributes struct {
	nLength              uint32
	lpSecurityDescriptor uintptr
	bInheritHandle       bool // BOOL
}

type TSHItemID struct {
	Cb   uint16  // Size of the ID (including cb itself)
	AbID [1]byte // The item ID (variable length)
}

type TItemIDList struct {
	Mkid TSHItemID
}

// TShellExecuteInfo ShellExecuteEx
type TShellExecuteInfo struct {
	CbSize       uint32
	FMask        uint32
	Wnd          types.HWND
	LpVerb       types.LPCWSTR
	LpFile       types.LPCWSTR
	LpParameters types.LPCWSTR
	LpDirectory  types.LPCWSTR
	NShow        int32
	HInstApp     types.HINST
	// Optional fields
	LpIDList  uintptr
	LpClass   types.LPCWSTR
	HkeyClass types.HKEY
	DwHotKey  uint32

	//	0: (
	// HICON
	//	1: (
	HIcon_hMonitor uintptr

	HProcess uintptr
}

type TMemoryBasicInformation struct {
	BaseAddress       uintptr
	AllocationBase    uintptr
	AllocationProtect uint32
	RegionSize        types.SIZE_T
	State             uint32
	Protect           uint32
	Type_9            uint32
}

type TSHFileInfo struct {
	HIcon         types.HICON      /* out: icon */
	IIcon         types.Integer    /* out: icon index */
	DwAttributes  types.DWORD      /* out: SFGAO_ flags */
	SzDisplayName [MAX_PATH]uint16 /* out: display name (or path) */
	SzTypeName    [80]uint16       /* out: type name */
}

type WINDOWPLACEMENT struct {
	Length           uint32
	Flags            uint32
	ShowCmd          uint32
	PtMinPosition    types.TPoint
	PtMaxPosition    types.TPoint
	RcNormalPosition types.TRect
}

type WINDOWPOS struct {
	_hwnd           types.HWND
	hwndInsertAfter types.HWND
	x               types.LongInt
	y               types.LongInt
	cx              types.LongInt
	cy              types.LongInt
	flags           types.UINT
}
type NCCALCSIZE_PARAMS struct {
	Rgrc  [3]types.TRect // 三个矩形区域
	Lppos *WINDOWPOS     // 指向WINDOWPOS结构的指针
}
