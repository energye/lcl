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

// ITrackBar Parent: ICustomTrackBar
type ITrackBar interface {
	ICustomTrackBar
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
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
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TTrackBar Parent: TCustomTrackBar
type TTrackBar struct {
	TCustomTrackBar
	contextPopupPtr    uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
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
	startDragPtr       uintptr
}

func NewTrackBar(AOwner IComponent) ITrackBar {
	r1 := rackBarImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsTrackBar(r1)
}

func (m *TTrackBar) DragCursor() TCursor {
	r1 := rackBarImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TTrackBar) SetDragCursor(AValue TCursor) {
	rackBarImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TTrackBar) DragMode() TDragMode {
	r1 := rackBarImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TTrackBar) SetDragMode(AValue TDragMode) {
	rackBarImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TTrackBar) ParentColor() bool {
	r1 := rackBarImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTrackBar) SetParentColor(AValue bool) {
	rackBarImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTrackBar) ParentFont() bool {
	r1 := rackBarImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTrackBar) SetParentFont(AValue bool) {
	rackBarImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TTrackBar) ParentShowHint() bool {
	r1 := rackBarImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TTrackBar) SetParentShowHint(AValue bool) {
	rackBarImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func TrackBarClass() TClass {
	ret := rackBarImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TTrackBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(7, m.Instance(), m.contextPopupPtr)
}

func (m *TTrackBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(8, m.Instance(), m.dragDropPtr)
}

func (m *TTrackBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(9, m.Instance(), m.dragOverPtr)
}

func (m *TTrackBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(10, m.Instance(), m.endDragPtr)
}

func (m *TTrackBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(11, m.Instance(), m.mouseDownPtr)
}

func (m *TTrackBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(12, m.Instance(), m.mouseEnterPtr)
}

func (m *TTrackBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(13, m.Instance(), m.mouseLeavePtr)
}

func (m *TTrackBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(14, m.Instance(), m.mouseMovePtr)
}

func (m *TTrackBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(15, m.Instance(), m.mouseUpPtr)
}

func (m *TTrackBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(16, m.Instance(), m.mouseWheelPtr)
}

func (m *TTrackBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(17, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TTrackBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(21, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TTrackBar) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(18, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TTrackBar) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(19, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TTrackBar) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(20, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TTrackBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	rackBarImportAPI().SysCallN(22, m.Instance(), m.startDragPtr)
}

var (
	rackBarImport       *imports.Imports = nil
	rackBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TrackBar_Class", 0),
		/*1*/ imports.NewTable("TrackBar_Create", 0),
		/*2*/ imports.NewTable("TrackBar_DragCursor", 0),
		/*3*/ imports.NewTable("TrackBar_DragMode", 0),
		/*4*/ imports.NewTable("TrackBar_ParentColor", 0),
		/*5*/ imports.NewTable("TrackBar_ParentFont", 0),
		/*6*/ imports.NewTable("TrackBar_ParentShowHint", 0),
		/*7*/ imports.NewTable("TrackBar_SetOnContextPopup", 0),
		/*8*/ imports.NewTable("TrackBar_SetOnDragDrop", 0),
		/*9*/ imports.NewTable("TrackBar_SetOnDragOver", 0),
		/*10*/ imports.NewTable("TrackBar_SetOnEndDrag", 0),
		/*11*/ imports.NewTable("TrackBar_SetOnMouseDown", 0),
		/*12*/ imports.NewTable("TrackBar_SetOnMouseEnter", 0),
		/*13*/ imports.NewTable("TrackBar_SetOnMouseLeave", 0),
		/*14*/ imports.NewTable("TrackBar_SetOnMouseMove", 0),
		/*15*/ imports.NewTable("TrackBar_SetOnMouseUp", 0),
		/*16*/ imports.NewTable("TrackBar_SetOnMouseWheel", 0),
		/*17*/ imports.NewTable("TrackBar_SetOnMouseWheelDown", 0),
		/*18*/ imports.NewTable("TrackBar_SetOnMouseWheelHorz", 0),
		/*19*/ imports.NewTable("TrackBar_SetOnMouseWheelLeft", 0),
		/*20*/ imports.NewTable("TrackBar_SetOnMouseWheelRight", 0),
		/*21*/ imports.NewTable("TrackBar_SetOnMouseWheelUp", 0),
		/*22*/ imports.NewTable("TrackBar_SetOnStartDrag", 0),
	}
)

func rackBarImportAPI() *imports.Imports {
	if rackBarImport == nil {
		rackBarImport = NewDefaultImports()
		rackBarImport.SetImportTable(rackBarImportTables)
		rackBarImportTables = nil
	}
	return rackBarImport
}
