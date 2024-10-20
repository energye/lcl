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

// ICheckComboBox Parent: ICustomCheckCombo
type ICheckComboBox interface {
	ICustomCheckCombo
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
	SetOnEndDrag(fn TEndDragEvent)                 // property event
	SetOnDropDown(fn TNotifyEvent)                 // property event
	SetOnEditingDone(fn TNotifyEvent)              // property event
	SetOnGetItems(fn TNotifyEvent)                 // property event
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

// TCheckComboBox Parent: TCustomCheckCombo
type TCheckComboBox struct {
	TCustomCheckCombo
	changePtr         uintptr
	closeUpPtr        uintptr
	contextPopupPtr   uintptr
	dblClickPtr       uintptr
	dragDropPtr       uintptr
	dragOverPtr       uintptr
	endDragPtr        uintptr
	dropDownPtr       uintptr
	editingDonePtr    uintptr
	getItemsPtr       uintptr
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

func NewCheckComboBox(AOwner IComponent) ICheckComboBox {
	r1 := checkComboBoxImportAPI().SysCallN(2, GetObjectUintptr(AOwner))
	return AsCheckComboBox(r1)
}

func (m *TCheckComboBox) BorderStyle() TBorderStyle {
	r1 := checkComboBoxImportAPI().SysCallN(0, 0, m.Instance(), 0)
	return TBorderStyle(r1)
}

func (m *TCheckComboBox) SetBorderStyle(AValue TBorderStyle) {
	checkComboBoxImportAPI().SysCallN(0, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) DragCursor() TCursor {
	r1 := checkComboBoxImportAPI().SysCallN(3, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TCheckComboBox) SetDragCursor(AValue TCursor) {
	checkComboBoxImportAPI().SysCallN(3, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) DragKind() TDragKind {
	r1 := checkComboBoxImportAPI().SysCallN(4, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TCheckComboBox) SetDragKind(AValue TDragKind) {
	checkComboBoxImportAPI().SysCallN(4, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) DragMode() TDragMode {
	r1 := checkComboBoxImportAPI().SysCallN(5, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TCheckComboBox) SetDragMode(AValue TDragMode) {
	checkComboBoxImportAPI().SysCallN(5, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) ItemHeight() int32 {
	r1 := checkComboBoxImportAPI().SysCallN(6, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCheckComboBox) SetItemHeight(AValue int32) {
	checkComboBoxImportAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) ItemWidth() int32 {
	r1 := checkComboBoxImportAPI().SysCallN(7, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCheckComboBox) SetItemWidth(AValue int32) {
	checkComboBoxImportAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) MaxLength() int32 {
	r1 := checkComboBoxImportAPI().SysCallN(8, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCheckComboBox) SetMaxLength(AValue int32) {
	checkComboBoxImportAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCheckComboBox) ParentColor() bool {
	r1 := checkComboBoxImportAPI().SysCallN(9, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckComboBox) SetParentColor(AValue bool) {
	checkComboBoxImportAPI().SysCallN(9, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckComboBox) ParentFont() bool {
	r1 := checkComboBoxImportAPI().SysCallN(10, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckComboBox) SetParentFont(AValue bool) {
	checkComboBoxImportAPI().SysCallN(10, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckComboBox) ParentShowHint() bool {
	r1 := checkComboBoxImportAPI().SysCallN(11, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckComboBox) SetParentShowHint(AValue bool) {
	checkComboBoxImportAPI().SysCallN(11, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCheckComboBox) Sorted() bool {
	r1 := checkComboBoxImportAPI().SysCallN(32, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCheckComboBox) SetSorted(AValue bool) {
	checkComboBoxImportAPI().SysCallN(32, 1, m.Instance(), PascalBool(AValue))
}

func CheckComboBoxClass() TClass {
	ret := checkComboBoxImportAPI().SysCallN(1)
	return TClass(ret)
}

func (m *TCheckComboBox) SetOnChange(fn TNotifyEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(12, m.Instance(), m.changePtr)
}

func (m *TCheckComboBox) SetOnCloseUp(fn TNotifyEvent) {
	if m.closeUpPtr != 0 {
		RemoveEventElement(m.closeUpPtr)
	}
	m.closeUpPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(13, m.Instance(), m.closeUpPtr)
}

func (m *TCheckComboBox) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(14, m.Instance(), m.contextPopupPtr)
}

func (m *TCheckComboBox) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(15, m.Instance(), m.dblClickPtr)
}

func (m *TCheckComboBox) SetOnDragDrop(fn TDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(16, m.Instance(), m.dragDropPtr)
}

func (m *TCheckComboBox) SetOnDragOver(fn TDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(17, m.Instance(), m.dragOverPtr)
}

func (m *TCheckComboBox) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(20, m.Instance(), m.endDragPtr)
}

func (m *TCheckComboBox) SetOnDropDown(fn TNotifyEvent) {
	if m.dropDownPtr != 0 {
		RemoveEventElement(m.dropDownPtr)
	}
	m.dropDownPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(18, m.Instance(), m.dropDownPtr)
}

func (m *TCheckComboBox) SetOnEditingDone(fn TNotifyEvent) {
	if m.editingDonePtr != 0 {
		RemoveEventElement(m.editingDonePtr)
	}
	m.editingDonePtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(19, m.Instance(), m.editingDonePtr)
}

func (m *TCheckComboBox) SetOnGetItems(fn TNotifyEvent) {
	if m.getItemsPtr != 0 {
		RemoveEventElement(m.getItemsPtr)
	}
	m.getItemsPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(21, m.Instance(), m.getItemsPtr)
}

func (m *TCheckComboBox) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(22, m.Instance(), m.mouseDownPtr)
}

func (m *TCheckComboBox) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(23, m.Instance(), m.mouseEnterPtr)
}

func (m *TCheckComboBox) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(24, m.Instance(), m.mouseLeavePtr)
}

func (m *TCheckComboBox) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(25, m.Instance(), m.mouseMovePtr)
}

func (m *TCheckComboBox) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(26, m.Instance(), m.mouseUpPtr)
}

func (m *TCheckComboBox) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(27, m.Instance(), m.mouseWheelPtr)
}

func (m *TCheckComboBox) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(28, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCheckComboBox) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(29, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCheckComboBox) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(31, m.Instance(), m.startDragPtr)
}

func (m *TCheckComboBox) SetOnSelect(fn TNotifyEvent) {
	if m.selectPtr != 0 {
		RemoveEventElement(m.selectPtr)
	}
	m.selectPtr = MakeEventDataPtr(fn)
	checkComboBoxImportAPI().SysCallN(30, m.Instance(), m.selectPtr)
}

var (
	checkComboBoxImport       *imports.Imports = nil
	checkComboBoxImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CheckComboBox_BorderStyle", 0),
		/*1*/ imports.NewTable("CheckComboBox_Class", 0),
		/*2*/ imports.NewTable("CheckComboBox_Create", 0),
		/*3*/ imports.NewTable("CheckComboBox_DragCursor", 0),
		/*4*/ imports.NewTable("CheckComboBox_DragKind", 0),
		/*5*/ imports.NewTable("CheckComboBox_DragMode", 0),
		/*6*/ imports.NewTable("CheckComboBox_ItemHeight", 0),
		/*7*/ imports.NewTable("CheckComboBox_ItemWidth", 0),
		/*8*/ imports.NewTable("CheckComboBox_MaxLength", 0),
		/*9*/ imports.NewTable("CheckComboBox_ParentColor", 0),
		/*10*/ imports.NewTable("CheckComboBox_ParentFont", 0),
		/*11*/ imports.NewTable("CheckComboBox_ParentShowHint", 0),
		/*12*/ imports.NewTable("CheckComboBox_SetOnChange", 0),
		/*13*/ imports.NewTable("CheckComboBox_SetOnCloseUp", 0),
		/*14*/ imports.NewTable("CheckComboBox_SetOnContextPopup", 0),
		/*15*/ imports.NewTable("CheckComboBox_SetOnDblClick", 0),
		/*16*/ imports.NewTable("CheckComboBox_SetOnDragDrop", 0),
		/*17*/ imports.NewTable("CheckComboBox_SetOnDragOver", 0),
		/*18*/ imports.NewTable("CheckComboBox_SetOnDropDown", 0),
		/*19*/ imports.NewTable("CheckComboBox_SetOnEditingDone", 0),
		/*20*/ imports.NewTable("CheckComboBox_SetOnEndDrag", 0),
		/*21*/ imports.NewTable("CheckComboBox_SetOnGetItems", 0),
		/*22*/ imports.NewTable("CheckComboBox_SetOnMouseDown", 0),
		/*23*/ imports.NewTable("CheckComboBox_SetOnMouseEnter", 0),
		/*24*/ imports.NewTable("CheckComboBox_SetOnMouseLeave", 0),
		/*25*/ imports.NewTable("CheckComboBox_SetOnMouseMove", 0),
		/*26*/ imports.NewTable("CheckComboBox_SetOnMouseUp", 0),
		/*27*/ imports.NewTable("CheckComboBox_SetOnMouseWheel", 0),
		/*28*/ imports.NewTable("CheckComboBox_SetOnMouseWheelDown", 0),
		/*29*/ imports.NewTable("CheckComboBox_SetOnMouseWheelUp", 0),
		/*30*/ imports.NewTable("CheckComboBox_SetOnSelect", 0),
		/*31*/ imports.NewTable("CheckComboBox_SetOnStartDrag", 0),
		/*32*/ imports.NewTable("CheckComboBox_Sorted", 0),
	}
)

func checkComboBoxImportAPI() *imports.Imports {
	if checkComboBoxImport == nil {
		checkComboBoxImport = NewDefaultImports()
		checkComboBoxImport.SetImportTable(checkComboBoxImportTables)
		checkComboBoxImportTables = nil
	}
	return checkComboBoxImport
}
