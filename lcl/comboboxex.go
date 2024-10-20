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

// IComboBoxEx Parent: ICustomComboBoxEx
type IComboBoxEx interface {
	ICustomComboBoxEx
	BorderStyle() TBorderStyle                     // property
	SetBorderStyle(AValue TBorderStyle)            // property
	DragCursor() TCursor                           // property
	SetDragCursor(AValue TCursor)                  // property
	DragKind() TDragKind                           // property
	SetDragKind(AValue TDragKind)                  // property
	DragMode() TDragMode                           // property
	SetDragMode(AValue TDragMode)                  // property
	ItemHeight() int32                             // property
	SetItemHeight(AValue int32)                    // property
	ItemWidth() int32                              // property
	SetItemWidth(AValue int32)                     // property
	MaxLength() int32                              // property
	SetMaxLength(AValue int32)                     // property
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
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnEndDock(fn TEndDragEvent)                 // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnGetItems(fn TNotifyEvent)                 // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnSelect(fn TNotifyEvent)                   // property event
	SetOnStartDock(fn TStartDockEvent)             // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TComboBoxEx Parent: TCustomComboBoxEx
type TComboBoxEx struct {
	TCustomComboBoxEx
	changePtr         uintptr
	closeUpPtr        uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	dropDownPtr       uintptr
	editingDonePtr    uintptr
	endDockPtr        uintptr
	endDragPtr        uintptr
	getItemsPtr       uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	selectPtr         uintptr
	startDockPtr      uintptr
	startDragPtr      uintptr
}

func NewComboBoxEx(TheOwner IComponent) IComboBoxEx {
	r1 := comboBoxExImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsComboBoxEx(r1)
}

func (m *TComboBoxEx) BorderStyle() TBorderStyle {
	r1 := comboBoxExImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TComboBoxEx) SetBorderStyle(AValue TBorderStyle) {
	comboBoxExImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) DragCursor() TCursor {
	r1 := comboBoxExImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TComboBoxEx) SetDragCursor(AValue TCursor) {
	comboBoxExImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) DragKind() TDragKind {
	r1 := comboBoxExImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TComboBoxEx) SetDragKind(AValue TDragKind) {
	comboBoxExImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) DragMode() TDragMode {
	r1 := comboBoxExImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TComboBoxEx) SetDragMode(AValue TDragMode) {
	comboBoxExImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) ItemHeight() int32 {
	r1 := comboBoxExImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBoxEx) SetItemHeight(AValue int32) {
	comboBoxExImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) ItemWidth() int32 {
	r1 := comboBoxExImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBoxEx) SetItemWidth(AValue int32) {
	comboBoxExImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) MaxLength() int32 {
	r1 := comboBoxExImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBoxEx) SetMaxLength(AValue int32) {
	comboBoxExImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBoxEx) ParentColor() bool {
	r1 := comboBoxExImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBoxEx) SetParentColor(AValue bool) {
	comboBoxExImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TComboBoxEx) ParentFont() bool {
	r1 := comboBoxExImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBoxEx) SetParentFont(AValue bool) {
	comboBoxExImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TComboBoxEx) ParentShowHint() bool {
	r1 := comboBoxExImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBoxEx) SetParentShowHint(AValue bool) {
	comboBoxExImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func ComboBoxExClass() TClass {
	ret := comboBoxExImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TComboBoxEx) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(12, m.Instance(), m.changePtr)
}

func (m *TComboBoxEx) SetOnCloseUp(fn TNotifyEvent) {
	if m.closeUpPtr != 0 {
		RemoveEventElement(m.closeUpPtr)
	}
	m.closeUpPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(13, m.Instance(), m.closeUpPtr)
}

func (m *TComboBoxEx) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(14, m.Instance(), m.contextPopupPtr)
}

func (m *TComboBoxEx) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(15, m.Instance(), m.dblClickPtr)
}

func (m *TComboBoxEx) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(16, m.Instance(), m.dragDropPtr)
}

func (m *TComboBoxEx) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(17, m.Instance(), m.dragOverPtr)
}

func (m *TComboBoxEx) SetOnDropDown(fn TNotifyEvent) {
	if m.dropDownPtr != 0 {
		RemoveEventElement(m.dropDownPtr)
	}
	m.dropDownPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(18, m.Instance(), m.dropDownPtr)
}

func (m *TComboBoxEx) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(19, m.Instance(), m.editingDonePtr)
}

func (m *TComboBoxEx) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(20, m.Instance(), m.endDockPtr)
}

func (m *TComboBoxEx) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(21, m.Instance(), m.endDragPtr)
}

func (m *TComboBoxEx) SetOnGetItems(fn TNotifyEvent) {
	if m.getItemsPtr != 0 {
		RemoveEventElement(m.getItemsPtr)
	}
	m.getItemsPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(22, m.Instance(), m.getItemsPtr)
}

func (m *TComboBoxEx) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(23, m.Instance(), m.mouseDownPtr)
}

func (m *TComboBoxEx) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(24, m.Instance(), m.mouseEnterPtr)
}

func (m *TComboBoxEx) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(25, m.Instance(), m.mouseLeavePtr)
}

func (m *TComboBoxEx) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(26, m.Instance(), m.mouseMovePtr)
}

func (m *TComboBoxEx) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(27, m.Instance(), m.mouseUpPtr)
}

func (m *TComboBoxEx) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(28, m.Instance(), m.mouseWheelPtr)
}

func (m *TComboBoxEx) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(29, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TComboBoxEx) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(30, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TComboBoxEx) SetOnSelect(fn TNotifyEvent) {
	if m.selectPtr != 0 {
		RemoveEventElement(m.selectPtr)
	}
	m.selectPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(31, m.Instance(), m.selectPtr)
}

func (m *TComboBoxEx) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(32, m.Instance(), m.startDockPtr)
}

func (m *TComboBoxEx) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	comboBoxExImportAPI().SysCallN(33, m.Instance(), m.startDragPtr)
}

var (
	comboBoxExImport       *imports.Imports = nil
	comboBoxExImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ComboBoxEx_BorderStyle", 0),
		/*1*/ imports.NewTable("ComboBoxEx_Class", 0),
		/*2*/ imports.NewTable("ComboBoxEx_Create", 0),
		/*3*/ imports.NewTable("ComboBoxEx_DragCursor", 0),
		/*4*/ imports.NewTable("ComboBoxEx_DragKind", 0),
		/*5*/ imports.NewTable("ComboBoxEx_DragMode", 0),
		/*6*/ imports.NewTable("ComboBoxEx_ItemHeight", 0),
		/*7*/ imports.NewTable("ComboBoxEx_ItemWidth", 0),
		/*8*/ imports.NewTable("ComboBoxEx_MaxLength", 0),
		/*9*/ imports.NewTable("ComboBoxEx_ParentColor", 0),
		/*10*/ imports.NewTable("ComboBoxEx_ParentFont", 0),
		/*11*/ imports.NewTable("ComboBoxEx_ParentShowHint", 0),
		/*12*/ imports.NewTable("ComboBoxEx_SetOnChange", 0),
		/*13*/ imports.NewTable("ComboBoxEx_SetOnCloseUp", 0),
		/*14*/ imports.NewTable("ComboBoxEx_SetOnContextPopup", 0),
		/*15*/ imports.NewTable("ComboBoxEx_SetOnDblClick", 0),
		/*16*/ imports.NewTable("ComboBoxEx_SetOnDragDrop", 0),
		/*17*/ imports.NewTable("ComboBoxEx_SetOnDragOver", 0),
		/*18*/ imports.NewTable("ComboBoxEx_SetOnDropDown", 0),
		/*19*/ imports.NewTable("ComboBoxEx_SetOnEditingDone", 0),
		/*20*/ imports.NewTable("ComboBoxEx_SetOnEndDock", 0),
		/*21*/ imports.NewTable("ComboBoxEx_SetOnEndDrag", 0),
		/*22*/ imports.NewTable("ComboBoxEx_SetOnGetItems", 0),
		/*23*/ imports.NewTable("ComboBoxEx_SetOnMouseDown", 0),
		/*24*/ imports.NewTable("ComboBoxEx_SetOnMouseEnter", 0),
		/*25*/ imports.NewTable("ComboBoxEx_SetOnMouseLeave", 0),
		/*26*/ imports.NewTable("ComboBoxEx_SetOnMouseMove", 0),
		/*27*/ imports.NewTable("ComboBoxEx_SetOnMouseUp", 0),
		/*28*/ imports.NewTable("ComboBoxEx_SetOnMouseWheel", 0),
		/*29*/ imports.NewTable("ComboBoxEx_SetOnMouseWheelDown", 0),
		/*30*/ imports.NewTable("ComboBoxEx_SetOnMouseWheelUp", 0),
		/*31*/ imports.NewTable("ComboBoxEx_SetOnSelect", 0),
		/*32*/ imports.NewTable("ComboBoxEx_SetOnStartDock", 0),
		/*33*/ imports.NewTable("ComboBoxEx_SetOnStartDrag", 0),
	}
)

func comboBoxExImportAPI() *imports.Imports {
	if comboBoxExImport == nil {
		comboBoxExImport = NewDefaultImports()
		comboBoxExImport.SetImportTable(comboBoxExImportTables)
		comboBoxExImportTables = nil
	}
	return comboBoxExImport
}
