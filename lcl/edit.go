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

// IEdit Parent: ICustomEdit
type IEdit interface {
	ICustomEdit
	AutoSelected() bool                            // property
	SetAutoSelected(AValue bool)                   // property
	AutoSelect() bool                              // property
	SetAutoSelect(AValue bool)                     // property
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

// TEdit Parent: TCustomEdit
type TEdit struct {
	TCustomEdit
	contextPopupPtr   uintptr
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

func NewEdit(AOwner IComponent) IEdit {
	r1 := editImportAPI().SysCallN(3, GetObjectUintptr(AOwner))
	return AsEdit(r1)
}

func (m *TEdit) AutoSelected() bool {
	r1 := editImportAPI().SysCallN(1, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEdit) SetAutoSelected(AValue bool) {
	editImportAPI().SysCallN(1, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEdit) AutoSelect() bool {
	r1 := editImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEdit) SetAutoSelect(AValue bool) {
	editImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEdit) DragCursor() TCursor {
	r1 := editImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TEdit) SetDragCursor(AValue TCursor) {
	editImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TEdit) DragKind() TDragKind {
	r1 := editImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TEdit) SetDragKind(AValue TDragKind) {
	editImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TEdit) DragMode() TDragMode {
	r1 := editImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TEdit) SetDragMode(AValue TDragMode) {
	editImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TEdit) ParentColor() bool {
	r1 := editImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEdit) SetParentColor(AValue bool) {
	editImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEdit) ParentFont() bool {
	r1 := editImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEdit) SetParentFont(AValue bool) {
	editImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TEdit) ParentShowHint() bool {
	r1 := editImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TEdit) SetParentShowHint(AValue bool) {
	editImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func EditClass() TClass {
	ret := editImportAPI().SysCallN(2)
	return TClass(ret)
}

func (m *TEdit) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(10, m.Instance(), m.contextPopupPtr)
}

func (m *TEdit) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(11, m.Instance(), m.dblClickPtr)
}

func (m *TEdit) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(12, m.Instance(), m.dragDropPtr)
}

func (m *TEdit) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(13, m.Instance(), m.dragOverPtr)
}

func (m *TEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(14, m.Instance(), m.editingDonePtr)
}

func (m *TEdit) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(15, m.Instance(), m.endDragPtr)
}

func (m *TEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(16, m.Instance(), m.mouseDownPtr)
}

func (m *TEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(17, m.Instance(), m.mouseEnterPtr)
}

func (m *TEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(18, m.Instance(), m.mouseLeavePtr)
}

func (m *TEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(19, m.Instance(), m.mouseMovePtr)
}

func (m *TEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(20, m.Instance(), m.mouseUpPtr)
}

func (m *TEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(21, m.Instance(), m.mouseWheelPtr)
}

func (m *TEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(22, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(23, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TEdit) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	editImportAPI().SysCallN(24, m.Instance(), m.startDragPtr)
}

var (
	editImport       *imports.Imports = nil
	editImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Edit_AutoSelect", 0),
		/*1*/ imports.NewTable("Edit_AutoSelected", 0),
		/*2*/ imports.NewTable("Edit_Class", 0),
		/*3*/ imports.NewTable("Edit_Create", 0),
		/*4*/ imports.NewTable("Edit_DragCursor", 0),
		/*5*/ imports.NewTable("Edit_DragKind", 0),
		/*6*/ imports.NewTable("Edit_DragMode", 0),
		/*7*/ imports.NewTable("Edit_ParentColor", 0),
		/*8*/ imports.NewTable("Edit_ParentFont", 0),
		/*9*/ imports.NewTable("Edit_ParentShowHint", 0),
		/*10*/ imports.NewTable("Edit_SetOnContextPopup", 0),
		/*11*/ imports.NewTable("Edit_SetOnDblClick", 0),
		/*12*/ imports.NewTable("Edit_SetOnDragDrop", 0),
		/*13*/ imports.NewTable("Edit_SetOnDragOver", 0),
		/*14*/ imports.NewTable("Edit_SetOnEditingDone", 0),
		/*15*/ imports.NewTable("Edit_SetOnEndDrag", 0),
		/*16*/ imports.NewTable("Edit_SetOnMouseDown", 0),
		/*17*/ imports.NewTable("Edit_SetOnMouseEnter", 0),
		/*18*/ imports.NewTable("Edit_SetOnMouseLeave", 0),
		/*19*/ imports.NewTable("Edit_SetOnMouseMove", 0),
		/*20*/ imports.NewTable("Edit_SetOnMouseUp", 0),
		/*21*/ imports.NewTable("Edit_SetOnMouseWheel", 0),
		/*22*/ imports.NewTable("Edit_SetOnMouseWheelDown", 0),
		/*23*/ imports.NewTable("Edit_SetOnMouseWheelUp", 0),
		/*24*/ imports.NewTable("Edit_SetOnStartDrag", 0),
	}
)

func editImportAPI() *imports.Imports {
	if editImport == nil {
		editImport = NewDefaultImports()
		editImport.SetImportTable(editImportTables)
		editImportTables = nil
	}
	return editImport
}
