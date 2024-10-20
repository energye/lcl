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

// ITabSheet Parent: ICustomPage
type ITabSheet interface {
	ICustomPage
	PageControl() IPageControl                     // property
	SetPageControl(AValue IPageControl)            // property
	TabIndex() int32                               // property
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

// TTabSheet Parent: TCustomPage
type TTabSheet struct {
	TCustomPage
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

func NewTabSheet(TheOwner IComponent) ITabSheet {
	r1 := abSheetImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsTabSheet(r1)
}

func (m *TTabSheet) PageControl() IPageControl {
	r1 := abSheetImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return AsPageControl(r1)
}

func (m *TTabSheet) SetPageControl(AValue IPageControl) {
	abSheetImportAPI().SysCallN(2, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TTabSheet) TabIndex() int32 {
	r1 := abSheetImportAPI().SysCallN(18, m.Instance())
	return int32(r1)
}

func (m *TTabSheet) ParentFont() bool {
	r1 := abSheetImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTabSheet) SetParentFont(AValue bool) {
	abSheetImportAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTabSheet) ParentShowHint() bool {
	r1 := abSheetImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTabSheet) SetParentShowHint(AValue bool) {
	abSheetImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func TabSheetClass() TClass {
	ret := abSheetImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TTabSheet) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(5, m.Instance(), m.contextPopupPtr)
}

func (m *TTabSheet) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(6, m.Instance(), m.dragDropPtr)
}

func (m *TTabSheet) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(7, m.Instance(), m.dragOverPtr)
}

func (m *TTabSheet) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(8, m.Instance(), m.endDragPtr)
}

func (m *TTabSheet) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(9, m.Instance(), m.mouseDownPtr)
}

func (m *TTabSheet) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(10, m.Instance(), m.mouseEnterPtr)
}

func (m *TTabSheet) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(11, m.Instance(), m.mouseLeavePtr)
}

func (m *TTabSheet) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(12, m.Instance(), m.mouseMovePtr)
}

func (m *TTabSheet) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(13, m.Instance(), m.mouseUpPtr)
}

func (m *TTabSheet) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(14, m.Instance(), m.mouseWheelPtr)
}

func (m *TTabSheet) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(15, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TTabSheet) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(16, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TTabSheet) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	abSheetImportAPI().SysCallN(17, m.Instance(), m.startDragPtr)
}

var (
	abSheetImport       *imports.Imports = nil
	abSheetImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TabSheet_Class", 0),
		/*1*/ imports.NewTable("TabSheet_Create", 0),
		/*2*/ imports.NewTable("TabSheet_PageControl", 0),
		/*3*/ imports.NewTable("TabSheet_ParentFont", 0),
		/*4*/ imports.NewTable("TabSheet_ParentShowHint", 0),
		/*5*/ imports.NewTable("TabSheet_SetOnContextPopup", 0),
		/*6*/ imports.NewTable("TabSheet_SetOnDragDrop", 0),
		/*7*/ imports.NewTable("TabSheet_SetOnDragOver", 0),
		/*8*/ imports.NewTable("TabSheet_SetOnEndDrag", 0),
		/*9*/ imports.NewTable("TabSheet_SetOnMouseDown", 0),
		/*10*/ imports.NewTable("TabSheet_SetOnMouseEnter", 0),
		/*11*/ imports.NewTable("TabSheet_SetOnMouseLeave", 0),
		/*12*/ imports.NewTable("TabSheet_SetOnMouseMove", 0),
		/*13*/ imports.NewTable("TabSheet_SetOnMouseUp", 0),
		/*14*/ imports.NewTable("TabSheet_SetOnMouseWheel", 0),
		/*15*/ imports.NewTable("TabSheet_SetOnMouseWheelDown", 0),
		/*16*/ imports.NewTable("TabSheet_SetOnMouseWheelUp", 0),
		/*17*/ imports.NewTable("TabSheet_SetOnStartDrag", 0),
		/*18*/ imports.NewTable("TabSheet_TabIndex", 0),
	}
)

func abSheetImportAPI() *imports.Imports {
	if abSheetImport == nil {
		abSheetImport = NewDefaultImports()
		abSheetImport.SetImportTable(abSheetImportTables)
		abSheetImportTables = nil
	}
	return abSheetImport
}
