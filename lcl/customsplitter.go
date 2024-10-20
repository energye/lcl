//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// ICustomSplitter Parent: ICustomControl
type ICustomSplitter interface {
	ICustomControl
	ResizeControl() IControl                            // property
	SetResizeControl(AValue IControl)                   // property
	AutoSnap() bool                                     // property
	SetAutoSnap(AValue bool)                            // property
	Beveled() bool                                      // property
	SetBeveled(AValue bool)                             // property
	MinSize() int32                                     // property
	SetMinSize(AValue int32)                            // property
	ResizeAnchor() TAnchorKind                          // property
	SetResizeAnchor(AValue TAnchorKind)                 // property
	ResizeStyle() TResizeStyle                          // property
	SetResizeStyle(AValue TResizeStyle)                 // property
	GetOtherResizeControl() IControl                    // function
	GetSplitterPosition() int32                         // function
	AnchorSplitter(Kind TAnchorKind, AControl IControl) // procedure
	MoveSplitter(Offset int32)                          // procedure
	SetSplitterPosition(NewPosition int32)              // procedure
	SetOnCanOffset(fn TCanOffsetEvent)                  // property event
	SetOnCanResize(fn TCanResizeEvent)                  // property event
	SetOnMoved(fn TNotifyEvent)                         // property event
}

// TCustomSplitter Parent: TCustomControl
type TCustomSplitter struct {
	TCustomControl
	canOffsetPtr uintptr
	canResizePtr uintptr
	movedPtr     uintptr
}

func NewCustomSplitter(TheOwner IComponent) ICustomSplitter {
	r1 := customSplitterImportAPI().SysCallN(4, GetObjectUintptr(TheOwner))
	return AsCustomSplitter(r1)
}

func (m *TCustomSplitter) ResizeControl() IControl {
	r1 := customSplitterImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return AsControl(r1)
}

func (m *TCustomSplitter) SetResizeControl(AValue IControl) {
	customSplitterImportAPI().SysCallN(10, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomSplitter) AutoSnap() bool {
	r1 := customSplitterImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSplitter) SetAutoSnap(AValue bool) {
	customSplitterImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSplitter) Beveled() bool {
	r1 := customSplitterImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomSplitter) SetBeveled(AValue bool) {
	customSplitterImportAPI().SysCallN(2, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomSplitter) MinSize() int32 {
	r1 := customSplitterImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomSplitter) SetMinSize(AValue int32) {
	customSplitterImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) ResizeAnchor() TAnchorKind {
	r1 := customSplitterImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return TAnchorKind(r1)
}

func (m *TCustomSplitter) SetResizeAnchor(AValue TAnchorKind) {
	customSplitterImportAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) ResizeStyle() TResizeStyle {
	r1 := customSplitterImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return TResizeStyle(r1)
}

func (m *TCustomSplitter) SetResizeStyle(AValue TResizeStyle) {
	customSplitterImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomSplitter) GetOtherResizeControl() IControl {
	r1 := customSplitterImportAPI().SysCallN(5, m.Instance())
	return AsControl(r1)
}

func (m *TCustomSplitter) GetSplitterPosition() int32 {
	r1 := customSplitterImportAPI().SysCallN(6, m.Instance())
	return int32(r1)
}

func CustomSplitterClass() TClass {
	ret := customSplitterImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TCustomSplitter) AnchorSplitter(Kind TAnchorKind, AControl IControl) {
	customSplitterImportAPI().SysCallN(0, m.Instance(), uintptr(Kind), GetObjectUintptr(AControl))
}

func (m *TCustomSplitter) MoveSplitter(Offset int32) {
	customSplitterImportAPI().SysCallN(8, m.Instance(), uintptr(Offset))
}

func (m *TCustomSplitter) SetSplitterPosition(NewPosition int32) {
	customSplitterImportAPI().SysCallN(15, m.Instance(), uintptr(NewPosition))
}

func (m *TCustomSplitter) SetOnCanOffset(fn TCanOffsetEvent) {
	if m.canOffsetPtr != 0 {
		RemoveEventElement(m.canOffsetPtr)
	}
	m.canOffsetPtr = MakeEventDataPtr(fn)
	customSplitterImportAPI().SysCallN(12, m.Instance(), m.canOffsetPtr)
}

func (m *TCustomSplitter) SetOnCanResize(fn TCanResizeEvent) {
	if m.canResizePtr != 0 {
		RemoveEventElement(m.canResizePtr)
	}
	m.canResizePtr = MakeEventDataPtr(fn)
	customSplitterImportAPI().SysCallN(13, m.Instance(), m.canResizePtr)
}

func (m *TCustomSplitter) SetOnMoved(fn TNotifyEvent) {
	if m.movedPtr != 0 {
		RemoveEventElement(m.movedPtr)
	}
	m.movedPtr = MakeEventDataPtr(fn)
	customSplitterImportAPI().SysCallN(14, m.Instance(), m.movedPtr)
}

var (
	customSplitterImport       *imports.Imports = nil
	customSplitterImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomSplitter_AnchorSplitter", 0),
		/*1*/ imports.NewTable("CustomSplitter_AutoSnap", 0),
		/*2*/ imports.NewTable("CustomSplitter_Beveled", 0),
		/*3*/ imports.NewTable("CustomSplitter_Class", 0),
		/*4*/ imports.NewTable("CustomSplitter_Create", 0),
		/*5*/ imports.NewTable("CustomSplitter_GetOtherResizeControl", 0),
		/*6*/ imports.NewTable("CustomSplitter_GetSplitterPosition", 0),
		/*7*/ imports.NewTable("CustomSplitter_MinSize", 0),
		/*8*/ imports.NewTable("CustomSplitter_MoveSplitter", 0),
		/*9*/ imports.NewTable("CustomSplitter_ResizeAnchor", 0),
		/*10*/ imports.NewTable("CustomSplitter_ResizeControl", 0),
		/*11*/ imports.NewTable("CustomSplitter_ResizeStyle", 0),
		/*12*/ imports.NewTable("CustomSplitter_SetOnCanOffset", 0),
		/*13*/ imports.NewTable("CustomSplitter_SetOnCanResize", 0),
		/*14*/ imports.NewTable("CustomSplitter_SetOnMoved", 0),
		/*15*/ imports.NewTable("CustomSplitter_SetSplitterPosition", 0),
	}
)

func customSplitterImportAPI() *imports.Imports {
	if customSplitterImport == nil {
		customSplitterImport = NewDefaultImports()
		customSplitterImport.SetImportTable(customSplitterImportTables)
		customSplitterImportTables = nil
	}
	return customSplitterImport
}
