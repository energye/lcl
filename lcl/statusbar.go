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

// IStatusBar Parent: IWinControl
type IStatusBar interface {
	IWinControl
	Canvas() ICanvas                                          // property
	AutoHint() bool                                           // property
	SetAutoHint(AValue bool)                                  // property
	DragCursor() TCursor                                      // property
	SetDragCursor(AValue TCursor)                             // property
	DragKind() TDragKind                                      // property
	SetDragKind(AValue TDragKind)                             // property
	DragMode() TDragMode                                      // property
	SetDragMode(AValue TDragMode)                             // property
	Panels() IStatusPanels                                    // property
	SetPanels(AValue IStatusPanels)                           // property
	ParentColor() bool                                        // property
	SetParentColor(AValue bool)                               // property
	ParentFont() bool                                         // property
	SetParentFont(AValue bool)                                // property
	ParentShowHint() bool                                     // property
	SetParentShowHint(AValue bool)                            // property
	SimpleText() string                                       // property
	SetSimpleText(AValue string)                              // property
	SimplePanel() bool                                        // property
	SetSimplePanel(AValue bool)                               // property
	SizeGrip() bool                                           // property
	SetSizeGrip(AValue bool)                                  // property
	UseSystemFont() bool                                      // property
	SetUseSystemFont(AValue bool)                             // property
	GetPanelIndexAt(X, Y int32) int32                         // function
	SizeGripEnabled() bool                                    // function
	UpdatingStatusBar() bool                                  // function
	InvalidatePanel(PanelIndex int32, PanelParts TPanelParts) // procedure
	BeginUpdate()                                             // procedure
	EndUpdate()                                               // procedure
	SetOnContextPopup(fn TContextPopupEvent)                  // property event
	SetOnCreatePanelClass(fn TSBCreatePanelClassEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                            // property event
	SetOnDragDrop(fn TDragDropEvent)                          // property event
	SetOnDragOver(fn TDragOverEvent)                          // property event
	SetOnDrawPanel(fn TDrawPanelEvent)                        // property event
	SetOnEndDock(fn TEndDragEvent)                            // property event
	SetOnEndDrag(fn TEndDragEvent)                            // property event
	SetOnHint(fn TNotifyEvent)                                // property event
	SetOnMouseDown(fn TMouseEvent)                            // property event
	SetOnMouseEnter(fn TNotifyEvent)                          // property event
	SetOnMouseLeave(fn TNotifyEvent)                          // property event
	SetOnMouseMove(fn TMouseMoveEvent)                        // property event
	SetOnMouseUp(fn TMouseEvent)                              // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                      // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)            // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)              // property event
	SetOnStartDock(fn TStartDockEvent)                        // property event
	SetOnStartDrag(fn TStartDragEvent)                        // property event
}

// TStatusBar Parent: TWinControl
type TStatusBar struct {
	TWinControl
	contextPopupPtr     uintptr
	createPanelClassPtr uintptr
	dblClickPtr         uintptr
	dragDropPtr         uintptr
	dragOverPtr         uintptr
	drawPanelPtr        uintptr
	endDockPtr          uintptr
	endDragPtr          uintptr
	hintPtr             uintptr
	mouseDownPtr        uintptr
	mouseEnterPtr       uintptr
	mouseLeavePtr       uintptr
	mouseMovePtr        uintptr
	mouseUpPtr          uintptr
	mouseWheelPtr       uintptr
	mouseWheelDownPtr   uintptr
	mouseWheelUpPtr     uintptr
	startDockPtr        uintptr
	startDragPtr        uintptr
}

func NewStatusBar(TheOwner IComponent) IStatusBar {
	r1 := statusBarImportAPI().SysCallN(4, GetObjectUintptr(TheOwner))
	return AsStatusBar(r1)
}

func (m *TStatusBar) Canvas() ICanvas {
	r1 := statusBarImportAPI().SysCallN(2, m.Instance())
	return AsCanvas(r1)
}

func (m *TStatusBar) AutoHint() bool {
	r1 := statusBarImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetAutoHint(AValue bool) {
	statusBarImportAPI().SysCallN(0, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) DragCursor() TCursor {
	r1 := statusBarImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStatusBar) SetDragCursor(AValue TCursor) {
	statusBarImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) DragKind() TDragKind {
	r1 := statusBarImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TStatusBar) SetDragKind(AValue TDragKind) {
	statusBarImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) DragMode() TDragMode {
	r1 := statusBarImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TStatusBar) SetDragMode(AValue TDragMode) {
	statusBarImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) Panels() IStatusPanels {
	r1 := statusBarImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return AsStatusPanels(r1)
}

func (m *TStatusBar) SetPanels(AValue IStatusPanels) {
	statusBarImportAPI().SysCallN(11, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TStatusBar) ParentColor() bool {
	r1 := statusBarImportAPI().SysCallN(12, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentColor(AValue bool) {
	statusBarImportAPI().SysCallN(12, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) ParentFont() bool {
	r1 := statusBarImportAPI().SysCallN(13, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentFont(AValue bool) {
	statusBarImportAPI().SysCallN(13, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) ParentShowHint() bool {
	r1 := statusBarImportAPI().SysCallN(14, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentShowHint(AValue bool) {
	statusBarImportAPI().SysCallN(14, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) SimpleText() string {
	r1 := statusBarImportAPI().SysCallN(35, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStatusBar) SetSimpleText(AValue string) {
	statusBarImportAPI().SysCallN(35, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStatusBar) SimplePanel() bool {
	r1 := statusBarImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetSimplePanel(AValue bool) {
	statusBarImportAPI().SysCallN(34, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) SizeGrip() bool {
	r1 := statusBarImportAPI().SysCallN(36, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetSizeGrip(AValue bool) {
	statusBarImportAPI().SysCallN(36, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) UseSystemFont() bool {
	r1 := statusBarImportAPI().SysCallN(39, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetUseSystemFont(AValue bool) {
	statusBarImportAPI().SysCallN(39, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) GetPanelIndexAt(X, Y int32) int32 {
	r1 := statusBarImportAPI().SysCallN(9, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r1)
}

func (m *TStatusBar) SizeGripEnabled() bool {
	r1 := statusBarImportAPI().SysCallN(37, m.Instance())
	return GoBool(r1)
}

func (m *TStatusBar) UpdatingStatusBar() bool {
	r1 := statusBarImportAPI().SysCallN(38, m.Instance())
	return GoBool(r1)
}

func StatusBarClass() TClass {
	ret := statusBarImportAPI().SysCallN(3)
	return TClass(ret)
}

func (m *TStatusBar) InvalidatePanel(PanelIndex int32, PanelParts TPanelParts) {
	statusBarImportAPI().SysCallN(10, m.Instance(), uintptr(PanelIndex), uintptr(PanelParts))
}

func (m *TStatusBar) BeginUpdate() {
	statusBarImportAPI().SysCallN(1, m.Instance())
}

func (m *TStatusBar) EndUpdate() {
	statusBarImportAPI().SysCallN(8, m.Instance())
}

func (m *TStatusBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(15, m.Instance(), m.contextPopupPtr)
}

func (m *TStatusBar) SetOnCreatePanelClass(fn TSBCreatePanelClassEvent) {
	if m.createPanelClassPtr != 0 {
		RemoveEventElement(m.createPanelClassPtr)
	}
	m.createPanelClassPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(16, m.Instance(), m.createPanelClassPtr)
}

func (m *TStatusBar) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(17, m.Instance(), m.dblClickPtr)
}

func (m *TStatusBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(18, m.Instance(), m.dragDropPtr)
}

func (m *TStatusBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(19, m.Instance(), m.dragOverPtr)
}

func (m *TStatusBar) SetOnDrawPanel(fn TDrawPanelEvent) {
	if m.drawPanelPtr != 0 {
		RemoveEventElement(m.drawPanelPtr)
	}
	m.drawPanelPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(20, m.Instance(), m.drawPanelPtr)
}

func (m *TStatusBar) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(21, m.Instance(), m.endDockPtr)
}

func (m *TStatusBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(22, m.Instance(), m.endDragPtr)
}

func (m *TStatusBar) SetOnHint(fn TNotifyEvent) {
	if m.hintPtr != 0 {
		RemoveEventElement(m.hintPtr)
	}
	m.hintPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(23, m.Instance(), m.hintPtr)
}

func (m *TStatusBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(24, m.Instance(), m.mouseDownPtr)
}

func (m *TStatusBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(25, m.Instance(), m.mouseEnterPtr)
}

func (m *TStatusBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(26, m.Instance(), m.mouseLeavePtr)
}

func (m *TStatusBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(27, m.Instance(), m.mouseMovePtr)
}

func (m *TStatusBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(28, m.Instance(), m.mouseUpPtr)
}

func (m *TStatusBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(29, m.Instance(), m.mouseWheelPtr)
}

func (m *TStatusBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(30, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TStatusBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(31, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TStatusBar) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(32, m.Instance(), m.startDockPtr)
}

func (m *TStatusBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	statusBarImportAPI().SysCallN(33, m.Instance(), m.startDragPtr)
}

var (
	statusBarImport       *imports.Imports = nil
	statusBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("StatusBar_AutoHint", 0),
		/*1*/ imports.NewTable("StatusBar_BeginUpdate", 0),
		/*2*/ imports.NewTable("StatusBar_Canvas", 0),
		/*3*/ imports.NewTable("StatusBar_Class", 0),
		/*4*/ imports.NewTable("StatusBar_Create", 0),
		/*5*/ imports.NewTable("StatusBar_DragCursor", 0),
		/*6*/ imports.NewTable("StatusBar_DragKind", 0),
		/*7*/ imports.NewTable("StatusBar_DragMode", 0),
		/*8*/ imports.NewTable("StatusBar_EndUpdate", 0),
		/*9*/ imports.NewTable("StatusBar_GetPanelIndexAt", 0),
		/*10*/ imports.NewTable("StatusBar_InvalidatePanel", 0),
		/*11*/ imports.NewTable("StatusBar_Panels", 0),
		/*12*/ imports.NewTable("StatusBar_ParentColor", 0),
		/*13*/ imports.NewTable("StatusBar_ParentFont", 0),
		/*14*/ imports.NewTable("StatusBar_ParentShowHint", 0),
		/*15*/ imports.NewTable("StatusBar_SetOnContextPopup", 0),
		/*16*/ imports.NewTable("StatusBar_SetOnCreatePanelClass", 0),
		/*17*/ imports.NewTable("StatusBar_SetOnDblClick", 0),
		/*18*/ imports.NewTable("StatusBar_SetOnDragDrop", 0),
		/*19*/ imports.NewTable("StatusBar_SetOnDragOver", 0),
		/*20*/ imports.NewTable("StatusBar_SetOnDrawPanel", 0),
		/*21*/ imports.NewTable("StatusBar_SetOnEndDock", 0),
		/*22*/ imports.NewTable("StatusBar_SetOnEndDrag", 0),
		/*23*/ imports.NewTable("StatusBar_SetOnHint", 0),
		/*24*/ imports.NewTable("StatusBar_SetOnMouseDown", 0),
		/*25*/ imports.NewTable("StatusBar_SetOnMouseEnter", 0),
		/*26*/ imports.NewTable("StatusBar_SetOnMouseLeave", 0),
		/*27*/ imports.NewTable("StatusBar_SetOnMouseMove", 0),
		/*28*/ imports.NewTable("StatusBar_SetOnMouseUp", 0),
		/*29*/ imports.NewTable("StatusBar_SetOnMouseWheel", 0),
		/*30*/ imports.NewTable("StatusBar_SetOnMouseWheelDown", 0),
		/*31*/ imports.NewTable("StatusBar_SetOnMouseWheelUp", 0),
		/*32*/ imports.NewTable("StatusBar_SetOnStartDock", 0),
		/*33*/ imports.NewTable("StatusBar_SetOnStartDrag", 0),
		/*34*/ imports.NewTable("StatusBar_SimplePanel", 0),
		/*35*/ imports.NewTable("StatusBar_SimpleText", 0),
		/*36*/ imports.NewTable("StatusBar_SizeGrip", 0),
		/*37*/ imports.NewTable("StatusBar_SizeGripEnabled", 0),
		/*38*/ imports.NewTable("StatusBar_UpdatingStatusBar", 0),
		/*39*/ imports.NewTable("StatusBar_UseSystemFont", 0),
	}
)

func statusBarImportAPI() *imports.Imports {
	if statusBarImport == nil {
		statusBarImport = NewDefaultImports()
		statusBarImport.SetImportTable(statusBarImportTables)
		statusBarImportTables = nil
	}
	return statusBarImport
}
