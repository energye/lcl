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

// IScrollBar Parent: ICustomScrollBar
type IScrollBar interface {
	ICustomScrollBar
	DragCursor() TCursor                     // property
	SetDragCursor(AValue TCursor)            // property
	DragKind() TDragKind                     // property
	SetDragKind(AValue TDragKind)            // property
	DragMode() TDragMode                     // property
	SetDragMode(AValue TDragMode)            // property
	ParentShowHint() bool                    // property
	SetParentShowHint(AValue bool)           // property
	SetOnContextPopup(fn TContextPopupEvent) // property event
	SetOnDragDrop(fn TDragDropEvent)         // property event
	SetOnDragOver(fn TDragOverEvent)         // property event
	SetOnEndDrag(fn TEndDragEvent)           // property event
	SetOnStartDrag(fn TStartDragEvent)       // property event
}

// TScrollBar Parent: TCustomScrollBar
type TScrollBar struct {
	TCustomScrollBar
	contextPopupPtr uintptr
	dragDropPtr     uintptr
	dragOverPtr     uintptr
	endDragPtr      uintptr
	startDragPtr    uintptr
}

func NewScrollBar(AOwner IComponent) IScrollBar {
	r1 := LCL().SysCallN(4966, GetObjectUintptr(AOwner))
	return AsScrollBar(r1)
}

func (m *TScrollBar) DragCursor() TCursor {
	r1 := LCL().SysCallN(4967, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TScrollBar) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4967, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBar) DragKind() TDragKind {
	r1 := LCL().SysCallN(4968, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TScrollBar) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4968, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBar) DragMode() TDragMode {
	r1 := LCL().SysCallN(4969, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TScrollBar) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4969, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBar) ParentShowHint() bool {
	r1 := LCL().SysCallN(4970, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBar) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4970, 1, m.Instance(), PascalBool(AValue))
}

func ScrollBarClass() TClass {
	ret := LCL().SysCallN(4965)
	return TClass(ret)
}

func (m *TScrollBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4971, m.Instance(), m.contextPopupPtr)
}

func (m *TScrollBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4972, m.Instance(), m.dragDropPtr)
}

func (m *TScrollBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4973, m.Instance(), m.dragOverPtr)
}

func (m *TScrollBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4974, m.Instance(), m.endDragPtr)
}

func (m *TScrollBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4975, m.Instance(), m.startDragPtr)
}
