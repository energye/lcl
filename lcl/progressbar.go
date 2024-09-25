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

// IProgressBar Parent: ICustomProgressBar
type IProgressBar interface {
	ICustomProgressBar
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
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TProgressBar Parent: TCustomProgressBar
type TProgressBar struct {
	TCustomProgressBar
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
	startDockPtr      uintptr
	startDragPtr      uintptr
}

func NewProgressBar(AOwner IComponent) IProgressBar {
	r1 := LCL().SysCallN(4690, GetObjectUintptr(AOwner))
	return AsProgressBar(r1)
}

func (m *TProgressBar) DragCursor() TCursor {
	r1 := LCL().SysCallN(4691, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TProgressBar) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4691, 1, m.Instance(), uintptr(AValue))
}

func (m *TProgressBar) DragKind() TDragKind {
	r1 := LCL().SysCallN(4692, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TProgressBar) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4692, 1, m.Instance(), uintptr(AValue))
}

func (m *TProgressBar) DragMode() TDragMode {
	r1 := LCL().SysCallN(4693, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TProgressBar) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4693, 1, m.Instance(), uintptr(AValue))
}

func (m *TProgressBar) ParentColor() bool {
	r1 := LCL().SysCallN(4694, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TProgressBar) SetParentColor(AValue bool) {
	LCL().SysCallN(4694, 1, m.Instance(), PascalBool(AValue))
}

func (m *TProgressBar) ParentFont() bool {
	r1 := LCL().SysCallN(4695, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TProgressBar) SetParentFont(AValue bool) {
	LCL().SysCallN(4695, 1, m.Instance(), PascalBool(AValue))
}

func (m *TProgressBar) ParentShowHint() bool {
	r1 := LCL().SysCallN(4696, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TProgressBar) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4696, 1, m.Instance(), PascalBool(AValue))
}

func ProgressBarClass() TClass {
	ret := LCL().SysCallN(4689)
	return TClass(ret)
}

func (m *TProgressBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4697, m.Instance(), m.contextPopupPtr)
}

func (m *TProgressBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4698, m.Instance(), m.dragDropPtr)
}

func (m *TProgressBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4699, m.Instance(), m.dragOverPtr)
}

func (m *TProgressBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4700, m.Instance(), m.endDragPtr)
}

func (m *TProgressBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4701, m.Instance(), m.mouseDownPtr)
}

func (m *TProgressBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4702, m.Instance(), m.mouseEnterPtr)
}

func (m *TProgressBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4703, m.Instance(), m.mouseLeavePtr)
}

func (m *TProgressBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4704, m.Instance(), m.mouseMovePtr)
}

func (m *TProgressBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4705, m.Instance(), m.mouseUpPtr)
}

func (m *TProgressBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4706, m.Instance(), m.mouseWheelPtr)
}

func (m *TProgressBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4707, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TProgressBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4708, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TProgressBar) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4709, m.Instance(), m.startDockPtr)
}

func (m *TProgressBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4710, m.Instance(), m.startDragPtr)
}
