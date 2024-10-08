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
	r1 := LCL().SysCallN(5149, GetObjectUintptr(TheOwner))
	return AsStatusBar(r1)
}

func (m *TStatusBar) Canvas() ICanvas {
	r1 := LCL().SysCallN(5147, m.Instance())
	return AsCanvas(r1)
}

func (m *TStatusBar) AutoHint() bool {
	r1 := LCL().SysCallN(5145, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetAutoHint(AValue bool) {
	LCL().SysCallN(5145, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) DragCursor() TCursor {
	r1 := LCL().SysCallN(5150, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TStatusBar) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(5150, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) DragKind() TDragKind {
	r1 := LCL().SysCallN(5151, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TStatusBar) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(5151, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) DragMode() TDragMode {
	r1 := LCL().SysCallN(5152, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TStatusBar) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(5152, 1, m.Instance(), uintptr(AValue))
}

func (m *TStatusBar) Panels() IStatusPanels {
	r1 := LCL().SysCallN(5156, 0, m.Instance(), 0)
	return AsStatusPanels(r1)
}

func (m *TStatusBar) SetPanels(AValue IStatusPanels) {
	LCL().SysCallN(5156, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TStatusBar) ParentColor() bool {
	r1 := LCL().SysCallN(5157, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentColor(AValue bool) {
	LCL().SysCallN(5157, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) ParentFont() bool {
	r1 := LCL().SysCallN(5158, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentFont(AValue bool) {
	LCL().SysCallN(5158, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) ParentShowHint() bool {
	r1 := LCL().SysCallN(5159, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetParentShowHint(AValue bool) {
	LCL().SysCallN(5159, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) SimpleText() string {
	r1 := LCL().SysCallN(5180, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TStatusBar) SetSimpleText(AValue string) {
	LCL().SysCallN(5180, 1, m.Instance(), PascalStr(AValue))
}

func (m *TStatusBar) SimplePanel() bool {
	r1 := LCL().SysCallN(5179, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetSimplePanel(AValue bool) {
	LCL().SysCallN(5179, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) SizeGrip() bool {
	r1 := LCL().SysCallN(5181, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetSizeGrip(AValue bool) {
	LCL().SysCallN(5181, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) UseSystemFont() bool {
	r1 := LCL().SysCallN(5184, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TStatusBar) SetUseSystemFont(AValue bool) {
	LCL().SysCallN(5184, 1, m.Instance(), PascalBool(AValue))
}

func (m *TStatusBar) GetPanelIndexAt(X, Y int32) int32 {
	r1 := LCL().SysCallN(5154, m.Instance(), uintptr(X), uintptr(Y))
	return int32(r1)
}

func (m *TStatusBar) SizeGripEnabled() bool {
	r1 := LCL().SysCallN(5182, m.Instance())
	return GoBool(r1)
}

func (m *TStatusBar) UpdatingStatusBar() bool {
	r1 := LCL().SysCallN(5183, m.Instance())
	return GoBool(r1)
}

func StatusBarClass() TClass {
	ret := LCL().SysCallN(5148)
	return TClass(ret)
}

func (m *TStatusBar) InvalidatePanel(PanelIndex int32, PanelParts TPanelParts) {
	LCL().SysCallN(5155, m.Instance(), uintptr(PanelIndex), uintptr(PanelParts))
}

func (m *TStatusBar) BeginUpdate() {
	LCL().SysCallN(5146, m.Instance())
}

func (m *TStatusBar) EndUpdate() {
	LCL().SysCallN(5153, m.Instance())
}

func (m *TStatusBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5160, m.Instance(), m.contextPopupPtr)
}

func (m *TStatusBar) SetOnCreatePanelClass(fn TSBCreatePanelClassEvent) {
	if m.createPanelClassPtr != 0 {
		RemoveEventElement(m.createPanelClassPtr)
	}
	m.createPanelClassPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5161, m.Instance(), m.createPanelClassPtr)
}

func (m *TStatusBar) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5162, m.Instance(), m.dblClickPtr)
}

func (m *TStatusBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5163, m.Instance(), m.dragDropPtr)
}

func (m *TStatusBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5164, m.Instance(), m.dragOverPtr)
}

func (m *TStatusBar) SetOnDrawPanel(fn TDrawPanelEvent) {
	if m.drawPanelPtr != 0 {
		RemoveEventElement(m.drawPanelPtr)
	}
	m.drawPanelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5165, m.Instance(), m.drawPanelPtr)
}

func (m *TStatusBar) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5166, m.Instance(), m.endDockPtr)
}

func (m *TStatusBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5167, m.Instance(), m.endDragPtr)
}

func (m *TStatusBar) SetOnHint(fn TNotifyEvent) {
	if m.hintPtr != 0 {
		RemoveEventElement(m.hintPtr)
	}
	m.hintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5168, m.Instance(), m.hintPtr)
}

func (m *TStatusBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5169, m.Instance(), m.mouseDownPtr)
}

func (m *TStatusBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5170, m.Instance(), m.mouseEnterPtr)
}

func (m *TStatusBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5171, m.Instance(), m.mouseLeavePtr)
}

func (m *TStatusBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5172, m.Instance(), m.mouseMovePtr)
}

func (m *TStatusBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5173, m.Instance(), m.mouseUpPtr)
}

func (m *TStatusBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5174, m.Instance(), m.mouseWheelPtr)
}

func (m *TStatusBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5175, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TStatusBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5176, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TStatusBar) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5177, m.Instance(), m.startDockPtr)
}

func (m *TStatusBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(5178, m.Instance(), m.startDragPtr)
}
