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
	r1 := LCL().SysCallN(568, GetObjectUintptr(TheOwner))
	return AsCheckBox(r1)
}

func (m *TCheckBox) Checked() bool {
	r1 := LCL().SysCallN(566, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckBox) SetChecked(AValue bool) {
	LCL().SysCallN(566, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckBox) DragCursor() TCursor {
	r1 := LCL().SysCallN(569, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCheckBox) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(569, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckBox) DragKind() TDragKind {
	r1 := LCL().SysCallN(570, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TCheckBox) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(570, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckBox) DragMode() TDragMode {
	r1 := LCL().SysCallN(571, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCheckBox) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(571, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckBox) ParentColor() bool {
	r1 := LCL().SysCallN(572, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckBox) SetParentColor(AValue bool) {
	LCL().SysCallN(572, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckBox) ParentFont() bool {
	r1 := LCL().SysCallN(573, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckBox) SetParentFont(AValue bool) {
	LCL().SysCallN(573, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckBox) ParentShowHint() bool {
	r1 := LCL().SysCallN(574, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckBox) SetParentShowHint(AValue bool) {
	LCL().SysCallN(574, 1, m.Instance(), PascalBool(AValue))
}

func CheckBoxClass() TClass {
	ret := LCL().SysCallN(567)
	return TClass(ret)
}

func (m *TCheckBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(575, m.Instance(), m.contextPopupPtr)
}

func (m *TCheckBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(576, m.Instance(), m.dragDropPtr)
}

func (m *TCheckBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(577, m.Instance(), m.dragOverPtr)
}

func (m *TCheckBox) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(578, m.Instance(), m.editingDonePtr)
}

func (m *TCheckBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(579, m.Instance(), m.endDragPtr)
}

func (m *TCheckBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(580, m.Instance(), m.mouseDownPtr)
}

func (m *TCheckBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(581, m.Instance(), m.mouseEnterPtr)
}

func (m *TCheckBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(582, m.Instance(), m.mouseLeavePtr)
}

func (m *TCheckBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(583, m.Instance(), m.mouseMovePtr)
}

func (m *TCheckBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(584, m.Instance(), m.mouseUpPtr)
}

func (m *TCheckBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(585, m.Instance(), m.mouseWheelPtr)
}

func (m *TCheckBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(586, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCheckBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(587, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCheckBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(588, m.Instance(), m.startDragPtr)
}
