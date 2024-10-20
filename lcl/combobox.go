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

// IComboBox Parent: ICustomComboBox
type IComboBox interface {
	ICustomComboBox
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
	Sorted() bool                                  // property
	SetSorted(AValue bool)                         // property
	SetOnChange(fn TNotifyEvent)                   // property event
	SetOnCloseUp(fn TNotifyEvent)                  // property event
	SetOnContextPopup(fn TContextPopupEvent)       // property event
	SetOnDblClick(fn TNotifyEvent)                 // property event
	SetOnDragDrop(fn TDragDropEvent)               // property event
	SetOnDragOver(fn TDragOverEvent)               // property event
	SetOnDrawItem(fn TDrawItemEvent)               // property event
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnGetItems(fn TNotifyEvent)                 // property event
	SetOnMeasureItem(fn TMeasureItemEvent)         // property event
	SetOnMouseDown(fn TMouseEvent)                 // property event
	SetOnMouseEnter(fn TNotifyEvent)               // property event
	SetOnMouseLeave(fn TNotifyEvent)               // property event
	SetOnMouseMove(fn TMouseMoveEvent)             // property event
	SetOnMouseUp(fn TMouseEvent)                   // property event
	SetOnMouseWheel(fn TMouseWheelEvent)           // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)   // property event
	SetOnSelect(fn TNotifyEvent)                   // property event
	SetOnStartDrag(fn TStartDragEvent)             // property event
}

// TComboBox Parent: TCustomComboBox
type TComboBox struct {
	TCustomComboBox
	changePtr         uintptr
	closeUpPtr        uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	drawItemPtr       uintptr
	endDragPtr        uintptr
	dropDownPtr       uintptr
	editingDonePtr    uintptr
	getItemsPtr       uintptr
	measureItemPtr    uintptr
	mouseDownPtr      uintptr
	mouseEnterPtr     uintptr
	mouseLeavePtr     uintptr
	mouseMovePtr      uintptr
	mouseUpPtr        uintptr
	mouseWheelPtr     uintptr
	mouseWheelDownPtr uintptr
	mouseWheelUpPtr   uintptr
	selectPtr         uintptr
	startDragPtr      uintptr
}

func NewComboBox(TheOwner IComponent) IComboBox {
	r1 := comboBoxImportAPI().SysCallN(2, GetObjectUintptr(TheOwner))
	return AsComboBox(r1)
}

func (m *TComboBox) BorderStyle() TBorderStyle {
	r1 := comboBoxImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TComboBox) SetBorderStyle(AValue TBorderStyle) {
	comboBoxImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) DragCursor() TCursor {
	r1 := comboBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TComboBox) SetDragCursor(AValue TCursor) {
	comboBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) DragKind() TDragKind {
	r1 := comboBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TComboBox) SetDragKind(AValue TDragKind) {
	comboBoxImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) DragMode() TDragMode {
	r1 := comboBoxImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TComboBox) SetDragMode(AValue TDragMode) {
	comboBoxImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) ItemHeight() int32 {
	r1 := comboBoxImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBox) SetItemHeight(AValue int32) {
	comboBoxImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) ItemWidth() int32 {
	r1 := comboBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBox) SetItemWidth(AValue int32) {
	comboBoxImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) MaxLength() int32 {
	r1 := comboBoxImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TComboBox) SetMaxLength(AValue int32) {
	comboBoxImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TComboBox) ParentColor() bool {
	r1 := comboBoxImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBox) SetParentColor(AValue bool) {
	comboBoxImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TComboBox) ParentFont() bool {
	r1 := comboBoxImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBox) SetParentFont(AValue bool) {
	comboBoxImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TComboBox) ParentShowHint() bool {
	r1 := comboBoxImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBox) SetParentShowHint(AValue bool) {
	comboBoxImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TComboBox) Sorted() bool {
	r1 := comboBoxImportAPI().SysCallN(34, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TComboBox) SetSorted(AValue bool) {
	comboBoxImportAPI().SysCallN(34, 1, m.Instance(), PascalBool(AValue))
}

func ComboBoxClass() TClass {
	ret := comboBoxImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TComboBox) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(12, m.Instance(), m.changePtr)
}

func (m *TComboBox) SetOnCloseUp(fn TNotifyEvent) {
	if m.closeUpPtr != 0 {
		RemoveEventElement(m.closeUpPtr)
	}
	m.closeUpPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(13, m.Instance(), m.closeUpPtr)
}

func (m *TComboBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(14, m.Instance(), m.contextPopupPtr)
}

func (m *TComboBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(15, m.Instance(), m.dblClickPtr)
}

func (m *TComboBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(16, m.Instance(), m.dragDropPtr)
}

func (m *TComboBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(17, m.Instance(), m.dragOverPtr)
}

func (m *TComboBox) SetOnDrawItem(fn TDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(18, m.Instance(), m.drawItemPtr)
}

func (m *TComboBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(21, m.Instance(), m.endDragPtr)
}

func (m *TComboBox) SetOnDropDown(fn TNotifyEvent) {
	if m.dropDownPtr != 0 {
		RemoveEventElement(m.dropDownPtr)
	}
	m.dropDownPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(19, m.Instance(), m.dropDownPtr)
}

func (m *TComboBox) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(20, m.Instance(), m.editingDonePtr)
}

func (m *TComboBox) SetOnGetItems(fn TNotifyEvent) {
	if m.getItemsPtr != 0 {
		RemoveEventElement(m.getItemsPtr)
	}
	m.getItemsPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(22, m.Instance(), m.getItemsPtr)
}

func (m *TComboBox) SetOnMeasureItem(fn TMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(23, m.Instance(), m.measureItemPtr)
}

func (m *TComboBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(24, m.Instance(), m.mouseDownPtr)
}

func (m *TComboBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(25, m.Instance(), m.mouseEnterPtr)
}

func (m *TComboBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(26, m.Instance(), m.mouseLeavePtr)
}

func (m *TComboBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(27, m.Instance(), m.mouseMovePtr)
}

func (m *TComboBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(28, m.Instance(), m.mouseUpPtr)
}

func (m *TComboBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(29, m.Instance(), m.mouseWheelPtr)
}

func (m *TComboBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(30, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TComboBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(31, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TComboBox) SetOnSelect(fn TNotifyEvent) {
	if m.selectPtr != 0 {
		RemoveEventElement(m.selectPtr)
	}
	m.selectPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(32, m.Instance(), m.selectPtr)
}

func (m *TComboBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	comboBoxImportAPI().SysCallN(33, m.Instance(), m.startDragPtr)
}

var (
	comboBoxImport       *imports.Imports = nil
	comboBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ComboBox_BorderStyle", 0),
		/*1*/ imports.NewTable("ComboBox_Class", 0),
		/*2*/ imports.NewTable("ComboBox_Create", 0),
		/*3*/ imports.NewTable("ComboBox_DragCursor", 0),
		/*4*/ imports.NewTable("ComboBox_DragKind", 0),
		/*5*/ imports.NewTable("ComboBox_DragMode", 0),
		/*6*/ imports.NewTable("ComboBox_ItemHeight", 0),
		/*7*/ imports.NewTable("ComboBox_ItemWidth", 0),
		/*8*/ imports.NewTable("ComboBox_MaxLength", 0),
		/*9*/ imports.NewTable("ComboBox_ParentColor", 0),
		/*10*/ imports.NewTable("ComboBox_ParentFont", 0),
		/*11*/ imports.NewTable("ComboBox_ParentShowHint", 0),
		/*12*/ imports.NewTable("ComboBox_SetOnChange", 0),
		/*13*/ imports.NewTable("ComboBox_SetOnCloseUp", 0),
		/*14*/ imports.NewTable("ComboBox_SetOnContextPopup", 0),
		/*15*/ imports.NewTable("ComboBox_SetOnDblClick", 0),
		/*16*/ imports.NewTable("ComboBox_SetOnDragDrop", 0),
		/*17*/ imports.NewTable("ComboBox_SetOnDragOver", 0),
		/*18*/ imports.NewTable("ComboBox_SetOnDrawItem", 0),
		/*19*/ imports.NewTable("ComboBox_SetOnDropDown", 0),
		/*20*/ imports.NewTable("ComboBox_SetOnEditingDone", 0),
		/*21*/ imports.NewTable("ComboBox_SetOnEndDrag", 0),
		/*22*/ imports.NewTable("ComboBox_SetOnGetItems", 0),
		/*23*/ imports.NewTable("ComboBox_SetOnMeasureItem", 0),
		/*24*/ imports.NewTable("ComboBox_SetOnMouseDown", 0),
		/*25*/ imports.NewTable("ComboBox_SetOnMouseEnter", 0),
		/*26*/ imports.NewTable("ComboBox_SetOnMouseLeave", 0),
		/*27*/ imports.NewTable("ComboBox_SetOnMouseMove", 0),
		/*28*/ imports.NewTable("ComboBox_SetOnMouseUp", 0),
		/*29*/ imports.NewTable("ComboBox_SetOnMouseWheel", 0),
		/*30*/ imports.NewTable("ComboBox_SetOnMouseWheelDown", 0),
		/*31*/ imports.NewTable("ComboBox_SetOnMouseWheelUp", 0),
		/*32*/ imports.NewTable("ComboBox_SetOnSelect", 0),
		/*33*/ imports.NewTable("ComboBox_SetOnStartDrag", 0),
		/*34*/ imports.NewTable("ComboBox_Sorted", 0),
	}
)

func comboBoxImportAPI() *imports.Imports {
	if comboBoxImport == nil {
		comboBoxImport = NewDefaultImports()
		comboBoxImport.SetImportTable(comboBoxImportTables)
		comboBoxImportTables = nil
	}
	return comboBoxImport
}
