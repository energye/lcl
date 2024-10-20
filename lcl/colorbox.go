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

// IColorBox Parent: ICustomColorBox
type IColorBox interface {
	ICustomColorBox
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ItemHeight() int32                             // property
	SetItemHeight(AValue int32)                    // property
	ItemWidth() int32                              // property
	SetItemWidth(AValue int32)                     // property
	ParentColor() bool                             // property
	SetParentColor(AValue bool)                    // property
	ParentFont() bool                              // property
	SetParentFont(AValue bool)                     // property
	ParentShowHint() bool                          // property
	SetParentShowHint(AValue bool)                 // property
	SetOnChange(fn TNotifyEvent)                   // property event
	SetOnCloseUp(fn TNotifyEvent)                  // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
	SetOnSelect(fn TNotifyEvent)                   // property event
}

// TColorBox Parent: TCustomColorBox
type TColorBox struct {
	TCustomColorBox
	changePtr         uintptr
	closeUpPtr        uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDragPtr        uintptr
	dropDownPtr       uintptr
	editingDonePtr    uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	startDragPtr      uintptr
	selectPtr         uintptr
}

func NewColorBox(AOwner IComponent) IColorBox {
	r1 := colorBoxImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsColorBox(r1)
}

func (m *TColorBox) DragCursor() TCursor {
	r1 := colorBoxImportAPI().SysCallN(2, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TColorBox) SetDragCursor(AValue TCursor) {
	colorBoxImportAPI().SysCallN(2, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorBox) DragMode() TDragMode {
	r1 := colorBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TColorBox) SetDragMode(AValue TDragMode) {
	colorBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorBox) ItemHeight() int32 {
	r1 := colorBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TColorBox) SetItemHeight(AValue int32) {
	colorBoxImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorBox) ItemWidth() int32 {
	r1 := colorBoxImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TColorBox) SetItemWidth(AValue int32) {
	colorBoxImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TColorBox) ParentColor() bool {
	r1 := colorBoxImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorBox) SetParentColor(AValue bool) {
	colorBoxImportAPI().SysCallN(6, 1, m.Instance(), PascalBool(AValue))
}

func (m *TColorBox) ParentFont() bool {
	r1 := colorBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorBox) SetParentFont(AValue bool) {
	colorBoxImportAPI().SysCallN(7, 1, m.Instance(), PascalBool(AValue))
}

func (m *TColorBox) ParentShowHint() bool {
	r1 := colorBoxImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TColorBox) SetParentShowHint(AValue bool) {
	colorBoxImportAPI().SysCallN(8, 1, m.Instance(), PascalBool(AValue))
}

func ColorBoxClass() TClass {
	ret := colorBoxImportAPI().SysCallN(0)
	return TClass(ret)
}

func (m *TColorBox) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(9, m.Instance(), m.changePtr)
}

func (m *TColorBox) SetOnCloseUp(fn TNotifyEvent) {
	if m.closeUpPtr != 0 {
		RemoveEventElement(m.closeUpPtr)
	}
	m.closeUpPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(10, m.Instance(), m.closeUpPtr)
}

func (m *TColorBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(11, m.Instance(), m.contextPopupPtr)
}

func (m *TColorBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(12, m.Instance(), m.dblClickPtr)
}

func (m *TColorBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(13, m.Instance(), m.dragDropPtr)
}

func (m *TColorBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(14, m.Instance(), m.dragOverPtr)
}

func (m *TColorBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(17, m.Instance(), m.endDragPtr)
}

func (m *TColorBox) SetOnDropDown(fn TNotifyEvent) {
	if m.dropDownPtr != 0 {
		RemoveEventElement(m.dropDownPtr)
	}
	m.dropDownPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(15, m.Instance(), m.dropDownPtr)
}

func (m *TColorBox) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(16, m.Instance(), m.editingDonePtr)
}

func (m *TColorBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(18, m.Instance(), m.mouseDownPtr)
}

func (m *TColorBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(19, m.Instance(), m.mouseEnterPtr)
}

func (m *TColorBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(20, m.Instance(), m.mouseLeavePtr)
}

func (m *TColorBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(21, m.Instance(), m.mouseMovePtr)
}

func (m *TColorBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(22, m.Instance(), m.mouseUpPtr)
}

func (m *TColorBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(23, m.Instance(), m.mouseWheelPtr)
}

func (m *TColorBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(24, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TColorBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(25, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TColorBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(27, m.Instance(), m.startDragPtr)
}

func (m *TColorBox) SetOnSelect(fn TNotifyEvent) {
	if m.selectPtr != 0 {
		RemoveEventElement(m.selectPtr)
	}
	m.selectPtr = MakeEventDataPtr(fn)
	colorBoxImportAPI().SysCallN(26, m.Instance(), m.selectPtr)
}

var (
	colorBoxImport       *imports.Imports = nil
	colorBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ColorBox_Class", 0),
		/*1*/ imports.NewTable("ColorBox_Create", 0),
		/*2*/ imports.NewTable("ColorBox_DragCursor", 0),
		/*3*/ imports.NewTable("ColorBox_DragMode", 0),
		/*4*/ imports.NewTable("ColorBox_ItemHeight", 0),
		/*5*/ imports.NewTable("ColorBox_ItemWidth", 0),
		/*6*/ imports.NewTable("ColorBox_ParentColor", 0),
		/*7*/ imports.NewTable("ColorBox_ParentFont", 0),
		/*8*/ imports.NewTable("ColorBox_ParentShowHint", 0),
		/*9*/ imports.NewTable("ColorBox_SetOnChange", 0),
		/*10*/ imports.NewTable("ColorBox_SetOnCloseUp", 0),
		/*11*/ imports.NewTable("ColorBox_SetOnContextPopup", 0),
		/*12*/ imports.NewTable("ColorBox_SetOnDblClick", 0),
		/*13*/ imports.NewTable("ColorBox_SetOnDragDrop", 0),
		/*14*/ imports.NewTable("ColorBox_SetOnDragOver", 0),
		/*15*/ imports.NewTable("ColorBox_SetOnDropDown", 0),
		/*16*/ imports.NewTable("ColorBox_SetOnEditingDone", 0),
		/*17*/ imports.NewTable("ColorBox_SetOnEndDrag", 0),
		/*18*/ imports.NewTable("ColorBox_SetOnMouseDown", 0),
		/*19*/ imports.NewTable("ColorBox_SetOnMouseEnter", 0),
		/*20*/ imports.NewTable("ColorBox_SetOnMouseLeave", 0),
		/*21*/ imports.NewTable("ColorBox_SetOnMouseMove", 0),
		/*22*/ imports.NewTable("ColorBox_SetOnMouseUp", 0),
		/*23*/ imports.NewTable("ColorBox_SetOnMouseWheel", 0),
		/*24*/ imports.NewTable("ColorBox_SetOnMouseWheelDown", 0),
		/*25*/ imports.NewTable("ColorBox_SetOnMouseWheelUp", 0),
		/*26*/ imports.NewTable("ColorBox_SetOnSelect", 0),
		/*27*/ imports.NewTable("ColorBox_SetOnStartDrag", 0),
	}
)

func colorBoxImportAPI() *imports.Imports {
	if colorBoxImport == nil {
		colorBoxImport = NewDefaultImports()
		colorBoxImport.SetImportTable(colorBoxImportTables)
		colorBoxImportTables = nil
	}
	return colorBoxImport
}
