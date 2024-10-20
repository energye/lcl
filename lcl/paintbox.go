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

// IPaintBox Parent: IGraphicControl
type IPaintBox interface {
	IGraphicControl
	DragCursor() TCursor                            // property
	SetDragCursor(AValue TCursor)                   // property
	DragMode() TDragMode                            // property
	SetDragMode(AValue TDragMode)                   // property
	ParentColor() bool                              // property
	SetParentColor(AValue bool)                     // property
	ParentFont() bool                               // property
	SetParentFont(AValue bool)                      // property
	ParentShowHint() bool                           // property
	SetParentShowHint(AValue bool)                  // property
	SetOnContextPopup(fn TContextPopupEvent)        // property event
	SetOnDblClick(fn TNotifyEvent)                  // property event
	SetOnDragDrop(fn TDragDropEvent)                // property event
	SetOnDragOver(fn TDragOverEvent)                // property event
	SetOnEndDrag(fn TEndDragEvent)                  // property event
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
	SetOnPaint(fn TNotifyEvent)                     // property event
	SetOnStartDrag(fn TStartDragEvent)              // property event
}

// TPaintBox Parent: TGraphicControl
type TPaintBox struct {
	TGraphicControl
	contextPopupPtr    uintptr
	dblClickPtr        uintptr
	dragDropPtr        uintptr
	dragOverPtr        uintptr
	endDragPtr         uintptr
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
	paintPtr           uintptr
	startDragPtr       uintptr
}

func NewPaintBox(AOwner IComponent) IPaintBox {
	r1 := paintBoxImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsPaintBox(r1)
}

func (m *TPaintBox) DragCursor() TCursor {
	r1 := paintBoxImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TPaintBox) SetDragCursor(AValue TCursor) {
	paintBoxImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TPaintBox) DragMode() TDragMode {
	r1 := paintBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TPaintBox) SetDragMode(AValue TDragMode) {
	paintBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TPaintBox) ParentColor() bool {
	r1 := paintBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPaintBox) SetParentColor(AValue bool) {
	paintBoxImportAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPaintBox) ParentFont() bool {
	r1 := paintBoxImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPaintBox) SetParentFont(AValue bool) {
	paintBoxImportAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TPaintBox) ParentShowHint() bool {
	r1 := paintBoxImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TPaintBox) SetParentShowHint(AValue bool) {
	paintBoxImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func PaintBoxClass() TClass {
	ret := paintBoxImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TPaintBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(7, m.Instance(), m.contextPopupPtr)
}

func (m *TPaintBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(8, m.Instance(), m.dblClickPtr)
}

func (m *TPaintBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(9, m.Instance(), m.dragDropPtr)
}

func (m *TPaintBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(10, m.Instance(), m.dragOverPtr)
}

func (m *TPaintBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(11, m.Instance(), m.endDragPtr)
}

func (m *TPaintBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(12, m.Instance(), m.mouseDownPtr)
}

func (m *TPaintBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(13, m.Instance(), m.mouseEnterPtr)
}

func (m *TPaintBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(14, m.Instance(), m.mouseLeavePtr)
}

func (m *TPaintBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(15, m.Instance(), m.mouseMovePtr)
}

func (m *TPaintBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(16, m.Instance(), m.mouseUpPtr)
}

func (m *TPaintBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(17, m.Instance(), m.mouseWheelPtr)
}

func (m *TPaintBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(18, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TPaintBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(22, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TPaintBox) SetOnMouseWheelHorz(fn TMouseWheelEvent) {
	if m.mouseWheelHorzPtr != 0 {
		RemoveEventElement(m.mouseWheelHorzPtr)
	}
	m.mouseWheelHorzPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(19, m.Instance(), m.mouseWheelHorzPtr)
}

func (m *TPaintBox) SetOnMouseWheelLeft(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelLeftPtr != 0 {
		RemoveEventElement(m.mouseWheelLeftPtr)
	}
	m.mouseWheelLeftPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(20, m.Instance(), m.mouseWheelLeftPtr)
}

func (m *TPaintBox) SetOnMouseWheelRight(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelRightPtr != 0 {
		RemoveEventElement(m.mouseWheelRightPtr)
	}
	m.mouseWheelRightPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(21, m.Instance(), m.mouseWheelRightPtr)
}

func (m *TPaintBox) SetOnPaint(fn TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(23, m.Instance(), m.paintPtr)
}

func (m *TPaintBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	paintBoxImportAPI().SysCallN(24, m.Instance(), m.startDragPtr)
}

var (
	paintBoxImport       *imports.Imports = nil
	paintBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("PaintBox_Class", 0),
		/*1*/ imports.NewTable("PaintBox_Create", 0),
		/*2*/ imports.NewTable("PaintBox_DragCursor", 0),
		/*3*/ imports.NewTable("PaintBox_DragMode", 0),
		/*4*/ imports.NewTable("PaintBox_ParentColor", 0),
		/*5*/ imports.NewTable("PaintBox_ParentFont", 0),
		/*6*/ imports.NewTable("PaintBox_ParentShowHint", 0),
		/*7*/ imports.NewTable("PaintBox_SetOnContextPopup", 0),
		/*8*/ imports.NewTable("PaintBox_SetOnDblClick", 0),
		/*9*/ imports.NewTable("PaintBox_SetOnDragDrop", 0),
		/*10*/ imports.NewTable("PaintBox_SetOnDragOver", 0),
		/*11*/ imports.NewTable("PaintBox_SetOnEndDrag", 0),
		/*12*/ imports.NewTable("PaintBox_SetOnMouseDown", 0),
		/*13*/ imports.NewTable("PaintBox_SetOnMouseEnter", 0),
		/*14*/ imports.NewTable("PaintBox_SetOnMouseLeave", 0),
		/*15*/ imports.NewTable("PaintBox_SetOnMouseMove", 0),
		/*16*/ imports.NewTable("PaintBox_SetOnMouseUp", 0),
		/*17*/ imports.NewTable("PaintBox_SetOnMouseWheel", 0),
		/*18*/ imports.NewTable("PaintBox_SetOnMouseWheelDown", 0),
		/*19*/ imports.NewTable("PaintBox_SetOnMouseWheelHorz", 0),
		/*20*/ imports.NewTable("PaintBox_SetOnMouseWheelLeft", 0),
		/*21*/ imports.NewTable("PaintBox_SetOnMouseWheelRight", 0),
		/*22*/ imports.NewTable("PaintBox_SetOnMouseWheelUp", 0),
		/*23*/ imports.NewTable("PaintBox_SetOnPaint", 0),
		/*24*/ imports.NewTable("PaintBox_SetOnStartDrag", 0),
	}
)

func paintBoxImportAPI() *imports.Imports {
	if paintBoxImport == nil {
		paintBoxImport = NewDefaultImports()
		paintBoxImport.SetImportTable(paintBoxImportTables)
		paintBoxImportTables = nil
	}
	return paintBoxImport
}
