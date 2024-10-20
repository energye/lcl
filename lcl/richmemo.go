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

// IRichMemo Parent: ICustomRichMemo
type IRichMemo interface {
	ICustomRichMemo
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragKind() TDragKind                           // property
	SetDragKind(AValue TDragKind)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	Rtf() string                                   // property
	SetRtf(AValue string)                          // property
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
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

// TRichMemo Parent: TCustomRichMemo
type TRichMemo struct {
	TCustomRichMemo
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	editingDonePtr    uintptr
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

func NewRichMemo(AOwner IComponent) IRichMemo {
	r1 := richMemoImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsRichMemo(r1)
}

func (m *TRichMemo) DragCursor() TCursor {
	r1 := richMemoImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TRichMemo) SetDragCursor(AValue TCursor) {
	richMemoImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TRichMemo) DragKind() TDragKind {
	r1 := richMemoImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TRichMemo) SetDragKind(AValue TDragKind) {
	richMemoImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TRichMemo) DragMode() TDragMode {
	r1 := richMemoImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TRichMemo) SetDragMode(AValue TDragMode) {
	richMemoImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TRichMemo) ParentColor() bool {
	r1 := richMemoImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRichMemo) SetParentColor(AValue bool) {
	richMemoImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRichMemo) ParentFont() bool {
	r1 := richMemoImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRichMemo) SetParentFont(AValue bool) {
	richMemoImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRichMemo) ParentShowHint() bool {
	r1 := richMemoImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TRichMemo) SetParentShowHint(AValue bool) {
	richMemoImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TRichMemo) Rtf() string {
	r1 := richMemoImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TRichMemo) SetRtf(AValue string) {
	richMemoImportAPI().SysCallN(8, 1, m.Instance(), PascalStr(AValue))
}

func RichMemoClass() TClass {
	ret := richMemoImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TRichMemo) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(9, m.Instance(), m.contextPopupPtr)
}

func (m *TRichMemo) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(10, m.Instance(), m.dblClickPtr)
}

func (m *TRichMemo) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(11, m.Instance(), m.dragDropPtr)
}

func (m *TRichMemo) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(12, m.Instance(), m.dragOverPtr)
}

func (m *TRichMemo) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(13, m.Instance(), m.editingDonePtr)
}

func (m *TRichMemo) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(14, m.Instance(), m.endDragPtr)
}

func (m *TRichMemo) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(15, m.Instance(), m.mouseDownPtr)
}

func (m *TRichMemo) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(16, m.Instance(), m.mouseEnterPtr)
}

func (m *TRichMemo) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(17, m.Instance(), m.mouseLeavePtr)
}

func (m *TRichMemo) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(18, m.Instance(), m.mouseMovePtr)
}

func (m *TRichMemo) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(19, m.Instance(), m.mouseUpPtr)
}

func (m *TRichMemo) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(20, m.Instance(), m.mouseWheelPtr)
}

func (m *TRichMemo) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(21, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TRichMemo) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(22, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TRichMemo) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	richMemoImportAPI().SysCallN(23, m.Instance(), m.startDragPtr)
}

var (
	richMemoImport       *imports.Imports = nil
	richMemoImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("RichMemo_Class", 0),
		/*1*/ imports.NewTable("RichMemo_Create", 0),
		/*2*/ imports.NewTable("RichMemo_DragCursor", 0),
		/*3*/ imports.NewTable("RichMemo_DragKind", 0),
		/*4*/ imports.NewTable("RichMemo_DragMode", 0),
		/*5*/ imports.NewTable("RichMemo_ParentColor", 0),
		/*6*/ imports.NewTable("RichMemo_ParentFont", 0),
		/*7*/ imports.NewTable("RichMemo_ParentShowHint", 0),
		/*8*/ imports.NewTable("RichMemo_Rtf", 0),
		/*9*/ imports.NewTable("RichMemo_SetOnContextPopup", 0),
		/*10*/ imports.NewTable("RichMemo_SetOnDblClick", 0),
		/*11*/ imports.NewTable("RichMemo_SetOnDragDrop", 0),
		/*12*/ imports.NewTable("RichMemo_SetOnDragOver", 0),
		/*13*/ imports.NewTable("RichMemo_SetOnEditingDone", 0),
		/*14*/ imports.NewTable("RichMemo_SetOnEndDrag", 0),
		/*15*/ imports.NewTable("RichMemo_SetOnMouseDown", 0),
		/*16*/ imports.NewTable("RichMemo_SetOnMouseEnter", 0),
		/*17*/ imports.NewTable("RichMemo_SetOnMouseLeave", 0),
		/*18*/ imports.NewTable("RichMemo_SetOnMouseMove", 0),
		/*19*/ imports.NewTable("RichMemo_SetOnMouseUp", 0),
		/*20*/ imports.NewTable("RichMemo_SetOnMouseWheel", 0),
		/*21*/ imports.NewTable("RichMemo_SetOnMouseWheelDown", 0),
		/*22*/ imports.NewTable("RichMemo_SetOnMouseWheelUp", 0),
		/*23*/ imports.NewTable("RichMemo_SetOnStartDrag", 0),
	}
)

func richMemoImportAPI() *imports.Imports {
	if richMemoImport == nil {
		richMemoImport = NewDefaultImports()
		richMemoImport.SetImportTable(richMemoImportTables)
		richMemoImportTables = nil
	}
	return richMemoImport
}
