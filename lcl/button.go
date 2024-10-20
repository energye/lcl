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

// IButton Parent: ICustomButton
type IButton interface {
	ICustomButton
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

// TButton Parent: TCustomButton
type TButton struct {
	TCustomButton
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

func NewButton(TheOwner IComponent) IButton {
	r1 := buttonImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsButton(r1)
}

func (m *TButton) DragCursor() TCursor {
	r1 := buttonImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TButton) SetDragCursor(AValue TCursor) {
	buttonImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TButton) DragKind() TDragKind {
	r1 := buttonImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TButton) SetDragKind(AValue TDragKind) {
	buttonImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TButton) DragMode() TDragMode {
	r1 := buttonImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TButton) SetDragMode(AValue TDragMode) {
	buttonImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TButton) ParentFont() bool {
	r1 := buttonImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TButton) SetParentFont(AValue bool) {
	buttonImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TButton) ParentShowHint() bool {
	r1 := buttonImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TButton) SetParentShowHint(AValue bool) {
	buttonImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func ButtonClass() TClass {
	ret := buttonImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TButton) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(7, m.Instance(), m.contextPopupPtr)
}

func (m *TButton) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(8, m.Instance(), m.dragDropPtr)
}

func (m *TButton) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(9, m.Instance(), m.dragOverPtr)
}

func (m *TButton) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(10, m.Instance(), m.endDragPtr)
}

func (m *TButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(11, m.Instance(), m.mouseDownPtr)
}

func (m *TButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(12, m.Instance(), m.mouseEnterPtr)
}

func (m *TButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(13, m.Instance(), m.mouseLeavePtr)
}

func (m *TButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(14, m.Instance(), m.mouseMovePtr)
}

func (m *TButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(15, m.Instance(), m.mouseUpPtr)
}

func (m *TButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(16, m.Instance(), m.mouseWheelPtr)
}

func (m *TButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(17, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(18, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TButton) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	buttonImportAPI().SysCallN(19, m.Instance(), m.startDragPtr)
}

var (
	buttonImport       *imports.Imports = nil
	buttonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Button_Class", 0),
		/*1*/ imports.NewTable("Button_Create", 0),
		/*2*/ imports.NewTable("Button_DragCursor", 0),
		/*3*/ imports.NewTable("Button_DragKind", 0),
		/*4*/ imports.NewTable("Button_DragMode", 0),
		/*5*/ imports.NewTable("Button_ParentFont", 0),
		/*6*/ imports.NewTable("Button_ParentShowHint", 0),
		/*7*/ imports.NewTable("Button_SetOnContextPopup", 0),
		/*8*/ imports.NewTable("Button_SetOnDragDrop", 0),
		/*9*/ imports.NewTable("Button_SetOnDragOver", 0),
		/*10*/ imports.NewTable("Button_SetOnEndDrag", 0),
		/*11*/ imports.NewTable("Button_SetOnMouseDown", 0),
		/*12*/ imports.NewTable("Button_SetOnMouseEnter", 0),
		/*13*/ imports.NewTable("Button_SetOnMouseLeave", 0),
		/*14*/ imports.NewTable("Button_SetOnMouseMove", 0),
		/*15*/ imports.NewTable("Button_SetOnMouseUp", 0),
		/*16*/ imports.NewTable("Button_SetOnMouseWheel", 0),
		/*17*/ imports.NewTable("Button_SetOnMouseWheelDown", 0),
		/*18*/ imports.NewTable("Button_SetOnMouseWheelUp", 0),
		/*19*/ imports.NewTable("Button_SetOnStartDrag", 0),
	}
)

func buttonImportAPI() *imports.Imports {
	if buttonImport == nil {
		buttonImport = NewDefaultImports()
		buttonImport.SetImportTable(buttonImportTables)
		buttonImportTables = nil
	}
	return buttonImport
}
