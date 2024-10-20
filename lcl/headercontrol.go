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

// IHeaderControl Parent: ICustomHeaderControl
type IHeaderControl interface {
	ICustomHeaderControl
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragKind() TDragKind                            // property
	SetDragKind(AValue TDragKind)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDock(fn TEndDragEvent)                  // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnMouseDown(fn TMouseEvent)                  // property event
	SetOnMouseEnter(fn TNotifyEvent)                // property event
	SetOnMouseLeave(fn TNotifyEvent)                // property event
	SetOnMouseMove(fn TMouseMoveEvent)              // property event
	SetOnMouseUp(fn TMouseEvent)                    // property event
	SetOnMouseWheel(fn TMouseWheelEvent)            // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)    // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)        // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)  // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) // property event
}

// THeaderControl Parent: TCustomHeaderControl
type THeaderControl struct {
	TCustomHeaderControl
	contextPopupPtr    uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDockPtr         uintptr
	endDragPtr         uintptr
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	mouseWheelHorzPtr  uintptr
	mouseWheelLeftPtr  uintptr
	mouseWheelRightPtr uintptr
}

func NewHeaderControl(AOwner IComponent) IHeaderControl {
	r1 := headerControlImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsHeaderControl(r1)
}

func (m *THeaderControl) DragCursor() TCursor {
	r1 := headerControlImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *THeaderControl) SetDragCursor(AValue TCursor) {
	headerControlImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderControl) DragKind() TDragKind {
	r1 := headerControlImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *THeaderControl) SetDragKind(AValue TDragKind) {
	headerControlImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderControl) DragMode() TDragMode {
	r1 := headerControlImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *THeaderControl) SetDragMode(AValue TDragMode) {
	headerControlImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *THeaderControl) ParentFont() bool {
	r1 := headerControlImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *THeaderControl) SetParentFont(AValue bool) {
	headerControlImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *THeaderControl) ParentShowHint() bool {
	r1 := headerControlImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *THeaderControl) SetParentShowHint(AValue bool) {
	headerControlImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func HeaderControlClass() TClass {
	ret := headerControlImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *THeaderControl) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(7, m.Instance(), m.contextPopupPtr)
}

func (m *THeaderControl) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(8, m.Instance(), m.dragDropPtr)
}

func (m *THeaderControl) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(9, m.Instance(), m.dragOverPtr)
}

func (m *THeaderControl) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(10, m.Instance(), m.endDockPtr)
}

func (m *THeaderControl) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(11, m.Instance(), m.endDragPtr)
}

func (m *THeaderControl) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(12, m.Instance(), m.mouseDownPtr)
}

func (m *THeaderControl) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(13, m.Instance(), m.mouseEnterPtr)
}

func (m *THeaderControl) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(14, m.Instance(), m.mouseLeavePtr)
}

func (m *THeaderControl) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(15, m.Instance(), m.mouseMovePtr)
}

func (m *THeaderControl) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(16, m.Instance(), m.mouseUpPtr)
}

func (m *THeaderControl) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(17, m.Instance(), m.mouseWheelPtr)
}

func (m *THeaderControl) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(18, m.Instance(), m.mouseWheelDownPtr)
}

func (m *THeaderControl) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(22, m.Instance(), m.mouseWheelUpPtr)
}

func (m *THeaderControl) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(19, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *THeaderControl) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(20, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *THeaderControl) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	headerControlImportAPI().SysCallN(21, m.Instance(), m.mouseWheelRightPtr)
}

var (
	headerControlImport       *imports.Imports = nil
	headerControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("HeaderControl_Class", 0),
		/*1*/ imports.NewTable("HeaderControl_Create", 0),
		/*2*/ imports.NewTable("HeaderControl_DragCursor", 0),
		/*3*/ imports.NewTable("HeaderControl_DragKind", 0),
		/*4*/ imports.NewTable("HeaderControl_DragMode", 0),
		/*5*/ imports.NewTable("HeaderControl_ParentFont", 0),
		/*6*/ imports.NewTable("HeaderControl_ParentShowHint", 0),
		/*7*/ imports.NewTable("HeaderControl_SetOnContextPopup", 0),
		/*8*/ imports.NewTable("HeaderControl_SetOnDragDrop", 0),
		/*9*/ imports.NewTable("HeaderControl_SetOnDragOver", 0),
		/*10*/ imports.NewTable("HeaderControl_SetOnEndDock", 0),
		/*11*/ imports.NewTable("HeaderControl_SetOnEndDrag", 0),
		/*12*/ imports.NewTable("HeaderControl_SetOnMouseDown", 0),
		/*13*/ imports.NewTable("HeaderControl_SetOnMouseEnter", 0),
		/*14*/ imports.NewTable("HeaderControl_SetOnMouseLeave", 0),
		/*15*/ imports.NewTable("HeaderControl_SetOnMouseMove", 0),
		/*16*/ imports.NewTable("HeaderControl_SetOnMouseUp", 0),
		/*17*/ imports.NewTable("HeaderControl_SetOnMouseWheel", 0),
		/*18*/ imports.NewTable("HeaderControl_SetOnMouseWheelDown", 0),
		/*19*/ imports.NewTable("HeaderControl_SetOnMouseWheelHorz", 0),
		/*20*/ imports.NewTable("HeaderControl_SetOnMouseWheelLeft", 0),
		/*21*/ imports.NewTable("HeaderControl_SetOnMouseWheelRight", 0),
		/*22*/ imports.NewTable("HeaderControl_SetOnMouseWheelUp", 0),
	}
)

func headerControlImportAPI() *imports.Imports {
	if headerControlImport == nil {
		headerControlImport = NewDefaultImports()
		headerControlImport.SetImportTable(headerControlImportTables)
		headerControlImportTables = nil
	}
	return headerControlImport
}
