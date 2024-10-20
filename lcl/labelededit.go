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

// ILabeledEdit Parent: ICustomLabeledEdit
type ILabeledEdit interface {
	ICustomLabeledEdit
	AutoSelect() bool                              // property
	SetAutoSelect(AValue bool)                     // property
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnDblClick(fn TNotifyEvent)                 // property event
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

// TLabeledEdit Parent: TCustomLabeledEdit
type TLabeledEdit struct {
	TCustomLabeledEdit
	dblClickPtr       uintptr
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

func NewLabeledEdit(TheOwner IComponent) ILabeledEdit {
	r1 := labeledEditImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsLabeledEdit(r1)
}

func (m *TLabeledEdit) AutoSelect() bool {
	r1 := labeledEditImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabeledEdit) SetAutoSelect(AValue bool) {
	labeledEditImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabeledEdit) DragCursor() TCursor {
	r1 := labeledEditImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLabeledEdit) SetDragCursor(AValue TCursor) {
	labeledEditImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabeledEdit) DragMode() TDragMode {
	r1 := labeledEditImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TLabeledEdit) SetDragMode(AValue TDragMode) {
	labeledEditImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabeledEdit) ParentColor() bool {
	r1 := labeledEditImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabeledEdit) SetParentColor(AValue bool) {
	labeledEditImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabeledEdit) ParentFont() bool {
	r1 := labeledEditImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabeledEdit) SetParentFont(AValue bool) {
	labeledEditImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabeledEdit) ParentShowHint() bool {
	r1 := labeledEditImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabeledEdit) SetParentShowHint(AValue bool) {
	labeledEditImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func LabeledEditClass() TClass {
	ret := labeledEditImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TLabeledEdit) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(8, m.Instance(), m.dblClickPtr)
}

func (m *TLabeledEdit) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(9, m.Instance(), m.dragDropPtr)
}

func (m *TLabeledEdit) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(10, m.Instance(), m.dragOverPtr)
}

func (m *TLabeledEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(11, m.Instance(), m.editingDonePtr)
}

func (m *TLabeledEdit) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(12, m.Instance(), m.endDragPtr)
}

func (m *TLabeledEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(13, m.Instance(), m.mouseDownPtr)
}

func (m *TLabeledEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(14, m.Instance(), m.mouseEnterPtr)
}

func (m *TLabeledEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(15, m.Instance(), m.mouseLeavePtr)
}

func (m *TLabeledEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(16, m.Instance(), m.mouseMovePtr)
}

func (m *TLabeledEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(17, m.Instance(), m.mouseUpPtr)
}

func (m *TLabeledEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(18, m.Instance(), m.mouseWheelPtr)
}

func (m *TLabeledEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(19, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TLabeledEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(20, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TLabeledEdit) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	labeledEditImportAPI().SysCallN(21, m.Instance(), m.startDragPtr)
}

var (
	labeledEditImport       *imports.Imports = nil
	labeledEditImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("LabeledEdit_AutoSelect", 0),
		/*1*/ imports.NewTable("LabeledEdit_Class", 0),
		/*2*/ imports.NewTable("LabeledEdit_Create", 0),
		/*3*/ imports.NewTable("LabeledEdit_DragCursor", 0),
		/*4*/ imports.NewTable("LabeledEdit_DragMode", 0),
		/*5*/ imports.NewTable("LabeledEdit_ParentColor", 0),
		/*6*/ imports.NewTable("LabeledEdit_ParentFont", 0),
		/*7*/ imports.NewTable("LabeledEdit_ParentShowHint", 0),
		/*8*/ imports.NewTable("LabeledEdit_SetOnDblClick", 0),
		/*9*/ imports.NewTable("LabeledEdit_SetOnDragDrop", 0),
		/*10*/ imports.NewTable("LabeledEdit_SetOnDragOver", 0),
		/*11*/ imports.NewTable("LabeledEdit_SetOnEditingDone", 0),
		/*12*/ imports.NewTable("LabeledEdit_SetOnEndDrag", 0),
		/*13*/ imports.NewTable("LabeledEdit_SetOnMouseDown", 0),
		/*14*/ imports.NewTable("LabeledEdit_SetOnMouseEnter", 0),
		/*15*/ imports.NewTable("LabeledEdit_SetOnMouseLeave", 0),
		/*16*/ imports.NewTable("LabeledEdit_SetOnMouseMove", 0),
		/*17*/ imports.NewTable("LabeledEdit_SetOnMouseUp", 0),
		/*18*/ imports.NewTable("LabeledEdit_SetOnMouseWheel", 0),
		/*19*/ imports.NewTable("LabeledEdit_SetOnMouseWheelDown", 0),
		/*20*/ imports.NewTable("LabeledEdit_SetOnMouseWheelUp", 0),
		/*21*/ imports.NewTable("LabeledEdit_SetOnStartDrag", 0),
	}
)

func labeledEditImportAPI() *imports.Imports {
	if labeledEditImport == nil {
		labeledEditImport = NewDefaultImports()
		labeledEditImport.SetImportTable(labeledEditImportTables)
		labeledEditImportTables = nil
	}
	return labeledEditImport
}
