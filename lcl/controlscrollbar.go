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

// IControlScrollBar Parent: IPersistent
type IControlScrollBar interface {
	IPersistent
	Kind() TScrollBarKind                 // property
	Size() int32                          // property
	Increment() TScrollBarInc             // property
	SetIncrement(AValue TScrollBarInc)    // property
	Page() TScrollBarInc                  // property
	SetPage(AValue TScrollBarInc)         // property
	Smooth() bool                         // property
	SetSmooth(AValue bool)                // property
	Position() int32                      // property
	SetPosition(AValue int32)             // property
	Range() int32                         // property
	SetRange(AValue int32)                // property
	Tracking() bool                       // property
	SetTracking(AValue bool)              // property
	Visible() bool                        // property
	SetVisible(AValue bool)               // property
	IsScrollBarVisible() bool             // function
	ScrollPos() int32                     // function
	GetOtherScrollBar() IControlScrollBar // function
	ControlSize() int32                   // function
	ClientSize() int32                    // function
	ClientSizeWithBar() int32             // function
	ClientSizeWithoutBar() int32          // function
}

// TControlScrollBar Parent: TPersistent
type TControlScrollBar struct {
	TPersistent
}

func NewControlScrollBar(AControl IWinControl, AKind TScrollBarKind) IControlScrollBar {
	r1 := controlScrollBarImportAPI().SysCallN(5, GetObjectUintptr(AControl), uintptr(AKind))
	return AsControlScrollBar(r1)
}

func (m *TControlScrollBar) Kind() TScrollBarKind {
	r1 := controlScrollBarImportAPI().SysCallN(9, m.Instance())
	return TScrollBarKind(r1)
}

func (m *TControlScrollBar) Size() int32 {
	r1 := controlScrollBarImportAPI().SysCallN(14, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) Increment() TScrollBarInc {
	r1 := controlScrollBarImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TScrollBarInc(r1)
}

func (m *TControlScrollBar) SetIncrement(AValue TScrollBarInc) {
	controlScrollBarImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Page() TScrollBarInc {
	r1 := controlScrollBarImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return TScrollBarInc(r1)
}

func (m *TControlScrollBar) SetPage(AValue TScrollBarInc) {
	controlScrollBarImportAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Smooth() bool {
	r1 := controlScrollBarImportAPI().SysCallN(15, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControlScrollBar) SetSmooth(AValue bool) {
	controlScrollBarImportAPI().SysCallN(15, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControlScrollBar) Position() int32 {
	r1 := controlScrollBarImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlScrollBar) SetPosition(AValue int32) {
	controlScrollBarImportAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Range() int32 {
	r1 := controlScrollBarImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TControlScrollBar) SetRange(AValue int32) {
	controlScrollBarImportAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TControlScrollBar) Tracking() bool {
	r1 := controlScrollBarImportAPI().SysCallN(16, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControlScrollBar) SetTracking(AValue bool) {
	controlScrollBarImportAPI().SysCallN(16, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControlScrollBar) Visible() bool {
	r1 := controlScrollBarImportAPI().SysCallN(17, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TControlScrollBar) SetVisible(AValue bool) {
	controlScrollBarImportAPI().SysCallN(17, 1, m.Instance(), PascalBool(AValue))
}

func (m *TControlScrollBar) IsScrollBarVisible() bool {
	r1 := controlScrollBarImportAPI().SysCallN(8, m.Instance())
	return GoBool(r1)
}

func (m *TControlScrollBar) ScrollPos() int32 {
	r1 := controlScrollBarImportAPI().SysCallN(13, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) GetOtherScrollBar() IControlScrollBar {
	r1 := controlScrollBarImportAPI().SysCallN(6, m.Instance())
	return AsControlScrollBar(r1)
}

func (m *TControlScrollBar) ControlSize() int32 {
	r1 := controlScrollBarImportAPI().SysCallN(4, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) ClientSize() int32 {
	r1 := controlScrollBarImportAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) ClientSizeWithBar() int32 {
	r1 := controlScrollBarImportAPI().SysCallN(2, m.Instance())
	return int32(r1)
}

func (m *TControlScrollBar) ClientSizeWithoutBar() int32 {
	r1 := controlScrollBarImportAPI().SysCallN(3, m.Instance())
	return int32(r1)
}

func ControlScrollBarClass() TClass {
	ret := controlScrollBarImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	controlScrollBarImport       *imports.Imports = nil
	controlScrollBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ControlScrollBar_Class", 0),
		/*1*/ imports.NewTable("ControlScrollBar_ClientSize", 0),
		/*2*/ imports.NewTable("ControlScrollBar_ClientSizeWithBar", 0),
		/*3*/ imports.NewTable("ControlScrollBar_ClientSizeWithoutBar", 0),
		/*4*/ imports.NewTable("ControlScrollBar_ControlSize", 0),
		/*5*/ imports.NewTable("ControlScrollBar_Create", 0),
		/*6*/ imports.NewTable("ControlScrollBar_GetOtherScrollBar", 0),
		/*7*/ imports.NewTable("ControlScrollBar_Increment", 0),
		/*8*/ imports.NewTable("ControlScrollBar_IsScrollBarVisible", 0),
		/*9*/ imports.NewTable("ControlScrollBar_Kind", 0),
		/*10*/ imports.NewTable("ControlScrollBar_Page", 0),
		/*11*/ imports.NewTable("ControlScrollBar_Position", 0),
		/*12*/ imports.NewTable("ControlScrollBar_Range", 0),
		/*13*/ imports.NewTable("ControlScrollBar_ScrollPos", 0),
		/*14*/ imports.NewTable("ControlScrollBar_Size", 0),
		/*15*/ imports.NewTable("ControlScrollBar_Smooth", 0),
		/*16*/ imports.NewTable("ControlScrollBar_Tracking", 0),
		/*17*/ imports.NewTable("ControlScrollBar_Visible", 0),
	}
)

func controlScrollBarImportAPI() *imports.Imports {
	if controlScrollBarImport == nil {
		controlScrollBarImport = NewDefaultImports()
		controlScrollBarImport.SetImportTable(controlScrollBarImportTables)
		controlScrollBarImportTables = nil
	}
	return controlScrollBarImport
}
