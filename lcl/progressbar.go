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

// IProgressBar Parent: ICustomProgressBar
type IProgressBar interface {
	ICustomProgressBar
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
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TProgressBar Parent: TCustomProgressBar
type TProgressBar struct {
	TCustomProgressBar
	contextPopupPtr   uintptr
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
	startDockPtr      uintptr
	startDragPtr      uintptr
}

func NewProgressBar(AOwner IComponent) IProgressBar {
	r1 := progressBarImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsProgressBar(r1)
}

func (m *TProgressBar) DragCursor() TCursor {
	r1 := progressBarImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TProgressBar) SetDragCursor(AValue TCursor) {
	progressBarImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TProgressBar) DragKind() TDragKind {
	r1 := progressBarImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TProgressBar) SetDragKind(AValue TDragKind) {
	progressBarImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TProgressBar) DragMode() TDragMode {
	r1 := progressBarImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TProgressBar) SetDragMode(AValue TDragMode) {
	progressBarImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TProgressBar) ParentColor() bool {
	r1 := progressBarImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TProgressBar) SetParentColor(AValue bool) {
	progressBarImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TProgressBar) ParentFont() bool {
	r1 := progressBarImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TProgressBar) SetParentFont(AValue bool) {
	progressBarImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TProgressBar) ParentShowHint() bool {
	r1 := progressBarImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TProgressBar) SetParentShowHint(AValue bool) {
	progressBarImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func ProgressBarClass() TClass {
	ret := progressBarImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TProgressBar) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(8, m.Instance(), m.contextPopupPtr)
}

func (m *TProgressBar) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(9, m.Instance(), m.dragDropPtr)
}

func (m *TProgressBar) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(10, m.Instance(), m.dragOverPtr)
}

func (m *TProgressBar) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(11, m.Instance(), m.endDragPtr)
}

func (m *TProgressBar) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(12, m.Instance(), m.mouseDownPtr)
}

func (m *TProgressBar) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(13, m.Instance(), m.mouseEnterPtr)
}

func (m *TProgressBar) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(14, m.Instance(), m.mouseLeavePtr)
}

func (m *TProgressBar) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(15, m.Instance(), m.mouseMovePtr)
}

func (m *TProgressBar) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(16, m.Instance(), m.mouseUpPtr)
}

func (m *TProgressBar) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(17, m.Instance(), m.mouseWheelPtr)
}

func (m *TProgressBar) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(18, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TProgressBar) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(19, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TProgressBar) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(20, m.Instance(), m.startDockPtr)
}

func (m *TProgressBar) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	progressBarImportAPI().SysCallN(21, m.Instance(), m.startDragPtr)
}

var (
	progressBarImport       *imports.Imports = nil
	progressBarImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ProgressBar_Class", 0),
		/*1*/ imports.NewTable("ProgressBar_Create", 0),
		/*2*/ imports.NewTable("ProgressBar_DragCursor", 0),
		/*3*/ imports.NewTable("ProgressBar_DragKind", 0),
		/*4*/ imports.NewTable("ProgressBar_DragMode", 0),
		/*5*/ imports.NewTable("ProgressBar_ParentColor", 0),
		/*6*/ imports.NewTable("ProgressBar_ParentFont", 0),
		/*7*/ imports.NewTable("ProgressBar_ParentShowHint", 0),
		/*8*/ imports.NewTable("ProgressBar_SetOnContextPopup", 0),
		/*9*/ imports.NewTable("ProgressBar_SetOnDragDrop", 0),
		/*10*/ imports.NewTable("ProgressBar_SetOnDragOver", 0),
		/*11*/ imports.NewTable("ProgressBar_SetOnEndDrag", 0),
		/*12*/ imports.NewTable("ProgressBar_SetOnMouseDown", 0),
		/*13*/ imports.NewTable("ProgressBar_SetOnMouseEnter", 0),
		/*14*/ imports.NewTable("ProgressBar_SetOnMouseLeave", 0),
		/*15*/ imports.NewTable("ProgressBar_SetOnMouseMove", 0),
		/*16*/ imports.NewTable("ProgressBar_SetOnMouseUp", 0),
		/*17*/ imports.NewTable("ProgressBar_SetOnMouseWheel", 0),
		/*18*/ imports.NewTable("ProgressBar_SetOnMouseWheelDown", 0),
		/*19*/ imports.NewTable("ProgressBar_SetOnMouseWheelUp", 0),
		/*20*/ imports.NewTable("ProgressBar_SetOnStartDock", 0),
		/*21*/ imports.NewTable("ProgressBar_SetOnStartDrag", 0),
	}
)

func progressBarImportAPI() *imports.Imports {
	if progressBarImport == nil {
		progressBarImport = NewDefaultImports()
		progressBarImport.SetImportTable(progressBarImportTables)
		progressBarImportTables = nil
	}
	return progressBarImport
}
