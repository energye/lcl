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

// IScrollBox Parent: IScrollingWinControl
type IScrollBox interface {
	IScrollingWinControl
	AutoScroll() bool                                  // property
	SetAutoScroll(AValue bool)                         // property
	DragCursor() TCursor                               // property
	SetDragCursor(AValue TCursor)                      // property
	DragKind() TDragKind                               // property
	SetDragKind(AValue TDragKind)                      // property
	DragMode() TDragMode                               // property
	SetDragMode(AValue TDragMode)                      // property
	ParentBackground() bool                            // property
	SetParentBackground(AValue bool)                   // property
	ParentColor() bool                                 // property
	SetParentColor(AValue bool)                        // property
	ParentFont() bool                                  // property
	SetParentFont(AValue bool)                         // property
	ParentShowHint() bool                              // property
	SetParentShowHint(AValue bool)                     // property
	SetOnConstrainedResize(fn TConstrainedResizeEvent) // property event
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
	SetOnMouseWheel(fn TMouseWheelEvent)               // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)    // property event
	SetOnStartDock(fn TStartDockEvent)                 // property event
	SetOnStartDrag(fn TStartDragEvent)                 // property event
}

// TScrollBox Parent: TScrollingWinControl
type TScrollBox struct {
	TScrollingWinControl
	constrainedResizePtr uintptr
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
	mouseWheelPtr        uintptr
	mouseWheelDownPtr    uintptr
	mouseWheelUpPtr      uintptr
	mouseWheelHorzPtr    uintptr
	mouseWheelLeftPtr    uintptr
	mouseWheelRightPtr   uintptr
	startDockPtr         uintptr
	startDragPtr         uintptr
}

func NewScrollBox(AOwner IComponent) IScrollBox {
	r1 := scrollBoxImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsScrollBox(r1)
}

func (m *TScrollBox) AutoScroll() bool {
	r1 := scrollBoxImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBox) SetAutoScroll(AValue bool) {
	scrollBoxImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TScrollBox) DragCursor() TCursor {
	r1 := scrollBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TScrollBox) SetDragCursor(AValue TCursor) {
	scrollBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBox) DragKind() TDragKind {
	r1 := scrollBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TScrollBox) SetDragKind(AValue TDragKind) {
	scrollBoxImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBox) DragMode() TDragMode {
	r1 := scrollBoxImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TScrollBox) SetDragMode(AValue TDragMode) {
	scrollBoxImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TScrollBox) ParentBackground() bool {
	r1 := scrollBoxImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBox) SetParentBackground(AValue bool) {
	scrollBoxImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TScrollBox) ParentColor() bool {
	r1 := scrollBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBox) SetParentColor(AValue bool) {
	scrollBoxImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TScrollBox) ParentFont() bool {
	r1 := scrollBoxImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBox) SetParentFont(AValue bool) {
	scrollBoxImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TScrollBox) ParentShowHint() bool {
	r1 := scrollBoxImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TScrollBox) SetParentShowHint(AValue bool) {
	scrollBoxImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func ScrollBoxClass() TClass {
	ret := scrollBoxImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TScrollBox) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if m.constrainedResizePtr != 0 {
		RemoveEventElement(m.constrainedResizePtr)
	}
	m.constrainedResizePtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(10, m.Instance(), m.constrainedResizePtr)
}

func (m *TScrollBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(11, m.Instance(), m.dblClickPtr)
}

func (m *TScrollBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(12, m.Instance(), m.dragDropPtr)
}

func (m *TScrollBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(13, m.Instance(), m.dragOverPtr)
}

func (m *TScrollBox) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(14, m.Instance(), m.endDockPtr)
}

func (m *TScrollBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(15, m.Instance(), m.endDragPtr)
}

func (m *TScrollBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(16, m.Instance(), m.getSiteInfoPtr)
}

func (m *TScrollBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(17, m.Instance(), m.mouseDownPtr)
}

func (m *TScrollBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(18, m.Instance(), m.mouseEnterPtr)
}

func (m *TScrollBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(19, m.Instance(), m.mouseLeavePtr)
}

func (m *TScrollBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(20, m.Instance(), m.mouseMovePtr)
}

func (m *TScrollBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(21, m.Instance(), m.mouseUpPtr)
}

func (m *TScrollBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(22, m.Instance(), m.mouseWheelPtr)
}

func (m *TScrollBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(23, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TScrollBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(27, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TScrollBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(24, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TScrollBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(25, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TScrollBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(26, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TScrollBox) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(28, m.Instance(), m.startDockPtr)
}

func (m *TScrollBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	scrollBoxImportAPI().SysCallN(29, m.Instance(), m.startDragPtr)
}

var (
	scrollBoxImport       *imports.Imports = nil
	scrollBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ScrollBox_AutoScroll", 0),
		/*1*/ imports.NewTable("ScrollBox_Class", 0),
		/*2*/ imports.NewTable("ScrollBox_Create", 0),
		/*3*/ imports.NewTable("ScrollBox_DragCursor", 0),
		/*4*/ imports.NewTable("ScrollBox_DragKind", 0),
		/*5*/ imports.NewTable("ScrollBox_DragMode", 0),
		/*6*/ imports.NewTable("ScrollBox_ParentBackground", 0),
		/*7*/ imports.NewTable("ScrollBox_ParentColor", 0),
		/*8*/ imports.NewTable("ScrollBox_ParentFont", 0),
		/*9*/ imports.NewTable("ScrollBox_ParentShowHint", 0),
		/*10*/ imports.NewTable("ScrollBox_SetOnConstrainedResize", 0),
		/*11*/ imports.NewTable("ScrollBox_SetOnDblClick", 0),
		/*12*/ imports.NewTable("ScrollBox_SetOnDragDrop", 0),
		/*13*/ imports.NewTable("ScrollBox_SetOnDragOver", 0),
		/*14*/ imports.NewTable("ScrollBox_SetOnEndDock", 0),
		/*15*/ imports.NewTable("ScrollBox_SetOnEndDrag", 0),
		/*16*/ imports.NewTable("ScrollBox_SetOnGetSiteInfo", 0),
		/*17*/ imports.NewTable("ScrollBox_SetOnMouseDown", 0),
		/*18*/ imports.NewTable("ScrollBox_SetOnMouseEnter", 0),
		/*19*/ imports.NewTable("ScrollBox_SetOnMouseLeave", 0),
		/*20*/ imports.NewTable("ScrollBox_SetOnMouseMove", 0),
		/*21*/ imports.NewTable("ScrollBox_SetOnMouseUp", 0),
		/*22*/ imports.NewTable("ScrollBox_SetOnMouseWheel", 0),
		/*23*/ imports.NewTable("ScrollBox_SetOnMouseWheelDown", 0),
		/*24*/ imports.NewTable("ScrollBox_SetOnMouseWheelHorz", 0),
		/*25*/ imports.NewTable("ScrollBox_SetOnMouseWheelLeft", 0),
		/*26*/ imports.NewTable("ScrollBox_SetOnMouseWheelRight", 0),
		/*27*/ imports.NewTable("ScrollBox_SetOnMouseWheelUp", 0),
		/*28*/ imports.NewTable("ScrollBox_SetOnStartDock", 0),
		/*29*/ imports.NewTable("ScrollBox_SetOnStartDrag", 0),
	}
)

func scrollBoxImportAPI() *imports.Imports {
	if scrollBoxImport == nil {
		scrollBoxImport = NewDefaultImports()
		scrollBoxImport.SetImportTable(scrollBoxImportTables)
		scrollBoxImportTables = nil
	}
	return scrollBoxImport
}
