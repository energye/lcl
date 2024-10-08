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

// IMaskEdit Parent: ICustomMaskEdit
type IMaskEdit interface {
	ICustomMaskEdit
	IsMasked() bool                                // property
	EditText() string                              // property
	SetEditText(AValue string)                     // property
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
	EditMask() string                              // property
	SetEditMask(AValue string)                     // property
	SpaceChar() Char                               // property
	SetSpaceChar(AValue Char)                      // property
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnEndDock(fn TEndDragEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TMaskEdit Parent: TCustomMaskEdit
type TMaskEdit struct {
	TCustomMaskEdit
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	editingDonePtr    uintptr
	endDockPtr        uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDockPtr      uintptr
	startDragPtr      uintptr
}

func NewMaskEdit(TheOwner IComponent) IMaskEdit {
	r1 := LCL().SysCallN(4205, GetObjectUintptr(TheOwner))
	return AsMaskEdit(r1)
}

func (m *TMaskEdit) IsMasked() bool {
	r1 := LCL().SysCallN(4211, m.Instance())
	return GoBool(r1)
}

func (m *TMaskEdit) EditText() string {
	r1 := LCL().SysCallN(4210, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TMaskEdit) SetEditText(AValue string) {
	LCL().SysCallN(4210, 1, m.Instance(), PascalStr(AValue))
}

func (m *TMaskEdit) AutoSelect() bool {
	r1 := LCL().SysCallN(4203, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMaskEdit) SetAutoSelect(AValue bool) {
	LCL().SysCallN(4203, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMaskEdit) DragCursor() TCursor {
	r1 := LCL().SysCallN(4206, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TMaskEdit) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4206, 1, m.Instance(), uintptr(AValue))
}

func (m *TMaskEdit) DragKind() TDragKind {
	r1 := LCL().SysCallN(4207, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TMaskEdit) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4207, 1, m.Instance(), uintptr(AValue))
}

func (m *TMaskEdit) DragMode() TDragMode {
	r1 := LCL().SysCallN(4208, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TMaskEdit) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4208, 1, m.Instance(), uintptr(AValue))
}

func (m *TMaskEdit) ParentColor() bool {
	r1 := LCL().SysCallN(4212, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMaskEdit) SetParentColor(AValue bool) {
	LCL().SysCallN(4212, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMaskEdit) ParentFont() bool {
	r1 := LCL().SysCallN(4213, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMaskEdit) SetParentFont(AValue bool) {
	LCL().SysCallN(4213, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMaskEdit) ParentShowHint() bool {
	r1 := LCL().SysCallN(4214, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMaskEdit) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4214, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMaskEdit) EditMask() string {
	r1 := LCL().SysCallN(4209, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TMaskEdit) SetEditMask(AValue string) {
	LCL().SysCallN(4209, 1, m.Instance(), PascalStr(AValue))
}

func (m *TMaskEdit) SpaceChar() Char {
	r1 := LCL().SysCallN(4231, 0, m.Instance(), 0)
	return Char(r1)
}

func (m *TMaskEdit) SetSpaceChar(AValue Char) {
	LCL().SysCallN(4231, 1, m.Instance(), uintptr(AValue))
}

func MaskEditClass() TClass {
	ret := LCL().SysCallN(4204)
	return TClass(ret)
}

func (m *TMaskEdit) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4215, m.Instance(), m.dblClickPtr)
}

func (m *TMaskEdit) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4216, m.Instance(), m.dragDropPtr)
}

func (m *TMaskEdit) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4217, m.Instance(), m.dragOverPtr)
}

func (m *TMaskEdit) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4218, m.Instance(), m.editingDonePtr)
}

func (m *TMaskEdit) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4219, m.Instance(), m.endDockPtr)
}

func (m *TMaskEdit) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4220, m.Instance(), m.endDragPtr)
}

func (m *TMaskEdit) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4221, m.Instance(), m.mouseDownPtr)
}

func (m *TMaskEdit) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4222, m.Instance(), m.mouseEnterPtr)
}

func (m *TMaskEdit) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4223, m.Instance(), m.mouseLeavePtr)
}

func (m *TMaskEdit) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4224, m.Instance(), m.mouseMovePtr)
}

func (m *TMaskEdit) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4225, m.Instance(), m.mouseUpPtr)
}

func (m *TMaskEdit) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4226, m.Instance(), m.mouseWheelPtr)
}

func (m *TMaskEdit) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4227, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TMaskEdit) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4228, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TMaskEdit) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4229, m.Instance(), m.startDockPtr)
}

func (m *TMaskEdit) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4230, m.Instance(), m.startDragPtr)
}
