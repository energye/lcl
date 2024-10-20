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

// IPanel Parent: ICustomPanel
type IPanel interface {
	ICustomPanel
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
	ShowAccelChar() bool                            // property
	SetShowAccelChar(AValue bool)                   // property
	VerticalAlignment() TVerticalAlignment          // property
	SetVerticalAlignment(AValue TVerticalAlignment) // property
	Wordwrap() bool                                 // property
	SetWordwrap(AValue bool)                        // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDock(fn TEndDragEvent)                  // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)          // property event
	SetOnGetDockCaption(fn TGetDockCaptionEvent)    // property event
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
	SetOnStartDock(fn TStartDockEvent)              // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TPanel Parent: TCustomPanel
type TPanel struct {
	TCustomPanel
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDockPtr         uintptr
	endDragPtr         uintptr
	getSiteInfoPtr     uintptr
	getDockCaptionPtr  uintptr
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
	startDockPtr       uintptr
	startDragPtr       uintptr
}

func NewPanel(TheOwner IComponent) IPanel {
	r1 := panelImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsPanel(r1)
}

func (m *TPanel) DragCursor() TCursor {
	r1 := panelImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TPanel) SetDragCursor(AValue TCursor) {
	panelImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TPanel) DragKind() TDragKind {
	r1 := panelImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TPanel) SetDragKind(AValue TDragKind) {
	panelImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TPanel) DragMode() TDragMode {
	r1 := panelImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TPanel) SetDragMode(AValue TDragMode) {
	panelImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TPanel) ParentFont() bool {
	r1 := panelImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPanel) SetParentFont(AValue bool) {
	panelImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPanel) ParentShowHint() bool {
	r1 := panelImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPanel) SetParentShowHint(AValue bool) {
	panelImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPanel) ShowAccelChar() bool {
	r1 := panelImportAPI().SysCallN(28, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPanel) SetShowAccelChar(AValue bool) {
	panelImportAPI().SysCallN(28, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPanel) VerticalAlignment() TVerticalAlignment {
	r1 := panelImportAPI().SysCallN(29, 0, m.Instance(), 0)
	return TVerticalAlignment(r1)
}

func (m *TPanel) SetVerticalAlignment(AValue TVerticalAlignment) {
	panelImportAPI().SysCallN(29, 1, m.Instance(), uintptr(AValue))
}

func (m *TPanel) Wordwrap() bool {
	r1 := panelImportAPI().SysCallN(30, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPanel) SetWordwrap(AValue bool) {
	panelImportAPI().SysCallN(30, 1, m.Instance(), PascalBool(AValue))
}

func PanelClass() TClass {
	ret := panelImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TPanel) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(7, m.Instance(), m.contextPopupPtr)
}

func (m *TPanel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(8, m.Instance(), m.dblClickPtr)
}

func (m *TPanel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(9, m.Instance(), m.dragDropPtr)
}

func (m *TPanel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(10, m.Instance(), m.dragOverPtr)
}

func (m *TPanel) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(11, m.Instance(), m.endDockPtr)
}

func (m *TPanel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(12, m.Instance(), m.endDragPtr)
}

func (m *TPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(14, m.Instance(), m.getSiteInfoPtr)
}

func (m *TPanel) SetOnGetDockCaption(fn TGetDockCaptionEvent) {
	if m.getDockCaptionPtr != 0 {
		RemoveEventElement(m.getDockCaptionPtr)
	}
	m.getDockCaptionPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(13, m.Instance(), m.getDockCaptionPtr)
}

func (m *TPanel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(15, m.Instance(), m.mouseDownPtr)
}

func (m *TPanel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(16, m.Instance(), m.mouseEnterPtr)
}

func (m *TPanel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(17, m.Instance(), m.mouseLeavePtr)
}

func (m *TPanel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(18, m.Instance(), m.mouseMovePtr)
}

func (m *TPanel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(19, m.Instance(), m.mouseUpPtr)
}

func (m *TPanel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(20, m.Instance(), m.mouseWheelPtr)
}

func (m *TPanel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(21, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TPanel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(25, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TPanel) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(22, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TPanel) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(23, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TPanel) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(24, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TPanel) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(26, m.Instance(), m.startDockPtr)
}

func (m *TPanel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	panelImportAPI().SysCallN(27, m.Instance(), m.startDragPtr)
}

var (
	panelImport       *imports.Imports = nil
	panelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Panel_Class", 0),
		/*1*/ imports.NewTable("Panel_Create", 0),
		/*2*/ imports.NewTable("Panel_DragCursor", 0),
		/*3*/ imports.NewTable("Panel_DragKind", 0),
		/*4*/ imports.NewTable("Panel_DragMode", 0),
		/*5*/ imports.NewTable("Panel_ParentFont", 0),
		/*6*/ imports.NewTable("Panel_ParentShowHint", 0),
		/*7*/ imports.NewTable("Panel_SetOnContextPopup", 0),
		/*8*/ imports.NewTable("Panel_SetOnDblClick", 0),
		/*9*/ imports.NewTable("Panel_SetOnDragDrop", 0),
		/*10*/ imports.NewTable("Panel_SetOnDragOver", 0),
		/*11*/ imports.NewTable("Panel_SetOnEndDock", 0),
		/*12*/ imports.NewTable("Panel_SetOnEndDrag", 0),
		/*13*/ imports.NewTable("Panel_SetOnGetDockCaption", 0),
		/*14*/ imports.NewTable("Panel_SetOnGetSiteInfo", 0),
		/*15*/ imports.NewTable("Panel_SetOnMouseDown", 0),
		/*16*/ imports.NewTable("Panel_SetOnMouseEnter", 0),
		/*17*/ imports.NewTable("Panel_SetOnMouseLeave", 0),
		/*18*/ imports.NewTable("Panel_SetOnMouseMove", 0),
		/*19*/ imports.NewTable("Panel_SetOnMouseUp", 0),
		/*20*/ imports.NewTable("Panel_SetOnMouseWheel", 0),
		/*21*/ imports.NewTable("Panel_SetOnMouseWheelDown", 0),
		/*22*/ imports.NewTable("Panel_SetOnMouseWheelHorz", 0),
		/*23*/ imports.NewTable("Panel_SetOnMouseWheelLeft", 0),
		/*24*/ imports.NewTable("Panel_SetOnMouseWheelRight", 0),
		/*25*/ imports.NewTable("Panel_SetOnMouseWheelUp", 0),
		/*26*/ imports.NewTable("Panel_SetOnStartDock", 0),
		/*27*/ imports.NewTable("Panel_SetOnStartDrag", 0),
		/*28*/ imports.NewTable("Panel_ShowAccelChar", 0),
		/*29*/ imports.NewTable("Panel_VerticalAlignment", 0),
		/*30*/ imports.NewTable("Panel_Wordwrap", 0),
	}
)

func panelImportAPI() *imports.Imports {
	if panelImport == nil {
		panelImport = NewDefaultImports()
		panelImport.SetImportTable(panelImportTables)
		panelImportTables = nil
	}
	return panelImport
}
