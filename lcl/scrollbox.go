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

// IScrollBox Parent: IScrollingWinControl
type IScrollBox interface {
	IScrollingWinControl
	AutoScroll() bool                                  // property
	SetAutoScroll(AValue bool)                         // property
	DragCursor() TCursor                               // property
	SetDragCursor(AValue TCursor)                      // property
	DragKind() TDragKind                               // property
	SetDragKind(AValue TDragKind)                      // property
	DragMode() TDragMode                               // property
	SetDragMode(AValue TDragMode)                      // property
	ParentBackground() bool                            // property
	SetParentBackground(AValue bool)                   // property
	ParentColor() bool                                 // property
	SetParentColor(AValue bool)                        // property
	ParentFont() bool                                  // property
	SetParentFont(AValue bool)                         // property
	ParentShowHint() bool                              // property
	SetParentShowHint(AValue bool)                     // property
	SetOnConstrainedResize(fn TConstrainedResizeEvent) // property event
	SetOnDblClick(fn TNotifyEvent)                     // property event
	SetOnDragDrop(fn TDragDropEvent)                   // property event
	SetOnDragOver(fn TDragOverEvent)                   // property event
	SetOnEndDock(fn TEndDragEvent)                     // property event
	SetOnEndDrag(fn TEndDragEvent)                     // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)             // property event
	SetOnMouseDown(fn TMouseEvent)                     // property event
	SetOnMouseEnter(fn TNotifyEvent)                   // property event
	SetOnMouseLeave(fn TNotifyEvent)                   // property event
	SetOnMouseMove(fn TMouseMoveEvent)                 // property event
	SetOnMouseUp(fn TMouseEvent)                       // property event
	SetOnMouseWheel(fn TMouseWheelEvent)               // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)    // property event
	SetOnStartDock(fn TStartDockEvent)                 // property event
	SetOnStartDrag(fn TStartDragEvent)                 // property event
}

// TScrollBox Parent: TScrollingWinControl
type TScrollBox struct {
	TScrollingWinControl
	constrainedResizePtr uintptr
	dblClickPtr          uintptr
	dragDropPtr          uintptr
	dragOverPtr          uintptr
	endDockPtr           uintptr
	endDragPtr           uintptr
	getSiteInfoPtr       uintptr
	mouseDownPtr         uintptr
	mouseEnterPtr        uintptr
	mouseLeavePtr        uintptr
	mouseMovePtr         uintptr
	mouseUpPtr           uintptr
	mouseWheelPtr        uintptr
	mouseWheelDownPtr    uintptr
	mouseWheelUpPtr      uintptr
	mouseWheelHorzPtr    uintptr
	mouseWheelLeftPtr    uintptr
	mouseWheelRightPtr   uintptr
	startDockPtr         uintptr
	startDragPtr         uintptr
}

func NewScrollBox(AOwner IComponent) IScrollBox {
	r1 := LCL().SysCallN(4978, GetObjectUintptr(AOwner))
	return AsScrollBox(r1)
}

func (m *TScrollBox) AutoScroll() bool {
	r1 := LCL().SysCallN(4976, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBox) SetAutoScroll(AValue bool) {
	LCL().SysCallN(4976, 1, m.Instance(), PascalBool(AValue))
}

func (m *TScrollBox) DragCursor() TCursor {
	r1 := LCL().SysCallN(4979, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TScrollBox) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4979, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBox) DragKind() TDragKind {
	r1 := LCL().SysCallN(4980, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TScrollBox) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4980, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBox) DragMode() TDragMode {
	r1 := LCL().SysCallN(4981, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TScrollBox) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4981, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBox) ParentBackground() bool {
	r1 := LCL().SysCallN(4982, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBox) SetParentBackground(AValue bool) {
	LCL().SysCallN(4982, 1, m.Instance(), PascalBool(AValue))
}

func (m *TScrollBox) ParentColor() bool {
	r1 := LCL().SysCallN(4983, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBox) SetParentColor(AValue bool) {
	LCL().SysCallN(4983, 1, m.Instance(), PascalBool(AValue))
}

func (m *TScrollBox) ParentFont() bool {
	r1 := LCL().SysCallN(4984, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBox) SetParentFont(AValue bool) {
	LCL().SysCallN(4984, 1, m.Instance(), PascalBool(AValue))
}

func (m *TScrollBox) ParentShowHint() bool {
	r1 := LCL().SysCallN(4985, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBox) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4985, 1, m.Instance(), PascalBool(AValue))
}

func ScrollBoxClass() TClass {
	ret := LCL().SysCallN(4977)
	return TClass(ret)
}

func (m *TScrollBox) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if m.constrainedResizePtr != 0 {
		RemoveEventElement(m.constrainedResizePtr)
	}
	m.constrainedResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4986, m.Instance(), m.constrainedResizePtr)
}

func (m *TScrollBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4987, m.Instance(), m.dblClickPtr)
}

func (m *TScrollBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4988, m.Instance(), m.dragDropPtr)
}

func (m *TScrollBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4989, m.Instance(), m.dragOverPtr)
}

func (m *TScrollBox) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4990, m.Instance(), m.endDockPtr)
}

func (m *TScrollBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4991, m.Instance(), m.endDragPtr)
}

func (m *TScrollBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4992, m.Instance(), m.getSiteInfoPtr)
}

func (m *TScrollBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4993, m.Instance(), m.mouseDownPtr)
}

func (m *TScrollBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4994, m.Instance(), m.mouseEnterPtr)
}

func (m *TScrollBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4995, m.Instance(), m.mouseLeavePtr)
}

func (m *TScrollBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4996, m.Instance(), m.mouseMovePtr)
}

func (m *TScrollBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4997, m.Instance(), m.mouseUpPtr)
}

func (m *TScrollBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4998, m.Instance(), m.mouseWheelPtr)
}

func (m *TScrollBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4999, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TScrollBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5003, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TScrollBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5000, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TScrollBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5001, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TScrollBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5002, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TScrollBox) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5004, m.Instance(), m.startDockPtr)
}

func (m *TScrollBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5005, m.Instance(), m.startDragPtr)
}
