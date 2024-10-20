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

// ISpeedButton Parent: ICustomSpeedButton
type ISpeedButton interface {
	ICustomSpeedButton
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragKind() TDragKind                           // property
	SetDragKind(AValue TDragKind)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnPaint(fn TNotifyEvent)                    // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TSpeedButton Parent: TCustomSpeedButton
type TSpeedButton struct {
	TCustomSpeedButton
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDragPtr        uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	paintPtr          uintptr
	startDragPtr      uintptr
}

func NewSpeedButton(AOwner IComponent) ISpeedButton {
	r1 := speedButtonImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsSpeedButton(r1)
}

func (m *TSpeedButton) DragCursor() TCursor {
	r1 := speedButtonImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TSpeedButton) SetDragCursor(AValue TCursor) {
	speedButtonImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TSpeedButton) DragKind() TDragKind {
	r1 := speedButtonImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TSpeedButton) SetDragKind(AValue TDragKind) {
	speedButtonImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TSpeedButton) DragMode() TDragMode {
	r1 := speedButtonImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TSpeedButton) SetDragMode(AValue TDragMode) {
	speedButtonImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TSpeedButton) ParentFont() bool {
	r1 := speedButtonImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpeedButton) SetParentFont(AValue bool) {
	speedButtonImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TSpeedButton) ParentShowHint() bool {
	r1 := speedButtonImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TSpeedButton) SetParentShowHint(AValue bool) {
	speedButtonImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func SpeedButtonClass() TClass {
	ret := speedButtonImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TSpeedButton) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(7, m.Instance(), m.contextPopupPtr)
}

func (m *TSpeedButton) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(8, m.Instance(), m.dblClickPtr)
}

func (m *TSpeedButton) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(9, m.Instance(), m.dragDropPtr)
}

func (m *TSpeedButton) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(10, m.Instance(), m.dragOverPtr)
}

func (m *TSpeedButton) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(11, m.Instance(), m.endDragPtr)
}

func (m *TSpeedButton) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(12, m.Instance(), m.mouseDownPtr)
}

func (m *TSpeedButton) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(13, m.Instance(), m.mouseEnterPtr)
}

func (m *TSpeedButton) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(14, m.Instance(), m.mouseLeavePtr)
}

func (m *TSpeedButton) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(15, m.Instance(), m.mouseMovePtr)
}

func (m *TSpeedButton) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(16, m.Instance(), m.mouseUpPtr)
}

func (m *TSpeedButton) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(17, m.Instance(), m.mouseWheelPtr)
}

func (m *TSpeedButton) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(18, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TSpeedButton) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(19, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TSpeedButton) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(20, m.Instance(), m.paintPtr)
}

func (m *TSpeedButton) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	speedButtonImportAPI().SysCallN(21, m.Instance(), m.startDragPtr)
}

var (
	speedButtonImport       *imports.Imports = nil
	speedButtonImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("SpeedButton_Class", 0),
		/*1*/ imports.NewTable("SpeedButton_Create", 0),
		/*2*/ imports.NewTable("SpeedButton_DragCursor", 0),
		/*3*/ imports.NewTable("SpeedButton_DragKind", 0),
		/*4*/ imports.NewTable("SpeedButton_DragMode", 0),
		/*5*/ imports.NewTable("SpeedButton_ParentFont", 0),
		/*6*/ imports.NewTable("SpeedButton_ParentShowHint", 0),
		/*7*/ imports.NewTable("SpeedButton_SetOnContextPopup", 0),
		/*8*/ imports.NewTable("SpeedButton_SetOnDblClick", 0),
		/*9*/ imports.NewTable("SpeedButton_SetOnDragDrop", 0),
		/*10*/ imports.NewTable("SpeedButton_SetOnDragOver", 0),
		/*11*/ imports.NewTable("SpeedButton_SetOnEndDrag", 0),
		/*12*/ imports.NewTable("SpeedButton_SetOnMouseDown", 0),
		/*13*/ imports.NewTable("SpeedButton_SetOnMouseEnter", 0),
		/*14*/ imports.NewTable("SpeedButton_SetOnMouseLeave", 0),
		/*15*/ imports.NewTable("SpeedButton_SetOnMouseMove", 0),
		/*16*/ imports.NewTable("SpeedButton_SetOnMouseUp", 0),
		/*17*/ imports.NewTable("SpeedButton_SetOnMouseWheel", 0),
		/*18*/ imports.NewTable("SpeedButton_SetOnMouseWheelDown", 0),
		/*19*/ imports.NewTable("SpeedButton_SetOnMouseWheelUp", 0),
		/*20*/ imports.NewTable("SpeedButton_SetOnPaint", 0),
		/*21*/ imports.NewTable("SpeedButton_SetOnStartDrag", 0),
	}
)

func speedButtonImportAPI() *imports.Imports {
	if speedButtonImport == nil {
		speedButtonImport = NewDefaultImports()
		speedButtonImport.SetImportTable(speedButtonImportTables)
		speedButtonImportTables = nil
	}
	return speedButtonImport
}
