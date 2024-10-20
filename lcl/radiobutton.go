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

// IRadioButton Parent: ICustomCheckBox
type IRadioButton interface {
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

// TRadioButton Parent: TCustomCheckBox
type TRadioButton struct {
	TCustomCheckBox
	contextPopupPtr   uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
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

func NewRadioButton(TheOwner IComponent) IRadioButton {
	r1 := radioButtonImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsRadioButton(r1)
}

func (m *TRadioButton) Checked() bool {
	r1 := radioButtonImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetChecked(AValue bool) {
	radioButtonImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioButton) DragCursor() TCursor {
	r1 := radioButtonImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TRadioButton) SetDragCursor(AValue TCursor) {
	radioButtonImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioButton) DragKind() TDragKind {
	r1 := radioButtonImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TRadioButton) SetDragKind(AValue TDragKind) {
	radioButtonImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioButton) DragMode() TDragMode {
	r1 := radioButtonImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TRadioButton) SetDragMode(AValue TDragMode) {
	radioButtonImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioButton) ParentColor() bool {
	r1 := radioButtonImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetParentColor(AValue bool) {
	radioButtonImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioButton) ParentFont() bool {
	r1 := radioButtonImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetParentFont(AValue bool) {
	radioButtonImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioButton) ParentShowHint() bool {
	r1 := radioButtonImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetParentShowHint(AValue bool) {
	radioButtonImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func RadioButtonClass() TClass {
	ret := radioButtonImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TRadioButton) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(9, m.Instance(), m.contextPopupPtr)
}

func (m *TRadioButton) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(10, m.Instance(), m.dragDropPtr)
}

func (m *TRadioButton) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(11, m.Instance(), m.dragOverPtr)
}

func (m *TRadioButton) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(12, m.Instance(), m.endDragPtr)
}

func (m *TRadioButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(13, m.Instance(), m.mouseDownPtr)
}

func (m *TRadioButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(14, m.Instance(), m.mouseEnterPtr)
}

func (m *TRadioButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(15, m.Instance(), m.mouseLeavePtr)
}

func (m *TRadioButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(16, m.Instance(), m.mouseMovePtr)
}

func (m *TRadioButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(17, m.Instance(), m.mouseUpPtr)
}

func (m *TRadioButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(18, m.Instance(), m.mouseWheelPtr)
}

func (m *TRadioButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(19, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TRadioButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(20, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TRadioButton) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	radioButtonImportAPI().SysCallN(21, m.Instance(), m.startDragPtr)
}

var (
	radioButtonImport       *imports.Imports = nil
	radioButtonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("RadioButton_Checked", 0),
		/*1*/ imports.NewTable("RadioButton_Class", 0),
		/*2*/ imports.NewTable("RadioButton_Create", 0),
		/*3*/ imports.NewTable("RadioButton_DragCursor", 0),
		/*4*/ imports.NewTable("RadioButton_DragKind", 0),
		/*5*/ imports.NewTable("RadioButton_DragMode", 0),
		/*6*/ imports.NewTable("RadioButton_ParentColor", 0),
		/*7*/ imports.NewTable("RadioButton_ParentFont", 0),
		/*8*/ imports.NewTable("RadioButton_ParentShowHint", 0),
		/*9*/ imports.NewTable("RadioButton_SetOnContextPopup", 0),
		/*10*/ imports.NewTable("RadioButton_SetOnDragDrop", 0),
		/*11*/ imports.NewTable("RadioButton_SetOnDragOver", 0),
		/*12*/ imports.NewTable("RadioButton_SetOnEndDrag", 0),
		/*13*/ imports.NewTable("RadioButton_SetOnMouseDown", 0),
		/*14*/ imports.NewTable("RadioButton_SetOnMouseEnter", 0),
		/*15*/ imports.NewTable("RadioButton_SetOnMouseLeave", 0),
		/*16*/ imports.NewTable("RadioButton_SetOnMouseMove", 0),
		/*17*/ imports.NewTable("RadioButton_SetOnMouseUp", 0),
		/*18*/ imports.NewTable("RadioButton_SetOnMouseWheel", 0),
		/*19*/ imports.NewTable("RadioButton_SetOnMouseWheelDown", 0),
		/*20*/ imports.NewTable("RadioButton_SetOnMouseWheelUp", 0),
		/*21*/ imports.NewTable("RadioButton_SetOnStartDrag", 0),
	}
)

func radioButtonImportAPI() *imports.Imports {
	if radioButtonImport == nil {
		radioButtonImport = NewDefaultImports()
		radioButtonImport.SetImportTable(radioButtonImportTables)
		radioButtonImportTables = nil
	}
	return radioButtonImport
}
