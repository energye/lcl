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

// ICheckListBox Parent: ICustomCheckListBox
type ICheckListBox interface {
	ICustomCheckListBox
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TCheckListBox Parent: TCustomCheckListBox
type TCheckListBox struct {
	TCustomCheckListBox
	contextPopupPtr    uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	startDragPtr       uintptr
}

func NewCheckListBox(AOwner IComponent) ICheckListBox {
	r1 := checkListBoxImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCheckListBox(r1)
}

func (m *TCheckListBox) DragCursor() TCursor {
	r1 := checkListBoxImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCheckListBox) SetDragCursor(AValue TCursor) {
	checkListBoxImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckListBox) DragMode() TDragMode {
	r1 := checkListBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCheckListBox) SetDragMode(AValue TDragMode) {
	checkListBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func CheckListBoxClass() TClass {
	ret := checkListBoxImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCheckListBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	checkListBoxImportAPI().SysCallN(4, m.Instance(), m.contextPopupPtr)
}

func (m *TCheckListBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	checkListBoxImportAPI().SysCallN(5, m.Instance(), m.dragDropPtr)
}

func (m *TCheckListBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	checkListBoxImportAPI().SysCallN(6, m.Instance(), m.dragOverPtr)
}

func (m *TCheckListBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	checkListBoxImportAPI().SysCallN(7, m.Instance(), m.endDragPtr)
}

func (m *TCheckListBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	checkListBoxImportAPI().SysCallN(8, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TCheckListBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	checkListBoxImportAPI().SysCallN(9, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TCheckListBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	checkListBoxImportAPI().SysCallN(10, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TCheckListBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	checkListBoxImportAPI().SysCallN(11, m.Instance(), m.startDragPtr)
}

var (
	checkListBoxImport       *imports.Imports = nil
	checkListBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CheckListBox_Class", 0),
		/*1*/ imports.NewTable("CheckListBox_Create", 0),
		/*2*/ imports.NewTable("CheckListBox_DragCursor", 0),
		/*3*/ imports.NewTable("CheckListBox_DragMode", 0),
		/*4*/ imports.NewTable("CheckListBox_SetOnContextPopup", 0),
		/*5*/ imports.NewTable("CheckListBox_SetOnDragDrop", 0),
		/*6*/ imports.NewTable("CheckListBox_SetOnDragOver", 0),
		/*7*/ imports.NewTable("CheckListBox_SetOnEndDrag", 0),
		/*8*/ imports.NewTable("CheckListBox_SetOnMouseWheelHorz", 0),
		/*9*/ imports.NewTable("CheckListBox_SetOnMouseWheelLeft", 0),
		/*10*/ imports.NewTable("CheckListBox_SetOnMouseWheelRight", 0),
		/*11*/ imports.NewTable("CheckListBox_SetOnStartDrag", 0),
	}
)

func checkListBoxImportAPI() *imports.Imports {
	if checkListBoxImport == nil {
		checkListBoxImport = NewDefaultImports()
		checkListBoxImport.SetImportTable(checkListBoxImportTables)
		checkListBoxImportTables = nil
	}
	return checkListBoxImport
}
