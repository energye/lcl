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

// IToggleBox Parent: ICustomCheckBox
type IToggleBox interface {
	ICustomCheckBox
	Checked() bool                                 // property
	SetChecked(AValue bool)                        // property
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

// TToggleBox Parent: TCustomCheckBox
type TToggleBox struct {
	TCustomCheckBox
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

func NewToggleBox(TheOwner IComponent) IToggleBox {
	r1 := oggleBoxImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsToggleBox(r1)
}

func (m *TToggleBox) Checked() bool {
	r1 := oggleBoxImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToggleBox) SetChecked(AValue bool) {
	oggleBoxImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToggleBox) DragCursor() TCursor {
	r1 := oggleBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TToggleBox) SetDragCursor(AValue TCursor) {
	oggleBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TToggleBox) DragKind() TDragKind {
	r1 := oggleBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TToggleBox) SetDragKind(AValue TDragKind) {
	oggleBoxImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TToggleBox) DragMode() TDragMode {
	r1 := oggleBoxImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TToggleBox) SetDragMode(AValue TDragMode) {
	oggleBoxImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TToggleBox) ParentFont() bool {
	r1 := oggleBoxImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToggleBox) SetParentFont(AValue bool) {
	oggleBoxImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TToggleBox) ParentShowHint() bool {
	r1 := oggleBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TToggleBox) SetParentShowHint(AValue bool) {
	oggleBoxImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func ToggleBoxClass() TClass {
	ret := oggleBoxImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TToggleBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(8, m.Instance(), m.contextPopupPtr)
}

func (m *TToggleBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(9, m.Instance(), m.dragDropPtr)
}

func (m *TToggleBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(10, m.Instance(), m.dragOverPtr)
}

func (m *TToggleBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(11, m.Instance(), m.endDragPtr)
}

func (m *TToggleBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(12, m.Instance(), m.mouseDownPtr)
}

func (m *TToggleBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(13, m.Instance(), m.mouseEnterPtr)
}

func (m *TToggleBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(14, m.Instance(), m.mouseLeavePtr)
}

func (m *TToggleBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(15, m.Instance(), m.mouseMovePtr)
}

func (m *TToggleBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(16, m.Instance(), m.mouseUpPtr)
}

func (m *TToggleBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(17, m.Instance(), m.mouseWheelPtr)
}

func (m *TToggleBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(18, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TToggleBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(19, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TToggleBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	oggleBoxImportAPI().SysCallN(20, m.Instance(), m.startDragPtr)
}

var (
	oggleBoxImport       *imports.Imports = nil
	oggleBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ToggleBox_Checked", 0),
		/*1*/ imports.NewTable("ToggleBox_Class", 0),
		/*2*/ imports.NewTable("ToggleBox_Create", 0),
		/*3*/ imports.NewTable("ToggleBox_DragCursor", 0),
		/*4*/ imports.NewTable("ToggleBox_DragKind", 0),
		/*5*/ imports.NewTable("ToggleBox_DragMode", 0),
		/*6*/ imports.NewTable("ToggleBox_ParentFont", 0),
		/*7*/ imports.NewTable("ToggleBox_ParentShowHint", 0),
		/*8*/ imports.NewTable("ToggleBox_SetOnContextPopup", 0),
		/*9*/ imports.NewTable("ToggleBox_SetOnDragDrop", 0),
		/*10*/ imports.NewTable("ToggleBox_SetOnDragOver", 0),
		/*11*/ imports.NewTable("ToggleBox_SetOnEndDrag", 0),
		/*12*/ imports.NewTable("ToggleBox_SetOnMouseDown", 0),
		/*13*/ imports.NewTable("ToggleBox_SetOnMouseEnter", 0),
		/*14*/ imports.NewTable("ToggleBox_SetOnMouseLeave", 0),
		/*15*/ imports.NewTable("ToggleBox_SetOnMouseMove", 0),
		/*16*/ imports.NewTable("ToggleBox_SetOnMouseUp", 0),
		/*17*/ imports.NewTable("ToggleBox_SetOnMouseWheel", 0),
		/*18*/ imports.NewTable("ToggleBox_SetOnMouseWheelDown", 0),
		/*19*/ imports.NewTable("ToggleBox_SetOnMouseWheelUp", 0),
		/*20*/ imports.NewTable("ToggleBox_SetOnStartDrag", 0),
	}
)

func oggleBoxImportAPI() *imports.Imports {
	if oggleBoxImport == nil {
		oggleBoxImport = NewDefaultImports()
		oggleBoxImport.SetImportTable(oggleBoxImportTables)
		oggleBoxImportTables = nil
	}
	return oggleBoxImport
}
