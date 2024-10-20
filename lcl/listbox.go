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

// IListBox Parent: ICustomListBox
type IListBox interface {
	ICustomListBox
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
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

// TListBox Parent: TCustomListBox
type TListBox struct {
	TCustomListBox
	contextPopupPtr    uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	startDragPtr       uintptr
}

func NewListBox(TheOwner IComponent) IListBox {
	r1 := listBoxImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsListBox(r1)
}

func (m *TListBox) DragCursor() TCursor {
	r1 := listBoxImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TListBox) SetDragCursor(AValue TCursor) {
	listBoxImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TListBox) DragKind() TDragKind {
	r1 := listBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TListBox) SetDragKind(AValue TDragKind) {
	listBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TListBox) DragMode() TDragMode {
	r1 := listBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TListBox) SetDragMode(AValue TDragMode) {
	listBoxImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func ListBoxClass() TClass {
	ret := listBoxImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TListBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	listBoxImportAPI().SysCallN(5, m.Instance(), m.contextPopupPtr)
}

func (m *TListBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	listBoxImportAPI().SysCallN(6, m.Instance(), m.dragDropPtr)
}

func (m *TListBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	listBoxImportAPI().SysCallN(7, m.Instance(), m.dragOverPtr)
}

func (m *TListBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	listBoxImportAPI().SysCallN(8, m.Instance(), m.endDragPtr)
}

func (m *TListBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	listBoxImportAPI().SysCallN(9, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TListBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	listBoxImportAPI().SysCallN(10, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TListBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	listBoxImportAPI().SysCallN(11, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TListBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	listBoxImportAPI().SysCallN(12, m.Instance(), m.startDragPtr)
}

var (
	listBoxImport       *imports.Imports = nil
	listBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ListBox_Class", 0),
		/*1*/ imports.NewTable("ListBox_Create", 0),
		/*2*/ imports.NewTable("ListBox_DragCursor", 0),
		/*3*/ imports.NewTable("ListBox_DragKind", 0),
		/*4*/ imports.NewTable("ListBox_DragMode", 0),
		/*5*/ imports.NewTable("ListBox_SetOnContextPopup", 0),
		/*6*/ imports.NewTable("ListBox_SetOnDragDrop", 0),
		/*7*/ imports.NewTable("ListBox_SetOnDragOver", 0),
		/*8*/ imports.NewTable("ListBox_SetOnEndDrag", 0),
		/*9*/ imports.NewTable("ListBox_SetOnMouseWheelHorz", 0),
		/*10*/ imports.NewTable("ListBox_SetOnMouseWheelLeft", 0),
		/*11*/ imports.NewTable("ListBox_SetOnMouseWheelRight", 0),
		/*12*/ imports.NewTable("ListBox_SetOnStartDrag", 0),
	}
)

func listBoxImportAPI() *imports.Imports {
	if listBoxImport == nil {
		listBoxImport = NewDefaultImports()
		listBoxImport.SetImportTable(listBoxImportTables)
		listBoxImportTables = nil
	}
	return listBoxImport
}
