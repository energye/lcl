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

// ISpinEdit Parent: ICustomSpinEdit
type ISpinEdit interface {
	ICustomSpinEdit
	AutoSelected() bool                             // property
	SetAutoSelected(AValue bool)                    // property
	AutoSelect() bool                               // property
	SetAutoSelect(AValue bool)                      // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnEditingDone(fn TNotifyEvent)               // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
}

// TSpinEdit Parent: TCustomSpinEdit
type TSpinEdit struct {
	TCustomSpinEdit
	editingDonePtr     uintptr
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
}

func NewSpinEdit(TheOwner IComponent) ISpinEdit {
	r1 := spinEditImportAPI().SysCallN(3, GetObjectUintptr(TheOwner))
	return AsSpinEdit(r1)
}

func (m *TSpinEdit) AutoSelected() bool {
	r1 := spinEditImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpinEdit) SetAutoSelected(AValue bool) {
	spinEditImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSpinEdit) AutoSelect() bool {
	r1 := spinEditImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpinEdit) SetAutoSelect(AValue bool) {
	spinEditImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSpinEdit) ParentColor() bool {
	r1 := spinEditImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpinEdit) SetParentColor(AValue bool) {
	spinEditImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSpinEdit) ParentFont() bool {
	r1 := spinEditImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpinEdit) SetParentFont(AValue bool) {
	spinEditImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSpinEdit) ParentShowHint() bool {
	r1 := spinEditImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpinEdit) SetParentShowHint(AValue bool) {
	spinEditImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func SpinEditClass() TClass {
	ret := spinEditImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TSpinEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(7, m.Instance(), m.editingDonePtr)
}

func (m *TSpinEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(8, m.Instance(), m.mouseDownPtr)
}

func (m *TSpinEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(9, m.Instance(), m.mouseEnterPtr)
}

func (m *TSpinEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(10, m.Instance(), m.mouseLeavePtr)
}

func (m *TSpinEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(11, m.Instance(), m.mouseMovePtr)
}

func (m *TSpinEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(12, m.Instance(), m.mouseUpPtr)
}

func (m *TSpinEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(13, m.Instance(), m.mouseWheelPtr)
}

func (m *TSpinEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(14, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TSpinEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(18, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TSpinEdit) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(15, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TSpinEdit) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(16, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TSpinEdit) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	spinEditImportAPI().SysCallN(17, m.Instance(), m.mouseWheelRightPtr)
}

var (
	spinEditImport       *imports.Imports = nil
	spinEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("SpinEdit_AutoSelect", 0),
		/*1*/ imports.NewTable("SpinEdit_AutoSelected", 0),
		/*2*/ imports.NewTable("SpinEdit_Class", 0),
		/*3*/ imports.NewTable("SpinEdit_Create", 0),
		/*4*/ imports.NewTable("SpinEdit_ParentColor", 0),
		/*5*/ imports.NewTable("SpinEdit_ParentFont", 0),
		/*6*/ imports.NewTable("SpinEdit_ParentShowHint", 0),
		/*7*/ imports.NewTable("SpinEdit_SetOnEditingDone", 0),
		/*8*/ imports.NewTable("SpinEdit_SetOnMouseDown", 0),
		/*9*/ imports.NewTable("SpinEdit_SetOnMouseEnter", 0),
		/*10*/ imports.NewTable("SpinEdit_SetOnMouseLeave", 0),
		/*11*/ imports.NewTable("SpinEdit_SetOnMouseMove", 0),
		/*12*/ imports.NewTable("SpinEdit_SetOnMouseUp", 0),
		/*13*/ imports.NewTable("SpinEdit_SetOnMouseWheel", 0),
		/*14*/ imports.NewTable("SpinEdit_SetOnMouseWheelDown", 0),
		/*15*/ imports.NewTable("SpinEdit_SetOnMouseWheelHorz", 0),
		/*16*/ imports.NewTable("SpinEdit_SetOnMouseWheelLeft", 0),
		/*17*/ imports.NewTable("SpinEdit_SetOnMouseWheelRight", 0),
		/*18*/ imports.NewTable("SpinEdit_SetOnMouseWheelUp", 0),
	}
)

func spinEditImportAPI() *imports.Imports {
	if spinEditImport == nil {
		spinEditImport = NewDefaultImports()
		spinEditImport.SetImportTable(spinEditImportTables)
		spinEditImportTables = nil
	}
	return spinEditImport
}
