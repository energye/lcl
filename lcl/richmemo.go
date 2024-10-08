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

// IRichMemo Parent: ICustomRichMemo
type IRichMemo interface {
	ICustomRichMemo
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
	Rtf() string                                   // property
	SetRtf(AValue string)                          // property
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
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

// TRichMemo Parent: TCustomRichMemo
type TRichMemo struct {
	TCustomRichMemo
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
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

func NewRichMemo(AOwner IComponent) IRichMemo {
	r1 := LCL().SysCallN(4865, GetObjectUintptr(AOwner))
	return AsRichMemo(r1)
}

func (m *TRichMemo) DragCursor() TCursor {
	r1 := LCL().SysCallN(4866, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TRichMemo) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(4866, 1, m.Instance(), uintptr(AValue))
}

func (m *TRichMemo) DragKind() TDragKind {
	r1 := LCL().SysCallN(4867, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TRichMemo) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(4867, 1, m.Instance(), uintptr(AValue))
}

func (m *TRichMemo) DragMode() TDragMode {
	r1 := LCL().SysCallN(4868, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TRichMemo) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(4868, 1, m.Instance(), uintptr(AValue))
}

func (m *TRichMemo) ParentColor() bool {
	r1 := LCL().SysCallN(4869, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRichMemo) SetParentColor(AValue bool) {
	LCL().SysCallN(4869, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRichMemo) ParentFont() bool {
	r1 := LCL().SysCallN(4870, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRichMemo) SetParentFont(AValue bool) {
	LCL().SysCallN(4870, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRichMemo) ParentShowHint() bool {
	r1 := LCL().SysCallN(4871, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRichMemo) SetParentShowHint(AValue bool) {
	LCL().SysCallN(4871, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRichMemo) Rtf() string {
	r1 := LCL().SysCallN(4872, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TRichMemo) SetRtf(AValue string) {
	LCL().SysCallN(4872, 1, m.Instance(), PascalStr(AValue))
}

func RichMemoClass() TClass {
	ret := LCL().SysCallN(4864)
	return TClass(ret)
}

func (m *TRichMemo) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4873, m.Instance(), m.contextPopupPtr)
}

func (m *TRichMemo) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4874, m.Instance(), m.dblClickPtr)
}

func (m *TRichMemo) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4875, m.Instance(), m.dragDropPtr)
}

func (m *TRichMemo) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4876, m.Instance(), m.dragOverPtr)
}

func (m *TRichMemo) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4877, m.Instance(), m.editingDonePtr)
}

func (m *TRichMemo) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4878, m.Instance(), m.endDragPtr)
}

func (m *TRichMemo) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4879, m.Instance(), m.mouseDownPtr)
}

func (m *TRichMemo) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4880, m.Instance(), m.mouseEnterPtr)
}

func (m *TRichMemo) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4881, m.Instance(), m.mouseLeavePtr)
}

func (m *TRichMemo) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4882, m.Instance(), m.mouseMovePtr)
}

func (m *TRichMemo) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4883, m.Instance(), m.mouseUpPtr)
}

func (m *TRichMemo) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4884, m.Instance(), m.mouseWheelPtr)
}

func (m *TRichMemo) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4885, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TRichMemo) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4886, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TRichMemo) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(4887, m.Instance(), m.startDragPtr)
}
