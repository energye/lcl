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
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// IMemo Parent: ICustomMemo
type IMemo interface {
	ICustomMemo
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
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
	SetOnEditingDone(fn TNotifyEvent)               // property event
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
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TMemo Parent: TCustomMemo
type TMemo struct {
	TCustomMemo
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	editingDonePtr     uintptr
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
	startDragPtr       uintptr
}

func NewMemo(AOwner IComponent) IMemo {
	r1 := memoImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsMemo(r1)
}

func (m *TMemo) DragCursor() TCursor {
	r1 := memoImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TMemo) SetDragCursor(AValue TCursor) {
	memoImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TMemo) DragKind() TDragKind {
	r1 := memoImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TMemo) SetDragKind(AValue TDragKind) {
	memoImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TMemo) DragMode() TDragMode {
	r1 := memoImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TMemo) SetDragMode(AValue TDragMode) {
	memoImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TMemo) ParentColor() bool {
	r1 := memoImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMemo) SetParentColor(AValue bool) {
	memoImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMemo) ParentFont() bool {
	r1 := memoImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMemo) SetParentFont(AValue bool) {
	memoImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMemo) ParentShowHint() bool {
	r1 := memoImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMemo) SetParentShowHint(AValue bool) {
	memoImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func MemoClass() TClass {
	ret := memoImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TMemo) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(8, m.Instance(), m.contextPopupPtr)
}

func (m *TMemo) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(9, m.Instance(), m.dblClickPtr)
}

func (m *TMemo) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(10, m.Instance(), m.dragDropPtr)
}

func (m *TMemo) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(11, m.Instance(), m.dragOverPtr)
}

func (m *TMemo) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(12, m.Instance(), m.editingDonePtr)
}

func (m *TMemo) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(13, m.Instance(), m.endDragPtr)
}

func (m *TMemo) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(14, m.Instance(), m.mouseDownPtr)
}

func (m *TMemo) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(15, m.Instance(), m.mouseEnterPtr)
}

func (m *TMemo) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(16, m.Instance(), m.mouseLeavePtr)
}

func (m *TMemo) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(17, m.Instance(), m.mouseMovePtr)
}

func (m *TMemo) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(18, m.Instance(), m.mouseUpPtr)
}

func (m *TMemo) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(19, m.Instance(), m.mouseWheelPtr)
}

func (m *TMemo) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(20, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TMemo) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(24, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TMemo) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(21, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TMemo) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(22, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TMemo) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(23, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TMemo) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	memoImportAPI().SysCallN(25, m.Instance(), m.startDragPtr)
}

var (
	memoImport       *imports.Imports = nil
	memoImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Memo_Class", 0),
		/*1*/ imports.NewTable("Memo_Create", 0),
		/*2*/ imports.NewTable("Memo_DragCursor", 0),
		/*3*/ imports.NewTable("Memo_DragKind", 0),
		/*4*/ imports.NewTable("Memo_DragMode", 0),
		/*5*/ imports.NewTable("Memo_ParentColor", 0),
		/*6*/ imports.NewTable("Memo_ParentFont", 0),
		/*7*/ imports.NewTable("Memo_ParentShowHint", 0),
		/*8*/ imports.NewTable("Memo_SetOnContextPopup", 0),
		/*9*/ imports.NewTable("Memo_SetOnDblClick", 0),
		/*10*/ imports.NewTable("Memo_SetOnDragDrop", 0),
		/*11*/ imports.NewTable("Memo_SetOnDragOver", 0),
		/*12*/ imports.NewTable("Memo_SetOnEditingDone", 0),
		/*13*/ imports.NewTable("Memo_SetOnEndDrag", 0),
		/*14*/ imports.NewTable("Memo_SetOnMouseDown", 0),
		/*15*/ imports.NewTable("Memo_SetOnMouseEnter", 0),
		/*16*/ imports.NewTable("Memo_SetOnMouseLeave", 0),
		/*17*/ imports.NewTable("Memo_SetOnMouseMove", 0),
		/*18*/ imports.NewTable("Memo_SetOnMouseUp", 0),
		/*19*/ imports.NewTable("Memo_SetOnMouseWheel", 0),
		/*20*/ imports.NewTable("Memo_SetOnMouseWheelDown", 0),
		/*21*/ imports.NewTable("Memo_SetOnMouseWheelHorz", 0),
		/*22*/ imports.NewTable("Memo_SetOnMouseWheelLeft", 0),
		/*23*/ imports.NewTable("Memo_SetOnMouseWheelRight", 0),
		/*24*/ imports.NewTable("Memo_SetOnMouseWheelUp", 0),
		/*25*/ imports.NewTable("Memo_SetOnStartDrag", 0),
	}
)

func memoImportAPI() *imports.Imports {
	if memoImport == nil {
		memoImport = NewDefaultImports()
		memoImport.SetImportTable(memoImportTables)
		memoImportTables = nil
	}
	return memoImport
}
