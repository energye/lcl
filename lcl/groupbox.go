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

// IGroupBox Parent: ICustomGroupBox
type IGroupBox interface {
	ICustomGroupBox
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
	SetOnEndDock(fn TEndDragEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)         // property event
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

// TGroupBox Parent: TCustomGroupBox
type TGroupBox struct {
	TCustomGroupBox
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDockPtr        uintptr
	endDragPtr        uintptr
	getSiteInfoPtr    uintptr
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

func NewGroupBox(AOwner IComponent) IGroupBox {
	r1 := LCL().SysCallN(3293, GetObjectUintptr(AOwner))
	return AsGroupBox(r1)
}

func (m *TGroupBox) DragCursor() TCursor {
	r1 := LCL().SysCallN(3294, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TGroupBox) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3294, 1, m.Instance(), uintptr(AValue))
}

func (m *TGroupBox) DragKind() TDragKind {
	r1 := LCL().SysCallN(3295, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TGroupBox) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3295, 1, m.Instance(), uintptr(AValue))
}

func (m *TGroupBox) DragMode() TDragMode {
	r1 := LCL().SysCallN(3296, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TGroupBox) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3296, 1, m.Instance(), uintptr(AValue))
}

func (m *TGroupBox) ParentColor() bool {
	r1 := LCL().SysCallN(3297, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGroupBox) SetParentColor(AValue bool) {
	LCL().SysCallN(3297, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGroupBox) ParentFont() bool {
	r1 := LCL().SysCallN(3298, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGroupBox) SetParentFont(AValue bool) {
	LCL().SysCallN(3298, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGroupBox) ParentShowHint() bool {
	r1 := LCL().SysCallN(3299, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGroupBox) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3299, 1, m.Instance(), PascalBool(AValue))
}

func GroupBoxClass() TClass {
	ret := LCL().SysCallN(3292)
	return TClass(ret)
}

func (m *TGroupBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3300, m.Instance(), m.contextPopupPtr)
}

func (m *TGroupBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3301, m.Instance(), m.dblClickPtr)
}

func (m *TGroupBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3302, m.Instance(), m.dragDropPtr)
}

func (m *TGroupBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3303, m.Instance(), m.dragOverPtr)
}

func (m *TGroupBox) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3304, m.Instance(), m.endDockPtr)
}

func (m *TGroupBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3305, m.Instance(), m.endDragPtr)
}

func (m *TGroupBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3306, m.Instance(), m.getSiteInfoPtr)
}

func (m *TGroupBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3307, m.Instance(), m.mouseDownPtr)
}

func (m *TGroupBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3308, m.Instance(), m.mouseEnterPtr)
}

func (m *TGroupBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3309, m.Instance(), m.mouseLeavePtr)
}

func (m *TGroupBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3310, m.Instance(), m.mouseMovePtr)
}

func (m *TGroupBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3311, m.Instance(), m.mouseUpPtr)
}

func (m *TGroupBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3312, m.Instance(), m.mouseWheelPtr)
}

func (m *TGroupBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3313, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TGroupBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3314, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TGroupBox) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3315, m.Instance(), m.startDockPtr)
}

func (m *TGroupBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3316, m.Instance(), m.startDragPtr)
}
