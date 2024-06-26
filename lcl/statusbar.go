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
	r1 := LCL().SysCallN(5106, GetObjectUintptr(TheOwner))
	return AsStatusBar(r1)
}

func (m *TStatusBar) Canvas() ICanvas {
	r1 := LCL().SysCallN(5104, m.Instance())
	return AsCanvas(r1)
}

func (m *TStatusBar) AutoHint() bool {
	r1 := LCL().SysCallN(5102, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetAutoHint(AValue bool) {
	LCL().SysCallN(5102, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) DragCursor() TCursor {
	r1 := LCL().SysCallN(5107, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStatusBar) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(5107, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) DragKind() TDragKind {
	r1 := LCL().SysCallN(5108, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TStatusBar) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(5108, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) DragMode() TDragMode {
	r1 := LCL().SysCallN(5109, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TStatusBar) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(5109, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) Panels() IStatusPanels {
	r1 := LCL().SysCallN(5113, 0, m.Instance(), 0)
	return AsStatusPanels(r1)
}

func (m *TStatusBar) SetPanels(AValue IStatusPanels) {
	LCL().SysCallN(5113, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TStatusBar) ParentColor() bool {
	r1 := LCL().SysCallN(5114, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentColor(AValue bool) {
	LCL().SysCallN(5114, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) ParentFont() bool {
	r1 := LCL().SysCallN(5115, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentFont(AValue bool) {
	LCL().SysCallN(5115, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) ParentShowHint() bool {
	r1 := LCL().SysCallN(5116, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5116, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) SimpleText() string {
	r1 := LCL().SysCallN(5137, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStatusBar) SetSimpleText(AValue string) {
	LCL().SysCallN(5137, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStatusBar) SimplePanel() bool {
	r1 := LCL().SysCallN(5136, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetSimplePanel(AValue bool) {
	LCL().SysCallN(5136, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) SizeGrip() bool {
	r1 := LCL().SysCallN(5138, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetSizeGrip(AValue bool) {
	LCL().SysCallN(5138, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) UseSystemFont() bool {
	r1 := LCL().SysCallN(5141, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetUseSystemFont(AValue bool) {
	LCL().SysCallN(5141, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) GetPanelIndexAt(X, Y int32) int32 {
	r1 := LCL().SysCallN(5111, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r1)
}

func (m *TStatusBar) SizeGripEnabled() bool {
	r1 := LCL().SysCallN(5139, m.Instance())
	return GoBool(r1)
}

func (m *TStatusBar) UpdatingStatusBar() bool {
	r1 := LCL().SysCallN(5140, m.Instance())
	return GoBool(r1)
}

func StatusBarClass() TClass {
	ret := LCL().SysCallN(5105)
	return TClass(ret)
}

func (m *TStatusBar) InvalidatePanel(PanelIndex int32, PanelParts TPanelParts) {
	LCL().SysCallN(5112, m.Instance(), uintptr(PanelIndex), uintptr(PanelParts))
}

func (m *TStatusBar) BeginUpdate() {
	LCL().SysCallN(5103, m.Instance())
}

func (m *TStatusBar) EndUpdate() {
	LCL().SysCallN(5110, m.Instance())
}

func (m *TStatusBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5117, m.Instance(), m.contextPopupPtr)
}

func (m *TStatusBar) SetOnCreatePanelClass(fn TSBCreatePanelClassEvent) {
	if m.createPanelClassPtr != 0 {
		RemoveEventElement(m.createPanelClassPtr)
	}
	m.createPanelClassPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5118, m.Instance(), m.createPanelClassPtr)
}

func (m *TStatusBar) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5119, m.Instance(), m.dblClickPtr)
}

func (m *TStatusBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5120, m.Instance(), m.dragDropPtr)
}

func (m *TStatusBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5121, m.Instance(), m.dragOverPtr)
}

func (m *TStatusBar) SetOnDrawPanel(fn TDrawPanelEvent) {
	if m.drawPanelPtr != 0 {
		RemoveEventElement(m.drawPanelPtr)
	}
	m.drawPanelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5122, m.Instance(), m.drawPanelPtr)
}

func (m *TStatusBar) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5123, m.Instance(), m.endDockPtr)
}

func (m *TStatusBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5124, m.Instance(), m.endDragPtr)
}

func (m *TStatusBar) SetOnHint(fn TNotifyEvent) {
	if m.hintPtr != 0 {
		RemoveEventElement(m.hintPtr)
	}
	m.hintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5125, m.Instance(), m.hintPtr)
}

func (m *TStatusBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5126, m.Instance(), m.mouseDownPtr)
}

func (m *TStatusBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5127, m.Instance(), m.mouseEnterPtr)
}

func (m *TStatusBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5128, m.Instance(), m.mouseLeavePtr)
}

func (m *TStatusBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5129, m.Instance(), m.mouseMovePtr)
}

func (m *TStatusBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5130, m.Instance(), m.mouseUpPtr)
}

func (m *TStatusBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5131, m.Instance(), m.mouseWheelPtr)
}

func (m *TStatusBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5132, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TStatusBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5133, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TStatusBar) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5134, m.Instance(), m.startDockPtr)
}

func (m *TStatusBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5135, m.Instance(), m.startDragPtr)
}
