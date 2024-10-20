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

// IGroupBox Parent: ICustomGroupBox
type IGroupBox interface {
	ICustomGroupBox
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
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDock(fn TEndDragEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)         // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TGroupBox Parent: TCustomGroupBox
type TGroupBox struct {
	TCustomGroupBox
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDockPtr        uintptr
	endDragPtr        uintptr
	getSiteInfoPtr    uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDockPtr      uintptr
	startDragPtr      uintptr
}

func NewGroupBox(AOwner IComponent) IGroupBox {
	r1 := groupBoxImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsGroupBox(r1)
}

func (m *TGroupBox) DragCursor() TCursor {
	r1 := groupBoxImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TGroupBox) SetDragCursor(AValue TCursor) {
	groupBoxImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TGroupBox) DragKind() TDragKind {
	r1 := groupBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TGroupBox) SetDragKind(AValue TDragKind) {
	groupBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TGroupBox) DragMode() TDragMode {
	r1 := groupBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TGroupBox) SetDragMode(AValue TDragMode) {
	groupBoxImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TGroupBox) ParentColor() bool {
	r1 := groupBoxImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGroupBox) SetParentColor(AValue bool) {
	groupBoxImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGroupBox) ParentFont() bool {
	r1 := groupBoxImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGroupBox) SetParentFont(AValue bool) {
	groupBoxImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TGroupBox) ParentShowHint() bool {
	r1 := groupBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TGroupBox) SetParentShowHint(AValue bool) {
	groupBoxImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func GroupBoxClass() TClass {
	ret := groupBoxImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TGroupBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(8, m.Instance(), m.contextPopupPtr)
}

func (m *TGroupBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(9, m.Instance(), m.dblClickPtr)
}

func (m *TGroupBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(10, m.Instance(), m.dragDropPtr)
}

func (m *TGroupBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(11, m.Instance(), m.dragOverPtr)
}

func (m *TGroupBox) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(12, m.Instance(), m.endDockPtr)
}

func (m *TGroupBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(13, m.Instance(), m.endDragPtr)
}

func (m *TGroupBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(14, m.Instance(), m.getSiteInfoPtr)
}

func (m *TGroupBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(15, m.Instance(), m.mouseDownPtr)
}

func (m *TGroupBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(16, m.Instance(), m.mouseEnterPtr)
}

func (m *TGroupBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(17, m.Instance(), m.mouseLeavePtr)
}

func (m *TGroupBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(18, m.Instance(), m.mouseMovePtr)
}

func (m *TGroupBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(19, m.Instance(), m.mouseUpPtr)
}

func (m *TGroupBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(20, m.Instance(), m.mouseWheelPtr)
}

func (m *TGroupBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(21, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TGroupBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(22, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TGroupBox) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(23, m.Instance(), m.startDockPtr)
}

func (m *TGroupBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	groupBoxImportAPI().SysCallN(24, m.Instance(), m.startDragPtr)
}

var (
	groupBoxImport       *imports.Imports = nil
	groupBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("GroupBox_Class", 0),
		/*1*/ imports.NewTable("GroupBox_Create", 0),
		/*2*/ imports.NewTable("GroupBox_DragCursor", 0),
		/*3*/ imports.NewTable("GroupBox_DragKind", 0),
		/*4*/ imports.NewTable("GroupBox_DragMode", 0),
		/*5*/ imports.NewTable("GroupBox_ParentColor", 0),
		/*6*/ imports.NewTable("GroupBox_ParentFont", 0),
		/*7*/ imports.NewTable("GroupBox_ParentShowHint", 0),
		/*8*/ imports.NewTable("GroupBox_SetOnContextPopup", 0),
		/*9*/ imports.NewTable("GroupBox_SetOnDblClick", 0),
		/*10*/ imports.NewTable("GroupBox_SetOnDragDrop", 0),
		/*11*/ imports.NewTable("GroupBox_SetOnDragOver", 0),
		/*12*/ imports.NewTable("GroupBox_SetOnEndDock", 0),
		/*13*/ imports.NewTable("GroupBox_SetOnEndDrag", 0),
		/*14*/ imports.NewTable("GroupBox_SetOnGetSiteInfo", 0),
		/*15*/ imports.NewTable("GroupBox_SetOnMouseDown", 0),
		/*16*/ imports.NewTable("GroupBox_SetOnMouseEnter", 0),
		/*17*/ imports.NewTable("GroupBox_SetOnMouseLeave", 0),
		/*18*/ imports.NewTable("GroupBox_SetOnMouseMove", 0),
		/*19*/ imports.NewTable("GroupBox_SetOnMouseUp", 0),
		/*20*/ imports.NewTable("GroupBox_SetOnMouseWheel", 0),
		/*21*/ imports.NewTable("GroupBox_SetOnMouseWheelDown", 0),
		/*22*/ imports.NewTable("GroupBox_SetOnMouseWheelUp", 0),
		/*23*/ imports.NewTable("GroupBox_SetOnStartDock", 0),
		/*24*/ imports.NewTable("GroupBox_SetOnStartDrag", 0),
	}
)

func groupBoxImportAPI() *imports.Imports {
	if groupBoxImport == nil {
		groupBoxImport = NewDefaultImports()
		groupBoxImport.SetImportTable(groupBoxImportTables)
		groupBoxImportTables = nil
	}
	return groupBoxImport
}
