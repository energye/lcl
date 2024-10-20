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

// ICoolBar Parent: ICustomCoolBar
type ICoolBar interface {
	ICustomCoolBar
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

// TCoolBar Parent: TCustomCoolBar
type TCoolBar struct {
	TCustomCoolBar
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

func NewCoolBar(AOwner IComponent) ICoolBar {
	r1 := coolBarImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCoolBar(r1)
}

func (m *TCoolBar) DragCursor() TCursor {
	r1 := coolBarImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCoolBar) SetDragCursor(AValue TCursor) {
	coolBarImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBar) DragKind() TDragKind {
	r1 := coolBarImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TCoolBar) SetDragKind(AValue TDragKind) {
	coolBarImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBar) DragMode() TDragMode {
	r1 := coolBarImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCoolBar) SetDragMode(AValue TDragMode) {
	coolBarImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoolBar) ParentColor() bool {
	r1 := coolBarImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBar) SetParentColor(AValue bool) {
	coolBarImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBar) ParentFont() bool {
	r1 := coolBarImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBar) SetParentFont(AValue bool) {
	coolBarImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoolBar) ParentShowHint() bool {
	r1 := coolBarImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoolBar) SetParentShowHint(AValue bool) {
	coolBarImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func CoolBarClass() TClass {
	ret := coolBarImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TCoolBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(8, m.Instance(), m.contextPopupPtr)
}

func (m *TCoolBar) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(9, m.Instance(), m.dblClickPtr)
}

func (m *TCoolBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(10, m.Instance(), m.dragDropPtr)
}

func (m *TCoolBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(11, m.Instance(), m.dragOverPtr)
}

func (m *TCoolBar) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(12, m.Instance(), m.endDockPtr)
}

func (m *TCoolBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(13, m.Instance(), m.endDragPtr)
}

func (m *TCoolBar) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(14, m.Instance(), m.getSiteInfoPtr)
}

func (m *TCoolBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(15, m.Instance(), m.mouseDownPtr)
}

func (m *TCoolBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(16, m.Instance(), m.mouseEnterPtr)
}

func (m *TCoolBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(17, m.Instance(), m.mouseLeavePtr)
}

func (m *TCoolBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(18, m.Instance(), m.mouseMovePtr)
}

func (m *TCoolBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(19, m.Instance(), m.mouseUpPtr)
}

func (m *TCoolBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(20, m.Instance(), m.mouseWheelPtr)
}

func (m *TCoolBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(21, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCoolBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(22, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCoolBar) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(23, m.Instance(), m.startDockPtr)
}

func (m *TCoolBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	coolBarImportAPI().SysCallN(24, m.Instance(), m.startDragPtr)
}

var (
	coolBarImport       *imports.Imports = nil
	coolBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CoolBar_Class", 0),
		/*1*/ imports.NewTable("CoolBar_Create", 0),
		/*2*/ imports.NewTable("CoolBar_DragCursor", 0),
		/*3*/ imports.NewTable("CoolBar_DragKind", 0),
		/*4*/ imports.NewTable("CoolBar_DragMode", 0),
		/*5*/ imports.NewTable("CoolBar_ParentColor", 0),
		/*6*/ imports.NewTable("CoolBar_ParentFont", 0),
		/*7*/ imports.NewTable("CoolBar_ParentShowHint", 0),
		/*8*/ imports.NewTable("CoolBar_SetOnContextPopup", 0),
		/*9*/ imports.NewTable("CoolBar_SetOnDblClick", 0),
		/*10*/ imports.NewTable("CoolBar_SetOnDragDrop", 0),
		/*11*/ imports.NewTable("CoolBar_SetOnDragOver", 0),
		/*12*/ imports.NewTable("CoolBar_SetOnEndDock", 0),
		/*13*/ imports.NewTable("CoolBar_SetOnEndDrag", 0),
		/*14*/ imports.NewTable("CoolBar_SetOnGetSiteInfo", 0),
		/*15*/ imports.NewTable("CoolBar_SetOnMouseDown", 0),
		/*16*/ imports.NewTable("CoolBar_SetOnMouseEnter", 0),
		/*17*/ imports.NewTable("CoolBar_SetOnMouseLeave", 0),
		/*18*/ imports.NewTable("CoolBar_SetOnMouseMove", 0),
		/*19*/ imports.NewTable("CoolBar_SetOnMouseUp", 0),
		/*20*/ imports.NewTable("CoolBar_SetOnMouseWheel", 0),
		/*21*/ imports.NewTable("CoolBar_SetOnMouseWheelDown", 0),
		/*22*/ imports.NewTable("CoolBar_SetOnMouseWheelUp", 0),
		/*23*/ imports.NewTable("CoolBar_SetOnStartDock", 0),
		/*24*/ imports.NewTable("CoolBar_SetOnStartDrag", 0),
	}
)

func coolBarImportAPI() *imports.Imports {
	if coolBarImport == nil {
		coolBarImport = NewDefaultImports()
		coolBarImport.SetImportTable(coolBarImportTables)
		coolBarImportTables = nil
	}
	return coolBarImport
}
