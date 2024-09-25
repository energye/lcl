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

// IImage Parent: ICustomImage
type IImage interface {
	ICustomImage
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnPaint(fn TNotifyEvent)                     // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TImage Parent: TCustomImage
type TImage struct {
	TCustomImage
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	paintPtr           uintptr
	startDragPtr       uintptr
}

func NewImage(AOwner IComponent) IImage {
	r1 := LCL().SysCallN(3410, GetObjectUintptr(AOwner))
	return AsImage(r1)
}

func (m *TImage) DragCursor() TCursor {
	r1 := LCL().SysCallN(3411, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TImage) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3411, 1, m.Instance(), uintptr(AValue))
}

func (m *TImage) DragMode() TDragMode {
	r1 := LCL().SysCallN(3412, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TImage) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3412, 1, m.Instance(), uintptr(AValue))
}

func (m *TImage) ParentShowHint() bool {
	r1 := LCL().SysCallN(3413, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TImage) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3413, 1, m.Instance(), PascalBool(AValue))
}

func ImageClass() TClass {
	ret := LCL().SysCallN(3409)
	return TClass(ret)
}

func (m *TImage) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3414, m.Instance(), m.contextPopupPtr)
}

func (m *TImage) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3415, m.Instance(), m.dblClickPtr)
}

func (m *TImage) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3416, m.Instance(), m.dragDropPtr)
}

func (m *TImage) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3417, m.Instance(), m.dragOverPtr)
}

func (m *TImage) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3418, m.Instance(), m.endDragPtr)
}

func (m *TImage) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3419, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TImage) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3420, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TImage) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3421, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TImage) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3422, m.Instance(), m.paintPtr)
}

func (m *TImage) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3423, m.Instance(), m.startDragPtr)
}
