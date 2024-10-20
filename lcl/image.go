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

// IImage Parent: ICustomImage
type IImage interface {
	ICustomImage
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
	SetOnPaint(fn TNotifyEvent)                     // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TImage Parent: TCustomImage
type TImage struct {
	TCustomImage
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
	paintPtr           uintptr
	startDragPtr       uintptr
}

func NewImage(AOwner IComponent) IImage {
	r1 := mageImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsImage(r1)
}

func (m *TImage) DragCursor() TCursor {
	r1 := mageImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TImage) SetDragCursor(AValue TCursor) {
	mageImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TImage) DragMode() TDragMode {
	r1 := mageImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TImage) SetDragMode(AValue TDragMode) {
	mageImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TImage) ParentShowHint() bool {
	r1 := mageImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TImage) SetParentShowHint(AValue bool) {
	mageImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func ImageClass() TClass {
	ret := mageImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TImage) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	mageImportAPI().SysCallN(5, m.Instance(), m.contextPopupPtr)
}

func (m *TImage) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	mageImportAPI().SysCallN(6, m.Instance(), m.dblClickPtr)
}

func (m *TImage) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	mageImportAPI().SysCallN(7, m.Instance(), m.dragDropPtr)
}

func (m *TImage) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	mageImportAPI().SysCallN(8, m.Instance(), m.dragOverPtr)
}

func (m *TImage) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	mageImportAPI().SysCallN(9, m.Instance(), m.endDragPtr)
}

func (m *TImage) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	mageImportAPI().SysCallN(10, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TImage) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	mageImportAPI().SysCallN(11, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TImage) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	mageImportAPI().SysCallN(12, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TImage) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	mageImportAPI().SysCallN(13, m.Instance(), m.paintPtr)
}

func (m *TImage) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	mageImportAPI().SysCallN(14, m.Instance(), m.startDragPtr)
}

var (
	mageImport       *imports.Imports = nil
	mageImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Image_Class", 0),
		/*1*/ imports.NewTable("Image_Create", 0),
		/*2*/ imports.NewTable("Image_DragCursor", 0),
		/*3*/ imports.NewTable("Image_DragMode", 0),
		/*4*/ imports.NewTable("Image_ParentShowHint", 0),
		/*5*/ imports.NewTable("Image_SetOnContextPopup", 0),
		/*6*/ imports.NewTable("Image_SetOnDblClick", 0),
		/*7*/ imports.NewTable("Image_SetOnDragDrop", 0),
		/*8*/ imports.NewTable("Image_SetOnDragOver", 0),
		/*9*/ imports.NewTable("Image_SetOnEndDrag", 0),
		/*10*/ imports.NewTable("Image_SetOnMouseWheelHorz", 0),
		/*11*/ imports.NewTable("Image_SetOnMouseWheelLeft", 0),
		/*12*/ imports.NewTable("Image_SetOnMouseWheelRight", 0),
		/*13*/ imports.NewTable("Image_SetOnPaint", 0),
		/*14*/ imports.NewTable("Image_SetOnStartDrag", 0),
	}
)

func mageImportAPI() *imports.Imports {
	if mageImport == nil {
		mageImport = NewDefaultImports()
		mageImport.SetImportTable(mageImportTables)
		mageImportTables = nil
	}
	return mageImport
}
