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

// ICheckBox Parent: ICustomCheckBox
type ICheckBox interface {
	ICustomCheckBox
	Checked() bool                                 // property
	SetChecked(AValue bool)                        // property
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragKind() TDragKind                           // property
	SetDragKind(AValue TDragKind)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TCheckBox Parent: TCustomCheckBox
type TCheckBox struct {
	TCustomCheckBox
	contextPopupPtr   uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	editingDonePtr    uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDragPtr      uintptr
}

func NewCheckBox(TheOwner IComponent) ICheckBox {
	r1 := checkBoxImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsCheckBox(r1)
}

func (m *TCheckBox) Checked() bool {
	r1 := checkBoxImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckBox) SetChecked(AValue bool) {
	checkBoxImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckBox) DragCursor() TCursor {
	r1 := checkBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCheckBox) SetDragCursor(AValue TCursor) {
	checkBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckBox) DragKind() TDragKind {
	r1 := checkBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TCheckBox) SetDragKind(AValue TDragKind) {
	checkBoxImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckBox) DragMode() TDragMode {
	r1 := checkBoxImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCheckBox) SetDragMode(AValue TDragMode) {
	checkBoxImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckBox) ParentColor() bool {
	r1 := checkBoxImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckBox) SetParentColor(AValue bool) {
	checkBoxImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckBox) ParentFont() bool {
	r1 := checkBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckBox) SetParentFont(AValue bool) {
	checkBoxImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckBox) ParentShowHint() bool {
	r1 := checkBoxImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckBox) SetParentShowHint(AValue bool) {
	checkBoxImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func CheckBoxClass() TClass {
	ret := checkBoxImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TCheckBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(9, m.Instance(), m.contextPopupPtr)
}

func (m *TCheckBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(10, m.Instance(), m.dragDropPtr)
}

func (m *TCheckBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(11, m.Instance(), m.dragOverPtr)
}

func (m *TCheckBox) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(12, m.Instance(), m.editingDonePtr)
}

func (m *TCheckBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(13, m.Instance(), m.endDragPtr)
}

func (m *TCheckBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(14, m.Instance(), m.mouseDownPtr)
}

func (m *TCheckBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(15, m.Instance(), m.mouseEnterPtr)
}

func (m *TCheckBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(16, m.Instance(), m.mouseLeavePtr)
}

func (m *TCheckBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(17, m.Instance(), m.mouseMovePtr)
}

func (m *TCheckBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(18, m.Instance(), m.mouseUpPtr)
}

func (m *TCheckBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(19, m.Instance(), m.mouseWheelPtr)
}

func (m *TCheckBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(20, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCheckBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(21, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCheckBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	checkBoxImportAPI().SysCallN(22, m.Instance(), m.startDragPtr)
}

var (
	checkBoxImport       *imports.Imports = nil
	checkBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CheckBox_Checked", 0),
		/*1*/ imports.NewTable("CheckBox_Class", 0),
		/*2*/ imports.NewTable("CheckBox_Create", 0),
		/*3*/ imports.NewTable("CheckBox_DragCursor", 0),
		/*4*/ imports.NewTable("CheckBox_DragKind", 0),
		/*5*/ imports.NewTable("CheckBox_DragMode", 0),
		/*6*/ imports.NewTable("CheckBox_ParentColor", 0),
		/*7*/ imports.NewTable("CheckBox_ParentFont", 0),
		/*8*/ imports.NewTable("CheckBox_ParentShowHint", 0),
		/*9*/ imports.NewTable("CheckBox_SetOnContextPopup", 0),
		/*10*/ imports.NewTable("CheckBox_SetOnDragDrop", 0),
		/*11*/ imports.NewTable("CheckBox_SetOnDragOver", 0),
		/*12*/ imports.NewTable("CheckBox_SetOnEditingDone", 0),
		/*13*/ imports.NewTable("CheckBox_SetOnEndDrag", 0),
		/*14*/ imports.NewTable("CheckBox_SetOnMouseDown", 0),
		/*15*/ imports.NewTable("CheckBox_SetOnMouseEnter", 0),
		/*16*/ imports.NewTable("CheckBox_SetOnMouseLeave", 0),
		/*17*/ imports.NewTable("CheckBox_SetOnMouseMove", 0),
		/*18*/ imports.NewTable("CheckBox_SetOnMouseUp", 0),
		/*19*/ imports.NewTable("CheckBox_SetOnMouseWheel", 0),
		/*20*/ imports.NewTable("CheckBox_SetOnMouseWheelDown", 0),
		/*21*/ imports.NewTable("CheckBox_SetOnMouseWheelUp", 0),
		/*22*/ imports.NewTable("CheckBox_SetOnStartDrag", 0),
	}
)

func checkBoxImportAPI() *imports.Imports {
	if checkBoxImport == nil {
		checkBoxImport = NewDefaultImports()
		checkBoxImport.SetImportTable(checkBoxImportTables)
		checkBoxImportTables = nil
	}
	return checkBoxImport
}
