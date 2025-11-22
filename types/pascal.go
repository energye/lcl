//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

import (
	"fmt"
	"unsafe"
)

type TGridCoord struct {
	X int32
	Y int32
}

// PRawImageDescription > ^TRawImageDescription = object > TRawImageDescriptionWrap
type PRawImageDescription = uintptr

// PRawImage > ^TRawImage = object > TRawImageWrap
type PRawImage = uintptr

// TCustomData Pointer
type TCustomData = uintptr

// TGridRect TRect
type TGridRect = TRect

// PUint64Array = ^TUint64Array = array of uint64
type PUint64Array uintptr

// PInt64Array = ^TInt64Array = array of uint64
type PInt64Array uintptr

// PUint32Array = ^TUint32Array = array of uint32
type PUint32Array uintptr

// PStringArray = ^TStringArray = array of string
type PStringArray uintptr

// PBrushPattern = ^TBrushPattern = array[0..PatternBitCount-1] of TPenPattern;
// PatternBitCount = SizeOf(uint32) * 8;
type PBrushPattern uintptr

// TObjectDynArray array of TObject;
type TObjectDynArray uintptr

// TColumnsArray array of TVirtualTreeColumn
// TVirtualTreeColumn = class
type TColumnsArray uintptr

// TNodeArray = array of PVirtualNode = array of TVirtualNodeWrap
type TNodeArray uintptr

// PVTVirtualNodeEnumeration = ^TVTVirtualNodeEnumeration > TVTVirtualNodeEnumerationWrap
type PVTVirtualNodeEnumeration uintptr

// PScaledImageListResolution = ^TScaledImageListResolution > TScaledImageListResolutionWrap
type PScaledImageListResolution uintptr

// PPointArray = ^TPointArray
type PPointArray uintptr

// PPoint = ^TPoint = []TPoint pointer
//
//	Use: types.PPoint(unsafe.Pointer(&points[0]))
type PPoint uintptr

// PFormatEtcArray = ^TFormatEtcArray
type PFormatEtcArray uintptr

// PWVCustomSchemeInfoArray = ^TWVCustomSchemeInfoArray
type PWVCustomSchemeInfoArray uintptr

// PFormatEtc = ^TFormatEtc
type PFormatEtc uintptr

// PPointerList = ^TPointerList = array[0..MaxListSize - 1] of Pointer
// MaxListSize = Maxint / 16
type PPointerList uintptr

type TSysLocale struct {
	//Delphi compat fields
	DefaultLCID int32
	PriLangID   int32
	SubLangID   int32
	// win32 names
	FarEast    bool
	MiddleEast bool
	// LCL
	// real meaning  2: (MBCS: boolean; RightToLeft: Boolean);
}

// PVirtualNode = ^TVirtualNode = TVirtualNode
// 2: = ^TVirtualNode = TVirtualNodeWrap
type PVirtualNode uintptr

func (m PVirtualNode) ToGo() *TVirtualNode {
	if m != 0 {
		return (*TVirtualNode)(unsafe.Pointer(m))
	}
	return nil
}

type States uint16

func (m States) Value() TSet {
	return TSet(m)
}

type CheckState byte

func (m CheckState) Value() TCheckState {
	return TCheckState(m)
}

type CheckType byte

func (m CheckType) Value() TCheckType {
	return TCheckState(m)
}

type TVirtualNode struct {
	Index       Cardinal   // index of node with regard to its parent
	ChildCount  Cardinal   // number of child nodes
	NodeHeight  uint16     // height in pixels
	States      States     // states describing various properties of the node (expanded, initialized etc.)
	Align       byte       // line/button alignment
	CheckState  CheckState // indicates the current check state (e.g. checked, pressed etc.)
	CheckType   CheckType  // indicates which check type shall be used for this node
	Dummy       byte       // dummy value to fill DWORD boundary
	TotalCount  Cardinal   // sum of this node, all of its child nodes and their child nodes etc.
	TotalHeight Cardinal   // height in pixels this node covers on screen including the height of all of its
	// children
	// Note: Some copy routines require that all pointers (as well as the data area) in a node are
	//       located at the end of the node! Hence if you want to add new member fields (except pointers to internal
	//       data) then put them before field Parent.
	Parent      PVirtualNode // reference to the node's parent (for the root this contains the treeview)
	PrevSibling PVirtualNode // link to the node's previous sibling or nil if it is the first node
	NextSibling PVirtualNode // link to the node's next sibling or nil if it is the last node
	FirstChild  PVirtualNode // link to the node's first child...
	LastChild   PVirtualNode // link to the node's last child...
}

type TGUID struct {
	D1 uint32
	D2 uint16
	D3 uint16
	D4 [8]uint8
}

// TLibResource
type TLibResource struct {
	Name string
	Ptr  uintptr
}

type TResItem struct {
	Name  uintptr
	Value uintptr
}

// TConstraintSize = 0..MaxInt;
type TConstraintSize = int32

// TBorderWidth = 0..MaxInt;
type TBorderWidth = int32

type TAlignInfo struct {
	AlignList    uintptr
	ControlIndex int32
	Align        TAlign
	Scratch      int32
}

// TCreateParams
//
// Moved from Controls to avoid circles
// Since it is part of the interface now
type TCreateParams struct {
	Caption      LPCWSTR
	Style        uint32
	ExStyle      uint32
	X            int32
	Y            int32
	Width        int32
	Height       int32
	WndParent    HWND
	Param        uintptr
	WindowClass  TWndClass
	WinClassName [64]int8
}

// TGUID

func (g TGUID) FromString(str string) (result TGUID) {
	fmt.Sscanf(str, "{%8X-%4X-%4X-%2X%2X-%2X%2X%2X%2X%2X%2X}", &result.D1, &result.D2, &result.D3, &result.D4[0],
		&result.D4[1], &result.D4[2], &result.D4[3], &result.D4[4], &result.D4[5], &result.D4[6], &result.D4[7])
	return
}

func (g TGUID) ToString() string {
	return fmt.Sprintf("{%.8X-%.4X-%.4X-%.2X%.2X-%.2X%.2X%.2X%.2X%.2X%.2X}",
		g.D1, g.D2, g.D3, g.D4[0], g.D4[1], g.D4[2], g.D4[3], g.D4[4], g.D4[5], g.D4[6], g.D4[7])
}

func (g TGUID) Empty() TGUID {
	return TGUID{}
}

func (g TGUID) IsEqual(val TGUID) bool {
	return (g.D1 == val.D1) && (g.D2 == val.D2) && (g.D3 == val.D3) && (g.D4 == val.D4)
}
