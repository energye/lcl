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

// IShape Parent: IGraphicControl
type IShape interface {
	IGraphicControl
	Brush() IBrush                                  // property
	SetBrush(AValue IBrush)                         // property
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	Pen() IPen                                      // property
	SetPen(AValue IPen)                             // property
	Shape() TShapeType                              // property
	SetShape(AValue TShapeType)                     // property
	Paint()                                         // procedure
	StyleChanged(Sender IObject)                    // procedure
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDock(fn TEndDragEvent)                  // property event
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
	SetOnStartDock(fn TStartDockEvent)              // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TShape Parent: TGraphicControl
type TShape struct {
	TGraphicControl
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDockPtr         uintptr
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
	startDockPtr       uintptr
	startDragPtr       uintptr
}

func NewShape(TheOwner IComponent) IShape {
	r1 := LCL().SysCallN(5016, GetObjectUintptr(TheOwner))
	return AsShape(r1)
}

func (m *TShape) Brush() IBrush {
	r1 := LCL().SysCallN(5014, 0, m.Instance(), 0)
	return AsBrush(r1)
}

func (m *TShape) SetBrush(AValue IBrush) {
	LCL().SysCallN(5014, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TShape) DragCursor() TCursor {
	r1 := LCL().SysCallN(5017, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TShape) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(5017, 1, m.Instance(), uintptr(AValue))
}

func (m *TShape) DragKind() TDragKind {
	r1 := LCL().SysCallN(5018, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TShape) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(5018, 1, m.Instance(), uintptr(AValue))
}

func (m *TShape) DragMode() TDragMode {
	r1 := LCL().SysCallN(5019, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TShape) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(5019, 1, m.Instance(), uintptr(AValue))
}

func (m *TShape) ParentShowHint() bool {
	r1 := LCL().SysCallN(5021, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TShape) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5021, 1, m.Instance(), PascalBool(AValue))
}

func (m *TShape) Pen() IPen {
	r1 := LCL().SysCallN(5022, 0, m.Instance(), 0)
	return AsPen(r1)
}

func (m *TShape) SetPen(AValue IPen) {
	LCL().SysCallN(5022, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TShape) Shape() TShapeType {
	r1 := LCL().SysCallN(5041, 0, m.Instance(), 0)
	return TShapeType(r1)
}

func (m *TShape) SetShape(AValue TShapeType) {
	LCL().SysCallN(5041, 1, m.Instance(), uintptr(AValue))
}

func ShapeClass() TClass {
	ret := LCL().SysCallN(5015)
	return TClass(ret)
}

func (m *TShape) Paint() {
	LCL().SysCallN(5020, m.Instance())
}

func (m *TShape) StyleChanged(Sender IObject) {
	LCL().SysCallN(5042, m.Instance(), GetObjectUintptr(Sender))
}

func (m *TShape) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5023, m.Instance(), m.dragDropPtr)
}

func (m *TShape) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5024, m.Instance(), m.dragOverPtr)
}

func (m *TShape) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5025, m.Instance(), m.endDockPtr)
}

func (m *TShape) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5026, m.Instance(), m.endDragPtr)
}

func (m *TShape) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5027, m.Instance(), m.mouseDownPtr)
}

func (m *TShape) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5028, m.Instance(), m.mouseEnterPtr)
}

func (m *TShape) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5029, m.Instance(), m.mouseLeavePtr)
}

func (m *TShape) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5030, m.Instance(), m.mouseMovePtr)
}

func (m *TShape) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5031, m.Instance(), m.mouseUpPtr)
}

func (m *TShape) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5032, m.Instance(), m.mouseWheelPtr)
}

func (m *TShape) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5033, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TShape) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5037, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TShape) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5034, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TShape) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5035, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TShape) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5036, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TShape) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5038, m.Instance(), m.paintPtr)
}

func (m *TShape) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5039, m.Instance(), m.startDockPtr)
}

func (m *TShape) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5040, m.Instance(), m.startDragPtr)
}
