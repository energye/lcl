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
	r1 := LCL().SysCallN(4713, GetObjectUintptr(TheOwner))
	return AsRadioButton(r1)
}

func (m *TRadioButton) Checked() bool {
	r1 := LCL().SysCallN(4711, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetChecked(AValue bool) {
	LCL().SysCallN(4711, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioButton) DragCursor() TCursor {
	r1 := LCL().SysCallN(4714, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TRadioButton) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4714, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioButton) DragKind() TDragKind {
	r1 := LCL().SysCallN(4715, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TRadioButton) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4715, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioButton) DragMode() TDragMode {
	r1 := LCL().SysCallN(4716, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TRadioButton) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4716, 1, m.Instance(), uintptr(AValue))
}

func (m *TRadioButton) ParentColor() bool {
	r1 := LCL().SysCallN(4717, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetParentColor(AValue bool) {
	LCL().SysCallN(4717, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioButton) ParentFont() bool {
	r1 := LCL().SysCallN(4718, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetParentFont(AValue bool) {
	LCL().SysCallN(4718, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRadioButton) ParentShowHint() bool {
	r1 := LCL().SysCallN(4719, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRadioButton) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4719, 1, m.Instance(), PascalBool(AValue))
}

func RadioButtonClass() TClass {
	ret := LCL().SysCallN(4712)
	return TClass(ret)
}

func (m *TRadioButton) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4720, m.Instance(), m.contextPopupPtr)
}

func (m *TRadioButton) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4721, m.Instance(), m.dragDropPtr)
}

func (m *TRadioButton) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4722, m.Instance(), m.dragOverPtr)
}

func (m *TRadioButton) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4723, m.Instance(), m.endDragPtr)
}

func (m *TRadioButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4724, m.Instance(), m.mouseDownPtr)
}

func (m *TRadioButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4725, m.Instance(), m.mouseEnterPtr)
}

func (m *TRadioButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4726, m.Instance(), m.mouseLeavePtr)
}

func (m *TRadioButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4727, m.Instance(), m.mouseMovePtr)
}

func (m *TRadioButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4728, m.Instance(), m.mouseUpPtr)
}

func (m *TRadioButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4729, m.Instance(), m.mouseWheelPtr)
}

func (m *TRadioButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4730, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TRadioButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4731, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TRadioButton) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4732, m.Instance(), m.startDragPtr)
}
