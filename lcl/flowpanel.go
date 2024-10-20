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

// IFlowPanel Parent: ICustomFlowPanel
type IFlowPanel interface {
	ICustomFlowPanel
	DragCursor() TCursor                               // property
	SetDragCursor(AValue TCursor)                      // property
	DragKind() TDragKind                               // property
	SetDragKind(AValue TDragKind)                      // property
	DragMode() TDragMode                               // property
	SetDragMode(AValue TDragMode)                      // property
	ParentFont() bool                                  // property
	SetParentFont(AValue bool)                         // property
	ParentShowHint() bool                              // property
	SetParentShowHint(AValue bool)                     // property
	SetOnConstrainedResize(fn TConstrainedResizeEvent) // property event
	SetOnContextPopup(fn TContextPopupEvent)           // property event
	SetOnDblClick(fn TNotifyEvent)                     // property event
	SetOnDragDrop(fn TDragDropEvent)                   // property event
	SetOnDragOver(fn TDragOverEvent)                   // property event
	SetOnEndDock(fn TEndDragEvent)                     // property event
	SetOnEndDrag(fn TEndDragEvent)                     // property event
	SetOnGetSiteInfo(fn TGetSiteInfoEvent)             // property event
	SetOnMouseDown(fn TMouseEvent)                     // property event
	SetOnMouseEnter(fn TNotifyEvent)                   // property event
	SetOnMouseLeave(fn TNotifyEvent)                   // property event
	SetOnMouseMove(fn TMouseMoveEvent)                 // property event
	SetOnMouseUp(fn TMouseEvent)                       // property event
	SetOnStartDock(fn TStartDockEvent)                 // property event
	SetOnStartDrag(fn TStartDragEvent)                 // property event
}

// TFlowPanel Parent: TCustomFlowPanel
type TFlowPanel struct {
	TCustomFlowPanel
	constrainedResizePtr uintptr
	contextPopupPtr      uintptr
	dblClickPtr          uintptr
	dragDropPtr          uintptr
	dragOverPtr          uintptr
	endDockPtr           uintptr
	endDragPtr           uintptr
	getSiteInfoPtr       uintptr
	mouseDownPtr         uintptr
	mouseEnterPtr        uintptr
	mouseLeavePtr        uintptr
	mouseMovePtr         uintptr
	mouseUpPtr           uintptr
	startDockPtr         uintptr
	startDragPtr         uintptr
}

func NewFlowPanel(AOwner IComponent) IFlowPanel {
	r1 := flowPanelImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsFlowPanel(r1)
}

func (m *TFlowPanel) DragCursor() TCursor {
	r1 := flowPanelImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TFlowPanel) SetDragCursor(AValue TCursor) {
	flowPanelImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TFlowPanel) DragKind() TDragKind {
	r1 := flowPanelImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TFlowPanel) SetDragKind(AValue TDragKind) {
	flowPanelImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TFlowPanel) DragMode() TDragMode {
	r1 := flowPanelImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TFlowPanel) SetDragMode(AValue TDragMode) {
	flowPanelImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TFlowPanel) ParentFont() bool {
	r1 := flowPanelImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFlowPanel) SetParentFont(AValue bool) {
	flowPanelImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFlowPanel) ParentShowHint() bool {
	r1 := flowPanelImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFlowPanel) SetParentShowHint(AValue bool) {
	flowPanelImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func FlowPanelClass() TClass {
	ret := flowPanelImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TFlowPanel) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if m.constrainedResizePtr != 0 {
		RemoveEventElement(m.constrainedResizePtr)
	}
	m.constrainedResizePtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(7, m.Instance(), m.constrainedResizePtr)
}

func (m *TFlowPanel) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(8, m.Instance(), m.contextPopupPtr)
}

func (m *TFlowPanel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(9, m.Instance(), m.dblClickPtr)
}

func (m *TFlowPanel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(10, m.Instance(), m.dragDropPtr)
}

func (m *TFlowPanel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(11, m.Instance(), m.dragOverPtr)
}

func (m *TFlowPanel) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(12, m.Instance(), m.endDockPtr)
}

func (m *TFlowPanel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(13, m.Instance(), m.endDragPtr)
}

func (m *TFlowPanel) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(14, m.Instance(), m.getSiteInfoPtr)
}

func (m *TFlowPanel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(15, m.Instance(), m.mouseDownPtr)
}

func (m *TFlowPanel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(16, m.Instance(), m.mouseEnterPtr)
}

func (m *TFlowPanel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(17, m.Instance(), m.mouseLeavePtr)
}

func (m *TFlowPanel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(18, m.Instance(), m.mouseMovePtr)
}

func (m *TFlowPanel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(19, m.Instance(), m.mouseUpPtr)
}

func (m *TFlowPanel) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(20, m.Instance(), m.startDockPtr)
}

func (m *TFlowPanel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	flowPanelImportAPI().SysCallN(21, m.Instance(), m.startDragPtr)
}

var (
	flowPanelImport       *imports.Imports = nil
	flowPanelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("FlowPanel_Class", 0),
		/*1*/ imports.NewTable("FlowPanel_Create", 0),
		/*2*/ imports.NewTable("FlowPanel_DragCursor", 0),
		/*3*/ imports.NewTable("FlowPanel_DragKind", 0),
		/*4*/ imports.NewTable("FlowPanel_DragMode", 0),
		/*5*/ imports.NewTable("FlowPanel_ParentFont", 0),
		/*6*/ imports.NewTable("FlowPanel_ParentShowHint", 0),
		/*7*/ imports.NewTable("FlowPanel_SetOnConstrainedResize", 0),
		/*8*/ imports.NewTable("FlowPanel_SetOnContextPopup", 0),
		/*9*/ imports.NewTable("FlowPanel_SetOnDblClick", 0),
		/*10*/ imports.NewTable("FlowPanel_SetOnDragDrop", 0),
		/*11*/ imports.NewTable("FlowPanel_SetOnDragOver", 0),
		/*12*/ imports.NewTable("FlowPanel_SetOnEndDock", 0),
		/*13*/ imports.NewTable("FlowPanel_SetOnEndDrag", 0),
		/*14*/ imports.NewTable("FlowPanel_SetOnGetSiteInfo", 0),
		/*15*/ imports.NewTable("FlowPanel_SetOnMouseDown", 0),
		/*16*/ imports.NewTable("FlowPanel_SetOnMouseEnter", 0),
		/*17*/ imports.NewTable("FlowPanel_SetOnMouseLeave", 0),
		/*18*/ imports.NewTable("FlowPanel_SetOnMouseMove", 0),
		/*19*/ imports.NewTable("FlowPanel_SetOnMouseUp", 0),
		/*20*/ imports.NewTable("FlowPanel_SetOnStartDock", 0),
		/*21*/ imports.NewTable("FlowPanel_SetOnStartDrag", 0),
	}
)

func flowPanelImportAPI() *imports.Imports {
	if flowPanelImport == nil {
		flowPanelImport = NewDefaultImports()
		flowPanelImport.SetImportTable(flowPanelImportTables)
		flowPanelImportTables = nil
	}
	return flowPanelImport
}
