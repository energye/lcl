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
	r1 := scrollBarImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsScrollBar(r1)
}

func (m *TScrollBar) DragCursor() TCursor {
	r1 := scrollBarImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TScrollBar) SetDragCursor(AValue TCursor) {
	scrollBarImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBar) DragKind() TDragKind {
	r1 := scrollBarImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TScrollBar) SetDragKind(AValue TDragKind) {
	scrollBarImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBar) DragMode() TDragMode {
	r1 := scrollBarImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TScrollBar) SetDragMode(AValue TDragMode) {
	scrollBarImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBar) ParentShowHint() bool {
	r1 := scrollBarImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBar) SetParentShowHint(AValue bool) {
	scrollBarImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func ScrollBarClass() TClass {
	ret := scrollBarImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TScrollBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	scrollBarImportAPI().SysCallN(6, m.Instance(), m.contextPopupPtr)
}

func (m *TScrollBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	scrollBarImportAPI().SysCallN(7, m.Instance(), m.dragDropPtr)
}

func (m *TScrollBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	scrollBarImportAPI().SysCallN(8, m.Instance(), m.dragOverPtr)
}

func (m *TScrollBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	scrollBarImportAPI().SysCallN(9, m.Instance(), m.endDragPtr)
}

func (m *TScrollBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	scrollBarImportAPI().SysCallN(10, m.Instance(), m.startDragPtr)
}

var (
	scrollBarImport       *imports.Imports = nil
	scrollBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ScrollBar_Class", 0),
		/*1*/ imports.NewTable("ScrollBar_Create", 0),
		/*2*/ imports.NewTable("ScrollBar_DragCursor", 0),
		/*3*/ imports.NewTable("ScrollBar_DragKind", 0),
		/*4*/ imports.NewTable("ScrollBar_DragMode", 0),
		/*5*/ imports.NewTable("ScrollBar_ParentShowHint", 0),
		/*6*/ imports.NewTable("ScrollBar_SetOnContextPopup", 0),
		/*7*/ imports.NewTable("ScrollBar_SetOnDragDrop", 0),
		/*8*/ imports.NewTable("ScrollBar_SetOnDragOver", 0),
		/*9*/ imports.NewTable("ScrollBar_SetOnEndDrag", 0),
		/*10*/ imports.NewTable("ScrollBar_SetOnStartDrag", 0),
	}
)

func scrollBarImportAPI() *imports.Imports {
	if scrollBarImport == nil {
		scrollBarImport = NewDefaultImports()
		scrollBarImport.SetImportTable(scrollBarImportTables)
		scrollBarImportTables = nil
	}
	return scrollBarImport
}
