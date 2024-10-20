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

// IFrame Parent: ICustomFrame
type IFrame interface {
	ICustomFrame
	AutoScroll() bool                                  // property
	SetAutoScroll(AValue bool)                         // property
	DragCursor() TCursor                               // property
	SetDragCursor(AValue TCursor)                      // property
	DragKind() TDragKind                               // property
	SetDragKind(AValue TDragKind)                      // property
	DragMode() TDragMode                               // property
	SetDragMode(AValue TDragMode)                      // property
	LCLVersion() string                                // property
	SetLCLVersion(AValue string)                       // property
	ParentColor() bool                                 // property
	SetParentColor(AValue bool)                        // property
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
	SetOnMouseWheel(fn TMouseWheelEvent)               // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)       // property event
	SetOnMouseWheelHorz(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent)     // property event
	SetOnMouseWheelRight(fn TMouseWheelUpDownEvent)    // property event
	SetOnStartDock(fn TStartDockEvent)                 // property event
	SetOnStartDrag(fn TStartDragEvent)                 // property event
}

// TFrame Parent: TCustomFrame
type TFrame struct {
	TCustomFrame
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
	mouseWheelPtr        uintptr
	mouseWheelDownPtr    uintptr
	mouseWheelUpPtr      uintptr
	mouseWheelHorzPtr    uintptr
	mouseWheelLeftPtr    uintptr
	mouseWheelRightPtr   uintptr
	startDockPtr         uintptr
	startDragPtr         uintptr
}

func NewFrame(TheOwner IComponent) IFrame {
	r1 := frameImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsFrame(r1)
}

func (m *TFrame) AutoScroll() bool {
	r1 := frameImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFrame) SetAutoScroll(AValue bool) {
	frameImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFrame) DragCursor() TCursor {
	r1 := frameImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TFrame) SetDragCursor(AValue TCursor) {
	frameImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TFrame) DragKind() TDragKind {
	r1 := frameImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TFrame) SetDragKind(AValue TDragKind) {
	frameImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TFrame) DragMode() TDragMode {
	r1 := frameImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TFrame) SetDragMode(AValue TDragMode) {
	frameImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TFrame) LCLVersion() string {
	r1 := frameImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TFrame) SetLCLVersion(AValue string) {
	frameImportAPI().SysCallN(6, 1, m.Instance(), PascalStr(AValue))
}

func (m *TFrame) ParentColor() bool {
	r1 := frameImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFrame) SetParentColor(AValue bool) {
	frameImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFrame) ParentFont() bool {
	r1 := frameImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFrame) SetParentFont(AValue bool) {
	frameImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TFrame) ParentShowHint() bool {
	r1 := frameImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TFrame) SetParentShowHint(AValue bool) {
	frameImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func FrameClass() TClass {
	ret := frameImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TFrame) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	if m.constrainedResizePtr != 0 {
		RemoveEventElement(m.constrainedResizePtr)
	}
	m.constrainedResizePtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(10, m.Instance(), m.constrainedResizePtr)
}

func (m *TFrame) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(11, m.Instance(), m.contextPopupPtr)
}

func (m *TFrame) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(12, m.Instance(), m.dblClickPtr)
}

func (m *TFrame) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(13, m.Instance(), m.dragDropPtr)
}

func (m *TFrame) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(14, m.Instance(), m.dragOverPtr)
}

func (m *TFrame) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(15, m.Instance(), m.endDockPtr)
}

func (m *TFrame) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(16, m.Instance(), m.endDragPtr)
}

func (m *TFrame) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
	if m.getSiteInfoPtr != 0 {
		RemoveEventElement(m.getSiteInfoPtr)
	}
	m.getSiteInfoPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(17, m.Instance(), m.getSiteInfoPtr)
}

func (m *TFrame) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(18, m.Instance(), m.mouseDownPtr)
}

func (m *TFrame) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(19, m.Instance(), m.mouseEnterPtr)
}

func (m *TFrame) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(20, m.Instance(), m.mouseLeavePtr)
}

func (m *TFrame) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(21, m.Instance(), m.mouseMovePtr)
}

func (m *TFrame) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(22, m.Instance(), m.mouseUpPtr)
}

func (m *TFrame) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(23, m.Instance(), m.mouseWheelPtr)
}

func (m *TFrame) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(24, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TFrame) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(28, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TFrame) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(25, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TFrame) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(26, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TFrame) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(27, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TFrame) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(29, m.Instance(), m.startDockPtr)
}

func (m *TFrame) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	frameImportAPI().SysCallN(30, m.Instance(), m.startDragPtr)
}

var (
	frameImport       *imports.Imports = nil
	frameImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("Frame_AutoScroll", 0),
		/*1*/ imports.NewTable("Frame_Class", 0),
		/*2*/ imports.NewTable("Frame_Create", 0),
		/*3*/ imports.NewTable("Frame_DragCursor", 0),
		/*4*/ imports.NewTable("Frame_DragKind", 0),
		/*5*/ imports.NewTable("Frame_DragMode", 0),
		/*6*/ imports.NewTable("Frame_LCLVersion", 0),
		/*7*/ imports.NewTable("Frame_ParentColor", 0),
		/*8*/ imports.NewTable("Frame_ParentFont", 0),
		/*9*/ imports.NewTable("Frame_ParentShowHint", 0),
		/*10*/ imports.NewTable("Frame_SetOnConstrainedResize", 0),
		/*11*/ imports.NewTable("Frame_SetOnContextPopup", 0),
		/*12*/ imports.NewTable("Frame_SetOnDblClick", 0),
		/*13*/ imports.NewTable("Frame_SetOnDragDrop", 0),
		/*14*/ imports.NewTable("Frame_SetOnDragOver", 0),
		/*15*/ imports.NewTable("Frame_SetOnEndDock", 0),
		/*16*/ imports.NewTable("Frame_SetOnEndDrag", 0),
		/*17*/ imports.NewTable("Frame_SetOnGetSiteInfo", 0),
		/*18*/ imports.NewTable("Frame_SetOnMouseDown", 0),
		/*19*/ imports.NewTable("Frame_SetOnMouseEnter", 0),
		/*20*/ imports.NewTable("Frame_SetOnMouseLeave", 0),
		/*21*/ imports.NewTable("Frame_SetOnMouseMove", 0),
		/*22*/ imports.NewTable("Frame_SetOnMouseUp", 0),
		/*23*/ imports.NewTable("Frame_SetOnMouseWheel", 0),
		/*24*/ imports.NewTable("Frame_SetOnMouseWheelDown", 0),
		/*25*/ imports.NewTable("Frame_SetOnMouseWheelHorz", 0),
		/*26*/ imports.NewTable("Frame_SetOnMouseWheelLeft", 0),
		/*27*/ imports.NewTable("Frame_SetOnMouseWheelRight", 0),
		/*28*/ imports.NewTable("Frame_SetOnMouseWheelUp", 0),
		/*29*/ imports.NewTable("Frame_SetOnStartDock", 0),
		/*30*/ imports.NewTable("Frame_SetOnStartDrag", 0),
	}
)

func frameImportAPI() *imports.Imports {
	if frameImport == nil {
		frameImport = NewDefaultImports()
		frameImport.SetImportTable(frameImportTables)
		frameImportTables = nil
	}
	return frameImport
}
