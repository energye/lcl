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

// IBoundLabel Parent: ICustomLabel
type IBoundLabel interface {
	ICustomLabel
	FocusControl() IWinControl                     // property
	SetFocusControl(AValue IWinControl)            // property
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	ShowAccelChar() bool                           // property
	SetShowAccelChar(AValue bool)                  // property
	Layout() TTextLayout                           // property
	SetLayout(AValue TTextLayout)                  // property
	WordWrap() bool                                // property
	SetWordWrap(AValue bool)                       // property
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
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TBoundLabel Parent: TCustomLabel
type TBoundLabel struct {
	TCustomLabel
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
	startDragPtr      uintptr
}

func NewBoundLabel(TheOwner IComponent) IBoundLabel {
	r1 := boundLabelImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsBoundLabel(r1)
}

func (m *TBoundLabel) FocusControl() IWinControl {
	r1 := boundLabelImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return AsWinControl(r1)
}

func (m *TBoundLabel) SetFocusControl(AValue IWinControl) {
	boundLabelImportAPI().SysCallN(4, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TBoundLabel) DragCursor() TCursor {
	r1 := boundLabelImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TBoundLabel) SetDragCursor(AValue TCursor) {
	boundLabelImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TBoundLabel) DragMode() TDragMode {
	r1 := boundLabelImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TBoundLabel) SetDragMode(AValue TDragMode) {
	boundLabelImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TBoundLabel) ParentColor() bool {
	r1 := boundLabelImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBoundLabel) SetParentColor(AValue bool) {
	boundLabelImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBoundLabel) ParentFont() bool {
	r1 := boundLabelImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBoundLabel) SetParentFont(AValue bool) {
	boundLabelImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBoundLabel) ParentShowHint() bool {
	r1 := boundLabelImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBoundLabel) SetParentShowHint(AValue bool) {
	boundLabelImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBoundLabel) ShowAccelChar() bool {
	r1 := boundLabelImportAPI().SysCallN(22, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBoundLabel) SetShowAccelChar(AValue bool) {
	boundLabelImportAPI().SysCallN(22, 1, m.Instance(), PascalBool(AValue))
}

func (m *TBoundLabel) Layout() TTextLayout {
	r1 := boundLabelImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TTextLayout(r1)
}

func (m *TBoundLabel) SetLayout(AValue TTextLayout) {
	boundLabelImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TBoundLabel) WordWrap() bool {
	r1 := boundLabelImportAPI().SysCallN(23, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TBoundLabel) SetWordWrap(AValue bool) {
	boundLabelImportAPI().SysCallN(23, 1, m.Instance(), PascalBool(AValue))
}

func BoundLabelClass() TClass {
	ret := boundLabelImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TBoundLabel) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(9, m.Instance(), m.dblClickPtr)
}

func (m *TBoundLabel) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(10, m.Instance(), m.dragDropPtr)
}

func (m *TBoundLabel) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(11, m.Instance(), m.dragOverPtr)
}

func (m *TBoundLabel) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(12, m.Instance(), m.endDragPtr)
}

func (m *TBoundLabel) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(13, m.Instance(), m.mouseDownPtr)
}

func (m *TBoundLabel) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(14, m.Instance(), m.mouseEnterPtr)
}

func (m *TBoundLabel) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(15, m.Instance(), m.mouseLeavePtr)
}

func (m *TBoundLabel) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(16, m.Instance(), m.mouseMovePtr)
}

func (m *TBoundLabel) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(17, m.Instance(), m.mouseUpPtr)
}

func (m *TBoundLabel) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(18, m.Instance(), m.mouseWheelPtr)
}

func (m *TBoundLabel) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(19, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TBoundLabel) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(20, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TBoundLabel) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	boundLabelImportAPI().SysCallN(21, m.Instance(), m.startDragPtr)
}

var (
	boundLabelImport       *imports.Imports = nil
	boundLabelImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("BoundLabel_Class", 0),
		/*1*/ imports.NewTable("BoundLabel_Create", 0),
		/*2*/ imports.NewTable("BoundLabel_DragCursor", 0),
		/*3*/ imports.NewTable("BoundLabel_DragMode", 0),
		/*4*/ imports.NewTable("BoundLabel_FocusControl", 0),
		/*5*/ imports.NewTable("BoundLabel_Layout", 0),
		/*6*/ imports.NewTable("BoundLabel_ParentColor", 0),
		/*7*/ imports.NewTable("BoundLabel_ParentFont", 0),
		/*8*/ imports.NewTable("BoundLabel_ParentShowHint", 0),
		/*9*/ imports.NewTable("BoundLabel_SetOnDblClick", 0),
		/*10*/ imports.NewTable("BoundLabel_SetOnDragDrop", 0),
		/*11*/ imports.NewTable("BoundLabel_SetOnDragOver", 0),
		/*12*/ imports.NewTable("BoundLabel_SetOnEndDrag", 0),
		/*13*/ imports.NewTable("BoundLabel_SetOnMouseDown", 0),
		/*14*/ imports.NewTable("BoundLabel_SetOnMouseEnter", 0),
		/*15*/ imports.NewTable("BoundLabel_SetOnMouseLeave", 0),
		/*16*/ imports.NewTable("BoundLabel_SetOnMouseMove", 0),
		/*17*/ imports.NewTable("BoundLabel_SetOnMouseUp", 0),
		/*18*/ imports.NewTable("BoundLabel_SetOnMouseWheel", 0),
		/*19*/ imports.NewTable("BoundLabel_SetOnMouseWheelDown", 0),
		/*20*/ imports.NewTable("BoundLabel_SetOnMouseWheelUp", 0),
		/*21*/ imports.NewTable("BoundLabel_SetOnStartDrag", 0),
		/*22*/ imports.NewTable("BoundLabel_ShowAccelChar", 0),
		/*23*/ imports.NewTable("BoundLabel_WordWrap", 0),
	}
)

func boundLabelImportAPI() *imports.Imports {
	if boundLabelImport == nil {
		boundLabelImport = NewDefaultImports()
		boundLabelImport.SetImportTable(boundLabelImportTables)
		boundLabelImportTables = nil
	}
	return boundLabelImport
}
