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

// IFloatSpinEdit Parent: ICustomFloatSpinEdit
type IFloatSpinEdit interface {
	ICustomFloatSpinEdit
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

// TFloatSpinEdit Parent: TCustomFloatSpinEdit
type TFloatSpinEdit struct {
	TCustomFloatSpinEdit
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

func NewFloatSpinEdit(TheOwner IComponent) IFloatSpinEdit {
	r1 := floatSpinEditImportAPI().SysCallN(3, GetObjectUintptr(TheOwner))
	return AsFloatSpinEdit(r1)
}

func (m *TFloatSpinEdit) AutoSelected() bool {
	r1 := floatSpinEditImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFloatSpinEdit) SetAutoSelected(AValue bool) {
	floatSpinEditImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFloatSpinEdit) AutoSelect() bool {
	r1 := floatSpinEditImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFloatSpinEdit) SetAutoSelect(AValue bool) {
	floatSpinEditImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFloatSpinEdit) ParentColor() bool {
	r1 := floatSpinEditImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFloatSpinEdit) SetParentColor(AValue bool) {
	floatSpinEditImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFloatSpinEdit) ParentFont() bool {
	r1 := floatSpinEditImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFloatSpinEdit) SetParentFont(AValue bool) {
	floatSpinEditImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFloatSpinEdit) ParentShowHint() bool {
	r1 := floatSpinEditImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFloatSpinEdit) SetParentShowHint(AValue bool) {
	floatSpinEditImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func FloatSpinEditClass() TClass {
	ret := floatSpinEditImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TFloatSpinEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(7, m.Instance(), m.editingDonePtr)
}

func (m *TFloatSpinEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(8, m.Instance(), m.mouseDownPtr)
}

func (m *TFloatSpinEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(9, m.Instance(), m.mouseEnterPtr)
}

func (m *TFloatSpinEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(10, m.Instance(), m.mouseLeavePtr)
}

func (m *TFloatSpinEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(11, m.Instance(), m.mouseMovePtr)
}

func (m *TFloatSpinEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(12, m.Instance(), m.mouseUpPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(13, m.Instance(), m.mouseWheelPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(14, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(18, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(15, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(16, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TFloatSpinEdit) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	floatSpinEditImportAPI().SysCallN(17, m.Instance(), m.mouseWheelRightPtr)
}

var (
	floatSpinEditImport       *imports.Imports = nil
	floatSpinEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FloatSpinEdit_AutoSelect", 0),
		/*1*/ imports.NewTable("FloatSpinEdit_AutoSelected", 0),
		/*2*/ imports.NewTable("FloatSpinEdit_Class", 0),
		/*3*/ imports.NewTable("FloatSpinEdit_Create", 0),
		/*4*/ imports.NewTable("FloatSpinEdit_ParentColor", 0),
		/*5*/ imports.NewTable("FloatSpinEdit_ParentFont", 0),
		/*6*/ imports.NewTable("FloatSpinEdit_ParentShowHint", 0),
		/*7*/ imports.NewTable("FloatSpinEdit_SetOnEditingDone", 0),
		/*8*/ imports.NewTable("FloatSpinEdit_SetOnMouseDown", 0),
		/*9*/ imports.NewTable("FloatSpinEdit_SetOnMouseEnter", 0),
		/*10*/ imports.NewTable("FloatSpinEdit_SetOnMouseLeave", 0),
		/*11*/ imports.NewTable("FloatSpinEdit_SetOnMouseMove", 0),
		/*12*/ imports.NewTable("FloatSpinEdit_SetOnMouseUp", 0),
		/*13*/ imports.NewTable("FloatSpinEdit_SetOnMouseWheel", 0),
		/*14*/ imports.NewTable("FloatSpinEdit_SetOnMouseWheelDown", 0),
		/*15*/ imports.NewTable("FloatSpinEdit_SetOnMouseWheelHorz", 0),
		/*16*/ imports.NewTable("FloatSpinEdit_SetOnMouseWheelLeft", 0),
		/*17*/ imports.NewTable("FloatSpinEdit_SetOnMouseWheelRight", 0),
		/*18*/ imports.NewTable("FloatSpinEdit_SetOnMouseWheelUp", 0),
	}
)

func floatSpinEditImportAPI() *imports.Imports {
	if floatSpinEditImport == nil {
		floatSpinEditImport = NewDefaultImports()
		floatSpinEditImport.SetImportTable(floatSpinEditImportTables)
		floatSpinEditImportTables = nil
	}
	return floatSpinEditImport
}
