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

// IColorListBox Parent: ICustomColorListBox
type IColorListBox interface {
	ICustomColorListBox
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

// TColorListBox Parent: TCustomColorListBox
type TColorListBox struct {
	TCustomColorListBox
	contextPopupPtr    uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	startDragPtr       uintptr
}

func NewColorListBox(AOwner IComponent) IColorListBox {
	r1 := colorListBoxImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsColorListBox(r1)
}

func (m *TColorListBox) DragCursor() TCursor {
	r1 := colorListBoxImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TColorListBox) SetDragCursor(AValue TCursor) {
	colorListBoxImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorListBox) DragKind() TDragKind {
	r1 := colorListBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TColorListBox) SetDragKind(AValue TDragKind) {
	colorListBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorListBox) DragMode() TDragMode {
	r1 := colorListBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TColorListBox) SetDragMode(AValue TDragMode) {
	colorListBoxImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func ColorListBoxClass() TClass {
	ret := colorListBoxImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TColorListBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	colorListBoxImportAPI().SysCallN(5, m.Instance(), m.contextPopupPtr)
}

func (m *TColorListBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	colorListBoxImportAPI().SysCallN(6, m.Instance(), m.dragDropPtr)
}

func (m *TColorListBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	colorListBoxImportAPI().SysCallN(7, m.Instance(), m.dragOverPtr)
}

func (m *TColorListBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	colorListBoxImportAPI().SysCallN(8, m.Instance(), m.endDragPtr)
}

func (m *TColorListBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	colorListBoxImportAPI().SysCallN(9, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TColorListBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	colorListBoxImportAPI().SysCallN(10, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TColorListBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	colorListBoxImportAPI().SysCallN(11, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TColorListBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	colorListBoxImportAPI().SysCallN(12, m.Instance(), m.startDragPtr)
}

var (
	colorListBoxImport       *imports.Imports = nil
	colorListBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ColorListBox_Class", 0),
		/*1*/ imports.NewTable("ColorListBox_Create", 0),
		/*2*/ imports.NewTable("ColorListBox_DragCursor", 0),
		/*3*/ imports.NewTable("ColorListBox_DragKind", 0),
		/*4*/ imports.NewTable("ColorListBox_DragMode", 0),
		/*5*/ imports.NewTable("ColorListBox_SetOnContextPopup", 0),
		/*6*/ imports.NewTable("ColorListBox_SetOnDragDrop", 0),
		/*7*/ imports.NewTable("ColorListBox_SetOnDragOver", 0),
		/*8*/ imports.NewTable("ColorListBox_SetOnEndDrag", 0),
		/*9*/ imports.NewTable("ColorListBox_SetOnMouseWheelHorz", 0),
		/*10*/ imports.NewTable("ColorListBox_SetOnMouseWheelLeft", 0),
		/*11*/ imports.NewTable("ColorListBox_SetOnMouseWheelRight", 0),
		/*12*/ imports.NewTable("ColorListBox_SetOnStartDrag", 0),
	}
)

func colorListBoxImportAPI() *imports.Imports {
	if colorListBoxImport == nil {
		colorListBoxImport = NewDefaultImports()
		colorListBoxImport.SetImportTable(colorListBoxImportTables)
		colorListBoxImportTables = nil
	}
	return colorListBoxImport
}
