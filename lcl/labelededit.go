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
	r1 := LCL().SysCallN(3493, GetObjectUintptr(TheOwner))
	return AsLabeledEdit(r1)
}

func (m *TLabeledEdit) AutoSelect() bool {
	r1 := LCL().SysCallN(3491, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabeledEdit) SetAutoSelect(AValue bool) {
	LCL().SysCallN(3491, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabeledEdit) DragCursor() TCursor {
	r1 := LCL().SysCallN(3494, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLabeledEdit) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3494, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabeledEdit) DragMode() TDragMode {
	r1 := LCL().SysCallN(3495, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TLabeledEdit) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3495, 1, m.Instance(), uintptr(AValue))
}

func (m *TLabeledEdit) ParentColor() bool {
	r1 := LCL().SysCallN(3496, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabeledEdit) SetParentColor(AValue bool) {
	LCL().SysCallN(3496, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabeledEdit) ParentFont() bool {
	r1 := LCL().SysCallN(3497, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabeledEdit) SetParentFont(AValue bool) {
	LCL().SysCallN(3497, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLabeledEdit) ParentShowHint() bool {
	r1 := LCL().SysCallN(3498, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLabeledEdit) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3498, 1, m.Instance(), PascalBool(AValue))
}

func LabeledEditClass() TClass {
	ret := LCL().SysCallN(3492)
	return TClass(ret)
}

func (m *TLabeledEdit) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3499, m.Instance(), m.dblClickPtr)
}

func (m *TLabeledEdit) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3500, m.Instance(), m.dragDropPtr)
}

func (m *TLabeledEdit) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3501, m.Instance(), m.dragOverPtr)
}

func (m *TLabeledEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3502, m.Instance(), m.editingDonePtr)
}

func (m *TLabeledEdit) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3503, m.Instance(), m.endDragPtr)
}

func (m *TLabeledEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3504, m.Instance(), m.mouseDownPtr)
}

func (m *TLabeledEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3505, m.Instance(), m.mouseEnterPtr)
}

func (m *TLabeledEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3506, m.Instance(), m.mouseLeavePtr)
}

func (m *TLabeledEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3507, m.Instance(), m.mouseMovePtr)
}

func (m *TLabeledEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3508, m.Instance(), m.mouseUpPtr)
}

func (m *TLabeledEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3509, m.Instance(), m.mouseWheelPtr)
}

func (m *TLabeledEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3510, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TLabeledEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3511, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TLabeledEdit) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3512, m.Instance(), m.startDragPtr)
}
