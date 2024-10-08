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

// IPaintBox Parent: IGraphicControl
type IPaintBox interface {
	IGraphicControl
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
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
	SetOnPaint(fn TNotifyEvent)                     // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TPaintBox Parent: TGraphicControl
type TPaintBox struct {
	TGraphicControl
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
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
	paintPtr           uintptr
	startDragPtr       uintptr
}

func NewPaintBox(AOwner IComponent) IPaintBox {
	r1 := LCL().SysCallN(4481, GetObjectUintptr(AOwner))
	return AsPaintBox(r1)
}

func (m *TPaintBox) DragCursor() TCursor {
	r1 := LCL().SysCallN(4482, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TPaintBox) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4482, 1, m.Instance(), uintptr(AValue))
}

func (m *TPaintBox) DragMode() TDragMode {
	r1 := LCL().SysCallN(4483, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TPaintBox) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4483, 1, m.Instance(), uintptr(AValue))
}

func (m *TPaintBox) ParentColor() bool {
	r1 := LCL().SysCallN(4484, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPaintBox) SetParentColor(AValue bool) {
	LCL().SysCallN(4484, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPaintBox) ParentFont() bool {
	r1 := LCL().SysCallN(4485, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPaintBox) SetParentFont(AValue bool) {
	LCL().SysCallN(4485, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPaintBox) ParentShowHint() bool {
	r1 := LCL().SysCallN(4486, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPaintBox) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4486, 1, m.Instance(), PascalBool(AValue))
}

func PaintBoxClass() TClass {
	ret := LCL().SysCallN(4480)
	return TClass(ret)
}

func (m *TPaintBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4487, m.Instance(), m.contextPopupPtr)
}

func (m *TPaintBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4488, m.Instance(), m.dblClickPtr)
}

func (m *TPaintBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4489, m.Instance(), m.dragDropPtr)
}

func (m *TPaintBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4490, m.Instance(), m.dragOverPtr)
}

func (m *TPaintBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4491, m.Instance(), m.endDragPtr)
}

func (m *TPaintBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4492, m.Instance(), m.mouseDownPtr)
}

func (m *TPaintBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4493, m.Instance(), m.mouseEnterPtr)
}

func (m *TPaintBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4494, m.Instance(), m.mouseLeavePtr)
}

func (m *TPaintBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4495, m.Instance(), m.mouseMovePtr)
}

func (m *TPaintBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4496, m.Instance(), m.mouseUpPtr)
}

func (m *TPaintBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4497, m.Instance(), m.mouseWheelPtr)
}

func (m *TPaintBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4498, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TPaintBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4502, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TPaintBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4499, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TPaintBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4500, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TPaintBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4501, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TPaintBox) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4503, m.Instance(), m.paintPtr)
}

func (m *TPaintBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4504, m.Instance(), m.startDragPtr)
}
