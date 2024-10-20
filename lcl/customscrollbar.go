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

// ICustomScrollBar Parent: IWinControl
type ICustomScrollBar interface {
	IWinControl
	Kind() TScrollBarKind                             // property
	SetKind(AValue TScrollBarKind)                    // property
	LargeChange() TScrollBarInc                       // property
	SetLargeChange(AValue TScrollBarInc)              // property
	Max() int32                                       // property
	SetMax(AValue int32)                              // property
	Min() int32                                       // property
	SetMin(AValue int32)                              // property
	PageSize() int32                                  // property
	SetPageSize(AValue int32)                         // property
	Position() int32                                  // property
	SetPosition(AValue int32)                         // property
	SmallChange() TScrollBarInc                       // property
	SetSmallChange(AValue TScrollBarInc)              // property
	SetParams(APosition, AMin, AMax, APageSize int32) // procedure
	SetParams1(APosition, AMin, AMax int32)           // procedure
	SetOnChange(fn TNotifyEvent)                      // property event
	SetOnScroll(fn TScrollEvent)                      // property event
}

// TCustomScrollBar Parent: TWinControl
type TCustomScrollBar struct {
	TWinControl
	changePtr uintptr
	scrollPtr uintptr
}

func NewCustomScrollBar(AOwner IComponent) ICustomScrollBar {
	r1 := customScrollBarImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCustomScrollBar(r1)
}

func (m *TCustomScrollBar) Kind() TScrollBarKind {
	r1 := customScrollBarImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TScrollBarKind(r1)
}

func (m *TCustomScrollBar) SetKind(AValue TScrollBarKind) {
	customScrollBarImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) LargeChange() TScrollBarInc {
	r1 := customScrollBarImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TScrollBarInc(r1)
}

func (m *TCustomScrollBar) SetLargeChange(AValue TScrollBarInc) {
	customScrollBarImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) Max() int32 {
	r1 := customScrollBarImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomScrollBar) SetMax(AValue int32) {
	customScrollBarImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) Min() int32 {
	r1 := customScrollBarImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomScrollBar) SetMin(AValue int32) {
	customScrollBarImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) PageSize() int32 {
	r1 := customScrollBarImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomScrollBar) SetPageSize(AValue int32) {
	customScrollBarImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) Position() int32 {
	r1 := customScrollBarImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomScrollBar) SetPosition(AValue int32) {
	customScrollBarImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomScrollBar) SmallChange() TScrollBarInc {
	r1 := customScrollBarImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return TScrollBarInc(r1)
}

func (m *TCustomScrollBar) SetSmallChange(AValue TScrollBarInc) {
	customScrollBarImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func CustomScrollBarClass() TClass {
	ret := customScrollBarImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCustomScrollBar) SetParams(APosition, AMin, AMax, APageSize int32) {
	customScrollBarImportAPI().SysCallN(10, m.Instance(), uintptr(APosition), uintptr(AMin), uintptr(AMax), uintptr(APageSize))
}

func (m *TCustomScrollBar) SetParams1(APosition, AMin, AMax int32) {
	customScrollBarImportAPI().SysCallN(11, m.Instance(), uintptr(APosition), uintptr(AMin), uintptr(AMax))
}

func (m *TCustomScrollBar) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	customScrollBarImportAPI().SysCallN(8, m.Instance(), m.changePtr)
}

func (m *TCustomScrollBar) SetOnScroll(fn TScrollEvent) {
	if m.scrollPtr != 0 {
		RemoveEventElement(m.scrollPtr)
	}
	m.scrollPtr = MakeEventDataPtr(fn)
	customScrollBarImportAPI().SysCallN(9, m.Instance(), m.scrollPtr)
}

var (
	customScrollBarImport       *imports.Imports = nil
	customScrollBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomScrollBar_Class", 0),
		/*1*/ imports.NewTable("CustomScrollBar_Create", 0),
		/*2*/ imports.NewTable("CustomScrollBar_Kind", 0),
		/*3*/ imports.NewTable("CustomScrollBar_LargeChange", 0),
		/*4*/ imports.NewTable("CustomScrollBar_Max", 0),
		/*5*/ imports.NewTable("CustomScrollBar_Min", 0),
		/*6*/ imports.NewTable("CustomScrollBar_PageSize", 0),
		/*7*/ imports.NewTable("CustomScrollBar_Position", 0),
		/*8*/ imports.NewTable("CustomScrollBar_SetOnChange", 0),
		/*9*/ imports.NewTable("CustomScrollBar_SetOnScroll", 0),
		/*10*/ imports.NewTable("CustomScrollBar_SetParams", 0),
		/*11*/ imports.NewTable("CustomScrollBar_SetParams1", 0),
		/*12*/ imports.NewTable("CustomScrollBar_SmallChange", 0),
	}
)

func customScrollBarImportAPI() *imports.Imports {
	if customScrollBarImport == nil {
		customScrollBarImport = NewDefaultImports()
		customScrollBarImport.SetImportTable(customScrollBarImportTables)
		customScrollBarImportTables = nil
	}
	return customScrollBarImport
}
