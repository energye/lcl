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

// IBitBtn Parent: ICustomBitBtn
type IBitBtn interface {
	ICustomBitBtn
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragKind() TDragKind                           // property
	SetDragKind(AValue TDragKind)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
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

// TBitBtn Parent: TCustomBitBtn
type TBitBtn struct {
	TCustomBitBtn
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

func NewBitBtn(TheOwner IComponent) IBitBtn {
	r1 := bitBtnImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsBitBtn(r1)
}

func (m *TBitBtn) DragCursor() TCursor {
	r1 := bitBtnImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TBitBtn) SetDragCursor(AValue TCursor) {
	bitBtnImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TBitBtn) DragKind() TDragKind {
	r1 := bitBtnImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TBitBtn) SetDragKind(AValue TDragKind) {
	bitBtnImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TBitBtn) DragMode() TDragMode {
	r1 := bitBtnImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TBitBtn) SetDragMode(AValue TDragMode) {
	bitBtnImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TBitBtn) ParentFont() bool {
	r1 := bitBtnImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBitBtn) SetParentFont(AValue bool) {
	bitBtnImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBitBtn) ParentShowHint() bool {
	r1 := bitBtnImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBitBtn) SetParentShowHint(AValue bool) {
	bitBtnImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func BitBtnClass() TClass {
	ret := bitBtnImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TBitBtn) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(7, m.Instance(), m.contextPopupPtr)
}

func (m *TBitBtn) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(8, m.Instance(), m.dragDropPtr)
}

func (m *TBitBtn) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(9, m.Instance(), m.dragOverPtr)
}

func (m *TBitBtn) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(10, m.Instance(), m.endDragPtr)
}

func (m *TBitBtn) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(11, m.Instance(), m.mouseDownPtr)
}

func (m *TBitBtn) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(12, m.Instance(), m.mouseEnterPtr)
}

func (m *TBitBtn) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(13, m.Instance(), m.mouseLeavePtr)
}

func (m *TBitBtn) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(14, m.Instance(), m.mouseMovePtr)
}

func (m *TBitBtn) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(15, m.Instance(), m.mouseUpPtr)
}

func (m *TBitBtn) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(16, m.Instance(), m.mouseWheelPtr)
}

func (m *TBitBtn) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(17, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TBitBtn) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(18, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TBitBtn) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	bitBtnImportAPI().SysCallN(19, m.Instance(), m.startDragPtr)
}

var (
	bitBtnImport       *imports.Imports = nil
	bitBtnImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("BitBtn_Class", 0),
		/*1*/ imports.NewTable("BitBtn_Create", 0),
		/*2*/ imports.NewTable("BitBtn_DragCursor", 0),
		/*3*/ imports.NewTable("BitBtn_DragKind", 0),
		/*4*/ imports.NewTable("BitBtn_DragMode", 0),
		/*5*/ imports.NewTable("BitBtn_ParentFont", 0),
		/*6*/ imports.NewTable("BitBtn_ParentShowHint", 0),
		/*7*/ imports.NewTable("BitBtn_SetOnContextPopup", 0),
		/*8*/ imports.NewTable("BitBtn_SetOnDragDrop", 0),
		/*9*/ imports.NewTable("BitBtn_SetOnDragOver", 0),
		/*10*/ imports.NewTable("BitBtn_SetOnEndDrag", 0),
		/*11*/ imports.NewTable("BitBtn_SetOnMouseDown", 0),
		/*12*/ imports.NewTable("BitBtn_SetOnMouseEnter", 0),
		/*13*/ imports.NewTable("BitBtn_SetOnMouseLeave", 0),
		/*14*/ imports.NewTable("BitBtn_SetOnMouseMove", 0),
		/*15*/ imports.NewTable("BitBtn_SetOnMouseUp", 0),
		/*16*/ imports.NewTable("BitBtn_SetOnMouseWheel", 0),
		/*17*/ imports.NewTable("BitBtn_SetOnMouseWheelDown", 0),
		/*18*/ imports.NewTable("BitBtn_SetOnMouseWheelUp", 0),
		/*19*/ imports.NewTable("BitBtn_SetOnStartDrag", 0),
	}
)

func bitBtnImportAPI() *imports.Imports {
	if bitBtnImport == nil {
		bitBtnImport = NewDefaultImports()
		bitBtnImport.SetImportTable(bitBtnImportTables)
		bitBtnImportTables = nil
	}
	return bitBtnImport
}
